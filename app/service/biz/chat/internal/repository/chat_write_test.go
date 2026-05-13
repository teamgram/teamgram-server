package repository

import (
	"context"
	"errors"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlc"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/internal/repository/model"
)

func TestBackfillBulkInsertIDsAssignsSequentialIDs(t *testing.T) {
	rows := []*model.ChatParticipants{
		{UserId: 1},
		{UserId: 2},
		{UserId: 3},
	}
	if err := backfillBulkInsertIDs(rows, 41, 3); err != nil {
		t.Fatalf("backfillBulkInsertIDs error: %v", err)
	}
	for i, row := range rows {
		want := int64(41 + i)
		if row.Id != want {
			t.Fatalf("row %d id = %d, want %d", i, row.Id, want)
		}
	}
}

func TestBackfillBulkInsertIDsRejectsMissingInsertMetadata(t *testing.T) {
	rows := []*model.ChatParticipants{{UserId: 1}}
	if err := backfillBulkInsertIDs(rows, 0, 1); err == nil {
		t.Fatal("backfillBulkInsertIDs accepted zero last insert id")
	}
	if err := backfillBulkInsertIDs(rows, 10, 0); err == nil {
		t.Fatal("backfillBulkInsertIDs accepted wrong rows affected")
	}
}

func TestCreateChatOperationIDUsesStableCreateChatReplayKey(t *testing.T) {
	if got := CreateChatReplayKey(11, 22); got != "create_chat:11:22" {
		t.Fatalf("replay key = %q", got)
	}
	first := CreateChatOperationID(11, 22)
	second := CreateChatOperationID(11, 22)
	other := CreateChatOperationID(11, 23)
	if first == "" || first != second {
		t.Fatalf("operation id should be stable, got %q and %q", first, second)
	}
	if first == other {
		t.Fatalf("operation id should include client msg id, got same %q", first)
	}
}

func TestCreateChatReplayReturnsSameChat(t *testing.T) {
	repo := openChatIntegrationRepo(t)
	ctx := context.Background()
	creatorID := int64(810000001)
	cleanupCreateChatTestRows(t, repo.db, creatorID)
	t.Cleanup(func() { cleanupCreateChatTestRows(t, repo.db, creatorID) })

	arg := CreateChatArg{
		CreatorID:   creatorID,
		UserIDs:     []int64{creatorID + 1},
		Title:       "replay same chat",
		ClientMsgID: 101,
		OperationID: CreateChatOperationID(creatorID, 101),
	}
	first, err := repo.CreateChat(ctx, arg)
	if err != nil {
		t.Fatalf("first CreateChat error: %v", err)
	}
	second, err := repo.CreateChat(ctx, arg)
	if err != nil {
		t.Fatalf("replay CreateChat error: %v", err)
	}
	if first.Chat.Id == 0 || second.Chat.Id != first.Chat.Id {
		t.Fatalf("replay chat id = %d, want %d", second.Chat.Id, first.Chat.Id)
	}
}

func TestCreateChatReplayDoesNotConsumeFloodWindow(t *testing.T) {
	repo := openChatIntegrationRepo(t)
	ctx := context.Background()
	creatorID := int64(810000002)
	cleanupCreateChatTestRows(t, repo.db, creatorID)
	t.Cleanup(func() { cleanupCreateChatTestRows(t, repo.db, creatorID) })

	arg := CreateChatArg{
		CreatorID:   creatorID,
		Title:       "replay bypasses flood",
		ClientMsgID: 201,
	}
	first, err := repo.CreateChat(ctx, arg)
	if err != nil {
		t.Fatalf("first CreateChat error: %v", err)
	}
	for i := 0; i < 2; i++ {
		replayed, err := repo.CreateChat(ctx, arg)
		if err != nil {
			t.Fatalf("replay %d CreateChat error: %v", i, err)
		}
		if replayed.Chat.Id != first.Chat.Id {
			t.Fatalf("replay %d chat id = %d, want %d", i, replayed.Chat.Id, first.Chat.Id)
		}
	}
}

func TestCreateChatDifferentClientMsgIDWithinWindowReturnsFlood(t *testing.T) {
	repo := openChatIntegrationRepo(t)
	ctx := context.Background()
	creatorID := int64(810000003)
	cleanupCreateChatTestRows(t, repo.db, creatorID)
	t.Cleanup(func() { cleanupCreateChatTestRows(t, repo.db, creatorID) })

	if _, err := repo.CreateChat(ctx, CreateChatArg{
		CreatorID:   creatorID,
		Title:       "first idempotent create",
		ClientMsgID: 301,
	}); err != nil {
		t.Fatalf("first CreateChat error: %v", err)
	}
	_, err := repo.CreateChat(ctx, CreateChatArg{
		CreatorID:   creatorID,
		Title:       "different idempotent create",
		ClientMsgID: 302,
	})
	if !errors.Is(err, chatpb.ErrCreateChatFlood) {
		t.Fatalf("second CreateChat error = %v, want ErrCreateChatFlood", err)
	}
}

func TestCreateChatConcurrentDifferentClientMsgIDWithinWindowCreatesOneChat(t *testing.T) {
	repo := openChatIntegrationRepo(t)
	ctx := context.Background()
	creatorID := int64(810000005)
	cleanupCreateChatTestRows(t, repo.db, creatorID)
	t.Cleanup(func() { cleanupCreateChatTestRows(t, repo.db, creatorID) })

	const attempts = 5
	var (
		wg          sync.WaitGroup
		createdIDs  = make(chan int64, attempts)
		floodErrors = make(chan error, attempts)
		otherErrors = make(chan error, attempts)
	)
	for i := 0; i < attempts; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mChat, err := repo.CreateChat(ctx, CreateChatArg{
				CreatorID:   creatorID,
				Title:       "concurrent create",
				ClientMsgID: int64(500 + i),
			})
			if err != nil {
				if errors.Is(err, chatpb.ErrCreateChatFlood) {
					floodErrors <- err
					return
				}
				otherErrors <- err
				return
			}
			createdIDs <- mChat.Chat.Id
		}(i)
	}
	wg.Wait()
	close(createdIDs)
	close(floodErrors)
	close(otherErrors)

	if len(otherErrors) != 0 {
		t.Fatalf("unexpected create errors: %v", <-otherErrors)
	}
	if len(createdIDs) != 1 {
		t.Fatalf("created chat count = %d, want 1", len(createdIDs))
	}
	if len(floodErrors) != attempts-1 {
		t.Fatalf("flood count = %d, want %d", len(floodErrors), attempts-1)
	}
}

func TestCreateChatActivePendingReplayReturnsTypedWait(t *testing.T) {
	repo := openChatIntegrationRepo(t)
	ctx := context.Background()
	creatorID := int64(810000006)
	clientMsgID := int64(601)
	cleanupCreateChatTestRows(t, repo.db, creatorID)
	t.Cleanup(func() { cleanupCreateChatTestRows(t, repo.db, creatorID) })

	now := time.Now().Unix()
	insertCreateChatOperationTestRow(t, repo.db, &model.ChatCreateOperations{
		OperationId:  CreateChatOperationID(creatorID, clientMsgID),
		ReplayKey:    CreateChatReplayKey(creatorID, clientMsgID),
		ActorUserId:  creatorID,
		ClientMsgId:  clientMsgID,
		Title:        "pending create",
		Status:       CreateChatOperationStatusPending,
		Date:         now,
		UpdatedAtSec: now,
		ExpiresAt:    now + 31,
	})

	_, err := repo.CreateChat(ctx, CreateChatArg{
		CreatorID:   creatorID,
		Title:       "pending create",
		ClientMsgID: clientMsgID,
	})
	if !errors.Is(err, chatpb.ErrCreateChatOperationPending) {
		t.Fatalf("CreateChat error = %v, want ErrCreateChatOperationPending", err)
	}
	var pendingErr *chatpb.CreateChatOperationPendingError
	if !errors.As(err, &pendingErr) {
		t.Fatalf("CreateChat error = %T, want CreateChatOperationPendingError", err)
	}
	if pendingErr.WaitSeconds <= 0 || pendingErr.WaitSeconds > 31 {
		t.Fatalf("pending wait = %d, want 1..31", pendingErr.WaitSeconds)
	}
}

func TestCreateChatExpiredPendingReplayIsReused(t *testing.T) {
	repo := openChatIntegrationRepo(t)
	ctx := context.Background()
	creatorID := int64(810000007)
	clientMsgID := int64(701)
	cleanupCreateChatTestRows(t, repo.db, creatorID)
	t.Cleanup(func() { cleanupCreateChatTestRows(t, repo.db, creatorID) })

	now := time.Now().Unix()
	insertCreateChatOperationTestRow(t, repo.db, &model.ChatCreateOperations{
		OperationId:  CreateChatOperationID(creatorID, clientMsgID),
		ReplayKey:    CreateChatReplayKey(creatorID, clientMsgID),
		ActorUserId:  creatorID,
		ClientMsgId:  clientMsgID,
		Title:        "expired pending create",
		Status:       CreateChatOperationStatusPending,
		Date:         now - 120,
		UpdatedAtSec: now - 120,
		ExpiresAt:    now - 60,
	})

	mChat, err := repo.CreateChat(ctx, CreateChatArg{
		CreatorID:   creatorID,
		Title:       "expired pending create",
		ClientMsgID: clientMsgID,
	})
	if err != nil {
		t.Fatalf("CreateChat error: %v", err)
	}
	if mChat.Chat.Id == 0 {
		t.Fatalf("created chat id = 0")
	}
}

func TestCreateChatFailedReplayIsReused(t *testing.T) {
	repo := openChatIntegrationRepo(t)
	ctx := context.Background()
	creatorID := int64(810000008)
	clientMsgID := int64(801)
	cleanupCreateChatTestRows(t, repo.db, creatorID)
	t.Cleanup(func() { cleanupCreateChatTestRows(t, repo.db, creatorID) })

	now := time.Now().Unix()
	insertCreateChatOperationTestRow(t, repo.db, &model.ChatCreateOperations{
		OperationId:  CreateChatOperationID(creatorID, clientMsgID),
		ReplayKey:    CreateChatReplayKey(creatorID, clientMsgID),
		ActorUserId:  creatorID,
		ClientMsgId:  clientMsgID,
		Title:        "failed create",
		Status:       CreateChatOperationStatusFailed,
		Date:         now - 120,
		UpdatedAtSec: now - 120,
		ExpiresAt:    now - 60,
	})

	mChat, err := repo.CreateChat(ctx, CreateChatArg{
		CreatorID:   creatorID,
		Title:       "failed create",
		ClientMsgID: clientMsgID,
	})
	if err != nil {
		t.Fatalf("CreateChat error: %v", err)
	}
	if mChat.Chat.Id == 0 {
		t.Fatalf("created chat id = 0")
	}
}

func TestCreateChatWithoutClientMsgIDKeepsLegacyCreatePathWithSixtySecondFlood(t *testing.T) {
	repo := openChatIntegrationRepo(t)
	ctx := context.Background()
	creatorID := int64(810000004)
	cleanupCreateChatTestRows(t, repo.db, creatorID)
	t.Cleanup(func() { cleanupCreateChatTestRows(t, repo.db, creatorID) })

	if _, err := repo.CreateChat(ctx, CreateChatArg{
		CreatorID: creatorID,
		Title:     "legacy create",
	}); err != nil {
		t.Fatalf("first CreateChat error: %v", err)
	}
	_, err := repo.CreateChat(ctx, CreateChatArg{
		CreatorID: creatorID,
		Title:     "legacy create flood",
	})
	if !errors.Is(err, chatpb.ErrCreateChatFlood) {
		t.Fatalf("second CreateChat error = %v, want ErrCreateChatFlood", err)
	}
	var floodErr *chatpb.CreateChatFloodError
	if !errors.As(err, &floodErr) || floodErr.WaitSeconds <= 0 || floodErr.WaitSeconds > 60 {
		t.Fatalf("flood error = %#v, want wait in 1..60 seconds", err)
	}
}

func openChatIntegrationRepo(t *testing.T) *Repository {
	t.Helper()
	if testing.Short() {
		t.Skip("integration test skipped in short mode")
	}
	dsn := os.Getenv("TEAMGRAM_TEST_MYSQL_DSN")
	if dsn == "" {
		t.Skip("TEAMGRAM_TEST_MYSQL_DSN is required for chat repository integration tests")
	}
	db, err := sqlx.Open(&sqlx.Config{DSN: dsn})
	if err != nil {
		t.Fatalf("open mysql: %v", err)
	}
	if _, err := db.Exec(context.Background(), "SELECT 1"); err != nil {
		t.Fatalf("ping mysql: %v", err)
	}
	return &Repository{
		CachedConn: sqlc.NewConnWithCache(db, noopCache{}),
		db:         db,
		model:      model.NewModels(db),
	}
}

func cleanupCreateChatTestRows(t *testing.T, db *sqlx.DB, creatorID int64) {
	t.Helper()
	ctx := context.Background()
	if _, err := db.Exec(ctx, "DELETE FROM chat_invites WHERE admin_id = ?", creatorID); err != nil {
		t.Fatalf("cleanup chat_invites: %v", err)
	}
	if _, err := db.Exec(ctx, "DELETE FROM chat_participants WHERE chat_id IN (SELECT id FROM chats WHERE creator_user_id = ?)", creatorID); err != nil {
		t.Fatalf("cleanup chat_participants: %v", err)
	}
	if _, err := db.Exec(ctx, "DELETE FROM chats WHERE creator_user_id = ?", creatorID); err != nil {
		t.Fatalf("cleanup chats: %v", err)
	}
	if _, err := db.Exec(ctx, "DELETE FROM chat_create_operations WHERE actor_user_id = ?", creatorID); err != nil {
		t.Fatalf("cleanup chat_create_operations: %v", err)
	}
}

func insertCreateChatOperationTestRow(t *testing.T, db *sqlx.DB, row *model.ChatCreateOperations) {
	t.Helper()
	if _, err := db.Exec(context.Background(), `
		INSERT INTO chat_create_operations
			(operation_id, replay_key, actor_user_id, client_msg_id, title, invitee_ids, ttl_period, chat_id, participants_version, status, date, updated_at_sec, expires_at)
		VALUES
			(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		row.OperationId,
		row.ReplayKey,
		row.ActorUserId,
		row.ClientMsgId,
		row.Title,
		row.InviteeIds,
		row.TtlPeriod,
		row.ChatId,
		row.ParticipantsVersion,
		row.Status,
		row.Date,
		row.UpdatedAtSec,
		row.ExpiresAt,
	); err != nil {
		t.Fatalf("insert chat_create_operations: %v", err)
	}
}

type noopCache struct{}

func (noopCache) Del(keys ...string) error {
	return nil
}

func (noopCache) DelCtx(ctx context.Context, keys ...string) error {
	return nil
}

func (noopCache) Get(key string, val any) error {
	return sqlc.ErrNotFound
}

func (noopCache) GetCtx(ctx context.Context, key string, val any) error {
	return sqlc.ErrNotFound
}

func (noopCache) IsNotFound(err error) bool {
	return errors.Is(err, sqlc.ErrNotFound)
}

func (noopCache) Set(key string, val any) error {
	return nil
}

func (noopCache) SetCtx(ctx context.Context, key string, val any) error {
	return nil
}

func (noopCache) SetWithExpire(key string, val any, expire time.Duration) error {
	return nil
}

func (noopCache) SetWithExpireCtx(ctx context.Context, key string, val any, expire time.Duration) error {
	return nil
}

func (noopCache) Take(val any, key string, query func(val any) error) error {
	return query(val)
}

func (noopCache) TakeCtx(ctx context.Context, val any, key string, query func(val any) error) error {
	return query(val)
}

func (noopCache) TakeWithExpire(val any, key string, query func(val any, expire time.Duration) error) error {
	return query(val, time.Minute)
}

func (noopCache) TakeWithExpireCtx(ctx context.Context, val any, key string, query func(val any, expire time.Duration) error) error {
	return query(val, time.Minute)
}

func (noopCache) Takes(query func(keys ...string) (map[string]any, error), cacheF func(k, v string) (any, error), keys ...string) error {
	_, err := query(keys...)
	return err
}

func (noopCache) TakesCtx(ctx context.Context, query func(keys ...string) (map[string]any, error), cacheF func(k, v string) (any, error), keys ...string) error {
	_, err := query(keys...)
	return err
}

func (noopCache) TakesWithExpire(query func(expire time.Duration, keys ...string) (map[string]any, error), cacheF func(k, v string) (any, error), keys ...string) error {
	_, err := query(time.Minute, keys...)
	return err
}

func (noopCache) TakesWithExpireCtx(ctx context.Context, query func(expire time.Duration, keys ...string) (map[string]any, error), cacheF func(k, v string) (any, error), keys ...string) error {
	_, err := query(time.Minute, keys...)
	return err
}
