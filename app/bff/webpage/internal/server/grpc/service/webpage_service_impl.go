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
	"github.com/teamgram/teamgram-server/app/bff/webpage/internal/core"
)

// MessagesGetWebPagePreview
// messages.getWebPagePreview#8b68b0cc flags:# message:string entities:flags.3?Vector<MessageEntity> = MessageMedia;
func (s *Service) MessagesGetWebPagePreview(ctx context.Context, request *mtproto.TLMessagesGetWebPagePreview) (*mtproto.MessageMedia, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getWebPagePreview - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetWebPagePreview(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getWebPagePreview - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetWebPage
// messages.getWebPage#32ca8f91 url:string hash:int = WebPage;
func (s *Service) MessagesGetWebPage(ctx context.Context, request *mtproto.TLMessagesGetWebPage) (*mtproto.WebPage, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getWebPage - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetWebPage(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getWebPage - reply: %s", r.DebugString())
	return r, err
}
