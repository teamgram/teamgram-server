/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package user_client

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
	UserGetImmutableUser(ctx context.Context, in *user.TLUserGetImmutableUser) (*user.ImmutableUser, error)
	UserGetMutableUsers(ctx context.Context, in *user.TLUserGetMutableUsers) (*user.Vector_ImmutableUser, error)
	UserGetImmutableUserByPhone(ctx context.Context, in *user.TLUserGetImmutableUserByPhone) (*user.ImmutableUser, error)
	UserGetImmutableUserByToken(ctx context.Context, in *user.TLUserGetImmutableUserByToken) (*user.ImmutableUser, error)
	UserSetAccountDaysTTL(ctx context.Context, in *user.TLUserSetAccountDaysTTL) (*mtproto.Bool, error)
	UserGetAccountDaysTTL(ctx context.Context, in *user.TLUserGetAccountDaysTTL) (*mtproto.AccountDaysTTL, error)
	UserGetNotifySettings(ctx context.Context, in *user.TLUserGetNotifySettings) (*mtproto.PeerNotifySettings, error)
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
	UserCreateNewPredefinedUser(ctx context.Context, in *user.TLUserCreateNewPredefinedUser) (*mtproto.PredefinedUser, error)
	UserGetPredefinedUser(ctx context.Context, in *user.TLUserGetPredefinedUser) (*mtproto.PredefinedUser, error)
	UserGetAllPredefinedUser(ctx context.Context, in *user.TLUserGetAllPredefinedUser) (*user.Vector_PredefinedUser, error)
	UserUpdatePredefinedFirstAndLastName(ctx context.Context, in *user.TLUserUpdatePredefinedFirstAndLastName) (*mtproto.PredefinedUser, error)
	UserUpdatePredefinedVerified(ctx context.Context, in *user.TLUserUpdatePredefinedVerified) (*mtproto.PredefinedUser, error)
	UserUpdatePredefinedUsername(ctx context.Context, in *user.TLUserUpdatePredefinedUsername) (*mtproto.PredefinedUser, error)
	UserUpdatePredefinedCode(ctx context.Context, in *user.TLUserUpdatePredefinedCode) (*mtproto.PredefinedUser, error)
	UserPredefinedBindRegisteredUserId(ctx context.Context, in *user.TLUserPredefinedBindRegisteredUserId) (*mtproto.Bool, error)
	UserCreateNewUser(ctx context.Context, in *user.TLUserCreateNewUser) (*user.ImmutableUser, error)
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
	UserGetContact(ctx context.Context, in *user.TLUserGetContact) (*user.ContactData, error)
	UserAddContact(ctx context.Context, in *user.TLUserAddContact) (*mtproto.Bool, error)
	UserCheckContact(ctx context.Context, in *user.TLUserCheckContact) (*mtproto.Bool, error)
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
	UserGetFullUser(ctx context.Context, in *user.TLUserGetFullUser) (*mtproto.Users_UserFull, error)
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
// user.updateLastSeen id:long last_seen_at:long expries:int = Bool;
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
// user.getImmutableUser id:long = ImmutableUser;
func (m *defaultUserClient) UserGetImmutableUser(ctx context.Context, in *user.TLUserGetImmutableUser) (*user.ImmutableUser, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetImmutableUser(ctx, in)
}

// UserGetMutableUsers
// user.getMutableUsers id:Vector<long> = Vector<ImmutableUser>;
func (m *defaultUserClient) UserGetMutableUsers(ctx context.Context, in *user.TLUserGetMutableUsers) (*user.Vector_ImmutableUser, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetMutableUsers(ctx, in)
}

// UserGetImmutableUserByPhone
// user.getImmutableUserByPhone phone:string = ImmutableUser;
func (m *defaultUserClient) UserGetImmutableUserByPhone(ctx context.Context, in *user.TLUserGetImmutableUserByPhone) (*user.ImmutableUser, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetImmutableUserByPhone(ctx, in)
}

// UserGetImmutableUserByToken
// user.getImmutableUserByToken token:string = ImmutableUser;
func (m *defaultUserClient) UserGetImmutableUserByToken(ctx context.Context, in *user.TLUserGetImmutableUserByToken) (*user.ImmutableUser, error) {
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

// UserCreateNewPredefinedUser
// user.createNewPredefinedUser flags:# phone:string first_name:string last_name:flags.0?string username:string code:string verified:flags.1?true = PredefinedUser;
func (m *defaultUserClient) UserCreateNewPredefinedUser(ctx context.Context, in *user.TLUserCreateNewPredefinedUser) (*mtproto.PredefinedUser, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserCreateNewPredefinedUser(ctx, in)
}

// UserGetPredefinedUser
// user.getPredefinedUser phone:string = PredefinedUser;
func (m *defaultUserClient) UserGetPredefinedUser(ctx context.Context, in *user.TLUserGetPredefinedUser) (*mtproto.PredefinedUser, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetPredefinedUser(ctx, in)
}

// UserGetAllPredefinedUser
// user.getAllPredefinedUser = Vector<PredefinedUser>;
func (m *defaultUserClient) UserGetAllPredefinedUser(ctx context.Context, in *user.TLUserGetAllPredefinedUser) (*user.Vector_PredefinedUser, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetAllPredefinedUser(ctx, in)
}

// UserUpdatePredefinedFirstAndLastName
// user.updatePredefinedFirstAndLastName flags:# phone:string first_name:string last_name:flags.0?string = PredefinedUser;
func (m *defaultUserClient) UserUpdatePredefinedFirstAndLastName(ctx context.Context, in *user.TLUserUpdatePredefinedFirstAndLastName) (*mtproto.PredefinedUser, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserUpdatePredefinedFirstAndLastName(ctx, in)
}

// UserUpdatePredefinedVerified
// user.updatePredefinedVerified flags:# phone:string verified:flags.1?true = PredefinedUser;
func (m *defaultUserClient) UserUpdatePredefinedVerified(ctx context.Context, in *user.TLUserUpdatePredefinedVerified) (*mtproto.PredefinedUser, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserUpdatePredefinedVerified(ctx, in)
}

// UserUpdatePredefinedUsername
// user.updatePredefinedUsername flags:# phone:string username:flags.1?string = PredefinedUser;
func (m *defaultUserClient) UserUpdatePredefinedUsername(ctx context.Context, in *user.TLUserUpdatePredefinedUsername) (*mtproto.PredefinedUser, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserUpdatePredefinedUsername(ctx, in)
}

// UserUpdatePredefinedCode
// user.updatePredefinedCode phone:string code:string = PredefinedUser;
func (m *defaultUserClient) UserUpdatePredefinedCode(ctx context.Context, in *user.TLUserUpdatePredefinedCode) (*mtproto.PredefinedUser, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserUpdatePredefinedCode(ctx, in)
}

// UserPredefinedBindRegisteredUserId
// user.predefinedBindRegisteredUserId phone:string registered_userId:long = Bool;
func (m *defaultUserClient) UserPredefinedBindRegisteredUserId(ctx context.Context, in *user.TLUserPredefinedBindRegisteredUserId) (*mtproto.Bool, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserPredefinedBindRegisteredUserId(ctx, in)
}

// UserCreateNewUser
// user.createNewUser secret_key_id:long phone:string country_code:string first_name:string last_name:string = ImmutableUser;
func (m *defaultUserClient) UserCreateNewUser(ctx context.Context, in *user.TLUserCreateNewUser) (*user.ImmutableUser, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserCreateNewUser(ctx, in)
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
func (m *defaultUserClient) UserGetContact(ctx context.Context, in *user.TLUserGetContact) (*user.ContactData, error) {
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

// UserGetFullUser
// user.getFullUser self_user_id:long id:long = users.UserFull;
func (m *defaultUserClient) UserGetFullUser(ctx context.Context, in *user.TLUserGetFullUser) (*mtproto.Users_UserFull, error) {
	client := user.NewRPCUserClient(m.cli.Conn())
	return client.UserGetFullUser(ctx, in)
}
