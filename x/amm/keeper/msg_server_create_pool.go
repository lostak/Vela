package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v2/modules/core/02-client/types"
	"github.com/lostak/Vela/x/amm/types"
)

func (k msgServer) SendCreatePool(goCtx context.Context, msg *types.MsgSendCreatePool) (*types.MsgSendCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.CreatePoolPacketData

	packet.PoolParam = msg.PoolParam
	packet.PoolAssets = msg.PoolAssets

	// Transmit the packet
	err := k.TransmitCreatePoolPacket(
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

	return &types.MsgSendCreatePoolResponse{}, nil
}
