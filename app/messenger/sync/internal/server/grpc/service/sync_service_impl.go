/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/sync/internal/core"
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"
)

// SyncUpdatesMe
// sync.updatesMe flags:# user_id:long auth_key_id:long server_id:string session_id:flags.0?long updates:Updates = Void;
func (s *Service) SyncUpdatesMe(ctx context.Context, request *sync.TLSyncUpdatesMe) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("sync.updatesMe - metadata: %s, request: %s", c.MD, request)

	r, err := c.SyncUpdatesMe(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("sync.updatesMe - reply: %s", r)
	return r, err
}

// SyncUpdatesNotMe
// sync.updatesNotMe user_id:long auth_key_id:long updates:Updates = Void;
func (s *Service) SyncUpdatesNotMe(ctx context.Context, request *sync.TLSyncUpdatesNotMe) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("sync.updatesNotMe - metadata: %s, request: %s", c.MD, request)

	r, err := c.SyncUpdatesNotMe(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("sync.updatesNotMe - reply: %s", r)
	return r, err
}

// SyncPushUpdates
// sync.pushUpdates user_id:long updates:Updates = Void;
func (s *Service) SyncPushUpdates(ctx context.Context, request *sync.TLSyncPushUpdates) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("sync.pushUpdates - metadata: %s, request: %s", c.MD, request)

	r, err := c.SyncPushUpdates(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("sync.pushUpdates - reply: %s", r)
	return r, err
}

// SyncPushUpdatesIfNot
// sync.pushUpdatesIfNot user_id:long excludes:Vector<long> updates:Updates = Void;
func (s *Service) SyncPushUpdatesIfNot(ctx context.Context, request *sync.TLSyncPushUpdatesIfNot) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("sync.pushUpdatesIfNot - metadata: %s, request: %s", c.MD, request)

	r, err := c.SyncPushUpdatesIfNot(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("sync.pushUpdatesIfNot - reply: %s", r)
	return r, err
}

// SyncPushBotUpdates
// sync.pushBotUpdates user_id:long updates:Updates = Void;
func (s *Service) SyncPushBotUpdates(ctx context.Context, request *sync.TLSyncPushBotUpdates) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("sync.pushBotUpdates - metadata: %s, request: %s", c.MD, request)

	c.Logger.Errorf("sync.pushBotUpdates blocked, License key from https://teamgram.net required to unlock enterprise features.")

	return mtproto.EmptyVoid, nil
}

// SyncPushRpcResult
// sync.pushRpcResult auth_key_id:long server_id:string session_id:long client_req_msg_id:long rpc_result:bytes = Void;
func (s *Service) SyncPushRpcResult(ctx context.Context, request *sync.TLSyncPushRpcResult) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("sync.pushRpcResult - metadata: %s, request: %s", c.MD, request)

	r, err := c.SyncPushRpcResult(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("sync.pushRpcResult - reply: %s", r)
	return r, err
}

// SyncBroadcastUpdates
// sync.broadcastUpdates broadcast_type:int chat_id:long exclude_id_list:Vector<long> updates:Updates = Void;
func (s *Service) SyncBroadcastUpdates(ctx context.Context, request *sync.TLSyncBroadcastUpdates) (*mtproto.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("sync.broadcastUpdates - metadata: %s, request: %s", c.MD, request)

	r, err := c.SyncBroadcastUpdates(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("sync.broadcastUpdates - reply: %s", r)
	return r, err
}
