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

package server

import (
	"github.com/teamgram/proto/mtproto"
)

func parseFromIncomingMessage(b []byte) (msgId int64, obj mtproto.TLObject, err error) {
	dBuf := mtproto.NewDecodeBuf(b)

	msgId = dBuf.Long()
	_ = dBuf.Int()
	obj = dBuf.Object()
	err = dBuf.GetError()

	return
}

func serializeToBuffer(x *mtproto.EncodeBuf, msgId int64, obj mtproto.TLObject) error {
	//obj.Encode(x, 0)
	// x := mtproto.NewEncodeBuf(8 + 4 + len(oBuf))
	x.Long(0)
	x.Long(msgId)
	offset := x.GetOffset()
	x.Int(0)
	err := obj.Encode(x, 0)
	if err != nil {
		return err
	}
	//x.Bytes(oBuf)
	x.IntOffset(offset, int32(x.GetOffset()-offset-4))
	return nil
}

const (
	kMsgContainerBufLen = 8
)

var (
	kMsgContainerBuf = func() []byte {
		x := mtproto.NewEncodeBuf(8)
		x.Int(int32(mtproto.CRC32_msg_container))
		x.Int(0)
		return x.GetBuf()
	}()
)

func serializeToBuffer2(salt, sessionId int64, seqNo int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	x.Long(salt)
	x.Long(sessionId)
	x.Long(mtproto.GenerateMessageId())
	x.Int(seqNo)
	x.Int(kMsgContainerBufLen)
	x.Bytes(kMsgContainerBuf)

	return x.GetBuf()
}
