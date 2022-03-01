/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package gif_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/gif/gif"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type GifClient interface {
	GifSaveGif(ctx context.Context, in *gif.TLGifSaveGif) (*mtproto.Bool, error)
	GifGetSavedGifs(ctx context.Context, in *gif.TLGifGetSavedGifs) (*gif.Vector_Long, error)
	GifDeleteSavedGif(ctx context.Context, in *gif.TLGifDeleteSavedGif) (*mtproto.Bool, error)
}

type defaultGifClient struct {
	cli zrpc.Client
}

func NewGifClient(cli zrpc.Client) GifClient {
	return &defaultGifClient{
		cli: cli,
	}
}

// GifSaveGif
// gif.saveGif user_id:long gif_id:long = Bool;
func (m *defaultGifClient) GifSaveGif(ctx context.Context, in *gif.TLGifSaveGif) (*mtproto.Bool, error) {
	client := gif.NewRPCGifClient(m.cli.Conn())
	return client.GifSaveGif(ctx, in)
}

// GifGetSavedGifs
// gif.getSavedGifs user_id:long = Vector<long>;
func (m *defaultGifClient) GifGetSavedGifs(ctx context.Context, in *gif.TLGifGetSavedGifs) (*gif.Vector_Long, error) {
	client := gif.NewRPCGifClient(m.cli.Conn())
	return client.GifGetSavedGifs(ctx, in)
}

// GifDeleteSavedGif
// gif.deleteSavedGif user_id:long gif_id:long = Bool;
func (m *defaultGifClient) GifDeleteSavedGif(ctx context.Context, in *gif.TLGifDeleteSavedGif) (*mtproto.Bool, error) {
	client := gif.NewRPCGifClient(m.cli.Conn())
	return client.GifDeleteSavedGif(ctx, in)
}
