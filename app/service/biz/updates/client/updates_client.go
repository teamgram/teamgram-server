/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package updatesclient

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/updates/updates"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/updates/updates/updatesservice"

	"github.com/cloudwego/kitex/client"
)

var _ *tg.Bool

type UpdatesClient interface {
	UpdatesGetStateV2(ctx context.Context, in *updates.TLUpdatesGetStateV2) (*tg.UpdatesState, error)
	UpdatesGetDifferenceV2(ctx context.Context, in *updates.TLUpdatesGetDifferenceV2) (*updates.Difference, error)
	UpdatesGetChannelDifferenceV2(ctx context.Context, in *updates.TLUpdatesGetChannelDifferenceV2) (*updates.ChannelDifference, error)
}

type defaultUpdatesClient struct {
	cli client.Client
}

func NewUpdatesClient(cli client.Client) UpdatesClient {
	return &defaultUpdatesClient{
		cli: cli,
	}
}

// UpdatesGetStateV2
// updates.getStateV2 auth_key_id:long user_id:long = updates.State;
func (m *defaultUpdatesClient) UpdatesGetStateV2(ctx context.Context, in *updates.TLUpdatesGetStateV2) (*tg.UpdatesState, error) {
	cli := updatesservice.NewRPCUpdatesClient(m.cli)
	return cli.UpdatesGetStateV2(ctx, in)
}

// UpdatesGetDifferenceV2
// updates.getDifferenceV2 flags:# auth_key_id:long user_id:long pts:int pts_total_limit:flags.0?int date:long = Difference;
func (m *defaultUpdatesClient) UpdatesGetDifferenceV2(ctx context.Context, in *updates.TLUpdatesGetDifferenceV2) (*updates.Difference, error) {
	cli := updatesservice.NewRPCUpdatesClient(m.cli)
	return cli.UpdatesGetDifferenceV2(ctx, in)
}

// UpdatesGetChannelDifferenceV2
// updates.getChannelDifferenceV2 auth_key_id:long user_id:long channel_id:long pts:int limit:int = ChannelDifference;
func (m *defaultUpdatesClient) UpdatesGetChannelDifferenceV2(ctx context.Context, in *updates.TLUpdatesGetChannelDifferenceV2) (*updates.ChannelDifference, error) {
	cli := updatesservice.NewRPCUpdatesClient(m.cli)
	return cli.UpdatesGetChannelDifferenceV2(ctx, in)
}
