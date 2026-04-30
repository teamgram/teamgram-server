//go:build integration

package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
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

func TestApplyUserOperationFinalTransaction(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000

	t.Run("commit idempotency and projections", func(t *testing.T) {
		userID := base + 101
		repo := NewForTest(db, &testIDGenerator{next: base + 10_000}, "local-userupdates")
		in := buildApplyInput(t, userID, userID, base+201, true, "hello")

		if _, err := repo.ClaimPartitionOwner(ctx, in.PartitionID); err != nil {
			t.Fatalf("ClaimPartitionOwner() error = %v", err)
		}

		result, err := repo.ApplyUserOperation(ctx, in)
		if err != nil {
			t.Fatalf("ApplyUserOperation() error = %v", err)
		}
		if result.Pts != 1 || result.PtsCount != 1 || result.AlreadyApplied {
			t.Fatalf("first result mismatch: %+v", result)
		}

		again, err := repo.ApplyUserOperation(ctx, in)
		if err != nil {
			t.Fatalf("ApplyUserOperation() idempotent error = %v", err)
		}
		if again.Pts != result.Pts || !again.AlreadyApplied {
			t.Fatalf("idempotent result mismatch: first=%+v again=%+v", result, again)
		}

		state, err := repo.GetState(ctx, userID)
		if err != nil {
			t.Fatalf("GetState() error = %v", err)
		}
		if state.Pts != 1 {
			t.Fatalf("state pts = %d, want 1", state.Pts)
		}

		diff, err := repo.GetDifference(ctx, GetDifferenceInput{UserID: userID, Pts: 0, Limit: 10})
		if err != nil {
			t.Fatalf("GetDifference() error = %v", err)
		}
		if len(diff.Events) != 1 {
			t.Fatalf("events length = %d, want 1", len(diff.Events))
		}
		if diff.Events[0].Pts != 1 || string(diff.Events[0].EventPayload) == "" {
			t.Fatalf("event mismatch: %+v", diff.Events[0])
		}

		opResult, err := repo.GetOperationResult(ctx, userID, in.OperationID)
		if err != nil {
			t.Fatalf("GetOperationResult() error = %v", err)
		}
		if opResult.Pts != 1 || !bytes.Equal(opResult.PayloadHash, in.PayloadHash) {
			t.Fatalf("operation result mismatch: %+v", opResult)
		}
	})

	t.Run("push task is immediately eligible by available_at", func(t *testing.T) {
		userID := base + 1101
		repo := NewForTest(db, &testIDGenerator{next: base + 60_000}, "local-userupdates")
		in := buildApplyInput(t, userID, userID, base+1201, true, "hello")
		if _, err := repo.ClaimPartitionOwner(ctx, in.PartitionID); err != nil {
			t.Fatalf("ClaimPartitionOwner() error = %v", err)
		}

		before := time.Now().UTC().Add(-time.Second)
		if _, err := repo.ApplyUserOperation(ctx, in); err != nil {
			t.Fatalf("ApplyUserOperation() error = %v", err)
		}
		after := time.Now().UTC().Add(time.Second)

		rows, err := repo.models.PushTaskOutboxModel.SelectPending(ctx, PushTaskStatusPending, mysqlTestTime(after), 10_000)
		if err != nil {
			t.Fatalf("SelectPending() error = %v", err)
		}
		var matched *model.PushTaskOutbox
		for i := range rows {
			if rows[i].UserId == userID && rows[i].OperationId == in.OperationID {
				matched = &rows[i]
				break
			}
		}
		if matched == nil {
			t.Fatalf("SelectPending() did not return push task for user_id=%d operation_id=%s; row_count=%d", userID, in.OperationID, len(rows))
		}
		availableAt := normalizeDBTestTime(t, matched.AvailableAt)
		if availableAt < mysqlTestTime(before) || availableAt > mysqlTestTime(after) {
			t.Fatalf("AvailableAt = %q, want between %q and %q", matched.AvailableAt, mysqlTestTime(before), mysqlTestTime(after))
		}
	})

	t.Run("push pending selection uses available_at", func(t *testing.T) {
		now := time.Now().UTC().Truncate(time.Second)
		past := now.Add(-time.Minute)
		future := now.Add(time.Hour)
		firstTaskID := base + 70_001
		secondTaskID := base + 70_002
		futureTaskID := base + 70_003
		status := int32(100_000 + base%1_000_000_000)

		insertTestPushTask(t, ctx, db, secondTaskID, status, past)
		insertTestPushTask(t, ctx, db, firstTaskID, status, past)
		insertTestPushTask(t, ctx, db, futureTaskID, status, future)

		rows, err := NewForTest(db, &testIDGenerator{next: base + 70_000}, "local-userupdates").
			models.PushTaskOutboxModel.SelectPending(ctx, status, mysqlTestTime(now), 10)
		if err != nil {
			t.Fatalf("SelectPending() error = %v", err)
		}

		positions := map[int64]int{}
		for i := range rows {
			positions[rows[i].TaskId] = i
			if rows[i].TaskId == futureTaskID {
				t.Fatalf("future task_id=%d returned before available_at", futureTaskID)
			}
		}
		firstPos, firstFound := positions[firstTaskID]
		secondPos, secondFound := positions[secondTaskID]
		if !firstFound || !secondFound {
			t.Fatalf("due tasks not returned: first_found=%t second_found=%t row_count=%d", firstFound, secondFound, len(rows))
		}
		if firstPos > secondPos {
			t.Fatalf("tasks with same available_at not ordered by task_id: first_pos=%d second_pos=%d", firstPos, secondPos)
		}

		rows, err = NewForTest(db, &testIDGenerator{next: base + 71_000}, "local-userupdates").
			models.PushTaskOutboxModel.SelectPending(ctx, status, mysqlTestTime(future), 10)
		if err != nil {
			t.Fatalf("SelectPending(future) error = %v", err)
		}
		foundFuture := false
		for i := range rows {
			if rows[i].TaskId == futureTaskID {
				foundFuture = true
				break
			}
		}
		if !foundFuture {
			t.Fatalf("future task_id=%d not returned once available; row_count=%d", futureTaskID, len(rows))
		}
	})

	t.Run("operation id conflict", func(t *testing.T) {
		userID := base + 301
		repo := NewForTest(db, &testIDGenerator{next: base + 20_000}, "local-userupdates")
		in := buildApplyInput(t, userID, userID, base+401, true, "hello")
		if _, err := repo.ClaimPartitionOwner(ctx, in.PartitionID); err != nil {
			t.Fatalf("ClaimPartitionOwner() error = %v", err)
		}
		if _, err := repo.ApplyUserOperation(ctx, in); err != nil {
			t.Fatalf("ApplyUserOperation() error = %v", err)
		}

		conflict := buildApplyInput(t, userID, userID, base+401, true, "different")
		conflict.OperationID = in.OperationID
		_, err := repo.ApplyUserOperation(ctx, conflict)
		if !errors.Is(err, userupdates.ErrOperationPayloadConflict) {
			t.Fatalf("ApplyUserOperation() error = %v, want ErrOperationPayloadConflict", err)
		}
	})

	t.Run("owner mismatch rejects without event", func(t *testing.T) {
		userID := base + 501
		goodRepo := NewForTest(db, &testIDGenerator{next: base + 30_000}, "local-userupdates")
		badRepo := NewForTest(db, &testIDGenerator{next: base + 31_000}, "other-userupdates")
		in := buildApplyInput(t, userID, userID, base+601, true, "hello")
		if _, err := goodRepo.ClaimPartitionOwner(ctx, in.PartitionID); err != nil {
			t.Fatalf("ClaimPartitionOwner() error = %v", err)
		}

		_, err := badRepo.ApplyUserOperation(ctx, in)
		if !errors.Is(err, userupdates.ErrNotOwner) {
			t.Fatalf("ApplyUserOperation() error = %v, want ErrNotOwner", err)
		}
		diff, err := goodRepo.GetDifference(ctx, GetDifferenceInput{UserID: userID, Pts: 0, Limit: 10})
		if err != nil {
			t.Fatalf("GetDifference() error = %v", err)
		}
		if len(diff.Events) != 0 {
			t.Fatalf("events length = %d, want 0", len(diff.Events))
		}
	})

	t.Run("dependency is terminal in first slice", func(t *testing.T) {
		userID := base + 701
		repo := NewForTest(db, &testIDGenerator{next: base + 40_000}, "local-userupdates")
		in := buildApplyInput(t, userID, userID, base+801, true, "hello")
		in.DependencyPts = []int64{1}
		if _, err := repo.ClaimPartitionOwner(ctx, in.PartitionID); err != nil {
			t.Fatalf("ClaimPartitionOwner() error = %v", err)
		}

		_, err := repo.ApplyUserOperation(ctx, in)
		if !errors.Is(err, userupdates.ErrOperationTerminal) {
			t.Fatalf("ApplyUserOperation() error = %v, want ErrOperationTerminal", err)
		}
	})

	t.Run("push task failure rolls back transaction", func(t *testing.T) {
		userID := base + 901
		repo := NewForTest(db, &testIDGenerator{next: base + 50_000}, "local-userupdates")
		in := buildApplyInput(t, userID, userID, base+1001, true, "hello")
		if _, err := repo.ClaimPartitionOwner(ctx, in.PartitionID); err != nil {
			t.Fatalf("ClaimPartitionOwner() error = %v", err)
		}

		failingRepo := NewForTest(db, &testIDGenerator{err: errors.New("idgen down")}, "local-userupdates")
		_, err := failingRepo.ApplyUserOperation(ctx, in)
		if !errors.Is(err, userupdates.ErrUserupdatesStorage) {
			t.Fatalf("ApplyUserOperation() error = %v, want ErrUserupdatesStorage", err)
		}

		state, err := repo.GetState(ctx, userID)
		if err != nil {
			t.Fatalf("GetState() error = %v", err)
		}
		if state.Pts != 0 {
			t.Fatalf("state pts = %d, want rollback to 0", state.Pts)
		}
	})
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

func mysqlTestTime(t time.Time) string {
	return t.UTC().Format("2006-01-02 15:04:05.000000")
}

func normalizeDBTestTime(t *testing.T, value string) string {
	t.Helper()
	if parsed, err := time.Parse(time.RFC3339Nano, value); err == nil {
		return parsed.Format("2006-01-02 15:04:05.000000")
	}
	if parsed, err := time.Parse("2006-01-02 15:04:05.999999", value); err == nil {
		return parsed.Format("2006-01-02 15:04:05.000000")
	}
	t.Fatalf("unsupported DB time format %q", value)
	return ""
}

func insertTestPushTask(t *testing.T, ctx context.Context, db *sqlx.DB, taskID int64, status int32, availableAt time.Time) {
	t.Helper()
	_, err := db.Exec(ctx, `
INSERT INTO push_task_outbox
	(task_id, user_id, pts, push_type, peer_type, peer_id, operation_id,
	 push_partition_id, task_schema_version, task_codec, task_payload, status,
	 publish_attempts, available_at, published_topic, published_partition,
	 published_offset, last_error_code)
VALUES
	(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		taskID,
		taskID+10_000,
		taskID,
		PushTypeUserUpdate,
		payload.PeerTypeUser,
		taskID+20_000,
		fmt.Sprintf("test-op-%d", taskID),
		int32(payload.RouteUser(taskID+10_000).PushPartitionID),
		int32(1),
		PayloadCodecJSON,
		[]byte(`{"test":true}`),
		status,
		int32(0),
		mysqlTestTime(availableAt),
		"",
		int32(0),
		int64(0),
		"",
	)
	if err != nil {
		t.Fatalf("insert test push task: %v", err)
	}
}

func buildApplyInput(t *testing.T, userID, fromUserID, toUserID int64, out bool, text string) ApplyUserOperationInput {
	t.Helper()
	route := payload.RouteUser(userID)
	op := payload.MessageOperationV1{
		SchemaVersion:      payload.MessageOperationSchemaVersion,
		OperationKind:      payload.OperationKindSendMessage,
		CanonicalMessageID: userID*10 + 1,
		PeerType:           payload.PeerTypeUser,
		PeerID:             toUserID,
		PeerSeq:            1,
		FromUserID:         fromUserID,
		ToUserID:           toUserID,
		Date:               int32(time.Now().Unix()),
		Out:                out,
		MessageText:        text,
	}
	body, err := json.Marshal(op)
	if err != nil {
		t.Fatalf("marshal operation: %v", err)
	}
	return ApplyUserOperationInput{
		UserID:       userID,
		OperationID:  payload.SenderOperationID(op.CanonicalMessageID, userID),
		OpType:       OpTypeSendMessage,
		PeerType:     op.PeerType,
		PeerID:       op.PeerID,
		PayloadCodec: PayloadCodecJSON,
		Payload:      body,
		PayloadHash:  payload.HashBytes(body),
		BucketID:     int32(route.BucketID),
		PartitionID:  int32(route.ReceiverPartitionID),
	}
}
