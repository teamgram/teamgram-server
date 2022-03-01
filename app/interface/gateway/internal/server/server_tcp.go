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
	"strconv"
	"strings"
	"time"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/interface/gateway/internal/server/codec"
	"github.com/teamgram/teamgram-server/app/interface/session/client"
	"github.com/teamgram/teamgram-server/app/interface/session/session"

	"github.com/panjf2000/gnet"
	"github.com/zeromicro/go-zero/core/logx"
)

type connContext struct {
	authKey   *authKeyUtil
	sessionId int64
	clientIp  string
}

func newConnContext(clientIp string) *connContext {
	return &connContext{
		authKey:   nil,
		sessionId: 0,
		clientIp:  clientIp,
	}
}

func (s *Server) asyncRun(connId int64, execb func() error, retcb func(c gnet.Conn)) {
	s.pool.Submit(func() {
		if err := execb(); err == nil {
			s.svr.Trigger(connId, func(c gnet.Conn) {
				retcb(c)
			})
		} else {
			// do nothing
		}
	})
}

// OnInitComplete fires when the server is ready for accepting connections.
// The parameter:server has information and various utilities.
func (s *Server) OnInitComplete(svr gnet.Server) (action gnet.Action) {
	logx.Infof("egate server is listening on [%s] (multi-cores: %t, loops: %d)",
		svr.AddrsString(), svr.Multicore, svr.NumEventLoop)
	s.svr = svr
	return
}

// OnShutdown fires when the server is being shut down, it is called right after
// all event-loops and connections are closed.
func (s *Server) OnShutdown(svr gnet.Server) {
	logx.Infof("egate server shutdown")
}

// OnOpened fires when a new connection has been opened.
// The parameter:c has information about the connection such as it's local and remote address.
// Parameter:out is the return value which is going to be sent back to the client.
func (s *Server) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	logx.Infof("onNewConn - conn(%s)", c.DebugString())
	return
}

// OnClosed fires when a connection has been closed.
// The parameter:err is the last known connection error.
func (s *Server) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	logx.Infof("onConnClosed - conn(%s), err: %v", c.DebugString(), err)
	if ctx, ok := c.Context().(*connContext); ok && ctx != nil {
		if ctx.authKey != nil {
			bDeleted := s.authSessionMgr.RemoveSession(ctx.authKey.AuthKeyId(), ctx.sessionId, c.ConnID())
			if bDeleted {
				s.pool.Submit(func() {
					s.sendToSessionClientClosed(ctx.authKey.AuthKeyId(), ctx.sessionId, ctx.clientIp)
				})
			}
		}
	}

	c.SetContext(nil)

	return
}

// PreWrite fires just before any data is written to any client socket, this event function is usually used to
// put some code of logging/counting/reporting or any prepositive operations before writing data to client.
func (s *Server) PreWrite(c gnet.Conn) {
	// log.Debugf("preWrite")
}

// React fires when a connection sends the server data.
// Call c.Read() or c.ReadN(n) within the parameter:c to read incoming data from client.
// Parameter:out is the return value which is going to be sent back to the client.
func (s *Server) React(frame interface{}, c gnet.Conn) (out interface{}, action gnet.Action) {
	if frame == nil {
		logx.Errorf("conn(%s) frame is nil", c.DebugString())
		return
	}

	msg2, ok := frame.(*codec.MTPRawMessage)
	if !ok {
		logx.Errorf("onReact - conn(%s) recv error: msg2 not codec.MTPRawMessage type", c.DebugString())
		action = gnet.Close
		out = nil
		return
	}

	logx.Infof("onReact - conn(%s) recv frame: %s", c.DebugString(), msg2.String())

	var err error
	if msg2.GetAuthKeyId() == 0 {
		out, err = s.onHandshake(c, msg2)
		if err != nil {
			action = gnet.Close
			out = nil
			return
		}
	} else {
		ctx, ok := c.Context().(*connContext)
		if !ok || ctx == nil {
			ctx = newConnContext(strings.Split(c.RemoteAddr().String(), ":")[0])
			c.SetContext(ctx)
		}

		authKey := ctx.authKey
		if authKey == nil {
			key := s.GetAuthKey(msg2.GetAuthKeyId())
			if key != nil {
				authKey = newAuthKeyUtil(key)
				ctx.authKey = authKey
			}
		}

		if authKey == nil {
			//s.asyncRun(c.ConnID(),
			//	func() error {
			var (
				err2       error
				sessClient session_client.SessionClient
				key3       *mtproto.AuthKeyInfo
				key2       *mtproto.AuthKeyInfo
			)
			key2 = s.GetAuthKey(msg2.GetAuthKeyId())
			if key2 != nil {
				authKey = newAuthKeyUtil(key2)
			} else {
				sessClient, err2 = s.session.getSessionClient(strconv.FormatInt(msg2.GetAuthKeyId(), 10))
				if err2 != nil {
					logx.Errorf("getSessionClient error: %v, {authKeyId: %d}", err2, msg2.GetAuthKeyId())
					//// return err2
					//out2 := &codec.MTPRawMessage{
					//	Payload: make([]byte, 4),
					//}
					//var code = int32(-404)
					//binary.LittleEndian.PutUint32(out2.Payload, uint32(code))
					//out = out2
					action = gnet.Close
					out = nil
					return
				}

				if key3, err2 = sessClient.SessionQueryAuthKey(context.Background(), &session.TLSessionQueryAuthKey{
					AuthKeyId: msg2.GetAuthKeyId(),
				}); err2 != nil {
					logx.Errorf("conn(%s) sessionQueryAuthKey error: %v", c.DebugString(), err2)
					//// return err2
					out2 := &codec.MTPRawMessage{
						Payload: make([]byte, 4),
					}
					var code = int32(-404)
					binary.LittleEndian.PutUint32(out2.Payload, uint32(code))
					out = out2
					action = gnet.Close
					// out = nil
					return
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
				}
			}
			//	return nil
			//},
			//func(c gnet.Conn) {
			ctx.authKey = authKey
			s.onEncryptedMessage(c, ctx, authKey, msg2)
			//})

			//if authKey == nil {
			//	out2 := &codec.MTPRawMessage{
			//		Payload: make([]byte, 4),
			//	}
			//	var code = int32(-404)
			//	binary.LittleEndian.PutUint32(out2.Payload, uint32(code))
			//	out = out2
			//}
		} else {
			err = s.onEncryptedMessage(c, ctx, authKey, msg2)
		}
	}

	return
}

// Tick fires immediately after the server starts and will fire again
// following the duration specified by the delay return value.
func (s *Server) Tick() (delay time.Duration, action gnet.Action) {
	// log.Debugf("tick")
	return
}

func (s *Server) onEncryptedMessage(c gnet.Conn, ctx *connContext, authKey *authKeyUtil, mmsg *codec.MTPRawMessage) error {
	logx.Infof("conn(%s) onEncryptedMessage: len(%d)", c.DebugString(), len(mmsg.Payload))
	mtpRwaData, err := authKey.AesIgeDecrypt(mmsg.Payload[8:8+16], mmsg.Payload[24:])
	if err != nil {
		logx.Errorf("conn(%s) decrypt error: {%v}", c.DebugString(), err)
		return err
	}

	var (
		sessionId = int64(binary.LittleEndian.Uint64(mtpRwaData[8:]))
		isNew     = ctx.sessionId == 0
		authKeyId = mmsg.GetAuthKeyId()
	)
	if isNew {
		ctx.sessionId = sessionId
	} else {
		// check sessionId??
	}

	s.pool.Submit(func() {
		sessClient, err2 := s.session.getSessionClient(strconv.FormatInt(mmsg.GetAuthKeyId(), 10))
		if err2 != nil {
			logx.Errorf("conn(%s) getSessionClient error: %v, {authKeyId: %d}", c.DebugString(), err, mmsg.GetAuthKeyId())
			return
		}

		if isNew {
			if s.authSessionMgr.AddNewSession(authKey, sessionId, c.ConnID()) {
				sessClient.SessionCreateSession(context.Background(),
					&session.TLSessionCreateSession{
						Client: session.MakeTLSessionClientEvent(&session.SessionClientEvent{
							ServerId:  s.session.gatewayId,
							AuthKeyId: authKeyId,
							SessionId: sessionId,
							ClientIp:  ctx.clientIp,
						}).To_SessionClientEvent(),
					})
			}
		}

		_, _ = sessClient.SessionSendDataToSession(context.Background(), &session.TLSessionSendDataToSession{
			Data: &session.SessionClientData{
				ServerId:  s.session.gatewayId,
				AuthKeyId: authKey.AuthKeyId(),
				SessionId: sessionId,
				Salt:      int64(binary.LittleEndian.Uint64(mtpRwaData)),
				Payload:   mtpRwaData[16:],
				ClientIp:  ctx.clientIp,
			},
		})
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
	return s.svr.CountConnections()
}
