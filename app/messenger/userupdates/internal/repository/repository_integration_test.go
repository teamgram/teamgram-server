//go:build integration

package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
)

type testIDGenerator struct {
	next      int64
	err       error
	failAfter int
	calls     int
}

func (g *testIDGenerator) NextID(context.Context) (int64, error) {
	if g.err != nil {
		return 0, g.err
	}
	g.calls++
	if g.failAfter > 0 && g.calls > g.failAfter {
		return 0, errors.New("idgen down")
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

		state, err := repo.GetState(ctx, userID, 0)
		if err != nil {
			t.Fatalf("GetState() error = %v", err)
		}
		if state.Pts != 1 {
			t.Fatalf("state pts = %d, want 1", state.Pts)
		}
		if state.UnreadCount != 0 {
			t.Fatalf("state unread_count = %d, want 0", state.UnreadCount)
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

	t.Run("state unread count includes first incoming message", func(t *testing.T) {
		receiverID := base + 2101
		senderID := base + 2102
		repo := NewForTest(db, &testIDGenerator{next: base + 70_000}, "local-userupdates")
		op := payload.MessageOperationV1{
			SchemaVersion:      payload.MessageOperationSchemaVersion,
			OperationKind:      payload.OperationKindSendMessage,
			CanonicalMessageID: receiverID*10 + 1,
			PeerType:           payload.PeerTypeUser,
			PeerID:             senderID,
			PeerSeq:            1,
			FromUserID:         senderID,
			ToUserID:           receiverID,
			Date:               int32(time.Now().Unix()),
			Out:                false,
			MessageText:        "incoming hello",
		}
		in := buildOperationApplyInput(t, receiverID, op, "incoming")
		if _, err := repo.ClaimPartitionOwner(ctx, in.PartitionID); err != nil {
			t.Fatalf("ClaimPartitionOwner() error = %v", err)
		}
		if _, err := repo.ApplyUserOperation(ctx, in); err != nil {
			t.Fatalf("ApplyUserOperation() error = %v", err)
		}

		state, err := repo.GetState(ctx, receiverID, 0)
		if err != nil {
			t.Fatalf("GetState() error = %v", err)
		}
		if state.Pts != 1 || state.UnreadCount != 1 {
			t.Fatalf("state = %+v, want pts=1 unread_count=1", state)
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

	t.Run("failed operation replay uses failed_at", func(t *testing.T) {
		now := time.Now().UTC().Truncate(time.Second)
		bucketID := int32(base % 10_000)
		status := int32(210_000 + base%1_000_000)
		firstFailedID := base + 80_001
		secondFailedID := base + 80_002
		futureFailedID := base + 80_003

		insertTestDeliveryFailedOperation(t, ctx, db, secondFailedID, bucketID, status, now.Add(-time.Minute))
		insertTestDeliveryFailedOperation(t, ctx, db, firstFailedID, bucketID, status, now.Add(-time.Minute))
		// SelectByBucketStatus is a bucket FIFO replay listing. It orders by failed_at
		// but does not filter by failed_at <= now.
		insertTestDeliveryFailedOperation(t, ctx, db, futureFailedID, bucketID, status, now.Add(time.Hour))

		rows, err := NewForTest(db, &testIDGenerator{next: base + 80_000}, "local-userupdates").
			models.DeliveryFailedOperationsModel.SelectByBucketStatus(ctx, bucketID, status, 10)
		if err != nil {
			t.Fatalf("SelectByBucketStatus() error = %v", err)
		}
		if len(rows) < 3 {
			t.Fatalf("SelectByBucketStatus() row_count=%d, want at least 3", len(rows))
		}
		if rows[0].FailedId != firstFailedID || rows[1].FailedId != secondFailedID || rows[2].FailedId != futureFailedID {
			t.Fatalf("failed replay order = [%d %d %d], want [%d %d %d]",
				rows[0].FailedId, rows[1].FailedId, rows[2].FailedId,
				firstFailedID, secondFailedID, futureFailedID)
		}
	})

	t.Run("terminal operation retention uses completed_at", func(t *testing.T) {
		now := time.Now().UTC().Truncate(time.Second)
		status := int32(220_000 + base%1_000_000)
		oldUserID := base + 90_001
		newUserID := base + 90_002
		oldOperationID := fmt.Sprintf("retention-old-%d", base)
		newOperationID := fmt.Sprintf("retention-new-%d", base)

		insertTestOperationResult(t, ctx, db, oldUserID, oldOperationID, status, now.Add(-time.Hour))
		insertTestOperationResult(t, ctx, db, newUserID, newOperationID, status, now.Add(time.Hour))

		rows, err := NewForTest(db, &testIDGenerator{next: base + 90_000}, "local-userupdates").
			models.UserOperationResultsModel.SelectByStatusCompletedBefore(ctx, status, mysqlTestTime(now), 10)
		if err != nil {
			t.Fatalf("SelectByStatusCompletedBefore() error = %v", err)
		}
		if len(rows) != 1 {
			t.Fatalf("SelectByStatusCompletedBefore() row_count=%d, want 1: %+v", len(rows), rows)
		}
		if rows[0].UserId != oldUserID || rows[0].OperationId != oldOperationID {
			t.Fatalf("retention row = %+v, want user_id=%d operation_id=%s", rows[0], oldUserID, oldOperationID)
		}
	})

	t.Run("pts event business reads do not expose created_at", func(t *testing.T) {
		eventType := reflect.TypeOf(model.UserPtsEvents{})
		if _, ok := eventType.FieldByName("CreatedAt"); ok {
			t.Fatalf("model.UserPtsEvents exposes CreatedAt; lifecycle fields must not be normal business read fields")
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

	t.Run("sync operation claims missing or unassigned fence", func(t *testing.T) {
		userID := base + 601
		repo := NewForTest(db, &testIDGenerator{next: base + 35_000}, "local-userupdates")
		in := buildApplyInput(t, userID, userID, base+701, true, "hello")
		_ = repo.models.UserupdatesPartitionFencesModel.Delete2(ctx, in.PartitionID)

		result, err := repo.ApplyUserOperation(ctx, in)
		if err != nil {
			t.Fatalf("ApplyUserOperation() error = %v", err)
		}
		if result.Pts != 1 {
			t.Fatalf("result pts = %d, want 1", result.Pts)
		}
		fence, err := repo.models.UserupdatesPartitionFencesModel.SelectByPartitionId(ctx, in.PartitionID)
		if err != nil {
			t.Fatalf("SelectByPartitionId() error = %v", err)
		}
		if fence.OwnerInstanceId != "local-userupdates" || fence.OwnerEpoch != 1 {
			t.Fatalf("fence = %+v, want local owner epoch 1", fence)
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

		state, err := repo.GetState(ctx, userID, 0)
		if err != nil {
			t.Fatalf("GetState() error = %v", err)
		}
		if state.Pts != 0 {
			t.Fatalf("state pts = %d, want rollback to 0", state.Pts)
		}
	})
}

func TestApplyUserOperationAffectedOutboxIdempotency(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	requesterID := base + 3101
	targetID := base + 3102
	repo := NewForTest(db, &testIDGenerator{next: base + 31000}, "local-userupdates")
	in := buildApplyInput(t, requesterID, requesterID, targetID, true, "hello")
	in.AffectedOutboxes = []AffectedOutbox{buildAffectedOutbox(t, in, targetID, "same")}
	if _, err := repo.ClaimPartitionOwner(ctx, in.PartitionID); err != nil {
		t.Fatalf("ClaimPartitionOwner() error = %v", err)
	}
	if _, err := repo.ApplyUserOperation(ctx, in); err != nil {
		t.Fatalf("ApplyUserOperation() error = %v", err)
	}

	duplicate := buildOperationApplyInput(t, requesterID, payload.MessageOperationV1{
		SchemaVersion:      payload.MessageOperationSchemaVersion,
		OperationKind:      payload.OperationKindSendMessage,
		CanonicalMessageID: requesterID*10 + 2,
		PeerType:           payload.PeerTypeUser,
		PeerID:             targetID,
		PeerSeq:            2,
		FromUserID:         requesterID,
		ToUserID:           targetID,
		Date:               int32(time.Now().Unix()),
		Out:                true,
		MessageText:        "second",
	}, "duplicate-affected")
	duplicate.AffectedOutboxes = []AffectedOutbox{buildAffectedOutbox(t, in, targetID, "same")}
	if _, err := repo.ApplyUserOperation(ctx, duplicate); err != nil {
		t.Fatalf("ApplyUserOperation() duplicate same hash error = %v", err)
	}

	row, err := repo.models.AffectedOperationOutboxModel.SelectByUserOperation(ctx, targetID, in.AffectedOutboxes[0].OperationID)
	if err != nil {
		t.Fatalf("SelectByUserOperation() error = %v", err)
	}
	if row.Status != AffectedOutboxStatusPending || row.UserId != targetID || row.RequesterUserId != requesterID {
		t.Fatalf("affected outbox row = %+v, want pending target/requester", row)
	}

	conflict := buildOperationApplyInput(t, requesterID, payload.MessageOperationV1{
		SchemaVersion:      payload.MessageOperationSchemaVersion,
		OperationKind:      payload.OperationKindSendMessage,
		CanonicalMessageID: requesterID*10 + 3,
		PeerType:           payload.PeerTypeUser,
		PeerID:             targetID,
		PeerSeq:            3,
		FromUserID:         requesterID,
		ToUserID:           targetID,
		Date:               int32(time.Now().Unix()),
		Out:                true,
		MessageText:        "third",
	}, "conflict-affected")
	conflict.AffectedOutboxes = []AffectedOutbox{buildAffectedOutbox(t, in, targetID, "different")}
	_, err = repo.ApplyUserOperation(ctx, conflict)
	if !errors.Is(err, userupdates.ErrOperationPayloadConflict) {
		t.Fatalf("ApplyUserOperation() duplicate different hash error = %v, want ErrOperationPayloadConflict", err)
	}
}

func TestApplyUserOperationRollsBackWhenAffectedOutboxInsertFails(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	requesterID := base + 3201
	targetID := base + 3202
	goodRepo := NewForTest(db, &testIDGenerator{next: base + 32000}, "local-userupdates")
	in := buildApplyInput(t, requesterID, requesterID, targetID, true, "hello")
	in.AffectedOutboxes = []AffectedOutbox{buildAffectedOutbox(t, in, targetID, "rollback")}
	if _, err := goodRepo.ClaimPartitionOwner(ctx, in.PartitionID); err != nil {
		t.Fatalf("ClaimPartitionOwner() error = %v", err)
	}

	failingRepo := NewForTest(db, &testIDGenerator{next: base + 33000, failAfter: 1}, "local-userupdates")
	_, err := failingRepo.ApplyUserOperation(ctx, in)
	if !errors.Is(err, userupdates.ErrUserupdatesStorage) {
		t.Fatalf("ApplyUserOperation() error = %v, want ErrUserupdatesStorage", err)
	}

	state, err := goodRepo.GetState(ctx, requesterID, 0)
	if err != nil {
		t.Fatalf("GetState() error = %v", err)
	}
	if state.Pts != 0 {
		t.Fatalf("state pts = %d, want rollback to 0", state.Pts)
	}
	diff, err := goodRepo.GetDifference(ctx, GetDifferenceInput{UserID: requesterID, Pts: 0, Limit: 10})
	if err != nil {
		t.Fatalf("GetDifference() error = %v", err)
	}
	if len(diff.Events) != 0 {
		t.Fatalf("events length = %d, want 0", len(diff.Events))
	}
	if _, err := goodRepo.GetOperationResult(ctx, requesterID, in.OperationID); !errors.Is(err, userupdates.ErrOperationTerminal) {
		t.Fatalf("GetOperationResult() error = %v, want ErrOperationTerminal", err)
	}
	if _, err := goodRepo.models.AffectedOperationOutboxModel.SelectByUserOperation(ctx, targetID, in.AffectedOutboxes[0].OperationID); !errors.Is(err, model.ErrNotFound) {
		t.Fatalf("SelectByUserOperation() error = %v, want model.ErrNotFound", err)
	}
}

func TestApplyReadHistoryUpdatesReadStateUnreadMarkAndPTS(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	userID := base + 2101
	peerID := base + 2102
	repo := NewForTest(db, &testIDGenerator{next: base + 21000}, "local-userupdates")
	send := buildOperationApplyInput(t, userID, payload.MessageOperationV1{
		SchemaVersion:      payload.MessageOperationSchemaVersion,
		OperationKind:      payload.OperationKindSendMessage,
		CanonicalMessageID: userID*10 + 1,
		PeerType:           payload.PeerTypeUser,
		PeerID:             peerID,
		PeerSeq:            1,
		FromUserID:         peerID,
		ToUserID:           userID,
		Date:               int32(time.Now().Unix()),
		Out:                false,
		MessageText:        "incoming",
	}, "send")
	if _, err := repo.ClaimPartitionOwner(ctx, send.PartitionID); err != nil {
		t.Fatalf("ClaimPartitionOwner() error = %v", err)
	}
	if _, err := repo.ApplyUserOperation(ctx, send); err != nil {
		t.Fatalf("ApplyUserOperation(send) error = %v", err)
	}
	read := buildOperationApplyInput(t, userID, payload.MessageOperationV1{
		SchemaVersion:        payload.MessageOperationSchemaVersion,
		OperationKind:        payload.OperationKindReadHistory,
		PeerType:             payload.PeerTypeUser,
		PeerID:               peerID,
		PeerSeq:              1,
		ReadInboxMaxPeerSeq:  1,
		ReadOutboxMaxPeerSeq: 0,
		FromUserID:           userID,
		ToUserID:             userID,
		Date:                 int32(time.Now().Unix()),
	}, "read")
	if _, err := repo.ApplyUserOperation(ctx, read); err != nil {
		t.Fatalf("ApplyUserOperation(read) error = %v", err)
	}
	row, err := repo.models.UserDialogsModel.SelectByUserPeer(ctx, userID, payload.PeerTypeUser, peerID)
	if err != nil {
		t.Fatalf("SelectByUserPeer() error = %v", err)
	}
	if row.UnreadCount != 0 || row.UnreadMark || row.ReadInboxMaxPeerSeq != 1 || row.LastPts != 2 {
		t.Fatalf("dialog read state = %+v, want unread cleared, read_inbox=1, last_pts=2", row)
	}
	readOutbox := buildOperationApplyInput(t, userID, payload.MessageOperationV1{
		SchemaVersion:        payload.MessageOperationSchemaVersion,
		OperationKind:        payload.OperationKindReadHistory,
		PeerType:             payload.PeerTypeUser,
		PeerID:               peerID,
		PeerSeq:              2,
		ReadOutboxMaxPeerSeq: 2,
		FromUserID:           peerID,
		ToUserID:             userID,
		Date:                 int32(time.Now().Unix()),
		Out:                  true,
	}, "read-outbox")
	if _, err := repo.ApplyUserOperation(ctx, readOutbox); err != nil {
		t.Fatalf("ApplyUserOperation(read outbox) error = %v", err)
	}
	row, err = repo.models.UserDialogsModel.SelectByUserPeer(ctx, userID, payload.PeerTypeUser, peerID)
	if err != nil {
		t.Fatalf("SelectByUserPeer() after read outbox error = %v", err)
	}
	if row.ReadInboxMaxPeerSeq != 1 || row.ReadOutboxMaxPeerSeq != 2 || row.LastPts != 3 {
		t.Fatalf("dialog outbox read state = %+v, want read_inbox preserved and read_outbox=2", row)
	}
}

func TestApplyReadOutboxPreservesIncomingUnreadCount(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	userID := base + 2151
	peerID := base + 2152
	repo := NewForTest(db, &testIDGenerator{next: base + 21500}, "local-userupdates")
	send := buildOperationApplyInput(t, userID, payload.MessageOperationV1{
		SchemaVersion:      payload.MessageOperationSchemaVersion,
		OperationKind:      payload.OperationKindSendMessage,
		CanonicalMessageID: userID*10 + 1,
		PeerType:           payload.PeerTypeUser,
		PeerID:             peerID,
		PeerSeq:            1,
		FromUserID:         peerID,
		ToUserID:           userID,
		Date:               int32(time.Now().Unix()),
		Out:                false,
		MessageText:        "incoming",
	}, "send")
	if _, err := repo.ClaimPartitionOwner(ctx, send.PartitionID); err != nil {
		t.Fatalf("ClaimPartitionOwner() error = %v", err)
	}
	if _, err := repo.ApplyUserOperation(ctx, send); err != nil {
		t.Fatalf("ApplyUserOperation(send) error = %v", err)
	}

	readOutbox := buildOperationApplyInput(t, userID, payload.MessageOperationV1{
		SchemaVersion:        payload.MessageOperationSchemaVersion,
		OperationKind:        payload.OperationKindReadHistory,
		PeerType:             payload.PeerTypeUser,
		PeerID:               peerID,
		PeerSeq:              1,
		ReadOutboxMaxPeerSeq: 1,
		FromUserID:           peerID,
		ToUserID:             userID,
		Date:                 int32(time.Now().Unix()),
		Out:                  true,
	}, "read-outbox")
	if _, err := repo.ApplyUserOperation(ctx, readOutbox); err != nil {
		t.Fatalf("ApplyUserOperation(read outbox) error = %v", err)
	}
	row, err := repo.models.UserDialogsModel.SelectByUserPeer(ctx, userID, payload.PeerTypeUser, peerID)
	if err != nil {
		t.Fatalf("SelectByUserPeer() error = %v", err)
	}
	if row.UnreadCount != 1 || row.ReadInboxMaxPeerSeq != 0 || row.ReadOutboxMaxPeerSeq != 1 || row.LastPts != 2 {
		t.Fatalf("dialog read outbox state = %+v, want unread preserved, read_inbox=0, read_outbox=1, last_pts=2", row)
	}
	state, err := repo.GetState(ctx, userID, 0)
	if err != nil {
		t.Fatalf("GetState() error = %v", err)
	}
	if state.UnreadCount != 1 {
		t.Fatalf("state unread_count = %d, want 1", state.UnreadCount)
	}
}

func TestApplyDeleteMessagesMaterializesPublicIDsAndDialogTop(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	userID := base + 2311
	peerID := base + 2312
	cleanupDeleteMessagesIntegrationRows(t, ctx, db, userID)
	repo := NewForTest(db, &testIDGenerator{next: base + 51000}, "local-userupdates")
	if _, err := repo.ClaimPartitionOwner(ctx, int32(payload.RouteUser(userID).ReceiverPartitionID)); err != nil {
		t.Fatalf("ClaimPartitionOwner() error = %v", err)
	}
	for i := int64(1); i <= 2; i++ {
		in := buildOperationApplyInput(t, userID, payload.MessageOperationV1{
			SchemaVersion:      payload.MessageOperationSchemaVersion,
			OperationKind:      payload.OperationKindSendMessage,
			CanonicalMessageID: userID*10 + i,
			PeerType:           payload.PeerTypeUser,
			PeerID:             peerID,
			PeerSeq:            i,
			FromUserID:         peerID,
			ToUserID:           userID,
			Date:               int32(time.Now().Unix()),
			Out:                false,
			MessageText:        fmt.Sprintf("incoming %d", i),
		}, fmt.Sprintf("delete-materialize-send-%d", i))
		if _, err := repo.ApplyUserOperation(ctx, in); err != nil {
			t.Fatalf("ApplyUserOperation(send %d) error = %v", i, err)
		}
	}

	deleteInput := buildOperationApplyInput(t, userID, payload.MessageOperationV1{
		SchemaVersion: payload.MessageOperationSchemaVersion,
		OperationKind: payload.OperationKindDeleteMessages,
		PeerType:      payload.PeerTypeUser,
		PeerID:        peerID,
		PeerSeq:       2,
		DeletePeerSeqs: []int64{
			2,
		},
		FromUserID: userID,
		ToUserID:   userID,
		Date:       int32(time.Now().Unix()),
	}, "delete-materialize")
	result, err := repo.ApplyUserOperation(ctx, deleteInput)
	if err != nil {
		t.Fatalf("ApplyUserOperation(delete) error = %v", err)
	}
	if result.PtsCount != 1 {
		t.Fatalf("delete pts_count = %d, want 1", result.PtsCount)
	}
	storedEvent, err := repo.models.UserPtsEventsModel.SelectByOperation(ctx, userID, deleteInput.OperationID)
	if err != nil {
		t.Fatalf("select delete event: %v", err)
	}
	var event payload.MessageEventV2
	if err := json.Unmarshal(storedEvent.EventPayload, &event); err != nil {
		t.Fatalf("decode delete event: %v", err)
	}
	if len(event.DeleteUserMessageIDs) != 1 || event.DeleteUserMessageIDs[0] != 2 {
		t.Fatalf("delete event ids = %v, want [2]", event.DeleteUserMessageIDs)
	}
	dialog, err := repo.models.UserDialogsModel.SelectByUserPeer(ctx, userID, payload.PeerTypeUser, peerID)
	if err != nil {
		t.Fatalf("select dialog: %v", err)
	}
	if dialog.TopUserMessageId != 1 || dialog.TopPeerSeq != 1 {
		t.Fatalf("dialog top = user_msg:%d peer_seq:%d, want 1/1", dialog.TopUserMessageId, dialog.TopPeerSeq)
	}
}

func TestApplyDeleteMessagesDecrementsUnreadCount(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	userID := base + 2321
	peerID := base + 2322
	cleanupDeleteMessagesIntegrationRows(t, ctx, db, userID)
	repo := NewForTest(db, &testIDGenerator{next: base + 52000}, "local-userupdates")
	if _, err := repo.ClaimPartitionOwner(ctx, int32(payload.RouteUser(userID).ReceiverPartitionID)); err != nil {
		t.Fatalf("ClaimPartitionOwner() error = %v", err)
	}
	for i := int64(1); i <= 2; i++ {
		in := buildOperationApplyInput(t, userID, payload.MessageOperationV1{
			SchemaVersion:      payload.MessageOperationSchemaVersion,
			OperationKind:      payload.OperationKindSendMessage,
			CanonicalMessageID: userID*10 + i,
			PeerType:           payload.PeerTypeUser,
			PeerID:             peerID,
			PeerSeq:            i,
			FromUserID:         peerID,
			ToUserID:           userID,
			Date:               int32(time.Now().Unix()),
			Out:                false,
			MessageText:        fmt.Sprintf("unread incoming %d", i),
		}, fmt.Sprintf("delete-unread-send-%d", i))
		if _, err := repo.ApplyUserOperation(ctx, in); err != nil {
			t.Fatalf("ApplyUserOperation(send %d) error = %v", i, err)
		}
	}

	firstDelete := buildOperationApplyInput(t, userID, payload.MessageOperationV1{
		SchemaVersion: payload.MessageOperationSchemaVersion,
		OperationKind: payload.OperationKindDeleteMessages,
		PeerType:      payload.PeerTypeUser,
		PeerID:        peerID,
		PeerSeq:       2,
		DeletePeerSeqs: []int64{
			2,
		},
		FromUserID: userID,
		ToUserID:   userID,
		Date:       int32(time.Now().Unix()),
	}, "delete-unread-first")
	if _, err := repo.ApplyUserOperation(ctx, firstDelete); err != nil {
		t.Fatalf("ApplyUserOperation(first delete) error = %v", err)
	}
	dialog, err := repo.models.UserDialogsModel.SelectByUserPeer(ctx, userID, payload.PeerTypeUser, peerID)
	if err != nil {
		t.Fatalf("select dialog after first delete: %v", err)
	}
	if dialog.UnreadCount != 1 {
		t.Fatalf("unread_count = %d, want 1", dialog.UnreadCount)
	}

	secondDelete := buildOperationApplyInput(t, userID, payload.MessageOperationV1{
		SchemaVersion: payload.MessageOperationSchemaVersion,
		OperationKind: payload.OperationKindDeleteMessages,
		PeerType:      payload.PeerTypeUser,
		PeerID:        peerID,
		PeerSeq:       1,
		DeletePeerSeqs: []int64{
			1,
		},
		FromUserID: userID,
		ToUserID:   userID,
		Date:       int32(time.Now().Unix()),
	}, "delete-unread-second")
	if _, err := repo.ApplyUserOperation(ctx, secondDelete); err != nil {
		t.Fatalf("ApplyUserOperation(second delete) error = %v", err)
	}
	dialog, err = repo.models.UserDialogsModel.SelectByUserPeer(ctx, userID, payload.PeerTypeUser, peerID)
	if err != nil {
		t.Fatalf("select dialog after second delete: %v", err)
	}
	if dialog.UnreadCount != 0 {
		t.Fatalf("unread_count = %d, want 0", dialog.UnreadCount)
	}
}

func TestApplyDeleteHistoryPersistsAvailableMinPublicUserMessageID(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	userID := base + 2161
	peerID := base + 2162
	repo := NewForTest(db, &testIDGenerator{next: base + 21600}, "local-userupdates")
	first := buildOperationApplyInput(t, userID, payload.MessageOperationV1{
		SchemaVersion:      payload.MessageOperationSchemaVersion,
		OperationKind:      payload.OperationKindSendMessage,
		CanonicalMessageID: userID*10 + 1,
		PeerType:           payload.PeerTypeUser,
		PeerID:             peerID,
		PeerSeq:            1,
		FromUserID:         userID,
		ToUserID:           peerID,
		Date:               int32(time.Now().Unix()),
		Out:                true,
		MessageText:        "first",
	}, "delete-history-first")
	if _, err := repo.ClaimPartitionOwner(ctx, first.PartitionID); err != nil {
		t.Fatalf("ClaimPartitionOwner() error = %v", err)
	}
	if _, err := repo.ApplyUserOperation(ctx, first); err != nil {
		t.Fatalf("ApplyUserOperation(first) error = %v", err)
	}
	second := buildOperationApplyInput(t, userID, payload.MessageOperationV1{
		SchemaVersion:      payload.MessageOperationSchemaVersion,
		OperationKind:      payload.OperationKindSendMessage,
		CanonicalMessageID: userID*10 + 2,
		PeerType:           payload.PeerTypeUser,
		PeerID:             peerID,
		PeerSeq:            7,
		FromUserID:         userID,
		ToUserID:           peerID,
		Date:               int32(time.Now().Unix()),
		Out:                true,
		MessageText:        "second",
	}, "delete-history-second")
	if _, err := repo.ApplyUserOperation(ctx, second); err != nil {
		t.Fatalf("ApplyUserOperation(second) error = %v", err)
	}
	secondView, err := repo.models.UserMessageViewsModel.SelectByUserPeerSeq(ctx, userID, payload.PeerTypeUser, peerID, 7)
	if err != nil {
		t.Fatalf("SelectByUserPeerSeq(second) error = %v", err)
	}
	if secondView.UserMessageId == 0 || secondView.UserMessageId == secondView.PeerSeq {
		t.Fatalf("test setup requires public id distinct from peer seq, view=%+v", secondView)
	}

	deleteHistory := buildOperationApplyInput(t, userID, payload.MessageOperationV1{
		SchemaVersion:    payload.MessageOperationSchemaVersion,
		OperationKind:    payload.OperationKindDeleteHistory,
		PeerType:         payload.PeerTypeUser,
		PeerID:           peerID,
		PeerSeq:          7,
		DeleteMaxPeerSeq: 7,
		FromUserID:       userID,
		ToUserID:         peerID,
		Date:             int32(time.Now().Unix()),
		JustClear:        true,
	}, "delete-history")
	if _, err := repo.ApplyUserOperation(ctx, deleteHistory); err != nil {
		t.Fatalf("ApplyUserOperation(delete history) error = %v", err)
	}

	dialog, err := repo.models.UserDialogsModel.SelectByUserPeer(ctx, userID, payload.PeerTypeUser, peerID)
	if err != nil {
		t.Fatalf("SelectByUserPeer() error = %v", err)
	}
	if dialog.AvailableMinPeerSeq != 7 || dialog.AvailableMinUserMessageId != secondView.UserMessageId {
		t.Fatalf("available min mirrors = peer_seq:%d user_message_id:%d, want peer_seq=7 user_message_id=%d",
			dialog.AvailableMinPeerSeq, dialog.AvailableMinUserMessageId, secondView.UserMessageId)
	}
}

func TestGetOutboxReadDateResolvesPublicUserMessageID(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	userID := base + 2171
	peerID := base + 2172
	repo := NewForTest(db, &testIDGenerator{next: base + 21700}, "local-userupdates")
	send := buildOperationApplyInput(t, userID, payload.MessageOperationV1{
		SchemaVersion:      payload.MessageOperationSchemaVersion,
		OperationKind:      payload.OperationKindSendMessage,
		CanonicalMessageID: userID*10 + 1,
		PeerType:           payload.PeerTypeUser,
		PeerID:             peerID,
		PeerSeq:            7,
		FromUserID:         userID,
		ToUserID:           peerID,
		Date:               int32(time.Now().Unix()),
		Out:                true,
		MessageText:        "outgoing",
	}, "send-public-read-date")
	if _, err := repo.ClaimPartitionOwner(ctx, send.PartitionID); err != nil {
		t.Fatalf("ClaimPartitionOwner() error = %v", err)
	}
	if _, err := repo.ApplyUserOperation(ctx, send); err != nil {
		t.Fatalf("ApplyUserOperation(send) error = %v", err)
	}
	view, err := repo.models.UserMessageViewsModel.SelectByUserCanonical(ctx, userID, userID*10+1)
	if err != nil {
		t.Fatalf("SelectByUserCanonical() error = %v", err)
	}
	if view.UserMessageId == view.PeerSeq {
		t.Fatalf("test setup requires public id != peer_seq, got %d", view.UserMessageId)
	}

	readDate := int32(time.Now().Unix())
	readOutbox := buildOperationApplyInput(t, userID, payload.MessageOperationV1{
		SchemaVersion:        payload.MessageOperationSchemaVersion,
		OperationKind:        payload.OperationKindReadHistory,
		PeerType:             payload.PeerTypeUser,
		PeerID:               peerID,
		PeerSeq:              7,
		ReadOutboxMaxPeerSeq: 7,
		FromUserID:           peerID,
		ToUserID:             userID,
		Date:                 readDate,
		Out:                  true,
	}, "read-outbox-public")
	if _, err := repo.ApplyUserOperation(ctx, readOutbox); err != nil {
		t.Fatalf("ApplyUserOperation(read outbox) error = %v", err)
	}

	got, err := repo.GetOutboxReadDate(ctx, OutboxReadDateInput{
		UserID:   userID,
		PeerType: payload.PeerTypeUser,
		PeerID:   peerID,
		MsgID:    int32(view.UserMessageId),
	})
	if err != nil {
		t.Fatalf("GetOutboxReadDate(public id) error = %v", err)
	}
	if got != int64(readDate) {
		t.Fatalf("read date = %d, want %d", got, readDate)
	}

	_, err = repo.GetOutboxReadDate(ctx, OutboxReadDateInput{
		UserID:   userID,
		PeerType: payload.PeerTypeUser,
		PeerID:   peerID,
		MsgID:    int32(view.PeerSeq),
	})
	if !errors.Is(err, userupdates.ErrOutboxReadMessageInvalid) {
		t.Fatalf("GetOutboxReadDate(peer seq input) error = %v, want ErrOutboxReadMessageInvalid", err)
	}
}

func TestGetOutboxReadDateSupportsLegacyPeerSeqInput(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	userID := base + 2181
	peerID := base + 2182
	repo := NewForTest(db, &testIDGenerator{next: base + 21800}, "local-userupdates")
	legacyPayload := mustMarshalMessageEvent(t, payload.MessageEventV1{
		SchemaVersion:      payload.MessageEventSchemaVersionV1,
		EventKind:          payload.EventKindNewMessage,
		CanonicalMessageID: userID*10 + 1,
		MessageID:          9,
		PeerType:           payload.PeerTypeUser,
		PeerID:             peerID,
		FromUserID:         userID,
		ToUserID:           peerID,
		Date:               int32(time.Now().Unix()),
		Out:                true,
		MessageText:        "legacy outgoing",
	})
	if _, _, err := repo.models.UserMessageViewsModel.InsertOrUpdate(ctx, &model.UserMessageViews{
		UserId:             userID,
		PeerType:           payload.PeerTypeUser,
		PeerId:             peerID,
		PeerSeq:            9,
		UserMessageId:      0,
		CanonicalMessageId: userID*10 + 1,
		FromUserId:         userID,
		Outgoing:           true,
		MessageKind:        MessageKindText,
		MessageStatus:      MessageStatusLive,
		Date:               int64(time.Now().Unix()),
		ViewSchemaVersion:  payload.MessageEventSchemaVersionV1,
		ViewPayload:        legacyPayload,
	}); err != nil {
		t.Fatalf("InsertOrUpdate(legacy view) error = %v", err)
	}
	readDate := int64(time.Now().Unix())
	if _, _, err := repo.models.MessageReadOutboxModel.InsertOrUpdate(ctx, &model.MessageReadOutbox{
		UserId:            userID,
		PeerType:          payload.PeerTypeUser,
		PeerId:            peerID,
		ReadUserId:        peerID,
		ReadOutboxMaxId:   9,
		ReadOutboxMaxDate: readDate,
	}); err != nil {
		t.Fatalf("InsertOrUpdate(read outbox) error = %v", err)
	}

	got, err := repo.GetOutboxReadDate(ctx, OutboxReadDateInput{
		UserID:   userID,
		PeerType: payload.PeerTypeUser,
		PeerID:   peerID,
		MsgID:    9,
	})
	if err != nil {
		t.Fatalf("GetOutboxReadDate(legacy peer seq) error = %v", err)
	}
	if got != readDate {
		t.Fatalf("read date = %d, want %d", got, readDate)
	}
}

func TestApplyUpdatePinnedMessageWritesProjectionAndPTSEvent(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	userID := base + 2201
	peerID := base + 2202
	repo := NewForTest(db, &testIDGenerator{next: base + 22000}, "local-userupdates")
	send := buildApplyInput(t, userID, userID, peerID, true, "outgoing")
	if _, err := repo.ClaimPartitionOwner(ctx, send.PartitionID); err != nil {
		t.Fatalf("ClaimPartitionOwner() error = %v", err)
	}
	if _, err := repo.ApplyUserOperation(ctx, send); err != nil {
		t.Fatalf("ApplyUserOperation(send) error = %v", err)
	}
	pin := buildOperationApplyInput(t, userID, payload.MessageOperationV1{
		SchemaVersion:            payload.MessageOperationSchemaVersion,
		OperationKind:            payload.OperationKindUpdatePinnedMessage,
		CanonicalMessageID:       userID*10 + 1,
		PeerType:                 payload.PeerTypeUser,
		PeerID:                   peerID,
		PeerSeq:                  1,
		PinnedPeerSeq:            1,
		PinnedCanonicalMessageID: userID*10 + 1,
		FromUserID:               userID,
		ToUserID:                 peerID,
		Date:                     int32(time.Now().Unix()),
	}, "pin")
	if _, err := repo.ApplyUserOperation(ctx, pin); err != nil {
		t.Fatalf("ApplyUserOperation(pin) error = %v", err)
	}
	row, err := repo.models.UserDialogsModel.SelectByUserPeer(ctx, userID, payload.PeerTypeUser, peerID)
	if err != nil {
		t.Fatalf("SelectByUserPeer() error = %v", err)
	}
	if row.PinnedPeerSeq != 1 || row.PinnedCanonicalMessageId != userID*10+1 || row.LastPts != 2 {
		t.Fatalf("pinned projection = %+v, want pinned peer seq/canonical and last_pts=2", row)
	}
}

func TestApplyEditMessageCarriesPublicUserMessageID(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	userID := base + 2251
	peerID := base + 2252
	repo := NewForTest(db, &testIDGenerator{next: base + 22500}, "local-userupdates")
	send := buildApplyInput(t, userID, userID, peerID, true, "before edit")
	if _, err := repo.ClaimPartitionOwner(ctx, send.PartitionID); err != nil {
		t.Fatalf("ClaimPartitionOwner() error = %v", err)
	}
	sendResult, err := repo.ApplyUserOperation(ctx, send)
	if err != nil {
		t.Fatalf("ApplyUserOperation(send) error = %v", err)
	}
	var sendResponse payload.OperationResponseV2
	if err := json.Unmarshal(sendResult.ResponsePayload, &sendResponse); err != nil {
		t.Fatalf("decode send response: %v", err)
	}
	if sendResponse.UserMessageID == 0 {
		t.Fatalf("send response user_message_id = 0")
	}
	sendView, err := repo.models.UserMessageViewsModel.SelectByUserCanonical(ctx, userID, userID*10+1)
	if err != nil {
		t.Fatalf("SelectByUserCanonical(send) error = %v", err)
	}
	if sendView.ViewSchemaVersion != payload.MessageEventSchemaVersion {
		t.Fatalf("send view schema = %d, want %d", sendView.ViewSchemaVersion, payload.MessageEventSchemaVersion)
	}

	edit := buildOperationApplyInput(t, userID, payload.MessageOperationV1{
		SchemaVersion:      payload.MessageOperationSchemaVersion,
		OperationKind:      payload.OperationKindEditMessage,
		CanonicalMessageID: userID*10 + 1,
		PeerType:           payload.PeerTypeUser,
		PeerID:             peerID,
		PeerSeq:            1,
		FromUserID:         userID,
		ToUserID:           peerID,
		Date:               int32(time.Now().Unix()),
		EditDate:           int32(time.Now().Unix()),
		EditVersion:        2,
		Out:                true,
		MessageText:        "after edit",
	}, "edit")
	editResult, err := repo.ApplyUserOperation(ctx, edit)
	if err != nil {
		t.Fatalf("ApplyUserOperation(edit) error = %v", err)
	}
	var editResponse payload.OperationResponseV2
	if err := json.Unmarshal(editResult.ResponsePayload, &editResponse); err != nil {
		t.Fatalf("decode edit response: %v", err)
	}
	if editResponse.UserMessageID != sendResponse.UserMessageID {
		t.Fatalf("edit response user_message_id = %d, want existing id %d", editResponse.UserMessageID, sendResponse.UserMessageID)
	}
	event, err := repo.models.UserPtsEventsModel.SelectByOperation(ctx, userID, edit.OperationID)
	if err != nil {
		t.Fatalf("SelectByOperation(edit) error = %v", err)
	}
	var editEvent payload.MessageEventV2
	if err := json.Unmarshal(event.EventPayload, &editEvent); err != nil {
		t.Fatalf("decode edit event: %v", err)
	}
	if editEvent.MessageID != sendResponse.UserMessageID {
		t.Fatalf("edit event message_id = %d, want existing public id %d", editEvent.MessageID, sendResponse.UserMessageID)
	}
	editView, err := repo.models.UserMessageViewsModel.SelectByUserCanonical(ctx, userID, userID*10+1)
	if err != nil {
		t.Fatalf("SelectByUserCanonical(edit) error = %v", err)
	}
	if editView.ViewSchemaVersion != payload.MessageEventSchemaVersion {
		t.Fatalf("edit view schema = %d, want %d", editView.ViewSchemaVersion, payload.MessageEventSchemaVersion)
	}
}

func TestApplyUserOperationConflictsOnDifferentPayloadForExistingCanonicalView(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	userID := base + 4201
	peerID := base + 4202
	repo := NewForTest(db, &testIDGenerator{next: base + 42000}, "local-userupdates")
	first := buildOperationApplyInput(t, userID, payload.MessageOperationV1{
		SchemaVersion:      payload.MessageOperationSchemaVersion,
		OperationKind:      payload.OperationKindSendMessage,
		CanonicalMessageID: userID*10 + 1,
		PeerType:           payload.PeerTypeUser,
		PeerID:             peerID,
		PeerSeq:            1,
		FromUserID:         userID,
		ToUserID:           peerID,
		Date:               int32(time.Now().Unix()),
		Out:                true,
		MessageText:        "original",
	}, "canonical-original")
	if _, err := repo.ClaimPartitionOwner(ctx, first.PartitionID); err != nil {
		t.Fatalf("ClaimPartitionOwner() error = %v", err)
	}
	if _, err := repo.ApplyUserOperation(ctx, first); err != nil {
		t.Fatalf("ApplyUserOperation(first) error = %v", err)
	}

	conflict := buildOperationApplyInput(t, userID, payload.MessageOperationV1{
		SchemaVersion:      payload.MessageOperationSchemaVersion,
		OperationKind:      payload.OperationKindSendMessage,
		CanonicalMessageID: userID*10 + 1,
		PeerType:           payload.PeerTypeUser,
		PeerID:             peerID,
		PeerSeq:            2,
		FromUserID:         userID,
		ToUserID:           peerID,
		Date:               int32(time.Now().Unix()),
		Out:                true,
		MessageText:        "different",
	}, "canonical-conflict")
	_, err := repo.ApplyUserOperation(ctx, conflict)
	if !errors.Is(err, userupdates.ErrOperationPayloadConflict) {
		t.Fatalf("ApplyUserOperation(conflict) error = %v, want ErrOperationPayloadConflict", err)
	}

	view, err := repo.models.UserMessageViewsModel.SelectByUserCanonical(ctx, userID, userID*10+1)
	if err != nil {
		t.Fatalf("SelectByUserCanonical() error = %v", err)
	}
	var event payload.MessageEventV2
	if err := json.Unmarshal(view.ViewPayload, &event); err != nil {
		t.Fatalf("decode view payload: %v", err)
	}
	if event.PeerSeq != 1 || event.MessageText != "original" {
		t.Fatalf("existing view was overwritten: %+v", event)
	}
	if _, err := repo.GetOperationResult(ctx, userID, conflict.OperationID); !errors.Is(err, userupdates.ErrOperationTerminal) {
		t.Fatalf("GetOperationResult(conflict) error = %v, want ErrOperationTerminal", err)
	}
}

func TestApplyMarkDialogUnreadLivesWithReadStateOwner(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	userID := base + 2301
	peerID := base + 2302
	repo := NewForTest(db, &testIDGenerator{next: base + 23000}, "local-userupdates")
	send := buildApplyInput(t, userID, userID, peerID, true, "outgoing")
	if _, err := repo.ClaimPartitionOwner(ctx, send.PartitionID); err != nil {
		t.Fatalf("ClaimPartitionOwner() error = %v", err)
	}
	if _, err := repo.ApplyUserOperation(ctx, send); err != nil {
		t.Fatalf("ApplyUserOperation(send) error = %v", err)
	}
	mark := true
	in := buildOperationApplyInput(t, userID, payload.MessageOperationV1{
		SchemaVersion: payload.MessageOperationSchemaVersion,
		OperationKind: payload.OperationKindMarkDialogUnread,
		PeerType:      payload.PeerTypeUser,
		PeerID:        peerID,
		PeerSeq:       1,
		UnreadMark:    &mark,
		FromUserID:    userID,
		ToUserID:      peerID,
		Date:          int32(time.Now().Unix()),
	}, "mark-unread")
	if _, err := repo.ApplyUserOperation(ctx, in); err != nil {
		t.Fatalf("ApplyUserOperation(mark unread) error = %v", err)
	}
	row, err := repo.models.UserDialogsModel.SelectByUserPeer(ctx, userID, payload.PeerTypeUser, peerID)
	if err != nil {
		t.Fatalf("SelectByUserPeer() error = %v", err)
	}
	if !row.UnreadMark || row.LastPts != 2 {
		t.Fatalf("dialog unread mark = %+v, want unread_mark and last_pts=2", row)
	}
}

func TestApplyScheduledMarkerDefaultsFalseUntilScheduledAPIsExist(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	userID := base + 2401
	peerID := base + 2402
	repo := NewForTest(db, &testIDGenerator{next: base + 24000}, "local-userupdates")
	in := buildOperationApplyInput(t, userID, payload.MessageOperationV1{
		SchemaVersion: payload.MessageOperationSchemaVersion,
		OperationKind: payload.OperationKindScheduledMarker,
		PeerType:      payload.PeerTypeUser,
		PeerID:        peerID,
		PeerSeq:       1,
		FromUserID:    userID,
		ToUserID:      peerID,
		Date:          int32(time.Now().Unix()),
	}, "scheduled")
	if _, err := repo.ClaimPartitionOwner(ctx, in.PartitionID); err != nil {
		t.Fatalf("ClaimPartitionOwner() error = %v", err)
	}
	if _, err := repo.ApplyUserOperation(ctx, in); err != nil {
		t.Fatalf("ApplyUserOperation(scheduled) error = %v", err)
	}
	row, err := repo.models.UserDialogsModel.SelectByUserPeer(ctx, userID, payload.PeerTypeUser, peerID)
	if err != nil {
		t.Fatalf("SelectByUserPeer() error = %v", err)
	}
	if row.HasScheduled {
		t.Fatalf("HasScheduled = true, want default false")
	}
}

func TestApplyUserOperationV3PersistsMediaAttrsForwardEvent(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	userID := base + 2601
	repo := NewForTest(db, &testIDGenerator{next: base + 26000}, "local-userupdates")
	op := payload.MessageOperationV3{
		SchemaVersion:      payload.MessageOperationSchemaVersionV3,
		OperationKind:      payload.OperationKindSendMessage,
		CanonicalMessageID: userID*10 + 1,
		PeerType:           payload.PeerTypeUser,
		PeerID:             base + 2602,
		PeerSeq:            1,
		FromUserID:         userID,
		ToUserID:           base + 2602,
		Date:               int32(time.Now().Unix()),
		Out:                true,
		MessageText:        "caption",
		MediaRef:           &payload.MediaRefV1{SchemaVersion: payload.MediaRefSchemaVersionV1, Kind: "photo", ID: 333},
		Attrs:              &payload.MessageAttrsV1{SchemaVersion: payload.MessageAttrsSchemaVersionV1, GroupedID: 444},
		ForwardRef:         &payload.ForwardRefV1{SchemaVersion: payload.ForwardRefSchemaVersionV1, FromUserID: base + 2603, Date: int64(time.Now().Unix())},
	}
	in := buildOperationApplyInputV3(t, userID, op, "operation-v3-media")
	if _, err := repo.ClaimPartitionOwner(ctx, in.PartitionID); err != nil {
		t.Fatalf("ClaimPartitionOwner() error = %v", err)
	}
	result, err := repo.ApplyUserOperation(ctx, in)
	if err != nil {
		t.Fatalf("ApplyUserOperation(V3) error = %v", err)
	}
	if result.UserID != userID || result.PtsCount != 1 {
		t.Fatalf("result = %+v, want user_id=%d pts_count=1", result, userID)
	}
	view, err := repo.models.UserMessageViewsModel.SelectByUserCanonical(ctx, userID, op.CanonicalMessageID)
	if err != nil {
		t.Fatalf("SelectByUserCanonical() error = %v", err)
	}
	if view.ViewSchemaVersion != payload.MessageEventSchemaVersionV3 {
		t.Fatalf("view schema = %d, want V3", view.ViewSchemaVersion)
	}
	var event payload.MessageEventV3
	if err := json.Unmarshal(view.ViewPayload, &event); err != nil {
		t.Fatalf("unmarshal V3 view payload: %v", err)
	}
	if event.MediaRef == nil || event.Attrs == nil || event.ForwardRef == nil {
		t.Fatalf("V3 event lost media/attrs/forward: %+v", event)
	}
	if view.MessageKind != MessageKindMedia {
		t.Fatalf("message kind = %d, want MessageKindMedia", view.MessageKind)
	}
	edit := op
	edit.OperationKind = payload.OperationKindEditMessage
	edit.EditDate = int32(time.Now().Unix())
	edit.EditVersion = 2
	edit.MessageText = "edited caption"
	edit.MediaRef = &payload.MediaRefV1{SchemaVersion: payload.MediaRefSchemaVersionV1, Kind: "document", ID: 444}
	edit.Attrs = &payload.MessageAttrsV1{SchemaVersion: payload.MessageAttrsSchemaVersionV1, GroupedID: 555, Noforwards: true}
	edit.ForwardRef = &payload.ForwardRefV1{SchemaVersion: payload.ForwardRefSchemaVersionV1, FromUserID: base + 2604, Date: int64(time.Now().Unix())}
	editIn := buildOperationApplyInputV3(t, userID, edit, "operation-v3-edit")
	editResult, err := repo.ApplyUserOperation(ctx, editIn)
	if err != nil {
		t.Fatalf("ApplyUserOperation(V3 edit) error = %v", err)
	}
	var editResponse payload.OperationResponseV2
	if err := json.Unmarshal(editResult.ResponsePayload, &editResponse); err != nil {
		t.Fatalf("unmarshal V3 edit response: %v", err)
	}
	if editResponse.UserMessageID != view.UserMessageId {
		t.Fatalf("edit response user_message_id = %d, want %d", editResponse.UserMessageID, view.UserMessageId)
	}
	editView, err := repo.models.UserMessageViewsModel.SelectByUserCanonical(ctx, userID, op.CanonicalMessageID)
	if err != nil {
		t.Fatalf("SelectByUserCanonical(edit) error = %v", err)
	}
	if editView.ViewSchemaVersion != payload.MessageEventSchemaVersionV3 {
		t.Fatalf("edit view schema = %d, want V3", editView.ViewSchemaVersion)
	}
	var editEvent payload.MessageEventV3
	if err := json.Unmarshal(editView.ViewPayload, &editEvent); err != nil {
		t.Fatalf("unmarshal V3 edit view payload: %v", err)
	}
	if editEvent.SchemaVersion != payload.MessageEventSchemaVersionV3 || editEvent.MessageID != view.UserMessageId || editEvent.MediaRef == nil || editEvent.Attrs == nil || editEvent.ForwardRef == nil {
		t.Fatalf("V3 edit event lost schema/media/attrs/forward: %+v", editEvent)
	}
}

func TestGetDifferenceLegacyMessageHydrationRequiresExactEventPeerSeq(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	userID := base + 2501
	peerID := base + 2502
	repo := NewForTest(db, &testIDGenerator{next: base + 25000}, "local-userupdates")

	viewPayload := []byte(`{"schema_version":1}`)
	if _, _, err := repo.models.UserMessageViewsModel.InsertOrUpdate(ctx, &model.UserMessageViews{
		UserId:             userID,
		PeerType:           payload.PeerTypeUser,
		PeerId:             peerID,
		PeerSeq:            1,
		UserMessageId:      101,
		CanonicalMessageId: userID*10 + 1,
		FromUserId:         peerID,
		Outgoing:           false,
		MessageKind:        MessageKindText,
		MessageStatus:      MessageStatusLive,
		Date:               time.Now().Unix(),
		ViewSchemaVersion:  1,
		ViewPayload:        viewPayload,
	}); err != nil {
		t.Fatalf("insert message view: %v", err)
	}
	legacy := payload.MessageEventV1{
		SchemaVersion:      payload.MessageEventSchemaVersionV1,
		EventKind:          payload.EventKindNewMessage,
		CanonicalMessageID: userID*10 + 2,
		MessageID:          2,
		PeerType:           payload.PeerTypeUser,
		PeerID:             peerID,
		FromUserID:         peerID,
		ToUserID:           userID,
		Date:               int32(time.Now().Unix()),
		MessageText:        "missing exact row",
	}
	body, err := json.Marshal(legacy)
	if err != nil {
		t.Fatalf("marshal legacy event: %v", err)
	}
	if _, _, err := repo.models.UserPtsEventsModel.Insert(ctx, &model.UserPtsEvents{
		UserId:             userID,
		Pts:                1,
		PtsCount:           1,
		OperationId:        fmt.Sprintf("legacy-exact-%d", base),
		OpType:             OpTypeSendMessage,
		EventType:          EventTypeNewMessage,
		PeerType:           payload.PeerTypeUser,
		PeerId:             peerID,
		CanonicalMessageId: userID*10 + 2,
		PeerSeq:            2,
		ActorUserId:        peerID,
		EventSchemaVersion: payload.MessageEventSchemaVersionV1,
		EventCodec:         PayloadCodecJSON,
		EventPayload:       body,
		EventPayloadHash:   payload.HashBytes(body),
	}); err != nil {
		t.Fatalf("insert legacy event: %v", err)
	}

	_, err = repo.GetDifference(ctx, GetDifferenceInput{UserID: userID, Pts: 0, Limit: 10})
	if !errors.Is(err, userupdates.ErrUserupdatesStorage) {
		t.Fatalf("GetDifference() error = %v, want ErrUserupdatesStorage for missing exact peer_seq", err)
	}
}

func TestGetDifferenceHydratesLegacyPinnedEventPublicMessageID(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	userID := base + 2521
	peerID := base + 2522
	canonicalID := userID*10 + 1
	repo := NewForTest(db, &testIDGenerator{next: base + 25200}, "local-userupdates")

	legacy := payload.MessageEventV1{
		SchemaVersion:      payload.MessageEventSchemaVersionV1,
		EventKind:          payload.OperationKindUpdatePinnedMessage,
		CanonicalMessageID: canonicalID,
		MessageID:          9,
		PeerType:           payload.PeerTypeUser,
		PeerID:             peerID,
		FromUserID:         userID,
		ToUserID:           peerID,
		Date:               int32(time.Now().Unix()),
		Out:                true,
	}
	body := mustMarshalMessageEvent(t, legacy)
	if _, _, err := repo.models.UserMessageViewsModel.InsertOrUpdate(ctx, &model.UserMessageViews{
		UserId:             userID,
		PeerType:           payload.PeerTypeUser,
		PeerId:             peerID,
		PeerSeq:            9,
		UserMessageId:      101,
		CanonicalMessageId: canonicalID,
		FromUserId:         userID,
		Outgoing:           true,
		MessageKind:        MessageKindText,
		MessageStatus:      MessageStatusLive,
		Date:               int64(legacy.Date),
		ViewSchemaVersion:  payload.MessageEventSchemaVersionV1,
		ViewPayload:        body,
	}); err != nil {
		t.Fatalf("insert message view: %v", err)
	}
	if _, _, err := repo.models.UserPtsEventsModel.Insert(ctx, &model.UserPtsEvents{
		UserId:             userID,
		Pts:                1,
		PtsCount:           1,
		OperationId:        fmt.Sprintf("legacy-pinned-%d", base),
		OpType:             OpTypeSendMessage,
		EventType:          EventTypeUpdatePinnedMessage,
		PeerType:           payload.PeerTypeUser,
		PeerId:             peerID,
		CanonicalMessageId: canonicalID,
		PeerSeq:            9,
		ActorUserId:        userID,
		EventSchemaVersion: payload.MessageEventSchemaVersionV1,
		EventCodec:         PayloadCodecJSON,
		EventPayload:       body,
		EventPayloadHash:   payload.HashBytes(body),
	}); err != nil {
		t.Fatalf("insert legacy pinned event: %v", err)
	}

	diff, err := repo.GetDifference(ctx, GetDifferenceInput{UserID: userID, Pts: 0, Limit: 10})
	if err != nil {
		t.Fatalf("GetDifference() error = %v", err)
	}
	if len(diff.Events) != 1 {
		t.Fatalf("events len = %d, want 1", len(diff.Events))
	}
	got := diff.Events[0]
	if got.EventSchemaVersion != payload.MessageEventSchemaVersion {
		t.Fatalf("event schema = %d, want %d", got.EventSchemaVersion, payload.MessageEventSchemaVersion)
	}
	var hydrated payload.MessageEventV2
	if err := json.Unmarshal(got.EventPayload, &hydrated); err != nil {
		t.Fatalf("decode hydrated event: %v", err)
	}
	if hydrated.EventKind != payload.OperationKindUpdatePinnedMessage {
		t.Fatalf("event kind = %q, want %q", hydrated.EventKind, payload.OperationKindUpdatePinnedMessage)
	}
	if hydrated.PeerSeq != 9 {
		t.Fatalf("peer_seq = %d, want 9", hydrated.PeerSeq)
	}
	if hydrated.MessageID != 101 {
		t.Fatalf("message_id = %d, want public user_message_id 101", hydrated.MessageID)
	}
	if hydrated.PinnedUserMessageID != 101 {
		t.Fatalf("pinned_user_message_id = %d, want public user_message_id 101", hydrated.PinnedUserMessageID)
	}
}

func openIntegrationDB(t *testing.T) *sqlx.DB {
	t.Helper()
	if testing.Short() {
		t.Skip("integration test skipped in short mode")
	}
	dsn := os.Getenv("TEAMGRAM_TEST_MYSQL_DSN")
	explicit := dsn != ""
	if dsn == "" {
		dsn = "root:@tcp(127.0.0.1:3306)/teamgooo?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
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

func cleanupDeleteMessagesIntegrationRows(t *testing.T, ctx context.Context, db *sqlx.DB, userID int64) {
	t.Helper()
	statements := []string{
		"DELETE FROM push_task_outbox WHERE user_id = ?",
		"DELETE FROM user_pts_events WHERE user_id = ?",
		"DELETE FROM user_operation_results WHERE user_id = ?",
		"DELETE FROM user_message_views WHERE user_id = ?",
		"DELETE FROM user_message_sequences WHERE user_id = ?",
		"DELETE FROM user_dialogs WHERE user_id = ?",
		"DELETE FROM user_pts_state WHERE user_id = ?",
	}
	for _, statement := range statements {
		if _, err := db.Exec(ctx, statement, userID); err != nil {
			t.Fatalf("cleanup %q: %v", statement, err)
		}
	}
}

func mysqlTestTime(t time.Time) int64 {
	return t.UTC().Unix()
}

func mysqlTestTimeValue(t time.Time) int64 {
	return t.UTC().Unix()
}

func normalizeDBTestTime(t *testing.T, value any) int64 {
	t.Helper()
	switch v := value.(type) {
	case time.Time:
		return mysqlTestTime(v)
	case int64:
		return v
	case int32:
		return int64(v)
	case int:
		return int64(v)
	default:
		t.Fatalf("unsupported DB time value %T", value)
	}
	return 0
}

func insertTestPushTask(t *testing.T, ctx context.Context, db *sqlx.DB, taskID int64, status int32, availableAt time.Time) {
	t.Helper()
	_, err := db.Exec(ctx, "DELETE FROM push_task_outbox WHERE task_id = ?", taskID)
	if err != nil {
		t.Fatalf("clean test push task: %v", err)
	}

	_, err = db.Exec(ctx, `
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

func insertTestDeliveryFailedOperation(t *testing.T, ctx context.Context, db *sqlx.DB, failedID int64, bucketID int32, status int32, failedAt time.Time) {
	t.Helper()
	_, err := db.Exec(ctx, `
INSERT INTO delivery_failed_operations
	(failed_id, user_id, operation_id, op_type, bucket_id, kafka_topic,
	 kafka_partition, kafka_offset, payload_schema_version, payload_hash,
	 failure_category, failure_code, failure_message, retry_count, status,
	 failed_at)
VALUES
	(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		failedID,
		failedID+10_000,
		fmt.Sprintf("failed-op-%d", failedID),
		int32(1),
		bucketID,
		"test-topic",
		int32(failedID%16),
		failedID,
		int32(1),
		[]byte("01234567890123456789012345678901"),
		int32(1),
		"test_failure",
		"test failure",
		int32(0),
		status,
		mysqlTestTime(failedAt),
	)
	if err != nil {
		t.Fatalf("insert failed operation: %v", err)
	}
}

func insertTestOperationResult(t *testing.T, ctx context.Context, db *sqlx.DB, userID int64, operationID string, status int32, completedAt time.Time) {
	t.Helper()
	_, err := db.Exec(ctx, `
INSERT INTO user_operation_results
	(user_id, operation_id, op_type, status, pts, pts_count, payload_hash,
	 response_schema_version, response_codec, response_payload,
	 response_payload_hash, terminal_error_code, completed_at)
VALUES
	(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		userID,
		operationID,
		int32(1),
		status,
		int64(1),
		int32(1),
		[]byte("01234567890123456789012345678901"),
		int32(1),
		PayloadCodecJSON,
		[]byte(`{"test":true}`),
		[]byte("12345678901234567890123456789012"),
		"",
		mysqlTestTime(completedAt),
	)
	if err != nil {
		t.Fatalf("insert operation result: %v", err)
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

func buildOperationApplyInput(t *testing.T, userID int64, op payload.MessageOperationV1, suffix string) ApplyUserOperationInput {
	t.Helper()
	route := payload.RouteUser(userID)
	body, err := json.Marshal(op)
	if err != nil {
		t.Fatalf("marshal operation: %v", err)
	}
	return ApplyUserOperationInput{
		UserID:       userID,
		OperationID:  fmt.Sprintf("v1:%s:user:%d:%s:%d", op.OperationKind, userID, suffix, time.Now().UnixNano()),
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

func buildOperationApplyInputV3(t *testing.T, userID int64, op payload.MessageOperationV3, suffix string) ApplyUserOperationInput {
	t.Helper()
	route := payload.RouteUser(userID)
	body, err := json.Marshal(op)
	if err != nil {
		t.Fatalf("marshal V3 operation: %v", err)
	}
	return ApplyUserOperationInput{
		UserID:       userID,
		OperationID:  fmt.Sprintf("v3:%s:user:%d:%s:%d", op.OperationKind, userID, suffix, time.Now().UnixNano()),
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

func mustMarshalMessageEvent(t *testing.T, event payload.MessageEventV1) []byte {
	t.Helper()
	body, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("marshal message event: %v", err)
	}
	return body
}

func buildAffectedOutbox(t *testing.T, in ApplyUserOperationInput, targetUserID int64, suffix string) AffectedOutbox {
	t.Helper()
	targetRoute := payload.RouteUser(targetUserID)
	body := []byte(fmt.Sprintf(`{"schema_version":1,"suffix":%q}`, suffix))
	return AffectedOutbox{
		RequesterUserID:   in.UserID,
		TargetUserID:      targetUserID,
		TargetBucketID:    int32(targetRoute.BucketID),
		TargetPartitionID: int32(targetRoute.ReceiverPartitionID),
		OperationID:       "affected-" + in.OperationID,
		OpType:            in.OpType,
		OperationKind:     payload.OperationKindSendMessage,
		PeerType:          in.PeerType,
		PeerID:            in.PeerID,
		PayloadCodec:      PayloadCodecJSON,
		Payload:           body,
		PayloadHash:       payload.HashBytes(body),
		DeliveryPolicy:    DeliveryPolicyDurableAsync,
	}
}
