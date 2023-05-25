package keeper

import (
	"math/big"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/irisnet/erc721-bridge/x/converter/types"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetNativeNftIdMap stores a token id pair
// classId/nftId = tokenId
func (k Keeper) SetNativeNftIdMap(ctx sdk.Context, classId, nftId string, tokenId *big.Int) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixNativeTokenIDByERC721TokenID)
	key := types.KeyTokenIdPair(classId, nftId)
	store.Set(key, tokenId.Bytes())
}

// SetERC721TokenIdMap stores a token id pair
// erc721/tokenId = sha256.sum(classId/nftId)
func (k Keeper) SetERC721TokenIdMap(ctx sdk.Context, erc721 common.Address, tokenId *big.Int, nftId string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixERC721TokenIDByNativeTokenID)
	key := types.KeyTokenIdPair(erc721.String(), tokenId.String())
	store.Set(key, []byte(nftId))
}

// DeleteNativeNftIdMap deletes a token id pair
// classId/nftId = tokenId
func (k Keeper) DeleteNativeNftIdMap(ctx sdk.Context, classId, nftId string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixNativeTokenIDByERC721TokenID)
	key := types.KeyTokenIdPair(classId, nftId)
	store.Delete(key)
}

// DeleteERC721TokenIdMap deletes a token id pair
// erc721/tokenId = sha256.sum(classId/nftId)
func (k Keeper) DeleteERC721TokenIdMap(ctx sdk.Context, erc721 common.Address, tokenId *big.Int) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixERC721TokenIDByNativeTokenID)
	key := types.KeyTokenIdPair(erc721.String(), tokenId.String())
	store.Delete(key)
}

// GetNativeNftIdMap
// erc721/tokenId = sha256.sum(classId/nftId)
func (k Keeper) GetNativeNftIdMap(ctx sdk.Context, erc721 common.Address, tokenId *big.Int) string {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixERC721TokenIDByNativeTokenID)
	key := types.KeyTokenIdPair(erc721.String(), tokenId.String())
	bz := store.Get(key)
	return string(bz)
}

// GetERC721TokenIdMap
// classId/nftId = tokenId
func (k Keeper) GetERC721TokenIdMap(ctx sdk.Context, classId, nftId string) *big.Int {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixNativeTokenIDByERC721TokenID)
	key := types.KeyTokenIdPair(classId, nftId)
	bz := store.Get(key)
	if bz == nil {
		return nil
	}
	return new(big.Int).SetBytes(bz)
}
