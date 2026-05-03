package core

import (
	"context"
	"errors"
	"testing"
	"time"

	codepb "github.com/teamgram/teamgram-server/v2/app/service/biz/code/code"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestAccountSendChangePhoneCodeRejectsOccupiedPhone(t *testing.T) {
	c := newAccountCoreForTest(&fakeAccountUserClient{
		getImmutableUserByPhone: func(_ context.Context, in *userpb.TLUserGetImmutableUserByPhone) (*tg.ImmutableUser, error) {
			if in.Phone != "8613711112222" {
				t.Fatalf("phone = %q, want 8613711112222", in.Phone)
			}
			return accountImmutableUserFixture(2002, in.Phone, ""), nil
		},
	}, nil, &fakeAccountCodeClient{
		createPhoneCode: func(context.Context, *codepb.TLCodeCreatePhoneCode) (*codepb.PhoneCodeTransaction, error) {
			t.Fatal("CodeCreatePhoneCode should not be called for occupied phone")
			return nil, nil
		},
	})

	_, err := c.AccountSendChangePhoneCode(&tg.TLAccountSendChangePhoneCode{
		PhoneNumber: "+86 137 1111 2222",
		Settings:    tg.MakeTLCodeSettings(&tg.TLCodeSettings{}).ToCodeSettings(),
	})
	if !errors.Is(err, tg.ErrPhoneNumberOccupied) {
		t.Fatalf("error = %v, want ErrPhoneNumberOccupied", err)
	}
}

func TestAccountSendChangePhoneCodeCreatesCodeTransaction(t *testing.T) {
	var gotCreate *codepb.TLCodeCreatePhoneCode
	c := newAccountCoreForTest(&fakeAccountUserClient{
		getImmutableUserByPhone: func(context.Context, *userpb.TLUserGetImmutableUserByPhone) (*tg.ImmutableUser, error) {
			return nil, userpb.ErrUserNotFound
		},
	}, nil, &fakeAccountCodeClient{
		createPhoneCode: func(_ context.Context, in *codepb.TLCodeCreatePhoneCode) (*codepb.PhoneCodeTransaction, error) {
			gotCreate = in
			return codepb.MakeTLPhoneCodeTransaction(&codepb.TLPhoneCodeTransaction{
				AuthKeyId:        in.AuthKeyId,
				SessionId:        in.SessionId,
				Phone:            in.Phone,
				PhoneCode:        "12345",
				PhoneCodeHash:    "hash-1",
				PhoneCodeExpired: int32(time.Now().Unix() + 120),
				SentCodeType:     1,
				NextCodeType:     1,
				State:            1,
				FlashCallPattern: "*",
			}).ToPhoneCodeTransaction(), nil
		},
	})

	got, err := c.AccountSendChangePhoneCode(&tg.TLAccountSendChangePhoneCode{
		PhoneNumber: "+86 137 1111 2222",
		Settings:    tg.MakeTLCodeSettings(&tg.TLCodeSettings{}).ToCodeSettings(),
	})
	if err != nil {
		t.Fatalf("AccountSendChangePhoneCode error = %v", err)
	}
	if gotCreate == nil {
		t.Fatal("CodeCreatePhoneCode was not called")
	}
	if gotCreate.AuthKeyId != 9001 || gotCreate.SessionId != 7001 || gotCreate.Phone != "8613711112222" {
		t.Fatalf("create request = %+v, want auth_key_id 9001 session_id 7001 phone 8613711112222", gotCreate)
	}
	sent, ok := got.ToAuthSentCode()
	if !ok {
		t.Fatalf("result = %s, want auth.sentCode", got.ClazzName())
	}
	if sent.PhoneCodeHash != "hash-1" {
		t.Fatalf("phone_code_hash = %q, want hash-1", sent.PhoneCodeHash)
	}
	appType, ok := sent.Type.(*tg.TLAuthSentCodeTypeApp)
	if !ok {
		t.Fatalf("sent code type = %T, want *tg.TLAuthSentCodeTypeApp", sent.Type)
	}
	if appType.Length != 5 {
		t.Fatalf("app code length = %d, want 5", appType.Length)
	}
}
