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

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/bff/configuration/internal/core"
)

// HelpGetConfig
// help.getConfig#c4f9186b = Config;
func (s *Service) HelpGetConfig(ctx context.Context, request *tg.TLHelpGetConfig) (*tg.Config, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("help.getConfig - metadata: %s, request: %s", c.MD, request)

	r, err := c.HelpGetConfig(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("help.getConfig - reply: %s", r)
	return r, err
}

// HelpGetNearestDc
// help.getNearestDc#1fb33026 = NearestDc;
func (s *Service) HelpGetNearestDc(ctx context.Context, request *tg.TLHelpGetNearestDc) (*tg.NearestDc, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("help.getNearestDc - metadata: %s, request: %s", c.MD, request)

	r, err := c.HelpGetNearestDc(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("help.getNearestDc - reply: %s", r)
	return r, err
}

// HelpGetAppUpdate
// help.getAppUpdate#522d5a7d source:string = help.AppUpdate;
func (s *Service) HelpGetAppUpdate(ctx context.Context, request *tg.TLHelpGetAppUpdate) (*tg.HelpAppUpdate, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("help.getAppUpdate - metadata: %s, request: %s", c.MD, request)

	r, err := c.HelpGetAppUpdate(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("help.getAppUpdate - reply: %s", r)
	return r, err
}

// HelpGetInviteText
// help.getInviteText#4d392343 = help.InviteText;
func (s *Service) HelpGetInviteText(ctx context.Context, request *tg.TLHelpGetInviteText) (*tg.HelpInviteText, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("help.getInviteText - metadata: %s, request: %s", c.MD, request)

	r, err := c.HelpGetInviteText(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("help.getInviteText - reply: %s", r)
	return r, err
}

// HelpGetSupport
// help.getSupport#9cdf08cd = help.Support;
func (s *Service) HelpGetSupport(ctx context.Context, request *tg.TLHelpGetSupport) (*tg.HelpSupport, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("help.getSupport - metadata: %s, request: %s", c.MD, request)

	r, err := c.HelpGetSupport(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("help.getSupport - reply: %s", r)
	return r, err
}

// HelpGetAppConfig
// help.getAppConfig#61e3f854 hash:int = help.AppConfig;
func (s *Service) HelpGetAppConfig(ctx context.Context, request *tg.TLHelpGetAppConfig) (*tg.HelpAppConfig, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("help.getAppConfig - metadata: %s, request: %s", c.MD, request)

	r, err := c.HelpGetAppConfig(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("help.getAppConfig - reply: %s", r)
	return r, err
}

// HelpGetSupportName
// help.getSupportName#d360e72c = help.SupportName;
func (s *Service) HelpGetSupportName(ctx context.Context, request *tg.TLHelpGetSupportName) (*tg.HelpSupportName, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("help.getSupportName - metadata: %s, request: %s", c.MD, request)

	r, err := c.HelpGetSupportName(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("help.getSupportName - reply: %s", r)
	return r, err
}

// HelpDismissSuggestion
// help.dismissSuggestion#f50dbaa1 peer:InputPeer suggestion:string = Bool;
func (s *Service) HelpDismissSuggestion(ctx context.Context, request *tg.TLHelpDismissSuggestion) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("help.dismissSuggestion - metadata: %s, request: %s", c.MD, request)

	r, err := c.HelpDismissSuggestion(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("help.dismissSuggestion - reply: %s", r)
	return r, err
}

// HelpGetCountriesList
// help.getCountriesList#735787a8 lang_code:string hash:int = help.CountriesList;
func (s *Service) HelpGetCountriesList(ctx context.Context, request *tg.TLHelpGetCountriesList) (*tg.HelpCountriesList, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("help.getCountriesList - metadata: %s, request: %s", c.MD, request)

	r, err := c.HelpGetCountriesList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("help.getCountriesList - reply: %s", r)
	return r, err
}
