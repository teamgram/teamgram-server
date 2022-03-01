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
	"github.com/teamgram/teamgram-server/app/bff/emoji/internal/core"
)

// MessagesGetEmojiKeywords
// messages.getEmojiKeywords#35a0e062 lang_code:string = EmojiKeywordsDifference;
func (s *Service) MessagesGetEmojiKeywords(ctx context.Context, request *mtproto.TLMessagesGetEmojiKeywords) (*mtproto.EmojiKeywordsDifference, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getEmojiKeywords - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetEmojiKeywords(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getEmojiKeywords - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetEmojiKeywordsDifference
// messages.getEmojiKeywordsDifference#1508b6af lang_code:string from_version:int = EmojiKeywordsDifference;
func (s *Service) MessagesGetEmojiKeywordsDifference(ctx context.Context, request *mtproto.TLMessagesGetEmojiKeywordsDifference) (*mtproto.EmojiKeywordsDifference, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getEmojiKeywordsDifference - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetEmojiKeywordsDifference(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getEmojiKeywordsDifference - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetEmojiKeywordsLanguages
// messages.getEmojiKeywordsLanguages#4e9963b2 lang_codes:Vector<string> = Vector<EmojiLanguage>;
func (s *Service) MessagesGetEmojiKeywordsLanguages(ctx context.Context, request *mtproto.TLMessagesGetEmojiKeywordsLanguages) (*mtproto.Vector_EmojiLanguage, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getEmojiKeywordsLanguages - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetEmojiKeywordsLanguages(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getEmojiKeywordsLanguages - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetEmojiURL
// messages.getEmojiURL#d5b10c26 lang_code:string = EmojiURL;
func (s *Service) MessagesGetEmojiURL(ctx context.Context, request *mtproto.TLMessagesGetEmojiURL) (*mtproto.EmojiURL, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getEmojiURL - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetEmojiURL(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getEmojiURL - reply: %s", r.DebugString())
	return r, err
}
