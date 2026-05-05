// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
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
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// DialogSetChatTheme
// dialog.setChatTheme user_id:long peer_type:int peer_id:long theme_emoticon:string = Bool;
func (c *DialogCore) DialogSetChatTheme(in *dialog.TLDialogSetChatTheme) (*tg.Bool, error) {
	if in.PeerType != repository.PeerTypeUser {
		return nil, dialog.ErrWrongOwner
	}
	sourceAuth, err := c.sourcePermAuthKeyID()
	if err != nil {
		return nil, err
	}
	operationID := deterministicOperationID("set_private_theme", in.UserId, in.PeerId, in.ThemeEmoticon)
	if _, err := c.svcCtx.Repo.SetPrivatePeerPolicy(c.ctx, repository.PrivatePeerPolicyInput{
		UserID:              in.UserId,
		PeerUserID:          in.PeerId,
		ThemeEmoticon:       in.ThemeEmoticon,
		SourcePermAuthKeyID: sourceAuth,
		OperationID:         operationID,
		ActorOutboxID:       deterministicOutboxID(operationID, "actor"),
		PeerOutboxID:        deterministicOutboxID(operationID, "peer"),
		DeliveryPath:        repository.DeliveryPathUserupdatesPTS,
		PublicUpdateType:    "messageActionSetChatTheme",
		Payload:             []byte(`{"schema_version":1}`),
	}); err != nil {
		return nil, err
	}
	return tg.BoolTrue, nil
}
