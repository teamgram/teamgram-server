/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package sponsoredmessagesclient

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type SponsoredMessagesClient interface {
	AccountToggleSponsoredMessages(ctx context.Context, in *mtproto.TLAccountToggleSponsoredMessages) (*mtproto.Bool, error)
	ChannelsViewSponsoredMessage(ctx context.Context, in *mtproto.TLChannelsViewSponsoredMessage) (*mtproto.Bool, error)
	ChannelsGetSponsoredMessages(ctx context.Context, in *mtproto.TLChannelsGetSponsoredMessages) (*mtproto.Messages_SponsoredMessages, error)
	ChannelsClickSponsoredMessage(ctx context.Context, in *mtproto.TLChannelsClickSponsoredMessage) (*mtproto.Bool, error)
	ChannelsReportSponsoredMessage(ctx context.Context, in *mtproto.TLChannelsReportSponsoredMessage) (*mtproto.Channels_SponsoredMessageReportResult, error)
	ChannelsRestrictSponsoredMessages(ctx context.Context, in *mtproto.TLChannelsRestrictSponsoredMessages) (*mtproto.Updates, error)
}

type defaultSponsoredMessagesClient struct {
	cli zrpc.Client
}

func NewSponsoredMessagesClient(cli zrpc.Client) SponsoredMessagesClient {
	return &defaultSponsoredMessagesClient{
		cli: cli,
	}
}

// AccountToggleSponsoredMessages
// account.toggleSponsoredMessages#b9d9a38d enabled:Bool = Bool;
func (m *defaultSponsoredMessagesClient) AccountToggleSponsoredMessages(ctx context.Context, in *mtproto.TLAccountToggleSponsoredMessages) (*mtproto.Bool, error) {
	client := mtproto.NewRPCSponsoredMessagesClient(m.cli.Conn())
	return client.AccountToggleSponsoredMessages(ctx, in)
}

// ChannelsViewSponsoredMessage
// channels.viewSponsoredMessage#beaedb94 channel:InputChannel random_id:bytes = Bool;
func (m *defaultSponsoredMessagesClient) ChannelsViewSponsoredMessage(ctx context.Context, in *mtproto.TLChannelsViewSponsoredMessage) (*mtproto.Bool, error) {
	client := mtproto.NewRPCSponsoredMessagesClient(m.cli.Conn())
	return client.ChannelsViewSponsoredMessage(ctx, in)
}

// ChannelsGetSponsoredMessages
// channels.getSponsoredMessages#ec210fbf channel:InputChannel = messages.SponsoredMessages;
func (m *defaultSponsoredMessagesClient) ChannelsGetSponsoredMessages(ctx context.Context, in *mtproto.TLChannelsGetSponsoredMessages) (*mtproto.Messages_SponsoredMessages, error) {
	client := mtproto.NewRPCSponsoredMessagesClient(m.cli.Conn())
	return client.ChannelsGetSponsoredMessages(ctx, in)
}

// ChannelsClickSponsoredMessage
// channels.clickSponsoredMessage#18afbc93 channel:InputChannel random_id:bytes = Bool;
func (m *defaultSponsoredMessagesClient) ChannelsClickSponsoredMessage(ctx context.Context, in *mtproto.TLChannelsClickSponsoredMessage) (*mtproto.Bool, error) {
	client := mtproto.NewRPCSponsoredMessagesClient(m.cli.Conn())
	return client.ChannelsClickSponsoredMessage(ctx, in)
}

// ChannelsReportSponsoredMessage
// channels.reportSponsoredMessage#af8ff6b9 channel:InputChannel random_id:bytes option:bytes = channels.SponsoredMessageReportResult;
func (m *defaultSponsoredMessagesClient) ChannelsReportSponsoredMessage(ctx context.Context, in *mtproto.TLChannelsReportSponsoredMessage) (*mtproto.Channels_SponsoredMessageReportResult, error) {
	client := mtproto.NewRPCSponsoredMessagesClient(m.cli.Conn())
	return client.ChannelsReportSponsoredMessage(ctx, in)
}

// ChannelsRestrictSponsoredMessages
// channels.restrictSponsoredMessages#9ae91519 channel:InputChannel restricted:Bool = Updates;
func (m *defaultSponsoredMessagesClient) ChannelsRestrictSponsoredMessages(ctx context.Context, in *mtproto.TLChannelsRestrictSponsoredMessages) (*mtproto.Updates, error) {
	client := mtproto.NewRPCSponsoredMessagesClient(m.cli.Conn())
	return client.ChannelsRestrictSponsoredMessages(ctx, in)
}
