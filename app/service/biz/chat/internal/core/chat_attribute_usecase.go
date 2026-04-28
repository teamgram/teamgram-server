package core

import (
	"math"

	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/internal/repository"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func (c *ChatCore) editChatTitle(chatID, editUserID int64, title string) (*tg.MutableChat, error) {
	if title == "" {
		return nil, chatpb.ErrChatTitleEmpty
	}
	mChat, me, err := c.chatWithNormalParticipant(chatID, editUserID, chatpb.ErrInputUserDeactivated)
	if err != nil {
		return nil, err
	}
	if chatpb.ChatTitle(mChat) == title {
		return nil, chatpb.ErrChatNotModified
	}
	if !chatpb.CanChangeInfo(me) {
		return nil, chatpb.ErrChatAdminRequired
	}
	date, err := c.writeRepository().UpdateChatTitle(c.ctx, chatID, title)
	if err != nil {
		return nil, err
	}
	mChat.Chat.Title = title
	bumpMutableChat(mChat, date)
	return mChat, nil
}

func (c *ChatCore) editChatAbout(chatID, editUserID int64, about string) (*tg.MutableChat, error) {
	mChat, me, err := c.chatWithNormalParticipant(chatID, editUserID, chatpb.ErrInputUserDeactivated)
	if err != nil {
		return nil, err
	}
	if chatpb.ChatAbout(mChat) == about {
		return nil, chatpb.ErrChatNotModified
	}
	if !chatpb.CanChangeInfo(me) {
		return nil, chatpb.ErrChatAdminRequired
	}
	date, err := c.writeRepository().UpdateChatAbout(c.ctx, chatID, about)
	if err != nil {
		return nil, err
	}
	mChat.Chat.About = about
	bumpMutableChat(mChat, date)
	return mChat, nil
}

func (c *ChatCore) editChatPhoto(chatID, editUserID int64, photo tg.PhotoClazz) (*tg.MutableChat, error) {
	mChat, me, err := c.chatWithNormalParticipant(chatID, editUserID, chatpb.ErrInputUserDeactivated)
	if err != nil {
		return nil, err
	}
	if !chatpb.CanChangeInfo(me) {
		return nil, chatpb.ErrChatAdminRequired
	}
	date, err := c.writeRepository().UpdateChatPhoto(c.ctx, chatID, photoID(photo))
	if err != nil {
		return nil, err
	}
	mChat.Chat.Photo = photo
	bumpMutableChat(mChat, date)
	return mChat, nil
}

func (c *ChatCore) editChatAdmin(chatID, operatorID, targetID int64, isAdmin bool) (*tg.MutableChat, error) {
	mChat, err := c.repo().GetMutableChat(c.ctx, chatID, operatorID, targetID)
	if err != nil {
		return nil, err
	}
	me, ok := chatpb.GetImmutableChatParticipant(mChat, operatorID)
	if !ok || !chatpb.IsChatMemberStateNormal(me) {
		return nil, chatpb.ErrUserNotParticipant
	}
	target, ok := chatpb.GetImmutableChatParticipant(mChat, targetID)
	if !ok {
		return nil, chatpb.ErrUserNotParticipant
	}
	if !chatpb.IsChatMemberStateNormal(target) {
		return nil, chatpb.ErrParticipantInvalid
	}
	if !chatpb.CanAdminAddAdmins(me) {
		return nil, chatpb.ErrChatAdminRequired
	}
	if chatpb.IsChatMemberCreator(target) {
		return nil, chatpb.ErrParticipantInvalid
	}
	updated, date, err := c.writeRepository().UpdateChatAdmin(c.ctx, repository.UpdateChatAdminArg{
		ChatID:      chatID,
		Participant: target,
		IsAdmin:     isAdmin,
	})
	if err != nil {
		return nil, err
	}
	replaceParticipant(mChat, updated)
	bumpMutableChat(mChat, date)
	return mChat, nil
}

func (c *ChatCore) editChatDefaultBannedRights(chatID, operatorID int64, rights tg.ChatBannedRightsClazz) (*tg.MutableChat, error) {
	mChat, me, err := c.chatWithNormalParticipant(chatID, operatorID, chatpb.ErrInputUserDeactivated)
	if err != nil {
		return nil, err
	}
	if !chatpb.CanAdminBanUsers(me) {
		return nil, chatpb.ErrChatAdminRequired
	}
	if rights != nil && rights.UntilDate == 0 {
		rights.UntilDate = math.MaxInt32
	}
	date, err := c.writeRepository().UpdateChatDefaultBannedRights(c.ctx, chatID, rights)
	if err != nil {
		return nil, err
	}
	mChat.Chat.DefaultBannedRights = rights
	bumpMutableChat(mChat, date)
	return mChat, nil
}

func (c *ChatCore) toggleNoForwards(chatID, operatorID int64, enabled bool) (*tg.MutableChat, error) {
	mChat, me, err := c.chatWithNormalParticipant(chatID, operatorID, chatpb.ErrInputUserDeactivated)
	if err != nil {
		return nil, err
	}
	if !chatpb.IsChatMemberCreator(me) {
		return nil, chatpb.ErrChatAdminRequired
	}
	date, err := c.writeRepository().UpdateChatNoForwards(c.ctx, chatID, enabled)
	if err != nil {
		return nil, err
	}
	mChat.Chat.Noforwards = enabled
	bumpMutableChat(mChat, date)
	return mChat, nil
}

func (c *ChatCore) setHistoryTTL(selfID, chatID int64, ttlPeriod int32) (*tg.MutableChat, error) {
	mChat, me, err := c.chatWithNormalParticipant(chatID, selfID, chatpb.ErrInputUserDeactivated)
	if err != nil {
		return nil, err
	}
	if !chatpb.IsChatMemberCreator(me) {
		return nil, chatpb.ErrChatAdminRequired
	}
	date, err := c.writeRepository().UpdateChatTTLPeriod(c.ctx, chatID, ttlPeriod)
	if err != nil {
		return nil, err
	}
	mChat.Chat.TtlPeriod = ttlPeriod
	bumpMutableChat(mChat, date)
	return mChat, nil
}

func (c *ChatCore) setChatAvailableReactions(selfID, chatID int64, kind int32, reactions []string) (*tg.MutableChat, error) {
	mChat, me, err := c.chatWithNormalParticipant(chatID, selfID, chatpb.ErrParticipantInvalid)
	if err != nil {
		return nil, err
	}
	if !chatpb.CanAdminAddAdmins(me) {
		return nil, chatpb.ErrChatAdminRequired
	}
	date, err := c.writeRepository().UpdateChatAvailableReactions(c.ctx, chatID, kind, reactions)
	if err != nil {
		return nil, err
	}
	mChat.Chat.AvailableReactionsType = kind
	mChat.Chat.AvailableReactions = reactions
	bumpMutableChat(mChat, date)
	return mChat, nil
}

func (c *ChatCore) chatWithNormalParticipant(chatID, userID int64, missingErr error) (*tg.MutableChat, *tg.ImmutableChatParticipant, error) {
	mChat, err := c.repo().GetMutableChat(c.ctx, chatID, userID)
	if err != nil {
		return nil, nil, err
	}
	me, ok := chatpb.GetImmutableChatParticipant(mChat, userID)
	if !ok || !chatpb.IsChatMemberStateNormal(me) {
		return nil, nil, missingErr
	}
	return mChat, me, nil
}

func bumpMutableChat(mChat *tg.MutableChat, _ int64) {
	if mChat == nil || mChat.Chat == nil {
		return
	}
	mChat.Chat.Version++
}

func replaceParticipant(mChat *tg.MutableChat, participant *tg.ImmutableChatParticipant) {
	if mChat == nil || participant == nil {
		return
	}
	for i, p := range mChat.ChatParticipants {
		if p != nil && p.UserId == participant.UserId {
			mChat.ChatParticipants[i] = participant
			return
		}
	}
	mChat.ChatParticipants = append(mChat.ChatParticipants, participant)
}

func photoID(photo tg.PhotoClazz) int64 {
	switch p := photo.(type) {
	case *tg.TLPhoto:
		return p.Id
	case *tg.TLPhotoEmpty:
		return p.Id
	default:
		return 0
	}
}
