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

// Package repository owns aggregate persistence orchestration for idgen.
//
// Two persistence sources back the service:
//
//   - a snowflake.Node that produces process-local 64-bit ids (used by
//     idgen.nextId / idgen.nextIds and the scalar branches of
//     idgen.getNextIdValList).
//   - an alloc.Allocator that maintains per-(table,id) monotonic sequences
//     in MySQL with an optional Redis segment cache (used by all sequence
//     operations).
//
// Repository hides those internal dependencies behind semantic methods such
// as NextID, ReserveNextSeqID, QueryCurrentSeqID, ResetSeqID. internal/core
// depends only on these methods and on the public service errors exported
// from app/service/idgen/idgen.
package repository

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/kv"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	idgenpb "github.com/teamgram/teamgram-server/v2/app/service/idgen/idgen"
	"github.com/teamgram/teamgram-server/v2/app/service/idgen/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/service/idgen/internal/repository/alloc"

	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/core/stores/cache"
)

// Repository aggregates the persistence-facing dependencies for idgen and
// exposes semantic methods to internal/core. Fields are intentionally
// unexported so callers cannot reach into the underlying allocator or
// snowflake node directly.
type Repository struct {
	node     *snowflake.Node
	seqAlloc *alloc.Allocator
}

// New constructs a Repository from prebuilt dependencies. seqAlloc may be
// nil to indicate a degraded mode in which sequence operations return
// idgenpb.ErrSeqAllocatorUnavailable. Mainly intended for tests; production
// wiring should use NewRepository.
func New(node *snowflake.Node, seqAlloc *alloc.Allocator) *Repository {
	return &Repository{node: node, seqAlloc: seqAlloc}
}

// NewRepository builds a Repository from configuration. It reports an error
// rather than panicking on invalid configuration so the caller (Server.
// Initialize) can fail loudly at startup.
//
// Configuration drives a one-step degradation:
//
//	If KV nodes are missing or have zero total weight, the allocator is
//	wired without a cache (DB-direct mode). This keeps correctness intact
//	but caps single-key throughput at MySQL's row-lock QPS (~1k); only
//	suitable for low-frequency keys or local development.
//
// An empty Mysql.DSN is rejected: idgen has no useful operating mode
// without a backing store, and silently nil'ing the allocator would let
// callers discover the misconfiguration only on the first sequence call.
func NewRepository(c config.Config) (*Repository, error) {
	node, err := snowflake.NewNode(c.NodeId)
	if err != nil {
		return nil, fmt.Errorf("idgen repository: new snowflake node (node_id=%d): %w", c.NodeId, err)
	}

	if c.Mysql.DSN == "" {
		return nil, fmt.Errorf("idgen repository: Mysql.DSN is required")
	}

	store := alloc.NewMySQLStore(sqlx.NewMySQL(&c.Mysql))
	var seqAlloc *alloc.Allocator
	if len(c.KV) == 0 || cache.TotalWeights(c.KV) <= 0 {
		seqAlloc = alloc.NewAllocator(nil, store)
	} else {
		seqAlloc = alloc.NewAllocator(alloc.NewXKVCache(kv.NewStore(c.KV)), store)
	}

	return New(node, seqAlloc), nil
}

// NextID returns a fresh snowflake id.
func (r *Repository) NextID() int64 {
	return r.node.Generate().Int64()
}

// NextIDs returns num fresh snowflake ids. The caller is responsible for
// validating num against any service-level limits before calling.
func (r *Repository) NextIDs(num int32) []int64 {
	if num <= 0 {
		return []int64{}
	}
	ids := make([]int64, num)
	for i := int32(0); i < num; i++ {
		ids[i] = r.node.Generate().Int64()
	}
	return ids
}

// QueryCurrentSeqID returns the current next-to-be-allocated seq value for
// the given key. The key must be of the form "<seq_table>_<id>" where
// <seq_table> is one of alloc.SeqTables().
//
// Returns idgenpb.ErrSeqAllocatorUnavailable when the allocator is not
// configured, idgenpb.ErrInvalidArgument for malformed keys, and
// idgenpb.ErrSeqStorage wrapping the underlying storage error otherwise.
func (r *Repository) QueryCurrentSeqID(ctx context.Context, key string) (int64, error) {
	if r.seqAlloc == nil {
		return 0, idgenpb.ErrSeqAllocatorUnavailable
	}
	table, id, err := parseSeqKey(key)
	if err != nil {
		return 0, err
	}
	seq, err := r.seqAlloc.GetMaxSeq(ctx, table, id)
	if err != nil {
		return 0, fmt.Errorf("%w: get max seq: %w", idgenpb.ErrSeqStorage, err)
	}
	return seq, nil
}

// ReserveNextSeqID atomically advances the seq cursor for key by n and
// returns the new cursor value (i.e. the highest seq id assigned to the
// caller's batch). n must be > 0.
//
// Returns idgenpb.ErrSeqAllocatorUnavailable when the allocator is not
// configured, idgenpb.ErrInvalidArgument when n <= 0 or the key is
// malformed, and idgenpb.ErrSeqStorage wrapping the underlying storage
// error otherwise.
func (r *Repository) ReserveNextSeqID(ctx context.Context, key string, n int32) (int64, error) {
	if r.seqAlloc == nil {
		return 0, idgenpb.ErrSeqAllocatorUnavailable
	}
	if n <= 0 {
		return 0, fmt.Errorf("%w: seq n %d must be > 0", idgenpb.ErrInvalidArgument, n)
	}
	table, id, err := parseSeqKey(key)
	if err != nil {
		return 0, err
	}
	start, err := r.seqAlloc.Malloc(ctx, table, id, int64(n))
	if err != nil {
		return 0, fmt.Errorf("%w: malloc seq: %w", idgenpb.ErrSeqStorage, err)
	}
	return start + int64(n), nil
}

// ResetSeqID forces the persisted max_seq for key to seq. The underlying
// store rejects regressions so a value smaller than the current max_seq is
// a no-op.
//
// Returns the same error contract as QueryCurrentSeqID / ReserveNextSeqID.
func (r *Repository) ResetSeqID(ctx context.Context, key string, seq int64) error {
	if r.seqAlloc == nil {
		return idgenpb.ErrSeqAllocatorUnavailable
	}
	table, id, err := parseSeqKey(key)
	if err != nil {
		return err
	}
	if err := r.seqAlloc.SetMaxSeq(ctx, table, id, seq); err != nil {
		return fmt.Errorf("%w: set max seq: %w", idgenpb.ErrSeqStorage, err)
	}
	return nil
}

// parseSeqKey decomposes a "<seq_table>_<id>" key into its table and id
// parts. The table must appear in alloc.SeqTables(); unknown prefixes and
// non-integer ids are rejected with idgenpb.ErrInvalidArgument so callers
// cannot trigger arbitrary allocator routes.
func parseSeqKey(key string) (string, int64, error) {
	for _, table := range alloc.SeqTables() {
		prefix := table + "_"
		if !strings.HasPrefix(key, prefix) {
			continue
		}
		id, err := strconv.ParseInt(strings.TrimPrefix(key, prefix), 10, 64)
		if err != nil {
			return "", 0, fmt.Errorf("%w: parse seq key %q: %v", idgenpb.ErrInvalidArgument, key, err)
		}
		return table, id, nil
	}
	return "", 0, fmt.Errorf("%w: invalid seq key %q", idgenpb.ErrInvalidArgument, key)
}
