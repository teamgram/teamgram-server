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
	"errors"
	"fmt"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	dialogpb "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository/model"
)

// Repository is the dependency container for repository instances.
type Repository struct {
	db    *sqlx.DB
	model *model.Models
}

// NewRepository creates a new Repository.
func NewRepository(c config.Config) *Repository {
	db := sqlx.NewMySQL(&c.Mysql)
	return &Repository{
		db:    db,
		model: model.NewModels(db),
	}
}

func NewRepositoryForTest(models *model.Models) *Repository {
	return &Repository{
		model: models,
	}
}

func NewRepositoryWithDBForTest(db *sqlx.DB) *Repository {
	return &Repository{
		db:    db,
		model: model.NewModels(db),
	}
}

func (r *Repository) requireDB() (*sqlx.DB, error) {
	if r == nil || r.db == nil {
		return nil, dialogpb.WrapDialogStorage("require db", errors.New("dialog mysql is not configured"))
	}
	return r.db, nil
}

func (r *Repository) requireModels() (*model.Models, error) {
	if r == nil || r.model == nil {
		return nil, dialogpb.WrapDialogStorage("require models", errors.New("dialog models are not configured"))
	}
	return r.model, nil
}

func storageError(op string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%w: %s: %w", dialogpb.ErrDialogStorage, op, err)
}
