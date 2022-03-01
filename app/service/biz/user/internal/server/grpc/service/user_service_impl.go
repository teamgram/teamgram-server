/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
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
	c.Infof("user.getLastSeens - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserGetLastSeens(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.getLastSeens - reply: %s", r.DebugString())
	return r, err
}

// UserUpdateLastSeen
// user.updateLastSeen id:long last_seen_at:long expries:int = Bool;
func (s *Service) UserUpdateLastSeen(ctx context.Context, request *user.TLUserUpdateLastSeen) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.updateLastSeen - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserUpdateLastSeen(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.updateLastSeen - reply: %s", r.DebugString())
	return r, err
}

// UserGetLastSeen
// user.getLastSeen id:long = LastSeenData;
func (s *Service) UserGetLastSeen(ctx context.Context, request *user.TLUserGetLastSeen) (*user.LastSeenData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.getLastSeen - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserGetLastSeen(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.getLastSeen - reply: %s", r.DebugString())
	return r, err
}

// UserGetImmutableUser
// user.getImmutableUser id:long = ImmutableUser;
func (s *Service) UserGetImmutableUser(ctx context.Context, request *user.TLUserGetImmutableUser) (*user.ImmutableUser, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.getImmutableUser - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserGetImmutableUser(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.getImmutableUser - reply: %s", r.DebugString())
	return r, err
}

// UserGetMutableUsers
// user.getMutableUsers id:Vector<long> = Vector<ImmutableUser>;
func (s *Service) UserGetMutableUsers(ctx context.Context, request *user.TLUserGetMutableUsers) (*user.Vector_ImmutableUser, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.getMutableUsers - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserGetMutableUsers(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.getMutableUsers - reply: %s", r.DebugString())
	return r, err
}

// UserGetImmutableUserByPhone
// user.getImmutableUserByPhone phone:string = ImmutableUser;
func (s *Service) UserGetImmutableUserByPhone(ctx context.Context, request *user.TLUserGetImmutableUserByPhone) (*user.ImmutableUser, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.getImmutableUserByPhone - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserGetImmutableUserByPhone(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.getImmutableUserByPhone - reply: %s", r.DebugString())
	return r, err
}

// UserGetImmutableUserByToken
// user.getImmutableUserByToken token:string = ImmutableUser;
func (s *Service) UserGetImmutableUserByToken(ctx context.Context, request *user.TLUserGetImmutableUserByToken) (*user.ImmutableUser, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.getImmutableUserByToken - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserGetImmutableUserByToken(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.getImmutableUserByToken - reply: %s", r.DebugString())
	return r, err
}

// UserSetAccountDaysTTL
// user.setAccountDaysTTL user_id:long ttl:int = Bool;
func (s *Service) UserSetAccountDaysTTL(ctx context.Context, request *user.TLUserSetAccountDaysTTL) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.setAccountDaysTTL - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserSetAccountDaysTTL(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.setAccountDaysTTL - reply: %s", r.DebugString())
	return r, err
}

// UserGetAccountDaysTTL
// user.getAccountDaysTTL user_id:long = AccountDaysTTL;
func (s *Service) UserGetAccountDaysTTL(ctx context.Context, request *user.TLUserGetAccountDaysTTL) (*mtproto.AccountDaysTTL, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.getAccountDaysTTL - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserGetAccountDaysTTL(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.getAccountDaysTTL - reply: %s", r.DebugString())
	return r, err
}

// UserGetNotifySettings
// user.getNotifySettings user_id:long peer_type:int peer_id:long = PeerNotifySettings;
func (s *Service) UserGetNotifySettings(ctx context.Context, request *user.TLUserGetNotifySettings) (*mtproto.PeerNotifySettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.getNotifySettings - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserGetNotifySettings(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.getNotifySettings - reply: %s", r.DebugString())
	return r, err
}

// UserSetNotifySettings
// user.setNotifySettings user_id:long peer_type:int peer_id:long settings:PeerNotifySettings = Bool;
func (s *Service) UserSetNotifySettings(ctx context.Context, request *user.TLUserSetNotifySettings) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.setNotifySettings - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserSetNotifySettings(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.setNotifySettings - reply: %s", r.DebugString())
	return r, err
}

// UserResetNotifySettings
// user.resetNotifySettings user_id:long = Bool;
func (s *Service) UserResetNotifySettings(ctx context.Context, request *user.TLUserResetNotifySettings) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.resetNotifySettings - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserResetNotifySettings(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.resetNotifySettings - reply: %s", r.DebugString())
	return r, err
}

// UserGetAllNotifySettings
// user.getAllNotifySettings user_id:long = Vector<PeerPeerNotifySettings>;
func (s *Service) UserGetAllNotifySettings(ctx context.Context, request *user.TLUserGetAllNotifySettings) (*user.Vector_PeerPeerNotifySettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.getAllNotifySettings - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserGetAllNotifySettings(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.getAllNotifySettings - reply: %s", r.DebugString())
	return r, err
}

// UserGetGlobalPrivacySettings
// user.getGlobalPrivacySettings user_id:long = GlobalPrivacySettings;
func (s *Service) UserGetGlobalPrivacySettings(ctx context.Context, request *user.TLUserGetGlobalPrivacySettings) (*mtproto.GlobalPrivacySettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.getGlobalPrivacySettings - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserGetGlobalPrivacySettings(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.getGlobalPrivacySettings - reply: %s", r.DebugString())
	return r, err
}

// UserSetGlobalPrivacySettings
// user.setGlobalPrivacySettings user_id:long settings:GlobalPrivacySettings = Bool;
func (s *Service) UserSetGlobalPrivacySettings(ctx context.Context, request *user.TLUserSetGlobalPrivacySettings) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.setGlobalPrivacySettings - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserSetGlobalPrivacySettings(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.setGlobalPrivacySettings - reply: %s", r.DebugString())
	return r, err
}

// UserGetPrivacy
// user.getPrivacy user_id:long key_type:int = Vector<PrivacyRule>;
func (s *Service) UserGetPrivacy(ctx context.Context, request *user.TLUserGetPrivacy) (*user.Vector_PrivacyRule, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.getPrivacy - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserGetPrivacy(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.getPrivacy - reply: %s", r.DebugString())
	return r, err
}

// UserSetPrivacy
// user.setPrivacy user_id:long key_type:int rules:Vector<PrivacyRule> = Bool;
func (s *Service) UserSetPrivacy(ctx context.Context, request *user.TLUserSetPrivacy) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.setPrivacy - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserSetPrivacy(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.setPrivacy - reply: %s", r.DebugString())
	return r, err
}

// UserCheckPrivacy
// user.checkPrivacy flags:# user_id:long key_type:int peer_id:long = Bool;
func (s *Service) UserCheckPrivacy(ctx context.Context, request *user.TLUserCheckPrivacy) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.checkPrivacy - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserCheckPrivacy(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.checkPrivacy - reply: %s", r.DebugString())
	return r, err
}

// UserAddPeerSettings
// user.addPeerSettings user_id:long peer_type:int peer_id:long settings:PeerSettings = Bool;
func (s *Service) UserAddPeerSettings(ctx context.Context, request *user.TLUserAddPeerSettings) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.addPeerSettings - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserAddPeerSettings(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.addPeerSettings - reply: %s", r.DebugString())
	return r, err
}

// UserGetPeerSettings
// user.getPeerSettings user_id:long peer_type:int peer_id:long = PeerSettings;
func (s *Service) UserGetPeerSettings(ctx context.Context, request *user.TLUserGetPeerSettings) (*mtproto.PeerSettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.getPeerSettings - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserGetPeerSettings(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.getPeerSettings - reply: %s", r.DebugString())
	return r, err
}

// UserDeletePeerSettings
// user.deletePeerSettings user_id:long peer_type:int peer_id:long = Bool;
func (s *Service) UserDeletePeerSettings(ctx context.Context, request *user.TLUserDeletePeerSettings) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.deletePeerSettings - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserDeletePeerSettings(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.deletePeerSettings - reply: %s", r.DebugString())
	return r, err
}

// UserChangePhone
// user.changePhone user_id:long phone:string = Bool;
func (s *Service) UserChangePhone(ctx context.Context, request *user.TLUserChangePhone) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.changePhone - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserChangePhone(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.changePhone - reply: %s", r.DebugString())
	return r, err
}

// UserCreateNewPredefinedUser
// user.createNewPredefinedUser flags:# phone:string first_name:string last_name:flags.0?string username:string code:string verified:flags.1?true = PredefinedUser;
func (s *Service) UserCreateNewPredefinedUser(ctx context.Context, request *user.TLUserCreateNewPredefinedUser) (*mtproto.PredefinedUser, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.createNewPredefinedUser - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserCreateNewPredefinedUser(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.createNewPredefinedUser - reply: %s", r.DebugString())
	return r, err
}

// UserGetPredefinedUser
// user.getPredefinedUser phone:string = PredefinedUser;
func (s *Service) UserGetPredefinedUser(ctx context.Context, request *user.TLUserGetPredefinedUser) (*mtproto.PredefinedUser, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.getPredefinedUser - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserGetPredefinedUser(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.getPredefinedUser - reply: %s", r.DebugString())
	return r, err
}

// UserGetAllPredefinedUser
// user.getAllPredefinedUser = Vector<PredefinedUser>;
func (s *Service) UserGetAllPredefinedUser(ctx context.Context, request *user.TLUserGetAllPredefinedUser) (*user.Vector_PredefinedUser, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.getAllPredefinedUser - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserGetAllPredefinedUser(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.getAllPredefinedUser - reply: %s", r.DebugString())
	return r, err
}

// UserUpdatePredefinedFirstAndLastName
// user.updatePredefinedFirstAndLastName flags:# phone:string first_name:string last_name:flags.0?string = PredefinedUser;
func (s *Service) UserUpdatePredefinedFirstAndLastName(ctx context.Context, request *user.TLUserUpdatePredefinedFirstAndLastName) (*mtproto.PredefinedUser, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.updatePredefinedFirstAndLastName - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserUpdatePredefinedFirstAndLastName(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.updatePredefinedFirstAndLastName - reply: %s", r.DebugString())
	return r, err
}

// UserUpdatePredefinedVerified
// user.updatePredefinedVerified flags:# phone:string verified:flags.1?true = PredefinedUser;
func (s *Service) UserUpdatePredefinedVerified(ctx context.Context, request *user.TLUserUpdatePredefinedVerified) (*mtproto.PredefinedUser, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.updatePredefinedVerified - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserUpdatePredefinedVerified(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.updatePredefinedVerified - reply: %s", r.DebugString())
	return r, err
}

// UserUpdatePredefinedUsername
// user.updatePredefinedUsername flags:# phone:string username:flags.1?string = PredefinedUser;
func (s *Service) UserUpdatePredefinedUsername(ctx context.Context, request *user.TLUserUpdatePredefinedUsername) (*mtproto.PredefinedUser, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.updatePredefinedUsername - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserUpdatePredefinedUsername(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.updatePredefinedUsername - reply: %s", r.DebugString())
	return r, err
}

// UserUpdatePredefinedCode
// user.updatePredefinedCode phone:string code:string = PredefinedUser;
func (s *Service) UserUpdatePredefinedCode(ctx context.Context, request *user.TLUserUpdatePredefinedCode) (*mtproto.PredefinedUser, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.updatePredefinedCode - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserUpdatePredefinedCode(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.updatePredefinedCode - reply: %s", r.DebugString())
	return r, err
}

// UserPredefinedBindRegisteredUserId
// user.predefinedBindRegisteredUserId phone:string registered_userId:long = Bool;
func (s *Service) UserPredefinedBindRegisteredUserId(ctx context.Context, request *user.TLUserPredefinedBindRegisteredUserId) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.predefinedBindRegisteredUserId - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserPredefinedBindRegisteredUserId(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.predefinedBindRegisteredUserId - reply: %s", r.DebugString())
	return r, err
}

// UserCreateNewUser
// user.createNewUser secret_key_id:long phone:string country_code:string first_name:string last_name:string = ImmutableUser;
func (s *Service) UserCreateNewUser(ctx context.Context, request *user.TLUserCreateNewUser) (*user.ImmutableUser, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.createNewUser - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserCreateNewUser(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.createNewUser - reply: %s", r.DebugString())
	return r, err
}

// UserBlockPeer
// user.blockPeer user_id:long peer_type:int peer_id:long = Bool;
func (s *Service) UserBlockPeer(ctx context.Context, request *user.TLUserBlockPeer) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.blockPeer - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserBlockPeer(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.blockPeer - reply: %s", r.DebugString())
	return r, err
}

// UserUnBlockPeer
// user.unBlockPeer user_id:long peer_type:int peer_id:long = Bool;
func (s *Service) UserUnBlockPeer(ctx context.Context, request *user.TLUserUnBlockPeer) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.unBlockPeer - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserUnBlockPeer(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.unBlockPeer - reply: %s", r.DebugString())
	return r, err
}

// UserBlockedByUser
// user.blockedByUser user_id:long peer_user_id:long = Bool;
func (s *Service) UserBlockedByUser(ctx context.Context, request *user.TLUserBlockedByUser) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.blockedByUser - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserBlockedByUser(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.blockedByUser - reply: %s", r.DebugString())
	return r, err
}

// UserIsBlockedByUser
// user.isBlockedByUser user_id:long peer_user_id:long = Bool;
func (s *Service) UserIsBlockedByUser(ctx context.Context, request *user.TLUserIsBlockedByUser) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.isBlockedByUser - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserIsBlockedByUser(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.isBlockedByUser - reply: %s", r.DebugString())
	return r, err
}

// UserCheckBlockUserList
// user.checkBlockUserList user_id:long id:Vector<long> = Vector<long>;
func (s *Service) UserCheckBlockUserList(ctx context.Context, request *user.TLUserCheckBlockUserList) (*user.Vector_Long, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.checkBlockUserList - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserCheckBlockUserList(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.checkBlockUserList - reply: %s", r.DebugString())
	return r, err
}

// UserGetBlockedList
// user.getBlockedList user_id:long offset:int limit:int = Vector<PeerBlocked>;
func (s *Service) UserGetBlockedList(ctx context.Context, request *user.TLUserGetBlockedList) (*user.Vector_PeerBlocked, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.getBlockedList - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserGetBlockedList(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.getBlockedList - reply: %s", r.DebugString())
	return r, err
}

// UserGetContactSignUpNotification
// user.getContactSignUpNotification user_id:long = Bool;
func (s *Service) UserGetContactSignUpNotification(ctx context.Context, request *user.TLUserGetContactSignUpNotification) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.getContactSignUpNotification - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserGetContactSignUpNotification(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.getContactSignUpNotification - reply: %s", r.DebugString())
	return r, err
}

// UserSetContactSignUpNotification
// user.setContactSignUpNotification user_id:long silent:Bool = Bool;
func (s *Service) UserSetContactSignUpNotification(ctx context.Context, request *user.TLUserSetContactSignUpNotification) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.setContactSignUpNotification - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserSetContactSignUpNotification(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.setContactSignUpNotification - reply: %s", r.DebugString())
	return r, err
}

// UserGetContentSettings
// user.getContentSettings user_id:long = account.ContentSettings;
func (s *Service) UserGetContentSettings(ctx context.Context, request *user.TLUserGetContentSettings) (*mtproto.Account_ContentSettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.getContentSettings - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserGetContentSettings(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.getContentSettings - reply: %s", r.DebugString())
	return r, err
}

// UserSetContentSettings
// user.setContentSettings flags:# user_id:long sensitive_enabled:flags.0?true = Bool;
func (s *Service) UserSetContentSettings(ctx context.Context, request *user.TLUserSetContentSettings) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.setContentSettings - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserSetContentSettings(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.setContentSettings - reply: %s", r.DebugString())
	return r, err
}

// UserDeleteContact
// user.deleteContact user_id:long id:long = Bool;
func (s *Service) UserDeleteContact(ctx context.Context, request *user.TLUserDeleteContact) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.deleteContact - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserDeleteContact(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.deleteContact - reply: %s", r.DebugString())
	return r, err
}

// UserGetContactList
// user.getContactList user_id:long = Vector<ContactData>;
func (s *Service) UserGetContactList(ctx context.Context, request *user.TLUserGetContactList) (*user.Vector_ContactData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.getContactList - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserGetContactList(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.getContactList - reply: %s", r.DebugString())
	return r, err
}

// UserGetContactIdList
// user.getContactIdList user_id:long = Vector<long>;
func (s *Service) UserGetContactIdList(ctx context.Context, request *user.TLUserGetContactIdList) (*user.Vector_Long, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.getContactIdList - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserGetContactIdList(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.getContactIdList - reply: %s", r.DebugString())
	return r, err
}

// UserGetContact
// user.getContact user_id:long id:long = ContactData;
func (s *Service) UserGetContact(ctx context.Context, request *user.TLUserGetContact) (*user.ContactData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.getContact - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserGetContact(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.getContact - reply: %s", r.DebugString())
	return r, err
}

// UserAddContact
// user.addContact user_id:long add_phone_privacy_exception:Bool id:long first_name:string last_name:string phone:string = Bool;
func (s *Service) UserAddContact(ctx context.Context, request *user.TLUserAddContact) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.addContact - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserAddContact(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.addContact - reply: %s", r.DebugString())
	return r, err
}

// UserCheckContact
// user.checkContact user_id:long id:long = Bool;
func (s *Service) UserCheckContact(ctx context.Context, request *user.TLUserCheckContact) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.checkContact - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserCheckContact(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.checkContact - reply: %s", r.DebugString())
	return r, err
}

// UserImportContacts
// user.importContacts user_id:long contacts:Vector<InputContact> = UserImportedContacts;
func (s *Service) UserImportContacts(ctx context.Context, request *user.TLUserImportContacts) (*user.UserImportedContacts, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.importContacts - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserImportContacts(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.importContacts - reply: %s", r.DebugString())
	return r, err
}

// UserGetCountryCode
// user.getCountryCode user_id:long = String;
func (s *Service) UserGetCountryCode(ctx context.Context, request *user.TLUserGetCountryCode) (*mtproto.String, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.getCountryCode - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserGetCountryCode(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.getCountryCode - reply: %s", r.DebugString())
	return r, err
}

// UserUpdateAbout
// user.updateAbout user_id:long about:string = Bool;
func (s *Service) UserUpdateAbout(ctx context.Context, request *user.TLUserUpdateAbout) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.updateAbout - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserUpdateAbout(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.updateAbout - reply: %s", r.DebugString())
	return r, err
}

// UserUpdateFirstAndLastName
// user.updateFirstAndLastName user_id:long first_name:string last_name:string = Bool;
func (s *Service) UserUpdateFirstAndLastName(ctx context.Context, request *user.TLUserUpdateFirstAndLastName) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.updateFirstAndLastName - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserUpdateFirstAndLastName(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.updateFirstAndLastName - reply: %s", r.DebugString())
	return r, err
}

// UserUpdateVerified
// user.updateVerified user_id:long verified:Bool = Bool;
func (s *Service) UserUpdateVerified(ctx context.Context, request *user.TLUserUpdateVerified) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.updateVerified - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserUpdateVerified(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.updateVerified - reply: %s", r.DebugString())
	return r, err
}

// UserUpdateUsername
// user.updateUsername user_id:long username:string = Bool;
func (s *Service) UserUpdateUsername(ctx context.Context, request *user.TLUserUpdateUsername) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.updateUsername - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserUpdateUsername(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.updateUsername - reply: %s", r.DebugString())
	return r, err
}

// UserUpdateProfilePhoto
// user.updateProfilePhoto user_id:long id:long = Int64;
func (s *Service) UserUpdateProfilePhoto(ctx context.Context, request *user.TLUserUpdateProfilePhoto) (*mtproto.Int64, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.updateProfilePhoto - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserUpdateProfilePhoto(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.updateProfilePhoto - reply: %s", r.DebugString())
	return r, err
}

// UserDeleteProfilePhotos
// user.deleteProfilePhotos user_id:long id:Vector<long> = Int64;
func (s *Service) UserDeleteProfilePhotos(ctx context.Context, request *user.TLUserDeleteProfilePhotos) (*mtproto.Int64, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.deleteProfilePhotos - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserDeleteProfilePhotos(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.deleteProfilePhotos - reply: %s", r.DebugString())
	return r, err
}

// UserGetProfilePhotos
// user.getProfilePhotos user_id:long = Vector<long>;
func (s *Service) UserGetProfilePhotos(ctx context.Context, request *user.TLUserGetProfilePhotos) (*user.Vector_Long, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.getProfilePhotos - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserGetProfilePhotos(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.getProfilePhotos - reply: %s", r.DebugString())
	return r, err
}

// UserSetBotCommands
// user.setBotCommands user_id:long bot_id:long commands:Vector<BotCommand> = Bool;
func (s *Service) UserSetBotCommands(ctx context.Context, request *user.TLUserSetBotCommands) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.setBotCommands - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserSetBotCommands(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.setBotCommands - reply: %s", r.DebugString())
	return r, err
}

// UserIsBot
// user.isBot id:long = Bool;
func (s *Service) UserIsBot(ctx context.Context, request *user.TLUserIsBot) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.isBot - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserIsBot(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.isBot - reply: %s", r.DebugString())
	return r, err
}

// UserGetBotInfo
// user.getBotInfo bot_id:long = BotInfo;
func (s *Service) UserGetBotInfo(ctx context.Context, request *user.TLUserGetBotInfo) (*mtproto.BotInfo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.getBotInfo - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserGetBotInfo(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.getBotInfo - reply: %s", r.DebugString())
	return r, err
}

// UserGetFullUser
// user.getFullUser self_user_id:long id:long = users.UserFull;
func (s *Service) UserGetFullUser(ctx context.Context, request *user.TLUserGetFullUser) (*mtproto.Users_UserFull, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("user.getFullUser - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.UserGetFullUser(request)
	if err != nil {
		return nil, err
	}

	c.Infof("user.getFullUser - reply: %s", r.DebugString())
	return r, err
}
