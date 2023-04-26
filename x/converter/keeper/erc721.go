package keeper

import (
	"math/big"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/irisnet/erc721-bridge/x/converter/contracts"
	"github.com/irisnet/erc721-bridge/x/converter/types"
)

func (k Keeper) DeployERC721Contract(ctx sdk.Context, class nfttypes.Class) (common.Address, error) {
	contractArgs, err := contracts.ERC721PresetMinterPauserContract.ABI.Pack(
		"",
		class.Name,
		class.Symbol,
		class.Uri,
	)
	if err != nil {
		return common.Address{}, errorsmod.Wrapf(types.ErrABIPack, "class metadata is invalid %s: %s", class.Name, err.Error())
	}
	data := make([]byte, len(contracts.ERC721PresetMinterPauserContract.Bin)+len(contractArgs))
	copy(data[:len(contracts.ERC721PresetMinterPauserContract.Bin)], contracts.ERC721PresetMinterPauserContract.Bin)
	copy(data[len(contracts.ERC721PresetMinterPauserContract.Bin):], contractArgs)

	nonce, err := k.accountKeeper.GetSequence(ctx, types.ModuleAddress.Bytes())
	if err != nil {
		return common.Address{}, err
	}
	contractAddr := crypto.CreateAddress(types.ModuleAddress, nonce)
	_, err = k.CallEVMWithData(ctx, types.ModuleAddress, nil, data, true)
	if err != nil {
		return common.Address{}, errorsmod.Wrapf(err, "failed to deploy contract for %s", class.Name)
	}

	return contractAddr, nil
}

// SetClass sets a class
func (k Keeper) SetClass(ctx sdk.Context, contract common.Address, class nfttypes.Class) error {
	erc721 := contracts.ERC721PresetMinterPauserContract.ABI
	_, err := k.CallEVM(ctx,
		erc721,
		types.ModuleAddress, contract, true,
		types.ERC721MethodSetClass, class.GetUri(), class.GetData())
	if err != nil {
		return err
	}
	return nil
}

// QueryERC721 queries an ERC721 contract
func (k Keeper) QueryERC721(
	ctx sdk.Context,
	contract common.Address,
) (types.ERC721Data, error) {

	var (
		nameRes   types.ERC721StringResponse
		symbolRes types.ERC721StringResponse
	)

	erc721 := contracts.ERC721PresetMinterPauserContract.ABI
	// Name
	res, err := k.CallEVM(ctx, erc721, types.ModuleAddress, contract, false, types.ERC721MethodName)
	if err != nil {
		return types.ERC721Data{}, err
	}

	if err := erc721.UnpackIntoInterface(&nameRes, "name", res.Ret); err != nil {
		return types.ERC721Data{}, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to unpack name: %s", err.Error(),
		)
	}

	// Symbol
	res, err = k.CallEVM(ctx, erc721, types.ModuleAddress, contract, false, types.ERC721MethodSymbol)
	if err != nil {
		return types.ERC721Data{}, err
	}

	if err := erc721.UnpackIntoInterface(&symbolRes, "symbol", res.Ret); err != nil {
		return types.ERC721Data{}, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to unpack symbol: %s", err.Error(),
		)
	}

	return types.NewERC721Data(nameRes.Value, symbolRes.Value), nil
}

// ClassData queries an account's class data for a given ERC721 contract
func (k Keeper) ClassData(
	ctx sdk.Context,
	abi abi.ABI,
	contract common.Address,
) (string, error) {
	res, err := k.CallEVM(ctx, abi, types.ModuleAddress, contract, false, types.ERC721MethodClassData)
	if err != nil {
		return "", err
	}

	unpacked, err := abi.Unpack("classData", res.Ret)
	if err != nil || len(unpacked) == 0 {
		return "", err
	}

	classData, ok := unpacked[0].(string)
	if !ok {
		return "", err
	}

	return classData, nil
}

// ClassURI queries an account's class URI for a given ERC721 contract
func (k Keeper) ClassURI(
	ctx sdk.Context,
	abi abi.ABI,
	contract common.Address,
) (string, error) {
	res, err := k.CallEVM(ctx, abi, types.ModuleAddress, contract, false, types.ERC721MethodClassURI)
	if err != nil {
		return "", err
	}

	unpacked, err := abi.Unpack("classURI", res.Ret)
	if err != nil || len(unpacked) == 0 {
		return "", err
	}

	classURI, ok := unpacked[0].(string)
	if !ok {
		return "", err
	}

	return classURI, nil
}

// TokenData queries an account's token data for a given ERC721 contract
func (k Keeper) TokenData(
	ctx sdk.Context,
	abi abi.ABI,
	contract common.Address,
	tokenID *big.Int,
) (string, error) {
	res, err := k.CallEVM(ctx, abi, types.ModuleAddress, contract, false, types.ERC721MethodTokenData, tokenID)
	if err != nil {
		return "", err
	}

	unpacked, err := abi.Unpack("tokenData", res.Ret)
	if err != nil || len(unpacked) == 0 {
		return "", err
	}

	tokenData, ok := unpacked[0].(string)
	if !ok {
		return "", err
	}

	return tokenData, nil
}

// OwnerOf queries an account's owner for a given ERC721 contract
func (k Keeper) OwnerOf(
	ctx sdk.Context,
	abi abi.ABI,
	contract common.Address,
	tokenID *big.Int,
) (common.Address, error) {
	res, err := k.CallEVM(ctx, abi, types.ModuleAddress, contract, false, types.ERC721MethodOwnerOf, tokenID)
	if err != nil {
		return common.Address{}, err
	}

	unpacked, err := abi.Unpack("ownerOf", res.Ret)
	if err != nil || len(unpacked) == 0 {
		return common.Address{}, err
	}

	owner, ok := unpacked[0].(common.Address)
	if !ok {
		return common.Address{}, err
	}

	return owner, nil
}

// TokenURI queries an account's token URI for a given ERC721 contract
func (k Keeper) TokenURI(
	ctx sdk.Context,
	abi abi.ABI,
	contract common.Address,
	tokenID *big.Int,
) (string, error) {
	res, err := k.CallEVM(ctx, abi, types.ModuleAddress, contract, false, types.ERC721MethodTokenURI, tokenID)
	if err != nil {
		return "", err
	}

	unpacked, err := abi.Unpack("tokenURI", res.Ret)
	if err != nil || len(unpacked) == 0 {
		return "", err
	}

	uri, ok := unpacked[0].(string)
	if !ok {
		return "", err
	}

	return uri, nil
}
