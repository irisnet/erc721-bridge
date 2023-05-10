package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ConverterKeeper interface {
	ContractToClass(ctx sdk.Context, contract common.Address) (string, bool)

	ClassToContract(ctx sdk.Context, classId string) (common.Address, bool)

	HasContract(ctx sdk.Context, contract common.Address) bool

	ERC721ToNFT(ctx sdk.Context, contract common.Address, erc721TokenId *big.Int) (string, bool)

	DeleteTokenMapping(ctx sdk.Context, classId string, nftId []string) error
}

type QueryTokenTrace = func(ctx client.Context, classId, tokenId string) (traceClassId string, traceTokenId string, err error)
