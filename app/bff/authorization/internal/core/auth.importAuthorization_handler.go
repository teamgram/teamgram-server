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
	"github.com/teamgram/proto/mtproto"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

// AuthImportAuthorization
// auth.importAuthorization#a57a7dad id:long bytes:bytes = auth.Authorization;
func (c *AuthorizationCore) AuthImportAuthorization(in *mtproto.TLAuthImportAuthorization) (*mtproto.Auth_Authorization, error) {
	// TODO: make tmp_session ????
	rValue := mtproto.MakeTLAuthAuthorization(&mtproto.Auth_Authorization{
		SetupPasswordRequired: false,
		OtherwiseReloginDays:  nil,
		TmpSessions:           &wrapperspb.Int32Value{Value: int32(in.GetId())},
		FutureAuthToken:       nil,
		User:                  mtproto.MakeTLUserEmpty(nil).To_User(),
	}).To_Auth_Authorization()

	return rValue, nil
}
