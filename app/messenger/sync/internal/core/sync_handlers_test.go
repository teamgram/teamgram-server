package core

import (
	"context"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/sync/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/messenger/sync/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/sync/internal/svc"
	syncpb "github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/identity"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestSyncPushUpdatesIfNotRejectsOversizedIncludes(t *testing.T) {
	includes := make([]int64, repository.MaxIncludePermKeys+1)
	for i := range includes {
		includes[i] = int64(i + 1)
	}
	c := newTestSyncCore(t, context.Background(), &fakeRepo{}, config.Config{})
	_, err := c.SyncPushUpdatesIfNot(&syncpb.TLSyncPushUpdatesIfNot{
		UserId:   42,
		Includes: includes,
		Updates:  emptyUpdates(),
	})
	if !errors.Is(err, syncpb.ErrSyncInvalidArgument) {
		t.Fatalf("error = %v, want ErrSyncInvalidArgument", err)
	}
}

func TestSyncUpdatesMeRejectsHalfSpecifiedPreciseSession(t *testing.T) {
	c := newTestSyncCore(t, context.Background(), &fakeRepo{}, config.Config{})
	authKeyID := int64(1001)
	_, err := c.SyncUpdatesMe(&syncpb.TLSyncUpdatesMe{
		UserId:        42,
		PermAuthKeyId: 9001,
		AuthKeyId:     &authKeyID,
		Updates:       emptyUpdates(),
	})
	if !errors.Is(err, syncpb.ErrSyncInvalidArgument) {
		t.Fatalf("error = %v, want ErrSyncInvalidArgument", err)
	}
}

func TestSyncPushRpcResultRejectsMissingGatewayRoute(t *testing.T) {
	c := newTestSyncCore(t, context.Background(), &fakeRepo{}, config.Config{})
	_, err := c.SyncPushRpcResult(&syncpb.TLSyncPushRpcResult{
		UserId:         42,
		PermAuthKeyId:  9001,
		AuthKeyId:      1001,
		SessionId:      2002,
		ClientReqMsgId: 3003,
		RpcResult:      []byte{1},
	})
	if !errors.Is(err, syncpb.ErrSyncInvalidArgument) {
		t.Fatalf("error = %v, want ErrSyncInvalidArgument", err)
	}
}

func TestSyncRejectsUnauthorizedCallerWhenRequired(t *testing.T) {
	ctx := identity.WithCallerService(context.Background(), "bad.caller")
	c := newTestSyncCore(t, ctx, &fakeRepo{}, config.Config{
		RequireCallerIdentity: true,
		AllowedCallers:        []string{"bff.dialogs"},
	})
	_, err := c.SyncPushUpdates(&syncpb.TLSyncPushUpdates{UserId: 42, Updates: emptyUpdates()})
	if !errors.Is(err, syncpb.ErrSyncPermissionDenied) {
		t.Fatalf("error = %v, want ErrSyncPermissionDenied", err)
	}
}

func TestSyncPushUpdatesCallsRepositoryForAllowedCaller(t *testing.T) {
	repo := &fakeRepo{}
	ctx := identity.WithCallerService(context.Background(), "bff.dialogs")
	c := newTestSyncCore(t, ctx, repo, config.Config{
		RequireCallerIdentity: true,
		AllowedCallers:        []string{"bff.dialogs"},
	})
	_, err := c.SyncPushUpdates(&syncpb.TLSyncPushUpdates{UserId: 42, Updates: emptyUpdates()})
	if err != nil {
		t.Fatalf("SyncPushUpdates() error = %v", err)
	}
	if repo.pushUpdatesCount != 1 {
		t.Fatalf("pushUpdatesCount = %d, want 1", repo.pushUpdatesCount)
	}
}

func newTestSyncCore(t *testing.T, ctx context.Context, repo svc.SyncRepository, c config.Config) *SyncCore {
	t.Helper()
	return New(ctx, &svc.ServiceContext{Config: c, Repo: repo})
}

func emptyUpdates() tg.UpdatesClazz {
	return tg.MakeTLUpdates(&tg.TLUpdates{})
}

type fakeRepo struct {
	pushUpdatesCount int
}

func (f *fakeRepo) PushUpdates(ctx context.Context, userID int64, updates tg.UpdatesClazz) error {
	f.pushUpdatesCount++
	return nil
}

func (f *fakeRepo) UpdatesNotMe(ctx context.Context, userID, permAuthKeyID int64, updates tg.UpdatesClazz) error {
	return nil
}

func (f *fakeRepo) PushUpdatesIfNot(ctx context.Context, userID int64, includesSet bool, includes []int64, excludesSet bool, excludes []int64, updates tg.UpdatesClazz) error {
	return nil
}

func (f *fakeRepo) UpdatesMe(ctx context.Context, userID, permAuthKeyID int64, authKeyID, sessionID int64, precise bool, updates tg.UpdatesClazz) error {
	return nil
}

func (f *fakeRepo) PushRpcResult(ctx context.Context, route repository.RpcResultRoute) error {
	return nil
}
