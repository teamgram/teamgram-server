/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package sessionclient

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/interface/session/session"
	"github.com/teamgram/teamgram-server/v2/app/interface/session/session/sessionservice"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
)

var _ *tg.Bool

type SessionClient interface {
	SessionQueryAuthKey(ctx context.Context, in *session.TLSessionQueryAuthKey) (*tg.AuthKeyInfo, error)
	SessionSetAuthKey(ctx context.Context, in *session.TLSessionSetAuthKey) (*tg.Bool, error)
	SessionCreateSession(ctx context.Context, in *session.TLSessionCreateSession) (*tg.Bool, error)
	SessionSendDataToSession(ctx context.Context, in *session.TLSessionSendDataToSession) (*tg.Bool, error)
	SessionSendHttpDataToSession(ctx context.Context, in *session.TLSessionSendHttpDataToSession) (*session.HttpSessionData, error)
	SessionCloseSession(ctx context.Context, in *session.TLSessionCloseSession) (*tg.Bool, error)
	SessionPushUpdatesData(ctx context.Context, in *session.TLSessionPushUpdatesData) (*tg.Bool, error)
	SessionPushSessionUpdatesData(ctx context.Context, in *session.TLSessionPushSessionUpdatesData) (*tg.Bool, error)
	SessionPushRpcResultData(ctx context.Context, in *session.TLSessionPushRpcResultData) (*tg.Bool, error)
}

type defaultSessionClient struct {
	cli client.Client
	rpc sessionservice.Client
}

func NewSessionClient(cli client.Client) SessionClient {
	return &defaultSessionClient{
		cli: cli,
		rpc: sessionservice.NewRPCSessionClient(cli),
	}
}

// SessionQueryAuthKey
// session.queryAuthKey auth_key_id:long = AuthKeyInfo;
func (m *defaultSessionClient) SessionQueryAuthKey(ctx context.Context, in *session.TLSessionQueryAuthKey) (*tg.AuthKeyInfo, error) {
	return m.rpc.SessionQueryAuthKey(ctx, in)
}

// SessionSetAuthKey
// session.setAuthKey auth_key:AuthKeyInfo future_salt:FutureSalt expires_in:int = Bool;
func (m *defaultSessionClient) SessionSetAuthKey(ctx context.Context, in *session.TLSessionSetAuthKey) (*tg.Bool, error) {
	return m.rpc.SessionSetAuthKey(ctx, in)
}

// SessionCreateSession
// session.createSession client:SessionClientEvent = Bool;
func (m *defaultSessionClient) SessionCreateSession(ctx context.Context, in *session.TLSessionCreateSession) (*tg.Bool, error) {
	return m.rpc.SessionCreateSession(ctx, in)
}

// SessionSendDataToSession
// session.sendDataToSession data:SessionClientData = Bool;
func (m *defaultSessionClient) SessionSendDataToSession(ctx context.Context, in *session.TLSessionSendDataToSession) (*tg.Bool, error) {
	return m.rpc.SessionSendDataToSession(ctx, in)
}

// SessionSendHttpDataToSession
// session.sendHttpDataToSession client:SessionClientData = HttpSessionData;
func (m *defaultSessionClient) SessionSendHttpDataToSession(ctx context.Context, in *session.TLSessionSendHttpDataToSession) (*session.HttpSessionData, error) {
	return m.rpc.SessionSendHttpDataToSession(ctx, in)
}

// SessionCloseSession
// session.closeSession client:SessionClientEvent = Bool;
func (m *defaultSessionClient) SessionCloseSession(ctx context.Context, in *session.TLSessionCloseSession) (*tg.Bool, error) {
	return m.rpc.SessionCloseSession(ctx, in)
}

// SessionPushUpdatesData
// session.pushUpdatesData flags:# perm_auth_key_id:long notification:flags.0?true updates:Updates = Bool;
func (m *defaultSessionClient) SessionPushUpdatesData(ctx context.Context, in *session.TLSessionPushUpdatesData) (*tg.Bool, error) {
	return m.rpc.SessionPushUpdatesData(ctx, in)
}

// SessionPushSessionUpdatesData
// session.pushSessionUpdatesData flags:# perm_auth_key_id:long auth_key_id:long session_id:long updates:Updates = Bool;
func (m *defaultSessionClient) SessionPushSessionUpdatesData(ctx context.Context, in *session.TLSessionPushSessionUpdatesData) (*tg.Bool, error) {
	return m.rpc.SessionPushSessionUpdatesData(ctx, in)
}

// SessionPushRpcResultData
// session.pushRpcResultData perm_auth_key_id:long auth_key_id:long session_id:long client_req_msg_id:long rpc_result_data:bytes = Bool;
func (m *defaultSessionClient) SessionPushRpcResultData(ctx context.Context, in *session.TLSessionPushRpcResultData) (*tg.Bool, error) {
	return m.rpc.SessionPushRpcResultData(ctx, in)
}
