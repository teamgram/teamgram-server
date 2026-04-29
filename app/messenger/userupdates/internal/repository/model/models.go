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
	PushTaskOutboxModel             PushTaskOutboxModel
	UserDialogsModel                UserDialogsModel
	UserMessageViewsModel           UserMessageViewsModel
	UserOperationResultsModel       UserOperationResultsModel
	UserPtsEventsModel              UserPtsEventsModel
	UserPtsStateModel               UserPtsStateModel
	UserupdatesPartitionFencesModel UserupdatesPartitionFencesModel
}

func NewModels(db *sqlx.DB) *Models {
	return &Models{
		DeliveryFailedOperationsModel:   NewDeliveryFailedOperationsModel(db),
		PushTaskOutboxModel:             NewPushTaskOutboxModel(db),
		UserDialogsModel:                NewUserDialogsModel(db),
		UserMessageViewsModel:           NewUserMessageViewsModel(db),
		UserOperationResultsModel:       NewUserOperationResultsModel(db),
		UserPtsEventsModel:              NewUserPtsEventsModel(db),
		UserPtsStateModel:               NewUserPtsStateModel(db),
		UserupdatesPartitionFencesModel: NewUserupdatesPartitionFencesModel(db),
	}
}
