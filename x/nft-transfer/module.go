package bridge

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/spf13/cobra"

	porttypes "github.com/cosmos/ibc-go/v5/modules/core/05-port/types"

	nfttransfer "github.com/bianjieai/nft-transfer"
	"github.com/bianjieai/nft-transfer/client/cli"
	"github.com/bianjieai/nft-transfer/types"

	"github.com/irisnet/erc721-bridge/x/nft-transfer/keeper"
)

var (
	_ module.AppModule    = AppModule{}
	_ porttypes.IBCModule = IBCModule{}
)

// AppModule represents the AppModule for this module
type AppModule struct {
	nfttransfer.AppModule
	k keeper.Keeper
}

// NewAppModule creates a new nft-transfer module
func NewAppModule(app nfttransfer.AppModule, k keeper.Keeper) AppModule {
	return AppModule{
		AppModule: app,
		k:         k,
	}
}

// GetQueryCmd override the nft-transfer module AppModuleBasic.GetQueryCmd
func (am AppModule) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

// RegisterServices override the nft-transfer module AppModule.RegisterServices
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), am.k)
	types.RegisterQueryServer(cfg.QueryServer(), am.k.ISC721Keeper())
}
