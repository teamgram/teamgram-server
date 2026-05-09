package core

import (
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
)

func (c *MsgCore) MsgGetUserMessageList(in *msg.TLMsgGetUserMessageList) (*msg.VectorMessageBox, error) {
	return nil, fmt.Errorf("%w: msg.getUserMessageList is not wired until Task 8", msg.ErrMethodNotImplemented)
}
