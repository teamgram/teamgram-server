/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package userclient

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type UserClient interface {
	UserGetLastSeens(ctx context.Context, in *user.TLUserGetLastSeens) (*user.Vector_LastSeenData, error)
	UserUpdateLastSeen(ctx context.Context, in *user.TLUserUpdateLastSeen) (*mtproto.Bool, error)
	UserGetLastSeen(ctx context.Context, in *user.TLUserGetLastSeen) (*user.LastSeenData, error)
	UserGetImmutableUser(ctx context.Context, in *user.TLUserGetImmutableUser) (*mtproto.ImmutableUser, error)
	UserGetMutableUsers(ctx context.Context, in *user.TLUserGetMutableUsers) (*user.Vector_ImmutableUser, error)
	UserGetImmutableUserByPhone(ctx context.Context, in *user.TLUserGetImmutableUserByPhone) (*mtproto.ImmutableUser, error)
	UserGetImmutableUserByToken(ctx context.Context, in *user.TLUserGetImmutableUserByToken) (*mtproto.ImmutableUser, error)
	UserSetAccountDaysTTL(ctx context.Context, in *user.TLUserSetAccountDaysTTL) (*mtproto.Bool, error)
	UserGetAccountDaysTTL(ctx context.Context, in *user.TLUserGetAccountDaysTTL) (*mtproto.AccountDaysTTL, error)
	UserGetNotifySettings(ctx context.Context, in *user.TLUserGetNotifySettings) (*mtproto.PeerNotifySettings, error)
	UserGetNotifySettingsList(ctx context.Context, in *user.TLUserGetNotifySettingsList) (*user.Vector_PeerPeerNotifySettings, error)
	UserSetNotifySettings(ctx context.Context, in *user.TLUserSetNotifySettings) (*mtproto.Bool, error)
	UserResetNotifySettings(ctx context.Context, in *user.TLUserResetNotifySettings) (*mtproto.Bool, error)
	UserGetAllNotifySettings(ctx context.Context, in *user.TLUserGetAllNotifySettings) (*user.Vector_PeerPeerNotifySettings, error)
	UserGetGlobalPrivacySettings(ctx context.Context, in *user.TLUserGetGlobalPrivacySettings) (*mtproto.GlobalPrivacySettings, error)
	UserSetGlobalPrivacySettings(ctx context.Context, in *user.TLUserSetGlobalPrivacySettings) (*mtproto.Bool, error)
	UserGetPrivacy(ctx context.Context, in *user.TLUserGetPrivacy) (*user.Vector_PrivacyRule, error)
	UserSetPrivacy(ctx context.Context, in *user.TLUserSetPrivacy) (*mtproto.Bool, error)
	UserCheckPrivacy(ctx context.Context, in *user.TLUserCheckPrivacy) (*mtproto.Bool, error)
	UserAddPeerSettings(ctx context.Context, in *user.TLUserAddPeerSettings) (*mtproto.Bool, error)
	UserGetPeerSettings(ctx context.Context, in *user.TLUserGetPeerSettings) (*mtproto.PeerSettings, error)
	UserDeletePeerSettings(ctx context.Context, in *user.TLUserDeletePeerSettings) (*mtproto.Bool, error)
	UserChangePhone(ctx context.Context, in *user.TLUserChangePhone) (*mtproto.Bool, error)
	UserCreateNewUser(ctx context.Context, in *user.TLUserCreateNewUser) (*mtproto.ImmutableUser, error)
	UserDeleteUser(ctx context.Context, in *user.TLUserDeleteUser) (*mtproto.Bool, error)
	UserBlockPeer(ctx context.Context, in *user.TLUserBlockPeer) (*mtproto.Bool, error)
	UserUnBlockPeer(ctx context.Context, in *user.TLUserUnBlockPeer) (*mtproto.Bool, error)
	UserBlockedByUser(ctx context.Context, in *user.TLUserBlockedByUser) (*mtproto.Bool, error)
	UserIsBlockedByUser(ctx context.Context, in *user.TLUserIsBlockedByUser) (*mtproto.Bool, error)
	UserCheckBlockUserList(ctx context.Context, in *user.TLUserCheckBlockUserList) (*user.Vector_Long, error)
	UserGetBlockedList(ctx context.Context, in *user.TLUserGetBlockedList) (*user.Vector_PeerBlocked, error)
	UserGetContactSignUpNotification(ctx context.Context, in *user.TLUserGetContactSignUpNotification) (*mtproto.Bool, error)
	UserSetContactSignUpNotification(ctx context.Context, in *user.TLUserSetContactSignUpNotification) (*mtproto.Bool, error)
	UserGetContentSettings(ctx context.Context, in *user.TLUserGetContentSettings) (*mtproto.Account_ContentSettings, error)
	UserSetContentSettings(ctx context.Context, in *user.TLUserSetContentSettings) (*mtproto.Bool, error)
	UserDeleteContact(ctx context.Context, in *user.TLUserDeleteContact) (*mtproto.Bool, error)
	UserGetContactList(ctx context.Context, in *user.TLUserGetContactList) (*user.Vector_ContactData, error)
	UserGetContactIdList(ctx context.Context, in *user.TLUserGetContactIdList) (*user.Vector_Long, error)
	UserGetContact(ctx context.Context, in *user.TLUserGetContact) (*mtproto.ContactData, error)
	UserAddContact(ctx context.Context, in *user.TLUserAddContact) (*mtproto.Bool, error)
	UserCheckContact(ctx context.Context, in *user.TLUserCheckContact) (*mtproto.Bool, error)
	UserGetImportersByPhone(ctx context.Context, in *user.TLUserGetImportersByPhone) (*user.Vector_InputContact, error)
	UserDeleteImportersByPhone(ctx context.Context, in *user.TLUserDeleteImportersByPhone) (*mtproto.Bool, error)
	UserImportContacts(ctx context.Context, in *user.TLUserImportContacts) (*user.UserImportedContacts, error)
	UserGetCountryCode(ctx context.Context, in *user.TLUserGetCountryCode) (*mtproto.String, error)
	UserUpdateAbout(ctx context.Context, in *user.TLUserUpdateAbout) (*mtproto.Bool, error)
	UserUpdateFirstAndLastName(ctx context.Context, in *user.TLUserUpdateFirstAndLastName) (*mtproto.Bool, error)
	UserUpdateVerified(ctx context.Context, in *user.TLUserUpdateVerified) (*mtproto.Bool, error)
	UserUpdateUsername(ctx context.Context, in *user.TLUserUpdateUsername) (*mtproto.Bool, error)
	UserUpdateProfilePhoto(ctx context.Context, in *user.TLUserUpdateProfilePhoto) (*mtproto.Int64, error)
	UserDeleteProfilePhotos(ctx context.Context, in *user.TLUserDeleteProfilePhotos) (*mtproto.Int64, error)
	UserGetProfilePhotos(ctx context.Context, in *user.TLUserGetProfilePhotos) (*user.Vector_Long, error)
	UserSetBotCommands(ctx context.Context, in *user.TLUserSetBotCommands) (*mtproto.Bool, error)
	UserIsBot(ctx context.Context, in *user.TLUserIsBot) (*mtproto.Bool, error)
	UserGetBotInfo(ctx context.Context, in *user.TLUserGetBotInfo) (*mtproto.BotInfo, error)
	UserCheckBots(ctx context.Context, in *user.TLUserCheckBots) (*user.Vector_Long, error)
	UserGetFullUser(ctx context.Context, in *user.TLUserGetFullUser) (*mtproto.Users_UserFull, error)
	UserUpdateEmojiStatus(ctx context.Context, in *user.TLUserUpdateEmojiStatus) (*mtproto.Bool, error)
	UserGetUserDataById(ctx context.Context, in *user.TLUserGetUserDataById) (*mtproto.UserData, error)
	UserGetUserDataListByIdList(ctx context.Context, in *user.TLUserGetUserDataListByIdList) (*user.Vector_UserData, error)
	UserGetUserDataByToken(ctx context.Context, in *user.TLUserGetUserDataByToken) (*mtproto.UserData, error)
	UserSearch(ctx context.Context, in *user.TLUserSearch) (*user.UsersFound, error)
	UserUpdateBotData(ctx context.Context, in *user.TLUserUpdateBotData) (*mtproto.Bool, error)
	UserGetImmutableUserV2(ctx context.Context, in *user.TLUserGetImmutableUserV2) (*mtproto.ImmutableUser, error)
	UserGetMutableUsersV2(ctx context.Context, in *user.TLUserGetMutableUsersV2) (*mtproto.MutableUsers, error)
	UserCreateNewTestUser(ctx context.Context, in *user.TLUserCreateNewTestUser) (*mtproto.ImmutableUser, error)
	UserEditCloseFriends(ctx context.Context, in *user.TLUserEditCloseFriends) (*mtproto.Bool, error)
	UserSetStoriesMaxId(ctx context.Context, in *user.TLUserSetStoriesMaxId) (*mtproto.Bool, error)
	UserSetColor(ctx context.Context, in *user.TLUserSetColor) (*mtproto.Bool, error)
	UserUpdateBirthday(ctx context.Context, in *user.TLUserUpdateBirthday) (*mtproto.Bool, error)
	UserGetBirthdays(ctx context.Context, in *user.TLUserGetBirthdays) (*user.Vector_ContactBirthday, error)
	UserSetStoriesHidden(ctx context.Context, in *user.TLUserSetStoriesHidden) (*mtproto.Bool, error)
}

type defaultUserClient struct {
	cli zrpc.Client
}

func NewUserClient(cli zrpc.Client) UserClient {
	return &defaultUserClient{
		cli: cli,
	}
}

// UserGetLastSeens
// user.getLastSeens id:Vector<long> = Vector<LastSeenData>;
func (m *defaultUserClient) UserGetLastSeens(ctx context.Context, in *user.TLUserGetLastSeens) (*user.Vector_LastSeenData, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetLastSeens(ctx, in)
}

// UserUpdateLastSeen
// user.updateLastSeen id:long last_seen_at:long expires:int = Bool;
func (m *defaultUserClient) UserUpdateLastSeen(ctx context.Context, in *user.TLUserUpdateLastSeen) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserUpdateLastSeen(ctx, in)
}

// UserGetLastSeen
// user.getLastSeen id:long = LastSeenData;
func (m *defaultUserClient) UserGetLastSeen(ctx context.Context, in *user.TLUserGetLastSeen) (*user.LastSeenData, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetLastSeen(ctx, in)
}

// UserGetImmutableUser
// user.getImmutableUser flags:# id:long privacy:flags.1?true contacts:Vector<long> = ImmutableUser;
func (m *defaultUserClient) UserGetImmutableUser(ctx context.Context, in *user.TLUserGetImmutableUser) (*mtproto.ImmutableUser, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetImmutableUser(ctx, in)
}

// UserGetMutableUsers
// user.getMutableUsers id:Vector<long> to:Vector<long> = Vector<ImmutableUser>;
func (m *defaultUserClient) UserGetMutableUsers(ctx context.Context, in *user.TLUserGetMutableUsers) (*user.Vector_ImmutableUser, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetMutableUsers(ctx, in)
}

// UserGetImmutableUserByPhone
// user.getImmutableUserByPhone phone:string = ImmutableUser;
func (m *defaultUserClient) UserGetImmutableUserByPhone(ctx context.Context, in *user.TLUserGetImmutableUserByPhone) (*mtproto.ImmutableUser, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetImmutableUserByPhone(ctx, in)
}

// UserGetImmutableUserByToken
// user.getImmutableUserByToken token:string = ImmutableUser;
func (m *defaultUserClient) UserGetImmutableUserByToken(ctx context.Context, in *user.TLUserGetImmutableUserByToken) (*mtproto.ImmutableUser, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetImmutableUserByToken(ctx, in)
}

// UserSetAccountDaysTTL
// user.setAccountDaysTTL user_id:long ttl:int = Bool;
func (m *defaultUserClient) UserSetAccountDaysTTL(ctx context.Context, in *user.TLUserSetAccountDaysTTL) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserSetAccountDaysTTL(ctx, in)
}

// UserGetAccountDaysTTL
// user.getAccountDaysTTL user_id:long = AccountDaysTTL;
func (m *defaultUserClient) UserGetAccountDaysTTL(ctx context.Context, in *user.TLUserGetAccountDaysTTL) (*mtproto.AccountDaysTTL, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetAccountDaysTTL(ctx, in)
}

// UserGetNotifySettings
// user.getNotifySettings user_id:long peer_type:int peer_id:long = PeerNotifySettings;
func (m *defaultUserClient) UserGetNotifySettings(ctx context.Context, in *user.TLUserGetNotifySettings) (*mtproto.PeerNotifySettings, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetNotifySettings(ctx, in)
}

// UserGetNotifySettingsList
// user.getNotifySettingsList user_id:long peers:Vector<PeerUtil> = Vector<PeerPeerNotifySettings>;
func (m *defaultUserClient) UserGetNotifySettingsList(ctx context.Context, in *user.TLUserGetNotifySettingsList) (*user.Vector_PeerPeerNotifySettings, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetNotifySettingsList(ctx, in)
}

// UserSetNotifySettings
// user.setNotifySettings user_id:long peer_type:int peer_id:long settings:PeerNotifySettings = Bool;
func (m *defaultUserClient) UserSetNotifySettings(ctx context.Context, in *user.TLUserSetNotifySettings) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserSetNotifySettings(ctx, in)
}

// UserResetNotifySettings
// user.resetNotifySettings user_id:long = Bool;
func (m *defaultUserClient) UserResetNotifySettings(ctx context.Context, in *user.TLUserResetNotifySettings) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserResetNotifySettings(ctx, in)
}

// UserGetAllNotifySettings
// user.getAllNotifySettings user_id:long = Vector<PeerPeerNotifySettings>;
func (m *defaultUserClient) UserGetAllNotifySettings(ctx context.Context, in *user.TLUserGetAllNotifySettings) (*user.Vector_PeerPeerNotifySettings, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetAllNotifySettings(ctx, in)
}

// UserGetGlobalPrivacySettings
// user.getGlobalPrivacySettings user_id:long = GlobalPrivacySettings;
func (m *defaultUserClient) UserGetGlobalPrivacySettings(ctx context.Context, in *user.TLUserGetGlobalPrivacySettings) (*mtproto.GlobalPrivacySettings, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetGlobalPrivacySettings(ctx, in)
}

// UserSetGlobalPrivacySettings
// user.setGlobalPrivacySettings user_id:long settings:GlobalPrivacySettings = Bool;
func (m *defaultUserClient) UserSetGlobalPrivacySettings(ctx context.Context, in *user.TLUserSetGlobalPrivacySettings) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserSetGlobalPrivacySettings(ctx, in)
}

// UserGetPrivacy
// user.getPrivacy user_id:long key_type:int = Vector<PrivacyRule>;
func (m *defaultUserClient) UserGetPrivacy(ctx context.Context, in *user.TLUserGetPrivacy) (*user.Vector_PrivacyRule, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetPrivacy(ctx, in)
}

// UserSetPrivacy
// user.setPrivacy user_id:long key_type:int rules:Vector<PrivacyRule> = Bool;
func (m *defaultUserClient) UserSetPrivacy(ctx context.Context, in *user.TLUserSetPrivacy) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserSetPrivacy(ctx, in)
}

// UserCheckPrivacy
// user.checkPrivacy flags:# user_id:long key_type:int peer_id:long = Bool;
func (m *defaultUserClient) UserCheckPrivacy(ctx context.Context, in *user.TLUserCheckPrivacy) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserCheckPrivacy(ctx, in)
}

// UserAddPeerSettings
// user.addPeerSettings user_id:long peer_type:int peer_id:long settings:PeerSettings = Bool;
func (m *defaultUserClient) UserAddPeerSettings(ctx context.Context, in *user.TLUserAddPeerSettings) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserAddPeerSettings(ctx, in)
}

// UserGetPeerSettings
// user.getPeerSettings user_id:long peer_type:int peer_id:long = PeerSettings;
func (m *defaultUserClient) UserGetPeerSettings(ctx context.Context, in *user.TLUserGetPeerSettings) (*mtproto.PeerSettings, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetPeerSettings(ctx, in)
}

// UserDeletePeerSettings
// user.deletePeerSettings user_id:long peer_type:int peer_id:long = Bool;
func (m *defaultUserClient) UserDeletePeerSettings(ctx context.Context, in *user.TLUserDeletePeerSettings) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserDeletePeerSettings(ctx, in)
}

// UserChangePhone
// user.changePhone user_id:long phone:string = Bool;
func (m *defaultUserClient) UserChangePhone(ctx context.Context, in *user.TLUserChangePhone) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserChangePhone(ctx, in)
}

// UserCreateNewUser
// user.createNewUser secret_key_id:long phone:string country_code:string first_name:string last_name:string = ImmutableUser;
func (m *defaultUserClient) UserCreateNewUser(ctx context.Context, in *user.TLUserCreateNewUser) (*mtproto.ImmutableUser, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserCreateNewUser(ctx, in)
}

// UserDeleteUser
// user.deleteUser user_id:long reason:string = Bool;
func (m *defaultUserClient) UserDeleteUser(ctx context.Context, in *user.TLUserDeleteUser) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserDeleteUser(ctx, in)
}

// UserBlockPeer
// user.blockPeer user_id:long peer_type:int peer_id:long = Bool;
func (m *defaultUserClient) UserBlockPeer(ctx context.Context, in *user.TLUserBlockPeer) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserBlockPeer(ctx, in)
}

// UserUnBlockPeer
// user.unBlockPeer user_id:long peer_type:int peer_id:long = Bool;
func (m *defaultUserClient) UserUnBlockPeer(ctx context.Context, in *user.TLUserUnBlockPeer) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserUnBlockPeer(ctx, in)
}

// UserBlockedByUser
// user.blockedByUser user_id:long peer_user_id:long = Bool;
func (m *defaultUserClient) UserBlockedByUser(ctx context.Context, in *user.TLUserBlockedByUser) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserBlockedByUser(ctx, in)
}

// UserIsBlockedByUser
// user.isBlockedByUser user_id:long peer_user_id:long = Bool;
func (m *defaultUserClient) UserIsBlockedByUser(ctx context.Context, in *user.TLUserIsBlockedByUser) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserIsBlockedByUser(ctx, in)
}

// UserCheckBlockUserList
// user.checkBlockUserList user_id:long id:Vector<long> = Vector<long>;
func (m *defaultUserClient) UserCheckBlockUserList(ctx context.Context, in *user.TLUserCheckBlockUserList) (*user.Vector_Long, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserCheckBlockUserList(ctx, in)
}

// UserGetBlockedList
// user.getBlockedList user_id:long offset:int limit:int = Vector<PeerBlocked>;
func (m *defaultUserClient) UserGetBlockedList(ctx context.Context, in *user.TLUserGetBlockedList) (*user.Vector_PeerBlocked, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetBlockedList(ctx, in)
}

// UserGetContactSignUpNotification
// user.getContactSignUpNotification user_id:long = Bool;
func (m *defaultUserClient) UserGetContactSignUpNotification(ctx context.Context, in *user.TLUserGetContactSignUpNotification) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetContactSignUpNotification(ctx, in)
}

// UserSetContactSignUpNotification
// user.setContactSignUpNotification user_id:long silent:Bool = Bool;
func (m *defaultUserClient) UserSetContactSignUpNotification(ctx context.Context, in *user.TLUserSetContactSignUpNotification) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserSetContactSignUpNotification(ctx, in)
}

// UserGetContentSettings
// user.getContentSettings user_id:long = account.ContentSettings;
func (m *defaultUserClient) UserGetContentSettings(ctx context.Context, in *user.TLUserGetContentSettings) (*mtproto.Account_ContentSettings, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetContentSettings(ctx, in)
}

// UserSetContentSettings
// user.setContentSettings flags:# user_id:long sensitive_enabled:flags.0?true = Bool;
func (m *defaultUserClient) UserSetContentSettings(ctx context.Context, in *user.TLUserSetContentSettings) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserSetContentSettings(ctx, in)
}

// UserDeleteContact
// user.deleteContact user_id:long id:long = Bool;
func (m *defaultUserClient) UserDeleteContact(ctx context.Context, in *user.TLUserDeleteContact) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserDeleteContact(ctx, in)
}

// UserGetContactList
// user.getContactList user_id:long = Vector<ContactData>;
func (m *defaultUserClient) UserGetContactList(ctx context.Context, in *user.TLUserGetContactList) (*user.Vector_ContactData, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetContactList(ctx, in)
}

// UserGetContactIdList
// user.getContactIdList user_id:long = Vector<long>;
func (m *defaultUserClient) UserGetContactIdList(ctx context.Context, in *user.TLUserGetContactIdList) (*user.Vector_Long, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetContactIdList(ctx, in)
}

// UserGetContact
// user.getContact user_id:long id:long = ContactData;
func (m *defaultUserClient) UserGetContact(ctx context.Context, in *user.TLUserGetContact) (*mtproto.ContactData, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetContact(ctx, in)
}

// UserAddContact
// user.addContact user_id:long add_phone_privacy_exception:Bool id:long first_name:string last_name:string phone:string = Bool;
func (m *defaultUserClient) UserAddContact(ctx context.Context, in *user.TLUserAddContact) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserAddContact(ctx, in)
}

// UserCheckContact
// user.checkContact user_id:long id:long = Bool;
func (m *defaultUserClient) UserCheckContact(ctx context.Context, in *user.TLUserCheckContact) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserCheckContact(ctx, in)
}

// UserGetImportersByPhone
// user.getImportersByPhone phone:string = Vector<InputContact>;
func (m *defaultUserClient) UserGetImportersByPhone(ctx context.Context, in *user.TLUserGetImportersByPhone) (*user.Vector_InputContact, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetImportersByPhone(ctx, in)
}

// UserDeleteImportersByPhone
// user.deleteImportersByPhone phone:string = Bool;
func (m *defaultUserClient) UserDeleteImportersByPhone(ctx context.Context, in *user.TLUserDeleteImportersByPhone) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserDeleteImportersByPhone(ctx, in)
}

// UserImportContacts
// user.importContacts user_id:long contacts:Vector<InputContact> = UserImportedContacts;
func (m *defaultUserClient) UserImportContacts(ctx context.Context, in *user.TLUserImportContacts) (*user.UserImportedContacts, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserImportContacts(ctx, in)
}

// UserGetCountryCode
// user.getCountryCode user_id:long = String;
func (m *defaultUserClient) UserGetCountryCode(ctx context.Context, in *user.TLUserGetCountryCode) (*mtproto.String, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetCountryCode(ctx, in)
}

// UserUpdateAbout
// user.updateAbout user_id:long about:string = Bool;
func (m *defaultUserClient) UserUpdateAbout(ctx context.Context, in *user.TLUserUpdateAbout) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserUpdateAbout(ctx, in)
}

// UserUpdateFirstAndLastName
// user.updateFirstAndLastName user_id:long first_name:string last_name:string = Bool;
func (m *defaultUserClient) UserUpdateFirstAndLastName(ctx context.Context, in *user.TLUserUpdateFirstAndLastName) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserUpdateFirstAndLastName(ctx, in)
}

// UserUpdateVerified
// user.updateVerified user_id:long verified:Bool = Bool;
func (m *defaultUserClient) UserUpdateVerified(ctx context.Context, in *user.TLUserUpdateVerified) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserUpdateVerified(ctx, in)
}

// UserUpdateUsername
// user.updateUsername user_id:long username:string = Bool;
func (m *defaultUserClient) UserUpdateUsername(ctx context.Context, in *user.TLUserUpdateUsername) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserUpdateUsername(ctx, in)
}

// UserUpdateProfilePhoto
// user.updateProfilePhoto user_id:long id:long = Int64;
func (m *defaultUserClient) UserUpdateProfilePhoto(ctx context.Context, in *user.TLUserUpdateProfilePhoto) (*mtproto.Int64, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserUpdateProfilePhoto(ctx, in)
}

// UserDeleteProfilePhotos
// user.deleteProfilePhotos user_id:long id:Vector<long> = Int64;
func (m *defaultUserClient) UserDeleteProfilePhotos(ctx context.Context, in *user.TLUserDeleteProfilePhotos) (*mtproto.Int64, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserDeleteProfilePhotos(ctx, in)
}

// UserGetProfilePhotos
// user.getProfilePhotos user_id:long = Vector<long>;
func (m *defaultUserClient) UserGetProfilePhotos(ctx context.Context, in *user.TLUserGetProfilePhotos) (*user.Vector_Long, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetProfilePhotos(ctx, in)
}

// UserSetBotCommands
// user.setBotCommands user_id:long bot_id:long commands:Vector<BotCommand> = Bool;
func (m *defaultUserClient) UserSetBotCommands(ctx context.Context, in *user.TLUserSetBotCommands) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserSetBotCommands(ctx, in)
}

// UserIsBot
// user.isBot id:long = Bool;
func (m *defaultUserClient) UserIsBot(ctx context.Context, in *user.TLUserIsBot) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserIsBot(ctx, in)
}

// UserGetBotInfo
// user.getBotInfo bot_id:long = BotInfo;
func (m *defaultUserClient) UserGetBotInfo(ctx context.Context, in *user.TLUserGetBotInfo) (*mtproto.BotInfo, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetBotInfo(ctx, in)
}

// UserCheckBots
// user.checkBots id:Vector<long> = Vector<long>;
func (m *defaultUserClient) UserCheckBots(ctx context.Context, in *user.TLUserCheckBots) (*user.Vector_Long, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserCheckBots(ctx, in)
}

// UserGetFullUser
// user.getFullUser self_user_id:long id:long = users.UserFull;
func (m *defaultUserClient) UserGetFullUser(ctx context.Context, in *user.TLUserGetFullUser) (*mtproto.Users_UserFull, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetFullUser(ctx, in)
}

// UserUpdateEmojiStatus
// user.updateEmojiStatus user_id:long emoji_status_document_id:long emoji_status_until:int = Bool;
func (m *defaultUserClient) UserUpdateEmojiStatus(ctx context.Context, in *user.TLUserUpdateEmojiStatus) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserUpdateEmojiStatus(ctx, in)
}

// UserGetUserDataById
// user.getUserDataById user_id:long = UserData;
func (m *defaultUserClient) UserGetUserDataById(ctx context.Context, in *user.TLUserGetUserDataById) (*mtproto.UserData, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetUserDataById(ctx, in)
}

// UserGetUserDataListByIdList
// user.getUserDataListByIdList user_id_list:Vector<long> = Vector<UserData>;
func (m *defaultUserClient) UserGetUserDataListByIdList(ctx context.Context, in *user.TLUserGetUserDataListByIdList) (*user.Vector_UserData, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetUserDataListByIdList(ctx, in)
}

// UserGetUserDataByToken
// user.getUserDataByToken token:string = UserData;
func (m *defaultUserClient) UserGetUserDataByToken(ctx context.Context, in *user.TLUserGetUserDataByToken) (*mtproto.UserData, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetUserDataByToken(ctx, in)
}

// UserSearch
// user.search q:string excluded_contacts:Vector<long> offset:long limit:int = UsersFound;
func (m *defaultUserClient) UserSearch(ctx context.Context, in *user.TLUserSearch) (*user.UsersFound, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserSearch(ctx, in)
}

// UserUpdateBotData
// user.updateBotData flags:# bot_id:long bot_chat_history:flags.15?Bool bot_nochats:flags.16?Bool bot_inline_geo:flags.21?Bool bot_attach_menu:flags.27?Bool bot_inline_placeholder:flags.19?Bool = Bool;
func (m *defaultUserClient) UserUpdateBotData(ctx context.Context, in *user.TLUserUpdateBotData) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserUpdateBotData(ctx, in)
}

// UserGetImmutableUserV2
// user.getImmutableUserV2 flags:# id:long privacy:flags.0?true has_to:flags.2?true to:flags.2?Vector<long> = ImmutableUser;
func (m *defaultUserClient) UserGetImmutableUserV2(ctx context.Context, in *user.TLUserGetImmutableUserV2) (*mtproto.ImmutableUser, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetImmutableUserV2(ctx, in)
}

// UserGetMutableUsersV2
// user.getMutableUsersV2 flags:# id:Vector<long> privacy:flags.0?true has_to:flags.2?true to:flags.2?Vector<long> = MutableUsers;
func (m *defaultUserClient) UserGetMutableUsersV2(ctx context.Context, in *user.TLUserGetMutableUsersV2) (*mtproto.MutableUsers, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetMutableUsersV2(ctx, in)
}

// UserCreateNewTestUser
// user.createNewTestUser secret_key_id:long min_id:long max_id:long = ImmutableUser;
func (m *defaultUserClient) UserCreateNewTestUser(ctx context.Context, in *user.TLUserCreateNewTestUser) (*mtproto.ImmutableUser, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserCreateNewTestUser(ctx, in)
}

// UserEditCloseFriends
// user.editCloseFriends user_id:long id:Vector<long> = Bool;
func (m *defaultUserClient) UserEditCloseFriends(ctx context.Context, in *user.TLUserEditCloseFriends) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserEditCloseFriends(ctx, in)
}

// UserSetStoriesMaxId
// user.setStoriesMaxId user_id:long id:int = Bool;
func (m *defaultUserClient) UserSetStoriesMaxId(ctx context.Context, in *user.TLUserSetStoriesMaxId) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserSetStoriesMaxId(ctx, in)
}

// UserSetColor
// user.setColor flags:# user_id:long for_profile:flags.1?true color:int background_emoji_id:long = Bool;
func (m *defaultUserClient) UserSetColor(ctx context.Context, in *user.TLUserSetColor) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserSetColor(ctx, in)
}

// UserUpdateBirthday
// user.updateBirthday flags:# user_id:long birthday:flags.1?Birthday = Bool;
func (m *defaultUserClient) UserUpdateBirthday(ctx context.Context, in *user.TLUserUpdateBirthday) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserUpdateBirthday(ctx, in)
}

// UserGetBirthdays
// user.getBirthdays user_id:long = Vector<ContactBirthday>;
func (m *defaultUserClient) UserGetBirthdays(ctx context.Context, in *user.TLUserGetBirthdays) (*user.Vector_ContactBirthday, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetBirthdays(ctx, in)
}

// UserSetStoriesHidden
// user.setStoriesHidden user_id:long  id:long hidden:Bool = Bool;
func (m *defaultUserClient) UserSetStoriesHidden(ctx context.Context, in *user.TLUserSetStoriesHidden) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserSetStoriesHidden(ctx, in)
}
