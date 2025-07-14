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
	"github.com/teamgram/teamgram-server/v2/app/interface/session/session"
)

var _ *tg.Bool

// SessionCloseSession
// session.closeSession client:SessionClientEvent = Bool;
func (c *SessionCore) SessionCloseSession(in *session.TLSessionCloseSession) (*tg.Bool, error) {
	cli, _ := in.Client.(*session.TLSessionClientEvent)

	if cli == nil {
		err := tg.ErrInputRequestInvalid
		c.Logger.Errorf("session.closeSession - error: %v", err)
		return nil, err
	}

	mainAuth := c.svcCtx.MainAuthMgr.GetMainAuthWrapper(cli.PermAuthKeyId)
	if mainAuth == nil {
		c.Logger.Errorf("session.closeSession - not found sessList by keyId: %s", cli)
	} else {
		_ = mainAuth.SessionClientClosed(c.ctx, int(cli.KeyType), cli.AuthKeyId, cli.ServerId, cli.SessionId)
	}

	return tg.BoolTrue, nil
}
