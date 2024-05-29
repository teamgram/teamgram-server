// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package sess

import (
	"strconv"
	"sync"

	"github.com/teamgram/teamgram-server/app/interface/session/internal/dao"
)

type MainAuthWrapperManager struct {
	mu      sync.Mutex
	authMgr map[int64]*MainAuthWrapper
	*dao.Dao
}

func NewMainAuthWrapperManager(d *dao.Dao) *MainAuthWrapperManager {
	return &MainAuthWrapperManager{
		authMgr: make(map[int64]*MainAuthWrapper),
		Dao:     d,
	}
}

func (m *MainAuthWrapperManager) GetMainAuthWrapper(mainAuthKeyId int64) *MainAuthWrapper {
	m.mu.Lock()
	defer m.mu.Unlock()

	v, ok := m.authMgr[mainAuthKeyId]
	if ok {
		return v
	}

	return nil
}

func (m *MainAuthWrapperManager) DeleteByAuthKeyId(mainAuthKeyId int64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.authMgr, mainAuthKeyId)
}

func (m *MainAuthWrapperManager) AllocMainAuthWrapper(mainAuth *MainAuthWrapper) *MainAuthWrapper {
	m.mu.Lock()
	defer m.mu.Unlock()

	v, ok := m.authMgr[mainAuth.authKeyId]
	if ok {
		return v
	} else {
		m.authMgr[mainAuth.authKeyId] = mainAuth
		return mainAuth
	}
}

func (m *MainAuthWrapperManager) OnShardingCB(sharding *dao.RpcShardingManager, oldList, addList []string, removeList []string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for k, v := range m.authMgr {
		if !sharding.ShardingVIsListenOn(strconv.FormatInt(k, 10)) {
			delete(m.authMgr, k)
			v.Stop()
		}
	}
}
