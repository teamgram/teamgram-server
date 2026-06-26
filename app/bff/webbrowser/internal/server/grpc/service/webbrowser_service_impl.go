/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026 The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/webbrowser/internal/core"
)

// AccountGetWebBrowserSettings
// account.getWebBrowserSettings#56655768 hash:long = account.WebBrowserSettings;
func (s *Service) AccountGetWebBrowserSettings(ctx context.Context, request *mtproto.TLAccountGetWebBrowserSettings) (*mtproto.Account_WebBrowserSettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.getWebBrowserSettings - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountGetWebBrowserSettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.getWebBrowserSettings - reply: {%s}", r)
	return r, err
}

// AccountUpdateWebBrowserSettings
// account.updateWebBrowserSettings#9adf82fe flags:# open_external_browser:flags.0?true display_close_button:flags.1?true = account.WebBrowserSettings;
func (s *Service) AccountUpdateWebBrowserSettings(ctx context.Context, request *mtproto.TLAccountUpdateWebBrowserSettings) (*mtproto.Account_WebBrowserSettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.updateWebBrowserSettings - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountUpdateWebBrowserSettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.updateWebBrowserSettings - reply: {%s}", r)
	return r, err
}

// AccountToggleWebBrowserSettingsException60ED4229
// account.toggleWebBrowserSettingsException#60ed4229 flags:# delete:flags.1?true open_external_browser:flags.0?Bool url:string = Updates;
func (s *Service) AccountToggleWebBrowserSettingsException60ED4229(ctx context.Context, request *mtproto.TLAccountToggleWebBrowserSettingsException60ED4229) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.toggleWebBrowserSettingsException60ED4229 - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountToggleWebBrowserSettingsException60ED4229(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.toggleWebBrowserSettingsException60ED4229 - reply: {%s}", r)
	return r, err
}

// AccountDeleteWebBrowserSettingsExceptions
// account.deleteWebBrowserSettingsExceptions#86a0765d = account.WebBrowserSettings;
func (s *Service) AccountDeleteWebBrowserSettingsExceptions(ctx context.Context, request *mtproto.TLAccountDeleteWebBrowserSettingsExceptions) (*mtproto.Account_WebBrowserSettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.deleteWebBrowserSettingsExceptions - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountDeleteWebBrowserSettingsExceptions(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.deleteWebBrowserSettingsExceptions - reply: {%s}", r)
	return r, err
}

// AccountToggleWebBrowserSettingsException2D0A0571
// account.toggleWebBrowserSettingsException#2d0a0571 flags:# delete:flags.1?true open_external_browser:flags.0?Bool url:string = Bool;
func (s *Service) AccountToggleWebBrowserSettingsException2D0A0571(ctx context.Context, request *mtproto.TLAccountToggleWebBrowserSettingsException2D0A0571) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.toggleWebBrowserSettingsException2D0A0571 - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountToggleWebBrowserSettingsException2D0A0571(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.toggleWebBrowserSettingsException2D0A0571 - reply: {%s}", r)
	return r, err
}
