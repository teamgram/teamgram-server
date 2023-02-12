// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package dao

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlc"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/app/service/media/media"

	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

const (
	cacheUserDataKeyPrefix = "user_data2"
)

var (
	GenCacheUserDataCacheKey   = genCacheUserDataCacheKey
	ParseCacheUserDataCacheKey = parseCacheUserDataCacheKey
	IsCacheUserDataCacheKey    = isCacheUserDataCacheKey
)

func genCacheUserDataCacheKey(id int64) string {
	return fmt.Sprintf("%s_%d", cacheUserDataKeyPrefix, id)
}

func parseCacheUserDataCacheKey(k string) int64 {
	if strings.HasPrefix(k, cacheUserDataKeyPrefix+"_") {
		v, _ := strconv.ParseInt(k[len(cacheUserDataKeyPrefix)+1:], 10, 64)
		return v
	}

	return 0
}

func isCacheUserDataCacheKey(k string) bool {
	return strings.HasPrefix(k, cacheUserDataKeyPrefix)
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

func (d *Dao) GetNoCacheUserData(ctx context.Context, id int64) (*CacheUserData, error) {
	var (
		rules0, rules1, rules2 *user.PrivacyKeyRules
		cacheData              = NewCacheUserData()
	)

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

			cacheData.UserData = userData
			return nil
		},
		func() error {
			idList2, err := d.UserContactsDAO.SelectUserContactIdList(ctx, id)
			if err != nil {
				return err
			}
			cacheData.ContactIdList = idList2

			return nil
		},
		func() error {
			rules0, _ = d.GetUserPrivacyRules(ctx, id, user.STATUS_TIMESTAMP)
			return nil
		},
		func() error {
			rules1, _ = d.GetUserPrivacyRules(ctx, id, user.PROFILE_PHOTO)
			return nil
		},
		func() error {
			rules2, _ = d.GetUserPrivacyRules(ctx, id, user.PHONE_NUMBER)
			return nil
		})

	if err2 != nil {
		return nil, err2
	}

	if rules0 != nil {
		cacheData.CachesPrivacyKeyRules = append(cacheData.CachesPrivacyKeyRules, rules0)
	}
	if rules1 != nil {
		cacheData.CachesPrivacyKeyRules = append(cacheData.CachesPrivacyKeyRules, rules1)
	}
	if rules2 != nil {
		cacheData.CachesPrivacyKeyRules = append(cacheData.CachesPrivacyKeyRules, rules2)
	}

	return cacheData, nil
}
