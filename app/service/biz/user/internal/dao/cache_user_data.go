// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package dao

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlc"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/app/service/media/media"

	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

const (
	cacheUserDataKeyPrefix = "user_data.2"
	cachePhoneUserPrefix   = "phone_user.1"
)

var (
	GenCacheUserDataCacheKey = genCacheUserDataCacheKey
)

func genCacheUserDataCacheKey(id int64) string {
	return fmt.Sprintf("%s#%d", cacheUserDataKeyPrefix, id)
}

func genCachePhoneUserKey(id string) string {
	return fmt.Sprintf("%s#%s", cachePhoneUserPrefix, id)
}

func parseCacheUserDataCacheKey(k string) int64 {
	if strings.HasPrefix(k, cacheUserDataKeyPrefix+"#") {
		v, _ := strconv.ParseInt(k[len(cacheUserDataKeyPrefix)+1:], 10, 64)
		return v
	}

	return 0
}

func isCacheUserDataCacheKey(k string) bool {
	return strings.HasPrefix(k, cacheUserDataKeyPrefix)
}

type CacheUserData struct {
	UserData              *tg.TLUserData            `json:"user_data"`
	ContactIdList         []int64                   `json:"contact_id_list"`
	CachesPrivacyKeyRules []tg.PrivacyKeyRulesClazz `json:"caches_privacy_key_rules"`
	ReverseContactIdList  []int64                   `json:"reverse_contact_id_list"`
}

func NewCacheUserData() *CacheUserData {
	return &CacheUserData{
		UserData:              nil,
		ContactIdList:         []int64{},
		CachesPrivacyKeyRules: []tg.PrivacyKeyRulesClazz{},
		ReverseContactIdList:  []int64{},
	}
}

func (m *CacheUserData) GetContactIdList() []int64 {
	if m == nil {
		return nil
	}
	return m.ContactIdList
}

func (m *CacheUserData) GetReverseContactIdList() []int64 {
	if m == nil {
		return nil
	}
	return m.ReverseContactIdList
}

func (m *CacheUserData) GetUserData() *tg.TLUserData {
	if m == nil {
		return nil
	}
	return m.UserData
}

func (d *Dao) GetCacheUserData(ctx context.Context, id int64) *CacheUserData {
	var (
		cacheUserData *CacheUserData
	)

	err := d.CachedConn.QueryRow(
		ctx,
		&cacheUserData,
		genCacheUserDataCacheKey(id),
		func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
			cacheData, err2 := d.GetNoCacheUserData(ctx, id)
			if err2 != nil {
				return err2
			}
			*v.(**CacheUserData) = cacheData

			return nil
		})
	if err != nil {
		logx.WithContext(ctx).Errorf("GetCacheUserData(%d) - error: %v", id, err)
		return nil
	}

	return cacheUserData
}

func makeEmojiStatus(documentId int64, until int32) tg.EmojiStatusClazz {
	//if documentId == 0 {
	//	return nil
	//}
	//if until > 0 {
	return tg.MakeTLEmojiStatus(&tg.TLEmojiStatus{
		DocumentId: documentId,
	})
	//} else {
	//	return tg.MakeEmojiStatus(&tg.TLEmojiStatusUntil{
	//		DocumentId:      documentId,
	//		Until_INT32:     until,
	//		Until_FLAGINT32: mtproto.MakeFlagsInt32(until),
	//	}).To_EmojiStatus()
	//}
}

func makePeerColor(color int32, backgroundEmojiId int64) tg.PeerColorClazz {
	if color == 0 && backgroundEmojiId == 0 {
		return nil
	}

	return tg.MakeTLPeerColor(&tg.TLPeerColor{
		Color:             tg.MakeFlagsInt32(color),
		BackgroundEmojiId: tg.MakeFlagsInt64(backgroundEmojiId),
	})
}

func (d *Dao) MakeUserDataByDO(userDO *dataobject.UsersDO) *tg.TLUserData {
	userData := tg.MakeTLUserData(&tg.TLUserData{
		Id:                 userDO.Id,
		AccessHash:         userDO.AccessHash,
		Deleted:            userDO.Deleted,
		UserType:           userDO.UserType,
		SceretKeyId:        userDO.SecretKeyId,
		FirstName:          userDO.FirstName,
		LastName:           userDO.LastName,
		Username:           userDO.Username,
		Phone:              userDO.Phone,
		ProfilePhoto:       nil,
		Bot:                nil,
		CountryCode:        userDO.CountryCode,
		Verified:           userDO.Verified,
		Support:            userDO.Support,
		Scam:               userDO.Scam,
		Fake:               userDO.Fake,
		About:              tg.MakeFlagsString(userDO.About),
		Restricted:         userDO.Restricted,
		RestrictionReason:  nil,
		ContactsVersion:    1,
		PrivaciesVersion:   1,
		Premium:            userDO.Premium,
		EmojiStatus:        makeEmojiStatus(userDO.EmojiStatusDocumentId, userDO.EmojiStatusUntil),
		StoriesUnavailable: true,
		StoriesMaxId:       userDO.StoriesMaxId,
		Color:              makePeerColor(userDO.Color, userDO.ColorBackgroundEmojiId),
		ProfileColor:       makePeerColor(userDO.ProfileColor, userDO.ProfileColorBackgroundEmojiId),
		Birthday:           userDO.Birthday,
		PersonalChannelId:  userDO.PersonalChannelId,
		PremiumExpireDate:  tg.MakeFlagsInt64(userDO.PremiumExpireDate),
	})

	return userData
}

func (d *Dao) GetNoCacheUserData(ctx context.Context, id int64) (*CacheUserData, error) {
	do, err3 := d.UsersDAO.SelectById(ctx, id)
	if err3 != nil {
		return nil, err3
	}
	if do == nil {
		return nil, sqlc.ErrNotFound
	}

	var (
		rules0, rules1, rules2 *tg.PrivacyKeyRules
		cacheData              = NewCacheUserData()
	)

	userData := d.MakeUserDataByDO(do)
	if do.Restricted {
		_ = jsonx.UnmarshalFromString(do.RestrictionReason, &userData.RestrictionReason)
	}
	cacheData.UserData = userData

	if do.UserType == user.UserTypeUnknown ||
		do.UserType == user.UserTypeDeleted ||
		do.Deleted {
		return cacheData, nil
	}

	if do.UserType == user.UserTypeBot {
		if do.PhotoId != 0 {
			mr.FinishVoid(
				func() {
					userData.Bot = d.getBotData(ctx, do.Id)
				},
				func() {
					profilePhoto, _ := d.MediaClient.MediaGetPhoto(ctx, &media.TLMediaGetPhoto{
						PhotoId: do.PhotoId,
					})
					userData.ProfilePhoto = profilePhoto.Clazz
				})
		} else {
			userData.Bot = d.getBotData(ctx, do.Id)
		}

		return cacheData, nil
	}

	mr.FinishVoid(
		func() {
			if do.PhotoId != 0 {
				profilePhoto, _ := d.MediaClient.MediaGetPhoto(ctx, &media.TLMediaGetPhoto{
					PhotoId: do.PhotoId,
				})
				userData.ProfilePhoto = profilePhoto.Clazz
			}
		},
		func() {
			cacheData.ContactIdList, _ = d.UserContactsDAO.SelectUserContactIdList(ctx, id)
		},
		func() {
			cacheData.ReverseContactIdList, _ = d.UserContactsDAO.SelectUserReverseContactIdList(ctx, id)
			if len(cacheData.ReverseContactIdList) > 0 {
				sort.Slice(cacheData.ReverseContactIdList, func(i, j int) bool { return cacheData.ReverseContactIdList[i] < cacheData.ReverseContactIdList[j] })
			}
		},
		func() {
			rules0, _ = d.GetUserPrivacyRules(ctx, id, tg.STATUS_TIMESTAMP)
		},
		func() {
			rules1, _ = d.GetUserPrivacyRules(ctx, id, tg.PROFILE_PHOTO)
		},
		func() {
			rules2, _ = d.GetUserPrivacyRules(ctx, id, tg.PHONE_NUMBER)
		})

	if rules0 != nil {
		cacheData.CachesPrivacyKeyRules = append(cacheData.CachesPrivacyKeyRules, rules0.Clazz)
	}
	if rules1 != nil {
		cacheData.CachesPrivacyKeyRules = append(cacheData.CachesPrivacyKeyRules, rules1.Clazz)
	}
	if rules2 != nil {
		cacheData.CachesPrivacyKeyRules = append(cacheData.CachesPrivacyKeyRules, rules2.Clazz)
	}

	// TODO
	// close_friends:Vector<long>

	// stories_hiddens:Vector<long>

	return cacheData, nil
}

func (d *Dao) GetCacheUserDataListByIdList(ctx context.Context, idList []int64) []*CacheUserData {
	var (
		keyList   = make([]string, 0, len(idList))
		cDataList = make([]*CacheUserData, 0, len(idList))
	)

	for _, id := range idList {
		keyList = append(keyList, genCacheUserDataCacheKey(id))
	}

	_ = d.CachedConn.QueryRows(
		ctx,
		func(ctx context.Context, conn *sqlx.DB, keys ...string) (map[string]interface{}, error) {
			vList := make(map[string]interface{}, len(keys))

			// TODO: mr
			for _, key := range keys {
				id := parseCacheUserDataCacheKey(key)
				if cacheData, err2 := d.GetNoCacheUserData(ctx, id); err2 != nil {
					continue
				} else {
					vList[key] = cacheData
					cDataList = append(cDataList, cacheData)
				}
			}

			return vList, nil
		},
		func(k, v string) (interface{}, error) {
			var (
				c   *CacheUserData
				err error
			)
			err = jsonx.UnmarshalFromString(v, &c)
			if err != nil {
				return nil, err
			}

			cDataList = append(cDataList, c)

			return c, nil
		},
		keyList...)

	return cDataList
}

func (d *Dao) GetUserIdByPhone(ctx context.Context, phone string) (int64, error) {
	var (
		id int64
	)

	err := d.CachedConn.QueryRow(
		ctx,
		&id,
		genCachePhoneUserKey(phone),
		func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
			do, err := d.UsersDAO.SelectByPhoneNumber(ctx, phone)
			if err != nil {
				return err
			} else if do == nil {
				return tg.ErrPhoneNotOccupied
			}

			*v.(*int64) = do.Id

			return nil
		})

	if err != nil {
		// logx.WithContext(ctx).Errorf("GetUserIdByPhone(%s) - error: %v", phone, err)
		return 0, err
	}

	return id, nil
}
