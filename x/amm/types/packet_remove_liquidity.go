package types

// ValidateBasic is used for validating the packet
func (p RemoveLiquidityPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p RemoveLiquidityPacketData) GetBytes() ([]byte, error) {
	var modulePacket AmmPacketData

	modulePacket.Packet = &AmmPacketData_RemoveLiquidityPacket{&p}

	return modulePacket.Marshal()
}
