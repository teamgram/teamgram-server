/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/updates/internal/core"
)

// UpdatesGetState
// updates.getState#edd4882a = updates.State;
func (s *Service) UpdatesGetState(ctx context.Context, request *mtproto.TLUpdatesGetState) (*mtproto.Updates_State, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("updates.getState - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UpdatesGetState(request)
	if err != nil {
		return nil, err
	}

	c.Infof("updates.getState - reply: %s", r.DebugString())
	return r, err
}

// UpdatesGetDifference
// updates.getDifference#25939651 flags:# pts:int pts_total_limit:flags.0?int date:int qts:int = updates.Difference;
func (s *Service) UpdatesGetDifference(ctx context.Context, request *mtproto.TLUpdatesGetDifference) (*mtproto.Updates_Difference, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("updates.getDifference - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UpdatesGetDifference(request)
	if err != nil {
		return nil, err
	}

	c.Infof("updates.getDifference - reply: %s", r.DebugString())
	return r, err
}

// UpdatesGetChannelDifference
// updates.getChannelDifference#3173d78 flags:# force:flags.0?true channel:InputChannel filter:ChannelMessagesFilter pts:int limit:int = updates.ChannelDifference;
func (s *Service) UpdatesGetChannelDifference(ctx context.Context, request *mtproto.TLUpdatesGetChannelDifference) (*mtproto.Updates_ChannelDifference, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("updates.getChannelDifference - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UpdatesGetChannelDifference(request)
	if err != nil {
		return nil, err
	}

	c.Infof("updates.getChannelDifference - reply: %s", r.DebugString())
	return r, err
}
