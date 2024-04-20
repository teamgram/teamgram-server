/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Studio (https://teamgram.io).
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
	var (
		archiveAndMuteNewNoncontactPeers bool
	)

	if in.GetSettings().GetArchiveAndMuteNewNoncontactPeers_FLAGBOOL() != nil {
		archiveAndMuteNewNoncontactPeers = mtproto.FromBool(in.GetSettings().GetArchiveAndMuteNewNoncontactPeers_FLAGBOOL())
	} else {
		archiveAndMuteNewNoncontactPeers = in.GetSettings().GetArchiveAndMuteNewNoncontactPeers_FLAGBOOLEAN()
	}

	c.svcCtx.Dao.UserGlobalPrivacySettingsDAO.InsertOrUpdate(c.ctx, &dataobject.UserGlobalPrivacySettingsDO{
		UserId:                           in.UserId,
		ArchiveAndMuteNewNoncontactPeers: archiveAndMuteNewNoncontactPeers,
		KeepArchivedUnmuted:              in.GetSettings().GetKeepArchivedUnmuted(),
		KeepArchivedFolders:              in.GetSettings().GetKeepArchivedFolders(),
		HideReadMarks:                    in.GetSettings().GetHideReadMarks(),
		NewNoncontactPeersRequirePremium: in.GetSettings().GetNewNoncontactPeersRequirePremium(),
	})

	return mtproto.BoolTrue, nil
}
