/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package gifs_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type GifsClient interface {
	MessagesGetSavedGifs(ctx context.Context, in *mtproto.TLMessagesGetSavedGifs) (*mtproto.Messages_SavedGifs, error)
	MessagesSaveGif(ctx context.Context, in *mtproto.TLMessagesSaveGif) (*mtproto.Bool, error)
}

type defaultGifsClient struct {
	cli zrpc.Client
}

func NewGifsClient(cli zrpc.Client) GifsClient {
	return &defaultGifsClient{
		cli: cli,
	}
}

// MessagesGetSavedGifs
// messages.getSavedGifs#5cf09635 hash:long = messages.SavedGifs;
func (m *defaultGifsClient) MessagesGetSavedGifs(ctx context.Context, in *mtproto.TLMessagesGetSavedGifs) (*mtproto.Messages_SavedGifs, error) {
	client := mtproto.NewRPCGifsClient(m.cli.Conn())
	return client.MessagesGetSavedGifs(ctx, in)
}

// MessagesSaveGif
// messages.saveGif#327a30cb id:InputDocument unsave:Bool = Bool;
func (m *defaultGifsClient) MessagesSaveGif(ctx context.Context, in *mtproto.TLMessagesSaveGif) (*mtproto.Bool, error) {
	client := mtproto.NewRPCGifsClient(m.cli.Conn())
	return client.MessagesSaveGif(ctx, in)
}
