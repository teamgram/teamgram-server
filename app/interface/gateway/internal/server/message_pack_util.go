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

func serializeToBuffer(msgId int64, obj mtproto.TLObject) []byte {
	oBuf := obj.Encode(0)
	x := mtproto.NewEncodeBuf(8 + 4 + len(oBuf))
	x.Long(0)
	x.Long(msgId)
	x.Int(int32(len(oBuf)))
	x.Bytes(oBuf)

	return x.GetBuf()
}
