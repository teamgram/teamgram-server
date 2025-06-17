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
	"github.com/teamgram/teamgram-server/v2/app/bff/premium/internal/core"
)

// HelpGetPremiumPromo
// help.getPremiumPromo#b81b93d4 = help.PremiumPromo;
func (s *Service) HelpGetPremiumPromo(ctx context.Context, request *tg.TLHelpGetPremiumPromo) (*tg.HelpPremiumPromo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("help.getPremiumPromo - metadata: {}, request: {%v}", request)

	r, err := c.HelpGetPremiumPromo(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("help.getPremiumPromo - reply: {%v}", r)
	return r, err
}

// PaymentsAssignAppStoreTransaction
// payments.assignAppStoreTransaction#80ed747d receipt:bytes purpose:InputStorePaymentPurpose = Updates;
func (s *Service) PaymentsAssignAppStoreTransaction(ctx context.Context, request *tg.TLPaymentsAssignAppStoreTransaction) (*tg.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("payments.assignAppStoreTransaction - metadata: {}, request: {%v}", request)

	r, err := c.PaymentsAssignAppStoreTransaction(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("payments.assignAppStoreTransaction - reply: {%v}", r)
	return r, err
}

// PaymentsAssignPlayMarketTransaction
// payments.assignPlayMarketTransaction#dffd50d3 receipt:DataJSON purpose:InputStorePaymentPurpose = Updates;
func (s *Service) PaymentsAssignPlayMarketTransaction(ctx context.Context, request *tg.TLPaymentsAssignPlayMarketTransaction) (*tg.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("payments.assignPlayMarketTransaction - metadata: {}, request: {%v}", request)

	r, err := c.PaymentsAssignPlayMarketTransaction(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("payments.assignPlayMarketTransaction - reply: {%v}", r)
	return r, err
}

// PaymentsCanPurchaseStore
// payments.canPurchaseStore#4fdc5ea7 purpose:InputStorePaymentPurpose = Bool;
func (s *Service) PaymentsCanPurchaseStore(ctx context.Context, request *tg.TLPaymentsCanPurchaseStore) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("payments.canPurchaseStore - metadata: {}, request: {%v}", request)

	r, err := c.PaymentsCanPurchaseStore(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("payments.canPurchaseStore - reply: {%v}", r)
	return r, err
}

// PaymentsCanPurchasePremium
// payments.canPurchasePremium#9fc19eb6 purpose:InputStorePaymentPurpose = Bool;
func (s *Service) PaymentsCanPurchasePremium(ctx context.Context, request *tg.TLPaymentsCanPurchasePremium) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("payments.canPurchasePremium - metadata: {}, request: {%v}", request)

	r, err := c.PaymentsCanPurchasePremium(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("payments.canPurchasePremium - reply: {%v}", r)
	return r, err
}
