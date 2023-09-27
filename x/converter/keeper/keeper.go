package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/irisnet/erc721-bridge/x/converter/types"

	"github.com/cometbft/cometbft/libs/log"
)

// Keeper of this module maintains collections of erc721.
type Keeper struct {
	storeKey storetypes.StoreKey
	cdc      codec.BinaryCodec

	accountKeeper types.AccountKeeper
	evmKeeper     types.EVMKeeper
	nftKeeper     types.NFTKeeper
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	ak types.AccountKeeper,
	evmKeeper types.EVMKeeper,
	nftKeeper types.NFTKeeper,
) Keeper {

	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,

		accountKeeper: ak,
		evmKeeper:     evmKeeper,
		nftKeeper:     nftKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("erc721-bridge/%s", types.ModuleName))
}
