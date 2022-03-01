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
	"strconv"

	"github.com/teamgram/marmota/pkg/cache"
	"github.com/teamgram/teamgram-server/app/interface/gateway/internal/config"
	"github.com/teamgram/teamgram-server/app/interface/gateway/internal/server/codec"

	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pkg/pool/goroutine"
)

var (
	//etcdPrefix is a etcd globe key prefix
	endpoints string
)

type Server struct {
	*gnet.EventServer
	svr            gnet.Server
	pool           *goroutine.Pool
	cache          *cache.LRUCache
	c              *config.Config
	handshake      *handshake
	session        *Session
	authSessionMgr *authSessionManager
}

func New(c config.Config) *Server {
	var (
		err error
		s   = new(Server)
	)

	s.authSessionMgr = NewAuthSessionManager()

	keyFingerprint, err := strconv.ParseUint(c.KeyFingerprint, 10, 64)
	if err != nil {
		panic(err)
	}
	s.handshake, err = newHandshake(c.KeyFile, keyFingerprint)
	if err != nil {
		panic(err)
	}

	s.cache = cache.NewLRUCache(10 * 1024 * 1024) // cache capacity: 10MB
	s.pool = goroutine.Default()

	s.session = NewSession(c)
	s.c = &c

	return s
}

func (s *Server) Close() {
	gnet.StopServer(s.c.Server.Addrs)
}

// Ping ping the resource.
func (s *Server) Ping(ctx context.Context) (err error) {
	return nil
}

func (s *Server) Serve() error {
	return gnet.Serve(s,
		s.c.Server.Addrs,
		gnet.WithMulticore(s.c.Server.Multicore),
		gnet.WithSocketRecvBuffer(s.c.Server.ReceiveBuf),
		gnet.WithSocketSendBuffer(s.c.Server.SendBuf),
		gnet.WithCodec(codec.NewMTProtoCodec()),
		gnet.WithLockOSThread(true),
		gnet.WithReusePort(true))
}
