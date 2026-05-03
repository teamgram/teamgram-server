package core

import (
	"context"
	"testing"

	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestAccountUpdateProfileUpdatesChangedFields(t *testing.T) {
	var gotAbout *userpb.TLUserUpdateAbout
	var gotName *userpb.TLUserUpdateFirstAndLastName
	core := newUserChannelProfilesCoreForTest(&fakeUserClient{
		getImmutableUser: func(context.Context, *userpb.TLUserGetImmutableUser) (*tg.ImmutableUser, error) {
			return immutableUserFixture(1001, "Old", "Name", "old"), nil
		},
		updateAbout: func(_ context.Context, in *userpb.TLUserUpdateAbout) (*tg.Bool, error) {
			gotAbout = in
			return tg.BoolTrue, nil
		},
		updateFirstAndLastName: func(_ context.Context, in *userpb.TLUserUpdateFirstAndLastName) (*tg.Bool, error) {
			gotName = in
			return tg.BoolTrue, nil
		},
	}, nil, 1001)

	first, last, about := "Ada", "Lovelace", "analytical engine"
	got, err := core.AccountUpdateProfile(&tg.TLAccountUpdateProfile{FirstName: &first, LastName: &last, About: &about})
	if err != nil {
		t.Fatalf("AccountUpdateProfile returned error: %v", err)
	}
	if gotAbout == nil || gotAbout.UserId != 1001 || gotAbout.About != about {
		t.Fatalf("about request = %+v", gotAbout)
	}
	if gotName == nil || gotName.UserId != 1001 || gotName.FirstName != first || gotName.LastName != last {
		t.Fatalf("name request = %+v", gotName)
	}
	user, ok := got.Clazz.(*tg.TLUser)
	if !ok || !user.Self || user.FirstName == nil || *user.FirstName != first {
		t.Fatalf("returned user = %#v", got)
	}
}

func TestAccountUpdateProfileRejectsLongAbout(t *testing.T) {
	core := newUserChannelProfilesCoreForTest(&fakeUserClient{
		getImmutableUser: func(context.Context, *userpb.TLUserGetImmutableUser) (*tg.ImmutableUser, error) {
			return immutableUserFixture(1001, "Ada", "", "ada"), nil
		},
	}, nil, 1001)
	about := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	if _, err := core.AccountUpdateProfile(&tg.TLAccountUpdateProfile{About: &about}); err != tg.ErrAboutTooLong {
		t.Fatalf("error = %v, want ABOUT_TOO_LONG", err)
	}
}

func TestAccountUpdateStatusUpdatesLastSeen(t *testing.T) {
	var got *userpb.TLUserUpdateLastSeen
	core := newUserChannelProfilesCoreForTest(&fakeUserClient{
		updateLastSeen: func(_ context.Context, in *userpb.TLUserUpdateLastSeen) (*tg.Bool, error) {
			got = in
			return tg.BoolTrue, nil
		},
	}, nil, 1001)

	if _, err := core.AccountUpdateStatus(&tg.TLAccountUpdateStatus{Offline: tg.BoolFalseClazz}); err != nil {
		t.Fatalf("AccountUpdateStatus returned error: %v", err)
	}
	if got == nil || got.Id != 1001 || got.Expires != 300 || got.LastSeenAt <= 0 {
		t.Fatalf("last seen request = %+v, want user 1001 online expires 300", got)
	}

	if _, err := core.AccountUpdateStatus(&tg.TLAccountUpdateStatus{Offline: tg.BoolTrueClazz}); err != nil {
		t.Fatalf("AccountUpdateStatus offline returned error: %v", err)
	}
	if got == nil || got.Id != 1001 || got.Expires != 0 || got.LastSeenAt <= 0 {
		t.Fatalf("last seen request = %+v, want user 1001 offline expires 0", got)
	}
}
