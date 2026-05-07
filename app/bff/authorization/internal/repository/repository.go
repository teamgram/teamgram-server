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
	"fmt"
	"strings"

	"github.com/teamgram/teamgram-server/v2/app/bff/authorization/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/bff/internal/userprojection"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	authsessionclient "github.com/teamgram/teamgram-server/v2/app/service/authsession/client"
	userclient "github.com/teamgram/teamgram-server/v2/app/service/biz/user/client"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// Repository is the dependency container for repository instances.
type Repository struct {
	AuthsessionClient authsessionclient.AuthsessionClient
	UserClient        userclient.UserClient
}

// NewRepository creates a new Repository.
func NewRepository(c config.Config) *Repository {
	r := &Repository{}
	if hasRPCClientConfig(c.AuthsessionClient) {
		r.AuthsessionClient = authsessionclient.NewAuthsessionClient(authsessionclient.MustNewKitexClient(c.AuthsessionClient))
	}
	if hasRPCClientConfig(c.UserClient) {
		r.UserClient = userclient.NewUserClient(userclient.MustNewKitexClient(c.UserClient))
	}
	return r
}

func (r *Repository) BindAuthKeyUser(ctx context.Context, authKeyId int64, userId int64) error {
	if r.AuthsessionClient == nil {
		return fmt.Errorf("authorization repository: authsession client is not configured")
	}
	if _, err := r.AuthsessionClient.AuthsessionBindAuthKeyUser(ctx, &authsession.TLAuthsessionBindAuthKeyUser{
		AuthKeyId: authKeyId,
		UserId:    userId,
	}); err != nil {
		return fmt.Errorf("authorization repository: bind auth key user %d/%d: %w", authKeyId, userId, err)
	}
	return nil
}

func (r *Repository) UnbindAuthKeyUser(ctx context.Context, authKeyId int64, userId int64) error {
	if r.AuthsessionClient == nil {
		return fmt.Errorf("authorization repository: authsession client is not configured")
	}
	if _, err := r.AuthsessionClient.AuthsessionUnbindAuthKeyUser(ctx, &authsession.TLAuthsessionUnbindAuthKeyUser{
		AuthKeyId: authKeyId,
		UserId:    userId,
	}); err != nil {
		return fmt.Errorf("authorization repository: unbind auth key user %d/%d: %w", authKeyId, userId, err)
	}
	return nil
}

func (r *Repository) BindTempAuthKey(ctx context.Context, permAuthKeyId int64, nonce int64, expiresAt int32, encryptedMessage []byte) error {
	if r.AuthsessionClient == nil {
		return fmt.Errorf("authorization repository: authsession client is not configured")
	}
	if _, err := r.AuthsessionClient.AuthsessionBindTempAuthKey(ctx, &authsession.TLAuthsessionBindTempAuthKey{
		PermAuthKeyId:    permAuthKeyId,
		Nonce:            nonce,
		ExpiresAt:        expiresAt,
		EncryptedMessage: encryptedMessage,
	}); err != nil {
		if isEncryptedMessageInvalid(err) {
			return ErrEncryptedMessageInvalid
		}
		return fmt.Errorf("authorization repository: bind temp auth key %d: %w", permAuthKeyId, err)
	}
	return nil
}

func (r *Repository) GetUserByPhone(ctx context.Context, phone string) (*tg.ImmutableUser, error) {
	if r.UserClient == nil {
		return nil, fmt.Errorf("authorization repository: user client is not configured")
	}
	user, err := r.UserClient.UserGetImmutableUserByPhone(ctx, &userpb.TLUserGetImmutableUserByPhone{Phone: phone})
	if err != nil {
		if isUserNotFound(err) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("authorization repository: get user by phone %s: %w", phone, err)
	}
	if user == nil || user.User == nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (r *Repository) CreateUser(ctx context.Context, secretKeyId int64, phone string, countryCode string, firstName string, lastName string) (*tg.ImmutableUser, error) {
	if r.UserClient == nil {
		return nil, fmt.Errorf("authorization repository: user client is not configured")
	}
	user, err := r.UserClient.UserCreateNewUser(ctx, &userpb.TLUserCreateNewUser{
		SecretKeyId: secretKeyId,
		Phone:       phone,
		CountryCode: countryCode,
		FirstName:   firstName,
		LastName:    lastName,
	})
	if err != nil {
		return nil, fmt.Errorf("authorization repository: create user %s: %w", phone, err)
	}
	if user == nil || user.User == nil {
		return nil, fmt.Errorf("authorization repository: create user %s returned nil user data", phone)
	}
	return user, nil
}

func (r *Repository) ProjectSelfUser(ctx context.Context, userId int64) (tg.UserClazz, error) {
	if r.UserClient == nil {
		return nil, fmt.Errorf("authorization repository: user client is not configured")
	}
	if userId <= 0 {
		return nil, fmt.Errorf("authorization repository: project self user: invalid user id %d", userId)
	}
	users, err := userprojection.ProjectUsers(ctx, r.UserClient, userId, []int64{userId}, userprojection.MissingStoredReference)
	if err != nil {
		return nil, fmt.Errorf("authorization repository: project self user %d: %w", userId, err)
	}
	if len(users) == 0 {
		return nil, fmt.Errorf("authorization repository: project self user %d returned empty users", userId)
	}
	return users[0], nil
}

func isUserNotFound(err error) bool {
	return errors.Is(err, userpb.ErrUserNotFound) || errors.Is(err, ErrUserNotFound) || strings.Contains(err.Error(), userpb.ErrUserNotFound.Error())
}

func isEncryptedMessageInvalid(err error) bool {
	return errors.Is(err, authsession.ErrEncryptedMessageInvalid) ||
		errors.Is(err, ErrEncryptedMessageInvalid) ||
		strings.Contains(err.Error(), authsession.ErrEncryptedMessageInvalid.Error())
}

func hasRPCClientConfig(c kitex.RpcClientConf) bool {
	return len(c.Endpoints) > 0 || len(c.Target) > 0 || c.HasEtcd()
}
