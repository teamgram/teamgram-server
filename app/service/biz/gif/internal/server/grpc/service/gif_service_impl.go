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
	"github.com/teamgram/teamgram-server/app/service/biz/gif/gif"
	"github.com/teamgram/teamgram-server/app/service/biz/gif/internal/core"
)

// GifSaveGif
// gif.saveGif user_id:long gif_id:long = Bool;
func (s *Service) GifSaveGif(ctx context.Context, request *gif.TLGifSaveGif) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("gif.saveGif - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.GifSaveGif(request)
	if err != nil {
		return nil, err
	}

	c.Infof("gif.saveGif - reply: %s", r.DebugString())
	return r, err
}

// GifGetSavedGifs
// gif.getSavedGifs user_id:long = Vector<long>;
func (s *Service) GifGetSavedGifs(ctx context.Context, request *gif.TLGifGetSavedGifs) (*gif.Vector_Long, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("gif.getSavedGifs - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.GifGetSavedGifs(request)
	if err != nil {
		return nil, err
	}

	c.Infof("gif.getSavedGifs - reply: %s", r.DebugString())
	return r, err
}

// GifDeleteSavedGif
// gif.deleteSavedGif user_id:long gif_id:long = Bool;
func (s *Service) GifDeleteSavedGif(ctx context.Context, request *gif.TLGifDeleteSavedGif) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("gif.deleteSavedGif - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.GifDeleteSavedGif(request)
	if err != nil {
		return nil, err
	}

	c.Infof("gif.deleteSavedGif - reply: %s", r.DebugString())
	return r, err
}
