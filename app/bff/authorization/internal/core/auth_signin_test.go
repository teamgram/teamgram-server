package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestAuthSignInReturnsPhoneCodeEmptyWhenCodeMissing(t *testing.T) {
	c := New(context.Background(), nil)

	_, err := c.AuthSignIn(&tg.TLAuthSignIn{
		PhoneNumber:   "+8613812345678",
		PhoneCodeHash: "hash",
		PhoneCode:     nil,
	})
	if err != tg.ErrPhoneCodeEmpty {
		t.Fatalf("expected ErrPhoneCodeEmpty, got %v", err)
	}
}

func TestAuthSignInReturnsPhoneCodeEmptyWhenHashMissing(t *testing.T) {
	c := New(context.Background(), nil)
	code := "12345"

	_, err := c.AuthSignIn(&tg.TLAuthSignIn{
		PhoneNumber:   "+8613812345678",
		PhoneCodeHash: "",
		PhoneCode:     &code,
	})
	if err != tg.ErrPhoneCodeEmpty {
		t.Fatalf("expected ErrPhoneCodeEmpty, got %v", err)
	}
}
