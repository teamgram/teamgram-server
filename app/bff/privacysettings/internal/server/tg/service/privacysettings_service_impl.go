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
	"github.com/teamgram/teamgram-server/v2/app/bff/privacysettings/internal/core"
)

// AccountGetPrivacy
// account.getPrivacy#dadbc950 key:InputPrivacyKey = account.PrivacyRules;
func (s *Service) AccountGetPrivacy(ctx context.Context, request *tg.TLAccountGetPrivacy) (*tg.AccountPrivacyRules, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.getPrivacy - metadata: {}, request: {%v}", request)

	r, err := c.AccountGetPrivacy(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.getPrivacy - reply: {%v}", r)
	return r, err
}

// AccountSetPrivacy
// account.setPrivacy#c9f81ce8 key:InputPrivacyKey rules:Vector<InputPrivacyRule> = account.PrivacyRules;
func (s *Service) AccountSetPrivacy(ctx context.Context, request *tg.TLAccountSetPrivacy) (*tg.AccountPrivacyRules, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.setPrivacy - metadata: {}, request: {%v}", request)

	r, err := c.AccountSetPrivacy(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.setPrivacy - reply: {%v}", r)
	return r, err
}

// AccountGetGlobalPrivacySettings
// account.getGlobalPrivacySettings#eb2b4cf6 = GlobalPrivacySettings;
func (s *Service) AccountGetGlobalPrivacySettings(ctx context.Context, request *tg.TLAccountGetGlobalPrivacySettings) (*tg.GlobalPrivacySettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.getGlobalPrivacySettings - metadata: {}, request: {%v}", request)

	r, err := c.AccountGetGlobalPrivacySettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.getGlobalPrivacySettings - reply: {%v}", r)
	return r, err
}

// AccountSetGlobalPrivacySettings
// account.setGlobalPrivacySettings#1edaaac2 settings:GlobalPrivacySettings = GlobalPrivacySettings;
func (s *Service) AccountSetGlobalPrivacySettings(ctx context.Context, request *tg.TLAccountSetGlobalPrivacySettings) (*tg.GlobalPrivacySettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.setGlobalPrivacySettings - metadata: {}, request: {%v}", request)

	r, err := c.AccountSetGlobalPrivacySettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.setGlobalPrivacySettings - reply: {%v}", r)
	return r, err
}

// UsersGetIsPremiumRequiredToContact
// users.getIsPremiumRequiredToContact#a622aa10 id:Vector<InputUser> = Vector<Bool>;
func (s *Service) UsersGetIsPremiumRequiredToContact(ctx context.Context, request *tg.TLUsersGetIsPremiumRequiredToContact) (*tg.VectorBool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("users.getIsPremiumRequiredToContact - metadata: {}, request: {%v}", request)

	r, err := c.UsersGetIsPremiumRequiredToContact(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("users.getIsPremiumRequiredToContact - reply: {%v}", r)
	return r, err
}

// MessagesSetDefaultHistoryTTL
// messages.setDefaultHistoryTTL#9eb51445 period:int = Bool;
func (s *Service) MessagesSetDefaultHistoryTTL(ctx context.Context, request *tg.TLMessagesSetDefaultHistoryTTL) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.setDefaultHistoryTTL - metadata: {}, request: {%v}", request)

	r, err := c.MessagesSetDefaultHistoryTTL(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.setDefaultHistoryTTL - reply: {%v}", r)
	return r, err
}

// MessagesGetDefaultHistoryTTL
// messages.getDefaultHistoryTTL#658b7188 = DefaultHistoryTTL;
func (s *Service) MessagesGetDefaultHistoryTTL(ctx context.Context, request *tg.TLMessagesGetDefaultHistoryTTL) (*tg.DefaultHistoryTTL, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getDefaultHistoryTTL - metadata: {}, request: {%v}", request)

	r, err := c.MessagesGetDefaultHistoryTTL(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getDefaultHistoryTTL - reply: {%v}", r)
	return r, err
}
