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
	PtsUpdatesNgenModel PtsUpdatesNgenModel
	IdgenNgenQueries    IdgenNgenQueriesModel
}

type TxModels struct {
	PtsUpdatesNgenModel PtsUpdatesNgenTxModel
	IdgenNgenQueries    IdgenNgenQueriesTxModel
}

func NewModels(db *sqlx.DB) *Models {
	return &Models{
		PtsUpdatesNgenModel: NewPtsUpdatesNgenModel(db),
		IdgenNgenQueries:    NewIdgenNgenQueriesModel(db),
	}
}

func (m *Models) WithTx(tx *sqlx.Tx) *TxModels {
	return &TxModels{
		PtsUpdatesNgenModel: NewPtsUpdatesNgenTxModel(tx),
		IdgenNgenQueries:    NewIdgenNgenQueriesTxModel(tx),
	}
}
