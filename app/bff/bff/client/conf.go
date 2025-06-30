// Copyright © 2025 The Teamgram Authors.
//  All Rights Reserved.
//
// Author: Benqi (wubenqi@gmail.com)

package bffproxyclient

import (
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
)

type BFFProxyClientListConf struct {
	Clients []BFFProxyClientConf
}

type BFFProxyClientConf struct {
	kitex.RpcClientConf
	ServiceNameList []string
}
