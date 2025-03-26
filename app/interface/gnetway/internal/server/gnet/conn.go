// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package gnet

import (
	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/teamgram-server/v2/app/interface/gnetway/internal/server/gnet/codec"
	"github.com/teamgram/teamgram-server/v2/app/interface/gnetway/internal/server/gnet/ws"

	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

type HandshakeStateCtx struct {
	State         int32      `json:"state,omitempty"`
	ResState      int32      `json:"res_state,omitempty"`
	Nonce         bin.Int128 `json:"nonce,omitempty"`
	ServerNonce   bin.Int128 `json:"server_nonce,omitempty"`
	NewNonce      bin.Int256 `json:"new_nonce,omitempty"`
	A             []byte     `json:"a,omitempty"`
	P             []byte     `json:"p,omitempty"`
	HandshakeType int        `json:"handshake_type"`
	ExpiresIn     int32      `json:"expires_in,omitempty"`
}

func (m *HandshakeStateCtx) DebugString() string {
	s, _ := jsonx.MarshalToString(m)
	return s
}

type connContext struct {
	codec      codec.Codec
	authKey    *authKeyUtil
	sessionId  int64
	handshakes []*HandshakeStateCtx
	clientIp   string
	tcp        bool
	websocket  bool
	wsCodec    *ws.WsCodec
	logx.Logger
	newSession bool
	nextSeqNo  int32
	closeDate  int64
}

func newConnContext() *connContext {
	return &connContext{
		codec:    nil,
		clientIp: "",
	}
}

func (ctx *connContext) generateMessageSeqNo(increment bool) int32 {
	value := ctx.nextSeqNo
	if increment {
		ctx.nextSeqNo++
		return value*2 + 1
	} else {
		return value * 2
	}
}

func (ctx *connContext) setClientIp(ip string) {
	ctx.clientIp = ip
}

func (ctx *connContext) getAuthKey() *authKeyUtil {
	return ctx.authKey
}

func (ctx *connContext) putAuthKey(k *authKeyUtil) {
	ctx.authKey = k
}

func (ctx *connContext) getHandshakeStateCtx(nonce bin.Int128) *HandshakeStateCtx {
	for _, state := range ctx.handshakes {
		if nonce == state.Nonce {
			return state
		}
	}

	return nil
}

func (ctx *connContext) putHandshakeStateCt(state *HandshakeStateCtx) {
	ctx.handshakes = append(ctx.handshakes, state)
}
