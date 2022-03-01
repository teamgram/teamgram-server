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

package service

import (
	"github.com/teamgram/proto/mtproto"
)

func ParseFromIncomingMessage(b []byte) (salt, sessionId int64, msg2 *mtproto.TLMessage2, err error) {
	dbuf := mtproto.NewDecodeBuf(b)
	msg2 = &mtproto.TLMessage2{}

	salt = dbuf.Long()      // salt
	sessionId = dbuf.Long() // session_id
	err = msg2.Decode(dbuf)

	return
}

func SerializeToBuffer(salt, sessionId int64, msg2 *mtproto.TLMessage2, layer int32) []byte {
	oBuf := msg2.Encode(layer)

	x := mtproto.NewEncodeBuf(16 + len(oBuf))
	x.Long(salt)
	x.Long(sessionId)
	x.Bytes(oBuf)

	return x.GetBuf()
}

func SerializeToBuffer2(salt, sessionId int64, msg2 *mtproto.TLMessageRawData) []byte {
	x := mtproto.NewEncodeBuf(32 + len(msg2.Body))

	x.Long(salt)
	x.Long(sessionId)
	x.Long(msg2.MsgId)
	x.Int(msg2.Seqno)
	x.Int(msg2.Bytes)
	x.Bytes(msg2.Body)

	return x.GetBuf()
}
