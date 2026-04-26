package core

import (
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/status/status"
)

func TestStatusSetSessionOnlineInvalidUserIDReturnsStatusInvalidArgument(t *testing.T) {
	c := &StatusCore{}

	_, err := c.StatusSetSessionOnline(&status.TLStatusSetSessionOnline{
		UserId: 0,
		Session: &status.TLSessionEntry{
			UserId:    1,
			AuthKeyId: 1,
		},
	})
	if !errors.Is(err, status.ErrStatusInvalidArgument) {
		t.Fatalf("StatusSetSessionOnline() error = %v, want ErrStatusInvalidArgument", err)
	}
}

func TestStatusSetSessionOnlineMismatchedSessionUserReturnsStatusInvalidArgument(t *testing.T) {
	c := &StatusCore{}

	_, err := c.StatusSetSessionOnline(&status.TLStatusSetSessionOnline{
		UserId: 1,
		Session: &status.TLSessionEntry{
			UserId:    2,
			AuthKeyId: 1,
		},
	})
	if !errors.Is(err, status.ErrStatusInvalidArgument) {
		t.Fatalf("StatusSetSessionOnline() error = %v, want ErrStatusInvalidArgument", err)
	}
}

func TestStatusGetUsersOnlineSessionsListTooManyUsersReturnsStatusInvalidArgument(t *testing.T) {
	c := &StatusCore{}

	_, err := c.StatusGetUsersOnlineSessionsList(&status.TLStatusGetUsersOnlineSessionsList{
		Users: make([]int64, maxBatchUsers+1),
	})
	if !errors.Is(err, status.ErrStatusInvalidArgument) {
		t.Fatalf("StatusGetUsersOnlineSessionsList() error = %v, want ErrStatusInvalidArgument", err)
	}
}
