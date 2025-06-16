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
	"github.com/teamgram/teamgram-server/v2/app/bff/updates/internal/core"
)

// UpdatesGetState
// updates.getState#edd4882a = updates.State;
func (s *Service) UpdatesGetState(ctx context.Context, request *tg.TLUpdatesGetState) (*tg.UpdatesState, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("updates.getState - metadata: {}, request: {%v}", request)

	r, err := c.UpdatesGetState(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("updates.getState - reply: {%v}", r)
	return r, err
}

// UpdatesGetDifference
// updates.getDifference#19c2f763 flags:# pts:int pts_limit:flags.1?int pts_total_limit:flags.0?int date:int qts:int qts_limit:flags.2?int = updates.Difference;
func (s *Service) UpdatesGetDifference(ctx context.Context, request *tg.TLUpdatesGetDifference) (*tg.UpdatesDifference, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("updates.getDifference - metadata: {}, request: {%v}", request)

	r, err := c.UpdatesGetDifference(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("updates.getDifference - reply: {%v}", r)
	return r, err
}

// UpdatesGetChannelDifference
// updates.getChannelDifference#3173d78 flags:# force:flags.0?true channel:InputChannel filter:ChannelMessagesFilter pts:int limit:int = updates.ChannelDifference;
func (s *Service) UpdatesGetChannelDifference(ctx context.Context, request *tg.TLUpdatesGetChannelDifference) (*tg.UpdatesChannelDifference, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("updates.getChannelDifference - metadata: {}, request: {%v}", request)

	r, err := c.UpdatesGetChannelDifference(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("updates.getChannelDifference - reply: {%v}", r)
	return r, err
}
