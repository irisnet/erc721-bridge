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
) (types.ClassPair, error) {

	id := k.GetClassPairID(ctx, token)
	if len(id) == 0 {
		return types.ClassPair{}, errorsmod.Wrapf(
			types.ErrClassPairNotFound, "class '%s' not registered by id", token,
		)
	}
	pair, found := k.GetClassPair(ctx, id)
	if !found {
		return types.ClassPair{}, errorsmod.Wrapf(
			types.ErrClassPairNotFound, "class '%s' not registered", token,
		)
	}
	if !sender.Equals(receiver) {
		return types.ClassPair{}, errorsmod.Wrapf(
			types.ErrUnauthorized, "sender must be equal to receiver",
		)
	}

	return pair, nil
}
