package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/ethermint/x/evm/statedb"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
)

// AccountKeeper defines the expected interface needed to retrieve account info.
type AccountKeeper interface {
	GetModuleAddress(moduleName string) sdk.AccAddress
	GetSequence(sdk.Context, sdk.AccAddress) (uint64, error)
	GetAccount(sdk.Context, sdk.AccAddress) authtypes.AccountI
}

// EVMKeeper defines the expected EVM keeper interface used on erc20
type EVMKeeper interface {
	GetParams(ctx sdk.Context) evmtypes.Params
	GetAccountWithoutBalance(ctx sdk.Context, addr common.Address) *statedb.Account
	EstimateGas(c context.Context, req *evmtypes.EthCallRequest) (*evmtypes.EstimateGasResponse, error)
	ApplyMessage(ctx sdk.Context, msg core.Message, tracer vm.EVMLogger, commit bool) (*evmtypes.MsgEthereumTxResponse, error)
}

// NFTKeeper defines the expected interface needed to retrieve the IRSMod NFT denom.
type NFTKeeper interface {
	SaveClass(ctx sdk.Context, classID, classURI string, classData string) error
	Mint(ctx sdk.Context, classID, tokenID, tokenURI string, tokenData string, receiver sdk.AccAddress) error
	Transfer(ctx sdk.Context, classID string, tokenID string, tokenData string, receiver sdk.AccAddress) error
	Burn(ctx sdk.Context, classID string, tokenID string) error

	GetOwner(ctx sdk.Context, classID string, tokenID string) sdk.AccAddress
	HasClass(ctx sdk.Context, classID string) bool
	GetClass(ctx sdk.Context, classID string) (Class, bool)
}

// Class defines the interface specifications of collection that can be transferred across chains
type Class interface {
	GetID() string
	GetURI() string
	GetData() string
}

// NFT defines the interface specification of nft that can be transferred across chains
type NFT interface {
	GetClassID() string
	GetID() string
	GetURI() string
	GetData() string
}
