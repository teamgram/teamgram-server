package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/authorization/model"
	kitexmetadata "github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestAuthResendCodeReturnsPhoneCodeHashEmptyWhenHashMissing(t *testing.T) {
	c := New(context.Background(), nil)

	_, err := c.AuthResendCode(&tg.TLAuthResendCode{
		PhoneNumber:   "+8613812345678",
		PhoneCodeHash: "",
	})
	if err != tg.ErrPhoneCodeHashEmpty {
		t.Fatalf("expected ErrPhoneCodeHashEmpty, got %v", err)
	}
}

func TestAuthResendCodeReturnsPhoneNumberInvalidForBadPhone(t *testing.T) {
	c := New(context.Background(), nil)

	_, err := c.AuthResendCode(&tg.TLAuthResendCode{
		PhoneNumber:   "bad-phone",
		PhoneCodeHash: "hash",
	})
	if err != tg.Err406PhoneNumberInvalid {
		t.Fatalf("expected Err406PhoneNumberInvalid, got %v", err)
	}
}

func TestAuthResendCodeReturnsSentCodeAfterResend(t *testing.T) {
	c, ctx, d := newAuthorizationCoreForAuthTestWithVerifier(t, &kitexmetadata.RpcMetadata{
		PermAuthKeyId: 101,
		SessionId:     202,
		Layer:         181,
	}, &fakeVerifyCode{sendExtraData: "sent-via-test"})
	phoneNumber := seedPhoneCodeTransaction(t, ctx, d, 101, 202, false, "12345", "hash", model.CodeStateSent)

	result, err := c.AuthResendCode(&tg.TLAuthResendCode{
		PhoneNumber:   "+8613812345678",
		PhoneCodeHash: "hash",
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if result == nil {
		t.Fatal("expected auth sent code, got nil")
	}
	sentCode, ok := result.ToAuthSentCode()
	if !ok {
		t.Fatalf("expected auth.sentCode, got %T", result.Clazz)
	}
	if sentCode.PhoneCodeHash != "hash" {
		t.Fatalf("expected phone_code_hash hash, got %s", sentCode.PhoneCodeHash)
	}
	if _, ok := sentCode.Type.(*tg.TLAuthSentCodeTypeSms); !ok {
		t.Fatalf("expected auth.sentCodeTypeSms, got %T", sentCode.Type)
	}

	codeData, err := d.GetPhoneCode(ctx, 101, phoneNumber, "hash")
	if err != nil {
		t.Fatalf("get phone code after resend: %v", err)
	}
	if codeData.State != model.CodeStateSent {
		t.Fatalf("expected state %d, got %d", model.CodeStateSent, codeData.State)
	}
	if codeData.PhoneCodeExtraData != "sent-via-test" {
		t.Fatalf("expected extra data sent-via-test, got %s", codeData.PhoneCodeExtraData)
	}
	if codeData.SentCodeType != model.SentCodeTypeSms {
		t.Fatalf("expected sent code type %d, got %d", model.SentCodeTypeSms, codeData.SentCodeType)
	}
}
