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

type UserGlobalPrivacySettingsDO struct {
	Id                               int64 `db:"id" json:"id"`
	UserId                           int64 `db:"user_id" json:"user_id"`
	ArchiveAndMuteNewNoncontactPeers bool  `db:"archive_and_mute_new_noncontact_peers" json:"archive_and_mute_new_noncontact_peers"`
	KeepArchivedUnmuted              bool  `db:"keep_archived_unmuted" json:"keep_archived_unmuted"`
	KeepArchivedFolders              bool  `db:"keep_archived_folders" json:"keep_archived_folders"`
	HideReadMarks                    bool  `db:"hide_read_marks" json:"hide_read_marks"`
	NewNoncontactPeersRequirePremium bool  `db:"new_noncontact_peers_require_premium" json:"new_noncontact_peers_require_premium"`
}
