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
	"github.com/teamgram/teamgram-server/v2/app/bff/nsfw/internal/core"
)

// AccountSetContentSettings
// account.setContentSettings#b574b16b flags:# sensitive_enabled:flags.0?true = Bool;
func (s *Service) AccountSetContentSettings(ctx context.Context, request *tg.TLAccountSetContentSettings) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.setContentSettings - metadata: %s, request: %s", c.MD, request)

	r, err := c.AccountSetContentSettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.setContentSettings - reply: %s", r)
	return r, err
}

// AccountGetContentSettings
// account.getContentSettings#8b9b4dae = account.ContentSettings;
func (s *Service) AccountGetContentSettings(ctx context.Context, request *tg.TLAccountGetContentSettings) (*tg.AccountContentSettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.getContentSettings - metadata: %s, request: %s", c.MD, request)

	r, err := c.AccountGetContentSettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.getContentSettings - reply: %s", r)
	return r, err
}
