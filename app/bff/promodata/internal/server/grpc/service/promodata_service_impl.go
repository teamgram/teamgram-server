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
	"github.com/teamgram/teamgram-server/app/bff/promodata/internal/core"
)

// HelpGetPromoData
// help.getPromoData#c0977421 = help.PromoData;
func (s *Service) HelpGetPromoData(ctx context.Context, request *mtproto.TLHelpGetPromoData) (*mtproto.Help_PromoData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("help.getPromoData - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.HelpGetPromoData(request)
	if err != nil {
		return nil, err
	}

	c.Infof("help.getPromoData - reply: %s", r.DebugString())
	return r, err
}

// HelpHidePromoData
// help.hidePromoData#1e251c95 peer:InputPeer = Bool;
func (s *Service) HelpHidePromoData(ctx context.Context, request *mtproto.TLHelpHidePromoData) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("help.hidePromoData - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.HelpHidePromoData(request)
	if err != nil {
		return nil, err
	}

	c.Infof("help.hidePromoData - reply: %s", r.DebugString())
	return r, err
}
