// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package alloc

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

// DefaultSeqTable is the default name of the table backing the SeqStore.
//
// The DDL for this table lives in app/service/idgen/sql/seq_conversations.sql
// and must be applied to every database referenced by the idgen service's
// Mysql.DSN before the service is started — the allocator will fail at
// runtime with "table doesn't exist" otherwise.
const DefaultSeqTable = "seq_conversations"

// mysqlStore implements SeqStore on top of a MySQL row that holds
// (conversation_id, min_seq, max_seq). max_seq is the next id to be
// allocated.
type mysqlStore struct {
	db    *sqlx.DB
	table string
}

// NewMySQLStore returns a SeqStore backed by the default seq table.
//
// See app/service/idgen/sql/seq_conversations.sql for the required schema.
func NewMySQLStore(db *sqlx.DB) SeqStore {
	return NewMySQLStoreWithTable(db, DefaultSeqTable)
}

// NewMySQLStoreWithTable returns a SeqStore backed by a caller-provided
// table name. The table schema must expose at least (conversation_id PK,
// min_seq BIGINT, max_seq BIGINT); use app/service/idgen/sql/seq_conversations.sql
// as the canonical template.
func NewMySQLStoreWithTable(db *sqlx.DB, table string) SeqStore {
	return &mysqlStore{db: db, table: table}
}

// Malloc atomically advances max_seq by `size` and returns the start of the
// newly reserved [firstSeq, firstSeq+size) range.
//
// First-time inserts use INSERT IGNORE inside a transaction so two callers
// racing on the same cold key cannot conflict on the primary key: at most
// one INSERT succeeds, both callers then converge on the same SELECT FOR
// UPDATE row lock and serialize.
func (s *mysqlStore) Malloc(ctx context.Context, key string, size int64) (int64, error) {
	if size < 0 {
		return 0, ErrInvalidSize
	}
	if size == 0 {
		return s.GetMaxSeq(ctx, key)
	}

	var firstSeq int64
	err := s.db.Transact(ctx, func(tx *sqlx.Tx) error {
		if err := s.ensureRowTx(tx, key); err != nil {
			return err
		}
		curr, err := s.getMaxSeqTx(tx, key)
		if err != nil {
			return err
		}
		firstSeq = curr
		return s.setMaxSeqTx(tx, key, curr+size)
	})
	if err != nil {
		return 0, err
	}
	return firstSeq, nil
}

// GetMaxSeq returns the current max_seq, treating a missing row as 0.
func (s *mysqlStore) GetMaxSeq(ctx context.Context, key string) (int64, error) {
	var maxSeq int64
	query := fmt.Sprintf("select max_seq from %s where conversation_id = ? limit 1", s.table)
	if err := s.db.QueryRow(ctx, &maxSeq, query, key); err != nil {
		if isSQLNotFound(err) {
			return 0, nil
		}
		return 0, err
	}
	return maxSeq, nil
}

// SetMaxSeq forces max_seq to the given value but never lets it regress: a
// call with a value smaller than the current max_seq is a no-op (no rows
// affected). This protects against accidental rewinds that would create
// duplicate ids.
func (s *mysqlStore) SetMaxSeq(ctx context.Context, key string, seq int64) error {
	if seq < 0 {
		return fmt.Errorf("alloc: SetMaxSeq seq must be >= 0, got %d", seq)
	}
	// GREATEST keeps the larger of the existing and the new max_seq.
	// On INSERT (cold key) max_seq is set to seq directly.
	query := fmt.Sprintf(
		"insert into %s(conversation_id, min_seq, max_seq) values (?, 0, ?) "+
			"on duplicate key update max_seq = greatest(max_seq, values(max_seq))",
		s.table,
	)
	_, err := s.db.Exec(ctx, query, key, seq)
	return err
}

// ensureRowTx makes sure the row for key exists. Concurrent callers race on
// INSERT IGNORE; at most one wins, the others observe the row already exists
// and proceed without error.
func (s *mysqlStore) ensureRowTx(tx *sqlx.Tx, key string) error {
	query := fmt.Sprintf(
		"insert ignore into %s(conversation_id, min_seq, max_seq) values (?, 0, 0)",
		s.table,
	)
	_, err := tx.Exec(query, key)
	return err
}

func (s *mysqlStore) getMaxSeqTx(tx *sqlx.Tx, key string) (int64, error) {
	var maxSeq int64
	query := fmt.Sprintf("select max_seq from %s where conversation_id = ? for update", s.table)
	if err := tx.QueryRow(&maxSeq, query, key); err != nil {
		return 0, err
	}
	return maxSeq, nil
}

func (s *mysqlStore) setMaxSeqTx(tx *sqlx.Tx, key string, maxSeq int64) error {
	query := fmt.Sprintf("update %s set max_seq = ? where conversation_id = ?", s.table)
	_, err := tx.Exec(query, maxSeq, key)
	return err
}

func isSQLNotFound(err error) bool {
	return errors.Is(err, sql.ErrNoRows) || errors.Is(err, sqlx.ErrNotFound)
}
