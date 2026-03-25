/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgooo Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/bff/drafts/internal/core"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesSaveDraft
// messages.saveDraft#54ae308e flags:# no_webpage:flags.1?true invert_media:flags.6?true reply_to:flags.4?InputReplyTo peer:InputPeer message:string entities:flags.3?Vector<MessageEntity> media:flags.5?InputMedia effect:flags.7?long suggested_post:flags.8?SuggestedPost = Bool;
func (s *Service) MessagesSaveDraft(ctx context.Context, request *tg.TLMessagesSaveDraft) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.saveDraft - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessagesSaveDraft(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.saveDraft - reply: %s", r)
	return r, err
}

// MessagesGetAllDrafts
// messages.getAllDrafts#6a3f8d65 = Updates;
func (s *Service) MessagesGetAllDrafts(ctx context.Context, request *tg.TLMessagesGetAllDrafts) (*tg.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getAllDrafts - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessagesGetAllDrafts(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getAllDrafts - reply: %s", r)
	return r, err
}

// MessagesClearAllDrafts
// messages.clearAllDrafts#7e58ee9c = Bool;
func (s *Service) MessagesClearAllDrafts(ctx context.Context, request *tg.TLMessagesClearAllDrafts) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.clearAllDrafts - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessagesClearAllDrafts(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.clearAllDrafts - reply: %s", r)
	return r, err
}
