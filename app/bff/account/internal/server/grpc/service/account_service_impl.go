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
	"github.com/teamgram/teamgram-server/app/bff/account/internal/core"
)

// AccountUpdateProfile
// account.updateProfile#78515775 flags:# first_name:flags.0?string last_name:flags.1?string about:flags.2?string = User;
func (s *Service) AccountUpdateProfile(ctx context.Context, request *mtproto.TLAccountUpdateProfile) (*mtproto.User, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.updateProfile - metadata: %s, request: %s", c.MD, request)

	r, err := c.AccountUpdateProfile(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.updateProfile - reply: %s", r)
	return r, err
}

// AccountUpdateStatus
// account.updateStatus#6628562c offline:Bool = Bool;
func (s *Service) AccountUpdateStatus(ctx context.Context, request *mtproto.TLAccountUpdateStatus) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.updateStatus - metadata: %s, request: %s", c.MD, request)

	r, err := c.AccountUpdateStatus(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.updateStatus - reply: %s", r)
	return r, err
}

// AccountGetPrivacy
// account.getPrivacy#dadbc950 key:InputPrivacyKey = account.PrivacyRules;
func (s *Service) AccountGetPrivacy(ctx context.Context, request *mtproto.TLAccountGetPrivacy) (*mtproto.Account_PrivacyRules, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.getPrivacy - metadata: %s, request: %s", c.MD, request)

	r, err := c.AccountGetPrivacy(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.getPrivacy - reply: %s", r)
	return r, err
}

// AccountSetPrivacy
// account.setPrivacy#c9f81ce8 key:InputPrivacyKey rules:Vector<InputPrivacyRule> = account.PrivacyRules;
func (s *Service) AccountSetPrivacy(ctx context.Context, request *mtproto.TLAccountSetPrivacy) (*mtproto.Account_PrivacyRules, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.setPrivacy - metadata: %s, request: %s", c.MD, request)

	r, err := c.AccountSetPrivacy(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.setPrivacy - reply: %s", r)
	return r, err
}

// AccountDeleteAccount
// account.deleteAccount#a2c0cf74 flags:# reason:string password:flags.0?InputCheckPasswordSRP = Bool;
func (s *Service) AccountDeleteAccount(ctx context.Context, request *mtproto.TLAccountDeleteAccount) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.deleteAccount - metadata: %s, request: %s", c.MD, request)

	r, err := c.AccountDeleteAccount(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.deleteAccount - reply: %s", r)
	return r, err
}

// AccountGetAccountTTL
// account.getAccountTTL#8fc711d = AccountDaysTTL;
func (s *Service) AccountGetAccountTTL(ctx context.Context, request *mtproto.TLAccountGetAccountTTL) (*mtproto.AccountDaysTTL, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.getAccountTTL - metadata: %s, request: %s", c.MD, request)

	r, err := c.AccountGetAccountTTL(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.getAccountTTL - reply: %s", r)
	return r, err
}

// AccountSetAccountTTL
// account.setAccountTTL#2442485e ttl:AccountDaysTTL = Bool;
func (s *Service) AccountSetAccountTTL(ctx context.Context, request *mtproto.TLAccountSetAccountTTL) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.setAccountTTL - metadata: %s, request: %s", c.MD, request)

	r, err := c.AccountSetAccountTTL(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.setAccountTTL - reply: %s", r)
	return r, err
}

// AccountSendChangePhoneCode
// account.sendChangePhoneCode#82574ae5 phone_number:string settings:CodeSettings = auth.SentCode;
func (s *Service) AccountSendChangePhoneCode(ctx context.Context, request *mtproto.TLAccountSendChangePhoneCode) (*mtproto.Auth_SentCode, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.sendChangePhoneCode - metadata: %s, request: %s", c.MD, request)

	r, err := c.AccountSendChangePhoneCode(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.sendChangePhoneCode - reply: %s", r)
	return r, err
}

// AccountChangePhone
// account.changePhone#70c32edb phone_number:string phone_code_hash:string phone_code:string = User;
func (s *Service) AccountChangePhone(ctx context.Context, request *mtproto.TLAccountChangePhone) (*mtproto.User, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.changePhone - metadata: %s, request: %s", c.MD, request)

	r, err := c.AccountChangePhone(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.changePhone - reply: %s", r)
	return r, err
}

// AccountResetAuthorization
// account.resetAuthorization#df77f3bc hash:long = Bool;
func (s *Service) AccountResetAuthorization(ctx context.Context, request *mtproto.TLAccountResetAuthorization) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.resetAuthorization - metadata: %s, request: %s", c.MD, request)

	r, err := c.AccountResetAuthorization(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.resetAuthorization - reply: %s", r)
	return r, err
}

// AccountSendConfirmPhoneCode
// account.sendConfirmPhoneCode#1b3faa88 hash:string settings:CodeSettings = auth.SentCode;
func (s *Service) AccountSendConfirmPhoneCode(ctx context.Context, request *mtproto.TLAccountSendConfirmPhoneCode) (*mtproto.Auth_SentCode, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.sendConfirmPhoneCode - metadata: %s, request: %s", c.MD, request)

	r, err := c.AccountSendConfirmPhoneCode(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.sendConfirmPhoneCode - reply: %s", r)
	return r, err
}

// AccountConfirmPhone
// account.confirmPhone#5f2178c3 phone_code_hash:string phone_code:string = Bool;
func (s *Service) AccountConfirmPhone(ctx context.Context, request *mtproto.TLAccountConfirmPhone) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.confirmPhone - metadata: %s, request: %s", c.MD, request)

	r, err := c.AccountConfirmPhone(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.confirmPhone - reply: %s", r)
	return r, err
}

// AccountGetGlobalPrivacySettings
// account.getGlobalPrivacySettings#eb2b4cf6 = GlobalPrivacySettings;
func (s *Service) AccountGetGlobalPrivacySettings(ctx context.Context, request *mtproto.TLAccountGetGlobalPrivacySettings) (*mtproto.GlobalPrivacySettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.getGlobalPrivacySettings - metadata: %s, request: %s", c.MD, request)

	r, err := c.AccountGetGlobalPrivacySettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.getGlobalPrivacySettings - reply: %s", r)
	return r, err
}

// AccountSetGlobalPrivacySettings
// account.setGlobalPrivacySettings#1edaaac2 settings:GlobalPrivacySettings = GlobalPrivacySettings;
func (s *Service) AccountSetGlobalPrivacySettings(ctx context.Context, request *mtproto.TLAccountSetGlobalPrivacySettings) (*mtproto.GlobalPrivacySettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.setGlobalPrivacySettings - metadata: %s, request: %s", c.MD, request)

	r, err := c.AccountSetGlobalPrivacySettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.setGlobalPrivacySettings - reply: %s", r)
	return r, err
}

// AccountUpdateBirthday
// account.updateBirthday#cc6e0c11 flags:# birthday:flags.0?Birthday = Bool;
func (s *Service) AccountUpdateBirthday(ctx context.Context, request *mtproto.TLAccountUpdateBirthday) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.updateBirthday - metadata: %s, request: %s", c.MD, request)

	r, err := c.AccountUpdateBirthday(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.updateBirthday - reply: %s", r)
	return r, err
}

// MessagesSetDefaultHistoryTTL
// messages.setDefaultHistoryTTL#9eb51445 period:int = Bool;
func (s *Service) MessagesSetDefaultHistoryTTL(ctx context.Context, request *mtproto.TLMessagesSetDefaultHistoryTTL) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.setDefaultHistoryTTL - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessagesSetDefaultHistoryTTL(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.setDefaultHistoryTTL - reply: %s", r)
	return r, err
}

// MessagesGetDefaultHistoryTTL
// messages.getDefaultHistoryTTL#658b7188 = DefaultHistoryTTL;
func (s *Service) MessagesGetDefaultHistoryTTL(ctx context.Context, request *mtproto.TLMessagesGetDefaultHistoryTTL) (*mtproto.DefaultHistoryTTL, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getDefaultHistoryTTL - metadata: %s, request: %s", c.MD, request)

	r, err := c.MessagesGetDefaultHistoryTTL(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getDefaultHistoryTTL - reply: %s", r)
	return r, err
}

// AccountUpdateVerified
// account.updateVerified flags:# id:long verified:flags.0?true = User;
func (s *Service) AccountUpdateVerified(ctx context.Context, request *mtproto.TLAccountUpdateVerified) (*mtproto.User, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.updateVerified - metadata: %s, request: %s", c.MD, request)

	r, err := c.AccountUpdateVerified(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.updateVerified - reply: %s", r)
	return r, err
}
