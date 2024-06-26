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
	"github.com/teamgram/teamgram-server/app/bff/miscellaneous/internal/core"
)

// HelpSaveAppLog
// help.saveAppLog#6f02f748 events:Vector<InputAppEvent> = Bool;
func (s *Service) HelpSaveAppLog(ctx context.Context, request *mtproto.TLHelpSaveAppLog) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("help.saveAppLog - metadata: %s, request: %s", c.MD, request)

	r, err := c.HelpSaveAppLog(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("help.saveAppLog - reply: %s", r)
	return r, err
}

// HelpTest
// help.test#c0e202f7 = Bool;
func (s *Service) HelpTest(ctx context.Context, request *mtproto.TLHelpTest) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("help.test - metadata: %s, request: %s", c.MD, request)

	r, err := c.HelpTest(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("help.test - reply: %s", r)
	return r, err
}
