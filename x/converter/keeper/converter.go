package keeper

import (
	"math/big"

	"github.com/irisnet/erc721-bridge/x/converter/contracts"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/irisnet/erc721-bridge/x/converter/types"
)

// ConvertNFTMint converts a native Cosmos token to an ERC721 token
func (k Keeper) ConvertNFTMint(
	ctx sdk.Context,
	pair types.TokenPair,
	msg *types.MsgConvertNFT,
	receiver common.Address,
	sender sdk.AccAddress,
) (*big.Int, error) {

	// Check native nft owner
	nativeNFTOwner := k.nftKeeper.GetOwner(ctx, msg.ClassId, msg.TokenId)
	if !nativeNFTOwner.Equals(sender) {
		return nil, errorsmod.Wrapf(
			types.ErrNativeNFTOwner, "native nft %s owner is not %s", msg.TokenId, sender)
	}

	nftInfo, found := k.nftKeeper.GetNft(ctx, msg.ClassId, msg.TokenId)
	if !found {
		return nil, errorsmod.Wrapf(
			types.ErrNativeNftNotFound, "native nft %s not found", msg.TokenId)
	}

	// Escrow native token on module account
	if err := k.nftKeeper.Transfer(ctx,
		msg.ClassId, msg.TokenId, nftInfo.GetData(), types.ModuleAddress.Bytes()); err != nil {
		return nil, errorsmod.Wrapf(
			types.ErrNativeNFTTransfer, "native nft %s transfer failed", msg.TokenId)
	}

	contract := pair.GetERC721Contract()
	erc721Abi := contracts.ERC721PresetMinterPauserContract.ABI

	// Get Token ID
	newTokenId := k.GenerateERC721TokenID(ctx, msg.ClassId, msg.TokenId)

	// Mint ERC721 Token To Receiver
	if err := k.Mint(ctx,
		contract, erc721Abi, receiver, newTokenId, nftInfo.GetURI(), nftInfo.GetData()); err != nil {
		return nil, errorsmod.Wrap(
			types.ErrERC721TokenMint, "erc721 token mint failed")
	}

	// Check expected receiver nft owner after mint
	owner, err := k.OwnerOf(ctx, erc721Abi, contract, newTokenId)
	if err != nil {
		return nil, errorsmod.Wrap(
			types.ErrERC721TokenOwner, "erc721 token owner check failed")
	}

	if owner != receiver {
		return nil, errorsmod.Wrap(
			types.ErrERC721TokenOwner, "erc721 token owner check failed")
	}

	return newTokenId, nil
}

// ConvertNFTBurn converts a erc721 token to an native Cosmos token
func (k Keeper) ConvertNFTBurn(ctx sdk.Context,
	pair types.TokenPair,
	msg *types.MsgConvertNFT,
	receiver common.Address,
	sender sdk.AccAddress) (*big.Int, error) {

	// Check native nft owner
	nativeNFTOwner := k.nftKeeper.GetOwner(ctx, msg.ClassId, msg.TokenId)
	if !nativeNFTOwner.Equals(sender) {
		return nil, errorsmod.Wrapf(
			types.ErrNativeNFTOwner, "native nft %s owner is not %s", msg.TokenId, sender)
	}

	// Burn Native NFT owner
	if err := k.nftKeeper.Burn(ctx, msg.ClassId, msg.TokenId); err != nil {
		return nil, errorsmod.Wrapf(
			types.ErrNativeNFTBurn, "native nft %s burn failed", msg.TokenId)
	}

	// Get Token ID
	newTokenId := k.GenerateERC721TokenID(ctx, msg.ClassId, msg.TokenId)

	contract := pair.GetERC721Contract()
	erc721Abi := contracts.ERC721PresetMinterPauserContract.ABI

	// transfer erc721 token to module account
	if err := k.TransferFrom(ctx,
		contract, erc721Abi, types.ModuleAddress, receiver, newTokenId); err != nil {
		return nil, errorsmod.Wrapf(types.ErrERC721TokenTransfer,
			"erc721 token %s transfer failed", newTokenId)
	}
	// Check expected receiver nft owner after mint
	owner, err := k.OwnerOf(ctx, erc721Abi, contract, newTokenId)
	if err != nil {
		return nil, errorsmod.Wrap(
			types.ErrERC721TokenOwner, "erc721 token owner check failed")
	}

	if owner != receiver {
		return nil, errorsmod.Wrap(
			types.ErrERC721TokenOwner, "erc721 token owner check failed")
	}

	return newTokenId, nil

}

// ConvertERC721Mint converts a erc721 token to an native Cosmos token
func (k Keeper) ConvertERC721Mint(ctx sdk.Context,
	pair types.TokenPair,
	msg *types.MsgConvertERC721,
	receiver sdk.AccAddress,
	sender common.Address,
) (string, error) {

	contract := pair.GetERC721Contract()
	erc721Abi := contracts.ERC721PresetMinterPauserContract.ABI

	// Check ERC721 nft owner
	erc721NFTOwner, err := k.OwnerOf(ctx, erc721Abi, contract, msg.TokenId.BigInt())
	if err != nil || erc721NFTOwner != sender {
		return "", errorsmod.Wrapf(
			types.ErrERC721TokenOwner, "erc721 nft %s owner check failed", msg.TokenId)
	}

	// transfer erc721 token to module account
	if err := k.TransferFrom(ctx,
		contract, erc721Abi, sender, types.ModuleAddress, msg.TokenId.BigInt()); err != nil {
		return "", errorsmod.Wrapf(types.ErrERC721TokenTransfer,
			"erc721 token %s transfer failed", msg.TokenId)
	}

	// Check expected receiver nft owner after mint
	tokenURI, err := k.TokenURI(ctx, erc721Abi, contract, msg.TokenId.BigInt())
	if err != nil {
		return "", errorsmod.Wrapf(types.ErrERC721TokenURI,
			"erc721 token %s tokenURI failed", msg.TokenId)
	}

	tokenData, err := k.TokenData(ctx, erc721Abi, contract, msg.TokenId.BigInt())
	if err != nil {
		return "", errorsmod.Wrapf(types.ErrERC721TokenData,
			"erc721 token %s tokenURI failed", msg.TokenId)
	}

	classId := pair.GetClassId()
	if err := k.nftKeeper.Mint(ctx,
		classId, msg.TokenId.String(), tokenURI, tokenData, receiver); err != nil {
		return "", errorsmod.Wrapf(types.ErrNativeNFTMint,
			"native nft %s mint failed", msg.TokenId,
		)
	}

	// Check expected receiver nft owner after mint
	newContractNFTOwner, err := k.OwnerOf(ctx, erc721Abi, contract, msg.TokenId.BigInt())
	if err != nil || receiver.Equals(sdk.AccAddress(newContractNFTOwner.Bytes())) {
		return "", errorsmod.Wrap(
			types.ErrERC721TokenOwner, "erc721 token owner check failed")
	}

	return msg.TokenId.String(), nil
}

// ConvertERC721Burn converts a native Cosmos token to an erc721 token
func (k Keeper) ConvertERC721Burn(ctx sdk.Context,
	pair types.TokenPair,
	msg *types.MsgConvertERC721,
	receiver sdk.AccAddress,
	sender common.Address,
) (string, error) {

	contract := pair.GetERC721Contract()
	erc721Abi := contracts.ERC721PresetMinterPauserContract.ABI

	// Check ERC721 nft owner
	erc721NFTOwner, err := k.OwnerOf(ctx, erc721Abi, contract, msg.TokenId.BigInt())
	if err != nil || erc721NFTOwner != sender {
		return "", errorsmod.Wrapf(
			types.ErrERC721TokenOwner, "erc721 nft %s owner check failed", msg.TokenId)
	}

	// transfer native nft to receiver from module account
	if err := k.nftKeeper.Transfer(ctx,
		pair.GetClassId(), msg.TokenId.String(), "", receiver); err != nil {
		return "", errorsmod.Wrapf(
			types.ErrNativeNFTTransfer, "native nft %s transfer failed", msg.TokenId)
	}

	// burn ERC721 token
	if err := k.Burn(ctx, contract, erc721Abi, sender, msg.TokenId.BigInt()); err != nil {
		return "", errorsmod.Wrapf(
			types.ErrERC721Brun, "erc721 token %s burn failed", msg.TokenId)
	}

	// todo: check nft is not exist

	return msg.TokenId.String(), nil
}
