/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package draftsclient

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type DraftsClient interface {
	MessagesSaveDraft(ctx context.Context, in *mtproto.TLMessagesSaveDraft) (*mtproto.Bool, error)
	MessagesGetAllDrafts(ctx context.Context, in *mtproto.TLMessagesGetAllDrafts) (*mtproto.Updates, error)
	MessagesClearAllDrafts(ctx context.Context, in *mtproto.TLMessagesClearAllDrafts) (*mtproto.Bool, error)
}

type defaultDraftsClient struct {
	cli zrpc.Client
}

func NewDraftsClient(cli zrpc.Client) DraftsClient {
	return &defaultDraftsClient{
		cli: cli,
	}
}

// MessagesSaveDraft
// messages.saveDraft#d372c5ce flags:# no_webpage:flags.1?true invert_media:flags.6?true reply_to:flags.4?InputReplyTo peer:InputPeer message:string entities:flags.3?Vector<MessageEntity> media:flags.5?InputMedia effect:flags.7?long = Bool;
func (m *defaultDraftsClient) MessagesSaveDraft(ctx context.Context, in *mtproto.TLMessagesSaveDraft) (*mtproto.Bool, error) {
	client := mtproto.NewRPCDraftsClient(m.cli.Conn())
	return client.MessagesSaveDraft(ctx, in)
}

// MessagesGetAllDrafts
// messages.getAllDrafts#6a3f8d65 = Updates;
func (m *defaultDraftsClient) MessagesGetAllDrafts(ctx context.Context, in *mtproto.TLMessagesGetAllDrafts) (*mtproto.Updates, error) {
	client := mtproto.NewRPCDraftsClient(m.cli.Conn())
	return client.MessagesGetAllDrafts(ctx, in)
}

// MessagesClearAllDrafts
// messages.clearAllDrafts#7e58ee9c = Bool;
func (m *defaultDraftsClient) MessagesClearAllDrafts(ctx context.Context, in *mtproto.TLMessagesClearAllDrafts) (*mtproto.Bool, error) {
	client := mtproto.NewRPCDraftsClient(m.cli.Conn())
	return client.MessagesClearAllDrafts(ctx, in)
}
