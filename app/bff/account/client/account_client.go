/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package account_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type AccountClient interface {
	AccountUpdateProfile(ctx context.Context, in *mtproto.TLAccountUpdateProfile) (*mtproto.User, error)
	AccountUpdateStatus(ctx context.Context, in *mtproto.TLAccountUpdateStatus) (*mtproto.Bool, error)
	AccountGetPrivacy(ctx context.Context, in *mtproto.TLAccountGetPrivacy) (*mtproto.Account_PrivacyRules, error)
	AccountSetPrivacy(ctx context.Context, in *mtproto.TLAccountSetPrivacy) (*mtproto.Account_PrivacyRules, error)
	AccountDeleteAccount(ctx context.Context, in *mtproto.TLAccountDeleteAccount) (*mtproto.Bool, error)
	AccountGetAccountTTL(ctx context.Context, in *mtproto.TLAccountGetAccountTTL) (*mtproto.AccountDaysTTL, error)
	AccountSetAccountTTL(ctx context.Context, in *mtproto.TLAccountSetAccountTTL) (*mtproto.Bool, error)
	AccountSendChangePhoneCode(ctx context.Context, in *mtproto.TLAccountSendChangePhoneCode) (*mtproto.Auth_SentCode, error)
	AccountChangePhone(ctx context.Context, in *mtproto.TLAccountChangePhone) (*mtproto.User, error)
	AccountResetAuthorization(ctx context.Context, in *mtproto.TLAccountResetAuthorization) (*mtproto.Bool, error)
	AccountSendConfirmPhoneCode(ctx context.Context, in *mtproto.TLAccountSendConfirmPhoneCode) (*mtproto.Auth_SentCode, error)
	AccountConfirmPhone(ctx context.Context, in *mtproto.TLAccountConfirmPhone) (*mtproto.Bool, error)
	AccountGetGlobalPrivacySettings(ctx context.Context, in *mtproto.TLAccountGetGlobalPrivacySettings) (*mtproto.GlobalPrivacySettings, error)
	AccountSetGlobalPrivacySettings(ctx context.Context, in *mtproto.TLAccountSetGlobalPrivacySettings) (*mtproto.GlobalPrivacySettings, error)
	AccountSetAuthorizationTTL(ctx context.Context, in *mtproto.TLAccountSetAuthorizationTTL) (*mtproto.Bool, error)
	AccountChangeAuthorizationSettings(ctx context.Context, in *mtproto.TLAccountChangeAuthorizationSettings) (*mtproto.Bool, error)
	AccountCreatePredefinedUser(ctx context.Context, in *mtproto.TLAccountCreatePredefinedUser) (*mtproto.PredefinedUser, error)
	AccountUpdatePredefinedUsername(ctx context.Context, in *mtproto.TLAccountUpdatePredefinedUsername) (*mtproto.PredefinedUser, error)
	AccountUpdatePredefinedProfile(ctx context.Context, in *mtproto.TLAccountUpdatePredefinedProfile) (*mtproto.PredefinedUser, error)
	AccountUpdateVerified(ctx context.Context, in *mtproto.TLAccountUpdateVerified) (*mtproto.User, error)
	AccountUpdatePredefinedVerified(ctx context.Context, in *mtproto.TLAccountUpdatePredefinedVerified) (*mtproto.PredefinedUser, error)
	AccountUpdatePredefinedCode(ctx context.Context, in *mtproto.TLAccountUpdatePredefinedCode) (*mtproto.PredefinedUser, error)
}

type defaultAccountClient struct {
	cli zrpc.Client
}

func NewAccountClient(cli zrpc.Client) AccountClient {
	return &defaultAccountClient{
		cli: cli,
	}
}

// AccountUpdateProfile
// account.updateProfile#78515775 flags:# first_name:flags.0?string last_name:flags.1?string about:flags.2?string = User;
func (m *defaultAccountClient) AccountUpdateProfile(ctx context.Context, in *mtproto.TLAccountUpdateProfile) (*mtproto.User, error) {
	client := mtproto.NewRPCAccountClient(m.cli.Conn())
	return client.AccountUpdateProfile(ctx, in)
}

// AccountUpdateStatus
// account.updateStatus#6628562c offline:Bool = Bool;
func (m *defaultAccountClient) AccountUpdateStatus(ctx context.Context, in *mtproto.TLAccountUpdateStatus) (*mtproto.Bool, error) {
	client := mtproto.NewRPCAccountClient(m.cli.Conn())
	return client.AccountUpdateStatus(ctx, in)
}

// AccountGetPrivacy
// account.getPrivacy#dadbc950 key:InputPrivacyKey = account.PrivacyRules;
func (m *defaultAccountClient) AccountGetPrivacy(ctx context.Context, in *mtproto.TLAccountGetPrivacy) (*mtproto.Account_PrivacyRules, error) {
	client := mtproto.NewRPCAccountClient(m.cli.Conn())
	return client.AccountGetPrivacy(ctx, in)
}

// AccountSetPrivacy
// account.setPrivacy#c9f81ce8 key:InputPrivacyKey rules:Vector<InputPrivacyRule> = account.PrivacyRules;
func (m *defaultAccountClient) AccountSetPrivacy(ctx context.Context, in *mtproto.TLAccountSetPrivacy) (*mtproto.Account_PrivacyRules, error) {
	client := mtproto.NewRPCAccountClient(m.cli.Conn())
	return client.AccountSetPrivacy(ctx, in)
}

// AccountDeleteAccount
// account.deleteAccount#418d4e0b reason:string = Bool;
func (m *defaultAccountClient) AccountDeleteAccount(ctx context.Context, in *mtproto.TLAccountDeleteAccount) (*mtproto.Bool, error) {
	client := mtproto.NewRPCAccountClient(m.cli.Conn())
	return client.AccountDeleteAccount(ctx, in)
}

// AccountGetAccountTTL
// account.getAccountTTL#8fc711d = AccountDaysTTL;
func (m *defaultAccountClient) AccountGetAccountTTL(ctx context.Context, in *mtproto.TLAccountGetAccountTTL) (*mtproto.AccountDaysTTL, error) {
	client := mtproto.NewRPCAccountClient(m.cli.Conn())
	return client.AccountGetAccountTTL(ctx, in)
}

// AccountSetAccountTTL
// account.setAccountTTL#2442485e ttl:AccountDaysTTL = Bool;
func (m *defaultAccountClient) AccountSetAccountTTL(ctx context.Context, in *mtproto.TLAccountSetAccountTTL) (*mtproto.Bool, error) {
	client := mtproto.NewRPCAccountClient(m.cli.Conn())
	return client.AccountSetAccountTTL(ctx, in)
}

// AccountSendChangePhoneCode
// account.sendChangePhoneCode#82574ae5 phone_number:string settings:CodeSettings = auth.SentCode;
func (m *defaultAccountClient) AccountSendChangePhoneCode(ctx context.Context, in *mtproto.TLAccountSendChangePhoneCode) (*mtproto.Auth_SentCode, error) {
	client := mtproto.NewRPCAccountClient(m.cli.Conn())
	return client.AccountSendChangePhoneCode(ctx, in)
}

// AccountChangePhone
// account.changePhone#70c32edb phone_number:string phone_code_hash:string phone_code:string = User;
func (m *defaultAccountClient) AccountChangePhone(ctx context.Context, in *mtproto.TLAccountChangePhone) (*mtproto.User, error) {
	client := mtproto.NewRPCAccountClient(m.cli.Conn())
	return client.AccountChangePhone(ctx, in)
}

// AccountResetAuthorization
// account.resetAuthorization#df77f3bc hash:long = Bool;
func (m *defaultAccountClient) AccountResetAuthorization(ctx context.Context, in *mtproto.TLAccountResetAuthorization) (*mtproto.Bool, error) {
	client := mtproto.NewRPCAccountClient(m.cli.Conn())
	return client.AccountResetAuthorization(ctx, in)
}

// AccountSendConfirmPhoneCode
// account.sendConfirmPhoneCode#1b3faa88 hash:string settings:CodeSettings = auth.SentCode;
func (m *defaultAccountClient) AccountSendConfirmPhoneCode(ctx context.Context, in *mtproto.TLAccountSendConfirmPhoneCode) (*mtproto.Auth_SentCode, error) {
	client := mtproto.NewRPCAccountClient(m.cli.Conn())
	return client.AccountSendConfirmPhoneCode(ctx, in)
}

// AccountConfirmPhone
// account.confirmPhone#5f2178c3 phone_code_hash:string phone_code:string = Bool;
func (m *defaultAccountClient) AccountConfirmPhone(ctx context.Context, in *mtproto.TLAccountConfirmPhone) (*mtproto.Bool, error) {
	client := mtproto.NewRPCAccountClient(m.cli.Conn())
	return client.AccountConfirmPhone(ctx, in)
}

// AccountGetGlobalPrivacySettings
// account.getGlobalPrivacySettings#eb2b4cf6 = GlobalPrivacySettings;
func (m *defaultAccountClient) AccountGetGlobalPrivacySettings(ctx context.Context, in *mtproto.TLAccountGetGlobalPrivacySettings) (*mtproto.GlobalPrivacySettings, error) {
	client := mtproto.NewRPCAccountClient(m.cli.Conn())
	return client.AccountGetGlobalPrivacySettings(ctx, in)
}

// AccountSetGlobalPrivacySettings
// account.setGlobalPrivacySettings#1edaaac2 settings:GlobalPrivacySettings = GlobalPrivacySettings;
func (m *defaultAccountClient) AccountSetGlobalPrivacySettings(ctx context.Context, in *mtproto.TLAccountSetGlobalPrivacySettings) (*mtproto.GlobalPrivacySettings, error) {
	client := mtproto.NewRPCAccountClient(m.cli.Conn())
	return client.AccountSetGlobalPrivacySettings(ctx, in)
}

// AccountSetAuthorizationTTL
// account.setAuthorizationTTL#bf899aa0 authorization_ttl_days:int = Bool;
func (m *defaultAccountClient) AccountSetAuthorizationTTL(ctx context.Context, in *mtproto.TLAccountSetAuthorizationTTL) (*mtproto.Bool, error) {
	client := mtproto.NewRPCAccountClient(m.cli.Conn())
	return client.AccountSetAuthorizationTTL(ctx, in)
}

// AccountChangeAuthorizationSettings
// account.changeAuthorizationSettings#40f48462 flags:# hash:long encrypted_requests_disabled:flags.0?Bool call_requests_disabled:flags.1?Bool = Bool;
func (m *defaultAccountClient) AccountChangeAuthorizationSettings(ctx context.Context, in *mtproto.TLAccountChangeAuthorizationSettings) (*mtproto.Bool, error) {
	client := mtproto.NewRPCAccountClient(m.cli.Conn())
	return client.AccountChangeAuthorizationSettings(ctx, in)
}

// AccountCreatePredefinedUser
// account.createPredefinedUser flags:# phone:string first_name:flags.0?string last_name:flags.1?string username:flags.2?string code:string verified:flags.3?true = PredefinedUser;
func (m *defaultAccountClient) AccountCreatePredefinedUser(ctx context.Context, in *mtproto.TLAccountCreatePredefinedUser) (*mtproto.PredefinedUser, error) {
	client := mtproto.NewRPCAccountClient(m.cli.Conn())
	return client.AccountCreatePredefinedUser(ctx, in)
}

// AccountUpdatePredefinedUsername
// account.updatePredefinedUsername phone:string username:string = PredefinedUser;
func (m *defaultAccountClient) AccountUpdatePredefinedUsername(ctx context.Context, in *mtproto.TLAccountUpdatePredefinedUsername) (*mtproto.PredefinedUser, error) {
	client := mtproto.NewRPCAccountClient(m.cli.Conn())
	return client.AccountUpdatePredefinedUsername(ctx, in)
}

// AccountUpdatePredefinedProfile
// account.updatePredefinedProfile flags:# phone:string first_name:flags.0?string last_name:flags.1?string about:flags.2?string = PredefinedUser;
func (m *defaultAccountClient) AccountUpdatePredefinedProfile(ctx context.Context, in *mtproto.TLAccountUpdatePredefinedProfile) (*mtproto.PredefinedUser, error) {
	client := mtproto.NewRPCAccountClient(m.cli.Conn())
	return client.AccountUpdatePredefinedProfile(ctx, in)
}

// AccountUpdateVerified
// account.updateVerified flags:# id:long verified:flags.0?true = User;
func (m *defaultAccountClient) AccountUpdateVerified(ctx context.Context, in *mtproto.TLAccountUpdateVerified) (*mtproto.User, error) {
	client := mtproto.NewRPCAccountClient(m.cli.Conn())
	return client.AccountUpdateVerified(ctx, in)
}

// AccountUpdatePredefinedVerified
// account.updatePredefinedVerified flags:# phone:string verified:flags.0?true = PredefinedUser;
func (m *defaultAccountClient) AccountUpdatePredefinedVerified(ctx context.Context, in *mtproto.TLAccountUpdatePredefinedVerified) (*mtproto.PredefinedUser, error) {
	client := mtproto.NewRPCAccountClient(m.cli.Conn())
	return client.AccountUpdatePredefinedVerified(ctx, in)
}

// AccountUpdatePredefinedCode
// account.updatePredefinedCode phone:string code:string = PredefinedUser;
func (m *defaultAccountClient) AccountUpdatePredefinedCode(ctx context.Context, in *mtproto.TLAccountUpdatePredefinedCode) (*mtproto.PredefinedUser, error) {
	client := mtproto.NewRPCAccountClient(m.cli.Conn())
	return client.AccountUpdatePredefinedCode(ctx, in)
}
