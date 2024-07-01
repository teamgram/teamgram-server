/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/core"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserGetLastSeens
// user.getLastSeens id:Vector<long> = Vector<LastSeenData>;
func (s *Service) UserGetLastSeens(ctx context.Context, request *user.TLUserGetLastSeens) (*user.Vector_LastSeenData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getLastSeens - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserGetLastSeens(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getLastSeens - reply: %s", r)
	return r, err
}

// UserUpdateLastSeen
// user.updateLastSeen id:long last_seen_at:long expires:int = Bool;
func (s *Service) UserUpdateLastSeen(ctx context.Context, request *user.TLUserUpdateLastSeen) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.updateLastSeen - metadata: %s, request: %s", c.MD, request)

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
	c.Logger.Debugf("user.getLastSeen - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserGetLastSeen(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getLastSeen - reply: %s", r)
	return r, err
}

// UserGetImmutableUser
// user.getImmutableUser flags:# id:long privacy:flags.1?true contacts:Vector<long> = ImmutableUser;
func (s *Service) UserGetImmutableUser(ctx context.Context, request *user.TLUserGetImmutableUser) (*mtproto.ImmutableUser, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getImmutableUser - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserGetImmutableUser(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getImmutableUser - reply: %s", r)
	return r, err
}

// UserGetMutableUsers
// user.getMutableUsers id:Vector<long> to:Vector<long> = Vector<ImmutableUser>;
func (s *Service) UserGetMutableUsers(ctx context.Context, request *user.TLUserGetMutableUsers) (*user.Vector_ImmutableUser, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getMutableUsers - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserGetMutableUsers(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getMutableUsers - reply: %s", r)
	return r, err
}

// UserGetImmutableUserByPhone
// user.getImmutableUserByPhone phone:string = ImmutableUser;
func (s *Service) UserGetImmutableUserByPhone(ctx context.Context, request *user.TLUserGetImmutableUserByPhone) (*mtproto.ImmutableUser, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getImmutableUserByPhone - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserGetImmutableUserByPhone(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getImmutableUserByPhone - reply: %s", r)
	return r, err
}

// UserGetImmutableUserByToken
// user.getImmutableUserByToken token:string = ImmutableUser;
func (s *Service) UserGetImmutableUserByToken(ctx context.Context, request *user.TLUserGetImmutableUserByToken) (*mtproto.ImmutableUser, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getImmutableUserByToken - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserGetImmutableUserByToken(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getImmutableUserByToken - reply: %s", r)
	return r, err
}

// UserSetAccountDaysTTL
// user.setAccountDaysTTL user_id:long ttl:int = Bool;
func (s *Service) UserSetAccountDaysTTL(ctx context.Context, request *user.TLUserSetAccountDaysTTL) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.setAccountDaysTTL - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserSetAccountDaysTTL(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.setAccountDaysTTL - reply: %s", r)
	return r, err
}

// UserGetAccountDaysTTL
// user.getAccountDaysTTL user_id:long = AccountDaysTTL;
func (s *Service) UserGetAccountDaysTTL(ctx context.Context, request *user.TLUserGetAccountDaysTTL) (*mtproto.AccountDaysTTL, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getAccountDaysTTL - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserGetAccountDaysTTL(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getAccountDaysTTL - reply: %s", r)
	return r, err
}

// UserGetNotifySettings
// user.getNotifySettings user_id:long peer_type:int peer_id:long = PeerNotifySettings;
func (s *Service) UserGetNotifySettings(ctx context.Context, request *user.TLUserGetNotifySettings) (*mtproto.PeerNotifySettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getNotifySettings - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserGetNotifySettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getNotifySettings - reply: %s", r)
	return r, err
}

// UserGetNotifySettingsList
// user.getNotifySettingsList user_id:long peers:Vector<PeerUtil> = Vector<PeerPeerNotifySettings>;
func (s *Service) UserGetNotifySettingsList(ctx context.Context, request *user.TLUserGetNotifySettingsList) (*user.Vector_PeerPeerNotifySettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getNotifySettingsList - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserGetNotifySettingsList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getNotifySettingsList - reply: %s", r)
	return r, err
}

// UserSetNotifySettings
// user.setNotifySettings user_id:long peer_type:int peer_id:long settings:PeerNotifySettings = Bool;
func (s *Service) UserSetNotifySettings(ctx context.Context, request *user.TLUserSetNotifySettings) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.setNotifySettings - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserSetNotifySettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.setNotifySettings - reply: %s", r)
	return r, err
}

// UserResetNotifySettings
// user.resetNotifySettings user_id:long = Bool;
func (s *Service) UserResetNotifySettings(ctx context.Context, request *user.TLUserResetNotifySettings) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.resetNotifySettings - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserResetNotifySettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.resetNotifySettings - reply: %s", r)
	return r, err
}

// UserGetAllNotifySettings
// user.getAllNotifySettings user_id:long = Vector<PeerPeerNotifySettings>;
func (s *Service) UserGetAllNotifySettings(ctx context.Context, request *user.TLUserGetAllNotifySettings) (*user.Vector_PeerPeerNotifySettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getAllNotifySettings - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserGetAllNotifySettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getAllNotifySettings - reply: %s", r)
	return r, err
}

// UserGetGlobalPrivacySettings
// user.getGlobalPrivacySettings user_id:long = GlobalPrivacySettings;
func (s *Service) UserGetGlobalPrivacySettings(ctx context.Context, request *user.TLUserGetGlobalPrivacySettings) (*mtproto.GlobalPrivacySettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getGlobalPrivacySettings - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserGetGlobalPrivacySettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getGlobalPrivacySettings - reply: %s", r)
	return r, err
}

// UserSetGlobalPrivacySettings
// user.setGlobalPrivacySettings user_id:long settings:GlobalPrivacySettings = Bool;
func (s *Service) UserSetGlobalPrivacySettings(ctx context.Context, request *user.TLUserSetGlobalPrivacySettings) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.setGlobalPrivacySettings - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserSetGlobalPrivacySettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.setGlobalPrivacySettings - reply: %s", r)
	return r, err
}

// UserGetPrivacy
// user.getPrivacy user_id:long key_type:int = Vector<PrivacyRule>;
func (s *Service) UserGetPrivacy(ctx context.Context, request *user.TLUserGetPrivacy) (*user.Vector_PrivacyRule, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getPrivacy - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserGetPrivacy(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getPrivacy - reply: %s", r)
	return r, err
}

// UserSetPrivacy
// user.setPrivacy user_id:long key_type:int rules:Vector<PrivacyRule> = Bool;
func (s *Service) UserSetPrivacy(ctx context.Context, request *user.TLUserSetPrivacy) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.setPrivacy - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserSetPrivacy(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.setPrivacy - reply: %s", r)
	return r, err
}

// UserCheckPrivacy
// user.checkPrivacy flags:# user_id:long key_type:int peer_id:long = Bool;
func (s *Service) UserCheckPrivacy(ctx context.Context, request *user.TLUserCheckPrivacy) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.checkPrivacy - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserCheckPrivacy(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.checkPrivacy - reply: %s", r)
	return r, err
}

// UserAddPeerSettings
// user.addPeerSettings user_id:long peer_type:int peer_id:long settings:PeerSettings = Bool;
func (s *Service) UserAddPeerSettings(ctx context.Context, request *user.TLUserAddPeerSettings) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.addPeerSettings - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserAddPeerSettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.addPeerSettings - reply: %s", r)
	return r, err
}

// UserGetPeerSettings
// user.getPeerSettings user_id:long peer_type:int peer_id:long = PeerSettings;
func (s *Service) UserGetPeerSettings(ctx context.Context, request *user.TLUserGetPeerSettings) (*mtproto.PeerSettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getPeerSettings - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserGetPeerSettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getPeerSettings - reply: %s", r)
	return r, err
}

// UserDeletePeerSettings
// user.deletePeerSettings user_id:long peer_type:int peer_id:long = Bool;
func (s *Service) UserDeletePeerSettings(ctx context.Context, request *user.TLUserDeletePeerSettings) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.deletePeerSettings - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserDeletePeerSettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.deletePeerSettings - reply: %s", r)
	return r, err
}

// UserChangePhone
// user.changePhone user_id:long phone:string = Bool;
func (s *Service) UserChangePhone(ctx context.Context, request *user.TLUserChangePhone) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.changePhone - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserChangePhone(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.changePhone - reply: %s", r)
	return r, err
}

// UserCreateNewUser
// user.createNewUser secret_key_id:long phone:string country_code:string first_name:string last_name:string = ImmutableUser;
func (s *Service) UserCreateNewUser(ctx context.Context, request *user.TLUserCreateNewUser) (*mtproto.ImmutableUser, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.createNewUser - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserCreateNewUser(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.createNewUser - reply: %s", r)
	return r, err
}

// UserDeleteUser
// user.deleteUser user_id:long reason:string = Bool;
func (s *Service) UserDeleteUser(ctx context.Context, request *user.TLUserDeleteUser) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.deleteUser - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserDeleteUser(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.deleteUser - reply: %s", r)
	return r, err
}

// UserBlockPeer
// user.blockPeer user_id:long peer_type:int peer_id:long = Bool;
func (s *Service) UserBlockPeer(ctx context.Context, request *user.TLUserBlockPeer) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.blockPeer - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserBlockPeer(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.blockPeer - reply: %s", r)
	return r, err
}

// UserUnBlockPeer
// user.unBlockPeer user_id:long peer_type:int peer_id:long = Bool;
func (s *Service) UserUnBlockPeer(ctx context.Context, request *user.TLUserUnBlockPeer) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.unBlockPeer - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserUnBlockPeer(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.unBlockPeer - reply: %s", r)
	return r, err
}

// UserBlockedByUser
// user.blockedByUser user_id:long peer_user_id:long = Bool;
func (s *Service) UserBlockedByUser(ctx context.Context, request *user.TLUserBlockedByUser) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.blockedByUser - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserBlockedByUser(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.blockedByUser - reply: %s", r)
	return r, err
}

// UserIsBlockedByUser
// user.isBlockedByUser user_id:long peer_user_id:long = Bool;
func (s *Service) UserIsBlockedByUser(ctx context.Context, request *user.TLUserIsBlockedByUser) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.isBlockedByUser - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserIsBlockedByUser(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.isBlockedByUser - reply: %s", r)
	return r, err
}

// UserCheckBlockUserList
// user.checkBlockUserList user_id:long id:Vector<long> = Vector<long>;
func (s *Service) UserCheckBlockUserList(ctx context.Context, request *user.TLUserCheckBlockUserList) (*user.Vector_Long, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.checkBlockUserList - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserCheckBlockUserList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.checkBlockUserList - reply: %s", r)
	return r, err
}

// UserGetBlockedList
// user.getBlockedList user_id:long offset:int limit:int = Vector<PeerBlocked>;
func (s *Service) UserGetBlockedList(ctx context.Context, request *user.TLUserGetBlockedList) (*user.Vector_PeerBlocked, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getBlockedList - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserGetBlockedList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getBlockedList - reply: %s", r)
	return r, err
}

// UserGetContactSignUpNotification
// user.getContactSignUpNotification user_id:long = Bool;
func (s *Service) UserGetContactSignUpNotification(ctx context.Context, request *user.TLUserGetContactSignUpNotification) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getContactSignUpNotification - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserGetContactSignUpNotification(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getContactSignUpNotification - reply: %s", r)
	return r, err
}

// UserSetContactSignUpNotification
// user.setContactSignUpNotification user_id:long silent:Bool = Bool;
func (s *Service) UserSetContactSignUpNotification(ctx context.Context, request *user.TLUserSetContactSignUpNotification) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.setContactSignUpNotification - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserSetContactSignUpNotification(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.setContactSignUpNotification - reply: %s", r)
	return r, err
}

// UserGetContentSettings
// user.getContentSettings user_id:long = account.ContentSettings;
func (s *Service) UserGetContentSettings(ctx context.Context, request *user.TLUserGetContentSettings) (*mtproto.Account_ContentSettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getContentSettings - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserGetContentSettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getContentSettings - reply: %s", r)
	return r, err
}

// UserSetContentSettings
// user.setContentSettings flags:# user_id:long sensitive_enabled:flags.0?true = Bool;
func (s *Service) UserSetContentSettings(ctx context.Context, request *user.TLUserSetContentSettings) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.setContentSettings - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserSetContentSettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.setContentSettings - reply: %s", r)
	return r, err
}

// UserDeleteContact
// user.deleteContact user_id:long id:long = Bool;
func (s *Service) UserDeleteContact(ctx context.Context, request *user.TLUserDeleteContact) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.deleteContact - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserDeleteContact(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.deleteContact - reply: %s", r)
	return r, err
}

// UserGetContactList
// user.getContactList user_id:long = Vector<ContactData>;
func (s *Service) UserGetContactList(ctx context.Context, request *user.TLUserGetContactList) (*user.Vector_ContactData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getContactList - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserGetContactList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getContactList - reply: %s", r)
	return r, err
}

// UserGetContactIdList
// user.getContactIdList user_id:long = Vector<long>;
func (s *Service) UserGetContactIdList(ctx context.Context, request *user.TLUserGetContactIdList) (*user.Vector_Long, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getContactIdList - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserGetContactIdList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getContactIdList - reply: %s", r)
	return r, err
}

// UserGetContact
// user.getContact user_id:long id:long = ContactData;
func (s *Service) UserGetContact(ctx context.Context, request *user.TLUserGetContact) (*mtproto.ContactData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getContact - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserGetContact(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getContact - reply: %s", r)
	return r, err
}

// UserAddContact
// user.addContact user_id:long add_phone_privacy_exception:Bool id:long first_name:string last_name:string phone:string = Bool;
func (s *Service) UserAddContact(ctx context.Context, request *user.TLUserAddContact) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.addContact - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserAddContact(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.addContact - reply: %s", r)
	return r, err
}

// UserCheckContact
// user.checkContact user_id:long id:long = Bool;
func (s *Service) UserCheckContact(ctx context.Context, request *user.TLUserCheckContact) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.checkContact - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserCheckContact(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.checkContact - reply: %s", r)
	return r, err
}

// UserGetImportersByPhone
// user.getImportersByPhone phone:string = Vector<InputContact>;
func (s *Service) UserGetImportersByPhone(ctx context.Context, request *user.TLUserGetImportersByPhone) (*user.Vector_InputContact, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getImportersByPhone - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserGetImportersByPhone(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getImportersByPhone - reply: %s", r)
	return r, err
}

// UserDeleteImportersByPhone
// user.deleteImportersByPhone phone:string = Bool;
func (s *Service) UserDeleteImportersByPhone(ctx context.Context, request *user.TLUserDeleteImportersByPhone) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.deleteImportersByPhone - metadata: %s, request: %s", c.MD, request)

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
	c.Logger.Debugf("user.importContacts - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserImportContacts(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.importContacts - reply: %s", r)
	return r, err
}

// UserGetCountryCode
// user.getCountryCode user_id:long = String;
func (s *Service) UserGetCountryCode(ctx context.Context, request *user.TLUserGetCountryCode) (*mtproto.String, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getCountryCode - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserGetCountryCode(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getCountryCode - reply: %s", r)
	return r, err
}

// UserUpdateAbout
// user.updateAbout user_id:long about:string = Bool;
func (s *Service) UserUpdateAbout(ctx context.Context, request *user.TLUserUpdateAbout) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.updateAbout - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserUpdateAbout(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.updateAbout - reply: %s", r)
	return r, err
}

// UserUpdateFirstAndLastName
// user.updateFirstAndLastName user_id:long first_name:string last_name:string = Bool;
func (s *Service) UserUpdateFirstAndLastName(ctx context.Context, request *user.TLUserUpdateFirstAndLastName) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.updateFirstAndLastName - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserUpdateFirstAndLastName(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.updateFirstAndLastName - reply: %s", r)
	return r, err
}

// UserUpdateVerified
// user.updateVerified user_id:long verified:Bool = Bool;
func (s *Service) UserUpdateVerified(ctx context.Context, request *user.TLUserUpdateVerified) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.updateVerified - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserUpdateVerified(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.updateVerified - reply: %s", r)
	return r, err
}

// UserUpdateUsername
// user.updateUsername user_id:long username:string = Bool;
func (s *Service) UserUpdateUsername(ctx context.Context, request *user.TLUserUpdateUsername) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.updateUsername - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserUpdateUsername(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.updateUsername - reply: %s", r)
	return r, err
}

// UserUpdateProfilePhoto
// user.updateProfilePhoto user_id:long id:long = Int64;
func (s *Service) UserUpdateProfilePhoto(ctx context.Context, request *user.TLUserUpdateProfilePhoto) (*mtproto.Int64, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.updateProfilePhoto - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserUpdateProfilePhoto(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.updateProfilePhoto - reply: %s", r)
	return r, err
}

// UserDeleteProfilePhotos
// user.deleteProfilePhotos user_id:long id:Vector<long> = Int64;
func (s *Service) UserDeleteProfilePhotos(ctx context.Context, request *user.TLUserDeleteProfilePhotos) (*mtproto.Int64, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.deleteProfilePhotos - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserDeleteProfilePhotos(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.deleteProfilePhotos - reply: %s", r)
	return r, err
}

// UserGetProfilePhotos
// user.getProfilePhotos user_id:long = Vector<long>;
func (s *Service) UserGetProfilePhotos(ctx context.Context, request *user.TLUserGetProfilePhotos) (*user.Vector_Long, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getProfilePhotos - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserGetProfilePhotos(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getProfilePhotos - reply: %s", r)
	return r, err
}

// UserSetBotCommands
// user.setBotCommands user_id:long bot_id:long commands:Vector<BotCommand> = Bool;
func (s *Service) UserSetBotCommands(ctx context.Context, request *user.TLUserSetBotCommands) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.setBotCommands - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserSetBotCommands(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.setBotCommands - reply: %s", r)
	return r, err
}

// UserIsBot
// user.isBot id:long = Bool;
func (s *Service) UserIsBot(ctx context.Context, request *user.TLUserIsBot) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.isBot - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserIsBot(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.isBot - reply: %s", r)
	return r, err
}

// UserGetBotInfo
// user.getBotInfo bot_id:long = BotInfo;
func (s *Service) UserGetBotInfo(ctx context.Context, request *user.TLUserGetBotInfo) (*mtproto.BotInfo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getBotInfo - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserGetBotInfo(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getBotInfo - reply: %s", r)
	return r, err
}

// UserCheckBots
// user.checkBots id:Vector<long> = Vector<long>;
func (s *Service) UserCheckBots(ctx context.Context, request *user.TLUserCheckBots) (*user.Vector_Long, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.checkBots - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserCheckBots(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.checkBots - reply: %s", r)
	return r, err
}

// UserGetFullUser
// user.getFullUser self_user_id:long id:long = users.UserFull;
func (s *Service) UserGetFullUser(ctx context.Context, request *user.TLUserGetFullUser) (*mtproto.Users_UserFull, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getFullUser - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserGetFullUser(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getFullUser - reply: %s", r)
	return r, err
}

// UserUpdateEmojiStatus
// user.updateEmojiStatus user_id:long emoji_status_document_id:long emoji_status_until:int = Bool;
func (s *Service) UserUpdateEmojiStatus(ctx context.Context, request *user.TLUserUpdateEmojiStatus) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.updateEmojiStatus - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserUpdateEmojiStatus(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.updateEmojiStatus - reply: %s", r)
	return r, err
}

// UserGetUserDataById
// user.getUserDataById user_id:long = UserData;
func (s *Service) UserGetUserDataById(ctx context.Context, request *user.TLUserGetUserDataById) (*mtproto.UserData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getUserDataById - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserGetUserDataById(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getUserDataById - reply: %s", r)
	return r, err
}

// UserGetUserDataListByIdList
// user.getUserDataListByIdList user_id_list:Vector<long> = Vector<UserData>;
func (s *Service) UserGetUserDataListByIdList(ctx context.Context, request *user.TLUserGetUserDataListByIdList) (*user.Vector_UserData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getUserDataListByIdList - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserGetUserDataListByIdList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getUserDataListByIdList - reply: %s", r)
	return r, err
}

// UserGetUserDataByToken
// user.getUserDataByToken token:string = UserData;
func (s *Service) UserGetUserDataByToken(ctx context.Context, request *user.TLUserGetUserDataByToken) (*mtproto.UserData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getUserDataByToken - metadata: %s, request: %s", c.MD, request)

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
	c.Logger.Debugf("user.search - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserSearch(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.search - reply: %s", r)
	return r, err
}

// UserUpdateBotData
// user.updateBotData flags:# bot_id:long bot_chat_history:flags.15?Bool bot_nochats:flags.16?Bool bot_inline_geo:flags.21?Bool bot_attach_menu:flags.27?Bool bot_inline_placeholder:flags.19?Bool = Bool;
func (s *Service) UserUpdateBotData(ctx context.Context, request *user.TLUserUpdateBotData) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.updateBotData - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserUpdateBotData(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.updateBotData - reply: %s", r)
	return r, err
}

// UserGetImmutableUserV2
// user.getImmutableUserV2 flags:# id:long privacy:flags.0?true has_to:flags.2?true to:flags.2?Vector<long> = ImmutableUser;
func (s *Service) UserGetImmutableUserV2(ctx context.Context, request *user.TLUserGetImmutableUserV2) (*mtproto.ImmutableUser, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getImmutableUserV2 - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserGetImmutableUserV2(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getImmutableUserV2 - reply: %s", r)
	return r, err
}

// UserGetMutableUsersV2
// user.getMutableUsersV2 flags:# id:Vector<long> privacy:flags.0?true has_to:flags.2?true to:flags.2?Vector<long> = MutableUsers;
func (s *Service) UserGetMutableUsersV2(ctx context.Context, request *user.TLUserGetMutableUsersV2) (*mtproto.MutableUsers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getMutableUsersV2 - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserGetMutableUsersV2(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getMutableUsersV2 - reply: %s", r)
	return r, err
}

// UserCreateNewTestUser
// user.createNewTestUser secret_key_id:long min_id:long max_id:long = ImmutableUser;
func (s *Service) UserCreateNewTestUser(ctx context.Context, request *user.TLUserCreateNewTestUser) (*mtproto.ImmutableUser, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.createNewTestUser - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserCreateNewTestUser(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.createNewTestUser - reply: %s", r)
	return r, err
}

// UserEditCloseFriends
// user.editCloseFriends user_id:long id:Vector<long> = Bool;
func (s *Service) UserEditCloseFriends(ctx context.Context, request *user.TLUserEditCloseFriends) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.editCloseFriends - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserEditCloseFriends(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.editCloseFriends - reply: %s", r)
	return r, err
}

// UserSetStoriesMaxId
// user.setStoriesMaxId user_id:long id:int = Bool;
func (s *Service) UserSetStoriesMaxId(ctx context.Context, request *user.TLUserSetStoriesMaxId) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.setStoriesMaxId - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserSetStoriesMaxId(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.setStoriesMaxId - reply: %s", r)
	return r, err
}

// UserSetColor
// user.setColor flags:# user_id:long for_profile:flags.1?true color:int background_emoji_id:long = Bool;
func (s *Service) UserSetColor(ctx context.Context, request *user.TLUserSetColor) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.setColor - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserSetColor(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.setColor - reply: %s", r)
	return r, err
}

// UserUpdateBirthday
// user.updateBirthday flags:# user_id:long birthday:flags.1?Birthday = Bool;
func (s *Service) UserUpdateBirthday(ctx context.Context, request *user.TLUserUpdateBirthday) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.updateBirthday - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserUpdateBirthday(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.updateBirthday - reply: %s", r)
	return r, err
}

// UserGetBirthdays
// user.getBirthdays user_id:long = Vector<ContactBirthday>;
func (s *Service) UserGetBirthdays(ctx context.Context, request *user.TLUserGetBirthdays) (*user.Vector_ContactBirthday, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.getBirthdays - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserGetBirthdays(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.getBirthdays - reply: %s", r)
	return r, err
}

// UserSetStoriesHidden
// user.setStoriesHidden user_id:long  id:long hidden:Bool = Bool;
func (s *Service) UserSetStoriesHidden(ctx context.Context, request *user.TLUserSetStoriesHidden) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("user.setStoriesHidden - metadata: %s, request: %s", c.MD, request)

	r, err := c.UserSetStoriesHidden(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("user.setStoriesHidden - reply: %s", r)
	return r, err
}
