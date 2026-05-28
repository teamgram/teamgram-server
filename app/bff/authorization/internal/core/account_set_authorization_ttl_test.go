package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/authorization/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/bff/authorization/internal/svc"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeAuthorizationTTLRepository struct {
	repository.AuthorizationRepository
	setAuthorizationTTL func(context.Context, int64, int32) error
}

func (f *fakeAuthorizationTTLRepository) SetAuthorizationTTL(ctx context.Context, userId int64, ttl int32) error {
	return f.setAuthorizationTTL(ctx, userId, ttl)
}

func TestAccountSetAuthorizationTTLDelegatesToUserService(t *testing.T) {
	var gotUserID int64
	var gotTTL int32

	c := New(context.Background(), &svc.ServiceContext{
		Repo: &fakeAuthorizationTTLRepository{
			setAuthorizationTTL: func(_ context.Context, userId int64, ttl int32) error {
				gotUserID = userId
				gotTTL = ttl
				return nil
			},
		},
	})
	c.MD = &metadata.RpcMetadata{UserId: 1001}

	got, err := c.AccountSetAuthorizationTTL(&tg.TLAccountSetAuthorizationTTL{AuthorizationTtlDays: 90})
	if err != nil {
		t.Fatalf("AccountSetAuthorizationTTL error = %v", err)
	}
	if got != tg.BoolTrue {
		t.Fatalf("AccountSetAuthorizationTTL result = %v, want BoolTrue", got)
	}
	if gotUserID != 1001 || gotTTL != 90 {
		t.Fatalf("set ttl request = user_id:%d ttl:%d, want user_id:1001 ttl:90", gotUserID, gotTTL)
	}
}
