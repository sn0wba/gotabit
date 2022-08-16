package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
	"gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ proto.Message = &Msg{}
)

// NewMsg constructs a new Msg instance
func NewMsg(from, to, message string) *Msg {
	return &Msg{
		From:    from,
		To:      to,
		Message: message,
	}
}

// GetFrom implements exported.MsgI
func (sm Msg) GetId() uint64 {
	return sm.Id
}

// GetFrom implements exported.MsgI
func (sm Msg) GetFrom() string {
	return sm.From
}

// GetDenom implements exported.MsgI
func (sm Msg) GetTo() string {
	return sm.To
}

// GetMessage implements exported.MsgI
func (sm Msg) GetMessage() string {
	return sm.Message
}

func (sm Msg) String() string {
	bz, _ := yaml.Marshal(sm)
	return string(bz)
}

func (sm Msg) Validate() error {
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
