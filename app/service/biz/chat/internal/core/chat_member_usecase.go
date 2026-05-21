package core

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/internal/repository"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const maxChatParticipants = 200

type addChatUserArg struct {
	chatID                  int64
	inviterID               int64
	userID                  int64
	isBot                   bool
	recordInviteParticipant bool
	inviteLink              string
	inviteRequested         bool
	approveJoinRequest      bool
	approvedBy              int64
}

type deleteChatUserArg struct {
	chatID       int64
	operatorID   int64
	deleteUserID int64
}

func (c *ChatCore) addChatUser(ctx context.Context, arg addChatUserArg) (*tg.MutableChat, error) {
	mChat, err := c.repo().GetMutableChat(ctx, arg.chatID)
	if err != nil {
		return nil, err
	}
	if data := chat.MutableChatData(mChat); data != nil && data.Deactivated && data.MigratedTo != nil {
		return nil, chat.ErrChatMigrated
	}

	inviterID := arg.inviterID
	var inviter *tg.ImmutableChatParticipant
	if inviterID == 0 {
		inviterID = chat.ChatCreator(mChat)
	} else {
		inviter, _ = chat.GetImmutableChatParticipant(mChat, inviterID)
		if inviter == nil || (!chat.IsChatMemberStateNormal(inviter) && !chat.IsChatMemberCreator(inviter)) {
			return nil, chat.ErrInputUserDeactivated
		}
	}

	willAdd, _ := chat.GetImmutableChatParticipant(mChat, arg.userID)
	if chat.IsChatMemberStateNormal(willAdd) {
		return nil, chat.ErrUserAlreadyParticipant
	}
	if chat.ChatParticipantsCount(mChat) >= maxChatParticipants {
		return nil, chat.ErrUsersTooFew
	}
	if inviter != nil && !chat.CanInviteUsers(inviter) && !chatDefaultAllowsInvite(mChat) {
		return nil, chat.ErrChatAdminRequired
	}

	participantType := chat.ChatMemberNormal
	if arg.userID == chat.ChatCreator(mChat) {
		participantType = chat.ChatMemberCreator
	}
	added, err := c.writeRepository().AddChatUser(ctx, repository.AddChatUserArg{
		ChatID:                  arg.chatID,
		InviterID:               inviterID,
		UserID:                  arg.userID,
		ParticipantID:           chatParticipantID(willAdd),
		ParticipantType:         participantType,
		IsBot:                   arg.isBot,
		Count:                   chat.ChatParticipantsCount(mChat) + 1,
		RecordInviteParticipant: arg.recordInviteParticipant,
		InviteLink:              arg.inviteLink,
		InviteRequested:         arg.inviteRequested,
		ApproveJoinRequest:      arg.approveJoinRequest,
		ApprovedBy:              arg.approvedBy,
	})
	if err != nil {
		return nil, err
	}
	if added != nil {
		upsertMutableChatParticipant(mChat, added)
	}
	if mChat.Chat != nil {
		mChat.Chat.ParticipantsCount++
		mChat.Chat.Version++
	}
	return mChat, nil
}

func (c *ChatCore) deleteChatUser(ctx context.Context, arg deleteChatUserArg) (*tg.MutableChat, error) {
	mChat, err := c.repo().GetMutableChat(ctx, arg.chatID)
	if err != nil {
		return nil, err
	}

	operatorID := arg.operatorID
	if operatorID == 0 {
		operatorID = chat.ChatCreator(mChat)
	}
	me, _ := chat.GetImmutableChatParticipant(mChat, operatorID)
	if me == nil {
		return nil, chat.ErrInputUserDeactivated
	}

	kicked := operatorID != arg.deleteUserID
	var deletedUser *tg.ImmutableChatParticipant
	if kicked {
		if !chat.IsChatMemberStateNormal(me) {
			return nil, chat.ErrParticipantInvalid
		}
		if !chat.CanAdminBanUsers(me) {
			return nil, chat.ErrChatAdminRequired
		}
		deletedUser, _ = chat.GetImmutableChatParticipant(mChat, arg.deleteUserID)
		if deletedUser == nil {
			return nil, chat.ErrUserNotParticipant
		}
		if !chat.IsChatMemberStateNormal(deletedUser) {
			return nil, chat.ErrParticipantInvalid
		}
		if chat.IsChatMemberAdmin(deletedUser) {
			return nil, chat.ErrChatAdminRequired
		}
	} else {
		deletedUser = me
		if !chat.IsChatMemberStateNormal(me) {
			return nil, chat.ErrParticipantInvalid
		}
	}
	if chat.IsChatMemberCreator(deletedUser) || deletedUser.UserId == chat.ChatCreator(mChat) {
		return nil, chat.ErrChatAdminRequired
	}

	count := chat.ChatParticipantsCount(mChat) - 1
	if count < 0 {
		count = 0
	}
	at, err := c.writeRepository().DeleteChatUser(ctx, repository.DeleteChatUserArg{
		ChatID:        arg.chatID,
		DeleteUserID:  arg.deleteUserID,
		ParticipantID: deletedUser.Id,
		Kicked:        kicked,
		Count:         count,
	})
	if err != nil {
		return nil, err
	}
	if kicked {
		deletedUser.State = chat.ChatMemberStateKicked
		deletedUser.KickedAt = at
	} else {
		deletedUser.State = chat.ChatMemberStateLeft
		deletedUser.LeftAt = at
	}
	deletedUser.Date = at
	if mChat.Chat != nil {
		mChat.Chat.ParticipantsCount = count
		mChat.Chat.Version++
	}
	return mChat, nil
}

func chatDefaultAllowsInvite(mChat *tg.MutableChat) bool {
	rights := chat.ChatDefaultBannedRights(mChat)
	return rights == nil || !rights.InviteUsers
}

func chatParticipantID(p *tg.ImmutableChatParticipant) int64 {
	if p == nil {
		return 0
	}
	return p.Id
}

func upsertMutableChatParticipant(mChat *tg.MutableChat, participant *tg.ImmutableChatParticipant) {
	if mChat == nil || participant == nil {
		return
	}
	for i, existing := range mChat.ChatParticipants {
		if existing != nil && existing.UserId == participant.UserId {
			mChat.ChatParticipants[i] = participant
			return
		}
	}
	mChat.ChatParticipants = append(mChat.ChatParticipants, participant)
}
