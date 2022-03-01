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
	"time"

	"github.com/teamgram/proto/mtproto"
)

// HelpGetPromoData
// help.getPromoData#c0977421 = help.PromoData;
func (c *PromoDataCore) HelpGetPromoData(in *mtproto.TLHelpGetPromoData) (*mtproto.Help_PromoData, error) {
	// TODO: return help.getPromoData

	rValue := mtproto.MakeTLHelpPromoDataEmpty(&mtproto.Help_PromoData{
		Expires: int32(time.Now().Unix() + 60*60),
	}).To_Help_PromoData()

	return rValue, nil
}
