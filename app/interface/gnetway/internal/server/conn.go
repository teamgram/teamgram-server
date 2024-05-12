// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package server

import (
	"bytes"

	"github.com/teamgram/teamgram-server/app/interface/gnetway/internal/server/codec"
	"github.com/teamgram/teamgram-server/app/interface/gnetway/internal/server/ws"

	"github.com/zeromicro/go-zero/core/jsonx"
)

type HandshakeStateCtx struct {
	State         int32  `json:"state,omitempty"`
	ResState      int32  `json:"res_state,omitempty"`
	Nonce         []byte `json:"nonce,omitempty"`
	ServerNonce   []byte `json:"server_nonce,omitempty"`
	NewNonce      []byte `json:"new_nonce,omitempty"`
	A             []byte `json:"a,omitempty"`
	P             []byte `json:"p,omitempty"`
	HandshakeType int    `json:"handshake_type"`
	ExpiresIn     int32  `json:"expires_in,omitempty"`
}

func (m *HandshakeStateCtx) DebugString() string {
	s, _ := jsonx.MarshalToString(m)
	return s
}

type connContext struct {
	codec         codec.Codec
	authKeys      []*authKeyUtil
	sessionId     int64
	handshakes    []*HandshakeStateCtx
	clientIp      string
	tcp           bool
	websocket     bool
	wsCodec       *ws.WsCodec
	permAuthKeyId int64
}

func newConnContext() *connContext {
	return &connContext{
		codec:    nil,
		clientIp: "",
	}
}

func (ctx *connContext) setClientIp(ip string) {
	ctx.clientIp = ip
}

func (ctx *connContext) getAuthKey(id int64) *authKeyUtil {
	for _, key := range ctx.authKeys {
		if key.AuthKeyId() == id {
			return key
		}
	}

	return nil
}

func (ctx *connContext) putAuthKey(k *authKeyUtil) {
	for _, key := range ctx.authKeys {
		if key.Equal(k) {
			return
		}
	}

	ctx.authKeys = append(ctx.authKeys, k)
}

func (ctx *connContext) getAllAuthKeyId() (idList []int64) {
	idList = make([]int64, len(ctx.authKeys))
	for i, key := range ctx.authKeys {
		idList[i] = key.AuthKeyId()
	}

	return
}

func (ctx *connContext) getHandshakeStateCtx(nonce []byte) *HandshakeStateCtx {
	for _, state := range ctx.handshakes {
		if bytes.Equal(nonce, state.Nonce) {
			return state
		}
	}

	return nil
}

func (ctx *connContext) putHandshakeStateCt(state *HandshakeStateCtx) {
	ctx.handshakes = append(ctx.handshakes, state)
}
