// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package server

import (
	"bytes"
	"fmt"
	"strings"
	"sync"

	"github.com/teamgram/marmota/pkg/timer2"
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
	handshakeType int
	ExpiresIn     int32 `json:"expires_in,omitempty"`
}

func (m *HandshakeStateCtx) DebugString() string {
	s, _ := jsonx.MarshalToString(m)
	return s
}

type connContext struct {
	// TODO(@benqi): lock
	sync.Mutex
	codec           codec.Codec
	state           int // 是否握手阶段
	authKeys        []*authKeyUtil
	sessionId       int64
	isHttp          bool
	canSend         bool
	trd             *timer2.TimerData
	handshakes      []*HandshakeStateCtx
	clientIp        string
	xForwardedForIp string
	tcp             bool
	websocket       bool
	wsCodec         *ws.WsCodec
}

func newConnContext() *connContext {
	return &connContext{
		codec:           nil,
		state:           STATE_CONNECTED2,
		clientIp:        "",
		xForwardedForIp: "*",
	}
}

func (ctx *connContext) getClientIp(xForwarderForIp interface{}) string {
	ctx.Lock()
	defer ctx.Unlock()
	if ctx.xForwardedForIp == "*" {
		ctx.xForwardedForIp = ""
		if xForwarderForIp != nil {
			ctx.xForwardedForIp, _ = xForwarderForIp.(string)
		}
	}

	if ctx.xForwardedForIp != "" {
		return ctx.xForwardedForIp
	}
	return ctx.clientIp
}

func (ctx *connContext) setClientIp(ip string) {
	ctx.Lock()
	defer ctx.Unlock()

	ctx.clientIp = ip
}

func (ctx *connContext) getState() int {
	ctx.Lock()
	defer ctx.Unlock()
	return ctx.state
}

func (ctx *connContext) setState(state int) {
	ctx.Lock()
	defer ctx.Unlock()
	if ctx.state != state {
		ctx.state = state
	}
}

func (ctx *connContext) getAuthKey(id int64) *authKeyUtil {
	ctx.Lock()
	defer ctx.Unlock()
	for _, key := range ctx.authKeys {
		if key.AuthKeyId() == id {
			return key
		}
	}

	return nil
}

func (ctx *connContext) putAuthKey(k *authKeyUtil) {
	ctx.Lock()
	defer ctx.Unlock()
	for _, key := range ctx.authKeys {
		if key.Equal(k) {
			return
		}
	}

	ctx.authKeys = append(ctx.authKeys, k)
}

func (ctx *connContext) getAllAuthKeyId() (idList []int64) {
	ctx.Lock()
	defer ctx.Unlock()

	idList = make([]int64, len(ctx.authKeys))
	for i, key := range ctx.authKeys {
		idList[i] = key.AuthKeyId()
	}

	return
}

func (ctx *connContext) getHandshakeStateCtx(nonce []byte) *HandshakeStateCtx {
	ctx.Lock()
	defer ctx.Unlock()

	for _, state := range ctx.handshakes {
		if bytes.Equal(nonce, state.Nonce) {
			return state
		}
	}

	return nil
}

func (ctx *connContext) putHandshakeStateCt(state *HandshakeStateCtx) {
	ctx.Lock()
	defer ctx.Unlock()

	ctx.handshakes = append(ctx.handshakes, state)
}

func (ctx *connContext) encryptedMessageAble() bool {
	ctx.Lock()
	defer ctx.Unlock()
	//return ctx.state == STATE_CONNECTED2 ||
	//	ctx.state == STATE_AUTH_KEY ||
	//	(ctx.state == STATE_HANDSHAKE &&
	//		(ctx.handshakeCtx.State == STATE_pq_res ||
	//			(ctx.handshakeCtx.State == STATE_dh_gen_res &&
	//				ctx.handshakeCtx.ResState == RES_STATE_OK)))
	return true
}

func (ctx *connContext) DebugString() string {
	s := make([]string, 0, 4)
	s = append(s, fmt.Sprintf(`"state":%d`, ctx.state))
	// s = append(s, fmt.Sprintf(`"handshake_ctx":%s`, ctx.handshakeCtx.DebugString()))
	//if ctx.authKey != nil {
	//	s = append(s, fmt.Sprintf(`"auth_key_id":%d`, ctx.authKey.AuthKeyId()))
	//}
	return "{" + strings.Join(s, ",") + "}"
}
