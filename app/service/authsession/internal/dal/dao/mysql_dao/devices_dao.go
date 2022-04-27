/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mysql_dao

import (
	"context"
	"database/sql"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/authsession/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type DevicesDAO struct {
	db *sqlx.DB
}

func NewDevicesDAO(db *sqlx.DB) *DevicesDAO {
	return &DevicesDAO{db}
}

// InsertOrUpdate
// insert into devices(auth_key_id, user_id, token_type, token, app_sandbox, secret, other_uids) values (:auth_key_id, :user_id, :token_type, :token, :app_sandbox, :secret, :other_uids) on duplicate key update token = values(token), secret = values(secret), other_uids = values(other_uids), state = 0
// TODO(@benqi): sqlmap
func (dao *DevicesDAO) InsertOrUpdate(ctx context.Context, do *dataobject.DevicesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into devices(auth_key_id, user_id, token_type, token, app_sandbox, secret, other_uids) values (:auth_key_id, :user_id, :token_type, :token, :app_sandbox, :secret, :other_uids) on duplicate key update token = values(token), secret = values(secret), other_uids = values(other_uids), state = 0"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", do, err)
	}

	return
}

// InsertOrUpdateTx
// insert into devices(auth_key_id, user_id, token_type, token, app_sandbox, secret, other_uids) values (:auth_key_id, :user_id, :token_type, :token, :app_sandbox, :secret, :other_uids) on duplicate key update token = values(token), secret = values(secret), other_uids = values(other_uids), state = 0
// TODO(@benqi): sqlmap
func (dao *DevicesDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.DevicesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into devices(auth_key_id, user_id, token_type, token, app_sandbox, secret, other_uids) values (:auth_key_id, :user_id, :token_type, :token, :app_sandbox, :secret, :other_uids) on duplicate key update token = values(token), secret = values(secret), other_uids = values(other_uids), state = 0"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", do, err)
	}

	return
}

// Select
// select id, auth_key_id, user_id, token_type, token, app_sandbox, secret, other_uids from devices where auth_key_id = :auth_key_id and user_id = :user_id and token_type = :token_type and state = 0
// TODO(@benqi): sqlmap
func (dao *DevicesDAO) Select(ctx context.Context, auth_key_id int64, user_id int64, token_type int32) (rValue *dataobject.DevicesDO, err error) {
	var (
		query = "select id, auth_key_id, user_id, token_type, token, app_sandbox, secret, other_uids from devices where auth_key_id = ? and user_id = ? and token_type = ? and state = 0"
		do    = &dataobject.DevicesDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, auth_key_id, user_id, token_type)

	if err != nil {
		if err != sqlx.ErrNotFound {
			logx.WithContext(ctx).Errorf("queryx in Select(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectListById
// select id, auth_key_id, user_id, token_type, token from devices where token_type = :token_type and token = :token and state = 1
// TODO(@benqi): sqlmap
func (dao *DevicesDAO) SelectListById(ctx context.Context, token_type int32, token string) (rList []dataobject.DevicesDO, err error) {
	var (
		query  = "select id, auth_key_id, user_id, token_type, token from devices where token_type = ? and token = ? and state = 1"
		values []dataobject.DevicesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, token_type, token)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListById(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByIdWithCB
// select id, auth_key_id, user_id, token_type, token from devices where token_type = :token_type and token = :token and state = 1
// TODO(@benqi): sqlmap
func (dao *DevicesDAO) SelectListByIdWithCB(ctx context.Context, token_type int32, token string, cb func(i int, v *dataobject.DevicesDO)) (rList []dataobject.DevicesDO, err error) {
	var (
		query  = "select id, auth_key_id, user_id, token_type, token from devices where token_type = ? and token = ? and state = 1"
		values []dataobject.DevicesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, token_type, token)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListById(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		for i := 0; i < len(rList); i++ {
			cb(i, &rList[i])
		}
	}

	return
}

// UpdateState
// update devices set state = :state where auth_key_id = :auth_key_id and user_id = :user_id and token_type
// TODO(@benqi): sqlmap
func (dao *DevicesDAO) UpdateState(ctx context.Context, state bool, auth_key_id int64, user_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update devices set state = ? where auth_key_id = ? and user_id = ? and token_type"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, state, auth_key_id, user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateState(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateState(_), error: %v", err)
	}

	return
}

// update devices set state = :state where auth_key_id = :auth_key_id and user_id = :user_id and token_type
// UpdateStateTx
// TODO(@benqi): sqlmap
func (dao *DevicesDAO) UpdateStateTx(tx *sqlx.Tx, state bool, auth_key_id int64, user_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update devices set state = ? where auth_key_id = ? and user_id = ? and token_type"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, state, auth_key_id, user_id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateState(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateState(_), error: %v", err)
	}

	return
}

// UpdateStateById
// update devices set state = :state where id = :id
// TODO(@benqi): sqlmap
func (dao *DevicesDAO) UpdateStateById(ctx context.Context, state bool, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update devices set state = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, state, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateStateById(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateStateById(_), error: %v", err)
	}

	return
}

// update devices set state = :state where id = :id
// UpdateStateByIdTx
// TODO(@benqi): sqlmap
func (dao *DevicesDAO) UpdateStateByIdTx(tx *sqlx.Tx, state bool, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update devices set state = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, state, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateStateById(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateStateById(_), error: %v", err)
	}

	return
}

// UpdateStateByToken
// update devices set state = :state where token_type = :token_type and token = :token
// TODO(@benqi): sqlmap
func (dao *DevicesDAO) UpdateStateByToken(ctx context.Context, state bool, token_type int32, token string) (rowsAffected int64, err error) {
	var (
		query   = "update devices set state = ? where token_type = ? and token = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, state, token_type, token)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateStateByToken(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateStateByToken(_), error: %v", err)
	}

	return
}

// update devices set state = :state where token_type = :token_type and token = :token
// UpdateStateByTokenTx
// TODO(@benqi): sqlmap
func (dao *DevicesDAO) UpdateStateByTokenTx(tx *sqlx.Tx, state bool, token_type int32, token string) (rowsAffected int64, err error) {
	var (
		query   = "update devices set state = ? where token_type = ? and token = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, state, token_type, token)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateStateByToken(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateStateByToken(_), error: %v", err)
	}

	return
}
