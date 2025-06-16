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
	"github.com/teamgram/teamgram-server/v2/app/bff/updates/updates/updatesservice"

	"github.com/cloudwego/kitex/client"
)

type UpdatesClient interface {
	UpdatesGetState(ctx context.Context, in *tg.TLUpdatesGetState) (*tg.UpdatesState, error)
	UpdatesGetDifference(ctx context.Context, in *tg.TLUpdatesGetDifference) (*tg.UpdatesDifference, error)
	UpdatesGetChannelDifference(ctx context.Context, in *tg.TLUpdatesGetChannelDifference) (*tg.UpdatesChannelDifference, error)
}

type defaultUpdatesClient struct {
	cli client.Client
}

func NewUpdatesClient(cli client.Client) UpdatesClient {
	return &defaultUpdatesClient{
		cli: cli,
	}
}

// UpdatesGetState
// updates.getState#edd4882a = updates.State;
func (m *defaultUpdatesClient) UpdatesGetState(ctx context.Context, in *tg.TLUpdatesGetState) (*tg.UpdatesState, error) {
	cli := updatesservice.NewRPCUpdatesClient(m.cli)
	return cli.UpdatesGetState(ctx, in)
}

// UpdatesGetDifference
// updates.getDifference#19c2f763 flags:# pts:int pts_limit:flags.1?int pts_total_limit:flags.0?int date:int qts:int qts_limit:flags.2?int = updates.Difference;
func (m *defaultUpdatesClient) UpdatesGetDifference(ctx context.Context, in *tg.TLUpdatesGetDifference) (*tg.UpdatesDifference, error) {
	cli := updatesservice.NewRPCUpdatesClient(m.cli)
	return cli.UpdatesGetDifference(ctx, in)
}

// UpdatesGetChannelDifference
// updates.getChannelDifference#3173d78 flags:# force:flags.0?true channel:InputChannel filter:ChannelMessagesFilter pts:int limit:int = updates.ChannelDifference;
func (m *defaultUpdatesClient) UpdatesGetChannelDifference(ctx context.Context, in *tg.TLUpdatesGetChannelDifference) (*tg.UpdatesChannelDifference, error) {
	cli := updatesservice.NewRPCUpdatesClient(m.cli)
	return cli.UpdatesGetChannelDifference(ctx, in)
}
