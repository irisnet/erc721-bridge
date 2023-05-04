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

// prefix bytes for the EVM persistent store
const (
	prefixTokenPair = iota + 1
	prefixTokenPairByERC721
	prefixTokenPairByClass
	prefixERC721TokenIDByNativeTokenID
	prefixNativeTokenIDByERC721TokenID
)

// KVStore key prefixes
var (
	KeyPrefixTokenPair                    = []byte{prefixTokenPair}
	KeyPrefixTokenPairByERC721            = []byte{prefixTokenPairByERC721}
	KeyPrefixTokenPairByClass             = []byte{prefixTokenPairByClass}
	KeyPrefixNativeTokenIDByERC721TokenID = []byte{prefixNativeTokenIDByERC721TokenID}
	KeyPrefixERC721TokenIDByNativeTokenID = []byte{prefixERC721TokenIDByNativeTokenID}
)

// ERC721 Method Names
const (
	ERC721MethodMintNFT           = "mint"
	ERC721MethodBurnNFT           = "burn"
	ERC721MethodOwnerOf           = "ownerOf"
	ERC721MethodName              = "name"
	ERC721MethodSymbol            = "symbol"
	ERC721MethodClassData         = "classData"
	ERC721MethodClassURI          = "baseURI"
	ERC721MethodTransferFrom      = "transferFrom"
	ERC721MethodTokenURI          = "tokenURI"
	ERC721MethodTokenData         = "tokenData"
	ERC721MethodSetClass          = "setClass"
	ERC165MethodSupportsInterface = "supportsInterface"
)

// Supported Interface ID
const (
	IERC721InterfaceId                   = "0x80ac58cd"
	IERC721MeatadataInterfaceId          = "0x5b5e139f"
	IERC721PresetMinterPauserInterfaceId = "0x9f1bf2d9"
)

func KeyTokenIdPair(classId, nftId string) []byte {
	return append(KeyPrefixTokenPair, []byte(classId+"/"+nftId)...)
}
