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
	"github.com/teamgram/teamgram-server/app/service/status/internal/config"
	"github.com/zeromicro/go-zero/core/stores/kv"
)

type Dao struct {
	KV kv.Store
}

func New(c config.Config) *Dao {
	return &Dao{
		KV: kv.NewStore(c.Status),
	}
}
