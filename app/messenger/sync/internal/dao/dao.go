/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dao

import (
	kafka "github.com/teamgram/marmota/pkg/mq"
	"github.com/teamgram/marmota/pkg/net/rpcx"
	"github.com/teamgram/marmota/pkg/stores/sqlc"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	sync_client "github.com/teamgram/teamgram-server/app/messenger/sync/client"
	"github.com/teamgram/teamgram-server/app/messenger/sync/internal/config"
	chat_client "github.com/teamgram/teamgram-server/app/service/biz/chat/client"
	idgen_client "github.com/teamgram/teamgram-server/app/service/idgen/client"
	status_client "github.com/teamgram/teamgram-server/app/service/status/client"
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/zrpc"
)

type Dao struct {
	*Mysql
	sqlc.CachedConn
	kv             kv.Store
	conf           *config.Config
	sessionServers map[string]*Session
	idgen_client.IDGenClient2
	status_client.StatusClient
	// channel_client.ChannelClient
	chat_client.ChatClient
	BotsClient sync_client.SyncClient
	PushClient sync_client.SyncClient
}

func New(c config.Config) *Dao {
	db := sqlx.NewMySQL(&c.Mysql)
	d := &Dao{
		Mysql:          newMysqlDao(db),
		CachedConn:     sqlc.NewConn(db, c.Cache),
		kv:             kv.NewStore(c.KV),
		conf:           &c,
		sessionServers: make(map[string]*Session),
		IDGenClient2:   idgen_client.NewIDGenClient2(zrpc.MustNewClient(c.IdgenClient)),
		StatusClient:   status_client.NewStatusClient(zrpc.MustNewClient(c.StatusClient)),
		// ChannelClient:  channel_client.NewChannelClient(rpcx.GetCachedRpcClient(c.ChannelClient)),
		ChatClient: chat_client.NewChatClient(rpcx.GetCachedRpcClient(c.ChatClient)),
		BotsClient: sync_client.NewSyncMqClient(kafka.MustKafkaProducer(c.BotsClient)),
		PushClient: sync_client.NewSyncMqClient(kafka.MustKafkaProducer(c.PushClient)),
	}
	go d.watch(c.SessionClient)
	return d
}
