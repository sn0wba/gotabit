package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
	"gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ proto.Message = &SendMsg{}
)

// NewSendMsg constructs a new SendMsg instance
func NewSendMsg(from, to, message string) *SendMsg {
	return &SendMsg{
		From:    from,
		To:      to,
		Message: message,
	}
}

// GetFrom implements exported.SendMsgI
func (sm SendMsg) GetId() uint64 {
	return sm.Id
}

// GetFrom implements exported.SendMsgI
func (sm SendMsg) GetFrom() string {
	return sm.From
}

// GetDenom implements exported.SendMsgI
func (sm SendMsg) GetTo() string {
	return sm.To
}

// GetMessage implements exported.SendMsgI
func (sm SendMsg) GetMessage() string {
	return sm.Message
}

func (sm SendMsg) String() string {
	bz, _ := yaml.Marshal(sm)
	return string(bz)
}

func (sm SendMsg) Validate() error {
	if _, err := sdk.AccAddressFromBech32(sm.From); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address (%s)", err)
	}
	if _, err := sdk.AccAddressFromBech32(sm.To); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid to address (%s)", err)
	}
	if len(sm.Message) == 0 {
		return sdkerrors.Wrapf(sdkerrors.Error{}, "missing message")
	}

	return nil
}