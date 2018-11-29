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
	"fmt"
	"github.com/nebula-chat/chatengine/pkg/util"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/net2"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/mtproto/rpc"
	"github.com/nebula-chat/chatengine/service/idgen/client"
	"sync"
	"github.com/gogo/protobuf/proto"
	"time"
)

func init() {
	//
	proto.RegisterType((*mtproto.TLSessionClientCreated)(nil), "mtproto.TLSessionClientCreated")
	proto.RegisterType((*mtproto.TLSessionClientClosed)(nil), "mtproto.TLSessionClientClosed")
	proto.RegisterType((*mtproto.SessionClientEvent)(nil), "mtproto.SessionClientEvent")

	proto.RegisterType((*mtproto.TLSessionMessageData)(nil), "mtproto.TLSessionMessageData")
	proto.RegisterType((*mtproto.TLHandshakeData)(nil), "mtproto.TLHandshakeData")
	proto.RegisterType((*mtproto.RawMessageData)(nil), "mtproto.RawMessageData")
}

type connContext struct {
	// TODO(@benqi): lock
	sync.Mutex
	state          int // 是否握手阶段
	handshakeState *mtproto.TLHandshakeData
	sessionAddr    string
	authKeyId      int64
}

func (ctx *connContext) getState() int {
	ctx.Lock()
	defer ctx.Unlock()
	return ctx.state
}

func (ctx *connContext) setState(state int) {
	ctx.Lock()
	defer ctx.Unlock()
	if ctx.state != state {
		ctx.state = state
	}
}

func (ctx *connContext) encryptedMessageAble() bool {
	ctx.Lock()
	defer ctx.Unlock()
	return ctx.state == mtproto.STATE_CONNECTED2 ||
		ctx.state == mtproto.STATE_AUTH_KEY ||
		(ctx.state == mtproto.STATE_HANDSHAKE &&
			(ctx.handshakeState.GetState() == mtproto.STATE_pq_res ||
				(ctx.handshakeState.GetState() == mtproto.STATE_dh_gen_res &&
					ctx.handshakeState.GetResState() == mtproto.RES_STATE_OK)))

}

type frontendServer struct {
	idgen      idgen.UUIDGen
	server80   *mtproto.MTProtoServer
	server443  *mtproto.MTProtoServer
	server5222 *mtproto.MTProtoServer
	client     *zrpc.ZRpcClient
}

func NewFrontendServer() *frontendServer {
	return &frontendServer{}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// AppInstance interface
func (s *frontendServer) Initialize() error {
	err := InitializeConfig()
	if err != nil {
		glog.Fatal(err)
		return err
	}
	glog.Info("load conf: ", Conf)

	// idgen
	s.idgen, _ = idgen.NewUUIDGen("snowflake", util.Int32ToString(Conf.ServerId))

	// mtproto_server
	s.server80 = mtproto.NewMTProtoServer(Conf.Server80, s)
	s.server443 = mtproto.NewMTProtoServer(Conf.Server443, s)
	s.server5222 = mtproto.NewMTProtoServer(Conf.Server5222, s)

	// client
	s.client = zrpc.NewZRpcClient("brpc", Conf.Clients, s)
	return nil
}

func (s *frontendServer) RunLoop() {
	s.server80.Serve()
	s.server443.Serve()
	s.server5222.Serve()
	s.client.Serve()
}

func (s *frontendServer) Destroy() {
	s.server80.Stop()
	s.server443.Stop()
	s.server5222.Stop()
	s.client.Stop()
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// MTProtoServerCallback
func (s *frontendServer) OnServerNewConnection(conn *net2.TcpConnection) {
	conn.Context = &connContext{
		state: mtproto.STATE_CONNECTED2,
		handshakeState: &mtproto.TLHandshakeData{Data2: &mtproto.RawMessageData_Data{
			State:    mtproto.STATE_CONNECTED2,
			ResState: mtproto.RES_STATE_NONE,
			Ctx:      mtproto.NewTLHandshakeContext().To_HandshakeContext(),
		}},
	}

	glog.Infof("onServerNewConnection - {peer: %s, ctx: {%v}}", conn, conn.Context)
}

func (s *frontendServer) OnServerMessageDataArrived(conn *net2.TcpConnection, msg *mtproto.MTPRawMessage) error {
	//md := s.newMetadata(conn)
	glog.Infof("onServerMessageDataArrived - receive data: {peer: %s, md: %s, msg: %s}", conn, msg)

	ctx, _ := conn.Context.(*connContext)

	var err error
	if msg.AuthKeyId() == 0 {
		// check ctx.AuthKeyId
		if ctx.getState() == mtproto.STATE_AUTH_KEY {
			err = fmt.Errorf("invalid state STATE_AUTH_KEY: %d", ctx.getState())
			glog.Errorf("process msg error: {%v} - {peer: %s, md: %s, msg: %s}", err, conn, msg)
			conn.Close()
		} else {
			err = s.onServerUnencryptedRawMessage(ctx, conn, msg)
		}
	} else {
		if !ctx.encryptedMessageAble() {
			err = fmt.Errorf("invalid state: {state: %d, handshakeState: {%v}}", ctx.state, ctx.handshakeState)
			glog.Errorf("process msg error: {%v} - {peer: %s, md: %s, msg: %s}", err, conn, msg)
			conn.Close()
		} else {
			err = s.onServerEncryptedRawMessage(ctx, conn, msg)
			if ctx.authKeyId == 0 {
				ctx.authKeyId = msg.AuthKeyId()
			}
		}
	}

	return err
}

func (s *frontendServer) OnServerConnectionClosed(conn *net2.TcpConnection) {
	glog.Infof("onServerConnectionClosed - {peer: %s}", conn)
	// s.sendClientClosed(conn)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// ZProtoClientCallBack
func (s *frontendServer) OnNewClient(client *net2.TcpClient) {
	glog.Infof("onNewClient - peer(%s)", client.GetConnection())
}


func (s *frontendServer) OnClientMessageArrived(client *net2.TcpClient, cntl *zrpc.ZRpcController, msg proto.Message) error {
//func (s *frontendServer) OnClientMessageArrived(client *net2.TcpClient, md *zproto.ZProtoMetadata, sessionId, messageId uint64, seqNo uint32, msg zproto.MessageBase) error {
	var err error

	switch msg.(type) {
	case *mtproto.RawMessageData:
		rawMsg, _ := msg.(*mtproto.RawMessageData)
		switch rawMsg.GetConstructor() {
		case mtproto.TLConstructor_CRC32_handshakeData:
			err = s.onClientHandshakeMessage(client, cntl, rawMsg.To_HandshakeData())
		case mtproto.TLConstructor_CRC32_sessionMessageData:
			err = s.onClientSessionData(client, cntl, rawMsg.To_SessionMessageData())
		default:
			err = fmt.Errorf("invalid error, md: %s, msg: %v", cntl, msg)
		}
	default:
		err = fmt.Errorf("invalid msg, md: %s, msg: %v", cntl, msg)
	}

	if err != nil {
		glog.Error(err)
	}
	return err
}

func (s *frontendServer) OnClientClosed(client *net2.TcpClient) {
	glog.Infof("onClientClosed - peer(%s)", client.GetConnection())
}

func (s *frontendServer) OnClientTimer(client *net2.TcpClient) {
	// Impl timer logic
	glog.Info("onClientTimer")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
func (s *frontendServer) onClientHandshakeMessage(client *net2.TcpClient, cntl *zrpc.ZRpcController, handshake *mtproto.TLHandshakeData) error {
	glog.Infof("onClientHandshakeMessage - handshake: peer(%s), state: {%v}",
		client.GetConnection(),
		handshake)

	///////////////////////////////////////////////////////////////////
	conn := s.getConnBySessionID(uint64(handshake.GetClientConnId()))
	if conn == nil {
		glog.Warning("conn closed, handshake: ", handshake)
		return nil
	}

	if handshake.GetResState() == mtproto.RES_STATE_ERROR {
		// TODO(@benqi): Close.
		glog.Warning(" handshake.State.ResState error, handshake: ", handshake)
		// conn.Close()
		return nil
	} else {
		ctx := conn.Context.(*connContext)
		ctx.Lock()
		ctx.handshakeState = handshake
		ctx.Unlock()

		glog.Infof("onClientHandshakeMessage - sendToClient to: {peer: %s, handshake: %s}",
			conn,
			handshake)

		return conn.Send(&mtproto.MTPRawMessage{Payload: cntl.Attachment})
	}
}

func (s *frontendServer) onClientSessionData(client *net2.TcpClient, cntl *zrpc.ZRpcController, sessData *mtproto.TLSessionMessageData) error {
	///////////////////////////////////////////////////////////////////
	conn := s.getConnBySessionID(uint64(sessData.GetClientConnId()))
	// s.server443.GetConnection(zmsg.SessionId)
	if conn == nil {
		glog.Warning("conn closed, connID = ", sessData.GetClientConnId())
		return nil
	}

	glog.Infof("onClientSessionData - sendToClient to: {peer: %s, sessData: %s}",
		conn,
		sessData)
	return conn.Send(&mtproto.MTPRawMessage{Payload: cntl.Attachment})
}

func (s *frontendServer) genSessionId(conn *net2.TcpConnection) uint64 {
	var sid = conn.GetConnID()
	if conn.Name() == "frontend443" {
		// sid = sid | 0 << 60
	} else if conn.Name() == "frontend80" {
		sid = sid | 1<<60
	} else if conn.Name() == "frontend5222" {
		sid = sid | 2<<60
	}

	return sid
}

func (s *frontendServer) getConnBySessionID(id uint64) *net2.TcpConnection {
	//
	var server *mtproto.MTProtoServer
	sid := id >> 60
	if sid == 0 {
		server = s.server443
	} else if sid == 1 {
		server = s.server80
	} else if sid == 2 {
		server = s.server5222
	} else {
		return nil
	}

	id = id & 0xffffffffffffff
	return server.GetConnection(id)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
func (s *frontendServer) onServerUnencryptedRawMessage(ctx *connContext, conn *net2.TcpConnection, mmsg *mtproto.MTPRawMessage) error {
	glog.Infof("onServerUnencryptedRawMessage - receive data: {peer: %s, ctx: %s, msg: %s}", conn, ctx, mmsg)

	zmsg := mtproto.NewTLHandshakeData()
	zmsg.SetClientConnId(int64(conn.GetConnID()))

	ctx.Lock()
	if ctx.state == mtproto.STATE_CONNECTED2 {
		ctx.state = mtproto.STATE_HANDSHAKE
	}
	if ctx.handshakeState.GetState() == mtproto.STATE_CONNECTED2 {
		ctx.handshakeState.SetState(mtproto.STATE_pq)
	}
	zmsg.SetState(int32(ctx.handshakeState.GetState()))
	zmsg.SetCtx(ctx.handshakeState.GetCtx())
	ctx.Unlock()

	cntl := zrpc.NewController()

	cntl.SetAttachment(mmsg.Payload)
	cntl.SetServiceName("handshake")
	cntl.SetMethodName(proto.MessageName(zmsg))

	var id int64
	id, _ = s.idgen.GetUUID()
	cntl.SetLogId(id)
	id, _ = s.idgen.GetUUID()
	cntl.SetCorrelationId(id)
	id, _ = s.idgen.GetUUID()
	cntl.SetTraceId(id)
	id, _ = s.idgen.GetUUID()
	cntl.SetSpanId(id)

	cntl.SetAuthKeyId(mmsg.AuthKeyId())
	cntl.SetServerId(Conf.ServerId)
	cntl.SetClientConnId(int64(conn.GetConnID()))
	cntl.SetClientAddr(conn.RemoteAddr().String())
	cntl.SetFrom("frontend")
	cntl.SetReceiveTime(time.Now().UnixNano() / 1e6)

	glog.Infof("sendMessage - handshake: {peer: %s, md: %s, msg: %v}", conn, cntl, zmsg)
	return s.client.SendMessage("handshake", cntl, zmsg)
}

func (s *frontendServer) onServerEncryptedRawMessage(ctx *connContext, conn *net2.TcpConnection, mmsg *mtproto.MTPRawMessage) error {
	glog.Infof("onServerEncryptedRawMessage - receive data: {peer: %s, ctx: %s, msg: %s}", conn, ctx, mmsg)

	zmsg := mtproto.NewTLSessionMessageData()
	zmsg.SetConnType(int32(mmsg.ConnType()))
	zmsg.SetClientConnId(int64(conn.GetConnID()))
	zmsg.SetAuthKeyId(mmsg.AuthKeyId())
	zmsg.SetQuickAck(mmsg.QuickAckId())

	cntl := zrpc.NewController()

	cntl.SetAttachment(mmsg.Payload)
	cntl.SetServiceName("session")
	cntl.SetMethodName(proto.MessageName(zmsg))

	var id int64
	id, _ = s.idgen.GetUUID()
	cntl.SetLogId(id)
	id, _ = s.idgen.GetUUID()
	cntl.SetCorrelationId(id)
	id, _ = s.idgen.GetUUID()
	cntl.SetTraceId(id)
	id, _ = s.idgen.GetUUID()
	cntl.SetSpanId(id)

	cntl.SetAuthKeyId(mmsg.AuthKeyId())
	cntl.SetServerId(Conf.ServerId)
	cntl.SetClientConnId(int64(conn.GetConnID()))
	cntl.SetClientAddr(conn.RemoteAddr().String())
	cntl.SetFrom("frontend")
	cntl.SetReceiveTime(time.Now().UnixNano() / 1e6)

	return s.client.SendKetamaMessage("session", util.Int64ToString(mmsg.AuthKeyId()), cntl, zmsg, func(addr string) {
		// s.checkAndSendClientNew(ctx, conn, addr, mmsg.AuthKeyId(), md)
	})
}

func (s *frontendServer) checkAndSendClientNew(ctx *connContext, conn *net2.TcpConnection, kaddr string, authKeyId int64) error {
	var err error
	if ctx.sessionAddr == "" {
		//clientNew := &zproto.ZProtoSessionClientNew{
		//	// MTPMessage: mmsg,
		//}
		//err = s.client.SendMessageToAddress("session", kaddr, s.newMetadata(conn), clientNew)
		//if err == nil {
		//	ctx.sessionAddr = kaddr
		//	ctx.authKeyId = authKeyId
		//} else {
		//	glog.Error(err)
		//}
	} else {
		// TODO(@benqi): check ctx.sessionAddr == kaddr
	}

	return err
}

func (s *frontendServer) sendClientClosed(conn *net2.TcpConnection) {
	if conn.Context == nil {
		return
	}

	ctx, _ := conn.Context.(*connContext)
	if ctx.sessionAddr == "" || ctx.authKeyId == 0 {
		return
	}

	//s.client.SendKetamaMessage("session", utils.Int64ToString(ctx.authKeyId), s.newMetadata(conn), &zproto.ZProtoSessionClientClosed{}, nil)
}
