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
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// AccountSaveMusic
// account.saveMusic#b26732a9 flags:# unsave:flags.0?true id:InputDocument after_id:flags.1?InputDocument = Bool;
func (c *UserChannelProfilesCore) AccountSaveMusic(in *tg.TLAccountSaveMusic) (*tg.Bool, error) {
	selfID, err := requireSelfID(c)
	if err != nil {
		return nil, err
	}
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}
	documentID, err := documentIDFromInputDocument(in.Id)
	if err != nil {
		return nil, err
	}
	if err := requireUserClient(c); err != nil {
		return nil, err
	}

	if _, err = c.svcCtx.Repo.UserClient.UserSaveMusic(c.ctx, &userpb.TLUserSaveMusic{
		Unsave:  in.Unsave,
		UserId:  selfID,
		Id:      documentID,
		AfterId: optionalDocumentID(in.AfterId),
	}); err != nil {
		c.Logger.Errorf("account.saveMusic - save failed: user_id: %d, document_id: %d, err: %v", selfID, documentID, err)
		return tg.BoolFalse, nil
	}

	return tg.BoolTrue, nil
}
