package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v2/modules/core/02-client/types"
	"github.com/lostak/Vela/x/amm/types"
)

func (k msgServer) SendRemoveLiquidity(goCtx context.Context, msg *types.MsgSendRemoveLiquidity) (*types.MsgSendRemoveLiquidityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.RemoveLiquidityPacketData

	packet.AmountShare = msg.AmountShare
	packet.AmountOther = msg.AmountOther

	// Transmit the packet
	err := k.TransmitRemoveLiquidityPacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendRemoveLiquidityResponse{}, nil
}
