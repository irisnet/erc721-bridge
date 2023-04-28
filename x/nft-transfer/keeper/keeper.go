package keeper

import (
	nfttransferkeeper "github.com/bianjieai/nft-transfer/keeper"

	"github.com/irisnet/erc721-bridge/x/nft-transfer/types"
)

type Keeper struct {
	erc721Port      string
	ics721Keeper    nfttransferkeeper.Keeper
	converterKeeper types.ConverterKeeper
}
