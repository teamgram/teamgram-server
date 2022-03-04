/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package scheduledmessages_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type ScheduledMessagesClient interface {
	MessagesGetScheduledHistory(ctx context.Context, in *mtproto.TLMessagesGetScheduledHistory) (*mtproto.Messages_Messages, error)
	MessagesGetScheduledMessages(ctx context.Context, in *mtproto.TLMessagesGetScheduledMessages) (*mtproto.Messages_Messages, error)
	MessagesSendScheduledMessages(ctx context.Context, in *mtproto.TLMessagesSendScheduledMessages) (*mtproto.Updates, error)
	MessagesDeleteScheduledMessages(ctx context.Context, in *mtproto.TLMessagesDeleteScheduledMessages) (*mtproto.Updates, error)
}

type defaultScheduledMessagesClient struct {
	cli zrpc.Client
}

func NewScheduledMessagesClient(cli zrpc.Client) ScheduledMessagesClient {
	return &defaultScheduledMessagesClient{
		cli: cli,
	}
}

// MessagesGetScheduledHistory
// messages.getScheduledHistory#f516760b peer:InputPeer hash:long = messages.Messages;
func (m *defaultScheduledMessagesClient) MessagesGetScheduledHistory(ctx context.Context, in *mtproto.TLMessagesGetScheduledHistory) (*mtproto.Messages_Messages, error) {
	client := mtproto.NewRPCScheduledMessagesClient(m.cli.Conn())
	return client.MessagesGetScheduledHistory(ctx, in)
}

// MessagesGetScheduledMessages
// messages.getScheduledMessages#bdbb0464 peer:InputPeer id:Vector<int> = messages.Messages;
func (m *defaultScheduledMessagesClient) MessagesGetScheduledMessages(ctx context.Context, in *mtproto.TLMessagesGetScheduledMessages) (*mtproto.Messages_Messages, error) {
	client := mtproto.NewRPCScheduledMessagesClient(m.cli.Conn())
	return client.MessagesGetScheduledMessages(ctx, in)
}

// MessagesSendScheduledMessages
// messages.sendScheduledMessages#bd38850a peer:InputPeer id:Vector<int> = Updates;
func (m *defaultScheduledMessagesClient) MessagesSendScheduledMessages(ctx context.Context, in *mtproto.TLMessagesSendScheduledMessages) (*mtproto.Updates, error) {
	client := mtproto.NewRPCScheduledMessagesClient(m.cli.Conn())
	return client.MessagesSendScheduledMessages(ctx, in)
}

// MessagesDeleteScheduledMessages
// messages.deleteScheduledMessages#59ae2b16 peer:InputPeer id:Vector<int> = Updates;
func (m *defaultScheduledMessagesClient) MessagesDeleteScheduledMessages(ctx context.Context, in *mtproto.TLMessagesDeleteScheduledMessages) (*mtproto.Updates, error) {
	client := mtproto.NewRPCScheduledMessagesClient(m.cli.Conn())
	return client.MessagesDeleteScheduledMessages(ctx, in)
}
