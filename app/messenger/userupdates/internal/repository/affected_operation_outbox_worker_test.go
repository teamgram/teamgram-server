package repository

import (
	"context"
	"errors"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
)

type fakeAffectedOutboxStore struct {
	rows              []AffectedOutboxRow
	claim             bool
	applyErr          error
	applied           []ApplyUserOperationInput
	completed         []fakeCompletedMark
	retryable         []fakeRetryableMark
	terminal          []fakeTerminalMark
	resetCalls        int
	resetCount        int64
	claimCalls        int
	lastClaimDeadline time.Time
	applyCalls        atomic.Int32
}

type fakeCompletedMark struct {
	outboxID int64
	deadline time.Time
}

type fakeRetryableMark struct {
	outboxID        int64
	deadline        time.Time
	code            string
	message         string
	nextAvailableAt time.Time
}

type fakeTerminalMark struct {
	outboxID int64
	deadline time.Time
	code     string
	message  string
}

func (s *fakeAffectedOutboxStore) ListPendingAffectedOutboxes(ctx context.Context, now time.Time, limit int32) ([]AffectedOutboxRow, error) {
	return s.rows, nil
}

func (s *fakeAffectedOutboxStore) TryMarkAffectedOutboxProcessing(ctx context.Context, outboxID int64, now time.Time, deadline time.Time) (bool, error) {
	s.claimCalls++
	s.lastClaimDeadline = deadline
	return s.claim, nil
}

func (s *fakeAffectedOutboxStore) ApplyUserOperation(ctx context.Context, in ApplyUserOperationInput) (*ApplyUserOperationResult, error) {
	s.applyCalls.Add(1)
	s.applied = append(s.applied, in)
	if s.applyErr != nil {
		return nil, s.applyErr
	}
	return &ApplyUserOperationResult{UserID: in.UserID, OperationID: in.OperationID, Pts: 7, PtsCount: 1}, nil
}

func (s *fakeAffectedOutboxStore) MarkAffectedOutboxCompleted(ctx context.Context, outboxID int64, processingDeadline time.Time) error {
	s.completed = append(s.completed, fakeCompletedMark{outboxID: outboxID, deadline: processingDeadline})
	return nil
}

func (s *fakeAffectedOutboxStore) MarkAffectedOutboxRetryable(ctx context.Context, outboxID int64, processingDeadline time.Time, code string, message string, nextAvailableAt time.Time) error {
	s.retryable = append(s.retryable, fakeRetryableMark{outboxID: outboxID, deadline: processingDeadline, code: code, message: message, nextAvailableAt: nextAvailableAt})
	return nil
}

func (s *fakeAffectedOutboxStore) MarkAffectedOutboxFailedTerminal(ctx context.Context, outboxID int64, processingDeadline time.Time, code string, message string) error {
	s.terminal = append(s.terminal, fakeTerminalMark{outboxID: outboxID, deadline: processingDeadline, code: code, message: message})
	return nil
}

func (s *fakeAffectedOutboxStore) ResetExpiredAffectedOutboxes(ctx context.Context, now time.Time, limit int32) (int64, error) {
	s.resetCalls++
	return s.resetCount, nil
}

func testAffectedOutboxRow(id int64) AffectedOutboxRow {
	return AffectedOutboxRow{
		OutboxID:     id,
		UserID:       1001,
		OperationID:  "op-peer-read",
		OpType:       OpTypeSendMessage,
		PeerType:     1,
		PeerID:       2002,
		PayloadCodec: PayloadCodecJSON,
		Payload:      []byte(`{"schema_version":1,"operation_kind":"readHistory"}`),
		PayloadHash:  []byte("hash"),
		BucketID:     10,
		PartitionID:  20,
	}
}

func TestAffectedOutboxWorkerAppliesPendingRow(t *testing.T) {
	store := &fakeAffectedOutboxStore{rows: []AffectedOutboxRow{testAffectedOutboxRow(1)}, claim: true}
	worker := NewAffectedOutboxWorker(store, AffectedOutboxWorkerOptions{Interval: time.Second, BatchSize: 10, ProcessingTimeout: time.Minute})

	worker.drain(context.Background())

	if got := store.applyCalls.Load(); got != 1 {
		t.Fatalf("apply calls = %d, want 1", got)
	}
	if len(store.applied) != 1 {
		t.Fatalf("applied = %+v", store.applied)
	}
	applied := store.applied[0]
	if applied.UserID != 1001 || applied.OperationID != "op-peer-read" || applied.PartitionID != 20 {
		t.Fatalf("applied input = %+v", applied)
	}
	if len(store.completed) != 1 || store.completed[0].outboxID != 1 {
		t.Fatalf("completed = %+v", store.completed)
	}
	if store.completed[0].deadline.IsZero() {
		t.Fatal("completed mark missing processing deadline fence")
	}
	if !store.completed[0].deadline.Equal(store.lastClaimDeadline) {
		t.Fatalf("completed deadline = %s, want claim deadline %s", store.completed[0].deadline, store.lastClaimDeadline)
	}
	if len(store.retryable) != 0 || len(store.terminal) != 0 {
		t.Fatalf("unexpected retryable=%+v terminal=%+v", store.retryable, store.terminal)
	}
}

func TestAffectedOutboxWorkerSkipsUnclaimedRow(t *testing.T) {
	store := &fakeAffectedOutboxStore{rows: []AffectedOutboxRow{testAffectedOutboxRow(2)}, claim: false}
	worker := NewAffectedOutboxWorker(store, AffectedOutboxWorkerOptions{Interval: time.Second, BatchSize: 10, ProcessingTimeout: time.Minute})

	worker.drain(context.Background())

	if got := store.applyCalls.Load(); got != 0 {
		t.Fatalf("apply calls = %d, want 0", got)
	}
	if len(store.completed) != 0 || len(store.retryable) != 0 || len(store.terminal) != 0 {
		t.Fatalf("unexpected completed=%+v retryable=%+v terminal=%+v", store.completed, store.retryable, store.terminal)
	}
}

func TestAffectedOutboxWorkerMarksRetryableOnInfrastructureError(t *testing.T) {
	store := &fakeAffectedOutboxStore{
		rows:     []AffectedOutboxRow{testAffectedOutboxRow(3)},
		claim:    true,
		applyErr: errors.New("mysql timeout " + strings.Repeat("x", 512)),
	}
	worker := NewAffectedOutboxWorker(store, AffectedOutboxWorkerOptions{Interval: 50 * time.Millisecond, BatchSize: 10, ProcessingTimeout: time.Minute})

	before := time.Now().UTC()
	worker.drain(context.Background())

	if len(store.retryable) != 1 {
		t.Fatalf("retryable = %+v, want one mark", store.retryable)
	}
	if store.retryable[0].outboxID != 3 || store.retryable[0].code != "affected_apply_retryable" {
		t.Fatalf("retryable mark = %+v", store.retryable[0])
	}
	if store.retryable[0].deadline.IsZero() {
		t.Fatal("retryable mark missing processing deadline fence")
	}
	if !store.retryable[0].nextAvailableAt.After(before) {
		t.Fatalf("next available = %s, want after %s", store.retryable[0].nextAvailableAt, before)
	}
	if len(store.retryable[0].message) > affectedOutboxWorkerMaxErrorMessage {
		t.Fatalf("retry message length = %d, want <= %d", len(store.retryable[0].message), affectedOutboxWorkerMaxErrorMessage)
	}
	if len(store.completed) != 0 || len(store.terminal) != 0 {
		t.Fatalf("unexpected completed=%+v terminal=%+v", store.completed, store.terminal)
	}
}

func TestAffectedOutboxWorkerMarksTerminalOnPayloadConflict(t *testing.T) {
	store := &fakeAffectedOutboxStore{
		rows:     []AffectedOutboxRow{testAffectedOutboxRow(4)},
		claim:    true,
		applyErr: userupdates.ErrOperationPayloadConflict,
	}
	worker := NewAffectedOutboxWorker(store, AffectedOutboxWorkerOptions{Interval: time.Second, BatchSize: 10, ProcessingTimeout: time.Minute})

	worker.drain(context.Background())

	if len(store.terminal) != 1 {
		t.Fatalf("terminal = %+v, want one mark", store.terminal)
	}
	if store.terminal[0].outboxID != 4 || store.terminal[0].code != "affected_apply_terminal" {
		t.Fatalf("terminal mark = %+v", store.terminal[0])
	}
	if store.terminal[0].deadline.IsZero() {
		t.Fatal("terminal mark missing processing deadline fence")
	}
	if len(store.completed) != 0 || len(store.retryable) != 0 {
		t.Fatalf("unexpected completed=%+v retryable=%+v", store.completed, store.retryable)
	}
}

func TestAffectedOutboxWorkerResetsExpiredProcessingRows(t *testing.T) {
	store := &fakeAffectedOutboxStore{resetCount: 2}
	worker := NewAffectedOutboxWorker(store, AffectedOutboxWorkerOptions{Interval: time.Second, BatchSize: 10, ProcessingTimeout: time.Minute})

	worker.drain(context.Background())

	if store.resetCalls != 1 {
		t.Fatalf("reset calls = %d, want 1", store.resetCalls)
	}
}

type blockingAffectedOutboxStore struct {
	started chan struct{}
	release chan struct{}
	done    atomic.Bool
}

func (s *blockingAffectedOutboxStore) ListPendingAffectedOutboxes(ctx context.Context, now time.Time, limit int32) ([]AffectedOutboxRow, error) {
	close(s.started)
	<-s.release
	s.done.Store(true)
	return nil, nil
}
func (s *blockingAffectedOutboxStore) TryMarkAffectedOutboxProcessing(context.Context, int64, time.Time, time.Time) (bool, error) {
	return false, nil
}
func (s *blockingAffectedOutboxStore) ApplyUserOperation(context.Context, ApplyUserOperationInput) (*ApplyUserOperationResult, error) {
	return nil, nil
}
func (s *blockingAffectedOutboxStore) MarkAffectedOutboxCompleted(context.Context, int64, time.Time) error {
	return nil
}
func (s *blockingAffectedOutboxStore) MarkAffectedOutboxRetryable(context.Context, int64, time.Time, string, string, time.Time) error {
	return nil
}
func (s *blockingAffectedOutboxStore) MarkAffectedOutboxFailedTerminal(context.Context, int64, time.Time, string, string) error {
	return nil
}
func (s *blockingAffectedOutboxStore) ResetExpiredAffectedOutboxes(context.Context, time.Time, int32) (int64, error) {
	return 0, nil
}

func TestAffectedOutboxWorkerWaitsForDrain(t *testing.T) {
	store := &blockingAffectedOutboxStore{
		started: make(chan struct{}),
		release: make(chan struct{}),
	}
	worker := NewAffectedOutboxWorker(store, AffectedOutboxWorkerOptions{Interval: time.Hour, BatchSize: 10, ProcessingTimeout: time.Minute})

	go worker.runDrain(context.Background())
	<-store.started

	done := make(chan struct{})
	go func() {
		worker.Wait()
		close(done)
	}()

	select {
	case <-done:
		t.Fatal("Wait returned while drain was still running")
	case <-time.After(20 * time.Millisecond):
	}

	close(store.release)

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("Wait did not return after drain completed")
	}
	if !store.done.Load() {
		t.Fatal("drain did not complete")
	}
}
