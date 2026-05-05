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
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// UserupdatesGetMessageViewsByPeerSeqs
// userupdates.getMessageViewsByPeerSeqs user_id:long peers:Vector<MessageViewPeerSeq> = MessageViewList;
func (c *UserupdatesCore) UserupdatesGetMessageViewsByPeerSeqs(in *userupdates.TLUserupdatesGetMessageViewsByPeerSeqs) (*userupdates.MessageViewList, error) {
	if in == nil || in.UserId == 0 {
		return nil, fmt.Errorf("%w: invalid get message views by peer seqs request", userupdates.ErrOperationTerminal)
	}
	if len(in.Peers) > int(paging.DialogMaxHydratePeersPerRequest) {
		return nil, userupdates.ErrDialogQueryTooLarge
	}

	peers := make([]repository.MessageViewPeerSeq, 0, len(in.Peers))
	for _, peer := range in.Peers {
		if peer == nil {
			return nil, fmt.Errorf("%w: nil message view peer seq", userupdates.ErrOperationTerminal)
		}
		peers = append(peers, repository.MessageViewPeerSeq{PeerType: peer.PeerType, PeerID: peer.PeerId, PeerSeq: peer.PeerSeq})
	}
	views, err := c.svcCtx.Repo.GetMessageViewsByPeerSeqs(c.ctx, in.UserId, peers)
	if err != nil {
		return nil, err
	}

	messages := make([]tg.MessageClazz, 0, len(peers))
	for _, peer := range peers {
		view, ok := views[peer]
		if !ok {
			continue
		}
		message, err := messageViewToTLMessage(view)
		if err != nil {
			return nil, err
		}
		if message != nil {
			messages = append(messages, message)
		}
	}
	return userupdates.MakeTLMessageViewList(&userupdates.TLMessageViewList{Messages: messages}).ToMessageViewList(), nil
}
