/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2026 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/internal/core"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

var _ *tg.Bool

// UserGetLastSeens
// user.getLastSeens id:Vector<long> = Vector<LastSeenData>;
func (s *Service) UserGetLastSeens(ctx context.Context, request *user.TLUserGetLastSeens) (*user.VectorLastSeenData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getLastSeens - request: %s", request)

	r, err := c.UserGetLastSeens(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getLastSeens - reply: %s", r)
	return r, err
}

// UserUpdateLastSeen
// user.updateLastSeen id:long last_seen_at:long expires:int = Bool;
func (s *Service) UserUpdateLastSeen(ctx context.Context, request *user.TLUserUpdateLastSeen) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.updateLastSeen - request: %s", request)

	r, err := c.UserUpdateLastSeen(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.updateLastSeen - reply: %s", r)
	return r, err
}

// UserGetLastSeen
// user.getLastSeen id:long = LastSeenData;
func (s *Service) UserGetLastSeen(ctx context.Context, request *user.TLUserGetLastSeen) (*user.LastSeenData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getLastSeen - request: %s", request)

	r, err := c.UserGetLastSeen(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getLastSeen - reply: %s", r)
	return r, err
}

// UserGetImmutableUser
// user.getImmutableUser flags:# id:long privacy:flags.1?true contacts:Vector<long> = ImmutableUser;
func (s *Service) UserGetImmutableUser(ctx context.Context, request *user.TLUserGetImmutableUser) (*tg.ImmutableUser, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getImmutableUser - request: %s", request)

	r, err := c.UserGetImmutableUser(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getImmutableUser - reply: %s", r)
	return r, err
}

// UserGetMutableUsers
// user.getMutableUsers id:Vector<long> to:Vector<long> = Vector<ImmutableUser>;
func (s *Service) UserGetMutableUsers(ctx context.Context, request *user.TLUserGetMutableUsers) (*user.VectorImmutableUser, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getMutableUsers - request: %s", request)

	r, err := c.UserGetMutableUsers(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getMutableUsers - reply: %s", r)
	return r, err
}

// UserGetImmutableUserByPhone
// user.getImmutableUserByPhone phone:string = ImmutableUser;
func (s *Service) UserGetImmutableUserByPhone(ctx context.Context, request *user.TLUserGetImmutableUserByPhone) (*tg.ImmutableUser, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getImmutableUserByPhone - request: %s", request)

	r, err := c.UserGetImmutableUserByPhone(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getImmutableUserByPhone - reply: %s", r)
	return r, err
}

// UserGetImmutableUserByToken
// user.getImmutableUserByToken token:string = ImmutableUser;
func (s *Service) UserGetImmutableUserByToken(ctx context.Context, request *user.TLUserGetImmutableUserByToken) (*tg.ImmutableUser, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getImmutableUserByToken - request: %s", request)

	r, err := c.UserGetImmutableUserByToken(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getImmutableUserByToken - reply: %s", r)
	return r, err
}

// UserSetAccountDaysTTL
// user.setAccountDaysTTL user_id:long ttl:int = Bool;
func (s *Service) UserSetAccountDaysTTL(ctx context.Context, request *user.TLUserSetAccountDaysTTL) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.setAccountDaysTTL - request: %s", request)

	r, err := c.UserSetAccountDaysTTL(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.setAccountDaysTTL - reply: %s", r)
	return r, err
}

// UserGetAccountDaysTTL
// user.getAccountDaysTTL user_id:long = AccountDaysTTL;
func (s *Service) UserGetAccountDaysTTL(ctx context.Context, request *user.TLUserGetAccountDaysTTL) (*tg.AccountDaysTTL, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getAccountDaysTTL - request: %s", request)

	r, err := c.UserGetAccountDaysTTL(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getAccountDaysTTL - reply: %s", r)
	return r, err
}

// UserGetNotifySettings
// user.getNotifySettings user_id:long peer_type:int peer_id:long = PeerNotifySettings;
func (s *Service) UserGetNotifySettings(ctx context.Context, request *user.TLUserGetNotifySettings) (*tg.PeerNotifySettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getNotifySettings - request: %s", request)

	r, err := c.UserGetNotifySettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getNotifySettings - reply: %s", r)
	return r, err
}

// UserGetNotifySettingsList
// user.getNotifySettingsList user_id:long peers:Vector<PeerUtil> = Vector<PeerPeerNotifySettings>;
func (s *Service) UserGetNotifySettingsList(ctx context.Context, request *user.TLUserGetNotifySettingsList) (*user.VectorPeerPeerNotifySettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getNotifySettingsList - request: %s", request)

	r, err := c.UserGetNotifySettingsList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getNotifySettingsList - reply: %s", r)
	return r, err
}

// UserSetNotifySettings
// user.setNotifySettings user_id:long peer_type:int peer_id:long settings:PeerNotifySettings = Bool;
func (s *Service) UserSetNotifySettings(ctx context.Context, request *user.TLUserSetNotifySettings) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.setNotifySettings - request: %s", request)

	r, err := c.UserSetNotifySettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.setNotifySettings - reply: %s", r)
	return r, err
}

// UserResetNotifySettings
// user.resetNotifySettings user_id:long = Bool;
func (s *Service) UserResetNotifySettings(ctx context.Context, request *user.TLUserResetNotifySettings) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.resetNotifySettings - request: %s", request)

	r, err := c.UserResetNotifySettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.resetNotifySettings - reply: %s", r)
	return r, err
}

// UserGetAllNotifySettings
// user.getAllNotifySettings user_id:long = Vector<PeerPeerNotifySettings>;
func (s *Service) UserGetAllNotifySettings(ctx context.Context, request *user.TLUserGetAllNotifySettings) (*user.VectorPeerPeerNotifySettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getAllNotifySettings - request: %s", request)

	r, err := c.UserGetAllNotifySettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getAllNotifySettings - reply: %s", r)
	return r, err
}

// UserGetGlobalPrivacySettings
// user.getGlobalPrivacySettings user_id:long = GlobalPrivacySettings;
func (s *Service) UserGetGlobalPrivacySettings(ctx context.Context, request *user.TLUserGetGlobalPrivacySettings) (*tg.GlobalPrivacySettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getGlobalPrivacySettings - request: %s", request)

	r, err := c.UserGetGlobalPrivacySettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getGlobalPrivacySettings - reply: %s", r)
	return r, err
}

// UserSetGlobalPrivacySettings
// user.setGlobalPrivacySettings user_id:long settings:GlobalPrivacySettings = Bool;
func (s *Service) UserSetGlobalPrivacySettings(ctx context.Context, request *user.TLUserSetGlobalPrivacySettings) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.setGlobalPrivacySettings - request: %s", request)

	r, err := c.UserSetGlobalPrivacySettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.setGlobalPrivacySettings - reply: %s", r)
	return r, err
}

// UserGetPrivacy
// user.getPrivacy user_id:long key_type:int = Vector<PrivacyRule>;
func (s *Service) UserGetPrivacy(ctx context.Context, request *user.TLUserGetPrivacy) (*user.VectorPrivacyRule, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getPrivacy - request: %s", request)

	r, err := c.UserGetPrivacy(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getPrivacy - reply: %s", r)
	return r, err
}

// UserSetPrivacy
// user.setPrivacy user_id:long key_type:int rules:Vector<PrivacyRule> = Bool;
func (s *Service) UserSetPrivacy(ctx context.Context, request *user.TLUserSetPrivacy) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.setPrivacy - request: %s", request)

	r, err := c.UserSetPrivacy(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.setPrivacy - reply: %s", r)
	return r, err
}

// UserCheckPrivacy
// user.checkPrivacy flags:# user_id:long key_type:int peer_id:long = Bool;
func (s *Service) UserCheckPrivacy(ctx context.Context, request *user.TLUserCheckPrivacy) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.checkPrivacy - request: %s", request)

	r, err := c.UserCheckPrivacy(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.checkPrivacy - reply: %s", r)
	return r, err
}

// UserAddPeerSettings
// user.addPeerSettings user_id:long peer_type:int peer_id:long settings:PeerSettings = Bool;
func (s *Service) UserAddPeerSettings(ctx context.Context, request *user.TLUserAddPeerSettings) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.addPeerSettings - request: %s", request)

	r, err := c.UserAddPeerSettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.addPeerSettings - reply: %s", r)
	return r, err
}

// UserGetPeerSettings
// user.getPeerSettings user_id:long peer_type:int peer_id:long = PeerSettings;
func (s *Service) UserGetPeerSettings(ctx context.Context, request *user.TLUserGetPeerSettings) (*tg.PeerSettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getPeerSettings - request: %s", request)

	r, err := c.UserGetPeerSettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getPeerSettings - reply: %s", r)
	return r, err
}

// UserDeletePeerSettings
// user.deletePeerSettings user_id:long peer_type:int peer_id:long = Bool;
func (s *Service) UserDeletePeerSettings(ctx context.Context, request *user.TLUserDeletePeerSettings) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.deletePeerSettings - request: %s", request)

	r, err := c.UserDeletePeerSettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.deletePeerSettings - reply: %s", r)
	return r, err
}

// UserChangePhone
// user.changePhone user_id:long phone:string = Bool;
func (s *Service) UserChangePhone(ctx context.Context, request *user.TLUserChangePhone) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.changePhone - request: %s", request)

	r, err := c.UserChangePhone(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.changePhone - reply: %s", r)
	return r, err
}

// UserCreateNewUser
// user.createNewUser secret_key_id:long phone:string country_code:string first_name:string last_name:string = ImmutableUser;
func (s *Service) UserCreateNewUser(ctx context.Context, request *user.TLUserCreateNewUser) (*tg.ImmutableUser, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.createNewUser - request: %s", request)

	r, err := c.UserCreateNewUser(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.createNewUser - reply: %s", r)
	return r, err
}

// UserDeleteUser
// user.deleteUser user_id:long reason:string phone:string = Bool;
func (s *Service) UserDeleteUser(ctx context.Context, request *user.TLUserDeleteUser) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.deleteUser - request: %s", request)

	r, err := c.UserDeleteUser(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.deleteUser - reply: %s", r)
	return r, err
}

// UserBlockPeer
// user.blockPeer user_id:long peer_type:int peer_id:long = Bool;
func (s *Service) UserBlockPeer(ctx context.Context, request *user.TLUserBlockPeer) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.blockPeer - request: %s", request)

	r, err := c.UserBlockPeer(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.blockPeer - reply: %s", r)
	return r, err
}

// UserUnBlockPeer
// user.unBlockPeer user_id:long peer_type:int peer_id:long = Bool;
func (s *Service) UserUnBlockPeer(ctx context.Context, request *user.TLUserUnBlockPeer) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.unBlockPeer - request: %s", request)

	r, err := c.UserUnBlockPeer(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.unBlockPeer - reply: %s", r)
	return r, err
}

// UserBlockedByUser
// user.blockedByUser user_id:long peer_user_id:long = Bool;
func (s *Service) UserBlockedByUser(ctx context.Context, request *user.TLUserBlockedByUser) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.blockedByUser - request: %s", request)

	r, err := c.UserBlockedByUser(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.blockedByUser - reply: %s", r)
	return r, err
}

// UserIsBlockedByUser
// user.isBlockedByUser user_id:long peer_user_id:long = Bool;
func (s *Service) UserIsBlockedByUser(ctx context.Context, request *user.TLUserIsBlockedByUser) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.isBlockedByUser - request: %s", request)

	r, err := c.UserIsBlockedByUser(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.isBlockedByUser - reply: %s", r)
	return r, err
}

// UserCheckBlockUserList
// user.checkBlockUserList user_id:long id:Vector<long> = Vector<long>;
func (s *Service) UserCheckBlockUserList(ctx context.Context, request *user.TLUserCheckBlockUserList) (*user.VectorLong, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.checkBlockUserList - request: %s", request)

	r, err := c.UserCheckBlockUserList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.checkBlockUserList - reply: %s", r)
	return r, err
}

// UserGetBlockedList
// user.getBlockedList user_id:long offset:int limit:int = Vector<PeerBlocked>;
func (s *Service) UserGetBlockedList(ctx context.Context, request *user.TLUserGetBlockedList) (*user.VectorPeerBlocked, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getBlockedList - request: %s", request)

	r, err := c.UserGetBlockedList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getBlockedList - reply: %s", r)
	return r, err
}

// UserGetContactSignUpNotification
// user.getContactSignUpNotification user_id:long = Bool;
func (s *Service) UserGetContactSignUpNotification(ctx context.Context, request *user.TLUserGetContactSignUpNotification) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getContactSignUpNotification - request: %s", request)

	r, err := c.UserGetContactSignUpNotification(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getContactSignUpNotification - reply: %s", r)
	return r, err
}

// UserSetContactSignUpNotification
// user.setContactSignUpNotification user_id:long silent:Bool = Bool;
func (s *Service) UserSetContactSignUpNotification(ctx context.Context, request *user.TLUserSetContactSignUpNotification) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.setContactSignUpNotification - request: %s", request)

	r, err := c.UserSetContactSignUpNotification(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.setContactSignUpNotification - reply: %s", r)
	return r, err
}

// UserGetContentSettings
// user.getContentSettings user_id:long = account.ContentSettings;
func (s *Service) UserGetContentSettings(ctx context.Context, request *user.TLUserGetContentSettings) (*tg.AccountContentSettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getContentSettings - request: %s", request)

	r, err := c.UserGetContentSettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getContentSettings - reply: %s", r)
	return r, err
}

// UserSetContentSettings
// user.setContentSettings flags:# user_id:long sensitive_enabled:flags.0?true = Bool;
func (s *Service) UserSetContentSettings(ctx context.Context, request *user.TLUserSetContentSettings) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.setContentSettings - request: %s", request)

	r, err := c.UserSetContentSettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.setContentSettings - reply: %s", r)
	return r, err
}

// UserDeleteContact
// user.deleteContact user_id:long id:long = Bool;
func (s *Service) UserDeleteContact(ctx context.Context, request *user.TLUserDeleteContact) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.deleteContact - request: %s", request)

	r, err := c.UserDeleteContact(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.deleteContact - reply: %s", r)
	return r, err
}

// UserGetContactList
// user.getContactList user_id:long = Vector<ContactData>;
func (s *Service) UserGetContactList(ctx context.Context, request *user.TLUserGetContactList) (*user.VectorContactData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getContactList - request: %s", request)

	r, err := c.UserGetContactList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getContactList - reply: %s", r)
	return r, err
}

// UserGetContactIdList
// user.getContactIdList user_id:long = Vector<long>;
func (s *Service) UserGetContactIdList(ctx context.Context, request *user.TLUserGetContactIdList) (*user.VectorLong, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getContactIdList - request: %s", request)

	r, err := c.UserGetContactIdList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getContactIdList - reply: %s", r)
	return r, err
}

// UserGetContact
// user.getContact user_id:long id:long = ContactData;
func (s *Service) UserGetContact(ctx context.Context, request *user.TLUserGetContact) (*tg.ContactData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getContact - request: %s", request)

	r, err := c.UserGetContact(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getContact - reply: %s", r)
	return r, err
}

// UserAddContact
// user.addContact user_id:long add_phone_privacy_exception:Bool id:long first_name:string last_name:string phone:string = Bool;
func (s *Service) UserAddContact(ctx context.Context, request *user.TLUserAddContact) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.addContact - request: %s", request)

	r, err := c.UserAddContact(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.addContact - reply: %s", r)
	return r, err
}

// UserCheckContact
// user.checkContact user_id:long id:long = Bool;
func (s *Service) UserCheckContact(ctx context.Context, request *user.TLUserCheckContact) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.checkContact - request: %s", request)

	r, err := c.UserCheckContact(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.checkContact - reply: %s", r)
	return r, err
}

// UserGetImportersByPhone
// user.getImportersByPhone phone:string = Vector<InputContact>;
func (s *Service) UserGetImportersByPhone(ctx context.Context, request *user.TLUserGetImportersByPhone) (*user.VectorInputContact, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getImportersByPhone - request: %s", request)

	r, err := c.UserGetImportersByPhone(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getImportersByPhone - reply: %s", r)
	return r, err
}

// UserDeleteImportersByPhone
// user.deleteImportersByPhone phone:string = Bool;
func (s *Service) UserDeleteImportersByPhone(ctx context.Context, request *user.TLUserDeleteImportersByPhone) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.deleteImportersByPhone - request: %s", request)

	r, err := c.UserDeleteImportersByPhone(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.deleteImportersByPhone - reply: %s", r)
	return r, err
}

// UserImportContacts
// user.importContacts user_id:long contacts:Vector<InputContact> = UserImportedContacts;
func (s *Service) UserImportContacts(ctx context.Context, request *user.TLUserImportContacts) (*user.UserImportedContacts, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.importContacts - request: %s", request)

	r, err := c.UserImportContacts(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.importContacts - reply: %s", r)
	return r, err
}

// UserGetCountryCode
// user.getCountryCode user_id:long = String;
func (s *Service) UserGetCountryCode(ctx context.Context, request *user.TLUserGetCountryCode) (*tg.String, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getCountryCode - request: %s", request)

	r, err := c.UserGetCountryCode(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getCountryCode - reply: %s", r)
	return r, err
}

// UserUpdateAbout
// user.updateAbout user_id:long about:string = Bool;
func (s *Service) UserUpdateAbout(ctx context.Context, request *user.TLUserUpdateAbout) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.updateAbout - request: %s", request)

	r, err := c.UserUpdateAbout(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.updateAbout - reply: %s", r)
	return r, err
}

// UserUpdateFirstAndLastName
// user.updateFirstAndLastName user_id:long first_name:string last_name:string = Bool;
func (s *Service) UserUpdateFirstAndLastName(ctx context.Context, request *user.TLUserUpdateFirstAndLastName) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.updateFirstAndLastName - request: %s", request)

	r, err := c.UserUpdateFirstAndLastName(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.updateFirstAndLastName - reply: %s", r)
	return r, err
}

// UserUpdateVerified
// user.updateVerified user_id:long verified:Bool = Bool;
func (s *Service) UserUpdateVerified(ctx context.Context, request *user.TLUserUpdateVerified) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.updateVerified - request: %s", request)

	r, err := c.UserUpdateVerified(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.updateVerified - reply: %s", r)
	return r, err
}

// UserUpdateUsername
// user.updateUsername user_id:long username:string = Bool;
func (s *Service) UserUpdateUsername(ctx context.Context, request *user.TLUserUpdateUsername) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.updateUsername - request: %s", request)

	r, err := c.UserUpdateUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.updateUsername - reply: %s", r)
	return r, err
}

// UserUpdateProfilePhoto
// user.updateProfilePhoto user_id:long id:long = Int64;
func (s *Service) UserUpdateProfilePhoto(ctx context.Context, request *user.TLUserUpdateProfilePhoto) (*tg.Int64, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.updateProfilePhoto - request: %s", request)

	r, err := c.UserUpdateProfilePhoto(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.updateProfilePhoto - reply: %s", r)
	return r, err
}

// UserDeleteProfilePhotos
// user.deleteProfilePhotos user_id:long id:Vector<long> = Int64;
func (s *Service) UserDeleteProfilePhotos(ctx context.Context, request *user.TLUserDeleteProfilePhotos) (*tg.Int64, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.deleteProfilePhotos - request: %s", request)

	r, err := c.UserDeleteProfilePhotos(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.deleteProfilePhotos - reply: %s", r)
	return r, err
}

// UserGetProfilePhotos
// user.getProfilePhotos user_id:long = Vector<long>;
func (s *Service) UserGetProfilePhotos(ctx context.Context, request *user.TLUserGetProfilePhotos) (*user.VectorLong, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getProfilePhotos - request: %s", request)

	r, err := c.UserGetProfilePhotos(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getProfilePhotos - reply: %s", r)
	return r, err
}

// UserSetBotCommands
// user.setBotCommands user_id:long bot_id:long commands:Vector<BotCommand> = Bool;
func (s *Service) UserSetBotCommands(ctx context.Context, request *user.TLUserSetBotCommands) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.setBotCommands - request: %s", request)

	r, err := c.UserSetBotCommands(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.setBotCommands - reply: %s", r)
	return r, err
}

// UserIsBot
// user.isBot id:long = Bool;
func (s *Service) UserIsBot(ctx context.Context, request *user.TLUserIsBot) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.isBot - request: %s", request)

	r, err := c.UserIsBot(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.isBot - reply: %s", r)
	return r, err
}

// UserGetBotInfo
// user.getBotInfo bot_id:long = BotInfo;
func (s *Service) UserGetBotInfo(ctx context.Context, request *user.TLUserGetBotInfo) (*tg.BotInfo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getBotInfo - request: %s", request)

	r, err := c.UserGetBotInfo(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getBotInfo - reply: %s", r)
	return r, err
}

// UserCheckBots
// user.checkBots id:Vector<long> = Vector<long>;
func (s *Service) UserCheckBots(ctx context.Context, request *user.TLUserCheckBots) (*user.VectorLong, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.checkBots - request: %s", request)

	r, err := c.UserCheckBots(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.checkBots - reply: %s", r)
	return r, err
}

// UserGetFullUser
// user.getFullUser self_user_id:long id:long = users.UserFull;
func (s *Service) UserGetFullUser(ctx context.Context, request *user.TLUserGetFullUser) (*tg.UsersUserFull, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getFullUser - request: %s", request)

	r, err := c.UserGetFullUser(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getFullUser - reply: %s", r)
	return r, err
}

// UserUpdateEmojiStatus
// user.updateEmojiStatus user_id:long emoji_status_document_id:long emoji_status_until:int = Bool;
func (s *Service) UserUpdateEmojiStatus(ctx context.Context, request *user.TLUserUpdateEmojiStatus) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.updateEmojiStatus - request: %s", request)

	r, err := c.UserUpdateEmojiStatus(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.updateEmojiStatus - reply: %s", r)
	return r, err
}

// UserGetUserDataById
// user.getUserDataById user_id:long = UserData;
func (s *Service) UserGetUserDataById(ctx context.Context, request *user.TLUserGetUserDataById) (*tg.UserData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getUserDataById - request: %s", request)

	r, err := c.UserGetUserDataById(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getUserDataById - reply: %s", r)
	return r, err
}

// UserGetUserDataListByIdList
// user.getUserDataListByIdList user_id_list:Vector<long> = Vector<UserData>;
func (s *Service) UserGetUserDataListByIdList(ctx context.Context, request *user.TLUserGetUserDataListByIdList) (*user.VectorUserData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getUserDataListByIdList - request: %s", request)

	r, err := c.UserGetUserDataListByIdList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getUserDataListByIdList - reply: %s", r)
	return r, err
}

// UserGetUserDataByToken
// user.getUserDataByToken token:string = UserData;
func (s *Service) UserGetUserDataByToken(ctx context.Context, request *user.TLUserGetUserDataByToken) (*tg.UserData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getUserDataByToken - request: %s", request)

	r, err := c.UserGetUserDataByToken(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getUserDataByToken - reply: %s", r)
	return r, err
}

// UserSearch
// user.search q:string excluded_contacts:Vector<long> offset:long limit:int = UsersFound;
func (s *Service) UserSearch(ctx context.Context, request *user.TLUserSearch) (*user.UsersFound, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.search - request: %s", request)

	r, err := c.UserSearch(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.search - reply: %s", r)
	return r, err
}

// UserUpdateBotData
// user.updateBotData flags:# bot_id:long bot_chat_history:flags.15?Bool bot_nochats:flags.16?Bool bot_inline_geo:flags.21?Bool bot_attach_menu:flags.27?Bool bot_inline_placeholder:flags.19?string bot_has_main_app:flags.13?Bool = Bool;
func (s *Service) UserUpdateBotData(ctx context.Context, request *user.TLUserUpdateBotData) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.updateBotData - request: %s", request)

	r, err := c.UserUpdateBotData(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.updateBotData - reply: %s", r)
	return r, err
}

// UserGetImmutableUserV2
// user.getImmutableUserV2 flags:# id:long privacy:flags.0?true has_to:flags.2?true to:flags.2?Vector<long> = ImmutableUser;
func (s *Service) UserGetImmutableUserV2(ctx context.Context, request *user.TLUserGetImmutableUserV2) (*tg.ImmutableUser, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getImmutableUserV2 - request: %s", request)

	r, err := c.UserGetImmutableUserV2(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getImmutableUserV2 - reply: %s", r)
	return r, err
}

// UserGetMutableUsersV2
// user.getMutableUsersV2 flags:# id:Vector<long> privacy:flags.0?true has_to:flags.2?true to:flags.2?Vector<long> = MutableUsers;
func (s *Service) UserGetMutableUsersV2(ctx context.Context, request *user.TLUserGetMutableUsersV2) (*tg.MutableUsers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getMutableUsersV2 - request: %s", request)

	r, err := c.UserGetMutableUsersV2(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getMutableUsersV2 - reply: %s", r)
	return r, err
}

// UserCreateNewTestUser
// user.createNewTestUser secret_key_id:long min_id:long max_id:long = ImmutableUser;
func (s *Service) UserCreateNewTestUser(ctx context.Context, request *user.TLUserCreateNewTestUser) (*tg.ImmutableUser, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.createNewTestUser - request: %s", request)

	r, err := c.UserCreateNewTestUser(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.createNewTestUser - reply: %s", r)
	return r, err
}

// UserEditCloseFriends
// user.editCloseFriends user_id:long id:Vector<long> = Bool;
func (s *Service) UserEditCloseFriends(ctx context.Context, request *user.TLUserEditCloseFriends) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.editCloseFriends - request: %s", request)

	r, err := c.UserEditCloseFriends(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.editCloseFriends - reply: %s", r)
	return r, err
}

// UserSetStoriesMaxId
// user.setStoriesMaxId user_id:long id:int = Bool;
func (s *Service) UserSetStoriesMaxId(ctx context.Context, request *user.TLUserSetStoriesMaxId) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.setStoriesMaxId - request: %s", request)

	r, err := c.UserSetStoriesMaxId(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.setStoriesMaxId - reply: %s", r)
	return r, err
}

// UserSetColor
// user.setColor flags:# user_id:long for_profile:flags.1?true color:int background_emoji_id:long = Bool;
func (s *Service) UserSetColor(ctx context.Context, request *user.TLUserSetColor) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.setColor - request: %s", request)

	r, err := c.UserSetColor(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.setColor - reply: %s", r)
	return r, err
}

// UserUpdateBirthday
// user.updateBirthday flags:# user_id:long birthday:flags.1?Birthday = Bool;
func (s *Service) UserUpdateBirthday(ctx context.Context, request *user.TLUserUpdateBirthday) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.updateBirthday - request: %s", request)

	r, err := c.UserUpdateBirthday(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.updateBirthday - reply: %s", r)
	return r, err
}

// UserGetBirthdays
// user.getBirthdays user_id:long = Vector<ContactBirthday>;
func (s *Service) UserGetBirthdays(ctx context.Context, request *user.TLUserGetBirthdays) (*user.VectorContactBirthday, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getBirthdays - request: %s", request)

	r, err := c.UserGetBirthdays(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getBirthdays - reply: %s", r)
	return r, err
}

// UserSetStoriesHidden
// user.setStoriesHidden user_id:long id:long hidden:Bool = Bool;
func (s *Service) UserSetStoriesHidden(ctx context.Context, request *user.TLUserSetStoriesHidden) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.setStoriesHidden - request: %s", request)

	r, err := c.UserSetStoriesHidden(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.setStoriesHidden - reply: %s", r)
	return r, err
}

// UserUpdatePersonalChannel
// user.updatePersonalChannel user_id:long channel_id:long = Bool;
func (s *Service) UserUpdatePersonalChannel(ctx context.Context, request *user.TLUserUpdatePersonalChannel) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.updatePersonalChannel - request: %s", request)

	r, err := c.UserUpdatePersonalChannel(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.updatePersonalChannel - reply: %s", r)
	return r, err
}

// UserGetUserIdByPhone
// user.getUserIdByPhone phone:string = Int64;
func (s *Service) UserGetUserIdByPhone(ctx context.Context, request *user.TLUserGetUserIdByPhone) (*tg.Int64, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getUserIdByPhone - request: %s", request)

	r, err := c.UserGetUserIdByPhone(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getUserIdByPhone - reply: %s", r)
	return r, err
}

// UserSetAuthorizationTTL
// user.setAuthorizationTTL user_id:long ttl:int = Bool;
func (s *Service) UserSetAuthorizationTTL(ctx context.Context, request *user.TLUserSetAuthorizationTTL) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.setAuthorizationTTL - request: %s", request)

	r, err := c.UserSetAuthorizationTTL(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.setAuthorizationTTL - reply: %s", r)
	return r, err
}

// UserGetAuthorizationTTL
// user.getAuthorizationTTL user_id:long = AccountDaysTTL;
func (s *Service) UserGetAuthorizationTTL(ctx context.Context, request *user.TLUserGetAuthorizationTTL) (*tg.AccountDaysTTL, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getAuthorizationTTL - request: %s", request)

	r, err := c.UserGetAuthorizationTTL(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getAuthorizationTTL - reply: %s", r)
	return r, err
}

// UserUpdatePremium
// user.updatePremium flags:# user_id:long premium:Bool months:flags.1?int = Bool;
func (s *Service) UserUpdatePremium(ctx context.Context, request *user.TLUserUpdatePremium) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.updatePremium - request: %s", request)

	r, err := c.UserUpdatePremium(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.updatePremium - reply: %s", r)
	return r, err
}

// UserGetBotInfoV2
// user.getBotInfoV2 bot_id:long = BotInfoData;
func (s *Service) UserGetBotInfoV2(ctx context.Context, request *user.TLUserGetBotInfoV2) (*user.BotInfoData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getBotInfoV2 - request: %s", request)

	r, err := c.UserGetBotInfoV2(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getBotInfoV2 - reply: %s", r)
	return r, err
}

// UserSaveMusic
// user.saveMusic flags:# unsave:flags.0?true user_id:long id:long after_id:flags.15?long = Bool;
func (s *Service) UserSaveMusic(ctx context.Context, request *user.TLUserSaveMusic) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.saveMusic - request: %s", request)

	r, err := c.UserSaveMusic(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.saveMusic - reply: %s", r)
	return r, err
}

// UserGetSavedMusicIdList
// user.getSavedMusicIdList user_id:long = Vector<long>;
func (s *Service) UserGetSavedMusicIdList(ctx context.Context, request *user.TLUserGetSavedMusicIdList) (*user.VectorLong, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getSavedMusicIdList - request: %s", request)

	r, err := c.UserGetSavedMusicIdList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getSavedMusicIdList - reply: %s", r)
	return r, err
}

// UserSetMainProfileTab
// user.setMainProfileTab user_id:long tab:ProfileTab = Bool;
func (s *Service) UserSetMainProfileTab(ctx context.Context, request *user.TLUserSetMainProfileTab) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.setMainProfileTab - request: %s", request)

	r, err := c.UserSetMainProfileTab(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.setMainProfileTab - reply: %s", r)
	return r, err
}

// UserSetDefaultHistoryTTL
// user.setDefaultHistoryTTL user_id:long ttl:int = Bool;
func (s *Service) UserSetDefaultHistoryTTL(ctx context.Context, request *user.TLUserSetDefaultHistoryTTL) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.setDefaultHistoryTTL - request: %s", request)

	r, err := c.UserSetDefaultHistoryTTL(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.setDefaultHistoryTTL - reply: %s", r)
	return r, err
}

// UserGetDefaultHistoryTTL
// user.getDefaultHistoryTTL user_id:long = DefaultHistoryTTL;
func (s *Service) UserGetDefaultHistoryTTL(ctx context.Context, request *user.TLUserGetDefaultHistoryTTL) (*tg.DefaultHistoryTTL, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getDefaultHistoryTTL - request: %s", request)

	r, err := c.UserGetDefaultHistoryTTL(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getDefaultHistoryTTL - reply: %s", r)
	return r, err
}

// UserGetAccountUsername
// user.getAccountUsername user_id:long = UsernameData;
func (s *Service) UserGetAccountUsername(ctx context.Context, request *user.TLUserGetAccountUsername) (*user.UsernameData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getAccountUsername - request: %s", request)

	r, err := c.UserGetAccountUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getAccountUsername - reply: %s", r)
	return r, err
}

// UserCheckAccountUsername
// user.checkAccountUsername user_id:long username:string = UsernameExist;
func (s *Service) UserCheckAccountUsername(ctx context.Context, request *user.TLUserCheckAccountUsername) (*user.UsernameExist, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.checkAccountUsername - request: %s", request)

	r, err := c.UserCheckAccountUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.checkAccountUsername - reply: %s", r)
	return r, err
}

// UserGetChannelUsername
// user.getChannelUsername channel_id:long = UsernameData;
func (s *Service) UserGetChannelUsername(ctx context.Context, request *user.TLUserGetChannelUsername) (*user.UsernameData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getChannelUsername - request: %s", request)

	r, err := c.UserGetChannelUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getChannelUsername - reply: %s", r)
	return r, err
}

// UserCheckChannelUsername
// user.checkChannelUsername channel_id:long username:string = UsernameExist;
func (s *Service) UserCheckChannelUsername(ctx context.Context, request *user.TLUserCheckChannelUsername) (*user.UsernameExist, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.checkChannelUsername - request: %s", request)

	r, err := c.UserCheckChannelUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.checkChannelUsername - reply: %s", r)
	return r, err
}

// UserUpdateUsernameByPeer
// user.updateUsernameByPeer peer_type:int peer_id:long username:string = Bool;
func (s *Service) UserUpdateUsernameByPeer(ctx context.Context, request *user.TLUserUpdateUsernameByPeer) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.updateUsernameByPeer - request: %s", request)

	r, err := c.UserUpdateUsernameByPeer(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.updateUsernameByPeer - reply: %s", r)
	return r, err
}

// UserCheckUsername
// user.checkUsername username:string = UsernameExist;
func (s *Service) UserCheckUsername(ctx context.Context, request *user.TLUserCheckUsername) (*user.UsernameExist, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.checkUsername - request: %s", request)

	r, err := c.UserCheckUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.checkUsername - reply: %s", r)
	return r, err
}

// UserUpdateUsernameByUsername
// user.updateUsernameByUsername peer_type:int peer_id:long username:string = Bool;
func (s *Service) UserUpdateUsernameByUsername(ctx context.Context, request *user.TLUserUpdateUsernameByUsername) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.updateUsernameByUsername - request: %s", request)

	r, err := c.UserUpdateUsernameByUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.updateUsernameByUsername - reply: %s", r)
	return r, err
}

// UserDeleteUsername
// user.deleteUsername username:string = Bool;
func (s *Service) UserDeleteUsername(ctx context.Context, request *user.TLUserDeleteUsername) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.deleteUsername - request: %s", request)

	r, err := c.UserDeleteUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.deleteUsername - reply: %s", r)
	return r, err
}

// UserResolveUsername
// user.resolveUsername username:string = Peer;
func (s *Service) UserResolveUsername(ctx context.Context, request *user.TLUserResolveUsername) (*tg.Peer, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.resolveUsername - request: %s", request)

	r, err := c.UserResolveUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.resolveUsername - reply: %s", r)
	return r, err
}

// UserGetListByUsernameList
// user.getListByUsernameList names:Vector<string> = Vector<UsernameData>;
func (s *Service) UserGetListByUsernameList(ctx context.Context, request *user.TLUserGetListByUsernameList) (*user.VectorUsernameData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getListByUsernameList - request: %s", request)

	r, err := c.UserGetListByUsernameList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getListByUsernameList - reply: %s", r)
	return r, err
}

// UserDeleteUsernameByPeer
// user.deleteUsernameByPeer peer_type:int peer_id:long = Bool;
func (s *Service) UserDeleteUsernameByPeer(ctx context.Context, request *user.TLUserDeleteUsernameByPeer) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.deleteUsernameByPeer - request: %s", request)

	r, err := c.UserDeleteUsernameByPeer(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.deleteUsernameByPeer - reply: %s", r)
	return r, err
}

// UserSearchUsername
// user.searchUsername q:string excluded_contacts:Vector<long> limit:int = Vector<UsernameData>;
func (s *Service) UserSearchUsername(ctx context.Context, request *user.TLUserSearchUsername) (*user.VectorUsernameData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.searchUsername - request: %s", request)

	r, err := c.UserSearchUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.searchUsername - reply: %s", r)
	return r, err
}

// UserToggleUsername
// user.toggleUsername peer_type:int peer_id:long username:string active:Bool = Bool;
func (s *Service) UserToggleUsername(ctx context.Context, request *user.TLUserToggleUsername) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.toggleUsername - request: %s", request)

	r, err := c.UserToggleUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.toggleUsername - reply: %s", r)
	return r, err
}

// UserReorderUsernames
// user.reorderUsernames peer_type:int peer_id:long username_list:Vector<string> = Bool;
func (s *Service) UserReorderUsernames(ctx context.Context, request *user.TLUserReorderUsernames) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.reorderUsernames - request: %s", request)

	r, err := c.UserReorderUsernames(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.reorderUsernames - reply: %s", r)
	return r, err
}

// UserDeactivateAllChannelUsernames
// user.deactivateAllChannelUsernames channel_id:long = Bool;
func (s *Service) UserDeactivateAllChannelUsernames(ctx context.Context, request *user.TLUserDeactivateAllChannelUsernames) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.deactivateAllChannelUsernames - request: %s", request)

	r, err := c.UserDeactivateAllChannelUsernames(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.deactivateAllChannelUsernames - reply: %s", r)
	return r, err
}
