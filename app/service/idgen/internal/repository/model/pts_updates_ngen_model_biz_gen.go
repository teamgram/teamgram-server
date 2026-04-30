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
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is
var _ *sqlx.DB
var _ *sqlx.Tx

type bizPtsUpdatesNgenModel interface {
}

type PtsUpdatesNgenTxModel interface {
}

type defaultPtsUpdatesNgenTxModel struct {
	tx *sqlx.Tx
}

func NewPtsUpdatesNgenTxModel(tx *sqlx.Tx) PtsUpdatesNgenTxModel {
	return &defaultPtsUpdatesNgenTxModel{tx: tx}
}
