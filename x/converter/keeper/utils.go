package keeper

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// GenerateERC721TokenID generates an ERC721 token ID
func GenerateERC721TokenID(classId string, nftId string) *big.Int {
	// save mapping from classId+nftId to erc721TokenId
	sourceData := fmt.Sprintf("%s/%s", classId, nftId)
	res := sha256.Sum256([]byte(sourceData))

	tokenId := new(big.Int).SetBytes(res[:])

	return tokenId
}

// GenerateNativeTokenID generates an ERC721 token ID
func GenerateNativeTokenID(erc721 common.Address, tokenId *big.Int) string {
	// save mapping from classId+nftId to erc721TokenId
	sourceData := fmt.Sprintf("%s/%s", erc721.String(), tokenId.String())
	hashData := sha256.Sum256([]byte(sourceData))
	newTokenId := hex.EncodeToString(hashData[:])
	return newTokenId
}
