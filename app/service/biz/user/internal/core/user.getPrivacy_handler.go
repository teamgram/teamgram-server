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
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserGetPrivacy
// user.getPrivacy user_id:int key_type:int = Vector<PrivacyRule>;
func (c *UserCore) UserGetPrivacy(in *user.TLUserGetPrivacy) (*user.Vector_PrivacyRule, error) {
	var (
		rValues = &user.Vector_PrivacyRule{
			Datas: nil,
		}
	)

	//// TODO(@benqi): check keyType
	//// TODO(@benqi): check bot rule
	//if rules, err = m.Dao.Redis.GetPrivacy(ctx, userId, keyType); err != nil {
	//	cacheError = true
	//} else if rules != nil {
	//	// hit
	//	return
	//}

	// miss or redis error
	if do, err := c.svcCtx.Dao.UserPrivaciesDAO.SelectPrivacy(c.ctx, in.UserId, in.KeyType); err != nil {
		c.Logger.Errorf("user.getPrivacy - error: %v", err)
		return nil, err
	} else if do == nil {
		// TODO(@benqi): make default
		rValues.Datas = append(rValues.Datas, mtproto.MakeTLPrivacyValueAllowAll(nil).To_PrivacyRule())
	} else {
		if err = json.Unmarshal([]byte(do.Rules), &rValues.Datas); err != nil {
			c.Logger.Errorf("user.getPrivacy - Unmarshal PrivacyRulesData(%d)error: %v", do.Id, err)
			return nil, err
		}
	}

	//if !cacheError {
	//	// put cache
	//	m.refreshPrivacyToCache(ctx, userId)
	//}

	return rValues, nil
}
