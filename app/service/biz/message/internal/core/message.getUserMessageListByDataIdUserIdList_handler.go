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
	"github.com/teamgram/teamgram-server/app/service/biz/message/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/message/message"
)

// MessageGetUserMessageListByDataIdUserIdList
// message.getUserMessageListByDataIdUserIdList id:long user_id_list:Vector<long> = Vector<MessageBox>;
func (c *MessageCore) MessageGetUserMessageListByDataIdUserIdList(in *message.TLMessageGetUserMessageListByDataIdUserIdList) (*message.Vector_MessageBox, error) {
	var (
		tables     = make(map[string][]int64)
		rValueList = &message.Vector_MessageBox{
			Datas: make([]*mtproto.MessageBox, 0),
		}
	)

	for _, uid := range in.GetUserIdList() {
		table := c.svcCtx.Dao.MessagesDAO.CalcTableName(uid)
		if idList, ok := tables[table]; ok {
			tables[table] = append(idList, uid)
		} else {
			tables[table] = []int64{uid}
		}
	}

	for k, v := range tables {
		c.svcCtx.Dao.MessagesDAO.SelectByMessageDataIdUserIdListWithCB(
			c.ctx,
			k,
			in.GetId(),
			v,
			func(sz, i int, v *dataobject.MessagesDO) {
				rValueList.Datas = append(rValueList.Datas, c.svcCtx.Dao.MakeMessageBox(c.ctx, 0, v))
			})
	}

	return rValueList, nil
}
