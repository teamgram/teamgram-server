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

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userupdatesPartitionFencesFieldNames          = builder.RawFieldNames(&UserupdatesPartitionFences{})
	userupdatesPartitionFencesRows                = strings.Join(userupdatesPartitionFencesFieldNames, ",")
	userupdatesPartitionFencesRowsExpectAutoSet   = strings.Join(stringx.Remove(userupdatesPartitionFencesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userupdatesPartitionFencesRowsWithPlaceHolder = strings.Join(stringx.Remove(userupdatesPartitionFencesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	userupdatesPartitionFencesModel interface {
		Insert2(ctx context.Context, data *UserupdatesPartitionFences) (sql.Result, error)
		FindOne(ctx context.Context, partitionId int32) (*UserupdatesPartitionFences, error)
		FindListByPartitionIdList(ctx context.Context, partitionId ...int32) ([]UserupdatesPartitionFences, error)
		Update2(ctx context.Context, data *UserupdatesPartitionFences) error
		Delete2(ctx context.Context, partitionId int32) error
	}

	defaultUserupdatesPartitionFencesModel struct {
		db *sqlx.DB
	}

	UserupdatesPartitionFences struct {
		PartitionId     int32  `db:"partition_id" json:"partition_id"`
		OwnerEpoch      int64  `db:"owner_epoch" json:"owner_epoch"`
		OwnerInstanceId string `db:"owner_instance_id" json:"owner_instance_id"`
		LeaseId         string `db:"lease_id" json:"lease_id"`
		LeaseExpiresAt  string `db:"lease_expires_at" json:"lease_expires_at"`
	}
)

func newUserupdatesPartitionFencesModel(db *sqlx.DB) *defaultUserupdatesPartitionFencesModel {
	return &defaultUserupdatesPartitionFencesModel{
		db: db,
	}
}

func (m *defaultUserupdatesPartitionFencesModel) Insert2(ctx context.Context, data *UserupdatesPartitionFences) (sql.Result, error) {
	tableName := "userupdates_partition_fences"
	query := fmt.Sprintf("insert into `%s` (%s) values (?, ?, ?, ?)", tableName, userupdatesPartitionFencesRowsExpectAutoSet)

	r, err := m.db.Exec(ctx, query, data.OwnerEpoch, data.OwnerInstanceId, data.LeaseId, data.LeaseExpiresAt)
	if err != nil {
		return nil, fmt.Errorf("userupdates_partition_fences.Insert2 exec: %w", err)
	}

	return r, nil
}

func (m *defaultUserupdatesPartitionFencesModel) Delete2(ctx context.Context, partitionId int32) error {
	tableName := "userupdates_partition_fences"
	query := fmt.Sprintf("delete from `%s` where `partition_id` = ?", tableName)

	_, err := m.db.Exec(ctx, query, partitionId)
	if err != nil {
		return fmt.Errorf("userupdates_partition_fences.Delete2 exec: %w", err)
	}

	return nil
}

func (m *defaultUserupdatesPartitionFencesModel) FindOne(ctx context.Context, partitionId int32) (*UserupdatesPartitionFences, error) {
	tableName := "userupdates_partition_fences"
	query := fmt.Sprintf("select %s from %s where partition_id = ? limit 1", userupdatesPartitionFencesRows, tableName)
	var resp UserupdatesPartitionFences

	err := m.db.QueryRowPartial(ctx, &resp, query, partitionId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "userupdates_partition_fences",
				Key:      fmt.Sprintf("partition_id=%v", partitionId),
				Cause:    err,
			}
		}
		return nil, fmt.Errorf("userupdates_partition_fences.FindOne: %w", err)
	}

	return &resp, nil
}

func (m *defaultUserupdatesPartitionFencesModel) FindListByPartitionIdList(ctx context.Context, partitionId ...int32) ([]UserupdatesPartitionFences, error) {
	if len(partitionId) == 0 {
		return []UserupdatesPartitionFences{}, nil
	}
	tableName := "userupdates_partition_fences"
	query := fmt.Sprintf("select %s from %s where partition_id in (%s)", userupdatesPartitionFencesRows, tableName, sqlx.InInt32List(partitionId))

	var resp []UserupdatesPartitionFences
	err := m.db.QueryRowsPartial(ctx, &resp, query)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []UserupdatesPartitionFences{}, nil
		}
		return nil, fmt.Errorf("userupdates_partition_fences.FindListByPartitionIdList: %w", err)
	}

	return resp, nil
}

func (m *defaultUserupdatesPartitionFencesModel) Update2(ctx context.Context, data *UserupdatesPartitionFences) error {
	tableName := "userupdates_partition_fences"
	query := fmt.Sprintf("update `%s` set %s where `partition_id` = ?", tableName, userupdatesPartitionFencesRowsWithPlaceHolder)

	_, err := m.db.Exec(ctx, query, data.OwnerEpoch, data.OwnerInstanceId, data.LeaseId, data.LeaseExpiresAt, data.PartitionId)
	if err != nil {
		return fmt.Errorf("userupdates_partition_fences.Update2 exec: %w", err)
	}

	return nil
}
