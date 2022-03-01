/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package sponsoredmessages_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type SponsoredMessagesClient interface {
	ChannelsViewSponsoredMessage(ctx context.Context, in *mtproto.TLChannelsViewSponsoredMessage) (*mtproto.Bool, error)
	ChannelsGetSponsoredMessages(ctx context.Context, in *mtproto.TLChannelsGetSponsoredMessages) (*mtproto.Messages_SponsoredMessages, error)
}

type defaultSponsoredMessagesClient struct {
	cli zrpc.Client
}

func NewSponsoredMessagesClient(cli zrpc.Client) SponsoredMessagesClient {
	return &defaultSponsoredMessagesClient{
		cli: cli,
	}
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
