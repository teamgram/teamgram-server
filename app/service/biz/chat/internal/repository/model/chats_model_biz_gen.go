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

		UpdateDefaultBannedRights(ctx context.Context, defaultBannedRights int64, id int64) (rowsAffected int64, err error)
		UpdateDefaultBannedRightsTx(tx *sqlx.Tx, defaultBannedRights int64, id int64) (rowsAffected int64, err error)

		UpdateVersion(ctx context.Context, id int64) (rowsAffected int64, err error)
		UpdateVersionTx(tx *sqlx.Tx, id int64) (rowsAffected int64, err error)

		UpdateDeactivated(ctx context.Context, deactivated bool, id int64) (rowsAffected int64, err error)
		UpdateDeactivatedTx(tx *sqlx.Tx, deactivated bool, id int64) (rowsAffected int64, err error)

		UpdateMigratedTo(ctx context.Context, migratedToId int64, migratedToAccessHash int64, id int64) (rowsAffected int64, err error)
		UpdateMigratedToTx(tx *sqlx.Tx, migratedToId int64, migratedToAccessHash int64, id int64) (rowsAffected int64, err error)

		UpdateAvailableReactions(ctx context.Context, availableReactionsType int32, availableReactions string, id int64) (rowsAffected int64, err error)
		UpdateAvailableReactionsTx(tx *sqlx.Tx, availableReactionsType int32, availableReactions string, id int64) (rowsAffected int64, err error)

		UpdateNoforwards(ctx context.Context, noforwards bool, id int64) (rowsAffected int64, err error)
		UpdateNoforwardsTx(tx *sqlx.Tx, noforwards bool, id int64) (rowsAffected int64, err error)

		UpdateTTLPeriod(ctx context.Context, ttlPeriod int32, id int64) (rowsAffected int64, err error)
		UpdateTTLPeriodTx(tx *sqlx.Tx, ttlPeriod int32, id int64) (rowsAffected int64, err error)

		SearchByQueryString(ctx context.Context, q string, limit int32) ([]int64, error)
		SearchByQueryStringWithCB(ctx context.Context, q string, limit int32, cb func(sz, i int, v int64)) ([]int64, error)
	}
)

// Insert
// insert into chats(creator_user_id, access_hash, random_id, participant_count, title, about, default_banned_rights, `date`) values (:creator_user_id, :access_hash, :random_id, :participant_count, :title, :about, :default_banned_rights, :date)
func (m *defaultChatsModel) Insert(ctx context.Context, data *Chats) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chats(creator_user_id, access_hash, random_id, participant_count, title, about, default_banned_rights, `date`) values (:creator_user_id, :access_hash, :random_id, :participant_count, :title, :about, :default_banned_rights, :date)"
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
// insert into chats(creator_user_id, access_hash, random_id, participant_count, title, about, default_banned_rights, `date`) values (:creator_user_id, :access_hash, :random_id, :participant_count, :title, :about, :default_banned_rights, :date)
func (m *defaultChatsModel) InsertTx(tx *sqlx.Tx, data *Chats) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chats(creator_user_id, access_hash, random_id, participant_count, title, about, default_banned_rights, `date`) values (:creator_user_id, :access_hash, :random_id, :participant_count, :title, :about, :default_banned_rights, :date)"
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
// select id, creator_user_id, access_hash, participant_count, title, about, photo_id, default_banned_rights, migrated_to_id, migrated_to_access_hash, noforwards, available_reactions_type, available_reactions, deactivated, ttl_period, version, `date` from chats where id = :id
func (m *defaultChatsModel) Select(ctx context.Context, id int64) (rValue *Chats, err error) {

	var (
		query = "select id, creator_user_id, access_hash, participant_count, title, about, photo_id, default_banned_rights, migrated_to_id, migrated_to_access_hash, noforwards, available_reactions_type, available_reactions, deactivated, ttl_period, version, `date` from chats where id = ?"
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

// SelectLastCreator
// select id, creator_user_id, access_hash, participant_count, title, about, photo_id, default_banned_rights, migrated_to_id, migrated_to_access_hash, noforwards, available_reactions_type, available_reactions, deactivated, ttl_period, version, `date` from chats where creator_user_id = :creator_user_id order by `date` desc limit 1
func (m *defaultChatsModel) SelectLastCreator(ctx context.Context, creatorUserId int64) (rValue *Chats, err error) {

	var (
		query = "select id, creator_user_id, access_hash, participant_count, title, about, photo_id, default_banned_rights, migrated_to_id, migrated_to_access_hash, noforwards, available_reactions_type, available_reactions, deactivated, ttl_period, version, `date` from chats where creator_user_id = ? order by `date` desc limit 1"
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
		return
	}

	return
}

// SelectByIdList
// select id, creator_user_id, access_hash, participant_count, title, about, photo_id, default_banned_rights, migrated_to_id, migrated_to_access_hash, noforwards, available_reactions_type, available_reactions, deactivated, ttl_period, version, `date` from chats where id in (:idList)
func (m *defaultChatsModel) SelectByIdList(ctx context.Context, idList []int32) (rList []Chats, err error) {
	var (
		query  = fmt.Sprintf("select id, creator_user_id, access_hash, participant_count, title, about, photo_id, default_banned_rights, migrated_to_id, migrated_to_access_hash, noforwards, available_reactions_type, available_reactions, deactivated, ttl_period, version, `date` from chats where id in (%s)", sqlx.InInt32List(idList))
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

// SelectByIdListWithCB
// select id, creator_user_id, access_hash, participant_count, title, about, photo_id, default_banned_rights, migrated_to_id, migrated_to_access_hash, noforwards, available_reactions_type, available_reactions, deactivated, ttl_period, version, `date` from chats where id in (:idList)
func (m *defaultChatsModel) SelectByIdListWithCB(ctx context.Context, idList []int32, cb func(sz, i int, v *Chats)) (rList []Chats, err error) {
	var (
		query  = fmt.Sprintf("select id, creator_user_id, access_hash, participant_count, title, about, photo_id, default_banned_rights, migrated_to_id, migrated_to_access_hash, noforwards, available_reactions_type, available_reactions, deactivated, ttl_period, version, `date` from chats where id in (%s)", sqlx.InInt32List(idList))
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
		return
	}

	return
}

// UpdateAvailableReactions
// update chats set available_reactions_type = :available_reactions_type, available_reactions = :available_reactions where id = :id
func (m *defaultChatsModel) UpdateAvailableReactions(ctx context.Context, availableReactionsType int32, availableReactions string, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update chats set available_reactions_type = ?, available_reactions = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, availableReactionsType, availableReactions, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateAvailableReactions exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateAvailableReactions rows affected: %w", err)
		return
	}

	return
}

// UpdateAvailableReactionsTx
// update chats set available_reactions_type = :available_reactions_type, available_reactions = :available_reactions where id = :id
func (m *defaultChatsModel) UpdateAvailableReactionsTx(tx *sqlx.Tx, availableReactionsType int32, availableReactions string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set available_reactions_type = ?, available_reactions = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, availableReactionsType, availableReactions, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateAvailableReactionsTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateAvailableReactionsTx rows affected: %w", err)
		return
	}

	return
}

// UpdateNoforwards
// update chats set noforwards = :noforwards, version = version + 1 where id = :id
func (m *defaultChatsModel) UpdateNoforwards(ctx context.Context, noforwards bool, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update chats set noforwards = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, noforwards, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateNoforwards exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateNoforwards rows affected: %w", err)
		return
	}

	return
}

// UpdateNoforwardsTx
// update chats set noforwards = :noforwards, version = version + 1 where id = :id
func (m *defaultChatsModel) UpdateNoforwardsTx(tx *sqlx.Tx, noforwards bool, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set noforwards = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, noforwards, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateNoforwardsTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateNoforwardsTx rows affected: %w", err)
		return
	}

	return
}

// UpdateTTLPeriod
// update chats set ttl_period = :ttl_period where id = :id
func (m *defaultChatsModel) UpdateTTLPeriod(ctx context.Context, ttlPeriod int32, id int64) (rowsAffected int64, err error) {

	var (
		query   = "update chats set ttl_period = ? where id = ?"
		rResult sql.Result
	)

	rResult, err = m.db.Exec(ctx, query, ttlPeriod, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateTTLPeriod exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateTTLPeriod rows affected: %w", err)
		return
	}

	return
}

// UpdateTTLPeriodTx
// update chats set ttl_period = :ttl_period where id = :id
func (m *defaultChatsModel) UpdateTTLPeriodTx(tx *sqlx.Tx, ttlPeriod int32, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set ttl_period = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, ttlPeriod, id)

	if err != nil {
		err = fmt.Errorf("chats.UpdateTTLPeriodTx exec: %w", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		err = fmt.Errorf("chats.UpdateTTLPeriodTx rows affected: %w", err)
		return
	}

	return
}

// SearchByQueryString
// select id from chats where title like :q limit :limit
func (m *defaultChatsModel) SearchByQueryString(ctx context.Context, q string, limit int32) (rList []int64, err error) {
	var query = "select id from chats where title like ? limit ?"
	err = m.db.QueryRowsPartial(ctx, &rList, query, q, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []int64{}
			err = nil
			return
		}
		err = fmt.Errorf("chats.SearchByQueryString: %w", err)
	}

	return
}

// SearchByQueryStringWithCB
// select id from chats where title like :q limit :limit
func (m *defaultChatsModel) SearchByQueryStringWithCB(ctx context.Context, q string, limit int32, cb func(sz, i int, v int64)) (rList []int64, err error) {
	var query = "select id from chats where title like ? limit ?"
	err = m.db.QueryRowsPartial(ctx, &rList, query, q, limit)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			rList = []int64{}
			err = nil
			return
		}
		err = fmt.Errorf("chats.SearchByQueryStringWithCB: %w", err)
		return
	}

	if cb != nil {
		sz := len(rList)
		for i := 0; i < sz; i++ {
			cb(sz, i, rList[i])
		}
	}

	return
}
