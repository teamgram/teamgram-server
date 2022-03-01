/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package emoji_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type EmojiClient interface {
	MessagesGetEmojiKeywords(ctx context.Context, in *mtproto.TLMessagesGetEmojiKeywords) (*mtproto.EmojiKeywordsDifference, error)
	MessagesGetEmojiKeywordsDifference(ctx context.Context, in *mtproto.TLMessagesGetEmojiKeywordsDifference) (*mtproto.EmojiKeywordsDifference, error)
	MessagesGetEmojiKeywordsLanguages(ctx context.Context, in *mtproto.TLMessagesGetEmojiKeywordsLanguages) (*mtproto.Vector_EmojiLanguage, error)
	MessagesGetEmojiURL(ctx context.Context, in *mtproto.TLMessagesGetEmojiURL) (*mtproto.EmojiURL, error)
}

type defaultEmojiClient struct {
	cli zrpc.Client
}

func NewEmojiClient(cli zrpc.Client) EmojiClient {
	return &defaultEmojiClient{
		cli: cli,
	}
}

// MessagesGetEmojiKeywords
// messages.getEmojiKeywords#35a0e062 lang_code:string = EmojiKeywordsDifference;
func (m *defaultEmojiClient) MessagesGetEmojiKeywords(ctx context.Context, in *mtproto.TLMessagesGetEmojiKeywords) (*mtproto.EmojiKeywordsDifference, error) {
	client := mtproto.NewRPCEmojiClient(m.cli.Conn())
	return client.MessagesGetEmojiKeywords(ctx, in)
}

// MessagesGetEmojiKeywordsDifference
// messages.getEmojiKeywordsDifference#1508b6af lang_code:string from_version:int = EmojiKeywordsDifference;
func (m *defaultEmojiClient) MessagesGetEmojiKeywordsDifference(ctx context.Context, in *mtproto.TLMessagesGetEmojiKeywordsDifference) (*mtproto.EmojiKeywordsDifference, error) {
	client := mtproto.NewRPCEmojiClient(m.cli.Conn())
	return client.MessagesGetEmojiKeywordsDifference(ctx, in)
}

// MessagesGetEmojiKeywordsLanguages
// messages.getEmojiKeywordsLanguages#4e9963b2 lang_codes:Vector<string> = Vector<EmojiLanguage>;
func (m *defaultEmojiClient) MessagesGetEmojiKeywordsLanguages(ctx context.Context, in *mtproto.TLMessagesGetEmojiKeywordsLanguages) (*mtproto.Vector_EmojiLanguage, error) {
	client := mtproto.NewRPCEmojiClient(m.cli.Conn())
	return client.MessagesGetEmojiKeywordsLanguages(ctx, in)
}

// MessagesGetEmojiURL
// messages.getEmojiURL#d5b10c26 lang_code:string = EmojiURL;
func (m *defaultEmojiClient) MessagesGetEmojiURL(ctx context.Context, in *mtproto.TLMessagesGetEmojiURL) (*mtproto.EmojiURL, error) {
	client := mtproto.NewRPCEmojiClient(m.cli.Conn())
	return client.MessagesGetEmojiURL(ctx, in)
}
