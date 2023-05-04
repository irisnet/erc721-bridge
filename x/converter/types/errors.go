package types

import errorsmod "cosmossdk.io/errors"

// errors
var (
	ErrInternalTokenPair       = errorsmod.Register(ModuleName, 1, "internal ethereum token mapping error")
	ErrTokenPairNotFound       = errorsmod.Register(ModuleName, 2, "token pair not found")
	ErrABIPack                 = errorsmod.Register(ModuleName, 3, "contract ABI pack failed")
	ErrABIUnpack               = errorsmod.Register(ModuleName, 4, "contract ABI unpack failed")
	ErrERC721TokenPairDisabled = errorsmod.Register(ModuleName, 5, "erc721 token pair is disabled")
	ErrUnauthorized            = errorsmod.Register(ModuleName, 6, "unauthorized address")
	ErrRegisterTokenPair       = errorsmod.Register(ModuleName, 7, "register token pair error")
	ErrUndefinedOwner          = errorsmod.Register(ModuleName, 8, "undefined owner of contract pair")
	ErrClassNotFound           = errorsmod.Register(ModuleName, 9, "class not found")
	ErrSaveClass               = errorsmod.Register(ModuleName, 10, "save class error")
	ErrNativeNftNotFound       = errorsmod.Register(ModuleName, 11, "native nft not found")
	ErrNativeNFTTransfer       = errorsmod.Register(ModuleName, 12, "native nft transfer error")
	ErrNativeNFTOwner          = errorsmod.Register(ModuleName, 13, "unauthorized owner")
	ErrNativeNFTMint           = errorsmod.Register(ModuleName, 14, "naive nft mint error")
	ErrNativeNFTBurn           = errorsmod.Register(ModuleName, 15, "naive nft burn error")
	ErrERC721TokenMint         = errorsmod.Register(ModuleName, 16, "mint erc721 nft error")
	ErrERC721TokenOwner        = errorsmod.Register(ModuleName, 17, "erc721 token owner check failed")
	ErrERC721TokenTransfer     = errorsmod.Register(ModuleName, 18, "erc721 token transfer error")
	ErrERC721TokenURI          = errorsmod.Register(ModuleName, 19, "erc721 token uri error")
	ErrERC721TokenData         = errorsmod.Register(ModuleName, 20, "erc721 token data error")
	ErrERC721Brun              = errorsmod.Register(ModuleName, 21, "erc721 token data error")
)
