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
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// AccountResetAuthorization
// account.resetAuthorization#df77f3bc hash:long = Bool;
func (c *AccountCore) AccountResetAuthorization(in *tg.TLAccountResetAuthorization) (*tg.Bool, error) {
	selfID, err := requireSelfID(c)
	if err != nil {
		return nil, err
	}
	if in == nil || in.Hash == 0 {
		return tg.BoolFalse, nil
	}
	if err := requireAuthsessionClient(c); err != nil {
		return nil, err
	}

	if _, err = c.svcCtx.Repo.AuthsessionClient.AuthsessionResetAuthorization(c.ctx, &authsession.TLAuthsessionResetAuthorization{
		UserId:    selfID,
		AuthKeyId: accountAuthKeyID(c),
		Hash:      in.Hash,
	}); err != nil {
		return nil, err
	}

	// TODO(v2 account): master sent updatesTooLong and updateAccountResetAuthorization through sync; do not migrate sync calls until v2 userupdates/gateway delivery is defined.
	return tg.BoolTrue, nil
}
