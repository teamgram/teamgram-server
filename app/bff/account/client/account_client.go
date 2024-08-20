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

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type AccountClient interface {
	AccountDeleteAccount(ctx context.Context, in *mtproto.TLAccountDeleteAccount) (*mtproto.Bool, error)
	AccountGetAccountTTL(ctx context.Context, in *mtproto.TLAccountGetAccountTTL) (*mtproto.AccountDaysTTL, error)
	AccountSetAccountTTL(ctx context.Context, in *mtproto.TLAccountSetAccountTTL) (*mtproto.Bool, error)
	AccountSendChangePhoneCode(ctx context.Context, in *mtproto.TLAccountSendChangePhoneCode) (*mtproto.Auth_SentCode, error)
	AccountChangePhone(ctx context.Context, in *mtproto.TLAccountChangePhone) (*mtproto.User, error)
	AccountResetAuthorization(ctx context.Context, in *mtproto.TLAccountResetAuthorization) (*mtproto.Bool, error)
	AccountSendConfirmPhoneCode(ctx context.Context, in *mtproto.TLAccountSendConfirmPhoneCode) (*mtproto.Auth_SentCode, error)
	AccountConfirmPhone(ctx context.Context, in *mtproto.TLAccountConfirmPhone) (*mtproto.Bool, error)
}

type defaultAccountClient struct {
	cli zrpc.Client
}

func NewAccountClient(cli zrpc.Client) AccountClient {
	return &defaultAccountClient{
		cli: cli,
	}
}

// AccountDeleteAccount
// account.deleteAccount#a2c0cf74 flags:# reason:string password:flags.0?InputCheckPasswordSRP = Bool;
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
