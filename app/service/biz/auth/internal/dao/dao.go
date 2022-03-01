/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dao

import "github.com/teamgram/teamgram-server/app/service/biz/auth/internal/config"

type Dao struct {
}

func New(c config.Config) *Dao {
	return new(Dao)
}
