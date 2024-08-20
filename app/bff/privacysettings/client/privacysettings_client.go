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

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type PrivacySettingsClient interface {
	AccountGetPrivacy(ctx context.Context, in *mtproto.TLAccountGetPrivacy) (*mtproto.Account_PrivacyRules, error)
	AccountSetPrivacy(ctx context.Context, in *mtproto.TLAccountSetPrivacy) (*mtproto.Account_PrivacyRules, error)
	AccountGetGlobalPrivacySettings(ctx context.Context, in *mtproto.TLAccountGetGlobalPrivacySettings) (*mtproto.GlobalPrivacySettings, error)
	AccountSetGlobalPrivacySettings(ctx context.Context, in *mtproto.TLAccountSetGlobalPrivacySettings) (*mtproto.GlobalPrivacySettings, error)
	UsersGetIsPremiumRequiredToContact(ctx context.Context, in *mtproto.TLUsersGetIsPremiumRequiredToContact) (*mtproto.Vector_Bool, error)
	MessagesSetDefaultHistoryTTL(ctx context.Context, in *mtproto.TLMessagesSetDefaultHistoryTTL) (*mtproto.Bool, error)
	MessagesGetDefaultHistoryTTL(ctx context.Context, in *mtproto.TLMessagesGetDefaultHistoryTTL) (*mtproto.DefaultHistoryTTL, error)
}

type defaultPrivacySettingsClient struct {
	cli zrpc.Client
}

func NewPrivacySettingsClient(cli zrpc.Client) PrivacySettingsClient {
	return &defaultPrivacySettingsClient{
		cli: cli,
	}
}

// AccountGetPrivacy
// account.getPrivacy#dadbc950 key:InputPrivacyKey = account.PrivacyRules;
func (m *defaultPrivacySettingsClient) AccountGetPrivacy(ctx context.Context, in *mtproto.TLAccountGetPrivacy) (*mtproto.Account_PrivacyRules, error) {
	client := mtproto.NewRPCPrivacySettingsClient(m.cli.Conn())
	return client.AccountGetPrivacy(ctx, in)
}

// AccountSetPrivacy
// account.setPrivacy#c9f81ce8 key:InputPrivacyKey rules:Vector<InputPrivacyRule> = account.PrivacyRules;
func (m *defaultPrivacySettingsClient) AccountSetPrivacy(ctx context.Context, in *mtproto.TLAccountSetPrivacy) (*mtproto.Account_PrivacyRules, error) {
	client := mtproto.NewRPCPrivacySettingsClient(m.cli.Conn())
	return client.AccountSetPrivacy(ctx, in)
}

// AccountGetGlobalPrivacySettings
// account.getGlobalPrivacySettings#eb2b4cf6 = GlobalPrivacySettings;
func (m *defaultPrivacySettingsClient) AccountGetGlobalPrivacySettings(ctx context.Context, in *mtproto.TLAccountGetGlobalPrivacySettings) (*mtproto.GlobalPrivacySettings, error) {
	client := mtproto.NewRPCPrivacySettingsClient(m.cli.Conn())
	return client.AccountGetGlobalPrivacySettings(ctx, in)
}

// AccountSetGlobalPrivacySettings
// account.setGlobalPrivacySettings#1edaaac2 settings:GlobalPrivacySettings = GlobalPrivacySettings;
func (m *defaultPrivacySettingsClient) AccountSetGlobalPrivacySettings(ctx context.Context, in *mtproto.TLAccountSetGlobalPrivacySettings) (*mtproto.GlobalPrivacySettings, error) {
	client := mtproto.NewRPCPrivacySettingsClient(m.cli.Conn())
	return client.AccountSetGlobalPrivacySettings(ctx, in)
}

// UsersGetIsPremiumRequiredToContact
// users.getIsPremiumRequiredToContact#a622aa10 id:Vector<InputUser> = Vector<Bool>;
func (m *defaultPrivacySettingsClient) UsersGetIsPremiumRequiredToContact(ctx context.Context, in *mtproto.TLUsersGetIsPremiumRequiredToContact) (*mtproto.Vector_Bool, error) {
	client := mtproto.NewRPCPrivacySettingsClient(m.cli.Conn())
	return client.UsersGetIsPremiumRequiredToContact(ctx, in)
}

// MessagesSetDefaultHistoryTTL
// messages.setDefaultHistoryTTL#9eb51445 period:int = Bool;
func (m *defaultPrivacySettingsClient) MessagesSetDefaultHistoryTTL(ctx context.Context, in *mtproto.TLMessagesSetDefaultHistoryTTL) (*mtproto.Bool, error) {
	client := mtproto.NewRPCPrivacySettingsClient(m.cli.Conn())
	return client.MessagesSetDefaultHistoryTTL(ctx, in)
}

// MessagesGetDefaultHistoryTTL
// messages.getDefaultHistoryTTL#658b7188 = DefaultHistoryTTL;
func (m *defaultPrivacySettingsClient) MessagesGetDefaultHistoryTTL(ctx context.Context, in *mtproto.TLMessagesGetDefaultHistoryTTL) (*mtproto.DefaultHistoryTTL, error) {
	client := mtproto.NewRPCPrivacySettingsClient(m.cli.Conn())
	return client.MessagesGetDefaultHistoryTTL(ctx, in)
}
