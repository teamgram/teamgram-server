package core

import (
	"errors"
	"fmt"
	"time"

	codepb "github.com/teamgram/teamgram-server/v2/app/service/biz/code/code"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/phonenumber"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const fallbackAccountPhoneCodeLength = int32(5)

func requireSelfID(c *AccountCore) (int64, error) {
	if c == nil || c.MD == nil || c.MD.UserId <= 0 {
		return 0, tg.ErrUserIdInvalid
	}
	return c.MD.UserId, nil
}

func accountAuthKeyID(c *AccountCore) int64 {
	if c == nil || c.MD == nil {
		return 0
	}
	if c.MD.PermAuthKeyId != 0 {
		return c.MD.PermAuthKeyId
	}
	return c.MD.AuthId
}

func accountSessionID(c *AccountCore) int64 {
	if c == nil || c.MD == nil {
		return 0
	}
	return c.MD.SessionId
}

func requireUserClient(c *AccountCore) error {
	if c == nil || c.svcCtx == nil || c.svcCtx.Repo == nil || c.svcCtx.Repo.UserClient == nil {
		return fmt.Errorf("account: user client is nil")
	}
	return nil
}

func requireAuthsessionClient(c *AccountCore) error {
	if c == nil || c.svcCtx == nil || c.svcCtx.Repo == nil || c.svcCtx.Repo.AuthsessionClient == nil {
		return fmt.Errorf("account: authsession client is nil")
	}
	return nil
}

func requireCodeClient(c *AccountCore) error {
	if c == nil || c.svcCtx == nil || c.svcCtx.Repo == nil || c.svcCtx.Repo.CodeClient == nil {
		return fmt.Errorf("account: code client is nil")
	}
	return nil
}

func normalizeAccountPhone(phone string) (string, error) {
	_, normalized, err := phonenumber.CheckPhoneNumberInvalid(phone)
	if err != nil {
		return "", err
	}
	return normalized, nil
}

func isUserNotFound(err error) bool {
	return errors.Is(err, userpb.ErrUserNotFound)
}

func mapPhoneCodeError(err error) error {
	switch {
	case err == nil:
		return nil
	case errors.Is(err, codepb.ErrPhoneCodeExpired):
		return tg.ErrPhoneCodeExpired
	case errors.Is(err, codepb.ErrPhoneCodeInvalid):
		return tg.ErrPhoneCodeInvalid
	default:
		return err
	}
}

func makeAccountSentCode(codeData *codepb.PhoneCodeTransaction) (*tg.AuthSentCode, error) {
	if codeData == nil {
		return nil, fmt.Errorf("account: code service returned nil phone code transaction")
	}
	length := int32(len(codeData.PhoneCode))
	if length == 0 {
		length = fallbackAccountPhoneCodeLength
	}
	timeout := int32(60)
	if codeData.PhoneCodeExpired > 0 {
		remaining := codeData.PhoneCodeExpired - int32(time.Now().Unix())
		if remaining > 0 {
			timeout = remaining
		}
	}
	return tg.MakeTLAuthSentCode(&tg.TLAuthSentCode{
		Type:          tg.MakeTLAuthSentCodeTypeApp(&tg.TLAuthSentCodeTypeApp{Length: length}),
		PhoneCodeHash: codeData.PhoneCodeHash,
		Timeout:       &timeout,
	}).ToAuthSentCode(), nil
}
