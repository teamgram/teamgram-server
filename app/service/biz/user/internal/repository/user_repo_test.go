package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/internal/repository/model"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestGetImmutableUserRejectsZeroID(t *testing.T) {
	r := &Repository{}
	_, err := r.GetImmutableUser(context.Background(), 0, false)
	if !errors.Is(err, userpb.ErrUserNotFound) {
		t.Fatalf("expected ErrUserNotFound, got %v", err)
	}
}

func TestUserUpdateTargetErrorAllowsNoOpUpdate(t *testing.T) {
	if err := userUpdateTargetError("update about", 1234, nil); err != nil {
		t.Fatalf("expected no error for existing no-op update, got %v", err)
	}
}

func TestUserUpdateTargetErrorMapsMissingUser(t *testing.T) {
	err := userUpdateTargetError("update about", 1234, model.ErrNotFound)
	if !errors.Is(err, userpb.ErrUserNotFound) {
		t.Fatalf("expected ErrUserNotFound, got %v", err)
	}
}

func TestUserUpdateTargetErrorWrapsStorageFailure(t *testing.T) {
	err := userUpdateTargetError("update about", 1234, errors.New("db unavailable"))
	if !errors.Is(err, userpb.ErrUserStorage) {
		t.Fatalf("expected ErrUserStorage, got %v", err)
	}
}

func TestUserFromModelPopulatesActiveEditableUsername(t *testing.T) {
	user := userFromModel(&model.Users{
		Id:         2002,
		AccessHash: 3003,
		FirstName:  "Ada",
		Username:   "ada",
	}, true, true, true, nil)

	full, ok := user.(*tg.TLUser)
	if !ok {
		t.Fatalf("user = %T, want *tg.TLUser", user)
	}
	if full.Username == nil || *full.Username != "ada" {
		t.Fatalf("username = %v, want ada", full.Username)
	}
	if len(full.Usernames) != 1 {
		t.Fatalf("usernames len = %d, want 1", len(full.Usernames))
	}
	username := full.Usernames[0]
	if username.Username != "ada" || !username.Active || !username.Editable {
		t.Fatalf("usernames[0] = %+v, want active editable ada", username)
	}
}
