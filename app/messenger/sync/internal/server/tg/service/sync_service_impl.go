/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/messenger/sync/internal/core"
	"github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

var _ *tg.Bool

// SyncUpdatesMe
// sync.updatesMe flags:# user_id:long perm_auth_key_id:long auth_key_id:flags.0?long session_id:flags.1?long updates:Updates = Void;
func (s *Service) SyncUpdatesMe(ctx context.Context, request *sync.TLSyncUpdatesMe) (*tg.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("sync.updatesMe - request: %s", request)

	r, err := c.SyncUpdatesMe(request)
	if err != nil {
		c.Logger.Errorf("sync.updatesMe - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("sync.updatesMe - reply: %s", r)
	return r, err
}

// SyncUpdatesNotMe
// sync.updatesNotMe user_id:long perm_auth_key_id:long updates:Updates = Void;
func (s *Service) SyncUpdatesNotMe(ctx context.Context, request *sync.TLSyncUpdatesNotMe) (*tg.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("sync.updatesNotMe - request: %s", request)

	r, err := c.SyncUpdatesNotMe(request)
	if err != nil {
		c.Logger.Errorf("sync.updatesNotMe - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("sync.updatesNotMe - reply: %s", r)
	return r, err
}

// SyncPushUpdates
// sync.pushUpdates user_id:long updates:Updates = Void;
func (s *Service) SyncPushUpdates(ctx context.Context, request *sync.TLSyncPushUpdates) (*tg.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("sync.pushUpdates - request: %s", request)

	r, err := c.SyncPushUpdates(request)
	if err != nil {
		c.Logger.Errorf("sync.pushUpdates - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("sync.pushUpdates - reply: %s", r)
	return r, err
}

// SyncPushUpdatesIfNot
// sync.pushUpdatesIfNot flags:# user_id:long includes:flags.0?Vector<long> excludes:flags.1?Vector<long> updates:Updates = Void;
func (s *Service) SyncPushUpdatesIfNot(ctx context.Context, request *sync.TLSyncPushUpdatesIfNot) (*tg.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("sync.pushUpdatesIfNot - request: %s", request)

	r, err := c.SyncPushUpdatesIfNot(request)
	if err != nil {
		c.Logger.Errorf("sync.pushUpdatesIfNot - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("sync.pushUpdatesIfNot - reply: %s", r)
	return r, err
}

// SyncPushRpcResult
// sync.pushRpcResult user_id:long perm_auth_key_id:long auth_key_id:long gateway_id:string gateway_generation:string gateway_rpc_addr:string session_id:long client_req_msg_id:long rpc_result:bytes = Void;
func (s *Service) SyncPushRpcResult(ctx context.Context, request *sync.TLSyncPushRpcResult) (*tg.Void, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("sync.pushRpcResult - request: %s", request)

	r, err := c.SyncPushRpcResult(request)
	if err != nil {
		c.Logger.Errorf("sync.pushRpcResult - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("sync.pushRpcResult - reply: %s", r)
	return r, err
}
