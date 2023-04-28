package types

import (
	errorsmod "cosmossdk.io/errors"
	nfttransfer "github.com/bianjieai/nft-transfer/types"
)

var (
	ErrInvalidErc721TokenId  = errorsmod.Register(nfttransfer.ModuleName, 20, "invalid erc721 tokenId")
	ErrNotExistErc721TokenId = errorsmod.Register(nfttransfer.ModuleName, 20, "erc721 tokenId not exist")
)
