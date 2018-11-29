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

package server

import (
	"fmt"
	"github.com/golang/glog"
	"sync"
	"github.com/nebula-chat/chatengine/mtproto/rpc"
	"github.com/nebula-chat/chatengine/mtproto"
)

type sessionManager struct {
	sessions sync.Map // map[int64]*sessionClientList
}

func newSessionManager() *sessionManager {
	return &sessionManager{}
}

func (s *sessionManager) onSessionClientNew(connID uint64, cntl *zrpc.ZRpcController, sessData *mtproto.TLSessionClientCreated) error {
	glog.Infof("onSessionClientNew - receive data: {client_conn_id: %s, md: %s, sess_data: %s}", connID, cntl.RpcMeta, sessData)

	var sessList *clientSessionManager

	if vv, ok := s.sessions.Load(sessData.GetAuthKeyId()); !ok {
		err := fmt.Errorf("onSessionClientNew - not find sessionList by authKeyId: {%d}", sessData.GetAuthKeyId())
		glog.Warning(err)
		return err
	} else {
		sessList, _ = vv.(*clientSessionManager)
	}

	clientConnID := makeClientConnID(int(sessData.GetConnType()), connID, uint64(sessData.GetClientConnId()))
	return sessList.onSessionClientNew(clientConnID)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////
func (s *sessionManager) onSessionData(connID uint64, cntl *zrpc.ZRpcController, sessData *mtproto.TLSessionMessageData) error {
	glog.Infof("onSessionData - receive data: {conn_id: %d, md: %s, sess_data: %s}",
		connID,
		cntl.RpcMeta,
		sessData)

	authKeyId := sessData.GetAuthKeyId()
	var sessList *clientSessionManager
	if vv, ok := s.sessions.Load(authKeyId); !ok {
		sessList = newClientSessionManager(authKeyId)
		s.sessions.Store(authKeyId, sessList)
		s.onNewSessionClientManager(sessList)
	} else {
		sessList, _ = vv.(*clientSessionManager)
	}

	clientConnID := makeClientConnID(int(sessData.GetConnType()), connID, uint64(sessData.GetClientConnId()))
	return sessList.OnSessionDataArrived(clientConnID, cntl, cntl.MoveAttachment())
}

func (s *sessionManager) onSessionClientClosed(connID uint64, cntl *zrpc.ZRpcController, sessData *mtproto.TLSessionClientClosed) error {
	glog.Infof("onSessionClientClosed - receive data: {client_conn_id: %d, md: %s, sess_data: %s}",
		connID,
		cntl,
		sessData)

	var sessList *clientSessionManager

	if vv, ok := s.sessions.Load(sessData.GetAuthKeyId()); !ok {
		err := fmt.Errorf("onSessionClientClosed - not find sessionList by authKeyId: {%d}", sessData.GetAuthKeyId())
		glog.Warning(err)
		return err
	} else {
		sessList, _ = vv.(*clientSessionManager)
	}

	clientConnID := makeClientConnID(int(sessData.GetConnType()), connID, uint64(sessData.GetClientConnId()))
	return sessList.onSessionClientClosed(clientConnID)
}

func (s *sessionManager) onSyncRpcResultData(authKeyId, clientMsgId int64, cntl *zrpc.ZRpcController) error {
	glog.Infof("onSyncRpcResultData - receive data: {auth_key_id: %d, client_msg_id: %d, md: %s}",
		authKeyId,
		clientMsgId,
		cntl)

	rawData := cntl.MoveAttachment()

	var sessList *clientSessionManager
	if vv, ok := s.sessions.Load(authKeyId); !ok {
		err := fmt.Errorf("pushToSessionData - not find sessionList by authKeyId: {%d}", authKeyId)
		glog.Warning(err)
		return err
	} else {
		sessList, _ = vv.(*clientSessionManager)
	}

	return sessList.OnSyncRpcResultDataArrived(clientMsgId, cntl, rawData)
}

func (s *sessionManager) onSyncData(authKeyId int64, cntl *zrpc.ZRpcController) error {
	glog.Infof("onSyncData - receive data: {auth_key_id: %d, md: %s}",
		authKeyId,
		cntl)

	dbuf := mtproto.NewDecodeBuf(cntl.MoveAttachment())
	obj := dbuf.Object()
	if obj == nil {
		return dbuf.GetError()
	}

	var sessList *clientSessionManager
	if vv, ok := s.sessions.Load(authKeyId); !ok {
		err := fmt.Errorf("pushToSessionData - not find sessionList by authKeyId: {%d}", authKeyId)
		glog.Warning(err)
		return err
	} else {
		sessList, _ = vv.(*clientSessionManager)
	}

	return sessList.OnSyncDataArrived(cntl, &messageData{obj: obj})
}

////////////////////////////////////////////////////////////////////////////////////////////////////////
// session event
func (s *sessionManager) onNewSessionClientManager(sess *clientSessionManager) {
	sess.Start()
}

func (s *sessionManager) onCloseSessionClientManager(authKeyId int64) {
	if vv, ok := s.sessions.Load(authKeyId); ok {
		vv.(*clientSessionManager).Stop()
		s.sessions.Delete(authKeyId)
	}
}
