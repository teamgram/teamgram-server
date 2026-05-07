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
	HashTagsModel                   HashTagsModel
	MessageReadOutboxModel          MessageReadOutboxModel
	PushTaskOutboxModel             PushTaskOutboxModel
	UserAuthSeqEventsModel          UserAuthSeqEventsModel
	UserAuthSeqStateModel           UserAuthSeqStateModel
	UserDialogsModel                UserDialogsModel
	UserMessageViewsModel           UserMessageViewsModel
	UserOperationResultsModel       UserOperationResultsModel
	UserPtsEventsModel              UserPtsEventsModel
	UserPtsStateModel               UserPtsStateModel
	UserupdatesPartitionFencesModel UserupdatesPartitionFencesModel
	UserupdatesQueries              UserupdatesQueriesModel
}

type TxModels struct {
	DeliveryFailedOperationsModel   DeliveryFailedOperationsTxModel
	DialogSideEffectOutboxModel     DialogSideEffectOutboxTxModel
	HashTagsModel                   HashTagsTxModel
	MessageReadOutboxModel          MessageReadOutboxTxModel
	PushTaskOutboxModel             PushTaskOutboxTxModel
	UserAuthSeqEventsModel          UserAuthSeqEventsTxModel
	UserAuthSeqStateModel           UserAuthSeqStateTxModel
	UserDialogsModel                UserDialogsTxModel
	UserMessageViewsModel           UserMessageViewsTxModel
	UserOperationResultsModel       UserOperationResultsTxModel
	UserPtsEventsModel              UserPtsEventsTxModel
	UserPtsStateModel               UserPtsStateTxModel
	UserupdatesPartitionFencesModel UserupdatesPartitionFencesTxModel
	UserupdatesQueries              UserupdatesQueriesTxModel
}

func NewModels(db *sqlx.DB) *Models {
	return &Models{
		DeliveryFailedOperationsModel:   NewDeliveryFailedOperationsModel(db),
		DialogSideEffectOutboxModel:     NewDialogSideEffectOutboxModel(db),
		HashTagsModel:                   NewHashTagsModel(db),
		MessageReadOutboxModel:          NewMessageReadOutboxModel(db),
		PushTaskOutboxModel:             NewPushTaskOutboxModel(db),
		UserAuthSeqEventsModel:          NewUserAuthSeqEventsModel(db),
		UserAuthSeqStateModel:           NewUserAuthSeqStateModel(db),
		UserDialogsModel:                NewUserDialogsModel(db),
		UserMessageViewsModel:           NewUserMessageViewsModel(db),
		UserOperationResultsModel:       NewUserOperationResultsModel(db),
		UserPtsEventsModel:              NewUserPtsEventsModel(db),
		UserPtsStateModel:               NewUserPtsStateModel(db),
		UserupdatesPartitionFencesModel: NewUserupdatesPartitionFencesModel(db),
		UserupdatesQueries:              NewUserupdatesQueriesModel(db),
	}
}

func (m *Models) WithTx(tx *sqlx.Tx) *TxModels {
	return &TxModels{
		DeliveryFailedOperationsModel:   NewDeliveryFailedOperationsTxModel(tx),
		DialogSideEffectOutboxModel:     NewDialogSideEffectOutboxTxModel(tx),
		HashTagsModel:                   NewHashTagsTxModel(tx),
		MessageReadOutboxModel:          NewMessageReadOutboxTxModel(tx),
		PushTaskOutboxModel:             NewPushTaskOutboxTxModel(tx),
		UserAuthSeqEventsModel:          NewUserAuthSeqEventsTxModel(tx),
		UserAuthSeqStateModel:           NewUserAuthSeqStateTxModel(tx),
		UserDialogsModel:                NewUserDialogsTxModel(tx),
		UserMessageViewsModel:           NewUserMessageViewsTxModel(tx),
		UserOperationResultsModel:       NewUserOperationResultsTxModel(tx),
		UserPtsEventsModel:              NewUserPtsEventsTxModel(tx),
		UserPtsStateModel:               NewUserPtsStateTxModel(tx),
		UserupdatesPartitionFencesModel: NewUserupdatesPartitionFencesTxModel(tx),
		UserupdatesQueries:              NewUserupdatesQueriesTxModel(tx),
	}
}
