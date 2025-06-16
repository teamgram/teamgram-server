/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/sync/internal/core"
	"github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync"
)

// SyncUpdatesMe
// sync.updatesMe flags:# user_id:long perm_auth_key_id:long server_id:flags.0?string auth_key_id:flags.1?long session_id:flags.1?long updates:Updates = Void;
func (s *Service) SyncUpdatesMe(ctx context.Context, request *sync.TLSyncUpdatesMe) (*tg.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("sync.updatesMe - metadata: {}, request: %v", request)

	r, err := c.SyncUpdatesMe(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("echos.echo - reply: %v", r)
	return r, err
}

// SyncUpdatesNotMe
// sync.updatesNotMe user_id:long perm_auth_key_id:long updates:Updates = Void;
func (s *Service) SyncUpdatesNotMe(ctx context.Context, request *sync.TLSyncUpdatesNotMe) (*tg.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("sync.updatesNotMe - metadata: {}, request: %v", request)

	r, err := c.SyncUpdatesNotMe(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("echos.echo - reply: %v", r)
	return r, err
}

// SyncPushUpdates
// sync.pushUpdates user_id:long updates:Updates = Void;
func (s *Service) SyncPushUpdates(ctx context.Context, request *sync.TLSyncPushUpdates) (*tg.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("sync.pushUpdates - metadata: {}, request: %v", request)

	r, err := c.SyncPushUpdates(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("echos.echo - reply: %v", r)
	return r, err
}

// SyncPushUpdatesIfNot
// sync.pushUpdatesIfNot user_id:long excludes:Vector<long> updates:Updates = Void;
func (s *Service) SyncPushUpdatesIfNot(ctx context.Context, request *sync.TLSyncPushUpdatesIfNot) (*tg.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("sync.pushUpdatesIfNot - metadata: {}, request: %v", request)

	r, err := c.SyncPushUpdatesIfNot(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("echos.echo - reply: %v", r)
	return r, err
}

// SyncPushBotUpdates
// sync.pushBotUpdates user_id:long updates:Updates = Void;
func (s *Service) SyncPushBotUpdates(ctx context.Context, request *sync.TLSyncPushBotUpdates) (*tg.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("sync.pushBotUpdates - metadata: {}, request: %v", request)

	r, err := c.SyncPushBotUpdates(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("echos.echo - reply: %v", r)
	return r, err
}

// SyncPushRpcResult
// sync.pushRpcResult user_id:long auth_key_id:long perm_auth_key_id:long server_id:string session_id:long client_req_msg_id:long rpc_result:bytes = Void;
func (s *Service) SyncPushRpcResult(ctx context.Context, request *sync.TLSyncPushRpcResult) (*tg.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("sync.pushRpcResult - metadata: {}, request: %v", request)

	r, err := c.SyncPushRpcResult(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("echos.echo - reply: %v", r)
	return r, err
}

// SyncBroadcastUpdates
// sync.broadcastUpdates broadcast_type:int chat_id:long exclude_id_list:Vector<long> updates:Updates = Void;
func (s *Service) SyncBroadcastUpdates(ctx context.Context, request *sync.TLSyncBroadcastUpdates) (*tg.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("sync.broadcastUpdates - metadata: {}, request: %v", request)

	r, err := c.SyncBroadcastUpdates(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("echos.echo - reply: %v", r)
	return r, err
}
