/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package wallpapers_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type WallpapersClient interface {
	AccountGetWallPapers(ctx context.Context, in *mtproto.TLAccountGetWallPapers) (*mtproto.Account_WallPapers, error)
	AccountGetWallPaper(ctx context.Context, in *mtproto.TLAccountGetWallPaper) (*mtproto.WallPaper, error)
	AccountUploadWallPaper(ctx context.Context, in *mtproto.TLAccountUploadWallPaper) (*mtproto.WallPaper, error)
	AccountSaveWallPaper(ctx context.Context, in *mtproto.TLAccountSaveWallPaper) (*mtproto.Bool, error)
	AccountInstallWallPaper(ctx context.Context, in *mtproto.TLAccountInstallWallPaper) (*mtproto.Bool, error)
	AccountResetWallPapers(ctx context.Context, in *mtproto.TLAccountResetWallPapers) (*mtproto.Bool, error)
	AccountGetMultiWallPapers(ctx context.Context, in *mtproto.TLAccountGetMultiWallPapers) (*mtproto.Vector_WallPaper, error)
}

type defaultWallpapersClient struct {
	cli zrpc.Client
}

func NewWallpapersClient(cli zrpc.Client) WallpapersClient {
	return &defaultWallpapersClient{
		cli: cli,
	}
}

// AccountGetWallPapers
// account.getWallPapers#7967d36 hash:long = account.WallPapers;
func (m *defaultWallpapersClient) AccountGetWallPapers(ctx context.Context, in *mtproto.TLAccountGetWallPapers) (*mtproto.Account_WallPapers, error) {
	client := mtproto.NewRPCWallpapersClient(m.cli.Conn())
	return client.AccountGetWallPapers(ctx, in)
}

// AccountGetWallPaper
// account.getWallPaper#fc8ddbea wallpaper:InputWallPaper = WallPaper;
func (m *defaultWallpapersClient) AccountGetWallPaper(ctx context.Context, in *mtproto.TLAccountGetWallPaper) (*mtproto.WallPaper, error) {
	client := mtproto.NewRPCWallpapersClient(m.cli.Conn())
	return client.AccountGetWallPaper(ctx, in)
}

// AccountUploadWallPaper
// account.uploadWallPaper#dd853661 file:InputFile mime_type:string settings:WallPaperSettings = WallPaper;
func (m *defaultWallpapersClient) AccountUploadWallPaper(ctx context.Context, in *mtproto.TLAccountUploadWallPaper) (*mtproto.WallPaper, error) {
	client := mtproto.NewRPCWallpapersClient(m.cli.Conn())
	return client.AccountUploadWallPaper(ctx, in)
}

// AccountSaveWallPaper
// account.saveWallPaper#6c5a5b37 wallpaper:InputWallPaper unsave:Bool settings:WallPaperSettings = Bool;
func (m *defaultWallpapersClient) AccountSaveWallPaper(ctx context.Context, in *mtproto.TLAccountSaveWallPaper) (*mtproto.Bool, error) {
	client := mtproto.NewRPCWallpapersClient(m.cli.Conn())
	return client.AccountSaveWallPaper(ctx, in)
}

// AccountInstallWallPaper
// account.installWallPaper#feed5769 wallpaper:InputWallPaper settings:WallPaperSettings = Bool;
func (m *defaultWallpapersClient) AccountInstallWallPaper(ctx context.Context, in *mtproto.TLAccountInstallWallPaper) (*mtproto.Bool, error) {
	client := mtproto.NewRPCWallpapersClient(m.cli.Conn())
	return client.AccountInstallWallPaper(ctx, in)
}

// AccountResetWallPapers
// account.resetWallPapers#bb3b9804 = Bool;
func (m *defaultWallpapersClient) AccountResetWallPapers(ctx context.Context, in *mtproto.TLAccountResetWallPapers) (*mtproto.Bool, error) {
	client := mtproto.NewRPCWallpapersClient(m.cli.Conn())
	return client.AccountResetWallPapers(ctx, in)
}

// AccountGetMultiWallPapers
// account.getMultiWallPapers#65ad71dc wallpapers:Vector<InputWallPaper> = Vector<WallPaper>;
func (m *defaultWallpapersClient) AccountGetMultiWallPapers(ctx context.Context, in *mtproto.TLAccountGetMultiWallPapers) (*mtproto.Vector_WallPaper, error) {
	client := mtproto.NewRPCWallpapersClient(m.cli.Conn())
	return client.AccountGetMultiWallPapers(ctx, in)
}
