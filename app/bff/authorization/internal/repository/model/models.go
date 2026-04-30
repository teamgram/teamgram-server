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
	AuthOpLogsModel AuthOpLogsModel
}

type TxModels struct {
	AuthOpLogsModel AuthOpLogsTxModel
}

func NewModels(db *sqlx.DB) *Models {
	return &Models{
		AuthOpLogsModel: NewAuthOpLogsModel(db),
	}
}

func (m *Models) WithTx(tx *sqlx.Tx) *TxModels {
	return &TxModels{
		AuthOpLogsModel: NewAuthOpLogsTxModel(tx),
	}
}
