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
	"github.com/teamgram/teamgram-server/v2/app/service/idgen/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/service/idgen/internal/repository/alloc"

	"github.com/teamgram/marmota/pkg/stores/kv"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stores/cache"
)

// Repository is the dependency container for repository instances.
type Repository struct {
	SeqAlloc *alloc.Allocator
}

// NewRepository creates a new Repository.
//
// Configuration drives a two-step degradation:
//
//  1. If c.Mysql.DSN is empty, no SeqStore can be built. SeqAlloc is left nil
//     and any call site that depends on it must guard with a nil check —
//     callers should fail loudly rather than silently mis-allocate.
//  2. If KV nodes are missing or have zero total weight, the allocator is
//     wired without a cache (DB-direct mode). This keeps correctness intact
//     but caps single-key throughput at MySQL's row-lock QPS (~1k); only
//     suitable for low-frequency keys or local development.
func NewRepository(c config.Config) *Repository {
	r := &Repository{}
	if c.Mysql.DSN == "" {
		return r
	}

	store := alloc.NewMySQLStore(sqlx.NewMySQL(&c.Mysql))
	if len(c.KV) == 0 || cache.TotalWeights(c.KV) <= 0 {
		r.SeqAlloc = alloc.NewAllocator(nil, store)
		return r
	}

	r.SeqAlloc = alloc.NewAllocator(alloc.NewXKVCache(kv.NewStore(c.KV)), store)
	return r
}
