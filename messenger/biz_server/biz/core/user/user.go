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

package user

import (
	"math/rand"
	"encoding/hex"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/pkg/util"
	"github.com/nebula-chat/chatengine/pkg/crypto"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dataobject"
)

const (
	USER_TYPE_REGULAR 	= 0		// 普通用户
	USER_TYPE_SERVICE 	= 1		// 服务通知(Telegram等)
	USER_TYPE_BOT 		= 2		// BOT
	USER_TYPE_DELETED 	= 3		// 已经删除用户
	USER_TYPE_UNKNOWN 	= 4		// 未知用户
)

type userItem struct {
	*UserModel
	*dataobject.UsersDO
	*dataobject.UserPresencesDO
	*dataobject.UserContactsDO
	*dataobject.BotsDO
	selfUserId int32
}

func (m *UserModel) makeUserItem(selfId, userId int32) *userItem {
	usersDO := m.dao.UsersDAO.SelectById(userId)
	if usersDO == nil {
		return nil
	}

	return m.makeUserItemByUsersDO(selfId, usersDO)
}

func (m *UserModel) makeUserItemByUsersDO(selfId int32, usersDO *dataobject.UsersDO) *userItem {
	u := &userItem{
		UserModel:       m,
		UsersDO:		 usersDO,
		UserPresencesDO: m.dao.UserPresencesDAO.Select(usersDO.Id),
		// UserContactsDO:  m.dao.UserContactsDAO.SelectUserContact(selfId, usersDO.Id),
		BotsDO:          m.dao.BotsDAO.Select(usersDO.Id),
		selfUserId:      selfId,
	}

	if selfId != usersDO.Id {
		u.UserContactsDO = m.dao.UserContactsDAO.SelectByContactId(selfId, usersDO.Id)
	}

	return u
}

func (m *userItem) IsSelf() bool {
	if m == nil || m.UsersDO == nil {
		return false
	}
	return m.selfUserId == m.UsersDO.Id
}

func (m *userItem) ToUser() *mtproto.User {
	if m.UsersDO == nil {
		return nil
	}

	user := mtproto.NewTLUser()

	user.SetId(m.UsersDO.Id)

	// TODO(@benqi): access_hash algorithm
	user.SetAccessHash(m.UsersDO.AccessHash)

	user.SetFirstName(m.UsersDO.FirstName)
	user.SetLastName(m.UsersDO.LastName)
	user.SetVerified(util.Int8ToBool(m.UsersDO.Verified))
	user.SetMin(util.Int8ToBool(m.UsersDO.Min))
	// user.SetBot(util.Int8ToBool(m.UsersDO.IsBot))
	user.SetRestricted(util.Int8ToBool(m.UsersDO.Restricted))
	user.SetRestrictionReason(m.UsersDO.RestrictionReason)
	user.SetDeleted(util.Int8ToBool(m.UsersDO.Deleted))

	// set photo
	photoId := m.UserModel.GetDefaultUserPhotoID(m.UsersDO.Id)
	if photoId == 0 {
		user.SetPhoto(mtproto.NewTLUserProfilePhotoEmpty().To_UserProfilePhoto())
	} else {
		user.SetPhoto(m.UserModel.photoCallback.GetUserProfilePhoto(photoId))
	}

	// TODO(@benqi): verified
	// TODO(@benqi): min
	// TODO(@benqi): restricted
	// TODO(@benqi): restriction_reason
	// TODO(@benqi): deleted

	if m.IsSelf() {
		user.SetSelf(true)
		user.SetContact(true)
		user.SetMutualContact(true)
		user.SetPhone(m.UsersDO.Phone)

		user.SetStatus(makeUserStatusOnline())
	} else {
		user.SetSelf(false)
		if m.UserContactsDO != nil {
			if m.UserContactsDO.IsDeleted == 0 {
				user.SetContact(true)
				user.SetMutualContact(util.Int8ToBool(m.UserContactsDO.Mutual))
				user.SetFirstName(m.UserContactsDO.ContactFirstName)
				user.SetLastName(m.UserContactsDO.ContactLastName)
			} else {
				user.SetContact(false)
				user.SetMutualContact(false)
			}

			user.SetPhone(m.UsersDO.Phone)
		} else {
			user.SetContact(false)
			user.SetMutualContact(false)
			// user.SetFirstName(m.UsersDO.FirstName)
			// user.SetLastName(m.UsersDO.LastName)
		}

		if m.UserPresencesDO != nil {
			user.SetStatus(makeUserStatus(m.UserPresencesDO))
		}
	}

	user.SetUsername(m.usernameCallback.GetAccountUsername(m.UsersDO.Id))

	if m.UsersDO.IsBot == 1 && m.BotsDO != nil {
		user.SetBot(true)
		user.SetBotChatHistory(util.Int8ToBool(m.BotsDO.BotChatHistory))
		user.SetBotNochats(util.Int8ToBool(m.BotsDO.BotNochats))
		user.SetBotInlineGeo(util.Int8ToBool(m.BotsDO.BotInlineGeo))
		user.SetBotInfoVersion(m.BotsDO.BotInfoVersion)
		user.SetBotInlinePlaceholder(m.BotsDO.BotInlinePlaceholder)
	} else {
		user.SetBot(false)
	}

	user.SetLangCode("")

	return user.To_User()
}

func (m *UserModel) GetUserById(selfId int32, userId int32) *mtproto.User {
	u := m.makeUserItem(selfId, userId)
	return u.ToUser()
}

func (m *UserModel) getUserListByUsersDOList(selfUserId int32, userIdList []int32, usersDOList []dataobject.UsersDO) (users []*mtproto.User) {
	var userMap = make(map[int32]*userItem, len(usersDOList))
	for i := 0; i < len(usersDOList); i++ {
		userMap[usersDOList[i].Id] = &userItem{UserModel: m, UsersDO: &usersDOList[i], selfUserId: selfUserId}
	}

	conatcts := m.dao.UserContactsDAO.SelectListByIdList(selfUserId, userIdList)
	for i := 0; i < len(conatcts); i++ {
		if u, ok := userMap[conatcts[i].ContactUserId]; ok {
			u.UserContactsDO = &conatcts[i]
		}
	}

	statuses := m.dao.UserPresencesDAO.SelectList(userIdList)
	for i := 0; i < len(statuses); i++ {
		if u, ok := userMap[statuses[i].UserId]; ok {
			u.UserPresencesDO = &statuses[i]
		}
	}

	bots := m.dao.BotsDAO.SelectByIdList(userIdList)
	for i := 0; i < len(bots); i++ {
		if u, ok := userMap[bots[i].BotId]; ok {
			u.BotsDO = &bots[i]
		}
	}

	for _, v := range userMap {
		users = append(users, v.ToUser())
	}

	return
}

func (m *UserModel) GetUserListByIdList(selfUserId int32, userIdList []int32) (users []*mtproto.User) {
	if len(userIdList) == 0 {
		users = []*mtproto.User{}
		return
	}

	// TODO(@benqi):  需要优化，makeUserDataByDO需要查询用户状态以及获取Mutual和Contact状态信息而导致多次查询
	usersDOList := m.dao.UsersDAO.SelectUsersByIdList(userIdList)
	if len(usersDOList) == 0 {
		users = []*mtproto.User{}
		return
	}

	users = m.getUserListByUsersDOList(selfUserId, userIdList, usersDOList)
	return
}

func CheckUserAccessHash(id int32, hash int64) bool {
	return true
}

func (m *UserModel) CheckPhoneNumberExist(phoneNumber string) bool {
	return nil != m.dao.UsersDAO.SelectByPhoneNumber(phoneNumber)
}

func (m *UserModel) CreateNewUser(phoneNumber, countryCode, firstName, lastName string) *mtproto.TLUser {
	do := &dataobject.UsersDO{
		AccessHash:  rand.Int63(),
		Phone:       phoneNumber,
		FirstName:   firstName,
		LastName:    lastName,
		CountryCode: countryCode,
	}
	do.Id = int32(m.dao.UsersDAO.Insert(do))
	user := &mtproto.TLUser{Data2: &mtproto.User_Data{
		Id:            do.Id,
		Self:          true,
		Contact:       true,
		MutualContact: true,
		AccessHash:    do.AccessHash,
		FirstName:     do.FirstName,
		LastName:      do.LastName,
		Username:      do.Username,
		Phone:         phoneNumber,
		// TODO(@benqi): Load from db
		Photo:  mtproto.NewTLUserProfilePhotoEmpty().To_UserProfilePhoto(),
		Status: makeUserStatusOnline(),
	}}
	return user
}

func (m *UserModel) CreateNewUserPassword(userId int32) {
	// gen server_nonce
	do := &dataobject.UserPasswordsDO{
		UserId:     userId,
		ServerSalt: hex.EncodeToString(crypto.GenerateNonce(8)),
	}
	m.dao.UserPasswordsDAO.Insert(do)
}

func (m *UserModel) CheckAccessHashByUserId(userId int32, accessHash int64) bool {
	params := map[string]interface{}{
		"id":          userId,
		"access_hash": accessHash,
	}
	return m.dao.CommonDAO.CheckExists("users", params)
}

func (m *UserModel) GetCountryCodeByUser(userId int32) string {
	do := m.dao.UsersDAO.SelectCountryCode(userId)
	if do == nil {
		return ""
	} else {
		return do.CountryCode
	}
}

func (m *UserModel) GetUserByPhoneNumber(selfId int32, phoneNumber string) *mtproto.User {
	do := m.dao.UsersDAO.SelectByPhoneNumber(phoneNumber)
	if do == nil {
		return nil
	}

	u := m.makeUserItemByUsersDO(selfId, do)
	return u.ToUser()
}

func (m *UserModel) GetUserListByPhoneNumberList(selfId int32, phoneNumberList []string) []*mtproto.User {
	usersDOList := m.dao.UsersDAO.SelectUsersByPhoneList(phoneNumberList)
	if len(usersDOList) == 0 {
		return []*mtproto.User{}
	}

	userIdList := make([]int32, 0, len(usersDOList))
	for i := 0; i < len(usersDOList); i++ {
		userIdList = append(userIdList, usersDOList[i].Id)
	}

	return m.getUserListByUsersDOList(selfId, userIdList, usersDOList)
}

func (m *UserModel) GetUserByUsername(selfId int32, username string) *mtproto.User {
	do := m.dao.UsersDAO.SelectByUsername(username)
	if do == nil {
		return nil
	}

	u := m.makeUserItemByUsersDO(selfId, do)
	return u.ToUser()
}

func (m *UserModel) GetMyUserByPhoneNumber(phoneNumber string) *mtproto.User {
	do := m.dao.UsersDAO.SelectByPhoneNumber(phoneNumber)
	if do == nil {
		return nil
	}

	u := m.makeUserItemByUsersDO(do.Id, do)
	return u.ToUser()
}

