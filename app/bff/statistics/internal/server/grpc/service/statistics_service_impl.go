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
	"github.com/teamgram/teamgram-server/app/bff/statistics/internal/core"
)

// StatsGetBroadcastStats
// stats.getBroadcastStats#ab42441a flags:# dark:flags.0?true channel:InputChannel = stats.BroadcastStats;
func (s *Service) StatsGetBroadcastStats(ctx context.Context, request *mtproto.TLStatsGetBroadcastStats) (*mtproto.Stats_BroadcastStats, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("stats.getBroadcastStats - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.StatsGetBroadcastStats(request)
	if err != nil {
		return nil, err
	}

	c.Infof("stats.getBroadcastStats - reply: %s", r.DebugString())
	return r, err
}

// StatsLoadAsyncGraph
// stats.loadAsyncGraph#621d5fa0 flags:# token:string x:flags.0?long = StatsGraph;
func (s *Service) StatsLoadAsyncGraph(ctx context.Context, request *mtproto.TLStatsLoadAsyncGraph) (*mtproto.StatsGraph, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("stats.loadAsyncGraph - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.StatsLoadAsyncGraph(request)
	if err != nil {
		return nil, err
	}

	c.Infof("stats.loadAsyncGraph - reply: %s", r.DebugString())
	return r, err
}

// StatsGetMegagroupStats
// stats.getMegagroupStats#dcdf8607 flags:# dark:flags.0?true channel:InputChannel = stats.MegagroupStats;
func (s *Service) StatsGetMegagroupStats(ctx context.Context, request *mtproto.TLStatsGetMegagroupStats) (*mtproto.Stats_MegagroupStats, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("stats.getMegagroupStats - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.StatsGetMegagroupStats(request)
	if err != nil {
		return nil, err
	}

	c.Infof("stats.getMegagroupStats - reply: %s", r.DebugString())
	return r, err
}

// StatsGetMessagePublicForwards
// stats.getMessagePublicForwards#5630281b channel:InputChannel msg_id:int offset_rate:int offset_peer:InputPeer offset_id:int limit:int = messages.Messages;
func (s *Service) StatsGetMessagePublicForwards(ctx context.Context, request *mtproto.TLStatsGetMessagePublicForwards) (*mtproto.Messages_Messages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("stats.getMessagePublicForwards - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.StatsGetMessagePublicForwards(request)
	if err != nil {
		return nil, err
	}

	c.Infof("stats.getMessagePublicForwards - reply: %s", r.DebugString())
	return r, err
}

// StatsGetMessageStats
// stats.getMessageStats#b6e0a3f5 flags:# dark:flags.0?true channel:InputChannel msg_id:int = stats.MessageStats;
func (s *Service) StatsGetMessageStats(ctx context.Context, request *mtproto.TLStatsGetMessageStats) (*mtproto.Stats_MessageStats, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("stats.getMessageStats - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.StatsGetMessageStats(request)
	if err != nil {
		return nil, err
	}

	c.Infof("stats.getMessageStats - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetStatsURL
// messages.getStatsURL#812c2ae6 flags:# dark:flags.0?true peer:InputPeer params:string = StatsURL;
func (s *Service) MessagesGetStatsURL(ctx context.Context, request *mtproto.TLMessagesGetStatsURL) (*mtproto.StatsURL, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getStatsURL - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetStatsURL(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getStatsURL - reply: %s", r.DebugString())
	return r, err
}
