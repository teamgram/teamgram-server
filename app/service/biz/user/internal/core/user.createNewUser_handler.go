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
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserCreateNewUser
// user.createNewUser secret_key_id:long phone:string country_code:string first_name:string last_name:string = ImmutableUser;
func (c *UserCore) UserCreateNewUser(in *user.TLUserCreateNewUser) (*mtproto.ImmutableUser, error) {
	user, err := c.svcCtx.Dao.CreateNewUserV2(
		c.ctx,
		in.SecretKeyId,
		in.Phone,
		in.CountryCode,
		in.FirstName,
		in.LastName)
	if err != nil {
		c.Logger.Errorf("user.createNewUser - error: %v", err)
		return nil, err
	}

	return user, nil
}
