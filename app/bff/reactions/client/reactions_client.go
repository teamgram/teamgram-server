/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package reactions_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type ReactionsClient interface {
	MessagesSendReaction(ctx context.Context, in *mtproto.TLMessagesSendReaction) (*mtproto.Updates, error)
	MessagesGetMessagesReactions(ctx context.Context, in *mtproto.TLMessagesGetMessagesReactions) (*mtproto.Updates, error)
	MessagesGetMessageReactionsList(ctx context.Context, in *mtproto.TLMessagesGetMessageReactionsList) (*mtproto.Messages_MessageReactionsList, error)
	MessagesSetChatAvailableReactions(ctx context.Context, in *mtproto.TLMessagesSetChatAvailableReactions) (*mtproto.Updates, error)
	MessagesGetAvailableReactions(ctx context.Context, in *mtproto.TLMessagesGetAvailableReactions) (*mtproto.Messages_AvailableReactions, error)
	MessagesSetDefaultReaction(ctx context.Context, in *mtproto.TLMessagesSetDefaultReaction) (*mtproto.Bool, error)
	MessagesGetUnreadReactions(ctx context.Context, in *mtproto.TLMessagesGetUnreadReactions) (*mtproto.Messages_Messages, error)
	MessagesReadReactions(ctx context.Context, in *mtproto.TLMessagesReadReactions) (*mtproto.Messages_AffectedHistory, error)
}

type defaultReactionsClient struct {
	cli zrpc.Client
}

func NewReactionsClient(cli zrpc.Client) ReactionsClient {
	return &defaultReactionsClient{
		cli: cli,
	}
}

// MessagesSendReaction
// messages.sendReaction#25690ce4 flags:# big:flags.1?true peer:InputPeer msg_id:int reaction:flags.0?string = Updates;
func (m *defaultReactionsClient) MessagesSendReaction(ctx context.Context, in *mtproto.TLMessagesSendReaction) (*mtproto.Updates, error) {
	client := mtproto.NewRPCReactionsClient(m.cli.Conn())
	return client.MessagesSendReaction(ctx, in)
}

// MessagesGetMessagesReactions
// messages.getMessagesReactions#8bba90e6 peer:InputPeer id:Vector<int> = Updates;
func (m *defaultReactionsClient) MessagesGetMessagesReactions(ctx context.Context, in *mtproto.TLMessagesGetMessagesReactions) (*mtproto.Updates, error) {
	client := mtproto.NewRPCReactionsClient(m.cli.Conn())
	return client.MessagesGetMessagesReactions(ctx, in)
}

// MessagesGetMessageReactionsList
// messages.getMessageReactionsList#e0ee6b77 flags:# peer:InputPeer id:int reaction:flags.0?string offset:flags.1?string limit:int = messages.MessageReactionsList;
func (m *defaultReactionsClient) MessagesGetMessageReactionsList(ctx context.Context, in *mtproto.TLMessagesGetMessageReactionsList) (*mtproto.Messages_MessageReactionsList, error) {
	client := mtproto.NewRPCReactionsClient(m.cli.Conn())
	return client.MessagesGetMessageReactionsList(ctx, in)
}

// MessagesSetChatAvailableReactions
// messages.setChatAvailableReactions#14050ea6 peer:InputPeer available_reactions:Vector<string> = Updates;
func (m *defaultReactionsClient) MessagesSetChatAvailableReactions(ctx context.Context, in *mtproto.TLMessagesSetChatAvailableReactions) (*mtproto.Updates, error) {
	client := mtproto.NewRPCReactionsClient(m.cli.Conn())
	return client.MessagesSetChatAvailableReactions(ctx, in)
}

// MessagesGetAvailableReactions
// messages.getAvailableReactions#18dea0ac hash:int = messages.AvailableReactions;
func (m *defaultReactionsClient) MessagesGetAvailableReactions(ctx context.Context, in *mtproto.TLMessagesGetAvailableReactions) (*mtproto.Messages_AvailableReactions, error) {
	client := mtproto.NewRPCReactionsClient(m.cli.Conn())
	return client.MessagesGetAvailableReactions(ctx, in)
}

// MessagesSetDefaultReaction
// messages.setDefaultReaction#d960c4d4 reaction:string = Bool;
func (m *defaultReactionsClient) MessagesSetDefaultReaction(ctx context.Context, in *mtproto.TLMessagesSetDefaultReaction) (*mtproto.Bool, error) {
	client := mtproto.NewRPCReactionsClient(m.cli.Conn())
	return client.MessagesSetDefaultReaction(ctx, in)
}

// MessagesGetUnreadReactions
// messages.getUnreadReactions#e85bae1a peer:InputPeer offset_id:int add_offset:int limit:int max_id:int min_id:int = messages.Messages;
func (m *defaultReactionsClient) MessagesGetUnreadReactions(ctx context.Context, in *mtproto.TLMessagesGetUnreadReactions) (*mtproto.Messages_Messages, error) {
	client := mtproto.NewRPCReactionsClient(m.cli.Conn())
	return client.MessagesGetUnreadReactions(ctx, in)
}

// MessagesReadReactions
// messages.readReactions#82e251d7 peer:InputPeer = messages.AffectedHistory;
func (m *defaultReactionsClient) MessagesReadReactions(ctx context.Context, in *mtproto.TLMessagesReadReactions) (*mtproto.Messages_AffectedHistory, error) {
	client := mtproto.NewRPCReactionsClient(m.cli.Conn())
	return client.MessagesReadReactions(ctx, in)
}
