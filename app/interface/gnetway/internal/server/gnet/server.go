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
	"errors"
	"sync/atomic"
	"time"

	"github.com/teamgram/marmota/pkg/cache"
	"github.com/teamgram/teamgram-server/app/interface/gnetway/internal/config"
	"github.com/teamgram/teamgram-server/app/interface/gnetway/internal/server/gnet/codec"
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
	cachedNow      atomic.Int64 // cached unix timestamp, updated every OnTick
	timeoutWheel   *timeoutWheel
}

// CachedNow returns the cached unix timestamp (updated every OnTick interval).
// Falls back to time.Now() if not yet initialized.
func (s *Server) CachedNow() int64 {
	if n := s.cachedNow.Load(); n > 0 {
		return n
	}
	return time.Now().Unix()
}

func New(svcCtx *svc.ServiceContext, c config.Config) *Server {
	var (
		s = new(Server)
	)

	s.authSessionMgr = NewAuthSessionManager()
	s.timeoutWheel = newTimeoutWheel()

	s.handshake = mustNewHandshake(c.RSAKey)

	cacheSizeMB := c.Gnetway.AuthKeyCacheM
	if cacheSizeMB <= 0 {
		cacheSizeMB = 10
	}
	s.cache = cache.NewLRUCache(int64(cacheSizeMB) * 1024 * 1024)

	s.pool = goroutine.Default()

	s.c = &c
	s.svcCtx = svcCtx

	go func() {
		s.Serve()
	}()
	return s
}

// classifyCodecError maps codec errors to a stable reason label for metrics.
func classifyCodecError(err error) string {
	switch {
	case errors.Is(err, codec.ErrProtoBadMagic):
		return "bad_magic"
	case errors.Is(err, codec.ErrProtoBadLength):
		return "bad_len"
	case errors.Is(err, codec.ErrProtoBadCRC):
		return "bad_crc"
	case errors.Is(err, codec.ErrProtoBadSeq):
		return "bad_seq"
	case errors.Is(err, codec.ErrProtoDecrypt):
		return "decrypt"
	case errors.Is(err, codec.ErrTransportNotSupported):
		return "transport_unsupported"
	case errors.Is(err, codec.ErrUnexpectedEOF):
		return "unexpected_eof"
	default:
		return "other"
	}
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
