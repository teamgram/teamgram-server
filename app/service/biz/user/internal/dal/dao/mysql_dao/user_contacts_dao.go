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
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type UserContactsDAO struct {
	db *sqlx.DB
}

func NewUserContactsDAO(db *sqlx.DB) *UserContactsDAO {
	return &UserContactsDAO{db}
}

// InsertOrUpdate
// insert into user_contacts(owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, date2) values (:owner_user_id, :contact_user_id, :contact_phone, :contact_first_name, :contact_last_name, :mutual, :date2) on duplicate key update contact_phone = values(contact_phone), contact_first_name = values(contact_first_name), contact_last_name = values(contact_last_name), mutual = values(mutual), date2 = values(date2), is_deleted = 0
// TODO(@benqi): sqlmap
func (dao *UserContactsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.UserContactsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_contacts(owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, date2) values (:owner_user_id, :contact_user_id, :contact_phone, :contact_first_name, :contact_last_name, :mutual, :date2) on duplicate key update contact_phone = values(contact_phone), contact_first_name = values(contact_first_name), contact_last_name = values(contact_last_name), mutual = values(mutual), date2 = values(date2), is_deleted = 0"
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
// insert into user_contacts(owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, date2) values (:owner_user_id, :contact_user_id, :contact_phone, :contact_first_name, :contact_last_name, :mutual, :date2) on duplicate key update contact_phone = values(contact_phone), contact_first_name = values(contact_first_name), contact_last_name = values(contact_last_name), mutual = values(mutual), date2 = values(date2), is_deleted = 0
// TODO(@benqi): sqlmap
func (dao *UserContactsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.UserContactsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_contacts(owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, date2) values (:owner_user_id, :contact_user_id, :contact_phone, :contact_first_name, :contact_last_name, :mutual, :date2) on duplicate key update contact_phone = values(contact_phone), contact_first_name = values(contact_first_name), contact_last_name = values(contact_last_name), mutual = values(mutual), date2 = values(date2), is_deleted = 0"
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

// SelectContact
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, is_deleted from user_contacts where is_deleted = 0 and owner_user_id = :owner_user_id and contact_user_id = :contact_user_id
// TODO(@benqi): sqlmap
func (dao *UserContactsDAO) SelectContact(ctx context.Context, owner_user_id int64, contact_user_id int64) (rValue *dataobject.UserContactsDO, err error) {
	var (
		query = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, is_deleted from user_contacts where is_deleted = 0 and owner_user_id = ? and contact_user_id = ?"
		do    = &dataobject.UserContactsDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, owner_user_id, contact_user_id)

	if err != nil {
		if err != sqlx.ErrNotFound {
			logx.WithContext(ctx).Errorf("queryx in SelectContact(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectByContactId
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, is_deleted from user_contacts where owner_user_id = :owner_user_id and contact_user_id = :contact_user_id
// TODO(@benqi): sqlmap
func (dao *UserContactsDAO) SelectByContactId(ctx context.Context, owner_user_id int64, contact_user_id int64) (rValue *dataobject.UserContactsDO, err error) {
	var (
		query = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, is_deleted from user_contacts where owner_user_id = ? and contact_user_id = ?"
		do    = &dataobject.UserContactsDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, owner_user_id, contact_user_id)

	if err != nil {
		if err != sqlx.ErrNotFound {
			logx.WithContext(ctx).Errorf("queryx in SelectByContactId(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// SelectListByPhoneList
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, is_deleted from user_contacts where is_deleted = 0 and owner_user_id = :owner_user_id and contact_phone in (:phoneList)
// TODO(@benqi): sqlmap
func (dao *UserContactsDAO) SelectListByPhoneList(ctx context.Context, owner_user_id int64, phoneList []string) (rList []dataobject.UserContactsDO, err error) {
	var (
		query  = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, is_deleted from user_contacts where is_deleted = 0 and owner_user_id = ? and contact_phone in (?)"
		a      []interface{}
		values []dataobject.UserContactsDO
	)

	if len(phoneList) == 0 {
		rList = []dataobject.UserContactsDO{}
		return
	}

	query, a, err = sqlx.In(query, owner_user_id, phoneList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectListByPhoneList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByPhoneList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByPhoneListWithCB
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, is_deleted from user_contacts where is_deleted = 0 and owner_user_id = :owner_user_id and contact_phone in (:phoneList)
// TODO(@benqi): sqlmap
func (dao *UserContactsDAO) SelectListByPhoneListWithCB(ctx context.Context, owner_user_id int64, phoneList []string, cb func(i int, v *dataobject.UserContactsDO)) (rList []dataobject.UserContactsDO, err error) {
	var (
		query  = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, is_deleted from user_contacts where is_deleted = 0 and owner_user_id = ? and contact_phone in (?)"
		a      []interface{}
		values []dataobject.UserContactsDO
	)

	if len(phoneList) == 0 {
		rList = []dataobject.UserContactsDO{}
		return
	}

	query, a, err = sqlx.In(query, owner_user_id, phoneList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectListByPhoneList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByPhoneList(_), error: %v", err)
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

// SelectAllUserContacts
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, is_deleted from user_contacts where owner_user_id = :owner_user_id
// TODO(@benqi): sqlmap
func (dao *UserContactsDAO) SelectAllUserContacts(ctx context.Context, owner_user_id int64) (rList []dataobject.UserContactsDO, err error) {
	var (
		query  = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, is_deleted from user_contacts where owner_user_id = ?"
		values []dataobject.UserContactsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, owner_user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectAllUserContacts(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectAllUserContactsWithCB
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, is_deleted from user_contacts where owner_user_id = :owner_user_id
// TODO(@benqi): sqlmap
func (dao *UserContactsDAO) SelectAllUserContactsWithCB(ctx context.Context, owner_user_id int64, cb func(i int, v *dataobject.UserContactsDO)) (rList []dataobject.UserContactsDO, err error) {
	var (
		query  = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, is_deleted from user_contacts where owner_user_id = ?"
		values []dataobject.UserContactsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, owner_user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectAllUserContacts(_), error: %v", err)
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

// SelectUserContacts
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, is_deleted from user_contacts where owner_user_id = :owner_user_id and is_deleted = 0 order by contact_user_id asc
// TODO(@benqi): sqlmap
func (dao *UserContactsDAO) SelectUserContacts(ctx context.Context, owner_user_id int64) (rList []dataobject.UserContactsDO, err error) {
	var (
		query  = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, is_deleted from user_contacts where owner_user_id = ? and is_deleted = 0 order by contact_user_id asc"
		values []dataobject.UserContactsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, owner_user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectUserContacts(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectUserContactsWithCB
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, is_deleted from user_contacts where owner_user_id = :owner_user_id and is_deleted = 0 order by contact_user_id asc
// TODO(@benqi): sqlmap
func (dao *UserContactsDAO) SelectUserContactsWithCB(ctx context.Context, owner_user_id int64, cb func(i int, v *dataobject.UserContactsDO)) (rList []dataobject.UserContactsDO, err error) {
	var (
		query  = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, is_deleted from user_contacts where owner_user_id = ? and is_deleted = 0 order by contact_user_id asc"
		values []dataobject.UserContactsDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query, owner_user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectUserContacts(_), error: %v", err)
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

// SelectUserContactIdList
// select contact_user_id from user_contacts where owner_user_id = :owner_user_id and is_deleted = 0 order by contact_user_id asc
// TODO(@benqi): sqlmap
func (dao *UserContactsDAO) SelectUserContactIdList(ctx context.Context, owner_user_id int64) (rList []int64, err error) {
	var query = "select contact_user_id from user_contacts where owner_user_id = ? and is_deleted = 0 order by contact_user_id asc"
	err = dao.db.QueryRowsPartial(ctx, &rList, query, owner_user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectUserContactIdList(_), error: %v", err)
	}

	return
}

// SelectUserContactIdListWithCB
// select contact_user_id from user_contacts where owner_user_id = :owner_user_id and is_deleted = 0 order by contact_user_id asc
// TODO(@benqi): sqlmap
func (dao *UserContactsDAO) SelectUserContactIdListWithCB(ctx context.Context, owner_user_id int64, cb func(i int, v int64)) (rList []int64, err error) {
	var query = "select contact_user_id from user_contacts where owner_user_id = ? and is_deleted = 0 order by contact_user_id asc"
	err = dao.db.QueryRowsPartial(ctx, &rList, query, owner_user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectUserContactIdList(_), error: %v", err)
	}

	if cb != nil {
		for i := 0; i < len(rList); i++ {
			cb(i, rList[i])
		}
	}

	return
}

// SelectListByIdList
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, is_deleted from user_contacts where owner_user_id = :owner_user_id and contact_user_id in (:id_list) and is_deleted = 0
// TODO(@benqi): sqlmap
func (dao *UserContactsDAO) SelectListByIdList(ctx context.Context, owner_user_id int64, id_list []int64) (rList []dataobject.UserContactsDO, err error) {
	var (
		query  = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, is_deleted from user_contacts where owner_user_id = ? and contact_user_id in (?) and is_deleted = 0"
		a      []interface{}
		values []dataobject.UserContactsDO
	)

	if len(id_list) == 0 {
		rList = []dataobject.UserContactsDO{}
		return
	}

	query, a, err = sqlx.In(query, owner_user_id, id_list)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectListByIdList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByIdListWithCB
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, is_deleted from user_contacts where owner_user_id = :owner_user_id and contact_user_id in (:id_list) and is_deleted = 0
// TODO(@benqi): sqlmap
func (dao *UserContactsDAO) SelectListByIdListWithCB(ctx context.Context, owner_user_id int64, id_list []int64, cb func(i int, v *dataobject.UserContactsDO)) (rList []dataobject.UserContactsDO, err error) {
	var (
		query  = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, is_deleted from user_contacts where owner_user_id = ? and contact_user_id in (?) and is_deleted = 0"
		a      []interface{}
		values []dataobject.UserContactsDO
	)

	if len(id_list) == 0 {
		rList = []dataobject.UserContactsDO{}
		return
	}

	query, a, err = sqlx.In(query, owner_user_id, id_list)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectListByIdList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByIdList(_), error: %v", err)
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

// SelectListByOwnerListAndContactList
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, is_deleted from user_contacts where owner_user_id in (:idList1) and contact_user_id in (:idList2) and is_deleted = 0
// TODO(@benqi): sqlmap
func (dao *UserContactsDAO) SelectListByOwnerListAndContactList(ctx context.Context, idList1 []int64, idList2 []int64) (rList []dataobject.UserContactsDO, err error) {
	var (
		query  = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, is_deleted from user_contacts where owner_user_id in (?) and contact_user_id in (?) and is_deleted = 0"
		a      []interface{}
		values []dataobject.UserContactsDO
	)
	if len(idList1) == 0 {
		rList = []dataobject.UserContactsDO{}
		return
	}
	if len(idList2) == 0 {
		rList = []dataobject.UserContactsDO{}
		return
	}

	query, a, err = sqlx.In(query, idList1, idList2)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectListByOwnerListAndContactList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByOwnerListAndContactList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByOwnerListAndContactListWithCB
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, is_deleted from user_contacts where owner_user_id in (:idList1) and contact_user_id in (:idList2) and is_deleted = 0
// TODO(@benqi): sqlmap
func (dao *UserContactsDAO) SelectListByOwnerListAndContactListWithCB(ctx context.Context, idList1 []int64, idList2 []int64, cb func(i int, v *dataobject.UserContactsDO)) (rList []dataobject.UserContactsDO, err error) {
	var (
		query  = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, is_deleted from user_contacts where owner_user_id in (?) and contact_user_id in (?) and is_deleted = 0"
		a      []interface{}
		values []dataobject.UserContactsDO
	)
	if len(idList1) == 0 {
		rList = []dataobject.UserContactsDO{}
		return
	}
	if len(idList2) == 0 {
		rList = []dataobject.UserContactsDO{}
		return
	}

	query, a, err = sqlx.In(query, idList1, idList2)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectListByOwnerListAndContactList(_), error: %v", err)
		return
	}
	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByOwnerListAndContactList(_), error: %v", err)
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

// UpdateContactNameById
// update user_contacts set contact_first_name = :contact_first_name, contact_last_name = :contact_last_name, is_deleted = 0 where id = :id
// TODO(@benqi): sqlmap
func (dao *UserContactsDAO) UpdateContactNameById(ctx context.Context, contact_first_name string, contact_last_name string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set contact_first_name = ?, contact_last_name = ?, is_deleted = 0 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, contact_first_name, contact_last_name, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateContactNameById(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateContactNameById(_), error: %v", err)
	}

	return
}

// update user_contacts set contact_first_name = :contact_first_name, contact_last_name = :contact_last_name, is_deleted = 0 where id = :id
// UpdateContactNameByIdTx
// TODO(@benqi): sqlmap
func (dao *UserContactsDAO) UpdateContactNameByIdTx(tx *sqlx.Tx, contact_first_name string, contact_last_name string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set contact_first_name = ?, contact_last_name = ?, is_deleted = 0 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, contact_first_name, contact_last_name, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateContactNameById(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateContactNameById(_), error: %v", err)
	}

	return
}

// UpdateContactName
// update user_contacts set contact_first_name = :contact_first_name, contact_last_name = :contact_last_name, is_deleted = 0 where contact_user_id != 0 and (owner_user_id = :owner_user_id and contact_user_id = :contact_user_id)
// TODO(@benqi): sqlmap
func (dao *UserContactsDAO) UpdateContactName(ctx context.Context, contact_first_name string, contact_last_name string, owner_user_id int64, contact_user_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set contact_first_name = ?, contact_last_name = ?, is_deleted = 0 where contact_user_id != 0 and (owner_user_id = ? and contact_user_id = ?)"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, contact_first_name, contact_last_name, owner_user_id, contact_user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateContactName(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateContactName(_), error: %v", err)
	}

	return
}

// update user_contacts set contact_first_name = :contact_first_name, contact_last_name = :contact_last_name, is_deleted = 0 where contact_user_id != 0 and (owner_user_id = :owner_user_id and contact_user_id = :contact_user_id)
// UpdateContactNameTx
// TODO(@benqi): sqlmap
func (dao *UserContactsDAO) UpdateContactNameTx(tx *sqlx.Tx, contact_first_name string, contact_last_name string, owner_user_id int64, contact_user_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set contact_first_name = ?, contact_last_name = ?, is_deleted = 0 where contact_user_id != 0 and (owner_user_id = ? and contact_user_id = ?)"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, contact_first_name, contact_last_name, owner_user_id, contact_user_id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateContactName(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateContactName(_), error: %v", err)
	}

	return
}

// UpdateMutual
// update user_contacts set mutual = :mutual where contact_user_id != 0 and (owner_user_id = :owner_user_id and contact_user_id = :contact_user_id)
// TODO(@benqi): sqlmap
func (dao *UserContactsDAO) UpdateMutual(ctx context.Context, mutual bool, owner_user_id int64, contact_user_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set mutual = ? where contact_user_id != 0 and (owner_user_id = ? and contact_user_id = ?)"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, mutual, owner_user_id, contact_user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateMutual(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateMutual(_), error: %v", err)
	}

	return
}

// update user_contacts set mutual = :mutual where contact_user_id != 0 and (owner_user_id = :owner_user_id and contact_user_id = :contact_user_id)
// UpdateMutualTx
// TODO(@benqi): sqlmap
func (dao *UserContactsDAO) UpdateMutualTx(tx *sqlx.Tx, mutual bool, owner_user_id int64, contact_user_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set mutual = ? where contact_user_id != 0 and (owner_user_id = ? and contact_user_id = ?)"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, mutual, owner_user_id, contact_user_id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateMutual(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateMutual(_), error: %v", err)
	}

	return
}

// DeleteContacts
// update user_contacts set is_deleted = 1, mutual = 0 where contact_user_id != 0 and (owner_user_id = :owner_user_id and contact_user_id in (:id_list))
// TODO(@benqi): sqlmap
func (dao *UserContactsDAO) DeleteContacts(ctx context.Context, owner_user_id int64, id_list []int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set is_deleted = 1, mutual = 0 where contact_user_id != 0 and (owner_user_id = ? and contact_user_id in (?))"
		a       []interface{}
		rResult sql.Result
	)

	if len(id_list) == 0 {
		return
	}

	query, a, err = sqlx.In(query, owner_user_id, id_list)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in DeleteContacts(_), error: %v", err)
		return
	}
	rResult, err = dao.db.Exec(ctx, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in DeleteContacts(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in DeleteContacts(_), error: %v", err)
	}

	return
}

// update user_contacts set is_deleted = 1, mutual = 0 where contact_user_id != 0 and (owner_user_id = :owner_user_id and contact_user_id in (:id_list))
// DeleteContactsTx
// TODO(@benqi): sqlmap
func (dao *UserContactsDAO) DeleteContactsTx(tx *sqlx.Tx, owner_user_id int64, id_list []int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set is_deleted = 1, mutual = 0 where contact_user_id != 0 and (owner_user_id = ? and contact_user_id in (?))"
		a       []interface{}
		rResult sql.Result
	)

	if len(id_list) == 0 {
		return
	}

	query, a, err = sqlx.In(query, owner_user_id, id_list)
	if err != nil {
		// r sql.Result
		logx.WithContext(tx.Context()).Errorf("sqlx.In in DeleteContacts(_), error: %v", err)
		return
	}
	rResult, err = tx.Exec(query, a...)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in DeleteContacts(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in DeleteContacts(_), error: %v", err)
	}

	return
}

// UpdatePhoneByContactId
// update user_contacts set contact_phone = :contact_phone where contact_user_id = :contact_user_id
// TODO(@benqi): sqlmap
func (dao *UserContactsDAO) UpdatePhoneByContactId(ctx context.Context, contact_phone string, contact_user_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set contact_phone = ? where contact_user_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, contact_phone, contact_user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdatePhoneByContactId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdatePhoneByContactId(_), error: %v", err)
	}

	return
}

// update user_contacts set contact_phone = :contact_phone where contact_user_id = :contact_user_id
// UpdatePhoneByContactIdTx
// TODO(@benqi): sqlmap
func (dao *UserContactsDAO) UpdatePhoneByContactIdTx(tx *sqlx.Tx, contact_phone string, contact_user_id int64) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set contact_phone = ? where contact_user_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, contact_phone, contact_user_id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdatePhoneByContactId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdatePhoneByContactId(_), error: %v", err)
	}

	return
}
