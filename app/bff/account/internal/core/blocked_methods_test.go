package core

import (
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestAccountSendConfirmPhoneCodeEnterpriseBlocked(t *testing.T) {
	c := newAccountCoreForTest(nil, nil, nil)
	_, err := c.AccountSendConfirmPhoneCode(&tg.TLAccountSendConfirmPhoneCode{})
	if !errors.Is(err, tg.ErrEnterpriseIsBlocked) {
		t.Fatalf("error = %v, want ErrEnterpriseIsBlocked", err)
	}
}

func TestAccountConfirmPhoneEnterpriseBlocked(t *testing.T) {
	c := newAccountCoreForTest(nil, nil, nil)
	_, err := c.AccountConfirmPhone(&tg.TLAccountConfirmPhone{})
	if !errors.Is(err, tg.ErrEnterpriseIsBlocked) {
		t.Fatalf("error = %v, want ErrEnterpriseIsBlocked", err)
	}
}
