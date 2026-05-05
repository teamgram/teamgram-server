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

type Models struct {
	DeliveryFailedOperationsModel   DeliveryFailedOperationsModel
	DialogSideEffectOutboxModel     DialogSideEffectOutboxModel
	PushTaskOutboxModel             PushTaskOutboxModel
	UserDialogsModel                UserDialogsModel
	UserMessageViewsModel           UserMessageViewsModel
	UserOperationResultsModel       UserOperationResultsModel
	UserPtsEventsModel              UserPtsEventsModel
	UserPtsStateModel               UserPtsStateModel
	UserupdatesPartitionFencesModel UserupdatesPartitionFencesModel
}

type TxModels struct {
	DeliveryFailedOperationsModel   DeliveryFailedOperationsTxModel
	DialogSideEffectOutboxModel     DialogSideEffectOutboxTxModel
	PushTaskOutboxModel             PushTaskOutboxTxModel
	UserDialogsModel                UserDialogsTxModel
	UserMessageViewsModel           UserMessageViewsTxModel
	UserOperationResultsModel       UserOperationResultsTxModel
	UserPtsEventsModel              UserPtsEventsTxModel
	UserPtsStateModel               UserPtsStateTxModel
	UserupdatesPartitionFencesModel UserupdatesPartitionFencesTxModel
}

func NewModels(db *sqlx.DB) *Models {
	return &Models{
		DeliveryFailedOperationsModel:   NewDeliveryFailedOperationsModel(db),
		DialogSideEffectOutboxModel:     NewDialogSideEffectOutboxModel(db),
		PushTaskOutboxModel:             NewPushTaskOutboxModel(db),
		UserDialogsModel:                NewUserDialogsModel(db),
		UserMessageViewsModel:           NewUserMessageViewsModel(db),
		UserOperationResultsModel:       NewUserOperationResultsModel(db),
		UserPtsEventsModel:              NewUserPtsEventsModel(db),
		UserPtsStateModel:               NewUserPtsStateModel(db),
		UserupdatesPartitionFencesModel: NewUserupdatesPartitionFencesModel(db),
	}
}

func (m *Models) WithTx(tx *sqlx.Tx) *TxModels {
	return &TxModels{
		DeliveryFailedOperationsModel:   NewDeliveryFailedOperationsTxModel(tx),
		DialogSideEffectOutboxModel:     NewDialogSideEffectOutboxTxModel(tx),
		PushTaskOutboxModel:             NewPushTaskOutboxTxModel(tx),
		UserDialogsModel:                NewUserDialogsTxModel(tx),
		UserMessageViewsModel:           NewUserMessageViewsTxModel(tx),
		UserOperationResultsModel:       NewUserOperationResultsTxModel(tx),
		UserPtsEventsModel:              NewUserPtsEventsTxModel(tx),
		UserPtsStateModel:               NewUserPtsStateTxModel(tx),
		UserupdatesPartitionFencesModel: NewUserupdatesPartitionFencesTxModel(tx),
	}
}
