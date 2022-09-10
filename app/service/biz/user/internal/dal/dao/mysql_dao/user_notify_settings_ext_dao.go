/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teagramio (teagram.io@gmail.com)
 */

package mysql_dao

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/teamgram/proto/mtproto"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

// InsertOrUpdateExt
// insert into user_notify_settings(user_id, peer_type, peer_id, show_previews, silent, mute_until, sound) values (:user_id, :peer_type, :peer_id, :show_previews, :silent, :mute_until, :sound) on duplicate key update show_previews = values(show_previews), silent = values(silent), mute_until = values(mute_until), sound = values(sound), deleted = 0
// TODO(@benqi): sqlmap
func (dao *UserNotifySettingsDAO) InsertOrUpdateExt(ctx context.Context, userId int64, peerType int32, peerId int64, cMap map[string]interface{}) (lastInsertId, rowsAffected int64, err error) {
	var (
		s1 []string
		s2 []string
		s3 []string
		r  sql.Result
	)

	for k, _ := range cMap {
		s1 = append(s1, k)
		s2 = append(s2, ":"+k)
		s3 = append(s3, fmt.Sprintf("%s = :%s", k, k))
	}

	cMap["user_id"] = userId
	cMap["peer_type"] = peerType
	cMap["peer_id"] = peerId

	query := `
		insert into user_notify_settings
			(user_id, peer_type, peer_id, %s) 
		values 
			(:user_id, :peer_type, :peer_id, %s)
		on duplicate key update %s, deleted = 0`

	ss := fmt.Sprintf(query, strings.Join(s1, ","), strings.Join(s2, ","), strings.Join(s3, ", "))
	// logx.WithContext(ctx).Debugf("sql - %s", ss)
	r, err = dao.db.NamedExec(ctx, ss, cMap)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec (%s) in InsertOrUpdate(%v), error: %v", query, cMap, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId (%s) in InsertOrUpdate(%v)_error: %v", query, cMap, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected (%s) in InsertOrUpdate(%v)_error: %v", query, cMap, err)
	}

	return
}

// InsertOrUpdateExtTx
// insert into user_notify_settings(user_id, peer_type, peer_id, show_previews, silent, mute_until, sound) values (:user_id, :peer_type, :peer_id, :show_previews, :silent, :mute_until, :sound) on duplicate key update show_previews = values(show_previews), silent = values(silent), mute_until = values(mute_until), sound = values(sound), deleted = 0
// TODO(@benqi): sqlmap
func (dao *UserNotifySettingsDAO) InsertOrUpdateExtTx(tx *sqlx.Tx, userId int64, peerType int32, peerId int64, cMap map[string]interface{}) (lastInsertId, rowsAffected int64, err error) {
	var (
		s1 []string
		s2 []string
		s3 []string
		r  sql.Result
	)

	for k, _ := range cMap {
		s1 = append(s1, k)
		s2 = append(s2, ":"+k)
		s3 = append(s3, fmt.Sprintf("%s = (%s)", k, k))
	}

	cMap["user_id"] = userId
	cMap["peer_type"] = peerType
	cMap["peer_id"] = peerId

	query := `
		insert into user_notify_settings
			(user_id, peer_type, peer_id, %s) 
		values 
			(:user_id, :peer_type, :peer_id, %s)
		on duplicate key update %s, deleted = 0`

	r, err = tx.NamedExec(fmt.Sprintf(query, strings.Join(s1, ","), strings.Join(s2, ","), strings.Join(s3, ", ")), cMap)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec (%s) in InsertOrUpdate(%v), error: %v", query, cMap, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId (%s) in InsertOrUpdate(%v)_error: %v", query, cMap, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected (%s) in InsertOrUpdate(%v)_error: %v", query, cMap, err)
	}

	return
}

// SelectListWithCB
// select id, user_id, peer_type, peer_id, show_previews, silent, mute_until, sound from user_notify_settings where user_id = :user_id and deleted = 0
// TODO(@benqi): sqlmap
func (dao *UserNotifySettingsDAO) SelectListWithCB(ctx context.Context, userId int64, peers []*mtproto.PeerUtil, cb func(i int, v *dataobject.UserNotifySettingsDO)) (rList []dataobject.UserNotifySettingsDO, err error) {
	var (
		qVs                                   []string
		args                                  []interface{}
		a                                     []interface{}
		values                                []dataobject.UserNotifySettingsDO
		userIdList, chatIdList, channelIdList []int64
	)

	if len(peers) == 0 {
		logx.WithContext(ctx).Errorf("idList empty")
		return
	}

	for _, peer := range peers {
		switch peer.PeerType {
		case mtproto.PEER_SELF, mtproto.PEER_USER:
			userIdList = append(userIdList, peer.PeerId)
		case mtproto.PEER_CHAT:
			chatIdList = append(chatIdList, peer.PeerId)
		case mtproto.PEER_CHANNEL:
			channelIdList = append(channelIdList, peer.PeerId)
		}
	}

	query := `
	select 
		id, user_id, peer_type, peer_id, show_previews, silent, mute_until, sound 
	from 
		user_notify_settings 
	where 
		user_id = ? AND deleted = 0 AND (%s)`

	args = append(args, userId)
	if len(userIdList) > 0 {
		qVs = append(qVs, "(peer_type = 2 AND peer_id IN (?)) ")
		args = append(args, userIdList)
	}
	if len(chatIdList) > 0 {
		qVs = append(qVs, "(peer_type = 3 AND peer_id IN (?)) ")
		args = append(args, chatIdList)
	}
	if len(channelIdList) > 0 {
		qVs = append(qVs, "(peer_type = 4 AND peer_id IN (?)) ")
		args = append(args, channelIdList)
	}
	query = fmt.Sprintf(query, strings.Join(qVs, " OR "))

	query, a, err = sqlx.In(query, args...)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In (%s) in SelectNotifySettingsList(_), error: %v", query, err)
		return
	}

	err = dao.db.QueryRowsPartial(ctx, &values, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectNotifySettingsList(_), error: %v", err)
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
