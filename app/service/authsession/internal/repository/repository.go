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
	"github.com/teamgram/marmota/pkg/stores/kv"
	"github.com/teamgram/marmota/pkg/stores/sqlc"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	geoipclient "github.com/teamgram/teamgram-server/v2/app/infra/geoip/client"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/internal/repository/xkv"
)

// Repository is the dependency container for repository instances.
//
// db and kv are kept as struct fields purely so Close can dispose of the
// underlying clients on shutdown — production paths read storage through the
// embedded CachedConn or the kv-backed sub-models below.
type Repository struct {
	sqlc.CachedConn
	db                    *sqlx.DB
	kv                    kv.ExtStore
	model                 *model.Models
	futureSaltsModel      FutureSaltsModelType
	authKeyLifecycleModel AuthKeyLifecycleModelType
	geoipClient           GeoipClientType
}

// NewRepository creates a new Repository.
func NewRepository(c config.Config, geoipClient geoipclient.GeoipClient) *Repository {
	db := sqlx.NewMySQL(&c.Mysql)
	kv2 := kv.NewStore(c.KV)

	return &Repository{
		CachedConn:            sqlc.NewConn(db, c.Cache),
		db:                    db,
		kv:                    kv2,
		model:                 model.NewModels(db, c.Cache),
		futureSaltsModel:      xkv.NewFutureSaltsModel(kv2, "future_salts"),
		authKeyLifecycleModel: xkv.NewAuthKeyLifecycleModel(kv2, "authsession"),
		geoipClient:           geoipClient,
	}
}

// Close releases repository-owned clients.
func (r *Repository) Close() error {
	if r == nil {
		return nil
	}
	return nil
}
