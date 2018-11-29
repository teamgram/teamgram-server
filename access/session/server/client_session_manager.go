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
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/queue2"
	"github.com/nebula-chat/chatengine/pkg/sync2"
	"github.com/nebula-chat/chatengine/mtproto"
	"sync"
	"time"
	"github.com/nebula-chat/chatengine/mtproto/rpc"
)

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
	// receiveCount   int		// httpReq
	// sendCount      int		// httpRsp
}

func makeClientConnID(connType int, clientConnID, frontendConnID uint64) ClientConnID {
	connID := ClientConnID{
		connType:       connType,
		clientConnID:   clientConnID,
		frontendConnID: frontendConnID,
		// receiveCount:   0,
		// sendCount:      0,
	}
	return connID
}

func (c ClientConnID) String() string {
	return fmt.Sprintf("{conn_type: %d, client_conn_id: %d, frontend_conn_id: %d}", c.connType, c.clientConnID, c.frontendConnID)
}

func (c ClientConnID) Equal(id ClientConnID) bool {
	return c.connType == id.connType && c.clientConnID == id.clientConnID && c.frontendConnID == id.frontendConnID
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
	cntl   *zrpc.ZRpcController
	sessionId   int64
	rpcMessages []*networkApiMessage
}

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
	// sessionID int64
	cntl *zrpc.ZRpcController
	data *messageData
}

type connData struct {
	isNew  bool
	connID ClientConnID
}

////////////////////////////////////////
const (
// inited --> work --> idle --> quit
)

type clientSessionManager struct {
	Layer           int32
	authKeyId       int64
	authKey         []byte
	cacheSalt       *mtproto.TLFutureSalt
	cacheLastSalt   *mtproto.TLFutureSalt
	AuthUserId      int32
	sessions        map[int64]*clientSessionHandler
	updatesSession  *clientUpdatesHandler
	bizRPCClient    *grpc_util.RPCClient
	nbfsRPCClient   *grpc_util.RPCClient
	syncRpcClient   mtproto.RPCSyncClient
	closeChan       chan struct{}
	sessionDataChan chan interface{} // receive from client
	rpcDataChan     chan interface{} // rpc reply
	rpcQueue        *queue2.SyncQueue
	finish          sync.WaitGroup
	running         sync2.AtomicInt32
	state           int
}

func newClientSessionManager(authKeyId int64) *clientSessionManager {
// func newClientSessionManager(authKeyId int64, authKey []byte, userId int32) *clientSessionManager {
	bizRPCClient, _ := getBizRPCClient()
	nbfsRPCClient, _ := getNbfsRPCClient()
	syncRpcClient, _ := getSyncRPCClient()
	// cacheSalt, _ := getOrFetchNewSalt(authKeyId)

	return &clientSessionManager{
		authKeyId:       authKeyId,
		// authKey:         authKey,
		// cacheSalt:       cacheSalt,
		// AuthUserId:      userId,
		sessions:        make(map[int64]*clientSessionHandler),
		updatesSession:  newClientUpdatesHandler(),
		bizRPCClient:    bizRPCClient,
		nbfsRPCClient:   nbfsRPCClient,
		syncRpcClient:   syncRpcClient,
		closeChan:       make(chan struct{}),
		sessionDataChan: make(chan interface{}, 1024),
		rpcDataChan:     make(chan interface{}, 1024),
		rpcQueue:        queue2.NewSyncQueue(),
		finish:          sync.WaitGroup{},
	}
}

func (s *clientSessionManager) String() string {
	return fmt.Sprintf("{auth_key_id: %d, user_id: %d}", s.authKeyId, s.AuthUserId)
}

func (s *clientSessionManager) Start() {
	s.running.Set(1)
	s.finish.Add(1)
	go s.rpcRunLoop()
	go s.runLoop()
}

func (s *clientSessionManager) Stop() {
	s.running.Set(0)
	s.rpcQueue.Close()
	// close(s.closeChan)
}

func (s *clientSessionManager) runLoop() {
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

func (s *clientSessionManager) rpcRunLoop() {
	for {
		apiRequests := s.rpcQueue.Pop()
		if apiRequests == nil {
			glog.Info("quit rpcRunLoop...")
			return
		} else {
			requests, _ := apiRequests.(*rpcApiMessages)
			s.onRpcRequest(requests)
		}
	}
}

func (s *clientSessionManager) onSessionClientNew(connID ClientConnID) error {
	select {
	case s.sessionDataChan <- &connData{true, connID}:
		return nil
	}
	return nil
}

func (s *clientSessionManager) OnSessionDataArrived(connID ClientConnID, cntl *zrpc.ZRpcController, buf []byte) error {
	select {
	case s.sessionDataChan <- &sessionData{connID, cntl, buf}:
		return nil
	}
	return nil
}

func (s *clientSessionManager) onSessionClientClosed(connID ClientConnID) error {
	select {
	case s.sessionDataChan <- &connData{false, connID}:
		return nil
	}
	return nil
}

func (s *clientSessionManager) OnSyncRpcResultDataArrived(clientMsgId int64, cntl *zrpc.ZRpcController, data []byte) error {
	select {
	case s.sessionDataChan <- &syncRpcResultData{clientMsgId,cntl, data}:
		return nil
	}
	return nil
}

func (s *clientSessionManager) OnSyncDataArrived(cntl *zrpc.ZRpcController, data *messageData) error {
	select {
	case s.sessionDataChan <- &syncData{cntl, data}:
		return nil
	}
	return nil
}

type messageListWrapper struct {
	messages []*mtproto.TLMessage2
}

func (s *clientSessionManager) onSessionData(sessionMsg *sessionData) {
	glog.Infof("onSessionData - receive data: {sess: %s, conn_id: %s, md: %s}", s, sessionMsg.connID, sessionMsg.cntl)

	var (
		err error
		now = int32(time.Now().Unix())
	)

	if s.authKey == nil {
		s.authKey = getCacheAuthKey(s.authKeyId)
		if s.authKey == nil {
			err := fmt.Errorf("onSessionData - not found authKeyId")
			glog.Error(err)
			return
		}
	}

	message := mtproto.NewEncryptedMessage2(s.authKeyId)
	err = message.Decode(s.authKeyId, s.authKey, sessionMsg.buf[8:])
	if err != nil {
		// TODO(@benqi): close frontend conn??
		glog.Error(err)
		return
	}

	glog.Infof("sessionDataChan: ", message)

	//// check message_id
	//if message.MessageId % 4 != 0 {
	//	err = fmt.Errorf("invalid message id %d", message.MessageId)
	//	glog.Error(err)
	//	return
	//}
	//
	//if message.MessageId&0xffffffff == 0 {
	//	err = fmt.Errorf("the lower 32 bits of msg_id passed by the client must not be empty: %d", message.MessageId)
	//	glog.Error(err)
	//
	//	// TODO(@benqi): replay-attack, close client conn.
	//	return
	//}
	//

	// https://core.telegram.org/mtproto/description#server-salt

	// Server Salt
	//
	// A (random) 64-bit number periodically (say, every 24 hours) changed
	// (separately for each session) at the request of the server.
	// All subsequent messages must contain the new salt
	// (although, messages with the old salt are still accepted for a further 300 seconds).
	// Required to protect against replay attacks and certain tricks
	// associated with adjusting the client clock to a moment in the distant future.
	//
	if s.cacheSalt == nil {
		s.cacheSalt, s.cacheLastSalt, _ = getOrFetchNewSalt(s.authKeyId)
	} else {
		if now >= s.cacheSalt.GetValidUntil() {
			s.cacheSalt, s.cacheLastSalt, _ = getOrFetchNewSalt(s.authKeyId)
		}
	}

	if s.cacheSalt == nil {
		err = fmt.Errorf("getOrFetchNewSalt error")
		// TODO(@benqi): close client conn
		glog.Error(err)
		return
	}

	glog.Info("salt: ", message.Salt, ", cacheSalt: ", s.cacheSalt, ", cacheLastSalt: ", s.cacheLastSalt)
	sess, ok := s.sessions[message.SessionId]
	if !ok {
		sess = newClientSessionHandler(message.SessionId, s)
		s.sessions[message.SessionId] = sess
	}

	message2 := &mtproto.TLMessage2{
		MsgId:  message.MessageId,
		Seqno:  message.SeqNo,
		Object: message.Object,
	}

	sess.processMessage(sessionMsg.connID, sessionMsg.cntl, message.Salt, message2)

/*
	if !sess.CheckBadServerSalt(sessionMsg.connID, sessionMsg.cntl, message.MessageId, message.SeqNo, message.Salt) {
		glog.Infof("salt invalid - {sess: %s, conn_id: %s, md: %s}", s, sessionMsg.connID, sessionMsg.cntl)
		// glog.Error("salt invalid..")
		return
	}

	_, isContainer := message.Object.(*mtproto.TLMsgContainer)
	if !sess.CheckBadMsgNotification(sessionMsg.connID, sessionMsg.cntl, message.MessageId, message.SeqNo, isContainer) {
		glog.Infof("bad msg invalid - {sess: %s, conn_id: %s, md: %s}", s, sessionMsg.connID, sessionMsg.cntl)
		// glog.Error("bad msg invalid..")
		return
	}

	var messages = &messageListWrapper{[]*mtproto.TLMessage2{}}
	extractClientMessage(message.MessageId, message.SeqNo, message.Object, messages, func(layer int32) {
		s.Layer = layer
		// TODO(@benqi): clear session_manager
	})

	if !ok {
		s.sessions[message.SessionId] = sess
		glog.Info("newClientSession: ", sess)
		sess.onNewSessionCreated(sessionMsg.connID, sessionMsg.cntl, message.MessageId)
		// sess.clientConnID = sessionMsg.connID
		sess.clientState = kStateOnline
	} else {
		// New Session Creation Notification
		//
		// The server notifies the client that a new session (from the server’s standpoint)
		// had to be created to handle a client message.
		// If, after this, the server receives a message with an even smaller msg_id within the same session,
		// a similar notification will be generated for this msg_id as well.
		// No such notifications are generated for high msg_id values.
		//
		if message.MessageId < sess.firstMsgId {
			glog.Info("message.MessageId < sess.firstMsgId: ", message.MessageId, ", ", sess.firstMsgId, ", sessionId: ", message.SessionId)
			sess.firstMsgId = message.MessageId
			sess.onNewSessionCreated(sessionMsg.connID, sessionMsg.cntl, message.MessageId)
		}
	}

	// sess.onClientMessage(message.MessageId, message.SeqNo, message.Object, messages)
	sess.onMessageData(sessionMsg.connID, sessionMsg.cntl, messages.messages)
 */
}

func (s *clientSessionManager) onTimer() {
	var delList = []int64{}
	for k, v := range s.sessions {
		if !v.onTimer() {
			delList = append(delList, k)
		}
	}

	for _, id := range delList {
		delete(s.sessions, id)
	}

	if len(s.sessions) == 0 {
		deleteClientSessionManager(s.authKeyId)
	}
}


func (s *clientSessionManager) onSyncRpcResultData(syncMsg *syncRpcResultData) {
	glog.Infof("onSyncRpcResultData - receive data: {sess: %s, md: %s}",
		s, syncMsg.cntl)

	s.updatesSession.onSyncRpcResultData(syncMsg.cntl, syncMsg.data)
}

func (s *clientSessionManager) onSyncData(syncMsg *syncData) {
	glog.Infof("onSyncData - receive data: {sess: %s, md: %s, data: {%v}}",
		s, syncMsg.cntl, syncMsg.data)

	s.updatesSession.onSyncData(syncMsg.cntl, syncMsg.data.obj)
}

func (s *clientSessionManager) onConnData(connMsg *connData) {
	if connMsg.isNew {

	} else {
		s.updatesSession.UnSubscribeUpdates(connMsg.connID)
	}
}

func (s *clientSessionManager) onRpcResult(rpcResults *rpcApiMessages) {
	if sess, ok := s.sessions[rpcResults.sessionId]; ok {
		var hasAuthLogout = false
		msgList := sess.pendingMessages
		sess.pendingMessages = []*pendingMessage{}
		for _, m := range rpcResults.rpcMessages {
			msgList = append(msgList, &pendingMessage{mtproto.GenerateMessageId(), true, m.rpcResult})
			if _, ok := m.rpcRequest.Object.(*mtproto.TLAuthLogOut); ok {
				hasAuthLogout = true
				break
			}
		}
		if len(msgList) > 0 {
			sess.sendPendingMessagesToClient(rpcResults.connID, rpcResults.cntl, msgList)
		}

		if hasAuthLogout {
			deleteClientSessionManager(s.authKeyId)
		}
	}
}

func (s *clientSessionManager) PushApiRequest(apiRequest *mtproto.TLMessage2) {
	s.rpcQueue.Push(apiRequest)
}

func (s *clientSessionManager) onRpcRequest(requests *rpcApiMessages) {
	glog.Infof("onRpcRequest - receive data: {sess: %s, session_id: %d, conn_id: %d, md: %s, data: {%v}}",
		s, requests.sessionId, requests.connID, requests.cntl, requests.rpcMessages)

	rpcMessageList := make([]*networkApiMessage, 0, len(requests.rpcMessages))

	for i := 0; i < len(requests.rpcMessages); i++ {
		var (
			err         error
			rpcResult   mtproto.TLObject
		)

		// 初始化metadata
		rpcMetadata := &grpc_util.RpcMetadata{
			ServerId:        getServerID(),
			NetlibSessionId: int64(requests.connID.clientConnID),
			AuthId:          s.authKeyId,
			SessionId:       requests.sessionId,
			TraceId:         requests.cntl.RpcMeta.GetRequest().GetTraceId(),
			SpanId:          getUUID(),
			ReceiveTime:     time.Now().Unix(),
			UserId:          s.AuthUserId,
			ClientMsgId:     requests.rpcMessages[i].rpcRequest.MsgId,
		}

		if s.Layer == 0 {
			s.Layer = getCacheApiLayer(s.authKeyId)
		}
		rpcMetadata.Layer = s.Layer

		// TODO(@benqi): change state.
		requests.rpcMessages[i].state = kNetworkMessageStateRunning

		// TODO(@benqi): rpc proxy
		if checkRpcUploadRequest(requests.rpcMessages[i].rpcRequest.Object) {
			rpcResult, err = s.nbfsRPCClient.Invoke(rpcMetadata, requests.rpcMessages[i].rpcRequest.Object)
		} else if checkRpcDownloadRequest(requests.rpcMessages[i].rpcRequest.Object) {
			rpcResult, err = s.nbfsRPCClient.Invoke(rpcMetadata, requests.rpcMessages[i].rpcRequest.Object)
		} else {
			rpcResult, err = s.bizRPCClient.Invoke(rpcMetadata, requests.rpcMessages[i].rpcRequest.Object)
		}

		reply := &mtproto.TLRpcResult{
			ReqMsgId: requests.rpcMessages[i].rpcRequest.MsgId,
		}

		if err != nil {
			glog.Error(err)
			rpcErr, _ := err.(*mtproto.TLRpcError)
			if rpcErr.GetErrorCode() == int32(mtproto.TLRpcErrorCodes_NOTRETURN_CLIENT) {
				continue
			}
			reply.Result = rpcErr
		} else {
			// glog.Infof("OnMessage - rpc_result: {%v}\n", rpcResult)
			reply.Result = rpcResult
		}

		requests.rpcMessages[i].state = kNetworkMessageStateInvoked
		requests.rpcMessages[i].rpcResult = reply

		rpcMessageList = append(rpcMessageList, requests.rpcMessages[i])

		if _, ok := requests.rpcMessages[i].rpcRequest.Object.(*mtproto.TLAuthLogOut); ok {
			glog.Info("authLogOut - ", rpcMetadata)
			putCacheUserId(s.authKeyId, 0)
			break
		}
	}

	// TODO(@benqi): rseult metadata
	requests.rpcMessages = rpcMessageList
	s.rpcDataChan <- requests
}

// TODO(@benqi): status_client
func (s *clientSessionManager) setUserOnline(sessionId int64, connID ClientConnID) {
	defer func() {
		if r := recover(); r != nil {
			glog.Error(r)
		}
	}()

	setOnline(s.AuthUserId, s.authKeyId, getServerID(), s.Layer)
}

//==================================================================================================
type InitConnectionHandler func(layer int32)

