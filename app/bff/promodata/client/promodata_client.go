/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package promodata_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type PromoDataClient interface {
	HelpGetPromoData(ctx context.Context, in *mtproto.TLHelpGetPromoData) (*mtproto.Help_PromoData, error)
	HelpHidePromoData(ctx context.Context, in *mtproto.TLHelpHidePromoData) (*mtproto.Bool, error)
}

type defaultPromoDataClient struct {
	cli zrpc.Client
}

func NewPromoDataClient(cli zrpc.Client) PromoDataClient {
	return &defaultPromoDataClient{
		cli: cli,
	}
}

// HelpGetPromoData
// help.getPromoData#c0977421 = help.PromoData;
func (m *defaultPromoDataClient) HelpGetPromoData(ctx context.Context, in *mtproto.TLHelpGetPromoData) (*mtproto.Help_PromoData, error) {
	client := mtproto.NewRPCPromoDataClient(m.cli.Conn())
	return client.HelpGetPromoData(ctx, in)
}

// HelpHidePromoData
// help.hidePromoData#1e251c95 peer:InputPeer = Bool;
func (m *defaultPromoDataClient) HelpHidePromoData(ctx context.Context, in *mtproto.TLHelpHidePromoData) (*mtproto.Bool, error) {
	client := mtproto.NewRPCPromoDataClient(m.cli.Conn())
	return client.HelpHidePromoData(ctx, in)
}
