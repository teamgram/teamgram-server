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
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/teamgram/marmota/pkg/hack"
	"github.com/teamgram/marmota/pkg/net2"
	"github.com/teamgram/marmota/pkg/timer2"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/interface/gateway/internal/server/codec"
	sessionpb "github.com/teamgram/teamgram-server/app/interface/session/session"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandshakeStateCtx struct {
	State         int32  `json:"state,omitempty"`
	ResState      int32  `json:"res_state,omitempty"`
	Nonce         []byte `json:"nonce,omitempty"`
	ServerNonce   []byte `json:"server_nonce,omitempty"`
	NewNonce      []byte `json:"new_nonce,omitempty"`
	A             []byte `json:"a,omitempty"`
	P             []byte `json:"p,omitempty"`
	handshakeType int
	ExpiresIn     int32 `json:"expires_in,omitempty"`
}

func (m *HandshakeStateCtx) DebugString() string {
	s, _ := json.Marshal(m)
	return hack.String(s)
}

type connContext struct {
	// TODO(@benqi): lock
	sync.Mutex
	state           int // 是否握手阶段
	authKeys        []*authKeyUtil
	sessionId       int64
	isHttp          bool
	canSend         bool
	trd             *timer2.TimerData
	handshakes      []*HandshakeStateCtx
	clientIp        string
	xForwardedForIp string
}

func newConnContext() *connContext {
	return &connContext{
		state:           STATE_CONNECTED2,
		clientIp:        "",
		xForwardedForIp: "*",
	}
}

func (ctx *connContext) getClientIp(xForwarderForIp interface{}) string {
	ctx.Lock()
	defer ctx.Unlock()
	if ctx.xForwardedForIp == "*" {
		ctx.xForwardedForIp = ""
		if xForwarderForIp != nil {
			ctx.xForwardedForIp, _ = xForwarderForIp.(string)
		}
	}

	if ctx.xForwardedForIp != "" {
		return ctx.xForwardedForIp
	}
	return ctx.clientIp
}

func (ctx *connContext) setClientIp(ip string) {
	ctx.Lock()
	defer ctx.Unlock()

	ctx.clientIp = ip
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

func (ctx *connContext) getAuthKey(id int64) *authKeyUtil {
	ctx.Lock()
	defer ctx.Unlock()
	for _, key := range ctx.authKeys {
		if key.AuthKeyId() == id {
			return key
		}
	}

	return nil
}

func (ctx *connContext) putAuthKey(k *authKeyUtil) {
	ctx.Lock()
	defer ctx.Unlock()
	for _, key := range ctx.authKeys {
		if key.Equal(k) {
			return
		}
	}

	ctx.authKeys = append(ctx.authKeys, k)
}

func (ctx *connContext) getAllAuthKeyId() (idList []int64) {
	ctx.Lock()
	defer ctx.Unlock()

	idList = make([]int64, len(ctx.authKeys))
	for i, key := range ctx.authKeys {
		idList[i] = key.AuthKeyId()
	}

	return
}

func (ctx *connContext) getHandshakeStateCtx(nonce []byte) *HandshakeStateCtx {
	ctx.Lock()
	defer ctx.Unlock()

	for _, state := range ctx.handshakes {
		if bytes.Equal(nonce, state.Nonce) {
			return state
		}
	}

	return nil
}

func (ctx *connContext) putHandshakeStateCtx(state *HandshakeStateCtx) {
	ctx.Lock()
	defer ctx.Unlock()

	ctx.handshakes = append(ctx.handshakes, state)
}

func (ctx *connContext) encryptedMessageAble() bool {
	ctx.Lock()
	defer ctx.Unlock()
	//return ctx.state == STATE_CONNECTED2 ||
	//	ctx.state == STATE_AUTH_KEY ||
	//	(ctx.state == STATE_HANDSHAKE &&
	//		(ctx.handshakeCtx.State == STATE_pq_res ||
	//			(ctx.handshakeCtx.State == STATE_dh_gen_res &&
	//				ctx.handshakeCtx.ResState == RES_STATE_OK)))
	return true
}

func (ctx *connContext) DebugString() string {
	s := make([]string, 0, 4)
	s = append(s, fmt.Sprintf(`"state":%d`, ctx.state))
	// s = append(s, fmt.Sprintf(`"handshake_ctx":%s`, ctx.handshakeCtx))
	//if ctx.authKey != nil {
	//	s = append(s, fmt.Sprintf(`"auth_key_id":%d`, ctx.authKey.AuthKeyId()))
	//}
	return "{" + strings.Join(s, ",") + "}"
}

// OnNewConnection
// /////////////////////////////////////////////////////////////////////////////////////////////
func (s *Server) OnNewConnection(conn *net2.TcpConnection) {
	ctx := newConnContext()
	ctx.setClientIp(strings.Split(conn.RemoteAddr().String(), ":")[0])

	logx.Infof("onNewConnection - {peer: %s, ctx: {%s}}", conn, ctx)
	conn.Context = ctx
}

func (s *Server) OnConnectionDataArrived(conn *net2.TcpConnection, msg interface{}) error {
	msg2, ok := msg.(*mtproto.MTPRawMessage)
	if !ok {
		err := fmt.Errorf("recv invalid MTPRawMessage: {peer: %s, msg: %v", conn, msg2)
		logx.Error(err.Error())
		return err
	}

	ctx, _ := conn.Context.(*connContext)

	logx.Infof("onConnectionDataArrived - receive data: {peer: %s, ctx: %s, msg: %s}", conn, ctx, msg2)

	if msg2.ConnType() == codec.TRANSPORT_HTTP {
		ctx.isHttp = true
	}

	var err error
	if msg2.AuthKeyId() == 0 {
		//if ctx.getState() == STATE_AUTH_KEY {
		//	err = fmt.Errorf("invalid state STATE_AUTH_KEY: %d", ctx.getState())
		//	logx.Error("process msg error: {%v} - {peer: %s, ctx: %s, msg: %s}", err, conn, ctx, msg2)
		//	conn.Close()
		//} else {
		//	err = s.onUnencryptedRawMessage(ctx, conn, msg2)
		//}
		err = s.onUnencryptedMessage(ctx, conn, msg2)
	} else {
		//if !ctx.encryptedMessageAble() {
		//	err = fmt.Errorf("invalid state: {state: %d, handshakeState: {%v}}", ctx.state, ctx.handshakeCtx)
		//	logx.Error("process msg error: {%v} - {peer: %s, ctx: %s, msg: %s}", err, conn, ctx, msg2)
		//	conn.Close()
		//} else {
		//	if ctx.state != STATE_AUTH_KEY {
		authKey := ctx.getAuthKey(msg2.AuthKeyId())
		if authKey == nil {
			key := s.GetAuthKey(msg2.AuthKeyId())
			if key == nil {
				sessClient, err2 := s.session.getSessionClient(strconv.FormatInt(msg2.AuthKeyId(), 10))
				if err2 != nil {
					logx.Errorf("getSessionClient error: %v, {authKeyId: %d}", err2, msg2.AuthKeyId())
				} else {
					key, err2 = sessClient.SessionQueryAuthKey(context.Background(), &sessionpb.TLSessionQueryAuthKey{
						AuthKeyId: msg2.AuthKeyId(),
					})
					if err2 != nil {
						logx.Errorf("conn(%s) sessionQueryAuthKey error: %v", conn.String(), err2)
					}
				}
			}
			// key := s.GetAuthKey(msg2.AuthKeyId())
			if key == nil {
				err = fmt.Errorf("invalid auth_key_id: {%d}", msg2.AuthKeyId())
				logx.Error("invalid auth_key_id: {%v} - {peer: %s, ctx: %s, msg: %s}", err, conn, ctx, msg2)
				var code = int32(-404)
				cData := make([]byte, 4)
				binary.LittleEndian.PutUint32(cData, uint32(code))
				conn.Send(&mtproto.MTPRawMessage{Payload: cData})
				// conn.Close()
				return err
			}
			authKey = newAuthKeyUtil(key)
			s.PutAuthKey(key)
			ctx.putAuthKey(authKey)
		}

		err = s.onEncryptedMessage(ctx, conn, authKey, msg2)
	}

	return err
}

func (s *Server) OnConnectionClosed(conn *net2.TcpConnection) {
	ctx, _ := conn.Context.(*connContext)
	logx.Info("onServerConnectionClosed - {peer:%s, ctx:%s}", conn, ctx)

	if ctx.trd != nil {
		s.timer.Del(ctx.trd)
		ctx.trd = nil
	}

	sessId, connId := ctx.sessionId, conn.GetConnID()
	for _, id := range ctx.getAllAuthKeyId() {
		bDeleted := s.authSessionMgr.RemoveSession(id, sessId, connId)
		if bDeleted {
			s.sendToSessionClientClosed(id, ctx.sessionId, ctx.getClientIp(conn.Codec().Context()))
			logx.Infof("onServerConnectionClosed - sendClientClosed: {peer:%s, ctx:%s}", conn, ctx)
		}
	}
}

// //////////////////////////////////////////////////////////////////////////////////////////////////
func (s *Server) onUnencryptedMessage(ctx *connContext, conn *net2.TcpConnection, mmsg *mtproto.MTPRawMessage) error {
	logx.Info("receive unencryptedRawMessage: {peer: %s, ctx: %s, mmsg: %s}", conn, ctx, mmsg)

	if len(mmsg.Payload) < 8 {
		err := fmt.Errorf("invalid data len < 8")
		logx.Error(err.Error())
		return err
	}

	_, obj, err := parseFromIncomingMessage(mmsg.Payload[8:])
	if err != nil {
		err := fmt.Errorf("invalid data len < 8")
		logx.Errorf(err.Error())
	}

	x := mtproto.NewEncodeBuf(512)

	switch request := obj.(type) {
	case *mtproto.TLReqPq:
		logx.Infof("TLReqPq - {\"request\":%s", request)
		resPQ, err := s.handshake.onReqPq(request)
		if err != nil {
			logx.Errorf("onHandshake error: {%v} - {peer: %s, ctx: %s, mmsg: %s}", err, conn, ctx, mmsg)
			conn.Close()
			return err
		}
		ctx.putHandshakeStateCtx(&HandshakeStateCtx{
			State:       STATE_pq_res,
			Nonce:       resPQ.GetNonce(),
			ServerNonce: resPQ.GetServerNonce(),
		})
		serializeToBuffer(x, mtproto.GenerateMessageId(), resPQ)
	case *mtproto.TLReqPqMulti:
		logx.Infof("TLReqPqMulti - {\"request\":%s", request)
		resPQ, err := s.handshake.onReqPqMulti(request)
		if err != nil {
			logx.Errorf("onHandshake error: {%v} - {peer: %s, ctx: %s, mmsg: %s}", err, conn, ctx, mmsg)
			conn.Close()
			return err
		}
		ctx.putHandshakeStateCtx(&HandshakeStateCtx{
			State:       STATE_pq_res,
			Nonce:       resPQ.GetNonce(),
			ServerNonce: resPQ.GetServerNonce(),
		})
		serializeToBuffer(x, mtproto.GenerateMessageId(), resPQ)
	case *mtproto.TLReq_DHParams:
		logx.Infof("TLReq_DHParams - {\"request\":%s", request)
		if state := ctx.getHandshakeStateCtx(request.Nonce); state != nil {
			resServerDHParam, err := s.handshake.onReqDHParams(state, obj.(*mtproto.TLReq_DHParams))
			if err != nil {
				logx.Errorf("onHandshake error: {%v} - {peer: %s, ctx: %s, mmsg: %s}", err, conn, ctx, mmsg)
				conn.Close()
				return err
			}
			state.State = STATE_DH_params_res
			serializeToBuffer(x, mtproto.GenerateMessageId(), resServerDHParam)
		} else {
			logx.Errorf("onHandshake error: {invalid nonce} - {peer: %s, ctx: %s, mmsg: %s}", conn, ctx, mmsg)
			return conn.Close()
		}
	case *mtproto.TLSetClient_DHParams:
		logx.Infof("TLSetClient_DHParams - {\"request\":%s", request)
		if state := ctx.getHandshakeStateCtx(request.Nonce); state != nil {
			resSetClientDHParamsAnswer, err := s.handshake.onSetClientDHParams(state, obj.(*mtproto.TLSetClient_DHParams))
			if err != nil {
				logx.Errorf("onHandshake error: {%v} - {peer: %s, ctx: %s, mmsg: %s}", err, conn, ctx, mmsg)
				return conn.Close()
			}
			state.State = STATE_dh_gen_res
			serializeToBuffer(x, mtproto.GenerateMessageId(), resSetClientDHParamsAnswer)
		} else {
			logx.Errorf("onHandshake error: {invalid nonce} - {peer: %s, ctx: %s, mmsg: %s}", conn, ctx, mmsg)
			return conn.Close()
		}
	case *mtproto.TLMsgsAck:
		logx.Infof("TLMsgsAck - {\"request\":%s", request)
		//err = s.onMsgsAck(state, obj.(*mtproto.TLMsgsAck))
		//return nil, err
		return nil
	default:
		err = fmt.Errorf("invalid handshake type")
		return conn.Close()
	}
	return conn.Send(&mtproto.MTPRawMessage{Payload: x.GetBuf()})
}

func (s *Server) onEncryptedMessage(ctx *connContext, conn *net2.TcpConnection, authKey *authKeyUtil, mmsg *mtproto.MTPRawMessage) error {
	mtpRwaData, err := authKey.AesIgeDecrypt(mmsg.Payload[8:8+16], mmsg.Payload[24:])
	if err != nil {
		logx.Errorf("conn(%s) decrypt error: {%v}", conn.String(), err)
		return err
	}

	var (
		sessionId = int64(binary.LittleEndian.Uint64(mtpRwaData[8:]))
		isNew     = ctx.sessionId != sessionId
		authKeyId = mmsg.AuthKeyId()
	)

	sessClient, err2 := s.session.getSessionClient(strconv.FormatInt(mmsg.AuthKeyId(), 10))
	if err2 != nil {
		logx.Errorf("conn(%s) getSessionClient error: %v, {authKeyId: %d}", conn.String(), err, mmsg.AuthKeyId())
		return err2
	}

	if isNew {
		ctx.sessionId = sessionId
	} else {
		// check sessionId??
	}

	if isNew {
		if s.authSessionMgr.AddNewSession(authKey, sessionId, conn.GetConnID()) {
			sessClient.SessionCreateSession(context.Background(),
				&sessionpb.TLSessionCreateSession{
					Client: sessionpb.MakeTLSessionClientEvent(&sessionpb.SessionClientEvent{
						ServerId:  s.session.gatewayId,
						AuthKeyId: authKeyId,
						SessionId: sessionId,
						ClientIp:  ctx.getClientIp(conn.Codec().Context()),
					}).To_SessionClientEvent(),
				})
		}
	}

	_, _ = sessClient.SessionSendDataToSession(context.Background(), &sessionpb.TLSessionSendDataToSession{
		Data: &sessionpb.SessionClientData{
			ServerId:  s.session.gatewayId,
			AuthKeyId: authKey.AuthKeyId(),
			SessionId: sessionId,
			Salt:      int64(binary.LittleEndian.Uint64(mtpRwaData)),
			Payload:   mtpRwaData[16:],
			ClientIp:  ctx.getClientIp(conn.Codec().Context()),
		},
	})

	return nil
}

func (s *Server) GetConnByConnID(id uint64) *net2.TcpConnection {
	return s.server.GetConnection(id)
}

func (s *Server) SendToClient(conn *net2.TcpConnection, authKey *authKeyUtil, b []byte) error {
	ctx, _ := conn.Context.(*connContext)
	if ctx.trd != nil {
		logx.Info("del conn timeout")
		s.timer.Del(ctx.trd)
		ctx.trd = nil
	}

	msgKey, mtpRawData, _ := authKey.AesIgeEncrypt(b)
	x := mtproto.NewEncodeBuf(8 + len(msgKey) + len(mtpRawData))
	x.Long(authKey.AuthKeyId())
	x.Bytes(msgKey)
	x.Bytes(mtpRawData)
	//logx.Info("egate receiveData - ready sendToClient to: {peer: %s, auth_key_id = %d, session_id = %d}",
	//	conn,
	//	r.AuthKeyId,
	//	r.SessionId)

	msg := &mtproto.MTPRawMessage{Payload: x.GetBuf()}
	if ctx.isHttp {
		//if !ctx.canSend {
		//	s.authSessionMgr.PushBackHttpData(authKey.AuthKeyId(), ctx.sessionId, msg)
		//	return nil
		//}
		ctx.canSend = false
	}

	// err := conn.Send(&mtproto.MTPRawMessage{Payload: x.GetBuf()})
	err := conn.Send(msg)
	if err != nil {
		logx.Errorf("send error: %v", err)
		return err
	}

	return nil
}

func (s *Server) sendToSessionClientNew(authKeyId, sessionId int64, clientIp string) {
	c, err := s.session.getSessionClient(strconv.FormatInt(authKeyId, 10))
	if err != nil {
		logx.Errorf("getSessionClient error: {%v} - {authKeyId: %d}", err, authKeyId)
		return
	}

	c.SessionCreateSession(context.Background(), &sessionpb.TLSessionCreateSession{
		Client: &sessionpb.SessionClientEvent{
			ServerId:  s.session.gatewayId,
			AuthKeyId: authKeyId,
			SessionId: sessionId,
			ClientIp:  clientIp,
		},
	})
}

func (s *Server) sendToSessionClientClosed(authKeyId, sessionId int64, clientIp string) {
	c, err := s.session.getSessionClient(strconv.FormatInt(authKeyId, 10))
	if err != nil {
		logx.Errorf("getSessionClient error: {%v} - {authKeyId: %d}", err, authKeyId)
		return
	}

	c.SessionCloseSession(context.Background(), &sessionpb.TLSessionCloseSession{
		Client: &sessionpb.SessionClientEvent{
			ServerId:  s.session.gatewayId,
			AuthKeyId: authKeyId,
			SessionId: sessionId,
			ClientIp:  clientIp,
		},
	})
}
