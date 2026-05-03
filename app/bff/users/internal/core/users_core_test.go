package core

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/bff/users/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/bff/users/internal/svc"
	userclient "github.com/teamgram/teamgram-server/v2/app/service/biz/user/client"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeUserClient struct {
	userclient.UserClient
	getMutableUsersV2       func(context.Context, *userpb.TLUserGetMutableUsersV2) (*tg.MutableUsers, error)
	getFullUser             func(context.Context, *userpb.TLUserGetFullUser) (*tg.UsersUserFull, error)
	getImmutableUserByToken func(context.Context, *userpb.TLUserGetImmutableUserByToken) (*tg.ImmutableUser, error)
	getUserIDByPhone        func(context.Context, *userpb.TLUserGetUserIdByPhone) (*tg.Int64, error)
}

func (f *fakeUserClient) UserGetMutableUsersV2(ctx context.Context, in *userpb.TLUserGetMutableUsersV2) (*tg.MutableUsers, error) {
	return f.getMutableUsersV2(ctx, in)
}

func (f *fakeUserClient) UserGetFullUser(ctx context.Context, in *userpb.TLUserGetFullUser) (*tg.UsersUserFull, error) {
	return f.getFullUser(ctx, in)
}

func (f *fakeUserClient) UserGetImmutableUserByToken(ctx context.Context, in *userpb.TLUserGetImmutableUserByToken) (*tg.ImmutableUser, error) {
	return f.getImmutableUserByToken(ctx, in)
}

func (f *fakeUserClient) UserGetUserIdByPhone(ctx context.Context, in *userpb.TLUserGetUserIdByPhone) (*tg.Int64, error) {
	return f.getUserIDByPhone(ctx, in)
}

func newUsersCoreForTest(userClient userclient.UserClient, selfID int64) *UsersCore {
	c := New(context.Background(), &svc.ServiceContext{
		Repo: &repository.Repository{UserClient: userClient},
	})
	c.MD = &metadata.RpcMetadata{UserId: selfID, PermAuthKeyId: 9001}
	return c
}

func immutableUserFixture(id int64, firstName string, lastName string, username string) *tg.ImmutableUser {
	return tg.MakeTLImmutableUser(&tg.TLImmutableUser{
		User: tg.MakeTLUserData(&tg.TLUserData{
			Id:         id,
			AccessHash: id * 10,
			FirstName:  firstName,
			LastName:   lastName,
			Username:   username,
		}),
	})
}
