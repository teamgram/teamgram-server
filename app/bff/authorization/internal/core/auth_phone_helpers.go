// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package core

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const (
	startupAuthPhoneCode       = "22222"
	startupAuthPhoneCodeLength = int32(len(startupAuthPhoneCode))
	startupAuthSecretKeySalt   = int64(0x544741555448)
)

type startupPhoneCodeTransaction struct {
	Phone       string
	CountryCode string
	Code        string
}

var startupPhoneCodes = struct {
	sync.Mutex
	byHash map[string]startupPhoneCodeTransaction
}{
	byHash: map[string]startupPhoneCodeTransaction{},
}

func normalizeStartupPhone(phone string) (countryCode string, normalized string, err error) {
	var b strings.Builder
	for _, r := range phone {
		switch {
		case unicode.IsDigit(r):
			b.WriteRune(r)
		case r == '+' || unicode.IsSpace(r) || r == '-' || r == '(' || r == ')':
		default:
			return "", "", tg.Err406PhoneNumberInvalid
		}
	}

	normalized = b.String()
	if len(normalized) < 8 {
		return "", "", tg.Err406PhoneNumberInvalid
	}
	if strings.HasPrefix(normalized, "86") {
		countryCode = "CN"
	} else {
		countryCode = "ZZ"
	}
	return countryCode, normalized, nil
}

func storeStartupPhoneCode(phone string, countryCode string) string {
	seed := phone + ":" + time.Now().UTC().Format(time.RFC3339Nano)
	hashBytes := sha256.Sum256([]byte(seed))
	hash := hex.EncodeToString(hashBytes[:16])

	startupPhoneCodes.Lock()
	startupPhoneCodes.byHash[hash] = startupPhoneCodeTransaction{
		Phone:       phone,
		CountryCode: countryCode,
		Code:        startupAuthPhoneCode,
	}
	startupPhoneCodes.Unlock()

	return hash
}

func verifyStartupPhoneCode(phone string, hash string, code string) (countryCode string, err error) {
	if hash == "" {
		return "", tg.ErrPhoneCodeHashEmpty
	}
	if code == "" {
		return "", tg.ErrPhoneCodeEmpty
	}

	startupPhoneCodes.Lock()
	tx, ok := startupPhoneCodes.byHash[hash]
	startupPhoneCodes.Unlock()

	if !ok || tx.Phone != phone {
		return "", tg.ErrPhoneCodeExpired
	}
	if tx.Code != code {
		return "", tg.ErrPhoneCodeInvalid
	}
	return tx.CountryCode, nil
}

func lookupStartupPhoneCode(hash string) (startupPhoneCodeTransaction, error) {
	if hash == "" {
		return startupPhoneCodeTransaction{}, tg.ErrPhoneCodeHashEmpty
	}

	startupPhoneCodes.Lock()
	tx, ok := startupPhoneCodes.byHash[hash]
	startupPhoneCodes.Unlock()

	if !ok {
		return startupPhoneCodeTransaction{}, tg.ErrPhoneCodeExpired
	}
	return tx, nil
}

func deleteStartupPhoneCode(hash string) {
	startupPhoneCodes.Lock()
	delete(startupPhoneCodes.byHash, hash)
	startupPhoneCodes.Unlock()
}

func makeAuthSentCode(phoneCodeHash string) *tg.AuthSentCode {
	timeout := int32(60)
	return tg.MakeTLAuthSentCode(&tg.TLAuthSentCode{
		Type:          tg.MakeTLAuthSentCodeTypeSms(&tg.TLAuthSentCodeTypeSms{Length: startupAuthPhoneCodeLength}),
		PhoneCodeHash: phoneCodeHash,
		Timeout:       &timeout,
	}).ToAuthSentCode()
}

func makeSignupRequired() *tg.AuthAuthorization {
	return tg.MakeTLAuthAuthorizationSignUpRequired(&tg.TLAuthAuthorizationSignUpRequired{}).ToAuthAuthorization()
}

func makeAuthAuthorization(user tg.UserClazz) *tg.AuthAuthorization {
	if user == nil {
		return nil
	}
	return tg.MakeTLAuthAuthorization(&tg.TLAuthAuthorization{
		User: user,
	}).ToAuthAuthorization()
}

func startupAuthKeyID(c *AuthorizationCore) int64 {
	if c.MD == nil {
		return 0
	}
	if c.MD.PermAuthKeyId != 0 {
		return c.MD.PermAuthKeyId
	}
	return c.MD.AuthId
}

func startupUserID(c *AuthorizationCore) int64 {
	if c.MD == nil {
		return 0
	}
	return c.MD.UserId
}

func startupSecretKeyID(authKeyID int64) int64 {
	if authKeyID == 0 {
		return startupAuthSecretKeySalt
	}
	return authKeyID
}

func immutableUserID(user *tg.ImmutableUser) int64 {
	if user == nil || user.User == nil {
		return 0
	}
	return user.User.Id
}
