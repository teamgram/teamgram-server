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
	"github.com/teamgram/marmota/pkg/stores/kv"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/status/internal/config"
)

type Dao struct {
	KV kv.Store
	mtproto.RPCNotificationClient
}

func New(c config.Config) *Dao {
	d := &Dao{
		KV: kv.NewStore(c.Status),
	}

	return d
}
