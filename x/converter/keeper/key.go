package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	keyContractClass = []byte{0x01}
	keyERC721NFT     = []byte{0x02}
)

func (ek erc721Keeper) contractClassStore(ctx sdk.Context) prefix.Store {
	store := ctx.KVStore(ek.k.storeKey)
	return prefix.NewStore(store, keyContractClass)
}

func (ek erc721Keeper) erc721NFTStore(ctx sdk.Context, prefixBz []byte) prefix.Store {
	store := ctx.KVStore(ek.k.storeKey)
	return prefix.NewStore(store, append(keyERC721NFT, prefixBz...))
}
