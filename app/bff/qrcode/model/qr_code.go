// Copyright © 2025 The Teamgram Authors.
//  All Rights Reserved.
//
// Author: Benqi (wubenqi@gmail.com)

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
