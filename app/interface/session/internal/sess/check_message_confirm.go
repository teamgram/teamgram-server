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

package sess

import (
	"github.com/teamgram/proto/mtproto"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*
bool SecureRequest::needAck() const {
	if (_data->size() <= kMessageBodyPosition) {
		return false;
	}
	const auto type = mtpTypeId((*_data)[kMessageBodyPosition]);
	switch (type) {
	case mtpc_msg_container:
	case mtpc_msgs_ack:
	case mtpc_http_wait:
	case mtpc_bad_msg_notification:
	case mtpc_msgs_all_info:
	case mtpc_msgs_state_info:
	case mtpc_msg_detailed_info:
	case mtpc_msg_new_detailed_info:
		return false;
	}
	return true;
}
*/

func checkMessageConfirm(msg mtproto.TLObject) bool {
	switch msg.(type) {
	case *mtproto.TLMsgContainer,
		*mtproto.TLMsgsAck,
		*mtproto.TLHttpWait,
		*mtproto.TLBadMsgNotification,
		*mtproto.TLMsgsAllInfo,
		*mtproto.TLMsgsStateInfo,
		*mtproto.TLMsgDetailedInfo,
		*mtproto.TLMsgNewDetailedInfo:

		return false
	default:
		return true
	}
}
