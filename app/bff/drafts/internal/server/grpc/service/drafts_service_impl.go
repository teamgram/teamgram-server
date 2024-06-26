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
	"github.com/teamgram/teamgram-server/app/bff/drafts/internal/core"
)

// MessagesSaveDraft
// messages.saveDraft#bc39e14b flags:# no_webpage:flags.1?true reply_to_msg_id:flags.0?int peer:InputPeer message:string entities:flags.3?Vector<MessageEntity> = Bool;
func (s *Service) MessagesSaveDraft(ctx context.Context, request *mtproto.TLMessagesSaveDraft) (*mtproto.Bool, error) {
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
func (s *Service) MessagesGetAllDrafts(ctx context.Context, request *mtproto.TLMessagesGetAllDrafts) (*mtproto.Updates, error) {
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
func (s *Service) MessagesClearAllDrafts(ctx context.Context, request *mtproto.TLMessagesClearAllDrafts) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.clearAllDrafts - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessagesClearAllDrafts(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.clearAllDrafts - reply: %s", r)
	return r, err
}
