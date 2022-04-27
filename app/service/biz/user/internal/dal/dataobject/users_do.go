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

type UsersDO struct {
	Id                               int64  `db:"id"`
	UserType                         int32  `db:"user_type"`
	AccessHash                       int64  `db:"access_hash"`
	SecretKeyId                      int64  `db:"secret_key_id"`
	FirstName                        string `db:"first_name"`
	LastName                         string `db:"last_name"`
	Username                         string `db:"username"`
	Phone                            string `db:"phone"`
	CountryCode                      string `db:"country_code"`
	Verified                         bool   `db:"verified"`
	Support                          bool   `db:"support"`
	Scam                             bool   `db:"scam"`
	Fake                             bool   `db:"fake"`
	About                            string `db:"about"`
	State                            int32  `db:"state"`
	IsBot                            bool   `db:"is_bot"`
	AccountDaysTtl                   int32  `db:"account_days_ttl"`
	PhotoId                          int64  `db:"photo_id"`
	Restricted                       bool   `db:"restricted"`
	RestrictionReason                string `db:"restriction_reason"`
	ArchiveAndMuteNewNoncontactPeers bool   `db:"archive_and_mute_new_noncontact_peers"`
	Deleted                          bool   `db:"deleted"`
	DeleteReason                     string `db:"delete_reason"`
}
