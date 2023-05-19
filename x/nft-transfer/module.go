package bridge

import (
	"github.com/cosmos/cosmos-sdk/types/module"

	porttypes "github.com/cosmos/ibc-go/v5/modules/core/05-port/types"

	nfttransfer "github.com/bianjieai/nft-transfer"
	"github.com/bianjieai/nft-transfer/types"

	"github.com/irisnet/erc721-bridge/x/nft-transfer/keeper"
	bridgetypes "github.com/irisnet/erc721-bridge/x/nft-transfer/types"
)

var (
	_ module.AppModule    = AppModule{}
	_ porttypes.IBCModule = IBCModule{}
)

// AppModule represents the AppModule for this module
type AppModule struct {
	nfttransfer.AppModule
	k                 keeper.Keeper
	queryTokenTraceFn bridgetypes.QueryTokenTrace
}

// NewAppModule creates a new nft-transfer module
func NewAppModule(
	app nfttransfer.AppModule,
	k keeper.Keeper,
) AppModule {
	return AppModule{
		AppModule: app,
		k:         k,
	}
}

// RegisterServices override the nft-transfer module AppModule.RegisterServices
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), am.k)
	types.RegisterQueryServer(cfg.QueryServer(), am.k.ISC721Keeper())
}
