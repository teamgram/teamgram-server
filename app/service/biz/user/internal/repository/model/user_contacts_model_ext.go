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
	"errors"
	"fmt"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

type (
	extendUserContactsModel interface {
		SelectContactTx(tx *sqlx.Tx, ownerUserID, contactUserID int64) (*UserContacts, error)
	}
)

func (m *defaultUserContactsModel) SelectContactTx(tx *sqlx.Tx, ownerUserID, contactUserID int64) (*UserContacts, error) {
	do := &UserContacts{}
	err := tx.QueryRowPartial(do,
		"select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, is_deleted, date2 from user_contacts where owner_user_id = ? and contact_user_id = ? and is_deleted = 0 limit 1",
		ownerUserID, contactUserID)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("user_contacts.SelectContactTx: %w", err)
	}
	return do, nil
}
