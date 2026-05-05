package repository

import (
	"context"
	"errors"
	"fmt"
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
	repo := NewRepositoryWithDBForTest(openDialogIntegrationDB(t))
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
	repo := NewRepositoryWithDBForTest(openDialogIntegrationDB(t))
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
	repo := NewRepositoryWithDBForTest(openDialogIntegrationDB(t))
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
	repo := NewRepositoryWithDBForTest(openDialogIntegrationDB(t))
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
	repo := NewRepositoryWithDBForTest(openDialogIntegrationDB(t))
	base := time.Now().UnixNano()
	retryRow := newPublicUpdateOutboxRow(base, fmt.Sprintf("public-retry-%d", base), DeliveryPathUserupdatesPTS)
	blockRow := newPublicUpdateOutboxRow(base+1, fmt.Sprintf("public-block-%d", base), DeliveryPathUserupdatesPTS)
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

func TestDialogOutboxClaimLeasePreventsDoubleClaim(t *testing.T) {
	ctx := context.Background()
	repo := NewRepositoryWithDBForTest(openDialogIntegrationDB(t))
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

func newAuthSeqOutboxRow(base int64, operationID string) *model.DialogAuthSeqOutbox {
	payload := []byte(`{"schema_version":1}`)
	return &model.DialogAuthSeqOutbox{
		OutboxId:             base%1_000_000_000 + 10_000_000,
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
		NextRetryAt:          mysqlZeroTime(),
		LeaseUntil:           mysqlZeroTime(),
	}
}

func newPublicUpdateOutboxRow(base int64, operationID string, deliveryPath string) *model.DialogPublicUpdateOutbox {
	payload := []byte(`{"schema_version":1}`)
	return &model.DialogPublicUpdateOutbox{
		OutboxId:             base%1_000_000_000 + 20_000_000,
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
		NextRetryAt:          mysqlZeroTime(),
		LeaseUntil:           mysqlZeroTime(),
	}
}
