// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package dao

import (
	"context"
	"strconv"

	"github.com/teamgram/proto/mtproto"
	sessionclient "github.com/teamgram/teamgram-server/app/interface/session/client"
	"github.com/teamgram/teamgram-server/app/interface/session/session"
)

// SessionDispatcher abstracts the communication between gnetway and session.
// UnarySessionDispatcher wraps ShardingSessionClient (existing unary RPC).
// StreamingSessionDispatcher uses bidirectional gRPC streams.
type SessionDispatcher interface {
	SendData(ctx context.Context, permAuthKeyId int64, in *session.TLSessionSendDataToSession) error
	CloseSession(ctx context.Context, permAuthKeyId int64, in *session.TLSessionCloseSession) error
	QueryAuthKey(ctx context.Context, authKeyId int64, in *session.TLSessionQueryAuthKey) (*mtproto.AuthKeyInfo, error)
	Close()
}

// UnarySessionDispatcher wraps ShardingSessionClient to implement SessionDispatcher.
type UnarySessionDispatcher struct {
	client *ShardingSessionClient
}

func NewUnarySessionDispatcher(client *ShardingSessionClient) *UnarySessionDispatcher {
	return &UnarySessionDispatcher{client: client}
}

func (d *UnarySessionDispatcher) SendData(ctx context.Context, permAuthKeyId int64, in *session.TLSessionSendDataToSession) error {
	return d.client.InvokeByKey(
		strconv.FormatInt(permAuthKeyId, 10),
		func(client sessionclient.SessionClient) error {
			_, err := client.SessionSendDataToSession(ctx, in)
			return err
		})
}

func (d *UnarySessionDispatcher) CloseSession(ctx context.Context, permAuthKeyId int64, in *session.TLSessionCloseSession) error {
	return d.client.InvokeByKey(
		strconv.FormatInt(permAuthKeyId, 10),
		func(client sessionclient.SessionClient) error {
			_, err := client.SessionCloseSession(ctx, in)
			return err
		})
}

func (d *UnarySessionDispatcher) QueryAuthKey(ctx context.Context, authKeyId int64, in *session.TLSessionQueryAuthKey) (*mtproto.AuthKeyInfo, error) {
	var key *mtproto.AuthKeyInfo
	err := d.client.InvokeByKey(
		strconv.FormatInt(authKeyId, 10),
		func(client sessionclient.SessionClient) (err error) {
			key, err = client.SessionQueryAuthKey(ctx, in)
			return
		})
	return key, err
}

func (d *UnarySessionDispatcher) Close() {
	// unary client lifecycle managed by etcd watcher, nothing to do
}
