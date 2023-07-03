package keeper

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/irisnet/erc721-bridge/x/converter/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) {
	// ensure erc721 module account is set on genesis
	if acc := k.accountKeeper.GetModuleAccount(ctx, types.ModuleName); acc == nil {
		// NOTE: shouldn't occur
		panic("the erc721 module account has not been set")
	}
	for _, classCollection := range data.ClassCollections {
		classPair := classCollection.GetClassPair()
		id := classPair.GetID()
		k.SetClassPair(ctx, classPair)
		k.SetClassMap(ctx, classPair.ClassId, id)
		k.SetERC721Map(ctx, classPair.GetERC721Contract(), id)
		for _, tokenPair := range classCollection.Tokens {
			erc721TokenId, ok := new(big.Int).SetString(tokenPair.TokenId, 10)
			if !ok {
				panic("")
			}
			k.SetNativeNftIdMap(ctx,
				classPair.ClassId, tokenPair.TokenId, erc721TokenId)

			k.SetERC721TokenIdMap(ctx,
				common.HexToAddress(classPair.ContractAddress),
				erc721TokenId, tokenPair.TokenId)
		}
	}

	erc721Keeper := k.ERC721Keeper()
	for _, trace := range data.ClassTraces {
		erc721Keeper.SetClassTrace(ctx, trace)
	}
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return &types.GenesisState{
		ClassCollections: k.GetClassCollections(ctx),
		ClassTraces:      k.ERC721Keeper().GetClassTrace(ctx),
	}
}
