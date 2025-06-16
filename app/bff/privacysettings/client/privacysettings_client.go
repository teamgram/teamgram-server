/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package privacysettingsclient

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/bff/privacysettings/privacysettings/privacysettingsservice"

	"github.com/cloudwego/kitex/client"
)

type PrivacySettingsClient interface {
	AccountGetPrivacy(ctx context.Context, in *tg.TLAccountGetPrivacy) (*tg.AccountPrivacyRules, error)
	AccountSetPrivacy(ctx context.Context, in *tg.TLAccountSetPrivacy) (*tg.AccountPrivacyRules, error)
	AccountGetGlobalPrivacySettings(ctx context.Context, in *tg.TLAccountGetGlobalPrivacySettings) (*tg.GlobalPrivacySettings, error)
	AccountSetGlobalPrivacySettings(ctx context.Context, in *tg.TLAccountSetGlobalPrivacySettings) (*tg.GlobalPrivacySettings, error)
	UsersGetIsPremiumRequiredToContact(ctx context.Context, in *tg.TLUsersGetIsPremiumRequiredToContact) (*tg.VectorBool, error)
	MessagesSetDefaultHistoryTTL(ctx context.Context, in *tg.TLMessagesSetDefaultHistoryTTL) (*tg.Bool, error)
	MessagesGetDefaultHistoryTTL(ctx context.Context, in *tg.TLMessagesGetDefaultHistoryTTL) (*tg.DefaultHistoryTTL, error)
}

type defaultPrivacySettingsClient struct {
	cli client.Client
}

func NewPrivacySettingsClient(cli client.Client) PrivacySettingsClient {
	return &defaultPrivacySettingsClient{
		cli: cli,
	}
}

// AccountGetPrivacy
// account.getPrivacy#dadbc950 key:InputPrivacyKey = account.PrivacyRules;
func (m *defaultPrivacySettingsClient) AccountGetPrivacy(ctx context.Context, in *tg.TLAccountGetPrivacy) (*tg.AccountPrivacyRules, error) {
	cli := privacysettingsservice.NewRPCPrivacySettingsClient(m.cli)
	return cli.AccountGetPrivacy(ctx, in)
}

// AccountSetPrivacy
// account.setPrivacy#c9f81ce8 key:InputPrivacyKey rules:Vector<InputPrivacyRule> = account.PrivacyRules;
func (m *defaultPrivacySettingsClient) AccountSetPrivacy(ctx context.Context, in *tg.TLAccountSetPrivacy) (*tg.AccountPrivacyRules, error) {
	cli := privacysettingsservice.NewRPCPrivacySettingsClient(m.cli)
	return cli.AccountSetPrivacy(ctx, in)
}

// AccountGetGlobalPrivacySettings
// account.getGlobalPrivacySettings#eb2b4cf6 = GlobalPrivacySettings;
func (m *defaultPrivacySettingsClient) AccountGetGlobalPrivacySettings(ctx context.Context, in *tg.TLAccountGetGlobalPrivacySettings) (*tg.GlobalPrivacySettings, error) {
	cli := privacysettingsservice.NewRPCPrivacySettingsClient(m.cli)
	return cli.AccountGetGlobalPrivacySettings(ctx, in)
}

// AccountSetGlobalPrivacySettings
// account.setGlobalPrivacySettings#1edaaac2 settings:GlobalPrivacySettings = GlobalPrivacySettings;
func (m *defaultPrivacySettingsClient) AccountSetGlobalPrivacySettings(ctx context.Context, in *tg.TLAccountSetGlobalPrivacySettings) (*tg.GlobalPrivacySettings, error) {
	cli := privacysettingsservice.NewRPCPrivacySettingsClient(m.cli)
	return cli.AccountSetGlobalPrivacySettings(ctx, in)
}

// UsersGetIsPremiumRequiredToContact
// users.getIsPremiumRequiredToContact#a622aa10 id:Vector<InputUser> = Vector<Bool>;
func (m *defaultPrivacySettingsClient) UsersGetIsPremiumRequiredToContact(ctx context.Context, in *tg.TLUsersGetIsPremiumRequiredToContact) (*tg.VectorBool, error) {
	cli := privacysettingsservice.NewRPCPrivacySettingsClient(m.cli)
	return cli.UsersGetIsPremiumRequiredToContact(ctx, in)
}

// MessagesSetDefaultHistoryTTL
// messages.setDefaultHistoryTTL#9eb51445 period:int = Bool;
func (m *defaultPrivacySettingsClient) MessagesSetDefaultHistoryTTL(ctx context.Context, in *tg.TLMessagesSetDefaultHistoryTTL) (*tg.Bool, error) {
	cli := privacysettingsservice.NewRPCPrivacySettingsClient(m.cli)
	return cli.MessagesSetDefaultHistoryTTL(ctx, in)
}

// MessagesGetDefaultHistoryTTL
// messages.getDefaultHistoryTTL#658b7188 = DefaultHistoryTTL;
func (m *defaultPrivacySettingsClient) MessagesGetDefaultHistoryTTL(ctx context.Context, in *tg.TLMessagesGetDefaultHistoryTTL) (*tg.DefaultHistoryTTL, error) {
	cli := privacysettingsservice.NewRPCPrivacySettingsClient(m.cli)
	return cli.MessagesGetDefaultHistoryTTL(ctx, in)
}
