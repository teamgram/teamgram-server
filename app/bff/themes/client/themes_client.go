/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package themes_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type ThemesClient interface {
	AccountUploadTheme(ctx context.Context, in *mtproto.TLAccountUploadTheme) (*mtproto.Document, error)
	AccountCreateTheme(ctx context.Context, in *mtproto.TLAccountCreateTheme) (*mtproto.Theme, error)
	AccountUpdateTheme(ctx context.Context, in *mtproto.TLAccountUpdateTheme) (*mtproto.Theme, error)
	AccountSaveTheme(ctx context.Context, in *mtproto.TLAccountSaveTheme) (*mtproto.Bool, error)
	AccountInstallTheme(ctx context.Context, in *mtproto.TLAccountInstallTheme) (*mtproto.Bool, error)
	AccountGetTheme(ctx context.Context, in *mtproto.TLAccountGetTheme) (*mtproto.Theme, error)
	AccountGetThemes(ctx context.Context, in *mtproto.TLAccountGetThemes) (*mtproto.Account_Themes, error)
	AccountGetChatThemesD638DE89(ctx context.Context, in *mtproto.TLAccountGetChatThemesD638DE89) (*mtproto.Account_Themes, error)
	MessagesSetChatTheme(ctx context.Context, in *mtproto.TLMessagesSetChatTheme) (*mtproto.Updates, error)
	AccountGetChatThemesD6D71D7B(ctx context.Context, in *mtproto.TLAccountGetChatThemesD6D71D7B) (*mtproto.Account_ChatThemes, error)
}

type defaultThemesClient struct {
	cli zrpc.Client
}

func NewThemesClient(cli zrpc.Client) ThemesClient {
	return &defaultThemesClient{
		cli: cli,
	}
}

// AccountUploadTheme
// account.uploadTheme#1c3db333 flags:# file:InputFile thumb:flags.0?InputFile file_name:string mime_type:string = Document;
func (m *defaultThemesClient) AccountUploadTheme(ctx context.Context, in *mtproto.TLAccountUploadTheme) (*mtproto.Document, error) {
	client := mtproto.NewRPCThemesClient(m.cli.Conn())
	return client.AccountUploadTheme(ctx, in)
}

// AccountCreateTheme
// account.createTheme#652e4400 flags:# slug:string title:string document:flags.2?InputDocument settings:flags.3?Vector<InputThemeSettings> = Theme;
func (m *defaultThemesClient) AccountCreateTheme(ctx context.Context, in *mtproto.TLAccountCreateTheme) (*mtproto.Theme, error) {
	client := mtproto.NewRPCThemesClient(m.cli.Conn())
	return client.AccountCreateTheme(ctx, in)
}

// AccountUpdateTheme
// account.updateTheme#2bf40ccc flags:# format:string theme:InputTheme slug:flags.0?string title:flags.1?string document:flags.2?InputDocument settings:flags.3?Vector<InputThemeSettings> = Theme;
func (m *defaultThemesClient) AccountUpdateTheme(ctx context.Context, in *mtproto.TLAccountUpdateTheme) (*mtproto.Theme, error) {
	client := mtproto.NewRPCThemesClient(m.cli.Conn())
	return client.AccountUpdateTheme(ctx, in)
}

// AccountSaveTheme
// account.saveTheme#f257106c theme:InputTheme unsave:Bool = Bool;
func (m *defaultThemesClient) AccountSaveTheme(ctx context.Context, in *mtproto.TLAccountSaveTheme) (*mtproto.Bool, error) {
	client := mtproto.NewRPCThemesClient(m.cli.Conn())
	return client.AccountSaveTheme(ctx, in)
}

// AccountInstallTheme
// account.installTheme#c727bb3b flags:# dark:flags.0?true theme:flags.1?InputTheme format:flags.2?string base_theme:flags.3?BaseTheme = Bool;
func (m *defaultThemesClient) AccountInstallTheme(ctx context.Context, in *mtproto.TLAccountInstallTheme) (*mtproto.Bool, error) {
	client := mtproto.NewRPCThemesClient(m.cli.Conn())
	return client.AccountInstallTheme(ctx, in)
}

// AccountGetTheme
// account.getTheme#8d9d742b format:string theme:InputTheme document_id:long = Theme;
func (m *defaultThemesClient) AccountGetTheme(ctx context.Context, in *mtproto.TLAccountGetTheme) (*mtproto.Theme, error) {
	client := mtproto.NewRPCThemesClient(m.cli.Conn())
	return client.AccountGetTheme(ctx, in)
}

// AccountGetThemes
// account.getThemes#7206e458 format:string hash:long = account.Themes;
func (m *defaultThemesClient) AccountGetThemes(ctx context.Context, in *mtproto.TLAccountGetThemes) (*mtproto.Account_Themes, error) {
	client := mtproto.NewRPCThemesClient(m.cli.Conn())
	return client.AccountGetThemes(ctx, in)
}

// AccountGetChatThemesD638DE89
// account.getChatThemes#d638de89 hash:long = account.Themes;
func (m *defaultThemesClient) AccountGetChatThemesD638DE89(ctx context.Context, in *mtproto.TLAccountGetChatThemesD638DE89) (*mtproto.Account_Themes, error) {
	client := mtproto.NewRPCThemesClient(m.cli.Conn())
	return client.AccountGetChatThemesD638DE89(ctx, in)
}

// MessagesSetChatTheme
// messages.setChatTheme#e63be13f peer:InputPeer emoticon:string = Updates;
func (m *defaultThemesClient) MessagesSetChatTheme(ctx context.Context, in *mtproto.TLMessagesSetChatTheme) (*mtproto.Updates, error) {
	client := mtproto.NewRPCThemesClient(m.cli.Conn())
	return client.MessagesSetChatTheme(ctx, in)
}

// AccountGetChatThemesD6D71D7B
// account.getChatThemes#d6d71d7b hash:int = account.ChatThemes;
func (m *defaultThemesClient) AccountGetChatThemesD6D71D7B(ctx context.Context, in *mtproto.TLAccountGetChatThemesD6D71D7B) (*mtproto.Account_ChatThemes, error) {
	client := mtproto.NewRPCThemesClient(m.cli.Conn())
	return client.AccountGetChatThemesD6D71D7B(ctx, in)
}
