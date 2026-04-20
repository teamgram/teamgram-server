/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2026 The Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package model

import (
	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

var _ DefaultHistoryTtlModel = (*customDefaultHistoryTtlModel)(nil)

type (
	// DefaultHistoryTtlModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDefaultHistoryTtlModel.
	DefaultHistoryTtlModel interface {
		default_history_ttlModel
		bizDefaultHistoryTtlModel
		extendDefaultHistoryTtlModel
	}

	customDefaultHistoryTtlModel struct {
		*defaultDefaultHistoryTtlModel
	}
)

// NewDefaultHistoryTtlModel returns a model for the database table.
func NewDefaultHistoryTtlModel(db *sqlx.DB) DefaultHistoryTtlModel {
	return &customDefaultHistoryTtlModel{
		defaultDefaultHistoryTtlModel: newDefaultHistoryTtlModel(db),
	}
}
