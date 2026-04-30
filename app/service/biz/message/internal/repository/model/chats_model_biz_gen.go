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

type bizChatsModel interface {
	Insert(ctx context.Context, data *Chats) (lastInsertId, rowsAffected int64, err error)
	Select(ctx context.Context, id int64) (*Chats, error)
	SelectPhotoId(ctx context.Context, id int64) (int64, error)
	SelectLastCreator(ctx context.Context, creatorUserId int64) (*Chats, error)
	UpdateTitle(ctx context.Context, title string, id int64) (rowsAffected int64, err error)
	UpdateAbout(ctx context.Context, about string, id int64) (rowsAffected int64, err error)
	SelectByIdList(ctx context.Context, idList []int32) ([]Chats, error)
	SelectByIdListWithCB(ctx context.Context, idList []int32, cb func(sz, i int, v *Chats)) ([]Chats, error)
	UpdateParticipantCount(ctx context.Context, participantCount int32, id int64) (rowsAffected int64, err error)
	UpdatePhotoId(ctx context.Context, photoId int64, id int64) (rowsAffected int64, err error)
	UpdateAdminsEnabled(ctx context.Context, id int64) (rowsAffected int64, err error)
	UpdateDefaultBannedRights(ctx context.Context, defaultBannedRights int64, id int64) (rowsAffected int64, err error)
	UpdateVersion(ctx context.Context, id int64) (rowsAffected int64, err error)
	UpdateDeactivated(ctx context.Context, deactivated bool, id int64) (rowsAffected int64, err error)
	SelectByLink(ctx context.Context) (*Chats, error)
	UpdateLink(ctx context.Context, date int64, id int64) (rowsAffected int64, err error)
	UpdateMigratedTo(ctx context.Context, migratedToId int64, migratedToAccessHash int64, id int64) (rowsAffected int64, err error)
}

type ChatsTxModel interface {
	Insert(data *Chats) (lastInsertId, rowsAffected int64, err error)
	Select(id int64) (*Chats, error)
	SelectPhotoId(id int64) (int64, error)
	SelectLastCreator(creatorUserId int64) (*Chats, error)
	UpdateTitle(title string, id int64) (rowsAffected int64, err error)
	UpdateAbout(about string, id int64) (rowsAffected int64, err error)
	SelectByIdList(idList []int32) ([]Chats, error)
	UpdateParticipantCount(participantCount int32, id int64) (rowsAffected int64, err error)
	UpdatePhotoId(photoId int64, id int64) (rowsAffected int64, err error)
	UpdateAdminsEnabled(id int64) (rowsAffected int64, err error)
	UpdateDefaultBannedRights(defaultBannedRights int64, id int64) (rowsAffected int64, err error)
	UpdateVersion(id int64) (rowsAffected int64, err error)
	UpdateDeactivated(deactivated bool, id int64) (rowsAffected int64, err error)
	SelectByLink() (*Chats, error)
	UpdateLink(date int64, id int64) (rowsAffected int64, err error)
	UpdateMigratedTo(migratedToId int64, migratedToAccessHash int64, id int64) (rowsAffected int64, err error)
}

type defaultChatsTxModel struct {
	tx *sqlx.Tx
}

func NewChatsTxModel(tx *sqlx.Tx) ChatsTxModel {
	return &defaultChatsTxModel{tx: tx}
}

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

// Insert
// insert into chats(creator_user_id, access_hash, random_id, participant_count, title, about, `date`) values (:creator_user_id, :access_hash, :random_id, :participant_count, :title, :about, :date)
func (m *defaultChatsTxModel) Insert(data *Chats) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chats(creator_user_id, access_hash, random_id, participant_count, title, about, `date`) values (:creator_user_id, :access_hash, :random_id, :participant_count, :title, :about, :date)"
		r     sql.Result
	)

	r, err = m.tx.NamedExec(query, data)
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

// Select
// select id, creator_user_id, access_hash, participant_count, title, about, photo_id, link, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where id = :id
func (m *defaultChatsModel) Select(ctx context.Context, id int64) (rValue *Chats, err error) {

	var (
		query = "select id, creator_user_id, access_hash, participant_count, title, about, photo_id, link, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where id = ?"
		do    = &Chats{}
	)
	err = m.db.QueryRowPartial(ctx, do, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "chats",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		err = fmt.Errorf("chats.Select: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// Select
// select id, creator_user_id, access_hash, participant_count, title, about, photo_id, link, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where id = :id
func (m *defaultChatsTxModel) Select(id int64) (rValue *Chats, err error) {
	var (
		query = "select id, creator_user_id, access_hash, participant_count, title, about, photo_id, link, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where id = ?"
		do    = &Chats{}
	)
	err = m.tx.QueryRowPartial(do, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "chats",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
		}
		err = fmt.Errorf("chats.Select: %w", err)
		return
	}
	rValue = do

	return
}

// SelectPhotoId
// select photo_id from chats where id = :id
func (m *defaultChatsModel) SelectPhotoId(ctx context.Context, id int64) (rValue int64, err error) {
	var query = "select photo_id from chats where id = ?"
	err = m.db.QueryRowPartial(ctx, &rValue, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			err = &NotFoundError{
				Resource: "chats",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
			return
		}
		err = fmt.Errorf("chats.SelectPhotoId: %w", err)
		return
	}

	return
}

// SelectPhotoId
// select photo_id from chats where id = :id
func (m *defaultChatsTxModel) SelectPhotoId(id int64) (rValue int64, err error) {
	var query = "select photo_id from chats where id = ?"
	err = m.tx.QueryRowPartial(&rValue, query, id)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			err = &NotFoundError{
				Resource: "chats",
				Key:      fmt.Sprintf("id=%v", id),
				Cause:    err,
			}
			return
		}
		err = fmt.Errorf("chats.SelectPhotoId: %w", err)
		return
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
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "chats",
				Key:      fmt.Sprintf("creator_user_id=%v", creatorUserId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("chats.SelectLastCreator: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectLastCreator
// select id, creator_user_id, access_hash, participant_count, title, about, photo_id, link, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where creator_user_id = :creator_user_id order by `date` desc limit 1
func (m *defaultChatsTxModel) SelectLastCreator(creatorUserId int64) (rValue *Chats, err error) {
	var (
		query = "select id, creator_user_id, access_hash, participant_count, title, about, photo_id, link, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where creator_user_id = ? order by `date` desc limit 1"
		do    = &Chats{}
	)
	err = m.tx.QueryRowPartial(do, query, creatorUserId)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "chats",
				Key:      fmt.Sprintf("creator_user_id=%v", creatorUserId),
				Cause:    err,
			}
		}
		err = fmt.Errorf("chats.SelectLastCreator: %w", err)
		return
	}
	rValue = do

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
		return
	}

	return
}

// UpdateTitle
// update chats set title = :title, version = version + 1 where id = :id
func (m *defaultChatsTxModel) UpdateTitle(title string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set title = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, title, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateTitle exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateTitle rows affected: %w", err)
		return
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
		return
	}

	return
}

// UpdateAbout
// update chats set about = :about where id = :id
func (m *defaultChatsTxModel) UpdateAbout(about string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set about = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, about, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateAbout exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateAbout rows affected: %w", err)
		return
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Chats{}
			err = nil
			return
		}
		err = fmt.Errorf("chats.SelectByIdList: %w", err)
		return
	}

	rList = values

	return
}

// SelectByIdList
// select id, creator_user_id, access_hash, participant_count, title, about, photo_id, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where id in (:idList)
func (m *defaultChatsTxModel) SelectByIdList(idList []int32) (rList []Chats, err error) {
	var (
		query  = fmt.Sprintf("select id, creator_user_id, access_hash, participant_count, title, about, photo_id, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where id in (%s)", sqlx.InInt32List(idList))
		values []Chats
	)
	if len(idList) == 0 {
		rList = []Chats{}
		return
	}

	err = m.tx.QueryRowsPartial(&values, query)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Chats{}
			err = nil
			return
		}
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
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []Chats{}
			err = nil
			return
		}
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
		return
	}

	return
}

// UpdateParticipantCount
// update chats set participant_count = :participant_count, version = version + 1 where id = :id
func (m *defaultChatsTxModel) UpdateParticipantCount(participantCount int32, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set participant_count = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, participantCount, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateParticipantCount exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateParticipantCount rows affected: %w", err)
		return
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
		return
	}

	return
}

// UpdatePhotoId
// update chats set photo_id = :photo_id, version = version + 1 where id = :id
func (m *defaultChatsTxModel) UpdatePhotoId(photoId int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set photo_id = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, photoId, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdatePhotoId exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdatePhotoId rows affected: %w", err)
		return
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
		return
	}

	return
}

// UpdateAdminsEnabled
// update chats set admins_enabled = :admins_enabled, version = version + 1 where id = :id
func (m *defaultChatsTxModel) UpdateAdminsEnabled(id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set admins_enabled = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateAdminsEnabled exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateAdminsEnabled rows affected: %w", err)
		return
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
		return
	}

	return
}

// UpdateDefaultBannedRights
// update chats set default_banned_rights = :default_banned_rights, version = version + 1 where id = :id
func (m *defaultChatsTxModel) UpdateDefaultBannedRights(defaultBannedRights int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set default_banned_rights = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, defaultBannedRights, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateDefaultBannedRights exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateDefaultBannedRights rows affected: %w", err)
		return
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
		return
	}

	return
}

// UpdateVersion
// update chats set version = version + 1 where id = :id
func (m *defaultChatsTxModel) UpdateVersion(id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateVersion exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateVersion rows affected: %w", err)
		return
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
		return
	}

	return
}

// UpdateDeactivated
// update chats set deactivated = :deactivated, version = version + 1 where id = :id
func (m *defaultChatsTxModel) UpdateDeactivated(deactivated bool, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set deactivated = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, deactivated, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateDeactivated exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateDeactivated rows affected: %w", err)
		return
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
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "chats",
				Key:      fmt.Sprintf(""),
				Cause:    err,
			}
		}
		err = fmt.Errorf("chats.SelectByLink: %w", err)
		return
	} else {
		rValue = do
	}

	return
}

// SelectByLink
// select id, creator_user_id, access_hash, participant_count, title, about, photo_id, link, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where deactivated = 0 and link = :link
func (m *defaultChatsTxModel) SelectByLink() (rValue *Chats, err error) {
	var (
		query = "select id, creator_user_id, access_hash, participant_count, title, about, photo_id, link, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where deactivated = 0 and link = ?"
		do    = &Chats{}
	)
	err = m.tx.QueryRowPartial(do, query)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, &NotFoundError{
				Resource: "chats",
				Key:      fmt.Sprintf(""),
				Cause:    err,
			}
		}
		err = fmt.Errorf("chats.SelectByLink: %w", err)
		return
	}
	rValue = do

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
		return
	}

	return
}

// UpdateLink
// update chats set link = :link, `date` = :date, version = version + 1 where id = :id
func (m *defaultChatsTxModel) UpdateLink(date int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set link = ?, `date` = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, date, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateLink exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateLink rows affected: %w", err)
		return
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
		return
	}

	return
}

// UpdateMigratedTo
// update chats set migrated_to_id = :migrated_to_id, migrated_to_access_hash = :migrated_to_access_hash, participant_count = 0, deactivated = 1, version = version + 1 where id = :id
func (m *defaultChatsTxModel) UpdateMigratedTo(migratedToId int64, migratedToAccessHash int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set migrated_to_id = ?, migrated_to_access_hash = ?, participant_count = 0, deactivated = 1, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = m.tx.Exec(query, migratedToId, migratedToAccessHash, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateMigratedTo exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateMigratedTo rows affected: %w", err)
		return
	}

	return
}
