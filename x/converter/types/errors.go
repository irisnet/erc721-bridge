package types

import errorsmod "cosmossdk.io/errors"

// errors
var (
	ErrInternalTokenPair       = errorsmod.Register(ModuleName, 2, "internal ethereum token mapping error")
	ErrTokenPairNotFound       = errorsmod.Register(ModuleName, 3, "token pair not found")
	ErrABIPack                 = errorsmod.Register(ModuleName, 4, "contract ABI pack failed")
	ErrABIUnpack               = errorsmod.Register(ModuleName, 5, "contract ABI unpack failed")
	ErrERC721TokenPairDisabled = errorsmod.Register(ModuleName, 6, "erc721 token pair is disabled")
	ErrUnauthorized            = errorsmod.Register(ModuleName, 7, "unauthorized address")
	ErrRegisterTokenPair       = errorsmod.Register(ModuleName, 8, "register token pair error")
	ErrUndefinedOwner          = errorsmod.Register(ModuleName, 10, "undefined owner of contract pair")
	ErrClassNotFound           = errorsmod.Register(ModuleName, 11, "class not found")
	ErrSaveClass               = errorsmod.Register(ModuleName, 12, "save class error")
	ErrNativeNftNotFound       = errorsmod.Register(ModuleName, 13, "native nft not found")
	ErrNativeNFTTransfer       = errorsmod.Register(ModuleName, 14, "native nft transfer error")
	ErrNativeNFTOwner          = errorsmod.Register(ModuleName, 15, "unauthorized owner")
	ErrNativeNFTMint           = errorsmod.Register(ModuleName, 16, "naive nft mint error")
	ErrNativeNFTBurn           = errorsmod.Register(ModuleName, 17, "naive nft burn error")
	ErrERC721TokenMint         = errorsmod.Register(ModuleName, 18, "mint erc721 nft error")
	ErrERC721TokenOwner        = errorsmod.Register(ModuleName, 19, "erc721 token owner check failed")
	ErrERC721TokenTransfer     = errorsmod.Register(ModuleName, 20, "erc721 token transfer error")
	ErrERC721TokenURI          = errorsmod.Register(ModuleName, 21, "erc721 token uri error")
	ErrERC721TokenData         = errorsmod.Register(ModuleName, 22, "erc721 token data error")
	ErrERC721Brun              = errorsmod.Register(ModuleName, 23, "erc721 token data error")
)
