/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package webpage_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/webpage/webpage"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type WebpageClient interface {
	WebpageGetPendingWebPagePreview(ctx context.Context, in *webpage.TLWebpageGetPendingWebPagePreview) (*mtproto.WebPage, error)
	WebpageGetWebPagePreview(ctx context.Context, in *webpage.TLWebpageGetWebPagePreview) (*mtproto.WebPage, error)
	WebpageGetWebPage(ctx context.Context, in *webpage.TLWebpageGetWebPage) (*mtproto.WebPage, error)
}

type defaultWebpageClient struct {
	cli zrpc.Client
}

func NewWebpageClient(cli zrpc.Client) WebpageClient {
	return &defaultWebpageClient{
		cli: cli,
	}
}

// WebpageGetPendingWebPagePreview
// webpage.getPendingWebPagePreview flags:# message:string entities:flags.3?Vector<MessageEntity> = WebPage;
func (m *defaultWebpageClient) WebpageGetPendingWebPagePreview(ctx context.Context, in *webpage.TLWebpageGetPendingWebPagePreview) (*mtproto.WebPage, error) {
	client := webpage.NewRPCWebpageClient(m.cli.Conn())
	return client.WebpageGetPendingWebPagePreview(ctx, in)
}

// WebpageGetWebPagePreview
// webpage.getWebPagePreview flags:# message:string entities:flags.3?Vector<MessageEntity> = WebPage;
func (m *defaultWebpageClient) WebpageGetWebPagePreview(ctx context.Context, in *webpage.TLWebpageGetWebPagePreview) (*mtproto.WebPage, error) {
	client := webpage.NewRPCWebpageClient(m.cli.Conn())
	return client.WebpageGetWebPagePreview(ctx, in)
}

// WebpageGetWebPage
// webpage.getWebPage url:string hash:int = WebPage;
func (m *defaultWebpageClient) WebpageGetWebPage(ctx context.Context, in *webpage.TLWebpageGetWebPage) (*mtproto.WebPage, error) {
	client := webpage.NewRPCWebpageClient(m.cli.Conn())
	return client.WebpageGetWebPage(ctx, in)
}
