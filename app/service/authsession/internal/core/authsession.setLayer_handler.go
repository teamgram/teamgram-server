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
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
)

var _ *tg.Bool

// AuthsessionSetLayer
// authsession.setLayer auth_key_id:long ip:string layer:int = Bool;
func (c *AuthsessionCore) AuthsessionSetLayer(in *authsession.TLAuthsessionSetLayer) (*tg.Bool, error) {
	var (
		setLayer = in
		inKeyId  = in.AuthKeyId
	)

	keyData, err := c.svcCtx.Dao.QueryAuthKeyV2(c.ctx, inKeyId)
	if err != nil {
		c.Logger.Errorf("setLayer - queryAuthKeyV2(%d) is error: %v", inKeyId, err)
		return nil, err
	} else if keyData.PermAuthKeyId == 0 {
		c.Logger.Errorf("queryAuthKeyV2(%d) - PermAuthKeyId is empty", inKeyId)
		return nil, tg.ErrAuthKeyPermEmpty
	}

	setLayer.AuthKeyId = keyData.PermAuthKeyId
	err = c.svcCtx.Dao.SetLayer(c.ctx, setLayer)
	if err != nil {
		c.Logger.Errorf("setLayer(%d, %d) is error: %v", inKeyId, in.Layer, err)
		return nil, err
	}

	return tg.BoolTrue, nil
}
