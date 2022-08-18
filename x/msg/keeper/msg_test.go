package keeper_test

import (
	"github.com/gotabit/gotabit/x/msg/types"
)

func (suite *KeeperTestSuite) TestLastMsgIdGetSet() {
	// get default last msg id
	lastMsgId := suite.app.MsgKeeper.GetLastMsgId(suite.ctx)
	suite.Require().Equal(lastMsgId, uint64(0))

	// set last msg id to new value
	newMsgId := uint64(2)
	suite.app.MsgKeeper.SetLastMsgId(suite.ctx, newMsgId)

	// check last msg id update
	lastMsgId = suite.app.MsgKeeper.GetLastMsgId(suite.ctx)
	suite.Require().Equal(lastMsgId, newMsgId)
}

func (suite *KeeperTestSuite) TestMsgGetSet() {
	// get msg by not available id
	_, err := suite.app.MsgKeeper.GetMsgById(suite.ctx, 0)
	suite.Require().Error(err)

	// get sent msgs when not available
	sentMsgs := suite.app.MsgKeeper.GetMsgsBySender(suite.ctx, "gio13m350fvnk3s6y5n8ugxhmka277r0t7cw48ru47")
	suite.Require().Len(sentMsgs, 0)

	// get received msgs when not available
	receivedMsgs := suite.app.MsgKeeper.GetMsgsByReceiver(suite.ctx, "gio13m350fvnk3s6y5n8ugxhmka277r0t7cw48ru47")
	suite.Require().Len(receivedMsgs, 0)

	// create a new msg
	msgs := []*types.Msg{
		{
			Id:      1,
			From:    "gio13m350fvnk3s6y5n8ugxhmka277r0t7cw48ru47",
			To:      "gio1pwhmp2d53crcqervmv5c6h4chdnctkvjaya9vs",
			Message: "message 1",
		},
		{
			Id:      2,
			From:    "gio13m350fvnk3s6y5n8ugxhmka277r0t7cw48ru47",
			To:      "gio1daxjpnra6jpahzjg8e6c865hmtt7469n249ln2",
			Message: "message 2",
		},
		{
			Id:      3,
			From:    "gio13m350fvnk3s6y5n8ugxhmka277r0t7cw48ru47",
			To:      "gio1daxjpnra6jpahzjg8e6c865hmtt7469n249ln2",
			Message: "message 3",
		},
	}

	for _, msg := range msgs {
		suite.app.MsgKeeper.SetMsg(suite.ctx, msg)
	}

	for _, msg := range msgs {
		c, err := suite.app.MsgKeeper.GetMsgById(suite.ctx, msg.Id)
		suite.Require().NoError(err)
		suite.Require().Equal(*msg, *c)
	}

	sentMsgs = suite.app.MsgKeeper.GetMsgsBySender(suite.ctx, "gio13m350fvnk3s6y5n8ugxhmka277r0t7cw48ru47")
	suite.Require().Len(sentMsgs, 3)
	for i, msg := range msgs {
		suite.Require().Equal(*msg, *sentMsgs[i])
	}

	receivedMsgs = suite.app.MsgKeeper.GetMsgsByReceiver(suite.ctx, "gio1daxjpnra6jpahzjg8e6c865hmtt7469n249ln2")
	suite.Require().Len(receivedMsgs, 2)
}
