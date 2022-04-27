/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dataobject

type ChatsDO struct {
	Id                   int64  `db:"id"`
	CreatorUserId        int64  `db:"creator_user_id"`
	AccessHash           int64  `db:"access_hash"`
	RandomId             int64  `db:"random_id"`
	ParticipantCount     int32  `db:"participant_count"`
	Title                string `db:"title"`
	About                string `db:"about"`
	PhotoId              int64  `db:"photo_id"`
	DefaultBannedRights  int64  `db:"default_banned_rights"`
	MigratedToId         int64  `db:"migrated_to_id"`
	MigratedToAccessHash int64  `db:"migrated_to_access_hash"`
	AvailableReactions   string `db:"available_reactions"`
	Deactivated          bool   `db:"deactivated"`
	Noforwards           bool   `db:"noforwards"`
	Version              int32  `db:"version"`
	Date                 int64  `db:"date"`
}
