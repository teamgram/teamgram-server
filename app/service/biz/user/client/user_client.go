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

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/user/userservice"

	"github.com/cloudwego/kitex/client"
)

var _ *tg.Bool

type UserClient interface {
	UserGetLastSeens(ctx context.Context, in *user.TLUserGetLastSeens) (*user.VectorLastSeenData, error)
	UserUpdateLastSeen(ctx context.Context, in *user.TLUserUpdateLastSeen) (*tg.Bool, error)
	UserGetLastSeen(ctx context.Context, in *user.TLUserGetLastSeen) (*user.LastSeenData, error)
	UserGetImmutableUser(ctx context.Context, in *user.TLUserGetImmutableUser) (*tg.ImmutableUser, error)
	UserGetMutableUsers(ctx context.Context, in *user.TLUserGetMutableUsers) (*user.VectorImmutableUser, error)
	UserGetImmutableUserByPhone(ctx context.Context, in *user.TLUserGetImmutableUserByPhone) (*tg.ImmutableUser, error)
	UserGetImmutableUserByToken(ctx context.Context, in *user.TLUserGetImmutableUserByToken) (*tg.ImmutableUser, error)
	UserSetAccountDaysTTL(ctx context.Context, in *user.TLUserSetAccountDaysTTL) (*tg.Bool, error)
	UserGetAccountDaysTTL(ctx context.Context, in *user.TLUserGetAccountDaysTTL) (*tg.AccountDaysTTL, error)
	UserGetNotifySettings(ctx context.Context, in *user.TLUserGetNotifySettings) (*tg.PeerNotifySettings, error)
	UserGetNotifySettingsList(ctx context.Context, in *user.TLUserGetNotifySettingsList) (*user.VectorPeerPeerNotifySettings, error)
	UserSetNotifySettings(ctx context.Context, in *user.TLUserSetNotifySettings) (*tg.Bool, error)
	UserResetNotifySettings(ctx context.Context, in *user.TLUserResetNotifySettings) (*tg.Bool, error)
	UserGetAllNotifySettings(ctx context.Context, in *user.TLUserGetAllNotifySettings) (*user.VectorPeerPeerNotifySettings, error)
	UserGetGlobalPrivacySettings(ctx context.Context, in *user.TLUserGetGlobalPrivacySettings) (*tg.GlobalPrivacySettings, error)
	UserSetGlobalPrivacySettings(ctx context.Context, in *user.TLUserSetGlobalPrivacySettings) (*tg.Bool, error)
	UserGetPrivacy(ctx context.Context, in *user.TLUserGetPrivacy) (*user.VectorPrivacyRule, error)
	UserSetPrivacy(ctx context.Context, in *user.TLUserSetPrivacy) (*tg.Bool, error)
	UserCheckPrivacy(ctx context.Context, in *user.TLUserCheckPrivacy) (*tg.Bool, error)
	UserAddPeerSettings(ctx context.Context, in *user.TLUserAddPeerSettings) (*tg.Bool, error)
	UserGetPeerSettings(ctx context.Context, in *user.TLUserGetPeerSettings) (*tg.PeerSettings, error)
	UserDeletePeerSettings(ctx context.Context, in *user.TLUserDeletePeerSettings) (*tg.Bool, error)
	UserChangePhone(ctx context.Context, in *user.TLUserChangePhone) (*tg.Bool, error)
	UserCreateNewUser(ctx context.Context, in *user.TLUserCreateNewUser) (*tg.ImmutableUser, error)
	UserDeleteUser(ctx context.Context, in *user.TLUserDeleteUser) (*tg.Bool, error)
	UserBlockPeer(ctx context.Context, in *user.TLUserBlockPeer) (*tg.Bool, error)
	UserUnBlockPeer(ctx context.Context, in *user.TLUserUnBlockPeer) (*tg.Bool, error)
	UserBlockedByUser(ctx context.Context, in *user.TLUserBlockedByUser) (*tg.Bool, error)
	UserIsBlockedByUser(ctx context.Context, in *user.TLUserIsBlockedByUser) (*tg.Bool, error)
	UserCheckBlockUserList(ctx context.Context, in *user.TLUserCheckBlockUserList) (*user.VectorLong, error)
	UserGetBlockedList(ctx context.Context, in *user.TLUserGetBlockedList) (*user.VectorPeerBlocked, error)
	UserGetContactSignUpNotification(ctx context.Context, in *user.TLUserGetContactSignUpNotification) (*tg.Bool, error)
	UserSetContactSignUpNotification(ctx context.Context, in *user.TLUserSetContactSignUpNotification) (*tg.Bool, error)
	UserGetContentSettings(ctx context.Context, in *user.TLUserGetContentSettings) (*tg.AccountContentSettings, error)
	UserSetContentSettings(ctx context.Context, in *user.TLUserSetContentSettings) (*tg.Bool, error)
	UserDeleteContact(ctx context.Context, in *user.TLUserDeleteContact) (*tg.Bool, error)
	UserGetContactList(ctx context.Context, in *user.TLUserGetContactList) (*user.VectorContactData, error)
	UserGetContactIdList(ctx context.Context, in *user.TLUserGetContactIdList) (*user.VectorLong, error)
	UserGetContact(ctx context.Context, in *user.TLUserGetContact) (*tg.ContactData, error)
	UserAddContact(ctx context.Context, in *user.TLUserAddContact) (*tg.Bool, error)
	UserCheckContact(ctx context.Context, in *user.TLUserCheckContact) (*tg.Bool, error)
	UserGetImportersByPhone(ctx context.Context, in *user.TLUserGetImportersByPhone) (*user.VectorInputContact, error)
	UserDeleteImportersByPhone(ctx context.Context, in *user.TLUserDeleteImportersByPhone) (*tg.Bool, error)
	UserImportContacts(ctx context.Context, in *user.TLUserImportContacts) (*user.UserImportedContacts, error)
	UserGetCountryCode(ctx context.Context, in *user.TLUserGetCountryCode) (*tg.String, error)
	UserUpdateAbout(ctx context.Context, in *user.TLUserUpdateAbout) (*tg.Bool, error)
	UserUpdateFirstAndLastName(ctx context.Context, in *user.TLUserUpdateFirstAndLastName) (*tg.Bool, error)
	UserUpdateVerified(ctx context.Context, in *user.TLUserUpdateVerified) (*tg.Bool, error)
	UserUpdateUsername(ctx context.Context, in *user.TLUserUpdateUsername) (*tg.Bool, error)
	UserUpdateProfilePhoto(ctx context.Context, in *user.TLUserUpdateProfilePhoto) (*tg.Int64, error)
	UserDeleteProfilePhotos(ctx context.Context, in *user.TLUserDeleteProfilePhotos) (*tg.Int64, error)
	UserGetProfilePhotos(ctx context.Context, in *user.TLUserGetProfilePhotos) (*user.VectorLong, error)
	UserSetBotCommands(ctx context.Context, in *user.TLUserSetBotCommands) (*tg.Bool, error)
	UserIsBot(ctx context.Context, in *user.TLUserIsBot) (*tg.Bool, error)
	UserGetBotInfo(ctx context.Context, in *user.TLUserGetBotInfo) (*tg.BotInfo, error)
	UserCheckBots(ctx context.Context, in *user.TLUserCheckBots) (*user.VectorLong, error)
	UserGetFullUser(ctx context.Context, in *user.TLUserGetFullUser) (*tg.UsersUserFull, error)
	UserUpdateEmojiStatus(ctx context.Context, in *user.TLUserUpdateEmojiStatus) (*tg.Bool, error)
	UserGetUserDataById(ctx context.Context, in *user.TLUserGetUserDataById) (*tg.UserData, error)
	UserGetUserDataListByIdList(ctx context.Context, in *user.TLUserGetUserDataListByIdList) (*user.VectorUserData, error)
	UserGetUserDataByToken(ctx context.Context, in *user.TLUserGetUserDataByToken) (*tg.UserData, error)
	UserSearch(ctx context.Context, in *user.TLUserSearch) (*user.UsersFound, error)
	UserUpdateBotData(ctx context.Context, in *user.TLUserUpdateBotData) (*tg.Bool, error)
	UserGetImmutableUserV2(ctx context.Context, in *user.TLUserGetImmutableUserV2) (*tg.ImmutableUser, error)
	UserGetMutableUsersV2(ctx context.Context, in *user.TLUserGetMutableUsersV2) (*tg.MutableUsers, error)
	UserCreateNewTestUser(ctx context.Context, in *user.TLUserCreateNewTestUser) (*tg.ImmutableUser, error)
	UserEditCloseFriends(ctx context.Context, in *user.TLUserEditCloseFriends) (*tg.Bool, error)
	UserSetStoriesMaxId(ctx context.Context, in *user.TLUserSetStoriesMaxId) (*tg.Bool, error)
	UserSetColor(ctx context.Context, in *user.TLUserSetColor) (*tg.Bool, error)
	UserUpdateBirthday(ctx context.Context, in *user.TLUserUpdateBirthday) (*tg.Bool, error)
	UserGetBirthdays(ctx context.Context, in *user.TLUserGetBirthdays) (*user.VectorContactBirthday, error)
	UserSetStoriesHidden(ctx context.Context, in *user.TLUserSetStoriesHidden) (*tg.Bool, error)
	UserUpdatePersonalChannel(ctx context.Context, in *user.TLUserUpdatePersonalChannel) (*tg.Bool, error)
	UserGetUserIdByPhone(ctx context.Context, in *user.TLUserGetUserIdByPhone) (*tg.Int64, error)
	UserSetAuthorizationTTL(ctx context.Context, in *user.TLUserSetAuthorizationTTL) (*tg.Bool, error)
	UserGetAuthorizationTTL(ctx context.Context, in *user.TLUserGetAuthorizationTTL) (*tg.AccountDaysTTL, error)
	UserUpdatePremium(ctx context.Context, in *user.TLUserUpdatePremium) (*tg.Bool, error)
	UserGetBotInfoV2(ctx context.Context, in *user.TLUserGetBotInfoV2) (*user.BotInfoData, error)
}

type defaultUserClient struct {
	cli client.Client
}

func NewUserClient(cli client.Client) UserClient {
	return &defaultUserClient{
		cli: cli,
	}
}

// UserGetLastSeens
// user.getLastSeens id:Vector<long> = Vector<LastSeenData>;
func (m *defaultUserClient) UserGetLastSeens(ctx context.Context, in *user.TLUserGetLastSeens) (*user.VectorLastSeenData, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetLastSeens(ctx, in)
}

// UserUpdateLastSeen
// user.updateLastSeen id:long last_seen_at:long expires:int = Bool;
func (m *defaultUserClient) UserUpdateLastSeen(ctx context.Context, in *user.TLUserUpdateLastSeen) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserUpdateLastSeen(ctx, in)
}

// UserGetLastSeen
// user.getLastSeen id:long = LastSeenData;
func (m *defaultUserClient) UserGetLastSeen(ctx context.Context, in *user.TLUserGetLastSeen) (*user.LastSeenData, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetLastSeen(ctx, in)
}

// UserGetImmutableUser
// user.getImmutableUser flags:# id:long privacy:flags.1?true contacts:Vector<long> = ImmutableUser;
func (m *defaultUserClient) UserGetImmutableUser(ctx context.Context, in *user.TLUserGetImmutableUser) (*tg.ImmutableUser, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetImmutableUser(ctx, in)
}

// UserGetMutableUsers
// user.getMutableUsers id:Vector<long> to:Vector<long> = Vector<ImmutableUser>;
func (m *defaultUserClient) UserGetMutableUsers(ctx context.Context, in *user.TLUserGetMutableUsers) (*user.VectorImmutableUser, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetMutableUsers(ctx, in)
}

// UserGetImmutableUserByPhone
// user.getImmutableUserByPhone phone:string = ImmutableUser;
func (m *defaultUserClient) UserGetImmutableUserByPhone(ctx context.Context, in *user.TLUserGetImmutableUserByPhone) (*tg.ImmutableUser, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetImmutableUserByPhone(ctx, in)
}

// UserGetImmutableUserByToken
// user.getImmutableUserByToken token:string = ImmutableUser;
func (m *defaultUserClient) UserGetImmutableUserByToken(ctx context.Context, in *user.TLUserGetImmutableUserByToken) (*tg.ImmutableUser, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetImmutableUserByToken(ctx, in)
}

// UserSetAccountDaysTTL
// user.setAccountDaysTTL user_id:long ttl:int = Bool;
func (m *defaultUserClient) UserSetAccountDaysTTL(ctx context.Context, in *user.TLUserSetAccountDaysTTL) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserSetAccountDaysTTL(ctx, in)
}

// UserGetAccountDaysTTL
// user.getAccountDaysTTL user_id:long = AccountDaysTTL;
func (m *defaultUserClient) UserGetAccountDaysTTL(ctx context.Context, in *user.TLUserGetAccountDaysTTL) (*tg.AccountDaysTTL, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetAccountDaysTTL(ctx, in)
}

// UserGetNotifySettings
// user.getNotifySettings user_id:long peer_type:int peer_id:long = PeerNotifySettings;
func (m *defaultUserClient) UserGetNotifySettings(ctx context.Context, in *user.TLUserGetNotifySettings) (*tg.PeerNotifySettings, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetNotifySettings(ctx, in)
}

// UserGetNotifySettingsList
// user.getNotifySettingsList user_id:long peers:Vector<PeerUtil> = Vector<PeerPeerNotifySettings>;
func (m *defaultUserClient) UserGetNotifySettingsList(ctx context.Context, in *user.TLUserGetNotifySettingsList) (*user.VectorPeerPeerNotifySettings, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetNotifySettingsList(ctx, in)
}

// UserSetNotifySettings
// user.setNotifySettings user_id:long peer_type:int peer_id:long settings:PeerNotifySettings = Bool;
func (m *defaultUserClient) UserSetNotifySettings(ctx context.Context, in *user.TLUserSetNotifySettings) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserSetNotifySettings(ctx, in)
}

// UserResetNotifySettings
// user.resetNotifySettings user_id:long = Bool;
func (m *defaultUserClient) UserResetNotifySettings(ctx context.Context, in *user.TLUserResetNotifySettings) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserResetNotifySettings(ctx, in)
}

// UserGetAllNotifySettings
// user.getAllNotifySettings user_id:long = Vector<PeerPeerNotifySettings>;
func (m *defaultUserClient) UserGetAllNotifySettings(ctx context.Context, in *user.TLUserGetAllNotifySettings) (*user.VectorPeerPeerNotifySettings, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetAllNotifySettings(ctx, in)
}

// UserGetGlobalPrivacySettings
// user.getGlobalPrivacySettings user_id:long = GlobalPrivacySettings;
func (m *defaultUserClient) UserGetGlobalPrivacySettings(ctx context.Context, in *user.TLUserGetGlobalPrivacySettings) (*tg.GlobalPrivacySettings, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetGlobalPrivacySettings(ctx, in)
}

// UserSetGlobalPrivacySettings
// user.setGlobalPrivacySettings user_id:long settings:GlobalPrivacySettings = Bool;
func (m *defaultUserClient) UserSetGlobalPrivacySettings(ctx context.Context, in *user.TLUserSetGlobalPrivacySettings) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserSetGlobalPrivacySettings(ctx, in)
}

// UserGetPrivacy
// user.getPrivacy user_id:long key_type:int = Vector<PrivacyRule>;
func (m *defaultUserClient) UserGetPrivacy(ctx context.Context, in *user.TLUserGetPrivacy) (*user.VectorPrivacyRule, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetPrivacy(ctx, in)
}

// UserSetPrivacy
// user.setPrivacy user_id:long key_type:int rules:Vector<PrivacyRule> = Bool;
func (m *defaultUserClient) UserSetPrivacy(ctx context.Context, in *user.TLUserSetPrivacy) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserSetPrivacy(ctx, in)
}

// UserCheckPrivacy
// user.checkPrivacy flags:# user_id:long key_type:int peer_id:long = Bool;
func (m *defaultUserClient) UserCheckPrivacy(ctx context.Context, in *user.TLUserCheckPrivacy) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserCheckPrivacy(ctx, in)
}

// UserAddPeerSettings
// user.addPeerSettings user_id:long peer_type:int peer_id:long settings:PeerSettings = Bool;
func (m *defaultUserClient) UserAddPeerSettings(ctx context.Context, in *user.TLUserAddPeerSettings) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserAddPeerSettings(ctx, in)
}

// UserGetPeerSettings
// user.getPeerSettings user_id:long peer_type:int peer_id:long = PeerSettings;
func (m *defaultUserClient) UserGetPeerSettings(ctx context.Context, in *user.TLUserGetPeerSettings) (*tg.PeerSettings, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetPeerSettings(ctx, in)
}

// UserDeletePeerSettings
// user.deletePeerSettings user_id:long peer_type:int peer_id:long = Bool;
func (m *defaultUserClient) UserDeletePeerSettings(ctx context.Context, in *user.TLUserDeletePeerSettings) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserDeletePeerSettings(ctx, in)
}

// UserChangePhone
// user.changePhone user_id:long phone:string = Bool;
func (m *defaultUserClient) UserChangePhone(ctx context.Context, in *user.TLUserChangePhone) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserChangePhone(ctx, in)
}

// UserCreateNewUser
// user.createNewUser secret_key_id:long phone:string country_code:string first_name:string last_name:string = ImmutableUser;
func (m *defaultUserClient) UserCreateNewUser(ctx context.Context, in *user.TLUserCreateNewUser) (*tg.ImmutableUser, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserCreateNewUser(ctx, in)
}

// UserDeleteUser
// user.deleteUser user_id:long reason:string phone:string = Bool;
func (m *defaultUserClient) UserDeleteUser(ctx context.Context, in *user.TLUserDeleteUser) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserDeleteUser(ctx, in)
}

// UserBlockPeer
// user.blockPeer user_id:long peer_type:int peer_id:long = Bool;
func (m *defaultUserClient) UserBlockPeer(ctx context.Context, in *user.TLUserBlockPeer) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserBlockPeer(ctx, in)
}

// UserUnBlockPeer
// user.unBlockPeer user_id:long peer_type:int peer_id:long = Bool;
func (m *defaultUserClient) UserUnBlockPeer(ctx context.Context, in *user.TLUserUnBlockPeer) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserUnBlockPeer(ctx, in)
}

// UserBlockedByUser
// user.blockedByUser user_id:long peer_user_id:long = Bool;
func (m *defaultUserClient) UserBlockedByUser(ctx context.Context, in *user.TLUserBlockedByUser) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserBlockedByUser(ctx, in)
}

// UserIsBlockedByUser
// user.isBlockedByUser user_id:long peer_user_id:long = Bool;
func (m *defaultUserClient) UserIsBlockedByUser(ctx context.Context, in *user.TLUserIsBlockedByUser) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserIsBlockedByUser(ctx, in)
}

// UserCheckBlockUserList
// user.checkBlockUserList user_id:long id:Vector<long> = Vector<long>;
func (m *defaultUserClient) UserCheckBlockUserList(ctx context.Context, in *user.TLUserCheckBlockUserList) (*user.VectorLong, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserCheckBlockUserList(ctx, in)
}

// UserGetBlockedList
// user.getBlockedList user_id:long offset:int limit:int = Vector<PeerBlocked>;
func (m *defaultUserClient) UserGetBlockedList(ctx context.Context, in *user.TLUserGetBlockedList) (*user.VectorPeerBlocked, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetBlockedList(ctx, in)
}

// UserGetContactSignUpNotification
// user.getContactSignUpNotification user_id:long = Bool;
func (m *defaultUserClient) UserGetContactSignUpNotification(ctx context.Context, in *user.TLUserGetContactSignUpNotification) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetContactSignUpNotification(ctx, in)
}

// UserSetContactSignUpNotification
// user.setContactSignUpNotification user_id:long silent:Bool = Bool;
func (m *defaultUserClient) UserSetContactSignUpNotification(ctx context.Context, in *user.TLUserSetContactSignUpNotification) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserSetContactSignUpNotification(ctx, in)
}

// UserGetContentSettings
// user.getContentSettings user_id:long = account.ContentSettings;
func (m *defaultUserClient) UserGetContentSettings(ctx context.Context, in *user.TLUserGetContentSettings) (*tg.AccountContentSettings, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetContentSettings(ctx, in)
}

// UserSetContentSettings
// user.setContentSettings flags:# user_id:long sensitive_enabled:flags.0?true = Bool;
func (m *defaultUserClient) UserSetContentSettings(ctx context.Context, in *user.TLUserSetContentSettings) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserSetContentSettings(ctx, in)
}

// UserDeleteContact
// user.deleteContact user_id:long id:long = Bool;
func (m *defaultUserClient) UserDeleteContact(ctx context.Context, in *user.TLUserDeleteContact) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserDeleteContact(ctx, in)
}

// UserGetContactList
// user.getContactList user_id:long = Vector<ContactData>;
func (m *defaultUserClient) UserGetContactList(ctx context.Context, in *user.TLUserGetContactList) (*user.VectorContactData, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetContactList(ctx, in)
}

// UserGetContactIdList
// user.getContactIdList user_id:long = Vector<long>;
func (m *defaultUserClient) UserGetContactIdList(ctx context.Context, in *user.TLUserGetContactIdList) (*user.VectorLong, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetContactIdList(ctx, in)
}

// UserGetContact
// user.getContact user_id:long id:long = ContactData;
func (m *defaultUserClient) UserGetContact(ctx context.Context, in *user.TLUserGetContact) (*tg.ContactData, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetContact(ctx, in)
}

// UserAddContact
// user.addContact user_id:long add_phone_privacy_exception:Bool id:long first_name:string last_name:string phone:string = Bool;
func (m *defaultUserClient) UserAddContact(ctx context.Context, in *user.TLUserAddContact) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserAddContact(ctx, in)
}

// UserCheckContact
// user.checkContact user_id:long id:long = Bool;
func (m *defaultUserClient) UserCheckContact(ctx context.Context, in *user.TLUserCheckContact) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserCheckContact(ctx, in)
}

// UserGetImportersByPhone
// user.getImportersByPhone phone:string = Vector<InputContact>;
func (m *defaultUserClient) UserGetImportersByPhone(ctx context.Context, in *user.TLUserGetImportersByPhone) (*user.VectorInputContact, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetImportersByPhone(ctx, in)
}

// UserDeleteImportersByPhone
// user.deleteImportersByPhone phone:string = Bool;
func (m *defaultUserClient) UserDeleteImportersByPhone(ctx context.Context, in *user.TLUserDeleteImportersByPhone) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserDeleteImportersByPhone(ctx, in)
}

// UserImportContacts
// user.importContacts user_id:long contacts:Vector<InputContact> = UserImportedContacts;
func (m *defaultUserClient) UserImportContacts(ctx context.Context, in *user.TLUserImportContacts) (*user.UserImportedContacts, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserImportContacts(ctx, in)
}

// UserGetCountryCode
// user.getCountryCode user_id:long = String;
func (m *defaultUserClient) UserGetCountryCode(ctx context.Context, in *user.TLUserGetCountryCode) (*tg.String, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetCountryCode(ctx, in)
}

// UserUpdateAbout
// user.updateAbout user_id:long about:string = Bool;
func (m *defaultUserClient) UserUpdateAbout(ctx context.Context, in *user.TLUserUpdateAbout) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserUpdateAbout(ctx, in)
}

// UserUpdateFirstAndLastName
// user.updateFirstAndLastName user_id:long first_name:string last_name:string = Bool;
func (m *defaultUserClient) UserUpdateFirstAndLastName(ctx context.Context, in *user.TLUserUpdateFirstAndLastName) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserUpdateFirstAndLastName(ctx, in)
}

// UserUpdateVerified
// user.updateVerified user_id:long verified:Bool = Bool;
func (m *defaultUserClient) UserUpdateVerified(ctx context.Context, in *user.TLUserUpdateVerified) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserUpdateVerified(ctx, in)
}

// UserUpdateUsername
// user.updateUsername user_id:long username:string = Bool;
func (m *defaultUserClient) UserUpdateUsername(ctx context.Context, in *user.TLUserUpdateUsername) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserUpdateUsername(ctx, in)
}

// UserUpdateProfilePhoto
// user.updateProfilePhoto user_id:long id:long = Int64;
func (m *defaultUserClient) UserUpdateProfilePhoto(ctx context.Context, in *user.TLUserUpdateProfilePhoto) (*tg.Int64, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserUpdateProfilePhoto(ctx, in)
}

// UserDeleteProfilePhotos
// user.deleteProfilePhotos user_id:long id:Vector<long> = Int64;
func (m *defaultUserClient) UserDeleteProfilePhotos(ctx context.Context, in *user.TLUserDeleteProfilePhotos) (*tg.Int64, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserDeleteProfilePhotos(ctx, in)
}

// UserGetProfilePhotos
// user.getProfilePhotos user_id:long = Vector<long>;
func (m *defaultUserClient) UserGetProfilePhotos(ctx context.Context, in *user.TLUserGetProfilePhotos) (*user.VectorLong, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetProfilePhotos(ctx, in)
}

// UserSetBotCommands
// user.setBotCommands user_id:long bot_id:long commands:Vector<BotCommand> = Bool;
func (m *defaultUserClient) UserSetBotCommands(ctx context.Context, in *user.TLUserSetBotCommands) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserSetBotCommands(ctx, in)
}

// UserIsBot
// user.isBot id:long = Bool;
func (m *defaultUserClient) UserIsBot(ctx context.Context, in *user.TLUserIsBot) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserIsBot(ctx, in)
}

// UserGetBotInfo
// user.getBotInfo bot_id:long = BotInfo;
func (m *defaultUserClient) UserGetBotInfo(ctx context.Context, in *user.TLUserGetBotInfo) (*tg.BotInfo, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetBotInfo(ctx, in)
}

// UserCheckBots
// user.checkBots id:Vector<long> = Vector<long>;
func (m *defaultUserClient) UserCheckBots(ctx context.Context, in *user.TLUserCheckBots) (*user.VectorLong, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserCheckBots(ctx, in)
}

// UserGetFullUser
// user.getFullUser self_user_id:long id:long = users.UserFull;
func (m *defaultUserClient) UserGetFullUser(ctx context.Context, in *user.TLUserGetFullUser) (*tg.UsersUserFull, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetFullUser(ctx, in)
}

// UserUpdateEmojiStatus
// user.updateEmojiStatus user_id:long emoji_status_document_id:long emoji_status_until:int = Bool;
func (m *defaultUserClient) UserUpdateEmojiStatus(ctx context.Context, in *user.TLUserUpdateEmojiStatus) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserUpdateEmojiStatus(ctx, in)
}

// UserGetUserDataById
// user.getUserDataById user_id:long = UserData;
func (m *defaultUserClient) UserGetUserDataById(ctx context.Context, in *user.TLUserGetUserDataById) (*tg.UserData, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetUserDataById(ctx, in)
}

// UserGetUserDataListByIdList
// user.getUserDataListByIdList user_id_list:Vector<long> = Vector<UserData>;
func (m *defaultUserClient) UserGetUserDataListByIdList(ctx context.Context, in *user.TLUserGetUserDataListByIdList) (*user.VectorUserData, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetUserDataListByIdList(ctx, in)
}

// UserGetUserDataByToken
// user.getUserDataByToken token:string = UserData;
func (m *defaultUserClient) UserGetUserDataByToken(ctx context.Context, in *user.TLUserGetUserDataByToken) (*tg.UserData, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetUserDataByToken(ctx, in)
}

// UserSearch
// user.search q:string excluded_contacts:Vector<long> offset:long limit:int = UsersFound;
func (m *defaultUserClient) UserSearch(ctx context.Context, in *user.TLUserSearch) (*user.UsersFound, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserSearch(ctx, in)
}

// UserUpdateBotData
// user.updateBotData flags:# bot_id:long bot_chat_history:flags.15?Bool bot_nochats:flags.16?Bool bot_inline_geo:flags.21?Bool bot_attach_menu:flags.27?Bool bot_inline_placeholder:flags.19?string bot_has_main_app:flags.13?Bool = Bool;
func (m *defaultUserClient) UserUpdateBotData(ctx context.Context, in *user.TLUserUpdateBotData) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserUpdateBotData(ctx, in)
}

// UserGetImmutableUserV2
// user.getImmutableUserV2 flags:# id:long privacy:flags.0?true has_to:flags.2?true to:flags.2?Vector<long> = ImmutableUser;
func (m *defaultUserClient) UserGetImmutableUserV2(ctx context.Context, in *user.TLUserGetImmutableUserV2) (*tg.ImmutableUser, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetImmutableUserV2(ctx, in)
}

// UserGetMutableUsersV2
// user.getMutableUsersV2 flags:# id:Vector<long> privacy:flags.0?true has_to:flags.2?true to:flags.2?Vector<long> = MutableUsers;
func (m *defaultUserClient) UserGetMutableUsersV2(ctx context.Context, in *user.TLUserGetMutableUsersV2) (*tg.MutableUsers, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetMutableUsersV2(ctx, in)
}

// UserCreateNewTestUser
// user.createNewTestUser secret_key_id:long min_id:long max_id:long = ImmutableUser;
func (m *defaultUserClient) UserCreateNewTestUser(ctx context.Context, in *user.TLUserCreateNewTestUser) (*tg.ImmutableUser, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserCreateNewTestUser(ctx, in)
}

// UserEditCloseFriends
// user.editCloseFriends user_id:long id:Vector<long> = Bool;
func (m *defaultUserClient) UserEditCloseFriends(ctx context.Context, in *user.TLUserEditCloseFriends) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserEditCloseFriends(ctx, in)
}

// UserSetStoriesMaxId
// user.setStoriesMaxId user_id:long id:int = Bool;
func (m *defaultUserClient) UserSetStoriesMaxId(ctx context.Context, in *user.TLUserSetStoriesMaxId) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserSetStoriesMaxId(ctx, in)
}

// UserSetColor
// user.setColor flags:# user_id:long for_profile:flags.1?true color:int background_emoji_id:long = Bool;
func (m *defaultUserClient) UserSetColor(ctx context.Context, in *user.TLUserSetColor) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserSetColor(ctx, in)
}

// UserUpdateBirthday
// user.updateBirthday flags:# user_id:long birthday:flags.1?Birthday = Bool;
func (m *defaultUserClient) UserUpdateBirthday(ctx context.Context, in *user.TLUserUpdateBirthday) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserUpdateBirthday(ctx, in)
}

// UserGetBirthdays
// user.getBirthdays user_id:long = Vector<ContactBirthday>;
func (m *defaultUserClient) UserGetBirthdays(ctx context.Context, in *user.TLUserGetBirthdays) (*user.VectorContactBirthday, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetBirthdays(ctx, in)
}

// UserSetStoriesHidden
// user.setStoriesHidden user_id:long id:long hidden:Bool = Bool;
func (m *defaultUserClient) UserSetStoriesHidden(ctx context.Context, in *user.TLUserSetStoriesHidden) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserSetStoriesHidden(ctx, in)
}

// UserUpdatePersonalChannel
// user.updatePersonalChannel user_id:long channel_id:long = Bool;
func (m *defaultUserClient) UserUpdatePersonalChannel(ctx context.Context, in *user.TLUserUpdatePersonalChannel) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserUpdatePersonalChannel(ctx, in)
}

// UserGetUserIdByPhone
// user.getUserIdByPhone phone:string = Int64;
func (m *defaultUserClient) UserGetUserIdByPhone(ctx context.Context, in *user.TLUserGetUserIdByPhone) (*tg.Int64, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetUserIdByPhone(ctx, in)
}

// UserSetAuthorizationTTL
// user.setAuthorizationTTL user_id:long ttl:int = Bool;
func (m *defaultUserClient) UserSetAuthorizationTTL(ctx context.Context, in *user.TLUserSetAuthorizationTTL) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserSetAuthorizationTTL(ctx, in)
}

// UserGetAuthorizationTTL
// user.getAuthorizationTTL user_id:long = AccountDaysTTL;
func (m *defaultUserClient) UserGetAuthorizationTTL(ctx context.Context, in *user.TLUserGetAuthorizationTTL) (*tg.AccountDaysTTL, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetAuthorizationTTL(ctx, in)
}

// UserUpdatePremium
// user.updatePremium flags:# user_id:long premium:Bool months:flags.1?int = Bool;
func (m *defaultUserClient) UserUpdatePremium(ctx context.Context, in *user.TLUserUpdatePremium) (*tg.Bool, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserUpdatePremium(ctx, in)
}

// UserGetBotInfoV2
// user.getBotInfoV2 bot_id:long = BotInfoData;
func (m *defaultUserClient) UserGetBotInfoV2(ctx context.Context, in *user.TLUserGetBotInfoV2) (*user.BotInfoData, error) {
	cli := userservice.NewRPCUserClient(m.cli)
	return cli.UserGetBotInfoV2(ctx, in)
}
