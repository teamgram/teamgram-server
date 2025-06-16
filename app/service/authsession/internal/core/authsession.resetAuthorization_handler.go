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

// AuthsessionResetAuthorization
// authsession.resetAuthorization user_id:long auth_key_id:long hash:long = Vector<long>;
func (c *AuthsessionCore) AuthsessionResetAuthorization(in *authsession.TLAuthsessionResetAuthorization) (*authsession.VectorLong, error) {
	var (
		excludeKeyId = in.AuthKeyId
	)

	if excludeKeyId != 0 {
		myKeyData, err := c.svcCtx.Dao.QueryAuthKeyV2(c.ctx, in.AuthKeyId)
		if err != nil {
			c.Logger.Errorf("session.getAuthorizations - error: %v", err)
			return nil, err
		} else if myKeyData == nil {
			c.Logger.Errorf("session.getAuthorizations - error: %v", err)
			err = tg.ErrAuthKeyInvalid
			return nil, err
		} else {
			excludeKeyId = myKeyData.PermAuthKeyId
		}
	}

	keyIdList := c.svcCtx.Dao.ResetAuthorization(c.ctx, in.UserId, excludeKeyId, in.Hash)
	// log.Debugf("keyIdList: %v", keyIdList)

	keyIdL2ist := make([]int64, 0, len(keyIdList))
	for _, keyId := range keyIdList {
		keyData, _ := c.svcCtx.Dao.QueryAuthKeyV2(c.ctx, keyId)
		if keyData != nil {
			if keyData.TempAuthKeyId != 0 {
				keyIdL2ist = append(keyIdL2ist, keyData.TempAuthKeyId)
			} else {
				keyIdL2ist = append(keyIdL2ist, keyId)
			}
		}
	}

	return &authsession.VectorLong{
		Datas: keyIdList,
	}, nil
}
