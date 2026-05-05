// Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package core

import (
	"errors"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository"
)

// DialogGetDialogByPeerV2
// dialog.getDialogByPeerV2 user_id:long peer:DialogPeer = DialogExtV2;
func (c *DialogCore) DialogGetDialogByPeerV2(in *dialog.TLDialogGetDialogByPeerV2) (*dialog.DialogExtV2, error) {
	if c.svcCtx.Userupdates == nil {
		return nil, dialog.WrapDialogStorage("userupdates.getDialogsByPeers", errors.New("userupdates client is not configured"))
	}
	records, err := c.svcCtx.Userupdates.UserupdatesGetDialogsByPeers(c.ctx, &userupdates.TLUserupdatesGetDialogsByPeers{
		UserId: in.UserId,
		Peers: []userupdates.DialogProjectionPeerClazz{
			userupdates.MakeTLDialogProjectionPeer(&userupdates.TLDialogProjectionPeer{
				PeerType: in.Peer.PeerType,
				PeerId:   in.Peer.PeerId,
			}),
		},
	})
	if err != nil {
		return nil, dialog.WrapDialogStorage("userupdates.getDialogsByPeers", err)
	}
	if len(records.Datas) == 0 {
		return nil, dialog.ErrDialogNotFound
	}
	extras, err := c.svcCtx.Repo.BatchGetDialogExtras(c.ctx, in.UserId, []repository.PeerRef{{PeerType: in.Peer.PeerType, PeerID: in.Peer.PeerId}})
	if err != nil {
		return nil, err
	}
	var extra *dialog.DialogExtras
	if len(extras) > 0 {
		extra = makeDialogExtras(extras[0])
	}
	return makeDialogExtV2FromProjection(records.Datas[0], extra), nil
}
