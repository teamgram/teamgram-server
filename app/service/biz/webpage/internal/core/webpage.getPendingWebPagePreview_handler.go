/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/webpage/webpage"
)

// WebpageGetPendingWebPagePreview
// webpage.getPendingWebPagePreview flags:# message:string entities:flags.3?Vector<MessageEntity> = WebPage;
func (c *WebpageCore) WebpageGetPendingWebPagePreview(in *webpage.TLWebpageGetPendingWebPagePreview) (*mtproto.WebPage, error) {
	// TODO: not impl
	// c.Logger.Errorf("webpage.getPendingWebPagePreview - error: method WebpageGetPendingWebPagePreview not impl")

	return emptyWebpageEmpty, nil
}
