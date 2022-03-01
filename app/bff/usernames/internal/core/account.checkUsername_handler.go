// Copyright 2022 Teamgram Authors
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
	"github.com/teamgram/marmota/pkg/strings2"
	"github.com/teamgram/marmota/pkg/utils"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/username/username"
)

// AccountCheckUsername
// account.checkUsername#2714d86c username:string = Bool;
func (c *UsernamesCore) AccountCheckUsername(in *mtproto.TLAccountCheckUsername) (*mtproto.Bool, error) {
	// Check username format
	// You can choose a username on Telegram.
	// If you do, other people will be able to find
	// you by this username and contact you
	// without knowing your phone number.
	//
	// You can use a-z, 0-9 and underscores.
	// Minimum length is 5 characters.";
	//
	if len(in.Username) < username.MinUsernameLen ||
		!strings2.IsAlNumString(in.Username) ||
		utils.IsNumber(in.Username[0]) {
		err := mtproto.ErrUsernameInvalid
		c.Logger.Errorf("account.checkUsername#2714d86c - format error: %v", err)
		return nil, err
	} else {
		existed, err := c.svcCtx.Dao.UsernameClient.UsernameCheckAccountUsername(c.ctx, &username.TLUsernameCheckAccountUsername{
			UserId:   c.MD.UserId,
			Username: in.GetUsername(),
		})
		if err != nil {
			return nil, err
		}

		switch existed.GetPredicateName() {
		case username.Predicate_usernameExistedNotMe:
			err = mtproto.ErrUsernameOccupied
			c.Logger.Errorf("account.checkUsername#2714d86c - exists username: %v", err)
			return mtproto.BoolFalse, nil
		default:
			break
		}
	}

	return mtproto.BoolTrue, nil
}
