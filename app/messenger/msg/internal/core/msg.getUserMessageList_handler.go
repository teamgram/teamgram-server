package core

import (
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func (c *MsgCore) MsgGetUserMessageList(in *msg.TLMsgGetUserMessageList) (*msg.VectorMessageBox, error) {
	if in == nil {
		return nil, msg.ErrMsgIdInvalid
	}
	ids := make([]int64, 0, len(in.IdList))
	for _, id := range in.IdList {
		ids = append(ids, int64(id))
	}
	boxes, err := c.svcCtx.Repo.GetUserMessageList(c.ctx, in.UserId, ids)
	if err != nil {
		return nil, err
	}
	out := &msg.VectorMessageBox{Datas: make([]tg.MessageBoxClazz, 0, len(boxes))}
	for i := range boxes {
		box, err := messageBoxFromUserMessage(&boxes[i])
		if err != nil {
			return nil, err
		}
		out.Datas = append(out.Datas, box)
	}
	return out, nil
}
