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
	"fmt"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
)

// Repository is the dependency container for repository instances.
type Repository struct {
	db            *sqlx.DB
	models        *model.Models
	idgen         IDGenerator
	ownerInstance string
}

// NewRepository creates a new Repository.
func NewRepository(c config.Config) *Repository {
	var db *sqlx.DB
	if c.Mysql.DSN != "" {
		db = sqlx.NewMySQL(&c.Mysql)
	}
	owner := c.OwnerInstance
	if owner == "" {
		owner = "local-userupdates"
	}
	return NewForTest(db, unavailableIDGenerator{}, owner)
}

type IDGenerator interface {
	NextID(ctx context.Context) (int64, error)
}

type unavailableIDGenerator struct{}

func (unavailableIDGenerator) NextID(context.Context) (int64, error) {
	return 0, fmt.Errorf("%w: id generator unavailable", userupdates.ErrUserupdatesStorage)
}

func NewForTest(db *sqlx.DB, idgen IDGenerator, ownerInstance string) *Repository {
	if ownerInstance == "" {
		ownerInstance = "local-userupdates"
	}
	if idgen == nil {
		idgen = unavailableIDGenerator{}
	}
	var models *model.Models
	if db != nil {
		models = model.NewModels(db)
	}
	return &Repository{
		db:            db,
		models:        models,
		idgen:         idgen,
		ownerInstance: ownerInstance,
	}
}

// Close releases repository-owned clients.
func (r *Repository) Close() error {
	if r == nil {
		return nil
	}

	return nil
}

func (r *Repository) OwnerInstance() string {
	if r == nil || r.ownerInstance == "" {
		return "local-userupdates"
	}
	return r.ownerInstance
}

func (r *Repository) requireDB() (*sqlx.DB, error) {
	if r == nil || r.db == nil {
		return nil, fmt.Errorf("%w: mysql is not configured", userupdates.ErrUserupdatesStorage)
	}
	return r.db, nil
}

func storageError(op string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%w: %s: %w", userupdates.ErrUserupdatesStorage, op, err)
}
