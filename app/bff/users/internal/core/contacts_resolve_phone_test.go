package core

import (
	"context"
	"testing"

	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestContactsResolvePhoneReturnsResolvedPeer(t *testing.T) {
	var gotPhone *userpb.TLUserGetUserIdByPhone
	var gotUsers *userpb.TLUserGetMutableUsersV2
	core := newUsersCoreForTest(&fakeUserClient{
		getUserIDByPhone: func(_ context.Context, in *userpb.TLUserGetUserIdByPhone) (*tg.Int64, error) {
			gotPhone = in
			return &tg.Int64{V: 2002}, nil
		},
		getMutableUsersV2: func(_ context.Context, in *userpb.TLUserGetMutableUsersV2) (*tg.MutableUsers, error) {
			gotUsers = in
			return tg.MakeTLMutableUsers(&tg.TLMutableUsers{
				Users: []tg.ImmutableUserClazz{immutableUserFixture(2002, "Ada", "", "ada")},
			}).ToMutableUsers(), nil
		},
	}, 1001)

	got, err := core.ContactsResolvePhone(&tg.TLContactsResolvePhone{Phone: "15551230000"})
	if err != nil {
		t.Fatalf("ContactsResolvePhone returned error: %v", err)
	}
	if gotPhone == nil || gotPhone.Phone != "15551230000" {
		t.Fatalf("phone request = %+v, want phone 15551230000", gotPhone)
	}
	if gotUsers == nil || len(gotUsers.Id) != 2 || gotUsers.Id[0] != 2002 || gotUsers.Id[1] != 1001 {
		t.Fatalf("mutable users request = %+v, want ids [2002 1001]", gotUsers)
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
	if id, _ := userID(got.Users[0]); id != 2002 {
		t.Fatalf("resolved user id = %d, want 2002", id)
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
