// Copyright (c) 2024 The Teamgooo Authors. All rights reserved.
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
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

var _ *tg.Bool

// DialogGetDialogById
// dialog.getDialogById user_id:long peer_type:int peer_id:long = DialogExt;
func (c *DialogCore) DialogGetDialogById(in *dialog.TLDialogGetDialogById) (*dialog.DialogExt, error) {
	peerType := int64(tg.PEER_USER)
	peerID := int64(0)
	if in != nil {
		peerType = int64(in.PeerType)
		peerID = in.PeerId
		if peerID == 0 {
			peerID = in.UserId
		}
	}

	return dialog.MakeTLDialogExt(&dialog.TLDialogExt{
		Order:          10,
		Dialog:         makeDialogPlaceholder(peerType, peerID, 10),
		AvailableMinId: 1,
		Date:           10,
	}).ToDialogExt(), nil
}
