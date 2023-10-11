/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserSetGlobalPrivacySettings
// user.setGlobalPrivacySettings user_id:int settings:GlobalPrivacySettings = Bool;
func (c *UserCore) UserSetGlobalPrivacySettings(in *user.TLUserSetGlobalPrivacySettings) (*mtproto.Bool, error) {
	var archiveAndMuteNewNoncontactPeers bool
	if in.GetSettings().GetArchiveAndMuteNewNoncontactPeers_FLAGBOOL() != nil {
		archiveAndMuteNewNoncontactPeers = mtproto.FromBool(in.GetSettings().GetArchiveAndMuteNewNoncontactPeers_FLAGBOOL())
	} else {
		archiveAndMuteNewNoncontactPeers = in.GetSettings().GetArchiveAndMuteNewNoncontactPeers_FLAGBOOLEAN()
	}

	// TODO: globalPrivacySettings#734c4ccb flags:#
	// 	archive_and_mute_new_noncontact_peers:flags.0?true
	//	keep_archived_unmuted:flags.1?true
	//	keep_archived_folders:flags.2?true = GlobalPrivacySettings;

	c.svcCtx.Dao.UserGlobalPrivacySettingsDAO.InsertOrUpdate(c.ctx, &dataobject.UserGlobalPrivacySettingsDO{
		UserId:                           in.UserId,
		ArchiveAndMuteNewNoncontactPeers: archiveAndMuteNewNoncontactPeers,
	})

	return mtproto.BoolTrue, nil
}
