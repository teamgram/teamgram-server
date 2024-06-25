/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mysql_dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/message/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result
var _ = fmt.Sprintf
var _ = strings.Join
var _ = errors.Is

type ChatsDAO struct {
	db *sqlx.DB
}

func NewChatsDAO(db *sqlx.DB) *ChatsDAO {
	return &ChatsDAO{
		db: db,
	}
}

// Insert
// insert into chats(creator_user_id, access_hash, random_id, participant_count, title, about, `date`) values (:creator_user_id, :access_hash, :random_id, :participant_count, :title, :about, :date)
func (dao *ChatsDAO) Insert(ctx context.Context, do *dataobject.ChatsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chats(creator_user_id, access_hash, random_id, participant_count, title, about, `date`) values (:creator_user_id, :access_hash, :random_id, :participant_count, :title, :about, :date)"
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
// insert into chats(creator_user_id, access_hash, random_id, participant_count, title, about, `date`) values (:creator_user_id, :access_hash, :random_id, :participant_count, :title, :about, :date)
func (dao *ChatsDAO) InsertTx(tx *sqlx.Tx, do *dataobject.ChatsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chats(creator_user_id, access_hash, random_id, participant_count, title, about, `date`) values (:creator_user_id, :access_hash, :random_id, :participant_count, :title, :about, :date)"
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

// Select
// select id, creator_user_id, access_hash, participant_count, title, about, photo_id, link, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where id = :id
func (dao *ChatsDAO) Select(ctx context.Context, id int64) (rValue *dataobject.ChatsDO, err error) {
	var (
		query = "select id, creator_user_id, access_hash, participant_count, title, about, photo_id, link, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where id = ?"
		do    = &dataobject.ChatsDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, id)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
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

// SelectPhotoId
// select photo_id from chats where id = :id
func (dao *ChatsDAO) SelectPhotoId(ctx context.Context, id int64) (rValue int64, err error) {
	var query = "select photo_id from chats where id = ?"
	err = dao.db.QueryRowPartial(ctx, &rValue, query, id)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("get in SelectPhotoId(_), error: %v", err)
			return
		} else {
			err = nil
		}
	}

	return
}

// SelectLastCreator
// select id, creator_user_id, access_hash, participant_count, title, about, photo_id, link, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where creator_user_id = :creator_user_id order by `date` desc limit 1
func (dao *ChatsDAO) SelectLastCreator(ctx context.Context, creatorUserId int64) (rValue *dataobject.ChatsDO, err error) {
	var (
		query = "select id, creator_user_id, access_hash, participant_count, title, about, photo_id, link, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where creator_user_id = ? order by `date` desc limit 1"
		do    = &dataobject.ChatsDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query, creatorUserId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectLastCreator(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// UpdateTitle
// update chats set title = :title, version = version + 1 where id = :id
func (dao *ChatsDAO) UpdateTitle(ctx context.Context, title string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set title = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, title, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateTitle(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateTitle(_), error: %v", err)
	}

	return
}

// UpdateTitleTx
// update chats set title = :title, version = version + 1 where id = :id
func (dao *ChatsDAO) UpdateTitleTx(tx *sqlx.Tx, title string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set title = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, title, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateTitle(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateTitle(_), error: %v", err)
	}

	return
}

// UpdateAbout
// update chats set about = :about where id = :id
func (dao *ChatsDAO) UpdateAbout(ctx context.Context, about string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set about = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, about, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateAbout(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateAbout(_), error: %v", err)
	}

	return
}

// UpdateAboutTx
// update chats set about = :about where id = :id
func (dao *ChatsDAO) UpdateAboutTx(tx *sqlx.Tx, about string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set about = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, about, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateAbout(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateAbout(_), error: %v", err)
	}

	return
}

// SelectByIdList
// select id, creator_user_id, access_hash, participant_count, title, about, photo_id, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where id in (:idList)
func (dao *ChatsDAO) SelectByIdList(ctx context.Context, idList []int32) (rList []dataobject.ChatsDO, err error) {
	var (
		query  = fmt.Sprintf("select id, creator_user_id, access_hash, participant_count, title, about, photo_id, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where id in (%s)", sqlx.InInt32List(idList))
		values []dataobject.ChatsDO
	)
	if len(idList) == 0 {
		rList = []dataobject.ChatsDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectByIdListWithCB
// select id, creator_user_id, access_hash, participant_count, title, about, photo_id, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where id in (:idList)
func (dao *ChatsDAO) SelectByIdListWithCB(ctx context.Context, idList []int32, cb func(sz, i int, v *dataobject.ChatsDO)) (rList []dataobject.ChatsDO, err error) {
	var (
		query  = fmt.Sprintf("select id, creator_user_id, access_hash, participant_count, title, about, photo_id, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where id in (%s)", sqlx.InInt32List(idList))
		values []dataobject.ChatsDO
	)
	if len(idList) == 0 {
		rList = []dataobject.ChatsDO{}
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByIdList(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		sz := len(rList)
		for i := 0; i < sz; i++ {
			cb(sz, i, &rList[i])
		}
	}

	return
}

// UpdateParticipantCount
// update chats set participant_count = :participant_count, version = version + 1 where id = :id
func (dao *ChatsDAO) UpdateParticipantCount(ctx context.Context, participantCount int32, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set participant_count = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, participantCount, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateParticipantCount(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateParticipantCount(_), error: %v", err)
	}

	return
}

// UpdateParticipantCountTx
// update chats set participant_count = :participant_count, version = version + 1 where id = :id
func (dao *ChatsDAO) UpdateParticipantCountTx(tx *sqlx.Tx, participantCount int32, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set participant_count = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, participantCount, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateParticipantCount(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateParticipantCount(_), error: %v", err)
	}

	return
}

// UpdatePhotoId
// update chats set photo_id = :photo_id, version = version + 1 where id = :id
func (dao *ChatsDAO) UpdatePhotoId(ctx context.Context, photoId int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set photo_id = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, photoId, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdatePhotoId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdatePhotoId(_), error: %v", err)
	}

	return
}

// UpdatePhotoIdTx
// update chats set photo_id = :photo_id, version = version + 1 where id = :id
func (dao *ChatsDAO) UpdatePhotoIdTx(tx *sqlx.Tx, photoId int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set photo_id = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, photoId, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdatePhotoId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdatePhotoId(_), error: %v", err)
	}

	return
}

// UpdateAdminsEnabled
// update chats set admins_enabled = :admins_enabled, version = version + 1 where id = :id
func (dao *ChatsDAO) UpdateAdminsEnabled(ctx context.Context, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set admins_enabled = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateAdminsEnabled(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateAdminsEnabled(_), error: %v", err)
	}

	return
}

// UpdateAdminsEnabledTx
// update chats set admins_enabled = :admins_enabled, version = version + 1 where id = :id
func (dao *ChatsDAO) UpdateAdminsEnabledTx(tx *sqlx.Tx, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set admins_enabled = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateAdminsEnabled(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateAdminsEnabled(_), error: %v", err)
	}

	return
}

// UpdateDefaultBannedRights
// update chats set default_banned_rights = :default_banned_rights, version = version + 1 where id = :id
func (dao *ChatsDAO) UpdateDefaultBannedRights(ctx context.Context, defaultBannedRights int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set default_banned_rights = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, defaultBannedRights, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateDefaultBannedRights(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateDefaultBannedRights(_), error: %v", err)
	}

	return
}

// UpdateDefaultBannedRightsTx
// update chats set default_banned_rights = :default_banned_rights, version = version + 1 where id = :id
func (dao *ChatsDAO) UpdateDefaultBannedRightsTx(tx *sqlx.Tx, defaultBannedRights int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set default_banned_rights = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, defaultBannedRights, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateDefaultBannedRights(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateDefaultBannedRights(_), error: %v", err)
	}

	return
}

// UpdateVersion
// update chats set version = version + 1 where id = :id
func (dao *ChatsDAO) UpdateVersion(ctx context.Context, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set version = version + 1 where id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateVersion(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateVersion(_), error: %v", err)
	}

	return
}

// UpdateVersionTx
// update chats set version = version + 1 where id = :id
func (dao *ChatsDAO) UpdateVersionTx(tx *sqlx.Tx, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateVersion(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateVersion(_), error: %v", err)
	}

	return
}

// UpdateDeactivated
// update chats set deactivated = :deactivated, version = version + 1 where id = :id
func (dao *ChatsDAO) UpdateDeactivated(ctx context.Context, deactivated bool, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set deactivated = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, deactivated, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateDeactivated(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateDeactivated(_), error: %v", err)
	}

	return
}

// UpdateDeactivatedTx
// update chats set deactivated = :deactivated, version = version + 1 where id = :id
func (dao *ChatsDAO) UpdateDeactivatedTx(tx *sqlx.Tx, deactivated bool, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set deactivated = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, deactivated, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateDeactivated(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateDeactivated(_), error: %v", err)
	}

	return
}

// SelectByLink
// select id, creator_user_id, access_hash, participant_count, title, about, photo_id, link, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where deactivated = 0 and link = :link
func (dao *ChatsDAO) SelectByLink(ctx context.Context) (rValue *dataobject.ChatsDO, err error) {
	var (
		query = "select id, creator_user_id, access_hash, participant_count, title, about, photo_id, link, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where deactivated = 0 and link = ?"
		do    = &dataobject.ChatsDO{}
	)
	err = dao.db.QueryRowPartial(ctx, do, query)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectByLink(_), error: %v", err)
			return
		} else {
			err = nil
		}
	} else {
		rValue = do
	}

	return
}

// UpdateLink
// update chats set link = :link, `date` = :date, version = version + 1 where id = :id
func (dao *ChatsDAO) UpdateLink(ctx context.Context, date int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set link = ?, `date` = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, date, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateLink(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateLink(_), error: %v", err)
	}

	return
}

// UpdateLinkTx
// update chats set link = :link, `date` = :date, version = version + 1 where id = :id
func (dao *ChatsDAO) UpdateLinkTx(tx *sqlx.Tx, date int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set link = ?, `date` = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, date, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateLink(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateLink(_), error: %v", err)
	}

	return
}

// UpdateMigratedTo
// update chats set migrated_to_id = :migrated_to_id, migrated_to_access_hash = :migrated_to_access_hash, participant_count = 0, deactivated = 1, version = version + 1 where id = :id
func (dao *ChatsDAO) UpdateMigratedTo(ctx context.Context, migratedToId int64, migratedToAccessHash int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set migrated_to_id = ?, migrated_to_access_hash = ?, participant_count = 0, deactivated = 1, version = version + 1 where id = ?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, migratedToId, migratedToAccessHash, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateMigratedTo(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateMigratedTo(_), error: %v", err)
	}

	return
}

// UpdateMigratedToTx
// update chats set migrated_to_id = :migrated_to_id, migrated_to_access_hash = :migrated_to_access_hash, participant_count = 0, deactivated = 1, version = version + 1 where id = :id
func (dao *ChatsDAO) UpdateMigratedToTx(tx *sqlx.Tx, migratedToId int64, migratedToAccessHash int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set migrated_to_id = ?, migrated_to_access_hash = ?, participant_count = 0, deactivated = 1, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, migratedToId, migratedToAccessHash, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateMigratedTo(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateMigratedTo(_), error: %v", err)
	}

	return
}
