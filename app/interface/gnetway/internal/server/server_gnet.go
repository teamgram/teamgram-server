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
	"context"
	"encoding/binary"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/interface/gnetway/internal/server/ws"
	"github.com/teamgram/teamgram-server/app/interface/session/client"
	"github.com/teamgram/teamgram-server/app/interface/session/session"

	"github.com/gobwas/ws/wsutil"
	"github.com/panjf2000/gnet/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

func (s *Server) asyncRun(connId int64, execb func() error, retcb func(c gnet.Conn)) {
	s.pool.Submit(func() {
		if err := execb(); err == nil {
			s.eng.Trigger(connId, func(c gnet.Conn) {
				retcb(c)
			})
		} else {
			// do nothing
		}
	})
}

// OnBoot fires when the engine is ready for accepting connections.
// The parameter engine has information and various utilities.
func (s *Server) OnBoot(eng gnet.Engine) (action gnet.Action) {
	logx.Debugf("gnetway server is listening")
	s.eng = eng
	return gnet.None
}

// OnShutdown fires when the engine is being shut down, it is called right after
// all event-loops and connections are closed.
func (s *Server) OnShutdown(eng gnet.Engine) {
	_ = eng
	logx.Debugf("gnetway server shutdown")
}

// OnOpen fires when a new connection has been opened.
// The parameter out is the return value which is going to be sent back to the peer.
func (s *Server) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	logx.Debugf("onNewConn - conn(%s)", c)

	ctx := newConnContext()
	ctx.setClientIp(strings.Split(c.RemoteAddr().String(), ":")[0])
	ctx.tcp = s.c.Gnetway.IsTcp(c.LocalAddr().String())
	ctx.websocket = s.c.Gnetway.IsWebsocket(c.LocalAddr().String())
	if ctx.websocket {
		ctx.wsCodec = new(ws.WsCodec)
	}
	c.SetContext(ctx)

	return
}

// OnClose fires when a connection has been closed.
// The parameter err is the last known connection error.
func (s *Server) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	logx.Debugf("onConnClosed - conn(%s), err: %v", c, err)

	ctx, _ := c.Context().(*connContext)
	if ctx == nil {
		return
	}

	sessId, connId, clientIp := ctx.sessionId, c.ConnId(), ctx.clientIp
	for _, id := range ctx.getAllAuthKeyId() {
		bDeleted := s.authSessionMgr.RemoveSession(id, sessId, connId)
		if bDeleted {
			s.pool.Submit(func() {
				s.sendToSessionClientClosed(id, sessId, clientIp)
			})
			logx.Debugf("onServerConnectionClosed - sendClientClosed: {conn: %s}", c)
		}
	}

	if ctx.wsCodec != nil {
		ctx.wsCodec.Conn.Release()
	}

	c.SetContext(nil)

	return
}

// OnTraffic fires when a local socket receives data from the peer.
func (s *Server) OnTraffic(c gnet.Conn) (action gnet.Action) {
	ctx := c.Context().(*connContext)

	if ctx.websocket {
		return s.onWebsocketData(ctx, c)
	} else {
		return s.onTcpData(ctx, c)
	}
}

// OnTick fires immediately after the engine starts and will fire again
// following the duration specified by the delay return value.
func (s *Server) OnTick() (delay time.Duration, action gnet.Action) {
	logx.Statf("conn count: %d", s.eng.CountConnections())
	delay = time.Second * 5
	return
}

func (s *Server) onEncryptedMessage(c gnet.Conn, ctx *connContext, authKey *authKeyUtil, mmsg *mtproto.MTPRawMessage) error {
	mtpRwaData, err := authKey.AesIgeDecrypt(mmsg.Payload[8:8+16], mmsg.Payload[24:])
	if err != nil {
		logx.Errorf("conn(%s) decrypt error: {%v}", c, err)
		return err
	}

	var (
		sessionId = int64(binary.LittleEndian.Uint64(mtpRwaData[8:]))
		isNew     = ctx.sessionId != sessionId
		authKeyId = mmsg.AuthKeyId()
		clientIp  = ctx.clientIp
		connId    = c.ConnId()
	)
	if isNew {
		ctx.sessionId = sessionId
	} else {
		// check sessionId??
	}

	if isNew {
		isNew = s.authSessionMgr.AddNewSession(authKey, sessionId, connId)
	}

	s.pool.Submit(func() {
		sessClient, err2 := s.session.getSessionClient(strconv.FormatInt(mmsg.AuthKeyId(), 10))
		if err2 != nil {
			logx.Errorf("conn(%s) getSessionClient error: %v, {authKeyId: %d}", c, err, mmsg)
			return
		}
		// TODO: close

		if isNew {
			sessClient.SessionCreateSession(context.Background(),
				&session.TLSessionCreateSession{
					Client: session.MakeTLSessionClientEvent(&session.SessionClientEvent{
						ServerId:  s.session.gatewayId,
						AuthKeyId: authKeyId,
						SessionId: sessionId,
						ClientIp:  clientIp,
					}).To_SessionClientEvent(),
				})
		}

		_, err = sessClient.SessionSendDataToSession(context.Background(), &session.TLSessionSendDataToSession{
			Data: &session.SessionClientData{
				ServerId:  s.session.gatewayId,
				AuthKeyId: authKey.AuthKeyId(),
				SessionId: sessionId,
				Salt:      int64(binary.LittleEndian.Uint64(mtpRwaData)),
				Payload:   mtpRwaData[16:],
				ClientIp:  clientIp,
			},
		})
		if err != nil {
			logx.Errorf("session.sendDataToSession - error: %v", err)
		}
	})

	return nil
}

func (s *Server) sendToSessionClientNew(authKeyId, sessionId int64, clientIp string) {
	c, err := s.session.getSessionClient(strconv.FormatInt(authKeyId, 10))
	if err != nil {
		logx.Errorf("getSessionClient error: {%v} - {authKeyId: %d}", err, authKeyId)
		return
	}

	c.SessionCreateSession(context.Background(), &session.TLSessionCreateSession{
		Client: &session.SessionClientEvent{
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

	c.SessionCloseSession(context.Background(), &session.TLSessionCloseSession{
		Client: &session.SessionClientEvent{
			ServerId:  s.session.gatewayId,
			AuthKeyId: authKeyId,
			SessionId: sessionId,
			ClientIp:  clientIp,
		},
	})
}

func (s *Server) GetConnCounts() int {
	return s.eng.CountConnections()
}

func (s *Server) onMTPRawMessage(ctx *connContext, c gnet.Conn, msg2 *mtproto.MTPRawMessage) (action gnet.Action) {
	var (
		err error
		out interface{}
	)

	if msg2.AuthKeyId() == 0 {
		out, err = s.onHandshake(c, msg2)
		if err != nil {
			action = gnet.Close
			return
		}
	} else {
		authKey := ctx.getAuthKey(msg2.AuthKeyId())
		if authKey == nil {
			key := s.GetAuthKey(msg2.AuthKeyId())
			if key != nil {
				authKey = newAuthKeyUtil(key)
				ctx.putAuthKey(authKey)
			}
		}

		if authKey == nil {
			var (
				err2       error
				sessClient sessionclient.SessionClient
				key3       *mtproto.AuthKeyInfo
				key2       *mtproto.AuthKeyInfo
			)
			sessClient, err2 = s.session.getSessionClient(strconv.FormatInt(msg2.AuthKeyId(), 10))
			if err2 != nil {
				logx.Errorf("getSessionClient error: %v, {authKeyId: %d}", err2, msg2.AuthKeyId())
				action = gnet.Close
				return
			}

			if key3, err2 = sessClient.SessionQueryAuthKey(context.Background(), &session.TLSessionQueryAuthKey{
				AuthKeyId: msg2.AuthKeyId(),
			}); err2 != nil {
				logx.Errorf("conn(%s) sessionQueryAuthKey error: %v", c, err2)
				if errors.Is(err2, mtproto.ErrAuthKeyUnregistered) {
					out2 := &mtproto.MTPRawMessage{
						Payload: make([]byte, 4),
					}
					var (
						code = int32(-404)
					)
					binary.LittleEndian.PutUint32(out2.Payload, uint32(code))
					out = out2
					action = gnet.Close
				} else {
					out = nil
					action = gnet.Close
					return
				}
			} else {
				key2 = &mtproto.AuthKeyInfo{
					AuthKeyId:          key3.AuthKeyId,
					AuthKey:            key3.AuthKey,
					AuthKeyType:        key3.AuthKeyType,
					PermAuthKeyId:      key3.PermAuthKeyId,
					TempAuthKeyId:      key3.TempAuthKeyId,
					MediaTempAuthKeyId: key3.MediaTempAuthKeyId,
				}
				s.PutAuthKey(key2)
				authKey = newAuthKeyUtil(key2)
				ctx.putAuthKey(authKey)
			}
			err = s.onEncryptedMessage(c, ctx, authKey, msg2)
			if err != nil {
				action = gnet.Close
				return
			}
		} else {
			err = s.onEncryptedMessage(c, ctx, authKey, msg2)
			if err != nil {
				action = gnet.Close
				return
			}
		}
	}

	if out != nil {
		UnThreadSafeWrite(c, out)
	}

	return gnet.None
}

func UnThreadSafeWrite(c gnet.Conn, msg interface{}) error {
	ctx := c.Context().(*connContext)

	if ctx.codec == nil {
		return nil
	}
	data, err := ctx.codec.Encode(c, msg)
	if err != nil {
		return err
	}

	if ctx.websocket {
		// This is the echo server
		err = wsutil.WriteServerBinary(c, data)
		if err != nil {
			logx.Infof("conn[%v] [err=%v]", c.RemoteAddr().String(), err.Error())
			return err
		}
	} else {
		_, err = c.Write(data)
	}

	return nil
}
