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
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/teamgram/proto/mtproto"
	httpcodec "github.com/teamgram/teamgram-server/app/interface/gnetway/internal/server/gnet/http"
	"github.com/teamgram/teamgram-server/app/interface/gnetway/internal/server/gnet/pp"
	"github.com/teamgram/teamgram-server/app/interface/gnetway/internal/server/gnet/ws"
	"github.com/teamgram/teamgram-server/app/interface/session/session"

	"github.com/gobwas/ws/wsutil"
	"github.com/panjf2000/gnet/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/timex"
	"go.opentelemetry.io/otel"
	"google.golang.org/protobuf/proto"
)

func (s *Server) asyncRunIfError(connId int64, msgId int64, execb func() error, retcb func(c gnet.Conn, msgId int64, err error)) {
	if err := s.pool.Submit(func() {
		if err := execb(); err != nil {
			s.eng.Trigger(connId, func(c gnet.Conn) {
				retcb(c, msgId, err)
			})
		}
	}); err != nil {
		logx.Errorf("asyncRunIfError - pool.Submit error: %v, connId: %d", err, connId)
	}
}

func (s *Server) asyncRun(connId int64, execb func() error, retcb func(c gnet.Conn)) {
	if err := s.pool.Submit(func() {
		if err := execb(); err == nil {
			s.eng.Trigger(connId, func(c gnet.Conn) {
				retcb(c)
			})
		}
	}); err != nil {
		logx.Errorf("asyncRun - pool.Submit error: %v, connId: %d", err, connId)
	}
}

func (s *Server) asyncRun2(
	connId int64,
	kId int64,
	needAck bool,
	mmsg []byte,
	execb func(kId int64, mmsg []byte) (interface{}, error),
	retcb func(c gnet.Conn, needAck bool, mmsg []byte, in interface{}, err error)) {
	if err := s.pool.Submit(func() {
		r, err := execb(kId, mmsg)
		s.eng.Trigger(connId, func(c gnet.Conn) {
			retcb(c, needAck, mmsg, r, err)
		})
	}); err != nil {
		logx.Errorf("asyncRun2 - pool.Submit error: %v, connId: %d, keyId: %d", err, connId, kId)
	}
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
	if host, _, err := net.SplitHostPort(c.RemoteAddr().String()); err == nil {
		ctx.setClientIp(host)
	} else {
		ctx.setClientIp(c.RemoteAddr().String())
	}
	ctx.ppv1 = s.c.Gnetway.IsProxyProtocolV1(c.LocalAddr().String())
	ctx.tcp = s.c.Gnetway.IsTcp(c.LocalAddr().String())
	ctx.websocket = s.c.Gnetway.IsWebsocket(c.LocalAddr().String())
	ctx.http = s.c.Gnetway.IsHttp(c.LocalAddr().String())
	if ctx.websocket {
		ctx.wsCodec = new(ws.WsCodec)
	}
	if ctx.http {
		ctx.httpCodec = new(httpcodec.HttpCodec)
		ctx.closeDate = s.CachedNow() + 60
	} else {
		ctx.closeDate = s.CachedNow() + 30
	}
	s.timeoutWheel.Add(c.ConnId(), ctx.closeDate)
	c.SetContext(ctx)

	proto := "tcp"
	if ctx.websocket {
		proto = "websocket"
	} else if ctx.http {
		proto = "http"
	}
	metricConnOpen.Inc(proto)

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
		s.timeoutWheel.Remove(c.ConnId(), ctx.closeDate)
		if ctx.wsCodec != nil {
			ctx.wsCodec.Conn.Release()
			ctx.wsCodec = nil
		}

		proto := "tcp"
		if ctx.websocket {
			proto = "websocket"
		} else if ctx.http {
			proto = "http"
		}
		metricConnClose.Inc(proto)

		c.SetContext(nil)
	}()

	// HTTP connections are transient - no persistent session to close
	if ctx.http {
		return
	}

	if ctx.authKey == nil || ctx.authKey.PermAuthKeyId() == 0 {
		return
	}

	// kId, sessId, connId, clientIp := ctx.authKey.AuthKeyId(), ctx.sessionId, c.ConnId(), ctx.clientIp
	bDeleted := s.authSessionMgr.RemoveSession(ctx.authKey.AuthKeyId(), ctx.sessionId, c.ConnId())
	if !bDeleted {
		return
	}

	if err := s.pool.Submit(func() {
		closeCtx, span := otel.Tracer("gnetway").Start(context.Background(), "SessionCloseSession")
		defer span.End()
		err := s.svcCtx.Dao.SessionDispatcher.CloseSession(closeCtx, ctx.authKey.PermAuthKeyId(), &session.TLSessionCloseSession{
			Client: session.MakeTLSessionClientEvent(&session.SessionClientEvent{
				ServerId:      s.svcCtx.GatewayId,
				AuthKeyId:     ctx.authKey.AuthKeyId(),
				KeyType:       int32(ctx.authKey.AuthKeyType()),
				PermAuthKeyId: ctx.authKey.PermAuthKeyId(),
				SessionId:     ctx.sessionId,
				ClientIp:      ctx.clientIp,
			}).To_SessionClientEvent(),
		})
		if err != nil {
			logx.Errorf("client.SessionCloseSession - error: %v", err)
		}
	}); err != nil {
		logx.Errorf("OnClose - pool.Submit error: %v, authKeyId: %d", err, ctx.authKey.AuthKeyId())
	}

	return
}

// OnTraffic fires when a local socket receives data from the peer.
func (s *Server) OnTraffic(c gnet.Conn) (action gnet.Action) {
	ctx := c.Context().(*connContext)
	oldCloseDate := ctx.closeDate
	if ctx.http {
		ctx.closeDate = s.CachedNow() + 60 + rand.Int63()%10
	} else {
		ctx.closeDate = s.CachedNow() + 300 + rand.Int63()%10
	}
	s.timeoutWheel.Move(c.ConnId(), oldCloseDate, ctx.closeDate)
	if ctx.ppv1 {
		ppv1, err := c.Peek(-1)
		if err != nil {
			logx.Errorf("conn(%s) Peek fail: %v", c, err)
			return
		}
		if len(ppv1) < len(pp.V1Identifier) {
			logx.Errorf("conn(%s) ppv1 < len(pp.V1Identifier), data: %s", c, hex.EncodeToString(ppv1))
			return
		}

		if bytes.HasPrefix(ppv1, pp.V1Identifier) {
			logx.Errorf("conn(%s) ppv1 data: %s", c, hex.EncodeToString(ppv1))

			r := bytes.NewReader(ppv1)
			h, err := pp.ReadHeader(r)
			if err != nil {
				logx.Errorf("conn(%s) ReadHeader error: %v", c, err)
				if r.Len() > 107 {
					return gnet.Close
				} else {
					return
				}
			}
			if host, _, err := net.SplitHostPort(h.Source.String()); err == nil {
				ctx.setClientIp(host)
			} else {
				ctx.setClientIp(h.Source.String())
			}
			_, _ = c.Discard(len(ppv1) - r.Len())
			ctx.ppv1 = false

			if r.Len() == 0 {
				return
			}
		} else {
			ctx.ppv1 = false
		}
	}

	if ctx.http {
		return s.onHttpData(ctx, c)
	} else if ctx.websocket {
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
	s.cachedNow.Store(now)

	// Use time wheel to expire only connections in the current slot
	// instead of iterating all connections.
	expiredConnIds := s.timeoutWheel.ExpireSlot(now)
	for _, connId := range expiredConnIds {
		s.eng.Trigger(connId, func(c gnet.Conn) {
			ctx, _ := c.Context().(*connContext)
			if ctx == nil {
				return
			}
			if now >= ctx.closeDate {
				logx.Errorf("close conn(%s) by timeout", c)
				metricConnTimeout.Inc()
				_ = c.Close()
			} else {
				// Connection was refreshed but still in old slot; re-add to new slot
				s.timeoutWheel.Add(connId, ctx.closeDate)
			}
		})
	}
	return
}

// computeQuickAckToken computes the Quick ACK token per MTProto spec:
// first 32 bits of SHA256(authKey[88:88+32] + encrypted_data) | 0x80000000
func computeQuickAckToken(authKey []byte, encryptedData []byte) uint32 {
	h := sha256.New()
	h.Write(authKey[88 : 88+32])
	h.Write(encryptedData)
	var sum [32]byte
	h.Sum(sum[:0])
	return binary.LittleEndian.Uint32(sum[:4]) | 0x80000000
}

func (s *Server) onEncryptedMessage(c gnet.Conn, ctx *connContext, authKey *authKeyUtil, needAck bool, mmsg []byte) error {
	since := timex.Now()
	defer func() {
		metricMsgProcess.ObserveFloat(timex.Since(since).Seconds()*1000, "encrypted")
	}()

	mtpRwaData, err := authKey.AesIgeDecrypt(mmsg[8:8+16], mmsg[24:])
	if err != nil {
		logx.Errorf("conn(%s) decrypt data(%d) error: {%v}, payload: %s", c, len(mmsg)-24, err, hex.EncodeToString(mmsg))
		return err
	}

	// Send Quick ACK if requested by client.
	// The token MUST go through ctx.codec.EncodeQuickAck so that the CTR cipher
	// state is advanced correctly on obfuscated transports. Bypassing the codec
	// here would desynchronise the client's CTR counter and corrupt all subsequent
	// messages (causing decryptServerResponse failures on the Android client).
	if needAck {
		ackToken := computeQuickAckToken(authKey.AuthKey(), mmsg[24:])
		if ackData := ctx.codec.EncodeQuickAck(ackToken); len(ackData) > 0 {
			if ctx.websocket {
				_ = wsutil.WriteServerBinary(c, ackData)
			} else {
				_, _ = c.Write(ackData)
			}
		}
		metricQuickAck.Inc()
	}

	var (
		permAuthKeyId     = authKey.PermAuthKeyId()
		salt              = int64(binary.LittleEndian.Uint64(mtpRwaData))
		sessionId         = int64(binary.LittleEndian.Uint64(mtpRwaData[8:]))
		msgId             = int64(binary.LittleEndian.Uint64(mtpRwaData[16:24]))
		isBindTempAuthKey = false
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
				isBindTempAuthKey = true
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
				buf := func() []byte {
					x2 := mtproto.GetEncodeBuf()
					defer mtproto.PutEncodeBuf(x2)
					x2.Long(authKey.AuthKeyId())
					x2.Bytes(msgKey)
					x2.Bytes(mtpRawData)
					return append([]byte(nil), x2.GetBuf()...)
				}()
				_ = UnThreadSafeWrite(c, &mtproto.MTPRawMessage{Payload: buf})

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

	var connType int32
	if isNew && s.authSessionMgr.AddNewSession(authKey, sessionId, connId) {
		connType = 1 // signal session side to create session
	}

	s.asyncRunIfError(
		c.ConnId(),
		msgId,
		func() error {
			defer func() {
				logx.WithDuration(timex.Since(since)).Infof("onEncryptedMessage: %s", c)
			}()

			ctx2, span := otel.Tracer("gnetway").Start(context.Background(), "SessionSendDataToSession")
			defer span.End()
			err := s.svcCtx.Dao.SessionDispatcher.SendData(ctx2, permAuthKeyId, &session.TLSessionSendDataToSession{
				Data: &session.SessionClientData{
					ServerId:      s.svcCtx.GatewayId,
					ConnType:      connType,
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
			return err
		},
		func(c gnet.Conn, msgId int64, err error) {
			if isBindTempAuthKey && errors.Is(err, mtproto.ErrAuthKeyUnregistered) {
				payload := serializeToBuffer2(salt, sessionId, &mtproto.TLMessage2{
					MsgId: nextMessageId(false),
					Seqno: 0,
					Bytes: 0,
					Object: &mtproto.TLRpcResult{
						ReqMsgId: msgId,
						Result: mtproto.MakeTLRpcError(&mtproto.RpcError{
							ErrorCode:    400,
							ErrorMessage: "ENCRYPTED_MESSAGE_INVALID",
						}).To_RpcError(),
					}})

				msgKey, mtpRawData, _ := authKey.AesIgeEncrypt(payload)
				buf := func() []byte {
					x2 := mtproto.GetEncodeBuf()
					defer mtproto.PutEncodeBuf(x2)
					x2.Long(authKey.AuthKeyId())
					x2.Bytes(msgKey)
					x2.Bytes(mtpRawData)
					return append([]byte(nil), x2.GetBuf()...)
				}()
				_ = UnThreadSafeWrite(c, &mtproto.MTPRawMessage{Payload: buf})
			}
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
			logx.Errorf("conn(%s) onHandshake - error: %v", c, err)
			metricHandshake.Inc("error")
			action = gnet.Close
		} else if out != nil {
			metricHandshake.Inc("ok")
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
			logx.Errorf("conn(%s) getAuthKey - error: invalid key id %d ", c, authKeyId)
			action = gnet.Close
			return
		}

		if authKey != nil {
			err := s.onEncryptedMessage(c, ctx, authKey, needAck, msg2)
			if err != nil {
				logx.Errorf("conn(%s) onEncryptedMessage - error: %v ", c, err)
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
				authKeyId,
				needAck,
				msg2Clone,
				func(authKeyId2 int64, mmsg []byte) (interface{}, error) {
					queryCtx, span := otel.Tracer("gnetway").Start(context.Background(), "SessionQueryAuthKey")
					defer span.End()
					key3, err2 := s.svcCtx.Dao.SessionDispatcher.QueryAuthKey(queryCtx, authKeyId2, &session.TLSessionQueryAuthKey{
						AuthKeyId: authKeyId2,
					})
					if err2 != nil {
						logx.Errorf("conn(%s) sessionQueryAuthKey - error: %v", c, err2)
						return nil, err2
					} else {
						s.PutAuthKey(key3)
					}

					return newAuthKeyUtil(key3), nil
				},
				func(c2 gnet.Conn, needAck2 bool, mmsg []byte, in interface{}, err error) {
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

						logx.Errorf("conn(%s) sessionQueryAuthKey - error: %v", c2, err)
						_ = c2.Close()
					} else {
						authKey2 := in.(*authKeyUtil)
						ctx2 := c2.Context().(*connContext)
						ctx2.putAuthKey(authKey2)
						err = s.onEncryptedMessage(c2, ctx2, authKey2, needAck2, mmsg)
						if err != nil {
							logx.Errorf("conn(%s) onEncryptedMessage - error: %v ", c2, err)
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
		logx.Errorf("conn(%s) c.Write(data) - error: ctx.codec == nil ", c)
		return nil
	}

	if msg == nil {
		logx.Errorf("conn(%s) c.Write(data) - error: msg == ni ", c)
		return nil
	}

	data, err := ctx.codec.Encode(c, msg)
	if err != nil {
		logx.Errorf("conn(%s) ctx.codec.Encode(c, msg) - error: %v ", c, err)
		return err
	}

	if ctx.websocket {
		// This is the echo server
		err = wsutil.WriteServerBinary(c, data)
		if err != nil {
			logx.Errorf("conn[%v] [err=%v]", c.RemoteAddr().String(), err.Error())
			return err
		}
	} else {
		_, err = c.Write(data)
		if err != nil {
			logx.Errorf("conn(%s) c.Write(data) - error: %v ", c, err)
		}
	}

	return nil
}
