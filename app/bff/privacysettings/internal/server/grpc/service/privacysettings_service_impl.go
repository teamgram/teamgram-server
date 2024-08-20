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
	"github.com/teamgram/teamgram-server/app/bff/privacysettings/internal/core"
)

// AccountGetPrivacy
// account.getPrivacy#dadbc950 key:InputPrivacyKey = account.PrivacyRules;
func (s *Service) AccountGetPrivacy(ctx context.Context, request *mtproto.TLAccountGetPrivacy) (*mtproto.Account_PrivacyRules, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.getPrivacy - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountGetPrivacy(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.getPrivacy - reply: {%s}", r)
	return r, err
}

// AccountSetPrivacy
// account.setPrivacy#c9f81ce8 key:InputPrivacyKey rules:Vector<InputPrivacyRule> = account.PrivacyRules;
func (s *Service) AccountSetPrivacy(ctx context.Context, request *mtproto.TLAccountSetPrivacy) (*mtproto.Account_PrivacyRules, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.setPrivacy - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountSetPrivacy(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.setPrivacy - reply: {%s}", r)
	return r, err
}

// AccountGetGlobalPrivacySettings
// account.getGlobalPrivacySettings#eb2b4cf6 = GlobalPrivacySettings;
func (s *Service) AccountGetGlobalPrivacySettings(ctx context.Context, request *mtproto.TLAccountGetGlobalPrivacySettings) (*mtproto.GlobalPrivacySettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.getGlobalPrivacySettings - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountGetGlobalPrivacySettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.getGlobalPrivacySettings - reply: {%s}", r)
	return r, err
}

// AccountSetGlobalPrivacySettings
// account.setGlobalPrivacySettings#1edaaac2 settings:GlobalPrivacySettings = GlobalPrivacySettings;
func (s *Service) AccountSetGlobalPrivacySettings(ctx context.Context, request *mtproto.TLAccountSetGlobalPrivacySettings) (*mtproto.GlobalPrivacySettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.setGlobalPrivacySettings - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountSetGlobalPrivacySettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.setGlobalPrivacySettings - reply: {%s}", r)
	return r, err
}

// UsersGetIsPremiumRequiredToContact
// users.getIsPremiumRequiredToContact#a622aa10 id:Vector<InputUser> = Vector<Bool>;
func (s *Service) UsersGetIsPremiumRequiredToContact(ctx context.Context, request *mtproto.TLUsersGetIsPremiumRequiredToContact) (*mtproto.Vector_Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("users.getIsPremiumRequiredToContact - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.UsersGetIsPremiumRequiredToContact(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("users.getIsPremiumRequiredToContact - reply: {%s}", r)
	return r, err
}

// MessagesSetDefaultHistoryTTL
// messages.setDefaultHistoryTTL#9eb51445 period:int = Bool;
func (s *Service) MessagesSetDefaultHistoryTTL(ctx context.Context, request *mtproto.TLMessagesSetDefaultHistoryTTL) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.setDefaultHistoryTTL - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesSetDefaultHistoryTTL(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.setDefaultHistoryTTL - reply: {%s}", r)
	return r, err
}

// MessagesGetDefaultHistoryTTL
// messages.getDefaultHistoryTTL#658b7188 = DefaultHistoryTTL;
func (s *Service) MessagesGetDefaultHistoryTTL(ctx context.Context, request *mtproto.TLMessagesGetDefaultHistoryTTL) (*mtproto.DefaultHistoryTTL, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("messages.getDefaultHistoryTTL - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetDefaultHistoryTTL(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("messages.getDefaultHistoryTTL - reply: {%s}", r)
	return r, err
}
