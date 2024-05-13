/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package gateway_helper

import (
	"github.com/teamgram/teamgram-server/app/interface/gateway/internal/server/server"
	"github.com/zeromicro/go-zero/zrpc"
)

func init() {
	zrpc.DontLogContentForMethod("/gateway.RPCGateway/gateway_sendDataToGateway")

	zrpc.DontLogClientContentForMethod("/session.RPCSession/session_sendDataToSession")
}

type (
	Server = server.Server
)
