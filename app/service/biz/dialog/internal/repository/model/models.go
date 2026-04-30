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
	DialogFiltersModel DialogFiltersModel
	DialogsModel       DialogsModel
	DraftsModel        DraftsModel
	SavedDialogsModel  SavedDialogsModel
}

type TxModels struct {
	DialogFiltersModel DialogFiltersTxModel
	DialogsModel       DialogsTxModel
	DraftsModel        DraftsTxModel
	SavedDialogsModel  SavedDialogsTxModel
}

func NewModels(db *sqlx.DB) *Models {
	return &Models{
		DialogFiltersModel: NewDialogFiltersModel(db),
		DialogsModel:       NewDialogsModel(db),
		DraftsModel:        NewDraftsModel(db),
		SavedDialogsModel:  NewSavedDialogsModel(db),
	}
}

func (m *Models) WithTx(tx *sqlx.Tx) *TxModels {
	return &TxModels{
		DialogFiltersModel: NewDialogFiltersTxModel(tx),
		DialogsModel:       NewDialogsTxModel(tx),
		DraftsModel:        NewDraftsTxModel(tx),
		SavedDialogsModel:  NewSavedDialogsTxModel(tx),
	}
}
