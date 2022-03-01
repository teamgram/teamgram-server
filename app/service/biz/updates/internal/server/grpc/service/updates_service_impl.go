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
	"github.com/teamgram/teamgram-server/app/service/biz/updates/internal/core"
	"github.com/teamgram/teamgram-server/app/service/biz/updates/updates"
)

// UpdatesGetState
// updates.getState auth_key_id:long user_id:long = updates.State;
func (s *Service) UpdatesGetState(ctx context.Context, request *updates.TLUpdatesGetState) (*mtproto.Updates_State, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("updates.getState - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UpdatesGetState(request)
	if err != nil {
		return nil, err
	}

	c.Infof("updates.getState - reply: %s", r.DebugString())
	return r, err
}

// UpdatesGetDifferenceV2
// updates.getDifferenceV2 flags:# auth_key_id:long user_id:long pts:int pts_total_limit:flags.0?int date:long = Difference;
func (s *Service) UpdatesGetDifferenceV2(ctx context.Context, request *updates.TLUpdatesGetDifferenceV2) (*updates.Difference, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("updates.getDifferenceV2 - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UpdatesGetDifferenceV2(request)
	if err != nil {
		return nil, err
	}

	c.Infof("updates.getDifferenceV2 - reply: %s", r.DebugString())
	return r, err
}

// UpdatesGetChannelDifferenceV2
// updates.getChannelDifferenceV2 auth_key_id:long user_id:long channel_id:long pts:int limit:int = ChannelDifference;
func (s *Service) UpdatesGetChannelDifferenceV2(ctx context.Context, request *updates.TLUpdatesGetChannelDifferenceV2) (*updates.ChannelDifference, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("updates.getChannelDifferenceV2 - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UpdatesGetChannelDifferenceV2(request)
	if err != nil {
		return nil, err
	}

	c.Infof("updates.getChannelDifferenceV2 - reply: %s", r.DebugString())
	return r, err
}
