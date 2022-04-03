package types

// ValidateBasic is used for validating the packet
func (p AddLiquidityPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p AddLiquidityPacketData) GetBytes() ([]byte, error) {
	var modulePacket AmmPacketData

	modulePacket.Packet = &AmmPacketData_AddLiquidityPacket{&p}

	return modulePacket.Marshal()
}
