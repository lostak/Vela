package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v2/modules/core/02-client/types"
	"github.com/lostak/Vela/x/amm/types"
)

func (k msgServer) SendAddLiquidity(goCtx context.Context, msg *types.MsgSendAddLiquidity) (*types.MsgSendAddLiquidityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.AddLiquidityPacketData

	packet.AmountVela = msg.AmountVela
	packet.AmountOther = msg.AmountOther

	// Transmit the packet
	err := k.TransmitAddLiquidityPacket(
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

	return &types.MsgSendAddLiquidityResponse{}, nil
}
