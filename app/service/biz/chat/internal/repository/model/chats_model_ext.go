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
	extendChatsModel interface {
		InsertFullTx(tx *sqlx.Tx, data *Chats) (lastInsertId, rowsAffected int64, err error)
	}
)

func (m *defaultChatsModel) InsertFullTx(tx *sqlx.Tx, data *Chats) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chats(creator_user_id, access_hash, random_id, participant_count, title, about, photo_id, default_banned_rights, migrated_to_id, migrated_to_access_hash, available_reactions_type, available_reactions, deactivated, noforwards, ttl_period, version, `date`) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
		r     sql.Result
	)

	r, err = tx.Exec(
		query,
		data.CreatorUserId,
		data.AccessHash,
		data.RandomId,
		data.ParticipantCount,
		data.Title,
		data.About,
		data.PhotoId,
		data.DefaultBannedRights,
		data.MigratedToId,
		data.MigratedToAccessHash,
		data.AvailableReactionsType,
		data.AvailableReactions,
		data.Deactivated,
		data.Noforwards,
		data.TtlPeriod,
		data.Version,
		data.Date,
	)
	if err != nil {
		err = fmt.Errorf("chats.InsertFullTx exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("chats.InsertFullTx last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.InsertFullTx rows affected: %w", err)
	}
	return
}
