// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package sess

import (
	"context"
	"fmt"
)

type connDataCtx struct {
	ctx context.Context
	connData
}

type connData struct {
	authType  int
	authId    int64
	isNew     bool
	gatewayId string
	sessionId int64
}

func (c *connData) DebugString() string {
	return fmt.Sprintf("{isNew: %v, gatewayId: %s, sessionId: %d}", c.isNew, c.gatewayId, c.sessionId)
}

type sessionDataCtx struct {
	ctx context.Context
	sessionData
}

type sessionData struct {
	authType  int
	authId    int64
	gatewayId string
	clientIp  string
	sessionId int64
	salt      int64
	buf       []byte
}

type sessionHttpDataCtx struct {
	ctx context.Context
	sessionHttpData
}

type sessionHttpData struct {
	authType   int
	authId     int64
	gatewayId  string
	clientIp   string
	sessionId  int64
	salt       int64
	buf        []byte
	resChannel chan interface{}
}

type syncRpcResultDataCtx struct {
	ctx context.Context
	syncRpcResultData
}

type syncRpcResultData struct {
	authType    int
	authId      int64
	sessionId   int64
	clientMsgId int64
	data        []byte
}

type syncSessionDataCtx struct {
	ctx context.Context
	syncSessionData
}

type syncSessionData struct {
	authType  int
	authId    int64
	sessionId int64
	data      *messageData
}

type syncDataCtx struct {
	ctx context.Context
	syncData
}

type syncData struct {
	needAndroidPush bool
	data            *messageData
}
