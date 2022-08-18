package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gotabit/gotabit/x/msg/keeper"
	"github.com/gotabit/gotabit/x/msg/types"
)

func (suite *KeeperTestSuite) SendMsg(sender sdk.AccAddress, from, to, message string) *types.MsgMsgResponse {
	msgServer := keeper.NewMsgServerImpl(&suite.app.MsgKeeper)
	resp, err := msgServer.Msg(sdk.WrapSDKContext(suite.ctx), types.NewMsgMsg(
		sender.String(), from, to, message,
	))
	suite.Require().NoError(err)
	return resp
}

func (suite *KeeperTestSuite) TestMsgServerSendMsg() {
	tests := []struct {
		testCase      string
		from          string
		to            string
		message       string
		expectPass    bool
		expectedMsgId uint64
	}{
		{
			"send message",
			"from",
			"to",
			"test message",
			true,
			1,
		},
		{
			"send empty message",
			"from",
			"to",
			"",
			false,
			0,
		},
	}

	for _, tc := range tests {
		sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

		msgServer := keeper.NewMsgServerImpl(&suite.app.MsgKeeper)
		resp, err := msgServer.Msg(sdk.WrapSDKContext(suite.ctx), types.NewMsgMsg(
			sender.String(), tc.from, tc.to, tc.message,
		))
		if tc.expectPass {
			suite.Require().NoError(err)

			// test response is correct
			suite.Require().Equal(resp.Id, tc.expectedMsgId)

			// test lastMsgId is updated correctly
			lastMsgId := suite.app.MsgKeeper.GetLastMsgId(suite.ctx)
			suite.Require().Equal(lastMsgId, tc.expectedMsgId)

			// test metadataId and nftId to set correctly
			msg, err := suite.app.MsgKeeper.GetMsgById(suite.ctx, resp.Id)
			suite.Require().NoError(err)
			suite.Require().Equal(msg.Id, tc.expectedMsgId)
		} else {
			suite.Require().Error(err)
		}
	}
}
