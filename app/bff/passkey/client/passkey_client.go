/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2025 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package passkeyclient

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type PasskeyClient interface {
	AuthInitPasskeyLogin(ctx context.Context, in *mtproto.TLAuthInitPasskeyLogin) (*mtproto.Auth_PasskeyLoginOptions, error)
	AuthFinishPasskeyLogin(ctx context.Context, in *mtproto.TLAuthFinishPasskeyLogin) (*mtproto.Auth_Authorization, error)
	AccountInitPasskeyRegistration(ctx context.Context, in *mtproto.TLAccountInitPasskeyRegistration) (*mtproto.Account_PasskeyRegistrationOptions, error)
	AccountRegisterPasskey(ctx context.Context, in *mtproto.TLAccountRegisterPasskey) (*mtproto.Passkey, error)
	AccountGetPasskeys(ctx context.Context, in *mtproto.TLAccountGetPasskeys) (*mtproto.Account_Passkeys, error)
	AccountDeletePasskey(ctx context.Context, in *mtproto.TLAccountDeletePasskey) (*mtproto.Bool, error)
}

type defaultPasskeyClient struct {
	cli zrpc.Client
}

func NewPasskeyClient(cli zrpc.Client) PasskeyClient {
	return &defaultPasskeyClient{
		cli: cli,
	}
}

// AuthInitPasskeyLogin
// auth.initPasskeyLogin#518ad0b7 api_id:int api_hash:string = auth.PasskeyLoginOptions;
func (m *defaultPasskeyClient) AuthInitPasskeyLogin(ctx context.Context, in *mtproto.TLAuthInitPasskeyLogin) (*mtproto.Auth_PasskeyLoginOptions, error) {
	client := mtproto.NewRPCPasskeyClient(m.cli.Conn())
	return client.AuthInitPasskeyLogin(ctx, in)
}

// AuthFinishPasskeyLogin
// auth.finishPasskeyLogin#9857ad07 flags:# credential:InputPasskeyCredential from_dc_id:flags.0?int from_auth_key_id:flags.0?long = auth.Authorization;
func (m *defaultPasskeyClient) AuthFinishPasskeyLogin(ctx context.Context, in *mtproto.TLAuthFinishPasskeyLogin) (*mtproto.Auth_Authorization, error) {
	client := mtproto.NewRPCPasskeyClient(m.cli.Conn())
	return client.AuthFinishPasskeyLogin(ctx, in)
}

// AccountInitPasskeyRegistration
// account.initPasskeyRegistration#429547e8 = account.PasskeyRegistrationOptions;
func (m *defaultPasskeyClient) AccountInitPasskeyRegistration(ctx context.Context, in *mtproto.TLAccountInitPasskeyRegistration) (*mtproto.Account_PasskeyRegistrationOptions, error) {
	client := mtproto.NewRPCPasskeyClient(m.cli.Conn())
	return client.AccountInitPasskeyRegistration(ctx, in)
}

// AccountRegisterPasskey
// account.registerPasskey#55b41fd6 credential:InputPasskeyCredential = Passkey;
func (m *defaultPasskeyClient) AccountRegisterPasskey(ctx context.Context, in *mtproto.TLAccountRegisterPasskey) (*mtproto.Passkey, error) {
	client := mtproto.NewRPCPasskeyClient(m.cli.Conn())
	return client.AccountRegisterPasskey(ctx, in)
}

// AccountGetPasskeys
// account.getPasskeys#ea1f0c52 = account.Passkeys;
func (m *defaultPasskeyClient) AccountGetPasskeys(ctx context.Context, in *mtproto.TLAccountGetPasskeys) (*mtproto.Account_Passkeys, error) {
	client := mtproto.NewRPCPasskeyClient(m.cli.Conn())
	return client.AccountGetPasskeys(ctx, in)
}

// AccountDeletePasskey
// account.deletePasskey#f5b5563f id:string = Bool;
func (m *defaultPasskeyClient) AccountDeletePasskey(ctx context.Context, in *mtproto.TLAccountDeletePasskey) (*mtproto.Bool, error) {
	client := mtproto.NewRPCPasskeyClient(m.cli.Conn())
	return client.AccountDeletePasskey(ctx, in)
}
