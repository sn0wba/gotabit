package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// module sentinel errors
var (
	ErrSendMsgDoesNotExist = sdkerrors.Register(ModuleName, 1, "SendMsg does not exist")
)
