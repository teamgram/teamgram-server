package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/authorization/model"
	kitexmetadata "github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestAuthSignUpReturnsPhoneCodeHashEmptyWhenHashMissing(t *testing.T) {
	c := New(context.Background(), nil)

	_, err := c.AuthSignUp(&tg.TLAuthSignUp{
		PhoneNumber:   "+8613812345678",
		PhoneCodeHash: "",
		FirstName:     "Test",
	})
	if err != tg.ErrPhoneCodeHashEmpty {
		t.Fatalf("expected ErrPhoneCodeHashEmpty, got %v", err)
	}
}

func TestAuthSignUpReturnsFirstnameInvalidWhenFirstNameMissing(t *testing.T) {
	c := New(context.Background(), nil)

	_, err := c.AuthSignUp(&tg.TLAuthSignUp{
		PhoneNumber:   "+8613812345678",
		PhoneCodeHash: "hash",
		FirstName:     "",
	})
	if err != tg.ErrFirstnameInvalid {
		t.Fatalf("expected ErrFirstnameInvalid, got %v", err)
	}
}

func TestAuthSignUpPersistsPhoneCodeStateBeforeUserCreation(t *testing.T) {
	c, ctx, d := newAuthorizationCoreForAuthTest(t, &kitexmetadata.RpcMetadata{
		PermAuthKeyId: 101,
		SessionId:     202,
		Layer:         181,
	})
	phoneNumber := seedPhoneCodeTransaction(t, ctx, d, 101, 202, false, "12345", "hash", model.CodeStateSignIn)

	_, err := c.AuthSignUp(&tg.TLAuthSignUp{
		PhoneNumber:   "+8613812345678",
		PhoneCodeHash: "hash",
		FirstName:     "Test",
		LastName:      "User",
	})
	if err != tg.ErrInternalServerError {
		t.Fatalf("expected ErrInternalServerError until user creation is implemented, got %v", err)
	}

	codeData, err := d.GetPhoneCode(ctx, 101, phoneNumber, "hash")
	if err != nil {
		t.Fatalf("get phone code after sign up: %v", err)
	}
	if codeData.State != model.CodeStateOk {
		t.Fatalf("expected state %d, got %d", model.CodeStateOk, codeData.State)
	}
}
