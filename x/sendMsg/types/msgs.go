package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// MsgRoute identifies transaction types
	MsgRoute = "sendMsg"

	TypeMsgSendMsg = "sendMsg"
)

var (
	_ sdk.Msg = &MsgSendMsg{}
)

// NewSendMsg - construct token issue msg.
func NewMsgSendMsg(from, to, message string) *MsgSendMsg {
	return &MsgSendMsg{
		From:    from,
		To:      to,
		Message: message,
	}
}

// Route Implements Msg.
func (msg MsgSendMsg) Route() string { return MsgRoute }

// Type Implements Msg.
func (msg MsgSendMsg) Type() string { return TypeMsgSendMsg }

// ValidateBasic Implements Msg.
func (msg MsgSendMsg) ValidateBasic() error {
	sendmsg := &SendMsg{
		From:    msg.From,
		To:      msg.To,
		Message: msg.Message,
	}

	return sendmsg.Validate()
}

// GetSignBytes Implements Msg.
func (msg MsgSendMsg) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgSendMsg) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}
