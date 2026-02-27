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

package dao

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	sessionclient "github.com/teamgram/teamgram-server/app/interface/session/client"
	"github.com/teamgram/teamgram-server/app/interface/session/session"

	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc/codes"
	grpcStatus "google.golang.org/grpc/status"
)

type sessionDataCtx struct {
	ctx     context.Context
	updates any
}

// SessionOptions comet options.
type SessionOptions struct {
	RoutineSize uint64
	RoutineChan uint64
}

// Session is a gateway.
type Session struct {
	serverId       string
	client         sessionclient.SessionClient
	sessionChan    []chan sessionDataCtx
	sessionChanNum uint64
	options        SessionOptions
	ctx            context.Context
	cancel         context.CancelFunc
	unavailable    atomic.Bool // set when connection errors detected
}

func isSessionConnError(err error) bool {
	s, ok := grpcStatus.FromError(err)
	if !ok {
		return false
	}
	switch s.Code() {
	case codes.Unavailable, codes.DeadlineExceeded:
		return true
	}
	return false
}

// process
func (c *Session) process(sessionChan chan sessionDataCtx) {
	var err error
	for {
		select {
		case sessionData, ok := <-sessionChan:
			if !ok {
				logx.Errorf("process error")
				return
			}

			if c.unavailable.Load() {
				logx.Errorf("session(%s) unavailable, dropping push", c.serverId)
				continue
			}

			switch r := sessionData.updates.(type) {
			case *session.TLSessionPushSessionUpdatesData:
				_, err = c.client.SessionPushSessionUpdatesData(sessionData.ctx, r)
				if err != nil {
					logx.Errorf("c.client.PushSessionUpdates(%s, %v, reply) serverId:%s error(%v)", r, c.serverId, c.serverId, err)
					if isSessionConnError(err) {
						c.unavailable.Store(true)
						logx.Errorf("session(%s) marked unavailable due to conn error", c.serverId)
					}
				}
			case *session.TLSessionPushUpdatesData:
				_, err = c.client.SessionPushUpdatesData(sessionData.ctx, r)
				if err != nil {
					logx.Errorf("c.client.PushUpdates(%s, %v, reply) serverId:%s error(%v)", r, c.serverId, c.serverId, err)
					if isSessionConnError(err) {
						c.unavailable.Store(true)
						logx.Errorf("session(%s) marked unavailable due to conn error", c.serverId)
					}
				}
			case *session.TLSessionPushRpcResultData:
				_, err = c.client.SessionPushRpcResultData(sessionData.ctx, r)
				if err != nil {
					logx.Errorf("c.client.PushRpcResult(%s, %v, reply) serverId:%s error(%v)", r, c.serverId, c.serverId, err)
					if isSessionConnError(err) {
						c.unavailable.Store(true)
						logx.Errorf("session(%s) marked unavailable due to conn error", c.serverId)
					}
				}
			default:
				logx.Errorf("invalid type: %#v", r)
			}
		case <-c.ctx.Done():
			return
		}
	}
}

func (c *Session) Close() (err error) {
	finish := make(chan bool)
	go func() {
		for {
			n := len(c.sessionChan)
			for _, ch := range c.sessionChan {
				n += len(ch)
			}
			if n == 0 {
				finish <- true
				return
			}
			time.Sleep(time.Second)
		}
	}()
	select {
	case <-finish:
		logx.Info("close session client finish")
	case <-time.After(5 * time.Second):
		err = fmt.Errorf("close session(server:%s push:%d) timeout", c.serverId, len(c.sessionChan))
	}
	c.cancel()
	return
}

func (c *Session) PushUpdates(ctx context.Context, msg *session.TLSessionPushUpdatesData) (err error) {
	if c.unavailable.Load() {
		return fmt.Errorf("session(%s) unavailable", c.serverId)
	}
	idx := atomic.AddUint64(&c.sessionChanNum, 1) % c.options.RoutineSize
	c.sessionChan[idx] <- sessionDataCtx{ctx: ctx, updates: msg}
	return
}

func (c *Session) PushSessionUpdates(ctx context.Context, msg *session.TLSessionPushSessionUpdatesData) (err error) {
	if c.unavailable.Load() {
		return fmt.Errorf("session(%s) unavailable", c.serverId)
	}
	idx := atomic.AddUint64(&c.sessionChanNum, 1) % c.options.RoutineSize
	c.sessionChan[idx] <- sessionDataCtx{ctx: ctx, updates: msg}
	return
}

func (c *Session) PushRpcResult(ctx context.Context, msg *session.TLSessionPushRpcResultData) (err error) {
	if c.unavailable.Load() {
		return fmt.Errorf("session(%s) unavailable", c.serverId)
	}
	idx := atomic.AddUint64(&c.sessionChanNum, 1) % c.options.RoutineSize
	c.sessionChan[idx] <- sessionDataCtx{ctx: ctx, updates: msg}
	return
}

// NewSession new a comet.
func NewSession(c zrpc.RpcClientConf, options SessionOptions) (*Session, error) {
	sess := &Session{
		serverId:    c.Endpoints[0],
		sessionChan: make([]chan sessionDataCtx, options.RoutineSize),
		options:     options,
	}

	cli, err := zrpc.NewClient(c)
	if err != nil {
		logx.Errorf("watchComet NewClient(%+v) error(%v)", c, err)
		return nil, err
	}
	sess.client = sessionclient.NewSessionClient(cli)
	sess.ctx, sess.cancel = context.WithCancel(context.Background())

	for i := uint64(0); i < options.RoutineSize; i++ {
		sess.sessionChan[i] = make(chan sessionDataCtx, options.RoutineChan)
		go sess.process(sess.sessionChan[i])
	}
	return sess, nil
}

func (d *Dao) watch(c zrpc.RpcClientConf) {
	sub, _ := discov.NewSubscriber(c.Etcd.Hosts, c.Etcd.Key)
	update := func() {
		values := sub.Values()
		if len(values) == 0 {
			return
		}

		sessions := map[string]*Session{}
		for _, v := range values {
			if old, ok := d.sessionServers[v]; ok {
				sessions[v] = old
				continue
			}
			c.Endpoints = []string{v}
			// cli, err := zrpc.NewClient(c)
			cli, err := NewSession(c, SessionOptions{
				RoutineSize: d.conf.Routine.Size,
				RoutineChan: d.conf.Routine.Chan,
			})
			if err != nil {
				logx.Error("watchComet NewClient(%+v) error(%v)", values, err)
				return
			}
			sessions[v] = cli
		}

		for key, old := range d.sessionServers {
			if _, ok := sessions[key]; !ok {
				old.cancel()
				logx.Infof("watchComet DelComet:%s", key)
			}
		}

		d.sessionServers = sessions
	}

	sub.AddListener(update)
	update()
}

func (d *Dao) PushUpdatesToSession(ctx context.Context, serverId string, msg *session.TLSessionPushUpdatesData) (err error) {
	if c, ok := d.sessionServers[serverId]; ok {
		return c.PushUpdates(ctx, msg)
	} else {
		logx.WithContext(ctx).Errorf("PushUpdatesToSession - stale gateway, serverId %s not in active sessions (permAuthKeyId:%d)",
			serverId, msg.PermAuthKeyId)
		return fmt.Errorf("stale gateway %s", serverId)
	}
}

func (d *Dao) PushSessionUpdatesToSession(ctx context.Context, serverId string, msg *session.TLSessionPushSessionUpdatesData) (err error) {
	if c, ok := d.sessionServers[serverId]; ok {
		return c.PushSessionUpdates(ctx, msg)
	} else {
		logx.WithContext(ctx).Errorf("PushSessionUpdatesToSession - stale gateway, serverId %s not in active sessions (permAuthKeyId:%d)",
			serverId, msg.PermAuthKeyId)
		return fmt.Errorf("stale gateway %s", serverId)
	}
}

func (d *Dao) PushRpcResultToSession(ctx context.Context, serverId string, msg *session.TLSessionPushRpcResultData) (err error) {
	if c, ok := d.sessionServers[serverId]; ok {
		return c.PushRpcResult(ctx, msg)
	} else {
		logx.WithContext(ctx).Errorf("PushRpcResultToSession - stale gateway, serverId %s not in active sessions (permAuthKeyId:%d)",
			serverId, msg.PermAuthKeyId)
		return fmt.Errorf("stale gateway %s", serverId)
	}
}
