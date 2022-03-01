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

package service

import (
	"context"
	"fmt"
	"math"
	"reflect"
	"runtime/debug"
	"sync"
	"time"

	"github.com/teamgram/marmota/pkg/queue2"
	"github.com/teamgram/marmota/pkg/sync2"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/proto/mtproto/rpc/metadata"
	"github.com/teamgram/teamgram-server/app/service/authsession/authsession"
	"github.com/teamgram/teamgram-server/app/service/status/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// import "container/list"

const (
	clientUnknown  = 0
	clientAndroid  = 1
	clientiOS      = 2
	clientTdesktop = 3
	clientMacSwift = 4
	clientWebogram = 5
	clientReact    = 6
)

type rpcApiMessage struct {
	sessionId int64
	clientIp  string
	reqMsgId  int64
	reqMsg    mtproto.TLObject
	rpcResult *mtproto.TLRpcResult
}

func (m *rpcApiMessage) DebugString() string {
	if m.rpcResult == nil {
		return fmt.Sprintf("{session_id: %d, req_msg_id: %d, req_msg: %s}",
			m.sessionId,
			m.reqMsgId,
			m.reqMsg.DebugString())
	} else {
		return fmt.Sprintf("{session_id: %d, req_msg_id: %d, req_msg: %s, rpc_result: %s}",
			m.sessionId,
			m.reqMsgId,
			m.reqMsg.DebugString(),
			m.rpcResult.Result.DebugString())
	}
}

///////////////////////////////////////////////////////////////////////////////////
type sessionData struct {
	gatewayId string
	clientIp  string
	sessionId int64
	salt      int64
	buf       []byte
}

type sessionHttpData struct {
	gatewayId  string
	clientIp   string
	sessionId  int64
	salt       int64
	buf        []byte
	resChannel chan interface{}
}

type syncRpcResultData struct {
	sessionId   int64
	clientMsgId int64
	data        []byte
}

type syncSessionData struct {
	sessionId int64
	data      *messageData
}

type syncData struct {
	needAndroidPush bool
	data            *messageData
}

func makeSyncData(needAndroidPush bool, data *messageData) *syncData {
	return &syncData{
		needAndroidPush: needAndroidPush,
		data:            data,
	}
}

type connData struct {
	isNew     bool
	gatewayId string
	sessionId int64
}

func (c *connData) DebugString() string {
	return fmt.Sprintf("{isNew: %d, gatewayId: %s, sessionId: %d}", c.isNew, c.gatewayId, c.sessionId)
}

///////////////////////////////////////////////////////////////////////////////////
const (
	keyIdNew     = 0
	keyLoaded    = 1
	unauthorized = 2
	userIdLoaded = 3
	offline      = 4
	closed       = 5
	unknownError = 6
)

type authSessionsCallback interface {
	SendDataToGate(ctx context.Context, serverId int32, authKeyId, sessionId int64, payload []byte) error
}

type authSessions struct {
	authKeyId       int64
	Layer           int32
	Client          string
	Langpack        string
	AuthUserId      int64 // 不为0，则signIn
	cacheSalt       *mtproto.TLFutureSalt
	cacheLastSalt   *mtproto.TLFutureSalt
	pushSessionId   int64
	sessions        map[int64]*session
	closeChan       chan struct{}
	sessionDataChan chan interface{} // receive from client
	rpcDataChan     chan interface{} // rpc reply
	rpcQueue        *queue2.SyncQueue
	finish          sync.WaitGroup
	running         sync2.AtomicInt32
	state           int
	onlineExpired   int64
	clientType      int
	nextNotifyId    int64
	nextPushId      int64
	*Service
}

func newAuthSessions(authKeyId int64, s2 *Service) (*authSessions, error) {
	//keyData, err := s2.Dao.AuthsessionClient.AuthsessionGetAuthStateData(context.Background(), &authsession.TLAuthsessionGetAuthStateData{
	//	AuthKeyId: authKeyId,
	//})
	//if err != nil {
	//	logx.Errorf("getKeyStateData error: %v", err)
	//	return nil, err
	//}

	s := &authSessions{
		authKeyId:       authKeyId,
		Layer:           0,
		AuthUserId:      0,
		sessions:        make(map[int64]*session),
		closeChan:       make(chan struct{}),
		sessionDataChan: make(chan interface{}, 1024),
		rpcDataChan:     make(chan interface{}, 1024),
		rpcQueue:        queue2.NewSyncQueue(),
		finish:          sync.WaitGroup{},
		clientType:      clientUnknown,
		nextPushId:      0,
		nextNotifyId:    math.MaxInt32,
		Service:         s2,
	}

	s.Start()
	return s, nil
}

func (s *authSessions) getNextNotifyId() (id int64) {
	id = s.nextNotifyId
	s.nextNotifyId--
	return
}

func (s *authSessions) getNextPushId() (id int64) {
	id = s.nextPushId
	s.nextPushId++
	return
}

func (s *authSessions) getAuthKeyId() int64 {
	return s.authKeyId
}

func (s *authSessions) getTempAuthKeyId() int64 {
	return s.authKeyId
}

func (s *authSessions) getUserId() int64 {
	return s.AuthUserId
}

func (s *authSessions) setUserId(userId int64) {
	s.AuthUserId = userId
	s.onBindUser(userId)
}

func (s *authSessions) getCacheSalt() *mtproto.TLFutureSalt {
	return s.cacheSalt
}

func (s *authSessions) getLayer() int32 {
	if s.Layer == 0 {
		s.Layer, _ = s.Dao.GetCacheApiLayer(context.Background(), s.authKeyId)
	}
	return s.Layer
}

func (s *authSessions) setLayer(layer int32) {
	if layer != 0 {
		s.Layer = layer
		s.Dao.PutCacheApiLayer(context.Background(), s.authKeyId, layer)
	}
}

func (s *authSessions) getClient() string {
	if s.Client == "" {
		s.Client = s.Dao.GetCacheClient(context.Background(), s.authKeyId)
	}
	return s.Client
}

func (s *authSessions) setClient(c string) {
	if c != "" {
		s.Client = c
		s.Dao.PutCacheClient(context.Background(), s.authKeyId, c)
	}
}

func (s *authSessions) getLangpack() string {
	if s.Langpack == "" {
		s.Langpack = s.Dao.GetCacheLangpack(context.Background(), s.authKeyId)
	}
	return s.Langpack
}

func (s *authSessions) setLangpack(c string) {
	if c != "" {
		s.Langpack = c
		s.Dao.PutCacheLangpack(context.Background(), s.authKeyId, c)
	}
}

func (s *authSessions) destroySession(sessionId int64) bool {
	// TODO(@benqi):
	if _, ok := s.sessions[sessionId]; ok {
		// s.updates.onGenericSessionClose(sess)
		delete(s.sessions, sessionId)
	} else {
		//
	}
	return true
}

func (s *authSessions) sendToRpcQueue(rpcMessage *rpcApiMessage) {
	s.rpcQueue.Push(rpcMessage)
}

func (s *authSessions) getPushSessionId() int64 {
	if s.pushSessionId == 0 && s.AuthUserId != 0 {
		s.pushSessionId, _ = s.Dao.GetCachePushSessionID(context.Background(), s.AuthUserId, s.authKeyId)
	}
	return s.pushSessionId
}

func (s *authSessions) onBindUser(userId int64) {
	// TODO(@benqi):
	s.state = userIdLoaded
	s.AuthUserId = userId

	s.getPushSessionId()

	if s.Layer == 0 {
		layer, _ := s.Dao.GetCacheApiLayer(context.Background(), s.authKeyId)
		if layer != 0 {
			s.onBindLayer(layer)
		}
	}
}

func (s *authSessions) onBindPushSessionId(sessionId int64) {
	if s.pushSessionId == 0 {
		s.pushSessionId = sessionId
	}
	sess, _ := s.sessions[sessionId]
	if sess != nil {
		sess.isAndroidPush = true
		sess.cb.setOnline()
	}
}

func (s *authSessions) onBindLayer(layer int32) {
	s.Layer = layer
}

func (s *authSessions) setOnline() {
	//setOnlineTTL(s.AuthUserId, s.authKeyId, getServerID(), s.Layer, 60)
	date := time.Now().Unix()
	if (s.onlineExpired == 0 || date > s.onlineExpired-kPingAddTimeout) && s.AuthUserId != 0 {
		logx.Infof("DEBUG] setOnline - set online: (date: %d, userId:%d, onlineExpired: %d, authKeyId: %d)",
			date,
			s.AuthUserId,
			s.onlineExpired,
			s.authKeyId)

		// s.setOnlineTTL(s.AuthUserId, s.authKeyId, getServerID(), s.Layer, 60)
		s.Dao.StatusClient.StatusSetSessionOnline(context.Background(), &status.TLStatusSetSessionOnline{
			UserId:    s.AuthUserId,
			AuthKeyId: s.authKeyId,
			Gateway:   s.serverId,
			Expired:   0,
			Layer:     0,
		})
		//s.AuthUserId, s.authKeyId, env.Hostname)
		s.onlineExpired = int64(date + 60)
	} else {
		//log.Infof("DEBUG] setOnline - not set online: (date: %d, onlineExpired: %d, AuthUserId: %d)",
		//	date,
		//	s.onlineExpired,
		//	s.AuthUserId)
	}
}

func (s *authSessions) trySetOffline() {
	for _, sess := range s.sessions {
		if (sess.isGeneric && sess.sessionOnline()) ||
			(sess.isAndroidPush && sess.sessionOnline()) {
			return
		}
	}

	logx.Infof("authSessions]]>> offline: %s", s)

	s.Dao.StatusClient.StatusSetSessionOffline(context.Background(), &status.TLStatusSetSessionOffline{
		UserId:    s.AuthUserId,
		AuthKeyId: s.authKeyId,
	})
	s.onlineExpired = 0
}

func (s *authSessions) delOnline() {
	logx.Infof("authSessions]]>> delOnline: %s", s)

	s.Dao.StatusClient.StatusSetSessionOffline(context.Background(), &status.TLStatusSetSessionOffline{
		UserId:    s.AuthUserId,
		AuthKeyId: s.authKeyId,
	})
	s.onlineExpired = 0
}

///////////////////////////////////////////////////////////////////////////////////////////////
func (s *authSessions) String() string {
	return fmt.Sprintf("{auth_key_id: %d, user_id: %d, layer: %d}", s.authKeyId, s.AuthUserId, s.Layer)
}

func (s *authSessions) Start() {
	s.running.Set(1)
	s.finish.Add(1)
	go s.rpcRunLoop()
	go s.runLoop()
}

func (s *authSessions) Stop() {
	s.delOnline()
	s.running.Set(0)
	s.rpcQueue.Close()
}

func (s *authSessions) runLoop() {
	defer func() {
		//for _, sess := range s.sessions {
		//	sess.httpQueue.Clear()
		//}

		s.finish.Done()
		close(s.closeChan)
		s.finish.Wait()
	}()

	for s.running.Get() == 1 {
		select {
		case <-s.closeChan:
			// log.Info("runLoop -> To Close ", this.String())
			return

		case sessionMsg, _ := <-s.sessionDataChan:
			switch sessionMsg.(type) {
			case *sessionData:
				s.onSessionData(sessionMsg.(*sessionData))
			case *sessionHttpData:
				s.onSessionHttpData(sessionMsg.(*sessionHttpData))
			case *syncRpcResultData:
				s.onSyncRpcResultData(sessionMsg.(*syncRpcResultData))
			case *syncData:
				s.onSyncData(sessionMsg.(*syncData))
			case *syncSessionData:
				s.onSyncSessionData(sessionMsg.(*syncSessionData))
			case *connData:
				if sessionMsg.(*connData).isNew {
					s.onSessionNew(sessionMsg.(*connData))
				} else {
					s.onSessionClosed(sessionMsg.(*connData))
				}
			default:
				panic("receive invalid type msg")
			}
		case rpcMessages, _ := <-s.rpcDataChan:
			result, _ := rpcMessages.(*rpcApiMessage)
			s.onRpcResult(result)
		// case <-time.After(100 * time.Millisecond):
		case <-time.After(time.Second):
			s.onTimer()
		}
	}

	logx.Info("quit runLoop...")
}

func (s *authSessions) rpcRunLoop() {
	for {
		apiRequest := s.rpcQueue.Pop()
		if apiRequest == nil {
			logx.Info("quit rpcRunLoop...")
			return
		} else {
			request, _ := apiRequest.(*rpcApiMessage)
			// log.Debugf("apiRequests: %s", request.DebugString())
			if s.onRpcRequest(request) {
				s.rpcDataChan <- request
			}
		}
	}
}

func (s *authSessions) onTimer() {
	for _, sess := range s.sessions {
		if (sess.isGeneric && sess.sessionOnline()) ||
			sess.isAndroidPush && sess.sessionOnline() {
			s.setOnline()
		}

		sess.onTimer()
	}

	for _, sess := range s.sessions {
		if !sess.sessionClosed() {
			return
		}
	}

	go func() {
		s.DeleteByAuthKeyId(s.authKeyId)
	}()
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
// client
func (s *authSessions) sessionClientNew(gatewayId string, sessionId int64) error {
	select {
	case s.sessionDataChan <- &connData{true, gatewayId, sessionId}:
		return nil
	}
	return nil
}

func (s *authSessions) sessionDataArrived(gatewayId, clientIp string, sessionId, salt int64, buf []byte) error {
	select {
	case s.sessionDataChan <- &sessionData{gatewayId, clientIp, sessionId, salt, buf}:
		return nil
	}
	return nil
}

func (s *authSessions) sessionHttpDataArrived(gatewayId, clientIp string, sessionId, salt int64, buf []byte, resChan chan interface{}) error {
	select {
	case s.sessionDataChan <- &sessionHttpData{gatewayId, clientIp, sessionId, salt, buf, resChan}:
		return nil
	}
	return nil
}

func (s *authSessions) sessionClientClosed(gatewayId string, sessionId int64) error {
	select {
	case s.sessionDataChan <- &connData{false, gatewayId, sessionId}:
		return nil
	}
	return nil
}

// push
func (s *authSessions) syncRpcResultDataArrived(sessionId, clientMsgId int64, data []byte) error {
	select {
	case s.sessionDataChan <- &syncRpcResultData{sessionId, clientMsgId, data}:
		return nil
	}
	return nil
}

func (s *authSessions) syncSessionDataArrived(sessionId int64, data *messageData) error {
	select {
	case s.sessionDataChan <- &syncSessionData{sessionId, data}:
		return nil
	}
	return nil
}

func (s *authSessions) syncDataArrived(needAndroidPush bool, data *messageData) error {
	select {
	case s.sessionDataChan <- makeSyncData(needAndroidPush, data):
		return nil
	}
	return nil
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
func (s *authSessions) onSessionNew(connMsg *connData) {
	sess, ok := s.sessions[connMsg.sessionId]
	if !ok {
		logx.Infof("onSessionNew - newSession, conn: %s", connMsg.DebugString())
		sess = newSession(connMsg.sessionId, s)
		s.sessions[connMsg.sessionId] = sess
		// sess.onSessionConnNew(connMsg.gatewayId)
	} else {
		sess.sessionState = kSessionStateNew
		logx.Infof("onSessionNew - session found, conn: %s", connMsg.DebugString())
	}

	sess.onSessionConnNew(connMsg.gatewayId)
	//if sess.isGeneric() {
	//
	//}
}

func (s *authSessions) onSessionData(sessionMsg *sessionData) {
	var (
		err error
		// salt, sessionId int64
		message2 = &mtproto.TLMessage2{}
		now      = int32(time.Now().Unix())
	)

	// salt, sessionId, message2, err = ParseFromIncomingMessage(sessionMsg.buf)
	err = message2.Decode(mtproto.NewDecodeBuf(sessionMsg.buf))
	if err != nil {
		// TODO(@benqi): close frontend conn??
		// log.Error(err)
		logx.Errorf("onSessionData - error: {%s}, data: {sessions: %s, gate_id: %d}", err, s, sessionMsg.gatewayId)
		return
	}

	// TODO(@benqi): load onNew
	if s.cacheSalt == nil {
		s.cacheSalt, s.cacheLastSalt, _ = s.Dao.GetOrFetchNewSalt(context.Background(), s.authKeyId)
	} else {
		if now >= s.cacheSalt.GetValidUntil() {
			s.cacheSalt, s.cacheLastSalt, _ = s.Dao.GetOrFetchNewSalt(context.Background(), s.authKeyId)
		}
	}

	if s.cacheSalt == nil {
		logx.Infof("onSessionData - getOrFetchNewSalt nil error, data: {sessions: %s, conn_id: %s}", s, sessionMsg.gatewayId)
		return
	}

	sess, ok := s.sessions[sessionMsg.sessionId]
	if !ok {
		sess = newSession(sessionMsg.sessionId, s)
		s.sessions[sessionMsg.sessionId] = sess
	}

	sess.onSessionConnNew(sessionMsg.gatewayId)
	sess.onSessionMessageData(sessionMsg.gatewayId, sessionMsg.clientIp, sessionMsg.salt, message2)
}

func (s *authSessions) onSessionHttpData(sessionMsg *sessionHttpData) {
	var (
		err error
		// salt, sessionId int64
		message2 = &mtproto.TLMessage2{}
		now      = int32(time.Now().Unix())
	)

	// salt, sessionId, message2, err = ParseFromIncomingMessage(sessionMsg.buf)
	err = message2.Decode(mtproto.NewDecodeBuf(sessionMsg.buf))
	if err != nil {
		// TODO(@benqi): close frontend conn??
		// log.Error(err)
		logx.Errorf("onSessionData - error: {%s}, data: {sessions: %s, gate_id: %d}", err, s, sessionMsg.gatewayId)
		return
	}

	// TODO(@benqi): load onNew
	if s.cacheSalt == nil {
		s.cacheSalt, s.cacheLastSalt, _ = s.Dao.GetOrFetchNewSalt(context.Background(), s.authKeyId)
	} else {
		if now >= s.cacheSalt.GetValidUntil() {
			s.cacheSalt, s.cacheLastSalt, _ = s.Dao.GetOrFetchNewSalt(context.Background(), s.authKeyId)
		}
	}

	if s.cacheSalt == nil {
		logx.Errorf("onSessionData - getOrFetchNewSalt nil error, data: {sessions: %s, conn_id: %s}", s, sessionMsg.gatewayId)
		return
	}

	sess, ok := s.sessions[sessionMsg.sessionId]
	if !ok {
		sess = newSession(sessionMsg.sessionId, s)
		s.sessions[sessionMsg.sessionId] = sess
	}

	sess.isHttp = true
	sess.httpQueue.Push(sessionMsg.resChannel)
	sess.onSessionConnNew(sessionMsg.gatewayId)
	sess.onSessionMessageData(sessionMsg.gatewayId, sessionMsg.clientIp, sessionMsg.salt, message2)
}

func (s *authSessions) onSessionClosed(connMsg *connData) {
	if sess, ok := s.sessions[connMsg.sessionId]; !ok {
		logx.Errorf("onSessionClosed - session conn closed -  conn: %s", connMsg.DebugString())
	} else {
		logx.Infof("onSessionClosed - conn: %s, sess: %s", connMsg.DebugString(), sess)
		sess.onSessionConnClose(connMsg.gatewayId)
	}
}

func (s *authSessions) onSyncRpcResultData(syncMsg *syncRpcResultData) {
	logx.Infof("onSyncRpcResultData - receive data: {sess: %s}",
		s)

	sess, _ := s.sessions[syncMsg.sessionId]
	if sess != nil {
		sess.onSyncRpcResultData(syncMsg.clientMsgId, syncMsg.data)
	}
}

func (s *authSessions) onSyncSessionData(syncMsg *syncSessionData) {
	logx.Infof("onSyncSessionData - receive data: {sess: %s}",
		s)
	sess, _ := s.sessions[syncMsg.sessionId]
	if sess != nil {
		// s.syncQueue.PushBack(syncMsg.data.obj)
		sess.onSyncSessionData(syncMsg.data.obj)
	}
}

func (s *authSessions) onSyncData(syncMsg *syncData) {
	logx.Info("authSessions - ", reflect.TypeOf(syncMsg.data.obj))
	if upds, ok := syncMsg.data.obj.(*mtproto.Updates); ok {
		if upds.PredicateName == mtproto.Predicate_updateAccountResetAuthorization {
			logx.Infof("recv updateAccountResetAuthorization - ", reflect.TypeOf(syncMsg.data.obj))
			if s.AuthUserId != upds.GetUserId() {
				logx.Errorf("upds -- ", upds)
			}
			s.Dao.PutCacheUserId(context.Background(), s.authKeyId, 0)
			s.DeleteByAuthKeyId(s.authKeyId)
			s.AuthUserId = 0
			return
		} else {
			// s.syncQueue.PushBack(upds)
		}
	}

	// s.syncQueue.PushBack(new)
	// s.updates.onUpdatesSyncData(syncMsg)
	//var (
	//	genericSession     *session
	//	androidPushSession *session
	//)

	for _, sess2 := range s.sessions {
		if sess2.isGeneric {
			// genericSession = sess2
			sess2.onSyncData(syncMsg.data.obj)
		}
		if sess2.isAndroidPush {
			if syncMsg.needAndroidPush {
				sess2.onSyncData(nil)
			}
		}
	}
}

func (s *authSessions) onRpcResult(rpcResult *rpcApiMessage) {
	defer func() {
		if err := recover(); err != nil {
			logx.Errorf("tcp_server handle panic: %v\n%s", err, debug.Stack())
		}
	}()

	// log.Debugf("onRpcResult - sessionId: ", rpcResult.sessionId)
	if sess, ok := s.sessions[rpcResult.sessionId]; ok {
		// log.Debugf("onRpcResult result: %s", rpcResult.DebugString())
		sess.onRpcResult(rpcResult)
	} else {
		logx.Errorf("onRpcResult - not found rpcSession by sessionId: ", rpcResult.sessionId)
	}
}

func (s *authSessions) onRpcRequest(request *rpcApiMessage) bool {
	var (
		err       error
		rpcResult mtproto.TLObject
	)

	// 初始化metadata
	rpcMetadata := &metadata.RpcMetadata{
		ServerId:    s.serverId,
		ClientAddr:  request.clientIp,
		AuthId:      s.authKeyId,
		SessionId:   request.sessionId,
		ReceiveTime: time.Now().Unix(),
		UserId:      s.AuthUserId,
		ClientMsgId: request.reqMsgId,
		Layer:       s.Layer,
		Client:      s.getClient(),
		Langpack:    s.getLangpack(),
	}

	// TODO(@benqi): change state.
	switch request.reqMsg.(type) {
	case *mtproto.TLAuthBindTempAuthKey:
		r := request.reqMsg.(*mtproto.TLAuthBindTempAuthKey)
		rpcResult, err = s.Service.Dao.AuthsessionClient.AuthsessionBindTempAuthKey(context.Background(), &authsession.TLAuthsessionBindTempAuthKey{
			PermAuthKeyId:    r.PermAuthKeyId,
			Nonce:            r.Nonce,
			ExpiresAt:        r.ExpiresAt,
			EncryptedMessage: r.EncryptedMessage,
		})
		// request.reqMsg.(*mtproto.TLAuthBindTempAuthKey))
	default:
		rpcResult, err = s.Service.Dao.Invoke(rpcMetadata, request.reqMsg)
	}

	reply := &mtproto.TLRpcResult{
		ReqMsgId: request.reqMsgId,
	}

	if err != nil {
		logx.Error(err.Error())
		if rpcErr, ok := err.(*mtproto.TLRpcError); ok {
			reply.Result = rpcErr
		} else {
			reply.Result = mtproto.NewRpcError(mtproto.StatusInternelServerError)
		}
	} else {
		logx.Infof("invokeRpcRequest - rpc_result: {%s}\n", reflect.TypeOf(rpcResult))
		reply.Result = rpcResult
	}

	request.rpcResult = reply

	if _, ok := request.reqMsg.(*mtproto.TLAuthLogOut3E72BA19); ok {
		logx.Infof("authLogOut - %#v", rpcMetadata)
		s.Dao.PutCacheUserId(context.Background(), s.authKeyId, 0)
	}
	if _, ok := request.reqMsg.(*mtproto.TLAuthLogOut5717DA40); ok {
		logx.Infof("authLogOut - %#v", rpcMetadata)
		s.Dao.PutCacheUserId(context.Background(), s.authKeyId, 0)
	}
	return true
}
