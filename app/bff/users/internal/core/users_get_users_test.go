package core

import (
	"context"
	"reflect"
	"testing"

	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestUsersGetUsersProjectsUsersInRequestOrder(t *testing.T) {
	var gotReq *userpb.TLUserGetUserProjectionBundle
	core := newUsersCoreForTest(&fakeUserClient{
		getUserProjectionBundle: func(_ context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
			gotReq = in
			return userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
				ViewerUsers: []userpb.ViewerUsersClazz{
					userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 1001, Users: []tg.UserClazz{
						tg.MakeTLUser(&tg.TLUser{Id: 1001, Self: true}),
						tg.MakeTLUser(&tg.TLUser{Id: 2002, Contact: true}),
					}}),
				},
			}).ToUserProjectionBundle(), nil
		},
	}, 1001)

	got, err := core.UsersGetUsers(&tg.TLUsersGetUsers{
		Id: []tg.InputUserClazz{
			tg.MakeTLInputUserSelf(&tg.TLInputUserSelf{}),
			tg.MakeTLInputUser(&tg.TLInputUser{UserId: 2002, AccessHash: 22}),
		},
	})
	if err != nil {
		t.Fatalf("UsersGetUsers returned error: %v", err)
	}
	if gotReq == nil {
		t.Fatal("UserGetUserProjectionBundle was not called")
	}
	if !reflect.DeepEqual(gotReq.ViewerUserIds, []int64{1001}) {
		t.Fatalf("viewer ids = %v, want [1001]", gotReq.ViewerUserIds)
	}
	if !reflect.DeepEqual(gotReq.TargetUserIds, []int64{1001, 2002}) {
		t.Fatalf("target ids = %v, want [1001 2002]", gotReq.TargetUserIds)
	}
	if len(got.Datas) != 2 {
		t.Fatalf("len(VectorUser) = %d, want 2", len(got.Datas))
	}
	if user, ok := got.Datas[0].(*tg.TLUser); !ok || user.Id != 1001 || !user.Self {
		t.Fatalf("first user = %#v, want self user 1001", got.Datas[0])
	}
	if user, ok := got.Datas[1].(*tg.TLUser); !ok || user.Id != 2002 || !user.Contact {
		t.Fatalf("second user = %#v, want contact user 2002", got.Datas[1])
	}
}

func TestUsersGetUsersMapsMissingProjectionToUserIdInvalid(t *testing.T) {
	core := newUsersCoreForTest(&fakeUserClient{
		getUserProjectionBundle: func(context.Context, *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
			return userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
				MissingUserIds: []int64{3003},
			}).ToUserProjectionBundle(), nil
		},
	}, 1001)

	_, err := core.UsersGetUsers(&tg.TLUsersGetUsers{
		Id: []tg.InputUserClazz{tg.MakeTLInputUser(&tg.TLInputUser{UserId: 3003, AccessHash: 33})},
	})
	if err != tg.ErrUserIdInvalid {
		t.Fatalf("error = %v, want USER_ID_INVALID", err)
	}
}

func TestUsersGetUsersRejectsInvalidInputUser(t *testing.T) {
	core := newUsersCoreForTest(&fakeUserClient{}, 1001)

	_, err := core.UsersGetUsers(&tg.TLUsersGetUsers{
		Id: []tg.InputUserClazz{tg.MakeTLInputUserEmpty(&tg.TLInputUserEmpty{})},
	})
	if err != tg.ErrUserIdInvalid {
		t.Fatalf("error = %v, want USER_ID_INVALID", err)
	}
}
