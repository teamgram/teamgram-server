// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package service

import (
	"github.com/zeromicro/go-zero/core/netx"
	"os"
)

const (
	envPodIp = "POD_IP"
)

var (
	internalIp string
)

func init() {
	internalIp = os.Getenv(envPodIp)
	if len(internalIp) == 0 {
		internalIp = netx.InternalIp()
	}
}
