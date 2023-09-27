package types

import errorsmod "cosmossdk.io/errors"

// errors
var (
	ErrInternalClassPair    = errorsmod.Register(ModuleName, 2, "internal class pair error")
	ErrClassPairNotFound    = errorsmod.Register(ModuleName, 3, "class pair not found")
	ErrABIPack              = errorsmod.Register(ModuleName, 4, "contract ABI pack failed")
	ErrABIUnpack            = errorsmod.Register(ModuleName, 5, "contract ABI unpack failed")
	ErrUnauthorized         = errorsmod.Register(ModuleName, 6, "unauthorized address")
	ErrRegisterTokenPair    = errorsmod.Register(ModuleName, 7, "register token pair error")
	ErrUndefinedOwner       = errorsmod.Register(ModuleName, 8, "undefined owner of contract pair")
	ErrSaveNativeClass      = errorsmod.Register(ModuleName, 9, "save native class error")
	ErrNativeNftNotFound    = errorsmod.Register(ModuleName, 10, "native nft not found")
	ErrNativeNFTTransfer    = errorsmod.Register(ModuleName, 11, "native nft transfer error")
	ErrNativeNFTOwner       = errorsmod.Register(ModuleName, 12, "unauthorized owner")
	ErrNativeNFTMint        = errorsmod.Register(ModuleName, 13, "naive nft mint error")
	ErrNativeNFTBurn        = errorsmod.Register(ModuleName, 14, "naive nft burn error")
	ErrERC721TokenMint      = errorsmod.Register(ModuleName, 15, "mint erc721 nft error")
	ErrERC721TokenOwner     = errorsmod.Register(ModuleName, 16, "erc721 token owner check failed")
	ErrERC721TokenTransfer  = errorsmod.Register(ModuleName, 17, "erc721 token transfer error")
	ErrERC721TokenURI       = errorsmod.Register(ModuleName, 18, "erc721 token uri error")
	ErrERC721Brun           = errorsmod.Register(ModuleName, 19, "erc721 token data error")
	ErrNotFoundClassMapping = errorsmod.Register(ModuleName, 20, "not found class mapping")
	ErrNotFoundTokenMapping = errorsmod.Register(ModuleName, 21, "not found token mapping")
	ErrInvalidERC721TokenId = errorsmod.Register(ModuleName, 22, "invalid erc721 token id")
)
