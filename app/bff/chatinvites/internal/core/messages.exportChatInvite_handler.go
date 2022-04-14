// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"github.com/teamgram/proto/mtproto"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
)

// MessagesExportChatInvite
// messages.exportChatInvite#a02ce5d5 flags:# legacy_revoke_permanent:flags.2?true request_needed:flags.3?true peer:InputPeer expire_date:flags.0?int usage_limit:flags.1?int title:flags.4?string = ExportedChatInvite;
func (c *ChatInvitesCore) MessagesExportChatInvite(in *mtproto.TLMessagesExportChatInvite) (*mtproto.ExportedChatInvite, error) {
	var (
		peer               = mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
		err                error
		exportedChatInvite *mtproto.ExportedChatInvite
	)

	if !peer.IsChat() {
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("messages.exportChatInvite - error: ", err)
		return nil, err
	}

	exportedChatInvite, err = c.svcCtx.Dao.ChatClient.ChatExportChatInvite(c.ctx, &chatpb.TLChatExportChatInvite{
		ChatId:                peer.PeerId,
		AdminId:               c.MD.UserId,
		LegacyRevokePermanent: in.LegacyRevokePermanent,
		RequestNeeded:         in.RequestNeeded,
		ExpireDate:            in.ExpireDate,
		UsageLimit:            in.UsageLimit,
		Title:                 in.Title,
	})
	if err != nil {
		c.Logger.Errorf("messages.exportChatInvite - error: ", err)
		return nil, err
	}

	// TODO(FIXME): Check logic
	return exportedChatInvite, nil
}
