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

package message

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/golang/protobuf/proto"
)

const (
	duplicateMessageId      = "duplicate_message_id"
	duplicateMessageData    = "duplicate_message_data"
	expireTimeout           = 60 // 60s
)

func makeDuplicateMessageKey(prefix string, senderUserId int32, clientRandomId int64) string {
	return fmt.Sprintf("%s_%d_%d", prefix, senderUserId, clientRandomId)
}

func (m *MessageModel) HasDuplicateMessage(senderUserId int32, clientRandomId int64) (bool, error) {
	conn := m.RedisPool.Get()
	defer conn.Close()

	k := makeDuplicateMessageKey(duplicateMessageId, senderUserId, clientRandomId)
	seq, err := redis.Int64(conn.Do("INCR", k))
	if err != nil {
		glog.Errorf("checkDuplicateMessage - INCR {%s}, error: {%v}", k, err)
		return false, err
	}

	if _, err = conn.Do("EXPIRE", k, expireTimeout); err != nil {
		glog.Errorf("expire DuplicateMessage - EXPIRE {%s, %d}, error: %s", k, expireTimeout, err)
		return false, err
	}

	return seq > 1, nil
}

func (m *MessageModel) PutDuplicateMessage(senderUserId int32, clientRandomId int64, upd *mtproto.Updates) error {
	k := makeDuplicateMessageKey(duplicateMessageData, senderUserId, clientRandomId)
	cacheData, _ := proto.Marshal(upd)

	conn := m.RedisPool.Get()
	defer conn.Close()
	if _, err := conn.Do("SET", k, cacheData, "EX", expireTimeout); err != nil {
		glog.Errorf("putDuplicateMessage - SET {%s, %s, %d}, error: %s", k, cacheData, expireTimeout, err)
		return err
	}
	//
	//if _, err := conn.Do("EXPIRE", k, expireTimeout); err != nil {
	//	glog.Errorf("expire putDuplicateMessage - EXPIRE {%s, %d}, error: %s", k, expireTimeout, err)
	//	return err
	//}

	return nil
}

func (m *MessageModel) GetDuplicateMessage(senderUserId int32, clientRandomId int64) (*mtproto.Updates, error) {
	k := makeDuplicateMessageKey(duplicateMessageData, senderUserId, clientRandomId)
	// cacheData, _ := proto.Marshal(upd)

	conn := m.RedisPool.Get()
	var upd *mtproto.Updates

	defer conn.Close()
	if cacheData, err := redis.Bytes(conn.Do("GET", k)); err != nil {
		glog.Errorf("getDuplicateMessage - GET {%s}, error: %s", k, err)
		return nil, err
	} else {
		upd = &mtproto.Updates{}
		proto.Unmarshal(cacheData, upd)
	}

	return upd, nil
}
