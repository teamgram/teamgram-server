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
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/teamgram/marmota/pkg/cache"
	"github.com/teamgram/teamgram-server/v2/app/interface/gnetway/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/interface/gnetway/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type Server struct {
	svcCtx         *svc.ServiceContext
	c              *config.Config
	handshake      *handshake
	authSessionMgr *authSessionManager
	cache          *cache.LRUCache

	listeners     []net.Listener
	connMgr       *connectionManager
	shutdownCh    chan struct{}
	wg            sync.WaitGroup
	connIdCounter int64
	tickNumber    int64
}

func New(svcCtx *svc.ServiceContext, c config.Config) *Server {
	s := &Server{
		svcCtx:         svcCtx,
		c:              &c,
		authSessionMgr: NewAuthSessionManager(),
		connMgr:        newConnectionManager(),
		shutdownCh:     make(chan struct{}),
		listeners:      make([]net.Listener, 0),
	}

	s.handshake = mustNewHandshake(c.RSAKey)
	s.cache = cache.NewLRUCache(10 * 1024 * 1024) // cache capacity: 10MB

	go func() {
		s.Serve()
	}()

	return s
}

func (s *Server) Close() {
	close(s.shutdownCh)

	// Close all listeners
	for _, ln := range s.listeners {
		_ = ln.Close()
	}

	// Close all connections
	s.connMgr.closeAll()

	// Wait for all goroutines to finish
	s.wg.Wait()

	logx.Infof("netserver shutdown completed")
}

func (s *Server) Serve() {
	addrs := s.c.Gnetway.ToAddresses()
	logx.Infof("starting netserver on addresses: %v", addrs)

	for _, addr := range addrs {
		protocol, address := parseAddress(addr)

		ln, err := net.Listen("tcp", address)
		if err != nil {
			logx.Errorf("failed to listen on %s: %v", address, err)
			panic(err)
		}

		s.listeners = append(s.listeners, ln)
		logx.Infof("netserver listening on %s://%s", protocol, address)

		s.wg.Add(1)
		go s.acceptLoop(ln, protocol, address)
	}

	// Start connection timeout checker
	s.wg.Add(1)
	go s.connectionTimeoutChecker()
}

func (s *Server) acceptLoop(ln net.Listener, protocol, address string) {
	defer s.wg.Done()

	for {
		conn, err := ln.Accept()
		if err != nil {
			select {
			case <-s.shutdownCh:
				return
			default:
				logx.Errorf("accept error: %v", err)
				continue
			}
		}

		connId := atomic.AddInt64(&s.connIdCounter, 1)

		isTcp := s.c.Gnetway.IsTcp(address)
		isWebsocket := s.c.Gnetway.IsWebsocket(address)

		c := s.connMgr.newConnection(connId, conn, isTcp, isWebsocket)

		s.wg.Add(1)
		if isWebsocket {
			go s.handleWebSocketConnection(c)
		} else {
			go s.handleTCPConnection(c)
		}
	}
}

func (s *Server) connectionTimeoutChecker() {
	defer s.wg.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-s.shutdownCh:
			return
		case <-ticker.C:
			s.tickNumber++
			if s.tickNumber%5 == 0 {
				count := s.connMgr.count()
				logx.Statf("conn count: %d", count)
			}

			now := time.Now().Unix()
			s.connMgr.iterate(func(c *connection) {
				if now >= c.closeDate {
					logx.Debugf("closing conn(%d) by timeout", c.id)
					c.Close()
				}
			})
		}
	}
}

func (s *Server) GetConnCounts() int {
	return s.connMgr.count()
}

// parseAddress parses address like "tcp://0.0.0.0:10443" or "ws://0.0.0.0:8080"
func parseAddress(addr string) (protocol string, address string) {
	// Default to tcp
	protocol = "tcp"
	address = addr

	// Parse protocol prefix if present
	if len(addr) > 6 && addr[3:6] == "://" {
		protocol = addr[:3]
		address = addr[6:]
	} else if len(addr) > 7 && addr[4:7] == "://" {
		protocol = addr[:4]
		address = addr[7:]
	} else if len(addr) > 8 && addr[5:8] == "://" {
		protocol = addr[:5]
		address = addr[8:]
	}

	return
}
