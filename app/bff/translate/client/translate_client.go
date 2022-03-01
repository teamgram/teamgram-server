/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package translate_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type TranslateClient interface {
	MessagesTranslateText(ctx context.Context, in *mtproto.TLMessagesTranslateText) (*mtproto.Messages_TranslatedText, error)
}

type defaultTranslateClient struct {
	cli zrpc.Client
}

func NewTranslateClient(cli zrpc.Client) TranslateClient {
	return &defaultTranslateClient{
		cli: cli,
	}
}

// MessagesTranslateText
// messages.translateText#24ce6dee flags:# peer:flags.0?InputPeer msg_id:flags.0?int text:flags.1?string from_lang:flags.2?string to_lang:string = messages.TranslatedText;
func (m *defaultTranslateClient) MessagesTranslateText(ctx context.Context, in *mtproto.TLMessagesTranslateText) (*mtproto.Messages_TranslatedText, error) {
	client := mtproto.NewRPCTranslateClient(m.cli.Conn())
	return client.MessagesTranslateText(ctx, in)
}
