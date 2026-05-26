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
	DialogAuthSeqOutboxModel      DialogAuthSeqOutboxModel
	DialogDraftsModel             DialogDraftsModel
	DialogFilterTagsModel         DialogFilterTagsModel
	DialogFiltersModel            DialogFiltersModel
	DialogPeerPolicyModel         DialogPeerPolicyModel
	DialogPreferenceVersionsModel DialogPreferenceVersionsModel
	DialogPreferencesModel        DialogPreferencesModel
	DialogPublicUpdateOutboxModel DialogPublicUpdateOutboxModel
	DialogVisualSettingsModel     DialogVisualSettingsModel
	DialogsModel                  DialogsModel
	DraftsModel                   DraftsModel
	SavedDialogsModel             SavedDialogsModel
	DialogRepositoryQueries       DialogRepositoryQueriesModel
}

type TxModels struct {
	DialogAuthSeqOutboxModel      DialogAuthSeqOutboxTxModel
	DialogDraftsModel             DialogDraftsTxModel
	DialogFilterTagsModel         DialogFilterTagsTxModel
	DialogFiltersModel            DialogFiltersTxModel
	DialogPeerPolicyModel         DialogPeerPolicyTxModel
	DialogPreferenceVersionsModel DialogPreferenceVersionsTxModel
	DialogPreferencesModel        DialogPreferencesTxModel
	DialogPublicUpdateOutboxModel DialogPublicUpdateOutboxTxModel
	DialogVisualSettingsModel     DialogVisualSettingsTxModel
	DialogsModel                  DialogsTxModel
	DraftsModel                   DraftsTxModel
	SavedDialogsModel             SavedDialogsTxModel
	DialogRepositoryQueries       DialogRepositoryQueriesTxModel
}

func NewModels(db *sqlx.DB) *Models {
	return &Models{
		DialogAuthSeqOutboxModel:      NewDialogAuthSeqOutboxModel(db),
		DialogDraftsModel:             NewDialogDraftsModel(db),
		DialogFilterTagsModel:         NewDialogFilterTagsModel(db),
		DialogFiltersModel:            NewDialogFiltersModel(db),
		DialogPeerPolicyModel:         NewDialogPeerPolicyModel(db),
		DialogPreferenceVersionsModel: NewDialogPreferenceVersionsModel(db),
		DialogPreferencesModel:        NewDialogPreferencesModel(db),
		DialogPublicUpdateOutboxModel: NewDialogPublicUpdateOutboxModel(db),
		DialogVisualSettingsModel:     NewDialogVisualSettingsModel(db),
		DialogsModel:                  NewDialogsModel(db),
		DraftsModel:                   NewDraftsModel(db),
		SavedDialogsModel:             NewSavedDialogsModel(db),
		DialogRepositoryQueries:       NewDialogRepositoryQueriesModel(db),
	}
}

func (m *Models) WithTx(tx *sqlx.Tx) *TxModels {
	return &TxModels{
		DialogAuthSeqOutboxModel:      NewDialogAuthSeqOutboxTxModel(tx),
		DialogDraftsModel:             NewDialogDraftsTxModel(tx),
		DialogFilterTagsModel:         NewDialogFilterTagsTxModel(tx),
		DialogFiltersModel:            NewDialogFiltersTxModel(tx),
		DialogPeerPolicyModel:         NewDialogPeerPolicyTxModel(tx),
		DialogPreferenceVersionsModel: NewDialogPreferenceVersionsTxModel(tx),
		DialogPreferencesModel:        NewDialogPreferencesTxModel(tx),
		DialogPublicUpdateOutboxModel: NewDialogPublicUpdateOutboxTxModel(tx),
		DialogVisualSettingsModel:     NewDialogVisualSettingsTxModel(tx),
		DialogsModel:                  NewDialogsTxModel(tx),
		DraftsModel:                   NewDraftsTxModel(tx),
		SavedDialogsModel:             NewSavedDialogsTxModel(tx),
		DialogRepositoryQueries:       NewDialogRepositoryQueriesTxModel(tx),
	}
}
