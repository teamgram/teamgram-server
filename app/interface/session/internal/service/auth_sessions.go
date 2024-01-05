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
	"github.com/zeromicro/go-zero/core/threading"

	"github.com/zeromicro/go-zero/core/contextx"
	"github.com/zeromicro/go-zero/core/logx"
	status2 "google.golang.org/grpc/status"
)

type rpcApiMessage struct {
	sessionId int64
	clientIp  string
	reqMsgId  int64
	reqMsg    mtproto.TLObject
	rpcResult *mtproto.TLRpcResult
}

func (m *rpcApiMessage) MoveRpcResult() *mtproto.TLRpcResult {
	v := m.rpcResult
	m.rpcResult = nil
	return v
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

// /////////////////////////////////////////////////////////////////////////////////
type sessionData struct {
	gatewayId string
	clientIp  string
	sessionId int64
	salt      int64
	buf       []byte
}

type sessionDataCtx struct {
	ctx         context.Context
	sessionData sessionData
}

type sessionHttpData struct {
	gatewayId  string
	clientIp   string
	sessionId  int64
	salt       int64
	buf        []byte
	resChannel chan interface{}
}

type sessionHttpDataCtx struct {
	ctx             context.Context
	sessionHttpData sessionHttpData
}

type syncRpcResultData struct {
	sessionId   int64
	clientMsgId int64
	data        []byte
}

type syncRpcResultDataCtx struct {
	ctx               context.Context
	syncRpcResultData syncRpcResultData
}

type syncSessionData struct {
	sessionId int64
	data      *messageData
}

type syncSessionDataCtx struct {
	ctx             context.Context
	syncSessionData syncSessionData
}

type syncData struct {
	needAndroidPush bool
	data            *messageData
}

type syncDataCtx struct {
	ctx      context.Context
	syncData syncData
}

//func makeSyncData(needAndroidPush bool, data *messageData) *syncData {
//	return &syncData{
//		needAndroidPush: needAndroidPush,
//		data:            data,
//	}
//}

type connData struct {
	isNew     bool
	gatewayId string
	sessionId int64
}

type connDataCtx struct {
	ctx      context.Context
	connData connData
}

func (c *connData) DebugString() string {
	return fmt.Sprintf("{isNew: %d, gatewayId: %s, sessionId: %d}", c.isNew, c.gatewayId, c.sessionId)
}

// /////////////////////////////////////////////////////////////////////////////////
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
	permAuthKeyId   int64
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

func (s *authSessions) getAuthKeyId(ctx context.Context) int64 {
	return s.authKeyId
}

func (s *authSessions) getTempAuthKeyId(ctx context.Context) int64 {
	return s.authKeyId
}

func (s *authSessions) getPermAuthKeyId(ctx context.Context) int64 {
	if s.permAuthKeyId != 0 {
		return s.permAuthKeyId
	}
	s.permAuthKeyId = s.Dao.GetCachePermAuthKeyId(context.Background(), s.authKeyId)
	return s.permAuthKeyId
}

func (s *authSessions) setPermAuthKeyId(ctx context.Context, kId int64) {
	s.permAuthKeyId = kId
	if kId != 0 {
		s.permAuthKeyId = kId
		s.Dao.PutCachePermAuthKeyId(context.Background(), s.authKeyId, kId)
	}
}

func (s *authSessions) getUserId(ctx context.Context) int64 {
	return s.AuthUserId
}

func (s *authSessions) setUserId(ctx context.Context, userId int64) {
	s.AuthUserId = userId
	s.onBindUser(userId)
}

func (s *authSessions) getCacheSalt(ctx context.Context) *mtproto.TLFutureSalt {
	return s.cacheSalt
}

func (s *authSessions) getLayer(ctx context.Context) int32 {
	if s.Layer == 0 {
		s.Layer, _ = s.Dao.GetCacheApiLayer(ctx, s.authKeyId)
	}
	return s.Layer
}

func (s *authSessions) setLayer(ctx context.Context, layer int32) {
	if layer != 0 {
		s.Layer = layer
		s.Dao.PutCacheApiLayer(ctx, s.authKeyId, layer)
	}
}

func (s *authSessions) getClient(ctx context.Context) string {
	if s.Client == "" {
		s.Client = s.Dao.GetCacheClient(ctx, s.authKeyId)
	}
	return s.Client
}

func (s *authSessions) setClient(ctx context.Context, c string) {
	if c != "" {
		s.Client = c
		s.Dao.PutCacheClient(ctx, s.authKeyId, c)
	}
}

func (s *authSessions) getLangpack(ctx context.Context) string {
	if s.Langpack == "" {
		s.Langpack = s.Dao.GetCacheLangpack(ctx, s.authKeyId)
	}
	return s.Langpack
}

func (s *authSessions) setLangpack(ctx context.Context, c string) {
	if c != "" {
		s.Langpack = c
		s.Dao.PutCacheLangpack(ctx, s.authKeyId, c)
	}
}

func (s *authSessions) destroySession(ctx context.Context, sessionId int64) bool {
	// TODO(@benqi):
	if _, ok := s.sessions[sessionId]; ok {
		// s.updates.onGenericSessionClose(sess)
		delete(s.sessions, sessionId)
	} else {
		//
	}
	return true
}

func (s *authSessions) sendToRpcQueue(ctx context.Context, rpcMessage *rpcApiMessage) {
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

func (s *authSessions) onBindPushSessionId(ctx context.Context, sessionId int64) {
	if s.pushSessionId == 0 {
		s.pushSessionId = sessionId
	}
	sess, _ := s.sessions[sessionId]
	if sess != nil {
		sess.isAndroidPush = true
		sess.cb.setOnline(ctx)
	}
}

func (s *authSessions) onBindLayer(layer int32) {
	s.Layer = layer
}

func (s *authSessions) setOnline(ctx context.Context) {
	//setOnlineTTL(s.AuthUserId, s.authKeyId, getServerID(), s.Layer, 60)
	date := time.Now().Unix()
	if (s.onlineExpired == 0 || date > s.onlineExpired-kPingAddTimeout) && s.AuthUserId != 0 {
		logx.Infof("DEBUG] setOnline - set online: (date: %d, userId:%d, onlineExpired: %d, authKeyId: %d)",
			date,
			s.AuthUserId,
			s.onlineExpired,
			s.authKeyId)

		// s.setOnlineTTL(s.AuthUserId, s.authKeyId, getServerID(), s.Layer, 60)
		s.Dao.StatusClient.StatusSetSessionOnline(
			context.Background(),
			&status.TLStatusSetSessionOnline{
				UserId: s.AuthUserId,
				Session: &status.SessionEntry{
					UserId:        s.AuthUserId,
					AuthKeyId:     s.authKeyId,
					Gateway:       s.serverId,
					Expired:       date + 60,
					Layer:         s.getLayer(ctx),
					PermAuthKeyId: s.getPermAuthKeyId(ctx),
					Client:        s.getClient(ctx),
				},
			})
		s.onlineExpired = date + 60
	} else {
		//logx.Infof("DEBUG] setOnline - not set online: (date: %d, onlineExpired: %d, AuthUserId: %d)",
		//	date,
		//	s.onlineExpired,
		//	s.AuthUserId)
	}
}

func (s *authSessions) trySetOffline(ctx context.Context) {
	for _, sess := range s.sessions {
		if (sess.isGeneric && sess.sessionOnline()) ||
			(sess.isAndroidPush && sess.sessionOnline()) {
			return
		}
	}

	if s.AuthUserId > 0 {
		logx.Infof("authSessions]]>> offline: %s", s)
		s.Dao.StatusClient.StatusSetSessionOffline(context.Background(), &status.TLStatusSetSessionOffline{
			UserId:    s.AuthUserId,
			AuthKeyId: s.authKeyId,
		})
	}
	s.onlineExpired = 0
}

func (s *authSessions) delOnline() {
	if s.AuthUserId > 0 {
		logx.Infof("authSessions]]>> delOnline: %s", s)

		s.Dao.StatusClient.StatusSetSessionOffline(context.Background(), &status.TLStatusSetSessionOffline{
			UserId:    s.AuthUserId,
			AuthKeyId: s.authKeyId,
		})
	}
	s.onlineExpired = 0
}

// /////////////////////////////////////////////////////////////////////////////////////////////
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
		close(s.sessionDataChan)
		close(s.rpcDataChan)
		s.finish.Wait()
	}()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for s.running.Get() == 1 {
		select {
		case <-s.closeChan:
			// log.Info("runLoop -> To Close ", this.String())
			return

		case sessionMsg, _ := <-s.sessionDataChan:
			switch ctxData := sessionMsg.(type) {
			case *sessionDataCtx:
				threading.RunSafe(func() {
					s.onSessionData(ctxData.ctx, &ctxData.sessionData)
				})
			case *sessionHttpDataCtx:
				threading.RunSafe(func() {
					s.onSessionHttpData(ctxData.ctx, &ctxData.sessionHttpData)
				})
			case *syncRpcResultDataCtx:
				threading.RunSafe(func() {
					s.onSyncRpcResultData(ctxData.ctx, &ctxData.syncRpcResultData)
				})
			case *syncDataCtx:
				threading.RunSafe(func() {
					s.onSyncData(ctxData.ctx, &ctxData.syncData)
				})
			case *syncSessionDataCtx:
				threading.RunSafe(func() {
					s.onSyncSessionData(ctxData.ctx, &ctxData.syncSessionData)
				})
			case *connDataCtx:
				threading.RunSafe(func() {
					if ctxData.connData.isNew {
						s.onSessionNew(ctxData.ctx, &ctxData.connData)
					} else {
						s.onSessionClosed(ctxData.ctx, &ctxData.connData)
					}
				})
			default:
				panic("receive invalid type msg")
			}
		case rpcMessages, _ := <-s.rpcDataChan:
			threading.RunSafe(func() {
				result, _ := rpcMessages.(*rpcApiMessage)
				s.onRpcResult(context.Background(), result)
			})
			// case <-time.After(100 * time.Millisecond):
		case <-ticker.C:
			threading.RunSafe(func() {
				s.onTimer(context.Background())
			})
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
			threading.RunSafe(func() {
				// TODO: fix panic
				request, _ := apiRequest.(*rpcApiMessage)
				// log.Debugf("apiRequests: %s", request.DebugString())
				if s.onRpcRequest(context.Background(), request) {
					s.rpcDataChan <- request
				}
			})
		}
	}
}

func (s *authSessions) onTimer(ctx context.Context) {
	for _, sess := range s.sessions {
		if (sess.isGeneric && sess.sessionOnline()) ||
			sess.isAndroidPush && sess.sessionOnline() {
			s.setOnline(ctx)
		}

		sess.onTimer(ctx)
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

// ///////////////////////////////////////////////////////////////////////////////////////////////////
// client
func (s *authSessions) sessionClientNew(ctx context.Context, gatewayId string, sessionId int64) error {
	select {
	case s.sessionDataChan <- &connDataCtx{contextx.ValueOnlyFrom(ctx), connData{true, gatewayId, sessionId}}:
		return nil
	}
	return nil
}

func (s *authSessions) sessionDataArrived(ctx context.Context, gatewayId, clientIp string, sessionId, salt int64, buf []byte) error {
	select {
	case s.sessionDataChan <- &sessionDataCtx{contextx.ValueOnlyFrom(ctx), sessionData{gatewayId, clientIp, sessionId, salt, buf}}:
		return nil
	}
	return nil
}

func (s *authSessions) sessionHttpDataArrived(ctx context.Context, gatewayId, clientIp string, sessionId, salt int64, buf []byte, resChan chan interface{}) error {
	select {
	case s.sessionDataChan <- &sessionHttpDataCtx{contextx.ValueOnlyFrom(ctx), sessionHttpData{gatewayId, clientIp, sessionId, salt, buf, resChan}}:
		return nil
	}
	return nil
}

func (s *authSessions) sessionClientClosed(ctx context.Context, gatewayId string, sessionId int64) error {
	select {
	case s.sessionDataChan <- &connDataCtx{contextx.ValueOnlyFrom(ctx), connData{false, gatewayId, sessionId}}:
		return nil
	}
	return nil
}

// push
func (s *authSessions) syncRpcResultDataArrived(ctx context.Context, sessionId, clientMsgId int64, data []byte) error {
	select {
	case s.sessionDataChan <- &syncRpcResultDataCtx{contextx.ValueOnlyFrom(ctx), syncRpcResultData{sessionId, clientMsgId, data}}:
		return nil
	}
	return nil
}

func (s *authSessions) syncSessionDataArrived(ctx context.Context, sessionId int64, data *messageData) error {
	select {
	case s.sessionDataChan <- &syncSessionDataCtx{contextx.ValueOnlyFrom(ctx), syncSessionData{sessionId, data}}:
		return nil
	}
	return nil
}

func (s *authSessions) syncDataArrived(ctx context.Context, needAndroidPush bool, data *messageData) error {
	select {
	case s.sessionDataChan <- &syncDataCtx{contextx.ValueOnlyFrom(ctx), syncData{needAndroidPush, data}}:
		return nil
	}
	return nil
}

// ///////////////////////////////////////////////////////////////////////////////////////////////////
func (s *authSessions) onSessionNew(ctx context.Context, connMsg *connData) {
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

	sess.onSessionConnNew(ctx, connMsg.gatewayId)
}

func (s *authSessions) onSessionData(ctx context.Context, sessionMsg *sessionData) {
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
		logx.WithContext(ctx).Errorf("onSessionData - error: {%s}, data: {sessions: %s, gate_id: %d}", err, s, sessionMsg.gatewayId)
		return
	}

	// TODO(@benqi): load onNew
	if s.cacheSalt == nil {
		s.cacheSalt, s.cacheLastSalt, _ = s.Dao.GetOrFetchNewSalt(ctx, s.authKeyId)
	} else {
		if now >= s.cacheSalt.GetValidUntil() {
			s.cacheSalt, s.cacheLastSalt, _ = s.Dao.GetOrFetchNewSalt(ctx, s.authKeyId)
		}
	}

	if s.cacheSalt == nil {
		logx.WithContext(ctx).Infof("onSessionData - getOrFetchNewSalt nil error, data: {sessions: %s, conn_id: %s}", s, sessionMsg.gatewayId)
		return
	}

	sess, ok := s.sessions[sessionMsg.sessionId]
	if !ok {
		sess = newSession(sessionMsg.sessionId, s)
		s.sessions[sessionMsg.sessionId] = sess
	}

	sess.onSessionConnNew(ctx, sessionMsg.gatewayId)
	sess.onSessionMessageData(ctx, sessionMsg.gatewayId, sessionMsg.clientIp, sessionMsg.salt, message2)
}

func (s *authSessions) onSessionHttpData(ctx context.Context, sessionMsg *sessionHttpData) {
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
	sess.onSessionConnNew(ctx, sessionMsg.gatewayId)
	sess.onSessionMessageData(ctx, sessionMsg.gatewayId, sessionMsg.clientIp, sessionMsg.salt, message2)
}

func (s *authSessions) onSessionClosed(ctx context.Context, connMsg *connData) {
	if sess, ok := s.sessions[connMsg.sessionId]; !ok {
		logx.WithContext(ctx).Errorf("onSessionClosed - session conn closed -  conn: %s", connMsg.DebugString())
	} else {
		logx.WithContext(ctx).Infof("onSessionClosed - conn: %s, sess: %s", connMsg.DebugString(), sess)
		sess.onSessionConnClose(ctx, connMsg.gatewayId)
	}
}

func (s *authSessions) onSyncRpcResultData(ctx context.Context, syncMsg *syncRpcResultData) {
	logx.WithContext(ctx).Infof("onSyncRpcResultData - receive data: {sess: %s}",
		s)

	sess, _ := s.sessions[syncMsg.sessionId]
	if sess != nil {
		sess.onSyncRpcResultData(ctx, syncMsg.clientMsgId, syncMsg.data)
	}
}

func (s *authSessions) onSyncSessionData(ctx context.Context, syncMsg *syncSessionData) {
	logx.WithContext(ctx).Infof("onSyncSessionData - receive data: {sess: %s}",
		s)
	sess, _ := s.sessions[syncMsg.sessionId]
	if sess != nil {
		// s.syncQueue.PushBack(syncMsg.data.obj)
		sess.onSyncSessionData(ctx, syncMsg.data.obj)
	}
}

func (s *authSessions) onSyncData(ctx context.Context, syncMsg *syncData) {
	logx.WithContext(ctx).Info("authSessions - ", reflect.TypeOf(syncMsg.data.obj))
	if upds, ok := syncMsg.data.obj.(*mtproto.Updates); ok {
		if upds.PredicateName == mtproto.Predicate_updateAccountResetAuthorization {
			logx.WithContext(ctx).Infof("recv updateAccountResetAuthorization - ", reflect.TypeOf(syncMsg.data.obj))
			if s.AuthUserId != upds.GetUserId() {
				logx.WithContext(ctx).Errorf("upds -- ", upds)
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
			sess2.onSyncData(ctx, syncMsg.data.obj)
		}
		if sess2.isAndroidPush {
			if syncMsg.needAndroidPush {
				sess2.onSyncData(ctx, nil)
			}
		}
	}
}

func (s *authSessions) onRpcResult(ctx context.Context, rpcResult *rpcApiMessage) {
	defer func() {
		if err := recover(); err != nil {
			logx.Errorf("tcp_server handle panic: %v\n%s", err, debug.Stack())
		}
	}()

	// log.Debugf("onRpcResult - sessionId: ", rpcResult.sessionId)
	if sess, ok := s.sessions[rpcResult.sessionId]; ok {
		// log.Debugf("onRpcResult result: %s", rpcResult.DebugString())
		sess.onRpcResult(ctx, rpcResult)
	} else {
		logx.Errorf("onRpcResult - not found rpcSession by sessionId: ", rpcResult.sessionId)
	}
}

func (s *authSessions) onRpcRequest(ctx context.Context, request *rpcApiMessage) bool {
	var (
		err       error
		rpcResult mtproto.TLObject
	)

	// 初始化metadata
	rpcMetadata := &metadata.RpcMetadata{
		ServerId:      s.serverId,
		ClientAddr:    request.clientIp,
		AuthId:        s.authKeyId,
		SessionId:     request.sessionId,
		ReceiveTime:   time.Now().Unix(),
		UserId:        s.AuthUserId,
		ClientMsgId:   request.reqMsgId,
		Layer:         s.Layer,
		Client:        s.getClient(ctx),
		Langpack:      s.getLangpack(ctx),
		PermAuthKeyId: s.getPermAuthKeyId(ctx),
	}

	// TODO(@benqi): change state.
	switch request.reqMsg.(type) {
	case *mtproto.TLAuthBindTempAuthKey:
		r := request.reqMsg.(*mtproto.TLAuthBindTempAuthKey)
		rpcResult, err = s.Service.Dao.AuthsessionClient.AuthsessionBindTempAuthKey(
			context.Background(),
			&authsession.TLAuthsessionBindTempAuthKey{
				PermAuthKeyId:    r.PermAuthKeyId,
				Nonce:            r.Nonce,
				ExpiresAt:        r.ExpiresAt,
				EncryptedMessage: r.EncryptedMessage,
			})
		if err != nil {
			s2 := status2.Convert(err)
			if s2.Message() == "ENCRYPTED_MESSAGE_INVALID" {
				s.Dao.PutCacheUserId(context.Background(), s.authKeyId, 0)
				s.DeleteByAuthKeyId(s.authKeyId)
				s.AuthUserId = 0
			}
			err = mtproto.NewRpcError(s2)
		} else {
			s.Dao.PutCachePermAuthKeyId(context.Background(), s.authKeyId, r.PermAuthKeyId)
		}
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
			reply.Result = mtproto.NewRpcError(mtproto.StatusInternalServerError)
		}
	} else {
		logx.Infof("invokeRpcRequest - rpc_result: {%s}\n", reflect.TypeOf(rpcResult))
		reply.Result = rpcResult
	}

	request.rpcResult = reply

	if _, ok := request.reqMsg.(*mtproto.TLAuthLogOut); ok {
		logx.Infof("authLogOut - %#v", rpcMetadata)
		s.Dao.PutCacheUserId(context.Background(), s.authKeyId, 0)
	}
	return true
}
