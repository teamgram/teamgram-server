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

package net2

import (
	"errors"
	"fmt"
	// "github.com/golang/glog"
	"math/rand"
	"sync"
)

type TcpClientGroupManager struct {
	protoName     string
	clientMapLock sync.RWMutex
	clientMap     map[string]map[string]*TcpClient
	callback      TcpClientCallBack
}

func NewTcpClientGroupManager(protoName string, clients map[string][]string, cb TcpClientCallBack) *TcpClientGroupManager {
	group := &TcpClientGroupManager{
		protoName: protoName,
		clientMap: make(map[string]map[string]*TcpClient),
		callback:  cb,
	}

	for k, v := range clients {
		m := make(map[string]*TcpClient)

		for _, address := range v {
			client := NewTcpClient(k, 10*1024, group.protoName, address, group.callback)
			if client != nil {
				m[address] = client
			}
		}

		group.clientMapLock.Lock()
		group.clientMap[k] = m
		group.clientMapLock.Unlock()
	}

	// glog.Info("NewTcpClientGroup group : ", group.clientMap)
	return group
}

func (cgm *TcpClientGroupManager) Serve() bool {
	cgm.clientMapLock.Lock()
	defer cgm.clientMapLock.Unlock()

	for _, v := range cgm.clientMap {
		for _, c := range v {
			c.Serve()
		}
	}

	return true
}

func (cgm *TcpClientGroupManager) Stop() bool {
	cgm.clientMapLock.Lock()
	defer cgm.clientMapLock.Unlock()

	for _, v := range cgm.clientMap {
		for _, c := range v {
			c.Stop()
		}
	}

	return true
}

func (cgm *TcpClientGroupManager) GetConfig() interface{} {
	return nil
}

func (cgm *TcpClientGroupManager) AddClient(name string, address string) {
	// glog.Info("TcpClientGroup AddClient name ", name, " address ", address)
	cgm.clientMapLock.Lock()
	defer cgm.clientMapLock.Unlock()

	m, ok := cgm.clientMap[name]
	if !ok {
		m = make(map[string]*TcpClient)
		cgm.clientMap[name] = m
	} else {
		if _, ok = m[address]; ok {
			return
		}
	}

	client := NewTcpClient(name, 10*1024, cgm.protoName, address, cgm.callback)
	m[address] = client
	client.Serve()
}

func (cgm *TcpClientGroupManager) RemoveClient(name string, address string) {
	// glog.Info("TcpClientGroup RemoveClient name ", name, " address ", address)
	cgm.clientMapLock.Lock()
	defer cgm.clientMapLock.Unlock()

	m, ok := cgm.clientMap[name]
	if !ok {
		return
	}
	m, _ = cgm.clientMap[name]

	c, ok := m[address]
	if !ok {
		return
	}

	c.Stop()
	delete(cgm.clientMap[name], address)
}

func (cgm *TcpClientGroupManager) SendDataToAddress(name, address string, msg interface{}) error {
	cgm.clientMapLock.RLock()
	m, ok := cgm.clientMap[name]
	if !ok {
		cgm.clientMapLock.RUnlock()
		err := fmt.Errorf("sendDataToAddress - name not exists: %s", name)
		// glog.Error(err)
		return err
	}

	c, ok := m[address]
	if !ok {
		cgm.clientMapLock.RUnlock()
		err := fmt.Errorf("sendDataToAddress - address not exists: %s", address)
		// glog.Error(err)
		return err
	}

	cgm.clientMapLock.RUnlock()

	// glog.Infof("tcp_client_group_manager sendDataToAddress: {name: %s, conn: %s, msg: {%v}}", name, c, msg)
	return c.Send(msg)
}

func (cgm *TcpClientGroupManager) SendData(name string, msg interface{}) error {
	tcpConn := cgm.getRotationSession(name)
	if tcpConn == nil {
		return errors.New("can not get connection")
	}
	// glog.Info("tcp_client_group_manager SendData: {name: %s, conn: %s, msg: {%v}}", name, tcpConn, msg)
	return tcpConn.Send(msg)
}

func (cgm *TcpClientGroupManager) getRotationSession(name string) *TcpConnection {
	allConns := cgm.getTcpClientsByName(name)
	if allConns == nil || len(allConns) == 0 {
		return nil
	}

	index := rand.Int() % len(allConns)
	return allConns[index]
}

func (cgm *TcpClientGroupManager) BroadcastData(name string, msg interface{}) error {
	allConns := cgm.getTcpClientsByName(name)

	if allConns == nil || len(allConns) == 0 {
		return nil
	}

	for _, conn := range allConns {
		conn.Send(msg)
	}

	return nil
}

func (cgm *TcpClientGroupManager) getTcpClientsByName(name string) []*TcpConnection {
	var allConns []*TcpConnection

	cgm.clientMapLock.RLock()

	serviceMap, ok := cgm.clientMap[name]

	if !ok {
		cgm.clientMapLock.RUnlock()
		return nil
	}

	for _, c := range serviceMap {
		if c != nil && c.GetConnection() != nil && !c.GetConnection().IsClosed() {
			allConns = append(allConns, c.GetConnection())
		}
	}

	cgm.clientMapLock.RUnlock()

	return allConns
}
