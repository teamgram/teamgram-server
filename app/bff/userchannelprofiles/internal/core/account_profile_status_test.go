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
	var gotProjection *userpb.TLUserGetUserProjectionBundle
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
		getUserProjection: func(_ context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
			gotProjection = in
			return userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
				ViewerUsers: []userpb.ViewerUsersClazz{
					userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 1001, Users: []tg.UserClazz{
						tg.MakeTLUser(&tg.TLUser{Id: 1001, Self: true, FirstName: strPtr("Ada")}),
					}}),
				},
			}).ToUserProjectionBundle(), nil
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
	if gotProjection == nil || len(gotProjection.ViewerUserIds) != 1 || gotProjection.ViewerUserIds[0] != 1001 ||
		len(gotProjection.TargetUserIds) != 1 || gotProjection.TargetUserIds[0] != 1001 {
		t.Fatalf("projection request = %+v, want viewer/target 1001", gotProjection)
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

func strPtr(v string) *string {
	return &v
}
