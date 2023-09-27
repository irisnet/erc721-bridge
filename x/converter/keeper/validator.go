package keeper

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/irisnet/erc721-bridge/x/converter/types"
)

// ConvertValidator is the validator for Convert
func (k Keeper) ConvertValidator(
	ctx sdk.Context,
	sender, receiver sdk.AccAddress,
	token string,
) (types.TokenPair, error) {

	id := k.GetTokenPairID(ctx, token)
	if len(id) == 0 {
		return types.TokenPair{}, errorsmod.Wrapf(
			types.ErrTokenPairNotFound, "class '%s' not registered by id", token,
		)
	}
	pair, found := k.GetTokenPair(ctx, id)
	if !found {
		return types.TokenPair{}, errorsmod.Wrapf(
			types.ErrTokenPairNotFound, "class '%s' not registered", token,
		)
	}
	if !pair.Enabled {
		return types.TokenPair{}, errorsmod.Wrapf(
			types.ErrERC721TokenPairDisabled, "minting token '%s' is not enabled by governance", token,
		)
	}
	if !sender.Equals(receiver) {
		return types.TokenPair{}, errorsmod.Wrapf(
			types.ErrUnauthorized, "sender must be equal to receiver",
		)
	}

	return pair, nil
}
