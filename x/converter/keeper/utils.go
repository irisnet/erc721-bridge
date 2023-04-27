package keeper

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenerateERC721TokenID generates an ERC721 token ID
func (k Keeper) GenerateERC721TokenID(ctx sdk.Context, classId string, nftId string) *big.Int {
	panic("TODO: implement me")
}
