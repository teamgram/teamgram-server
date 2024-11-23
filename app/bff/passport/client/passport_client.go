/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package passportclient

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type PassportClient interface {
	AccountGetAuthorizations(ctx context.Context, in *mtproto.TLAccountGetAuthorizations) (*mtproto.Account_Authorizations, error)
	AccountGetAllSecureValues(ctx context.Context, in *mtproto.TLAccountGetAllSecureValues) (*mtproto.Vector_SecureValue, error)
	AccountGetSecureValue(ctx context.Context, in *mtproto.TLAccountGetSecureValue) (*mtproto.Vector_SecureValue, error)
	AccountSaveSecureValue(ctx context.Context, in *mtproto.TLAccountSaveSecureValue) (*mtproto.SecureValue, error)
	AccountDeleteSecureValue(ctx context.Context, in *mtproto.TLAccountDeleteSecureValue) (*mtproto.Bool, error)
	AccountGetAuthorizationForm(ctx context.Context, in *mtproto.TLAccountGetAuthorizationForm) (*mtproto.Account_AuthorizationForm, error)
	AccountAcceptAuthorization(ctx context.Context, in *mtproto.TLAccountAcceptAuthorization) (*mtproto.Bool, error)
	AccountSendVerifyPhoneCode(ctx context.Context, in *mtproto.TLAccountSendVerifyPhoneCode) (*mtproto.Auth_SentCode, error)
	AccountVerifyPhone(ctx context.Context, in *mtproto.TLAccountVerifyPhone) (*mtproto.Bool, error)
	UsersSetSecureValueErrors(ctx context.Context, in *mtproto.TLUsersSetSecureValueErrors) (*mtproto.Bool, error)
	HelpGetPassportConfig(ctx context.Context, in *mtproto.TLHelpGetPassportConfig) (*mtproto.Help_PassportConfig, error)
}

type defaultPassportClient struct {
	cli zrpc.Client
}

func NewPassportClient(cli zrpc.Client) PassportClient {
	return &defaultPassportClient{
		cli: cli,
	}
}

// AccountGetAuthorizations
// account.getAuthorizations#e320c158 = account.Authorizations;
func (m *defaultPassportClient) AccountGetAuthorizations(ctx context.Context, in *mtproto.TLAccountGetAuthorizations) (*mtproto.Account_Authorizations, error) {
	client := mtproto.NewRPCPassportClient(m.cli.Conn())
	return client.AccountGetAuthorizations(ctx, in)
}

// AccountGetAllSecureValues
// account.getAllSecureValues#b288bc7d = Vector<SecureValue>;
func (m *defaultPassportClient) AccountGetAllSecureValues(ctx context.Context, in *mtproto.TLAccountGetAllSecureValues) (*mtproto.Vector_SecureValue, error) {
	client := mtproto.NewRPCPassportClient(m.cli.Conn())
	return client.AccountGetAllSecureValues(ctx, in)
}

// AccountGetSecureValue
// account.getSecureValue#73665bc2 types:Vector<SecureValueType> = Vector<SecureValue>;
func (m *defaultPassportClient) AccountGetSecureValue(ctx context.Context, in *mtproto.TLAccountGetSecureValue) (*mtproto.Vector_SecureValue, error) {
	client := mtproto.NewRPCPassportClient(m.cli.Conn())
	return client.AccountGetSecureValue(ctx, in)
}

// AccountSaveSecureValue
// account.saveSecureValue#899fe31d value:InputSecureValue secure_secret_id:long = SecureValue;
func (m *defaultPassportClient) AccountSaveSecureValue(ctx context.Context, in *mtproto.TLAccountSaveSecureValue) (*mtproto.SecureValue, error) {
	client := mtproto.NewRPCPassportClient(m.cli.Conn())
	return client.AccountSaveSecureValue(ctx, in)
}

// AccountDeleteSecureValue
// account.deleteSecureValue#b880bc4b types:Vector<SecureValueType> = Bool;
func (m *defaultPassportClient) AccountDeleteSecureValue(ctx context.Context, in *mtproto.TLAccountDeleteSecureValue) (*mtproto.Bool, error) {
	client := mtproto.NewRPCPassportClient(m.cli.Conn())
	return client.AccountDeleteSecureValue(ctx, in)
}

// AccountGetAuthorizationForm
// account.getAuthorizationForm#a929597a bot_id:long scope:string public_key:string = account.AuthorizationForm;
func (m *defaultPassportClient) AccountGetAuthorizationForm(ctx context.Context, in *mtproto.TLAccountGetAuthorizationForm) (*mtproto.Account_AuthorizationForm, error) {
	client := mtproto.NewRPCPassportClient(m.cli.Conn())
	return client.AccountGetAuthorizationForm(ctx, in)
}

// AccountAcceptAuthorization
// account.acceptAuthorization#f3ed4c73 bot_id:long scope:string public_key:string value_hashes:Vector<SecureValueHash> credentials:SecureCredentialsEncrypted = Bool;
func (m *defaultPassportClient) AccountAcceptAuthorization(ctx context.Context, in *mtproto.TLAccountAcceptAuthorization) (*mtproto.Bool, error) {
	client := mtproto.NewRPCPassportClient(m.cli.Conn())
	return client.AccountAcceptAuthorization(ctx, in)
}

// AccountSendVerifyPhoneCode
// account.sendVerifyPhoneCode#a5a356f9 phone_number:string settings:CodeSettings = auth.SentCode;
func (m *defaultPassportClient) AccountSendVerifyPhoneCode(ctx context.Context, in *mtproto.TLAccountSendVerifyPhoneCode) (*mtproto.Auth_SentCode, error) {
	client := mtproto.NewRPCPassportClient(m.cli.Conn())
	return client.AccountSendVerifyPhoneCode(ctx, in)
}

// AccountVerifyPhone
// account.verifyPhone#4dd3a7f6 phone_number:string phone_code_hash:string phone_code:string = Bool;
func (m *defaultPassportClient) AccountVerifyPhone(ctx context.Context, in *mtproto.TLAccountVerifyPhone) (*mtproto.Bool, error) {
	client := mtproto.NewRPCPassportClient(m.cli.Conn())
	return client.AccountVerifyPhone(ctx, in)
}

// UsersSetSecureValueErrors
// users.setSecureValueErrors#90c894b5 id:InputUser errors:Vector<SecureValueError> = Bool;
func (m *defaultPassportClient) UsersSetSecureValueErrors(ctx context.Context, in *mtproto.TLUsersSetSecureValueErrors) (*mtproto.Bool, error) {
	client := mtproto.NewRPCPassportClient(m.cli.Conn())
	return client.UsersSetSecureValueErrors(ctx, in)
}

// HelpGetPassportConfig
// help.getPassportConfig#c661ad08 hash:int = help.PassportConfig;
func (m *defaultPassportClient) HelpGetPassportConfig(ctx context.Context, in *mtproto.TLHelpGetPassportConfig) (*mtproto.Help_PassportConfig, error) {
	client := mtproto.NewRPCPassportClient(m.cli.Conn())
	return client.HelpGetPassportConfig(ctx, in)
}
