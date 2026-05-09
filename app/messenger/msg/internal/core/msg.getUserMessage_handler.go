package core

import (
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func (c *MsgCore) MsgGetUserMessage(in *msg.TLMsgGetUserMessage) (*tg.MessageBox, error) {
	if in == nil {
		return nil, msg.ErrMsgIdInvalid
	}
	box, err := c.svcCtx.Repo.GetUserMessage(c.ctx, in.UserId, int64(in.Id))
	if err != nil {
		return nil, err
	}
	return messageBoxFromUserMessage(box)
}

func messageBoxFromUserMessage(box *repository.UserMessageBox) (*tg.MessageBox, error) {
	if box == nil {
		return nil, msg.ErrMsgIdInvalid
	}
	messageID, err := historyIDInt32(box.UserMessageID, "user message id")
	if err != nil {
		return nil, err
	}
	date, err := msgDateInt32FromUnixSeconds(box.MessageDate, "user message date")
	if err != nil {
		return nil, err
	}
	message := tg.MakeTLMessage(&tg.TLMessage{
		Out:     box.Outgoing,
		Id:      messageID,
		FromId:  tg.MakePeerUser(box.FromUserID),
		PeerId:  sentMessagePeerFromOptional(box.PeerType, box.PeerID),
		Date:    date,
		Message: box.MessageText,
	})
	return tg.MakeTLMessageBox(&tg.TLMessageBox{
		UserId:          box.UserID,
		MessageId:       messageID,
		SenderUserId:    box.FromUserID,
		PeerType:        box.PeerType,
		PeerId:          box.PeerID,
		DialogId1:       box.UserID,
		DialogId2:       box.PeerID,
		DialogMessageId: box.UserMessageID,
		Message:         message,
		Reaction:        "",
	}).ToMessageBox(), nil
}
