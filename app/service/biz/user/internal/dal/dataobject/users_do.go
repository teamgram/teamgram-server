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

type UsersDO struct {
	Id                               int64  `db:"id" json:"id"`
	UserType                         int32  `db:"user_type" json:"user_type"`
	AccessHash                       int64  `db:"access_hash" json:"access_hash"`
	SecretKeyId                      int64  `db:"secret_key_id" json:"secret_key_id"`
	FirstName                        string `db:"first_name" json:"first_name"`
	LastName                         string `db:"last_name" json:"last_name"`
	Username                         string `db:"username" json:"username"`
	Phone                            string `db:"phone" json:"phone"`
	CountryCode                      string `db:"country_code" json:"country_code"`
	Verified                         bool   `db:"verified" json:"verified"`
	Support                          bool   `db:"support" json:"support"`
	Scam                             bool   `db:"scam" json:"scam"`
	Fake                             bool   `db:"fake" json:"fake"`
	Premium                          bool   `db:"premium" json:"premium"`
	About                            string `db:"about" json:"about"`
	State                            int32  `db:"state" json:"state"`
	IsBot                            bool   `db:"is_bot" json:"is_bot"`
	AccountDaysTtl                   int32  `db:"account_days_ttl" json:"account_days_ttl"`
	PhotoId                          int64  `db:"photo_id" json:"photo_id"`
	Restricted                       bool   `db:"restricted" json:"restricted"`
	RestrictionReason                string `db:"restriction_reason" json:"restriction_reason"`
	ArchiveAndMuteNewNoncontactPeers bool   `db:"archive_and_mute_new_noncontact_peers" json:"archive_and_mute_new_noncontact_peers"`
	EmojiStatusDocumentId            int64  `db:"emoji_status_document_id" json:"emoji_status_document_id"`
	EmojiStatusUntil                 int32  `db:"emoji_status_until" json:"emoji_status_until"`
	StoriesMaxId                     int32  `db:"stories_max_id" json:"stories_max_id"`
	Color                            int32  `db:"color" json:"color"`
	ColorBackgroundEmojiId           int64  `db:"color_background_emoji_id" json:"color_background_emoji_id"`
	ProfileColor                     int32  `db:"profile_color" json:"profile_color"`
	ProfileColorBackgroundEmojiId    int64  `db:"profile_color_background_emoji_id" json:"profile_color_background_emoji_id"`
	Birthday                         string `db:"birthday" json:"birthday"`
	Deleted                          bool   `db:"deleted" json:"deleted"`
	DeleteReason                     string `db:"delete_reason" json:"delete_reason"`
}
