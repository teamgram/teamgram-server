//go:build integration

package repository

import (
	"bytes"
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
)

type testIDGenerator struct {
	next int64
	err  error
}

func (g *testIDGenerator) NextID(context.Context) (int64, error) {
	if g.err != nil {
		return 0, g.err
	}
	g.next++
	return g.next, nil
}

func TestMessageRepositoryCreateAndCommitSendState(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 10_000})

	senderID := base + 101
	receiverID := base + 102
	randomID := base + 103
	requestHash := payload.HashBytes([]byte("send request"))

	state, err := repo.CreateOrLoadSendState(ctx, CreateSendStateInput{
		SenderUserID:                senderID,
		PeerType:                    payload.PeerTypeUser,
		PeerID:                      receiverID,
		ClientRandomID:              randomID,
		RequestPayloadSchemaVersion: 1,
		RequestPayloadHash:          requestHash,
	})
	if err != nil {
		t.Fatalf("CreateOrLoadSendState() error = %v", err)
	}
	if state.SendStateID == 0 || state.Status != SendStateStatusInitialized {
		t.Fatalf("unexpected initial state: %+v", state)
	}

	canonical, err := repo.CreateOrGetByClientRandom(ctx, CreateCanonicalMessageInput{
		SendStateID:        state.SendStateID,
		SenderUserID:       senderID,
		PeerType:           payload.PeerTypeUser,
		PeerID:             receiverID,
		ClientRandomID:     randomID,
		RequestPayloadHash: requestHash,
		MessageText:        "hello",
		MessageDate:        int32(time.Now().Unix()),
	})
	if err != nil {
		t.Fatalf("CreateOrGetByClientRandom() error = %v", err)
	}
	if canonical.SendStateID != state.SendStateID || canonical.CanonicalMessageID == 0 || canonical.PeerSeq != 1 || !canonical.CreatedNew {
		t.Fatalf("unexpected canonical result: %+v", canonical)
	}

	again, err := repo.CreateOrGetByClientRandom(ctx, CreateCanonicalMessageInput{
		SendStateID:        state.SendStateID,
		SenderUserID:       senderID,
		PeerType:           payload.PeerTypeUser,
		PeerID:             receiverID,
		ClientRandomID:     randomID,
		RequestPayloadHash: requestHash,
		MessageText:        "hello retry",
		MessageDate:        int32(time.Now().Unix()),
	})
	if err != nil {
		t.Fatalf("CreateOrGetByClientRandom() retry error = %v", err)
	}
	if again.CanonicalMessageID != canonical.CanonicalMessageID || again.PeerSeq != canonical.PeerSeq || again.MessageDate != canonical.MessageDate || again.CreatedNew {
		t.Fatalf("idempotent canonical mismatch: first=%+v again=%+v", canonical, again)
	}

	if err := repo.MarkCanonicalCreated(ctx, canonical.SendStateID, canonical.CanonicalMessageID, canonical.PeerSeq); err != nil {
		t.Fatalf("MarkCanonicalCreated() error = %v", err)
	}
	senderUpdate := []byte(`{"schema_version":1,"pts":7,"pts_count":1}`)
	senderUpdateHash := payload.HashBytes(senderUpdate)
	if err := repo.MarkSenderCommitted(ctx, MarkSenderCommittedInput{
		SendStateID:               canonical.SendStateID,
		SenderOperationID:         payload.SenderOperationID(canonical.CanonicalMessageID, senderID),
		SenderPTS:                 7,
		SenderPTSCount:            1,
		SenderUpdateSchemaVersion: payload.OperationResponseSchemaVersion,
		SenderUpdatePayload:       senderUpdate,
		SenderUpdatePayloadHash:   senderUpdateHash,
	}); err != nil {
		t.Fatalf("MarkSenderCommitted() error = %v", err)
	}
	if err := repo.MarkReceiverOpsAcked(ctx, canonical.SendStateID, 0); err != nil {
		t.Fatalf("MarkReceiverOpsAcked() error = %v", err)
	}
	if err := repo.MarkCompleted(ctx, canonical.SendStateID); err != nil {
		t.Fatalf("MarkCompleted() error = %v", err)
	}

	committed, err := repo.CreateOrLoadSendState(ctx, CreateSendStateInput{
		SenderUserID:                senderID,
		PeerType:                    payload.PeerTypeUser,
		PeerID:                      receiverID,
		ClientRandomID:              randomID,
		RequestPayloadSchemaVersion: 1,
		RequestPayloadHash:          requestHash,
	})
	if err != nil {
		t.Fatalf("CreateOrLoadSendState() committed load error = %v", err)
	}
	if committed.Status != SendStateStatusCompleted ||
		committed.CanonicalMessageID != canonical.CanonicalMessageID ||
		committed.PeerSeq != canonical.PeerSeq ||
		committed.SenderPTS != 7 ||
		!bytes.Equal(committed.SenderUpdatePayloadHash, senderUpdateHash) {
		t.Fatalf("unexpected committed state: %+v", committed)
	}
}

func TestMessageRepositoryRandomIdConflict(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 20_000})

	in := CreateSendStateInput{
		SenderUserID:                base + 201,
		PeerType:                    payload.PeerTypeUser,
		PeerID:                      base + 202,
		ClientRandomID:              base + 203,
		RequestPayloadSchemaVersion: 1,
		RequestPayloadHash:          payload.HashBytes([]byte("first")),
	}
	if _, err := repo.CreateOrLoadSendState(ctx, in); err != nil {
		t.Fatalf("CreateOrLoadSendState() error = %v", err)
	}
	in.RequestPayloadHash = payload.HashBytes([]byte("different"))
	_, err := repo.CreateOrLoadSendState(ctx, in)
	if !errors.Is(err, msg.ErrRandomIdConflict) {
		t.Fatalf("CreateOrLoadSendState() error = %v, want ErrRandomIdConflict", err)
	}
}

func TestMessageRepositoryListHistoryMessages(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 40_000})

	senderID := base + 401
	receiverID := base + 402
	firstDate := int32(time.Now().Unix())
	secondDate := firstDate + 1

	first := createCanonicalMessageForTest(t, ctx, repo, senderID, receiverID, base+403, "first", firstDate)
	second := createCanonicalMessageForTest(t, ctx, repo, senderID, receiverID, base+404, "second", secondDate)

	history, err := repo.ListHistoryMessages(ctx, ListHistoryMessagesInput{
		PeerType: payload.PeerTypeUser,
		PeerID:   receiverID,
		Limit:    10,
	})
	if err != nil {
		t.Fatalf("ListHistoryMessages() error = %v", err)
	}
	if len(history) != 2 {
		t.Fatalf("ListHistoryMessages() len = %d, want 2: %+v", len(history), history)
	}
	if history[0].CanonicalMessageID != second.CanonicalMessageID ||
		history[0].PeerSeq != second.PeerSeq ||
		history[0].FromUserID != senderID ||
		history[0].MessageText != "second" ||
		history[0].MessageDate != secondDate {
		t.Fatalf("unexpected newest history row: %+v, canonical: %+v", history[0], second)
	}
	if history[1].CanonicalMessageID != first.CanonicalMessageID ||
		history[1].PeerSeq != first.PeerSeq ||
		history[1].FromUserID != senderID ||
		history[1].MessageText != "first" ||
		history[1].MessageDate != firstDate {
		t.Fatalf("unexpected older history row: %+v, canonical: %+v", history[1], first)
	}

	beforeSecond, err := repo.ListHistoryMessages(ctx, ListHistoryMessagesInput{
		PeerType: payload.PeerTypeUser,
		PeerID:   receiverID,
		OffsetID: int32(second.PeerSeq),
		Limit:    10,
	})
	if err != nil {
		t.Fatalf("ListHistoryMessages() offset error = %v", err)
	}
	if len(beforeSecond) != 1 || beforeSecond[0].CanonicalMessageID != first.CanonicalMessageID {
		t.Fatalf("ListHistoryMessages() offset = %+v, want first canonical id %d", beforeSecond, first.CanonicalMessageID)
	}
}

func TestMessageRepositoryMarkSenderCommittedRejectsOutOfRangePTS(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 30_000})

	state, err := repo.CreateOrLoadSendState(ctx, CreateSendStateInput{
		SenderUserID:                base + 301,
		PeerType:                    payload.PeerTypeUser,
		PeerID:                      base + 302,
		ClientRandomID:              base + 303,
		RequestPayloadSchemaVersion: 1,
		RequestPayloadHash:          payload.HashBytes([]byte("send request")),
	})
	if err != nil {
		t.Fatalf("CreateOrLoadSendState() error = %v", err)
	}

	err = repo.MarkSenderCommitted(ctx, MarkSenderCommittedInput{
		SendStateID:               state.SendStateID,
		SenderOperationID:         payload.SenderOperationID(base+304, base+301),
		SenderPTS:                 1 << 40,
		SenderPTSCount:            1,
		SenderUpdateSchemaVersion: payload.OperationResponseSchemaVersion,
		SenderUpdatePayload:       []byte(`{}`),
		SenderUpdatePayloadHash:   payload.HashBytes([]byte(`{}`)),
	})
	if !errors.Is(err, msg.ErrSenderSyncFailed) {
		t.Fatalf("MarkSenderCommitted() error = %v, want ErrSenderSyncFailed", err)
	}
}

func createCanonicalMessageForTest(
	t *testing.T,
	ctx context.Context,
	repo *Repository,
	senderID int64,
	receiverID int64,
	randomID int64,
	text string,
	date int32,
) *CanonicalMessageResult {
	t.Helper()
	requestHash := payload.HashBytes([]byte(text))
	state, err := repo.CreateOrLoadSendState(ctx, CreateSendStateInput{
		SenderUserID:                senderID,
		PeerType:                    payload.PeerTypeUser,
		PeerID:                      receiverID,
		ClientRandomID:              randomID,
		RequestPayloadSchemaVersion: 1,
		RequestPayloadHash:          requestHash,
	})
	if err != nil {
		t.Fatalf("CreateOrLoadSendState(%q) error = %v", text, err)
	}
	canonical, err := repo.CreateOrGetByClientRandom(ctx, CreateCanonicalMessageInput{
		SendStateID:        state.SendStateID,
		SenderUserID:       senderID,
		PeerType:           payload.PeerTypeUser,
		PeerID:             receiverID,
		ClientRandomID:     randomID,
		RequestPayloadHash: requestHash,
		MessageText:        text,
		MessageDate:        date,
	})
	if err != nil {
		t.Fatalf("CreateOrGetByClientRandom(%q) error = %v", text, err)
	}
	return canonical
}

func openIntegrationDB(t *testing.T) *sqlx.DB {
	t.Helper()
	if testing.Short() {
		t.Skip("integration test skipped in short mode")
	}
	dsn := os.Getenv("TEAMGRAM_TEST_MYSQL_DSN")
	explicit := dsn != ""
	if dsn == "" {
		dsn = "root:@tcp(127.0.0.1:3306)/teamgram?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	}
	db, err := sqlx.Open(&sqlx.Config{DSN: dsn})
	if err != nil {
		if explicit {
			t.Fatalf("open mysql: %v", err)
		}
		t.Skipf("mysql unavailable: %v", err)
	}
	if _, err := db.Exec(context.Background(), "SELECT 1"); err != nil {
		if explicit {
			t.Fatalf("ping mysql: %v", err)
		}
		t.Skipf("mysql unavailable: %v", err)
	}
	return db
}
