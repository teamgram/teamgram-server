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
	"github.com/teamgram/marmota/pkg/stores/sqlc"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/internal/repository/model"
	repositoryrpc "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/internal/repository/rpc"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/plugin"
)

// Repository is the dependency container for repository instances.
type Repository struct {
	sqlc.CachedConn
	db          *sqlx.DB
	model       *model.Models
	mediaReader repositoryrpc.MediaReader
	plugin      plugin.ChatPlugin
}

// NewRepository creates a new Repository.
func NewRepository(c config.Config, mediaReader repositoryrpc.MediaReader, p plugin.ChatPlugin) *Repository {
	db := sqlx.NewMySQL(&c.Mysql)
	return &Repository{
		CachedConn:  sqlc.NewConn(db, c.Cache),
		db:          db,
		model:       model.NewModels(db),
		mediaReader: mediaReader,
		plugin:      p,
	}
}

func (r *Repository) Close() error {
	return nil
}
