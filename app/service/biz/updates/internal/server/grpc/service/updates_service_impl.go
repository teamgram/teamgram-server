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

// UpdatesGetStateV2
// updates.getStateV2 auth_key_id:long user_id:long = updates.State;
func (s *Service) UpdatesGetStateV2(ctx context.Context, request *updates.TLUpdatesGetStateV2) (*mtproto.Updates_State, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("updates.getStateV2 - metadata: %s, request: %s", c.MD, request)

	r, err := c.UpdatesGetStateV2(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("updates.getStateV2 - reply: %s", r)
	return r, err
}

// UpdatesGetDifferenceV2
// updates.getDifferenceV2 flags:# auth_key_id:long user_id:long pts:int pts_total_limit:flags.0?int date:long = Difference;
func (s *Service) UpdatesGetDifferenceV2(ctx context.Context, request *updates.TLUpdatesGetDifferenceV2) (*updates.Difference, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("updates.getDifferenceV2 - metadata: %s, request: %s", c.MD, request)

	r, err := c.UpdatesGetDifferenceV2(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("updates.getDifferenceV2 - reply: %s", r)
	return r, err
}

// UpdatesGetChannelDifferenceV2
// updates.getChannelDifferenceV2 auth_key_id:long user_id:long channel_id:long pts:int limit:int = ChannelDifference;
func (s *Service) UpdatesGetChannelDifferenceV2(ctx context.Context, request *updates.TLUpdatesGetChannelDifferenceV2) (*updates.ChannelDifference, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("updates.getChannelDifferenceV2 - metadata: %s, request: %s", c.MD, request)

	r, err := c.UpdatesGetChannelDifferenceV2(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("updates.getChannelDifferenceV2 - reply: %s", r)
	return r, err
}
