package sessionstate

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	gmtproto "github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/mtproto"
	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/crypto"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/mt"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
	"github.com/zeromicro/go-zero/core/logx"
)

type ConnInfo struct {
	GatewayId  string
	ClientAddr string
}

type Dispatcher interface {
	Invoke(ctx context.Context, md *metadata.RpcMetadata, payload []byte) ([]byte, error)
}

type ActiveSession struct {
	PermAuthKeyId int64
	AuthKeyId     int64
	AuthKeyType   int32
	SessionId     int64
	Layer         int32
	Salt          int64
	AuthKey       *crypto.AuthKey
}

type SeqNoAllocator interface {
	NextSeqNo(contentRelated bool) int32
}

type SessionObserver func(ActiveSession) SeqNoAllocator

type Processor struct {
	store        repository.AuthKeyStore
	dispatch     Dispatcher
	authKeysMu   sync.RWMutex
	authKeys     map[int64]*crypto.AuthKey
	authKeyInfos map[int64]*tg.AuthKeyInfo
	metaMu       sync.RWMutex
	clientMeta   map[clientMetadataKey]gmtproto.WrapperMetadata
	seqMu        sync.Mutex
	seq          map[activeSessionKey]int32
}

type sessionKey struct {
	authKeyId int64
	sessionId int64
}

type activeSessionKey struct {
	authKeyId   int64
	authKeyType int32
	sessionId   int64
}

type clientMetadataKey struct {
	permAuthKeyId int64
}

func NewProcessor(store repository.AuthKeyStore, dispatch Dispatcher) *Processor {
	return &Processor{
		store:        store,
		dispatch:     dispatch,
		authKeys:     make(map[int64]*crypto.AuthKey),
		authKeyInfos: make(map[int64]*tg.AuthKeyInfo),
		clientMeta:   make(map[clientMetadataKey]gmtproto.WrapperMetadata),
		seq:          make(map[activeSessionKey]int32),
	}
}

func (p *Processor) HandleEncrypted(ctx context.Context, conn ConnInfo, payload []byte) ([]byte, error) {
	return p.HandleEncryptedWithSession(ctx, conn, payload, nil)
}

func (p *Processor) HandleEncryptedWithSession(ctx context.Context, conn ConnInfo, payload []byte, observe SessionObserver) ([]byte, error) {
	authKeyId, err := readAuthKeyID(payload)
	if err != nil {
		return nil, err
	}
	key, keyInfo, err := p.authKey(ctx, authKeyId)
	if err != nil {
		return nil, err
	}

	msg, err := gmtproto.DecodeEncryptedMessage(payload, key)
	if err != nil {
		return nil, err
	}
	var seq SeqNoAllocator
	if observe != nil {
		permAuthKeyId, layer, err := p.activeSessionMetadata(ctx, keyInfo, msg.AuthKeyId)
		if err != nil {
			return nil, err
		}
		seq = observe(ActiveSession{
			PermAuthKeyId: permAuthKeyId,
			AuthKeyId:     msg.AuthKeyId,
			AuthKeyType:   authKeyType(keyInfo),
			SessionId:     msg.SessionId,
			Layer:         layer,
			Salt:          msg.Salt,
			AuthKey:       key,
		})
	}

	body, err := p.handleMessageBody(ctx, conn, key, keyInfo, msg, seq)
	if err != nil || body == nil {
		return nil, err
	}
	return gmtproto.EncodeEncryptedMessage(gmtproto.EncryptedMessage{
		AuthKeyId: msg.AuthKeyId,
		Salt:      msg.Salt,
		SessionId: msg.SessionId,
		MsgId:     gmtproto.NextServerMsgId(msg.MsgId),
		SeqNo:     p.nextSeqNo(msg.AuthKeyId, authKeyType(keyInfo), msg.SessionId, true, seq),
		Body:      body,
	}, key)
}

func (p *Processor) authKey(ctx context.Context, authKeyId int64) (*crypto.AuthKey, *tg.AuthKeyInfo, error) {
	p.authKeysMu.RLock()
	if key, ok := p.authKeys[authKeyId]; ok {
		keyInfo := p.authKeyInfos[authKeyId]
		p.authKeysMu.RUnlock()
		return key, keyInfo, nil
	}
	p.authKeysMu.RUnlock()

	p.authKeysMu.Lock()
	defer p.authKeysMu.Unlock()
	if key, ok := p.authKeys[authKeyId]; ok {
		return key, p.authKeyInfos[authKeyId], nil
	}
	if p.store == nil {
		return nil, nil, fmt.Errorf("session processor: auth key store is nil")
	}
	keyInfo, err := p.store.QueryAuthKey(ctx, authKeyId)
	if err != nil {
		return nil, nil, fmt.Errorf("session processor: query auth key %d: %w", authKeyId, err)
	}
	if keyInfo == nil || len(keyInfo.AuthKey) == 0 {
		return nil, nil, fmt.Errorf("session processor: auth key %d not found", authKeyId)
	}
	key := crypto.NewAuthKey(keyInfo.AuthKeyId, keyInfo.AuthKey)
	p.authKeys[authKeyId] = key
	p.authKeyInfos[authKeyId] = keyInfo
	return key, keyInfo, nil
}

func (p *Processor) refreshAuthKeyInfo(ctx context.Context, authKeyId int64, current *tg.AuthKeyInfo) (*tg.AuthKeyInfo, error) {
	if p.store == nil {
		return nil, fmt.Errorf("session processor: auth key store is nil")
	}
	keyInfo, err := p.store.QueryAuthKey(ctx, authKeyId)
	if err != nil {
		return nil, fmt.Errorf("session processor: refresh auth key %d: %w", authKeyId, err)
	}
	if keyInfo == nil || len(keyInfo.AuthKey) == 0 {
		return nil, fmt.Errorf("session processor: auth key %d not found", authKeyId)
	}

	p.authKeysMu.Lock()
	defer p.authKeysMu.Unlock()
	cached := p.authKeyInfos[authKeyId]
	switch {
	case cached != nil:
		*cached = *keyInfo
		keyInfo = cached
	case current != nil:
		*current = *keyInfo
		p.authKeyInfos[authKeyId] = current
		keyInfo = current
	default:
		p.authKeyInfos[authKeyId] = keyInfo
	}
	if _, ok := p.authKeys[authKeyId]; !ok {
		p.authKeys[authKeyId] = crypto.NewAuthKey(keyInfo.AuthKeyId, keyInfo.AuthKey)
	}
	return keyInfo, nil
}

func (p *Processor) handleMessageBody(ctx context.Context, conn ConnInfo, key *crypto.AuthKey, keyInfo *tg.AuthKeyInfo, msg gmtproto.EncryptedMessage, seq SeqNoAllocator) ([]byte, error) {
	obj, err := iface.DecodeObject(bin.NewDecoder(msg.Body))
	if err != nil {
		return nil, fmt.Errorf("session processor: decode body: %w", err)
	}

	if body, handled, err := p.handleServiceMessage(obj, msg); handled || err != nil {
		return body, err
	}
	if container, ok := obj.(*mt.TLMsgContainer); ok {
		rawContainer, err := decodeRawContainer(msg.Body)
		if err != nil {
			return nil, err
		}
		return p.handleContainer(ctx, conn, keyInfo, msg, container, rawContainer, seq)
	}
	return p.dispatchRPC(ctx, conn, keyInfo, msg, msg.Body)
}

func (p *Processor) handleContainer(ctx context.Context, conn ConnInfo, keyInfo *tg.AuthKeyInfo, parent gmtproto.EncryptedMessage, container *mt.TLMsgContainer, rawContainer *mt.TLMsgRawDataContainer, seq SeqNoAllocator) ([]byte, error) {
	if len(container.Messages) != len(rawContainer.Messages) {
		return nil, fmt.Errorf("session processor: container raw message count %d does not match decoded count %d", len(rawContainer.Messages), len(container.Messages))
	}
	var responses []*mt.TLMessage2
	for i, child := range container.Messages {
		rawChild := rawContainer.Messages[i]
		if child.MsgId != rawChild.MsgId || child.Seqno != rawChild.Seqno {
			return nil, fmt.Errorf("session processor: container child %d raw metadata mismatch", i)
		}
		childMsg := gmtproto.EncryptedMessage{
			AuthKeyId: parent.AuthKeyId,
			Salt:      parent.Salt,
			SessionId: parent.SessionId,
			MsgId:     child.MsgId,
			SeqNo:     child.Seqno,
			Body:      rawMessageBody(rawChild),
		}
		if body, handled, err := p.handleServiceMessage(child.Object, childMsg); err != nil {
			return nil, err
		} else if handled {
			if body != nil {
				responses = append(responses, &mt.TLMessage2{
					MsgId:  gmtproto.NextServerMsgId(childMsg.MsgId),
					Seqno:  p.nextSeqNo(parent.AuthKeyId, authKeyType(keyInfo), parent.SessionId, true, seq),
					Object: codec.NewRawTLObject(body),
				})
			}
			continue
		}
		body, err := p.dispatchRPC(ctx, conn, keyInfo, childMsg, childMsg.Body)
		if err != nil {
			return nil, err
		}
		if body == nil {
			continue
		}
		responses = append(responses, &mt.TLMessage2{
			MsgId:  gmtproto.NextServerMsgId(childMsg.MsgId),
			Seqno:  p.nextSeqNo(parent.AuthKeyId, authKeyType(keyInfo), parent.SessionId, true, seq),
			Object: codec.NewRawTLObject(body),
		})
	}
	switch len(responses) {
	case 0:
		return nil, nil
	default:
		return encodeObject(&mt.TLMsgContainer{Messages: responses}), nil
	}
}

func decodeRawContainer(payload []byte) (*mt.TLMsgRawDataContainer, error) {
	d := bin.NewDecoder(payload)
	clazzID, err := d.ClazzID()
	if err != nil {
		return nil, fmt.Errorf("session processor: decode raw container constructor: %w", err)
	}
	if clazzID != mt.ClazzID_msg_container {
		return nil, fmt.Errorf("session processor: decode raw container: unexpected constructor %#x", clazzID)
	}
	container := new(mt.TLMsgRawDataContainer)
	if err := container.Decode(d); err != nil {
		return nil, fmt.Errorf("session processor: decode raw container: %w", err)
	}
	if d.Remaining() != 0 {
		return nil, fmt.Errorf("session processor: decode raw container: %d trailing bytes", d.Remaining())
	}
	return container, nil
}

func rawMessageBody(msg *mt.TLMessageRawData) []byte {
	x := bin.NewEncoder()
	defer x.End()
	x.PutClazzID(msg.ClazzID)
	x.PutRaw(msg.Body)
	return append([]byte(nil), x.Bytes()...)
}

func (p *Processor) dispatchRPC(ctx context.Context, conn ConnInfo, keyInfo *tg.AuthKeyInfo, msg gmtproto.EncryptedMessage, payload []byte) ([]byte, error) {
	if p.dispatch == nil {
		return nil, fmt.Errorf("session processor: dispatcher is nil")
	}
	inner, wrapperMD, err := gmtproto.UnwrapClientRPC(payload)
	if err != nil {
		return nil, err
	}
	cacheID := metadataCacheID(keyInfo, msg.AuthKeyId)
	incomingMD := wrapperMD
	if hasWrapperMetadata(incomingMD) {
		incomingMD.Ip = clientIP(conn.ClientAddr)
	}
	wrapperMD, changed := p.mergeClientMetadata(cacheID, incomingMD)
	if hasWrapperMetadata(incomingMD) && changed {
		if hasInitConnectionMetadata(wrapperMD) {
			if err := p.store.SetClientSessionInfo(ctx, wrapperMetadataToClientSession(msg.AuthKeyId, wrapperMD)); err != nil {
				return nil, fmt.Errorf("session processor: set client session info for auth key %d: %w", msg.AuthKeyId, err)
			}
		} else if wrapperMD.Layer != 0 {
			if err := p.store.SetLayer(ctx, msg.AuthKeyId, clientIP(conn.ClientAddr), wrapperMD.Layer); err != nil {
				return nil, fmt.Errorf("session processor: set client layer for auth key %d: %w", msg.AuthKeyId, err)
			}
		}
	}
	if !hasClientMetadata(wrapperMD) {
		if cached, ok := p.cachedClientMetadata(cacheID); ok {
			wrapperMD = cached
		} else if p.store != nil {
			client, err := p.store.GetClientSession(ctx, msg.AuthKeyId)
			if err != nil {
				return nil, fmt.Errorf("session processor: get client session for auth key %d: %w", msg.AuthKeyId, err)
			}
			if client != nil {
				wrapperMD, _ = p.mergeClientMetadata(cacheID, clientSessionToWrapperMetadata(client))
			}
		}
	}
	md := &metadata.RpcMetadata{
		ServerId:      conn.GatewayId,
		ClientAddr:    conn.ClientAddr,
		AuthId:        msg.AuthKeyId,
		SessionId:     msg.SessionId,
		ReceiveTime:   time.Now().Unix(),
		ClientMsgId:   msg.MsgId,
		Layer:         wrapperMD.Layer,
		Client:        wrapperMD.Client,
		Langpack:      wrapperMD.LangPack,
		LangCode:      wrapperMD.LangCode,
		PermAuthKeyId: msg.AuthKeyId,
	}
	if keyInfo != nil && keyInfo.PermAuthKeyId != 0 {
		md.PermAuthKeyId = keyInfo.PermAuthKeyId
	}
	if p.store != nil && msg.AuthKeyId != 0 {
		userID, err := p.store.GetUserId(ctx, msg.AuthKeyId)
		if err != nil {
			return nil, fmt.Errorf("session processor: get user id for auth key %d: %w", msg.AuthKeyId, err)
		}
		md.UserId = userID
	}
	method := rawRPCMethodName(inner)
	result, err := p.dispatch.Invoke(ctx, md, inner)
	if err != nil {
		var rpcErr interface {
			RPCError() *tg.TLRpcError
		}
		if errors.As(err, &rpcErr) {
			e := rpcErr.RPCError()
			if e != nil {
				logx.WithContext(ctx).Errorf(
					"gateway rpc request error: method=%s msg_id=%d auth_key_id=%d perm_auth_key_id=%d session_id=%d client_addr=%s error_code=%d error_message=%s",
					method,
					msg.MsgId,
					msg.AuthKeyId,
					md.PermAuthKeyId,
					msg.SessionId,
					conn.ClientAddr,
					e.ErrorCode,
					e.ErrorMessage,
				)
				return gmtproto.WrapRPCError(msg.MsgId, e.ErrorCode, e.ErrorMessage)
			}
		}
		logx.WithContext(ctx).Errorf(
			"gateway rpc request infrastructure error: method=%s msg_id=%d auth_key_id=%d perm_auth_key_id=%d session_id=%d client_addr=%s err=%v",
			method,
			msg.MsgId,
			msg.AuthKeyId,
			md.PermAuthKeyId,
			msg.SessionId,
			conn.ClientAddr,
			err,
		)
		return nil, err
	}
	if method == tg.ClazzName_auth_bindTempAuthKey {
		if _, err := p.refreshAuthKeyInfo(ctx, msg.AuthKeyId, keyInfo); err != nil {
			return nil, err
		}
	}
	return gmtproto.WrapRPCResult(msg.MsgId, result)
}

func metadataCacheID(keyInfo *tg.AuthKeyInfo, authKeyId int64) int64 {
	if keyInfo != nil && keyInfo.PermAuthKeyId != 0 {
		return keyInfo.PermAuthKeyId
	}
	return authKeyId
}

func authKeyType(keyInfo *tg.AuthKeyInfo) int32 {
	if keyInfo == nil {
		return tg.AuthKeyTypeUnknown
	}
	return keyInfo.AuthKeyType
}

func (p *Processor) activeSessionMetadata(ctx context.Context, keyInfo *tg.AuthKeyInfo, authKeyId int64) (int64, int32, error) {
	cacheID := metadataCacheID(keyInfo, authKeyId)
	if md, ok := p.cachedClientMetadata(cacheID); ok {
		return cacheID, md.Layer, nil
	}
	if p.store == nil {
		return cacheID, 0, nil
	}
	client, err := p.store.GetClientSession(ctx, authKeyId)
	if err != nil {
		return 0, 0, fmt.Errorf("session processor: get client session for auth key %d: %w", authKeyId, err)
	}
	if client == nil {
		return cacheID, 0, nil
	}
	md, _ := p.mergeClientMetadata(cacheID, clientSessionToWrapperMetadata(client))
	return cacheID, md.Layer, nil
}

func (p *Processor) mergeClientMetadata(cacheID int64, md gmtproto.WrapperMetadata) (gmtproto.WrapperMetadata, bool) {
	key := clientMetadataKey{permAuthKeyId: cacheID}
	p.metaMu.Lock()
	defer p.metaMu.Unlock()
	old := p.clientMeta[key]
	merged := mergeWrapperMetadata(old, md)
	changed := persistedMetadataChanged(old, merged)
	if hasClientMetadata(merged) {
		p.clientMeta[key] = merged
	}
	return merged, changed
}

func (p *Processor) cachedClientMetadata(cacheID int64) (gmtproto.WrapperMetadata, bool) {
	p.metaMu.RLock()
	defer p.metaMu.RUnlock()
	md, ok := p.clientMeta[clientMetadataKey{permAuthKeyId: cacheID}]
	return md, ok
}

func mergeWrapperMetadata(old, next gmtproto.WrapperMetadata) gmtproto.WrapperMetadata {
	if next.Layer == 0 {
		next.Layer = old.Layer
	}
	if next.Client == "" {
		next.Client = old.Client
	}
	if next.Ip == "" {
		next.Ip = old.Ip
	}
	if next.ApiId == 0 {
		next.ApiId = old.ApiId
	}
	if next.DeviceModel == "" {
		next.DeviceModel = old.DeviceModel
	}
	if next.SystemVersion == "" {
		next.SystemVersion = old.SystemVersion
	}
	if next.AppVersion == "" {
		next.AppVersion = old.AppVersion
	}
	if next.SystemLangCode == "" {
		next.SystemLangCode = old.SystemLangCode
	}
	if next.LangPack == "" {
		next.LangPack = old.LangPack
	}
	if next.LangCode == "" {
		next.LangCode = old.LangCode
	}
	if next.Proxy == "" {
		next.Proxy = old.Proxy
	}
	if next.Params == "" {
		next.Params = old.Params
	}
	return next
}

func persistedMetadataChanged(old, next gmtproto.WrapperMetadata) bool {
	return old.Layer != next.Layer ||
		old.Ip != next.Ip ||
		old.ApiId != next.ApiId ||
		old.DeviceModel != next.DeviceModel ||
		old.SystemVersion != next.SystemVersion ||
		old.AppVersion != next.AppVersion ||
		old.SystemLangCode != next.SystemLangCode ||
		old.LangPack != next.LangPack ||
		old.LangCode != next.LangCode ||
		old.Proxy != next.Proxy ||
		old.Params != next.Params
}

func hasWrapperMetadata(md gmtproto.WrapperMetadata) bool {
	return md.Layer != 0 || hasInitConnectionMetadata(md)
}

func hasClientMetadata(md gmtproto.WrapperMetadata) bool {
	return md.Layer != 0 || md.Client != "" || md.LangPack != "" || md.LangCode != "" || hasInitConnectionMetadata(md)
}

func hasInitConnectionMetadata(md gmtproto.WrapperMetadata) bool {
	return md.ApiId != 0 ||
		md.DeviceModel != "" ||
		md.SystemVersion != "" ||
		md.AppVersion != "" ||
		md.SystemLangCode != "" ||
		md.Proxy != "" ||
		md.Params != ""
}

func wrapperMetadataToClientSession(authKeyId int64, md gmtproto.WrapperMetadata) *authsession.ClientSession {
	return authsession.MakeTLClientSession(&authsession.TLClientSession{
		AuthKeyId:      authKeyId,
		Ip:             md.Ip,
		Layer:          md.Layer,
		ApiId:          md.ApiId,
		DeviceModel:    md.DeviceModel,
		SystemVersion:  md.SystemVersion,
		AppVersion:     md.AppVersion,
		SystemLangCode: md.SystemLangCode,
		LangPack:       md.LangPack,
		LangCode:       md.LangCode,
		Proxy:          md.Proxy,
		Params:         md.Params,
	}).ToClientSession()
}

func clientSessionToWrapperMetadata(client *authsession.ClientSession) gmtproto.WrapperMetadata {
	if client == nil {
		return gmtproto.WrapperMetadata{}
	}
	return gmtproto.WrapperMetadata{
		Layer:          client.Layer,
		Client:         strings.TrimSpace(client.DeviceModel + " " + client.SystemVersion + " " + client.AppVersion),
		Ip:             client.Ip,
		ApiId:          client.ApiId,
		DeviceModel:    client.DeviceModel,
		SystemVersion:  client.SystemVersion,
		AppVersion:     client.AppVersion,
		SystemLangCode: client.SystemLangCode,
		LangPack:       client.LangPack,
		LangCode:       client.LangCode,
		Proxy:          client.Proxy,
		Params:         client.Params,
	}
}

func clientIP(clientAddr string) string {
	host, _, err := net.SplitHostPort(clientAddr)
	if err != nil || host == "" {
		return clientAddr
	}
	return host
}

func rawRPCMethodName(payload []byte) string {
	constructorID, err := bin.NewDecoder(payload).PeekClazzID()
	if err != nil {
		return "unknown"
	}
	if name := iface.GetClazzNameByID(constructorID); name != "" {
		return name
	}
	return fmt.Sprintf("unknown#%08x", constructorID)
}

func readAuthKeyID(payload []byte) (int64, error) {
	if len(payload) < 8 {
		return 0, fmt.Errorf("session processor: encrypted payload too short")
	}
	v, err := bin.NewDecoder(payload).Int64()
	if err != nil {
		return 0, err
	}
	return v, nil
}

func (p *Processor) nextSeqNo(authKeyId int64, authKeyType int32, sessionId int64, contentRelated bool, allocator SeqNoAllocator) int32 {
	if allocator != nil {
		return allocator.NextSeqNo(contentRelated)
	}
	p.seqMu.Lock()
	defer p.seqMu.Unlock()
	key := activeSessionKey{authKeyId: authKeyId, authKeyType: authKeyType, sessionId: sessionId}
	seq := p.seq[key] * 2
	if contentRelated {
		seq++
		p.seq[key]++
	}
	return seq
}
