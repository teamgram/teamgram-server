// Copyright (c) 2026 The Teamgram Authors (https://teamgram.net).
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

package core

import (
	"github.com/teamgram/proto/mtproto"
)

// AccountToggleWebBrowserSettingsException2D0A0571
// account.toggleWebBrowserSettingsException#2d0a0571 flags:# delete:flags.1?true open_external_browser:flags.0?Bool url:string = Bool;
func (c *WebBrowserCore) AccountToggleWebBrowserSettingsException2D0A0571(in *mtproto.TLAccountToggleWebBrowserSettingsException2D0A0571) (*mtproto.Bool, error) {
	// TODO: not impl
	c.Logger.Errorf("account.toggleWebBrowserSettingsException2D0A0571 blocked, License key from https://teamgram.net required to unlock enterprise features.")

	return nil, mtproto.ErrEnterpriseIsBlocked
}
