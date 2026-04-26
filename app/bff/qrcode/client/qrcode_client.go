/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package qrcodeclient

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/bff/qrcode/qrcode/qrcodeservice"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
)

type QrCodeClient interface {
	AuthExportLoginToken(ctx context.Context, in *tg.TLAuthExportLoginToken) (*tg.AuthLoginToken, error)
	AuthImportLoginToken(ctx context.Context, in *tg.TLAuthImportLoginToken) (*tg.AuthLoginToken, error)
	AuthAcceptLoginToken(ctx context.Context, in *tg.TLAuthAcceptLoginToken) (*tg.Authorization, error)
}

type defaultQrCodeClient struct {
	cli client.Client
	rpc qrcodeservice.Client
}

func NewQrCodeClient(cli client.Client) QrCodeClient {
	return &defaultQrCodeClient{
		cli: cli,
		rpc: qrcodeservice.NewRPCQrCodeClient(cli),
	}
}

// AuthExportLoginToken
// auth.exportLoginToken#b7e085fe api_id:int api_hash:string except_ids:Vector<long> = auth.LoginToken;
func (m *defaultQrCodeClient) AuthExportLoginToken(ctx context.Context, in *tg.TLAuthExportLoginToken) (*tg.AuthLoginToken, error) {
	return m.rpc.AuthExportLoginToken(ctx, in)
}

// AuthImportLoginToken
// auth.importLoginToken#95ac5ce4 token:bytes = auth.LoginToken;
func (m *defaultQrCodeClient) AuthImportLoginToken(ctx context.Context, in *tg.TLAuthImportLoginToken) (*tg.AuthLoginToken, error) {
	return m.rpc.AuthImportLoginToken(ctx, in)
}

// AuthAcceptLoginToken
// auth.acceptLoginToken#e894ad4d token:bytes = Authorization;
func (m *defaultQrCodeClient) AuthAcceptLoginToken(ctx context.Context, in *tg.TLAuthAcceptLoginToken) (*tg.Authorization, error) {
	return m.rpc.AuthAcceptLoginToken(ctx, in)
}
