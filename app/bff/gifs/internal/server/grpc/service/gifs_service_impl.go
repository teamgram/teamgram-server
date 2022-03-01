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
	"github.com/teamgram/teamgram-server/app/bff/gifs/internal/core"
)

// MessagesGetSavedGifs
// messages.getSavedGifs#5cf09635 hash:long = messages.SavedGifs;
func (s *Service) MessagesGetSavedGifs(ctx context.Context, request *mtproto.TLMessagesGetSavedGifs) (*mtproto.Messages_SavedGifs, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getSavedGifs - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetSavedGifs(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getSavedGifs - reply: %s", r.DebugString())
	return r, err
}

// MessagesSaveGif
// messages.saveGif#327a30cb id:InputDocument unsave:Bool = Bool;
func (s *Service) MessagesSaveGif(ctx context.Context, request *mtproto.TLMessagesSaveGif) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.saveGif - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesSaveGif(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.saveGif - reply: %s", r.DebugString())
	return r, err
}
