package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/irisnet/erc721-bridge/x/converter/types"
)

// ConvertNFT converts a native Cosmos token to an ERC721 token
func (k Keeper) ConvertNFT(goCtx context.Context, msg *types.MsgConvertNFT) (*types.MsgConvertNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Error checked during msg validation
	receiver := common.HexToAddress(msg.Receiver)
	sender := sdk.MustAccAddressFromBech32(msg.Sender)
	// Check if the token pair not exists
	if !k.IsClassRegistered(ctx, msg.ClassId) {
		// Register the token pair
		_, err := k.SaveRegisteredClass(ctx, msg.ClassId)
		if err != nil {
			return nil, types.ErrRegisterTokenPair
		}
	}
	pair, err := k.ConvertValidator(ctx, sender, receiver.Bytes(), msg.ClassId)
	if err != nil {
		return nil, err
	}
	erc721 := common.HexToAddress(pair.Erc721Address)
	acc := k.evmKeeper.GetAccountWithoutBalance(ctx, erc721)
	if acc == nil || !acc.IsContract() {
		k.DeleteTokenPair(ctx, pair)
		k.Logger(ctx).Debug(
			"deleting selfdestructed token pair from state",
			"contract", pair.Erc721Address,
		)
		// NOTE: return nil error to persist the changes from the deletion
		return nil, nil
	}

	// Check ownership and execute conversion
	switch {
	case pair.IsNativeNFT():
		// Owner is module account
		// Convert NFT to ERC721
		return k.convertNFTNativeNFT(ctx, pair, msg, receiver, sender)
	case pair.IsNativeERC721():
		// owner is user account
		// Convert ERC721 to NFT
		return k.convertNFTERC721(ctx, pair, msg, receiver, sender)
	default:
		return nil, types.ErrUndefinedOwner
	}

}

// ConvertERC721 converts an ERC721 token to an native Cosmos token
func (k Keeper) ConvertERC721(goCtx context.Context, msg *types.MsgConvertERC721) (*types.MsgConvertERC721Response, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Error checked during msg validation
	receiver := sdk.MustAccAddressFromBech32(msg.Receiver)
	sender := common.HexToAddress(msg.Sender)

	contractAddress := common.HexToAddress(msg.ContractAddress)

	if !k.IsERC721Registered(ctx, contractAddress) {
		// Register the token pair
		_, err := k.SaveRegisteredERC721(ctx, contractAddress)
		if err != nil {
			return nil, types.ErrRegisterTokenPair
		}
	}

	pair, err := k.ConvertValidator(ctx, sender.Bytes(), receiver, msg.ContractAddress)
	if err != nil {
		return nil, err
	}
	erc721 := common.HexToAddress(pair.Erc721Address)
	acc := k.evmKeeper.GetAccountWithoutBalance(ctx, erc721)
	if acc == nil || !acc.IsContract() {
		k.DeleteTokenPair(ctx, pair)
		k.Logger(ctx).Debug(
			"deleting selfdestructed token pair from state",
			"contract", pair.Erc721Address,
		)
		// NOTE: return nil error to persist the changes from the deletion
		return nil, nil
	}

	// Check ownership and execute conversion
	switch {
	case pair.IsNativeNFT():
		// Contract Owner is Module Account
		// Convert ERC721 Token to Cosmos Native NFT
		return k.convertERC721NativeERC721(ctx, pair, msg, receiver, sender)
	case pair.IsNativeERC721():
		// Contract Owner is User Account
		// Convert Cosmos Native NFT to ERC721 Token
		return k.convertERC721NativeNFT(ctx, pair, msg, receiver, sender)

	default:
		return nil, types.ErrUndefinedOwner
	}
}

// convertNFTNativeNFT converts a native Cosmos token to an ERC721 token
// 1. Lock Native NFT To Module Account
// 2. Mint ERC721 Token To Receiver
func (k Keeper) convertNFTNativeNFT(ctx sdk.Context, pair types.TokenPair, msg *types.MsgConvertNFT, receiver common.Address, sender sdk.AccAddress) (*types.MsgConvertNFTResponse, error) {
	tokenId, err := k.ConvertNFTMint(ctx, pair, msg, receiver, sender)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(
		sdk.Events{
			sdk.NewEvent(
				types.EventTypeConvertNFT,
				sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
				sdk.NewAttribute(types.AttributeKeyReceiver, msg.Receiver),
				sdk.NewAttribute(types.AttributeKeyClass, msg.ClassId),
				sdk.NewAttribute(types.AttributeKeyCosmosNFT, msg.TokenId),
				sdk.NewAttribute(types.AttributeKeyERC721, pair.Erc721Address),
				sdk.NewAttribute(types.AttributeKeyERC721Token, tokenId.String()),
			),
		},
	)

	return &types.MsgConvertNFTResponse{}, nil
}

// convertNFTERC721 converts a native Cosmos token to an ERC721 token
// 1. Unlock ERC721 Token From ERC721 Contract
// 2. Transfer ERC721 Token To Receiver
// 3. Burn Native NFT
func (k Keeper) convertNFTERC721(ctx sdk.Context, pair types.TokenPair, msg *types.MsgConvertNFT, receiver common.Address, sender sdk.AccAddress) (*types.MsgConvertNFTResponse, error) {
	tokenId, err := k.ConvertNFTBurn(ctx, pair, msg, receiver, sender)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(
		sdk.Events{
			sdk.NewEvent(
				types.EventTypeConvertNFT,
				sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
				sdk.NewAttribute(types.AttributeKeyReceiver, msg.Receiver),
				sdk.NewAttribute(types.AttributeKeyClass, msg.ClassId),
				sdk.NewAttribute(types.AttributeKeyCosmosNFT, msg.TokenId),
				sdk.NewAttribute(types.AttributeKeyERC721, pair.Erc721Address),
				sdk.NewAttribute(types.AttributeKeyERC721Token, tokenId.String()),
			),
		},
	)

	return &types.MsgConvertNFTResponse{}, nil
}

// convertERC721NativeNFT converts an ERC721 token to a native Cosmos token
// 1. Lock ERC721 Token To Module Account
// 2. Mint Native NFT To Receiver
func (k Keeper) convertERC721NativeNFT(ctx sdk.Context, pair types.TokenPair, msg *types.MsgConvertERC721, receiver sdk.AccAddress, sender common.Address) (*types.MsgConvertERC721Response, error) {
	tokenId, err := k.ConvertERC721Mint(ctx, pair, msg, receiver, sender)
	if err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvents(
		sdk.Events{
			sdk.NewEvent(
				types.EventTypeConvertERC721,
				sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
				sdk.NewAttribute(types.AttributeKeyReceiver, msg.Receiver),
				sdk.NewAttribute(types.AttributeKeyClass, pair.ClassId),
				sdk.NewAttribute(types.AttributeKeyCosmosNFT, msg.TokenId.String()),
				sdk.NewAttribute(types.AttributeKeyERC721, pair.Erc721Address),
				sdk.NewAttribute(types.AttributeKeyERC721Token, tokenId),
			),
		},
	)
	return &types.MsgConvertERC721Response{}, nil
}

// convertERC721NativeERC721
// 1. Unlock ERC721 Token From ERC721 Contract
// 2. Transfer ERC721 Token To Receiver
// 3. Burn Native NFT
func (k Keeper) convertERC721NativeERC721(ctx sdk.Context, pair types.TokenPair, msg *types.MsgConvertERC721, receiver sdk.AccAddress, sender common.Address) (*types.MsgConvertERC721Response, error) {
	tokenId, err := k.ConvertERC721Burn(ctx, pair, msg, receiver, sender)
	if err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvents(
		sdk.Events{
			sdk.NewEvent(
				types.EventTypeConvertERC721,
				sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
				sdk.NewAttribute(types.AttributeKeyReceiver, msg.Receiver),
				sdk.NewAttribute(types.AttributeKeyClass, pair.ClassId),
				sdk.NewAttribute(types.AttributeKeyCosmosNFT, msg.TokenId.String()),
				sdk.NewAttribute(types.AttributeKeyERC721, pair.Erc721Address),
				sdk.NewAttribute(types.AttributeKeyERC721Token, tokenId),
			),
		},
	)
	return &types.MsgConvertERC721Response{}, nil
}
