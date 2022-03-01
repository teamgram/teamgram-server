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

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserAddContact
// user.addContact user_id:long add_phone_privacy_exception:Bool id:long first_name:string last_name:string phone:string = Bool;
func (c *UserCore) UserAddContact(in *user.TLUserAddContact) (*mtproto.Bool, error) {
	// 1. Check mutal_contact
	meDO, err := c.svcCtx.Dao.UserContactsDAO.SelectByContactId(c.ctx, in.UserId, in.Id)
	if err != nil {
		c.Logger.Errorf("user.addContact - error: %v", err)
		return mtproto.BoolFalse, nil
	}

	var (
		needCheckMutual = false
		changeMutual    = false
	)

	if meDO == nil || meDO.IsDeleted {
		needCheckMutual = true
	}

	if meDO == nil {
		meDO = &dataobject.UserContactsDO{
			OwnerUserId:      in.UserId,
			ContactUserId:    in.Id,
			ContactPhone:     "",
			ContactFirstName: in.FirstName,
			ContactLastName:  in.LastName,
			Mutual:           false,
			IsDeleted:        false,
			Date2:            time.Now().Unix(),
		}
	} else {
		meDO.ContactFirstName = in.FirstName
		meDO.ContactLastName = in.LastName
		meDO.IsDeleted = false
	}

	// not contact
	if needCheckMutual {
		contactDO, _ := c.svcCtx.Dao.UserContactsDAO.SelectByContactId(c.ctx, in.Id, in.UserId)
		if contactDO != nil && !contactDO.IsDeleted {
			meDO.Mutual = true
			changeMutual = true
		}
	}

	tR := sqlx.TxWrapper(c.ctx, c.svcCtx.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		if changeMutual {
			c.svcCtx.Dao.UserContactsDAO.UpdateMutualTx(tx, true, in.Id, in.UserId)
		}
		_, _, err = c.svcCtx.Dao.UserContactsDAO.InsertOrUpdateTx(tx, meDO)
		if err != nil {
			result.Err = err
			return
		}

		// TODO(@benqi): set addPhonePrivacyException
		if mtproto.FromBool(in.AddPhonePrivacyException) {

		}
	})

	return mtproto.ToBool(changeMutual), tR.Err
}
