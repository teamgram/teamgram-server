/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package statistics_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type StatisticsClient interface {
	StatsGetBroadcastStats(ctx context.Context, in *mtproto.TLStatsGetBroadcastStats) (*mtproto.Stats_BroadcastStats, error)
	StatsLoadAsyncGraph(ctx context.Context, in *mtproto.TLStatsLoadAsyncGraph) (*mtproto.StatsGraph, error)
	StatsGetMegagroupStats(ctx context.Context, in *mtproto.TLStatsGetMegagroupStats) (*mtproto.Stats_MegagroupStats, error)
	StatsGetMessagePublicForwards(ctx context.Context, in *mtproto.TLStatsGetMessagePublicForwards) (*mtproto.Messages_Messages, error)
	StatsGetMessageStats(ctx context.Context, in *mtproto.TLStatsGetMessageStats) (*mtproto.Stats_MessageStats, error)
	MessagesGetStatsURL(ctx context.Context, in *mtproto.TLMessagesGetStatsURL) (*mtproto.StatsURL, error)
}

type defaultStatisticsClient struct {
	cli zrpc.Client
}

func NewStatisticsClient(cli zrpc.Client) StatisticsClient {
	return &defaultStatisticsClient{
		cli: cli,
	}
}

// StatsGetBroadcastStats
// stats.getBroadcastStats#ab42441a flags:# dark:flags.0?true channel:InputChannel = stats.BroadcastStats;
func (m *defaultStatisticsClient) StatsGetBroadcastStats(ctx context.Context, in *mtproto.TLStatsGetBroadcastStats) (*mtproto.Stats_BroadcastStats, error) {
	client := mtproto.NewRPCStatisticsClient(m.cli.Conn())
	return client.StatsGetBroadcastStats(ctx, in)
}

// StatsLoadAsyncGraph
// stats.loadAsyncGraph#621d5fa0 flags:# token:string x:flags.0?long = StatsGraph;
func (m *defaultStatisticsClient) StatsLoadAsyncGraph(ctx context.Context, in *mtproto.TLStatsLoadAsyncGraph) (*mtproto.StatsGraph, error) {
	client := mtproto.NewRPCStatisticsClient(m.cli.Conn())
	return client.StatsLoadAsyncGraph(ctx, in)
}

// StatsGetMegagroupStats
// stats.getMegagroupStats#dcdf8607 flags:# dark:flags.0?true channel:InputChannel = stats.MegagroupStats;
func (m *defaultStatisticsClient) StatsGetMegagroupStats(ctx context.Context, in *mtproto.TLStatsGetMegagroupStats) (*mtproto.Stats_MegagroupStats, error) {
	client := mtproto.NewRPCStatisticsClient(m.cli.Conn())
	return client.StatsGetMegagroupStats(ctx, in)
}

// StatsGetMessagePublicForwards
// stats.getMessagePublicForwards#5630281b channel:InputChannel msg_id:int offset_rate:int offset_peer:InputPeer offset_id:int limit:int = messages.Messages;
func (m *defaultStatisticsClient) StatsGetMessagePublicForwards(ctx context.Context, in *mtproto.TLStatsGetMessagePublicForwards) (*mtproto.Messages_Messages, error) {
	client := mtproto.NewRPCStatisticsClient(m.cli.Conn())
	return client.StatsGetMessagePublicForwards(ctx, in)
}

// StatsGetMessageStats
// stats.getMessageStats#b6e0a3f5 flags:# dark:flags.0?true channel:InputChannel msg_id:int = stats.MessageStats;
func (m *defaultStatisticsClient) StatsGetMessageStats(ctx context.Context, in *mtproto.TLStatsGetMessageStats) (*mtproto.Stats_MessageStats, error) {
	client := mtproto.NewRPCStatisticsClient(m.cli.Conn())
	return client.StatsGetMessageStats(ctx, in)
}

// MessagesGetStatsURL
// messages.getStatsURL#812c2ae6 flags:# dark:flags.0?true peer:InputPeer params:string = StatsURL;
func (m *defaultStatisticsClient) MessagesGetStatsURL(ctx context.Context, in *mtproto.TLMessagesGetStatsURL) (*mtproto.StatsURL, error) {
	client := mtproto.NewRPCStatisticsClient(m.cli.Conn())
	return client.MessagesGetStatsURL(ctx, in)
}
