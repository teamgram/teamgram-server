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
	"github.com/teamgram/teamgram-server/app/bff/translate/internal/core"
)

// MessagesTranslateText
// messages.translateText#24ce6dee flags:# peer:flags.0?InputPeer msg_id:flags.0?int text:flags.1?string from_lang:flags.2?string to_lang:string = messages.TranslatedText;
func (s *Service) MessagesTranslateText(ctx context.Context, request *mtproto.TLMessagesTranslateText) (*mtproto.Messages_TranslatedText, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.translateText - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesTranslateText(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.translateText - reply: %s", r.DebugString())
	return r, err
}
