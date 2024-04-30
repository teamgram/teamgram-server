/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package httpserverhelper

import (
	"github.com/teamgram/teamgram-server/app/interface/httpserver/internal/server"
	"github.com/zeromicro/go-zero/zrpc"
)

func init() {
	zrpc.DontLogClientContentForMethod("/session.RPCSession/session_sendHttpDataToSession")
}

type (
	Server = server.Server
)
