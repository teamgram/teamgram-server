// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package dao

import (
	"errors"
	"strings"
	"sync"

	"github.com/teamgram/teamgram-server/app/interface/gnetway/internal/config"
	sessionclient "github.com/teamgram/teamgram-server/app/interface/session/client"

	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/hash"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stringx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrSessionNotFound = errors.New("not found session")
)

type ShardingSessionClient struct {
	mu         sync.RWMutex
	gatewayId  string
	dispatcher *hash.ConsistentHash
	sessions   map[string]sessionclient.SessionClient
}

func NewShardingSessionClient(c config.Config) *ShardingSessionClient {
	sess := &ShardingSessionClient{
		dispatcher: hash.NewConsistentHash(),
		sessions:   make(map[string]sessionclient.SessionClient),
	}
	sess.watch(c.Session)

	return sess
}

func (sess *ShardingSessionClient) watch(c zrpc.RpcClientConf) {
	sub, _ := discov.NewSubscriber(c.Etcd.Hosts, c.Etcd.Key)
	update := func() {
		var (
			addClis    []string
			removeClis []string
		)

		sess.mu.Lock()
		defer sess.mu.Unlock()

		values := sub.Values()
		sessions := map[string]sessionclient.SessionClient{}
		for _, v := range values {
			if old, ok := sess.sessions[v]; ok {
				sessions[v] = old
				continue
			}
			c.Endpoints = []string{v}
			cli, err := zrpc.NewClient(c)
			if err != nil {
				logx.Error("watchComet NewClient(%+v) error(%v)", values, err)
				return
			}
			sessionCli := sessionclient.NewSessionClient(cli)
			sessions[v] = sessionCli

			addClis = append(addClis, v)
		}

		for key, _ := range sess.sessions {
			if !stringx.Contains(values, key) {
				removeClis = append(removeClis, key)
			}
		}

		for _, n := range addClis {
			sess.dispatcher.Add(n)
		}

		for _, n := range removeClis {
			sess.dispatcher.Remove(n)
		}

		sess.sessions = sessions
	}

	sub.AddListener(update)
	update()
}

func (sess *ShardingSessionClient) InvokeByKey(key string, cb func(client sessionclient.SessionClient) (err error)) error {
	if cb == nil {
		return nil
	}

	sess.mu.RLock()
	val, ok := sess.dispatcher.Get(key)
	if !ok {
		sess.mu.RUnlock()
		return ErrSessionNotFound
	}

	node := val.(string)
	cli, ok := sess.sessions[node]
	sess.mu.RUnlock()

	if !ok {
		return ErrSessionNotFound
	}

	err := cb(cli)
	if err == nil {
		return nil
	}

	// redirect 错误：session 告知正确节点，直接重路由
	if target, ok := parseRedirectError(err); ok {
		logx.Infof("session node %s redirected to %s for key %s", node, target, key)

		sess.mu.RLock()
		redirectCli, ok := sess.sessions[target]
		sess.mu.RUnlock()

		if ok {
			return cb(redirectCli)
		}
		// 目标节点不在已知列表中，按连接错误处理走 fallback
	}

	// 非连接错误直接返回
	if !isConnError(err) {
		return err
	}

	// 连接错误：从哈希环移除故障节点，重试到新节点
	logx.Errorf("session node %s unreachable, removing from ring and retrying", node)

	sess.mu.Lock()
	sess.dispatcher.Remove(node)
	delete(sess.sessions, node)
	sess.mu.Unlock()

	sess.mu.RLock()
	newVal, ok := sess.dispatcher.Get(key)
	if !ok {
		sess.mu.RUnlock()
		return ErrSessionNotFound
	}
	newNode := newVal.(string)
	newCli, ok := sess.sessions[newNode]
	sess.mu.RUnlock()

	if !ok {
		return ErrSessionNotFound
	}

	return cb(newCli)
}

// parseRedirectError 解析 session 返回的 redirect 错误，提取目标节点地址
// 错误格式：gRPC code 700, message "REDIRECT_TO_{server_addr}"
func parseRedirectError(err error) (string, bool) {
	s, ok := status.FromError(err)
	if !ok {
		return "", false
	}
	if s.Code() != codes.Code(700) {
		return "", false
	}
	msg := s.Message()
	const prefix = "REDIRECT_TO_"
	if strings.HasPrefix(msg, prefix) {
		target := msg[len(prefix):]
		if target != "" {
			return target, true
		}
	}
	return "", false
}

// isConnError 判断是否为连接级错误（节点不可达）
func isConnError(err error) bool {
	s, ok := status.FromError(err)
	if !ok {
		return false
	}
	switch s.Code() {
	case codes.Unavailable, codes.DeadlineExceeded:
		return true
	}
	return false
}
