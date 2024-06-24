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
)

const (
	QRCodeStateNew      = 1
	QRCodeStateAccepted = 2
	QRCodeStateSuccess  = 3
)

type QRCodeTransaction struct {
	PermAuthKeyId int64  `json:"perm_auth_key_id"`
	ServerId      string `json:"server_id"`
	AuthKeyId     int64  `json:"auth_key_id"`
	SessionId     int64  `json:"session_id"`
	ApiId         int32  `json:"api_id"`
	ApiHash       string `json:"api_hash"`
	CodeHash      string `json:"code_hash"`
	ExpireAt      int64  `json:"expire_at"`
	UserId        int64  `json:"user_id"`
	State         int    `json:"state"`
}

func (m *QRCodeTransaction) Token() []byte {
	token := make([]byte, 8, 24)
	binary.BigEndian.PutUint64(token, uint64(m.PermAuthKeyId))
	m2 := md5.New()
	io.WriteString(m2, strconv.Itoa(int(m.AuthKeyId)))
	io.WriteString(m2, m.CodeHash)
	io.WriteString(m2, strconv.Itoa(int(m.ExpireAt)))
	return m2.Sum(token)
}

func (m *QRCodeTransaction) CheckByToken(token []byte) bool {
	return bytes.Equal(m.Token(), token)
}
