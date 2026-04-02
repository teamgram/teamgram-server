/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2026 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package messagesclient

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/bff/messages/messages/messagesservice"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
)

type MessagesClient interface {
	MessagesComposeMessageWithAI(ctx context.Context, in *tg.TLMessagesComposeMessageWithAI) (*tg.MessagesComposedMessageWithAI, error)
	MessagesReportReadMetrics(ctx context.Context, in *tg.TLMessagesReportReadMetrics) (*tg.Bool, error)
	MessagesReportMusicListen(ctx context.Context, in *tg.TLMessagesReportMusicListen) (*tg.Bool, error)
	MessagesAddPollAnswer(ctx context.Context, in *tg.TLMessagesAddPollAnswer) (*tg.Updates, error)
	MessagesDeletePollAnswer(ctx context.Context, in *tg.TLMessagesDeletePollAnswer) (*tg.Updates, error)
	MessagesGetUnreadPollVotes(ctx context.Context, in *tg.TLMessagesGetUnreadPollVotes) (*tg.MessagesMessages, error)
	MessagesReadPollVotes(ctx context.Context, in *tg.TLMessagesReadPollVotes) (*tg.MessagesAffectedHistory, error)
	Close() error
}

type defaultMessagesClient struct {
	cli client.Client
}

func NewMessagesClient(cli client.Client) MessagesClient {
	return &defaultMessagesClient{
		cli: cli,
	}
}

func (m *defaultMessagesClient) Close() error {
	if closer, ok := any(m.cli).(interface{ Close() error }); ok {
		return closer.Close()
	}
	return nil
}

// MessagesComposeMessageWithAI
// messages.composeMessageWithAI#fd426afe flags:# proofread:flags.0?true emojify:flags.3?true text:TextWithEntities translate_to_lang:flags.1?string change_tone:flags.2?string = messages.ComposedMessageWithAI;
func (m *defaultMessagesClient) MessagesComposeMessageWithAI(ctx context.Context, in *tg.TLMessagesComposeMessageWithAI) (*tg.MessagesComposedMessageWithAI, error) {
	cli := messagesservice.NewRPCMessagesClient(m.cli)
	return cli.MessagesComposeMessageWithAI(ctx, in)
}

// MessagesReportReadMetrics
// messages.reportReadMetrics#4067c5e6 peer:InputPeer metrics:Vector<InputMessageReadMetric> = Bool;
func (m *defaultMessagesClient) MessagesReportReadMetrics(ctx context.Context, in *tg.TLMessagesReportReadMetrics) (*tg.Bool, error) {
	cli := messagesservice.NewRPCMessagesClient(m.cli)
	return cli.MessagesReportReadMetrics(ctx, in)
}

// MessagesReportMusicListen
// messages.reportMusicListen#ddbcd819 id:InputDocument listened_duration:int = Bool;
func (m *defaultMessagesClient) MessagesReportMusicListen(ctx context.Context, in *tg.TLMessagesReportMusicListen) (*tg.Bool, error) {
	cli := messagesservice.NewRPCMessagesClient(m.cli)
	return cli.MessagesReportMusicListen(ctx, in)
}

// MessagesAddPollAnswer
// messages.addPollAnswer#19bc4b6d peer:InputPeer msg_id:int answer:PollAnswer = Updates;
func (m *defaultMessagesClient) MessagesAddPollAnswer(ctx context.Context, in *tg.TLMessagesAddPollAnswer) (*tg.Updates, error) {
	cli := messagesservice.NewRPCMessagesClient(m.cli)
	return cli.MessagesAddPollAnswer(ctx, in)
}

// MessagesDeletePollAnswer
// messages.deletePollAnswer#ac8505a5 peer:InputPeer msg_id:int option:bytes = Updates;
func (m *defaultMessagesClient) MessagesDeletePollAnswer(ctx context.Context, in *tg.TLMessagesDeletePollAnswer) (*tg.Updates, error) {
	cli := messagesservice.NewRPCMessagesClient(m.cli)
	return cli.MessagesDeletePollAnswer(ctx, in)
}

// MessagesGetUnreadPollVotes
// messages.getUnreadPollVotes#43286cf2 flags:# peer:InputPeer top_msg_id:flags.0?int offset_id:int add_offset:int limit:int max_id:int min_id:int = messages.Messages;
func (m *defaultMessagesClient) MessagesGetUnreadPollVotes(ctx context.Context, in *tg.TLMessagesGetUnreadPollVotes) (*tg.MessagesMessages, error) {
	cli := messagesservice.NewRPCMessagesClient(m.cli)
	return cli.MessagesGetUnreadPollVotes(ctx, in)
}

// MessagesReadPollVotes
// messages.readPollVotes#1720b4d8 flags:# peer:InputPeer top_msg_id:flags.0?int = messages.AffectedHistory;
func (m *defaultMessagesClient) MessagesReadPollVotes(ctx context.Context, in *tg.TLMessagesReadPollVotes) (*tg.MessagesAffectedHistory, error) {
	cli := messagesservice.NewRPCMessagesClient(m.cli)
	return cli.MessagesReadPollVotes(ctx, in)
}
