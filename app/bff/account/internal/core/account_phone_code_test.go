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

func TestAccountChangePhoneMapsExpiredCode(t *testing.T) {
	c := newAccountCoreForTest(&fakeAccountUserClient{
		getImmutableUserByPhone: func(context.Context, *userpb.TLUserGetImmutableUserByPhone) (*tg.ImmutableUser, error) {
			return nil, userpb.ErrUserNotFound
		},
	}, nil, &fakeAccountCodeClient{
		getPhoneCode: func(context.Context, *codepb.TLCodeGetPhoneCode) (*codepb.PhoneCodeTransaction, error) {
			return nil, codepb.ErrPhoneCodeExpired
		},
	})

	_, err := c.AccountChangePhone(&tg.TLAccountChangePhone{
		PhoneNumber:   "+86 137 1111 2222",
		PhoneCodeHash: "hash-1",
		PhoneCode:     "12345",
	})
	if !errors.Is(err, tg.ErrPhoneCodeExpired) {
		t.Fatalf("error = %v, want ErrPhoneCodeExpired", err)
	}
}

func TestAccountChangePhoneRejectsWrongCode(t *testing.T) {
	c := newAccountCoreForTest(&fakeAccountUserClient{
		getImmutableUserByPhone: func(context.Context, *userpb.TLUserGetImmutableUserByPhone) (*tg.ImmutableUser, error) {
			return nil, userpb.ErrUserNotFound
		},
	}, nil, &fakeAccountCodeClient{
		getPhoneCode: func(context.Context, *codepb.TLCodeGetPhoneCode) (*codepb.PhoneCodeTransaction, error) {
			return codepb.MakeTLPhoneCodeTransaction(&codepb.TLPhoneCodeTransaction{
				Phone:              "8613711112222",
				PhoneCode:          "12345",
				PhoneCodeHash:      "hash-1",
				PhoneCodeExpired:   int32(time.Now().Unix() + 60),
				PhoneCodeExtraData: "12345",
				FlashCallPattern:   "*",
			}).ToPhoneCodeTransaction(), nil
		},
	})

	_, err := c.AccountChangePhone(&tg.TLAccountChangePhone{
		PhoneNumber:   "+86 137 1111 2222",
		PhoneCodeHash: "hash-1",
		PhoneCode:     "54321",
	})
	if !errors.Is(err, tg.ErrPhoneCodeInvalid) {
		t.Fatalf("error = %v, want ErrPhoneCodeInvalid", err)
	}
}

func TestAccountChangePhoneUpdatesUserAndDeletesCode(t *testing.T) {
	var gotChange *userpb.TLUserChangePhone
	var gotDelete *codepb.TLCodeDeletePhoneCode
	var gotProjection *userpb.TLUserGetUserProjectionBundle
	c := newAccountCoreForTest(&fakeAccountUserClient{
		getImmutableUserByPhone: func(context.Context, *userpb.TLUserGetImmutableUserByPhone) (*tg.ImmutableUser, error) {
			return nil, userpb.ErrUserNotFound
		},
		getImmutableUser: func(context.Context, *userpb.TLUserGetImmutableUser) (*tg.ImmutableUser, error) {
			return accountImmutableUserFixture(1001, "8613000000000", "self"), nil
		},
		changePhone: func(_ context.Context, in *userpb.TLUserChangePhone) (*tg.Bool, error) {
			gotChange = in
			return tg.BoolTrue, nil
		},
		getUserProjection: func(_ context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
			gotProjection = in
			return userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
				ViewerUsers: []userpb.ViewerUsersClazz{
					userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 1001, Users: []tg.UserClazz{
						tg.MakeTLUser(&tg.TLUser{Id: 1001, Self: true, Phone: strPtr("8613711112222")}),
					}}),
				},
			}).ToUserProjectionBundle(), nil
		},
	}, nil, &fakeAccountCodeClient{
		getPhoneCode: func(_ context.Context, in *codepb.TLCodeGetPhoneCode) (*codepb.PhoneCodeTransaction, error) {
			return codepb.MakeTLPhoneCodeTransaction(&codepb.TLPhoneCodeTransaction{
				AuthKeyId:          in.AuthKeyId,
				Phone:              in.Phone,
				PhoneCode:          "12345",
				PhoneCodeHash:      in.PhoneCodeHash,
				PhoneCodeExpired:   int32(time.Now().Unix() + 60),
				PhoneCodeExtraData: "12345",
				FlashCallPattern:   "*",
			}).ToPhoneCodeTransaction(), nil
		},
		deletePhoneCode: func(_ context.Context, in *codepb.TLCodeDeletePhoneCode) (*tg.Bool, error) {
			gotDelete = in
			return tg.BoolTrue, nil
		},
	})

	got, err := c.AccountChangePhone(&tg.TLAccountChangePhone{
		PhoneNumber:   "+86 137 1111 2222",
		PhoneCodeHash: "hash-1",
		PhoneCode:     "12345",
	})
	if err != nil {
		t.Fatalf("AccountChangePhone error = %v", err)
	}
	if gotChange == nil || gotChange.UserId != 1001 || gotChange.Phone != "8613711112222" {
		t.Fatalf("change request = %+v, want user_id 1001 phone 8613711112222", gotChange)
	}
	if gotDelete == nil || gotDelete.AuthKeyId != 9001 || gotDelete.Phone != "8613711112222" || gotDelete.PhoneCodeHash != "hash-1" {
		t.Fatalf("delete request = %+v, want auth_key_id 9001 phone 8613711112222 hash-1", gotDelete)
	}
	if gotProjection == nil || len(gotProjection.ViewerUserIds) != 1 || gotProjection.ViewerUserIds[0] != 1001 ||
		len(gotProjection.TargetUserIds) != 1 || gotProjection.TargetUserIds[0] != 1001 {
		t.Fatalf("projection request = %+v, want viewer/target 1001", gotProjection)
	}
	user, ok := got.ToUser()
	if !ok {
		t.Fatalf("result = %s, want user", got.ClazzName())
	}
	if user.Phone == nil || *user.Phone != "8613711112222" {
		t.Fatalf("returned phone = %v, want 8613711112222", user.Phone)
	}
	if !user.Self {
		t.Fatal("returned user Self = false, want true")
	}
}

func strPtr(v string) *string {
	return &v
}
