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
	"strconv"
	"time"

	"github.com/teamgram/marmota/pkg/cache"
	"github.com/teamgram/teamgram-server/app/interface/gnetway/internal/config"
	"github.com/teamgram/teamgram-server/app/interface/gnetway/internal/svc"

	"github.com/panjf2000/gnet/v2"
	"github.com/panjf2000/gnet/v2/pkg/logging"
	"github.com/panjf2000/gnet/v2/pkg/pool/goroutine"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	//etcdPrefix is a etcd globe key prefix
	endpoints string
)

type Server struct {
	gnet.BuiltinEventEngine
	eng            gnet.Engine
	pool           *goroutine.Pool
	cache          *cache.LRUCache
	c              *config.Config
	handshake      *handshake
	authSessionMgr *authSessionManager
	svcCtx         *svc.ServiceContext
	tickNumber     int64
}

func New(svcCtx *svc.ServiceContext, c config.Config) *Server {
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

	s.c = &c
	s.svcCtx = svcCtx

	go func() {
		s.Serve()
	}()
	return s
}

func (s *Server) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	logx.Debugf("stop engine... error: %v", s.eng.Stop(ctx))
}

func (s *Server) Serve() {
	// addrs := strings.Join(s.c.Gnetway.ToAddresses(), ",")
	logx.Debugf("addrs: %s", s.c.Gnetway.ToAddresses())

	err := gnet.Rotate(
		s,
		s.c.Gnetway.ToAddresses(),
		gnet.WithMulticore(s.c.Gnetway.Multicore),
		gnet.WithSocketRecvBuffer(s.c.Gnetway.ReceiveBuf),
		gnet.WithSocketSendBuffer(s.c.Gnetway.SendBuf),
		gnet.WithLockOSThread(true),
		gnet.WithReuseAddr(true),
		gnet.WithTicker(true),
		gnet.WithLogLevel(logging.DebugLevel),
		gnet.WithLogger(NewLogger()))
	if err != nil {
		logx.Error(err)
		panic(err)
	}
}
