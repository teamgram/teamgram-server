package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/passport/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/bff/passport/internal/svc"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	authsessionclient "github.com/teamgram/teamgram-server/v2/app/service/authsession/client"
	userclient "github.com/teamgram/teamgram-server/v2/app/service/biz/user/client"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakePassportUserClient struct {
	userclient.UserClient
	getAuthorizationTTL func(context.Context, *userpb.TLUserGetAuthorizationTTL) (*tg.AccountDaysTTL, error)
}

type fakePassportAuthsessionClient struct {
	authsessionclient.AuthsessionClient
	getAuthorizations func(context.Context, *authsession.TLAuthsessionGetAuthorizations) (*tg.AccountAuthorizations, error)
}

func (f *fakePassportUserClient) UserGetAuthorizationTTL(ctx context.Context, in *userpb.TLUserGetAuthorizationTTL) (*tg.AccountDaysTTL, error) {
	return f.getAuthorizationTTL(ctx, in)
}

func (f *fakePassportAuthsessionClient) AuthsessionGetAuthorizations(ctx context.Context, in *authsession.TLAuthsessionGetAuthorizations) (*tg.AccountAuthorizations, error) {
	return f.getAuthorizations(ctx, in)
}

func newPassportCoreForTest(userClient userclient.UserClient, authClient authsessionclient.AuthsessionClient) *PassportCore {
	c := New(context.Background(), &svc.ServiceContext{
		Repo: &repository.Repository{
			UserClient:        userClient,
			AuthsessionClient: authClient,
		},
	})
	c.MD = &metadata.RpcMetadata{
		UserId:        1001,
		PermAuthKeyId: 9001,
	}
	return c
}

func TestAccountGetAuthorizationsDelegatesAndFillsTTL(t *testing.T) {
	var gotAuthReq *authsession.TLAuthsessionGetAuthorizations
	var gotTTLReq *userpb.TLUserGetAuthorizationTTL
	c := newPassportCoreForTest(&fakePassportUserClient{
		getAuthorizationTTL: func(_ context.Context, in *userpb.TLUserGetAuthorizationTTL) (*tg.AccountDaysTTL, error) {
			gotTTLReq = in
			return tg.MakeTLAccountDaysTTL(&tg.TLAccountDaysTTL{Days: 45}).ToAccountDaysTTL(), nil
		},
	}, &fakePassportAuthsessionClient{
		getAuthorizations: func(_ context.Context, in *authsession.TLAuthsessionGetAuthorizations) (*tg.AccountAuthorizations, error) {
			gotAuthReq = in
			return tg.MakeTLAccountAuthorizations(&tg.TLAccountAuthorizations{}).ToAccountAuthorizations(), nil
		},
	})

	got, err := c.AccountGetAuthorizations(&tg.TLAccountGetAuthorizations{})
	if err != nil {
		t.Fatalf("AccountGetAuthorizations error = %v", err)
	}
	if got.AuthorizationTtlDays != 45 {
		t.Fatalf("authorization_ttl_days = %d, want 45", got.AuthorizationTtlDays)
	}
	if gotAuthReq == nil || gotAuthReq.UserId != 1001 || gotAuthReq.ExcludeAuthKeyId != 9001 {
		t.Fatalf("authsession request = %+v, want user_id 1001 exclude_auth_key_id 9001", gotAuthReq)
	}
	if gotTTLReq == nil || gotTTLReq.UserId != 1001 {
		t.Fatalf("ttl request = %+v, want user_id 1001", gotTTLReq)
	}
}
