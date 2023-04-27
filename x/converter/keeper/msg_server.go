package keeper

import (
	"context"

	"github.com/irisnet/erc721-bridge/x/converter/types"
)

// ConvertNFT converts a native Cosmos token to an ERC721 token
func (k Keeper) ConvertNFT(goCtx context.Context, msg *types.MsgConvertNFT) (*types.MsgConvertNFTResponse, error) {
	// todo: implement
	return &types.MsgConvertNFTResponse{}, nil
}

// ConvertERC721 converts an ERC721 token to an native Cosmos token
func (k Keeper) ConvertERC721(goCtx context.Context, msg *types.MsgConvertERC721) (*types.MsgConvertERC721Response, error) {
	// todo: implement
	return &types.MsgConvertERC721Response{}, nil
}
