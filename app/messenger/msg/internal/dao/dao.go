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
	inbox_client "github.com/teamgram/teamgram-server/app/messenger/msg/inbox/client"
	sync_client "github.com/teamgram/teamgram-server/app/messenger/sync/client"
	chat_client "github.com/teamgram/teamgram-server/app/service/biz/chat/client"
	dialog_client "github.com/teamgram/teamgram-server/app/service/biz/dialog/client"
	user_client "github.com/teamgram/teamgram-server/app/service/biz/user/client"
	username_client "github.com/teamgram/teamgram-server/app/service/biz/username/client"
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
	deduplication.MessageDeDuplicate
	*redis.Redis
	username_client.UsernameClient
}

func (d *Dao) DoIdempotent(ctx context.Context, senderUserId int64, deDuplicateKey string, v any, cb func(ctx context.Context, v any) error) (bool, error) {
	return idempotent.DoIdempotent(
		ctx,
		d.Redis,
		fmt.Sprintf("idempotent#%d@%s", senderUserId, deDuplicateKey),
		v,
		5,
		90,
		cb)
}
