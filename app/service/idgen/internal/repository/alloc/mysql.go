package alloc

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

type mysqlStore struct {
	db    *sqlx.DB
	table string
}

func NewMySQLStore(db *sqlx.DB) SeqStore {
	return NewMySQLStoreWithTable(db, "seq_conversations")
}

func NewMySQLStoreWithTable(db *sqlx.DB, table string) SeqStore {
	return &mysqlStore{db: db, table: table}
}

func (s *mysqlStore) Malloc(ctx context.Context, key string, size int64) (int64, error) {
	if size < 0 {
		return 0, ErrInvalidSize
	}
	if size == 0 {
		return s.GetMaxSeq(ctx, key)
	}

	var firstSeq int64
	err := s.db.Transact(ctx, func(tx *sqlx.Tx) error {
		maxSeq, err := s.getMaxSeqTx(tx, key)
		if err != nil {
			if !isSQLNotFound(err) {
				return err
			}
			if err := s.insertSeqTx(tx, key, size); err != nil {
				return err
			}
			firstSeq = 0
			return nil
		}

		firstSeq = maxSeq
		return s.setMaxSeqTx(tx, key, maxSeq+size)
	})
	if err != nil {
		return 0, err
	}
	return firstSeq, nil
}

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

func (s *mysqlStore) SetMaxSeq(ctx context.Context, key string, seq int64) error {
	query := fmt.Sprintf(
		"insert into %s(conversation_id, min_seq, max_seq) values (?, 0, ?) on duplicate key update max_seq = values(max_seq)",
		s.table,
	)
	_, err := s.db.Exec(ctx, query, key, seq)
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

func (s *mysqlStore) insertSeqTx(tx *sqlx.Tx, key string, maxSeq int64) error {
	query := fmt.Sprintf("insert into %s(conversation_id, min_seq, max_seq) values (?, 0, ?)", s.table)
	_, err := tx.Exec(query, key, maxSeq)
	return err
}

func (s *mysqlStore) setMaxSeqTx(tx *sqlx.Tx, key string, maxSeq int64) error {
	query := fmt.Sprintf("update %s set max_seq = ? where conversation_id = ?", s.table)
	_, err := tx.Exec(query, maxSeq, key)
	return err
}

func isSQLNotFound(err error) bool {
	return errors.Is(err, sql.ErrNoRows) || errors.Is(err, sqlx.ErrNotFound)
}
