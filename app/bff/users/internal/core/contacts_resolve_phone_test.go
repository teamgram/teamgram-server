package core

import (
	"context"
	"testing"

	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestContactsResolvePhoneReturnsResolvedPeer(t *testing.T) {
	var gotPhone *userpb.TLUserGetUserIdByPhone
	var gotUsers *userpb.TLUserGetUserProjectionBundle
	core := newUsersCoreForTest(&fakeUserClient{
		getUserIDByPhone: func(_ context.Context, in *userpb.TLUserGetUserIdByPhone) (*tg.Int64, error) {
			gotPhone = in
			return &tg.Int64{V: 2002}, nil
		},
		getUserProjectionBundle: func(_ context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
			gotUsers = in
			return userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
				ViewerUsers: []userpb.ViewerUsersClazz{
					userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 1001, Users: []tg.UserClazz{
						tg.MakeTLUser(&tg.TLUser{Id: 2002, Contact: true}),
					}}),
				},
			}).ToUserProjectionBundle(), nil
		},
	}, 1001)

	got, err := core.ContactsResolvePhone(&tg.TLContactsResolvePhone{Phone: "15551230000"})
	if err != nil {
		t.Fatalf("ContactsResolvePhone returned error: %v", err)
	}
	if gotPhone == nil || gotPhone.Phone != "15551230000" {
		t.Fatalf("phone request = %+v, want phone 15551230000", gotPhone)
	}
	if gotUsers == nil || len(gotUsers.ViewerUserIds) != 1 || gotUsers.ViewerUserIds[0] != 1001 || len(gotUsers.TargetUserIds) != 1 || gotUsers.TargetUserIds[0] != 2002 {
		t.Fatalf("projection request = %+v, want viewer [1001] target [2002]", gotUsers)
	}
	if _, ok := got.Peer.(*tg.TLPeerUser); !ok {
		t.Fatalf("peer = %T, want *tg.TLPeerUser", got.Peer)
	}
	if len(got.Chats) != 0 {
		t.Fatalf("chats len = %d, want 0", len(got.Chats))
	}
	if len(got.Users) != 1 {
		t.Fatalf("users len = %d, want 1", len(got.Users))
	}
	if user, ok := got.Users[0].(*tg.TLUser); !ok || user.Id != 2002 || !user.Contact {
		t.Fatalf("resolved user = %#v, want contact user 2002", got.Users[0])
	}
}

func TestContactsResolvePhoneMapsMissingPhone(t *testing.T) {
	core := newUsersCoreForTest(&fakeUserClient{
		getUserIDByPhone: func(context.Context, *userpb.TLUserGetUserIdByPhone) (*tg.Int64, error) {
			return &tg.Int64{V: 0}, nil
		},
	}, 1001)

	_, err := core.ContactsResolvePhone(&tg.TLContactsResolvePhone{Phone: "15551230000"})
	if err != tg.ErrPhoneNotOccupied {
		t.Fatalf("error = %v, want PHONE_NOT_OCCUPIED", err)
	}
}
