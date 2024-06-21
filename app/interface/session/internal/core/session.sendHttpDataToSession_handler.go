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
	"github.com/teamgram/teamgram-server/app/interface/session/session"
)

// SessionSendHttpDataToSession
// session.sendHttpDataToSession client:SessionClientData = HttpSessionData;
func (c *SessionCore) SessionSendHttpDataToSession(in *session.TLSessionSendHttpDataToSession) (*session.HttpSessionData, error) {
	//var (
	//	data = in.GetClient()
	//)
	//
	//if data == nil {
	//	err := mtproto.ErrInputRequestInvalid
	//	c.Logger.Errorf("session.sendHttpDataToSession - error: %v", err)
	//	return nil, err
	//}
	//
	//mainAuth, err := c.getOrFetchMainAuthWrapper(data.PermAuthKeyId)
	//if err != nil {
	//	c.Logger.Errorf("session.sendHttpDataToSession - error: %v", err)
	//	return nil, err
	//}
	//
	//chData := make(chan interface{})
	//mainAuth.SessionHttpDataArrived(
	//	c.ctx,
	//	int(data.KeyType),
	//	data.AuthKeyId,
	//	data.ServerId,
	//	data.ClientIp,
	//	data.SessionId,
	//	data.Salt,
	//	data.Payload,
	//	chData)
	//
	//timer := time.NewTimer(time.Second * 7)
	//select {
	//case cData := <-chData:
	//	return &session.HttpSessionData{
	//		Payload: cData.([]byte),
	//	}, nil
	//case <-timer.C:
	//	c.Logger.Errorf("chData timeout...")
	//}

	c.Logger.Errorf("session.sendHttpDataToSession - error: not implement")

	return &session.HttpSessionData{
		Payload: []byte{},
	}, nil
}
