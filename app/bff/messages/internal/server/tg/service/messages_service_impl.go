/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2026 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/bff/messages/internal/core"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesComposeMessageWithAI
// messages.composeMessageWithAI#fd426afe flags:# proofread:flags.0?true emojify:flags.3?true text:TextWithEntities translate_to_lang:flags.1?string change_tone:flags.2?string = messages.ComposedMessageWithAI;
func (s *Service) MessagesComposeMessageWithAI(ctx context.Context, request *tg.TLMessagesComposeMessageWithAI) (*tg.MessagesComposedMessageWithAI, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.composeMessageWithAI - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessagesComposeMessageWithAI(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.composeMessageWithAI - reply: %s", r)
	return r, err
}

// MessagesReportReadMetrics
// messages.reportReadMetrics#4067c5e6 peer:InputPeer metrics:Vector<InputMessageReadMetric> = Bool;
func (s *Service) MessagesReportReadMetrics(ctx context.Context, request *tg.TLMessagesReportReadMetrics) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.reportReadMetrics - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessagesReportReadMetrics(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.reportReadMetrics - reply: %s", r)
	return r, err
}

// MessagesReportMusicListen
// messages.reportMusicListen#ddbcd819 id:InputDocument listened_duration:int = Bool;
func (s *Service) MessagesReportMusicListen(ctx context.Context, request *tg.TLMessagesReportMusicListen) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.reportMusicListen - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessagesReportMusicListen(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.reportMusicListen - reply: %s", r)
	return r, err
}

// MessagesAddPollAnswer
// messages.addPollAnswer#19bc4b6d peer:InputPeer msg_id:int answer:PollAnswer = Updates;
func (s *Service) MessagesAddPollAnswer(ctx context.Context, request *tg.TLMessagesAddPollAnswer) (*tg.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.addPollAnswer - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessagesAddPollAnswer(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.addPollAnswer - reply: %s", r)
	return r, err
}

// MessagesDeletePollAnswer
// messages.deletePollAnswer#ac8505a5 peer:InputPeer msg_id:int option:bytes = Updates;
func (s *Service) MessagesDeletePollAnswer(ctx context.Context, request *tg.TLMessagesDeletePollAnswer) (*tg.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.deletePollAnswer - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessagesDeletePollAnswer(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.deletePollAnswer - reply: %s", r)
	return r, err
}

// MessagesGetUnreadPollVotes
// messages.getUnreadPollVotes#43286cf2 flags:# peer:InputPeer top_msg_id:flags.0?int offset_id:int add_offset:int limit:int max_id:int min_id:int = messages.Messages;
func (s *Service) MessagesGetUnreadPollVotes(ctx context.Context, request *tg.TLMessagesGetUnreadPollVotes) (*tg.MessagesMessages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getUnreadPollVotes - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessagesGetUnreadPollVotes(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getUnreadPollVotes - reply: %s", r)
	return r, err
}

// MessagesReadPollVotes
// messages.readPollVotes#1720b4d8 flags:# peer:InputPeer top_msg_id:flags.0?int = messages.AffectedHistory;
func (s *Service) MessagesReadPollVotes(ctx context.Context, request *tg.TLMessagesReadPollVotes) (*tg.MessagesAffectedHistory, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.readPollVotes - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessagesReadPollVotes(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.readPollVotes - reply: %s", r)
	return r, err
}
