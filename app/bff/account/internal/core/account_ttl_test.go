package core

import (
	"context"
	"errors"
	"testing"

	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestAccountGetAccountTTLDelegatesToUserService(t *testing.T) {
	var gotUserID int64
	c := newAccountCoreForTest(&fakeAccountUserClient{
		getAccountDaysTTL: func(_ context.Context, in *userpb.TLUserGetAccountDaysTTL) (*tg.AccountDaysTTL, error) {
			gotUserID = in.UserId
			return tg.MakeTLAccountDaysTTL(&tg.TLAccountDaysTTL{Days: 365}).ToAccountDaysTTL(), nil
		},
	}, nil, nil)

	got, err := c.AccountGetAccountTTL(&tg.TLAccountGetAccountTTL{})
	if err != nil {
		t.Fatalf("AccountGetAccountTTL error = %v", err)
	}
	if gotUserID != 1001 {
		t.Fatalf("user_id = %d, want 1001", gotUserID)
	}
	if got.Days != 365 {
		t.Fatalf("Days = %d, want 365", got.Days)
	}
}

func TestAccountSetAccountTTLRejectsInvalidDays(t *testing.T) {
	c := newAccountCoreForTest(&fakeAccountUserClient{
		setAccountDaysTTL: func(context.Context, *userpb.TLUserSetAccountDaysTTL) (*tg.Bool, error) {
			t.Fatal("UserSetAccountDaysTTL should not be called for invalid TTL")
			return nil, nil
		},
	}, nil, nil)

	_, err := c.AccountSetAccountTTL(&tg.TLAccountSetAccountTTL{
		Ttl: tg.MakeTLAccountDaysTTL(&tg.TLAccountDaysTTL{Days: 31}).ToAccountDaysTTL(),
	})
	if !errors.Is(err, tg.ErrTtlDaysInvalid) {
		t.Fatalf("error = %v, want ErrTtlDaysInvalid", err)
	}
}

func TestAccountSetAccountTTLDelegatesValidDays(t *testing.T) {
	var gotUserID int64
	var gotTTL int32
	c := newAccountCoreForTest(&fakeAccountUserClient{
		setAccountDaysTTL: func(_ context.Context, in *userpb.TLUserSetAccountDaysTTL) (*tg.Bool, error) {
			gotUserID = in.UserId
			gotTTL = in.Ttl
			return tg.BoolTrue, nil
		},
	}, nil, nil)

	got, err := c.AccountSetAccountTTL(&tg.TLAccountSetAccountTTL{
		Ttl: tg.MakeTLAccountDaysTTL(&tg.TLAccountDaysTTL{Days: 180}).ToAccountDaysTTL(),
	})
	if err != nil {
		t.Fatalf("AccountSetAccountTTL error = %v", err)
	}
	if got != tg.BoolTrue {
		t.Fatalf("result = %v, want BoolTrue", got)
	}
	if gotUserID != 1001 || gotTTL != 180 {
		t.Fatalf("request = user_id:%d ttl:%d, want user_id:1001 ttl:180", gotUserID, gotTTL)
	}
}
