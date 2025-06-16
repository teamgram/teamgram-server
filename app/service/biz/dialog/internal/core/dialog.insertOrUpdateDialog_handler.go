// Copyright (c) 2024 The Teamgram Authors. All rights reserved.
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
	"errors"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
)

var _ *tg.Bool

// DialogInsertOrUpdateDialog
// dialog.insertOrUpdateDialog flags:# user_id:long peer_type:int peer_id:long top_message:flags.0?int read_outbox_max_id:flags.1?int read_inbox_max_id:flags.2?int unread_count:flags.3?int unread_mark:flags.4?true date2:flags.5?long pinned_msg_id:flags.6?int = Bool;
func (c *DialogCore) DialogInsertOrUpdateDialog(in *dialog.TLDialogInsertOrUpdateDialog) (*tg.Bool, error) {
	// TODO: not impl
	// c.Logger.Errorf("dialog.insertOrUpdateDialog blocked, License key from https://teamgram.net required to unlock enterprise features.")

	return nil, errors.New("dialog.insertOrUpdateDialog not implemented")
}
