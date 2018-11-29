// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
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

// Author: Benqi (wubenqi@gmail.com)

package mtproto

import (
	"fmt"
)

type MessageBase interface {
	Encode() []byte
	Decode(b []byte) error
}

//
//type CodecAble interface {
//	Encode() ([]byte, error)
//	Decode(dbuf *DecodeBuf) error
//}

func NewMTPRawMessage(authKeyId int64, quickAckId int32, connType int) *MTPRawMessage {
	return &MTPRawMessage{
		connType:   connType,
		authKeyId:  authKeyId,
		quickAckId: quickAckId,
	}
}

////////////////////////////////////////////////////////////////////////////
// 代理使用
type MTPRawMessage struct {
	connType   int
	authKeyId  int64 // 由原始数据解压获得
	quickAckId int32 // EncryptedMessage，则可能存在

	// 原始数据
	Payload []byte
}

func (m *MTPRawMessage) String() string {
	return fmt.Sprintf("{conn_type: %d, auth_key_id: %d, quick_ack_id: %d, payload_len: %d}",
		m.connType,
		m.authKeyId,
		m.quickAckId,
		len(m.Payload))
}

func (m *MTPRawMessage) ConnType() int {
	return m.connType
}

func (m *MTPRawMessage) AuthKeyId() int64 {
	return m.authKeyId
}

func (m *MTPRawMessage) QuickAckId() int32 {
	return m.quickAckId
}

////////////////////////////////////////////////////////////////////////////
func (m *MTPRawMessage) Encode() []byte {
	return m.Payload
}

func (m *MTPRawMessage) Decode(b []byte) error {
	m.Payload = b
	return nil
}

////////////////////////////////////////////////////////////////////////////
func NewUnencryptedRawMessage() *UnencryptedRawMessage {
	return &UnencryptedRawMessage{
		AuthKeyId: 0,
	}
}

type UnencryptedRawMessage struct {
	// TODO(@benqi): reportAck and quickAck
	// NeedAck bool
	AuthKeyId   int64
	MessageId   int64
	MessageData []byte
}

func (m *UnencryptedRawMessage) Encode() []byte {
	// 一次性分配
	x := NewEncodeBuf(20 + len(m.MessageData))
	x.Long(0)
	m.MessageId = GenerateMessageId()
	x.Long(m.MessageId)
	x.Int(int32(len(m.MessageData)))
	x.Bytes(m.MessageData)
	return x.buf
}

func (m *UnencryptedRawMessage) Decode(b []byte) error {
	dbuf := NewDecodeBuf(b)
	m.MessageId = dbuf.Long()
	messageLen := dbuf.Int()
	if int(messageLen) != dbuf.size-12 {
		return fmt.Errorf("invalid UnencryptedRawMessage len: %d (need %d)", messageLen, dbuf.size-12)
	}
	m.MessageData = dbuf.Bytes(int(messageLen))
	return dbuf.err
}

type EncryptedRawMessage struct {
	// TODO(@benqi): reportAck and quickAck
	// NeedAck bool
	AuthKeyId     int64
	MsgKey        []byte
	EncryptedData []byte
}

func NewEncryptedRawMessage(authKeyId int64) *EncryptedRawMessage {
	return &EncryptedRawMessage{
		AuthKeyId: authKeyId,
	}
}

func (m *EncryptedRawMessage) Encode() []byte {
	// 一次性分配
	x := NewEncodeBuf(24 + len(m.EncryptedData))
	x.Long(m.AuthKeyId)
	x.Bytes(m.MsgKey)
	x.Bytes(m.EncryptedData)
	return x.buf
}

func (m *EncryptedRawMessage) Decode(b []byte) error {
	dbuf := NewDecodeBuf(b)
	m.MsgKey = dbuf.Bytes(16)
	m.EncryptedData = dbuf.Bytes(len(b) - 16)
	return dbuf.err
}
