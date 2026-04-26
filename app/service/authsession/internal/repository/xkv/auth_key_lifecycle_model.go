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

package xkv

import (
	"context"
	"errors"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/kv"
)

// AuthKeyLifecycleModel tracks the validity window of a temporary auth key.
//
// MTProto temp / media-temp keys must expire after the negotiated duration.
// The auth_keys table itself does not carry an expires_at column, so we
// piggyback on the kv store: the presence of the lifecycle entry means the
// key is still alive, and the kv TTL guarantees automatic cleanup. Permanent
// keys are not tracked — they are valid until explicitly revoked.
type AuthKeyLifecycleModel interface {
	// Activate marks the key as alive for the given TTL (in seconds).
	// A non-positive TTL is rejected to avoid accidentally writing a key
	// that immediately disappears.
	Activate(ctx context.Context, authKeyId int64, ttlSeconds int) error
	// IsActive reports whether the key still has a valid lifecycle entry.
	IsActive(ctx context.Context, authKeyId int64) (bool, error)
	// Revoke removes the lifecycle entry, which lazily evicts subsequent
	// queries for the key.
	Revoke(ctx context.Context, authKeyId int64) error
}

// ErrInvalidTTL is returned when the caller asks to activate a key with a
// non-positive TTL.
var ErrInvalidTTL = errors.New("xkv: invalid auth key ttl")

const authKeyLifecycleValue = "1"

type authKeyLifecycleModel struct {
	kv     kv.Store
	prefix string
}

// NewAuthKeyLifecycleModel builds the kv-backed lifecycle model.
func NewAuthKeyLifecycleModel(store kv.Store, prefix string) AuthKeyLifecycleModel {
	return &authKeyLifecycleModel{
		kv:     store,
		prefix: prefix,
	}
}

func (m *authKeyLifecycleModel) cacheKey(authKeyId int64) string {
	if m.prefix == "" {
		return fmt.Sprintf("auth_key_ttl#%d", authKeyId)
	}
	return fmt.Sprintf("%s:auth_key_ttl#%d", m.prefix, authKeyId)
}

func (m *authKeyLifecycleModel) Activate(ctx context.Context, authKeyId int64, ttlSeconds int) error {
	if ttlSeconds <= 0 {
		return ErrInvalidTTL
	}
	return m.kv.SetexCtx(ctx, m.cacheKey(authKeyId), authKeyLifecycleValue, ttlSeconds)
}

func (m *authKeyLifecycleModel) IsActive(ctx context.Context, authKeyId int64) (bool, error) {
	val, err := m.kv.GetCtx(ctx, m.cacheKey(authKeyId))
	if err != nil {
		return false, err
	}
	return val == authKeyLifecycleValue, nil
}

func (m *authKeyLifecycleModel) Revoke(ctx context.Context, authKeyId int64) error {
	_, err := m.kv.DelCtx(ctx, m.cacheKey(authKeyId))
	return err
}
