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

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	idgenclient "github.com/teamgram/teamgram-server/v2/app/service/idgen/client"
	"github.com/teamgram/teamgram-server/v2/app/service/idgen/idgen"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

// Repository is the dependency container for repository instances.
type Repository struct {
	db     *sqlx.DB
	models *model.Models
	idgen  IDGenerator
}

// NewRepository creates a new Repository.
func NewRepository(c config.Config) *Repository {
	var db *sqlx.DB
	if c.Mysql.DSN != "" {
		db = sqlx.NewMySQL(&c.Mysql)
	}
	var idgen IDGenerator = unavailableIDGenerator{}
	if hasRPCClientConfig(c.Idgen) {
		idgen = &idgenRPCGenerator{
			client: idgenclient.NewIdgenClient(idgenclient.MustNewKitexClient(c.Idgen)),
		}
	}
	return NewForTest(db, idgen)
}

type IDGenerator interface {
	NextID(ctx context.Context) (int64, error)
}

type unavailableIDGenerator struct{}

func (unavailableIDGenerator) NextID(context.Context) (int64, error) {
	return 0, errors.New("idgen client unavailable")
}

type idgenRPCGenerator struct {
	client idgenclient.IdgenClient
}

func (g *idgenRPCGenerator) NextID(ctx context.Context) (int64, error) {
	if g == nil || g.client == nil {
		return 0, errors.New("idgen client unavailable")
	}
	id, err := g.client.IdgenNextId(ctx, &idgen.TLIdgenNextId{})
	if err != nil {
		return 0, err
	}
	if id == nil {
		return 0, errors.New("idgen returned nil")
	}
	return id.V, nil
}

func NewForTest(db *sqlx.DB, idgen IDGenerator) *Repository {
	if idgen == nil {
		idgen = unavailableIDGenerator{}
	}
	var models *model.Models
	if db != nil {
		models = model.NewModels(db)
	}
	return &Repository{
		db:     db,
		models: models,
		idgen:  idgen,
	}
}

// Close releases repository-owned clients.
func (r *Repository) Close() error {
	if r == nil {
		return nil
	}

	return nil
}

func (r *Repository) requireDB() (*sqlx.DB, error) {
	if r == nil || r.db == nil {
		return nil, storageError("require mysql", errors.New("mysql is not configured"))
	}
	return r.db, nil
}

func (r *Repository) nextID(ctx context.Context, op string) (int64, error) {
	if r == nil || r.idgen == nil {
		return 0, storageError(op, errors.New("idgen is not configured"))
	}
	id, err := r.idgen.NextID(ctx)
	if err != nil {
		return 0, storageError(op, err)
	}
	return id, nil
}

func storageError(op string, err error) error {
	return fmt.Errorf("%w: %s: %v", msg.ErrMsgStorage, op, err)
}

func hasRPCClientConfig(c kitex.RpcClientConf) bool {
	if c.DestService == "" {
		return false
	}
	return len(c.Endpoints) > 0 || c.Target != "" || c.HasEtcd()
}
