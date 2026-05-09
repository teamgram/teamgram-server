package repository

import (
	"errors"
	"testing"

	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestFullUserProjectedUserFromBundlePreservesProjectedPhoto(t *testing.T) {
	photo := tg.MakeTLUserProfilePhoto(&tg.TLUserProfilePhoto{PhotoId: 9002, DcId: 1})
	projected := tg.MakeTLUser(&tg.TLUser{Id: 2002, Photo: photo})

	got, err := fullUserProjectedUserFromBundle(&UserProjectionBundle{
		ViewerUsers: []ViewerUsers{
			{ViewerUserId: 1001, Users: []tg.UserClazz{projected}},
		},
	}, 1001, 2002)
	if err != nil {
		t.Fatalf("fullUserProjectedUserFromBundle returned error: %v", err)
	}
	user, ok := got.(*tg.TLUser)
	if !ok {
		t.Fatalf("projected user = %T, want *tg.TLUser", got)
	}
	if user.Photo != photo {
		t.Fatalf("projected photo = %#v, want preserved projection photo", user.Photo)
	}
}

func TestFullUserProjectedUserFromBundleMissingUser(t *testing.T) {
	_, err := fullUserProjectedUserFromBundle(&UserProjectionBundle{
		ViewerUsers: []ViewerUsers{
			{ViewerUserId: 1001, Users: []tg.UserClazz{tg.MakeTLUser(&tg.TLUser{Id: 3003})}},
		},
	}, 1001, 2002)
	if !errors.Is(err, userpb.ErrUserNotFound) {
		t.Fatalf("error = %v, want ErrUserNotFound", err)
	}
}
