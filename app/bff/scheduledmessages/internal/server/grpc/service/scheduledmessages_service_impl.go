/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/scheduledmessages/internal/core"
)

// MessagesGetScheduledHistory
// messages.getScheduledHistory#f516760b peer:InputPeer hash:long = messages.Messages;
func (s *Service) MessagesGetScheduledHistory(ctx context.Context, request *mtproto.TLMessagesGetScheduledHistory) (*mtproto.Messages_Messages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getScheduledHistory - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetScheduledHistory(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getScheduledHistory - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetScheduledMessages
// messages.getScheduledMessages#bdbb0464 peer:InputPeer id:Vector<int> = messages.Messages;
func (s *Service) MessagesGetScheduledMessages(ctx context.Context, request *mtproto.TLMessagesGetScheduledMessages) (*mtproto.Messages_Messages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getScheduledMessages - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetScheduledMessages(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getScheduledMessages - reply: %s", r.DebugString())
	return r, err
}

// MessagesSendScheduledMessages
// messages.sendScheduledMessages#bd38850a peer:InputPeer id:Vector<int> = Updates;
func (s *Service) MessagesSendScheduledMessages(ctx context.Context, request *mtproto.TLMessagesSendScheduledMessages) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.sendScheduledMessages - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesSendScheduledMessages(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.sendScheduledMessages - reply: %s", r.DebugString())
	return r, err
}

// MessagesDeleteScheduledMessages
// messages.deleteScheduledMessages#59ae2b16 peer:InputPeer id:Vector<int> = Updates;
func (s *Service) MessagesDeleteScheduledMessages(ctx context.Context, request *mtproto.TLMessagesDeleteScheduledMessages) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.deleteScheduledMessages - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesDeleteScheduledMessages(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.deleteScheduledMessages - reply: %s", r.DebugString())
	return r, err
}
