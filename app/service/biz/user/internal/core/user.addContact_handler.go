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
	"time"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserAddContact
// user.addContact user_id:long add_phone_privacy_exception:Bool id:long first_name:string last_name:string phone:string = Bool;
func (c *UserCore) UserAddContact(in *user.TLUserAddContact) (*mtproto.Bool, error) {
	var (
		needCheckMutual = false
		changeMutual    = false
	)

	// 1. Check mutal_contact
	meContact := c.svcCtx.Dao.GetUserContact(c.ctx, in.GetUserId(), in.GetId())
	if meContact == nil {
		needCheckMutual = true
	}

	meDO := &dataobject.UserContactsDO{
		OwnerUserId:      in.UserId,
		ContactUserId:    in.Id,
		ContactPhone:     in.GetPhone(),
		ContactFirstName: in.FirstName,
		ContactLastName:  in.LastName,
		Mutual:           false,
		IsDeleted:        false,
		Date2:            time.Now().Unix(),
	}
	if meContact != nil {
		meDO.Mutual = meContact.GetMutualContact()
	}

	// not contact
	if needCheckMutual {
		mutual := c.svcCtx.Dao.GetUserContact(c.ctx, in.Id, in.UserId)
		if mutual != nil {
			meDO.Mutual = true
			changeMutual = true
		}
	}

	err := c.svcCtx.Dao.PutUserContact(c.ctx, changeMutual, meDO)
	if err != nil {
		c.Logger.Errorf(" - error: %v", err)
		return nil, err
	}

	return mtproto.ToBool(changeMutual), nil
}
