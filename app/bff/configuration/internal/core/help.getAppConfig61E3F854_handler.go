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

// HelpGetAppConfig61E3F854
// help.getAppConfig#61e3f854 hash:int = help.AppConfig;
func (c *ConfigurationCore) HelpGetAppConfig61E3F854(in *mtproto.TLHelpGetAppConfig61E3F854) (*mtproto.Help_AppConfig, error) {
	// TODO: not impl
	c.Logger.Errorf("help.getAppConfig blocked, License key from https://teamgram.net required to unlock enterprise features.")

	return mtproto.MakeTLHelpAppConfig(&mtproto.Help_AppConfig{
		Hash: 0,
		Config: mtproto.MakeTLJsonObject(&mtproto.JSONValue{
			Value_VECTORJSONOBJECTVALUE: []*mtproto.JSONObjectValue{},
		}).To_JSONValue(),
	}).To_Help_AppConfig(), nil
}
