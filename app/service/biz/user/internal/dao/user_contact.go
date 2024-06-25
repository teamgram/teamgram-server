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
	"fmt"
	"strconv"
	"strings"

	"github.com/teamgram/marmota/pkg/container2"
	"github.com/teamgram/marmota/pkg/stores/sqlc"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/mr"
)

const (
	contactKeyPrefix = "user_contact.2"
)

var (
	GenContactCacheKey = genContactCacheKey
)

type ContactItem struct {
	C               *mtproto.InputContact
	Unregistered    bool  // 未注册
	UserId          int64 // 已经注册的用户ID
	ContactId       int64 // 已经注册是我的联系人
	ImportContactId int64 // 已经注册的反向联系人
}

func genContactCacheKey(selfId, contactId int64) string {
	return fmt.Sprintf("%s_%d_%d", contactKeyPrefix, selfId, contactId)
}

func isContactCacheKey(k string) bool {
	return strings.HasPrefix(k, contactKeyPrefix+"_")
}

func parseContactCacheKey(k string) (int64, int64) {
	if strings.HasPrefix(k, contactKeyPrefix+"_") {
		v := strings.Split(k[len(contactKeyPrefix)+1:], "_")
		if len(v) != 2 {
			return 0, 0
		}
		v0, _ := strconv.ParseInt(v[0], 10, 64)
		v1, _ := strconv.ParseInt(v[1], 10, 64)

		return v0, v1
	}

	return 0, 0
}

func (d *Dao) GetUserContactList(ctx context.Context, id int64) []*mtproto.ContactData {
	cacheUserData := d.GetCacheUserData(ctx, id)
	if len(cacheUserData.GetContactIdList()) == 0 {
		return nil
	}

	return d.getContactListByIdList(ctx, id, cacheUserData.GetContactIdList())
}

func (d *Dao) GetUserContact(ctx context.Context, id, contactId int64) *mtproto.ContactData {
	contacts := d.GetUserContactListByIdList(ctx, id, contactId)
	if len(contacts) == 0 {
		return nil
	}

	return contacts[0]
}

func (d *Dao) GetUserContactListByIdList(ctx context.Context, id int64, contactId ...int64) []*mtproto.ContactData {
	cacheUserData := d.GetCacheUserData(ctx, id)
	idList := cacheUserData.GetContactIdList()
	if len(idList) == 0 {
		return nil
	}

	idList2 := make([]int64, 0, len(idList))
	for _, id2 := range contactId {
		if ok := container2.ContainsInt64(idList, id2); ok {
			idList2 = append(idList2, id2)
		}
	}
	if len(idList2) == 0 {
		return nil
	}

	return d.getContactListByIdList(ctx, id, idList2)
}

func (d *Dao) DeleteUserContact(ctx context.Context, id int64, contactId int64) {
	_, affected, _ := d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			tR := sqlx.TxWrapper(
				ctx,
				d.DB,
				func(tx *sqlx.Tx, result *sqlx.StoreResult) {
					_, err := d.UserContactsDAO.DeleteContactsTx(tx, id, []int64{contactId})
					if err != nil {
						result.Err = err
						return
					}

					affected, err := d.UserContactsDAO.UpdateMutualTx(tx, false, contactId, id)
					result.Data = affected
					result.Err = err
				})
			return 0, tR.Data.(int64), tR.Err
		},
		genCacheUserDataCacheKey(id),
		genContactCacheKey(id, contactId))

	if affected != 0 {
		d.CachedConn.DelCache(ctx, genContactCacheKey(contactId, id), genCacheUserDataCacheKey(contactId))
	}
}

func (d *Dao) PutUserContact(ctx context.Context, changeMutual bool, do *dataobject.UserContactsDO) error {
	keys := []string{
		genCacheUserDataCacheKey(do.OwnerUserId),
		genCacheUserDataCacheKey(do.ContactUserId),
		genContactCacheKey(do.OwnerUserId, do.ContactUserId),
	}
	if changeMutual {
		keys = append(keys, genContactCacheKey(do.ContactUserId, do.OwnerUserId))
	}

	_, _, err := d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			tR := sqlx.TxWrapper(
				ctx,
				conn,
				func(tx *sqlx.Tx, result *sqlx.StoreResult) {
					if changeMutual {
						_, result.Err = d.UserContactsDAO.UpdateMutualTx(tx, true, do.ContactUserId, do.OwnerUserId)
						if result.Err != nil {
							return
						}
					}
					_, _, result.Err = d.UserContactsDAO.InsertOrUpdateTx(tx, do)
					if result.Err != nil {
						return
					}

					// // TODO(@benqi): set addPhonePrivacyException
					// if mtproto.FromBool(in.AddPhonePrivacyException) {
					// 	//
					// }
				})

			return 0, 0, tR.Err
		},
		keys...)

	return err
}

func (d *Dao) ClearContactCaches(ctx context.Context, userId int64, contactId ...int64) {
	keys := []string{genCacheUserDataCacheKey(userId)}
	for _, id := range contactId {
		keys = append(keys, genContactCacheKey(userId, id))
		keys = append(keys, genContactCacheKey(id, userId))
	}
	d.CachedConn.DelCache(ctx, keys...)
}

func (d *Dao) GetCloseFriendList(ctx context.Context, id int64) []*mtproto.ContactData {
	cacheUserData := d.GetCacheUserData(ctx, id)
	if len(cacheUserData.GetContactIdList()) == 0 {
		return nil
	}

	return d.getCloseFriendListByIdList(ctx, id, cacheUserData.GetContactIdList())
}

func (d *Dao) getCloseFriendListByIdList(ctx context.Context, id int64, idList []int64) []*mtproto.ContactData {
	closeFriendList := make([]*mtproto.ContactData, len(idList))
	mr.ForEach(
		func(source chan<- interface{}) {
			for i, v := range idList {
				source <- idxId{i, v}
			}
		},
		func(item interface{}) {
			idx := item.(idxId)
			do := new(dataobject.UserContactsDO)
			err := d.CachedConn.QueryRow(
				ctx,
				do,
				genContactCacheKey(id, idx.id),
				func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
					do2, _ := d.UserContactsDAO.SelectContact(ctx, id, idx.id)
					if do2 == nil {
						return sqlc.ErrNotFound
					}
					*v.(*dataobject.UserContactsDO) = *do2
					return nil
				})
			if err == nil && do.CloseFriend {
				closeFriendList[idx.idx] = mtproto.MakeTLContactData(&mtproto.ContactData{
					UserId:        id,
					ContactUserId: do.ContactUserId,
					FirstName:     mtproto.MakeFlagsString(do.ContactFirstName),
					LastName:      mtproto.MakeFlagsString(do.ContactLastName),
					MutualContact: do.Mutual,
					Phone:         mtproto.MakeFlagsString(do.ContactPhone),
					CloseFriend:   do.CloseFriend,
				}).To_ContactData()
			}
		})

	for _, v := range closeFriendList {
		if v == nil {
			// has hole, internal error
			return nil
		}
	}

	return closeFriendList
}

func (d *Dao) getContactListByIdList(ctx context.Context, id int64, idList []int64) []*mtproto.ContactData {
	var (
		cKeys = make([]string, 0, len(idList))
	)

	for _, v := range idList {
		cKeys = append(cKeys, genContactCacheKey(id, v))
	}

	return d.getContactListByKeyList(ctx, cKeys...)
}

func (d *Dao) getReverseContactListByIdList(ctx context.Context, id int64, idList []int64) []*mtproto.ContactData {
	var (
		cKeys = make([]string, 0, len(idList))
	)

	for _, v := range idList {
		cKeys = append(cKeys, genContactCacheKey(v, id))
	}

	return d.getContactListByKeyList(ctx, cKeys...)
}

func (d *Dao) getContactListByKeyList(ctx context.Context, cKeys ...string) []*mtproto.ContactData {
	var (
		contactList = make([]*mtproto.ContactData, len(cKeys))
	)

	if len(cKeys) == 0 {
		return contactList
	}

	d.CachedConn.QueryRows(
		ctx,
		func(ctx context.Context, conn *sqlx.DB, keys ...string) (map[string]interface{}, error) {
			noCaches := make(map[string]interface{}, len(keys))
			for _, key := range keys {
				id0, id1 := parseContactCacheKey(key)
				contact, _ := d.UserContactsDAO.SelectContact(ctx, id0, id1)
				if contact == nil {
					// return sqlc.ErrNotFound
					continue
				}
				noCaches[key] = contact
				contactList = append(contactList, mtproto.MakeTLContactData(&mtproto.ContactData{
					UserId:        contact.OwnerUserId,
					ContactUserId: contact.ContactUserId,
					FirstName:     mtproto.MakeFlagsString(contact.ContactFirstName),
					LastName:      mtproto.MakeFlagsString(contact.ContactLastName),
					MutualContact: contact.Mutual,
					Phone:         mtproto.MakeFlagsString(contact.ContactPhone),
					CloseFriend:   contact.CloseFriend,
				}).To_ContactData())
			}
			return noCaches, nil
		},
		func(k, v string) (interface{}, error) {
			var (
				contact *dataobject.UserContactsDO
			)

			err := jsonx.UnmarshalFromString(v, &contact)
			if err != nil {
				return nil, err
			}

			contactList = append(contactList, mtproto.MakeTLContactData(&mtproto.ContactData{
				UserId:        contact.OwnerUserId,
				ContactUserId: contact.ContactUserId,
				FirstName:     mtproto.MakeFlagsString(contact.ContactFirstName),
				LastName:      mtproto.MakeFlagsString(contact.ContactLastName),
				MutualContact: contact.Mutual,
				Phone:         mtproto.MakeFlagsString(contact.ContactPhone),
				CloseFriend:   contact.CloseFriend,
			}).To_ContactData())

			return contact, nil
		},
		cKeys...)

	return removeAllNil(contactList)
}
