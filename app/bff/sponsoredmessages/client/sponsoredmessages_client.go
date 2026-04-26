/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package sponsoredmessagesclient

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/bff/sponsoredmessages/sponsoredmessages/sponsoredmessagesservice"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
)

type SponsoredMessagesClient interface {
	AccountToggleSponsoredMessages(ctx context.Context, in *tg.TLAccountToggleSponsoredMessages) (*tg.Bool, error)
	ContactsGetSponsoredPeers(ctx context.Context, in *tg.TLContactsGetSponsoredPeers) (*tg.ContactsSponsoredPeers, error)
	MessagesViewSponsoredMessage(ctx context.Context, in *tg.TLMessagesViewSponsoredMessage) (*tg.Bool, error)
	MessagesClickSponsoredMessage(ctx context.Context, in *tg.TLMessagesClickSponsoredMessage) (*tg.Bool, error)
	MessagesReportSponsoredMessage(ctx context.Context, in *tg.TLMessagesReportSponsoredMessage) (*tg.ChannelsSponsoredMessageReportResult, error)
	MessagesGetSponsoredMessages(ctx context.Context, in *tg.TLMessagesGetSponsoredMessages) (*tg.MessagesSponsoredMessages, error)
	ChannelsRestrictSponsoredMessages(ctx context.Context, in *tg.TLChannelsRestrictSponsoredMessages) (*tg.Updates, error)
}

type defaultSponsoredMessagesClient struct {
	cli client.Client
	rpc sponsoredmessagesservice.Client
}

func NewSponsoredMessagesClient(cli client.Client) SponsoredMessagesClient {
	return &defaultSponsoredMessagesClient{
		cli: cli,
		rpc: sponsoredmessagesservice.NewRPCSponsoredMessagesClient(cli),
	}
}

// AccountToggleSponsoredMessages
// account.toggleSponsoredMessages#b9d9a38d enabled:Bool = Bool;
func (m *defaultSponsoredMessagesClient) AccountToggleSponsoredMessages(ctx context.Context, in *tg.TLAccountToggleSponsoredMessages) (*tg.Bool, error) {
	return m.rpc.AccountToggleSponsoredMessages(ctx, in)
}

// ContactsGetSponsoredPeers
// contacts.getSponsoredPeers#b6c8c393 q:string = contacts.SponsoredPeers;
func (m *defaultSponsoredMessagesClient) ContactsGetSponsoredPeers(ctx context.Context, in *tg.TLContactsGetSponsoredPeers) (*tg.ContactsSponsoredPeers, error) {
	return m.rpc.ContactsGetSponsoredPeers(ctx, in)
}

// MessagesViewSponsoredMessage
// messages.viewSponsoredMessage#269e3643 random_id:bytes = Bool;
func (m *defaultSponsoredMessagesClient) MessagesViewSponsoredMessage(ctx context.Context, in *tg.TLMessagesViewSponsoredMessage) (*tg.Bool, error) {
	return m.rpc.MessagesViewSponsoredMessage(ctx, in)
}

// MessagesClickSponsoredMessage
// messages.clickSponsoredMessage#8235057e flags:# media:flags.0?true fullscreen:flags.1?true random_id:bytes = Bool;
func (m *defaultSponsoredMessagesClient) MessagesClickSponsoredMessage(ctx context.Context, in *tg.TLMessagesClickSponsoredMessage) (*tg.Bool, error) {
	return m.rpc.MessagesClickSponsoredMessage(ctx, in)
}

// MessagesReportSponsoredMessage
// messages.reportSponsoredMessage#12cbf0c4 random_id:bytes option:bytes = channels.SponsoredMessageReportResult;
func (m *defaultSponsoredMessagesClient) MessagesReportSponsoredMessage(ctx context.Context, in *tg.TLMessagesReportSponsoredMessage) (*tg.ChannelsSponsoredMessageReportResult, error) {
	return m.rpc.MessagesReportSponsoredMessage(ctx, in)
}

// MessagesGetSponsoredMessages
// messages.getSponsoredMessages#3d6ce850 flags:# peer:InputPeer msg_id:flags.0?int = messages.SponsoredMessages;
func (m *defaultSponsoredMessagesClient) MessagesGetSponsoredMessages(ctx context.Context, in *tg.TLMessagesGetSponsoredMessages) (*tg.MessagesSponsoredMessages, error) {
	return m.rpc.MessagesGetSponsoredMessages(ctx, in)
}

// ChannelsRestrictSponsoredMessages
// channels.restrictSponsoredMessages#9ae91519 channel:InputChannel restricted:Bool = Updates;
func (m *defaultSponsoredMessagesClient) ChannelsRestrictSponsoredMessages(ctx context.Context, in *tg.TLChannelsRestrictSponsoredMessages) (*tg.Updates, error) {
	return m.rpc.ChannelsRestrictSponsoredMessages(ctx, in)
}
