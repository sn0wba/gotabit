package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gotabit/gotabit/x/msg/types"
)

type msgServer struct {
	*Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the token MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (m msgServer) Msg(goCtx context.Context, msg *types.MsgMsg) (*types.MsgMsgResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id, err := m.Keeper.Send(ctx, msg)
	if err != nil {
		return nil, err
	}

	return &types.MsgMsgResponse{
		Id: id,
	}, nil
}
