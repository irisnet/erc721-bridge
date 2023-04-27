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

// ERC721 Method Names
const (
	ERC721MethodMintNFT           = "mint"
	ERC721MethodBurnNFT           = "burn"
	ERC721MethodOwnerOf           = "ownerOf"
	ERC721MethodName              = "name"
	ERC721MethodSymbol            = "symbol"
	ERC721MethodClassData         = "classData"
	ERC721MethodClassURI          = "classURI"
	ERC721MethodTransferFrom      = "transferFrom"
	ERC721MethodTokenURI          = "tokenURI"
	ERC721MethodTokenData         = "tokenData"
	ERC721MethodSetClass          = "setClass"
	ERC165MethodSupportsInterface = "supportsInterface"
)

// Supported Interface Id
const (
	IERC721BaseInterfaceId               = "0x93254542"
	IERC721CommonInterfaceId             = "0xff2625a9"
	IERC721PresetMinterPauserInterfaceId = "0x0c3eb79b"
)
