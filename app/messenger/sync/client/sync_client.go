/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package syncclient

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync"
	"github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync/syncservice"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
)

var _ *tg.Bool

type SyncClient interface {
	SyncUpdatesMe(ctx context.Context, in *sync.TLSyncUpdatesMe) (*tg.Void, error)
	SyncUpdatesNotMe(ctx context.Context, in *sync.TLSyncUpdatesNotMe) (*tg.Void, error)
	SyncPushUpdates(ctx context.Context, in *sync.TLSyncPushUpdates) (*tg.Void, error)
	SyncPushUpdatesIfNot(ctx context.Context, in *sync.TLSyncPushUpdatesIfNot) (*tg.Void, error)
	SyncPushRpcResult(ctx context.Context, in *sync.TLSyncPushRpcResult) (*tg.Void, error)
}

type defaultSyncClient struct {
	cli client.Client
	rpc syncservice.Client
}

func NewSyncClient(cli client.Client) SyncClient {
	return &defaultSyncClient{
		cli: cli,
		rpc: syncservice.NewRPCSyncClient(cli),
	}
}

// SyncUpdatesMe
// sync.updatesMe flags:# user_id:long perm_auth_key_id:long auth_key_id:flags.0?long session_id:flags.1?long updates:Updates = Void;
func (m *defaultSyncClient) SyncUpdatesMe(ctx context.Context, in *sync.TLSyncUpdatesMe) (*tg.Void, error) {
	return m.rpc.SyncUpdatesMe(ctx, in)
}

// SyncUpdatesNotMe
// sync.updatesNotMe user_id:long perm_auth_key_id:long updates:Updates = Void;
func (m *defaultSyncClient) SyncUpdatesNotMe(ctx context.Context, in *sync.TLSyncUpdatesNotMe) (*tg.Void, error) {
	return m.rpc.SyncUpdatesNotMe(ctx, in)
}

// SyncPushUpdates
// sync.pushUpdates user_id:long updates:Updates = Void;
func (m *defaultSyncClient) SyncPushUpdates(ctx context.Context, in *sync.TLSyncPushUpdates) (*tg.Void, error) {
	return m.rpc.SyncPushUpdates(ctx, in)
}

// SyncPushUpdatesIfNot
// sync.pushUpdatesIfNot flags:# user_id:long includes:flags.0?Vector<long> excludes:flags.1?Vector<long> updates:Updates = Void;
func (m *defaultSyncClient) SyncPushUpdatesIfNot(ctx context.Context, in *sync.TLSyncPushUpdatesIfNot) (*tg.Void, error) {
	return m.rpc.SyncPushUpdatesIfNot(ctx, in)
}

// SyncPushRpcResult
// sync.pushRpcResult user_id:long perm_auth_key_id:long auth_key_id:long gateway_id:string gateway_generation:string gateway_rpc_addr:string session_id:long client_req_msg_id:long rpc_result:bytes = Void;
func (m *defaultSyncClient) SyncPushRpcResult(ctx context.Context, in *sync.TLSyncPushRpcResult) (*tg.Void, error) {
	return m.rpc.SyncPushRpcResult(ctx, in)
}
