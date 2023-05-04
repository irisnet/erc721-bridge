package keeper

import (
	"errors"
	"math/big"
	"strings"

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
	_, found := ek.classToContract(ctx, ibcClassId)
	if found {
		return nil
	}
	contractAddr, err := ek.k.DeployERC721Contract(ctx, "", "", "", classURI, classData)
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
	contractAddr, found := ek.classToContract(ctx, ibcClassId)
	if !found {
		return errors.New("ibcClassId not found")
	}

	var (
		erc721TokenId *big.Int
		ok            bool
	)

	erc721TokenId, ok = new(big.Int).SetString(tokenID, 10)
	if !ok {
		//TODO Use the unique algorithm to calculate the new tokenId
	}

	err := ek.k.Mint(ctx, contractAddr, common.BytesToAddress(receiver.Bytes()), *erc721TokenId, tokenURI, tokenData)
	if err != nil {
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
		contractAddr, found := ek.classToContract(ctx, classID)
		if !found {
			return errors.New("ibcClassId not found")
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
		return errors.New("invalid tokenID")
	}

	// Note: nft-transfer will verify the owner of the token, so here you can directly use the owner to operate
	owner, err := ek.k.OwnerOf(ctx, contracts.ERC721PresetMinterPauserContract.ABI, contractAddr, erc721TokenId)
	if err != nil {
		return err
	}
	return ek.k.TransferFrom(ctx, contractAddr, owner, common.BytesToAddress(receiver), *erc721TokenId)
}

// 该方法只会在sink链上执行，当且仅当token回到原链
func (ek erc721Keeper) Burn(ctx sdk.Context, ibcClassId string, nftId string) error {
	contractAddr, found := ek.classToContract(ctx, ibcClassId)
	if !found {
		return errors.New("ibcClassId not found")
	}

	erc721TokenId, err := ek.nftToERC721(ctx, ibcClassId, nftId)
	if err != nil {
		return err
	}
	return ek.k.Burn(ctx, contractAddr, *erc721TokenId)
}

// GetOwner会在origin,sink链上执行:
// 1. 当nft被跨链转出的时候，可能会在origin,sink链上执行owner校验
// 1. 当nft被接收时，可能会在origin,sink链上执行
func (ek erc721Keeper) GetOwner(ctx sdk.Context, classID string, tokenID string) sdk.AccAddress {
	var (
		contractAddr  common.Address
		erc721TokenId *big.Int
		ok            bool
	)

	if strings.HasPrefix(classID, nfttransfertypes.ClassPrefix+"/") {
		contractAddr, found := ek.classToContract(ctx, classID)
		if !found {
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

// HasClass not need to be implemented
func (ek erc721Keeper) HasClass(ctx sdk.Context, classID string) bool {
	return false
}

func (ek erc721Keeper) GetClass(ctx sdk.Context, classID string) (nfttransfertypes.Class, bool) {
	if strings.HasPrefix(classID, nfttransfertypes.ClassPrefix+"/") {
		contractAddr, found := ek.classToContract(ctx, classID)
		if !found {
			return nil, false
		}

		classID = contractAddr.Hex()
	}
	contractAddr := common.HexToAddress(classID)

	//TODO
	data, _ := ek.k.ClassData(ctx, contracts.ERC721PresetMinterPauserContract.ABI, contractAddr)
	uri, _ := ek.k.ClassURI(ctx, contracts.ERC721PresetMinterPauserContract.ABI, contractAddr)
	return types.ERC721Contract{
		Contract: contractAddr,
		URI:      uri,
		Data:     data,
	}, true
}

func (ek erc721Keeper) GetNFT(ctx sdk.Context, classID string, tokenID string) (nfttransfertypes.NFT, bool) {
	var (
		contractAddr  common.Address
		erc721TokenId *big.Int
		ok            bool
	)

	if strings.HasPrefix(classID, nfttransfertypes.ClassPrefix+"/") {
		contractAddr, found := ek.classToContract(ctx, classID)
		if !found {
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

	//TODO
	data, _ := ek.k.TokenData(ctx, contracts.ERC721PresetMinterPauserContract.ABI, contractAddr, erc721TokenId)
	uri, _ := ek.k.TokenURI(ctx, contracts.ERC721PresetMinterPauserContract.ABI, contractAddr, erc721TokenId)
	return types.ERC721Token{
		Contract: contractAddr,
		ID:       erc721TokenId,
		URI:      uri,
		Data:     data,
	}, true
}

func (ek erc721Keeper) ContractToClass(ctx sdk.Context, contractAddr common.Address) (string, bool) {
	store := ek.contractClassStore(ctx)
	bz := store.Get(contractAddr.Bytes())
	if bz == nil || len(bz) == 0 {
		return "", false
	}
	return string(bz), false
}

func (ek erc721Keeper) ERC721ToNFT(ctx sdk.Context, contract common.Address, erc721TokenId *big.Int) (string, error) {
	erc721Store := ek.erc721NFTStore(ctx, contract.Bytes())
	bz := erc721Store.Get(erc721TokenId.Bytes())
	if bz == nil || len(bz) == 0 {
		return "", errors.New("not found")
	}
	return string(bz), nil
}

func (ek erc721Keeper) mapClassAndContract(ctx sdk.Context, ibcClassId string, contractAddr common.Address) {
	store := ek.contractClassStore(ctx)
	store.Set([]byte(ibcClassId), contractAddr.Bytes())
	store.Set(contractAddr.Bytes(), []byte(ibcClassId))
}

func (ek erc721Keeper) mapERC721AndNFT(ctx sdk.Context,
	ibcClassId,
	nftId string,
	contractAddr common.Address,
	erc721TokenId *big.Int,
) {
	classStore := ek.erc721NFTStore(ctx, []byte(ibcClassId))
	classStore.Set([]byte(nftId), erc721TokenId.Bytes())

	erc721Store := ek.erc721NFTStore(ctx, contractAddr.Bytes())
	erc721Store.Set(erc721TokenId.Bytes(), []byte(nftId))
}

func (ek erc721Keeper) classToContract(ctx sdk.Context, ibcClassId string) (common.Address, bool) {
	store := ek.contractClassStore(ctx)
	contractBz := store.Get([]byte(ibcClassId))
	if contractBz == nil || len(contractBz) == 0 {
		return common.Address{}, false
	}
	return common.Address(contractBz), false
}

func (ek erc721Keeper) nftToERC721(ctx sdk.Context, ibcClassId string, nftId string) (*big.Int, error) {
	classStore := ek.erc721NFTStore(ctx, []byte(ibcClassId))
	bz := classStore.Get([]byte(nftId))
	if bz == nil || len(bz) == 0 {
		return nil, errors.New("not found")
	}
	return new(big.Int).SetBytes(bz), nil
}
