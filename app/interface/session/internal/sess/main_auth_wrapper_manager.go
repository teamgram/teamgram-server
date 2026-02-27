// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package sess

import (
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/teamgram/teamgram-server/app/interface/session/internal/dao"
	"github.com/teamgram/teamgram-server/app/service/status/status"

	"github.com/zeromicro/go-zero/core/logx"
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

func (m *MainAuthWrapperManager) AllocMainAuthWrapper(authKeyId int64, newMainAuth func(authKeyId int64) *MainAuthWrapper) *MainAuthWrapper {
	m.mu.Lock()
	defer m.mu.Unlock()

	v, ok := m.authMgr[authKeyId]
	if ok {
		return v
	} else {
		mainAuth := newMainAuth(authKeyId)
		m.authMgr[authKeyId] = mainAuth
		return mainAuth
	}
}

func (m *MainAuthWrapperManager) OnShardingCB(sharding *dao.RpcShardingManager, oldList, addList []string, removeList []string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for k, v := range m.authMgr {
		if !sharding.ShardingVIsListenOn(strconv.FormatInt(k, 10)) {
			// 主动清理 status 中的过期 Gateway，使 sync 推送路径尽快恢复
			if v.AuthUserId > 0 {
				_, err := m.Dao.StatusClient.StatusSetSessionOffline(
					context.Background(),
					&status.TLStatusSetSessionOffline{
						UserId:    v.AuthUserId,
						AuthKeyId: k,
					})
				if err != nil {
					logx.Errorf("OnShardingCB - StatusSetSessionOffline(userId:%d, authKeyId:%d) error: %v",
						v.AuthUserId, k, err)
				}
			}
			delete(m.authMgr, k)
			v.Stop()
		}
	}
}

// Drain 优雅排空所有会话。等待所有 MainAuthWrapper 的 runLoop 退出，
// 或超时后强制停止。用于 session 服务优雅退出。
func (m *MainAuthWrapperManager) Drain(timeout time.Duration) {
	m.mu.Lock()
	wrappers := make([]*MainAuthWrapper, 0, len(m.authMgr))
	for k, v := range m.authMgr {
		// 清理 status 中的 Gateway
		if v.AuthUserId > 0 {
			_, _ = m.Dao.StatusClient.StatusSetSessionOffline(
				context.Background(),
				&status.TLStatusSetSessionOffline{
					UserId:    v.AuthUserId,
					AuthKeyId: k,
				})
		}
		wrappers = append(wrappers, v)
	}
	m.authMgr = make(map[int64]*MainAuthWrapper)
	m.mu.Unlock()

	logx.Infof("Drain: stopping %d auth wrappers, timeout %v", len(wrappers), timeout)

	// 通知所有 wrapper 停止
	for _, w := range wrappers {
		w.Stop()
	}

	// 等待所有 wrapper 的 runLoop 退出（通过 finish.Wait）
	done := make(chan struct{})
	go func() {
		for _, w := range wrappers {
			w.finish.Wait()
		}
		close(done)
	}()

	select {
	case <-done:
		logx.Infof("Drain: all %d auth wrappers stopped gracefully", len(wrappers))
	case <-time.After(timeout):
		logx.Errorf("Drain: timeout after %v, %d auth wrappers may not have stopped", timeout, len(wrappers))
	}
}
