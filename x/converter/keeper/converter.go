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

	nftInfo, found := k.nftKeeper.GetNFT(ctx, msg.ClassId, msg.TokenId)
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

	erc721Contract := pair.GetERC721Contract()
	erc721Abi := contracts.ERC721PresetMinterPauserContract.ABI

	// Get Token ID
	newERC721TokenId := GenerateERC721TokenID(msg.ClassId, msg.TokenId)

	// Mint ERC721 Token To Receiver
	if err := k.Mint(ctx,
		erc721Contract, erc721Abi, receiver, newERC721TokenId, nftInfo.GetURI(), nftInfo.GetData()); err != nil {
		return nil, errorsmod.Wrap(
			types.ErrERC721TokenMint, "erc721 token mint failed")
	}

	// Check expected receiver nft owner after mint
	erc721TokenOwner, err := k.OwnerOf(ctx, erc721Abi, erc721Contract, newERC721TokenId)
	if err != nil {
		return nil, errorsmod.Wrap(
			types.ErrERC721TokenOwner, "erc721 token owner check failed")
	}

	if erc721TokenOwner != receiver {
		return nil, errorsmod.Wrap(
			types.ErrERC721TokenOwner, "erc721 token owner check failed")
	}

	// Store the mapping between nftId and tokenId
	k.SetNativeNftIdMap(ctx, msg.ClassId, msg.TokenId, newERC721TokenId)
	k.SetERC721TokenIdMap(ctx, pair.GetERC721Contract(), newERC721TokenId, msg.TokenId)

	return newERC721TokenId, nil
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
	// from store
	// k.GetERC721TokenID(ctx, msg.ClassId, msg.TokenId)
	erc721TokenId := k.GetERC721TokenIdMap(ctx, msg.ClassId, msg.TokenId)

	erc721Contract := pair.GetERC721Contract()
	erc721Abi := contracts.ERC721PresetMinterPauserContract.ABI

	// transfer erc721 token to module account
	if err := k.TransferFrom(ctx,
		erc721Contract, erc721Abi, types.ModuleAddress, receiver, erc721TokenId); err != nil {
		return nil, errorsmod.Wrapf(types.ErrERC721TokenTransfer,
			"erc721 token %s transfer failed", erc721TokenId)
	}
	// Check expected receiver nft owner after mint
	owner, err := k.OwnerOf(ctx, erc721Abi, erc721Contract, erc721TokenId)
	if err != nil {
		return nil, errorsmod.Wrap(
			types.ErrERC721TokenOwner, "erc721 token owner check failed")
	}

	if owner != receiver {
		return nil, errorsmod.Wrap(
			types.ErrERC721TokenOwner, "erc721 token owner check failed")
	}

	// delete token id pair
	k.DeleteERC721TokenIdMap(ctx, erc721Contract, erc721TokenId)
	k.DeleteNativeNftIdMap(ctx, pair.GetClassId(), msg.TokenId)

	return erc721TokenId, nil

}

// ConvertERC721Mint converts a erc721 token to an native Cosmos token
func (k Keeper) ConvertERC721Mint(ctx sdk.Context,
	pair types.TokenPair,
	msg *types.MsgConvertERC721,
	receiver sdk.AccAddress,
	sender common.Address,
) (string, error) {

	erc721Contract := pair.GetERC721Contract()
	erc721Abi := contracts.ERC721PresetMinterPauserContract.ABI

	// Check ERC721 nft owner
	erc721TokenOwner, err := k.OwnerOf(ctx, erc721Abi, erc721Contract, msg.TokenId.BigInt())
	if err != nil || erc721TokenOwner != sender {
		return "", errorsmod.Wrapf(
			types.ErrERC721TokenOwner, "erc721 nft %s owner check failed", msg.TokenId)
	}

	// transfer erc721 token to module account
	if err := k.TransferFrom(ctx,
		erc721Contract, erc721Abi, sender, types.ModuleAddress, msg.TokenId.BigInt()); err != nil {
		return "", errorsmod.Wrapf(types.ErrERC721TokenTransfer,
			"erc721 token %s transfer failed", msg.TokenId)
	}

	// Check expected receiver nft owner after mint
	tokenMetadata, err := k.QueryERC721Token(ctx, erc721Contract, erc721Abi, msg.TokenId.BigInt(), false)
	if err != nil {
		return "", errorsmod.Wrapf(types.ErrERC721TokenURI,
			"erc721 token %s tokenURI failed", msg.TokenId)
	}
	// Generator native token id
	newNativeTokenId := GenerateNativeTokenID(pair.GetERC721Contract(), msg.TokenId.BigInt())

	classId := pair.GetClassId()
	if err := k.nftKeeper.Mint(ctx,
		classId, newNativeTokenId, tokenMetadata.URI, tokenMetadata.Data, receiver); err != nil {
		return "", errorsmod.Wrapf(types.ErrNativeNFTMint,
			"native nft %s mint failed", msg.TokenId,
		)
	}

	// Check expected receiver nft owner after mint
	newContractTokenOwner, err := k.OwnerOf(ctx, erc721Abi, erc721Contract, msg.TokenId.BigInt())
	if err != nil || receiver.Equals(sdk.AccAddress(newContractTokenOwner.Bytes())) {
		return "", errorsmod.Wrap(
			types.ErrERC721TokenOwner, "erc721 token owner check failed")
	}

	// Store the mapping between nftId and tokenId
	k.SetERC721TokenIdMap(ctx, pair.GetERC721Contract(), msg.TokenId.BigInt(), newNativeTokenId)
	k.SetNativeNftIdMap(ctx, pair.GetClassId(), newNativeTokenId, msg.TokenId.BigInt())

	return newNativeTokenId, nil
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
	erc721TokenOwner, err := k.OwnerOf(ctx, erc721Abi, contract, msg.TokenId.BigInt())
	if err != nil || erc721TokenOwner != sender {
		return "", errorsmod.Wrapf(
			types.ErrERC721TokenOwner, "erc721 nft %s owner check failed", msg.TokenId)
	}

	// Query NativeNFTID from the mapping
	nativeTokenId := k.GetNativeNftIdMap(ctx, contract, msg.TokenId.BigInt())

	// transfer native nft to receiver from module account
	if err := k.nftKeeper.Transfer(ctx,
		pair.GetClassId(), nativeTokenId, "", receiver); err != nil {
		return "", errorsmod.Wrapf(
			types.ErrNativeNFTTransfer, "native nft %s transfer failed", msg.TokenId)
	}

	// burn ERC721 token
	if err := k.Burn(ctx, contract, erc721Abi, sender, msg.TokenId.BigInt()); err != nil {
		return "", errorsmod.Wrapf(
			types.ErrERC721Brun, "erc721 token %s burn failed", msg.TokenId)
	}

	// delete token id pair
	k.DeleteERC721TokenIdMap(ctx, contract, msg.TokenId.BigInt())
	k.DeleteNativeNftIdMap(ctx, pair.GetClassId(), nativeTokenId)

	return msg.TokenId.String(), nil
}
