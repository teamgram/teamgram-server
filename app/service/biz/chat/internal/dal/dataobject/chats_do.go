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

type ChatsDO struct {
	Id                     int64  `db:"id" json:"id"`
	CreatorUserId          int64  `db:"creator_user_id" json:"creator_user_id"`
	AccessHash             int64  `db:"access_hash" json:"access_hash"`
	RandomId               int64  `db:"random_id" json:"random_id"`
	ParticipantCount       int32  `db:"participant_count" json:"participant_count"`
	Title                  string `db:"title" json:"title"`
	About                  string `db:"about" json:"about"`
	PhotoId                int64  `db:"photo_id" json:"photo_id"`
	DefaultBannedRights    int64  `db:"default_banned_rights" json:"default_banned_rights"`
	MigratedToId           int64  `db:"migrated_to_id" json:"migrated_to_id"`
	MigratedToAccessHash   int64  `db:"migrated_to_access_hash" json:"migrated_to_access_hash"`
	AvailableReactionsType int32  `db:"available_reactions_type" json:"available_reactions_type"`
	AvailableReactions     string `db:"available_reactions" json:"available_reactions"`
	Deactivated            bool   `db:"deactivated" json:"deactivated"`
	Noforwards             bool   `db:"noforwards" json:"noforwards"`
	TtlPeriod              int32  `db:"ttl_period" json:"ttl_period"`
	Version                int32  `db:"version" json:"version"`
	Date                   int64  `db:"date" json:"date"`
}
