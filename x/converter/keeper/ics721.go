package keeper

import (
	"encoding/json"
	"errors"
	"math/big"
	"strings"

	errorsmod "cosmossdk.io/errors"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	nfttransfertypes "github.com/bianjieai/nft-transfer/types"

	"github.com/irisnet/erc721-bridge/x/converter/contracts"
	"github.com/irisnet/erc721-bridge/x/converter/types"
)

var _ nfttransfertypes.NFTKeeper = erc721Keeper{}

type erc721Keeper struct {
	k Keeper
}

// CreateOrUpdateClass deploys an erc721 contract.
// It will only be executed on the sink chain, and it will only be executed once
func (ek erc721Keeper) CreateOrUpdateClass(ctx sdk.Context, ibcClassId string, classURI string, classData string) error {
	if _, found := ek.ClassToContract(ctx, ibcClassId); found {
		return nil
	}

	var (
		erc721Data   types.ERC721Data
		name, symbol string
	)
	if err := json.Unmarshal([]byte(classData), &erc721Data); err == nil {
		name = erc721Data.Name
		symbol = erc721Data.Symbol
	}

	contractAddr, err := ek.k.DeployERC721Contract(ctx,
		types.ModuleAddress,
		name,
		symbol,
		classURI,
		classData,
		types.ModuleAddress,
	)
	if err != nil {
		return err
	}
	ek.mapClassAndContract(ctx, ibcClassId, contractAddr)
	return nil
}

// Mint mint a erc721 token. It will only be executed on the sink chain, there are the following scenarios:
// 1. when the token is far away from the original chain, the sink chain performs the mint operation
// 2. when the token returns to the sink chain (timeout, failure), the sink chain executes the mint operation
func (ek erc721Keeper) Mint(ctx sdk.Context, ibcClassId string, tokenID string, tokenURI string, tokenData string, receiver sdk.AccAddress) error {
	contractAddr, err := ek.classToContract(ctx, ibcClassId)
	if err != nil {
		return err
	}

	var (
		erc721TokenId *big.Int
		ok            bool
	)

	erc721TokenId, ok = new(big.Int).SetString(tokenID, 10)
	if !ok {
		erc721TokenId = GenerateERC721TokenID(ibcClassId, tokenID)
	}

	if err = ek.k.Mint(ctx,
		contractAddr,
		contracts.ERC721PresetMinterPauserContract.ABI,
		common.BytesToAddress(receiver.Bytes()),
		erc721TokenId,
		tokenURI,
		tokenData,
	); err != nil {
		return err
	}
	ek.mapERC721AndNFT(ctx, ibcClassId, tokenID, contractAddr, erc721TokenId)
	return nil
}

// Transfer will be executed on the origin and sink chains, and there are the following scenarios:
// 1. when the token is far away from the origin chain, the origin chain performs a locking operation (transfer out), (classId, tokenID)=(contractAddr, erc721TokenId)
// 2. when the token is far away from the origin chain, the origin chain executes the unlock operation (timeout, failure), (classId, tokenID)=(contractAddr, erc721TokenId)
// 3. when the token is far away from the sink chain, the sink chain performs a lock operation (transfer out), (classId, tokenID)=(ibcClassId, nftId)
// 4. when the token is far away from the sink chain, the sink chain performs the unlock operation (timeout, failure return), (classId, tokenID)=(ibcClassId, nftId)
// 5. when the token returns to the origin chain, the origin chain performs the unlock operation (transfer), (classId, tokenID)=(contractAddr, erc721TokenId)
// 6. when the token returns to the sink chain, the origin chain executes the unlock operation (transfer), (classId, tokenID)=(ibcClassId, nftId)
func (ek erc721Keeper) Transfer(ctx sdk.Context, classID string, tokenID string, tokenData string, receiver sdk.AccAddress) error {
	var (
		contractAddr  common.Address
		erc721TokenId *big.Int
		ok            bool
	)

	if strings.HasPrefix(classID, nfttransfertypes.ClassPrefix+"/") {
		contractAddr, err := ek.classToContract(ctx, classID)
		if err != nil {
			return err
		}

		classID = contractAddr.Hex()
		id, err := ek.nftToERC721(ctx, classID, tokenID)
		if err != nil {
			return err
		}
		tokenID = id.String()
	}

	contractAddr = common.HexToAddress(classID)
	erc721TokenId, ok = new(big.Int).SetString(tokenID, 10)
	if !ok {
		return errorsmod.Wrapf(types.ErrInvalidERC721TokenId, "token_id: %s", tokenID)
	}

	// Note: nft-transfer will verify the owner of the token, so here you can directly use the owner to operate
	owner, err := ek.k.OwnerOf(ctx, contracts.ERC721PresetMinterPauserContract.ABI, contractAddr, erc721TokenId)
	if err != nil {
		return err
	}
	return ek.k.TransferFrom(ctx, contractAddr, contracts.ERC721PresetMinterPauserContract.ABI, owner, common.BytesToAddress(receiver), erc721TokenId)
}

// Burn only be executed on the sink chain, if and only if the token returns to the original chain
func (ek erc721Keeper) Burn(ctx sdk.Context, ibcClassId string, nftId string) error {
	contractAddr, err := ek.classToContract(ctx, ibcClassId)
	if err != nil {
		return err
	}

	erc721TokenId, err := ek.nftToERC721(ctx, ibcClassId, nftId)
	if err != nil {
		return err
	}
	return ek.k.Burn(ctx, contractAddr, contracts.ERC721PresetMinterPauserContract.ABI, types.ModuleAddress, erc721TokenId)
}

// GetOwner will be executed on the origin and sink chains:
// 1. when nft is transferred across chains, owner verification may be performed on the origin and sink chains
// 2. when nft is received, it may be executed on the origin and sink chains
func (ek erc721Keeper) GetOwner(ctx sdk.Context, classID string, tokenID string) sdk.AccAddress {
	var (
		contractAddr  common.Address
		erc721TokenId *big.Int
		ok            bool
	)

	if strings.HasPrefix(classID, nfttransfertypes.ClassPrefix+"/") {
		contractAddr, err := ek.classToContract(ctx, classID)
		if err != nil {
			return sdk.AccAddress{}
		}

		classID = contractAddr.Hex()
		id, err := ek.nftToERC721(ctx, classID, tokenID)
		if err != nil {
			return sdk.AccAddress{}
		}
		tokenID = id.String()
	}

	contractAddr = common.HexToAddress(classID)
	erc721TokenId, ok = new(big.Int).SetString(tokenID, 10)
	if !ok {
		return sdk.AccAddress{}
	}
	owner, err := ek.k.OwnerOf(ctx, contracts.ERC721PresetMinterPauserContract.ABI, contractAddr, erc721TokenId)
	if err != nil {
		return sdk.AccAddress{}
	}
	return sdk.AccAddress(owner.Bytes())
}

// HasClass return whether the class or contract exists
func (ek erc721Keeper) HasClass(ctx sdk.Context, classID string) bool {
	if strings.HasPrefix(classID, nfttransfertypes.ClassPrefix+"/") {
		contractAddr, err := ek.classToContract(ctx, classID)
		if err != nil {
			return false
		}

		classID = contractAddr.Hex()
	}
	contractAddr := common.HexToAddress(classID)
	return ek.HasContract(ctx, contractAddr)
}

// GetClass return the basic information of a class or contract (class_uri, class_data or base_uri, name, symbol)
func (ek erc721Keeper) GetClass(ctx sdk.Context, classID string) (nfttransfertypes.Class, bool) {
	if strings.HasPrefix(classID, nfttransfertypes.ClassPrefix+"/") {
		contractAddr, err := ek.classToContract(ctx, classID)
		if err != nil {
			return nil, false
		}

		classID = contractAddr.Hex()
	}

	contractAddr := common.HexToAddress(classID)
	support := ek.supportSysInterface(ctx, contractAddr)
	data, err := ek.k.QueryERC721(ctx, contractAddr,
		contracts.ERC721PresetMinterPauserContract.ABI, support)
	if err != nil {
		return nil, false
	}

	var (
		classURI  = data.URI
		classData = data.Data
	)

	if !support {
		bz, err := json.Marshal(data)
		if err != nil {
			return nil, false
		}
		classData = string(bz)
	}
	return types.ERC721Contract{
		Contract: contractAddr,
		URI:      classURI,
		Data:     classData,
	}, true
}

// GetNFT return the basic information of a nft (token_uri,token_data)
func (ek erc721Keeper) GetNFT(ctx sdk.Context, classID string, tokenID string) (nfttransfertypes.NFT, bool) {
	var (
		contractAddr  common.Address
		erc721TokenId *big.Int
		ok            bool
	)

	if strings.HasPrefix(classID, nfttransfertypes.ClassPrefix+"/") {
		contractAddr, err := ek.classToContract(ctx, classID)
		if err != nil {
			return nil, false
		}

		classID = contractAddr.Hex()
		id, err := ek.nftToERC721(ctx, classID, tokenID)
		if err != nil {
			return nil, false
		}
		tokenID = id.String()
	}

	contractAddr = common.HexToAddress(classID)
	erc721TokenId, ok = new(big.Int).SetString(tokenID, 10)
	if !ok {
		return nil, false
	}

	support := ek.supportSysInterface(ctx, contractAddr)
	erc721TokenInfo, err := ek.k.QueryERC721Token(ctx, contractAddr,
		contracts.ERC721PresetMinterPauserContract.ABI, erc721TokenId, support)
	if err != nil {
		return nil, false
	}

	return types.ERC721Token{
		Contract: contractAddr,
		ID:       erc721TokenId,
		URI:      erc721TokenInfo.URI,
		Data:     erc721TokenInfo.Data,
	}, true
}

// ContractToClass returns the classId of the contract mapping
func (ek erc721Keeper) ContractToClass(ctx sdk.Context, contractAddr common.Address) (string, bool) {
	store := ek.classStore(ctx)
	bz := store.Get(contractAddr.Bytes())
	if bz == nil || len(bz) == 0 {
		return "", false
	}
	return string(bz), true
}

// ClassToContract returns the contract address of the classId mapping
func (ek erc721Keeper) ClassToContract(ctx sdk.Context, classId string) (common.Address, bool) {
	contractAddr, err := ek.classToContract(ctx, classId)
	return contractAddr, err == nil
}

// HasContract return whether the contract exists
func (ek erc721Keeper) HasContract(ctx sdk.Context, contract common.Address) bool {
	account := ek.k.evmKeeper.GetAccountWithoutBalance(ctx, contract)
	if account == nil {
		return false
	}
	return account.IsContract()
}

// ERC721ToNFT returns the nftId of the (contractAddr,erc721TokenId) mapping
func (ek erc721Keeper) ERC721ToNFT(ctx sdk.Context, contract common.Address, erc721TokenId *big.Int) (string, error) {
	erc721Store := ek.tokenStore(ctx, contract.Bytes())
	bz := erc721Store.Get(erc721TokenId.Bytes())
	if bz == nil || len(bz) == 0 {
		return "", errors.New("not found")
	}
	return string(bz), nil
}

// ERC721ToNFT delete the (ibcClassId,[]nftId) mapping
func (ek erc721Keeper) DeleteTokenMapping(ctx sdk.Context, ibcClassId string, nftId []string) error {
	classStore := ek.tokenStore(ctx, []byte(ibcClassId))
	for _, nftId := range nftId {
		classStore.Delete([]byte(nftId))
	}
	return nil
}

func (ek erc721Keeper) supportSysInterface(ctx sdk.Context, contract common.Address) bool {
	interfaceId := common.FromHex(types.IERC721PresetMinterPauserInterfaceId)
	support, err := ek.k.SupportsInterface(ctx,
		contracts.ERC721PresetMinterPauserContract.ABI,
		contract,
		[4]byte(interfaceId),
	)
	if err != nil {
		return false
	}
	return support
}

func (ek erc721Keeper) mapClassAndContract(ctx sdk.Context, ibcClassId string, contractAddr common.Address) {
	store := ek.classStore(ctx)
	store.Set([]byte(ibcClassId), contractAddr.Bytes())
	store.Set(contractAddr.Bytes(), []byte(ibcClassId))
}

func (ek erc721Keeper) mapERC721AndNFT(ctx sdk.Context,
	ibcClassId,
	nftId string,
	contractAddr common.Address,
	erc721TokenId *big.Int,
) {
	classStore := ek.tokenStore(ctx, []byte(ibcClassId))
	classStore.Set([]byte(nftId), erc721TokenId.Bytes())

	erc721Store := ek.tokenStore(ctx, contractAddr.Bytes())
	erc721Store.Set(erc721TokenId.Bytes(), []byte(nftId))
}

func (ek erc721Keeper) classToContract(ctx sdk.Context, ibcClassId string) (common.Address, error) {
	store := ek.classStore(ctx)
	contractBz := store.Get([]byte(ibcClassId))
	if contractBz == nil || len(contractBz) == 0 {
		return common.Address{}, errorsmod.Wrapf(types.ErrNotFoundClassMapping, "class_id: %s", ibcClassId)
	}
	return common.Address(contractBz), nil
}

func (ek erc721Keeper) nftToERC721(ctx sdk.Context, ibcClassId string, nftId string) (*big.Int, error) {
	classStore := ek.tokenStore(ctx, []byte(ibcClassId))
	bz := classStore.Get([]byte(nftId))
	if bz == nil || len(bz) == 0 {
		return nil, errorsmod.Wrapf(types.ErrNotFoundTokenMapping, "class_id: %s,token_id: %s", ibcClassId, nftId)
	}
	return new(big.Int).SetBytes(bz), nil
}

func (ek erc721Keeper) classStore(ctx sdk.Context) prefix.Store {
	store := ctx.KVStore(ek.k.storeKey)
	return prefix.NewStore(store, types.KeyPrefixContractClass)
}

func (ek erc721Keeper) tokenStore(ctx sdk.Context, idBz []byte) prefix.Store {
	store := ctx.KVStore(ek.k.storeKey)
	return prefix.NewStore(store, append(types.KeyPrefixERC721NFT, idBz...))
}
