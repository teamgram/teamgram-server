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
	"context"
	"errors"
	"fmt"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

type (
	extendUsersModel interface {
		SelectProfilePhotoTx(tx *sqlx.Tx, id int64) (int64, error)
		UpdatePhone(ctx context.Context, phone string, id int64) (rowsAffected int64, err error)
		UpdatePremium(ctx context.Context, premium bool, premiumExpireDate int64, updateExpireDate bool, id int64) (rowsAffected int64, err error)
		UpdateVerified(ctx context.Context, verified bool, id int64) (rowsAffected int64, err error)
	}
)

func (m *defaultUsersModel) SelectProfilePhotoTx(tx *sqlx.Tx, id int64) (int64, error) {
	var photoID int64
	err := tx.QueryRowPartial(&photoID, "select photo_id from users where id = ? limit 1", id)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return 0, nil
		}
		return 0, fmt.Errorf("users.SelectProfilePhotoTx: %w", err)
	}
	return photoID, nil
}

func (m *defaultUsersModel) UpdatePhone(ctx context.Context, phone string, id int64) (rowsAffected int64, err error) {
	result, err := m.db.Exec(ctx, "update users set phone = ? where id = ?", phone, id)
	if err != nil {
		return 0, fmt.Errorf("users.UpdatePhone exec: %w", err)
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("users.UpdatePhone rows affected: %w", err)
	}
	return rowsAffected, nil
}

func (m *defaultUsersModel) UpdatePremium(ctx context.Context, premium bool, premiumExpireDate int64, updateExpireDate bool, id int64) (rowsAffected int64, err error) {
	query := "update users set premium = ? where id = ?"
	args := []interface{}{premium, id}
	if updateExpireDate {
		query = "update users set premium = ?, premium_expire_date = ? where id = ?"
		args = []interface{}{premium, premiumExpireDate, id}
	}
	result, err := m.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, fmt.Errorf("users.UpdatePremium exec: %w", err)
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("users.UpdatePremium rows affected: %w", err)
	}
	return rowsAffected, nil
}

func (m *defaultUsersModel) UpdateVerified(ctx context.Context, verified bool, id int64) (rowsAffected int64, err error) {
	result, err := m.db.Exec(ctx, "update users set verified = ? where id = ?", verified, id)
	if err != nil {
		return 0, fmt.Errorf("users.UpdateVerified exec: %w", err)
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("users.UpdateVerified rows affected: %w", err)
	}
	return rowsAffected, nil
}
