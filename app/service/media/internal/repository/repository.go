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
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/service/media/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/service/media/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/service/media/internal/repository/rpc"
)

// Repository is the dependency container for repository instances.
type Repository struct {
	db                   *sqlx.DB
	model                *model.Models
	dfsClient            rpc.DfsMediaClient
	processorClient      rpc.MediaProcessorClient
	fileReferenceService *FileReferenceService
	fileReferenceTTL     time.Duration
	readLeaseSecret      []byte
	readLeaseTTL         time.Duration
}

// NewRepository creates a new Repository.
func NewRepository(c config.Config, dfsClient rpc.DfsMediaClient, processorClient rpc.MediaProcessorClient) *Repository {
	db := sqlx.NewMySQL(&c.Mysql)
	return &Repository{
		db:                   db,
		model:                model.NewModels(db),
		dfsClient:            dfsClient,
		processorClient:      processorClient,
		fileReferenceService: NewFileReferenceService([]byte(c.FileReference.Secret), time.Now),
		fileReferenceTTL:     time.Duration(c.FileReference.TTLSeconds) * time.Second,
		readLeaseSecret:      []byte(c.ReadLease.Secret),
		readLeaseTTL:         time.Duration(c.ReadLease.TTLSeconds) * time.Second,
	}
}
