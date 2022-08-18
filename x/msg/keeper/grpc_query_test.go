package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gotabit/gotabit/x/msg/types"
)

func (suite *KeeperTestSuite) TestGRPCSentMsgs() {
	// send messages
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	msg1 := suite.SendMsg(sender, "alice", "bob", "message1")
	msg2 := suite.SendMsg(sender, "alice", "carol", "message2")
	msg3 := suite.SendMsg(sender, "alice", "bob", "message3")

	tests := []struct {
		testCase     string
		sender       string
		expectedMsgs []*types.MsgMsgResponse
	}{
		{
			"query alice's sent messages",
			"alice",
			[]*types.MsgMsgResponse{
				{
					Id:      1,
					From:    msg1.From,
					To:      msg1.To,
					Message: msg1.Message,
				},
				{
					Id:      2,
					From:    msg2.From,
					To:      msg2.To,
					Message: msg2.Message,
				},
				{
					Id:      3,
					From:    msg3.From,
					To:      msg3.To,
					Message: msg3.Message,
				},
			},
		},
		{
			"query bob's sent messages",
			"bob",
			[]*types.MsgMsgResponse{},
		},
	}

	for _, tc := range tests {
		resp := suite.app.MsgKeeper.GetMsgsBySender(suite.ctx, tc.sender)
		suite.Require().Equal(len(resp), len(tc.expectedMsgs))
		for i, msg := range resp {
			suite.Require().Equal(msg.Id, tc.expectedMsgs[i].Id)
			suite.Require().Equal(msg.From, tc.expectedMsgs[i].From)
			suite.Require().Equal(msg.To, tc.expectedMsgs[i].To)
			suite.Require().Equal(msg.Message, tc.expectedMsgs[i].Message)
		}
	}
}

func (suite *KeeperTestSuite) TestGRPCReceivedMsgs() {
	// send messages
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	msg1 := suite.SendMsg(sender, "alice", "bob", "message1")
	msg2 := suite.SendMsg(sender, "alice", "carol", "message2")
	msg3 := suite.SendMsg(sender, "alice", "bob", "message3")

	tests := []struct {
		testCase     string
		receiver     string
		expectedMsgs []*types.MsgMsgResponse
	}{
		{
			"query bob's received messages",
			"bob",
			[]*types.MsgMsgResponse{
				{
					Id:      1,
					From:    msg1.From,
					To:      msg1.To,
					Message: msg1.Message,
				},
				{
					Id:      3,
					From:    msg3.From,
					To:      msg3.To,
					Message: msg3.Message,
				},
			},
		},
		{
			"query carol's received messages",
			"carol",
			[]*types.MsgMsgResponse{
				{
					Id:      2,
					From:    msg2.From,
					To:      msg2.To,
					Message: msg2.Message,
				},
			},
		},
	}

	for _, tc := range tests {
		resp := suite.app.MsgKeeper.GetMsgsByReceiver(suite.ctx, tc.receiver)
		suite.Require().Equal(len(resp), len(tc.expectedMsgs))
		for i, msg := range resp {
			suite.Require().Equal(msg.Id, tc.expectedMsgs[i].Id)
			suite.Require().Equal(msg.From, tc.expectedMsgs[i].From)
			suite.Require().Equal(msg.To, tc.expectedMsgs[i].To)
			suite.Require().Equal(msg.Message, tc.expectedMsgs[i].Message)
		}
	}
}
