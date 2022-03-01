/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mysql_dao

import (
	"context"
	"database/sql"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/updates/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type ChannelPtsUpdatesDAO struct {
	db *sqlx.DB
}

func NewChannelPtsUpdatesDAO(db *sqlx.DB) *ChannelPtsUpdatesDAO {
	return &ChannelPtsUpdatesDAO{db}
}

// Insert
// insert into channel_pts_updates(channel_id, pts, pts_count, update_type, new_message_id, update_data, date2) values (:channel_id, :pts, :pts_count, :update_type, :new_message_id, :update_data, :date2)
func (dao *ChannelPtsUpdatesDAO) Insert(ctx context.Context, do *dataobject.ChannelPtsUpdatesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into channel_pts_updates(channel_id, pts, pts_count, update_type, new_message_id, update_data, date2) values (:channel_id, :pts, :pts_count, :update_type, :new_message_id, :update_data, :date2)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in Insert(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Insert(%v)_error: %v", do, err)
	}

	return
}

// InsertTx
// insert into channel_pts_updates(channel_id, pts, pts_count, update_type, new_message_id, update_data, date2) values (:channel_id, :pts, :pts_count, :update_type, :new_message_id, :update_data, :date2)
func (dao *ChannelPtsUpdatesDAO) InsertTx(tx *sqlx.Tx, do *dataobject.ChannelPtsUpdatesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into channel_pts_updates(channel_id, pts, pts_count, update_type, new_message_id, update_data, date2) values (:channel_id, :pts, :pts_count, :update_type, :new_message_id, :update_data, :date2)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in Insert(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Insert(%v)_error: %v", do, err)
	}

	return
}

// SelectLastPts
// select pts from channel_pts_updates where channel_id = :channel_id order by pts desc limit 1
func (dao *ChannelPtsUpdatesDAO) SelectLastPts(ctx context.Context, channel_id int64) (rValue *dataobject.ChannelPtsUpdatesDO, err error) {
	var (
		query = "select pts from channel_pts_updates where channel_id = ? order by pts desc limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, channel_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectLastPts(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.ChannelPtsUpdatesDO{}
	if rows.Next() {
		// TODO(@benqi): not use reflect
		err = rows.StructScan(do)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectLastPts(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

// SelectByGtPts
// select channel_id, pts, pts_count, update_type, new_message_id, update_data, date2 from channel_pts_updates where channel_id = :channel_id and pts > :pts order by pts asc limit :limit
func (dao *ChannelPtsUpdatesDAO) SelectByGtPts(ctx context.Context, channel_id int64, pts int32, limit int32) (rList []dataobject.ChannelPtsUpdatesDO, err error) {
	var (
		query = "select channel_id, pts, pts_count, update_type, new_message_id, update_data, date2 from channel_pts_updates where channel_id = ? and pts > ? order by pts asc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, channel_id, pts, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByGtPts(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChannelPtsUpdatesDO
	for rows.Next() {
		v := dataobject.ChannelPtsUpdatesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectByGtPts(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectByGtPtsWithCB
// select channel_id, pts, pts_count, update_type, new_message_id, update_data, date2 from channel_pts_updates where channel_id = :channel_id and pts > :pts order by pts asc limit :limit
func (dao *ChannelPtsUpdatesDAO) SelectByGtPtsWithCB(ctx context.Context, channel_id int64, pts int32, limit int32, cb func(i int, v *dataobject.ChannelPtsUpdatesDO)) (rList []dataobject.ChannelPtsUpdatesDO, err error) {
	var (
		query = "select channel_id, pts, pts_count, update_type, new_message_id, update_data, date2 from channel_pts_updates where channel_id = ? and pts > ? order by pts asc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, channel_id, pts, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByGtPts(_), error: %v", err)
		return
	}

	defer func() {
		rows.Close()
		if err == nil && cb != nil {
			for i := 0; i < len(rList); i++ {
				cb(i, &rList[i])
			}
		}
	}()

	var values []dataobject.ChannelPtsUpdatesDO
	for rows.Next() {
		v := dataobject.ChannelPtsUpdatesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectByGtPts(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectByGtDate2
// select channel_id, pts, pts_count, update_type, new_message_id, update_data, date2 from channel_pts_updates where channel_id in (:idList) and date2 > :date2 order by date2 asc
func (dao *ChannelPtsUpdatesDAO) SelectByGtDate2(ctx context.Context, idList []int32, date2 int64) (rList []dataobject.ChannelPtsUpdatesDO, err error) {
	var (
		query = "select channel_id, pts, pts_count, update_type, new_message_id, update_data, date2 from channel_pts_updates where channel_id in (?) and date2 > ? order by date2 asc"
		a     []interface{}
		rows  *sqlx.Rows
	)
	if len(idList) == 0 {
		rList = []dataobject.ChannelPtsUpdatesDO{}
		return
	}

	query, a, err = sqlx.In(query, idList, date2)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectByGtDate2(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByGtDate2(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChannelPtsUpdatesDO
	for rows.Next() {
		v := dataobject.ChannelPtsUpdatesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectByGtDate2(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectByGtDate2WithCB
// select channel_id, pts, pts_count, update_type, new_message_id, update_data, date2 from channel_pts_updates where channel_id in (:idList) and date2 > :date2 order by date2 asc
func (dao *ChannelPtsUpdatesDAO) SelectByGtDate2WithCB(ctx context.Context, idList []int32, date2 int64, cb func(i int, v *dataobject.ChannelPtsUpdatesDO)) (rList []dataobject.ChannelPtsUpdatesDO, err error) {
	var (
		query = "select channel_id, pts, pts_count, update_type, new_message_id, update_data, date2 from channel_pts_updates where channel_id in (?) and date2 > ? order by date2 asc"
		a     []interface{}
		rows  *sqlx.Rows
	)
	if len(idList) == 0 {
		rList = []dataobject.ChannelPtsUpdatesDO{}
		return
	}

	query, a, err = sqlx.In(query, idList, date2)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectByGtDate2(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByGtDate2(_), error: %v", err)
		return
	}

	defer func() {
		rows.Close()
		if err == nil && cb != nil {
			for i := 0; i < len(rList); i++ {
				cb(i, &rList[i])
			}
		}
	}()

	var values []dataobject.ChannelPtsUpdatesDO
	for rows.Next() {
		v := dataobject.ChannelPtsUpdatesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectByGtDate2(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
