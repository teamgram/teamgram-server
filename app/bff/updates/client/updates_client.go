/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package updates_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type UpdatesClient interface {
	UpdatesGetState(ctx context.Context, in *mtproto.TLUpdatesGetState) (*mtproto.Updates_State, error)
	UpdatesGetDifference(ctx context.Context, in *mtproto.TLUpdatesGetDifference) (*mtproto.Updates_Difference, error)
	UpdatesGetChannelDifference(ctx context.Context, in *mtproto.TLUpdatesGetChannelDifference) (*mtproto.Updates_ChannelDifference, error)
}

type defaultUpdatesClient struct {
	cli zrpc.Client
}

func NewUpdatesClient(cli zrpc.Client) UpdatesClient {
	return &defaultUpdatesClient{
		cli: cli,
	}
}

// UpdatesGetState
// updates.getState#edd4882a = updates.State;
func (m *defaultUpdatesClient) UpdatesGetState(ctx context.Context, in *mtproto.TLUpdatesGetState) (*mtproto.Updates_State, error) {
	client := mtproto.NewRPCUpdatesClient(m.cli.Conn())
	return client.UpdatesGetState(ctx, in)
}

// UpdatesGetDifference
// updates.getDifference#25939651 flags:# pts:int pts_total_limit:flags.0?int date:int qts:int = updates.Difference;
func (m *defaultUpdatesClient) UpdatesGetDifference(ctx context.Context, in *mtproto.TLUpdatesGetDifference) (*mtproto.Updates_Difference, error) {
	client := mtproto.NewRPCUpdatesClient(m.cli.Conn())
	return client.UpdatesGetDifference(ctx, in)
}

// UpdatesGetChannelDifference
// updates.getChannelDifference#3173d78 flags:# force:flags.0?true channel:InputChannel filter:ChannelMessagesFilter pts:int limit:int = updates.ChannelDifference;
func (m *defaultUpdatesClient) UpdatesGetChannelDifference(ctx context.Context, in *mtproto.TLUpdatesGetChannelDifference) (*mtproto.Updates_ChannelDifference, error) {
	client := mtproto.NewRPCUpdatesClient(m.cli.Conn())
	return client.UpdatesGetChannelDifference(ctx, in)
}
