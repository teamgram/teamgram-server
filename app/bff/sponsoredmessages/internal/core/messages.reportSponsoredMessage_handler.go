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
	"github.com/teamgram/proto/mtproto"
)

// MessagesReportSponsoredMessage
// messages.reportSponsoredMessage#1af3dbb8 peer:InputPeer random_id:bytes option:bytes = channels.SponsoredMessageReportResult;
func (c *SponsoredMessagesCore) MessagesReportSponsoredMessage(in *mtproto.TLMessagesReportSponsoredMessage) (*mtproto.Channels_SponsoredMessageReportResult, error) {
	// TODO: not impl
	c.Logger.Errorf("messages.reportSponsoredMessage blocked, License key from https://teamgram.net required to unlock enterprise features.")

	return nil, mtproto.ErrEnterpriseIsBlocked
}
