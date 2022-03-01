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
	"github.com/teamgram/teamgram-server/app/service/biz/banned/banned"
)

// BannedGetBannedByPhoneList
// banned.getBannedByPhoneList phone:string = Vector<string>;
func (c *BannedCore) BannedGetBannedByPhoneList(in *banned.TLBannedGetBannedByPhoneList) (*banned.Vector_String, error) {
	pList, _ := c.svcCtx.BannedDAO.SelectPhoneList(c.ctx, in.GetPhones())
	if pList == nil {
		pList = []string{}
	}

	return &banned.Vector_String{
		Datas: pList,
	}, nil
}
