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

package dao

import (
	"context"
	"fmt"

	"github.com/teamgram/proto/mtproto"

	"github.com/gogo/protobuf/proto"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	duplicateMessageId   = "duplicate_message_id"
	duplicateMessageData = "duplicate_message_data"
	expireTimeout        = 60 // 60s
)

func makeDuplicateMessageKey(prefix string, senderUserId, clientRandomId int64) string {
	return fmt.Sprintf("%s_%d_%d", prefix, senderUserId, clientRandomId)
}

func (d *Dao) HasDuplicateMessage(ctx context.Context, senderUserId, clientRandomId int64) (bool, error) {
	//conn := d.redis.Redis.Get(ctx)
	//defer conn.Close()

	k := makeDuplicateMessageKey(duplicateMessageId, senderUserId, clientRandomId)

	seq, err := d.KV.Incr(k)
	if err != nil {
		logx.WithContext(ctx).Errorf("checkDuplicateMessage - INCR {%s}, error: {%v}", k, err)
		return false, err
	}

	if err = d.KV.Expire(k, expireTimeout); err != nil {
		logx.WithContext(ctx).Errorf("expire DuplicateMessage - EXPIRE {%s, %d}, error: %s", k, expireTimeout, err)
		return false, err
	}

	return seq > 1, nil
}

func (d *Dao) PutDuplicateMessage(ctx context.Context, senderUserId, clientRandomId int64, upd *mtproto.Updates) error {
	k := makeDuplicateMessageKey(duplicateMessageData, senderUserId, clientRandomId)
	cacheData, _ := proto.Marshal(upd)

	if err := d.KV.Setex(k, string(cacheData), expireTimeout); err != nil {
		logx.WithContext(ctx).Errorf("putDuplicateMessage - SET {%s, %s, %d}, error: %s", k, cacheData, expireTimeout, err)
		return err
	}

	return nil
}

func (d *Dao) GetDuplicateMessage(ctx context.Context, senderUserId, clientRandomId int64) (*mtproto.Updates, error) {
	k := makeDuplicateMessageKey(duplicateMessageData, senderUserId, clientRandomId)

	if cacheData, err := d.KV.Get(k); err != nil {
		if err.Error() == "redigo: nil returned" {
			return nil, nil
		}

		logx.WithContext(ctx).Errorf("getDuplicateMessage - GET {%s}, error: %s", k, err)
		return nil, err
	} else {
		upd := &mtproto.Updates{}
		proto.Unmarshal([]byte(cacheData), upd)

		return upd, nil
	}
}
