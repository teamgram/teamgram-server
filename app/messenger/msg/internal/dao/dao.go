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
	"context"
	"fmt"

	"github.com/teamgram/marmota/pkg/idempotent"
	"github.com/teamgram/proto/mtproto"
	inbox_client "github.com/teamgram/teamgram-server/app/messenger/msg/inbox/client"
	"github.com/teamgram/teamgram-server/app/messenger/msg/msg/plugin"
	sync_client "github.com/teamgram/teamgram-server/app/messenger/sync/client"
	chat_client "github.com/teamgram/teamgram-server/app/service/biz/chat/client"
	dialog_client "github.com/teamgram/teamgram-server/app/service/biz/dialog/client"
	user_client "github.com/teamgram/teamgram-server/app/service/biz/user/client"
	idgen_client "github.com/teamgram/teamgram-server/app/service/idgen/client"
	"github.com/teamgram/teamgram-server/pkg/deduplication"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Dao struct {
	*Mysql
	//KV kv.Store
	idgen_client.IDGenClient2
	user_client.UserClient
	chat_client.ChatClient
	inbox_client.InboxClient
	SyncClient    sync_client.SyncClient
	BotSyncClient sync_client.SyncClient
	dialog_client.DialogClient
	plugin.MsgPlugin
	deduplication.MessageDeDuplicate
	*redis.Redis
}

func (d *Dao) DoIdempotent(ctx context.Context, senderUserId, deDuplicateId int64, cb func(ctx context.Context) (*mtproto.Updates, error)) (*mtproto.Updates, bool, error) {
	v, ok, err := idempotent.DoIdempotent(
		ctx,
		d.Redis,
		fmt.Sprintf("%d@%d", senderUserId, deDuplicateId),
		5,
		90,
		func() (any, error) {
			return cb(ctx)
		})
	if err != nil {
		return nil, false, err
	}

	return v.(*mtproto.Updates), ok, nil
}
