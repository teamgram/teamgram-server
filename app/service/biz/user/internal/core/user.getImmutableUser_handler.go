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
	"github.com/zeromicro/go-zero/core/mr"
)

// UserGetImmutableUser
// user.getImmutableUser id:long = ImmutableUser;
func (c *UserCore) UserGetImmutableUser(in *user.TLUserGetImmutableUser) (*user.ImmutableUser, error) {
	imUser, err := c.getImmutableUser(in.GetId(), false, false)
	if err != nil {
		return nil, err
	}

	return imUser, nil
}

func (c *UserCore) getImmutableUser(id int64, contacts, privacy bool) (*user.ImmutableUser, error) {
	userData := c.svcCtx.Dao.GetUserData(c.ctx, id)

	// userDO, _ := c.svcCtx.Dao.UsersDAO.SelectById(c.ctx, in.Id)
	if userData == nil {
		err := mtproto.ErrUserIdInvalid
		c.Logger.Errorf("user.getImmutableUser - error: %v", err)
		return nil, err
	}

	immutableUser := user.MakeTLImmutableUser(&user.ImmutableUser{
		User:             userData,
		LastSeenAt:       0,
		Contacts:         nil,
		KeysPrivacyRules: nil,
	}).To_ImmutableUser()

	if !userData.Deleted {
		if int(userData.UserType) == user.UserTypeRegular {
			mr.FinishVoid(
				func() {
					lastSeenAt, _ := c.svcCtx.Dao.GetLastSeenAt(c.ctx, id)
					if lastSeenAt != nil {
						immutableUser.LastSeenAt = lastSeenAt.LastSeenAt
					}
				},
				func() {
					if contacts {
						// TODO: aaa
						immutableUser.Contacts = c.svcCtx.Dao.GetUserContactList(c.ctx, id)
					}
				},
				func() {
					if privacy {
						immutableUser.KeysPrivacyRules = c.svcCtx.Dao.GetUserPrivacyRulesListByKeys(
							c.ctx,
							id,
							user.STATUS_TIMESTAMP,
							user.PROFILE_PHOTO,
							user.PHONE_NUMBER)
					}
				})
		}
	}

	return immutableUser, nil
}
