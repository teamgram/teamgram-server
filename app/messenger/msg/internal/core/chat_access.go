package core

import (
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
)

func (c *MsgCore) checkChatAction(userID int64, chatID int64, action string, mediaKind string) error {
	if c == nil || c.svcCtx == nil || c.svcCtx.Chat == nil {
		return fmt.Errorf("%w: missing chat client", msg.ErrSendStateConflict)
	}
	_, err := c.svcCtx.Chat.ChatCheckMessageAction(c.ctx, &chatpb.TLChatCheckMessageAction{
		SelfId:    userID,
		ChatId:    chatID,
		Action:    action,
		MediaKind: mediaKind,
	})
	return err
}

func (c *MsgCore) checkChatAccess(userID int64, chatID int64, accessKind string) error {
	if c == nil || c.svcCtx == nil || c.svcCtx.Chat == nil {
		return fmt.Errorf("%w: missing chat client", msg.ErrSendStateConflict)
	}
	_, err := c.svcCtx.Chat.ChatCheckChatAccess(c.ctx, &chatpb.TLChatCheckChatAccess{
		SelfId:     userID,
		ChatId:     chatID,
		AccessKind: accessKind,
	})
	return err
}

func (c *MsgCore) activeChatMemberIDs(chatID int64) ([]int64, error) {
	if c == nil || c.svcCtx == nil || c.svcCtx.Chat == nil {
		return nil, fmt.Errorf("%w: missing chat client", msg.ErrSendStateConflict)
	}
	result, err := c.svcCtx.Chat.ChatGetChatParticipantIdList(c.ctx, &chatpb.TLChatGetChatParticipantIdList{ChatId: chatID})
	if err != nil {
		return nil, err
	}
	if result == nil {
		return []int64{}, nil
	}
	return append([]int64(nil), result.Datas...), nil
}

func chatSendActionForNormalized(normalized normalizedOutboxMessage) (action string, mediaKind string) {
	if normalized.MediaRef == nil {
		return chatpb.MessageActionSendText, ""
	}
	switch normalized.MediaRef.Kind {
	case "photo":
		return chatpb.MessageActionSendMediaPhoto, "photo"
	case "document":
		return chatpb.MessageActionSendMediaDoc, "document"
	default:
		return chatpb.MessageActionSendMediaDoc, normalized.MediaRef.Kind
	}
}
