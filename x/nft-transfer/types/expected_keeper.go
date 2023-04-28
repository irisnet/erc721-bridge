package types

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

type ConverterKeeper interface {
	ContractToClass(ctx sdk.Context, contract common.Address) (string, bool)
	HasContract(ctx sdk.Context, contract common.Address) bool
	ClassToContract(ctx sdk.Context, classId string) (common.Address, bool)

	ERC721ToNFT(ctx sdk.Context,
		contract common.Address,
		erc721TokenId *big.Int,
	) (string, bool)

	NFTToERC721(ctx sdk.Context,
		classId string,
		nftId string,
	) (*big.Int, bool)

	MapClassAndContract(ctx sdk.Context,
		classId string,
		contract common.Address,
	) error

	MapERC721AndNFT(ctx sdk.Context,
		classId string,
		nftId string,
		contract common.Address,
		erc721TokenId *big.Int,
	) error
}
