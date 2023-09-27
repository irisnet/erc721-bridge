package bridge

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"

	nfttransfer "github.com/bianjieai/nft-transfer"
	"github.com/bianjieai/nft-transfer/types"

	"github.com/irisnet/erc721-bridge/x/nft-transfer/keeper"
)

var (
	_ porttypes.IBCModule = IBCModule{}
)

// IBCModule implements the ICS26 interface for transfer given the transfer keeper.
type IBCModule struct {
	nfttransfer.IBCModule
	k keeper.Keeper
}

// NewIBCModule creates a new IBCModule given the keeper
func NewIBCModule(app nfttransfer.IBCModule, k keeper.Keeper) IBCModule {
	return IBCModule{app, k}
}

// OnAcknowledgementPacket implements the IBCModule interface
func (im IBCModule) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	if err := im.IBCModule.OnAcknowledgementPacket(ctx,
		packet, acknowledgement, relayer); err != nil {
		return err
	}

	var ack channeltypes.Acknowledgement
	if err := types.ModuleCdc.UnmarshalJSON(acknowledgement, &ack); err != nil {
		return im.IBCModule.OnAcknowledgementPacket(ctx,
			packet, acknowledgement, relayer,
		)
	}
	// If the cross-chain fails, the token mapping cannot be deleted
	if !ack.Success() {
		return nil
	}

	var data types.NonFungibleTokenPacketData
	if err := types.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		return im.IBCModule.OnAcknowledgementPacket(ctx,
			packet, acknowledgement, relayer,
		)
	}

	classTrace := types.ParseClassTrace(data.ClassId)
	ibcClassId := classTrace.IBCClassID()

	// If it is far away from the original chain, the token mapping cannot be deleted
	if types.IsAwayFromOrigin(packet.GetSourcePort(), packet.GetSourceChannel(), data.ClassId) {
		return nil
	}

	// If it is back to the original chain, delete the token mapping
	_, ok := im.k.ClassToContract(ctx, ibcClassId)
	// If the class mapped to the contract is not found, it may be the erc721 token of the original chain transferred
	if !ok {
		return nil
	}

	// The failure to delete the token mapping does not affect the cross-chain logic, so ignore the error
	_ = im.k.DeleteTokenMapping(ctx, ibcClassId, data.TokenIds)
	return nil
}
