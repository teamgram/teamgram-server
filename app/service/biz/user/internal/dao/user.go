// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package dao

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	"github.com/teamgram/marmota/pkg/container2"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"

	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

func (d *Dao) getBotData(ctx context.Context, botId int64) *mtproto.BotData {
	var (
		botData *mtproto.BotData
	)

	botDO, _ := d.BotsDAO.Select(ctx, botId)
	if botDO != nil {
		// userData.Bot
		botData = mtproto.MakeTLBotData(&mtproto.BotData{
			Id:                   botDO.BotId,
			BotType:              botDO.BotType,
			Creator:              botDO.CreatorUserId,
			Token:                botDO.Token,
			Description:          botDO.Description,
			BotChatHistory:       botDO.BotChatHistory,
			BotNochats:           botDO.BotNochats,
			BotInlineGeo:         botDO.BotInlineGeo,
			BotInfoVersion:       botDO.BotInfoVersion,
			BotInlinePlaceholder: mtproto.MakeFlagsString(botDO.BotInlinePlaceholder),
			BotAttachMenu:        false,
			AttachMenuEnabled:    false,
			BotCanEdit:           false,
		}).To_BotData()
	}

	return botData
}

func (d *Dao) CreateNewUserV2(
	ctx context.Context,
	secretKeyId int64,
	phone string,
	countryCode string,
	firstName string, lastName string) (*mtproto.ImmutableUser, error) {
	var (
		//err    error
		userDO        *dataobject.UsersDO
		now           = time.Now().Unix()
		cacheUserData = NewCacheUserData()
	)

	//
	//tR := sqlx.TxWrapper(ctx, d.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
	// var err error
	// user
	userDO = &dataobject.UsersDO{
		UserType:       user.UserTypeRegular,
		AccessHash:     rand.Int63(),
		Phone:          phone,
		SecretKeyId:    secretKeyId,
		FirstName:      firstName,
		LastName:       lastName,
		CountryCode:    countryCode,
		AccountDaysTtl: 180,
	}
	if lastInsertId, _, err2 := d.UsersDAO.Insert(ctx, userDO); err2 != nil {
		if sqlx.IsDuplicate(err2) {
			err2 = mtproto.ErrPhoneNumberOccupied
		}
		return nil, err2
		//result.Err = err2
		//return
	} else {
		userDO.Id = lastInsertId
	}

	cacheUserData.UserData = d.MakeUserDataByDO(userDO)
	cacheUserData.CachesPrivacyKeyRules = append(
		cacheUserData.CachesPrivacyKeyRules,
		mtproto.MakeTLPrivacyKeyRules(&mtproto.PrivacyKeyRules{
			Key:   mtproto.STATUS_TIMESTAMP,
			Rules: defaultRules,
		}).To_PrivacyKeyRules(),
		mtproto.MakeTLPrivacyKeyRules(&mtproto.PrivacyKeyRules{
			Key:   mtproto.PHONE_NUMBER,
			Rules: phoneNumberRules,
		}).To_PrivacyKeyRules(),
		mtproto.MakeTLPrivacyKeyRules(&mtproto.PrivacyKeyRules{
			Key:   mtproto.PROFILE_PHOTO,
			Rules: defaultRules,
		}).To_PrivacyKeyRules())

	// 1. cacheUserData
	d.CachedConn.SetCache(ctx, genCacheUserDataCacheKey(userDO.Id), cacheUserData)

	// 2. PutLastSeenAt
	d.PutLastSeenAt(ctx, userDO.Id, now, 300)

	return mtproto.MakeTLImmutableUser(&mtproto.ImmutableUser{
		User:             cacheUserData.UserData,
		LastSeenAt:       now,
		Contacts:         nil,
		KeysPrivacyRules: nil,
	}).To_ImmutableUser(), nil
}

func (d *Dao) UpdateUserFirstAndLastName(ctx context.Context, id int64, firstName, lastName string) bool {
	_, _, err := d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			rowsAffected, err := d.UsersDAO.UpdateUser(ctx, map[string]interface{}{
				"first_name": firstName,
				"last_name":  lastName,
			}, id)

			if err != nil {
				return 0, 0, err
			}

			return 0, rowsAffected, nil
		},
		genCacheUserDataCacheKey(id))
	if err != nil {
		logx.WithContext(ctx).Errorf("updateUserFirstAndLastName - error: %v", err)
		return false
	}

	return true
}

func (d *Dao) UpdateUserAbout(ctx context.Context, id int64, about string) bool {
	_, _, err := d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			rowsAffected, err := d.UsersDAO.UpdateUser(ctx, map[string]interface{}{
				"about": about,
			}, id)

			if err != nil {
				return 0, 0, err
			}

			return 0, rowsAffected, nil
		},
		genCacheUserDataCacheKey(id))
	if err != nil {
		logx.WithContext(ctx).Errorf("updateUserAbout - error: %v", err)
		return false
	}

	return true
}

func (d *Dao) UpdateUserUsername(ctx context.Context, id int64, username string) bool {
	_, _, err := d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			rowsAffected, err := d.UsersDAO.UpdateUser(ctx, map[string]interface{}{
				"username": username,
			}, id)

			if err != nil {
				return 0, 0, err
			}

			return 0, rowsAffected, nil
		},
		genCacheUserDataCacheKey(id))
	if err != nil {
		logx.WithContext(ctx).Errorf("updateUserUsername - error: %v", err)
		return false
	}

	return true
}

//func (d *Dao) DeleteProfilePhoto(ctx context.Context, userId, photoId int64) int64 {
//}
//
//func (d *Dao) DeleteMainProfilePhoto(ctx context.Context, userId int64) int64 {
//}

func (d *Dao) UpdateProfilePhoto(ctx context.Context, userId, photoId int64) int64 {
	var (
		mainPhotoId = photoId
	)

	_, _, err := d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			var err error
			if photoId == 0 {
				mainPhotoId, _ = d.UsersDAO.SelectProfilePhoto(ctx, userId)
				if mainPhotoId > 0 {
					nextPhotoId, _ := d.UserProfilePhotosDAO.SelectNext(ctx, userId, []int64{mainPhotoId})
					tR := sqlx.TxWrapper(ctx, d.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
						_, result.Err = d.UserProfilePhotosDAO.DeleteTx(tx, userId, []int64{mainPhotoId})
						if result.Err != nil {
							return
						}
						_, result.Err = d.UsersDAO.UpdateProfilePhotoTx(tx, nextPhotoId, userId)
					})
					mainPhotoId = nextPhotoId
					err = tR.Err
				} else {
					_, err = d.UsersDAO.UpdateProfilePhoto(ctx, 0, userId)
					mainPhotoId = 0
				}
			} else {
				tR := sqlx.TxWrapper(ctx, d.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
					_, _, result.Err = d.UserProfilePhotosDAO.InsertOrUpdateTx(tx, &dataobject.UserProfilePhotosDO{
						UserId:  userId,
						PhotoId: mainPhotoId,
						Date2:   time.Now().Unix(),
					})
					if result.Err != nil {
						return
					}
					_, result.Err = d.UsersDAO.UpdateProfilePhotoTx(tx, mainPhotoId, userId)
				})
				err = tR.Err
			}

			return 0, 0, err
		},
		genCacheUserDataCacheKey(userId))
	if err != nil {
		logx.WithContext(ctx).Errorf("updateProfilePhoto - error: %v", err)
		return 0
	}

	return mainPhotoId
}

func (d *Dao) GetImmutableUser(ctx context.Context, id int64, privacy bool, contacts ...int64) (*mtproto.ImmutableUser, error) {
	cacheUserData := d.GetCacheUserData(ctx, id)

	// userDO, _ := c.svcCtx.Dao.UsersDAO.SelectById(c.ctx, in.Id)
	if cacheUserData == nil {
		err := mtproto.ErrUserIdInvalid
		logx.WithContext(ctx).Errorf("user.getImmutableUser - error: %v", err)
		return nil, err
	}
	userData := cacheUserData.UserData
	immutableUser := mtproto.MakeTLImmutableUser(&mtproto.ImmutableUser{
		User:             userData,
		LastSeenAt:       0,
		Contacts:         nil,
		KeysPrivacyRules: nil,
	}).To_ImmutableUser()

	if userData.Deleted {
		return immutableUser, nil
	}

	if userData.UserType == user.UserTypeUnknown ||
		userData.UserType == user.UserTypeBot ||
		userData.UserType == user.UserTypeDeleted {
		// not need load
		return immutableUser, nil
	}

	mr.FinishVoid(
		func() {
			lastSeenAt, _ := d.GetLastSeenAt(ctx, id)
			if lastSeenAt != nil {
				immutableUser.LastSeenAt = lastSeenAt.LastSeenAt
			}
		},
		func() {
			// TODO: aaa
			// immutableUser.Contacts = c.svcCtx.Dao.GetUserContactListByIdList(c.ctx, id, contacts...)

			idList := cacheUserData.GetContactIdList()
			if len(idList) == 0 {
				return
			}

			idList2 := make([]int64, 0, len(idList))
			for _, id2 := range contacts {
				if ok := container2.ContainsInt64(idList, id2); ok && id2 != id {
					idList2 = append(idList2, id2)
				}
			}
			if len(idList2) == 0 {
				return
			}

			immutableUser.Contacts = d.getContactListByIdList(ctx, id, idList2)
		})
	//func() {
	//	if privacy {
	//		immutableUser.KeysPrivacyRules = c.svcCtx.Dao.GetUserPrivacyRulesListByKeys(
	//			c.ctx,
	//			id,
	//			user.STATUS_TIMESTAMP,
	//			user.PROFILE_PHOTO,
	//			user.PHONE_NUMBER)
	//	}
	//})
	if privacy {
		immutableUser.KeysPrivacyRules = cacheUserData.CachesPrivacyKeyRules
	}

	return immutableUser, nil
}

func (d *Dao) UpdateUserEmojiStatus(ctx context.Context, id int64, emojiStatusDocumentId int64, emojiStatusUntil int32) bool {
	_, _, err := d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			rowsAffected, err := d.UsersDAO.UpdateEmojiStatus(
				ctx,
				emojiStatusDocumentId,
				emojiStatusUntil,
				id)

			if err != nil {
				return 0, 0, err
			}

			return 0, rowsAffected, nil
		},
		genCacheUserDataCacheKey(id))
	if err != nil {
		logx.WithContext(ctx).Errorf("updateUserEmojiStatus - error: %v", err)
		return false
	}

	return true
}

func (d *Dao) DeleteUser(ctx context.Context, id int64, reason string) bool {
	_, _, err := d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			rowsAffected, err := d.UsersDAO.Delete(
				ctx,
				"-"+strconv.FormatInt(id, 10), // hack
				reason,
				id)
			if err != nil {
				return 0, 0, err
			}

			return 0, rowsAffected, nil
		},
		genCacheUserDataCacheKey(id))
	if err != nil {
		logx.WithContext(ctx).Errorf("DeleteUser - error: %v", err)
		return false
	}

	return true
}

func (d *Dao) GetCacheImmutableUserList(ctx context.Context, idList2 []int64, contacts []int64) []*mtproto.ImmutableUser {
	id := make([]int64, 0, len(idList2)+len(contacts))
	for _, v := range idList2 {
		if ok := container2.ContainsInt64(id, v); !ok {
			id = append(id, v)
		}
	}
	for _, v := range contacts {
		if ok := container2.ContainsInt64(id, v); !ok {
			id = append(id, v)
		}
	}

	if len(id) == 0 {
		return []*mtproto.ImmutableUser{}
	} else if len(id) == 1 {
		immutableUser, _ := d.GetImmutableUser(ctx, id[0], false)
		if immutableUser != nil {
			return []*mtproto.ImmutableUser{immutableUser}
		} else {
			return []*mtproto.ImmutableUser{}
		}
	}

	var (
		mUsers = make([]*mtproto.ImmutableUser, len(id))
	)

	mr.ForEach(
		func(source chan<- interface{}) {
			for idx := 0; idx < len(id); idx++ {
				source <- idx
			}
		},
		func(item interface{}) {
			var (
				idx = item.(int)
				err error
			)

			if ok := container2.ContainsInt64(contacts, id[idx]); ok {
				mUsers[idx], err = d.GetImmutableUser(ctx, id[idx], true, idList2...)
				if err != nil {
					logx.WithContext(ctx).Errorf("getImmutableUser - error: %v", err)
				}
			} else {
				if len(contacts) == 0 {
					mUsers[idx], err = d.GetImmutableUser(ctx, id[idx], true, idList2...)
					if err != nil {
						logx.WithContext(ctx).Errorf("getImmutableUser - error: %v", err)
					}
				} else {
					mUsers[idx], err = d.GetImmutableUser(ctx, id[idx], true, contacts...)
					if err != nil {
						logx.WithContext(ctx).Errorf("getImmutableUser - error: %v", err)
					}
				}
			}
		})

	for i := 0; i < len(mUsers); {
		if mUsers[i] != nil {
			i++
			continue
		}

		if i < len(mUsers)-1 {
			copy(mUsers[i:], mUsers[i+1:])
		}

		mUsers[len(mUsers)-1] = nil
		mUsers = mUsers[:len(mUsers)-1]
	}

	return mUsers
}

func (d *Dao) UpdateStoriesMaxId(ctx context.Context, id int64, maxId int32) bool {
	_, _, err := d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			rowsAffected, err := d.UsersDAO.UpdateStoriesMaxId(ctx, maxId, id)

			if err != nil {
				return 0, 0, err
			}

			return 0, rowsAffected, nil
		},
		genCacheUserDataCacheKey(id))
	if err != nil {
		logx.WithContext(ctx).Errorf("updateStoriesMaxId - error: %v", err)
		return false
	}

	return true
}

func (d *Dao) UpdateColor(ctx context.Context, id int64, forProfile bool, color int32, backgroundEmojiId int64) bool {
	_, _, err := d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			var (
				err          error
				rowsAffected int64
			)

			if forProfile {
				rowsAffected, err = d.UsersDAO.UpdateProfileColor(ctx, color, backgroundEmojiId, id)
			} else {
				rowsAffected, err = d.UsersDAO.UpdateColor(ctx, color, backgroundEmojiId, id)
			}

			if err != nil {
				return 0, 0, err
			}

			return 0, rowsAffected, nil
		},
		genCacheUserDataCacheKey(id))
	if err != nil {
		logx.WithContext(ctx).Errorf("updateColor - error: %v", err)
		return false
	}

	return true
}

func (d *Dao) UpdateBirthday(ctx context.Context, id int64, birthday *mtproto.Birthday) bool {
	_, _, err := d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			var (
				err          error
				rowsAffected int64
			)

			rowsAffected, err = d.UsersDAO.UpdateBirthday(ctx, birthday.ToBirthdayString(), id)

			if err != nil {
				return 0, 0, err
			}

			return 0, rowsAffected, nil
		},
		genCacheUserDataCacheKey(id))
	if err != nil {
		logx.WithContext(ctx).Errorf("updateBirthday - error: %v", err)
		return false
	}

	return true
}

func (d *Dao) GetCacheImmutableUserListV2(ctx context.Context, idList2 []int64, contacts []int64) []*mtproto.ImmutableUser {
	logger := logx.WithContext(ctx)

	logger.Infof("getCacheImmutableUserList - request: {id: %v, contacts: %v}", idList2, contacts)

	id2 := make([]int64, 0, len(idList2)+len(contacts))
	for _, v := range idList2 {
		if ok := container2.ContainsInt64(id2, v); !ok {
			id2 = append(id2, v)
		}
	}
	for _, v := range contacts {
		if ok := container2.ContainsInt64(id2, v); !ok {
			id2 = append(id2, v)
		}
	}

	if len(id2) == 0 {
		return []*mtproto.ImmutableUser{}
	}

	if len(id2) == 1 {
		immutableUser, _ := d.GetImmutableUser(ctx, id2[0], false)
		if immutableUser != nil {
			return []*mtproto.ImmutableUser{immutableUser}
		} else {
			return []*mtproto.ImmutableUser{}
		}
	}

	// len(id) > 1

	cDataList := d.GetCacheUserDataListByIdList(ctx, id2)
	if len(cDataList) == 0 {
		return []*mtproto.ImmutableUser{}
	}

	var (
		keyList   = make([]string, 0, len(id2))
		cUserList = make([]*mtproto.ImmutableUser, 0, len(cDataList))
		mUsers    = make(map[int64]*mtproto.ImmutableUser, len(cDataList))
	)

	for _, cData := range cDataList {
		id := cData.GetUserData().GetId()

		cUser := mtproto.MakeTLImmutableUser(&mtproto.ImmutableUser{
			User:             cData.GetUserData(),
			LastSeenAt:       0,
			Contacts:         nil,
			KeysPrivacyRules: cData.CachesPrivacyKeyRules,
		}).To_ImmutableUser()

		cUserList = append(cUserList, cUser)
		mUsers[id] = cUser

		// LastSeenAt
		keyList = append(keyList, genUserPresencesKey(cData.GetUserData().GetId()))

		// contacts
		var (
			myContacts []int64
		)
		if ok := container2.ContainsInt64(contacts, id); ok {
			myContacts = idList2
		} else {
			if len(contacts) == 0 {
				myContacts = idList2
			} else {
				myContacts = contacts
			}
		}

		cIdList := cData.GetContactIdList()
		if len(cIdList) == 0 {
			continue
		}
		for _, v := range myContacts {
			if ok := container2.ContainsInt64(cIdList, v); ok && v != id {
				keyList = append(keyList, genContactCacheKey(id, v))
			}
		}
	}
	logger.Infof("getCacheImmutableUserList - cDataList: %d", len(cDataList))

	d.CachedConn.QueryRows(
		ctx,
		func(ctx context.Context, conn *sqlx.DB, keys ...string) (map[string]interface{}, error) {
			noCaches := make(map[string]interface{}, len(keys))
			for _, key := range keys {
				logger.Infof("getCacheImmutableUserList - noCache: %v", keys)
				if isUserPresencesKey(key) {
					lastSeenAt, err := d.UserPresencesDAO.Select(ctx, parseUserPresencesKey(key))
					if err != nil {
						continue
					}
					if lastSeenAt != nil {
						noCaches[key] = lastSeenAt

						if cUser, ok := mUsers[lastSeenAt.UserId]; ok {
							cUser.LastSeenAt = lastSeenAt.LastSeenAt
						}
					}
				} else {
					id0, id1 := parseContactCacheKey(key)
					contact, _ := d.UserContactsDAO.SelectContact(ctx, id0, id1)
					if contact == nil {
						// return sqlc.ErrNotFound
						continue
					}
					noCaches[key] = contact
					if cUser, ok := mUsers[contact.OwnerUserId]; ok {
						cUser.Contacts = append(cUser.Contacts, mtproto.MakeTLContactData(&mtproto.ContactData{
							UserId:        contact.OwnerUserId,
							ContactUserId: contact.ContactUserId,
							FirstName:     mtproto.MakeFlagsString(contact.ContactFirstName),
							LastName:      mtproto.MakeFlagsString(contact.ContactLastName),
							MutualContact: contact.Mutual,
							Phone:         mtproto.MakeFlagsString(contact.ContactPhone),
							CloseFriend:   contact.CloseFriend,
						}).To_ContactData())
					}
				}
			}

			return noCaches, nil
		},
		func(k, v string) (interface{}, error) {
			logger.Infof("getCacheImmutableUserList - cache: {k: %s, v: %d}", k, len(v))

			if isUserPresencesKey(k) {
				var (
					lastSeenAt *dataobject.UserPresencesDO
				)
				err := jsonx.UnmarshalFromString(v, &lastSeenAt)
				if err != nil {
					return nil, err
				}

				if cUser, ok := mUsers[lastSeenAt.UserId]; ok {
					cUser.LastSeenAt = lastSeenAt.LastSeenAt
				}

				return lastSeenAt, nil
			} else {
				var (
					contact *dataobject.UserContactsDO
				)

				err := jsonx.UnmarshalFromString(v, &contact)
				if err != nil {
					return nil, err
				}

				if cUser, ok := mUsers[contact.OwnerUserId]; ok {
					cUser.Contacts = append(cUser.Contacts, mtproto.MakeTLContactData(&mtproto.ContactData{
						UserId:        contact.OwnerUserId,
						ContactUserId: contact.ContactUserId,
						FirstName:     mtproto.MakeFlagsString(contact.ContactFirstName),
						LastName:      mtproto.MakeFlagsString(contact.ContactLastName),
						MutualContact: contact.Mutual,
						Phone:         mtproto.MakeFlagsString(contact.ContactPhone),
						CloseFriend:   contact.CloseFriend,
					}).To_ContactData())
				}

				return contact, nil
			}
		},
		keyList...)

	logger.Infof("getCacheImmutableUserList - cUserList: %d", len(cUserList))

	return cUserList
}

func (d *Dao) GetImmutableUserV2(ctx context.Context, id int64, privacy bool, hasReverseContacts bool, reverseContacts []int64) (*mtproto.ImmutableUser, error) {
	cacheUserData := d.GetCacheUserData(ctx, id)

	if cacheUserData == nil {
		err := mtproto.ErrUserIdInvalid
		logx.WithContext(ctx).Errorf("user.getImmutableUser - error: %v", err)
		return nil, err
	}
	userData := cacheUserData.UserData
	immutableUser := mtproto.MakeTLImmutableUser(&mtproto.ImmutableUser{
		User:             userData,
		LastSeenAt:       0,
		Contacts:         nil,
		KeysPrivacyRules: nil,
		ReverseContacts:  nil,
	}).To_ImmutableUser()

	if userData.Deleted {
		return immutableUser, nil
	}

	if userData.UserType == user.UserTypeUnknown ||
		userData.UserType == user.UserTypeBot ||
		userData.UserType == user.UserTypeDeleted {
		// not load these data
		return immutableUser, nil
	}

	var (
		rIdList []int64
	)

	if hasReverseContacts && len(cacheUserData.ReverseContactIdList) > 0 {
		if len(reverseContacts) == 0 {
			rIdList = cacheUserData.ReverseContactIdList
		} else {
			for _, id2 := range reverseContacts {
				if ok := container2.ContainsInt64(cacheUserData.ReverseContactIdList, id2); ok && id2 != id {
					rIdList = append(rIdList, id2)
				}
			}
		}
	}

	fns := []func(){
		func() {
			lastSeenAt, _ := d.GetLastSeenAt(ctx, id)
			if lastSeenAt != nil {
				immutableUser.LastSeenAt = lastSeenAt.LastSeenAt
			}
		},
	}

	if len(rIdList) > 0 {
		fns = append(
			fns,
			func() {
				immutableUser.ReverseContacts = d.getReverseContactListByIdList(ctx, id, rIdList)
			})

	}

	if len(fns) == 1 {
		fns[0]()
	} else {
		mr.FinishVoid(fns...)
	}

	if privacy {
		immutableUser.KeysPrivacyRules = cacheUserData.CachesPrivacyKeyRules
	}

	return immutableUser, nil
}

func (d *Dao) GetMutableUsersV2(ctx context.Context, idList2 []int64, privacy bool, hasTo bool, to []int64) []*mtproto.ImmutableUser {
	if len(idList2) == 0 {
		return []*mtproto.ImmutableUser{}
	}

	if len(idList2) == 1 {
		immutableUser, _ := d.GetImmutableUserV2(ctx, idList2[0], privacy, hasTo, to)
		if immutableUser != nil {
			return []*mtproto.ImmutableUser{immutableUser}
		} else {
			return []*mtproto.ImmutableUser{}
		}
	}

	cDataList := d.GetCacheUserDataListByIdList(ctx, idList2)
	if len(cDataList) == 0 {
		return []*mtproto.ImmutableUser{}
	}

	var (
		keyList   = make([]string, 0, len(idList2))
		cUserList = make([]*mtproto.ImmutableUser, 0, len(cDataList))
		mUsers    = make(map[int64]*mtproto.ImmutableUser, len(cDataList))
	)

	for _, cData := range cDataList {
		id := cData.GetUserData().GetId()

		cUser := mtproto.MakeTLImmutableUser(&mtproto.ImmutableUser{
			User:             cData.GetUserData(),
			LastSeenAt:       0,
			Contacts:         nil,
			ReverseContacts:  nil,
			KeysPrivacyRules: nil,
		}).To_ImmutableUser()
		if privacy {
			cUser.KeysPrivacyRules = cData.CachesPrivacyKeyRules
		}

		cUserList = append(cUserList, cUser)
		mUsers[id] = cUser

		// LastSeenAt
		keyList = append(keyList, genUserPresencesKey(cData.GetUserData().GetId()))

		// contacts
		var (
			rContacts []int64
		)

		if hasTo && len(cData.ReverseContactIdList) > 0 {
			if len(to) == 0 {
				rContacts = cData.ReverseContactIdList
			} else {
				for _, rId := range to {
					if ok := container2.ContainsInt64(cData.ReverseContactIdList, rId); ok {
						rContacts = append(rContacts, rId)
					}
				}
			}
		}

		for _, v := range rContacts {
			keyList = append(keyList, genContactCacheKey(v, id))
		}
	}

	d.CachedConn.QueryRows(
		ctx,
		func(ctx context.Context, conn *sqlx.DB, keys ...string) (map[string]interface{}, error) {
			noCaches := make(map[string]interface{}, len(keys))
			for _, key := range keys {
				//logger.Infof("getCacheImmutableUserList - noCache: %v", keys)
				if isUserPresencesKey(key) {
					lastSeenAt, err := d.UserPresencesDAO.Select(ctx, parseUserPresencesKey(key))
					if err != nil {
						continue
					}
					if lastSeenAt != nil {
						noCaches[key] = lastSeenAt

						if cUser, ok := mUsers[lastSeenAt.UserId]; ok {
							cUser.LastSeenAt = lastSeenAt.LastSeenAt
						}
					}
				} else {
					id0, id1 := parseContactCacheKey(key)
					contact, _ := d.UserContactsDAO.SelectContact(ctx, id0, id1)
					if contact == nil {
						// return sqlc.ErrNotFound
						continue
					}
					noCaches[key] = contact
					if cUser, ok := mUsers[contact.ContactUserId]; ok {
						cUser.ReverseContacts = append(cUser.ReverseContacts, mtproto.MakeTLContactData(&mtproto.ContactData{
							UserId:        contact.OwnerUserId,
							ContactUserId: contact.ContactUserId,
							FirstName:     mtproto.MakeFlagsString(contact.ContactFirstName),
							LastName:      mtproto.MakeFlagsString(contact.ContactLastName),
							MutualContact: contact.Mutual,
							Phone:         mtproto.MakeFlagsString(contact.ContactPhone),
							CloseFriend:   contact.CloseFriend,
						}).To_ContactData())
					}
				}
			}

			return noCaches, nil
		},
		func(k, v string) (interface{}, error) {
			//logger.Infof("getCacheImmutableUserList - cache: {k: %s, v: %d}", k, len(v))
			if isUserPresencesKey(k) {
				var (
					lastSeenAt *dataobject.UserPresencesDO
				)
				err := jsonx.UnmarshalFromString(v, &lastSeenAt)
				if err != nil {
					return nil, err
				}

				if cUser, ok := mUsers[lastSeenAt.UserId]; ok {
					cUser.LastSeenAt = lastSeenAt.LastSeenAt
				}

				return lastSeenAt, nil
			} else {
				var (
					contact *dataobject.UserContactsDO
				)

				err := jsonx.UnmarshalFromString(v, &contact)
				if err != nil {
					return nil, err
				}

				if cUser, ok := mUsers[contact.ContactUserId]; ok {
					cUser.ReverseContacts = append(cUser.ReverseContacts, mtproto.MakeTLContactData(&mtproto.ContactData{
						UserId:        contact.OwnerUserId,
						ContactUserId: contact.ContactUserId,
						FirstName:     mtproto.MakeFlagsString(contact.ContactFirstName),
						LastName:      mtproto.MakeFlagsString(contact.ContactLastName),
						MutualContact: contact.Mutual,
						Phone:         mtproto.MakeFlagsString(contact.ContactPhone),
						CloseFriend:   contact.CloseFriend,
					}).To_ContactData())
				}

				return contact, nil
			}
		},
		keyList...)

	return cUserList
}
