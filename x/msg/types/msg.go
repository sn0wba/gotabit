package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
	"gopkg.in/yaml.v2"
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
	if len(sm.From) == 0 {
		return sdkerrors.Wrapf(sdkerrors.Error{}, "missing from")
	}
	if len(sm.From) > 64 {
		return sdkerrors.Wrapf(sdkerrors.Error{}, "from too long")
	}

	if len(sm.To) == 0 {
		return sdkerrors.Wrapf(sdkerrors.Error{}, "missing to")
	}
	if len(sm.To) > 64 {
		return sdkerrors.Wrapf(sdkerrors.Error{}, "to too long")
	}

	if len(sm.Message) == 0 {
		return sdkerrors.Wrapf(sdkerrors.Error{}, "missing message")
	}
	if len(sm.Message) > 512 {
		return sdkerrors.Wrapf(sdkerrors.Error{}, "message too long")
	}

	return nil
}
