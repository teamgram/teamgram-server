/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package accountclient

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/bff/account/account/accountservice"

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
}

func NewAccountClient(cli client.Client) AccountClient {
	return &defaultAccountClient{
		cli: cli,
	}
}

// AccountDeleteAccount
// account.deleteAccount#a2c0cf74 flags:# reason:string password:flags.0?InputCheckPasswordSRP = Bool;
func (m *defaultAccountClient) AccountDeleteAccount(ctx context.Context, in *tg.TLAccountDeleteAccount) (*tg.Bool, error) {
	cli := accountservice.NewRPCAccountClient(m.cli)
	return cli.AccountDeleteAccount(ctx, in)
}

// AccountGetAccountTTL
// account.getAccountTTL#8fc711d = AccountDaysTTL;
func (m *defaultAccountClient) AccountGetAccountTTL(ctx context.Context, in *tg.TLAccountGetAccountTTL) (*tg.AccountDaysTTL, error) {
	cli := accountservice.NewRPCAccountClient(m.cli)
	return cli.AccountGetAccountTTL(ctx, in)
}

// AccountSetAccountTTL
// account.setAccountTTL#2442485e ttl:AccountDaysTTL = Bool;
func (m *defaultAccountClient) AccountSetAccountTTL(ctx context.Context, in *tg.TLAccountSetAccountTTL) (*tg.Bool, error) {
	cli := accountservice.NewRPCAccountClient(m.cli)
	return cli.AccountSetAccountTTL(ctx, in)
}

// AccountSendChangePhoneCode
// account.sendChangePhoneCode#82574ae5 phone_number:string settings:CodeSettings = auth.SentCode;
func (m *defaultAccountClient) AccountSendChangePhoneCode(ctx context.Context, in *tg.TLAccountSendChangePhoneCode) (*tg.AuthSentCode, error) {
	cli := accountservice.NewRPCAccountClient(m.cli)
	return cli.AccountSendChangePhoneCode(ctx, in)
}

// AccountChangePhone
// account.changePhone#70c32edb phone_number:string phone_code_hash:string phone_code:string = User;
func (m *defaultAccountClient) AccountChangePhone(ctx context.Context, in *tg.TLAccountChangePhone) (*tg.User, error) {
	cli := accountservice.NewRPCAccountClient(m.cli)
	return cli.AccountChangePhone(ctx, in)
}

// AccountResetAuthorization
// account.resetAuthorization#df77f3bc hash:long = Bool;
func (m *defaultAccountClient) AccountResetAuthorization(ctx context.Context, in *tg.TLAccountResetAuthorization) (*tg.Bool, error) {
	cli := accountservice.NewRPCAccountClient(m.cli)
	return cli.AccountResetAuthorization(ctx, in)
}

// AccountSendConfirmPhoneCode
// account.sendConfirmPhoneCode#1b3faa88 hash:string settings:CodeSettings = auth.SentCode;
func (m *defaultAccountClient) AccountSendConfirmPhoneCode(ctx context.Context, in *tg.TLAccountSendConfirmPhoneCode) (*tg.AuthSentCode, error) {
	cli := accountservice.NewRPCAccountClient(m.cli)
	return cli.AccountSendConfirmPhoneCode(ctx, in)
}

// AccountConfirmPhone
// account.confirmPhone#5f2178c3 phone_code_hash:string phone_code:string = Bool;
func (m *defaultAccountClient) AccountConfirmPhone(ctx context.Context, in *tg.TLAccountConfirmPhone) (*tg.Bool, error) {
	cli := accountservice.NewRPCAccountClient(m.cli)
	return cli.AccountConfirmPhone(ctx, in)
}
