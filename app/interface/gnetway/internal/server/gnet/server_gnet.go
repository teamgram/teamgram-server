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

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/interface/gnetway/internal/server/gnet/ws"
	"github.com/teamgram/teamgram-server/app/interface/session/client"
	"github.com/teamgram/teamgram-server/app/interface/session/session"

	"github.com/gobwas/ws/wsutil"
	"github.com/panjf2000/gnet/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/timex"
	"google.golang.org/protobuf/proto"
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
					Client: session.MakeTLSessionClientEvent(&session.SessionClientEvent{
						ServerId:      s.svcCtx.GatewayId,
						AuthKeyId:     ctx.authKey.AuthKeyId(),
						KeyType:       int32(ctx.authKey.AuthKeyType()),
						PermAuthKeyId: ctx.authKey.PermAuthKeyId(),
						SessionId:     ctx.sessionId,
						ClientIp:      ctx.clientIp,
					}).To_SessionClientEvent(),
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
	since := timex.Now()

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
		// hack
		for _, unknown := range tryGetUnknownTLObject(mtpRwaData[16:]) {
			switch unknownMsg := unknown.(type) {
			case *mtproto.TLAuthBindTempAuthKey:
				permAuthKeyId = unknownMsg.PermAuthKeyId

				clone := proto.Clone(authKey.keyData).(*mtproto.AuthKeyInfo)
				clone.PermAuthKeyId = permAuthKeyId
				authKey.keyData = clone
				s.PutAuthKey(clone)
			case *mtproto.TLPing:
				payload := serializeToBuffer2(salt, sessionId, &mtproto.TLMessage2{
					MsgId: nextMessageId(false),
					Seqno: 0,
					Bytes: 0,
					Object: mtproto.MakeTLPong(&mtproto.Pong{
						MsgId:  int64(binary.LittleEndian.Uint64(mtpRwaData[16:])),
						PingId: unknownMsg.PingId,
					}).To_Pong(),
				})

				msgKey, mtpRawData, _ := authKey.AesIgeEncrypt(payload)
				x2 := mtproto.NewEncodeBuf(8 + len(msgKey) + len(mtpRawData))
				x2.Long(authKey.AuthKeyId())
				x2.Bytes(msgKey)
				x2.Bytes(mtpRawData)
				_ = UnThreadSafeWrite(c, &mtproto.MTPRawMessage{Payload: x2.GetBuf()})

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
		defer func() {
			logx.WithDuration(timex.Since(since)).Infof("onEncryptedMessage: %s", c)
		}()

		_ = s.svcCtx.Dao.ShardingSessionClient.InvokeByKey(
			strconv.FormatInt(permAuthKeyId, 10),
			func(client sessionclient.SessionClient) (err error) {
				if isNew {
					if s.authSessionMgr.AddNewSession(authKey, sessionId, connId) {
						_, err = client.SessionCreateSession(context.Background(), &session.TLSessionCreateSession{
							Client: session.MakeTLSessionClientEvent(&session.SessionClientEvent{
								ServerId:      s.svcCtx.GatewayId,
								AuthKeyId:     authKey.AuthKeyId(),
								KeyType:       int32(authKey.AuthKeyType()),
								PermAuthKeyId: permAuthKeyId,
								SessionId:     sessionId,
								ClientIp:      clientIp,
							}).To_SessionClientEvent(),
						})
					}
				}

				_, err = client.SessionSendDataToSession(context.Background(), &session.TLSessionSendDataToSession{
					Data: &session.SessionClientData{
						ServerId:      s.svcCtx.GatewayId,
						AuthKeyId:     authKey.AuthKeyId(),
						KeyType:       int32(authKey.AuthKeyType()),
						PermAuthKeyId: permAuthKeyId,
						SessionId:     sessionId,
						Salt:          salt,
						Payload:       mtpRwaData[16:],
						ClientIp:      clientIp,
					},
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
		out, err := s.onHandshake(c, msg2)
		if err != nil {
			action = gnet.Close
		} else if out != nil {
			_ = UnThreadSafeWrite(c, out)
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
		} else {
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
						key3 *mtproto.AuthKeyInfo
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
					} else {
						s.PutAuthKey(key3)
					}

					return newAuthKeyUtil(key3), nil
				},
				func(c2 gnet.Conn, mmsg []byte, in interface{}, err error) {
					if err != nil {
						if errors.Is(err, mtproto.ErrAuthKeyUnregistered) {
							out2 := &mtproto.MTPRawMessage{
								Payload: make([]byte, 4),
							}
							var (
								code = int32(-404)
							)
							binary.LittleEndian.PutUint32(out2.Payload, uint32(code))
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
	}

	return gnet.None
}

func UnThreadSafeWrite(c gnet.Conn, msg interface{}) error {
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
