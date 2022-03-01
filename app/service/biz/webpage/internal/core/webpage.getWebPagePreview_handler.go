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

// WebpageGetWebPagePreview
// webpage.getWebPagePreview flags:# message:string entities:flags.3?Vector<MessageEntity> = WebPage;
func (c *WebpageCore) WebpageGetWebPagePreview(in *webpage.TLWebpageGetWebPagePreview) (*mtproto.WebPage, error) {
	// TODO: not impl
	// c.Logger.Errorf("webpage.getWebPagePreview - error: method WebpageGetWebPagePreview not impl")

	return emptyWebpageEmpty, nil
}
