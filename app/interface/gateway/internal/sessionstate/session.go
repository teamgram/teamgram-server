package sessionstate

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	gmtproto "github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/mtproto"
	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/repository"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/crypto"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/mt"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type ConnInfo struct {
	GatewayId  string
	ClientAddr string
}

type Dispatcher interface {
	Invoke(ctx context.Context, md *metadata.RpcMetadata, payload []byte) ([]byte, error)
}

type ActiveSession struct {
	AuthKeyId int64
	SessionId int64
	Salt      int64
	AuthKey   *crypto.AuthKey
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
	seqMu        sync.Mutex
	seq          map[sessionKey]int32
}

type sessionKey struct {
	authKeyId int64
	sessionId int64
}

func NewProcessor(store repository.AuthKeyStore, dispatch Dispatcher) *Processor {
	return &Processor{
		store:        store,
		dispatch:     dispatch,
		authKeys:     make(map[int64]*crypto.AuthKey),
		authKeyInfos: make(map[int64]*tg.AuthKeyInfo),
		seq:          make(map[sessionKey]int32),
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
		seq = observe(ActiveSession{
			AuthKeyId: msg.AuthKeyId,
			SessionId: msg.SessionId,
			Salt:      msg.Salt,
			AuthKey:   key,
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
		SeqNo:     p.nextSeqNo(msg.AuthKeyId, msg.SessionId, true, seq),
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
					Seqno:  p.nextSeqNo(parent.AuthKeyId, parent.SessionId, true, seq),
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
			Seqno:  p.nextSeqNo(parent.AuthKeyId, parent.SessionId, true, seq),
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
	md := &metadata.RpcMetadata{
		ServerId:      conn.GatewayId,
		ClientAddr:    conn.ClientAddr,
		AuthId:        msg.AuthKeyId,
		SessionId:     msg.SessionId,
		ReceiveTime:   time.Now().Unix(),
		ClientMsgId:   msg.MsgId,
		Layer:         wrapperMD.Layer,
		Client:        wrapperMD.Client,
		Langpack:      wrapperMD.Langpack,
		LangCode:      wrapperMD.LangCode,
		PermAuthKeyId: msg.AuthKeyId,
	}
	if keyInfo != nil && keyInfo.PermAuthKeyId != 0 {
		md.PermAuthKeyId = keyInfo.PermAuthKeyId
	}
	result, err := p.dispatch.Invoke(ctx, md, inner)
	if err != nil {
		var rpcErr interface {
			RPCError() *tg.TLRpcError
		}
		if errors.As(err, &rpcErr) {
			e := rpcErr.RPCError()
			if e != nil {
				return gmtproto.WrapRPCError(msg.MsgId, e.ErrorCode, e.ErrorMessage)
			}
		}
		return nil, err
	}
	return gmtproto.WrapRPCResult(msg.MsgId, result)
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

func (p *Processor) nextSeqNo(authKeyId int64, sessionId int64, contentRelated bool, allocator SeqNoAllocator) int32 {
	if allocator != nil {
		return allocator.NextSeqNo(contentRelated)
	}
	p.seqMu.Lock()
	defer p.seqMu.Unlock()
	key := sessionKey{authKeyId: authKeyId, sessionId: sessionId}
	seq := p.seq[key] * 2
	if contentRelated {
		seq++
		p.seq[key]++
	}
	return seq
}
