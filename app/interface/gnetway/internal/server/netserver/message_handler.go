// Copyright (c) 2024 The Teamgooo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package netserver

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"

	sessionclient "github.com/teamgram/teamgram-server/v2/app/interface/session/client"
	"github.com/teamgram/teamgram-server/v2/app/interface/session/session"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/mt"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/zeromicro/go-zero/core/logx"
)

func (s *Server) onClose(c *connection, err error) {
	logx.Debugf("onConnClosed - conn(%d), err: %v", c.id, err)

	if c.authKey == nil || c.authKey.PermAuthKeyId() == 0 {
		return
	}

	bDeleted := s.authSessionMgr.RemoveSession(c.authKey.AuthKeyId(), c.sessionId, c.id)
	if !bDeleted {
		return
	}

	// Synchronous call to close session
	_ = s.svcCtx.ShardingSessionClient.InvokeByKey(
		strconv.FormatInt(c.authKey.PermAuthKeyId(), 10),
		func(client sessionclient.SessionClient) (err error) {
			_, err = client.SessionCloseSession(context.Background(), &session.TLSessionCloseSession{
				Client: session.MakeTLSessionClientEvent(&session.TLSessionClientEvent{
					ServerId:      s.svcCtx.GatewayId,
					AuthKeyId:     c.authKey.AuthKeyId(),
					KeyType:       int32(c.authKey.AuthKeyType()),
					PermAuthKeyId: c.authKey.PermAuthKeyId(),
					SessionId:     c.sessionId,
					ClientIp:      c.clientIp,
				}),
			})
			return
		})
}

func (s *Server) onEncryptedMessage(c *connection, authKey *authKeyUtil, needAck bool, mmsg []byte) error {
	mtpRwaData, err := authKey.AesIgeDecrypt(mmsg[8:8+16], mmsg[24:])
	if err != nil {
		logx.Errorf("conn(%d) decrypt data(%d) error: {%v}, payload: %s", c.id, len(mmsg)-24, err, hex.EncodeToString(mmsg))
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

		logx.Debugf("mtpRwaData: %s", hex.EncodeToString(mtpRwaData))
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
					Object: mt.MakeTLPong(&mt.TLPong{
						MsgId:  int64(binary.LittleEndian.Uint64(mtpRwaData[16:])),
						PingId: unknownMsg.PingId,
					}),
				})

				msgKey, mtpRawData, _ := authKey.AesIgeEncrypt(payload)
				x3.PutInt64(authKey.AuthKeyId())
				x3.Put(msgKey)
				x3.Put(mtpRawData)
				_ = s.writeToConnection(c, x3.Bytes())
				x3.Reset()

				return nil
			default:
				if permAuthKeyId == 0 {
					logx.Errorf("recv unknown msg: %v, ignore it", unknownMsg)
					return fmt.Errorf("unknown msg")
				}
			}
		}
	}

	var (
		isNew    = c.sessionId != sessionId
		clientIp = c.clientIp
		connId   = c.id
	)
	if isNew {
		c.sessionId = sessionId
	}

	// Synchronous call to create session and send data
	_ = s.svcCtx.Dao.ShardingSessionClient.InvokeByKey(
		strconv.FormatInt(permAuthKeyId, 10),
		func(client sessionclient.SessionClient) (err error) {
			if isNew {
				if s.authSessionMgr.AddNewSession(authKey, sessionId, connId) {
					r := &session.TLSessionCreateSession{
						Client: session.MakeTLSessionClientEvent(&session.TLSessionClientEvent{
							ServerId:      s.svcCtx.GatewayId,
							AuthKeyId:     authKey.AuthKeyId(),
							KeyType:       int32(authKey.AuthKeyType()),
							PermAuthKeyId: permAuthKeyId,
							SessionId:     sessionId,
							ClientIp:      clientIp,
						}),
					}
					logx.Infof("conn(%d) create new session(%s)", connId, r)
					_, err = client.SessionCreateSession(context.Background(), r)
					if err != nil {
						logx.Errorf("session.createSession - error: %v", err)
					}
				}
			}

			_, err = client.SessionSendDataToSession(context.Background(), &session.TLSessionSendDataToSession{
				Data: session.MakeTLSessionClientData(&session.TLSessionClientData{
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

	return nil
}

func (s *Server) onMTPRawMessage(c *connection, authKeyId int64, needAck bool, msg2 []byte) (shouldClose bool) {
	if authKeyId == 0 {
		if len(msg2) < 8+8+4 {
			logx.Errorf("conn(%d) error: invalid data len < 8+8+4", c.id)
			return true
		}

		d := bin.NewDecoder(msg2[8:])
		msgId, _ := d.Int64()
		_ = msgId

		dataLen, _ := d.Int32()
		if len(msg2) < 8+8+4+int(dataLen) {
			logx.Errorf("conn(%d) error: invalid data len < 8+8+4+int(dataLen)", c.id)
			return true
		}

		err := s.onHandshake(c, d)
		if err != nil {
			logx.Errorf("conn(%d) error: %v", c.id, err)
			return true
		}
	} else {
		authKey := c.getAuthKey()
		if authKey == nil {
			key := s.GetAuthKey(authKeyId)
			if key != nil {
				authKey = newAuthKeyUtil(key)
				c.putAuthKey(authKey)
			}
		} else if authKey.AuthKeyId() != authKeyId {
			return true
		}

		if authKey != nil {
			err := s.onEncryptedMessage(c, authKey, needAck, msg2)
			if err != nil {
				return true
			}
			return false
		}

		if len(msg2) > 32 {
			logx.Debugf("conn(%d) data: %s", c.id, hex.EncodeToString(msg2[:32]))
		}

		// Synchronous call to query auth key
		var key3 *tg.AuthKeyInfo
		err2 := s.svcCtx.Dao.ShardingSessionClient.InvokeByKey(
			strconv.FormatInt(authKeyId, 10),
			func(client sessionclient.SessionClient) (err error) {
				key3, err = client.SessionQueryAuthKey(context.Background(), &session.TLSessionQueryAuthKey{
					AuthKeyId: authKeyId,
				})
				return
			})

		if err2 != nil {
			logx.Errorf("conn(%d) sessionQueryAuthKey error: %v", c.id, err2)
			if errors.Is(err2, tg.ErrAuthKeyUnregistered) {
				out2 := make([]byte, 4)
				var code int32 = -404
				binary.LittleEndian.PutUint32(out2, uint32(code))
				_ = s.writeToConnection(c, out2)
			}
			return true // shouldClose
		}

		key4 := key3
		s.PutAuthKey(key4)
		authKey2 := newAuthKeyUtil(key4)

		c.putAuthKey(authKey2)
		err := s.onEncryptedMessage(c, authKey2, needAck, msg2)
		if err != nil {
			return true // shouldClose
		}
	}

	return false
}

func (s *Server) writeToConnection(c *connection, msg []byte) error {
	if c.codec == nil {
		return nil
	}

	if msg == nil {
		return nil
	}

	data, err := c.codec.Encode(&tcpConnAdapter{c}, msg)
	if err != nil {
		return err
	}

	if c.websocket {
		if c.gwsConn != nil {
			// gws requires OpcodeBinary and the data
			return c.gwsConn.WriteMessage(1, data) // 1 = OpcodeBinary
		}
		return fmt.Errorf("websocket connection not initialized")
	}

	_, err = c.Write(data)
	return err
}
