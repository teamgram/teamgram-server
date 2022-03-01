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
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Cache       cache.CacheConf
	AuthSession zrpc.RpcClientConf
	BFFClient   zrpc.RpcClientConf
	// PhoneClient       zrpc.RpcClientConf
	StatusClient  zrpc.RpcClientConf
	GatewayClient zrpc.RpcClientConf
	// BotsClient        zrpc.RpcClientConf
	// MessagesBotClient zrpc.RpcClientConf
}

// Routine routine.
type Routine struct {
	Size uint64
	Chan uint64
}
