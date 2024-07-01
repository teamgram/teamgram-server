/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package user

import (
	"reflect"

	"github.com/teamgram/proto/mtproto"
)

var _ *mtproto.Bool

type newRPCReplyFunc func() interface{}

type RPCContextTuple struct {
	Method       string
	NewReplyFunc newRPCReplyFunc
}

var rpcContextRegisters = map[string]RPCContextTuple{
	"TLUserGetLastSeens":                 RPCContextTuple{"/mtproto.RPCUser/user_getLastSeens", func() interface{} { return new(Vector_LastSeenData) }},
	"TLUserUpdateLastSeen":               RPCContextTuple{"/mtproto.RPCUser/user_updateLastSeen", func() interface{} { return new(mtproto.Bool) }},
	"TLUserGetLastSeen":                  RPCContextTuple{"/mtproto.RPCUser/user_getLastSeen", func() interface{} { return new(LastSeenData) }},
	"TLUserGetImmutableUser":             RPCContextTuple{"/mtproto.RPCUser/user_getImmutableUser", func() interface{} { return new(mtproto.ImmutableUser) }},
	"TLUserGetMutableUsers":              RPCContextTuple{"/mtproto.RPCUser/user_getMutableUsers", func() interface{} { return new(Vector_ImmutableUser) }},
	"TLUserGetImmutableUserByPhone":      RPCContextTuple{"/mtproto.RPCUser/user_getImmutableUserByPhone", func() interface{} { return new(mtproto.ImmutableUser) }},
	"TLUserGetImmutableUserByToken":      RPCContextTuple{"/mtproto.RPCUser/user_getImmutableUserByToken", func() interface{} { return new(mtproto.ImmutableUser) }},
	"TLUserSetAccountDaysTTL":            RPCContextTuple{"/mtproto.RPCUser/user_setAccountDaysTTL", func() interface{} { return new(mtproto.Bool) }},
	"TLUserGetAccountDaysTTL":            RPCContextTuple{"/mtproto.RPCUser/user_getAccountDaysTTL", func() interface{} { return new(mtproto.AccountDaysTTL) }},
	"TLUserGetNotifySettings":            RPCContextTuple{"/mtproto.RPCUser/user_getNotifySettings", func() interface{} { return new(mtproto.PeerNotifySettings) }},
	"TLUserGetNotifySettingsList":        RPCContextTuple{"/mtproto.RPCUser/user_getNotifySettingsList", func() interface{} { return new(Vector_PeerPeerNotifySettings) }},
	"TLUserSetNotifySettings":            RPCContextTuple{"/mtproto.RPCUser/user_setNotifySettings", func() interface{} { return new(mtproto.Bool) }},
	"TLUserResetNotifySettings":          RPCContextTuple{"/mtproto.RPCUser/user_resetNotifySettings", func() interface{} { return new(mtproto.Bool) }},
	"TLUserGetAllNotifySettings":         RPCContextTuple{"/mtproto.RPCUser/user_getAllNotifySettings", func() interface{} { return new(Vector_PeerPeerNotifySettings) }},
	"TLUserGetGlobalPrivacySettings":     RPCContextTuple{"/mtproto.RPCUser/user_getGlobalPrivacySettings", func() interface{} { return new(mtproto.GlobalPrivacySettings) }},
	"TLUserSetGlobalPrivacySettings":     RPCContextTuple{"/mtproto.RPCUser/user_setGlobalPrivacySettings", func() interface{} { return new(mtproto.Bool) }},
	"TLUserGetPrivacy":                   RPCContextTuple{"/mtproto.RPCUser/user_getPrivacy", func() interface{} { return new(Vector_PrivacyRule) }},
	"TLUserSetPrivacy":                   RPCContextTuple{"/mtproto.RPCUser/user_setPrivacy", func() interface{} { return new(mtproto.Bool) }},
	"TLUserCheckPrivacy":                 RPCContextTuple{"/mtproto.RPCUser/user_checkPrivacy", func() interface{} { return new(mtproto.Bool) }},
	"TLUserAddPeerSettings":              RPCContextTuple{"/mtproto.RPCUser/user_addPeerSettings", func() interface{} { return new(mtproto.Bool) }},
	"TLUserGetPeerSettings":              RPCContextTuple{"/mtproto.RPCUser/user_getPeerSettings", func() interface{} { return new(mtproto.PeerSettings) }},
	"TLUserDeletePeerSettings":           RPCContextTuple{"/mtproto.RPCUser/user_deletePeerSettings", func() interface{} { return new(mtproto.Bool) }},
	"TLUserChangePhone":                  RPCContextTuple{"/mtproto.RPCUser/user_changePhone", func() interface{} { return new(mtproto.Bool) }},
	"TLUserCreateNewUser":                RPCContextTuple{"/mtproto.RPCUser/user_createNewUser", func() interface{} { return new(mtproto.ImmutableUser) }},
	"TLUserDeleteUser":                   RPCContextTuple{"/mtproto.RPCUser/user_deleteUser", func() interface{} { return new(mtproto.Bool) }},
	"TLUserBlockPeer":                    RPCContextTuple{"/mtproto.RPCUser/user_blockPeer", func() interface{} { return new(mtproto.Bool) }},
	"TLUserUnBlockPeer":                  RPCContextTuple{"/mtproto.RPCUser/user_unBlockPeer", func() interface{} { return new(mtproto.Bool) }},
	"TLUserBlockedByUser":                RPCContextTuple{"/mtproto.RPCUser/user_blockedByUser", func() interface{} { return new(mtproto.Bool) }},
	"TLUserIsBlockedByUser":              RPCContextTuple{"/mtproto.RPCUser/user_isBlockedByUser", func() interface{} { return new(mtproto.Bool) }},
	"TLUserCheckBlockUserList":           RPCContextTuple{"/mtproto.RPCUser/user_checkBlockUserList", func() interface{} { return new(Vector_Long) }},
	"TLUserGetBlockedList":               RPCContextTuple{"/mtproto.RPCUser/user_getBlockedList", func() interface{} { return new(Vector_PeerBlocked) }},
	"TLUserGetContactSignUpNotification": RPCContextTuple{"/mtproto.RPCUser/user_getContactSignUpNotification", func() interface{} { return new(mtproto.Bool) }},
	"TLUserSetContactSignUpNotification": RPCContextTuple{"/mtproto.RPCUser/user_setContactSignUpNotification", func() interface{} { return new(mtproto.Bool) }},
	"TLUserGetContentSettings":           RPCContextTuple{"/mtproto.RPCUser/user_getContentSettings", func() interface{} { return new(mtproto.Account_ContentSettings) }},
	"TLUserSetContentSettings":           RPCContextTuple{"/mtproto.RPCUser/user_setContentSettings", func() interface{} { return new(mtproto.Bool) }},
	"TLUserDeleteContact":                RPCContextTuple{"/mtproto.RPCUser/user_deleteContact", func() interface{} { return new(mtproto.Bool) }},
	"TLUserGetContactList":               RPCContextTuple{"/mtproto.RPCUser/user_getContactList", func() interface{} { return new(Vector_ContactData) }},
	"TLUserGetContactIdList":             RPCContextTuple{"/mtproto.RPCUser/user_getContactIdList", func() interface{} { return new(Vector_Long) }},
	"TLUserGetContact":                   RPCContextTuple{"/mtproto.RPCUser/user_getContact", func() interface{} { return new(mtproto.ContactData) }},
	"TLUserAddContact":                   RPCContextTuple{"/mtproto.RPCUser/user_addContact", func() interface{} { return new(mtproto.Bool) }},
	"TLUserCheckContact":                 RPCContextTuple{"/mtproto.RPCUser/user_checkContact", func() interface{} { return new(mtproto.Bool) }},
	"TLUserGetImportersByPhone":          RPCContextTuple{"/mtproto.RPCUser/user_getImportersByPhone", func() interface{} { return new(Vector_InputContact) }},
	"TLUserDeleteImportersByPhone":       RPCContextTuple{"/mtproto.RPCUser/user_deleteImportersByPhone", func() interface{} { return new(mtproto.Bool) }},
	"TLUserImportContacts":               RPCContextTuple{"/mtproto.RPCUser/user_importContacts", func() interface{} { return new(UserImportedContacts) }},
	"TLUserGetCountryCode":               RPCContextTuple{"/mtproto.RPCUser/user_getCountryCode", func() interface{} { return new(mtproto.String) }},
	"TLUserUpdateAbout":                  RPCContextTuple{"/mtproto.RPCUser/user_updateAbout", func() interface{} { return new(mtproto.Bool) }},
	"TLUserUpdateFirstAndLastName":       RPCContextTuple{"/mtproto.RPCUser/user_updateFirstAndLastName", func() interface{} { return new(mtproto.Bool) }},
	"TLUserUpdateVerified":               RPCContextTuple{"/mtproto.RPCUser/user_updateVerified", func() interface{} { return new(mtproto.Bool) }},
	"TLUserUpdateUsername":               RPCContextTuple{"/mtproto.RPCUser/user_updateUsername", func() interface{} { return new(mtproto.Bool) }},
	"TLUserUpdateProfilePhoto":           RPCContextTuple{"/mtproto.RPCUser/user_updateProfilePhoto", func() interface{} { return new(mtproto.Int64) }},
	"TLUserDeleteProfilePhotos":          RPCContextTuple{"/mtproto.RPCUser/user_deleteProfilePhotos", func() interface{} { return new(mtproto.Int64) }},
	"TLUserGetProfilePhotos":             RPCContextTuple{"/mtproto.RPCUser/user_getProfilePhotos", func() interface{} { return new(Vector_Long) }},
	"TLUserSetBotCommands":               RPCContextTuple{"/mtproto.RPCUser/user_setBotCommands", func() interface{} { return new(mtproto.Bool) }},
	"TLUserIsBot":                        RPCContextTuple{"/mtproto.RPCUser/user_isBot", func() interface{} { return new(mtproto.Bool) }},
	"TLUserGetBotInfo":                   RPCContextTuple{"/mtproto.RPCUser/user_getBotInfo", func() interface{} { return new(mtproto.BotInfo) }},
	"TLUserCheckBots":                    RPCContextTuple{"/mtproto.RPCUser/user_checkBots", func() interface{} { return new(Vector_Long) }},
	"TLUserGetFullUser":                  RPCContextTuple{"/mtproto.RPCUser/user_getFullUser", func() interface{} { return new(mtproto.Users_UserFull) }},
	"TLUserUpdateEmojiStatus":            RPCContextTuple{"/mtproto.RPCUser/user_updateEmojiStatus", func() interface{} { return new(mtproto.Bool) }},
	"TLUserGetUserDataById":              RPCContextTuple{"/mtproto.RPCUser/user_getUserDataById", func() interface{} { return new(mtproto.UserData) }},
	"TLUserGetUserDataListByIdList":      RPCContextTuple{"/mtproto.RPCUser/user_getUserDataListByIdList", func() interface{} { return new(Vector_UserData) }},
	"TLUserGetUserDataByToken":           RPCContextTuple{"/mtproto.RPCUser/user_getUserDataByToken", func() interface{} { return new(mtproto.UserData) }},
	"TLUserSearch":                       RPCContextTuple{"/mtproto.RPCUser/user_search", func() interface{} { return new(UsersFound) }},
	"TLUserUpdateBotData":                RPCContextTuple{"/mtproto.RPCUser/user_updateBotData", func() interface{} { return new(mtproto.Bool) }},
	"TLUserGetImmutableUserV2":           RPCContextTuple{"/mtproto.RPCUser/user_getImmutableUserV2", func() interface{} { return new(mtproto.ImmutableUser) }},
	"TLUserGetMutableUsersV2":            RPCContextTuple{"/mtproto.RPCUser/user_getMutableUsersV2", func() interface{} { return new(mtproto.MutableUsers) }},
	"TLUserCreateNewTestUser":            RPCContextTuple{"/mtproto.RPCUser/user_createNewTestUser", func() interface{} { return new(mtproto.ImmutableUser) }},
	"TLUserEditCloseFriends":             RPCContextTuple{"/mtproto.RPCUser/user_editCloseFriends", func() interface{} { return new(mtproto.Bool) }},
	"TLUserSetStoriesMaxId":              RPCContextTuple{"/mtproto.RPCUser/user_setStoriesMaxId", func() interface{} { return new(mtproto.Bool) }},
	"TLUserSetColor":                     RPCContextTuple{"/mtproto.RPCUser/user_setColor", func() interface{} { return new(mtproto.Bool) }},
	"TLUserUpdateBirthday":               RPCContextTuple{"/mtproto.RPCUser/user_updateBirthday", func() interface{} { return new(mtproto.Bool) }},
	"TLUserGetBirthdays":                 RPCContextTuple{"/mtproto.RPCUser/user_getBirthdays", func() interface{} { return new(Vector_ContactBirthday) }},
	"TLUserSetStoriesHidden":             RPCContextTuple{"/mtproto.RPCUser/user_setStoriesHidden", func() interface{} { return new(mtproto.Bool) }},
}

func FindRPCContextTuple(t interface{}) *RPCContextTuple {
	rt := reflect.TypeOf(t)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	m, ok := rpcContextRegisters[rt.Name()]
	if !ok {
		// log.Errorf("Can't find name: %s", rt.Name())
		return nil
	}
	return &m
}

func GetRPCContextRegisters() map[string]RPCContextTuple {
	return rpcContextRegisters
}
