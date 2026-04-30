//go:build integration

package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
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
