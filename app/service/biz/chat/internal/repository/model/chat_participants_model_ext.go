// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package model

import (
	"database/sql"
	"fmt"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

type (
	extendChatParticipantsModel interface {
		UpdateAdminRightsTx(tx *sqlx.Tx, participantType int32, adminRights int32, id int64) (rowsAffected int64, err error)
	}
)

func (m *customChatParticipantsModel) UpdateAdminRightsTx(tx *sqlx.Tx, participantType int32, adminRights int32, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chat_participants set participant_type = ?, admin_rights = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, participantType, adminRights, id)
	if err != nil {
		return 0, fmt.Errorf("chat_participants.UpdateAdminRightsTx exec: %w", err)
	}
	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("chat_participants.UpdateAdminRightsTx rows affected: %w", err)
	}
	return rowsAffected, nil
}
