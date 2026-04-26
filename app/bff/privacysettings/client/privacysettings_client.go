/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package privacysettingsclient

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/bff/privacysettings/privacysettings/privacysettingsservice"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
)

type PrivacySettingsClient interface {
	AccountGetPrivacy(ctx context.Context, in *tg.TLAccountGetPrivacy) (*tg.AccountPrivacyRules, error)
	AccountSetPrivacy(ctx context.Context, in *tg.TLAccountSetPrivacy) (*tg.AccountPrivacyRules, error)
	AccountGetGlobalPrivacySettings(ctx context.Context, in *tg.TLAccountGetGlobalPrivacySettings) (*tg.GlobalPrivacySettings, error)
	AccountSetGlobalPrivacySettings(ctx context.Context, in *tg.TLAccountSetGlobalPrivacySettings) (*tg.GlobalPrivacySettings, error)
	UsersGetRequirementsToContact(ctx context.Context, in *tg.TLUsersGetRequirementsToContact) (*tg.VectorRequirementToContact, error)
	MessagesSetDefaultHistoryTTL(ctx context.Context, in *tg.TLMessagesSetDefaultHistoryTTL) (*tg.Bool, error)
	MessagesGetDefaultHistoryTTL(ctx context.Context, in *tg.TLMessagesGetDefaultHistoryTTL) (*tg.DefaultHistoryTTL, error)
	Close() error
}

type defaultPrivacySettingsClient struct {
	cli client.Client
	rpc privacysettingsservice.Client
}

func NewPrivacySettingsClient(cli client.Client) PrivacySettingsClient {
	return &defaultPrivacySettingsClient{
		cli: cli,
		rpc: privacysettingsservice.NewRPCPrivacySettingsClient(cli),
	}
}

func (m *defaultPrivacySettingsClient) Close() error {
	if closer, ok := any(m.cli).(interface{ Close() error }); ok {
		return closer.Close()
	}
	return nil
}

// AccountGetPrivacy
// account.getPrivacy#dadbc950 key:InputPrivacyKey = account.PrivacyRules;
func (m *defaultPrivacySettingsClient) AccountGetPrivacy(ctx context.Context, in *tg.TLAccountGetPrivacy) (*tg.AccountPrivacyRules, error) {
	return m.rpc.AccountGetPrivacy(ctx, in)
}

// AccountSetPrivacy
// account.setPrivacy#c9f81ce8 key:InputPrivacyKey rules:Vector<InputPrivacyRule> = account.PrivacyRules;
func (m *defaultPrivacySettingsClient) AccountSetPrivacy(ctx context.Context, in *tg.TLAccountSetPrivacy) (*tg.AccountPrivacyRules, error) {
	return m.rpc.AccountSetPrivacy(ctx, in)
}

// AccountGetGlobalPrivacySettings
// account.getGlobalPrivacySettings#eb2b4cf6 = GlobalPrivacySettings;
func (m *defaultPrivacySettingsClient) AccountGetGlobalPrivacySettings(ctx context.Context, in *tg.TLAccountGetGlobalPrivacySettings) (*tg.GlobalPrivacySettings, error) {
	return m.rpc.AccountGetGlobalPrivacySettings(ctx, in)
}

// AccountSetGlobalPrivacySettings
// account.setGlobalPrivacySettings#1edaaac2 settings:GlobalPrivacySettings = GlobalPrivacySettings;
func (m *defaultPrivacySettingsClient) AccountSetGlobalPrivacySettings(ctx context.Context, in *tg.TLAccountSetGlobalPrivacySettings) (*tg.GlobalPrivacySettings, error) {
	return m.rpc.AccountSetGlobalPrivacySettings(ctx, in)
}

// UsersGetRequirementsToContact
// users.getRequirementsToContact#d89a83a3 id:Vector<InputUser> = Vector<RequirementToContact>;
func (m *defaultPrivacySettingsClient) UsersGetRequirementsToContact(ctx context.Context, in *tg.TLUsersGetRequirementsToContact) (*tg.VectorRequirementToContact, error) {
	return m.rpc.UsersGetRequirementsToContact(ctx, in)
}

// MessagesSetDefaultHistoryTTL
// messages.setDefaultHistoryTTL#9eb51445 period:int = Bool;
func (m *defaultPrivacySettingsClient) MessagesSetDefaultHistoryTTL(ctx context.Context, in *tg.TLMessagesSetDefaultHistoryTTL) (*tg.Bool, error) {
	return m.rpc.MessagesSetDefaultHistoryTTL(ctx, in)
}

// MessagesGetDefaultHistoryTTL
// messages.getDefaultHistoryTTL#658b7188 = DefaultHistoryTTL;
func (m *defaultPrivacySettingsClient) MessagesGetDefaultHistoryTTL(ctx context.Context, in *tg.TLMessagesGetDefaultHistoryTTL) (*tg.DefaultHistoryTTL, error) {
	return m.rpc.MessagesGetDefaultHistoryTTL(ctx, in)
}
