// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package dao

import (
	"github.com/teamgram/marmota/pkg/container2/sets"

	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/hash"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stringx"
)

type RpcShardingManager struct {
	pubListenOn  string
	c            discov.EtcdConf
	shardingList sets.Set
	dispatcher   *hash.ConsistentHash
	cb           func(sharding *RpcShardingManager, oldList, addList, removeList []string)
}

func NewRpcShardingManager(pubListenOn string, c discov.EtcdConf) *RpcShardingManager {
	sess := &RpcShardingManager{
		c:            c,
		pubListenOn:  pubListenOn,
		shardingList: sets.NewWithLength(16),
		dispatcher:   hash.NewConsistentHash(),
		cb:           nil,
	}

	return sess
}

func (sess *RpcShardingManager) GetShardingV(k string) (string, bool) {
	if v, ok := sess.dispatcher.Get(k); ok {
		return v.(string), true
	} else {
		return "", false
	}
}

func (sess *RpcShardingManager) ShardingVIsListenOn(k string) bool {
	if v, ok := sess.dispatcher.Get(k); ok {
		return v.(string) == sess.pubListenOn
	} else {
		return false
	}
}

func (sess *RpcShardingManager) RegisterCB(cb func(sharding *RpcShardingManager, oldList, addList, removeList []string)) {
	sess.cb = cb
}

func (sess *RpcShardingManager) Start() {
	sub, err := discov.NewSubscriber(sess.c.Hosts, sess.c.Key)
	if err != nil {
		logx.Must(err)
	}

	update := func() {
		var (
			oldList    = sess.shardingList.UnsortedList()
			addList    []string
			removeList []string
		)

		values := sub.Values()
		shardingList := sets.New(values...)
		for _, v := range values {
			shardingList.Insert(v)
			if sess.shardingList.Contains(v) {
				continue
			} else {
				addList = append(addList, v)
			}
		}

		for key, _ := range sess.shardingList {
			if !stringx.Contains(values, key) {
				removeList = append(removeList, key)
			}
		}

		for _, n := range addList {
			sess.dispatcher.Add(n)
		}

		for _, n := range removeList {
			sess.dispatcher.Remove(n)
		}

		sess.shardingList = shardingList

		if len(removeList) > 0 && sess.cb != nil {
			sess.cb(sess, oldList, addList, removeList)
		}
	}

	sub.AddListener(update)
	update()
}
