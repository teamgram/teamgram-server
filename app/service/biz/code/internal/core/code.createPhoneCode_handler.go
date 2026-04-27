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
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/code/code"
)

const (
	codeStateSend      = 1
	codeLen            = 5
	codeHashLen        = 16
	codeExpireDuration = 3 * 60 // seconds
)

func randomNumeric(n int) string {
	const digits = "0123456789"
	b := make([]byte, n)
	for i := range b {
		idx, _ := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		b[i] = digits[idx.Int64()]
	}
	return string(b)
}

func randomHexString(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

// CodeCreatePhoneCode
// code.createPhoneCode flags:# auth_key_id:long session_id:long phone:string phone_number_registered:flags.0?true sent_code_type:int next_code_type:int state:int = PhoneCodeTransaction;
func (c *CodeCore) CodeCreatePhoneCode(in *code.TLCodeCreatePhoneCode) (*code.PhoneCodeTransaction, error) {
	codeData, err := c.repo.GetCachePhoneCode(c.ctx, in.AuthKeyId, in.Phone)
	if err != nil {
		c.Logger.Errorf("code.createPhoneCode - get cache failed: auth_key_id: %d, phone: %s, err: %v",
			in.AuthKeyId, in.Phone, err)
	}

	if codeData == nil || in.SessionId != codeData.SessionId {
		codeData = code.MakeTLPhoneCodeTransaction(&code.TLPhoneCodeTransaction{
			AuthKeyId:             in.AuthKeyId,
			SessionId:             in.SessionId,
			Phone:                 in.Phone,
			PhoneNumberRegistered: in.PhoneNumberRegistered,
			PhoneCode:             randomNumeric(codeLen),
			PhoneCodeHash:         randomHexString(codeHashLen),
			PhoneCodeExpired:      int32(time.Now().Unix() + codeExpireDuration),
			SentCodeType:          in.SentCodeType,
			FlashCallPattern:      "*",
			NextCodeType:          in.NextCodeType,
			State:                 codeStateSend,
		})
	}

	return codeData, nil
}
