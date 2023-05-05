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

// KVStore key prefixes
var (
	KeyPrefixTokenPair                    = []byte{0x01}
	KeyPrefixTokenPairByERC721            = []byte{0x02}
	KeyPrefixTokenPairByClass             = []byte{0x03}
	KeyPrefixNativeTokenIDByERC721TokenID = []byte{0x04}
	KeyPrefixERC721TokenIDByNativeTokenID = []byte{0x05}
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
	IERC721InterfaceId                   = "0x80ac58cd" // 1. IERC721: https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.8.1/contracts/token/ERC721/IERC721.sol
	IERC721MeatadataInterfaceId          = "0x5b5e139f" // 2. IERC721MetaData: https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.8.1/contracts/token/ERC721/extensions/IERC721Metadata.sol
	IERC721PresetMinterPauserInterfaceId = "0x9f1bf2d9" // 3. System ERC721 Contract: https://github.com/irisnet/erc721-bridge/blob/main/x/converter/contracts/IERC721Interface.sol
)

func KeyTokenIdPair(classId, nftId string) []byte {
	return append(KeyPrefixTokenPair, []byte(classId+"/"+nftId)...)
}
