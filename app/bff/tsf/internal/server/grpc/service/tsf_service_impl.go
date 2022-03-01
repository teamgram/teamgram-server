/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/tsf/internal/core"
)

// HelpGetUserInfo
// help.getUserInfo#38a08d3 user_id:InputUser = help.UserInfo;
func (s *Service) HelpGetUserInfo(ctx context.Context, request *mtproto.TLHelpGetUserInfo) (*mtproto.Help_UserInfo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("help.getUserInfo - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.HelpGetUserInfo(request)
	if err != nil {
		return nil, err
	}

	c.Infof("help.getUserInfo - reply: %s", r.DebugString())
	return r, err
}

// HelpEditUserInfo
// help.editUserInfo#66b91b70 user_id:InputUser message:string entities:Vector<MessageEntity> = help.UserInfo;
func (s *Service) HelpEditUserInfo(ctx context.Context, request *mtproto.TLHelpEditUserInfo) (*mtproto.Help_UserInfo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("help.editUserInfo - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.HelpEditUserInfo(request)
	if err != nil {
		return nil, err
	}

	c.Infof("help.editUserInfo - reply: %s", r.DebugString())
	return r, err
}
