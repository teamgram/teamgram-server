/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package twofa_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type TwoFaClient interface {
	AccountGetPassword(ctx context.Context, in *mtproto.TLAccountGetPassword) (*mtproto.Account_Password, error)
	AccountGetPasswordSettings(ctx context.Context, in *mtproto.TLAccountGetPasswordSettings) (*mtproto.Account_PasswordSettings, error)
	AccountUpdatePasswordSettings(ctx context.Context, in *mtproto.TLAccountUpdatePasswordSettings) (*mtproto.Bool, error)
	AccountConfirmPasswordEmail(ctx context.Context, in *mtproto.TLAccountConfirmPasswordEmail) (*mtproto.Bool, error)
	AccountResendPasswordEmail(ctx context.Context, in *mtproto.TLAccountResendPasswordEmail) (*mtproto.Bool, error)
	AccountCancelPasswordEmail(ctx context.Context, in *mtproto.TLAccountCancelPasswordEmail) (*mtproto.Bool, error)
	AccountDeclinePasswordReset(ctx context.Context, in *mtproto.TLAccountDeclinePasswordReset) (*mtproto.Bool, error)
}

type defaultTwoFaClient struct {
	cli zrpc.Client
}

func NewTwoFaClient(cli zrpc.Client) TwoFaClient {
	return &defaultTwoFaClient{
		cli: cli,
	}
}

// AccountGetPassword
// account.getPassword#548a30f5 = account.Password;
func (m *defaultTwoFaClient) AccountGetPassword(ctx context.Context, in *mtproto.TLAccountGetPassword) (*mtproto.Account_Password, error) {
	client := mtproto.NewRPCTwoFaClient(m.cli.Conn())
	return client.AccountGetPassword(ctx, in)
}

// AccountGetPasswordSettings
// account.getPasswordSettings#9cd4eaf9 password:InputCheckPasswordSRP = account.PasswordSettings;
func (m *defaultTwoFaClient) AccountGetPasswordSettings(ctx context.Context, in *mtproto.TLAccountGetPasswordSettings) (*mtproto.Account_PasswordSettings, error) {
	client := mtproto.NewRPCTwoFaClient(m.cli.Conn())
	return client.AccountGetPasswordSettings(ctx, in)
}

// AccountUpdatePasswordSettings
// account.updatePasswordSettings#a59b102f password:InputCheckPasswordSRP new_settings:account.PasswordInputSettings = Bool;
func (m *defaultTwoFaClient) AccountUpdatePasswordSettings(ctx context.Context, in *mtproto.TLAccountUpdatePasswordSettings) (*mtproto.Bool, error) {
	client := mtproto.NewRPCTwoFaClient(m.cli.Conn())
	return client.AccountUpdatePasswordSettings(ctx, in)
}

// AccountConfirmPasswordEmail
// account.confirmPasswordEmail#8fdf1920 code:string = Bool;
func (m *defaultTwoFaClient) AccountConfirmPasswordEmail(ctx context.Context, in *mtproto.TLAccountConfirmPasswordEmail) (*mtproto.Bool, error) {
	client := mtproto.NewRPCTwoFaClient(m.cli.Conn())
	return client.AccountConfirmPasswordEmail(ctx, in)
}

// AccountResendPasswordEmail
// account.resendPasswordEmail#7a7f2a15 = Bool;
func (m *defaultTwoFaClient) AccountResendPasswordEmail(ctx context.Context, in *mtproto.TLAccountResendPasswordEmail) (*mtproto.Bool, error) {
	client := mtproto.NewRPCTwoFaClient(m.cli.Conn())
	return client.AccountResendPasswordEmail(ctx, in)
}

// AccountCancelPasswordEmail
// account.cancelPasswordEmail#c1cbd5b6 = Bool;
func (m *defaultTwoFaClient) AccountCancelPasswordEmail(ctx context.Context, in *mtproto.TLAccountCancelPasswordEmail) (*mtproto.Bool, error) {
	client := mtproto.NewRPCTwoFaClient(m.cli.Conn())
	return client.AccountCancelPasswordEmail(ctx, in)
}

// AccountDeclinePasswordReset
// account.declinePasswordReset#4c9409f6 = Bool;
func (m *defaultTwoFaClient) AccountDeclinePasswordReset(ctx context.Context, in *mtproto.TLAccountDeclinePasswordReset) (*mtproto.Bool, error) {
	client := mtproto.NewRPCTwoFaClient(m.cli.Conn())
	return client.AccountDeclinePasswordReset(ctx, in)
}
