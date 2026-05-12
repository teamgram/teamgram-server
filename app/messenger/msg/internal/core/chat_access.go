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

func (c *MsgCore) activeChatReceiverIDs(chatID int64, actorUserID int64) ([]int64, error) {
	memberIDs, err := c.activeChatMemberIDs(chatID)
	if err != nil {
		return nil, err
	}
	out := make([]int64, 0, len(memberIDs))
	seen := make(map[int64]struct{}, len(memberIDs))
	for _, memberID := range memberIDs {
		if memberID <= 0 || memberID == actorUserID {
			continue
		}
		if _, ok := seen[memberID]; ok {
			continue
		}
		seen[memberID] = struct{}{}
		out = append(out, memberID)
	}
	return out, nil
}

func chatSendActionForNormalized(normalized normalizedOutboxMessage) (action string, mediaKind string) {
	if normalized.ForwardRef != nil {
		return chatpb.MessageActionForwardToChat, "forward"
	}
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

func (c *MsgCore) checkChatBatchActions(userID int64, chatID int64, normalizedBatch []normalizedOutboxMessage) error {
	if len(normalizedBatch) > 1 {
		if err := c.checkChatAction(userID, chatID, chatpb.MessageActionSendAlbum, "album"); err != nil {
			return err
		}
	}
	checked := make(map[string]struct{})
	for _, normalized := range normalizedBatch {
		action, mediaKind := chatSendActionForNormalized(normalized)
		key := chatMessageActionKey(action, mediaKind)
		if _, ok := checked[key]; ok {
			continue
		}
		if err := c.checkChatAction(userID, chatID, action, mediaKind); err != nil {
			return err
		}
		checked[key] = struct{}{}
	}
	return nil
}

func chatMessageActionKey(action string, mediaKind string) string {
	return action + "\x00" + mediaKind
}
