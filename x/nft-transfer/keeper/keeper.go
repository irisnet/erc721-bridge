package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"

	nfttransferkeeper "github.com/bianjieai/nft-transfer/keeper"

	"github.com/irisnet/erc721-bridge/x/nft-transfer/types"
)

type Keeper struct {
	ics721Keeper    nfttransferkeeper.Keeper
	converterKeeper types.ConverterKeeper
}

func NewKeeper(ics721Keeper nfttransferkeeper.Keeper, converterKeeper types.ConverterKeeper) Keeper {
	return Keeper{ics721Keeper, converterKeeper}
}

func (k Keeper) ClassToContract(ctx sdk.Context, classId string) (common.Address, bool) {
	return k.converterKeeper.ClassToContract(ctx, classId)
}

func (k Keeper) DeleteTokenMapping(ctx sdk.Context, classId string, nftId []string) error {
	return k.converterKeeper.DeleteTokenMapping(ctx, classId, nftId)
}
