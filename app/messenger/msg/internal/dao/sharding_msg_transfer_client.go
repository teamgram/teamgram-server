// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package dao

import (
	"errors"

	msgtransferclient "github.com/teamgram/teamgram-server/app/messenger/msg/msgtransfer/client"

	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/hash"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stringx"
	"github.com/zeromicro/go-zero/zrpc"
)

var (
	ErrMsgClientNotFound = errors.New("not found clients")
)

type ShardingMsgTransferClient struct {
	gatewayId  string
	dispatcher *hash.ConsistentHash
	msgClients map[string]msgtransferclient.MsgtransferClient
}

func NewShardingMsgTransferClient(c zrpc.RpcClientConf) *ShardingMsgTransferClient {
	sess := &ShardingMsgTransferClient{
		dispatcher: hash.NewConsistentHash(),
		msgClients: make(map[string]msgtransferclient.MsgtransferClient),
	}
	sess.watch(c)

	return sess
}

func (c *ShardingMsgTransferClient) watch(config zrpc.RpcClientConf) {
	sub, _ := discov.NewSubscriber(config.Etcd.Hosts, config.Etcd.Key)
	update := func() {
		var (
			addClis    []string
			removeClis []string
		)

		values := sub.Values()
		clients := map[string]msgtransferclient.MsgtransferClient{}
		for _, v := range values {
			if old, ok := c.msgClients[v]; ok {
				clients[v] = old
				continue
			}
			config.Endpoints = []string{v}
			cli, err := zrpc.NewClient(config)
			if err != nil {
				logx.Error("watchComet NewClient(%+v) error(%v)", values, err)
				return
			}
			sessionCli := msgtransferclient.NewMsgtransferClient(cli)
			clients[v] = sessionCli

			addClis = append(addClis, v)
		}

		for key, _ := range c.msgClients {
			if !stringx.Contains(values, key) {
				removeClis = append(removeClis, key)
			}
		}

		for _, n := range addClis {
			c.dispatcher.Add(n)
		}

		for _, n := range removeClis {
			c.dispatcher.Remove(n)
		}

		c.msgClients = clients
	}

	sub.AddListener(update)
	update()
}

func (c *ShardingMsgTransferClient) InvokeByKey(key string, cb func(client msgtransferclient.MsgtransferClient) (err error)) error {
	val, ok := c.dispatcher.Get(key)
	if !ok {
		return ErrMsgClientNotFound
	}

	cli, ok := c.msgClients[val.(string)]
	if !ok {
		return ErrMsgClientNotFound
	}

	if cb == nil {
		return nil
	}

	return cb(cli)
}
