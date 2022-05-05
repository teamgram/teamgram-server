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
	"github.com/teamgram/teamgram-server/app/service/media/media"
)

// MediaGetDocumentList
// media.getDocumentList id_list:Vector<long> = Vector<Document>;
func (c *MediaCore) MediaGetDocumentList(in *media.TLMediaGetDocumentList) (*media.Vector_Document, error) {
	documents := c.svcCtx.Dao.GetDocumentListByIdList(c.ctx, in.IdList)

	return &media.Vector_Document{
		Datas: documents,
	}, nil
}
