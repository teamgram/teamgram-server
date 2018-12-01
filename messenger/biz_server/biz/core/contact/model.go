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
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dao"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dao/mysql_dao"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dataobject"
)

type ImportedContactData struct {
	UserId        int32
	Importers     int32
	MutualUpdated bool
}

type InputContactData struct {
	UserId    int32
	Phone     string
	FirstName string
	LastName  string
}

type DeleteResult struct {
	UserId int32
	State  int32
}

type contactsDAO struct {
	*mysql_dao.UserContactsDAO
	*mysql_dao.UsersDAO
	*mysql_dao.UnregisteredContactsDAO
	*mysql_dao.PopularContactsDAO
	*mysql_dao.ImportedContactsDAO
	*mysql_dao.PhoneBooksDAO
	*mysql_dao.UsernameDAO
}

type ContactModel struct {
	dao *contactsDAO
}

func (m *ContactModel) InstallModel() {
	m.dao.UserContactsDAO = dao.GetUserContactsDAO(dao.DB_MASTER)
	m.dao.UsersDAO = dao.GetUsersDAO(dao.DB_MASTER)
	m.dao.UnregisteredContactsDAO = dao.GetUnregisteredContactsDAO(dao.DB_MASTER)
	m.dao.PopularContactsDAO = dao.GetPopularContactsDAO(dao.DB_MASTER)
	m.dao.ImportedContactsDAO = dao.GetImportedContactsDAO(dao.DB_MASTER)
	m.dao.PhoneBooksDAO = dao.GetPhoneBooksDAO(dao.DB_MASTER)
	m.dao.UsernameDAO = dao.GetUsernameDAO(dao.DB_MASTER)
}

func (m *ContactModel) RegisterCallback(cb interface{}) {
}

func (m *ContactModel) CheckContactAndMutualByUserId(selfId, contactId int32) (bool, bool) {
	do := m.dao.UserContactsDAO.SelectContact(selfId, contactId)
	if do == nil {
		return false, false
	} else {
		return true, do.Mutual == 1
	}
}

// Impl ContactCallback
func (m *ContactModel) GetContactAndMutual(selfUserId, id int32) (bool, bool) {
	return m.CheckContactAndMutualByUserId(selfUserId, id)
}

func (m *ContactModel) BackupPhoneBooks(authKeyId int64, contacts []*mtproto.InputContact_Data) {
	do := &dataobject.PhoneBooksDO{
		AuthKeyId: authKeyId,
	}
	for _, c := range contacts {
		// TODO(@benqi): only phone.
		// filter addContact
		if c.GetClientId() == 0 {
			continue
		}

		do.ClientId = c.GetClientId()
		do.Phone = c.GetPhone()
		c.FirstName = c.GetFirstName()
		c.LastName = c.GetLastName()
		m.dao.PhoneBooksDAO.InsertOrUpdate(do)
	}
}

// check A block B
func (m *ContactModel) CheckBlockUser(selfUserId, id int32) bool {
	return m.dao.UserContactsDAO.SelectBlocked(selfUserId, id) != nil
}

func init() {
	core.RegisterCoreModel(&ContactModel{dao: &contactsDAO{}})
}
