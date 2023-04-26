package types

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/common"
)

const (
	// ModuleName is the name of the module
	ModuleName = "converter"

	// StoreKey is the string store representation
	StoreKey = ModuleName

	// RouterKey is the msg router key for the module
	RouterKey = ModuleName
)

// ModuleAddress is the native module address for the module
var ModuleAddress common.Address

func init() {
	ModuleAddress = common.BytesToAddress(authtypes.NewModuleAddress(ModuleName).Bytes())
}
