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
	"context"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/mt"
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/interface/gnetway/internal/server/gnet/ws"
	"github.com/teamgram/teamgram-server/v2/app/interface/session/client"
	"github.com/teamgram/teamgram-server/v2/app/interface/session/session"

	"github.com/gobwas/ws/wsutil"
	"github.com/panjf2000/gnet/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

func (s *Server) asyncRun(connId int64, execb func() error, retcb func(c gnet.Conn)) {
	_ = s.pool.Submit(func() {
		if err := execb(); err == nil {
			s.eng.Trigger(connId, func(c gnet.Conn) {
				retcb(c)
			})
		} else {
			// do nothing
		}
	})
}

func (s *Server) asyncRun2(
	connId int64,
	mmsg []byte,
	execb func(mmsg []byte) (interface{}, error),
	retcb func(c gnet.Conn, mmsg []byte, in interface{}, err error)) {
	_ = s.pool.Submit(func() {
		r, err := execb(mmsg)
		s.eng.Trigger(connId, func(c gnet.Conn) {
			retcb(c, mmsg, r, err)
		})
	})
}

// OnBoot fires when the engine is ready for accepting connections.
// The parameter engine has information and various utilities.
func (s *Server) OnBoot(eng gnet.Engine) (action gnet.Action) {
	logx.Infof("gnetway server is listening")
	s.eng = eng
	return gnet.None
}

// OnShutdown fires when the engine is being shut down, it is called right after
// all event-loops and connections are closed.
func (s *Server) OnShutdown(eng gnet.Engine) {
	_ = eng
	logx.Infof("gnetway server shutdown")
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
	ctx.closeDate = time.Now().Unix() + 30
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

	defer func() {
		if ctx.wsCodec != nil {
			ctx.wsCodec.Conn.Release()
		}

		c.SetContext(nil)
	}()

	if ctx.authKey == nil || ctx.authKey.PermAuthKeyId() == 0 {
		return
	}

	// kId, sessId, connId, clientIp := ctx.authKey.AuthKeyId(), ctx.sessionId, c.ConnId(), ctx.clientIp
	bDeleted := s.authSessionMgr.RemoveSession(ctx.authKey.AuthKeyId(), ctx.sessionId, c.ConnId())
	if !bDeleted {
		return
	}

	_ = s.pool.Submit(func() {
		_ = s.svcCtx.ShardingSessionClient.InvokeByKey(
			strconv.FormatInt(ctx.authKey.PermAuthKeyId(), 10),
			func(client sessionclient.SessionClient) (err error) {
				_, err = client.SessionCloseSession(context.Background(), &session.TLSessionCloseSession{
					Client: session.MakeSessionClientEvent(&session.TLSessionClientEvent{
						ServerId:      s.svcCtx.GatewayId,
						AuthKeyId:     ctx.authKey.AuthKeyId(),
						KeyType:       int32(ctx.authKey.AuthKeyType()),
						PermAuthKeyId: ctx.authKey.PermAuthKeyId(),
						SessionId:     ctx.sessionId,
						ClientIp:      ctx.clientIp,
					}),
				})
				return
			})
	})

	return
}

// OnTraffic fires when a local socket receives data from the peer.
func (s *Server) OnTraffic(c gnet.Conn) (action gnet.Action) {
	ctx := c.Context().(*connContext)
	ctx.closeDate = time.Now().Unix() + 300 + rand.Int63()%10
	if ctx.websocket {
		return s.onWebsocketData(ctx, c)
	} else {
		return s.onTcpData(ctx, c)
	}
}

// OnTick fires immediately after the engine starts and will fire again
// following the duration specified by the delay return value.
func (s *Server) OnTick() (delay time.Duration, action gnet.Action) {
	s.tickNumber = s.tickNumber + 1
	if s.tickNumber%5 == 0 {
		logx.Statf("conn count: %d", s.eng.CountConnections())
	}
	delay = time.Second * 1
	now := time.Now().Unix()

	s.eng.Iterate(func(c gnet.Conn) {
		ctx, _ := c.Context().(*connContext)
		if ctx == nil {
			return
		}
		if now >= ctx.closeDate {
			logx.Debugf("close conn(%s) by timeout", c)
			_ = c.Close()
		}
	})
	return
}

func (s *Server) onEncryptedMessage(c gnet.Conn, ctx *connContext, authKey *authKeyUtil, needAck bool, mmsg []byte) error {
	mtpRwaData, err := authKey.AesIgeDecrypt(mmsg[8:8+16], mmsg[24:])
	if err != nil {
		logx.Errorf("conn(%s) decrypt data(%d) error: {%v}, payload: %s", c, len(mmsg)-24, err, hex.EncodeToString(mmsg))
		return err
	}

	var (
		permAuthKeyId = authKey.PermAuthKeyId()
		salt          = int64(binary.LittleEndian.Uint64(mtpRwaData))
		sessionId     = int64(binary.LittleEndian.Uint64(mtpRwaData[8:]))
	)

	if permAuthKeyId == 0 {
		x3 := bin.NewEncoder()
		defer x3.End()

		// hack
		for _, unknown := range tryGetUnknownTLObject(mtpRwaData[16:]) {
			switch unknownMsg := unknown.(type) {
			case *tg.TLAuthBindTempAuthKey:
				permAuthKeyId = unknownMsg.PermAuthKeyId

				clone := authKey.CloneKeyData()
				clone.PermAuthKeyId = permAuthKeyId
				authKey.keyData = clone
				s.PutAuthKey(clone)
			case *mt.TLPing:
				payload := serializeToBuffer2(salt, sessionId, &mt.TLMessage2{
					MsgId: nextMessageId(false),
					Seqno: 0,
					Bytes: 0,
					Object: mt.MakePong(&mt.TLPong{
						MsgId:  int64(binary.LittleEndian.Uint64(mtpRwaData[16:])),
						PingId: unknownMsg.PingId,
					}),
				})

				msgKey, mtpRawData, _ := authKey.AesIgeEncrypt(payload)
				x3.PutInt64(authKey.AuthKeyId())
				x3.Put(msgKey)
				x3.Put(mtpRawData)
				_ = UnThreadSafeWrite(c, x3.Bytes())
				x3.Reset()

				return nil
			default:
				if permAuthKeyId == 0 {
					logx.Errorf("recv unknown msg: %v, ignore it", unknownMsg)
					return fmt.Errorf("unknown msg")
				} else {
					// ignore it
				}
			}
		}
	}

	var (
		isNew    = ctx.sessionId != sessionId
		clientIp = ctx.clientIp
		connId   = c.ConnId()
	)
	if isNew {
		ctx.sessionId = sessionId
	} else {
		// check sessionId??
	}

	_ = s.pool.Submit(func() {
		_ = s.svcCtx.Dao.ShardingSessionClient.InvokeByKey(
			strconv.FormatInt(permAuthKeyId, 10),
			func(client sessionclient.SessionClient) (err error) {
				if isNew {
					if s.authSessionMgr.AddNewSession(authKey, sessionId, connId) {
						_, err = client.SessionCreateSession(context.Background(), &session.TLSessionCreateSession{
							Client: session.MakeSessionClientEvent(&session.TLSessionClientEvent{
								ServerId:      s.svcCtx.GatewayId,
								AuthKeyId:     authKey.AuthKeyId(),
								KeyType:       int32(authKey.AuthKeyType()),
								PermAuthKeyId: permAuthKeyId,
								SessionId:     sessionId,
								ClientIp:      clientIp,
							}),
						})
					}
				}

				_, err = client.SessionSendDataToSession(context.Background(), &session.TLSessionSendDataToSession{
					Data: session.MakeSessionClientData(&session.TLSessionClientData{
						ServerId:      s.svcCtx.GatewayId,
						AuthKeyId:     authKey.AuthKeyId(),
						KeyType:       int32(authKey.AuthKeyType()),
						PermAuthKeyId: permAuthKeyId,
						SessionId:     sessionId,
						Salt:          salt,
						Payload:       mtpRwaData[16:],
						ClientIp:      clientIp,
					}),
				})
				if err != nil {
					logx.Errorf("session.sendDataToSession - error: %v", err)
				}

				return
			})
	})

	return nil
}

func (s *Server) GetConnCounts() int {
	return s.eng.CountConnections()
}

func (s *Server) onMTPRawMessage(ctx *connContext, c gnet.Conn, authKeyId int64, needAck bool, msg2 []byte) (action gnet.Action) {
	if authKeyId == 0 {
		/**
			### Unencrypted Messages

			Special plain-text messages may be used to create an authorization key as well as to perform a time
			synchronization. They begin with auth_key_id = 0 (64 bits) which means that there is no auth_key. This is followed
			directly by the message body in serialized format without internal or external headers. A message identifier (64 bits)
			and body length in bytes (32 bytes) are added before the message body.

			Only a very limited number of messages of special types can be transmitted as plain text.

			#### Unencrypted Message

			| **auth_key_id** = `0` | **message_id** | **message_data_length** | **message_data** |
			| --------------------- | -------------- | ----------------------- | ---------------- |
			| int64                 | int64          | int32                   | bytes            |
		**/

		if len(msg2) < 8+8+4 {
			// err := fmt.Errorf("invalid data len < 8")
			logx.Errorf("conn(%s) error: %v, invalid data len < 8+8+4", c)
			action = gnet.Close
			return
		}

		d := bin.NewDecoder(msg2[8:])
		msgId, _ := d.Int64()
		_ = msgId
		// Check msgId

		dataLen, _ := d.Int32()
		if len(msg2) < 8+8+4+int(dataLen) {
			logx.Errorf("conn(%s) error: invalid data len < 8+8+4+int(dataLen)", c)
			action = gnet.Close
			return
		}

		err := s.onHandshake(c, d)
		if err != nil {
			logx.Errorf("conn(%s) error: %v", c, err)
			action = gnet.Close
		}
	} else {
		authKey := ctx.getAuthKey()
		if authKey == nil {
			key := s.GetAuthKey(authKeyId)
			if key != nil {
				authKey = newAuthKeyUtil(key)
				ctx.putAuthKey(authKey)
			}
		} else if authKey.AuthKeyId() != authKeyId {
			//
			action = gnet.Close
			return
		}

		if authKey != nil {
			err := s.onEncryptedMessage(c, ctx, authKey, needAck, msg2)
			if err != nil {
				action = gnet.Close
			}
			return
		}
		if len(msg2) > 32 {
			logx.Debugf("conn(%s) data: %s", c, hex.EncodeToString(msg2[:32]))
		}

		msg2Clone := make([]byte, len(msg2))
		copy(msg2Clone, msg2)

		s.asyncRun2(
			c.ConnId(),
			msg2Clone,
			func(mmsg []byte) (interface{}, error) {
				var (
					key3 *tg.AuthKeyInfo
				)

				err2 := s.svcCtx.Dao.ShardingSessionClient.InvokeByKey(
					strconv.FormatInt(authKeyId, 10),
					func(client sessionclient.SessionClient) (err error) {
						key3, err = client.SessionQueryAuthKey(context.Background(), &session.TLSessionQueryAuthKey{
							AuthKeyId: authKeyId,
						})
						return
					})
				if err2 != nil {
					logx.Errorf("conn(%s) sessionQueryAuthKey error: %v", c, err2)
					return nil, err2
				}

				key4, _ := key3.ToAuthKeyInfo()

				s.PutAuthKey(key4)
				return newAuthKeyUtil(key4), nil
			},
			func(c2 gnet.Conn, mmsg []byte, in interface{}, err error) {
				if err != nil {
					if errors.Is(err, tg.ErrAuthKeyUnregistered) {
						out2 := make([]byte, 4)
						var (
							code = int32(-404)
						)
						binary.LittleEndian.PutUint32(out2, uint32(code))
						_ = UnThreadSafeWrite(c2, out2)
					}
					_ = c2.Close()
				} else {
					authKey2 := in.(*authKeyUtil)
					ctx2 := c2.Context().(*connContext)
					ctx2.putAuthKey(authKey2)
					err = s.onEncryptedMessage(c2, ctx2, authKey2, needAck, mmsg)
					if err != nil {
						_ = c2.Close()
					}
				}
			})
	}

	return gnet.None
}

func UnThreadSafeWrite(c gnet.Conn, msg []byte) error {
	ctx := c.Context().(*connContext)

	if ctx.codec == nil {
		return nil
	}

	if msg == nil {
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
