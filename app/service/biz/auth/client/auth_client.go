/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package auth_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/auth/auth"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type AuthClient interface {
	AuthExportLoginToken(ctx context.Context, in *auth.TLAuthExportLoginToken) (*mtproto.Auth_LoginToken, error)
	AuthImportLoginToken(ctx context.Context, in *auth.TLAuthImportLoginToken) (*mtproto.Auth_LoginToken, error)
	AuthAcceptLoginToken(ctx context.Context, in *auth.TLAuthAcceptLoginToken) (*mtproto.Authorization, error)
}

type defaultAuthClient struct {
	cli zrpc.Client
}

func NewAuthClient(cli zrpc.Client) AuthClient {
	return &defaultAuthClient{
		cli: cli,
	}
}

// AuthExportLoginToken
// auth.exportLoginToken api_id:int api_hash:string except_ids:Vector<long> = auth.LoginToken;
func (m *defaultAuthClient) AuthExportLoginToken(ctx context.Context, in *auth.TLAuthExportLoginToken) (*mtproto.Auth_LoginToken, error) {
	client := auth.NewRPCAuthClient(m.cli.Conn())
	return client.AuthExportLoginToken(ctx, in)
}

// AuthImportLoginToken
// auth.importLoginToken token:bytes = auth.LoginToken;
func (m *defaultAuthClient) AuthImportLoginToken(ctx context.Context, in *auth.TLAuthImportLoginToken) (*mtproto.Auth_LoginToken, error) {
	client := auth.NewRPCAuthClient(m.cli.Conn())
	return client.AuthImportLoginToken(ctx, in)
}

// AuthAcceptLoginToken
// auth.acceptLoginToken token:bytes = Authorization;
func (m *defaultAuthClient) AuthAcceptLoginToken(ctx context.Context, in *auth.TLAuthAcceptLoginToken) (*mtproto.Authorization, error) {
	client := auth.NewRPCAuthClient(m.cli.Conn())
	return client.AuthAcceptLoginToken(ctx, in)
}
