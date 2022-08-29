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
	"github.com/teamgram/teamgram-server/app/bff/premium/internal/core"
)

// HelpGetPremiumPromo
// help.getPremiumPromo#b81b93d4 = help.PremiumPromo;
func (s *Service) HelpGetPremiumPromo(ctx context.Context, request *mtproto.TLHelpGetPremiumPromo) (*mtproto.Help_PremiumPromo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("help.getPremiumPromo - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.HelpGetPremiumPromo(request)
	if err != nil {
		return nil, err
	}

	c.Infof("help.getPremiumPromo - reply: %s", r.DebugString())
	return r, err
}

// PaymentsAssignAppStoreTransaction
// payments.assignAppStoreTransaction#80ed747d receipt:bytes purpose:InputStorePaymentPurpose = Updates;
func (s *Service) PaymentsAssignAppStoreTransaction(ctx context.Context, request *mtproto.TLPaymentsAssignAppStoreTransaction) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("payments.assignAppStoreTransaction - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.PaymentsAssignAppStoreTransaction(request)
	if err != nil {
		return nil, err
	}

	c.Infof("payments.assignAppStoreTransaction - reply: %s", r.DebugString())
	return r, err
}

// PaymentsAssignPlayMarketTransaction
// payments.assignPlayMarketTransaction#dffd50d3 receipt:DataJSON purpose:InputStorePaymentPurpose = Updates;
func (s *Service) PaymentsAssignPlayMarketTransaction(ctx context.Context, request *mtproto.TLPaymentsAssignPlayMarketTransaction) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("payments.assignPlayMarketTransaction - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.PaymentsAssignPlayMarketTransaction(request)
	if err != nil {
		return nil, err
	}

	c.Infof("payments.assignPlayMarketTransaction - reply: %s", r.DebugString())
	return r, err
}

// PaymentsCanPurchasePremium
// payments.canPurchasePremium#9fc19eb6 purpose:InputStorePaymentPurpose = Bool;
func (s *Service) PaymentsCanPurchasePremium(ctx context.Context, request *mtproto.TLPaymentsCanPurchasePremium) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("payments.canPurchasePremium - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.PaymentsCanPurchasePremium(request)
	if err != nil {
		return nil, err
	}

	c.Infof("payments.canPurchasePremium - reply: %s", r.DebugString())
	return r, err
}

// PaymentsRequestRecurringPayment
// payments.requestRecurringPayment#146e958d user_id:InputUser recurring_init_charge:string invoice_media:InputMedia = Updates;
func (s *Service) PaymentsRequestRecurringPayment(ctx context.Context, request *mtproto.TLPaymentsRequestRecurringPayment) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("payments.requestRecurringPayment - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.PaymentsRequestRecurringPayment(request)
	if err != nil {
		return nil, err
	}

	c.Infof("payments.requestRecurringPayment - reply: %s", r.DebugString())
	return r, err
}

// PaymentsRestorePlayMarketReceipt
// payments.restorePlayMarketReceipt#d164e36a receipt:bytes = Updates;
func (s *Service) PaymentsRestorePlayMarketReceipt(ctx context.Context, request *mtproto.TLPaymentsRestorePlayMarketReceipt) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("payments.restorePlayMarketReceipt - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.PaymentsRestorePlayMarketReceipt(request)
	if err != nil {
		return nil, err
	}

	c.Infof("payments.restorePlayMarketReceipt - reply: %s", r.DebugString())
	return r, err
}
