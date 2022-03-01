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
	"github.com/teamgram/teamgram-server/app/bff/themes/internal/core"
)

// AccountUploadTheme
// account.uploadTheme#1c3db333 flags:# file:InputFile thumb:flags.0?InputFile file_name:string mime_type:string = Document;
func (s *Service) AccountUploadTheme(ctx context.Context, request *mtproto.TLAccountUploadTheme) (*mtproto.Document, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.uploadTheme - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountUploadTheme(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.uploadTheme - reply: %s", r.DebugString())
	return r, err
}

// AccountCreateTheme
// account.createTheme#652e4400 flags:# slug:string title:string document:flags.2?InputDocument settings:flags.3?Vector<InputThemeSettings> = Theme;
func (s *Service) AccountCreateTheme(ctx context.Context, request *mtproto.TLAccountCreateTheme) (*mtproto.Theme, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.createTheme - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountCreateTheme(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.createTheme - reply: %s", r.DebugString())
	return r, err
}

// AccountUpdateTheme
// account.updateTheme#2bf40ccc flags:# format:string theme:InputTheme slug:flags.0?string title:flags.1?string document:flags.2?InputDocument settings:flags.3?Vector<InputThemeSettings> = Theme;
func (s *Service) AccountUpdateTheme(ctx context.Context, request *mtproto.TLAccountUpdateTheme) (*mtproto.Theme, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.updateTheme - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountUpdateTheme(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.updateTheme - reply: %s", r.DebugString())
	return r, err
}

// AccountSaveTheme
// account.saveTheme#f257106c theme:InputTheme unsave:Bool = Bool;
func (s *Service) AccountSaveTheme(ctx context.Context, request *mtproto.TLAccountSaveTheme) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.saveTheme - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountSaveTheme(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.saveTheme - reply: %s", r.DebugString())
	return r, err
}

// AccountInstallTheme
// account.installTheme#c727bb3b flags:# dark:flags.0?true theme:flags.1?InputTheme format:flags.2?string base_theme:flags.3?BaseTheme = Bool;
func (s *Service) AccountInstallTheme(ctx context.Context, request *mtproto.TLAccountInstallTheme) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.installTheme - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountInstallTheme(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.installTheme - reply: %s", r.DebugString())
	return r, err
}

// AccountGetTheme
// account.getTheme#8d9d742b format:string theme:InputTheme document_id:long = Theme;
func (s *Service) AccountGetTheme(ctx context.Context, request *mtproto.TLAccountGetTheme) (*mtproto.Theme, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.getTheme - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountGetTheme(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.getTheme - reply: %s", r.DebugString())
	return r, err
}

// AccountGetThemes
// account.getThemes#7206e458 format:string hash:long = account.Themes;
func (s *Service) AccountGetThemes(ctx context.Context, request *mtproto.TLAccountGetThemes) (*mtproto.Account_Themes, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.getThemes - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountGetThemes(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.getThemes - reply: %s", r.DebugString())
	return r, err
}

// AccountGetChatThemesD638DE89
// account.getChatThemes#d638de89 hash:long = account.Themes;
func (s *Service) AccountGetChatThemesD638DE89(ctx context.Context, request *mtproto.TLAccountGetChatThemesD638DE89) (*mtproto.Account_Themes, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.getChatThemesD638DE89 - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountGetChatThemesD638DE89(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.getChatThemesD638DE89 - reply: %s", r.DebugString())
	return r, err
}

// MessagesSetChatTheme
// messages.setChatTheme#e63be13f peer:InputPeer emoticon:string = Updates;
func (s *Service) MessagesSetChatTheme(ctx context.Context, request *mtproto.TLMessagesSetChatTheme) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.setChatTheme - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesSetChatTheme(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.setChatTheme - reply: %s", r.DebugString())
	return r, err
}

// AccountGetChatThemesD6D71D7B
// account.getChatThemes#d6d71d7b hash:int = account.ChatThemes;
func (s *Service) AccountGetChatThemesD6D71D7B(ctx context.Context, request *mtproto.TLAccountGetChatThemesD6D71D7B) (*mtproto.Account_ChatThemes, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.getChatThemesD6D71D7B - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountGetChatThemesD6D71D7B(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.getChatThemesD6D71D7B - reply: %s", r.DebugString())
	return r, err
}
