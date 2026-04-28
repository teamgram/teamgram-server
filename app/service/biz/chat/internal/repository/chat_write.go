package repository

import (
	"context"
	"fmt"
	"math"
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
	ParticipantID   int64
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

type UpdateChatAdminArg struct {
	ChatID      int64
	Participant *tg.ImmutableChatParticipant
	IsAdmin     bool
}

type chatAttributeMutation string

const (
	chatAttributeMutationTitle               chatAttributeMutation = "title"
	chatAttributeMutationAbout               chatAttributeMutation = "about"
	chatAttributeMutationPhoto               chatAttributeMutation = "photo"
	chatAttributeMutationDefaultBannedRights chatAttributeMutation = "default_banned_rights"
	chatAttributeMutationNoForwards          chatAttributeMutation = "noforwards"
	chatAttributeMutationTTLPeriod           chatAttributeMutation = "ttl_period"
	chatAttributeMutationAvailableReactions  chatAttributeMutation = "available_reactions"
	chatAttributeMutationAdmin               chatAttributeMutation = "admin"
)

func (op chatAttributeMutation) needsExplicitVersionBump() bool {
	switch op {
	case chatAttributeMutationAbout,
		chatAttributeMutationNoForwards,
		chatAttributeMutationTTLPeriod,
		chatAttributeMutationAvailableReactions,
		chatAttributeMutationAdmin:
		return true
	default:
		return false
	}
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
		id, _, err := r.model.ChatsModel.InsertFullTx(tx, chatRow)
		if err != nil {
			return err
		}
		chatRow.Id = id
		for _, p := range participantRows {
			p.ChatId = id
		}
		lastInsertID, rowsAffected, err := r.model.ChatParticipantsModel.InsertBulkTx(tx, participantRows)
		if err != nil {
			return err
		}
		if err = backfillBulkInsertIDs(participantRows, lastInsertID, rowsAffected); err != nil {
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
	userIDs, err := r.GetChatParticipantIDList(ctx, chatID)
	if err != nil {
		return err
	}
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

func (r *Repository) AddChatUser(ctx context.Context, arg AddChatUserArg) (*tg.ImmutableChatParticipant, error) {
	now := time.Now().Unix()
	participantType := arg.ParticipantType
	if participantType == 0 {
		participantType = chatpb.ChatMemberNormal
	}
	row := &model.ChatParticipants{
		Id:              arg.ParticipantID,
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
		lastInsertID, _, err := r.model.ChatParticipantsModel.InsertOrUpdateTx(tx, row)
		if err != nil {
			return err
		}
		if row.Id == 0 && lastInsertID != 0 {
			row.Id = lastInsertID
		}
		if _, err := r.model.ChatsModel.UpdateParticipantCountTx(tx, arg.Count, arg.ChatID); err != nil {
			return err
		}
		_, err = r.model.ChatInviteParticipantsModel.DeleteTx(tx, arg.ChatID, arg.UserID)
		return err
	}); err != nil {
		return nil, wrapStorage("chat.AddChatUser transaction", err)
	}
	_ = r.CachedConn.DelCache(ctx, chatAggregateAndParticipantCacheKeys(arg.ChatID, []int64{arg.UserID})...)
	return makeImmutableChatParticipant(row), nil
}

func (r *Repository) DeleteChatUser(ctx context.Context, arg DeleteChatUserArg) (int64, error) {
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
		return 0, wrapStorage("chat.DeleteChatUser transaction", err)
	}
	_ = r.CachedConn.DelCache(ctx, chatAggregateAndParticipantCacheKeys(arg.ChatID, []int64{arg.DeleteUserID})...)
	return now, nil
}

func (r *Repository) MigratedToChannel(ctx context.Context, arg MigratedToChannelArg) error {
	userIDs, err := r.GetChatParticipantIDList(ctx, arg.ChatID)
	if err != nil {
		return err
	}
	if err := r.db.Transact(ctx, func(tx *sqlx.Tx) error {
		if _, err := r.model.ChatsModel.UpdateMigratedToTx(tx, arg.ChannelID, arg.AccessHash, arg.ChatID); err != nil {
			return err
		}
		_, err := r.model.ChatParticipantsModel.UpdateStateByChatIdTx(tx, chatpb.ChatMemberStateMigrated, arg.ChatID)
		return err
	}); err != nil {
		return wrapStorage("chat.MigratedToChannel transaction", err)
	}
	_ = r.CachedConn.DelCache(ctx, chatAggregateAndParticipantCacheKeys(arg.ChatID, userIDs)...)
	return nil
}

func (r *Repository) UpdateChatTitle(ctx context.Context, chatID int64, title string) (int64, error) {
	now := time.Now().Unix()
	if err := r.db.Transact(ctx, func(tx *sqlx.Tx) error {
		_, err := r.model.ChatsModel.UpdateTitleTx(tx, title, chatID)
		return err
	}); err != nil {
		return 0, wrapStorage("chat.UpdateChatTitle transaction", err)
	}
	_ = r.CachedConn.DelCache(ctx, chatAggregateCacheKey(chatID))
	return now, nil
}

func (r *Repository) UpdateChatAbout(ctx context.Context, chatID int64, about string) (int64, error) {
	now := time.Now().Unix()
	if err := r.db.Transact(ctx, func(tx *sqlx.Tx) error {
		if _, err := r.model.ChatsModel.UpdateAboutTx(tx, about, chatID); err != nil {
			return err
		}
		if chatAttributeMutationAbout.needsExplicitVersionBump() {
			_, err := r.model.ChatsModel.UpdateVersionTx(tx, chatID)
			return err
		}
		return nil
	}); err != nil {
		return 0, wrapStorage("chat.UpdateChatAbout transaction", err)
	}
	_ = r.CachedConn.DelCache(ctx, chatAggregateCacheKey(chatID))
	return now, nil
}

func (r *Repository) UpdateChatPhoto(ctx context.Context, chatID int64, photoID int64) (int64, error) {
	now := time.Now().Unix()
	if err := r.db.Transact(ctx, func(tx *sqlx.Tx) error {
		_, err := r.model.ChatsModel.UpdatePhotoIdTx(tx, photoID, chatID)
		return err
	}); err != nil {
		return 0, wrapStorage("chat.UpdateChatPhoto transaction", err)
	}
	_ = r.CachedConn.DelCache(ctx, chatAggregateCacheKey(chatID))
	return now, nil
}

func (r *Repository) UpdateChatAdmin(ctx context.Context, arg UpdateChatAdminArg) (*tg.ImmutableChatParticipant, int64, error) {
	if arg.Participant == nil {
		return nil, 0, chatpb.ErrParticipantInvalid
	}
	now := time.Now().Unix()
	row := &model.ChatParticipants{
		Id:              arg.Participant.Id,
		ChatId:          arg.ChatID,
		UserId:          arg.Participant.UserId,
		ParticipantType: chatpb.ChatMemberNormal,
		State:           arg.Participant.State,
		Link:            "",
		InviterUserId:   arg.Participant.InviterUserId,
		InvitedAt:       arg.Participant.InvitedAt,
		KickedAt:        arg.Participant.KickedAt,
		LeftAt:          arg.Participant.LeftAt,
		Date2:           arg.Participant.Date,
		IsBot:           arg.Participant.IsBot,
	}
	if arg.IsAdmin {
		row.ParticipantType = chatpb.ChatMemberAdmin
		row.AdminRights = defaultChatAdminRightsStorage()
		row.Link = arg.Participant.Link
		if row.Link == "" {
			row.Link = chatpb.NormalizeInviteHash(chatpb.BuildInviteLink(chatpb.GenChatInviteHash()))
		}
	}

	if err := r.db.Transact(ctx, func(tx *sqlx.Tx) error {
		if _, err := r.model.ChatParticipantsModel.UpdateAdminRightsTx(tx, row.ParticipantType, row.AdminRights, row.Id); err != nil {
			return err
		}
		if _, err := r.model.ChatParticipantsModel.UpdateLinkTx(tx, row.Link, arg.ChatID, row.UserId); err != nil {
			return err
		}
		if arg.IsAdmin && arg.Participant.Link == "" {
			if _, _, err := r.model.ChatInvitesModel.InsertTx(tx, &model.ChatInvites{
				ChatId:    arg.ChatID,
				AdminId:   row.UserId,
				Link:      row.Link,
				Permanent: true,
				Date2:     now,
			}); err != nil {
				return err
			}
		}
		if chatAttributeMutationAdmin.needsExplicitVersionBump() {
			_, err := r.model.ChatsModel.UpdateVersionTx(tx, arg.ChatID)
			return err
		}
		return nil
	}); err != nil {
		return nil, 0, wrapStorage("chat.UpdateChatAdmin transaction", err)
	}
	_ = r.CachedConn.DelCache(ctx, chatAggregateAndParticipantCacheKeys(arg.ChatID, []int64{row.UserId})...)
	return makeImmutableChatParticipant(row), now, nil
}

func (r *Repository) UpdateChatDefaultBannedRights(ctx context.Context, chatID int64, rights tg.ChatBannedRightsClazz) (int64, error) {
	now := time.Now().Unix()
	if rights != nil && rights.UntilDate == 0 {
		rights.UntilDate = math.MaxInt32
	}
	if err := r.db.Transact(ctx, func(tx *sqlx.Tx) error {
		_, err := r.model.ChatsModel.UpdateDefaultBannedRightsTx(tx, chatBannedRightsToStorage(rights), chatID)
		return err
	}); err != nil {
		return 0, wrapStorage("chat.UpdateChatDefaultBannedRights transaction", err)
	}
	_ = r.CachedConn.DelCache(ctx, chatAggregateCacheKey(chatID))
	return now, nil
}

func (r *Repository) UpdateChatNoForwards(ctx context.Context, chatID int64, noforwards bool) (int64, error) {
	now := time.Now().Unix()
	if err := r.db.Transact(ctx, func(tx *sqlx.Tx) error {
		if _, err := r.model.ChatsModel.UpdateNoforwardsTx(tx, noforwards, chatID); err != nil {
			return err
		}
		if chatAttributeMutationNoForwards.needsExplicitVersionBump() {
			_, err := r.model.ChatsModel.UpdateVersionTx(tx, chatID)
			return err
		}
		return nil
	}); err != nil {
		return 0, wrapStorage("chat.UpdateChatNoForwards transaction", err)
	}
	_ = r.CachedConn.DelCache(ctx, chatAggregateCacheKey(chatID))
	return now, nil
}

func (r *Repository) UpdateChatTTLPeriod(ctx context.Context, chatID int64, ttlPeriod int32) (int64, error) {
	now := time.Now().Unix()
	if err := r.db.Transact(ctx, func(tx *sqlx.Tx) error {
		if _, err := r.model.ChatsModel.UpdateTTLPeriodTx(tx, ttlPeriod, chatID); err != nil {
			return err
		}
		if chatAttributeMutationTTLPeriod.needsExplicitVersionBump() {
			_, err := r.model.ChatsModel.UpdateVersionTx(tx, chatID)
			return err
		}
		return nil
	}); err != nil {
		return 0, wrapStorage("chat.UpdateChatTTLPeriod transaction", err)
	}
	_ = r.CachedConn.DelCache(ctx, chatAggregateCacheKey(chatID))
	return now, nil
}

func (r *Repository) UpdateChatAvailableReactions(ctx context.Context, chatID int64, kind int32, reactions []string) (int64, error) {
	now := time.Now().Unix()
	payload := availableReactionsToStorage(reactions)
	if err := r.db.Transact(ctx, func(tx *sqlx.Tx) error {
		if _, err := r.model.ChatsModel.UpdateAvailableReactionsTx(tx, kind, payload, chatID); err != nil {
			return err
		}
		if chatAttributeMutationAvailableReactions.needsExplicitVersionBump() {
			_, err := r.model.ChatsModel.UpdateVersionTx(tx, chatID)
			return err
		}
		return nil
	}); err != nil {
		return 0, wrapStorage("chat.UpdateChatAvailableReactions transaction", err)
	}
	_ = r.CachedConn.DelCache(ctx, chatAggregateCacheKey(chatID))
	return now, nil
}

func defaultChatAdminRightsStorage() int32 {
	return chatAdminRightsToStorage(tg.MakeTLChatAdminRights(&tg.TLChatAdminRights{
		ChangeInfo:     true,
		DeleteMessages: true,
		BanUsers:       true,
		InviteUsers:    true,
		PinMessages:    true,
		AddAdmins:      true,
		ManageCall:     true,
		Other:          true,
	}).ToChatAdminRights())
}

func backfillBulkInsertIDs(rows []*model.ChatParticipants, lastInsertID, rowsAffected int64) error {
	if len(rows) == 0 {
		return nil
	}
	if lastInsertID <= 0 {
		return fmt.Errorf("chat_participants.InsertBulkTx last insert id %d", lastInsertID)
	}
	if rowsAffected != int64(len(rows)) {
		return fmt.Errorf("chat_participants.InsertBulkTx rows affected %d, want %d", rowsAffected, len(rows))
	}
	for i, row := range rows {
		if row != nil {
			row.Id = lastInsertID + int64(i)
		}
	}
	return nil
}
