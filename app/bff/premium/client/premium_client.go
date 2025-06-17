/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package premiumclient

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/bff/premium/premium/premiumservice"

	"github.com/cloudwego/kitex/client"
)

type PremiumClient interface {
	HelpGetPremiumPromo(ctx context.Context, in *tg.TLHelpGetPremiumPromo) (*tg.HelpPremiumPromo, error)
	PaymentsAssignAppStoreTransaction(ctx context.Context, in *tg.TLPaymentsAssignAppStoreTransaction) (*tg.Updates, error)
	PaymentsAssignPlayMarketTransaction(ctx context.Context, in *tg.TLPaymentsAssignPlayMarketTransaction) (*tg.Updates, error)
	PaymentsCanPurchaseStore(ctx context.Context, in *tg.TLPaymentsCanPurchaseStore) (*tg.Bool, error)
	PaymentsCanPurchasePremium(ctx context.Context, in *tg.TLPaymentsCanPurchasePremium) (*tg.Bool, error)
}

type defaultPremiumClient struct {
	cli client.Client
}

func NewPremiumClient(cli client.Client) PremiumClient {
	return &defaultPremiumClient{
		cli: cli,
	}
}

// HelpGetPremiumPromo
// help.getPremiumPromo#b81b93d4 = help.PremiumPromo;
func (m *defaultPremiumClient) HelpGetPremiumPromo(ctx context.Context, in *tg.TLHelpGetPremiumPromo) (*tg.HelpPremiumPromo, error) {
	cli := premiumservice.NewRPCPremiumClient(m.cli)
	return cli.HelpGetPremiumPromo(ctx, in)
}

// PaymentsAssignAppStoreTransaction
// payments.assignAppStoreTransaction#80ed747d receipt:bytes purpose:InputStorePaymentPurpose = Updates;
func (m *defaultPremiumClient) PaymentsAssignAppStoreTransaction(ctx context.Context, in *tg.TLPaymentsAssignAppStoreTransaction) (*tg.Updates, error) {
	cli := premiumservice.NewRPCPremiumClient(m.cli)
	return cli.PaymentsAssignAppStoreTransaction(ctx, in)
}

// PaymentsAssignPlayMarketTransaction
// payments.assignPlayMarketTransaction#dffd50d3 receipt:DataJSON purpose:InputStorePaymentPurpose = Updates;
func (m *defaultPremiumClient) PaymentsAssignPlayMarketTransaction(ctx context.Context, in *tg.TLPaymentsAssignPlayMarketTransaction) (*tg.Updates, error) {
	cli := premiumservice.NewRPCPremiumClient(m.cli)
	return cli.PaymentsAssignPlayMarketTransaction(ctx, in)
}

// PaymentsCanPurchaseStore
// payments.canPurchaseStore#4fdc5ea7 purpose:InputStorePaymentPurpose = Bool;
func (m *defaultPremiumClient) PaymentsCanPurchaseStore(ctx context.Context, in *tg.TLPaymentsCanPurchaseStore) (*tg.Bool, error) {
	cli := premiumservice.NewRPCPremiumClient(m.cli)
	return cli.PaymentsCanPurchaseStore(ctx, in)
}

// PaymentsCanPurchasePremium
// payments.canPurchasePremium#9fc19eb6 purpose:InputStorePaymentPurpose = Bool;
func (m *defaultPremiumClient) PaymentsCanPurchasePremium(ctx context.Context, in *tg.TLPaymentsCanPurchasePremium) (*tg.Bool, error) {
	cli := premiumservice.NewRPCPremiumClient(m.cli)
	return cli.PaymentsCanPurchasePremium(ctx, in)
}
