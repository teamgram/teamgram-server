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
	"sync"
	"time"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/sync2"
	"github.com/nebula-chat/chatengine/pkg/queue2"
	"github.com/nebula-chat/chatengine/mtproto/rpc"
	"github.com/nebula-chat/chatengine/mtproto"
	"container/list"
	"reflect"
)

// import "container/list"

const (
	kNetworkMessageStateNone             = 0 // created
	kNetworkMessageStateReceived         = 1 // received from client
	kNetworkMessageStateRunning          = 2 // invoke api
	kNetworkMessageStateWaitReplyTimeout = 3 // invoke timeout
	kNetworkMessageStateInvoked          = 4 // invoke ok, send to client
	kNetworkMessageStatePushSync         = 5 // invoke ok, send to client
	kNetworkMessageStateAck              = 6 // received client ack
	kNetworkMessageStateWaitAckTimeout   = 7 // wait ack timeout
	kNetworkMessageStateError            = 8 // invalid error
	kNetworkMessageStateEnd              = 9 // end state
)

//const (
//	kReqConnTypeTcp 			= 0
//	kReqConnTypeHttpRpc 		= 1
//	kReqConnTypeHttpWait 		= 2
//	kReqConnTypeHttpRpcAndWait 	= 3
//)

// client connID
//TRANSPORT_TCP  = 1 // TCP
//TRANSPORT_HTTP = 2 // HTTP
type ClientConnID struct {
	connType       int
	clientConnID   uint64 // client -> frontend netlib connID
	frontendConnID uint64 // frontend -> session netlib connID
	createdAt      int64
}

func makeClientConnID(connType int, clientConnID, frontendConnID uint64) ClientConnID {
	connID := ClientConnID{
		connType:       connType,
		clientConnID:   clientConnID,
		frontendConnID: frontendConnID,
		createdAt:      time.Now().Unix(),
	}
	return connID
}

func (c ClientConnID) String() string {
	return fmt.Sprintf("{conn_type: %d, client_conn_id: %d, frontend_conn_id: %d}", c.connType, c.clientConnID, c.frontendConnID)
}

func (c ClientConnID) Equal(id ClientConnID) bool {
	// return c.connType == id.connType && c.clientConnID == id.clientConnID && c.frontendConnID == id.frontendConnID
	return c.clientConnID == id.clientConnID && c.frontendConnID == id.frontendConnID
}

type networkApiMessage struct {
	date       int64
	quickAckId int32 // 0: not use
	rpcRequest *mtproto.TLMessage2
	state      int // TODO(@benqi): sync.AtomicInt32
	rpcMsgId   int64
	rpcResult  mtproto.TLObject
}

type networkSyncMessage struct {
	date   int64
	update *mtproto.TLMessage2
	state  int
}

type rpcApiMessages struct {
	connID      ClientConnID
	cntl        *zrpc.ZRpcController
	sessionId   int64
	rpcMessages []*networkApiMessage
}

///////////////////////////////////////////////////////////////////////////////////
type sessionData struct {
	connID ClientConnID
	cntl   *zrpc.ZRpcController
	buf    []byte
}

type syncRpcResultData struct {
	// sessionID int64
	clientMsgId int64
	cntl        *zrpc.ZRpcController
	data        []byte
}

type syncData struct {
	cntl     *zrpc.ZRpcController
	pts      int32
	ptsCount int32
	data     *messageData
}

func makeSyncData(cntl *zrpc.ZRpcController, pts, ptsCount int32, data *messageData) *syncData {
	return &syncData{
		cntl:     cntl,
		pts:      pts,
		ptsCount: ptsCount,
		data:     data,
	}
}

type connData struct {
	isNew  bool
	connID ClientConnID
}

///////////////////////////////////////////////////////////////////////////////////
const (
	keyIdNew = 0
	keyLoaded = 1
	unauthorized = 2
	userIdLoaded = 3
	offline = 4
	closed = 5
)

type authSessions struct {
	Layer           int32
	authKeyId       int64
	authKey         []byte
	AuthUserId      int32
	cacheSalt       *mtproto.TLFutureSalt
	cacheLastSalt   *mtproto.TLFutureSalt
	sessions        []sessionBase
	unknownSessions *list.List
	closeChan       chan struct{}
	sessionDataChan chan interface{} // receive from client
	rpcDataChan     chan interface{} // rpc reply
	rpcQueue        *queue2.SyncQueue
	finish          sync.WaitGroup
	running         sync2.AtomicInt32
	state           int
	onlineExpired   int64
	// pushSessionId   int64
	// cacheSalt       *mtproto.TLFutureSalt
	// cacheLastSalt   *mtproto.TLFutureSalt
}

func makeAuthSessions(authKeyId int64) *authSessions {
	ss := &authSessions{
		authKeyId:       authKeyId,
		sessions:        make([]sessionBase, kSessionMaxSize),
		unknownSessions: list.New(),
		closeChan:       make(chan struct{}),
		sessionDataChan: make(chan interface{}, 1024),
		rpcDataChan:     make(chan interface{}, 1024),
		rpcQueue:        queue2.NewSyncQueue(),
		finish:          sync.WaitGroup{},
		state:           keyIdNew,
	}

	for i := 0; i < kSessionMaxSize; i++ {
		ss.sessions[i] = newSession(i, ss)
	}

	return ss
}

func (s *authSessions) getAuthKeyId() int64 {
	return s.authKeyId
}

func (s *authSessions) getAuthKey() []byte {
	return s.authKey
}

func (s *authSessions) getUserId() int32 {
	return s.AuthUserId
}

func (s *authSessions) setUserId(userId int32) {
	s.AuthUserId = userId
	s.onBindUser(userId)
}

func (s *authSessions) getCacheSalt() *mtproto.TLFutureSalt {
	return s.cacheSalt
}

func (s *authSessions) getLayer() int32 {
	return s.Layer
}

func (s *authSessions) setLayer(layer int32) {
	s.Layer = layer
}

func (s *authSessions) destroySession(sessionId int64) bool {
	// TODO(@benqi):
	return true
}

func (s *authSessions) sendToRpcQueue(rpcMessage *rpcApiMessages) {
	s.rpcQueue.Push(rpcMessage)
}

///////////////////////////////////////////////////////////////////////////////////////////////
func (s *authSessions) onBindAuthKey(authKey []byte) {
	// TODO(@benqi):
	// return true
	s.state = keyLoaded
	s.authKey = authKey

	// try load userId
	authUserId := getCacheUserID(s.authKeyId)
	if authUserId != 0 {
		s.onBindUser(authUserId)
	}
}

func (s *authSessions) onBindUser(userId int32) {
	// TODO(@benqi):
	s.state = userIdLoaded
	s.AuthUserId = userId

	pushSessionId := getCachePushSessionID(userId, s.authKeyId)
	if pushSessionId != 0 {
		s.onBindPushSessionId(pushSessionId)
	}

	if s.Layer == 0 {
		layer := getCacheApiLayer(s.authKeyId)
		s.onBindLayer(layer)
	}

	//setOnlineTTL(s.AuthUserId, s.authKeyId, getServerID(), s.Layer, 180)
	//s.onlineExpired = int64(time.Now().Unix() + 180)
}

func (s *authSessions) onBindPushSessionId(sessionId int64) {
	var (
		unknownSess sessionBase
		unknownE *list.Element
	)

	// found unknown
	for e := s.unknownSessions.Front(); e != nil; e = e.Next() {
		v, _ := e.Value.(sessionBase)
		if v.SessionId() == sessionId {
			unknownE = e
			unknownSess = v
			break
		}
	}

	pushSess := s.sessions[kSessionPush]
	if unknownSess == nil {
		if pushSess.SessionId() != sessionId {
			pushSess.SetSessionId(sessionId)
		}
	} else {
		pushSess.MergeSession(unknownSess)
		s.unknownSessions.Remove(unknownE)
	}
}

func (s *authSessions) onBindLayer(layer int32) {
	s.Layer = layer
}

func (s *authSessions) setOnline() {
	date := time.Now().Unix()
	if (s.onlineExpired == 0 || date > s.onlineExpired - kPingAddTimeout) && s.AuthUserId != 0 {
		// glog.Info("DEBUG] setOnline - set online ", s.onlineExpired)
		setOnlineTTL(s.AuthUserId, s.authKeyId, getServerID(), s.Layer, 60)
		s.onlineExpired = int64(time.Now().Unix() + 60)
	} else {
		// glog.Info("DEBUG] setOnline - not set online ", s.onlineExpired)
	}
}

func (s *authSessions) trySetOffline() {
	if !s.sessions[kSessionGeneric].sessionOnline() && !s.sessions[kSessionPush].sessionOnline() && s.AuthUserId != 0 {
		// glog.Info("DEBUG] setOffline - set offline ", s.onlineExpired)
		setOfflineTTL(s.AuthUserId, s.authKeyId, getServerID())
		s.onlineExpired = 0
	} else {
		// glog.Info("DEBUG] setOffline - not set offline ", s.onlineExpired)
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////
func (s *authSessions) String() string {
	return fmt.Sprintf("{auth_key_id: %d, user_id: %d}", s.authKeyId, s.AuthUserId)
}

func (s *authSessions) Start() {
	s.running.Set(1)
	s.finish.Add(1)
	go s.rpcRunLoop()
	go s.runLoop()
}

func (s *authSessions) Stop() {
	s.running.Set(0)
	s.rpcQueue.Close()
	// close(s.closeChan)
}

func (s *authSessions) runLoop() {
	defer func() {
		s.finish.Done()
		close(s.closeChan)
		s.finish.Wait()
	}()

	for s.running.Get() == 1 {
		select {
		case <-s.closeChan:
			// glog.Info("runLoop -> To Close ", this.String())
			return

		case sessionMsg, _ := <-s.sessionDataChan:
			switch sessionMsg.(type) {
			case *sessionData:
				s.onSessionData(sessionMsg.(*sessionData))
			case *syncRpcResultData:
				s.onSyncRpcResultData(sessionMsg.(*syncRpcResultData))
			case *syncData:
				s.onSyncData(sessionMsg.(*syncData))
			case *connData:
				s.onConnData(sessionMsg.(*connData))
			default:
				panic("receive invalid type msg")
			}
		case rpcMessages, _ := <-s.rpcDataChan:
			results, _ := rpcMessages.(*rpcApiMessages)
			s.onRpcResult(results)
		case <-time.After(time.Second):
			s.onTimer()
		}
	}

	glog.Info("quit runLoop...")
}

func (s *authSessions) rpcRunLoop() {
	for {
		apiRequests := s.rpcQueue.Pop()
		if apiRequests == nil {
			glog.Info("quit rpcRunLoop...")
			//if s.AuthUserId != 0 {
			//	setOfflineTTL(s.AuthUserId, s.authKeyId, getServerID())
			//}
			return
		} else {
			requests, _ := apiRequests.(*rpcApiMessages)
			s.onRpcRequest(requests)
		}
	}
}

func (s *authSessions) onTimer() {
	if s.sessions[kSessionPush].sessionOnline() || s.sessions[kSessionGeneric].sessionOnline() {
		// glog.Info("DEBUG] onTimer - set online")
		s.setOnline()
	}

	for _, v := range s.sessions {
		v.onTimer()
	}

	for e := s.unknownSessions.Front(); e != nil; e = e.Next() {
		e.Value.(sessionBase).onTimer()
	}

	if s.sessions[kSessionPush].sessionClosed() && s.sessions[kSessionGeneric].sessionClosed() {
		deleteClientSessionManager(s.authKeyId)
	}
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
// client
func (s *authSessions) onSessionClientNew(connID ClientConnID) error {
	select {
	case s.sessionDataChan <- &connData{true, connID}:
		return nil
	}
	return nil
}

func (s *authSessions) onSessionDataArrived(connID ClientConnID, cntl *zrpc.ZRpcController, buf []byte) error {
	select {
	case s.sessionDataChan <- &sessionData{connID, cntl, buf}:
		return nil
	}
	return nil
}

func (s *authSessions) onSessionClientClosed(connID ClientConnID) error {
	select {
	case s.sessionDataChan <- &connData{false, connID}:
		return nil
	}
	return nil
}

// push
func (s *authSessions) onSyncRpcResultDataArrived(clientMsgId int64, cntl *zrpc.ZRpcController, data []byte) error {
	select {
	case s.sessionDataChan <- &syncRpcResultData{clientMsgId,cntl, data}:
		return nil
	}
	return nil
}

func (s *authSessions) onSyncDataArrived(cntl *zrpc.ZRpcController, pts, ptsCount int32, data *messageData) error {
	select {
	case s.sessionDataChan <- makeSyncData(cntl, pts, ptsCount, data):
		return nil
	}
	return nil
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
func (s *authSessions) onSessionData(sessionMsg *sessionData) {
	var (
		err error
		now = int32(time.Now().Unix())
	)

	if s.authKey == nil {
		authKey := getCacheAuthKey(s.authKeyId)
		if authKey == nil {
			// err := fmt.Errorf("onSessionData - not found authKeyId")
			glog.Errorf("onSessionData - error: {not found authKeyId}, data: {sess: %s, conn_id: %s, md: %s}", s, sessionMsg.connID, sessionMsg.cntl)
			return
		} else {
			s.onBindAuthKey(authKey)
		}
	}

	message := mtproto.NewEncryptedMessage2(s.authKeyId)
	err = message.Decode(s.authKeyId, s.authKey, sessionMsg.buf[8:])
	if err != nil {
		// TODO(@benqi): close frontend conn??
		glog.Error(err)
		glog.Errorf("onSessionData - error: {%v}, data: {sess: %s, conn_id: %s, md: %s}", err, s, sessionMsg.connID, sessionMsg.cntl)

		return
	}

	glog.Infof("onSessionData - message: {%s}, data: {sess: %s, conn_id: %s, md: %s}", message, s, sessionMsg.connID, sessionMsg.cntl)

	if s.cacheSalt == nil {
		s.cacheSalt, s.cacheLastSalt, _ = getOrFetchNewSalt(s.authKeyId)
	} else {
		if now >= s.cacheSalt.GetValidUntil() {
			s.cacheSalt, s.cacheLastSalt, _ = getOrFetchNewSalt(s.authKeyId)
		}
	}

	if s.cacheSalt == nil {
		glog.Errorf("onSessionData - getOrFetchNewSalt nil error, data: {sess: %s, conn_id: %s, md: %s}", s, sessionMsg.connID, sessionMsg.cntl)
		return
	}

	sess := s.getOrCreateSession(sessionMsg.connID, message.SessionId, message.Object)
	if sess.SessionType() == kSessionUnknown && !sess.sessionOnline() {
		pushSessionId := getCachePushSessionID(s.AuthUserId, s.authKeyId)
		if pushSessionId != 0 && message.SessionId == pushSessionId {
			s.onBindPushSessionId(pushSessionId)
			sess = s.sessions[kSessionPush]
		}
	}

	message2 := &mtproto.TLMessage2{
		MsgId:  message.MessageId,
		Seqno:  message.SeqNo,
		Object: message.Object,
	}

	sess.onMessageData(sessionMsg.connID, sessionMsg.cntl, message.Salt, message2)
}

func (s *authSessions) onSyncRpcResultData(syncMsg *syncRpcResultData) {
	glog.Infof("onSyncRpcResultData - receive data: {sess: %s, md: %s}",
		s, syncMsg.cntl)

	genericSess := s.sessions[kSessionGeneric]
	genericSess.(*genericSession).onSyncRpcResultData(syncMsg.cntl, syncMsg.data)
}

func (s *authSessions) onSyncData(syncMsg *syncData) {
	glog.Infof("onSyncData - generic session: {pts: %d, pts_count: %d, updates: %s}",
		syncMsg.pts, syncMsg.ptsCount, reflect.TypeOf(syncMsg.data.obj))

	genericSess := s.sessions[kSessionGeneric].(*genericSession)
	pushSess := s.sessions[kSessionPush].(*pushSession)

	if pushSess.sessionOnline() {
		if syncMsg.ptsCount > 0 {
			glog.Infof("onSyncData - push session: {pts: %d, pts_count: %d, updates: %s}",
				syncMsg.pts, syncMsg.ptsCount, reflect.TypeOf(syncMsg.data.obj))
			pushSess.onSyncData(syncMsg.cntl)

			if genericSess.sessionOnline() {
				glog.Infof("onSyncData - generic session: {pts: %d, pts_count: %d, updates: %s}",
					syncMsg.pts, syncMsg.ptsCount, reflect.TypeOf(syncMsg.data.obj))
				genericSess.onSyncData(syncMsg.cntl, syncMsg.data.obj)
			}
		}
	} else {
		if genericSess.sessionOnline() {
			glog.Infof("onSyncData - generic session: {pts: %d, pts_count: %d, updates: %s}",
				syncMsg.pts, syncMsg.ptsCount, reflect.TypeOf(syncMsg.data.obj))
			genericSess.onSyncData(syncMsg.cntl, syncMsg.data.obj)
		}
	}
}

func (s *authSessions) onConnData(connMsg *connData) {
	sess := s.getSessionByConnId(connMsg.connID)
	if sess != nil {
		if connMsg.isNew {
			glog.Warning("session conn created: ", connMsg)
			sess.onNewSession(connMsg.connID, 0)
		} else {
			glog.Warning("session conn closed: ", connMsg)
			sess.onSessionClose(connMsg.connID)
		}
	} else {
		glog.Warning("session not new: ", connMsg)
	}
}

func (s *authSessions) onRpcResult(rpcResults *rpcApiMessages) {
	switch rpcResults.sessionId {
	case s.sessions[kSessionGeneric].SessionId():
		s.sessions[kSessionGeneric].(*genericSession).onRpcResult(rpcResults)
	case s.sessions[kSessionUpload].SessionId():
		s.sessions[kSessionUpload].(*uploadSession).onRpcResult(rpcResults)
	case s.sessions[kSessionDownload].SessionId():
		s.sessions[kSessionDownload].(*downloadSession).onRpcResult(rpcResults)
	default:
		glog.Warning("onRpcResult - not found rpcSession by sessionId: ", rpcResults.sessionId)
		return
	}
}

func (s *authSessions) onRpcRequest(requests *rpcApiMessages) {
	var rpcMessageList []*networkApiMessage

	switch requests.sessionId {
	case s.sessions[kSessionGeneric].SessionId():
		rpcMessageList = s.sessions[kSessionGeneric].(*genericSession).onInvokeRpcRequest(s.AuthUserId, s.authKeyId, s.Layer, requests)
	case s.sessions[kSessionUpload].SessionId():
		rpcMessageList = s.sessions[kSessionUpload].(*uploadSession).onInvokeRpcRequest(s.AuthUserId, s.authKeyId, s.Layer,requests)
	case s.sessions[kSessionDownload].SessionId():
		rpcMessageList = s.sessions[kSessionDownload].(*downloadSession).onInvokeRpcRequest(s.AuthUserId, s.authKeyId, s.Layer,requests)
	default:
		glog.Warning("onRpcRequest - not found rpcSession by sessionId: ", requests.sessionId)
		return
	}

	// TODO(@benqi): rseult metadata
	requests.rpcMessages = rpcMessageList
	s.rpcDataChan <- requests
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
func (s *authSessions) getSessionByConnId(connId ClientConnID) sessionBase {
	for _, sess2 := range s.sessions {
		if sess2.checkConnIdExist(connId) {
			return sess2
		}
	}
	return nil
}

func (s *authSessions) getRpcSession(sessionId int64) sessionBase {
	var sess sessionBase
	switch sessionId {
	case s.sessions[kSessionGeneric].SessionId():
		sess = s.sessions[kSessionGeneric]
	case s.sessions[kSessionUpload].SessionId():
		sess = s.sessions[kSessionUpload]
	case s.sessions[kSessionDownload].SessionId():
		sess = s.sessions[kSessionDownload]
	}
	return sess
}

func (s *authSessions) getOrCreateSession(connId ClientConnID, sessionId int64, request mtproto.TLObject) sessionBase {
	var (
		sess sessionBase
		unknownE *list.Element
	)

	for _, sess2 := range s.sessions {
		if sess2 != nil && sess2.SessionId() == sessionId {
			sess = sess2
			break
		}
	}

	if sess == nil {
		// found unknown
		for e := s.unknownSessions.Front(); e != nil; e = e.Next() {
			v, _ := e.Value.(sessionBase)
			if v.SessionId() == sessionId {
				unknownE = e
				sess = v
				break
			}
		}
	}

	if sess != nil && sess.SessionType() != kSessionUnknown {
		return sess
	}

	sessType := kSessionUnknown
	getSessionType2(request, &sessType)

	if sessType != kSessionUnknown {
		// rpc or ...
		if sess == nil {
			sess = s.sessions[sessType]
			sess.onNewSession(connId, sessionId)
		} else {
			s.sessions[sessType].MergeSession(sess)
			s.unknownSessions.Remove(unknownE)
		}
	} else {
		if sess == nil {
			sess = newSession(kSessionUnknown, s)
			s.unknownSessions.PushBack(sess)
			sess.onNewSession(connId, sessionId)
		}
	}
	return sess
}