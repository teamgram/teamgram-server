/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package qrcodeclient

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/bff/qrcode/qrcode/qrcodeservice"

	"github.com/cloudwego/kitex/client"
)

type QrCodeClient interface {
	AuthExportLoginToken(ctx context.Context, in *tg.TLAuthExportLoginToken) (*tg.AuthLoginToken, error)
	AuthImportLoginToken(ctx context.Context, in *tg.TLAuthImportLoginToken) (*tg.AuthLoginToken, error)
	AuthAcceptLoginToken(ctx context.Context, in *tg.TLAuthAcceptLoginToken) (*tg.Authorization, error)
}

type defaultQrCodeClient struct {
	cli client.Client
}

func NewQrCodeClient(cli client.Client) QrCodeClient {
	return &defaultQrCodeClient{
		cli: cli,
	}
}

// AuthExportLoginToken
// auth.exportLoginToken#b7e085fe api_id:int api_hash:string except_ids:Vector<long> = auth.LoginToken;
func (m *defaultQrCodeClient) AuthExportLoginToken(ctx context.Context, in *tg.TLAuthExportLoginToken) (*tg.AuthLoginToken, error) {
	cli := qrcodeservice.NewRPCQrCodeClient(m.cli)
	return cli.AuthExportLoginToken(ctx, in)
}

// AuthImportLoginToken
// auth.importLoginToken#95ac5ce4 token:bytes = auth.LoginToken;
func (m *defaultQrCodeClient) AuthImportLoginToken(ctx context.Context, in *tg.TLAuthImportLoginToken) (*tg.AuthLoginToken, error) {
	cli := qrcodeservice.NewRPCQrCodeClient(m.cli)
	return cli.AuthImportLoginToken(ctx, in)
}

// AuthAcceptLoginToken
// auth.acceptLoginToken#e894ad4d token:bytes = Authorization;
func (m *defaultQrCodeClient) AuthAcceptLoginToken(ctx context.Context, in *tg.TLAuthAcceptLoginToken) (*tg.Authorization, error) {
	cli := qrcodeservice.NewRPCQrCodeClient(m.cli)
	return cli.AuthAcceptLoginToken(ctx, in)
}
