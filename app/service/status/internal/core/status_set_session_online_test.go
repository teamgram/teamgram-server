package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/status/status"
)

func TestStatusSetSessionOnlineRejectsInvalidSessionEntry(t *testing.T) {
	c := New(context.Background(), nil)

	_, err := c.StatusSetSessionOnline(&status.TLStatusSetSessionOnline{
		UserId:  1,
		Session: nil,
	})
	if err == nil {
		t.Fatal("expected error for nil session entry")
	}

	_, err = c.StatusSetSessionOnline(&status.TLStatusSetSessionOnline{
		UserId: 1,
		Session: &status.TLSessionEntry{
			UserId:    1,
			AuthKeyId: 0,
			Gateway:   "gw-1",
			Client:    "ios",
		},
	})
	if err == nil {
		t.Fatal("expected error for zero auth key")
	}
}
