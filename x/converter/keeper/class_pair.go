package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/irisnet/erc721-bridge/x/converter/types"
)

// GetClassPairs - get all registered class pair
func (k Keeper) GetClassPairs(ctx sdk.Context) []types.ClassCollection {
	classPairs := []types.ClassPair{}
	k.IterateClassPair(ctx, func(classPair types.ClassPair) (stop bool) {
		classPairs = append(classPairs, classPair)
		return false
	})

	return nil
}

// IterateClassPair iterates over all the stored class pair
func (k Keeper) IterateClassPair(ctx sdk.Context, cb func(ClassPair types.ClassPair) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefixClassPair)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var classPair types.ClassPair
		k.cdc.MustUnmarshal(iterator.Value(), &classPair)

		if cb(classPair) {
			break
		}
	}
}

// GetClassPairID returns the pair id from either of the registered tokens.
// Hex address or ClassId can be used as token argument.
func (k Keeper) GetClassPairID(ctx sdk.Context, class string) []byte {
	if common.IsHexAddress(class) {
		addr := common.HexToAddress(class)
		return k.GetERC721Map(ctx, addr)
	}
	return k.GetDenomMap(ctx, class)
}

// GetClassPair gets a registered token pair from the identifier.
func (k Keeper) GetClassPair(ctx sdk.Context, id []byte) (types.ClassPair, bool) {
	if id == nil {
		return types.ClassPair{}, false
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixClassPair)
	var ClassPair types.ClassPair
	bz := store.Get(id)
	if len(bz) == 0 {
		return types.ClassPair{}, false
	}

	k.cdc.MustUnmarshal(bz, &ClassPair)
	return ClassPair, true
}

// SetClassPair stores a token pair
func (k Keeper) SetClassPair(ctx sdk.Context, ClassPair types.ClassPair) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixClassPair)
	key := ClassPair.GetID()
	bz := k.cdc.MustMarshal(&ClassPair)
	store.Set(key, bz)
}

// DeleteClassPair removes a token pair.
func (k Keeper) DeleteClassPair(ctx sdk.Context, ClassPair types.ClassPair) {
	id := ClassPair.GetID()
	k.deleteClassPair(ctx, id)
	k.deleteERC721Map(ctx, ClassPair.GetERC721Contract())
	k.deleteClassMap(ctx, ClassPair.ClassId)
}

// deleteClassPair deletes the token pair for the given id
func (k Keeper) deleteClassPair(ctx sdk.Context, id []byte) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixClassPair)
	store.Delete(id)
}

// GetERC721Map returns the token pair id for the given address
func (k Keeper) GetERC721Map(ctx sdk.Context, erc721 common.Address) []byte {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixClassPairByERC721)
	return store.Get(erc721.Bytes())
}

// GetDenomMap returns the token pair id for the given denomination
func (k Keeper) GetDenomMap(ctx sdk.Context, denom string) []byte {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixClassPairByClass)
	return store.Get([]byte(denom))
}

// SetERC721Map sets the token pair id for the given address
func (k Keeper) SetERC721Map(ctx sdk.Context, erc721 common.Address, id []byte) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixClassPairByERC721)
	store.Set(erc721.Bytes(), id)
}

// deleteERC721Map deletes the token pair id for the given address
func (k Keeper) deleteERC721Map(ctx sdk.Context, erc721 common.Address) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixClassPairByERC721)
	store.Delete(erc721.Bytes())
}

// SetClassMap sets the token pair id for the given class
func (k Keeper) SetClassMap(ctx sdk.Context, classId string, id []byte) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixClassPairByClass)
	store.Set([]byte(classId), id)
}

// deleteClassMap deletes the token pair id for the given class
func (k Keeper) deleteClassMap(ctx sdk.Context, classId string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixClassPairByClass)
	store.Delete([]byte(classId))
}

// IsClassPairRegistered - check if registered token ClassPair is registered
func (k Keeper) IsClassPairRegistered(ctx sdk.Context, id []byte) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixClassPair)
	return store.Has(id)
}

// IsERC721Registered check if registered ERC20 token is registered
func (k Keeper) IsERC721Registered(ctx sdk.Context, erc721 common.Address) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixClassPairByERC721)
	return store.Has(erc721.Bytes())
}

// IsClassRegistered check if registered cosmos x/nft Class
func (k Keeper) IsClassRegistered(ctx sdk.Context, classId string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixClassPairByClass)
	return store.Has([]byte(classId))
}
