/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package updates_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/updates/updates"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type UpdatesClient interface {
	UpdatesGetStateV2(ctx context.Context, in *updates.TLUpdatesGetStateV2) (*mtproto.Updates_State, error)
	UpdatesGetDifferenceV2(ctx context.Context, in *updates.TLUpdatesGetDifferenceV2) (*updates.Difference, error)
	UpdatesGetChannelDifferenceV2(ctx context.Context, in *updates.TLUpdatesGetChannelDifferenceV2) (*updates.ChannelDifference, error)
}

type defaultUpdatesClient struct {
	cli zrpc.Client
}

func NewUpdatesClient(cli zrpc.Client) UpdatesClient {
	return &defaultUpdatesClient{
		cli: cli,
	}
}

// UpdatesGetStateV2
// updates.getStateV2 auth_key_id:long user_id:long = updates.State;
func (m *defaultUpdatesClient) UpdatesGetStateV2(ctx context.Context, in *updates.TLUpdatesGetStateV2) (*mtproto.Updates_State, error) {
	client := updates.NewRPCUpdatesClient(m.cli.Conn())
	return client.UpdatesGetStateV2(ctx, in)
}

// UpdatesGetDifferenceV2
// updates.getDifferenceV2 flags:# auth_key_id:long user_id:long pts:int pts_total_limit:flags.0?int date:long = Difference;
func (m *defaultUpdatesClient) UpdatesGetDifferenceV2(ctx context.Context, in *updates.TLUpdatesGetDifferenceV2) (*updates.Difference, error) {
	client := updates.NewRPCUpdatesClient(m.cli.Conn())
	return client.UpdatesGetDifferenceV2(ctx, in)
}

// UpdatesGetChannelDifferenceV2
// updates.getChannelDifferenceV2 auth_key_id:long user_id:long channel_id:long pts:int limit:int = ChannelDifference;
func (m *defaultUpdatesClient) UpdatesGetChannelDifferenceV2(ctx context.Context, in *updates.TLUpdatesGetChannelDifferenceV2) (*updates.ChannelDifference, error) {
	client := updates.NewRPCUpdatesClient(m.cli.Conn())
	return client.UpdatesGetChannelDifferenceV2(ctx, in)
}
