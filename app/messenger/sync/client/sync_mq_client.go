// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package sync_client

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"

	kafka "github.com/teamgram/marmota/pkg/mq"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"

	"github.com/gogo/protobuf/proto"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type defaultSyncMqClient struct {
	cli *kafka.Producer
}

func NewSyncMqClient(cli *kafka.Producer) SyncClient {
	return &defaultSyncMqClient{
		cli: cli,
	}
}

func (m *defaultSyncMqClient) sendMessage(ctx context.Context, k string, in interface{}) (*mtproto.Void, error) {
	var (
		b   []byte
		err error
	)

	b, err = jsonx.Marshal(in)
	if err != nil {
		return nil, err
	}

	_, _, err = m.cli.SendMessage(ctx, k, b)
	if err != nil {
		return nil, err
	}

	return mtproto.EmptyVoid, nil
}

// SyncUpdatesMe
// sync.updatesMe flags:# user_id:long auth_key_id:long server_id:string session_id:flags.0?long updates:Updates = Void;
func (m *defaultSyncMqClient) SyncUpdatesMe(ctx context.Context, in *sync.TLSyncUpdatesMe) (*mtproto.Void, error) {
	return m.sendMessage(ctx, proto.MessageName(in), in)
}

// SyncUpdatesNotMe
// sync.updatesNotMe user_id:long auth_key_id:long updates:Updates = Void;
func (m *defaultSyncMqClient) SyncUpdatesNotMe(ctx context.Context, in *sync.TLSyncUpdatesNotMe) (*mtproto.Void, error) {
	logx.Infof("send sync.updatesNotMe: %s", in.DebugString())
	return m.sendMessage(ctx, proto.MessageName(in), in)
}

// SyncPushUpdates
// sync.pushUpdates user_id:long updates:Updates = Void;
func (m *defaultSyncMqClient) SyncPushUpdates(ctx context.Context, in *sync.TLSyncPushUpdates) (*mtproto.Void, error) {
	return m.sendMessage(ctx, proto.MessageName(in), in)
}

// SyncPushUpdatesIfNot
// sync.pushUpdatesIfNot user_id:long excludes:Vector<long> updates:Updates = Void;
func (m *defaultSyncMqClient) SyncPushUpdatesIfNot(ctx context.Context, in *sync.TLSyncPushUpdatesIfNot) (*mtproto.Void, error) {
	return m.sendMessage(ctx, proto.MessageName(in), in)
}

// SyncPushBotUpdates
// sync.pushBotUpdates user_id:long updates:Updates = Void;
func (m *defaultSyncMqClient) SyncPushBotUpdates(ctx context.Context, in *sync.TLSyncPushBotUpdates) (*mtproto.Void, error) {
	return m.sendMessage(ctx, proto.MessageName(in), in)
}

// SyncPushRpcResult
// sync.pushRpcResult auth_key_id:long server_id:string session_id:long client_req_msg_id:long rpc_result:bytes = Void;
func (m *defaultSyncMqClient) SyncPushRpcResult(ctx context.Context, in *sync.TLSyncPushRpcResult) (*mtproto.Void, error) {
	return m.sendMessage(ctx, proto.MessageName(in), in)
}

// SyncBroadcastUpdates
// sync.broadcastUpdates broadcast_type:int chat_id:long exclude_id_list:Vector<long> updates:Updates = Void;
func (m *defaultSyncMqClient) SyncBroadcastUpdates(ctx context.Context, in *sync.TLSyncBroadcastUpdates) (*mtproto.Void, error) {
	return m.sendMessage(ctx, proto.MessageName(in), in)
}
