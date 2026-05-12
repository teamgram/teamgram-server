package core

import (
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func (c *MessagesCore) checkChatMessageAction(chatID int64, action string, mediaKind string) error {
	if c == nil || c.svcCtx == nil || c.svcCtx.Repo == nil || c.svcCtx.Repo.ChatClient == nil {
		return tg.ErrInternalServerError
	}
	_, err := c.svcCtx.Repo.ChatClient.ChatCheckMessageAction(c.ctx, &chatpb.TLChatCheckMessageAction{
		SelfId:    c.MD.UserId,
		ChatId:    chatID,
		Action:    action,
		MediaKind: mediaKind,
	})
	if err != nil {
		mapped := mapChatMessageError(err)
		c.Logger.Errorf("messages.chatMessageAction - chat error mapped: self_user_id=%d chat_id=%d action=%s err=%v mapped=%v", c.MD.UserId, chatID, action, err, mapped)
		return mapped
	}
	return nil
}

func (c *MessagesCore) checkChatAccess(chatID int64, accessKind string) error {
	if c == nil || c.svcCtx == nil || c.svcCtx.Repo == nil || c.svcCtx.Repo.ChatClient == nil {
		return tg.ErrInternalServerError
	}
	_, err := c.svcCtx.Repo.ChatClient.ChatCheckChatAccess(c.ctx, &chatpb.TLChatCheckChatAccess{
		SelfId:     c.MD.UserId,
		ChatId:     chatID,
		AccessKind: accessKind,
	})
	if err != nil {
		mapped := mapChatMessageError(err)
		c.Logger.Errorf("messages.chatAccess - chat error mapped: self_user_id=%d chat_id=%d access_kind=%s err=%v mapped=%v", c.MD.UserId, chatID, accessKind, err, mapped)
		return mapped
	}
	return nil
}
