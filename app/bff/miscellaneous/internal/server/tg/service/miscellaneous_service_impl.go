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
	"github.com/teamgram/teamgram-server/v2/app/bff/miscellaneous/internal/core"
)

// HelpSaveAppLog
// help.saveAppLog#6f02f748 events:Vector<InputAppEvent> = Bool;
func (s *Service) HelpSaveAppLog(ctx context.Context, request *tg.TLHelpSaveAppLog) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("help.saveAppLog - metadata: %s, request: %s", c.MD, request)

	r, err := c.HelpSaveAppLog(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("help.saveAppLog - reply: %s", r)
	return r, err
}
