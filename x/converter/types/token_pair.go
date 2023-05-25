package types

import (
	"github.com/ethereum/go-ethereum/common"
	etherminttypes "github.com/evmos/ethermint/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

func NewTokenPair(contractAddress common.Address, classId string, contractOwner Owner) TokenPair {
	return TokenPair{
		ContractAddress: contractAddress.String(),
		ClassId:         classId,
		ContractOwner:   contractOwner,
	}
}

// GetID returns the SHA256 hash of the ERC721 address and denomination
func (tp TokenPair) GetID() []byte {
	id := tp.ContractAddress + "|" + tp.ClassId
	return tmhash.Sum([]byte(id))
}

// GetERC721Contract casts the hex string address of the ERC721 to common.Address
func (tp TokenPair) GetERC721Contract() common.Address {
	return common.HexToAddress(tp.ContractAddress)
}

// IsNativeNFT returns true if the owner of the ERC721 contract is the
// erc721 module account
func (tp TokenPair) IsNativeNFT() bool {
	return tp.ContractOwner == OWNER_MODULE
}

// IsNativeERC721 returns true if the owner of the ERC721 contract not the
// erc721 module account
func (tp TokenPair) IsNativeERC721() bool {
	return tp.ContractOwner == OWNER_EXTERNAL
}

// Validate performs a stateless validation of a TokenPair
func (tp TokenPair) Validate() error {
	return etherminttypes.ValidateAddress(tp.ContractAddress)
}
