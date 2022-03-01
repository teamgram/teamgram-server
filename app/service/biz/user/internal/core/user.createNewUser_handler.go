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
	"math/rand"
	"time"

	"github.com/teamgram/marmota/pkg/hack"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserCreateNewUser
// user.createNewUser secret_key_id:long phone:string country_code:string first_name:string last_name:string = ImmutableUser;
func (c *UserCore) UserCreateNewUser(in *user.TLUserCreateNewUser) (*user.ImmutableUser, error) {
	var (
		err          error
		userDO       *dataobject.UsersDO
		now          = time.Now().Unix()
		defaultRules = []*mtproto.PrivacyRule{
			mtproto.MakeTLPrivacyValueAllowAll(nil).To_PrivacyRule(),
		}

		phoneNumberRules = []*mtproto.PrivacyRule{
			mtproto.MakeTLPrivacyValueDisallowAll(nil).To_PrivacyRule(),
		}
	)

	//
	tR := sqlx.TxWrapper(c.ctx, c.svcCtx.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		// var err error
		// user
		userDO = &dataobject.UsersDO{
			UserType:    user.UserTypeRegular,
			AccessHash:  rand.Int63(),
			Phone:       in.Phone,
			SecretKeyId: in.SecretKeyId,
			FirstName:   in.FirstName,
			LastName:    in.LastName,
			CountryCode: in.CountryCode,
			// AccountDaysTtl: 180,
		}
		if lastInsertId, _, err := c.svcCtx.Dao.UsersDAO.InsertTx(tx, userDO); err != nil {
			if sqlx.IsDuplicate(err) {
				result.Err = mtproto.ErrPhoneNumberOccupied
				return
			}
			result.Err = err
			return
		} else {
			userDO.Id = lastInsertId
		}

		// presences
		presencesDO := &dataobject.UserPresencesDO{
			UserId:     userDO.Id,
			LastSeenAt: now,
			Expires:    0,
		}

		if _, _, err := c.svcCtx.Dao.UserPresencesDAO.InsertTx(tx, presencesDO); err != nil {
			result.Err = err
			return
		}

		// account_days_ttl

		// privacy
		bData, _ := json.Marshal(defaultRules)
		bData2, _ := json.Marshal(phoneNumberRules)
		doList := make([]*dataobject.UserPrivaciesDO, 0, user.MAX_KEY_TYPE)
		for i := user.STATUS_TIMESTAMP; i <= user.MAX_KEY_TYPE; i++ {
			if i == user.PHONE_NUMBER {
				doList = append(doList, &dataobject.UserPrivaciesDO{
					Id:      1,
					UserId:  userDO.Id,
					KeyType: int32(i),
					Rules:   hack.String(bData2),
				})
			} else {
				doList = append(doList, &dataobject.UserPrivaciesDO{
					Id:      1,
					UserId:  userDO.Id,
					KeyType: int32(i),
					Rules:   hack.String(bData),
				})
			}
		}

		c.Logger.Infof("doList - %v", doList)
		_, _, err = c.svcCtx.Dao.UserPrivaciesDAO.InsertBulkTx(tx, doList)
		if err != nil {
			result.Err = err
			return
		}

		_, _, err = c.svcCtx.Dao.UserGlobalPrivacySettingsDAO.InsertOrUpdateTx(tx, &dataobject.UserGlobalPrivacySettingsDO{
			UserId:                           userDO.Id,
			ArchiveAndMuteNewNoncontactPeers: false,
		})
		if err != nil {
			result.Err = err
			return
		}
	})

	if tR.Err != nil {
		c.Logger.Errorf("createNewUser2 error: %v", tR.Err)
		return nil, tR.Err
	}

	//// put privacy to cache
	//var privacyList = make(map[int][]*mtproto.PrivacyRule, user.MAX_KEY_TYPE)
	//for i := user.STATUS_TIMESTAMP; i <= user.MAX_KEY_TYPE; i++ {
	//	if i == user.PHONE_NUMBER {
	//		privacyList[i] = phoneNumberRules
	//	} else {
	//		privacyList[i] = defaultRules
	//	}
	//}
	//
	//// user.MakeTLImmutableUser()
	//iUser := user.MakeTLImmutableUser(&user.ImmutableUser{
	//	User:             nil,
	//	LastSeenAt:       0,
	//	Contacts:         nil,
	//	KeysPrivacyRules: nil,
	//}).To_ImmutableUser()
	////makeUserDataByDO(userDO, int32(now)),
	////nil,
	////nil,
	////privacyList)
	//
	////// put to cache
	////err = m.Redis.PutCacheUser2(ctx, user)
	////err = m.Redis.SetPrivacyList(ctx, userDO.Id, privacyList)
	////m.Redis.SetContactList(ctx, userDO.Id)

	return c.UserGetImmutableUser(&user.TLUserGetImmutableUser{
		Id: userDO.Id,
	})
}
