/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package model

import (
	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

var _ DeliveryFailedOperationsModel = (*customDeliveryFailedOperationsModel)(nil)

type (
	// DeliveryFailedOperationsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDeliveryFailedOperationsModel.
	DeliveryFailedOperationsModel interface {
		deliveryFailedOperationsModel
		bizDeliveryFailedOperationsModel
		extendDeliveryFailedOperationsModel
	}

	customDeliveryFailedOperationsModel struct {
		*defaultDeliveryFailedOperationsModel
	}
)

// NewDeliveryFailedOperationsModel returns a model for the database table.
func NewDeliveryFailedOperationsModel(db *sqlx.DB) DeliveryFailedOperationsModel {
	return &customDeliveryFailedOperationsModel{
		defaultDeliveryFailedOperationsModel: newDeliveryFailedOperationsModel(db),
	}
}
