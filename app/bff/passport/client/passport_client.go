/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package passportclient

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/bff/passport/passport/passportservice"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
)

type PassportClient interface {
	AccountGetAuthorizations(ctx context.Context, in *tg.TLAccountGetAuthorizations) (*tg.AccountAuthorizations, error)
	AccountGetAllSecureValues(ctx context.Context, in *tg.TLAccountGetAllSecureValues) (*tg.VectorSecureValue, error)
	AccountGetSecureValue(ctx context.Context, in *tg.TLAccountGetSecureValue) (*tg.VectorSecureValue, error)
	AccountSaveSecureValue(ctx context.Context, in *tg.TLAccountSaveSecureValue) (*tg.SecureValue, error)
	AccountDeleteSecureValue(ctx context.Context, in *tg.TLAccountDeleteSecureValue) (*tg.Bool, error)
	AccountGetAuthorizationForm(ctx context.Context, in *tg.TLAccountGetAuthorizationForm) (*tg.AccountAuthorizationForm, error)
	AccountAcceptAuthorization(ctx context.Context, in *tg.TLAccountAcceptAuthorization) (*tg.Bool, error)
	AccountSendVerifyPhoneCode(ctx context.Context, in *tg.TLAccountSendVerifyPhoneCode) (*tg.AuthSentCode, error)
	AccountVerifyPhone(ctx context.Context, in *tg.TLAccountVerifyPhone) (*tg.Bool, error)
	UsersSetSecureValueErrors(ctx context.Context, in *tg.TLUsersSetSecureValueErrors) (*tg.Bool, error)
	HelpGetPassportConfig(ctx context.Context, in *tg.TLHelpGetPassportConfig) (*tg.HelpPassportConfig, error)
}

type defaultPassportClient struct {
	cli client.Client
	rpc passportservice.Client
}

func NewPassportClient(cli client.Client) PassportClient {
	return &defaultPassportClient{
		cli: cli,
		rpc: passportservice.NewRPCPassportClient(cli),
	}
}

// AccountGetAuthorizations
// account.getAuthorizations#e320c158 = account.Authorizations;
func (m *defaultPassportClient) AccountGetAuthorizations(ctx context.Context, in *tg.TLAccountGetAuthorizations) (*tg.AccountAuthorizations, error) {
	return m.rpc.AccountGetAuthorizations(ctx, in)
}

// AccountGetAllSecureValues
// account.getAllSecureValues#b288bc7d = Vector<SecureValue>;
func (m *defaultPassportClient) AccountGetAllSecureValues(ctx context.Context, in *tg.TLAccountGetAllSecureValues) (*tg.VectorSecureValue, error) {
	return m.rpc.AccountGetAllSecureValues(ctx, in)
}

// AccountGetSecureValue
// account.getSecureValue#73665bc2 types:Vector<SecureValueType> = Vector<SecureValue>;
func (m *defaultPassportClient) AccountGetSecureValue(ctx context.Context, in *tg.TLAccountGetSecureValue) (*tg.VectorSecureValue, error) {
	return m.rpc.AccountGetSecureValue(ctx, in)
}

// AccountSaveSecureValue
// account.saveSecureValue#899fe31d value:InputSecureValue secure_secret_id:long = SecureValue;
func (m *defaultPassportClient) AccountSaveSecureValue(ctx context.Context, in *tg.TLAccountSaveSecureValue) (*tg.SecureValue, error) {
	return m.rpc.AccountSaveSecureValue(ctx, in)
}

// AccountDeleteSecureValue
// account.deleteSecureValue#b880bc4b types:Vector<SecureValueType> = Bool;
func (m *defaultPassportClient) AccountDeleteSecureValue(ctx context.Context, in *tg.TLAccountDeleteSecureValue) (*tg.Bool, error) {
	return m.rpc.AccountDeleteSecureValue(ctx, in)
}

// AccountGetAuthorizationForm
// account.getAuthorizationForm#a929597a bot_id:long scope:string public_key:string = account.AuthorizationForm;
func (m *defaultPassportClient) AccountGetAuthorizationForm(ctx context.Context, in *tg.TLAccountGetAuthorizationForm) (*tg.AccountAuthorizationForm, error) {
	return m.rpc.AccountGetAuthorizationForm(ctx, in)
}

// AccountAcceptAuthorization
// account.acceptAuthorization#f3ed4c73 bot_id:long scope:string public_key:string value_hashes:Vector<SecureValueHash> credentials:SecureCredentialsEncrypted = Bool;
func (m *defaultPassportClient) AccountAcceptAuthorization(ctx context.Context, in *tg.TLAccountAcceptAuthorization) (*tg.Bool, error) {
	return m.rpc.AccountAcceptAuthorization(ctx, in)
}

// AccountSendVerifyPhoneCode
// account.sendVerifyPhoneCode#a5a356f9 phone_number:string settings:CodeSettings = auth.SentCode;
func (m *defaultPassportClient) AccountSendVerifyPhoneCode(ctx context.Context, in *tg.TLAccountSendVerifyPhoneCode) (*tg.AuthSentCode, error) {
	return m.rpc.AccountSendVerifyPhoneCode(ctx, in)
}

// AccountVerifyPhone
// account.verifyPhone#4dd3a7f6 phone_number:string phone_code_hash:string phone_code:string = Bool;
func (m *defaultPassportClient) AccountVerifyPhone(ctx context.Context, in *tg.TLAccountVerifyPhone) (*tg.Bool, error) {
	return m.rpc.AccountVerifyPhone(ctx, in)
}

// UsersSetSecureValueErrors
// users.setSecureValueErrors#90c894b5 id:InputUser errors:Vector<SecureValueError> = Bool;
func (m *defaultPassportClient) UsersSetSecureValueErrors(ctx context.Context, in *tg.TLUsersSetSecureValueErrors) (*tg.Bool, error) {
	return m.rpc.UsersSetSecureValueErrors(ctx, in)
}

// HelpGetPassportConfig
// help.getPassportConfig#c661ad08 hash:int = help.PassportConfig;
func (m *defaultPassportClient) HelpGetPassportConfig(ctx context.Context, in *tg.TLHelpGetPassportConfig) (*tg.HelpPassportConfig, error) {
	return m.rpc.HelpGetPassportConfig(ctx, in)
}
