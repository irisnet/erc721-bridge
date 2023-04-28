package contracts

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"

	"github.com/irisnet/erc721-bridge/x/converter/types"

	"github.com/ethereum/go-ethereum/common"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
)

var (
	//go:embed compiled_contracts/ERC721PresetMinterPauser.json
	ERC721PresetMinterPauserJSON []byte //nolint: golint

	// ERC721PresetMinterPauserContract is the compiled erc721 contract
	ERC721PresetMinterPauserContract evmtypes.CompiledContract

	// ERC721PresetMinterPauserAddress is the erc721 module address
	ERC721PresetMinterPauserAddress common.Address
)

func init() {
	ERC721PresetMinterPauserAddress = types.ModuleAddress

	err := json.Unmarshal(ERC721PresetMinterPauserJSON, &ERC721PresetMinterPauserContract)
	if err != nil {
		panic(err)
	}
}
