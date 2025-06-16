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
	"github.com/teamgram/teamgram-server/v2/app/bff/tos/internal/core"
)

// HelpGetTermsOfServiceUpdate
// help.getTermsOfServiceUpdate#2ca51fd1 = help.TermsOfServiceUpdate;
func (s *Service) HelpGetTermsOfServiceUpdate(ctx context.Context, request *tg.TLHelpGetTermsOfServiceUpdate) (*tg.HelpTermsOfServiceUpdate, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("help.getTermsOfServiceUpdate - metadata: {}, request: {%v}", request)

	r, err := c.HelpGetTermsOfServiceUpdate(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("help.getTermsOfServiceUpdate - reply: {%v}", r)
	return r, err
}

// HelpAcceptTermsOfService
// help.acceptTermsOfService#ee72f79a id:DataJSON = Bool;
func (s *Service) HelpAcceptTermsOfService(ctx context.Context, request *tg.TLHelpAcceptTermsOfService) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("help.acceptTermsOfService - metadata: {}, request: {%v}", request)

	r, err := c.HelpAcceptTermsOfService(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("help.acceptTermsOfService - reply: {%v}", r)
	return r, err
}
