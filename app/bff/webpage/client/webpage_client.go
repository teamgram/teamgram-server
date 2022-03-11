/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package webpage_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type WebPageClient interface {
	MessagesGetWebPagePreview(ctx context.Context, in *mtproto.TLMessagesGetWebPagePreview) (*mtproto.MessageMedia, error)
	MessagesGetWebPage(ctx context.Context, in *mtproto.TLMessagesGetWebPage) (*mtproto.WebPage, error)
}

type defaultWebPageClient struct {
	cli zrpc.Client
}

func NewWebPageClient(cli zrpc.Client) WebPageClient {
	return &defaultWebPageClient{
		cli: cli,
	}
}

// MessagesGetWebPagePreview
// messages.getWebPagePreview#8b68b0cc flags:# message:string entities:flags.3?Vector<MessageEntity> = MessageMedia;
func (m *defaultWebPageClient) MessagesGetWebPagePreview(ctx context.Context, in *mtproto.TLMessagesGetWebPagePreview) (*mtproto.MessageMedia, error) {
	client := mtproto.NewRPCWebPageClient(m.cli.Conn())
	return client.MessagesGetWebPagePreview(ctx, in)
}

// MessagesGetWebPage
// messages.getWebPage#32ca8f91 url:string hash:int = WebPage;
func (m *defaultWebPageClient) MessagesGetWebPage(ctx context.Context, in *mtproto.TLMessagesGetWebPage) (*mtproto.WebPage, error) {
	client := mtproto.NewRPCWebPageClient(m.cli.Conn())
	return client.MessagesGetWebPage(ctx, in)
}
