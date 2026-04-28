package repository

import (
	"context"
	"math/rand"
	"time"

	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

type CreateChatArg struct {
	CreatorID int64
	UserIDs   []int64
	Title     string
	BotIDs    []int64
	TTLPeriod int32
}

type AddChatUserArg struct {
	ChatID          int64
	InviterID       int64
	UserID          int64
	ParticipantType int32
	IsBot           bool
	Count           int32
}

type DeleteChatUserArg struct {
	ChatID        int64
	DeleteUserID  int64
	ParticipantID int64
	Kicked        bool
	Count         int32
}

type MigratedToChannelArg struct {
	ChatID     int64
	ChannelID  int64
	AccessHash int64
}

func (r *Repository) CreateChat(ctx context.Context, arg CreateChatArg) (*tg.MutableChat, error) {
	now := time.Now().Unix()
	last, err := r.model.ChatsModel.SelectLastCreator(ctx, arg.CreatorID)
	if err != nil && !isNotFound(err) {
		return nil, wrapStorage("chats.SelectLastCreator", err)
	}
	if last != nil {
		elapsed := time.Duration(now-last.Date) * time.Second
		if elapsed < createChatFloodInterval {
			return nil, chatpb.NewCreateChatFloodError(int32((createChatFloodInterval - elapsed).Seconds()))
		}
	}

	chatRow := &model.Chats{
		CreatorUserId:          arg.CreatorID,
		AccessHash:             rand.Int63(),
		ParticipantCount:       int32(1 + len(arg.UserIDs) + len(arg.BotIDs)),
		Title:                  arg.Title,
		DefaultBannedRights:    chatBannedRightsToStorage(tg.MakeTLChatBannedRights(&tg.TLChatBannedRights{}).ToChatBannedRights()),
		AvailableReactionsType: 0,
		AvailableReactions:     "",
		TtlPeriod:              arg.TTLPeriod,
		Version:                1,
		Date:                   now,
	}

	creatorHash := chatpb.NormalizeInviteHash(chatpb.BuildInviteLink(chatpb.GenChatInviteHash()))
	participantRows := make([]*model.ChatParticipants, 0, int(chatRow.ParticipantCount))
	participantRows = append(participantRows, &model.ChatParticipants{
		UserId:          arg.CreatorID,
		ParticipantType: chatpb.ChatMemberCreator,
		Link:            creatorHash,
		InvitedAt:       now,
		Date2:           now,
		State:           chatpb.ChatMemberStateNormal,
	})
	for _, userID := range arg.UserIDs {
		participantRows = append(participantRows, &model.ChatParticipants{
			UserId:          userID,
			ParticipantType: chatpb.ChatMemberNormal,
			InviterUserId:   arg.CreatorID,
			InvitedAt:       now,
			Date2:           now,
			State:           chatpb.ChatMemberStateNormal,
		})
	}
	for _, botID := range arg.BotIDs {
		participantRows = append(participantRows, &model.ChatParticipants{
			UserId:          botID,
			ParticipantType: chatpb.ChatMemberNormal,
			InviterUserId:   arg.CreatorID,
			InvitedAt:       now,
			Date2:           now,
			IsBot:           true,
			State:           chatpb.ChatMemberStateNormal,
		})
	}

	if err := r.db.Transact(ctx, func(tx *sqlx.Tx) error {
		id, _, err := r.model.ChatsModel.InsertTx(tx, chatRow)
		if err != nil {
			return err
		}
		chatRow.Id = id
		for _, p := range participantRows {
			p.ChatId = id
		}
		if _, _, err = r.model.ChatParticipantsModel.InsertBulkTx(tx, participantRows); err != nil {
			return err
		}
		_, _, err = r.model.ChatInvitesModel.InsertTx(tx, &model.ChatInvites{
			ChatId:    id,
			AdminId:   arg.CreatorID,
			Link:      creatorHash,
			Permanent: true,
			Date2:     now,
		})
		return err
	}); err != nil {
		return nil, wrapStorage("chat.CreateChat transaction", err)
	}

	outRows := make([]model.ChatParticipants, 0, len(participantRows))
	userIDs := make([]int64, 0, len(participantRows))
	for _, p := range participantRows {
		outRows = append(outRows, *p)
		userIDs = append(userIDs, p.UserId)
	}
	_ = r.CachedConn.DelCache(ctx, chatAggregateAndParticipantCacheKeys(chatRow.Id, userIDs)...)
	return r.makeMutableChatFromRows(ctx, chatRow, outRows), nil
}

func (r *Repository) DeleteChat(ctx context.Context, chatID int64) error {
	userIDs, _ := r.GetChatParticipantIDList(ctx, chatID)
	if err := r.db.Transact(ctx, func(tx *sqlx.Tx) error {
		if _, err := r.model.ChatParticipantsModel.UpdateStateByChatIdTx(tx, chatpb.ChatMemberStateKicked, chatID); err != nil {
			return err
		}
		if _, err := r.model.ChatsModel.UpdateParticipantCountTx(tx, 0, chatID); err != nil {
			return err
		}
		_, err := r.model.ChatsModel.UpdateDeactivatedTx(tx, true, chatID)
		return err
	}); err != nil {
		return wrapStorage("chat.DeleteChat transaction", err)
	}
	_ = r.CachedConn.DelCache(ctx, chatAggregateAndParticipantCacheKeys(chatID, userIDs)...)
	return nil
}

func (r *Repository) AddChatUser(ctx context.Context, arg AddChatUserArg) (*tg.MutableChat, error) {
	now := time.Now().Unix()
	participantType := arg.ParticipantType
	if participantType == 0 {
		participantType = chatpb.ChatMemberNormal
	}
	row := &model.ChatParticipants{
		ChatId:          arg.ChatID,
		UserId:          arg.UserID,
		ParticipantType: participantType,
		InviterUserId:   arg.InviterID,
		InvitedAt:       now,
		Date2:           now,
		IsBot:           arg.IsBot,
		State:           chatpb.ChatMemberStateNormal,
	}

	if err := r.db.Transact(ctx, func(tx *sqlx.Tx) error {
		if _, _, err := r.model.ChatParticipantsModel.InsertOrUpdateTx(tx, row); err != nil {
			return err
		}
		if _, err := r.model.ChatsModel.UpdateParticipantCountTx(tx, arg.Count, arg.ChatID); err != nil {
			return err
		}
		_, err := r.model.ChatInviteParticipantsModel.DeleteTx(tx, arg.ChatID, arg.UserID)
		return err
	}); err != nil {
		return nil, wrapStorage("chat.AddChatUser transaction", err)
	}
	_ = r.CachedConn.DelCache(ctx, chatAggregateAndParticipantCacheKeys(arg.ChatID, []int64{arg.UserID})...)
	return r.GetMutableChat(ctx, arg.ChatID, arg.InviterID, arg.UserID)
}

func (r *Repository) DeleteChatUser(ctx context.Context, arg DeleteChatUserArg) (*tg.MutableChat, error) {
	now := time.Now().Unix()
	if err := r.db.Transact(ctx, func(tx *sqlx.Tx) error {
		var err error
		if arg.Kicked {
			_, err = r.model.ChatParticipantsModel.UpdateKickedTx(tx, now, arg.ParticipantID)
		} else {
			_, err = r.model.ChatParticipantsModel.UpdateLeftTx(tx, now, arg.ParticipantID)
		}
		if err != nil {
			return err
		}
		if _, err = r.model.ChatsModel.UpdateParticipantCountTx(tx, arg.Count, arg.ChatID); err != nil {
			return err
		}
		_, err = r.model.ChatInviteParticipantsModel.DeleteTx(tx, arg.ChatID, arg.DeleteUserID)
		return err
	}); err != nil {
		return nil, wrapStorage("chat.DeleteChatUser transaction", err)
	}
	_ = r.CachedConn.DelCache(ctx, chatAggregateAndParticipantCacheKeys(arg.ChatID, []int64{arg.DeleteUserID})...)
	return r.GetMutableChat(ctx, arg.ChatID, arg.DeleteUserID)
}

func (r *Repository) MigratedToChannel(ctx context.Context, arg MigratedToChannelArg) (*tg.MutableChat, error) {
	userIDs, _ := r.GetChatParticipantIDList(ctx, arg.ChatID)
	if err := r.db.Transact(ctx, func(tx *sqlx.Tx) error {
		if _, err := r.model.ChatsModel.UpdateMigratedToTx(tx, arg.ChannelID, arg.AccessHash, arg.ChatID); err != nil {
			return err
		}
		_, err := r.model.ChatParticipantsModel.UpdateStateByChatIdTx(tx, chatpb.ChatMemberStateMigrated, arg.ChatID)
		return err
	}); err != nil {
		return nil, wrapStorage("chat.MigratedToChannel transaction", err)
	}
	_ = r.CachedConn.DelCache(ctx, chatAggregateAndParticipantCacheKeys(arg.ChatID, userIDs)...)
	return r.GetMutableChat(ctx, arg.ChatID)
}
