package repository

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository/model"
)

type fakeDialogOutboxUserupdatesClient struct {
	mu          sync.Mutex
	authSeqReqs []*userupdates.TLUserupdatesAppendDialogAuthSeqSideEffect
	ptsReqs     []*userupdates.TLUserupdatesAppendDialogPtsSideEffect
	authSeqErr  error
	ptsErr      error
}

func (f *fakeDialogOutboxUserupdatesClient) UserupdatesAppendDialogAuthSeqSideEffect(ctx context.Context, in *userupdates.TLUserupdatesAppendDialogAuthSeqSideEffect) (*userupdates.UserAuthSeqAppendResult, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.authSeqReqs = append(f.authSeqReqs, in)
	if f.authSeqErr != nil {
		return nil, f.authSeqErr
	}
	return userupdates.MakeTLUserAuthSeqAppendResult(&userupdates.TLUserAuthSeqAppendResult{
		UserId:      in.UserId,
		OperationId: in.OperationId,
		Seq:         901,
		Date:        1700000001,
	}), nil
}

func (f *fakeDialogOutboxUserupdatesClient) UserupdatesAppendDialogPtsSideEffect(ctx context.Context, in *userupdates.TLUserupdatesAppendDialogPtsSideEffect) (*userupdates.UserPtsAppendResult, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.ptsReqs = append(f.ptsReqs, in)
	if f.ptsErr != nil {
		return nil, f.ptsErr
	}
	return userupdates.MakeTLUserPtsAppendResult(&userupdates.TLUserPtsAppendResult{
		UserId:      in.UserId,
		OperationId: in.OperationId,
		Pts:         801,
		PtsCount:    1,
	}), nil
}

func TestDialogAuthSeqOutboxWorkerPublishesAndMarksPublished(t *testing.T) {
	ctx := context.Background()
	repo := newDialogOutboxWorkerTestRepo(t)
	base := time.Now().UnixNano()
	row := newAuthSeqOutboxRow(base, fmt.Sprintf("auth-publish-%d", base))
	if _, _, err := repo.model.DialogAuthSeqOutboxModel.Insert(ctx, row); err != nil {
		t.Fatalf("insert auth seq outbox: %v", err)
	}

	client := &fakeDialogOutboxUserupdatesClient{}
	worker := NewDialogAuthSeqOutboxWorker(repo, client, testOutboxOptions("auth-publish"))
	if err := worker.Drain(ctx); err != nil {
		t.Fatalf("Drain() error = %v", err)
	}

	got, err := repo.model.DialogAuthSeqOutboxModel.FindOne(ctx, row.OutboxId)
	if err != nil {
		t.Fatalf("FindOne() error = %v", err)
	}
	if got.Status != OutboxStatusPublished || got.PublishedSeq != 901 || got.PublishedDate != 1700000001 {
		t.Fatalf("published state = status:%d seq:%d date:%d, want status:%d seq:901 date:1700000001", got.Status, got.PublishedSeq, got.PublishedDate, OutboxStatusPublished)
	}
	if len(client.authSeqReqs) != 1 {
		t.Fatalf("auth seq requests = %d, want 1", len(client.authSeqReqs))
	}
	if client.authSeqReqs[0].SourcePermAuthKeyId != row.SourcePermAuthKeyId {
		t.Fatalf("SourcePermAuthKeyId = %d, want %d", client.authSeqReqs[0].SourcePermAuthKeyId, row.SourcePermAuthKeyId)
	}
}

func TestDialogAuthSeqOutboxWorkerBlocksMissingSourceForNotSourcePolicy(t *testing.T) {
	ctx := context.Background()
	repo := newDialogOutboxWorkerTestRepo(t)
	base := time.Now().UnixNano()
	row := newAuthSeqOutboxRow(base, fmt.Sprintf("auth-block-%d", base))
	row.SourcePermAuthKeyId = 0
	if _, _, err := repo.model.DialogAuthSeqOutboxModel.Insert(ctx, row); err != nil {
		t.Fatalf("insert auth seq outbox: %v", err)
	}

	client := &fakeDialogOutboxUserupdatesClient{}
	worker := NewDialogAuthSeqOutboxWorker(repo, client, testOutboxOptions("auth-block"))
	if err := worker.Drain(ctx); err != nil {
		t.Fatalf("Drain() error = %v", err)
	}

	got, err := repo.model.DialogAuthSeqOutboxModel.FindOne(ctx, row.OutboxId)
	if err != nil {
		t.Fatalf("FindOne() error = %v", err)
	}
	if got.Status != OutboxStatusBlocked || got.LastErrorKind != "invalid_target_auth_policy" {
		t.Fatalf("blocked state = status:%d kind:%q, want status:%d kind invalid_target_auth_policy", got.Status, got.LastErrorKind, OutboxStatusBlocked)
	}
	if len(client.authSeqReqs) != 0 {
		t.Fatalf("auth seq requests = %d, want 0", len(client.authSeqReqs))
	}
}

func TestDialogPublicUpdateOutboxWorkerPublishesPTS(t *testing.T) {
	ctx := context.Background()
	repo := newDialogOutboxWorkerTestRepo(t)
	base := time.Now().UnixNano()
	row := newPublicUpdateOutboxRow(base, fmt.Sprintf("public-pts-%d", base), DeliveryPathUserupdatesPTS)
	if _, _, err := repo.model.DialogPublicUpdateOutboxModel.Insert(ctx, row); err != nil {
		t.Fatalf("insert public update outbox: %v", err)
	}

	client := &fakeDialogOutboxUserupdatesClient{}
	worker := NewDialogPublicUpdateOutboxWorker(repo, client, testOutboxOptions("public-pts"))
	if err := worker.Drain(ctx); err != nil {
		t.Fatalf("Drain() error = %v", err)
	}

	got, err := repo.model.DialogPublicUpdateOutboxModel.FindOne(ctx, row.OutboxId)
	if err != nil {
		t.Fatalf("FindOne() error = %v", err)
	}
	if got.Status != OutboxStatusPublished || got.PublishedPts != 801 || got.PublishedPtsCount != 1 {
		t.Fatalf("published pts state = status:%d pts:%d count:%d, want status:%d pts:801 count:1", got.Status, got.PublishedPts, got.PublishedPtsCount, OutboxStatusPublished)
	}
	if len(client.ptsReqs) != 1 || len(client.authSeqReqs) != 0 {
		t.Fatalf("requests = pts:%d auth:%d, want pts:1 auth:0", len(client.ptsReqs), len(client.authSeqReqs))
	}
}

func TestDialogPublicUpdateOutboxWorkerPublishesAuthSeq(t *testing.T) {
	ctx := context.Background()
	repo := newDialogOutboxWorkerTestRepo(t)
	base := time.Now().UnixNano()
	row := newPublicUpdateOutboxRow(base, fmt.Sprintf("public-auth-%d", base), DeliveryPathUserupdatesAuthSeq)
	if _, _, err := repo.model.DialogPublicUpdateOutboxModel.Insert(ctx, row); err != nil {
		t.Fatalf("insert public update outbox: %v", err)
	}

	client := &fakeDialogOutboxUserupdatesClient{}
	worker := NewDialogPublicUpdateOutboxWorker(repo, client, testOutboxOptions("public-auth"))
	if err := worker.Drain(ctx); err != nil {
		t.Fatalf("Drain() error = %v", err)
	}

	got, err := repo.model.DialogPublicUpdateOutboxModel.FindOne(ctx, row.OutboxId)
	if err != nil {
		t.Fatalf("FindOne() error = %v", err)
	}
	if got.Status != OutboxStatusPublished || got.PublishedSeq != 901 || got.PublishedDate != 1700000001 {
		t.Fatalf("published auth seq state = status:%d seq:%d date:%d, want status:%d seq:901 date:1700000001", got.Status, got.PublishedSeq, got.PublishedDate, OutboxStatusPublished)
	}
	if len(client.authSeqReqs) != 1 || len(client.ptsReqs) != 0 {
		t.Fatalf("requests = auth:%d pts:%d, want auth:1 pts:0", len(client.authSeqReqs), len(client.ptsReqs))
	}
}

func TestDialogOutboxWorkerRetryAndBlockedState(t *testing.T) {
	ctx := context.Background()
	repo := newDialogOutboxWorkerTestRepo(t)
	base := time.Now().UnixNano()
	retryRow := newPublicUpdateOutboxRow(base, fmt.Sprintf("public-retry-%d", base), DeliveryPathUserupdatesPTS)
	blockRow := newPublicUpdateOutboxRow(base+1, fmt.Sprintf("public-block-%d", base), DeliveryPathUserupdatesPTS)
	blockRow.OutboxId = retryRow.OutboxId + 1
	blockRow.AttemptCount = OutboxWorkerBlockedAttempts - 1
	if _, _, err := repo.model.DialogPublicUpdateOutboxModel.Insert(ctx, retryRow); err != nil {
		t.Fatalf("insert retry public update outbox: %v", err)
	}
	if _, _, err := repo.model.DialogPublicUpdateOutboxModel.Insert(ctx, blockRow); err != nil {
		t.Fatalf("insert block public update outbox: %v", err)
	}

	client := &fakeDialogOutboxUserupdatesClient{ptsErr: errors.New("ledger unavailable")}
	options := testOutboxOptions("public-retry")
	options.BatchSize = 2
	worker := NewDialogPublicUpdateOutboxWorker(repo, client, options)
	if err := worker.Drain(ctx); err != nil {
		t.Fatalf("Drain() error = %v", err)
	}

	retryGot, err := repo.model.DialogPublicUpdateOutboxModel.FindOne(ctx, retryRow.OutboxId)
	if err != nil {
		t.Fatalf("FindOne(retry) error = %v", err)
	}
	if retryGot.Status != OutboxStatusFailedRetryable || retryGot.AttemptCount != 1 || retryGot.LastErrorKind != "userupdates_append_pts" {
		t.Fatalf("retry state = status:%d attempts:%d kind:%q, want status:%d attempts:1 kind userupdates_append_pts", retryGot.Status, retryGot.AttemptCount, retryGot.LastErrorKind, OutboxStatusFailedRetryable)
	}

	blockGot, err := repo.model.DialogPublicUpdateOutboxModel.FindOne(ctx, blockRow.OutboxId)
	if err != nil {
		t.Fatalf("FindOne(block) error = %v", err)
	}
	if blockGot.Status != OutboxStatusBlocked || blockGot.LastErrorKind != "userupdates_append_pts" {
		t.Fatalf("blocked state = status:%d kind:%q, want status:%d kind userupdates_append_pts", blockGot.Status, blockGot.LastErrorKind, OutboxStatusBlocked)
	}
}

func TestOutboxRetryDelayIsCapped(t *testing.T) {
	if got := nextOutboxRetry(99); got != time.Duration(OutboxWorkerMaxRetryDelay)*time.Second {
		t.Fatalf("nextOutboxRetry(99) = %v, want %ds", got, OutboxWorkerMaxRetryDelay)
	}
	if got := nextOutboxRetry(0); got != time.Duration(InitialRetryDelaySeconds)*time.Second {
		t.Fatalf("nextOutboxRetry(0) = %v, want initial delay", got)
	}
}

func TestOutboxBlockedAfterAge(t *testing.T) {
	ctx := context.Background()
	repo := newDialogOutboxWorkerTestRepo(t)
	base := time.Now().UnixNano()
	row := newPublicUpdateOutboxRow(base, fmt.Sprintf("public-age-block-%d", base), DeliveryPathUserupdatesPTS)
	old := time.Now().UTC().Add(-time.Duration(OutboxWorkerBlockedAgeSeconds+1) * time.Second)
	row.NextRetryAt = mysqlDateTimeForBind(old)
	if _, _, err := repo.model.DialogPublicUpdateOutboxModel.Insert(ctx, row); err != nil {
		t.Fatalf("insert public update outbox: %v", err)
	}

	worker := NewDialogPublicUpdateOutboxWorker(repo, nil, testOutboxOptions("public-age-block"))
	if err := worker.markPublicUpdateFailure(ctx, *row, "userupdates_append_pts", errors.New("ledger unavailable")); err != nil {
		t.Fatalf("markPublicUpdateFailure() error = %v", err)
	}

	got, err := repo.model.DialogPublicUpdateOutboxModel.FindOne(ctx, row.OutboxId)
	if err != nil {
		t.Fatalf("FindOne() error = %v", err)
	}
	if got.Status != OutboxStatusBlocked {
		t.Fatalf("status = %d, want blocked", got.Status)
	}
}

func TestOutboxResetBlockedToPendingPreservesOperationID(t *testing.T) {
	ctx := context.Background()
	repo := newDialogOutboxWorkerTestRepo(t)
	base := time.Now().UnixNano()
	row := newPublicUpdateOutboxRow(base, fmt.Sprintf("public-reset-%d", base), DeliveryPathUserupdatesPTS)
	row.Status = OutboxStatusBlocked
	row.AttemptCount = 9
	row.LastErrorKind = "blocked"
	row.LastErrorMessage = "payload preserved"
	if _, _, err := repo.model.DialogPublicUpdateOutboxModel.Insert(ctx, row); err != nil {
		t.Fatalf("insert public update outbox: %v", err)
	}

	if err := repo.ResetDialogPublicUpdateOutboxBlocked(ctx, []int64{row.OutboxId}); err != nil {
		t.Fatalf("ResetDialogPublicUpdateOutboxBlocked() error = %v", err)
	}

	got, err := repo.model.DialogPublicUpdateOutboxModel.FindOne(ctx, row.OutboxId)
	if err != nil {
		t.Fatalf("FindOne() error = %v", err)
	}
	if got.Status != OutboxStatusPending || got.AttemptCount != 0 || got.LastErrorKind != "" || got.LastErrorMessage != "" {
		t.Fatalf("reset state mismatch: %+v", got)
	}
	if got.OperationId != row.OperationId || string(got.Payload) != string(row.Payload) || string(got.PayloadHash) != string(row.PayloadHash) {
		t.Fatalf("reset changed idempotency fields: got=%+v want operation=%q", got, row.OperationId)
	}
}

func TestOutboxMetricsDoNotUsePayloadAsLabel(t *testing.T) {
	var hits []string
	for _, root := range []string{"."} {
		err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() || filepath.Ext(path) != ".go" || strings.HasSuffix(path, "_test.go") {
				return nil
			}
			b, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			if strings.Contains(string(b), "Payload") && strings.Contains(string(b), "Label") {
				hits = append(hits, path)
			}
			return nil
		})
		if err != nil {
			t.Fatalf("scan metrics labels: %v", err)
		}
	}
	if len(hits) != 0 {
		t.Fatalf("payload appears in metric label paths: %v", hits)
	}
}

func TestDialogOutboxClaimLeasePreventsDoubleClaim(t *testing.T) {
	ctx := context.Background()
	repo := newDialogOutboxWorkerTestRepo(t)
	base := time.Now().UnixNano()
	row := newAuthSeqOutboxRow(base, fmt.Sprintf("auth-lease-%d", base))
	if _, _, err := repo.model.DialogAuthSeqOutboxModel.Insert(ctx, row); err != nil {
		t.Fatalf("insert auth seq outbox: %v", err)
	}

	now := time.Now().UTC()
	first, err := repo.ClaimDialogAuthSeqOutbox(ctx, "owner-a", now, now.Add(time.Minute), 1)
	if err != nil {
		t.Fatalf("first claim error = %v", err)
	}
	if len(first) != 1 || first[0].OutboxId != row.OutboxId {
		t.Fatalf("first claim = %+v, want only outbox_id %d", first, row.OutboxId)
	}
	second, err := repo.ClaimDialogAuthSeqOutbox(ctx, "owner-b", now, now.Add(time.Minute), 1)
	if err != nil {
		t.Fatalf("second claim error = %v", err)
	}
	for _, claimed := range second {
		if claimed.OutboxId == row.OutboxId {
			t.Fatalf("second claim included leased outbox_id %d", row.OutboxId)
		}
	}
	got, err := repo.model.DialogAuthSeqOutboxModel.FindOne(ctx, row.OutboxId)
	if err != nil {
		t.Fatalf("FindOne() error = %v", err)
	}
	if got.Status != OutboxStatusPublishing || got.LeaseOwner != "owner-a" {
		t.Fatalf("leased row state = status:%d owner:%q, want status:%d owner owner-a", got.Status, got.LeaseOwner, OutboxStatusPublishing)
	}
}

func testOutboxOptions(owner string) DialogOutboxWorkerOptions {
	return DialogOutboxWorkerOptions{
		Owner:          owner,
		BatchSize:      1,
		LeaseSeconds:   30,
		PollInterval:   time.Hour,
		BlockedAttempt: OutboxWorkerBlockedAttempts,
	}
}

func newDialogOutboxWorkerTestRepo(t *testing.T) *Repository {
	t.Helper()
	repo := NewRepositoryWithDBForTest(openDialogIntegrationDB(t))
	cleanupDialogOutboxWorkerRows(t, repo)
	t.Cleanup(func() {
		cleanupDialogOutboxWorkerRows(t, repo)
	})
	return repo
}

func cleanupDialogOutboxWorkerRows(t *testing.T, repo *Repository) {
	t.Helper()
	db, err := repo.requireDB()
	if err != nil {
		t.Fatalf("require db: %v", err)
	}
	ctx := context.Background()
	if _, err := db.Exec(ctx, "DELETE FROM dialog_auth_seq_outbox WHERE operation_id LIKE 'auth-%'"); err != nil {
		t.Fatalf("cleanup auth seq outbox worker rows: %v", err)
	}
	if _, err := db.Exec(ctx, "DELETE FROM dialog_public_update_outbox WHERE operation_id LIKE 'public-%'"); err != nil {
		t.Fatalf("cleanup public update outbox worker rows: %v", err)
	}
}

func newAuthSeqOutboxRow(base int64, operationID string) *model.DialogAuthSeqOutbox {
	payload := []byte(`{"schema_version":1}`)
	return &model.DialogAuthSeqOutbox{
		OutboxId:             -base,
		UserId:               base%1_000_000_000 + 101,
		SourcePermAuthKeyId:  base%1_000_000_000 + 201,
		TargetAuthPolicy:     TargetAuthPolicyNotSourcePermAuthKey,
		OperationId:          operationID,
		EventType:            "dialog.preferencePinned",
		PeerType:             PeerTypeUser,
		PeerId:               base%1_000_000_000 + 301,
		PayloadSchemaVersion: 1,
		Payload:              payload,
		PayloadHash:          hashPayload(payload),
		Status:               OutboxStatusPending,
		AttemptCount:         0,
		NextRetryAt:          mysqlZeroDateTime(),
		LeaseUntil:           mysqlZeroDateTime(),
	}
}

func newPublicUpdateOutboxRow(base int64, operationID string, deliveryPath string) *model.DialogPublicUpdateOutbox {
	payload := []byte(`{"schema_version":1}`)
	return &model.DialogPublicUpdateOutbox{
		OutboxId:             -base,
		SourceUserId:         base%1_000_000_000 + 401,
		SourcePermAuthKeyId:  base%1_000_000_000 + 501,
		TargetUserId:         base%1_000_000_000 + 601,
		TargetAuthPolicy:     TargetAuthPolicyAll,
		OperationId:          operationID,
		DeliveryPath:         deliveryPath,
		PublicUpdateType:     "dialog.peerSettings",
		PeerType:             PeerTypeUser,
		PeerId:               base%1_000_000_000 + 701,
		PayloadSchemaVersion: 1,
		Payload:              payload,
		PayloadHash:          hashPayload(payload),
		Status:               OutboxStatusPending,
		AttemptCount:         0,
		NextRetryAt:          mysqlZeroDateTime(),
		LeaseUntil:           mysqlZeroDateTime(),
	}
}
