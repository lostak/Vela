package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendAddLiquidity = "send_add_liquidity"

var _ sdk.Msg = &MsgSendAddLiquidity{}

func NewMsgSendAddLiquidity(
	creator string,
	port string,
	channelID string,
	timeoutTimestamp uint64,
	amountVela string,
	amountOther string,
) *MsgSendAddLiquidity {
	return &MsgSendAddLiquidity{
		Creator:          creator,
		Port:             port,
		ChannelID:        channelID,
		TimeoutTimestamp: timeoutTimestamp,
		AmountVela:       amountVela,
		AmountOther:      amountOther,
	}
}

func (msg *MsgSendAddLiquidity) Route() string {
	return RouterKey
}

func (msg *MsgSendAddLiquidity) Type() string {
	return TypeMsgSendAddLiquidity
}

func (msg *MsgSendAddLiquidity) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendAddLiquidity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendAddLiquidity) ValidateBasic() error {
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
