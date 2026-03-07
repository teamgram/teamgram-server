/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package config

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Status       kv.KvConf
	StatusExpire int
}

func (c *Config) Validate() error {
	if c.StatusExpire <= 0 {
		return fmt.Errorf("StatusExpire must be positive, got %d", c.StatusExpire)
	}
	return nil
}
