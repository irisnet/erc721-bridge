package keeper

import (
	"math/big"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/irisnet/erc721-bridge/x/converter/contracts"
	"github.com/irisnet/erc721-bridge/x/converter/types"
)

func (k Keeper) DeployERC721Contract(ctx sdk.Context, deployer common.Address,
	name, symbol, baseURI, classData string, owner common.Address) (common.Address, error) {
	contractArgs, err := contracts.ERC721PresetMinterPauserContract.ABI.Pack(
		"",
		name,
		symbol,
		baseURI,
		classData,
		owner,
	)
	if err != nil {
		return common.Address{}, errorsmod.Wrapf(types.ErrABIPack, "class metadata is invalid %s: %s", name, err.Error())
	}
	data := make([]byte, len(contracts.ERC721PresetMinterPauserContract.Bin)+len(contractArgs))
	copy(data[:len(contracts.ERC721PresetMinterPauserContract.Bin)], contracts.ERC721PresetMinterPauserContract.Bin)
	copy(data[len(contracts.ERC721PresetMinterPauserContract.Bin):], contractArgs)

	nonce, err := k.accountKeeper.GetSequence(ctx, deployer.Bytes())
	if err != nil {
		return common.Address{}, err
	}
	contractAddr := crypto.CreateAddress(deployer, nonce)
	_, err = k.CallEVMWithData(ctx, deployer, nil, data, true)
	if err != nil {
		return common.Address{}, errorsmod.Wrapf(err, "failed to deploy contract for %s", name)
	}

	return contractAddr, nil
}

// Mint mints an NFT
func (k Keeper) Mint(ctx sdk.Context,
	contractAddress common.Address,
	contractABI abi.ABI,
	to common.Address,
	tokenId *big.Int,
	tokenURI string,
	tokenData string,
) error {
	_, err := k.CallEVM(ctx,
		contractABI,
		types.ModuleAddress, contractAddress, true,
		types.ERC721MethodMintNFT, to, tokenId, tokenURI, tokenData)
	if err != nil {
		return err
	}
	return nil
}

// Burn burns an NFT
func (k Keeper) Burn(ctx sdk.Context,
	contractAddress common.Address,
	contractABI abi.ABI,
	sender common.Address,
	tokenId *big.Int,
) error {
	_, err := k.CallEVM(ctx,
		contractABI,
		sender, contractAddress, true,
		types.ERC721MethodBurnNFT, tokenId)
	if err != nil {
		return err
	}
	return nil
}

// TransferFrom transfers an NFT
func (k Keeper) TransferFrom(ctx sdk.Context,
	contractAddress common.Address,
	contractABI abi.ABI,
	from common.Address,
	to common.Address,
	tokenId *big.Int,
) error {
	_, err := k.CallEVM(ctx,
		contractABI,
		from, contractAddress, true,
		types.ERC721MethodTransferFrom, from, to, tokenId)
	if err != nil {
		return err
	}
	return nil
}

// SetClass sets a class
func (k Keeper) SetClass(ctx sdk.Context, contract common.Address, uri string, data string) error {
	erc721 := contracts.ERC721PresetMinterPauserContract.ABI
	_, err := k.CallEVM(ctx,
		erc721,
		types.ModuleAddress, contract, true,
		types.ERC721MethodSetClass, uri, data)
	if err != nil {
		return err
	}
	return nil
}

// QueryERC721 queries an ERC721 contract
func (k Keeper) QueryERC721(
	ctx sdk.Context,
	contractAddress common.Address,
	contractABI abi.ABI,
	isSystem bool,
) (types.ERC721Data, error) {

	var (
		nameRes   types.ERC721StringResponse
		symbolRes types.ERC721StringResponse
	)

	// Name
	res, err := k.CallEVM(ctx,
		contractABI, types.ModuleAddress, contractAddress, false, types.ERC721MethodName)
	if err != nil {
		return types.ERC721Data{}, err
	}

	if err := contractABI.UnpackIntoInterface(&nameRes, types.ERC721MethodName, res.Ret); err != nil {
		return types.ERC721Data{}, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to unpack name: %s", err.Error(),
		)
	}

	// Symbol
	res, err = k.CallEVM(ctx,
		contractABI, types.ModuleAddress, contractAddress, false, types.ERC721MethodSymbol)
	if err != nil {
		return types.ERC721Data{}, err
	}

	if err := contractABI.UnpackIntoInterface(&symbolRes, types.ERC721MethodSymbol, res.Ret); err != nil {
		return types.ERC721Data{}, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to unpack symbol: %s", err.Error(),
		)
	}
	erc721Data := types.NewERC721Data(nameRes.Value, symbolRes.Value)
	if isSystem {

		var (
			classURIRes  types.ERC721StringResponse
			classDataRes types.ERC721StringResponse
		)

		// URI
		res, err = k.CallEVM(ctx,
			contractABI, types.ModuleAddress, contractAddress, false, types.ERC721MethodClassURI)
		if err != nil {
			return types.ERC721Data{}, err
		}

		if err := contractABI.UnpackIntoInterface(&classURIRes, types.ERC721MethodClassURI, res.Ret); err != nil {
			return types.ERC721Data{}, errorsmod.Wrapf(
				types.ErrABIUnpack, "failed to unpack symbol: %s", err.Error(),
			)
		}

		// Class Data
		res, err = k.CallEVM(ctx,
			contractABI, types.ModuleAddress, contractAddress, false, types.ERC721MethodClassData)
		if err != nil {
			return types.ERC721Data{}, err
		}
		//
		if err := contractABI.UnpackIntoInterface(&classDataRes, types.ERC721MethodClassData, res.Ret); err != nil {
			return types.ERC721Data{}, errorsmod.Wrapf(
				types.ErrABIUnpack, "failed to unpack symbol: %s", err.Error(),
			)
		}
		erc721Data.URI = classURIRes.Value
		erc721Data.Data = classDataRes.Value
	}

	return erc721Data, nil
}

// QueryERC721Token queries an ERC721 token
func (k Keeper) QueryERC721Token(
	ctx sdk.Context,
	contractAddress common.Address,
	contractABI abi.ABI,
	tokenId *big.Int,
	isSystem bool,
) (types.ERC721TokenData, error) {

	var (
		tokenURIRs types.ERC721StringResponse
		tokenData  types.ERC721StringResponse
	)

	// URI
	res, err := k.CallEVM(ctx,
		contractABI, types.ModuleAddress, contractAddress, false, types.ERC721MethodTokenURI, tokenId)
	if err != nil {
		return types.ERC721TokenData{}, err
	}

	if err := contractABI.UnpackIntoInterface(&tokenURIRs, types.ERC721MethodTokenURI, res.Ret); err != nil {
		return types.ERC721TokenData{}, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to unpack tokenURI: %s", err.Error(),
		)
	}

	erc721TokenData := types.NewERC721TokenData(tokenURIRs.Value)

	if isSystem {
		// Data
		res, err = k.CallEVM(ctx,
			contractABI, types.ModuleAddress, contractAddress, false, types.ERC721MethodTokenData, tokenId)
		if err != nil {
			return types.ERC721TokenData{}, err
		}

		if err := contractABI.UnpackIntoInterface(&tokenData, types.ERC721MethodTokenData, res.Ret); err != nil {
			return types.ERC721TokenData{}, errorsmod.Wrapf(
				types.ErrABIUnpack, "failed to unpack tokenData: %s", err.Error(),
			)
		}
		erc721TokenData.Data = tokenData.Value
	}

	return erc721TokenData, nil
}

// ClassData queries an account's class data for a given ERC721 contract
func (k Keeper) ClassData(
	ctx sdk.Context,
	contractABI abi.ABI,
	contractAddress common.Address,
) (string, error) {
	res, err := k.CallEVM(ctx,
		contractABI, types.ModuleAddress, contractAddress, false, types.ERC721MethodClassData)
	if err != nil {
		return "", err
	}

	unpacked, err := contractABI.Unpack(types.ERC721MethodClassData, res.Ret)
	if err != nil || len(unpacked) == 0 {
		return "", err
	}

	classData, ok := unpacked[0].(string)
	if !ok {
		return "", err
	}

	return classData, nil
}

// SupportsInterface checks if the contract supports an interface
func (k Keeper) SupportsInterface(ctx sdk.Context,
	contractABI abi.ABI,
	contractAddress common.Address,
	interfaceId [4]byte,
) (bool, error) {
	res, err := k.CallEVM(ctx,
		contractABI,
		types.ModuleAddress, contractAddress, false,
		types.ERC165MethodSupportsInterface, interfaceId)
	if err != nil {
		return false, err
	}

	unpacked, err := contractABI.Unpack(types.ERC165MethodSupportsInterface, res.Ret)
	if err != nil || len(unpacked) == 0 {
		return false, err
	}

	success, ok := unpacked[0].(bool)
	if !ok {
		return false, err
	}

	return success, nil
}

// ClassURI queries an account's class URI for a given ERC721 contract
func (k Keeper) ClassURI(
	ctx sdk.Context,
	contractABI abi.ABI,
	contractAddress common.Address,
) (string, error) {
	res, err := k.CallEVM(ctx,
		contractABI, types.ModuleAddress, contractAddress, false, types.ERC721MethodClassURI)
	if err != nil {
		return "", err
	}

	unpacked, err := contractABI.Unpack(types.ERC721MethodClassURI, res.Ret)
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
	contractABI abi.ABI,
	contractAddress common.Address,
	tokenID *big.Int,
) (string, error) {
	res, err := k.CallEVM(ctx,
		contractABI, types.ModuleAddress, contractAddress, false, types.ERC721MethodTokenData, tokenID)
	if err != nil {
		return "", err
	}

	unpacked, err := contractABI.Unpack(types.ERC721MethodTokenData, res.Ret)
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
	contractABI abi.ABI,
	contractAddress common.Address,
	tokenID *big.Int,
) (common.Address, error) {
	res, err := k.CallEVM(ctx,
		contractABI, types.ModuleAddress, contractAddress, false, types.ERC721MethodOwnerOf, tokenID)
	if err != nil {
		return common.Address{}, err
	}

	unpacked, err := contractABI.Unpack(types.ERC721MethodOwnerOf, res.Ret)
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
	contractABI abi.ABI,
	contractAddress common.Address,
	tokenID *big.Int,
) (string, error) {
	res, err := k.CallEVM(ctx,
		contractABI, types.ModuleAddress, contractAddress, false, types.ERC721MethodTokenURI, tokenID)
	if err != nil {
		return "", err
	}

	unpacked, err := contractABI.Unpack(types.ERC721MethodTokenURI, res.Ret)
	if err != nil || len(unpacked) == 0 {
		return "", err
	}

	uri, ok := unpacked[0].(string)
	if !ok {
		return "", err
	}

	return uri, nil
}
