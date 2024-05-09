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
	"context"
	"fmt"
	"strconv"

	"github.com/teamgram/marmota/pkg/cache"
	"github.com/teamgram/marmota/pkg/net2"
	"github.com/teamgram/marmota/pkg/timer2"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/interface/gateway/internal/config"
	sessionpb "github.com/teamgram/teamgram-server/app/interface/session/session"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	//etcdPrefix is a etcd globe key prefix
	endpoints string
)

type Server struct {
	c              *config.Config
	server         *net2.TcpServer2
	cache          *cache.LRUCache
	handshake      *handshake
	session        *Session
	authSessionMgr *authSessionManager
	timer          *timer2.Timer // 32 * 2048
}

func New(c config.Config) *Server {
	var (
		err error
		s   = new(Server)
	)

	s.timer = timer2.NewTimer(1024)

	s.authSessionMgr = NewAuthSessionManager()

	keyFingerprint, err := strconv.ParseUint(c.KeyFingerprint, 10, 64)
	if err != nil {
		panic(err)
	}
	s.handshake, err = newHandshake(c.KeyFile, keyFingerprint,
		func(ctx context.Context, key *mtproto.AuthKeyInfo, salt *mtproto.FutureSalt, expiresIn int32) error {
			sessClient, err2 := s.session.getSessionClient(strconv.FormatInt(key.AuthKeyId, 10))
			if err2 != nil {
				logx.Errorf("getSessionClient error: %v, {authKeyId: %d}", err, key.AuthKeyId)
				return err2
			}

			// Fix by @wuyun9527, 2018-12-21
			var (
				rB *mtproto.Bool
			)
			rB, err2 = sessClient.SessionSetAuthKey(context.Background(), &sessionpb.TLSessionSetAuthKey{
				AuthKey:    key,
				FutureSalt: salt,
				ExpiresIn:  expiresIn,
			})
			if err2 != nil {
				logx.Errorf("saveAuthKeyInfo not successful - auth_key_id:%d, err:%v", key.AuthKeyId, err2)
				return err2
			} else if !mtproto.FromBool(rB) {
				logx.Errorf("saveAuthKeyInfo not successful - auth_key_id:%d", key.AuthKeyId)
				err2 = fmt.Errorf("saveAuthKeyInfo error")
				return err2
			} else {
				s.PutAuthKey(&mtproto.AuthKeyInfo{
					AuthKeyId:          key.AuthKeyId,
					AuthKey:            key.AuthKey,
					AuthKeyType:        key.AuthKeyType,
					PermAuthKeyId:      key.PermAuthKeyId,
					TempAuthKeyId:      key.TempAuthKeyId,
					MediaTempAuthKeyId: key.MediaTempAuthKeyId})
			}
			return nil
		})
	// s.handshake, err = newHandshake(c.KeyFile, keyFingerprint)
	if err != nil {
		panic(err)
	}

	s.cache = cache.NewLRUCache(10 * 1024 * 1024) // cache capacity: 10MB
	// s.pool = goroutine.Default()

	s.session = NewSession(c)
	s.c = &c

	return s
}

func (s *Server) Close() {
	s.server.Stop()
}

// Ping ping the resource.
func (s *Server) Ping(ctx context.Context) (err error) {
	return nil
}

func (s *Server) Serve() error {
	serv, err := net2.NewTcpServer2(s.c.Server, s.c.MaxProc, s)
	if err != nil {
		panic(err)
	}
	s.server = serv
	s.server.Serve()

	return nil
}
