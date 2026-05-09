package core

import (
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func (c *MsgCore) MsgGetUserMessage(in *msg.TLMsgGetUserMessage) (*tg.MessageBox, error) {
	return nil, fmt.Errorf("%w: msg.getUserMessage is not wired until Task 8", msg.ErrMethodNotImplemented)
}
