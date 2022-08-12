package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/gotabit/gotabit/x/sendMsg/types"
)

func (k Keeper) GetLastSendMsgId(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyLastSendMsgId)
	if bz == nil {
		return 0
	}
	return sdk.BigEndianToUint64(bz)
}

func (k Keeper) SetLastSendMsgId(ctx sdk.Context, id uint64) {
	idBz := sdk.Uint64ToBigEndian(id)
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyLastSendMsgId, idBz)
}

func (k Keeper) GetSendMsgById(ctx sdk.Context, id uint64) (*types.SendMsg, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(append(types.PrefixSendMsg, sdk.Uint64ToBigEndian(id)...))
	if bz == nil {
		return &types.SendMsg{}, sdkerrors.Wrapf(types.ErrSendMsgDoesNotExist, "SendMsg: %d does not exist", id)
	}
	msg := types.SendMsg{}
	k.cdc.MustUnmarshal(bz, &msg)
	return &msg, nil
}

func (k Keeper) GetSendMsgsBySender(ctx sdk.Context, sender string) []*types.SendMsg {
	store := ctx.KVStore(k.storeKey)

	msgs := []*types.SendMsg{}
	it := sdk.KVStorePrefixIterator(store, append(types.PrefixSendMsgBySender, sender...))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		msgId := sdk.BigEndianToUint64(it.Value())
		if msgId == 0 {
			continue
		}
		msg, err := k.GetSendMsgById(ctx, msgId)
		if err != nil {
			panic(err)
		}

		msgs = append(msgs, msg)
	}
	return msgs
}

func (k Keeper) GetSendMsgsByReceiver(ctx sdk.Context, receiver string) []*types.SendMsg {
	store := ctx.KVStore(k.storeKey)

	msgs := []*types.SendMsg{}
	it := sdk.KVStorePrefixIterator(store, append(types.PrefixSendMsgByReceiver, receiver...))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		msgId := sdk.BigEndianToUint64(it.Value())
		if msgId == 0 {
			continue
		}
		msg, err := k.GetSendMsgById(ctx, msgId)
		if err != nil {
			panic(err)
		}

		msgs = append(msgs, msg)
	}
	return msgs
}

func (k Keeper) SetSendMsg(ctx sdk.Context, msg *types.SendMsg) {
	idBz := sdk.Uint64ToBigEndian(msg.Id)
	bz := k.cdc.MustMarshal(msg)
	store := ctx.KVStore(k.storeKey)
	store.Set(append(types.PrefixSendMsg, idBz...), bz)
	store.Set(append(append(types.PrefixSendMsgBySender, msg.From...), idBz...), idBz)
	store.Set(append(append(types.PrefixSendMsgByReceiver, msg.To...), idBz...), idBz)
}

// Send sends a new message
func (k Keeper) Send(ctx sdk.Context, msg *types.MsgSendMsg) (id uint64, err error) {
	sendMsg := types.NewSendMsg(msg.From, msg.To, msg.Message)

	if err := sendMsg.Validate(); err != nil {
		return id, err
	}

	msgId := k.GetLastSendMsgId(ctx) + 1
	k.SetLastSendMsgId(ctx, msgId)

	sendMsg.Id = msgId
	k.SetSendMsg(ctx, sendMsg)
	ctx.EventManager().EmitTypedEvent(&types.EventSend{
		Id: id,
	})

	return msgId, nil
}
