/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package sessionclient

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/interface/session/session"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type SessionClient interface {
	SessionQueryAuthKey(ctx context.Context, in *session.TLSessionQueryAuthKey) (*mtproto.AuthKeyInfo, error)
	SessionSetAuthKey(ctx context.Context, in *session.TLSessionSetAuthKey) (*mtproto.Bool, error)
	SessionCreateSession(ctx context.Context, in *session.TLSessionCreateSession) (*mtproto.Bool, error)
	SessionSendDataToSession(ctx context.Context, in *session.TLSessionSendDataToSession) (*mtproto.Bool, error)
	SessionSendHttpDataToSession(ctx context.Context, in *session.TLSessionSendHttpDataToSession) (*session.HttpSessionData, error)
	SessionCloseSession(ctx context.Context, in *session.TLSessionCloseSession) (*mtproto.Bool, error)
	SessionPushUpdatesData(ctx context.Context, in *session.TLSessionPushUpdatesData) (*mtproto.Bool, error)
	SessionPushSessionUpdatesData(ctx context.Context, in *session.TLSessionPushSessionUpdatesData) (*mtproto.Bool, error)
	SessionPushRpcResultData(ctx context.Context, in *session.TLSessionPushRpcResultData) (*mtproto.Bool, error)
}

type defaultSessionClient struct {
	cli zrpc.Client
}

func NewSessionClient(cli zrpc.Client) SessionClient {
	return &defaultSessionClient{
		cli: cli,
	}
}

// SessionQueryAuthKey
// session.queryAuthKey auth_key_id:long = AuthKeyInfo;
func (m *defaultSessionClient) SessionQueryAuthKey(ctx context.Context, in *session.TLSessionQueryAuthKey) (*mtproto.AuthKeyInfo, error) {
	client := session.NewRPCSessionClient(m.cli.Conn())
	return client.SessionQueryAuthKey(ctx, in)
}

// SessionSetAuthKey
// session.setAuthKey auth_key:AuthKeyInfo future_salt:FutureSalt expires_in:int = Bool;
func (m *defaultSessionClient) SessionSetAuthKey(ctx context.Context, in *session.TLSessionSetAuthKey) (*mtproto.Bool, error) {
	client := session.NewRPCSessionClient(m.cli.Conn())
	return client.SessionSetAuthKey(ctx, in)
}

// SessionCreateSession
// session.createSession client:SessionClientEvent = Bool;
func (m *defaultSessionClient) SessionCreateSession(ctx context.Context, in *session.TLSessionCreateSession) (*mtproto.Bool, error) {
	client := session.NewRPCSessionClient(m.cli.Conn())
	return client.SessionCreateSession(ctx, in)
}

// SessionSendDataToSession
// session.sendDataToSession data:SessionClientData = Bool;
func (m *defaultSessionClient) SessionSendDataToSession(ctx context.Context, in *session.TLSessionSendDataToSession) (*mtproto.Bool, error) {
	client := session.NewRPCSessionClient(m.cli.Conn())
	return client.SessionSendDataToSession(ctx, in)
}

// SessionSendHttpDataToSession
// session.sendHttpDataToSession client:SessionClientData = HttpSessionData;
func (m *defaultSessionClient) SessionSendHttpDataToSession(ctx context.Context, in *session.TLSessionSendHttpDataToSession) (*session.HttpSessionData, error) {
	client := session.NewRPCSessionClient(m.cli.Conn())
	return client.SessionSendHttpDataToSession(ctx, in)
}

// SessionCloseSession
// session.closeSession client:SessionClientEvent = Bool;
func (m *defaultSessionClient) SessionCloseSession(ctx context.Context, in *session.TLSessionCloseSession) (*mtproto.Bool, error) {
	client := session.NewRPCSessionClient(m.cli.Conn())
	return client.SessionCloseSession(ctx, in)
}

// SessionPushUpdatesData
// session.pushUpdatesData flags:# auth_key_id:long notification:flags.0?true updates:Updates = Bool;
func (m *defaultSessionClient) SessionPushUpdatesData(ctx context.Context, in *session.TLSessionPushUpdatesData) (*mtproto.Bool, error) {
	client := session.NewRPCSessionClient(m.cli.Conn())
	return client.SessionPushUpdatesData(ctx, in)
}

// SessionPushSessionUpdatesData
// session.pushSessionUpdatesData auth_key_id:long session_id:long updates:Updates = Bool;
func (m *defaultSessionClient) SessionPushSessionUpdatesData(ctx context.Context, in *session.TLSessionPushSessionUpdatesData) (*mtproto.Bool, error) {
	client := session.NewRPCSessionClient(m.cli.Conn())
	return client.SessionPushSessionUpdatesData(ctx, in)
}

// SessionPushRpcResultData
// session.pushRpcResultData auth_key_id:long session_id:long client_req_msg_id:long rpc_result_data:bytes = Bool;
func (m *defaultSessionClient) SessionPushRpcResultData(ctx context.Context, in *session.TLSessionPushRpcResultData) (*mtproto.Bool, error) {
	client := session.NewRPCSessionClient(m.cli.Conn())
	return client.SessionPushRpcResultData(ctx, in)
}
