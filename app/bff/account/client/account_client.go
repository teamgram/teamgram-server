/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package accountclient

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/bff/account/account/accountservice"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
)

type AccountClient interface {
	AccountDeleteAccount(ctx context.Context, in *tg.TLAccountDeleteAccount) (*tg.Bool, error)
	AccountGetAccountTTL(ctx context.Context, in *tg.TLAccountGetAccountTTL) (*tg.AccountDaysTTL, error)
	AccountSetAccountTTL(ctx context.Context, in *tg.TLAccountSetAccountTTL) (*tg.Bool, error)
	AccountSendChangePhoneCode(ctx context.Context, in *tg.TLAccountSendChangePhoneCode) (*tg.AuthSentCode, error)
	AccountChangePhone(ctx context.Context, in *tg.TLAccountChangePhone) (*tg.User, error)
	AccountResetAuthorization(ctx context.Context, in *tg.TLAccountResetAuthorization) (*tg.Bool, error)
	AccountSendConfirmPhoneCode(ctx context.Context, in *tg.TLAccountSendConfirmPhoneCode) (*tg.AuthSentCode, error)
	AccountConfirmPhone(ctx context.Context, in *tg.TLAccountConfirmPhone) (*tg.Bool, error)
}

type defaultAccountClient struct {
	cli client.Client
	rpc accountservice.Client
}

func NewAccountClient(cli client.Client) AccountClient {
	return &defaultAccountClient{
		cli: cli,
		rpc: accountservice.NewRPCAccountClient(cli),
	}
}

// AccountDeleteAccount
// account.deleteAccount#a2c0cf74 flags:# reason:string password:flags.0?InputCheckPasswordSRP = Bool;
func (m *defaultAccountClient) AccountDeleteAccount(ctx context.Context, in *tg.TLAccountDeleteAccount) (*tg.Bool, error) {
	return m.rpc.AccountDeleteAccount(ctx, in)
}

// AccountGetAccountTTL
// account.getAccountTTL#8fc711d = AccountDaysTTL;
func (m *defaultAccountClient) AccountGetAccountTTL(ctx context.Context, in *tg.TLAccountGetAccountTTL) (*tg.AccountDaysTTL, error) {
	return m.rpc.AccountGetAccountTTL(ctx, in)
}

// AccountSetAccountTTL
// account.setAccountTTL#2442485e ttl:AccountDaysTTL = Bool;
func (m *defaultAccountClient) AccountSetAccountTTL(ctx context.Context, in *tg.TLAccountSetAccountTTL) (*tg.Bool, error) {
	return m.rpc.AccountSetAccountTTL(ctx, in)
}

// AccountSendChangePhoneCode
// account.sendChangePhoneCode#82574ae5 phone_number:string settings:CodeSettings = auth.SentCode;
func (m *defaultAccountClient) AccountSendChangePhoneCode(ctx context.Context, in *tg.TLAccountSendChangePhoneCode) (*tg.AuthSentCode, error) {
	return m.rpc.AccountSendChangePhoneCode(ctx, in)
}

// AccountChangePhone
// account.changePhone#70c32edb phone_number:string phone_code_hash:string phone_code:string = User;
func (m *defaultAccountClient) AccountChangePhone(ctx context.Context, in *tg.TLAccountChangePhone) (*tg.User, error) {
	return m.rpc.AccountChangePhone(ctx, in)
}

// AccountResetAuthorization
// account.resetAuthorization#df77f3bc hash:long = Bool;
func (m *defaultAccountClient) AccountResetAuthorization(ctx context.Context, in *tg.TLAccountResetAuthorization) (*tg.Bool, error) {
	return m.rpc.AccountResetAuthorization(ctx, in)
}

// AccountSendConfirmPhoneCode
// account.sendConfirmPhoneCode#1b3faa88 hash:string settings:CodeSettings = auth.SentCode;
func (m *defaultAccountClient) AccountSendConfirmPhoneCode(ctx context.Context, in *tg.TLAccountSendConfirmPhoneCode) (*tg.AuthSentCode, error) {
	return m.rpc.AccountSendConfirmPhoneCode(ctx, in)
}

// AccountConfirmPhone
// account.confirmPhone#5f2178c3 phone_code_hash:string phone_code:string = Bool;
func (m *defaultAccountClient) AccountConfirmPhone(ctx context.Context, in *tg.TLAccountConfirmPhone) (*tg.Bool, error) {
	return m.rpc.AccountConfirmPhone(ctx, in)
}
