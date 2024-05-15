// Copyright 2024 Teamgram Authors
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
	"github.com/teamgram/teamgram-server/app/interface/session/session"
	"github.com/teamgram/teamgram-server/app/service/authsession/authsession"
)

// SessionSetAuthKey
// session.setAuthKey auth_key:AuthKeyInfo future_salt:FutureSalt expires_in:int = Bool;
func (c *SessionCore) SessionSetAuthKey(in *session.TLSessionSetAuthKey) (*mtproto.Bool, error) {
	if in.AuthKey == nil {
		c.Logger.Errorf("session.setAuthKey error: auth_key is nil")
		return nil, mtproto.ErrInputRequestInvalid
	}

	rV, err := c.svcCtx.Dao.AuthsessionClient.AuthsessionSetAuthKey(
		c.ctx,
		&authsession.TLAuthsessionSetAuthKey{
			AuthKey:    in.AuthKey,
			FutureSalt: in.GetFutureSalt(),
			ExpiresIn:  in.GetExpiresIn(),
		})
	if err != nil {
		c.Logger.Errorf("session.setAuthKey - error: %v", err)
		return nil, err
	}

	return rV, nil
}
