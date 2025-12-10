/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dataobject

type UserSavedMusicDO struct {
	Id           int32 `db:"id" json:"id"`
	UserId       int64 `db:"user_id" json:"user_id"`
	SavedMusicId int64 `db:"saved_music_id" json:"saved_music_id"`
	Order2       int64 `db:"order2" json:"order2"`
	Deleted      bool  `db:"deleted" json:"deleted"`
}
