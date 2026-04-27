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

package repository

import (
	"context"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/code/code"
)

// GetCachePhoneCode retrieves a phone code from cache.
func (r *Repository) GetCachePhoneCode(ctx context.Context, authKeyId int64, phone string) (*code.PhoneCodeTransaction, error) {
	txn, err := r.phoneCodeModel.GetPhoneCode(ctx, authKeyId, phone)
	if err != nil {
		return nil, fmt.Errorf("%w: get phone code cache: %w", code.ErrCodeStorage, err)
	}
	return txn, nil
}

// PutCachePhoneCode stores a phone code into cache.
func (r *Repository) PutCachePhoneCode(ctx context.Context, authKeyId int64, phone string, data *code.PhoneCodeTransaction) error {
	if err := r.phoneCodeModel.PutPhoneCode(ctx, authKeyId, phone, data); err != nil {
		return fmt.Errorf("%w: put phone code cache: %w", code.ErrCodeStorage, err)
	}
	return nil
}

// DeleteCachePhoneCode removes a phone code from cache.
func (r *Repository) DeleteCachePhoneCode(ctx context.Context, authKeyId int64, phone string) error {
	if err := r.phoneCodeModel.DeletePhoneCode(ctx, authKeyId, phone); err != nil {
		return fmt.Errorf("%w: delete phone code cache: %w", code.ErrCodeStorage, err)
	}
	return nil
}
