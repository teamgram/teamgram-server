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

type bizUserupdatesPartitionFencesModel interface {
	InsertIgnore(ctx context.Context, data *UserupdatesPartitionFences) (lastInsertId, rowsAffected int64, err error)
	SelectByPartitionId(ctx context.Context, partitionId int32) (*UserupdatesPartitionFences, error)
	CasAcquireOwner(ctx context.Context, ownerInstanceId string, leaseId string, partitionId int32, prevOwnerEpoch int64) (rowsAffected int64, err error)
}

type UserupdatesPartitionFencesTxModel interface {
	InsertIgnore(data *UserupdatesPartitionFences) (lastInsertId, rowsAffected int64, err error)
	SelectByPartitionId(partitionId int32) (*UserupdatesPartitionFences, error)
	CasAcquireOwner(ownerInstanceId string, leaseId string, partitionId int32, prevOwnerEpoch int64) (rowsAffected int64, err error)
}

type defaultUserupdatesPartitionFencesTxModel struct {
	tx *sqlx.Tx
}

func NewUserupdatesPartitionFencesTxModel(tx *sqlx.Tx) UserupdatesPartitionFencesTxModel {
	return &defaultUserupdatesPartitionFencesTxModel{tx: tx}
}

// InsertIgnore
// insert ignore into userupdates_partition_fences(partition_id, owner_epoch, owner_instance_id, lease_id) values (:partition_id, :owner_epoch, :owner_instance_id, :lease_id)
func (m *defaultUserupdatesPartitionFencesModel) InsertIgnore(ctx context.Context, data *UserupdatesPartitionFences) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into userupdates_partition_fences(partition_id, owner_epoch, owner_instance_id, lease_id) values (:partition_id, :owner_epoch, :owner_instance_id, :lease_id)"
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

// InsertIgnore
// insert ignore into userupdates_partition_fences(partition_id, owner_epoch, owner_instance_id, lease_id) values (:partition_id, :owner_epoch, :owner_instance_id, :lease_id)
func (m *defaultUserupdatesPartitionFencesTxModel) InsertIgnore(data *UserupdatesPartitionFences) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into userupdates_partition_fences(partition_id, owner_epoch, owner_instance_id, lease_id) values (:partition_id, :owner_epoch, :owner_instance_id, :lease_id)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
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

// SelectByPartitionId
// select partition_id, owner_epoch, owner_instance_id from userupdates_partition_fences where partition_id = :partition_id limit 1
func (m *defaultUserupdatesPartitionFencesModel) SelectByPartitionId(ctx context.Context, partitionId int32) (rValue *UserupdatesPartitionFences, err error) {

	var (
		query = "select partition_id, owner_epoch, owner_instance_id from userupdates_partition_fences where partition_id = ? limit 1"
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

// SelectByPartitionId
// select partition_id, owner_epoch, owner_instance_id from userupdates_partition_fences where partition_id = :partition_id limit 1
func (m *defaultUserupdatesPartitionFencesTxModel) SelectByPartitionId(partitionId int32) (rValue *UserupdatesPartitionFences, err error) {
	var (
		query = "select partition_id, owner_epoch, owner_instance_id from userupdates_partition_fences where partition_id = ? limit 1"
		do    = &UserupdatesPartitionFences{}
	)
	err = m.tx.QueryRowPartial(do, query, partitionId)

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
	}
	rValue = do

	return
}

// CasAcquireOwner
// update userupdates_partition_fences set owner_epoch = owner_epoch + 1, owner_instance_id = :owner_instance_id, lease_id = :lease_id where partition_id = :partition_id and owner_epoch = :prevOwnerEpoch
func (m *defaultUserupdatesPartitionFencesModel) CasAcquireOwner(ctx context.Context, ownerInstanceId string, leaseId string, partitionId int32, prevOwnerEpoch int64) (rowsAffected int64, err error) {

	var (
		query   = "update userupdates_partition_fences set owner_epoch = owner_epoch + 1, owner_instance_id = ?, lease_id = ? where partition_id = ? and owner_epoch = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, ownerInstanceId, leaseId, partitionId, prevOwnerEpoch)

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

// CasAcquireOwner
// update userupdates_partition_fences set owner_epoch = owner_epoch + 1, owner_instance_id = :owner_instance_id, lease_id = :lease_id where partition_id = :partition_id and owner_epoch = :prevOwnerEpoch
func (m *defaultUserupdatesPartitionFencesTxModel) CasAcquireOwner(ownerInstanceId string, leaseId string, partitionId int32, prevOwnerEpoch int64) (rowsAffected int64, err error) {
	var (
		query   = "update userupdates_partition_fences set owner_epoch = owner_epoch + 1, owner_instance_id = ?, lease_id = ? where partition_id = ? and owner_epoch = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, ownerInstanceId, leaseId, partitionId, prevOwnerEpoch)

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
