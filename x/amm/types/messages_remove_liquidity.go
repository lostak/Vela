package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendRemoveLiquidity = "send_remove_liquidity"

var _ sdk.Msg = &MsgSendRemoveLiquidity{}

func NewMsgSendRemoveLiquidity(
	creator string,
	port string,
	channelID string,
	timeoutTimestamp uint64,
	amountShare string,
	amountOther string,
) *MsgSendRemoveLiquidity {
	return &MsgSendRemoveLiquidity{
		Creator:          creator,
		Port:             port,
		ChannelID:        channelID,
		TimeoutTimestamp: timeoutTimestamp,
		AmountShare:      amountShare,
		AmountOther:      amountOther,
	}
}

func (msg *MsgSendRemoveLiquidity) Route() string {
	return RouterKey
}

func (msg *MsgSendRemoveLiquidity) Type() string {
	return TypeMsgSendRemoveLiquidity
}

func (msg *MsgSendRemoveLiquidity) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendRemoveLiquidity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendRemoveLiquidity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Port == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet port")
	}
	if msg.ChannelID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet channel")
	}
	if msg.TimeoutTimestamp == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet timeout")
	}
	return nil
}
