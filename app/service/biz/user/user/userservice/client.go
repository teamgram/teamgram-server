/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package userservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

var _ *tg.Bool

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	UserGetLastSeens(ctx context.Context, req *user.TLUserGetLastSeens, callOptions ...callopt.Option) (r *user.VectorLastSeenData, err error)
	UserUpdateLastSeen(ctx context.Context, req *user.TLUserUpdateLastSeen, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserGetLastSeen(ctx context.Context, req *user.TLUserGetLastSeen, callOptions ...callopt.Option) (r *user.LastSeenData, err error)
	UserGetImmutableUser(ctx context.Context, req *user.TLUserGetImmutableUser, callOptions ...callopt.Option) (r *tg.ImmutableUser, err error)
	UserGetMutableUsers(ctx context.Context, req *user.TLUserGetMutableUsers, callOptions ...callopt.Option) (r *user.VectorImmutableUser, err error)
	UserGetImmutableUserByPhone(ctx context.Context, req *user.TLUserGetImmutableUserByPhone, callOptions ...callopt.Option) (r *tg.ImmutableUser, err error)
	UserGetImmutableUserByToken(ctx context.Context, req *user.TLUserGetImmutableUserByToken, callOptions ...callopt.Option) (r *tg.ImmutableUser, err error)
	UserSetAccountDaysTTL(ctx context.Context, req *user.TLUserSetAccountDaysTTL, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserGetAccountDaysTTL(ctx context.Context, req *user.TLUserGetAccountDaysTTL, callOptions ...callopt.Option) (r *tg.AccountDaysTTL, err error)
	UserGetNotifySettings(ctx context.Context, req *user.TLUserGetNotifySettings, callOptions ...callopt.Option) (r *tg.PeerNotifySettings, err error)
	UserGetNotifySettingsList(ctx context.Context, req *user.TLUserGetNotifySettingsList, callOptions ...callopt.Option) (r *user.VectorPeerPeerNotifySettings, err error)
	UserSetNotifySettings(ctx context.Context, req *user.TLUserSetNotifySettings, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserResetNotifySettings(ctx context.Context, req *user.TLUserResetNotifySettings, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserGetAllNotifySettings(ctx context.Context, req *user.TLUserGetAllNotifySettings, callOptions ...callopt.Option) (r *user.VectorPeerPeerNotifySettings, err error)
	UserGetGlobalPrivacySettings(ctx context.Context, req *user.TLUserGetGlobalPrivacySettings, callOptions ...callopt.Option) (r *tg.GlobalPrivacySettings, err error)
	UserSetGlobalPrivacySettings(ctx context.Context, req *user.TLUserSetGlobalPrivacySettings, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserGetPrivacy(ctx context.Context, req *user.TLUserGetPrivacy, callOptions ...callopt.Option) (r *user.VectorPrivacyRule, err error)
	UserSetPrivacy(ctx context.Context, req *user.TLUserSetPrivacy, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserCheckPrivacy(ctx context.Context, req *user.TLUserCheckPrivacy, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserAddPeerSettings(ctx context.Context, req *user.TLUserAddPeerSettings, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserGetPeerSettings(ctx context.Context, req *user.TLUserGetPeerSettings, callOptions ...callopt.Option) (r *tg.PeerSettings, err error)
	UserDeletePeerSettings(ctx context.Context, req *user.TLUserDeletePeerSettings, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserChangePhone(ctx context.Context, req *user.TLUserChangePhone, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserCreateNewUser(ctx context.Context, req *user.TLUserCreateNewUser, callOptions ...callopt.Option) (r *tg.ImmutableUser, err error)
	UserDeleteUser(ctx context.Context, req *user.TLUserDeleteUser, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserBlockPeer(ctx context.Context, req *user.TLUserBlockPeer, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserUnBlockPeer(ctx context.Context, req *user.TLUserUnBlockPeer, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserBlockedByUser(ctx context.Context, req *user.TLUserBlockedByUser, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserIsBlockedByUser(ctx context.Context, req *user.TLUserIsBlockedByUser, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserCheckBlockUserList(ctx context.Context, req *user.TLUserCheckBlockUserList, callOptions ...callopt.Option) (r *user.VectorLong, err error)
	UserGetBlockedList(ctx context.Context, req *user.TLUserGetBlockedList, callOptions ...callopt.Option) (r *user.VectorPeerBlocked, err error)
	UserGetContactSignUpNotification(ctx context.Context, req *user.TLUserGetContactSignUpNotification, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserSetContactSignUpNotification(ctx context.Context, req *user.TLUserSetContactSignUpNotification, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserGetContentSettings(ctx context.Context, req *user.TLUserGetContentSettings, callOptions ...callopt.Option) (r *tg.AccountContentSettings, err error)
	UserSetContentSettings(ctx context.Context, req *user.TLUserSetContentSettings, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserDeleteContact(ctx context.Context, req *user.TLUserDeleteContact, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserGetContactList(ctx context.Context, req *user.TLUserGetContactList, callOptions ...callopt.Option) (r *user.VectorContactData, err error)
	UserGetContactIdList(ctx context.Context, req *user.TLUserGetContactIdList, callOptions ...callopt.Option) (r *user.VectorLong, err error)
	UserGetContact(ctx context.Context, req *user.TLUserGetContact, callOptions ...callopt.Option) (r *tg.ContactData, err error)
	UserAddContact(ctx context.Context, req *user.TLUserAddContact, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserCheckContact(ctx context.Context, req *user.TLUserCheckContact, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserGetImportersByPhone(ctx context.Context, req *user.TLUserGetImportersByPhone, callOptions ...callopt.Option) (r *user.VectorInputContact, err error)
	UserDeleteImportersByPhone(ctx context.Context, req *user.TLUserDeleteImportersByPhone, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserImportContacts(ctx context.Context, req *user.TLUserImportContacts, callOptions ...callopt.Option) (r *user.UserImportedContacts, err error)
	UserGetCountryCode(ctx context.Context, req *user.TLUserGetCountryCode, callOptions ...callopt.Option) (r *tg.String, err error)
	UserUpdateAbout(ctx context.Context, req *user.TLUserUpdateAbout, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserUpdateFirstAndLastName(ctx context.Context, req *user.TLUserUpdateFirstAndLastName, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserUpdateVerified(ctx context.Context, req *user.TLUserUpdateVerified, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserUpdateUsername(ctx context.Context, req *user.TLUserUpdateUsername, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserUpdateProfilePhoto(ctx context.Context, req *user.TLUserUpdateProfilePhoto, callOptions ...callopt.Option) (r *tg.Int64, err error)
	UserDeleteProfilePhotos(ctx context.Context, req *user.TLUserDeleteProfilePhotos, callOptions ...callopt.Option) (r *tg.Int64, err error)
	UserGetProfilePhotos(ctx context.Context, req *user.TLUserGetProfilePhotos, callOptions ...callopt.Option) (r *user.VectorLong, err error)
	UserSetBotCommands(ctx context.Context, req *user.TLUserSetBotCommands, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserIsBot(ctx context.Context, req *user.TLUserIsBot, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserGetBotInfo(ctx context.Context, req *user.TLUserGetBotInfo, callOptions ...callopt.Option) (r *tg.BotInfo, err error)
	UserCheckBots(ctx context.Context, req *user.TLUserCheckBots, callOptions ...callopt.Option) (r *user.VectorLong, err error)
	UserGetFullUser(ctx context.Context, req *user.TLUserGetFullUser, callOptions ...callopt.Option) (r *tg.UsersUserFull, err error)
	UserUpdateEmojiStatus(ctx context.Context, req *user.TLUserUpdateEmojiStatus, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserGetUserDataById(ctx context.Context, req *user.TLUserGetUserDataById, callOptions ...callopt.Option) (r *tg.UserData, err error)
	UserGetUserDataListByIdList(ctx context.Context, req *user.TLUserGetUserDataListByIdList, callOptions ...callopt.Option) (r *user.VectorUserData, err error)
	UserGetUserDataByToken(ctx context.Context, req *user.TLUserGetUserDataByToken, callOptions ...callopt.Option) (r *tg.UserData, err error)
	UserSearch(ctx context.Context, req *user.TLUserSearch, callOptions ...callopt.Option) (r *user.UsersFound, err error)
	UserUpdateBotData(ctx context.Context, req *user.TLUserUpdateBotData, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserGetImmutableUserV2(ctx context.Context, req *user.TLUserGetImmutableUserV2, callOptions ...callopt.Option) (r *tg.ImmutableUser, err error)
	UserGetMutableUsersV2(ctx context.Context, req *user.TLUserGetMutableUsersV2, callOptions ...callopt.Option) (r *tg.MutableUsers, err error)
	UserCreateNewTestUser(ctx context.Context, req *user.TLUserCreateNewTestUser, callOptions ...callopt.Option) (r *tg.ImmutableUser, err error)
	UserEditCloseFriends(ctx context.Context, req *user.TLUserEditCloseFriends, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserSetStoriesMaxId(ctx context.Context, req *user.TLUserSetStoriesMaxId, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserSetColor(ctx context.Context, req *user.TLUserSetColor, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserUpdateBirthday(ctx context.Context, req *user.TLUserUpdateBirthday, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserGetBirthdays(ctx context.Context, req *user.TLUserGetBirthdays, callOptions ...callopt.Option) (r *user.VectorContactBirthday, err error)
	UserSetStoriesHidden(ctx context.Context, req *user.TLUserSetStoriesHidden, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserUpdatePersonalChannel(ctx context.Context, req *user.TLUserUpdatePersonalChannel, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserGetUserIdByPhone(ctx context.Context, req *user.TLUserGetUserIdByPhone, callOptions ...callopt.Option) (r *tg.Int64, err error)
	UserSetAuthorizationTTL(ctx context.Context, req *user.TLUserSetAuthorizationTTL, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserGetAuthorizationTTL(ctx context.Context, req *user.TLUserGetAuthorizationTTL, callOptions ...callopt.Option) (r *tg.AccountDaysTTL, err error)
	UserUpdatePremium(ctx context.Context, req *user.TLUserUpdatePremium, callOptions ...callopt.Option) (r *tg.Bool, err error)
	UserGetBotInfoV2(ctx context.Context, req *user.TLUserGetBotInfoV2, callOptions ...callopt.Option) (r *user.BotInfoData, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserClient struct {
	*kClient
}

func NewRPCUserClient(cli client.Client) Client {
	return &kUserClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kUserClient) UserGetLastSeens(ctx context.Context, req *user.TLUserGetLastSeens, callOptions ...callopt.Option) (r *user.VectorLastSeenData, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetLastSeens(ctx, req)
}

func (p *kUserClient) UserUpdateLastSeen(ctx context.Context, req *user.TLUserUpdateLastSeen, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserUpdateLastSeen(ctx, req)
}

func (p *kUserClient) UserGetLastSeen(ctx context.Context, req *user.TLUserGetLastSeen, callOptions ...callopt.Option) (r *user.LastSeenData, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetLastSeen(ctx, req)
}

func (p *kUserClient) UserGetImmutableUser(ctx context.Context, req *user.TLUserGetImmutableUser, callOptions ...callopt.Option) (r *tg.ImmutableUser, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetImmutableUser(ctx, req)
}

func (p *kUserClient) UserGetMutableUsers(ctx context.Context, req *user.TLUserGetMutableUsers, callOptions ...callopt.Option) (r *user.VectorImmutableUser, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetMutableUsers(ctx, req)
}

func (p *kUserClient) UserGetImmutableUserByPhone(ctx context.Context, req *user.TLUserGetImmutableUserByPhone, callOptions ...callopt.Option) (r *tg.ImmutableUser, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetImmutableUserByPhone(ctx, req)
}

func (p *kUserClient) UserGetImmutableUserByToken(ctx context.Context, req *user.TLUserGetImmutableUserByToken, callOptions ...callopt.Option) (r *tg.ImmutableUser, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetImmutableUserByToken(ctx, req)
}

func (p *kUserClient) UserSetAccountDaysTTL(ctx context.Context, req *user.TLUserSetAccountDaysTTL, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserSetAccountDaysTTL(ctx, req)
}

func (p *kUserClient) UserGetAccountDaysTTL(ctx context.Context, req *user.TLUserGetAccountDaysTTL, callOptions ...callopt.Option) (r *tg.AccountDaysTTL, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetAccountDaysTTL(ctx, req)
}

func (p *kUserClient) UserGetNotifySettings(ctx context.Context, req *user.TLUserGetNotifySettings, callOptions ...callopt.Option) (r *tg.PeerNotifySettings, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetNotifySettings(ctx, req)
}

func (p *kUserClient) UserGetNotifySettingsList(ctx context.Context, req *user.TLUserGetNotifySettingsList, callOptions ...callopt.Option) (r *user.VectorPeerPeerNotifySettings, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetNotifySettingsList(ctx, req)
}

func (p *kUserClient) UserSetNotifySettings(ctx context.Context, req *user.TLUserSetNotifySettings, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserSetNotifySettings(ctx, req)
}

func (p *kUserClient) UserResetNotifySettings(ctx context.Context, req *user.TLUserResetNotifySettings, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserResetNotifySettings(ctx, req)
}

func (p *kUserClient) UserGetAllNotifySettings(ctx context.Context, req *user.TLUserGetAllNotifySettings, callOptions ...callopt.Option) (r *user.VectorPeerPeerNotifySettings, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetAllNotifySettings(ctx, req)
}

func (p *kUserClient) UserGetGlobalPrivacySettings(ctx context.Context, req *user.TLUserGetGlobalPrivacySettings, callOptions ...callopt.Option) (r *tg.GlobalPrivacySettings, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetGlobalPrivacySettings(ctx, req)
}

func (p *kUserClient) UserSetGlobalPrivacySettings(ctx context.Context, req *user.TLUserSetGlobalPrivacySettings, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserSetGlobalPrivacySettings(ctx, req)
}

func (p *kUserClient) UserGetPrivacy(ctx context.Context, req *user.TLUserGetPrivacy, callOptions ...callopt.Option) (r *user.VectorPrivacyRule, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetPrivacy(ctx, req)
}

func (p *kUserClient) UserSetPrivacy(ctx context.Context, req *user.TLUserSetPrivacy, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserSetPrivacy(ctx, req)
}

func (p *kUserClient) UserCheckPrivacy(ctx context.Context, req *user.TLUserCheckPrivacy, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserCheckPrivacy(ctx, req)
}

func (p *kUserClient) UserAddPeerSettings(ctx context.Context, req *user.TLUserAddPeerSettings, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserAddPeerSettings(ctx, req)
}

func (p *kUserClient) UserGetPeerSettings(ctx context.Context, req *user.TLUserGetPeerSettings, callOptions ...callopt.Option) (r *tg.PeerSettings, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetPeerSettings(ctx, req)
}

func (p *kUserClient) UserDeletePeerSettings(ctx context.Context, req *user.TLUserDeletePeerSettings, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserDeletePeerSettings(ctx, req)
}

func (p *kUserClient) UserChangePhone(ctx context.Context, req *user.TLUserChangePhone, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserChangePhone(ctx, req)
}

func (p *kUserClient) UserCreateNewUser(ctx context.Context, req *user.TLUserCreateNewUser, callOptions ...callopt.Option) (r *tg.ImmutableUser, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserCreateNewUser(ctx, req)
}

func (p *kUserClient) UserDeleteUser(ctx context.Context, req *user.TLUserDeleteUser, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserDeleteUser(ctx, req)
}

func (p *kUserClient) UserBlockPeer(ctx context.Context, req *user.TLUserBlockPeer, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserBlockPeer(ctx, req)
}

func (p *kUserClient) UserUnBlockPeer(ctx context.Context, req *user.TLUserUnBlockPeer, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserUnBlockPeer(ctx, req)
}

func (p *kUserClient) UserBlockedByUser(ctx context.Context, req *user.TLUserBlockedByUser, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserBlockedByUser(ctx, req)
}

func (p *kUserClient) UserIsBlockedByUser(ctx context.Context, req *user.TLUserIsBlockedByUser, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserIsBlockedByUser(ctx, req)
}

func (p *kUserClient) UserCheckBlockUserList(ctx context.Context, req *user.TLUserCheckBlockUserList, callOptions ...callopt.Option) (r *user.VectorLong, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserCheckBlockUserList(ctx, req)
}

func (p *kUserClient) UserGetBlockedList(ctx context.Context, req *user.TLUserGetBlockedList, callOptions ...callopt.Option) (r *user.VectorPeerBlocked, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetBlockedList(ctx, req)
}

func (p *kUserClient) UserGetContactSignUpNotification(ctx context.Context, req *user.TLUserGetContactSignUpNotification, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetContactSignUpNotification(ctx, req)
}

func (p *kUserClient) UserSetContactSignUpNotification(ctx context.Context, req *user.TLUserSetContactSignUpNotification, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserSetContactSignUpNotification(ctx, req)
}

func (p *kUserClient) UserGetContentSettings(ctx context.Context, req *user.TLUserGetContentSettings, callOptions ...callopt.Option) (r *tg.AccountContentSettings, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetContentSettings(ctx, req)
}

func (p *kUserClient) UserSetContentSettings(ctx context.Context, req *user.TLUserSetContentSettings, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserSetContentSettings(ctx, req)
}

func (p *kUserClient) UserDeleteContact(ctx context.Context, req *user.TLUserDeleteContact, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserDeleteContact(ctx, req)
}

func (p *kUserClient) UserGetContactList(ctx context.Context, req *user.TLUserGetContactList, callOptions ...callopt.Option) (r *user.VectorContactData, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetContactList(ctx, req)
}

func (p *kUserClient) UserGetContactIdList(ctx context.Context, req *user.TLUserGetContactIdList, callOptions ...callopt.Option) (r *user.VectorLong, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetContactIdList(ctx, req)
}

func (p *kUserClient) UserGetContact(ctx context.Context, req *user.TLUserGetContact, callOptions ...callopt.Option) (r *tg.ContactData, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetContact(ctx, req)
}

func (p *kUserClient) UserAddContact(ctx context.Context, req *user.TLUserAddContact, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserAddContact(ctx, req)
}

func (p *kUserClient) UserCheckContact(ctx context.Context, req *user.TLUserCheckContact, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserCheckContact(ctx, req)
}

func (p *kUserClient) UserGetImportersByPhone(ctx context.Context, req *user.TLUserGetImportersByPhone, callOptions ...callopt.Option) (r *user.VectorInputContact, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetImportersByPhone(ctx, req)
}

func (p *kUserClient) UserDeleteImportersByPhone(ctx context.Context, req *user.TLUserDeleteImportersByPhone, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserDeleteImportersByPhone(ctx, req)
}

func (p *kUserClient) UserImportContacts(ctx context.Context, req *user.TLUserImportContacts, callOptions ...callopt.Option) (r *user.UserImportedContacts, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserImportContacts(ctx, req)
}

func (p *kUserClient) UserGetCountryCode(ctx context.Context, req *user.TLUserGetCountryCode, callOptions ...callopt.Option) (r *tg.String, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetCountryCode(ctx, req)
}

func (p *kUserClient) UserUpdateAbout(ctx context.Context, req *user.TLUserUpdateAbout, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserUpdateAbout(ctx, req)
}

func (p *kUserClient) UserUpdateFirstAndLastName(ctx context.Context, req *user.TLUserUpdateFirstAndLastName, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserUpdateFirstAndLastName(ctx, req)
}

func (p *kUserClient) UserUpdateVerified(ctx context.Context, req *user.TLUserUpdateVerified, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserUpdateVerified(ctx, req)
}

func (p *kUserClient) UserUpdateUsername(ctx context.Context, req *user.TLUserUpdateUsername, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserUpdateUsername(ctx, req)
}

func (p *kUserClient) UserUpdateProfilePhoto(ctx context.Context, req *user.TLUserUpdateProfilePhoto, callOptions ...callopt.Option) (r *tg.Int64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserUpdateProfilePhoto(ctx, req)
}

func (p *kUserClient) UserDeleteProfilePhotos(ctx context.Context, req *user.TLUserDeleteProfilePhotos, callOptions ...callopt.Option) (r *tg.Int64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserDeleteProfilePhotos(ctx, req)
}

func (p *kUserClient) UserGetProfilePhotos(ctx context.Context, req *user.TLUserGetProfilePhotos, callOptions ...callopt.Option) (r *user.VectorLong, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetProfilePhotos(ctx, req)
}

func (p *kUserClient) UserSetBotCommands(ctx context.Context, req *user.TLUserSetBotCommands, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserSetBotCommands(ctx, req)
}

func (p *kUserClient) UserIsBot(ctx context.Context, req *user.TLUserIsBot, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserIsBot(ctx, req)
}

func (p *kUserClient) UserGetBotInfo(ctx context.Context, req *user.TLUserGetBotInfo, callOptions ...callopt.Option) (r *tg.BotInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetBotInfo(ctx, req)
}

func (p *kUserClient) UserCheckBots(ctx context.Context, req *user.TLUserCheckBots, callOptions ...callopt.Option) (r *user.VectorLong, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserCheckBots(ctx, req)
}

func (p *kUserClient) UserGetFullUser(ctx context.Context, req *user.TLUserGetFullUser, callOptions ...callopt.Option) (r *tg.UsersUserFull, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetFullUser(ctx, req)
}

func (p *kUserClient) UserUpdateEmojiStatus(ctx context.Context, req *user.TLUserUpdateEmojiStatus, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserUpdateEmojiStatus(ctx, req)
}

func (p *kUserClient) UserGetUserDataById(ctx context.Context, req *user.TLUserGetUserDataById, callOptions ...callopt.Option) (r *tg.UserData, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetUserDataById(ctx, req)
}

func (p *kUserClient) UserGetUserDataListByIdList(ctx context.Context, req *user.TLUserGetUserDataListByIdList, callOptions ...callopt.Option) (r *user.VectorUserData, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetUserDataListByIdList(ctx, req)
}

func (p *kUserClient) UserGetUserDataByToken(ctx context.Context, req *user.TLUserGetUserDataByToken, callOptions ...callopt.Option) (r *tg.UserData, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetUserDataByToken(ctx, req)
}

func (p *kUserClient) UserSearch(ctx context.Context, req *user.TLUserSearch, callOptions ...callopt.Option) (r *user.UsersFound, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserSearch(ctx, req)
}

func (p *kUserClient) UserUpdateBotData(ctx context.Context, req *user.TLUserUpdateBotData, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserUpdateBotData(ctx, req)
}

func (p *kUserClient) UserGetImmutableUserV2(ctx context.Context, req *user.TLUserGetImmutableUserV2, callOptions ...callopt.Option) (r *tg.ImmutableUser, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetImmutableUserV2(ctx, req)
}

func (p *kUserClient) UserGetMutableUsersV2(ctx context.Context, req *user.TLUserGetMutableUsersV2, callOptions ...callopt.Option) (r *tg.MutableUsers, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetMutableUsersV2(ctx, req)
}

func (p *kUserClient) UserCreateNewTestUser(ctx context.Context, req *user.TLUserCreateNewTestUser, callOptions ...callopt.Option) (r *tg.ImmutableUser, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserCreateNewTestUser(ctx, req)
}

func (p *kUserClient) UserEditCloseFriends(ctx context.Context, req *user.TLUserEditCloseFriends, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserEditCloseFriends(ctx, req)
}

func (p *kUserClient) UserSetStoriesMaxId(ctx context.Context, req *user.TLUserSetStoriesMaxId, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserSetStoriesMaxId(ctx, req)
}

func (p *kUserClient) UserSetColor(ctx context.Context, req *user.TLUserSetColor, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserSetColor(ctx, req)
}

func (p *kUserClient) UserUpdateBirthday(ctx context.Context, req *user.TLUserUpdateBirthday, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserUpdateBirthday(ctx, req)
}

func (p *kUserClient) UserGetBirthdays(ctx context.Context, req *user.TLUserGetBirthdays, callOptions ...callopt.Option) (r *user.VectorContactBirthday, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetBirthdays(ctx, req)
}

func (p *kUserClient) UserSetStoriesHidden(ctx context.Context, req *user.TLUserSetStoriesHidden, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserSetStoriesHidden(ctx, req)
}

func (p *kUserClient) UserUpdatePersonalChannel(ctx context.Context, req *user.TLUserUpdatePersonalChannel, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserUpdatePersonalChannel(ctx, req)
}

func (p *kUserClient) UserGetUserIdByPhone(ctx context.Context, req *user.TLUserGetUserIdByPhone, callOptions ...callopt.Option) (r *tg.Int64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetUserIdByPhone(ctx, req)
}

func (p *kUserClient) UserSetAuthorizationTTL(ctx context.Context, req *user.TLUserSetAuthorizationTTL, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserSetAuthorizationTTL(ctx, req)
}

func (p *kUserClient) UserGetAuthorizationTTL(ctx context.Context, req *user.TLUserGetAuthorizationTTL, callOptions ...callopt.Option) (r *tg.AccountDaysTTL, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetAuthorizationTTL(ctx, req)
}

func (p *kUserClient) UserUpdatePremium(ctx context.Context, req *user.TLUserUpdatePremium, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserUpdatePremium(ctx, req)
}

func (p *kUserClient) UserGetBotInfoV2(ctx context.Context, req *user.TLUserGetBotInfoV2, callOptions ...callopt.Option) (r *user.BotInfoData, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserGetBotInfoV2(ctx, req)
}
