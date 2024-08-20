/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/configuration/internal/core"
)

// HelpGetConfig
// help.getConfig#c4f9186b = Config;
func (s *Service) HelpGetConfig(ctx context.Context, request *mtproto.TLHelpGetConfig) (*mtproto.Config, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("help.getConfig - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.HelpGetConfig(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("help.getConfig - reply: {%s}", r)
	return r, err
}

// HelpGetNearestDc
// help.getNearestDc#1fb33026 = NearestDc;
func (s *Service) HelpGetNearestDc(ctx context.Context, request *mtproto.TLHelpGetNearestDc) (*mtproto.NearestDc, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("help.getNearestDc - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.HelpGetNearestDc(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("help.getNearestDc - reply: {%s}", r)
	return r, err
}

// HelpGetAppUpdate
// help.getAppUpdate#522d5a7d source:string = help.AppUpdate;
func (s *Service) HelpGetAppUpdate(ctx context.Context, request *mtproto.TLHelpGetAppUpdate) (*mtproto.Help_AppUpdate, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("help.getAppUpdate - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.HelpGetAppUpdate(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("help.getAppUpdate - reply: {%s}", r)
	return r, err
}

// HelpGetInviteText
// help.getInviteText#4d392343 = help.InviteText;
func (s *Service) HelpGetInviteText(ctx context.Context, request *mtproto.TLHelpGetInviteText) (*mtproto.Help_InviteText, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("help.getInviteText - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.HelpGetInviteText(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("help.getInviteText - reply: {%s}", r)
	return r, err
}

// HelpGetSupport
// help.getSupport#9cdf08cd = help.Support;
func (s *Service) HelpGetSupport(ctx context.Context, request *mtproto.TLHelpGetSupport) (*mtproto.Help_Support, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("help.getSupport - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.HelpGetSupport(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("help.getSupport - reply: {%s}", r)
	return r, err
}

// HelpGetAppConfig61E3F854
// help.getAppConfig#61e3f854 hash:int = help.AppConfig;
func (s *Service) HelpGetAppConfig61E3F854(ctx context.Context, request *mtproto.TLHelpGetAppConfig61E3F854) (*mtproto.Help_AppConfig, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("help.getAppConfig61E3F854 - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.HelpGetAppConfig61E3F854(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("help.getAppConfig61E3F854 - reply: {%s}", r)
	return r, err
}

// HelpGetSupportName
// help.getSupportName#d360e72c = help.SupportName;
func (s *Service) HelpGetSupportName(ctx context.Context, request *mtproto.TLHelpGetSupportName) (*mtproto.Help_SupportName, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("help.getSupportName - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.HelpGetSupportName(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("help.getSupportName - reply: {%s}", r)
	return r, err
}

// HelpDismissSuggestion
// help.dismissSuggestion#f50dbaa1 peer:InputPeer suggestion:string = Bool;
func (s *Service) HelpDismissSuggestion(ctx context.Context, request *mtproto.TLHelpDismissSuggestion) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("help.dismissSuggestion - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.HelpDismissSuggestion(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("help.dismissSuggestion - reply: {%s}", r)
	return r, err
}

// HelpGetCountriesList
// help.getCountriesList#735787a8 lang_code:string hash:int = help.CountriesList;
func (s *Service) HelpGetCountriesList(ctx context.Context, request *mtproto.TLHelpGetCountriesList) (*mtproto.Help_CountriesList, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("help.getCountriesList - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.HelpGetCountriesList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("help.getCountriesList - reply: {%s}", r)
	return r, err
}

// HelpGetAppChangelog
// help.getAppChangelog#9010ef6f prev_app_version:string = Updates;
func (s *Service) HelpGetAppChangelog(ctx context.Context, request *mtproto.TLHelpGetAppChangelog) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("help.getAppChangelog - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.HelpGetAppChangelog(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("help.getAppChangelog - reply: {%s}", r)
	return r, err
}

// HelpGetAppConfig98914110
// help.getAppConfig#98914110 = JSONValue;
func (s *Service) HelpGetAppConfig98914110(ctx context.Context, request *mtproto.TLHelpGetAppConfig98914110) (*mtproto.JSONValue, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("help.getAppConfig98914110 - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.HelpGetAppConfig98914110(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("help.getAppConfig98914110 - reply: {%s}", r)
	return r, err
}
