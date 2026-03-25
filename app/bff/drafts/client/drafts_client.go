/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgooo Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package draftsclient

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/bff/drafts/drafts/draftsservice"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
)

type DraftsClient interface {
	MessagesSaveDraft(ctx context.Context, in *tg.TLMessagesSaveDraft) (*tg.Bool, error)
	MessagesGetAllDrafts(ctx context.Context, in *tg.TLMessagesGetAllDrafts) (*tg.Updates, error)
	MessagesClearAllDrafts(ctx context.Context, in *tg.TLMessagesClearAllDrafts) (*tg.Bool, error)
}

type defaultDraftsClient struct {
	cli client.Client
}

func NewDraftsClient(cli client.Client) DraftsClient {
	return &defaultDraftsClient{
		cli: cli,
	}
}

// MessagesSaveDraft
// messages.saveDraft#54ae308e flags:# no_webpage:flags.1?true invert_media:flags.6?true reply_to:flags.4?InputReplyTo peer:InputPeer message:string entities:flags.3?Vector<MessageEntity> media:flags.5?InputMedia effect:flags.7?long suggested_post:flags.8?SuggestedPost = Bool;
func (m *defaultDraftsClient) MessagesSaveDraft(ctx context.Context, in *tg.TLMessagesSaveDraft) (*tg.Bool, error) {
	cli := draftsservice.NewRPCDraftsClient(m.cli)
	return cli.MessagesSaveDraft(ctx, in)
}

// MessagesGetAllDrafts
// messages.getAllDrafts#6a3f8d65 = Updates;
func (m *defaultDraftsClient) MessagesGetAllDrafts(ctx context.Context, in *tg.TLMessagesGetAllDrafts) (*tg.Updates, error) {
	cli := draftsservice.NewRPCDraftsClient(m.cli)
	return cli.MessagesGetAllDrafts(ctx, in)
}

// MessagesClearAllDrafts
// messages.clearAllDrafts#7e58ee9c = Bool;
func (m *defaultDraftsClient) MessagesClearAllDrafts(ctx context.Context, in *tg.TLMessagesClearAllDrafts) (*tg.Bool, error) {
	cli := draftsservice.NewRPCDraftsClient(m.cli)
	return cli.MessagesClearAllDrafts(ctx, in)
}
