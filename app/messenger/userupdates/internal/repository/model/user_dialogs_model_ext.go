// Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
//  All rights reserved.
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
	"context"
	"errors"
	"fmt"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

type (
	extendUserDialogsModel interface {
		SelectByUserPeersProjection(ctx context.Context, userId int64, peerIdList []int64) ([]UserDialogs, error)
		SelectByUserCursorProjection(ctx context.Context, userId int64, topMessageDate string, topPeerSeq int64, peerType int32, peerId int64, limit int32) ([]UserDialogs, error)
	}
)

const userDialogsProjectionRows = "user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id, COALESCE(top_message_date, '1970-01-01 00:00:00.000000') AS top_message_date, top_message_status, read_inbox_max_peer_seq, read_outbox_max_peer_seq, unread_count, unread_mentions_count, unread_reactions_count, unread_mark, pinned_peer_seq, pinned_canonical_message_id, has_scheduled, available_min_peer_seq, hidden, COALESCE(deleted_at, '1970-01-01 00:00:00.000000') AS deleted_at, last_pts, COALESCE(last_pts_at, '1970-01-01 00:00:00.000000') AS last_pts_at, dialog_schema_version, dialog_payload"

func (m *customUserDialogsModel) SelectByUserPeersProjection(ctx context.Context, userId int64, peerIdList []int64) ([]UserDialogs, error) {
	if len(peerIdList) == 0 {
		return []UserDialogs{}, nil
	}

	query := fmt.Sprintf("select %s from user_dialogs where user_id = ? and peer_id in (%s)", userDialogsProjectionRows, sqlx.InInt64List(peerIdList))
	values := []UserDialogs{}
	if err := m.db.QueryRowsPartial(ctx, &values, query, userId); err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []UserDialogs{}, nil
		}
		return nil, fmt.Errorf("user_dialogs.SelectByUserPeersProjection: %w", err)
	}
	return values, nil
}

func (m *customUserDialogsModel) SelectByUserCursorProjection(ctx context.Context, userId int64, topMessageDate string, topPeerSeq int64, peerType int32, peerId int64, limit int32) ([]UserDialogs, error) {
	query := fmt.Sprintf("select %s from user_dialogs where user_id = ? and hidden = 0 and (? = '' or top_message_date < ? or (top_message_date = ? and top_peer_seq < ?) or (top_message_date = ? and top_peer_seq = ? and peer_type > ?) or (top_message_date = ? and top_peer_seq = ? and peer_type = ? and peer_id > ?)) order by top_message_date desc, top_peer_seq desc, peer_type asc, peer_id asc limit ?", userDialogsProjectionRows)
	values := []UserDialogs{}
	err := m.db.QueryRowsPartial(ctx, &values, query, userId, topMessageDate, topMessageDate, topMessageDate, topPeerSeq, topMessageDate, topPeerSeq, peerType, topMessageDate, topPeerSeq, peerType, peerId, limit)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []UserDialogs{}, nil
		}
		return nil, fmt.Errorf("user_dialogs.SelectByUserCursorProjection: %w", err)
	}
	return values, nil
}
