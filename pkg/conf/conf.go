// Copyright 2022 Teamgooo Authors
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

type DialogOutboxWorkersConf struct {
	Enabled        bool  `json:",optional"`
	BatchSize      int32 `json:",optional"`
	LeaseSeconds   int32 `json:",optional"`
	PollIntervalMs int32 `json:",optional"`
}

type SmsVerifyCodeConfig struct {
	Name          string
	SendCodeUrl   string
	VerifyCodeUrl string
	Key           string
	Secret        string
	RegionId      string
}
