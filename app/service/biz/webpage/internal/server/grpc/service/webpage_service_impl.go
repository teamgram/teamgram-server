/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/webpage/internal/core"
	"github.com/teamgram/teamgram-server/app/service/biz/webpage/webpage"
)

// WebpageGetPendingWebPagePreview
// webpage.getPendingWebPagePreview flags:# message:string entities:flags.3?Vector<MessageEntity> = WebPage;
func (s *Service) WebpageGetPendingWebPagePreview(ctx context.Context, request *webpage.TLWebpageGetPendingWebPagePreview) (*mtproto.WebPage, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("webpage.getPendingWebPagePreview - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.WebpageGetPendingWebPagePreview(request)
	if err != nil {
		return nil, err
	}

	c.Infof("webpage.getPendingWebPagePreview - reply: %s", r.DebugString())
	return r, err
}

// WebpageGetWebPagePreview
// webpage.getWebPagePreview flags:# message:string entities:flags.3?Vector<MessageEntity> = WebPage;
func (s *Service) WebpageGetWebPagePreview(ctx context.Context, request *webpage.TLWebpageGetWebPagePreview) (*mtproto.WebPage, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("webpage.getWebPagePreview - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.WebpageGetWebPagePreview(request)
	if err != nil {
		return nil, err
	}

	c.Infof("webpage.getWebPagePreview - reply: %s", r.DebugString())
	return r, err
}

// WebpageGetWebPage
// webpage.getWebPage url:string hash:int = WebPage;
func (s *Service) WebpageGetWebPage(ctx context.Context, request *webpage.TLWebpageGetWebPage) (*mtproto.WebPage, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("webpage.getWebPage - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.WebpageGetWebPage(request)
	if err != nil {
		return nil, err
	}

	c.Infof("webpage.getWebPage - reply: %s", r.DebugString())
	return r, err
}
