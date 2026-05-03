package core

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/bff/account/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/bff/account/internal/svc"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	authsessionclient "github.com/teamgram/teamgram-server/v2/app/service/authsession/client"
	codeclient "github.com/teamgram/teamgram-server/v2/app/service/biz/code/client"
	codepb "github.com/teamgram/teamgram-server/v2/app/service/biz/code/code"
	userclient "github.com/teamgram/teamgram-server/v2/app/service/biz/user/client"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeAccountUserClient struct {
	userclient.UserClient
	getAccountDaysTTL       func(context.Context, *userpb.TLUserGetAccountDaysTTL) (*tg.AccountDaysTTL, error)
	setAccountDaysTTL       func(context.Context, *userpb.TLUserSetAccountDaysTTL) (*tg.Bool, error)
	getImmutableUserByPhone func(context.Context, *userpb.TLUserGetImmutableUserByPhone) (*tg.ImmutableUser, error)
	getImmutableUser        func(context.Context, *userpb.TLUserGetImmutableUser) (*tg.ImmutableUser, error)
	changePhone             func(context.Context, *userpb.TLUserChangePhone) (*tg.Bool, error)
	getUserDataByID         func(context.Context, *userpb.TLUserGetUserDataById) (*tg.UserData, error)
	deleteUsername          func(context.Context, *userpb.TLUserDeleteUsername) (*tg.Bool, error)
	deleteUser              func(context.Context, *userpb.TLUserDeleteUser) (*tg.Bool, error)
}

type fakeAccountAuthsessionClient struct {
	authsessionclient.AuthsessionClient
	resetAuthorization func(context.Context, *authsession.TLAuthsessionResetAuthorization) (*authsession.VectorLong, error)
	unbindAuthKeyUser  func(context.Context, *authsession.TLAuthsessionUnbindAuthKeyUser) (*tg.Bool, error)
}

type fakeAccountCodeClient struct {
	codeclient.CodeClient
	createPhoneCode func(context.Context, *codepb.TLCodeCreatePhoneCode) (*codepb.PhoneCodeTransaction, error)
	getPhoneCode    func(context.Context, *codepb.TLCodeGetPhoneCode) (*codepb.PhoneCodeTransaction, error)
	deletePhoneCode func(context.Context, *codepb.TLCodeDeletePhoneCode) (*tg.Bool, error)
}

func (f *fakeAccountUserClient) UserGetAccountDaysTTL(ctx context.Context, in *userpb.TLUserGetAccountDaysTTL) (*tg.AccountDaysTTL, error) {
	return f.getAccountDaysTTL(ctx, in)
}

func (f *fakeAccountUserClient) UserSetAccountDaysTTL(ctx context.Context, in *userpb.TLUserSetAccountDaysTTL) (*tg.Bool, error) {
	return f.setAccountDaysTTL(ctx, in)
}

func (f *fakeAccountUserClient) UserGetImmutableUserByPhone(ctx context.Context, in *userpb.TLUserGetImmutableUserByPhone) (*tg.ImmutableUser, error) {
	return f.getImmutableUserByPhone(ctx, in)
}

func (f *fakeAccountUserClient) UserGetImmutableUser(ctx context.Context, in *userpb.TLUserGetImmutableUser) (*tg.ImmutableUser, error) {
	return f.getImmutableUser(ctx, in)
}

func (f *fakeAccountUserClient) UserChangePhone(ctx context.Context, in *userpb.TLUserChangePhone) (*tg.Bool, error) {
	return f.changePhone(ctx, in)
}

func (f *fakeAccountUserClient) UserGetUserDataById(ctx context.Context, in *userpb.TLUserGetUserDataById) (*tg.UserData, error) {
	return f.getUserDataByID(ctx, in)
}

func (f *fakeAccountUserClient) UserDeleteUsername(ctx context.Context, in *userpb.TLUserDeleteUsername) (*tg.Bool, error) {
	return f.deleteUsername(ctx, in)
}

func (f *fakeAccountUserClient) UserDeleteUser(ctx context.Context, in *userpb.TLUserDeleteUser) (*tg.Bool, error) {
	return f.deleteUser(ctx, in)
}

func (f *fakeAccountAuthsessionClient) AuthsessionResetAuthorization(ctx context.Context, in *authsession.TLAuthsessionResetAuthorization) (*authsession.VectorLong, error) {
	return f.resetAuthorization(ctx, in)
}

func (f *fakeAccountAuthsessionClient) AuthsessionUnbindAuthKeyUser(ctx context.Context, in *authsession.TLAuthsessionUnbindAuthKeyUser) (*tg.Bool, error) {
	return f.unbindAuthKeyUser(ctx, in)
}

func (f *fakeAccountCodeClient) CodeCreatePhoneCode(ctx context.Context, in *codepb.TLCodeCreatePhoneCode) (*codepb.PhoneCodeTransaction, error) {
	return f.createPhoneCode(ctx, in)
}

func (f *fakeAccountCodeClient) CodeGetPhoneCode(ctx context.Context, in *codepb.TLCodeGetPhoneCode) (*codepb.PhoneCodeTransaction, error) {
	return f.getPhoneCode(ctx, in)
}

func (f *fakeAccountCodeClient) CodeDeletePhoneCode(ctx context.Context, in *codepb.TLCodeDeletePhoneCode) (*tg.Bool, error) {
	return f.deletePhoneCode(ctx, in)
}

func newAccountCoreForTest(userClient userclient.UserClient, authClient authsessionclient.AuthsessionClient, codeClient codeclient.CodeClient) *AccountCore {
	c := New(context.Background(), &svc.ServiceContext{
		Repo: &repository.Repository{
			UserClient:        userClient,
			AuthsessionClient: authClient,
			CodeClient:        codeClient,
		},
	})
	c.MD = &metadata.RpcMetadata{
		UserId:        1001,
		PermAuthKeyId: 9001,
		AuthId:        9001,
		SessionId:     7001,
	}
	return c
}

func accountImmutableUserFixture(id int64, phone string, username string) *tg.ImmutableUser {
	return tg.MakeTLImmutableUser(&tg.TLImmutableUser{
		User: tg.MakeTLUserData(&tg.TLUserData{
			Id:         id,
			AccessHash: id * 10,
			FirstName:  "Test",
			LastName:   "User",
			Username:   username,
			Phone:      phone,
		}),
	}).ToImmutableUser()
}
