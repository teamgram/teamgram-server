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
	"errors"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

var (
	ErrUserNotFound            = errors.New("authorization repository: user not found")
	ErrEncryptedMessageInvalid = errors.New("authorization repository: encrypted message invalid")
)

type (
	AuthorizationRepository interface {
		AuthsessionBinding
		UserDirectory
	}

	AuthsessionBinding interface {
		BindAuthKeyUser(ctx context.Context, authKeyId int64, userId int64) error
		UnbindAuthKeyUser(ctx context.Context, authKeyId int64, userId int64) error
		BindTempAuthKey(ctx context.Context, permAuthKeyId int64, nonce int64, expiresAt int32, encryptedMessage []byte) error
	}

	UserDirectory interface {
		GetUserByPhone(ctx context.Context, phone string) (*tg.ImmutableUser, error)
		CreateUser(ctx context.Context, secretKeyId int64, phone string, countryCode string, firstName string, lastName string) (*tg.ImmutableUser, error)
		ProjectSelfUser(ctx context.Context, userId int64) (tg.UserClazz, error)
		SetAuthorizationTTL(ctx context.Context, userId int64, ttl int32) error
	}
)
