// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package model

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"io"
	"strconv"

	"github.com/teamgram/proto/mtproto"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

// PhoneCodeTransaction
// TODO(@benqi): Add phone region
type PhoneCodeTransaction struct {
	AuthKeyId             int64  `json:"auth_key_id"`
	SessionId             int64  `json:"session_id"`
	PhoneNumber           string `json:"phone_number"`
	PhoneNumberRegistered bool   `json:"phone_number_registered"`
	PhoneCode             string `json:"phone_code"`
	PhoneCodeHash         string `json:"phone_code_hash"`
	PhoneCodeExpired      int32  `json:"phone_code_expired"`
	PhoneCodeExtraData    string `json:"phone_code_extra_data"`
	SentCodeType          int    `json:"sent_code_type"`
	FlashCallPattern      string `json:"flash_call_pattern"`
	NextCodeType          int    `json:"next_code_type"`
	State                 int    `json:"state"`
}

// ToAuthSentCode
// TODO(@benqi): 如果手机号已经注册，检查是否有其他设备在线，有则使用sentCodeTypeApp
//
//	否则使用sentCodeTypeSms
//
// TODO(@benqi): 有则使用sentCodeTypeFlashCall和entCodeTypeCall？？
func (m *PhoneCodeTransaction) ToAuthSentCode() *mtproto.Auth_SentCode {
	// TODO(@benqi): only use sms

	authSentCode := mtproto.MakeTLAuthSentCode(&mtproto.Auth_SentCode{
		Type:          makeAuthSentCodeType(m.SentCodeType, len(m.PhoneCode), m.FlashCallPattern),
		PhoneCodeHash: m.PhoneCodeHash,
		NextType:      makeAuthCodeType(m.NextCodeType),
		Timeout:       &wrapperspb.Int32Value{Value: 60}, // TODO(@benqi): 默认60s
	}).To_Auth_SentCode()
	if m.SentCodeType == SentCodeTypeApp {
		authSentCode.Timeout = nil
	}
	return authSentCode
}

const (
	QRCodeStateNew      = 1
	QRCodeStateAccepted = 2
	QRCodeStateSuccess  = 3
)

type QRCodeTransaction struct {
	AuthKeyId int64  `json:"auth_key_id"`
	ServerId  string `json:"server_id"`
	SessionId int64  `json:"session_id"`
	ApiId     int32  `json:"api_id"`
	ApiHash   string `json:"api_hash"`
	CodeHash  string `json:"code_hash"`
	ExpireAt  int64  `json:"expire_at"`
	UserId    int64  `json:"user_id"`
	State     int    `json:"state"`
}

func (m *QRCodeTransaction) Token() []byte {
	token := make([]byte, 8, 24)
	binary.BigEndian.PutUint64(token, uint64(m.AuthKeyId))
	m2 := md5.New()
	io.WriteString(m2, strconv.Itoa(int(m.AuthKeyId)))
	io.WriteString(m2, m.CodeHash)
	io.WriteString(m2, strconv.Itoa(int(m.ExpireAt)))
	return m2.Sum(token)
}

func (m *QRCodeTransaction) CheckByToken(token []byte) bool {
	return bytes.Equal(m.Token(), token)
}
