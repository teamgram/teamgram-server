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
	"encoding/json"
	"github.com/teamgram/marmota/pkg/hack"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserSetPrivacy
// user.setPrivacy user_id:int key_type:int rules:Vector<PrivacyRule> = Bool;
func (c *UserCore) UserSetPrivacy(in *user.TLUserSetPrivacy) (*mtproto.Bool, error) {
	bData, _ := json.Marshal(in.Rules)
	do := &dataobject.UserPrivaciesDO{
		UserId:  in.UserId,
		KeyType: in.KeyType,
		Rules:   hack.String(bData),
	}
	c.svcCtx.Dao.UserPrivaciesDAO.InsertOrUpdate(c.ctx, do)

	return mtproto.BoolTrue, nil
}
