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
)

/**
# auth.importLoginToken

Login using a redirected login token, generated in case of DC mismatch during QR code login.

For more info, see login via QR code.

**/

// AuthImportLoginToken
// auth.importLoginToken#95ac5ce4 token:bytes = auth.LoginToken;
func (c *QrCodeCore) AuthImportLoginToken(in *mtproto.TLAuthImportLoginToken) (*mtproto.Auth_LoginToken, error) {
	// TODO: not impl
	// teamgram does not implement multi-datacenter support, so this method is not implemented.

	c.Logger.Errorf("auth.importLoginToken - method not impl")

	return nil, mtproto.ErrAuthTokenAlreadyAccepted
}
