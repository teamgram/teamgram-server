package repository

import (
	"context"
	"errors"
	"testing"

	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
)

func TestGetImmutableUserRejectsZeroID(t *testing.T) {
	r := &Repository{}
	_, err := r.GetImmutableUser(context.Background(), 0, false)
	if !errors.Is(err, userpb.ErrUserNotFound) {
		t.Fatalf("expected ErrUserNotFound, got %v", err)
	}
}
