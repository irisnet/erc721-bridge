package keeper

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenerateERC721TokenID generates an ERC721 token ID
func (k Keeper) GenerateERC721TokenID(ctx sdk.Context, classId string, nftId string) *big.Int {
	// save mapping from classId+nftId to erc721TokenId
	sourceData := fmt.Sprintf("%s/%s", classId, nftId)
	res := sha256.Sum256([]byte(sourceData))

	tokenId := new(big.Int).SetBytes(res[:])

	return tokenId
}

// GenerateNativeTokenID generates an ERC721 token ID
func (k Keeper) GenerateNativeTokenID(ctx sdk.Context, erc721 common.Address, tokenId *big.Int) string {
	// save mapping from classId+nftId to erc721TokenId
	sourceData := fmt.Sprintf("%s/%s", erc721.String(), tokenId.String())
	newTokenId := sha256.Sum256([]byte(sourceData))
	return string(newTokenId[:])
}
