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
	"github.com/teamgram/teamgram-server/app/bff/autodownload/internal/core"
)

// AccountGetAutoDownloadSettings
// account.getAutoDownloadSettings#56da0b3f = account.AutoDownloadSettings;
func (s *Service) AccountGetAutoDownloadSettings(ctx context.Context, request *mtproto.TLAccountGetAutoDownloadSettings) (*mtproto.Account_AutoDownloadSettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.getAutoDownloadSettings - metadata: %s, request: %s", c.MD, request)

	r, err := c.AccountGetAutoDownloadSettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.getAutoDownloadSettings - reply: %s", r)
	return r, err
}

// AccountSaveAutoDownloadSettings
// account.saveAutoDownloadSettings#76f36233 flags:# low:flags.0?true high:flags.1?true settings:AutoDownloadSettings = Bool;
func (s *Service) AccountSaveAutoDownloadSettings(ctx context.Context, request *mtproto.TLAccountSaveAutoDownloadSettings) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.saveAutoDownloadSettings - metadata: %s, request: %s", c.MD, request)

	r, err := c.AccountSaveAutoDownloadSettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.saveAutoDownloadSettings - reply: %s", r)
	return r, err
}
