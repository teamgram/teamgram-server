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

package server

import (
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/golang/glog"
	"encoding/hex"
)

type messagePackUtil struct {
	messages     []*mtproto.TLMessage2
	errMsgIDList []int64
	// layer        int32
}

// TODO(@benqi): handle unpack error, send bad_msg_notification??
func (pack *messagePackUtil) unpackServiceMessage(msgId int64, seqNo int32, object mtproto.TLObject) {
	switch object.(type) {
	case *mtproto.TLMsgContainer:
		msgContainer, _ := object.(*mtproto.TLMsgContainer)
		for _, m2 := range msgContainer.Messages {
			pack.unpackServiceMessage(m2.MsgId, m2.Seqno, m2.Object)
		}
	case *mtproto.TLGzipPacked:
		gzipPacked, _ := object.(*mtproto.TLGzipPacked)
		glog.Info("processGzipPacked - request data: ", gzipPacked)
		dbuf := mtproto.NewDecodeBuf(gzipPacked.PackedData)
		o := dbuf.Object()
		if o == nil {
			glog.Errorf("decode query error: %s", hex.EncodeToString(gzipPacked.PackedData))
			pack.errMsgIDList = append(pack.errMsgIDList, msgId)
			return
		}
		pack.unpackServiceMessage(msgId, seqNo, o)
	case *mtproto.TLMsgCopy:
		// not use in client
		glog.Error("android client not use msg_copy: ", object)
	case *mtproto.TLInvokeAfterMsg:
		invokeAfterMsg := object.(*mtproto.TLInvokeAfterMsg)
		invokeAfterMsgExt := NewInvokeAfterMsgExt(invokeAfterMsg)
		if invokeAfterMsgExt.Query == nil {
			glog.Errorf("decode query error: %s", hex.EncodeToString(invokeAfterMsg.Query))
			pack.errMsgIDList = append(pack.errMsgIDList, msgId)
			return
		}

		pack.messages = append(pack.messages, &mtproto.TLMessage2{MsgId: msgId, Seqno: seqNo, Object: invokeAfterMsgExt})
	case *mtproto.TLInvokeAfterMsgs:
		invokeAfterMsgs := object.(*mtproto.TLInvokeAfterMsgs)
		invokeAfterMsgsExt := NewInvokeAfterMsgsExt(invokeAfterMsgs)
		if invokeAfterMsgsExt.Query == nil {
			glog.Errorf("decode query error: %s", hex.EncodeToString(invokeAfterMsgs.Query))
			pack.errMsgIDList = append(pack.errMsgIDList, msgId)
			return
		}

		pack.messages = append(pack.messages, &mtproto.TLMessage2{MsgId: msgId, Seqno: seqNo, Object: invokeAfterMsgsExt})

	case *mtproto.TLInvokeWithLayer:
		invokeWithLayer := object.(*mtproto.TLInvokeWithLayer)
		if invokeWithLayer.GetQuery() == nil {
			glog.Errorf("invokeWithLayer Query is nil, query: {%v}", invokeWithLayer)
			pack.errMsgIDList = append(pack.errMsgIDList, msgId)
			return
		} else {
			dbuf := mtproto.NewDecodeBuf(invokeWithLayer.Query)
			var initConnectionExt *TLInitConnectionExt

			classID := dbuf.Int()
			switch classID {
			case int32(mtproto.TLConstructor_CRC32_initConnection):
				initConnection := &mtproto.TLInitConnection{}
				err := initConnection.Decode(dbuf)
				if err != nil {
					glog.Error("decode initConnection error: ", err)
					pack.errMsgIDList = append(pack.errMsgIDList, msgId)
					return
				}
				initConnectionExt = NewInitConnectionExt(invokeWithLayer.Layer, initConnection)
			case int32(mtproto.TLConstructor_CRC32_initConnectionLayer68):
				initConnection := &mtproto.TLInitConnectionLayer68{}
				err := initConnection.Decode(dbuf)
				if err != nil {
					glog.Error("decode initConnectionLayer68 error: ", err)
					pack.errMsgIDList = append(pack.errMsgIDList, msgId)
					return
				}
				initConnectionExt = NewInitConnectionExtByLayer68(invokeWithLayer.Layer, initConnection)
			default:
				glog.Errorf("not initConnection classID: %d", classID)
				pack.errMsgIDList = append(pack.errMsgIDList, msgId)
			}
			pack.messages = append(pack.messages, &mtproto.TLMessage2{MsgId: msgId, Seqno: seqNo, Object: initConnectionExt})
		}

	case *mtproto.TLInvokeWithoutUpdates:
		// TODO(@benqi): macOS client used.
		// glog.Error("android client not use invokeWithoutUpdates: ", object)
		invokeWithoutUpdates := object.(*mtproto.TLInvokeWithoutUpdates)
		invokeWithoutUpdatesExt := NewInvokeWithoutUpdatesExt(invokeWithoutUpdates)
		if invokeWithoutUpdatesExt.Query == nil {
			glog.Errorf("decode query error: %s", hex.EncodeToString(invokeWithoutUpdates.Query))
			pack.errMsgIDList = append(pack.errMsgIDList, msgId)
			return
		}
		pack.messages = append(pack.messages, &mtproto.TLMessage2{MsgId: msgId, Seqno: seqNo, Object: invokeWithoutUpdatesExt})

	case *mtproto.TLInvokeWithMessagesRange:
		invokeWithMessagesRange := object.(*mtproto.TLInvokeWithMessagesRange)
		invokeWithMessagesRangeExt := NewTLInvokeWithMessagesRangeExt(invokeWithMessagesRange)
		if invokeWithMessagesRangeExt.Query == nil {
			glog.Errorf("decode query error: %s", hex.EncodeToString(invokeWithMessagesRange.Query))
			pack.errMsgIDList = append(pack.errMsgIDList, msgId)
			return
		}
		pack.messages = append(pack.messages, &mtproto.TLMessage2{MsgId: msgId, Seqno: seqNo, Object: invokeWithMessagesRangeExt})

	case *mtproto.TLInvokeWithTakeout:
		invokeWithTakeout := object.(*mtproto.TLInvokeWithTakeout)
		invokeWithTakeoutExt := NewTLInvokeWithTakeoutExt(invokeWithTakeout)
		if invokeWithTakeoutExt.Query == nil {
			glog.Errorf("decode query error: %s", hex.EncodeToString(invokeWithTakeout.Query))
			pack.errMsgIDList = append(pack.errMsgIDList, msgId)
			return
		}
		pack.messages = append(pack.messages, &mtproto.TLMessage2{MsgId: msgId, Seqno: seqNo, Object: invokeWithTakeoutExt})
	default:
		// glog.Info("processOthers - request data: ", object)
		pack.messages = append(pack.messages, &mtproto.TLMessage2{MsgId: msgId, Seqno: seqNo, Object: object})
	}
}
