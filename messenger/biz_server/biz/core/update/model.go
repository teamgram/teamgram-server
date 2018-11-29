// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
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

// Author: Benqi (wubenqi@gmail.com)

package updates

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dao"
	"github.com/nebula-chat/chatengine/service/idgen/client"
)

/**
  1. peer_user和peer_chat使用user_dialogs存储
  2. channel和super_chat用channel_participants
*/

type updatesDAO struct {
	idgen.SeqIDGen
}

type UpdateModel struct {
	dao *updatesDAO
}

func (m *UpdateModel) InstallModel() {
	var err error
	m.dao.SeqIDGen, _ = idgen.NewSeqIDGen("redis", dao.CACHE)
	if err != nil {
		glog.Fatal("seqidgen init error: ", err)
	}
}

func (m *UpdateModel) RegisterCallback(cb interface{}) {
}

func init() {
	core.RegisterCoreModel(&UpdateModel{dao: &updatesDAO{}})
}
