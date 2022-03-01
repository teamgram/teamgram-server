// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
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

package codec

import (
	"errors"
	"fmt"

	"github.com/teamgram/proto/mtproto/crypto"
)

var (
	isClientType = false
)

const (
	ERROR                       = -1
	INVALID                     = 0
	WAIT_FIRST_PACKET           = 1
	WAIT_PACKET_LENGTH_1        = 2
	WAIT_PACKET_LENGTH_1_PACKET = 3
	WAIT_PACKET_LENGTH_3        = 4
	WAIT_PACKET_LENGTH_3_PACKET = 5
	WAIT_PACKET_LENGTH          = 6
	WAIT_PACKET                 = 7
)

const (
	MAX_MTPRORO_FRAME_SIZE = 16777216
)

var (
	errUnexpectedEOF = errors.New("there is no enough data")
)

// innerBuffer
type innerBuffer []byte

func (in *innerBuffer) readN(n int) (buf []byte, err error) {
	if n <= 0 {
		return nil, errors.New("zero or negative length is invalid")
	} else if n > len(*in) {
		return nil, errors.New("exceeding buffer length")
	}
	buf = (*in)[:n]
	*in = (*in)[n:]

	return
}

// MTPRawMessage MTPRawMessage
type MTPRawMessage struct {
	IsHttp     bool
	AuthKeyId  int64  // 由原始数据解压获得
	QuickAckId int32  // EncryptedMessage，则可能存在
	Payload    []byte // 原始数据
}

func NewMTPRawMessage(isHttp bool, authKeyId int64, quickAckId int32, payload []byte) *MTPRawMessage {
	return &MTPRawMessage{
		IsHttp:     isHttp,
		AuthKeyId:  authKeyId,
		QuickAckId: quickAckId,
		Payload:    payload,
	}
}

func (m *MTPRawMessage) String() string {
	return fmt.Sprintf("{is_http: %v, auth_key_id: %d, quick_ack_id: %d, payload_len: %d}",
		m.IsHttp,
		m.AuthKeyId,
		m.QuickAckId,
		len(m.Payload))
}

func (m *MTPRawMessage) GetIsHttp() bool {
	return m.IsHttp
}

func (m *MTPRawMessage) GetAuthKeyId() int64 {
	return m.AuthKeyId
}

func (m *MTPRawMessage) GetQuickAckId() int32 {
	return m.QuickAckId
}

// AesCTR128Crypto AesCTR128Crypto
type AesCTR128Crypto struct {
	decrypt *crypto.AesCTR128Encrypt
	encrypt *crypto.AesCTR128Encrypt
}

func newAesCTR128Crypto(d, e *crypto.AesCTR128Encrypt) *AesCTR128Crypto {
	return &AesCTR128Crypto{
		decrypt: d,
		encrypt: e,
	}
}
func (e *AesCTR128Crypto) Encrypt(plaintext []byte) []byte {
	if e == nil {
		return plaintext
	} else {
		return e.encrypt.Encrypt(plaintext)
	}
}

func (e *AesCTR128Crypto) Decrypt(plaintext []byte) []byte {
	if e == nil {
		return plaintext
	} else {
		return e.decrypt.Encrypt(plaintext)
	}
}
