package core

import (
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/authorization/model"
	kitexmetadata "github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestAuthSendCodeReturnsPhoneNumberInvalidForBadPhone(t *testing.T) {
	c, _, _ := newAuthorizationCoreForAuthTest(t, &kitexmetadata.RpcMetadata{
		PermAuthKeyId: 101,
		SessionId:     202,
		Layer:         181,
	})

	_, err := c.AuthSendCode(&tg.TLAuthSendCode{
		PhoneNumber: "bad-phone",
		ApiId:       1,
		ApiHash:     "hash",
		Settings:    tg.MakeTLCodeSettings(&tg.TLCodeSettings{}),
	})
	if err != tg.Err406PhoneNumberInvalid {
		t.Fatalf("expected Err406PhoneNumberInvalid, got %v", err)
	}
}

func TestAuthSendCodeReturnsSentCodeForUnregisteredPhone(t *testing.T) {
	c, ctx, d := newAuthorizationCoreForAuthTestWithVerifier(t, &kitexmetadata.RpcMetadata{
		PermAuthKeyId: 101,
		SessionId:     202,
		Layer:         181,
	}, &fakeVerifyCode{sendExtraData: "sms-extra"})

	result, err := c.AuthSendCode(&tg.TLAuthSendCode{
		PhoneNumber: "+8613812345678",
		ApiId:       1,
		ApiHash:     "hash",
		Settings: tg.MakeTLCodeSettings(&tg.TLCodeSettings{
			AllowFlashcall: false,
			CurrentNumber:  false,
		}),
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	sentCode, ok := result.ToAuthSentCode()
	if !ok {
		t.Fatalf("expected auth.sentCode, got %T", result.Clazz)
	}
	if _, ok := sentCode.Type.(*tg.TLAuthSentCodeTypeSms); !ok {
		t.Fatalf("expected auth.sentCodeTypeSms, got %T", sentCode.Type)
	}

	_, phoneNumber, err := checkPhoneNumberInvalid("+8613812345678")
	if err != nil {
		t.Fatalf("normalize phone: %v", err)
	}

	codeData, err := d.GetPhoneCode(ctx, 101, phoneNumber, sentCode.PhoneCodeHash)
	if err != nil {
		t.Fatalf("get phone code after sendCode: %v", err)
	}
	if codeData.State != model.CodeStateSent {
		t.Fatalf("expected state %d, got %d", model.CodeStateSent, codeData.State)
	}
	if codeData.PhoneNumberRegistered {
		t.Fatal("expected unregistered phone")
	}
	if codeData.PhoneCodeExtraData != "sms-extra" {
		t.Fatalf("expected extra data sms-extra, got %s", codeData.PhoneCodeExtraData)
	}
	if codeData.SentCodeType != model.SentCodeTypeSms {
		t.Fatalf("expected sent code type %d, got %d", model.SentCodeTypeSms, codeData.SentCodeType)
	}
}
