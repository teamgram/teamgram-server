/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is
var _ *sqlx.DB
var _ *sqlx.Tx

type bizFileReferencesModel interface {
	Insert(ctx context.Context, data *FileReferences) (lastInsertId, rowsAffected int64, err error)
	SelectByRefHash(ctx context.Context, refHash []byte) (*FileReferences, error)
}

type FileReferencesTxModel interface {
	Insert(data *FileReferences) (lastInsertId, rowsAffected int64, err error)
	SelectByRefHash(refHash []byte) (*FileReferences, error)
}

type defaultFileReferencesTxModel struct {
	tx *sqlx.Tx
}

func NewFileReferencesTxModel(tx *sqlx.Tx) FileReferencesTxModel {
	return &defaultFileReferencesTxModel{tx: tx}
}

// Insert
// insert into file_references(ref_hash, domain, media_id, access_hash, object_id, origin_domain, origin_id, expire_at, revoked_at) values (:ref_hash, :domain, :media_id, :access_hash, :object_id, :origin_domain, :origin_id, :expire_at, :revoked_at)
func (m *defaultFileReferencesModel) Insert(ctx context.Context, data *FileReferences) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into file_references(ref_hash, domain, media_id, access_hash, object_id, origin_domain, origin_id, expire_at, revoked_at) values (:ref_hash, :domain, :media_id, :access_hash, :object_id, :origin_domain, :origin_id, :expire_at, :revoked_at)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("file_references.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("file_references.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("file_references.Insert rows affected: %w", err)
	}

	return

}

// Insert
// insert into file_references(ref_hash, domain, media_id, access_hash, object_id, origin_domain, origin_id, expire_at, revoked_at) values (:ref_hash, :domain, :media_id, :access_hash, :object_id, :origin_domain, :origin_id, :expire_at, :revoked_at)
func (m *defaultFileReferencesTxModel) Insert(data *FileReferences) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into file_references(ref_hash, domain, media_id, access_hash, object_id, origin_domain, origin_id, expire_at, revoked_at) values (:ref_hash, :domain, :media_id, :access_hash, :object_id, :origin_domain, :origin_id, :expire_at, :revoked_at)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("file_references.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("file_references.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("file_references.Insert rows affected: %w", err)
	}

	return
}

// SelectByRefHash
// select id, ref_hash, domain, media_id, access_hash, object_id, origin_domain, origin_id, expire_at, revoked_at, created_at from file_references where ref_hash = :ref_hash limit 1
func (m *defaultFileReferencesModel) SelectByRefHash(ctx context.Context, refHash []byte) (rValue *FileReferences, err error) {

	var (
		query = "select id, ref_hash, domain, media_id, access_hash, object_id, origin_domain, origin_id, expire_at, revoked_at, created_at from file_references where ref_hash = ? limit 1"
		do    = &FileReferences{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, refHash)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "file_references",
				Key:      fmt.Sprintf("ref_hash=%v", refHash),
				Cause:    err,
			}
		}
		err = fmt.Errorf("file_references.SelectByRefHash: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByRefHash
// select id, ref_hash, domain, media_id, access_hash, object_id, origin_domain, origin_id, expire_at, revoked_at, created_at from file_references where ref_hash = :ref_hash limit 1
func (m *defaultFileReferencesTxModel) SelectByRefHash(refHash []byte) (rValue *FileReferences, err error) {
	var (
		query = "select id, ref_hash, domain, media_id, access_hash, object_id, origin_domain, origin_id, expire_at, revoked_at, created_at from file_references where ref_hash = ? limit 1"
		do    = &FileReferences{}
	)
	err = m.tx.QueryRowPartial(do, query, refHash)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "file_references",
				Key:      fmt.Sprintf("ref_hash=%v", refHash),
				Cause:    err,
			}
		}
		err = fmt.Errorf("file_references.SelectByRefHash: %w", err)
		return
	}
	rValue = do

	return
}
