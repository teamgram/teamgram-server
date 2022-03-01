/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package gdpr_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type GdprClient interface {
	AccountInitTakeoutSession(ctx context.Context, in *mtproto.TLAccountInitTakeoutSession) (*mtproto.Account_Takeout, error)
	AccountFinishTakeoutSession(ctx context.Context, in *mtproto.TLAccountFinishTakeoutSession) (*mtproto.Bool, error)
	MessagesGetSplitRanges(ctx context.Context, in *mtproto.TLMessagesGetSplitRanges) (*mtproto.Vector_MessageRange, error)
	ChannelsGetLeftChannels(ctx context.Context, in *mtproto.TLChannelsGetLeftChannels) (*mtproto.Messages_Chats, error)
}

type defaultGdprClient struct {
	cli zrpc.Client
}

func NewGdprClient(cli zrpc.Client) GdprClient {
	return &defaultGdprClient{
		cli: cli,
	}
}

// AccountInitTakeoutSession
// account.initTakeoutSession#f05b4804 flags:# contacts:flags.0?true message_users:flags.1?true message_chats:flags.2?true message_megagroups:flags.3?true message_channels:flags.4?true files:flags.5?true file_max_size:flags.5?int = account.Takeout;
func (m *defaultGdprClient) AccountInitTakeoutSession(ctx context.Context, in *mtproto.TLAccountInitTakeoutSession) (*mtproto.Account_Takeout, error) {
	client := mtproto.NewRPCGdprClient(m.cli.Conn())
	return client.AccountInitTakeoutSession(ctx, in)
}

// AccountFinishTakeoutSession
// account.finishTakeoutSession#1d2652ee flags:# success:flags.0?true = Bool;
func (m *defaultGdprClient) AccountFinishTakeoutSession(ctx context.Context, in *mtproto.TLAccountFinishTakeoutSession) (*mtproto.Bool, error) {
	client := mtproto.NewRPCGdprClient(m.cli.Conn())
	return client.AccountFinishTakeoutSession(ctx, in)
}

// MessagesGetSplitRanges
// messages.getSplitRanges#1cff7e08 = Vector<MessageRange>;
func (m *defaultGdprClient) MessagesGetSplitRanges(ctx context.Context, in *mtproto.TLMessagesGetSplitRanges) (*mtproto.Vector_MessageRange, error) {
	client := mtproto.NewRPCGdprClient(m.cli.Conn())
	return client.MessagesGetSplitRanges(ctx, in)
}

// ChannelsGetLeftChannels
// channels.getLeftChannels#8341ecc0 offset:int = messages.Chats;
func (m *defaultGdprClient) ChannelsGetLeftChannels(ctx context.Context, in *mtproto.TLChannelsGetLeftChannels) (*mtproto.Messages_Chats, error) {
	client := mtproto.NewRPCGdprClient(m.cli.Conn())
	return client.ChannelsGetLeftChannels(ctx, in)
}
