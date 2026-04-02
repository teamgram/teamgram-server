package core

import (
	"context"
	"github.com/teamgram/teamgram-server/v2/app/bff/authorization/model"
	kitexmetadata "github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
	"testing"
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

func TestAuthSignInReturnsPhoneNumberInvalidForBadPhone(t *testing.T) {
	c := New(context.Background(), nil)
	code := "12345"

	_, err := c.AuthSignIn(&tg.TLAuthSignIn{
		PhoneNumber:   "bad-phone",
		PhoneCodeHash: "hash",
		PhoneCode:     &code,
	})
	if err != tg.Err406PhoneNumberInvalid {
		t.Fatalf("expected Err406PhoneNumberInvalid, got %v", err)
	}
}

func TestAuthSignInReturnsSignUpRequiredForUnregisteredPhone(t *testing.T) {
	c, ctx, d := newAuthorizationCoreForAuthTest(t, &kitexmetadata.RpcMetadata{
		PermAuthKeyId: 101,
		SessionId:     202,
		Layer:         181,
	})
	code := "12345"
	phoneNumber := seedPhoneCodeTransaction(t, ctx, d, 101, 202, false, code, "hash", model.CodeStateSent)

	result, err := c.AuthSignIn(&tg.TLAuthSignIn{
		PhoneNumber:   "+8613812345678",
		PhoneCodeHash: "hash",
		PhoneCode:     &code,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if _, ok := result.ToAuthAuthorizationSignUpRequired(); !ok {
		t.Fatalf("expected auth.authorizationSignUpRequired, got %T", result.Clazz)
	}

	codeData, err := d.GetPhoneCode(ctx, 101, phoneNumber, "hash")
	if err != nil {
		t.Fatalf("get phone code after sign in: %v", err)
	}
	if codeData.State != model.CodeStateSignIn {
		t.Fatalf("expected state %d, got %d", model.CodeStateSignIn, codeData.State)
	}
}

func TestAuthSignInReturnsPhoneCodeInvalidWhenCodeDoesNotMatch(t *testing.T) {
	c, ctx, d := newAuthorizationCoreForAuthTest(t, &kitexmetadata.RpcMetadata{
		PermAuthKeyId: 101,
		SessionId:     202,
		Layer:         181,
	})
	seedPhoneCodeTransaction(t, ctx, d, 101, 202, false, "12345", "hash", model.CodeStateSent)
	wrongCode := "54321"

	_, err := c.AuthSignIn(&tg.TLAuthSignIn{
		PhoneNumber:   "+8613812345678",
		PhoneCodeHash: "hash",
		PhoneCode:     &wrongCode,
	})
	if err != tg.ErrPhoneCodeInvalid {
		t.Fatalf("expected ErrPhoneCodeInvalid, got %v", err)
	}
}
