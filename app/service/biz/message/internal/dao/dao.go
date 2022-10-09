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
	"github.com/teamgram/marmota/pkg/stores/sqlc"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/message/internal/config"
	"github.com/teamgram/teamgram-server/app/service/biz/message/internal/plugin"
)

// Dao dao.
type Dao struct {
	*Mysql
	sqlc.CachedConn
	Plugin plugin.MessagePlugin
}

// New new a dao and return.
func New(c config.Config, plugin plugin.MessagePlugin) *Dao {
	db := sqlx.NewMySQL(&c.Mysql)
	return &Dao{
		Mysql:      newMysqlDao(db, c.MessageSharding),
		CachedConn: sqlc.NewConn(db, c.Cache),
		Plugin:     plugin,
	}
}
