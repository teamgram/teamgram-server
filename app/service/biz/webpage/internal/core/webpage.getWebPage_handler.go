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

// WebpageGetWebPage
// webpage.getWebPage url:string hash:int = WebPage;
func (c *WebpageCore) WebpageGetWebPage(in *webpage.TLWebpageGetWebPage) (*mtproto.WebPage, error) {
	// TODO: not impl
	// c.Logger.Errorf("webpage.getWebPage - error: method WebpageGetWebPage not impl")

	return emptyWebpageEmpty, nil
}
