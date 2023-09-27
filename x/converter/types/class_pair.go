package types

import (
	"github.com/ethereum/go-ethereum/common"
	etherminttypes "github.com/evmos/ethermint/types"
	"github.com/cometbft/cometbft/crypto/tmhash"
)

func NewClassPair(contractAddress common.Address, classId string, contractOwner Owner) ClassPair {
	return ClassPair{
		ContractAddress: contractAddress.String(),
		ClassId:         classId,
		ContractOwner:   contractOwner,
	}
}

// GetID returns the SHA256 hash of the ERC721 address and denomination
func (tp ClassPair) GetID() []byte {
	id := tp.ContractAddress + "|" + tp.ClassId
	return tmhash.Sum([]byte(id))
}

// GetERC721Contract casts the hex string address of the ERC721 to common.Address
func (tp ClassPair) GetERC721Contract() common.Address {
	return common.HexToAddress(tp.ContractAddress)
}

// IsNativeNFT returns true if the owner of the ERC721 contract is the
// erc721 module account
func (tp ClassPair) IsNativeNFT() bool {
	return tp.ContractOwner == OWNER_MODULE
}

// IsNativeERC721 returns true if the owner of the ERC721 contract not the
// erc721 module account
func (tp ClassPair) IsNativeERC721() bool {
	return tp.ContractOwner == OWNER_EXTERNAL
}

// Validate performs a stateless validation of a TokenPair
func (tp ClassPair) Validate() error {
	return etherminttypes.ValidateAddress(tp.ContractAddress)
}
