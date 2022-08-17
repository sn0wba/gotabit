package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// MsgRoute identifies transaction types
	MsgRoute = "msg"

	TypeMsgMsg = "msg"
)

var (
	_ sdk.Msg = &MsgMsg{}
)

// NewMsg - construct token issue msg.
func NewMsgMsg(sender, from, to, message string) *MsgMsg {
	return &MsgMsg{
		Sender:  sender,
		From:    from,
		To:      to,
		Message: message,
	}
}

// Route Implements Msg.
func (mm MsgMsg) Route() string { return MsgRoute }

// Type Implements Msg.
func (mm MsgMsg) Type() string { return TypeMsgMsg }

// ValidateBasic Implements Msg.
func (mm MsgMsg) ValidateBasic() error {
	msg := &Msg{
		From:    mm.From,
		To:      mm.To,
		Message: mm.Message,
	}

	return msg.Validate()
}

// GetSignBytes Implements Msg.
func (mm MsgMsg) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&mm)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (mm MsgMsg) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(mm.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}
