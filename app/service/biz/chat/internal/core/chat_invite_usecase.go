package core

import (
	"time"

	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func (c *ChatCore) checkChatInvite(selfID int64, hash string) (*chatpb.ChatInviteExt, error) {
	invite, mChat, err := c.loadUsableInvite(hash, selfID)
	if err != nil {
		return nil, err
	}
	me, _ := chatpb.GetImmutableChatParticipant(mChat, selfID)
	if chatpb.IsChatMemberStateNormal(me) {
		return chatpb.MakeTLChatInviteAlready(&chatpb.TLChatInviteAlready{
			Chat: mChat,
		}).ToChatInviteExt(), nil
	}
	return chatpb.MakeTLChatInvite(&chatpb.TLChatInvite{
		RequestNeeded:     invite.RequestNeeded,
		Title:             chatpb.ChatTitle(mChat),
		About:             stringPtr(chatpb.ChatAbout(mChat)),
		Photo:             invitePhoto(mChat),
		ParticipantsCount: chatpb.ChatParticipantsCount(mChat),
		Participants:      chatpb.ChatParticipantIDList(mChat),
	}).ToChatInviteExt(), nil
}

func (c *ChatCore) importChatInvite(selfID int64, hash string) (*tg.MutableChat, error) {
	invite, _, err := c.loadUsableInvite(hash, selfID)
	if err != nil {
		return nil, err
	}
	mChat, err := c.addChatUser(c.ctx, addChatUserArg{
		chatID:                  invite.ChatId,
		inviterID:               invite.AdminId,
		userID:                  selfID,
		recordInviteParticipant: true,
		inviteLink:              invite.Link,
	})
	if err != nil {
		return nil, err
	}
	return mChat, nil
}

func (c *ChatCore) importChatInvite2(selfID int64, hash string) (*chatpb.ChatInviteImported, error) {
	invite, mChat, err := c.loadUsableInvite(hash, selfID)
	if err != nil {
		return nil, err
	}
	if invite.RequestNeeded {
		if chatpb.ChatParticipantsCount(mChat) >= maxChatParticipants {
			return nil, chatpb.ErrUsersTooMuch
		}
		me, _ := chatpb.GetImmutableChatParticipant(mChat, selfID)
		if chatpb.IsChatMemberStateNormal(me) {
			return nil, chatpb.ErrUserAlreadyParticipant
		}
		if err := c.inviteRepository().RecordInviteParticipant(c.ctx, repository.InviteParticipantArg{
			ChatID:    invite.ChatId,
			Link:      invite.Link,
			UserID:    selfID,
			Requested: true,
		}); err != nil {
			return nil, err
		}
		requesters, err := c.inviteRepository().GetRecentChatInviteRequesters(c.ctx, invite.ChatId)
		if err != nil {
			return nil, err
		}
		return chatpb.MakeTLChatInviteImported(&chatpb.TLChatInviteImported{
			Chat:       mChat,
			Requesters: requesters,
		}).ToChatInviteImported(), nil
	}

	mChat, err = c.addChatUser(c.ctx, addChatUserArg{
		chatID:                  invite.ChatId,
		inviterID:               invite.AdminId,
		userID:                  selfID,
		recordInviteParticipant: true,
		inviteLink:              invite.Link,
	})
	if err != nil {
		return nil, err
	}
	return chatpb.MakeTLChatInviteImported(&chatpb.TLChatInviteImported{
		Chat: mChat,
	}).ToChatInviteImported(), nil
}

func (c *ChatCore) loadUsableInvite(hash string, participantIDs ...int64) (*model.ChatInvites, *tg.MutableChat, error) {
	invite, err := c.inviteRepository().GetChatInviteByLink(c.ctx, hash)
	if err != nil {
		return nil, nil, err
	}
	usage, err := c.inviteRepository().CountChatInviteParticipants(c.ctx, invite.Link, false)
	if err != nil {
		return nil, nil, err
	}
	if repository.IsInviteExpired(invite, usage, time.Now().Unix()) {
		return nil, nil, chatpb.ErrInviteHashExpired
	}
	ids := append([]int64{invite.AdminId}, participantIDs...)
	mChat, err := c.repo().GetMutableChat(c.ctx, invite.ChatId, ids...)
	if err != nil {
		return nil, nil, err
	}
	if data := chatpb.MutableChatData(mChat); data != nil && data.Deactivated && data.MigratedTo != nil {
		return nil, nil, chatpb.ErrChatMigrated
	}
	admin, _ := chatpb.GetImmutableChatParticipant(mChat, invite.AdminId)
	if admin == nil || !chatpb.CanInviteUsers(admin) {
		return nil, nil, chatpb.ErrInviteHashExpired
	}
	return invite, mChat, nil
}

func (c *ChatCore) requireCanInvite(chatID, selfID int64) (*tg.MutableChat, error) {
	mChat, err := c.repo().GetMutableChat(c.ctx, chatID, selfID)
	if err != nil {
		return nil, err
	}
	me, _ := chatpb.GetImmutableChatParticipant(mChat, selfID)
	if me == nil || !chatpb.IsChatMemberStateNormal(me) {
		return nil, chatpb.ErrParticipantInvalid
	}
	if !chatpb.CanInviteUsers(me) {
		return nil, chatpb.ErrChatAdminRequired
	}
	return mChat, nil
}

func (c *ChatCore) hideSingleJoinRequest(selfID, chatID, userID int64, approved bool) (*chatpb.RecentChatInviteRequesters, error) {
	if _, err := c.requireCanInvite(chatID, selfID); err != nil {
		return nil, err
	}
	if approved {
		if _, err := c.addChatUser(c.ctx, addChatUserArg{
			chatID:             chatID,
			inviterID:          selfID,
			userID:             userID,
			approveJoinRequest: true,
			approvedBy:         selfID,
		}); err != nil {
			return nil, err
		}
	} else if err := c.inviteRepository().HideChatJoinRequest(c.ctx, repository.HideJoinRequestsArg{
		ChatID:   chatID,
		UserID:   userID,
		Approver: selfID,
		Approve:  approved,
	}); err != nil {
		return nil, err
	}
	return c.inviteRepository().GetRecentChatInviteRequesters(c.ctx, chatID)
}

func (c *ChatCore) hideAllJoinRequests(selfID, chatID int64, link *string, approved bool) (*chatpb.RecentChatInviteRequesters, error) {
	if _, err := c.requireCanInvite(chatID, selfID); err != nil {
		return nil, err
	}
	requests, err := c.inviteRepository().GetPendingJoinRequests(c.ctx, chatID, link)
	if err != nil {
		return nil, err
	}
	for _, request := range requests {
		if approved {
			if _, err := c.addChatUser(c.ctx, addChatUserArg{
				chatID:             request.ChatID,
				inviterID:          selfID,
				userID:             request.UserID,
				approveJoinRequest: true,
				approvedBy:         selfID,
			}); err != nil {
				return nil, err
			}
			continue
		}
		if err := c.inviteRepository().HideChatJoinRequest(c.ctx, repository.HideJoinRequestsArg{
			ChatID:   request.ChatID,
			UserID:   request.UserID,
			Approver: selfID,
			Approve:  approved,
		}); err != nil {
			return nil, err
		}
	}
	return c.inviteRepository().GetRecentChatInviteRequesters(c.ctx, chatID)
}

func invitePhoto(mChat *tg.MutableChat) tg.PhotoClazz {
	if mChat != nil && mChat.Chat != nil && mChat.Chat.Photo != nil {
		return mChat.Chat.Photo
	}
	return tg.MakeTLPhotoEmpty(&tg.TLPhotoEmpty{})
}

func stringPtr(v string) *string {
	if v == "" {
		return nil
	}
	return &v
}

func boolFromTLBool(v tg.BoolClazz) *bool {
	if v == nil {
		return nil
	}
	_, ok := v.(*tg.TLBoolTrue)
	return &ok
}
