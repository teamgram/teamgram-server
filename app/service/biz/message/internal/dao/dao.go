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
	// "github.com/teamgram/marmota/pkg/net/rpcx"
	"github.com/teamgram/marmota/pkg/stores/sqlc"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/message/internal/config"
	// poll_client "github.com/teamgram/teamgram-server/app/service/poll/client"
)

// Dao dao.
type Dao struct {
	*Mysql
	sqlc.CachedConn
	// poll_client.PollClient
}

// New new a dao and return.
func New(c config.Config) (dao *Dao) {
	db := sqlx.NewMySQL(&c.Mysql)
	return &Dao{
		Mysql:      newMysqlDao(db),
		CachedConn: sqlc.NewConn(db, c.Cache),
		// PollClient: poll_client.NewPollClient(rpcx.GetCachedRpcClient(c.PollClient)),
	}
}
