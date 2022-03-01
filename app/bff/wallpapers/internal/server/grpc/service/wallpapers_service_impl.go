/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/wallpapers/internal/core"
)

// AccountGetWallPapers
// account.getWallPapers#7967d36 hash:long = account.WallPapers;
func (s *Service) AccountGetWallPapers(ctx context.Context, request *mtproto.TLAccountGetWallPapers) (*mtproto.Account_WallPapers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.getWallPapers - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountGetWallPapers(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.getWallPapers - reply: %s", r.DebugString())
	return r, err
}

// AccountGetWallPaper
// account.getWallPaper#fc8ddbea wallpaper:InputWallPaper = WallPaper;
func (s *Service) AccountGetWallPaper(ctx context.Context, request *mtproto.TLAccountGetWallPaper) (*mtproto.WallPaper, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.getWallPaper - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountGetWallPaper(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.getWallPaper - reply: %s", r.DebugString())
	return r, err
}

// AccountUploadWallPaper
// account.uploadWallPaper#dd853661 file:InputFile mime_type:string settings:WallPaperSettings = WallPaper;
func (s *Service) AccountUploadWallPaper(ctx context.Context, request *mtproto.TLAccountUploadWallPaper) (*mtproto.WallPaper, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.uploadWallPaper - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountUploadWallPaper(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.uploadWallPaper - reply: %s", r.DebugString())
	return r, err
}

// AccountSaveWallPaper
// account.saveWallPaper#6c5a5b37 wallpaper:InputWallPaper unsave:Bool settings:WallPaperSettings = Bool;
func (s *Service) AccountSaveWallPaper(ctx context.Context, request *mtproto.TLAccountSaveWallPaper) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.saveWallPaper - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountSaveWallPaper(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.saveWallPaper - reply: %s", r.DebugString())
	return r, err
}

// AccountInstallWallPaper
// account.installWallPaper#feed5769 wallpaper:InputWallPaper settings:WallPaperSettings = Bool;
func (s *Service) AccountInstallWallPaper(ctx context.Context, request *mtproto.TLAccountInstallWallPaper) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.installWallPaper - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountInstallWallPaper(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.installWallPaper - reply: %s", r.DebugString())
	return r, err
}

// AccountResetWallPapers
// account.resetWallPapers#bb3b9804 = Bool;
func (s *Service) AccountResetWallPapers(ctx context.Context, request *mtproto.TLAccountResetWallPapers) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.resetWallPapers - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountResetWallPapers(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.resetWallPapers - reply: %s", r.DebugString())
	return r, err
}

// AccountGetMultiWallPapers
// account.getMultiWallPapers#65ad71dc wallpapers:Vector<InputWallPaper> = Vector<WallPaper>;
func (s *Service) AccountGetMultiWallPapers(ctx context.Context, request *mtproto.TLAccountGetMultiWallPapers) (*mtproto.Vector_WallPaper, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.getMultiWallPapers - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountGetMultiWallPapers(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.getMultiWallPapers - reply: %s", r.DebugString())
	return r, err
}
