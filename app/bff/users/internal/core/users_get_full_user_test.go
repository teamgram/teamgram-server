package core

import (
	"context"
	"testing"

	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestUsersGetFullUserDelegatesToUserService(t *testing.T) {
	var gotReq *userpb.TLUserGetFullUser
	want := tg.MakeTLUsersUserFull(&tg.TLUsersUserFull{
		FullUser: tg.MakeTLUserFull(&tg.TLUserFull{Id: 2002}),
		Users:    []tg.UserClazz{tg.MakeTLUser(&tg.TLUser{Id: 2002})},
		Chats:    []tg.ChatClazz{},
	}).ToUsersUserFull()

	core := newUsersCoreForTest(&fakeUserClient{
		getFullUser: func(_ context.Context, in *userpb.TLUserGetFullUser) (*tg.UsersUserFull, error) {
			gotReq = in
			return want, nil
		},
	}, 1001)

	got, err := core.UsersGetFullUser(&tg.TLUsersGetFullUser{
		Id: tg.MakeTLInputUser(&tg.TLInputUser{UserId: 2002, AccessHash: 22}),
	})
	if err != nil {
		t.Fatalf("UsersGetFullUser returned error: %v", err)
	}
	if got != want {
		t.Fatal("UsersGetFullUser did not return user service result")
	}
	if gotReq == nil || gotReq.SelfUserId != 1001 || gotReq.Id != 2002 {
		t.Fatalf("request = %+v, want self=1001 id=2002", gotReq)
	}
}

func TestUsersGetFullUserAcceptsSelf(t *testing.T) {
	var gotReq *userpb.TLUserGetFullUser
	core := newUsersCoreForTest(&fakeUserClient{
		getFullUser: func(_ context.Context, in *userpb.TLUserGetFullUser) (*tg.UsersUserFull, error) {
			gotReq = in
			return tg.MakeTLUsersUserFull(&tg.TLUsersUserFull{
				FullUser: tg.MakeTLUserFull(&tg.TLUserFull{Id: 1001}),
				Users:    []tg.UserClazz{tg.MakeTLUser(&tg.TLUser{Id: 1001, Self: true})},
				Chats:    []tg.ChatClazz{},
			}).ToUsersUserFull(), nil
		},
	}, 1001)

	_, err := core.UsersGetFullUser(&tg.TLUsersGetFullUser{Id: tg.MakeTLInputUserSelf(&tg.TLInputUserSelf{})})
	if err != nil {
		t.Fatalf("UsersGetFullUser(self) returned error: %v", err)
	}
	if gotReq == nil || gotReq.SelfUserId != 1001 || gotReq.Id != 1001 {
		t.Fatalf("request = %+v, want self=1001 id=1001", gotReq)
	}
}
