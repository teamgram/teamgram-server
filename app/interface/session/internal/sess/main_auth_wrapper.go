// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package sess

import (
	"context"
	"fmt"
	"math"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/teamgram/marmota/pkg/queue2"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/proto/mtproto/rpc/metadata"
	"github.com/teamgram/teamgram-server/app/interface/session/internal/dao"
	"github.com/teamgram/teamgram-server/app/service/authsession/authsession"
	"github.com/teamgram/teamgram-server/app/service/status/status"

	"github.com/zeromicro/go-zero/core/contextx"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/syncx"
	"github.com/zeromicro/go-zero/core/threading"
)

type SessionList struct {
	authId        int64
	authType      int
	state         int
	cacheSalt     *mtproto.TLFutureSalt
	cacheLastSalt *mtproto.TLFutureSalt
	sessions      map[int64]*session
	cb            *MainAuthWrapper
}

func newSessionList(kType int, cb *MainAuthWrapper) *SessionList {
	return &SessionList{
		authId:        0,
		authType:      kType,
		state:         0,
		cacheSalt:     nil,
		cacheLastSalt: nil,
		sessions:      make(map[int64]*session),
		cb:            cb,
	}
}

func (s *SessionList) Reset(authId int64) (lastAuthId int64) {
	lastAuthId = s.authId

	s.authId = authId
	s.state = 0
	s.cacheSalt = nil
	s.cacheLastSalt = nil
	s.sessions = make(map[int64]*session)

	return
}

func (s *SessionList) destroySession(sessionId int64) bool {
	// TODO(@benqi):
	if _, ok := s.sessions[sessionId]; ok {
		// s.updates.onGenericSessionClose(sess)
		delete(s.sessions, sessionId)
	} else {
		//
	}
	return true
}

func (s *SessionList) changeAuthState(state int) {
	s.state = state
}

type MainAuthWrapper struct {
	authKeyId            int64
	state                int
	AuthUserId           int64 // 不为0，则signIn
	client               *authsession.ClientSession
	pushSessionId        int64
	mainAuth             *SessionList
	tempAuth             *SessionList
	mediaTempAuth        *SessionList
	mainUpdatesSession   *session
	androidPushSession   *session
	closeChan            chan struct{}
	sessionDataChan      chan interface{} // receive from client
	rpcDataChan          chan interface{}
	rpcQueue             *queue2.SyncQueue
	finish               sync.WaitGroup
	running              *syncx.AtomicBool
	onlineExpired        int64
	clientType           int
	nextNotifyId         int64
	nextPushId           int64
	cb                   *MainAuthWrapperManager
	tmpRpcApiMessageList []*rpcApiMessage
}

func NewMainAuthWrapper(mainAuthKeyId int64, authUserId int64, state int, client *authsession.ClientSession, androidPushId int64, cb *MainAuthWrapperManager) *MainAuthWrapper {
	mainAuth := &MainAuthWrapper{
		authKeyId:            mainAuthKeyId,
		state:                state,
		AuthUserId:           authUserId,
		client:               client,
		pushSessionId:        androidPushId,
		mainAuth:             nil,
		tempAuth:             nil,
		mediaTempAuth:        nil,
		mainUpdatesSession:   nil,
		androidPushSession:   nil,
		clientType:           clientUnknown,
		nextPushId:           0,
		nextNotifyId:         math.MaxInt32,
		closeChan:            make(chan struct{}),
		sessionDataChan:      make(chan interface{}, 1024),
		rpcDataChan:          make(chan interface{}, 1024),
		rpcQueue:             queue2.NewSyncQueue(),
		finish:               sync.WaitGroup{},
		running:              syncx.NewAtomicBool(),
		cb:                   cb,
		tmpRpcApiMessageList: make([]*rpcApiMessage, 0),
	}
	mainAuth.tempAuth = newSessionList(mtproto.AuthKeyTypeTemp, mainAuth)
	mainAuth.mediaTempAuth = newSessionList(mtproto.AuthKeyTypeMediaTemp, mainAuth)
	mainAuth.mainAuth = newSessionList(mtproto.AuthKeyTypePerm, mainAuth)

	mainAuth.Start()
	return mainAuth
}

func (m *MainAuthWrapper) changeAuthState(ctx context.Context, state int, stateData interface{}) {
	m.state = state

	switch state {
	case mtproto.AuthStateUnknown:
		// m.delOnline(ctx)
		// m.AuthUserId = 0
		m.cb.DeleteByAuthKeyId(m.authKeyId)
		m.Stop()
	case mtproto.AuthStateLogout:
		// m.delOnline(ctx)
		// m.AuthUserId = 0
		m.cb.DeleteByAuthKeyId(m.authKeyId)
		m.Stop()
	case mtproto.AuthStateDeleted:
		// m.delOnline(ctx)
		// m.AuthUserId = 0
		m.cb.DeleteByAuthKeyId(m.authKeyId)
		m.Stop()
	case mtproto.AuthStateNeedPassword:
		m.AuthUserId = stateData.(int64)
	case mtproto.AuthStateNormal:
		m.AuthUserId = stateData.(int64)
	default:
		m.AuthUserId = 0
	}
}

func (m *MainAuthWrapper) resetAuth(kType int, authId int64) (lastAuthId int64) {
	switch kType {
	case mtproto.AuthKeyTypeTemp:
		lastAuthId = m.tempAuth.Reset(authId)
		m.androidPushSession = nil
		m.mainUpdatesSession = nil
	case mtproto.AuthKeyTypeMediaTemp:
		lastAuthId = m.mediaTempAuth.Reset(authId)
	default:
		lastAuthId = m.mainAuth.Reset(authId)
		m.androidPushSession = nil
		m.mainUpdatesSession = nil
	}

	// TODO: notify

	return
}

func (m *MainAuthWrapper) setOnline(ctx context.Context) {
	//setOnlineTTL(s.AuthUserId, s.authKeyId, getServerID(), s.Layer, 60)
	date := time.Now().Unix()
	if (m.onlineExpired == 0 || date > m.onlineExpired-kPingAddTimeout) && m.AuthUserId != 0 {
		logx.WithContext(ctx).Infof("DEBUG] setOnline - set online: (date: %d, userId:%d, onlineExpired: %d, authKeyId: %d)",
			date,
			m.AuthUserId,
			m.onlineExpired,
			m.authKeyId)

		// s.setOnlineTTL(s.AuthUserId, s.authKeyId, getServerID(), s.Layer, 60)
		m.cb.Dao.StatusClient.StatusSetSessionOnline(
			ctx,
			&status.TLStatusSetSessionOnline{
				UserId: m.AuthUserId,
				Session: &status.SessionEntry{
					UserId:        m.AuthUserId,
					AuthKeyId:     m.authKeyId,
					Gateway:       m.cb.Dao.MyServerId,
					Expired:       date + 60,
					Layer:         m.Layer(),
					PermAuthKeyId: m.authKeyId,
					Client:        m.ClientName(),
				},
			})
		m.onlineExpired = date + 60
	} else {
		//logx.WithContext(ctx).Infof("DEBUG] setOnline - not set online: (date: %d, onlineExpired: %d, AuthUserId: %d)",
		//	date,
		//	s.onlineExpired,
		//	s.AuthUserId)
	}
}

func (m *MainAuthWrapper) trySetOffline(ctx context.Context) {
	if m.androidPushSession != nil && m.androidPushSession.sessionOnline() {
		return
	}
	if m.mainUpdatesSession != nil && m.mainUpdatesSession.sessionOnline() {
		return
	}

	if m.AuthUserId > 0 {
		logx.WithContext(ctx).Infof("authSessions]]>> offline: %s", m)
		m.cb.Dao.StatusClient.StatusSetSessionOffline(ctx, &status.TLStatusSetSessionOffline{
			UserId:    m.AuthUserId,
			AuthKeyId: m.authKeyId,
		})
	}
	m.onlineExpired = 0
}

func (m *MainAuthWrapper) delOnline(ctx context.Context) {
	if m.AuthUserId > 0 {
		logx.Infof("authSessions]]>> delOnline: %s", m)

		m.cb.Dao.StatusClient.StatusSetSessionOffline(ctx, &status.TLStatusSetSessionOffline{
			UserId:    m.AuthUserId,
			AuthKeyId: m.authKeyId,
		})
	}
	m.onlineExpired = 0
}

func (m *MainAuthWrapper) getNextNotifyId() (id int64) {
	id = m.nextNotifyId
	m.nextNotifyId--
	return
}

func (m *MainAuthWrapper) getNextPushId() (id int64) {
	id = m.nextPushId
	m.nextPushId++
	return
}

func (m *MainAuthWrapper) getClient() *authsession.ClientSession {
	if m.client != nil {
		return m.client
	}

	return nil
}

func (m *MainAuthWrapper) onUpdateLayer(ctx context.Context, layer int32) {
	if layer == 0 || m.Layer() == layer {
		return
	}

	if m.client == nil {
		m.client = authsession.MakeTLClientSession(&authsession.ClientSession{
			AuthKeyId:      m.authKeyId,
			Ip:             "",
			Layer:          layer,
			ApiId:          0,
			DeviceModel:    "",
			SystemVersion:  "",
			AppVersion:     "",
			SystemLangCode: "",
			LangPack:       "",
			LangCode:       "",
			Proxy:          "",
			Params:         "",
		}).To_ClientSession()
	} else {
		m.client.Layer = layer
	}

	m.cb.Dao.AuthsessionClient.AuthsessionSetLayer(ctx, &authsession.TLAuthsessionSetLayer{
		AuthKeyId: m.authKeyId,
		Ip:        "",
		Layer:     layer,
	})
}

func (m *MainAuthWrapper) onUpdateInitConnection(ctx context.Context, clientIp string, initConnection *mtproto.TLInitConnection) {
	if initConnection == nil {
		return
	}

	if m.state < mtproto.AuthStateUnauthorized {
		m.state = mtproto.AuthStateUnauthorized
	}

	clientNeedUpdate := false

	proxy, _ := jsonx.MarshalToString(initConnection.Proxy)
	params, _ := jsonx.MarshalToString(initConnection.Params)

	if m.client == nil {
		m.client = authsession.MakeTLClientSession(&authsession.ClientSession{
			AuthKeyId:      m.authKeyId,
			Ip:             clientIp,
			Layer:          0,
			ApiId:          initConnection.ApiId,
			DeviceModel:    initConnection.DeviceModel,
			SystemVersion:  initConnection.SystemVersion,
			AppVersion:     initConnection.AppVersion,
			SystemLangCode: initConnection.SystemLangCode,
			LangPack:       initConnection.LangPack,
			LangCode:       initConnection.LangPack,
			Proxy:          proxy,
			Params:         params,
		}).To_ClientSession()

		clientNeedUpdate = true
	} else {
		if m.client.Ip != clientIp {
			m.client.Ip = clientIp
			clientNeedUpdate = true
		}
		if m.client.ApiId != initConnection.ApiId {
			m.client.ApiId = initConnection.ApiId
			clientNeedUpdate = true
		}
		if m.client.DeviceModel != initConnection.DeviceModel {
			m.client.DeviceModel = initConnection.DeviceModel
			clientNeedUpdate = true
		}
		if m.client.SystemVersion != initConnection.SystemVersion {
			m.client.SystemVersion = initConnection.SystemVersion
			clientNeedUpdate = true
		}
		if m.client.AppVersion != initConnection.AppVersion {
			m.client.AppVersion = initConnection.AppVersion
			clientNeedUpdate = true
		}
		if m.client.SystemLangCode != initConnection.SystemLangCode {
			m.client.SystemLangCode = initConnection.SystemLangCode
			clientNeedUpdate = true
		}
		if m.client.LangPack != initConnection.LangPack {
			m.client.LangPack = initConnection.LangPack
			clientNeedUpdate = true
		}
		if m.client.LangCode != initConnection.LangCode {
			m.client.LangCode = initConnection.LangCode
			clientNeedUpdate = true
		}
		if m.client.Proxy != proxy {
			m.client.Proxy = proxy
			// clientNeedUpdate = true
		}
		if m.client.Params != params {
			m.client.Params = params
			// clientNeedUpdate = true
		}
	}

	if clientNeedUpdate {
		m.cb.Dao.AuthsessionClient.AuthsessionSetInitConnection(ctx, &authsession.TLAuthsessionSetInitConnection{
			AuthKeyId:      m.authKeyId,
			Ip:             clientIp,
			ApiId:          m.ApiId(),
			DeviceModel:    m.DeviceModel(),
			SystemVersion:  m.SystemVersion(),
			AppVersion:     m.AppVersion(),
			SystemLangCode: m.SystemLangCode(),
			LangPack:       m.LangPack(),
			LangCode:       m.LangCode(),
			Proxy:          m.Proxy(),
			Params:         m.Params(),
		})
	}
}

func (m *MainAuthWrapper) onBindPushSessionId(ctx context.Context, sList *SessionList, sessionId int64) {
	if m.pushSessionId == 0 {
		m.pushSessionId = sessionId
	}

	sess := m.androidPushSession
	if sess == nil {
		sess, _ = sList.sessions[sessionId]
	}
	if sess == nil {
		logx.WithContext(ctx).Errorf("onBindPushSessionId - unknown error(auth_key_id: %d, session_id: %d)", sList.authId, sessionId)
		return
	} else {
		sess.isAndroidPush = true
		m.androidPushSession = sess
	}

	m.setOnline(ctx)
}

func (m *MainAuthWrapper) onSetMainUpdatesSession(ctx context.Context, sess *session) {
	if !sess.isGeneric {
		sess.isGeneric = true
	}
	if m.mainUpdatesSession == nil || m.mainUpdatesSession.sessionId != sess.sessionId {
		m.mainUpdatesSession = sess
	}
	m.setOnline(ctx)
}

func (m *MainAuthWrapper) Layer() int32 {
	return m.getClient().GetLayer()
}

func (m *MainAuthWrapper) ApiId() int32 {
	return m.getClient().GetApiId()
}

func (m *MainAuthWrapper) DeviceModel() string {
	return m.getClient().GetDeviceModel()
}

func (m *MainAuthWrapper) SystemVersion() string {
	return m.getClient().GetSystemVersion()
}

func (m *MainAuthWrapper) AppVersion() string {
	return m.getClient().GetAppVersion()
}

func (m *MainAuthWrapper) SystemLangCode() string {
	return m.getClient().GetSystemLangCode()
}

func (m *MainAuthWrapper) LangPack() string {
	return m.getClient().GetLangPack()
}

func (m *MainAuthWrapper) LangCode() string {
	return m.getClient().GetLangCode()
}

func (m *MainAuthWrapper) ClientIp() string {
	return m.getClient().GetIp()
}

func (m *MainAuthWrapper) Proxy() string {
	return m.getClient().GetProxy()
}

func (m *MainAuthWrapper) Params() string {
	return m.getClient().GetParams()
}

func (m *MainAuthWrapper) ClientName() string {
	c := m.getClient().GetLangPack()

	if c == "android" {
		if strings.Index(m.getClient().GetAppVersion(), "TDLib") >= 0 {
			c = "react"
		}
	} else if c == "" {
		if strings.HasSuffix(m.getClient().GetAppVersion(), " A") {
			c = "weba"
		} else if strings.HasSuffix(m.getClient().GetAppVersion(), " Z") {
			c = "weba"
		}
	}

	return c
}

func (m *MainAuthWrapper) getSessionList(kType int) (sList *SessionList) {
	switch kType {
	case mtproto.AuthKeyTypeTemp:
		sList = m.tempAuth
	case mtproto.AuthKeyTypeMediaTemp:
		sList = m.mediaTempAuth
	default:
		sList = m.mainAuth
	}

	return
}

func (m *MainAuthWrapper) getSessionListById(authId int64) (sList *SessionList) {
	switch authId {
	case m.tempAuth.authId:
		sList = m.tempAuth
	case m.mediaTempAuth.authId:
		sList = m.mediaTempAuth
	case m.mainAuth.authId:
		sList = m.mainAuth
	}

	return
}

func (m *MainAuthWrapper) String() string {
	return fmt.Sprintf("{auth_key_id: %d, user_id: %d, layer: %d}", m.authKeyId, m.AuthUserId, m.Layer)
}

func (m *MainAuthWrapper) Start() {
	m.running.Set(true)
	m.finish.Add(1)
	go m.rpcRunLoop()
	go m.runLoop()
}

func (m *MainAuthWrapper) Stop() {
	m.running.Set(false)
	// m.rpcQueue.Close()
}

func (m *MainAuthWrapper) runLoop() {
	defer func() {
		if (m.mainUpdatesSession != nil && m.mainUpdatesSession.sessionOnline()) ||
			(m.androidPushSession != nil && m.androidPushSession.sessionOnline()) {
			m.delOnline(context.Background())
		}
		m.finish.Done()
		m.rpcQueue.Close()
		close(m.closeChan)
		close(m.sessionDataChan)
		m.finish.Wait()
	}()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for m.running.True() {
		select {
		case <-m.closeChan:
			// log.Info("runLoop -> To Close ", this.String())
			return

		case sessionMsg, _ := <-m.sessionDataChan:
			switch ctxData := sessionMsg.(type) {
			case *sessionDataCtx:
				threading.RunSafe(func() {
					m.onSessionData(ctxData.ctx, &ctxData.sessionData)
				})
			case *sessionHttpDataCtx:
				threading.RunSafe(func() {
					m.onSessionHttpData(ctxData.ctx, &ctxData.sessionHttpData)
				})
			case *syncRpcResultDataCtx:
				threading.RunSafe(func() {
					m.onSyncRpcResultData(ctxData.ctx, &ctxData.syncRpcResultData)
				})
			case *syncDataCtx:
				threading.RunSafe(func() {
					m.onSyncData(ctxData.ctx, &ctxData.syncData)
				})
			case *syncSessionDataCtx:
				threading.RunSafe(func() {
					m.onSyncSessionData(ctxData.ctx, &ctxData.syncSessionData)
				})
			case *connDataCtx:
				threading.RunSafe(func() {
					if ctxData.isNew {
						m.onSessionNew(ctxData.ctx, &ctxData.connData)
					} else {
						m.onSessionClosed(ctxData.ctx, &ctxData.connData)
					}
				})
			default:
				panic("receive invalid type msg")
			}
		case rpcMessages, _ := <-m.rpcDataChan:
			threading.RunSafe(func() {
				rpcResult, _ := rpcMessages.(*rpcApiMessage)
				_ = rpcResult
				if sess, ok := rpcResult.sessList.sessions[rpcResult.sessionId]; ok {
					// log.Debugf("onRpcResult result: %s", rpcResult)
					sess.onRpcResult(rpcResult.ctx, rpcResult)
				} else {
					logx.WithContext(rpcResult.ctx).Errorf("onRpcResult - not found rpcSession by sessionId: %d", rpcResult.sessionId)
				}
			})
		case <-ticker.C:
			threading.RunSafe(func() {
				m.onTimer(context.Background())
			})
		}
	}

	logx.Info("quit runLoop...")
}

func (m *MainAuthWrapper) rpcRunLoop() {
	defer func() {
		close(m.rpcDataChan)
	}()
	for {
		apiRequest := m.rpcQueue.Pop()
		if apiRequest == nil {
			logx.Info("quit rpcRunLoop...")
			return
		} else {
			threading.RunSafe(func() {
				for _, request := range apiRequest.([]*rpcApiMessage) {
					doRpcRequest(request.ctx,
						m.cb.Dao,
						&metadata.RpcMetadata{
							ServerId:      m.cb.Dao.MyServerId,
							ClientAddr:    request.clientIp,
							AuthId:        request.sessList.authId,
							SessionId:     request.sessionId,
							ReceiveTime:   time.Now().Unix(),
							UserId:        m.AuthUserId,
							ClientMsgId:   request.reqMsgId,
							Layer:         m.Layer(),
							Client:        m.ClientName(),
							Langpack:      m.LangPack(),
							PermAuthKeyId: m.authKeyId,
						},
						request)
					m.rpcDataChan <- request
				}
			})
		}
	}
}

func (m *MainAuthWrapper) sendToRpcQueue(ctx context.Context, rpcMessage []*rpcApiMessage) {
	m.rpcQueue.Push(rpcMessage)
}

func (m *MainAuthWrapper) onTimer(ctx context.Context) {
	for _, sess := range m.mainAuth.sessions {
		sess.onTimer(ctx)
	}
	for _, sess := range m.tempAuth.sessions {
		sess.onTimer(ctx)
	}
	for _, sess := range m.mediaTempAuth.sessions {
		sess.onTimer(ctx)
	}

	if (m.mainUpdatesSession != nil && m.mainUpdatesSession.sessionOnline()) ||
		(m.androidPushSession != nil && m.androidPushSession.sessionOnline()) {
		m.setOnline(ctx)
	}

	for _, sess := range m.mainAuth.sessions {
		if !sess.sessionClosed() {
			return
		}
	}
	for _, sess := range m.tempAuth.sessions {
		if !sess.sessionClosed() {
			return
		}
	}
	for _, sess := range m.mediaTempAuth.sessions {
		if !sess.sessionClosed() {
			return
		}
	}

	m.cb.DeleteByAuthKeyId(m.authKeyId)
	m.Stop()
}

func (m *MainAuthWrapper) SessionClientNew(ctx context.Context, kType int, kId int64, gatewayId string, sessionId int64) error {
	cData := &connDataCtx{
		ctx: contextx.ValueOnlyFrom(ctx),
		connData: connData{
			authType:  kType,
			authId:    kId,
			isNew:     true,
			gatewayId: gatewayId,
			sessionId: sessionId,
		},
	}

	select {
	case m.sessionDataChan <- cData:
	default:
	}

	return nil
}

func (m *MainAuthWrapper) SessionDataArrived(ctx context.Context, kType int, kId int64, gatewayId, clientIp string, sessionId, salt int64, buf []byte) error {
	sData := &sessionDataCtx{
		ctx: contextx.ValueOnlyFrom(ctx),
		sessionData: sessionData{
			authType:  kType,
			authId:    kId,
			gatewayId: gatewayId,
			clientIp:  clientIp,
			sessionId: sessionId,
			salt:      salt,
			buf:       buf,
		},
	}

	select {
	case m.sessionDataChan <- sData:
	default:
	}

	return nil
}

func (m *MainAuthWrapper) SessionHttpDataArrived(ctx context.Context, kType int, kId int64, gatewayId, clientIp string, sessionId, salt int64, buf []byte, resChan chan interface{}) error {
	sData := &sessionHttpDataCtx{
		ctx: contextx.ValueOnlyFrom(ctx),
		sessionHttpData: sessionHttpData{
			authType:   kType,
			authId:     kId,
			gatewayId:  gatewayId,
			clientIp:   clientIp,
			sessionId:  sessionId,
			salt:       salt,
			buf:        buf,
			resChannel: resChan,
		},
	}

	select {
	case m.sessionDataChan <- sData:
	default:
	}

	return nil
}

func (m *MainAuthWrapper) SessionClientClosed(ctx context.Context, kType int, kId int64, gatewayId string, sessionId int64) error {
	cData := &connDataCtx{
		ctx: contextx.ValueOnlyFrom(ctx),
		connData: connData{
			authType:  kType,
			authId:    kId,
			isNew:     false,
			gatewayId: gatewayId,
			sessionId: sessionId,
		},
	}

	select {
	case m.sessionDataChan <- cData:
	default:
	}

	return nil
}

func (m *MainAuthWrapper) SyncRpcResultDataArrived(ctx context.Context, kId int64, sessionId, clientMsgId int64, data []byte) error {
	rData := &syncRpcResultDataCtx{
		ctx: contextx.ValueOnlyFrom(ctx),
		syncRpcResultData: syncRpcResultData{
			authType:    mtproto.AuthKeyTypeUnknown,
			authId:      kId,
			sessionId:   sessionId,
			clientMsgId: clientMsgId,
			data:        data,
		},
	}

	select {
	case m.sessionDataChan <- rData:
	default:
	}

	return nil
}

func (m *MainAuthWrapper) SyncSessionDataArrived(ctx context.Context, kId int64, sessionId int64, updates *mtproto.Updates) error {
	sData := &syncSessionDataCtx{
		ctx: contextx.ValueOnlyFrom(ctx),
		syncSessionData: syncSessionData{
			authType:  mtproto.AuthKeyTypeUnknown,
			authId:    kId,
			sessionId: sessionId,
			data:      &messageData{obj: updates},
		},
	}

	select {
	case m.sessionDataChan <- sData:
	default:
	}

	return nil
}

func (m *MainAuthWrapper) SyncDataArrived(ctx context.Context, needAndroidPush bool, updates *mtproto.Updates) error {
	sData := &syncDataCtx{
		ctx: contextx.ValueOnlyFrom(ctx),
		syncData: syncData{
			needAndroidPush: needAndroidPush,
			data:            &messageData{obj: updates},
		},
	}

	select {
	case m.sessionDataChan <- sData:
	default:
	}

	return nil
}

func (m *MainAuthWrapper) onSessionNew(ctx context.Context, connMsg *connData) {
	sList := m.getSessionList(connMsg.authType)

	if sList.authId == 0 {
		m.resetAuth(connMsg.authType, connMsg.authId)
	} else if sList.authId != connMsg.authId {
		m.resetAuth(connMsg.authType, connMsg.authId)
	}

	sess, ok := sList.sessions[connMsg.sessionId]
	if !ok {
		logx.WithContext(ctx).Infof("onSessionNew - newSession(%d), conn: %s", m.authKeyId, connMsg)
		sess = newSession(connMsg.sessionId, sList)
		sList.sessions[connMsg.sessionId] = sess
	} else {
		sess.sessionState = kSessionStateNew
		logx.WithContext(ctx).Infof("onSessionNew - session(%d) found, conn: %s", m.authKeyId, connMsg)
	}

	sess.onSessionConnNew(ctx, connMsg.gatewayId)
}

func (m *MainAuthWrapper) onSessionData(ctx context.Context, sessionMsg *sessionData) {
	sList := m.getSessionList(sessionMsg.authType)

	if sList.authId == 0 {
		m.resetAuth(sessionMsg.authType, sessionMsg.authId)
	} else if sList.authId != sessionMsg.authId {
		m.resetAuth(sessionMsg.authType, sessionMsg.authId)
	}

	message2 := new(mtproto.TLMessage2)
	err := message2.Decode(mtproto.NewDecodeBuf(sessionMsg.buf))
	if err != nil {
		// TODO(@benqi): close frontend conn??
		logx.WithContext(ctx).Errorf("onSessionData - error: {%s}, data: {sessions: %s, gate_id: %d}", err, m, sessionMsg.gatewayId)
		return
	}

	// TODO(@benqi): load onNew
	if sList.cacheSalt == nil {
		sList.cacheSalt, sList.cacheLastSalt, _ = m.cb.Dao.GetOrFetchNewSalt(ctx, sList.authId)
	} else {
		if int32(time.Now().Unix()) >= sList.cacheSalt.GetValidUntil() {
			sList.cacheSalt, sList.cacheLastSalt, _ = m.cb.Dao.GetOrFetchNewSalt(ctx, sList.authId)
		}
	}

	if sList.cacheSalt == nil {
		logx.WithContext(ctx).Infof("onSessionData - getOrFetchNewSalt nil error, data: {sessions: %s, conn_id: %s}", m, sessionMsg.gatewayId)
		return
	}

	sess, ok := sList.sessions[sessionMsg.sessionId]
	if !ok {
		sess = newSession(sessionMsg.sessionId, sList)
		sList.sessions[sessionMsg.sessionId] = sess
	}

	sess.onSessionConnNew(ctx, sessionMsg.gatewayId)
	sess.onSessionMessageData(ctx, sessionMsg.gatewayId, sessionMsg.clientIp, sessionMsg.salt, message2)
}

func (m *MainAuthWrapper) onSessionHttpData(ctx context.Context, sessionMsg *sessionHttpData) {
	//sList := m.getSessionList(sessionMsg.authType)
	//
	//if sList.authId == 0 {
	//	m.resetAuth(sessionMsg.authType, sessionMsg.authId)
	//} else if sList.authId != sessionMsg.authId {
	//	m.resetAuth(sessionMsg.authType, sessionMsg.authId)
	//}
	//
	//message2 := new(mtproto.TLMessage2)
	//err := message2.Decode(mtproto.NewDecodeBuf(sessionMsg.buf))
	//if err != nil {
	//	// TODO(@benqi): close frontend conn??
	//	logx.WithContext(ctx).Errorf("onSessionHttpData - error: {%s}, data: {sessions: %s, gate_id: %d}", err, m, sessionMsg.gatewayId)
	//	return
	//}
	//
	//// TODO(@benqi): load onNew
	//if sList.cacheSalt == nil {
	//	sList.cacheSalt, sList.cacheLastSalt, _ = m.cb.Dao.GetOrFetchNewSalt(ctx, sList.authId)
	//} else {
	//	if int32(time.Now().Unix()) >= sList.cacheSalt.GetValidUntil() {
	//		sList.cacheSalt, sList.cacheLastSalt, _ = m.cb.Dao.GetOrFetchNewSalt(ctx, sList.authId)
	//	}
	//}
	//
	//if sList.cacheSalt == nil {
	//	logx.WithContext(ctx).Infof("onSessionHttpData - getOrFetchNewSalt nil error, data: {sessions: %s, conn_id: %s}", m, sessionMsg.gatewayId)
	//	return
	//}
	//
	//sess, ok := sList.sessions[sessionMsg.sessionId]
	//if !ok {
	//	sess = newSession(sessionMsg.sessionId, sList)
	//	sList.sessions[sessionMsg.sessionId] = sess
	//}
	//
	//sess.isHttp = true
	//sess.httpQueue.Push(sessionMsg.resChannel)
	//sess.onSessionConnNew(ctx, sessionMsg.gatewayId)
	//sess.onSessionMessageData(ctx, sessionMsg.gatewayId, sessionMsg.clientIp, sessionMsg.salt, message2)
}

func (m *MainAuthWrapper) onSessionClosed(ctx context.Context, connMsg *connData) {
	sList := m.getSessionList(connMsg.authType)

	if sess, ok := sList.sessions[connMsg.sessionId]; !ok {
		logx.WithContext(ctx).Errorf("onSessionClosed - session conn closed -  conn: %s", connMsg)
	} else {
		logx.WithContext(ctx).Infof("onSessionClosed - conn: %s, sess: %s", connMsg, sess)
		sess.onSessionConnClose(ctx, connMsg.gatewayId)
	}
}

func (m *MainAuthWrapper) onSyncRpcResultData(ctx context.Context, syncMsg *syncRpcResultData) {
	logx.WithContext(ctx).Infof("onSyncRpcResultData - receive data: {sess: %s}",
		m)

	sList := m.getSessionListById(syncMsg.authId)
	if sList == nil {
		logx.WithContext(ctx).Errorf("onSyncRpcResultData - not found sessionList by authId: %d", syncMsg.authId)
		return
	}

	sess, _ := sList.sessions[syncMsg.sessionId]
	if sess != nil {
		sess.onSyncRpcResultData(ctx, syncMsg.clientMsgId, syncMsg.data)
	}
}

func (m *MainAuthWrapper) onSyncSessionData(ctx context.Context, syncMsg *syncSessionData) {
	logx.WithContext(ctx).Infof("onSyncSessionData - receive data: {sess: %s}",
		m)

	sList := m.getSessionListById(syncMsg.authId)
	if sList == nil {
		logx.WithContext(ctx).Errorf("onSyncRpcResultData - not found sessionList by authId: %d", syncMsg.authId)
		return
	}

	sess, _ := sList.sessions[syncMsg.sessionId]
	if sess != nil {
		sess.onSyncSessionData(ctx, syncMsg.data.obj)
	}
}

func (m *MainAuthWrapper) onSyncData(ctx context.Context, syncMsg *syncData) {
	logx.WithContext(ctx).Info("authSessions - ", reflect.TypeOf(syncMsg.data.obj))

	if upds, ok := syncMsg.data.obj.(*mtproto.Updates); ok {
		if upds.PredicateName == mtproto.Predicate_updateAccountResetAuthorization {
			logx.WithContext(ctx).Info("recv updateAccountResetAuthorization - ", reflect.TypeOf(syncMsg.data.obj))
			if m.AuthUserId != upds.GetUserId() {
				logx.WithContext(ctx).Error("upds -- ", upds)
			}
			// m.cb.Dao.PutCacheUserId(context.Background(), m.authKeyId, 0)
			// m.cb.DeleteByAuthKeyId(m.authKeyId)
			m.changeAuthState(ctx, mtproto.AuthStateDeleted, 0)
			// m.AuthUserId = 0
			return
		}
	}

	if m.mainUpdatesSession != nil {
		m.mainUpdatesSession.onSyncData(ctx, syncMsg.data.obj)
	}

	if syncMsg.needAndroidPush && m.androidPushSession != nil {
		m.androidPushSession.onSyncData(ctx, nil)
	}
}

func doRpcRequest(ctx context.Context, dao *dao.Dao, md *metadata.RpcMetadata, request *rpcApiMessage) {
	var (
		err       error
		rpcResult mtproto.TLObject
	)

	// TODO(@benqi): change state.
	switch request.reqMsg.(type) {
	case *mtproto.TLAuthBindTempAuthKey:
		r := request.reqMsg.(*mtproto.TLAuthBindTempAuthKey)
		rpcResult, err = dao.AuthsessionClient.AuthsessionBindTempAuthKey(
			ctx,
			&authsession.TLAuthsessionBindTempAuthKey{
				PermAuthKeyId:    r.PermAuthKeyId,
				Nonce:            r.Nonce,
				ExpiresAt:        r.ExpiresAt,
				EncryptedMessage: r.EncryptedMessage,
			})
	default:
		rpcResult, err = dao.InvokeContext(ctx, md, request.reqMsg)
	}

	reply := &mtproto.TLRpcResult{
		ReqMsgId: request.reqMsgId,
		Result:   nil,
	}

	if err != nil {
		logx.WithContext(ctx).Error(err.Error())
		reply.Result = mtproto.NewRpcError(err)
	} else {
		logx.WithContext(ctx).Infof("invokeRpcRequest - rpc_result: {%s}", reflect.TypeOf(rpcResult))
		reply.Result = rpcResult
	}

	request.rpcResult = reply
}
