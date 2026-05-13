package repository

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

type CreateChatArg struct {
	CreatorID   int64
	UserIDs     []int64
	Title       string
	BotIDs      []int64
	TTLPeriod   int32
	ClientMsgID int64
	OperationID string
}

type AddChatUserArg struct {
	ChatID                  int64
	InviterID               int64
	UserID                  int64
	ParticipantID           int64
	ParticipantType         int32
	IsBot                   bool
	Count                   int32
	RecordInviteParticipant bool
	InviteLink              string
	InviteRequested         bool
	ApproveJoinRequest      bool
	ApprovedBy              int64
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
		chatAttributeMutationTTLPeriod,
		chatAttributeMutationAvailableReactions,
		chatAttributeMutationAdmin:
		return true
	default:
		return false
	}
}

func CreateChatReplayKey(actorUserID, clientMsgID int64) string {
	if clientMsgID == 0 {
		return ""
	}
	return fmt.Sprintf("create_chat:%d:%d", actorUserID, clientMsgID)
}

func CreateChatOperationID(actorUserID, clientMsgID int64) string {
	sum := sha256.Sum256([]byte(CreateChatReplayKey(actorUserID, clientMsgID)))
	return hex.EncodeToString(sum[:])
}

func createChatActorLockOperationID(actorUserID int64) string {
	sum := sha256.Sum256([]byte(fmt.Sprintf("create_chat_actor_lock:%d", actorUserID)))
	return hex.EncodeToString(sum[:])
}

func createChatActorLockReplayKey(actorUserID int64) string {
	return fmt.Sprintf("create_chat_actor_lock:%d", actorUserID)
}

func (r *Repository) CreateChat(ctx context.Context, arg CreateChatArg) (*tg.MutableChat, error) {
	now := time.Now().Unix()
	replayKey := CreateChatReplayKey(arg.CreatorID, arg.ClientMsgID)
	operationID := arg.OperationID
	if replayKey != "" {
		if operationID == "" {
			operationID = CreateChatOperationID(arg.CreatorID, arg.ClientMsgID)
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

	if err := r.ensureCreateChatActorLockRow(ctx, arg.CreatorID, now); err != nil {
		return nil, wrapStorage("chat_create_operations.EnsureActorLock", err)
	}

	var replayChatID int64
	if err := r.db.Transact(ctx, func(tx *sqlx.Tx) error {
		txModel := r.model.WithTx(tx)
		if err := lockCreateChatActorTx(txModel, arg.CreatorID); err != nil {
			return err
		}
		if replayKey != "" {
			existing, err := txModel.ChatCreateOperationsModel.SelectByReplayKeyForUpdate(replayKey)
			if err != nil && !isNotFound(err) {
				return err
			}
			if existing != nil {
				switch existing.Status {
				case CreateChatOperationStatusChatCreated, CreateChatOperationStatusCompleted:
					if existing.ChatId != 0 {
						replayChatID = existing.ChatId
						return nil
					}
				case CreateChatOperationStatusPending:
					if existing.ExpiresAt > now {
						return chatpb.NewCreateChatOperationPendingError(waitSecondsUntil(now, existing.ExpiresAt))
					}
				case CreateChatOperationStatusFailed:
				}
				if err := resetCreateChatOperationForRetry(txModel, existing, arg, operationID, now, chatRow.Version); err != nil {
					return err
				}
			} else {
				if err := enforceCreateChatOperationFloodTx(txModel, arg.CreatorID, now); err != nil {
					return err
				}
				_, _, err = txModel.ChatCreateOperationsModel.Insert(&model.ChatCreateOperations{
					OperationId:         operationID,
					ReplayKey:           replayKey,
					ActorUserId:         arg.CreatorID,
					ClientMsgId:         arg.ClientMsgID,
					Title:               arg.Title,
					InviteeIds:          encodeCreateChatInviteeIDs(arg.UserIDs, arg.BotIDs),
					TtlPeriod:           arg.TTLPeriod,
					Status:              CreateChatOperationStatusPending,
					Date:                now,
					UpdatedAtSec:        now,
					ExpiresAt:           now + int64(createChatFloodInterval.Seconds()),
					ParticipantsVersion: chatRow.Version,
				})
				if err != nil {
					return err
				}
			}
		} else if err := enforceLegacyCreateChatFloodTx(txModel, arg.CreatorID, now); err != nil {
			return err
		}
		id, _, err := r.model.ChatsModel.InsertFullTx(tx, chatRow)
		if err != nil {
			return err
		}
		chatRow.Id = id
		for _, p := range participantRows {
			p.ChatId = id
		}
		lastInsertID, rowsAffected, err := txModel.ChatParticipantsModel.InsertBulk(participantRows)
		if err != nil {
			return err
		}
		if err = backfillBulkInsertIDs(participantRows, lastInsertID, rowsAffected); err != nil {
			return err
		}
		_, _, err = txModel.ChatInvitesModel.Insert(&model.ChatInvites{
			ChatId:    id,
			AdminId:   arg.CreatorID,
			Link:      creatorHash,
			Permanent: true,
			Date2:     now,
		})
		if err != nil {
			return err
		}
		if replayKey != "" {
			_, err = txModel.ChatCreateOperationsModel.MarkChatCreated(
				chatRow.Id,
				chatRow.Version,
				CreateChatOperationStatusChatCreated,
				now,
				operationID,
			)
			return err
		}
		return nil
	}); err != nil {
		if errors.Is(err, chatpb.ErrCreateChatFlood) || errors.Is(err, chatpb.ErrCreateChatOperationPending) {
			return nil, err
		}
		return nil, wrapStorage("chat.CreateChat transaction", err)
	}
	if replayChatID != 0 {
		return r.GetMutableChat(ctx, replayChatID)
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

func resetCreateChatOperationForRetry(txModel *model.TxModels, existing *model.ChatCreateOperations, arg CreateChatArg, operationID string, now int64, participantsVersion int32) error {
	if existing == nil {
		return nil
	}
	_, err := txModel.ChatCreateOperationsModel.ResetForRetry(
		operationID,
		arg.CreatorID,
		arg.ClientMsgID,
		arg.Title,
		encodeCreateChatInviteeIDs(arg.UserIDs, arg.BotIDs),
		arg.TTLPeriod,
		participantsVersion,
		CreateChatOperationStatusPending,
		now,
		now,
		now+int64(createChatFloodInterval.Seconds()),
		existing.ReplayKey,
	)
	return err
}

func (r *Repository) ensureCreateChatActorLockRow(ctx context.Context, actorUserID, now int64) error {
	_, _, err := r.model.ChatCreateOperationsModel.EnsureActorLock(ctx, &model.ChatCreateOperations{
		OperationId:  createChatActorLockOperationID(actorUserID),
		ReplayKey:    createChatActorLockReplayKey(actorUserID),
		ActorUserId:  actorUserID,
		Status:       0,
		Date:         now,
		UpdatedAtSec: now,
		ExpiresAt:    0,
	})
	if err != nil {
		return err
	}
	return nil
}

func lockCreateChatActorTx(txModel *model.TxModels, actorUserID int64) error {
	_, err := txModel.ChatCreateOperationsModel.SelectActorLockForUpdate(createChatActorLockOperationID(actorUserID))
	return err
}

func waitSecondsUntil(now, expiresAt int64) int32 {
	if expiresAt <= now {
		return 0
	}
	remaining := expiresAt - now
	if remaining > math.MaxInt32 {
		return math.MaxInt32
	}
	return int32(remaining)
}

func enforceCreateChatOperationFloodTx(txModel *model.TxModels, actorUserID int64, now int64) error {
	last, err := selectLastCreateChatOperationTx(txModel, actorUserID, CreateChatOperationStatusChatCreated, CreateChatOperationStatusCompleted)
	if err != nil {
		return err
	}
	if last == nil {
		return nil
	}
	elapsed := time.Duration(now-last.UpdatedAtSec) * time.Second
	if elapsed < createChatFloodInterval {
		return chatpb.NewCreateChatFloodError(int32((createChatFloodInterval - elapsed).Seconds()))
	}
	return nil
}

func selectLastCreateChatOperationTx(txModel *model.TxModels, actorUserID int64, statuses ...int32) (*model.ChatCreateOperations, error) {
	var last *model.ChatCreateOperations
	for _, status := range statuses {
		row, err := txModel.ChatCreateOperationsModel.SelectLastCompletedByActor(actorUserID, status)
		if err != nil {
			if isNotFound(err) {
				continue
			}
			return nil, err
		}
		if last == nil || row.UpdatedAtSec > last.UpdatedAtSec {
			last = row
		}
	}
	return last, nil
}

func enforceLegacyCreateChatFloodTx(txModel *model.TxModels, creatorID int64, now int64) error {
	last, err := txModel.ChatsModel.SelectLastCreator(creatorID)
	if err != nil {
		if isNotFound(err) {
			return nil
		}
		return err
	}
	elapsed := time.Duration(now-last.Date) * time.Second
	if elapsed < createChatFloodInterval {
		return chatpb.NewCreateChatFloodError(int32((createChatFloodInterval - elapsed).Seconds()))
	}
	return nil
}

func encodeCreateChatInviteeIDs(userIDs, botIDs []int64) string {
	if len(userIDs) == 0 && len(botIDs) == 0 {
		return ""
	}
	ids := make([]string, 0, len(userIDs)+len(botIDs))
	for _, id := range userIDs {
		ids = append(ids, strconv.FormatInt(id, 10))
	}
	for _, id := range botIDs {
		ids = append(ids, strconv.FormatInt(id, 10))
	}
	return strings.Join(ids, ",")
}

func (r *Repository) DeleteChat(ctx context.Context, chatID int64) error {
	userIDs, err := r.GetChatParticipantIDList(ctx, chatID)
	if err != nil {
		return err
	}
	if err := r.db.Transact(ctx, func(tx *sqlx.Tx) error {
		txModel := r.model.WithTx(tx)
		if _, err := txModel.ChatParticipantsModel.UpdateStateByChatId(chatpb.ChatMemberStateKicked, chatID); err != nil {
			return err
		}
		if _, err := txModel.ChatsModel.UpdateParticipantCount(0, chatID); err != nil {
			return err
		}
		_, err := txModel.ChatsModel.UpdateDeactivated(true, chatID)
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
		txModel := r.model.WithTx(tx)
		lastInsertID, _, err := txModel.ChatParticipantsModel.InsertOrUpdate(row)
		if err != nil {
			return err
		}
		if row.Id == 0 && lastInsertID != 0 {
			row.Id = lastInsertID
		}
		if _, err := txModel.ChatsModel.UpdateParticipantCount(arg.Count, arg.ChatID); err != nil {
			return err
		}
		switch {
		case arg.ApproveJoinRequest:
			rowsAffected, err := txModel.ChatInviteParticipantsModel.UpdateApprovedBy(arg.ApprovedBy, arg.ChatID, arg.UserID)
			if err != nil {
				return err
			}
			return requireRowsAffected(rowsAffected)
		case arg.RecordInviteParticipant:
			_, _, err = txModel.ChatInviteParticipantsModel.Insert(&model.ChatInviteParticipants{
				ChatId:    arg.ChatID,
				Link:      chatpb.NormalizeInviteHash(arg.InviteLink),
				UserId:    arg.UserID,
				Requested: arg.InviteRequested,
				Date2:     now,
			})
			return err
		default:
			_, err = txModel.ChatInviteParticipantsModel.Delete(arg.ChatID, arg.UserID)
			return err
		}
	}); err != nil {
		return nil, wrapMutationError("chat.AddChatUser transaction", err)
	}
	_ = r.CachedConn.DelCache(ctx, chatAggregateAndParticipantCacheKeys(arg.ChatID, []int64{arg.UserID})...)
	return makeImmutableChatParticipant(row), nil
}

func (r *Repository) DeleteChatUser(ctx context.Context, arg DeleteChatUserArg) (int64, error) {
	now := time.Now().Unix()
	if err := r.db.Transact(ctx, func(tx *sqlx.Tx) error {
		txModel := r.model.WithTx(tx)
		var err error
		if arg.Kicked {
			_, err = txModel.ChatParticipantsModel.UpdateKicked(now, arg.ParticipantID)
		} else {
			_, err = txModel.ChatParticipantsModel.UpdateLeft(now, arg.ParticipantID)
		}
		if err != nil {
			return err
		}
		if _, err = txModel.ChatsModel.UpdateParticipantCount(arg.Count, arg.ChatID); err != nil {
			return err
		}
		_, err = txModel.ChatInviteParticipantsModel.Delete(arg.ChatID, arg.DeleteUserID)
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
		txModel := r.model.WithTx(tx)
		if _, err := txModel.ChatsModel.UpdateMigratedTo(arg.ChannelID, arg.AccessHash, arg.ChatID); err != nil {
			return err
		}
		_, err := txModel.ChatParticipantsModel.UpdateStateByChatId(chatpb.ChatMemberStateMigrated, arg.ChatID)
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
		txModel := r.model.WithTx(tx)
		_, err := txModel.ChatsModel.UpdateTitle(title, chatID)
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
		txModel := r.model.WithTx(tx)
		if _, err := txModel.ChatsModel.UpdateAbout(about, chatID); err != nil {
			return err
		}
		if chatAttributeMutationAbout.needsExplicitVersionBump() {
			_, err := txModel.ChatsModel.UpdateVersion(chatID)
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
		txModel := r.model.WithTx(tx)
		_, err := txModel.ChatsModel.UpdatePhotoId(photoID, chatID)
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
	updated, adminRights, link := applyChatAdminMutation(arg.Participant, arg.IsAdmin)

	if err := r.db.Transact(ctx, func(tx *sqlx.Tx) error {
		txModel := r.model.WithTx(tx)
		if _, err := r.model.ChatParticipantsModel.UpdateAdminRightsTx(tx, updated.ParticipantType, adminRights, updated.Id); err != nil {
			return err
		}
		if _, err := txModel.ChatParticipantsModel.UpdateLink(link, arg.ChatID, updated.UserId); err != nil {
			return err
		}
		if arg.IsAdmin && arg.Participant.Link == "" {
			if _, _, err := txModel.ChatInvitesModel.Insert(&model.ChatInvites{
				ChatId:    arg.ChatID,
				AdminId:   updated.UserId,
				Link:      link,
				Permanent: true,
				Date2:     now,
			}); err != nil {
				return err
			}
		}
		if chatAttributeMutationAdmin.needsExplicitVersionBump() {
			_, err := txModel.ChatsModel.UpdateVersion(arg.ChatID)
			return err
		}
		return nil
	}); err != nil {
		return nil, 0, wrapStorage("chat.UpdateChatAdmin transaction", err)
	}
	_ = r.CachedConn.DelCache(ctx, chatAggregateAndParticipantCacheKeys(arg.ChatID, []int64{updated.UserId})...)
	return updated, now, nil
}

func (r *Repository) UpdateChatDefaultBannedRights(ctx context.Context, chatID int64, rights tg.ChatBannedRightsClazz) (int64, error) {
	now := time.Now().Unix()
	if rights != nil && rights.UntilDate == 0 {
		rights.UntilDate = math.MaxInt32
	}
	if err := r.db.Transact(ctx, func(tx *sqlx.Tx) error {
		txModel := r.model.WithTx(tx)
		_, err := txModel.ChatsModel.UpdateDefaultBannedRights(chatBannedRightsToStorage(rights), chatID)
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
		txModel := r.model.WithTx(tx)
		if _, err := txModel.ChatsModel.UpdateNoforwards(noforwards, chatID); err != nil {
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
		txModel := r.model.WithTx(tx)
		if _, err := txModel.ChatsModel.UpdateTTLPeriod(ttlPeriod, chatID); err != nil {
			return err
		}
		if chatAttributeMutationTTLPeriod.needsExplicitVersionBump() {
			_, err := txModel.ChatsModel.UpdateVersion(chatID)
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
		txModel := r.model.WithTx(tx)
		if _, err := txModel.ChatsModel.UpdateAvailableReactions(kind, payload, chatID); err != nil {
			return err
		}
		if chatAttributeMutationAvailableReactions.needsExplicitVersionBump() {
			_, err := txModel.ChatsModel.UpdateVersion(chatID)
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

func defaultChatAdminRights() tg.ChatAdminRightsClazz {
	return tg.MakeTLChatAdminRights(&tg.TLChatAdminRights{
		ChangeInfo:     true,
		DeleteMessages: true,
		BanUsers:       true,
		InviteUsers:    true,
		PinMessages:    true,
		AddAdmins:      true,
		ManageCall:     true,
		Other:          true,
	}).ToChatAdminRights()
}

func applyChatAdminMutation(participant *tg.ImmutableChatParticipant, isAdmin bool) (*tg.ImmutableChatParticipant, int32, string) {
	if participant == nil {
		return nil, 0, ""
	}
	updated := *participant
	if !isAdmin {
		updated.ParticipantType = chatpb.ChatMemberNormal
		updated.AdminRights = nil
		updated.Link = ""
		return &updated, 0, ""
	}
	updated.ParticipantType = chatpb.ChatMemberAdmin
	updated.AdminRights = defaultChatAdminRights()
	if updated.Link == "" {
		updated.Link = chatpb.NormalizeInviteHash(chatpb.BuildInviteLink(chatpb.GenChatInviteHash()))
	}
	return &updated, chatAdminRightsToStorage(updated.AdminRights), updated.Link
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
