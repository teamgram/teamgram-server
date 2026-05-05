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
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository"
)

// DialogBatchGetDialogExtras
// dialog.batchGetDialogExtras user_id:long peers:Vector<DialogPeer> = Vector<DialogExtras>;
func (c *DialogCore) DialogBatchGetDialogExtras(in *dialog.TLDialogBatchGetDialogExtras) (*dialog.VectorDialogExtras, error) {
	peers := make([]repository.PeerRef, 0, len(in.Peers))
	for _, peer := range in.Peers {
		peers = append(peers, repository.PeerRef{PeerType: peer.PeerType, PeerID: peer.PeerId})
	}
	extras, err := c.svcCtx.Repo.BatchGetDialogExtras(c.ctx, in.UserId, peers)
	if err != nil {
		return nil, err
	}
	return makeDialogExtrasVector(extras), nil
}
