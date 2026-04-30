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
	"github.com/teamgram/marmota/pkg/stores/cache"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

type Models struct {
	AuthKeysModel  AuthKeysModel
	AuthUsersModel AuthUsersModel
	AuthsModel     AuthsModel
}

type TxModels struct {
	AuthKeysModel  AuthKeysTxModel
	AuthUsersModel AuthUsersTxModel
	AuthsModel     AuthsTxModel
}

func NewModels(db *sqlx.DB, c cache.CacheConf) *Models {
	return &Models{
		AuthKeysModel:  NewAuthKeysModel(db, c),
		AuthUsersModel: NewAuthUsersModel(db),
		AuthsModel:     NewAuthsModel(db),
	}
}

func (m *Models) WithTx(tx *sqlx.Tx) *TxModels {
	return &TxModels{
		AuthKeysModel:  NewAuthKeysTxModel(tx),
		AuthUsersModel: NewAuthUsersTxModel(tx),
		AuthsModel:     NewAuthsTxModel(tx),
	}
}
