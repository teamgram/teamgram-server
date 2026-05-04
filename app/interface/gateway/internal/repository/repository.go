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

package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	authsessionclient "github.com/teamgram/teamgram-server/v2/app/service/authsession/client"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// Repository is the dependency container for repository instances.
type Repository struct {
	AuthsessionClient AuthsessionClient
}

// NewRepository creates a new Repository.
func NewRepository(c config.Config) *Repository {
	r := &Repository{}
	if hasRPCClientConfig(c.AuthsessionClient) {
		r.AuthsessionClient = authsessionclient.NewAuthsessionClient(authsessionclient.MustNewKitexClient(c.AuthsessionClient))
	}
	return r
}

// Close releases repository-owned clients.
func (r *Repository) Close() error {
	if r == nil {
		return nil
	}

	return nil
}

func (r *Repository) QueryAuthKey(ctx context.Context, authKeyId int64) (*tg.AuthKeyInfo, error) {
	if r.AuthsessionClient == nil {
		return nil, fmt.Errorf("gateway repository: authsession client is not configured")
	}
	key, err := r.AuthsessionClient.AuthsessionQueryAuthKey(ctx, &authsession.TLAuthsessionQueryAuthKey{
		AuthKeyId: authKeyId,
	})
	if err != nil {
		return nil, fmt.Errorf("gateway repository: query auth key %d: %w", authKeyId, err)
	}
	if key == nil {
		return nil, fmt.Errorf("gateway repository: query auth key %d returned nil", authKeyId)
	}
	return key, nil
}

func (r *Repository) SetAuthKey(ctx context.Context, authKey *tg.AuthKeyInfo, futureSalt *tg.FutureSalt, expiresIn int32) error {
	if r.AuthsessionClient == nil {
		return fmt.Errorf("gateway repository: authsession client is not configured")
	}
	if authKey == nil {
		return fmt.Errorf("gateway repository: set auth key: auth key is nil")
	}
	if futureSalt == nil {
		return fmt.Errorf("gateway repository: set auth key %d: future salt is nil", authKey.AuthKeyId)
	}
	if _, err := r.AuthsessionClient.AuthsessionSetAuthKey(ctx, &authsession.TLAuthsessionSetAuthKey{
		AuthKey:    authKey,
		FutureSalt: futureSalt,
		ExpiresIn:  expiresIn,
	}); err != nil {
		return fmt.Errorf("gateway repository: set auth key %d: %w", authKey.AuthKeyId, err)
	}
	return nil
}

func (r *Repository) GetFutureSalts(ctx context.Context, authKeyId int64, num int32) (*tg.FutureSalts, error) {
	if r.AuthsessionClient == nil {
		return nil, fmt.Errorf("gateway repository: authsession client is not configured")
	}
	salts, err := r.AuthsessionClient.AuthsessionGetFutureSalts(ctx, &authsession.TLAuthsessionGetFutureSalts{
		AuthKeyId: authKeyId,
		Num:       num,
	})
	if err != nil {
		return nil, fmt.Errorf("gateway repository: get future salts %d: %w", authKeyId, err)
	}
	if salts == nil {
		return nil, fmt.Errorf("gateway repository: get future salts %d returned nil", authKeyId)
	}
	return salts, nil
}

func (r *Repository) GetUserId(ctx context.Context, authKeyId int64) (int64, error) {
	if r.AuthsessionClient == nil {
		return 0, fmt.Errorf("gateway repository: authsession client is not configured")
	}
	userID, err := r.AuthsessionClient.AuthsessionGetUserId(ctx, &authsession.TLAuthsessionGetUserId{
		AuthKeyId: authKeyId,
	})
	if err != nil {
		return 0, fmt.Errorf("gateway repository: get user id %d: %w", authKeyId, err)
	}
	if userID == nil {
		return 0, fmt.Errorf("gateway repository: get user id %d returned nil", authKeyId)
	}
	return userID.V, nil
}

func (r *Repository) SetClientSessionInfo(ctx context.Context, session *authsession.ClientSession) error {
	if r.AuthsessionClient == nil {
		return fmt.Errorf("gateway repository: authsession client is not configured")
	}
	if session == nil {
		return fmt.Errorf("gateway repository: set client session info: session is nil")
	}
	if _, err := r.AuthsessionClient.AuthsessionSetClientSessionInfo(ctx, &authsession.TLAuthsessionSetClientSessionInfo{
		Data: session,
	}); err != nil {
		return fmt.Errorf("gateway repository: set client session info %d: %w", session.AuthKeyId, err)
	}
	return nil
}

func (r *Repository) SetLayer(ctx context.Context, authKeyId int64, ip string, layer int32) error {
	if r.AuthsessionClient == nil {
		return fmt.Errorf("gateway repository: authsession client is not configured")
	}
	if _, err := r.AuthsessionClient.AuthsessionSetLayer(ctx, &authsession.TLAuthsessionSetLayer{
		AuthKeyId: authKeyId,
		Ip:        ip,
		Layer:     layer,
	}); err != nil {
		return fmt.Errorf("gateway repository: set layer %d: %w", authKeyId, err)
	}
	return nil
}

func (r *Repository) GetClientSession(ctx context.Context, authKeyId int64) (*authsession.ClientSession, error) {
	if r.AuthsessionClient == nil {
		return nil, fmt.Errorf("gateway repository: authsession client is not configured")
	}
	state, err := r.AuthsessionClient.AuthsessionGetAuthStateData(ctx, &authsession.TLAuthsessionGetAuthStateData{
		AuthKeyId: authKeyId,
	})
	if err != nil {
		if isPermAuthKeyEmpty(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("gateway repository: get client session %d: %w", authKeyId, err)
	}
	if state == nil || state.Client == nil {
		return nil, nil
	}
	return state.Client, nil
}

func hasRPCClientConfig(c kitex.RpcClientConf) bool {
	return len(c.Endpoints) > 0 || len(c.Target) > 0 || c.HasEtcd()
}

func isPermAuthKeyEmpty(err error) bool {
	return errors.Is(err, authsession.ErrPermAuthKeyEmpty) ||
		strings.Contains(err.Error(), authsession.ErrPermAuthKeyEmpty.Error())
}
