/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package qrcode_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type QrCodeClient interface {
	AuthExportLoginToken(ctx context.Context, in *mtproto.TLAuthExportLoginToken) (*mtproto.Auth_LoginToken, error)
	AuthImportLoginToken(ctx context.Context, in *mtproto.TLAuthImportLoginToken) (*mtproto.Auth_LoginToken, error)
	AuthAcceptLoginToken(ctx context.Context, in *mtproto.TLAuthAcceptLoginToken) (*mtproto.Authorization, error)
}

type defaultQrCodeClient struct {
	cli zrpc.Client
}

func NewQrCodeClient(cli zrpc.Client) QrCodeClient {
	return &defaultQrCodeClient{
		cli: cli,
	}
}

// AuthExportLoginToken
// auth.exportLoginToken#b7e085fe api_id:int api_hash:string except_ids:Vector<long> = auth.LoginToken;
func (m *defaultQrCodeClient) AuthExportLoginToken(ctx context.Context, in *mtproto.TLAuthExportLoginToken) (*mtproto.Auth_LoginToken, error) {
	client := mtproto.NewRPCQrCodeClient(m.cli.Conn())
	return client.AuthExportLoginToken(ctx, in)
}

// AuthImportLoginToken
// auth.importLoginToken#95ac5ce4 token:bytes = auth.LoginToken;
func (m *defaultQrCodeClient) AuthImportLoginToken(ctx context.Context, in *mtproto.TLAuthImportLoginToken) (*mtproto.Auth_LoginToken, error) {
	client := mtproto.NewRPCQrCodeClient(m.cli.Conn())
	return client.AuthImportLoginToken(ctx, in)
}

// AuthAcceptLoginToken
// auth.acceptLoginToken#e894ad4d token:bytes = Authorization;
func (m *defaultQrCodeClient) AuthAcceptLoginToken(ctx context.Context, in *mtproto.TLAuthAcceptLoginToken) (*mtproto.Authorization, error) {
	client := mtproto.NewRPCQrCodeClient(m.cli.Conn())
	return client.AuthAcceptLoginToken(ctx, in)
}
