// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package conf

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type BFFProxyClients struct {
	Clients []zrpc.RpcClientConf
	IDMap   map[string]string
}

type ZRpcServerConf struct {
	zrpc.RpcServerConf
	WriteBufferSize int `json:",default=32768"`
	ReadBufferSize  int `json:",default=32768"`
}
