// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package dao

import (
	"github.com/teamgram/teamgram-server/app/interface/gnetway/internal/config"
)

type Dao struct {
	*ShardingSessionClient
	SessionDispatcher SessionDispatcher
}

func New(c config.Config) *Dao {
	shardingClient := NewShardingSessionClient(c)
	d := &Dao{
		ShardingSessionClient: shardingClient,
	}

	if c.UseStreamSession {
		d.SessionDispatcher = NewStreamingSessionDispatcher(c)
	} else {
		d.SessionDispatcher = NewUnarySessionDispatcher(shardingClient)
	}

	return d
}
