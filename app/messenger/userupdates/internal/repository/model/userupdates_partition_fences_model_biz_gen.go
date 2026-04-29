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

type (
	bizUserupdatesPartitionFencesModel interface {
		InsertIgnore(ctx context.Context, data *UserupdatesPartitionFences) (lastInsertId, rowsAffected int64, err error)
		InsertIgnoreTx(tx *sqlx.Tx, data *UserupdatesPartitionFences) (lastInsertId, rowsAffected int64, err error)

		SelectByPartitionId(ctx context.Context, partitionId int32) (*UserupdatesPartitionFences, error)

		CasAcquireOwner(ctx context.Context, ownerInstanceId string, leaseId string, leaseExpiresAt string, partitionId int32, prevOwnerEpoch int64) (rowsAffected int64, err error)
		CasAcquireOwnerTx(tx *sqlx.Tx, ownerInstanceId string, leaseId string, leaseExpiresAt string, partitionId int32, prevOwnerEpoch int64) (rowsAffected int64, err error)
	}
)

// InsertIgnore
// insert ignore into userupdates_partition_fences(partition_id, owner_epoch, owner_instance_id, lease_id, lease_expires_at, created_at, updated_at) values (:partition_id, :owner_epoch, :owner_instance_id, :lease_id, :lease_expires_at, NOW(6), NOW(6))
func (m *defaultUserupdatesPartitionFencesModel) InsertIgnore(ctx context.Context, data *UserupdatesPartitionFences) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into userupdates_partition_fences(partition_id, owner_epoch, owner_instance_id, lease_id, lease_expires_at, created_at, updated_at) values (:partition_id, :owner_epoch, :owner_instance_id, :lease_id, :lease_expires_at, NOW(6), NOW(6))"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("userupdates_partition_fences.InsertIgnore named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("userupdates_partition_fences.InsertIgnore last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("userupdates_partition_fences.InsertIgnore rows affected: %w", err)
	}

	return

}

// InsertIgnoreTx
// insert ignore into userupdates_partition_fences(partition_id, owner_epoch, owner_instance_id, lease_id, lease_expires_at, created_at, updated_at) values (:partition_id, :owner_epoch, :owner_instance_id, :lease_id, :lease_expires_at, NOW(6), NOW(6))
func (m *defaultUserupdatesPartitionFencesModel) InsertIgnoreTx(tx *sqlx.Tx, data *UserupdatesPartitionFences) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into userupdates_partition_fences(partition_id, owner_epoch, owner_instance_id, lease_id, lease_expires_at, created_at, updated_at) values (:partition_id, :owner_epoch, :owner_instance_id, :lease_id, :lease_expires_at, NOW(6), NOW(6))"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("userupdates_partition_fences.InsertIgnoreTx named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("userupdates_partition_fences.InsertIgnoreTx last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("userupdates_partition_fences.InsertIgnoreTx rows affected: %w", err)
	}

	return
}

// SelectByPartitionId
// select partition_id, owner_epoch, owner_instance_id, lease_id, lease_expires_at, created_at, updated_at from userupdates_partition_fences where partition_id = :partition_id limit 1
func (m *defaultUserupdatesPartitionFencesModel) SelectByPartitionId(ctx context.Context, partitionId int32) (rValue *UserupdatesPartitionFences, err error) {

	var (
		query = "select partition_id, owner_epoch, owner_instance_id, lease_id, lease_expires_at, created_at, updated_at from userupdates_partition_fences where partition_id = ? limit 1"
		do    = &UserupdatesPartitionFences{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, partitionId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "userupdates_partition_fences",
				Key:      fmt.Sprintf("partition_id=%v", partitionId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("userupdates_partition_fences.SelectByPartitionId: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// CasAcquireOwner
// update userupdates_partition_fences set owner_epoch = owner_epoch + 1, owner_instance_id = :owner_instance_id, lease_id = :lease_id, lease_expires_at = :lease_expires_at, updated_at = NOW(6) where partition_id = :partition_id and owner_epoch = :prevOwnerEpoch
func (m *defaultUserupdatesPartitionFencesModel) CasAcquireOwner(ctx context.Context, ownerInstanceId string, leaseId string, leaseExpiresAt string, partitionId int32, prevOwnerEpoch int64) (rowsAffected int64, err error) {

	var (
		query   = "update userupdates_partition_fences set owner_epoch = owner_epoch + 1, owner_instance_id = ?, lease_id = ?, lease_expires_at = ?, updated_at = NOW(6) where partition_id = ? and owner_epoch = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, ownerInstanceId, leaseId, leaseExpiresAt, partitionId, prevOwnerEpoch)

	if err != nil {
		err = fmt.Errorf("userupdates_partition_fences.CasAcquireOwner exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("userupdates_partition_fences.CasAcquireOwner rows affected: %w", err)
		return
	}

	return
}

// CasAcquireOwnerTx
// update userupdates_partition_fences set owner_epoch = owner_epoch + 1, owner_instance_id = :owner_instance_id, lease_id = :lease_id, lease_expires_at = :lease_expires_at, updated_at = NOW(6) where partition_id = :partition_id and owner_epoch = :prevOwnerEpoch
func (m *defaultUserupdatesPartitionFencesModel) CasAcquireOwnerTx(tx *sqlx.Tx, ownerInstanceId string, leaseId string, leaseExpiresAt string, partitionId int32, prevOwnerEpoch int64) (rowsAffected int64, err error) {
	var (
		query   = "update userupdates_partition_fences set owner_epoch = owner_epoch + 1, owner_instance_id = ?, lease_id = ?, lease_expires_at = ?, updated_at = NOW(6) where partition_id = ? and owner_epoch = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, ownerInstanceId, leaseId, leaseExpiresAt, partitionId, prevOwnerEpoch)

	if err != nil {
		err = fmt.Errorf("userupdates_partition_fences.CasAcquireOwnerTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("userupdates_partition_fences.CasAcquireOwnerTx rows affected: %w", err)
		return
	}

	return
}
