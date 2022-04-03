package types

// ValidateBasic is used for validating the packet
func (p CreatePoolPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p CreatePoolPacketData) GetBytes() ([]byte, error) {
	var modulePacket AmmPacketData

	modulePacket.Packet = &AmmPacketData_CreatePoolPacket{&p}

	return modulePacket.Marshal()
}
