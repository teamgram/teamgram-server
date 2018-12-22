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

package zrpc

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/nebula-chat/chatengine/mtproto/rpc/brpc"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
)

type ZRpcController struct {
	*brpc.RpcMeta
	Attachment []byte
}

func NewController() *ZRpcController {
	return &ZRpcController{
		RpcMeta: new(brpc.RpcMeta),
	}
}

func (c *ZRpcController) Clone() *ZRpcController {
	md := proto.Clone(c.RpcMeta)
	return &ZRpcController{
		RpcMeta:    md.(*brpc.RpcMeta),
		Attachment: c.Attachment,
	}
}

func (c *ZRpcController) String() string {
	return c.RpcMeta.String()
}

func (c *ZRpcController) SetServiceName(v string) {
	if c.RpcMeta == nil {
		c.RpcMeta = new(brpc.RpcMeta)
	}
	if c.GetRequest() == nil {
		c.Request = new(brpc.RpcRequestMeta)
	}
	c.Request.ServiceName = &v
}

func (c *ZRpcController) SetMethodName(v string) {
	if c.RpcMeta == nil {
		c.RpcMeta = new(brpc.RpcMeta)
	}
	if c.GetRequest() == nil {
		c.Request = new(brpc.RpcRequestMeta)
	}
	c.Request.MethodName = &v
}

func (c *ZRpcController) SetLogId(v int64) {
	if c.RpcMeta == nil {
		c.RpcMeta = new(brpc.RpcMeta)
	}
	if c.GetRequest() == nil {
		c.Request = new(brpc.RpcRequestMeta)
	}
	c.Request.LogId = &v
}

func (c *ZRpcController) SetTraceId(v int64) {
	if c.RpcMeta == nil {
		c.RpcMeta = new(brpc.RpcMeta)
	}
	if c.GetRequest() == nil {
		c.Request = new(brpc.RpcRequestMeta)
	}
	c.Request.TraceId = &v
}

func (c *ZRpcController) SetSpanId(v int64) {
	if c.RpcMeta == nil {
		c.RpcMeta = new(brpc.RpcMeta)
	}
	if c.GetRequest() == nil {
		c.Request = new(brpc.RpcRequestMeta)
	}
	c.Request.SpanId = &v
}

func (c *ZRpcController) SetParentSpanId(v int64) {
	if c.RpcMeta == nil {
		c.RpcMeta = new(brpc.RpcMeta)
	}
	if c.GetRequest() == nil {
		c.Request = new(brpc.RpcRequestMeta)
	}
	c.Request.ParentSpanId = &v
}

func (c *ZRpcController) SetErrorCode(v int32) {
	if c.RpcMeta == nil {
		c.RpcMeta = new(brpc.RpcMeta)
	}
	if c.GetResponse() == nil {
		c.Response = new(brpc.RpcResponseMeta)
	}
	c.Response.ErrorCode = &v
}

func (c *ZRpcController) SetErrorText(v string) {
	if c.RpcMeta == nil {
		c.RpcMeta = new(brpc.RpcMeta)
	}
	if c.GetResponse() == nil {
		c.Response = new(brpc.RpcResponseMeta)
	}
	c.Response.ErrorText = &v
}

func (c *ZRpcController) SetCompressType(v int32) {
	if c.RpcMeta == nil {
		c.RpcMeta = new(brpc.RpcMeta)
	}
	c.CompressType = &v
}

func (c *ZRpcController) SetCorrelationId(v int64) {
	if c.RpcMeta == nil {
		c.RpcMeta = new(brpc.RpcMeta)
	}
	c.CorrelationId = &v
}

func (c *ZRpcController) SetAuthKeyId(v int64) {
	if c.RpcMeta == nil {
		c.RpcMeta = new(brpc.RpcMeta)
	}
	if c.GetMtprotoMeta() == nil {
		c.MtprotoMeta = new(brpc.MTProtoMeta)
	}
	c.MtprotoMeta.AuthKeyId = &v
}

func (c *ZRpcController) SetSessionId(v int64) {
	if c.RpcMeta == nil {
		c.RpcMeta = new(brpc.RpcMeta)
	}
	if c.GetMtprotoMeta() == nil {
		c.MtprotoMeta = new(brpc.MTProtoMeta)
	}
	c.MtprotoMeta.SessionId = &v
}

func (c *ZRpcController) SetMessageId(v int64) {
	if c.RpcMeta == nil {
		c.RpcMeta = new(brpc.RpcMeta)
	}
	if c.GetMtprotoMeta() == nil {
		c.MtprotoMeta = new(brpc.MTProtoMeta)
	}
	c.MtprotoMeta.MessageId = &v
}

func (c *ZRpcController) SetLayer(v int32) {
	if c.RpcMeta == nil {
		c.RpcMeta = new(brpc.RpcMeta)
	}
	if c.GetMtprotoMeta() == nil {
		c.MtprotoMeta = new(brpc.MTProtoMeta)
	}
	c.MtprotoMeta.Layer = &v
}

func (c *ZRpcController) SetUserId(v int32) {
	if c.RpcMeta == nil {
		c.RpcMeta = new(brpc.RpcMeta)
	}
	if c.GetMtprotoMeta() == nil {
		c.MtprotoMeta = new(brpc.MTProtoMeta)
	}
	c.MtprotoMeta.UserId = &v
}

func (c *ZRpcController) SetAccessHash(v int64) {
	if c.RpcMeta == nil {
		c.RpcMeta = new(brpc.RpcMeta)
	}
	if c.GetMtprotoMeta() == nil {
		c.MtprotoMeta = new(brpc.MTProtoMeta)
	}
	c.MtprotoMeta.AccessHash = &v
}

func (c *ZRpcController) SetServerId(v int32) {
	if c.RpcMeta == nil {
		c.RpcMeta = new(brpc.RpcMeta)
	}
	if c.GetMtprotoMeta() == nil {
		c.MtprotoMeta = new(brpc.MTProtoMeta)
	}
	c.MtprotoMeta.ServerId = &v
}

func (c *ZRpcController) SetClientConnId(v int64) {
	if c.RpcMeta == nil {
		c.RpcMeta = new(brpc.RpcMeta)
	}
	if c.GetMtprotoMeta() == nil {
		c.MtprotoMeta = new(brpc.MTProtoMeta)
	}
	c.MtprotoMeta.ClientConnId = &v
}

func (c *ZRpcController) SetClientAddr(v string) {
	if c.RpcMeta == nil {
		c.RpcMeta = new(brpc.RpcMeta)
	}
	if c.GetMtprotoMeta() == nil {
		c.MtprotoMeta = new(brpc.MTProtoMeta)
	}
	c.MtprotoMeta.ClientAddr = &v
}

func (c *ZRpcController) SetFrom(v string) {
	if c.RpcMeta == nil {
		c.RpcMeta = new(brpc.RpcMeta)
	}
	if c.GetMtprotoMeta() == nil {
		c.MtprotoMeta = new(brpc.MTProtoMeta)
	}
	c.MtprotoMeta.From = &v
}

func (c *ZRpcController) SetReceiveTime(v int64) {
	if c.RpcMeta == nil {
		c.RpcMeta = new(brpc.RpcMeta)
	}
	if c.GetMtprotoMeta() == nil {
		c.MtprotoMeta = new(brpc.MTProtoMeta)
	}
	c.MtprotoMeta.ReceiveTime = &v
}

func (c *ZRpcController) setAttachmentSize(v int32) {
	if c.RpcMeta == nil {
		c.RpcMeta = new(brpc.RpcMeta)
	}
	c.AttachmentSize = &v
}

func (c *ZRpcController) GetAttachment() []byte {
	return c.Attachment
}

func (c *ZRpcController) MoveAttachment() []byte {
	var move = c.Attachment
	c.setAttachmentSize(0)
	c.Attachment = nil
	return move
}

func (c *ZRpcController) SetAttachment(b []byte) {
	c.Attachment = b
	c.setAttachmentSize(int32(len(b)))
}

func MakeBaiduRpcMessage(cntl *ZRpcController, zmsg proto.Message) *brpc.BaiduRpcMessage {
	// cntl.SetMethodName(proto.MessageName(zmsg))
	payload, _ := proto.Marshal(zmsg)
	bmsg := &brpc.BaiduRpcMessage{
		Meta:       cntl.RpcMeta,
		Payload:    payload,
		Attachment: cntl.Attachment,
	}
	return bmsg
}

func SplitBaiduRpcMessage(bmsg *brpc.BaiduRpcMessage) (*ZRpcController, proto.Message, error) {
	methodName := bmsg.Meta.GetRequest().GetMethodName()
	if methodName == "" {
		err := fmt.Errorf("rpc method name is emtpty")
		return nil, nil, err
	}

	zmsg, err := grpc_util.NewMessageByName(methodName)
	if err != nil {
		// glog.Error(err)
		return nil, nil, err
	}

	err = proto.Unmarshal(bmsg.Payload, zmsg)
	if err != nil {
		// glog.Error(err)
		return nil, nil, err
	}

	cntl := &ZRpcController{
		RpcMeta:    bmsg.Meta,
		Attachment: bmsg.Attachment,
	}

	return cntl, zmsg, nil
}
