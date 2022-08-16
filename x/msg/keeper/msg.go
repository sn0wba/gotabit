package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/gotabit/gotabit/x/msg/types"
)

func (k Keeper) GetLastMsgId(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyLastMsgId)
	if bz == nil {
		return 0
	}
	return sdk.BigEndianToUint64(bz)
}

func (k Keeper) SetLastMsgId(ctx sdk.Context, id uint64) {
	idBz := sdk.Uint64ToBigEndian(id)
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyLastMsgId, idBz)
}

func (k Keeper) GetMsgById(ctx sdk.Context, id uint64) (*types.Msg, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(append(types.PrefixMsg, sdk.Uint64ToBigEndian(id)...))
	if bz == nil {
		return &types.Msg{}, sdkerrors.Wrapf(types.ErrMsgDoesNotExist, "Msg: %d does not exist", id)
	}
	msg := types.Msg{}
	k.cdc.MustUnmarshal(bz, &msg)
	return &msg, nil
}

func (k Keeper) GetMsgsBySender(ctx sdk.Context, sender string) []*types.Msg {
	store := ctx.KVStore(k.storeKey)

	msgs := []*types.Msg{}
	it := sdk.KVStorePrefixIterator(store, append(types.PrefixMsgBySender, sender...))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		msgId := sdk.BigEndianToUint64(it.Value())
		if msgId == 0 {
			continue
		}
		msg, err := k.GetMsgById(ctx, msgId)
		if err != nil {
			panic(err)
		}

		msgs = append(msgs, msg)
	}
	return msgs
}

func (k Keeper) GetMsgsByReceiver(ctx sdk.Context, receiver string) []*types.Msg {
	store := ctx.KVStore(k.storeKey)

	msgs := []*types.Msg{}
	it := sdk.KVStorePrefixIterator(store, append(types.PrefixMsgByReceiver, receiver...))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		msgId := sdk.BigEndianToUint64(it.Value())
		if msgId == 0 {
			continue
		}
		msg, err := k.GetMsgById(ctx, msgId)
		if err != nil {
			panic(err)
		}

		msgs = append(msgs, msg)
	}
	return msgs
}

func (k Keeper) SetMsg(ctx sdk.Context, msg *types.Msg) {
	idBz := sdk.Uint64ToBigEndian(msg.Id)
	bz := k.cdc.MustMarshal(msg)
	store := ctx.KVStore(k.storeKey)
	store.Set(append(types.PrefixMsg, idBz...), bz)
	store.Set(append(append(types.PrefixMsgBySender, msg.From...), idBz...), idBz)
	store.Set(append(append(types.PrefixMsgByReceiver, msg.To...), idBz...), idBz)
}

// Send sends a new message
func (k Keeper) Send(ctx sdk.Context, mm *types.MsgMsg) (id uint64, err error) {
	msg := types.NewMsg(mm.From, mm.To, mm.Message)

	if err := msg.Validate(); err != nil {
		return id, err
	}

	msgId := k.GetLastMsgId(ctx) + 1
	k.SetLastMsgId(ctx, msgId)

	msg.Id = msgId
	k.SetMsg(ctx, msg)
	ctx.EventManager().EmitTypedEvent(&types.EventSend{
		Id: id,
	})

	return msgId, nil
}
