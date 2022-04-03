package types

// IBC events
const (
	EventTypeTimeout            = "timeout"
	EventTypeCreatePoolPacket   = "createPool_packet"
	EventTypeAddLiquidityPacket = "addLiquidity_packet"
	// this line is used by starport scaffolding # ibc/packet/event

	AttributeKeyAckSuccess = "success"
	AttributeKeyAck        = "acknowledgement"
	AttributeKeyAckError   = "error"
)
