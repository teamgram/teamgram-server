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

package alloc

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"regexp"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

const (
	MessageDataNGen       = "message_data_ngen"
	MessageBoxNGen        = "message_box_ngen"
	ChannelMessageBoxNGen = "channel_message_box_ngen"
	SeqUpdatesNGen        = "seq_updates_ngen"
	PtsUpdatesNGen        = "pts_updates_ngen"
	QtsUpdatesNGen        = "qts_updates_ngen"
	ChannelPtsUpdatesNGen = "channel_pts_updates_ngen"
	ScheduledNGen         = "scheduled_ngen"
	BotUpdatesNGen        = "bot_updates_ngen"
	StoryNGen             = "story_ngen"
	ChannelStoryNGen      = "channel_story_ngen"
)

var (
	ErrInvalidTable = errors.New("alloc: invalid seq table")
	seqTableRE      = regexp.MustCompile(`^[a-z0-9_]+$`)
	seqTables       = map[string]struct{}{
		MessageDataNGen:       {},
		MessageBoxNGen:        {},
		ChannelMessageBoxNGen: {},
		SeqUpdatesNGen:        {},
		PtsUpdatesNGen:        {},
		QtsUpdatesNGen:        {},
		ChannelPtsUpdatesNGen: {},
		ScheduledNGen:         {},
		BotUpdatesNGen:        {},
		StoryNGen:             {},
		ChannelStoryNGen:      {},
	}
)

// mysqlStore implements SeqStore on top of a MySQL row that holds
// (id, min_seq, max_seq). max_seq is the next id to be allocated.
type mysqlStore struct {
	db *sqlx.DB
}

// NewMySQLStore returns a SeqStore backed by caller-supplied table names.
func NewMySQLStore(db *sqlx.DB) SeqStore {
	return &mysqlStore{db: db}
}

func quoteSeqTable(table string) (string, error) {
	if _, ok := seqTables[table]; !ok || !seqTableRE.MatchString(table) {
		return "", fmt.Errorf("%w: %s", ErrInvalidTable, table)
	}
	return "`" + table + "`", nil
}

func getMaxSeqQuery(quotedTable string) string {
	return fmt.Sprintf("select max_seq from %s where id = ? limit 1", quotedTable)
}

func insertMaxSeqQuery(quotedTable string) string {
	return fmt.Sprintf(
		"insert into %s(id, min_seq, max_seq) values (?, 0, ?) "+
			"on duplicate key update max_seq = greatest(max_seq, values(max_seq))",
		quotedTable,
	)
}

func ensureRowQuery(quotedTable string) string {
	return fmt.Sprintf("insert ignore into %s(id, min_seq, max_seq) values (?, 0, 0)", quotedTable)
}

func lockMaxSeqQuery(quotedTable string) string {
	return fmt.Sprintf("select max_seq from %s where id = ? for update", quotedTable)
}

func updateMaxSeqQuery(quotedTable string) string {
	return fmt.Sprintf("update %s set max_seq = ? where id = ?", quotedTable)
}

// Malloc atomically advances max_seq by `size` and returns the start of the
// newly reserved [firstSeq, firstSeq+size) range.
//
// First-time inserts use INSERT IGNORE inside a transaction so two callers
// racing on the same cold key cannot conflict on the primary key: at most
// one INSERT succeeds, both callers then converge on the same SELECT FOR
// UPDATE row lock and serialize.
func (s *mysqlStore) Malloc(ctx context.Context, table string, id int64, size int64) (int64, error) {
	if size < 0 {
		return 0, ErrInvalidSize
	}
	if size == 0 {
		return s.GetMaxSeq(ctx, table, id)
	}
	quotedTable, err := quoteSeqTable(table)
	if err != nil {
		return 0, err
	}

	var firstSeq int64
	err = s.db.Transact(ctx, func(tx *sqlx.Tx) error {
		if err := s.ensureRowTx(tx, quotedTable, id); err != nil {
			return err
		}
		curr, err := s.getMaxSeqTx(tx, quotedTable, id)
		if err != nil {
			return err
		}
		firstSeq = curr
		return s.setMaxSeqTx(tx, quotedTable, id, curr+size)
	})
	if err != nil {
		return 0, err
	}
	return firstSeq, nil
}

// GetMaxSeq returns the current max_seq, treating a missing row as 0.
func (s *mysqlStore) GetMaxSeq(ctx context.Context, table string, id int64) (int64, error) {
	quotedTable, err := quoteSeqTable(table)
	if err != nil {
		return 0, err
	}
	var maxSeq int64
	query := getMaxSeqQuery(quotedTable)
	if err := s.db.QueryRow(ctx, &maxSeq, query, id); err != nil {
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
func (s *mysqlStore) SetMaxSeq(ctx context.Context, table string, id int64, seq int64) error {
	if seq < 0 {
		return fmt.Errorf("alloc: SetMaxSeq seq must be >= 0, got %d", seq)
	}
	quotedTable, err := quoteSeqTable(table)
	if err != nil {
		return err
	}
	// GREATEST keeps the larger of the existing and the new max_seq.
	// On INSERT (cold key) max_seq is set to seq directly.
	query := insertMaxSeqQuery(quotedTable)
	_, err = s.db.Exec(ctx, query, id, seq)
	return err
}

// ensureRowTx makes sure the row for key exists. Concurrent callers race on
// INSERT IGNORE; at most one wins, the others observe the row already exists
// and proceed without error.
func (s *mysqlStore) ensureRowTx(tx *sqlx.Tx, quotedTable string, id int64) error {
	query := ensureRowQuery(quotedTable)
	_, err := tx.Exec(query, id)
	return err
}

func (s *mysqlStore) getMaxSeqTx(tx *sqlx.Tx, quotedTable string, id int64) (int64, error) {
	var maxSeq int64
	query := lockMaxSeqQuery(quotedTable)
	if err := tx.QueryRow(&maxSeq, query, id); err != nil {
		return 0, err
	}
	return maxSeq, nil
}

func (s *mysqlStore) setMaxSeqTx(tx *sqlx.Tx, quotedTable string, id int64, maxSeq int64) error {
	query := updateMaxSeqQuery(quotedTable)
	_, err := tx.Exec(query, maxSeq, id)
	return err
}

func isSQLNotFound(err error) bool {
	return errors.Is(err, sql.ErrNoRows) || errors.Is(err, sqlx.ErrNotFound)
}
