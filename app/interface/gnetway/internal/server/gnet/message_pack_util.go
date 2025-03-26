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
	"encoding/binary"
	"time"

	"github.com/teamgram/marmota/pkg/sync2"
	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/mt"
	"github.com/teamgram/proto/v2/tg"

	"github.com/zeromicro/go-zero/core/logx"
)

var msgIdSeq = sync2.NewAtomicInt64(0)

func nextMessageId(isRpc bool) int64 {
	unixNano := time.Now().UnixNano()
	ts := unixNano / 1e9
	ms := (unixNano % 1e9) / 1e6
	sid := msgIdSeq.Add(1) & 0x1ffff
	msgIdSeq.CompareAndSwap(0x1ffff, 0)
	last := 1
	if !isRpc {
		last = 3
	}
	msgId := ts<<32 | int64(ms)<<21 | sid<<3 | int64(last)
	return msgId
}

func parseFromIncomingMessage(b []byte) (msgId int64, obj iface.TLObject, err error) {
	dBuf := bin.NewDecoder(b)

	msgId, _ = dBuf.Long()
	_, _ = dBuf.Int()
	obj, err = iface.DecodeObject(dBuf)

	return
}

func serializeToBuffer(x *bin.Encoder, msgId int64, obj iface.TLObject) error {
	//obj.Encode(x, 0)
	// x := mtproto.NewEncodeBuf(8 + 4 + len(oBuf))
	x.PutLong(0)
	x.PutLong(msgId)
	offset := x.Len()
	x.PutInt(0)
	err := obj.Encode(x, 0)
	if err != nil {
		return err
	}

	b := x.Bytes()

	binary.LittleEndian.PutUint32(b[offset:], uint32(x.Len()-offset-4))

	return nil
}

func serializeToBuffer2(salt, sessionId int64, msg2 *mt.TLMessage2) []byte {
	x := bin.NewEncoder()
	defer x.End()

	x.PutLong(salt)
	x.PutLong(sessionId)
	_ = msg2.Encode(x, 0)

	return x.Bytes()
}

const (
// kMsgContainerBufLen = 8
)

var (
//	kMsgContainerBuf = func() []byte {
//		x := bin.NewEncoder()
//      defer x.End()

//		x.Int(int32(mtproto.CRC32_msg_container))
//		x.Int(0)
//		return x.GetBuf()
//	}()
)

//func serializeToBuffer2(salt, sessionId int64, seqNo int32) []byte {
//	x := mtproto.NewEncodeBuf(512)
//
//	x.Long(salt)
//	x.Long(sessionId)
//	x.Long(mtproto.GenerateMessageId())
//	x.Int(seqNo)
//	x.Int(kMsgContainerBufLen)
//	x.Bytes(kMsgContainerBuf)
//
//	return x.GetBuf()
//}

func getRpcMethod(in iface.TLObject) iface.TLObject {
	if in == nil {
		return nil
	}
	logx.Debugf("rpc: %s", in)

	switch r := in.(type) {
	case *mt.TLDestroyAuthKey: // 所有连接都有可能
		return nil
	case *mt.TLRpcDropAnswer: // 所有连接都有可能
		return nil
	case *mt.TLGetFutureSalts: // GENERIC
		return nil
	case *mt.TLPing: // android未用
		return nil
	case *mt.TLPingDelayDisconnect: // PUSH和GENERIC
		return nil
	case *mt.TLDestroySession: // GENERIC
		return nil
	case *mt.TLMsgsStateReq: // android未用
		return nil
	case *mt.TLMsgsStateInfo: // android未用
		return nil
	case *mt.TLMsgsAllInfo: // android未用
		return nil
	case *mt.TLMsgResendReq: // 都有可能
		return nil
	case *mt.TLMsgDetailedInfo: // 都有可能
		return nil
	case *mt.TLMsgNewDetailedInfo: // 都有可能
		return nil
	case *tg.TLInvokeWithLayer:
		dBuf := bin.NewDecoder(r.Query)
		o, _ := iface.DecodeObject(dBuf)
		return getRpcMethod(o)
	case *tg.TLInvokeAfterMsg:
		dBuf := bin.NewDecoder(r.Query)
		o, _ := iface.DecodeObject(dBuf)
		return getRpcMethod(o)
	case *tg.TLInvokeAfterMsgs:
		dBuf := bin.NewDecoder(r.Query)
		o, _ := iface.DecodeObject(dBuf)
		return getRpcMethod(o)
	case *tg.TLInvokeWithoutUpdates:
		dBuf := bin.NewDecoder(r.Query)
		o, _ := iface.DecodeObject(dBuf)
		return getRpcMethod(o)
	case *tg.TLInvokeWithMessagesRange:
		dBuf := bin.NewDecoder(r.Query)
		o, _ := iface.DecodeObject(dBuf)
		return getRpcMethod(o)
	case *tg.TLInvokeWithTakeout:
		dBuf := bin.NewDecoder(r.Query)
		o, _ := iface.DecodeObject(dBuf)
		return getRpcMethod(o)
	case *tg.TLInvokeWithBusinessConnection:
		dBuf := bin.NewDecoder(r.Query)
		o, _ := iface.DecodeObject(dBuf)
		return getRpcMethod(o)
	case *tg.TLInitConnection:
		dBuf := bin.NewDecoder(r.Query)
		o, _ := iface.DecodeObject(dBuf)
		return getRpcMethod(o)
	case *mt.TLGzipPacked:
		return r.Obj
	default:
		return r
	}
}

func tryGetUnknownTLObject(b []byte) (rList []iface.TLObject) {
	var (
		err  error
		msg  = &mt.TLMessage2{}
		msgs []*mt.TLMessage2
	)

	err = msg.Decode(bin.NewDecoder(b))
	if err != nil {
		return
	}

	if msgContainer, ok := msg.Object.(*mt.TLMsgContainer); ok {
		msgs = msgContainer.Messages
	} else {
		msgs = append(msgs, msg)
	}

	for _, m2 := range msgs {
		r := getRpcMethod(m2.Object)
		if r != nil {
			rList = append(rList, r)
		}
	}

	return
}

func tryGetPermAuthKeyId(b []byte) int64 {
	var (
		err  error
		msg  = &mt.TLMessage2{}
		msgs []*mt.TLMessage2
	)

	err = msg.Decode(bin.NewDecoder(b))
	if err != nil {
		return 0
	}

	if msgContainer, ok := msg.Object.(*mt.TLMsgContainer); ok {
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
			if request, ok := r.(*tg.TLAuthBindTempAuthKey); ok {
				return request.PermAuthKeyId
			}
		}
	}

	return 0
}

func encodeUnencryptedMessage(x *bin.Encoder, msgId int64, obj iface.TLObject) error {
	x.PutInt64(0)
	x.PutInt64(msgId)
	offset := x.Len()
	x.PutInt(0)
	if err := obj.Encode(x, 0); err != nil {
		return err
	}
	b := x.Bytes()
	binary.LittleEndian.PutUint32(b[offset:], uint32(x.Len()-offset-4))
	return nil
}
