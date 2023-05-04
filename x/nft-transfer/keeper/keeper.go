package keeper

import (
	nfttransferkeeper "github.com/bianjieai/nft-transfer/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/irisnet/erc721-bridge/x/nft-transfer/types"
)

type Keeper struct {
	erc721Port      string
	ics721Keeper    nfttransferkeeper.Keeper
	converterKeeper types.ConverterKeeper
}

func (k Keeper) IsERC721Port(port string) bool {
	return k.erc721Port == port
}

func (k Keeper) ClassToContract(ctx sdk.Context, classId string) (common.Address, bool) {
	return k.converterKeeper.ClassToContract(ctx, classId)
}

func (k Keeper) DeleteTokenMapping(ctx sdk.Context, classId string, nftId []string) error {
	return k.converterKeeper.DeleteTokenMapping(ctx, classId, nftId)
}
