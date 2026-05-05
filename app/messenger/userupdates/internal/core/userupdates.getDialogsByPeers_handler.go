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
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/pkg/paging"
)

// UserupdatesGetDialogsByPeers
// userupdates.getDialogsByPeers user_id:long peers:Vector<DialogProjectionPeer> = Vector<DialogProjection>;
func (c *UserupdatesCore) UserupdatesGetDialogsByPeers(in *userupdates.TLUserupdatesGetDialogsByPeers) (*userupdates.VectorDialogProjection, error) {
	if in == nil || in.UserId == 0 {
		return nil, fmt.Errorf("%w: invalid get dialogs by peers request", userupdates.ErrOperationTerminal)
	}
	if len(in.Peers) > int(paging.DialogMaxHydratePeersPerRequest) {
		return nil, userupdates.ErrDialogQueryTooLarge
	}
	peers := make([]repository.DialogProjectionPeer, 0, len(in.Peers))
	for _, peer := range in.Peers {
		if peer == nil {
			return nil, fmt.Errorf("%w: nil dialog projection peer", userupdates.ErrOperationTerminal)
		}
		peers = append(peers, repository.DialogProjectionPeer{PeerType: peer.PeerType, PeerID: peer.PeerId})
	}
	projections, err := c.svcCtx.Repo.GetDialogProjectionsByPeers(c.ctx, in.UserId, peers)
	if err != nil {
		return nil, err
	}

	out := &userupdates.VectorDialogProjection{Datas: make([]userupdates.DialogProjectionClazz, 0, len(peers))}
	for _, peer := range peers {
		projection, ok := projections[peer]
		if !ok {
			continue
		}
		out.Datas = append(out.Datas, dialogProjectionToTL(projection))
	}
	return out, nil
}
