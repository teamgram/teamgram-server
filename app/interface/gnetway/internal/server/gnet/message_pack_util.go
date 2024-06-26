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

package gnet

import (
	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/core/logx"
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

func getRpcMethod(in mtproto.TLObject) mtproto.TLObject {
	if in == nil {
		return nil
	}
	logx.Debugf("rpc: %s", in)

	switch r := in.(type) {
	case *mtproto.TLDestroyAuthKey: // 所有连接都有可能
		return nil
	case *mtproto.TLRpcDropAnswer: // 所有连接都有可能
		return nil
	case *mtproto.TLGetFutureSalts: // GENERIC
		return nil
	case *mtproto.TLPing: // android未用
		return nil
	case *mtproto.TLPingDelayDisconnect: // PUSH和GENERIC
		return nil
	case *mtproto.TLDestroySession: // GENERIC
		return nil
	case *mtproto.TLMsgsStateReq: // android未用
		return nil
	case *mtproto.TLMsgsStateInfo: // android未用
		return nil
	case *mtproto.TLMsgsAllInfo: // android未用
		return nil
	case *mtproto.TLMsgResendReq: // 都有可能
		return nil
	case *mtproto.TLMsgDetailedInfo: // 都有可能
		return nil
	case *mtproto.TLMsgNewDetailedInfo: // 都有可能
		return nil
	case *mtproto.TLInvokeWithLayer:
		dBuf := mtproto.NewDecodeBuf(r.Query)
		return getRpcMethod(dBuf.Object())
	case *mtproto.TLInvokeAfterMsg:
		dBuf := mtproto.NewDecodeBuf(r.Query)
		return getRpcMethod(dBuf.Object())
	case *mtproto.TLInvokeAfterMsgs:
		dBuf := mtproto.NewDecodeBuf(r.Query)
		return getRpcMethod(dBuf.Object())
	case *mtproto.TLInvokeWithoutUpdates:
		dBuf := mtproto.NewDecodeBuf(r.Query)
		return getRpcMethod(dBuf.Object())
	case *mtproto.TLInvokeWithMessagesRange:
		dBuf := mtproto.NewDecodeBuf(r.Query)
		return getRpcMethod(dBuf.Object())
	case *mtproto.TLInvokeWithTakeout:
		dBuf := mtproto.NewDecodeBuf(r.Query)
		return getRpcMethod(dBuf.Object())
	case *mtproto.TLInvokeWithBusinessConnection:
		dBuf := mtproto.NewDecodeBuf(r.Query)
		return getRpcMethod(dBuf.Object())
	case *mtproto.TLInitConnection:
		dBuf := mtproto.NewDecodeBuf(r.Query)
		return getRpcMethod(dBuf.Object())
	case *mtproto.TLGzipPacked:
		return r.Obj
	default:
		return r
	}
}

func tryGetPermAuthKeyId(b []byte) int64 {
	var (
		err  error
		msg  = &mtproto.TLMessage2{}
		msgs []*mtproto.TLMessage2
	)

	err = msg.Decode(mtproto.NewDecodeBuf(b))
	if err != nil {
		return 0
	}

	if msgContainer, ok := msg.Object.(*mtproto.TLMsgContainer); ok {
		msgs = msgContainer.Messages
	} else {
		msgs = append(msgs, msg)
	}

	//for i := 0; i < len(msgs); i++ {
	//	if packed, ok := msgs[i].Object.(*mtproto.TLGzipPacked); ok {
	//		msgs[i] = &mtproto.TLMessage2{
	//			MsgId:  msgs[i].MsgId,
	//			Seqno:  msgs[i].Seqno,
	//			Bytes:  int32(len(packed.PackedData)),
	//			Object: packed.Obj,
	//		}
	//	}
	//}

	for _, m2 := range msgs {
		r := getRpcMethod(m2.Object)
		if r != nil {
			if request, ok := r.(*mtproto.TLAuthBindTempAuthKey); ok {
				return request.GetPermAuthKeyId()
			}
		}
	}

	return 0
}
