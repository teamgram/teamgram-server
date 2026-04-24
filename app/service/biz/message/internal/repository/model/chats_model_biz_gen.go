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
	bizChatsModel interface {
		Insert(ctx context.Context, data *Chats) (lastInsertId, rowsAffected int64, err error)
		InsertTx(tx *sqlx.Tx, data *Chats) (lastInsertId, rowsAffected int64, err error)

		Select(ctx context.Context, id int64) (*Chats, error)

		SelectPhotoId(ctx context.Context, id int64) (int64, error)

		SelectLastCreator(ctx context.Context, creatorUserId int64) (*Chats, error)

		UpdateTitle(ctx context.Context, title string, id int64) (rowsAffected int64, err error)
		UpdateTitleTx(tx *sqlx.Tx, title string, id int64) (rowsAffected int64, err error)

		UpdateAbout(ctx context.Context, about string, id int64) (rowsAffected int64, err error)
		UpdateAboutTx(tx *sqlx.Tx, about string, id int64) (rowsAffected int64, err error)

		SelectByIdList(ctx context.Context, idList []int32) ([]Chats, error)
		SelectByIdListWithCB(ctx context.Context, idList []int32, cb func(sz, i int, v *Chats)) ([]Chats, error)

		UpdateParticipantCount(ctx context.Context, participantCount int32, id int64) (rowsAffected int64, err error)
		UpdateParticipantCountTx(tx *sqlx.Tx, participantCount int32, id int64) (rowsAffected int64, err error)

		UpdatePhotoId(ctx context.Context, photoId int64, id int64) (rowsAffected int64, err error)
		UpdatePhotoIdTx(tx *sqlx.Tx, photoId int64, id int64) (rowsAffected int64, err error)

		UpdateAdminsEnabled(ctx context.Context, id int64) (rowsAffected int64, err error)
		UpdateAdminsEnabledTx(tx *sqlx.Tx, id int64) (rowsAffected int64, err error)

		UpdateDefaultBannedRights(ctx context.Context, defaultBannedRights int64, id int64) (rowsAffected int64, err error)
		UpdateDefaultBannedRightsTx(tx *sqlx.Tx, defaultBannedRights int64, id int64) (rowsAffected int64, err error)

		UpdateVersion(ctx context.Context, id int64) (rowsAffected int64, err error)
		UpdateVersionTx(tx *sqlx.Tx, id int64) (rowsAffected int64, err error)

		UpdateDeactivated(ctx context.Context, deactivated bool, id int64) (rowsAffected int64, err error)
		UpdateDeactivatedTx(tx *sqlx.Tx, deactivated bool, id int64) (rowsAffected int64, err error)

		SelectByLink(ctx context.Context) (*Chats, error)

		UpdateLink(ctx context.Context, date int64, id int64) (rowsAffected int64, err error)
		UpdateLinkTx(tx *sqlx.Tx, date int64, id int64) (rowsAffected int64, err error)

		UpdateMigratedTo(ctx context.Context, migratedToId int64, migratedToAccessHash int64, id int64) (rowsAffected int64, err error)
		UpdateMigratedToTx(tx *sqlx.Tx, migratedToId int64, migratedToAccessHash int64, id int64) (rowsAffected int64, err error)
	}
)

// Insert
// insert into chats(creator_user_id, access_hash, random_id, participant_count, title, about, `date`) values (:creator_user_id, :access_hash, :random_id, :participant_count, :title, :about, :date)
func (m *defaultChatsModel) Insert(ctx context.Context, data *Chats) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chats(creator_user_id, access_hash, random_id, participant_count, title, about, `date`) values (:creator_user_id, :access_hash, :random_id, :participant_count, :title, :about, :date)"
		r     sql.Result
	)

	r, err = m.db.NamedExec(ctx, query, data)
	if err != nil {
		err = fmt.Errorf("chats.Insert named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("chats.Insert last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.Insert rows affected: %w", err)
	}

	return

}

// InsertTx
// insert into chats(creator_user_id, access_hash, random_id, participant_count, title, about, `date`) values (:creator_user_id, :access_hash, :random_id, :participant_count, :title, :about, :date)
func (m *defaultChatsModel) InsertTx(tx *sqlx.Tx, data *Chats) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chats(creator_user_id, access_hash, random_id, participant_count, title, about, `date`) values (:creator_user_id, :access_hash, :random_id, :participant_count, :title, :about, :date)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, data)
	if err != nil {
		err = fmt.Errorf("chats.InsertTx named exec: %w", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		err = fmt.Errorf("chats.InsertTx last insert id: %w", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.InsertTx rows affected: %w", err)
	}

	return
}

// Select
// select id, creator_user_id, access_hash, participant_count, title, about, photo_id, link, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where id = :id
func (m *defaultChatsModel) Select(ctx context.Context, id int64) (rValue *Chats, err error) {

	var (
		query = "select id, creator_user_id, access_hash, participant_count, title, about, photo_id, link, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where id = ?"
		do    = &Chats{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, id)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			err = fmt.Errorf("chats.Select: %w", err)
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
func (m *defaultChatsModel) SelectPhotoId(ctx context.Context, id int64) (rValue int64, err error) {
	var query = "select photo_id from chats where id = ?"
	err = m.db.QueryRowPartial(ctx, &rValue, query, id)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			err = fmt.Errorf("chats.SelectPhotoId: %w", err)
			return
		} else {
			err = nil
		}
	}

	return
}

// SelectLastCreator
// select id, creator_user_id, access_hash, participant_count, title, about, photo_id, link, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where creator_user_id = :creator_user_id order by `date` desc limit 1
func (m *defaultChatsModel) SelectLastCreator(ctx context.Context, creatorUserId int64) (rValue *Chats, err error) {

	var (
		query = "select id, creator_user_id, access_hash, participant_count, title, about, photo_id, link, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where creator_user_id = ? order by `date` desc limit 1"
		do    = &Chats{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, creatorUserId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			err = fmt.Errorf("chats.SelectLastCreator: %w", err)
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
func (m *defaultChatsModel) UpdateTitle(ctx context.Context, title string, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update chats set title = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, title, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateTitle exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateTitle rows affected: %w", err)
	}

	return
}

// UpdateTitleTx
// update chats set title = :title, version = version + 1 where id = :id
func (m *defaultChatsModel) UpdateTitleTx(tx *sqlx.Tx, title string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set title = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, title, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateTitleTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateTitleTx rows affected: %w", err)
	}

	return
}

// UpdateAbout
// update chats set about = :about where id = :id
func (m *defaultChatsModel) UpdateAbout(ctx context.Context, about string, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update chats set about = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, about, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateAbout exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateAbout rows affected: %w", err)
	}

	return
}

// UpdateAboutTx
// update chats set about = :about where id = :id
func (m *defaultChatsModel) UpdateAboutTx(tx *sqlx.Tx, about string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set about = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, about, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateAboutTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateAboutTx rows affected: %w", err)
	}

	return
}

// SelectByIdList
// select id, creator_user_id, access_hash, participant_count, title, about, photo_id, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where id in (:idList)
func (m *defaultChatsModel) SelectByIdList(ctx context.Context, idList []int32) (rList []Chats, err error) {
	var (
		query  = fmt.Sprintf("select id, creator_user_id, access_hash, participant_count, title, about, photo_id, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where id in (%s)", sqlx.InInt32List(idList))
		values []Chats
	)
	if len(idList) == 0 {
		rList = []Chats{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		err = fmt.Errorf("chats.SelectByIdList: %w", err)
		return
	}

	rList = values

	return
}

// SelectByIdListWithCB
// select id, creator_user_id, access_hash, participant_count, title, about, photo_id, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where id in (:idList)
func (m *defaultChatsModel) SelectByIdListWithCB(ctx context.Context, idList []int32, cb func(sz, i int, v *Chats)) (rList []Chats, err error) {
	var (
		query  = fmt.Sprintf("select id, creator_user_id, access_hash, participant_count, title, about, photo_id, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where id in (%s)", sqlx.InInt32List(idList))
		values []Chats
	)
	if len(idList) == 0 {
		rList = []Chats{}
		return
	}

	err = m.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		err = fmt.Errorf("chats.SelectByIdListWithCB: %w", err)
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
func (m *defaultChatsModel) UpdateParticipantCount(ctx context.Context, participantCount int32, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update chats set participant_count = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, participantCount, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateParticipantCount exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateParticipantCount rows affected: %w", err)
	}

	return
}

// UpdateParticipantCountTx
// update chats set participant_count = :participant_count, version = version + 1 where id = :id
func (m *defaultChatsModel) UpdateParticipantCountTx(tx *sqlx.Tx, participantCount int32, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set participant_count = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, participantCount, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateParticipantCountTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateParticipantCountTx rows affected: %w", err)
	}

	return
}

// UpdatePhotoId
// update chats set photo_id = :photo_id, version = version + 1 where id = :id
func (m *defaultChatsModel) UpdatePhotoId(ctx context.Context, photoId int64, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update chats set photo_id = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, photoId, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdatePhotoId exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdatePhotoId rows affected: %w", err)
	}

	return
}

// UpdatePhotoIdTx
// update chats set photo_id = :photo_id, version = version + 1 where id = :id
func (m *defaultChatsModel) UpdatePhotoIdTx(tx *sqlx.Tx, photoId int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set photo_id = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, photoId, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdatePhotoIdTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdatePhotoIdTx rows affected: %w", err)
	}

	return
}

// UpdateAdminsEnabled
// update chats set admins_enabled = :admins_enabled, version = version + 1 where id = :id
func (m *defaultChatsModel) UpdateAdminsEnabled(ctx context.Context, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update chats set admins_enabled = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateAdminsEnabled exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateAdminsEnabled rows affected: %w", err)
	}

	return
}

// UpdateAdminsEnabledTx
// update chats set admins_enabled = :admins_enabled, version = version + 1 where id = :id
func (m *defaultChatsModel) UpdateAdminsEnabledTx(tx *sqlx.Tx, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set admins_enabled = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateAdminsEnabledTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateAdminsEnabledTx rows affected: %w", err)
	}

	return
}

// UpdateDefaultBannedRights
// update chats set default_banned_rights = :default_banned_rights, version = version + 1 where id = :id
func (m *defaultChatsModel) UpdateDefaultBannedRights(ctx context.Context, defaultBannedRights int64, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update chats set default_banned_rights = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, defaultBannedRights, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateDefaultBannedRights exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateDefaultBannedRights rows affected: %w", err)
	}

	return
}

// UpdateDefaultBannedRightsTx
// update chats set default_banned_rights = :default_banned_rights, version = version + 1 where id = :id
func (m *defaultChatsModel) UpdateDefaultBannedRightsTx(tx *sqlx.Tx, defaultBannedRights int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set default_banned_rights = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, defaultBannedRights, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateDefaultBannedRightsTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateDefaultBannedRightsTx rows affected: %w", err)
	}

	return
}

// UpdateVersion
// update chats set version = version + 1 where id = :id
func (m *defaultChatsModel) UpdateVersion(ctx context.Context, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update chats set version = version + 1 where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateVersion exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateVersion rows affected: %w", err)
	}

	return
}

// UpdateVersionTx
// update chats set version = version + 1 where id = :id
func (m *defaultChatsModel) UpdateVersionTx(tx *sqlx.Tx, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateVersionTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateVersionTx rows affected: %w", err)
	}

	return
}

// UpdateDeactivated
// update chats set deactivated = :deactivated, version = version + 1 where id = :id
func (m *defaultChatsModel) UpdateDeactivated(ctx context.Context, deactivated bool, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update chats set deactivated = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, deactivated, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateDeactivated exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateDeactivated rows affected: %w", err)
	}

	return
}

// UpdateDeactivatedTx
// update chats set deactivated = :deactivated, version = version + 1 where id = :id
func (m *defaultChatsModel) UpdateDeactivatedTx(tx *sqlx.Tx, deactivated bool, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set deactivated = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, deactivated, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateDeactivatedTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateDeactivatedTx rows affected: %w", err)
	}

	return
}

// SelectByLink
// select id, creator_user_id, access_hash, participant_count, title, about, photo_id, link, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where deactivated = 0 and link = :link
func (m *defaultChatsModel) SelectByLink(ctx context.Context) (rValue *Chats, err error) {

	var (
		query = "select id, creator_user_id, access_hash, participant_count, title, about, photo_id, link, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where deactivated = 0 and link = ?"
		do    = &Chats{}
	)
	err = m.db.QueryRowPartial(ctx, do, query)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			err = fmt.Errorf("chats.SelectByLink: %w", err)
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
func (m *defaultChatsModel) UpdateLink(ctx context.Context, date int64, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update chats set link = ?, `date` = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, date, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateLink exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateLink rows affected: %w", err)
	}

	return
}

// UpdateLinkTx
// update chats set link = :link, `date` = :date, version = version + 1 where id = :id
func (m *defaultChatsModel) UpdateLinkTx(tx *sqlx.Tx, date int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set link = ?, `date` = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, date, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateLinkTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateLinkTx rows affected: %w", err)
	}

	return
}

// UpdateMigratedTo
// update chats set migrated_to_id = :migrated_to_id, migrated_to_access_hash = :migrated_to_access_hash, participant_count = 0, deactivated = 1, version = version + 1 where id = :id
func (m *defaultChatsModel) UpdateMigratedTo(ctx context.Context, migratedToId int64, migratedToAccessHash int64, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update chats set migrated_to_id = ?, migrated_to_access_hash = ?, participant_count = 0, deactivated = 1, version = version + 1 where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, migratedToId, migratedToAccessHash, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateMigratedTo exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateMigratedTo rows affected: %w", err)
	}

	return
}

// UpdateMigratedToTx
// update chats set migrated_to_id = :migrated_to_id, migrated_to_access_hash = :migrated_to_access_hash, participant_count = 0, deactivated = 1, version = version + 1 where id = :id
func (m *defaultChatsModel) UpdateMigratedToTx(tx *sqlx.Tx, migratedToId int64, migratedToAccessHash int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set migrated_to_id = ?, migrated_to_access_hash = ?, participant_count = 0, deactivated = 1, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, migratedToId, migratedToAccessHash, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateMigratedToTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateMigratedToTx rows affected: %w", err)
	}

	return
}
