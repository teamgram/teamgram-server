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
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dao"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

type contactItem = dao.ContactItem

//	c               *mtproto.InputContact
//	unregistered    bool  // 未注册
//	userId          int64 // 已经注册的用户ID
//	contactId       int64 // 已经注册是我的联系人
//	importContactId int64 // 已经注册的反向联系人
//}

// UserImportContacts
// user.importContacts user_id:long contacts:Vector<InputContact> = UserImportedContacts;
func (c *UserCore) UserImportContacts(in *user.TLUserImportContacts) (*user.UserImportedContacts, error) {
	var (
		contacts          = in.Contacts
		importedContacts  = make([]*mtproto.ImportedContact, 0, len(contacts))
		popularContactMap = make(map[string]*mtproto.TLPopularContact, len(contacts))
		updList           = make([]int64, 0, len(contacts))
		idList            = make([]int64, 0, len(contacts))
		//keys              = []string{
		//	genCacheUserDataCacheKey(do.OwnerUserId),
		//	genContactCacheKey(do.OwnerUserId, do.ContactUserId),
		//}
	)

	importContacts := make(map[string]*contactItem)
	// 1. 整理
	phoneList := make([]string, 0, len(contacts))
	for _, c2 := range contacts {
		phoneList = append(phoneList, c2.Phone)
		importContacts[c2.Phone] = &contactItem{Unregistered: true, C: c2}
	}

	// 2. 已注册
	registeredContacts, _ := c.svcCtx.Dao.UsersDAO.SelectUsersByPhoneList(c.ctx, phoneList)
	var contactIdList []int64

	// clear phoneList
	// phoneList = phoneList[0:0]
	for i := 0; i < len(registeredContacts); i++ {
		if c2, ok := importContacts[registeredContacts[i].Phone]; ok {
			c2.Unregistered = false
			c2.UserId = registeredContacts[i].Id
			phoneList = append(phoneList, registeredContacts[i].Phone)
			contactIdList = append(contactIdList, registeredContacts[i].Id)
		} else {
			c2.Unregistered = true
		}
	}

	if len(contactIdList) > 0 {
		// 3. 我的联系人
		myContacts, _ := c.svcCtx.Dao.UserContactsDAO.SelectListByIdList(c.ctx, in.UserId, contactIdList)
		c.Logger.Infof("myContacts - %v", myContacts)
		for i := 0; i < len(myContacts); i++ {
			if c2, ok := importContacts[myContacts[i].ContactPhone]; ok {
				c2.ContactId = myContacts[i].ContactUserId
			}
		}
	}

	if len(contactIdList) > 0 {
		// 4. 反向联系人
		importedMyContacts, _ := c.svcCtx.Dao.ImportedContactsDAO.SelectListByImportedList(c.ctx, in.UserId, contactIdList)
		c.Logger.Infof("importedMyContacts - %v", importedMyContacts)
		for i := 0; i < len(importedMyContacts); i++ {
			for _, c2 := range importContacts {
				if c2.UserId == importedMyContacts[i].ImportedUserId {
					c2.ImportContactId = c2.UserId
					break
				}
			}
		}
	}

	// clear phoneList
	phoneList = phoneList[0:0]
	for _, c2 := range importContacts {
		if c2.Unregistered {
			go func() {
				// 1. 未注册 - popular inviter
				unregisteredContactsDO := &dataobject.UnregisteredContactsDO{
					Phone:           c2.C.Phone,
					ImporterUserId:  in.UserId,
					ImportFirstName: c2.C.FirstName,
					ImportLastName:  c2.C.LastName,
				}
				c.svcCtx.Dao.UnregisteredContactsDAO.InsertOrUpdate(c.ctx, unregisteredContactsDO)
			}()

			//popularContactsDO := &dataobject.PopularContactsDO{
			//	Phone:     c2.c.Phone,
			//	Importers: 1,
			//}
			//c.dao.PopularContactsDAO.InsertOrUpdate(popularContactsDO)
			phoneList = append(phoneList, c2.C.Phone)
			popularContact := mtproto.MakeTLPopularContact(&mtproto.PopularContact{
				ClientId:  c2.C.ClientId,
				Importers: 1, // TODO(@benqi): get importers
			})
			popularContactMap[c2.C.Phone] = popularContact
			// &popularContactData{c2.c.Phone, c2.c.ClientId})
		} else {
			// 已经注册
			userContactsDO := &dataobject.UserContactsDO{
				OwnerUserId:      in.UserId,
				ContactUserId:    c2.UserId,
				ContactPhone:     c2.C.Phone,
				ContactFirstName: c2.C.FirstName,
				ContactLastName:  c2.C.LastName,
				Date2:            time.Now().Unix(),
			}

			if c2.ContactId > 0 {
				if c2.ImportContactId > 0 {
					updList = append(updList, c2.ImportContactId)
				}

				// 联系人已经存在，刷新first_name, last_name
				c.svcCtx.Dao.UserContactsDAO.UpdateContactName(c.ctx, userContactsDO.ContactFirstName,
					userContactsDO.ContactLastName,
					userContactsDO.OwnerUserId,
					userContactsDO.ContactUserId)
			} else {
				userContactsDO.IsDeleted = false
				if c2.ImportContactId > 0 {
					userContactsDO.Mutual = true

					// need update to contact
					updList = append(updList, c2.ImportContactId)

					c.svcCtx.Dao.UserContactsDAO.UpdateMutual(c.ctx, true, userContactsDO.ContactUserId, userContactsDO.OwnerUserId)
				} else {
					importedContactsDO := &dataobject.ImportedContactsDO{
						UserId:         userContactsDO.ContactUserId,
						ImportedUserId: userContactsDO.OwnerUserId,
					}
					c.svcCtx.Dao.ImportedContactsDAO.InsertOrUpdate(c.ctx, importedContactsDO)
				}
				c.svcCtx.Dao.UserContactsDAO.InsertOrUpdate(c.ctx, userContactsDO)
			}

			c.Logger.Infof("userContactsDO - %v", userContactsDO)
			c.Logger.Infof("c2 - %v", c2)

			importedContact := mtproto.MakeTLImportedContact(&mtproto.ImportedContact{
				UserId:   userContactsDO.ContactUserId,
				ClientId: c2.C.ClientId,
			})
			importedContacts = append(importedContacts, importedContact.To_ImportedContact())
			idList = append(idList, userContactsDO.ContactUserId)
		}
	}

	//
	popularContacts := make([]*mtproto.PopularContact, 0, len(phoneList))
	if len(phoneList) > 0 {
		popularDOList, _ := c.svcCtx.Dao.PopularContactsDAO.SelectImportersList(c.ctx, phoneList)
		for i := 0; i < len(popularDOList); i++ {
			if c2, ok := popularContactMap[popularDOList[i].Phone]; ok {
				c2.SetImporters(popularDOList[i].Importers + 1)
			}
		}

		for _, c2 := range popularContactMap {
			popularContacts = append(popularContacts, c2.To_PopularContact())
		}

		//go func() {
		//	// TODO:
		//	// m.PopularContactsDAO.IncreaseImportersList(context.Background(), phoneList)
		//}()
	}

	c.svcCtx.Dao.ClearContactCaches(c.ctx, in.UserId, idList...)
	users, _ := c.UserGetMutableUsers(&user.TLUserGetMutableUsers{
		Id: append(idList, in.UserId),
	})

	// importedContacts, popularContacts, updList
	rImportContacts := user.MakeTLUserImportedContacts(&user.UserImportedContacts{
		Imported:       importedContacts,
		PopularInvites: popularContacts,
		RetryContacts:  []int64{},
		Users:          users.GetUserListByIdList(in.UserId, idList...),
		UpdateIdList:   updList,
	}).To_UserImportedContacts()

	return rImportContacts, nil
}
