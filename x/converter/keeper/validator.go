package keeper

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/irisnet/erc721-bridge/x/converter/types"
)

// ConvertNFTValidator is the validator for ConvertNFT
func (k Keeper) ConvertNFTValidator(
	ctx sdk.Context,
	sender sdk.AccAddress,
	classId string,
	nftId string,
) (types.TokenPair, error) {

	id := k.GetTokenPairID(ctx, classId)
	if len(id) == 0 {
		return types.TokenPair{}, errorsmod.Wrapf(
			types.ErrTokenPairNotFound, "class '%s' not registered by id", classId,
		)
	}
	pair, found := k.GetTokenPair(ctx, id)
	if !found {
		return types.TokenPair{}, errorsmod.Wrapf(
			types.ErrTokenPairNotFound, "class '%s' not registered", classId,
		)
	}
	if !pair.Enabled {
		return types.TokenPair{}, errorsmod.Wrapf(
			types.ErrERC721TokenPairDisabled, "minting token '%s' is not enabled by governance", classId,
		)
	}

	return pair, nil
}
