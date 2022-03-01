/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package tsf_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type TsfClient interface {
	HelpGetUserInfo(ctx context.Context, in *mtproto.TLHelpGetUserInfo) (*mtproto.Help_UserInfo, error)
	HelpEditUserInfo(ctx context.Context, in *mtproto.TLHelpEditUserInfo) (*mtproto.Help_UserInfo, error)
}

type defaultTsfClient struct {
	cli zrpc.Client
}

func NewTsfClient(cli zrpc.Client) TsfClient {
	return &defaultTsfClient{
		cli: cli,
	}
}

// HelpGetUserInfo
// help.getUserInfo#38a08d3 user_id:InputUser = help.UserInfo;
func (m *defaultTsfClient) HelpGetUserInfo(ctx context.Context, in *mtproto.TLHelpGetUserInfo) (*mtproto.Help_UserInfo, error) {
	client := mtproto.NewRPCTsfClient(m.cli.Conn())
	return client.HelpGetUserInfo(ctx, in)
}

// HelpEditUserInfo
// help.editUserInfo#66b91b70 user_id:InputUser message:string entities:Vector<MessageEntity> = help.UserInfo;
func (m *defaultTsfClient) HelpEditUserInfo(ctx context.Context, in *mtproto.TLHelpEditUserInfo) (*mtproto.Help_UserInfo, error) {
	client := mtproto.NewRPCTsfClient(m.cli.Conn())
	return client.HelpEditUserInfo(ctx, in)
}
