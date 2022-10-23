// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package dao

import (
	"context"
	"fmt"

	"github.com/teamgram/marmota/pkg/stores/sqlc"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/app/service/media/media"

	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/mr"
)

const (
	cacheUserDataKeyPrefix = "user_data2"
)

func genCacheUserDataCacheKey(id int64) string {
	return fmt.Sprintf("%s_%d", cacheUserDataKeyPrefix, id)
}

type CacheUserData struct {
	UserData              *user.UserData          `json:"user_data"`
	ContactIdList         []int64                 `json:"contact_id_list"`
	CachesPrivacyKeyRules []*user.PrivacyKeyRules `json:"caches_privacy_key_rules"`
}

func NewCacheUserData() *CacheUserData {
	return &CacheUserData{
		UserData:              nil,
		ContactIdList:         []int64{},
		CachesPrivacyKeyRules: []*user.PrivacyKeyRules{},
	}
}

func (m *CacheUserData) GetContactIdList() []int64 {
	if m == nil {
		return nil
	}
	return m.ContactIdList
}

func (m *CacheUserData) GetUserData() *user.UserData {
	if m == nil {
		return nil
	}
	return m.UserData
}

func (d *Dao) GetCacheUserData(ctx context.Context, id int64) *CacheUserData {
	cacheUserData := NewCacheUserData()
	// user.MakeTLUserData(nil).To_UserData()

	err := d.CachedConn.QueryRow(
		ctx,
		cacheUserData,
		genCacheUserDataCacheKey(id),
		func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
			cacheData := v.(*CacheUserData)
			err2 := mr.Finish(
				func() error {
					do, err := d.UsersDAO.SelectById(ctx, id)
					if err != nil {
						return err
					}
					if do == nil {
						return sqlc.ErrNotFound
					}
					userData := d.makeUserDataByDO(do)

					// TODO: deleted
					if do.IsBot && do.PhotoId != 0 {
						mr.FinishVoid(
							func() {
								userData.Bot = d.getBotData(ctx, do.Id)
							},
							func() {
								userData.ProfilePhoto, _ = d.MediaClient.MediaGetPhoto(ctx, &media.TLMediaGetPhoto{
									PhotoId: do.PhotoId,
								})
							})
					} else {
						if do.IsBot {
							userData.Bot = d.getBotData(ctx, do.Id)
						}
						if do.PhotoId != 0 {
							userData.ProfilePhoto, _ = d.MediaClient.MediaGetPhoto(ctx, &media.TLMediaGetPhoto{
								PhotoId: do.PhotoId,
							})
						}
					}

					if do.Restricted {
						jsonx.UnmarshalFromString(do.RestrictionReason, &userData.RestrictionReason)
					}

					cacheUserData.UserData = userData
					return nil
				},
				func() error {
					idList, err := d.UserContactsDAO.SelectUserContactIdList(ctx, id)
					if err != nil {
						return err
					}
					cacheData.ContactIdList = idList

					return nil
				},
				func() error {
					d.UserPrivaciesDAO.SelectPrivacyListWithCB(
						ctx,
						id,
						[]int32{
							user.STATUS_TIMESTAMP,
							user.PROFILE_PHOTO,
							user.PHONE_NUMBER,
						},
						func(i int, v *dataobject.UserPrivaciesDO) {
							rules := user.MakeTLPrivacyKeyRules(&user.PrivacyKeyRules{
								Key:   v.KeyType,
								Rules: nil,
							}).To_PrivacyKeyRules()

							if err2 := jsonx.UnmarshalFromString(v.Rules, &rules.Rules); err2 != nil {
								// c.Logger.Errorf("user.getPrivacy - Unmarshal PrivacyRulesData(%d)error: %v", do.Id, err)
								// return err2
								return
							}

							cacheData.CachesPrivacyKeyRules = append(cacheData.CachesPrivacyKeyRules, rules)
						})

					return nil
				})

			return err2
		})
	if err != nil {
		return nil
	}

	return cacheUserData
}

func makeEmojiStatus(documentId int64, until int32) *mtproto.EmojiStatus {
	if documentId == 0 {
		return nil
	}
	if until > 0 {
		return mtproto.MakeTLEmojiStatus(&mtproto.EmojiStatus{
			DocumentId: documentId,
		}).To_EmojiStatus()
	} else {
		return mtproto.MakeTLEmojiStatusUntil(&mtproto.EmojiStatus{
			DocumentId: documentId,
			Until:      until,
		}).To_EmojiStatus()
	}
}

func (d *Dao) makeUserDataByDO(userDO *dataobject.UsersDO) *user.UserData {
	userData := user.MakeTLUserData(&user.UserData{
		Id:                userDO.Id,
		AccessHash:        userDO.AccessHash,
		Deleted:           userDO.Deleted,
		UserType:          userDO.UserType,
		SceretKeyId:       userDO.SecretKeyId,
		FirstName:         userDO.FirstName,
		LastName:          userDO.LastName,
		Username:          userDO.Username,
		Phone:             userDO.Phone,
		ProfilePhoto:      nil,
		Bot:               nil,
		CountryCode:       userDO.CountryCode,
		Verified:          userDO.Verified,
		Support:           userDO.Support,
		Scam:              userDO.Scam,
		Fake:              userDO.Fake,
		About:             mtproto.MakeFlagsString(userDO.About),
		Restricted:        userDO.Restricted,
		RestrictionReason: nil,
		ContactsVersion:   1,
		PrivaciesVersion:  1,
		BotAttachMenu:     false,
		Premium:           userDO.Premium,
		EmojiStatus:       makeEmojiStatus(userDO.EmojiStatusDocumentId, userDO.EmojiStatusUntil),
	}).To_UserData()

	return userData
}
