/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package userclient

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/user/userservice"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

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
	UserSaveMusic(ctx context.Context, in *user.TLUserSaveMusic) (*tg.Bool, error)
	UserGetSavedMusicIdList(ctx context.Context, in *user.TLUserGetSavedMusicIdList) (*user.VectorLong, error)
	UserSetMainProfileTab(ctx context.Context, in *user.TLUserSetMainProfileTab) (*tg.Bool, error)
	UserSetDefaultHistoryTTL(ctx context.Context, in *user.TLUserSetDefaultHistoryTTL) (*tg.Bool, error)
	UserGetDefaultHistoryTTL(ctx context.Context, in *user.TLUserGetDefaultHistoryTTL) (*tg.DefaultHistoryTTL, error)
	UserGetAccountUsername(ctx context.Context, in *user.TLUserGetAccountUsername) (*user.UsernameData, error)
	UserCheckAccountUsername(ctx context.Context, in *user.TLUserCheckAccountUsername) (*user.UsernameExist, error)
	UserGetChannelUsername(ctx context.Context, in *user.TLUserGetChannelUsername) (*user.UsernameData, error)
	UserCheckChannelUsername(ctx context.Context, in *user.TLUserCheckChannelUsername) (*user.UsernameExist, error)
	UserUpdateUsernameByPeer(ctx context.Context, in *user.TLUserUpdateUsernameByPeer) (*tg.Bool, error)
	UserCheckUsername(ctx context.Context, in *user.TLUserCheckUsername) (*user.UsernameExist, error)
	UserUpdateUsernameByUsername(ctx context.Context, in *user.TLUserUpdateUsernameByUsername) (*tg.Bool, error)
	UserDeleteUsername(ctx context.Context, in *user.TLUserDeleteUsername) (*tg.Bool, error)
	UserResolveUsername(ctx context.Context, in *user.TLUserResolveUsername) (*tg.Peer, error)
	UserGetListByUsernameList(ctx context.Context, in *user.TLUserGetListByUsernameList) (*user.VectorUsernameData, error)
	UserDeleteUsernameByPeer(ctx context.Context, in *user.TLUserDeleteUsernameByPeer) (*tg.Bool, error)
	UserSearchUsername(ctx context.Context, in *user.TLUserSearchUsername) (*user.VectorUsernameData, error)
	UserToggleUsername(ctx context.Context, in *user.TLUserToggleUsername) (*tg.Bool, error)
	UserReorderUsernames(ctx context.Context, in *user.TLUserReorderUsernames) (*tg.Bool, error)
	UserDeactivateAllChannelUsernames(ctx context.Context, in *user.TLUserDeactivateAllChannelUsernames) (*tg.Bool, error)
	Close() error
}

type defaultUserClient struct {
	cli client.Client
	rpc userservice.Client
}

func NewUserClient(cli client.Client) UserClient {
	return &defaultUserClient{
		cli: cli,
		rpc: userservice.NewRPCUserClient(cli),
	}
}

func (m *defaultUserClient) Close() error {
	if closer, ok := any(m.cli).(interface{ Close() error }); ok {
		return closer.Close()
	}
	return nil
}

// UserGetLastSeens
// user.getLastSeens id:Vector<long> = Vector<LastSeenData>;
func (m *defaultUserClient) UserGetLastSeens(ctx context.Context, in *user.TLUserGetLastSeens) (*user.VectorLastSeenData, error) {
	return m.rpc.UserGetLastSeens(ctx, in)
}

// UserUpdateLastSeen
// user.updateLastSeen id:long last_seen_at:long expires:int = Bool;
func (m *defaultUserClient) UserUpdateLastSeen(ctx context.Context, in *user.TLUserUpdateLastSeen) (*tg.Bool, error) {
	return m.rpc.UserUpdateLastSeen(ctx, in)
}

// UserGetLastSeen
// user.getLastSeen id:long = LastSeenData;
func (m *defaultUserClient) UserGetLastSeen(ctx context.Context, in *user.TLUserGetLastSeen) (*user.LastSeenData, error) {
	return m.rpc.UserGetLastSeen(ctx, in)
}

// UserGetImmutableUser
// user.getImmutableUser flags:# id:long privacy:flags.1?true contacts:Vector<long> = ImmutableUser;
func (m *defaultUserClient) UserGetImmutableUser(ctx context.Context, in *user.TLUserGetImmutableUser) (*tg.ImmutableUser, error) {
	return m.rpc.UserGetImmutableUser(ctx, in)
}

// UserGetMutableUsers
// user.getMutableUsers id:Vector<long> to:Vector<long> = Vector<ImmutableUser>;
func (m *defaultUserClient) UserGetMutableUsers(ctx context.Context, in *user.TLUserGetMutableUsers) (*user.VectorImmutableUser, error) {
	return m.rpc.UserGetMutableUsers(ctx, in)
}

// UserGetImmutableUserByPhone
// user.getImmutableUserByPhone phone:string = ImmutableUser;
func (m *defaultUserClient) UserGetImmutableUserByPhone(ctx context.Context, in *user.TLUserGetImmutableUserByPhone) (*tg.ImmutableUser, error) {
	return m.rpc.UserGetImmutableUserByPhone(ctx, in)
}

// UserGetImmutableUserByToken
// user.getImmutableUserByToken token:string = ImmutableUser;
func (m *defaultUserClient) UserGetImmutableUserByToken(ctx context.Context, in *user.TLUserGetImmutableUserByToken) (*tg.ImmutableUser, error) {
	return m.rpc.UserGetImmutableUserByToken(ctx, in)
}

// UserSetAccountDaysTTL
// user.setAccountDaysTTL user_id:long ttl:int = Bool;
func (m *defaultUserClient) UserSetAccountDaysTTL(ctx context.Context, in *user.TLUserSetAccountDaysTTL) (*tg.Bool, error) {
	return m.rpc.UserSetAccountDaysTTL(ctx, in)
}

// UserGetAccountDaysTTL
// user.getAccountDaysTTL user_id:long = AccountDaysTTL;
func (m *defaultUserClient) UserGetAccountDaysTTL(ctx context.Context, in *user.TLUserGetAccountDaysTTL) (*tg.AccountDaysTTL, error) {
	return m.rpc.UserGetAccountDaysTTL(ctx, in)
}

// UserGetNotifySettings
// user.getNotifySettings user_id:long peer_type:int peer_id:long = PeerNotifySettings;
func (m *defaultUserClient) UserGetNotifySettings(ctx context.Context, in *user.TLUserGetNotifySettings) (*tg.PeerNotifySettings, error) {
	return m.rpc.UserGetNotifySettings(ctx, in)
}

// UserGetNotifySettingsList
// user.getNotifySettingsList user_id:long peers:Vector<PeerUtil> = Vector<PeerPeerNotifySettings>;
func (m *defaultUserClient) UserGetNotifySettingsList(ctx context.Context, in *user.TLUserGetNotifySettingsList) (*user.VectorPeerPeerNotifySettings, error) {
	return m.rpc.UserGetNotifySettingsList(ctx, in)
}

// UserSetNotifySettings
// user.setNotifySettings user_id:long peer_type:int peer_id:long settings:PeerNotifySettings = Bool;
func (m *defaultUserClient) UserSetNotifySettings(ctx context.Context, in *user.TLUserSetNotifySettings) (*tg.Bool, error) {
	return m.rpc.UserSetNotifySettings(ctx, in)
}

// UserResetNotifySettings
// user.resetNotifySettings user_id:long = Bool;
func (m *defaultUserClient) UserResetNotifySettings(ctx context.Context, in *user.TLUserResetNotifySettings) (*tg.Bool, error) {
	return m.rpc.UserResetNotifySettings(ctx, in)
}

// UserGetAllNotifySettings
// user.getAllNotifySettings user_id:long = Vector<PeerPeerNotifySettings>;
func (m *defaultUserClient) UserGetAllNotifySettings(ctx context.Context, in *user.TLUserGetAllNotifySettings) (*user.VectorPeerPeerNotifySettings, error) {
	return m.rpc.UserGetAllNotifySettings(ctx, in)
}

// UserGetGlobalPrivacySettings
// user.getGlobalPrivacySettings user_id:long = GlobalPrivacySettings;
func (m *defaultUserClient) UserGetGlobalPrivacySettings(ctx context.Context, in *user.TLUserGetGlobalPrivacySettings) (*tg.GlobalPrivacySettings, error) {
	return m.rpc.UserGetGlobalPrivacySettings(ctx, in)
}

// UserSetGlobalPrivacySettings
// user.setGlobalPrivacySettings user_id:long settings:GlobalPrivacySettings = Bool;
func (m *defaultUserClient) UserSetGlobalPrivacySettings(ctx context.Context, in *user.TLUserSetGlobalPrivacySettings) (*tg.Bool, error) {
	return m.rpc.UserSetGlobalPrivacySettings(ctx, in)
}

// UserGetPrivacy
// user.getPrivacy user_id:long key_type:int = Vector<PrivacyRule>;
func (m *defaultUserClient) UserGetPrivacy(ctx context.Context, in *user.TLUserGetPrivacy) (*user.VectorPrivacyRule, error) {
	return m.rpc.UserGetPrivacy(ctx, in)
}

// UserSetPrivacy
// user.setPrivacy user_id:long key_type:int rules:Vector<PrivacyRule> = Bool;
func (m *defaultUserClient) UserSetPrivacy(ctx context.Context, in *user.TLUserSetPrivacy) (*tg.Bool, error) {
	return m.rpc.UserSetPrivacy(ctx, in)
}

// UserCheckPrivacy
// user.checkPrivacy flags:# user_id:long key_type:int peer_id:long = Bool;
func (m *defaultUserClient) UserCheckPrivacy(ctx context.Context, in *user.TLUserCheckPrivacy) (*tg.Bool, error) {
	return m.rpc.UserCheckPrivacy(ctx, in)
}

// UserAddPeerSettings
// user.addPeerSettings user_id:long peer_type:int peer_id:long settings:PeerSettings = Bool;
func (m *defaultUserClient) UserAddPeerSettings(ctx context.Context, in *user.TLUserAddPeerSettings) (*tg.Bool, error) {
	return m.rpc.UserAddPeerSettings(ctx, in)
}

// UserGetPeerSettings
// user.getPeerSettings user_id:long peer_type:int peer_id:long = PeerSettings;
func (m *defaultUserClient) UserGetPeerSettings(ctx context.Context, in *user.TLUserGetPeerSettings) (*tg.PeerSettings, error) {
	return m.rpc.UserGetPeerSettings(ctx, in)
}

// UserDeletePeerSettings
// user.deletePeerSettings user_id:long peer_type:int peer_id:long = Bool;
func (m *defaultUserClient) UserDeletePeerSettings(ctx context.Context, in *user.TLUserDeletePeerSettings) (*tg.Bool, error) {
	return m.rpc.UserDeletePeerSettings(ctx, in)
}

// UserChangePhone
// user.changePhone user_id:long phone:string = Bool;
func (m *defaultUserClient) UserChangePhone(ctx context.Context, in *user.TLUserChangePhone) (*tg.Bool, error) {
	return m.rpc.UserChangePhone(ctx, in)
}

// UserCreateNewUser
// user.createNewUser secret_key_id:long phone:string country_code:string first_name:string last_name:string = ImmutableUser;
func (m *defaultUserClient) UserCreateNewUser(ctx context.Context, in *user.TLUserCreateNewUser) (*tg.ImmutableUser, error) {
	return m.rpc.UserCreateNewUser(ctx, in)
}

// UserDeleteUser
// user.deleteUser user_id:long reason:string phone:string = Bool;
func (m *defaultUserClient) UserDeleteUser(ctx context.Context, in *user.TLUserDeleteUser) (*tg.Bool, error) {
	return m.rpc.UserDeleteUser(ctx, in)
}

// UserBlockPeer
// user.blockPeer user_id:long peer_type:int peer_id:long = Bool;
func (m *defaultUserClient) UserBlockPeer(ctx context.Context, in *user.TLUserBlockPeer) (*tg.Bool, error) {
	return m.rpc.UserBlockPeer(ctx, in)
}

// UserUnBlockPeer
// user.unBlockPeer user_id:long peer_type:int peer_id:long = Bool;
func (m *defaultUserClient) UserUnBlockPeer(ctx context.Context, in *user.TLUserUnBlockPeer) (*tg.Bool, error) {
	return m.rpc.UserUnBlockPeer(ctx, in)
}

// UserBlockedByUser
// user.blockedByUser user_id:long peer_user_id:long = Bool;
func (m *defaultUserClient) UserBlockedByUser(ctx context.Context, in *user.TLUserBlockedByUser) (*tg.Bool, error) {
	return m.rpc.UserBlockedByUser(ctx, in)
}

// UserIsBlockedByUser
// user.isBlockedByUser user_id:long peer_user_id:long = Bool;
func (m *defaultUserClient) UserIsBlockedByUser(ctx context.Context, in *user.TLUserIsBlockedByUser) (*tg.Bool, error) {
	return m.rpc.UserIsBlockedByUser(ctx, in)
}

// UserCheckBlockUserList
// user.checkBlockUserList user_id:long id:Vector<long> = Vector<long>;
func (m *defaultUserClient) UserCheckBlockUserList(ctx context.Context, in *user.TLUserCheckBlockUserList) (*user.VectorLong, error) {
	return m.rpc.UserCheckBlockUserList(ctx, in)
}

// UserGetBlockedList
// user.getBlockedList user_id:long offset:int limit:int = Vector<PeerBlocked>;
func (m *defaultUserClient) UserGetBlockedList(ctx context.Context, in *user.TLUserGetBlockedList) (*user.VectorPeerBlocked, error) {
	return m.rpc.UserGetBlockedList(ctx, in)
}

// UserGetContactSignUpNotification
// user.getContactSignUpNotification user_id:long = Bool;
func (m *defaultUserClient) UserGetContactSignUpNotification(ctx context.Context, in *user.TLUserGetContactSignUpNotification) (*tg.Bool, error) {
	return m.rpc.UserGetContactSignUpNotification(ctx, in)
}

// UserSetContactSignUpNotification
// user.setContactSignUpNotification user_id:long silent:Bool = Bool;
func (m *defaultUserClient) UserSetContactSignUpNotification(ctx context.Context, in *user.TLUserSetContactSignUpNotification) (*tg.Bool, error) {
	return m.rpc.UserSetContactSignUpNotification(ctx, in)
}

// UserGetContentSettings
// user.getContentSettings user_id:long = account.ContentSettings;
func (m *defaultUserClient) UserGetContentSettings(ctx context.Context, in *user.TLUserGetContentSettings) (*tg.AccountContentSettings, error) {
	return m.rpc.UserGetContentSettings(ctx, in)
}

// UserSetContentSettings
// user.setContentSettings flags:# user_id:long sensitive_enabled:flags.0?true = Bool;
func (m *defaultUserClient) UserSetContentSettings(ctx context.Context, in *user.TLUserSetContentSettings) (*tg.Bool, error) {
	return m.rpc.UserSetContentSettings(ctx, in)
}

// UserDeleteContact
// user.deleteContact user_id:long id:long = Bool;
func (m *defaultUserClient) UserDeleteContact(ctx context.Context, in *user.TLUserDeleteContact) (*tg.Bool, error) {
	return m.rpc.UserDeleteContact(ctx, in)
}

// UserGetContactList
// user.getContactList user_id:long = Vector<ContactData>;
func (m *defaultUserClient) UserGetContactList(ctx context.Context, in *user.TLUserGetContactList) (*user.VectorContactData, error) {
	return m.rpc.UserGetContactList(ctx, in)
}

// UserGetContactIdList
// user.getContactIdList user_id:long = Vector<long>;
func (m *defaultUserClient) UserGetContactIdList(ctx context.Context, in *user.TLUserGetContactIdList) (*user.VectorLong, error) {
	return m.rpc.UserGetContactIdList(ctx, in)
}

// UserGetContact
// user.getContact user_id:long id:long = ContactData;
func (m *defaultUserClient) UserGetContact(ctx context.Context, in *user.TLUserGetContact) (*tg.ContactData, error) {
	return m.rpc.UserGetContact(ctx, in)
}

// UserAddContact
// user.addContact user_id:long add_phone_privacy_exception:Bool id:long first_name:string last_name:string phone:string = Bool;
func (m *defaultUserClient) UserAddContact(ctx context.Context, in *user.TLUserAddContact) (*tg.Bool, error) {
	return m.rpc.UserAddContact(ctx, in)
}

// UserCheckContact
// user.checkContact user_id:long id:long = Bool;
func (m *defaultUserClient) UserCheckContact(ctx context.Context, in *user.TLUserCheckContact) (*tg.Bool, error) {
	return m.rpc.UserCheckContact(ctx, in)
}

// UserGetImportersByPhone
// user.getImportersByPhone phone:string = Vector<InputContact>;
func (m *defaultUserClient) UserGetImportersByPhone(ctx context.Context, in *user.TLUserGetImportersByPhone) (*user.VectorInputContact, error) {
	return m.rpc.UserGetImportersByPhone(ctx, in)
}

// UserDeleteImportersByPhone
// user.deleteImportersByPhone phone:string = Bool;
func (m *defaultUserClient) UserDeleteImportersByPhone(ctx context.Context, in *user.TLUserDeleteImportersByPhone) (*tg.Bool, error) {
	return m.rpc.UserDeleteImportersByPhone(ctx, in)
}

// UserImportContacts
// user.importContacts user_id:long contacts:Vector<InputContact> = UserImportedContacts;
func (m *defaultUserClient) UserImportContacts(ctx context.Context, in *user.TLUserImportContacts) (*user.UserImportedContacts, error) {
	return m.rpc.UserImportContacts(ctx, in)
}

// UserGetCountryCode
// user.getCountryCode user_id:long = String;
func (m *defaultUserClient) UserGetCountryCode(ctx context.Context, in *user.TLUserGetCountryCode) (*tg.String, error) {
	return m.rpc.UserGetCountryCode(ctx, in)
}

// UserUpdateAbout
// user.updateAbout user_id:long about:string = Bool;
func (m *defaultUserClient) UserUpdateAbout(ctx context.Context, in *user.TLUserUpdateAbout) (*tg.Bool, error) {
	return m.rpc.UserUpdateAbout(ctx, in)
}

// UserUpdateFirstAndLastName
// user.updateFirstAndLastName user_id:long first_name:string last_name:string = Bool;
func (m *defaultUserClient) UserUpdateFirstAndLastName(ctx context.Context, in *user.TLUserUpdateFirstAndLastName) (*tg.Bool, error) {
	return m.rpc.UserUpdateFirstAndLastName(ctx, in)
}

// UserUpdateVerified
// user.updateVerified user_id:long verified:Bool = Bool;
func (m *defaultUserClient) UserUpdateVerified(ctx context.Context, in *user.TLUserUpdateVerified) (*tg.Bool, error) {
	return m.rpc.UserUpdateVerified(ctx, in)
}

// UserUpdateUsername
// user.updateUsername user_id:long username:string = Bool;
func (m *defaultUserClient) UserUpdateUsername(ctx context.Context, in *user.TLUserUpdateUsername) (*tg.Bool, error) {
	return m.rpc.UserUpdateUsername(ctx, in)
}

// UserUpdateProfilePhoto
// user.updateProfilePhoto user_id:long id:long = Int64;
func (m *defaultUserClient) UserUpdateProfilePhoto(ctx context.Context, in *user.TLUserUpdateProfilePhoto) (*tg.Int64, error) {
	return m.rpc.UserUpdateProfilePhoto(ctx, in)
}

// UserDeleteProfilePhotos
// user.deleteProfilePhotos user_id:long id:Vector<long> = Int64;
func (m *defaultUserClient) UserDeleteProfilePhotos(ctx context.Context, in *user.TLUserDeleteProfilePhotos) (*tg.Int64, error) {
	return m.rpc.UserDeleteProfilePhotos(ctx, in)
}

// UserGetProfilePhotos
// user.getProfilePhotos user_id:long = Vector<long>;
func (m *defaultUserClient) UserGetProfilePhotos(ctx context.Context, in *user.TLUserGetProfilePhotos) (*user.VectorLong, error) {
	return m.rpc.UserGetProfilePhotos(ctx, in)
}

// UserSetBotCommands
// user.setBotCommands user_id:long bot_id:long commands:Vector<BotCommand> = Bool;
func (m *defaultUserClient) UserSetBotCommands(ctx context.Context, in *user.TLUserSetBotCommands) (*tg.Bool, error) {
	return m.rpc.UserSetBotCommands(ctx, in)
}

// UserIsBot
// user.isBot id:long = Bool;
func (m *defaultUserClient) UserIsBot(ctx context.Context, in *user.TLUserIsBot) (*tg.Bool, error) {
	return m.rpc.UserIsBot(ctx, in)
}

// UserGetBotInfo
// user.getBotInfo bot_id:long = BotInfo;
func (m *defaultUserClient) UserGetBotInfo(ctx context.Context, in *user.TLUserGetBotInfo) (*tg.BotInfo, error) {
	return m.rpc.UserGetBotInfo(ctx, in)
}

// UserCheckBots
// user.checkBots id:Vector<long> = Vector<long>;
func (m *defaultUserClient) UserCheckBots(ctx context.Context, in *user.TLUserCheckBots) (*user.VectorLong, error) {
	return m.rpc.UserCheckBots(ctx, in)
}

// UserGetFullUser
// user.getFullUser self_user_id:long id:long = users.UserFull;
func (m *defaultUserClient) UserGetFullUser(ctx context.Context, in *user.TLUserGetFullUser) (*tg.UsersUserFull, error) {
	return m.rpc.UserGetFullUser(ctx, in)
}

// UserUpdateEmojiStatus
// user.updateEmojiStatus user_id:long emoji_status_document_id:long emoji_status_until:int = Bool;
func (m *defaultUserClient) UserUpdateEmojiStatus(ctx context.Context, in *user.TLUserUpdateEmojiStatus) (*tg.Bool, error) {
	return m.rpc.UserUpdateEmojiStatus(ctx, in)
}

// UserGetUserDataById
// user.getUserDataById user_id:long = UserData;
func (m *defaultUserClient) UserGetUserDataById(ctx context.Context, in *user.TLUserGetUserDataById) (*tg.UserData, error) {
	return m.rpc.UserGetUserDataById(ctx, in)
}

// UserGetUserDataListByIdList
// user.getUserDataListByIdList user_id_list:Vector<long> = Vector<UserData>;
func (m *defaultUserClient) UserGetUserDataListByIdList(ctx context.Context, in *user.TLUserGetUserDataListByIdList) (*user.VectorUserData, error) {
	return m.rpc.UserGetUserDataListByIdList(ctx, in)
}

// UserGetUserDataByToken
// user.getUserDataByToken token:string = UserData;
func (m *defaultUserClient) UserGetUserDataByToken(ctx context.Context, in *user.TLUserGetUserDataByToken) (*tg.UserData, error) {
	return m.rpc.UserGetUserDataByToken(ctx, in)
}

// UserSearch
// user.search q:string excluded_contacts:Vector<long> offset:long limit:int = UsersFound;
func (m *defaultUserClient) UserSearch(ctx context.Context, in *user.TLUserSearch) (*user.UsersFound, error) {
	return m.rpc.UserSearch(ctx, in)
}

// UserUpdateBotData
// user.updateBotData flags:# bot_id:long bot_chat_history:flags.15?Bool bot_nochats:flags.16?Bool bot_inline_geo:flags.21?Bool bot_attach_menu:flags.27?Bool bot_inline_placeholder:flags.19?string bot_has_main_app:flags.13?Bool = Bool;
func (m *defaultUserClient) UserUpdateBotData(ctx context.Context, in *user.TLUserUpdateBotData) (*tg.Bool, error) {
	return m.rpc.UserUpdateBotData(ctx, in)
}

// UserGetImmutableUserV2
// user.getImmutableUserV2 flags:# id:long privacy:flags.0?true has_to:flags.2?true to:flags.2?Vector<long> = ImmutableUser;
func (m *defaultUserClient) UserGetImmutableUserV2(ctx context.Context, in *user.TLUserGetImmutableUserV2) (*tg.ImmutableUser, error) {
	return m.rpc.UserGetImmutableUserV2(ctx, in)
}

// UserGetMutableUsersV2
// user.getMutableUsersV2 flags:# id:Vector<long> privacy:flags.0?true has_to:flags.2?true to:flags.2?Vector<long> = MutableUsers;
func (m *defaultUserClient) UserGetMutableUsersV2(ctx context.Context, in *user.TLUserGetMutableUsersV2) (*tg.MutableUsers, error) {
	return m.rpc.UserGetMutableUsersV2(ctx, in)
}

// UserCreateNewTestUser
// user.createNewTestUser secret_key_id:long min_id:long max_id:long = ImmutableUser;
func (m *defaultUserClient) UserCreateNewTestUser(ctx context.Context, in *user.TLUserCreateNewTestUser) (*tg.ImmutableUser, error) {
	return m.rpc.UserCreateNewTestUser(ctx, in)
}

// UserEditCloseFriends
// user.editCloseFriends user_id:long id:Vector<long> = Bool;
func (m *defaultUserClient) UserEditCloseFriends(ctx context.Context, in *user.TLUserEditCloseFriends) (*tg.Bool, error) {
	return m.rpc.UserEditCloseFriends(ctx, in)
}

// UserSetStoriesMaxId
// user.setStoriesMaxId user_id:long id:int = Bool;
func (m *defaultUserClient) UserSetStoriesMaxId(ctx context.Context, in *user.TLUserSetStoriesMaxId) (*tg.Bool, error) {
	return m.rpc.UserSetStoriesMaxId(ctx, in)
}

// UserSetColor
// user.setColor flags:# user_id:long for_profile:flags.1?true color:int background_emoji_id:long = Bool;
func (m *defaultUserClient) UserSetColor(ctx context.Context, in *user.TLUserSetColor) (*tg.Bool, error) {
	return m.rpc.UserSetColor(ctx, in)
}

// UserUpdateBirthday
// user.updateBirthday flags:# user_id:long birthday:flags.1?Birthday = Bool;
func (m *defaultUserClient) UserUpdateBirthday(ctx context.Context, in *user.TLUserUpdateBirthday) (*tg.Bool, error) {
	return m.rpc.UserUpdateBirthday(ctx, in)
}

// UserGetBirthdays
// user.getBirthdays user_id:long = Vector<ContactBirthday>;
func (m *defaultUserClient) UserGetBirthdays(ctx context.Context, in *user.TLUserGetBirthdays) (*user.VectorContactBirthday, error) {
	return m.rpc.UserGetBirthdays(ctx, in)
}

// UserSetStoriesHidden
// user.setStoriesHidden user_id:long id:long hidden:Bool = Bool;
func (m *defaultUserClient) UserSetStoriesHidden(ctx context.Context, in *user.TLUserSetStoriesHidden) (*tg.Bool, error) {
	return m.rpc.UserSetStoriesHidden(ctx, in)
}

// UserUpdatePersonalChannel
// user.updatePersonalChannel user_id:long channel_id:long = Bool;
func (m *defaultUserClient) UserUpdatePersonalChannel(ctx context.Context, in *user.TLUserUpdatePersonalChannel) (*tg.Bool, error) {
	return m.rpc.UserUpdatePersonalChannel(ctx, in)
}

// UserGetUserIdByPhone
// user.getUserIdByPhone phone:string = Int64;
func (m *defaultUserClient) UserGetUserIdByPhone(ctx context.Context, in *user.TLUserGetUserIdByPhone) (*tg.Int64, error) {
	return m.rpc.UserGetUserIdByPhone(ctx, in)
}

// UserSetAuthorizationTTL
// user.setAuthorizationTTL user_id:long ttl:int = Bool;
func (m *defaultUserClient) UserSetAuthorizationTTL(ctx context.Context, in *user.TLUserSetAuthorizationTTL) (*tg.Bool, error) {
	return m.rpc.UserSetAuthorizationTTL(ctx, in)
}

// UserGetAuthorizationTTL
// user.getAuthorizationTTL user_id:long = AccountDaysTTL;
func (m *defaultUserClient) UserGetAuthorizationTTL(ctx context.Context, in *user.TLUserGetAuthorizationTTL) (*tg.AccountDaysTTL, error) {
	return m.rpc.UserGetAuthorizationTTL(ctx, in)
}

// UserUpdatePremium
// user.updatePremium flags:# user_id:long premium:Bool months:flags.1?int = Bool;
func (m *defaultUserClient) UserUpdatePremium(ctx context.Context, in *user.TLUserUpdatePremium) (*tg.Bool, error) {
	return m.rpc.UserUpdatePremium(ctx, in)
}

// UserGetBotInfoV2
// user.getBotInfoV2 bot_id:long = BotInfoData;
func (m *defaultUserClient) UserGetBotInfoV2(ctx context.Context, in *user.TLUserGetBotInfoV2) (*user.BotInfoData, error) {
	return m.rpc.UserGetBotInfoV2(ctx, in)
}

// UserSaveMusic
// user.saveMusic flags:# unsave:flags.0?true user_id:long id:long after_id:flags.15?long = Bool;
func (m *defaultUserClient) UserSaveMusic(ctx context.Context, in *user.TLUserSaveMusic) (*tg.Bool, error) {
	return m.rpc.UserSaveMusic(ctx, in)
}

// UserGetSavedMusicIdList
// user.getSavedMusicIdList user_id:long = Vector<long>;
func (m *defaultUserClient) UserGetSavedMusicIdList(ctx context.Context, in *user.TLUserGetSavedMusicIdList) (*user.VectorLong, error) {
	return m.rpc.UserGetSavedMusicIdList(ctx, in)
}

// UserSetMainProfileTab
// user.setMainProfileTab user_id:long tab:ProfileTab = Bool;
func (m *defaultUserClient) UserSetMainProfileTab(ctx context.Context, in *user.TLUserSetMainProfileTab) (*tg.Bool, error) {
	return m.rpc.UserSetMainProfileTab(ctx, in)
}

// UserSetDefaultHistoryTTL
// user.setDefaultHistoryTTL user_id:long ttl:int = Bool;
func (m *defaultUserClient) UserSetDefaultHistoryTTL(ctx context.Context, in *user.TLUserSetDefaultHistoryTTL) (*tg.Bool, error) {
	return m.rpc.UserSetDefaultHistoryTTL(ctx, in)
}

// UserGetDefaultHistoryTTL
// user.getDefaultHistoryTTL user_id:long = DefaultHistoryTTL;
func (m *defaultUserClient) UserGetDefaultHistoryTTL(ctx context.Context, in *user.TLUserGetDefaultHistoryTTL) (*tg.DefaultHistoryTTL, error) {
	return m.rpc.UserGetDefaultHistoryTTL(ctx, in)
}

// UserGetAccountUsername
// user.getAccountUsername user_id:long = UsernameData;
func (m *defaultUserClient) UserGetAccountUsername(ctx context.Context, in *user.TLUserGetAccountUsername) (*user.UsernameData, error) {
	return m.rpc.UserGetAccountUsername(ctx, in)
}

// UserCheckAccountUsername
// user.checkAccountUsername user_id:long username:string = UsernameExist;
func (m *defaultUserClient) UserCheckAccountUsername(ctx context.Context, in *user.TLUserCheckAccountUsername) (*user.UsernameExist, error) {
	return m.rpc.UserCheckAccountUsername(ctx, in)
}

// UserGetChannelUsername
// user.getChannelUsername channel_id:long = UsernameData;
func (m *defaultUserClient) UserGetChannelUsername(ctx context.Context, in *user.TLUserGetChannelUsername) (*user.UsernameData, error) {
	return m.rpc.UserGetChannelUsername(ctx, in)
}

// UserCheckChannelUsername
// user.checkChannelUsername channel_id:long username:string = UsernameExist;
func (m *defaultUserClient) UserCheckChannelUsername(ctx context.Context, in *user.TLUserCheckChannelUsername) (*user.UsernameExist, error) {
	return m.rpc.UserCheckChannelUsername(ctx, in)
}

// UserUpdateUsernameByPeer
// user.updateUsernameByPeer peer_type:int peer_id:long username:string = Bool;
func (m *defaultUserClient) UserUpdateUsernameByPeer(ctx context.Context, in *user.TLUserUpdateUsernameByPeer) (*tg.Bool, error) {
	return m.rpc.UserUpdateUsernameByPeer(ctx, in)
}

// UserCheckUsername
// user.checkUsername username:string = UsernameExist;
func (m *defaultUserClient) UserCheckUsername(ctx context.Context, in *user.TLUserCheckUsername) (*user.UsernameExist, error) {
	return m.rpc.UserCheckUsername(ctx, in)
}

// UserUpdateUsernameByUsername
// user.updateUsernameByUsername peer_type:int peer_id:long username:string = Bool;
func (m *defaultUserClient) UserUpdateUsernameByUsername(ctx context.Context, in *user.TLUserUpdateUsernameByUsername) (*tg.Bool, error) {
	return m.rpc.UserUpdateUsernameByUsername(ctx, in)
}

// UserDeleteUsername
// user.deleteUsername username:string = Bool;
func (m *defaultUserClient) UserDeleteUsername(ctx context.Context, in *user.TLUserDeleteUsername) (*tg.Bool, error) {
	return m.rpc.UserDeleteUsername(ctx, in)
}

// UserResolveUsername
// user.resolveUsername username:string = Peer;
func (m *defaultUserClient) UserResolveUsername(ctx context.Context, in *user.TLUserResolveUsername) (*tg.Peer, error) {
	return m.rpc.UserResolveUsername(ctx, in)
}

// UserGetListByUsernameList
// user.getListByUsernameList names:Vector<string> = Vector<UsernameData>;
func (m *defaultUserClient) UserGetListByUsernameList(ctx context.Context, in *user.TLUserGetListByUsernameList) (*user.VectorUsernameData, error) {
	return m.rpc.UserGetListByUsernameList(ctx, in)
}

// UserDeleteUsernameByPeer
// user.deleteUsernameByPeer peer_type:int peer_id:long = Bool;
func (m *defaultUserClient) UserDeleteUsernameByPeer(ctx context.Context, in *user.TLUserDeleteUsernameByPeer) (*tg.Bool, error) {
	return m.rpc.UserDeleteUsernameByPeer(ctx, in)
}

// UserSearchUsername
// user.searchUsername q:string excluded_contacts:Vector<long> limit:int = Vector<UsernameData>;
func (m *defaultUserClient) UserSearchUsername(ctx context.Context, in *user.TLUserSearchUsername) (*user.VectorUsernameData, error) {
	return m.rpc.UserSearchUsername(ctx, in)
}

// UserToggleUsername
// user.toggleUsername peer_type:int peer_id:long username:string active:Bool = Bool;
func (m *defaultUserClient) UserToggleUsername(ctx context.Context, in *user.TLUserToggleUsername) (*tg.Bool, error) {
	return m.rpc.UserToggleUsername(ctx, in)
}

// UserReorderUsernames
// user.reorderUsernames peer_type:int peer_id:long username_list:Vector<string> = Bool;
func (m *defaultUserClient) UserReorderUsernames(ctx context.Context, in *user.TLUserReorderUsernames) (*tg.Bool, error) {
	return m.rpc.UserReorderUsernames(ctx, in)
}

// UserDeactivateAllChannelUsernames
// user.deactivateAllChannelUsernames channel_id:long = Bool;
func (m *defaultUserClient) UserDeactivateAllChannelUsernames(ctx context.Context, in *user.TLUserDeactivateAllChannelUsernames) (*tg.Bool, error) {
	return m.rpc.UserDeactivateAllChannelUsernames(ctx, in)
}
