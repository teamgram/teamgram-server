/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package syncclient

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync"
	"github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync/syncservice"

	"github.com/cloudwego/kitex/client"
)

var _ *tg.Bool

type SyncClient interface {
	SyncUpdatesMe(ctx context.Context, in *sync.TLSyncUpdatesMe) (*tg.Void, error)
	SyncUpdatesNotMe(ctx context.Context, in *sync.TLSyncUpdatesNotMe) (*tg.Void, error)
	SyncPushUpdates(ctx context.Context, in *sync.TLSyncPushUpdates) (*tg.Void, error)
	SyncPushUpdatesIfNot(ctx context.Context, in *sync.TLSyncPushUpdatesIfNot) (*tg.Void, error)
	SyncPushBotUpdates(ctx context.Context, in *sync.TLSyncPushBotUpdates) (*tg.Void, error)
	SyncPushRpcResult(ctx context.Context, in *sync.TLSyncPushRpcResult) (*tg.Void, error)
	SyncBroadcastUpdates(ctx context.Context, in *sync.TLSyncBroadcastUpdates) (*tg.Void, error)
}

type defaultSyncClient struct {
	cli client.Client
}

func NewSyncClient(cli client.Client) SyncClient {
	return &defaultSyncClient{
		cli: cli,
	}
}

// SyncUpdatesMe
// sync.updatesMe flags:# user_id:long perm_auth_key_id:long server_id:flags.0?string auth_key_id:flags.1?long session_id:flags.1?long updates:Updates = Void;
func (m *defaultSyncClient) SyncUpdatesMe(ctx context.Context, in *sync.TLSyncUpdatesMe) (*tg.Void, error) {
	cli := syncservice.NewRPCSyncClient(m.cli)
	return cli.SyncUpdatesMe(ctx, in)
}

// SyncUpdatesNotMe
// sync.updatesNotMe user_id:long perm_auth_key_id:long updates:Updates = Void;
func (m *defaultSyncClient) SyncUpdatesNotMe(ctx context.Context, in *sync.TLSyncUpdatesNotMe) (*tg.Void, error) {
	cli := syncservice.NewRPCSyncClient(m.cli)
	return cli.SyncUpdatesNotMe(ctx, in)
}

// SyncPushUpdates
// sync.pushUpdates user_id:long updates:Updates = Void;
func (m *defaultSyncClient) SyncPushUpdates(ctx context.Context, in *sync.TLSyncPushUpdates) (*tg.Void, error) {
	cli := syncservice.NewRPCSyncClient(m.cli)
	return cli.SyncPushUpdates(ctx, in)
}

// SyncPushUpdatesIfNot
// sync.pushUpdatesIfNot flags:# user_id:long includes:flags.0?Vector<long> excludes:flags.1?Vector<long> updates:Updates = Void;
func (m *defaultSyncClient) SyncPushUpdatesIfNot(ctx context.Context, in *sync.TLSyncPushUpdatesIfNot) (*tg.Void, error) {
	cli := syncservice.NewRPCSyncClient(m.cli)
	return cli.SyncPushUpdatesIfNot(ctx, in)
}

// SyncPushBotUpdates
// sync.pushBotUpdates user_id:long updates:Updates = Void;
func (m *defaultSyncClient) SyncPushBotUpdates(ctx context.Context, in *sync.TLSyncPushBotUpdates) (*tg.Void, error) {
	cli := syncservice.NewRPCSyncClient(m.cli)
	return cli.SyncPushBotUpdates(ctx, in)
}

// SyncPushRpcResult
// sync.pushRpcResult user_id:long auth_key_id:long perm_auth_key_id:long server_id:string session_id:long client_req_msg_id:long rpc_result:bytes = Void;
func (m *defaultSyncClient) SyncPushRpcResult(ctx context.Context, in *sync.TLSyncPushRpcResult) (*tg.Void, error) {
	cli := syncservice.NewRPCSyncClient(m.cli)
	return cli.SyncPushRpcResult(ctx, in)
}

// SyncBroadcastUpdates
// sync.broadcastUpdates broadcast_type:int chat_id:long exclude_id_list:Vector<long> updates:Updates = Void;
func (m *defaultSyncClient) SyncBroadcastUpdates(ctx context.Context, in *sync.TLSyncBroadcastUpdates) (*tg.Void, error) {
	cli := syncservice.NewRPCSyncClient(m.cli)
	return cli.SyncBroadcastUpdates(ctx, in)
}
