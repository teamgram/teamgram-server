// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
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

// Author: Benqi (wubenqi@gmail.com)

package contact

import (
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dataobject"
	"github.com/nebula-chat/chatengine/mtproto"
	"time"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/base"
)

//type contactUser struct {
//	userId int32
//	phone string
//	firstName string
//	lastName string
//}

// exclude

type contactData *dataobject.UserContactsDO
type contactLogic struct {
	selfUserId int32
	dao        *contactsDAO
}

func (m *ContactModel) MakeContactLogic(userId int32) *contactLogic {
	return &contactLogic{
		selfUserId: userId,
		dao:        m.dao,
	}
}

func findContaceByPhone(contacts []contactData, phone string) *dataobject.UserContactsDO {
	for _, c := range contacts {
		if c.ContactPhone == phone {
			return c
		}
	}
	return nil
}

// include deleted
func (c contactLogic) GetAllContactList() []contactData {
	doList := c.dao.UserContactsDAO.SelectAllUserContacts(c.selfUserId)
	contactList := make([]contactData, 0, len(doList))
	for index, _ := range doList {
		contactList = append(contactList, &doList[index])
	}
	return contactList
}

// exclude deleted
func (c contactLogic) GetContactList() []contactData {
	doList := c.dao.UserContactsDAO.SelectUserContacts(c.selfUserId)
	contactList := make([]contactData, 0, len(doList))
	for index, _ := range doList {
		contactList = append(contactList, &doList[index])
	}
	return contactList
}

func (c contactLogic) DeleteContact(deleteId int32, mutual bool) bool {
	// A 删除 B
	// 如果AB is mutual，则BA设置为非mutual

	var needUpdate = false

	c.dao.UserContactsDAO.DeleteContacts(c.selfUserId, []int32{deleteId})

	if deleteId != c.selfUserId && mutual {
		c.dao.UserContactsDAO.UpdateMutual(0, deleteId, c.selfUserId)
		needUpdate = true
	}

	return needUpdate
}

/////////////////////////////////////////////////////////////////////////////////////////
func (c contactLogic) BlockUser(blockId int32) bool {
	c.dao.UserContactsDAO.UpdateBlock(1, c.selfUserId, blockId)
	return true
}

func (c contactLogic) UnBlockUser(blockedId int32) bool {
	c.dao.UserContactsDAO.UpdateBlock(0, c.selfUserId, blockedId)
	return true
}

func (c contactLogic) GetBlockedList(offset, limit int32) []*mtproto.ContactBlocked {
	// TODO(@benqi): enable offset
	doList := c.dao.UserContactsDAO.SelectBlockedList(c.selfUserId, limit)
	bockedList := make([]*mtproto.ContactBlocked, 0, len(doList))
	for _, do := range doList {
		blocked := &mtproto.ContactBlocked{
			Constructor: mtproto.TLConstructor_CRC32_contactBlocked,
			Data2: &mtproto.ContactBlocked_Data{
				UserId: do.ContactUserId,
				Date:   do.Date2,
			},
		}
		bockedList = append(bockedList, blocked)
	}
	return bockedList
}

func (c contactLogic) SearchContacts(q string, limit int32) ([]int32, []int32) {
	contactList := c.GetContactList()
	idList := make([]int32, 0, len(contactList)+1)
	idList = append(idList, c.selfUserId)
	for _, c2 := range contactList {
		idList = append(idList, c2.ContactUserId)
	}

	// TODO(@benqi): 区分大小写

	var (
		userIdList []int32
		channelIdList []int32
	)

	// 构造模糊查询字符串
	q = q + "%"
	doList := c.dao.UsernameDAO.SearchByQueryNotIdList(2, q, idList, limit)
	founds := make([]int32, 0, len(doList))
	for i := 0; i < len(doList); i++ {
		switch doList[i].PeerType {
		case base.PEER_USER:
			userIdList = append(userIdList, doList[i].PeerId)
		case base.PEER_CHANNEL:
			channelIdList = append(channelIdList, doList[i].PeerId)
		}
		founds = append(founds, doList[i].PeerId)
	}
	return userIdList, channelIdList
}

//const (
//	_unregisted = iota
//	_noneContact
//	_contact
//	_mutualContact
//)

type contactItem struct {
	c               *mtproto.InputContact_Data
	unregistered    bool  // 未注册
	userId          int32 // 已经注册的用户ID
	contactId       int32 // 已经注册是我的联系人
	importContactId int32 // 已经注册的反向联系人
}

//type popularContactData struct {
//	phone    string
//	clientId int64
//}

func (c *contactLogic) ImportContacts(contacts []*mtproto.InputContact_Data) ([]*mtproto.ImportedContact, []*mtproto.PopularContact, []int32) {
	var (
		importedContacts = make([]*mtproto.ImportedContact, 0, len(contacts))
		popularContactMap = make(map[string]*mtproto.TLPopularContact, len(contacts))
		updList = make([]int32, 0, len(contacts))
	)

	importContacts := make(map[string]*contactItem)
	// 1. 整理
	phoneList := make([]string, 0, len(contacts))
	for _, c2 := range contacts {
		phoneList = append(phoneList, c2.Phone)
		importContacts[c2.Phone] = &contactItem{unregistered: true, c: c2}
	}

	// 2. 已注册
	registeredContacts := c.dao.UsersDAO.SelectUsersByPhoneList(phoneList)
	var contactIdList []int32

	// clear phoneList
	// phoneList = phoneList[0:0]
	for i := 0; i < len(registeredContacts); i++ {
		if c2, ok := importContacts[registeredContacts[i].Phone]; ok {
			c2.unregistered = false
			c2.userId = registeredContacts[i].Id
			phoneList = append(phoneList, registeredContacts[i].Phone)
			contactIdList = append(contactIdList, registeredContacts[i].Id)
		} else {
			c2.unregistered = true
		}
	}

	if len(contactIdList) > 0 {
		// 3. 我的联系人
		myContacts := c.dao.UserContactsDAO.SelectListByIdList(c.selfUserId, contactIdList)
		glog.Info("myContacts - ", myContacts)
		for i := 0; i < len(myContacts); i++ {
			if c2, ok := importContacts[myContacts[i].ContactPhone]; ok {
				c2.contactId = myContacts[i].ContactUserId
			}
		}
	}

	if len(contactIdList) > 0 {
		// 4. 反向联系人
		importedMyContacts := c.dao.ImportedContactsDAO.SelectListByImportedList(c.selfUserId, contactIdList)
		glog.Info("importedMyContacts - ", importedMyContacts)
		for i := 0; i < len(importedMyContacts); i++ {
			for _, c2 := range importContacts {
				if c2.userId == importedMyContacts[i].ImportedUserId {
					c2.importContactId = c2.userId
					break
				}
			}
		}
	}

	// clear phoneList
	phoneList = phoneList[0:0]
	for _, c2 := range importContacts {
		if c2.unregistered {
			// 1. 未注册 - popular inviter
			unregisteredContactsDO := &dataobject.UnregisteredContactsDO{
				Phone:           c2.c.Phone,
				ImporterUserId:  c.selfUserId,
				ImportFirstName: c2.c.FirstName,
				ImportLastName:  c2.c.LastName,
			}
			c.dao.UnregisteredContactsDAO.InsertOrUpdate(unregisteredContactsDO)

			//popularContactsDO := &dataobject.PopularContactsDO{
			//	Phone:     c2.c.Phone,
			//	Importers: 1,
			//}
			//c.dao.PopularContactsDAO.InsertOrUpdate(popularContactsDO)
			phoneList = append(phoneList, c2.c.Phone)
			popularContact := &mtproto.TLPopularContact{Data2: &mtproto.PopularContact_Data{
				ClientId:  c2.c.ClientId,
				Importers: 1, // TODO(@benqi): get importers
			}}
			popularContactMap[c2.c.Phone] = popularContact
			// &popularContactData{c2.c.Phone, c2.c.ClientId})
		} else {
			// 已经注册
			userContactsDO := &dataobject.UserContactsDO{
				OwnerUserId:      c.selfUserId,
				ContactUserId:    c2.userId,
				ContactPhone:     c2.c.Phone,
				ContactFirstName: c2.c.FirstName,
				ContactLastName:  c2.c.LastName,
				Date2:            int32(time.Now().Unix()),
			}

			if c2.contactId > 0 {
				if c2.importContactId > 0 {
					updList = append(updList, c2.importContactId)
				}

				// 联系人已经存在，刷新first_name, last_name
				c.dao.UserContactsDAO.UpdateContactName(userContactsDO.ContactFirstName,
					userContactsDO.ContactLastName,
					userContactsDO.OwnerUserId,
					userContactsDO.ContactUserId)
			} else {
				userContactsDO.IsDeleted = 0
				if c2.importContactId > 0 {
					userContactsDO.Mutual = 1

					// need update to contact
					updList = append(updList, c2.importContactId)

					c.dao.UserContactsDAO.UpdateMutual(1, userContactsDO.ContactUserId, userContactsDO.OwnerUserId)
				} else {
					importedContactsDO := &dataobject.ImportedContactsDO{
						UserId:         userContactsDO.ContactUserId,
						ImportedUserId: userContactsDO.OwnerUserId,
					}
					c.dao.ImportedContactsDAO.InsertOrUpdate(importedContactsDO)
				}
				c.dao.UserContactsDAO.InsertOrUpdate(userContactsDO)
			}

			glog.Info("userContactsDO - ", userContactsDO)
			glog.Info("c2 - ", c2)

			importedContact := &mtproto.TLImportedContact{Data2: &mtproto.ImportedContact_Data{
				UserId:   userContactsDO.ContactUserId,
				ClientId: c2.c.ClientId,
			}}
			importedContacts = append(importedContacts, importedContact.To_ImportedContact())
		}
	}

	//
	popularContacts := make([]*mtproto.PopularContact, 0, len(phoneList))
	if len(phoneList) > 0 {
		popularDOList := c.dao.PopularContactsDAO.SelectImportersList(phoneList)
		for i := 0; i < len(popularDOList); i++ {
			if c2, ok := popularContactMap[popularDOList[i].Phone]; ok {
				c2.SetImporters(popularDOList[i].Importers + 1)
			}
		}

		for _, c2 := range popularContactMap {
			popularContacts = append(popularContacts, c2.To_PopularContact())
		}

		c.dao.PopularContactsDAO.IncreaseImportersList(phoneList)
	}

	return importedContacts, popularContacts, updList
}
