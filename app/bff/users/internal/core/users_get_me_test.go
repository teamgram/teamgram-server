package core

import (
	"context"
	"testing"

	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestUsersGetMeReturnsSelfUserForMatchingToken(t *testing.T) {
	var gotReq *userpb.TLUserGetImmutableUserByToken
	var gotProjection *userpb.TLUserGetUserProjectionBundle
	core := newUsersCoreForTest(&fakeUserClient{
		getImmutableUserByToken: func(_ context.Context, in *userpb.TLUserGetImmutableUserByToken) (*tg.ImmutableUser, error) {
			gotReq = in
			return immutableUserFixture(1001, "Grace", "", "grace"), nil
		},
		getUserProjectionBundle: func(_ context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
			gotProjection = in
			return userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
				ViewerUsers: []userpb.ViewerUsersClazz{
					userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 1001, Users: []tg.UserClazz{
						tg.MakeTLUser(&tg.TLUser{Id: 1001, Self: true, FirstName: strPtr("Grace")}),
					}}),
				},
			}).ToUserProjectionBundle(), nil
		},
	}, 0)

	got, err := core.UsersGetMe(&tg.TLUsersGetMe{Id: 1001, Token: "bot-token"})
	if err != nil {
		t.Fatalf("UsersGetMe returned error: %v", err)
	}
	if gotReq == nil || gotReq.Token != "bot-token" {
		t.Fatalf("request = %+v, want token bot-token", gotReq)
	}
	if gotProjection == nil || len(gotProjection.ViewerUserIds) != 1 || gotProjection.ViewerUserIds[0] != 1001 ||
		len(gotProjection.TargetUserIds) != 1 || gotProjection.TargetUserIds[0] != 1001 {
		t.Fatalf("projection request = %+v, want viewer/target 1001", gotProjection)
	}
	user, ok := got.Clazz.(*tg.TLUser)
	if !ok {
		t.Fatalf("got user clazz = %T, want *tg.TLUser", got.Clazz)
	}
	if !user.Self || user.Id != 1001 {
		t.Fatalf("self user = %+v, want self id 1001", user)
	}
}

func TestUsersGetMeRejectsTokenUserMismatch(t *testing.T) {
	core := newUsersCoreForTest(&fakeUserClient{
		getImmutableUserByToken: func(context.Context, *userpb.TLUserGetImmutableUserByToken) (*tg.ImmutableUser, error) {
			return immutableUserFixture(2002, "Other", "", "other"), nil
		},
	}, 0)

	_, err := core.UsersGetMe(&tg.TLUsersGetMe{Id: 1001, Token: "bot-token"})
	if err != tg.ErrTokenInvalid {
		t.Fatalf("error = %v, want TOKEN_INVALID", err)
	}
}

func strPtr(v string) *string {
	return &v
}
