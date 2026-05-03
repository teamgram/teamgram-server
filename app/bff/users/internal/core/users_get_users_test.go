package core

import (
	"context"
	"reflect"
	"testing"

	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestUsersGetUsersFetchesMutableUsersInRequestOrder(t *testing.T) {
	var gotReq *userpb.TLUserGetMutableUsersV2
	core := newUsersCoreForTest(&fakeUserClient{
		getMutableUsersV2: func(_ context.Context, in *userpb.TLUserGetMutableUsersV2) (*tg.MutableUsers, error) {
			gotReq = in
			return tg.MakeTLMutableUsers(&tg.TLMutableUsers{
				Users: []tg.ImmutableUserClazz{
					immutableUserFixture(2002, "Ada", "", "ada"),
				},
			}).ToMutableUsers(), nil
		},
	}, 1001)

	got, err := core.UsersGetUsers(&tg.TLUsersGetUsers{
		Id: []tg.InputUserClazz{
			tg.MakeTLInputUserSelf(&tg.TLInputUserSelf{}),
			tg.MakeTLInputUser(&tg.TLInputUser{UserId: 2002, AccessHash: 22}),
			tg.MakeTLInputUser(&tg.TLInputUser{UserId: 3003, AccessHash: 33}),
		},
	})
	if err != nil {
		t.Fatalf("UsersGetUsers returned error: %v", err)
	}
	if gotReq == nil {
		t.Fatal("UserGetMutableUsersV2 was not called")
	}
	if !reflect.DeepEqual(gotReq.Id, []int64{1001, 2002, 3003}) {
		t.Fatalf("request ids = %v, want [1001 2002 3003]", gotReq.Id)
	}
	if !gotReq.Privacy || !gotReq.HasTo || !reflect.DeepEqual(gotReq.To, []int64{1001}) {
		t.Fatalf("privacy request = %+v, want privacy=true has_to=true to=[1001]", gotReq)
	}
	if len(got.Datas) != 3 {
		t.Fatalf("len(VectorUser) = %d, want 3", len(got.Datas))
	}
	if id, _ := userID(got.Datas[0]); id != 1001 {
		t.Fatalf("first user id = %d, want 1001", id)
	}
	if id, _ := userID(got.Datas[1]); id != 2002 {
		t.Fatalf("second user id = %d, want 2002", id)
	}
	if empty, ok := got.Datas[2].(*tg.TLUserEmpty); !ok || empty.Id != 3003 {
		t.Fatalf("third user = %T %+v, want userEmpty id 3003", got.Datas[2], got.Datas[2])
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
