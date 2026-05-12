package core

import (
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func (c *ChatCore) loadMessageActionChat(chatID int64, userID int64) (*tg.MutableChat, *tg.ImmutableChatParticipant, error) {
	mChat, err := c.repo().GetMutableChat(c.ctx, chatID, userID)
	if err != nil {
		return nil, nil, err
	}
	if mChat == nil || mChat.Chat == nil || mChat.Chat.Deactivated {
		return nil, nil, chat.ErrChatNotFound
	}
	participant, ok := chat.GetImmutableChatParticipant(mChat, userID)
	if !ok {
		return mChat, nil, chat.ErrUserNotParticipant
	}
	return mChat, participant, nil
}

func checkSupportedMessageAction(action string) error {
	switch action {
	case chat.MessageActionSendPoll, chat.MessageActionSendInline, chat.MessageActionSendGame:
		return chat.ErrMessageActionUnsupported
	default:
		return nil
	}
}

func checkMessageBannedRights(mChat *tg.MutableChat, participant *tg.ImmutableChatParticipant, action string) error {
	switch action {
	case chat.MessageActionPinMessage, chat.MessageActionUnpinAll:
		if chat.CanPinMessages(participant) {
			return nil
		}
		return chat.ErrChatAdminRequired
	case chat.MessageActionDeleteRevoke:
		if chat.CanAdminBanUsers(participant) {
			return nil
		}
		return chat.ErrChatAdminRequired
	}
	if chat.IsChatMemberCreator(participant) || chat.IsChatMemberAdmin(participant) {
		return nil
	}
	rights := chat.ChatDefaultBannedRights(mChat)
	if rights == nil {
		return nil
	}
	switch action {
	case chat.MessageActionSendText, chat.MessageActionEditOwnMessage:
		if rights.SendMessages || rights.SendPlain {
			return chat.ErrChatWriteForbidden
		}
	case chat.MessageActionSendMediaPhoto:
		if rights.SendMessages || rights.SendMedia || rights.SendPhotos {
			return chat.ErrChatWriteForbidden
		}
	case chat.MessageActionSendMediaDoc:
		if rights.SendMessages || rights.SendMedia || rights.SendDocs {
			return chat.ErrChatWriteForbidden
		}
	case chat.MessageActionSendAlbum, chat.MessageActionForwardToChat:
		if rights.SendMessages || rights.SendMedia {
			return chat.ErrChatWriteForbidden
		}
	}
	return nil
}
