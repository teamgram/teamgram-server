package sessionstate

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/mtproto"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/mt"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const (
	defaultFutureSaltsCount = 32
	maxFutureSaltsCount     = 64
)

func (p *Processor) handleServiceMessage(ctx context.Context, keyInfo *tg.AuthKeyInfo, obj iface.TLObject, msg mtproto.EncryptedMessage) ([]byte, bool, error) {
	switch request := obj.(type) {
	case *mt.TLPing:
		return encodeObject(&mt.TLPong{MsgId: msg.MsgId, PingId: request.PingId}), true, nil
	case *mt.TLPingDelayDisconnect:
		return encodeObject(&mt.TLPong{MsgId: msg.MsgId, PingId: request.PingId}), true, nil
	case *mt.TLMsgsAck:
		if p != nil && p.runtime != nil {
			p.runtime.ackOutbound(runtimeSessionKey{
				authKeyId:   msg.AuthKeyId,
				authKeyType: authKeyType(keyInfo),
				sessionId:   msg.SessionId,
			}, request.MsgIds)
		}
		return nil, true, nil
	case *mt.TLGetFutureSalts:
		salts, err := p.getFutureSalts(ctx, msg.AuthKeyId, normalizeFutureSaltsCount(request.Num))
		if err != nil {
			return nil, true, err
		}
		salts.ReqMsgId = msg.MsgId
		if salts.Now == 0 {
			salts.Now = int32(time.Now().Unix())
		}
		return encodeObject(salts), true, nil
	case *mt.TLDestroySession:
		target := runtimeSessionKey{
			authKeyId:   msg.AuthKeyId,
			authKeyType: authKeyType(keyInfo),
			sessionId:   request.SessionId,
		}
		if request.SessionId != msg.SessionId && p != nil && p.runtime != nil && p.runtime.destroySession(target) {
			if p.roles != nil {
				p.roles.unregisterSession(roleSessionKey{
					permAuthKeyId: metadataCacheID(keyInfo, msg.AuthKeyId),
					authKeyId:     msg.AuthKeyId,
					authKeyType:   authKeyType(keyInfo),
					sessionId:     request.SessionId,
				})
			}
			return encodeObject(&mt.TLDestroySessionOk{SessionId: request.SessionId}), true, nil
		}
		return encodeObject(&mt.TLDestroySessionNone{SessionId: request.SessionId}), true, nil
	case *mt.TLMsgsStateReq:
		return encodeObject(&mt.TLMsgsStateInfo{
			ReqMsgId: msg.MsgId,
			Info:     strings.Repeat("\x01", len(request.MsgIds)),
		}), true, nil
	default:
		return nil, false, nil
	}
}

func (p *Processor) getFutureSalts(ctx context.Context, authKeyID int64, num int32) (*mt.TLFutureSalts, error) {
	if p == nil || p.store == nil {
		return nil, fmt.Errorf("session processor: auth key store is nil")
	}
	salts, err := p.store.GetFutureSalts(ctx, authKeyID, num)
	if err != nil {
		return nil, fmt.Errorf("session processor: get future salts for auth key %d: %w", authKeyID, err)
	}
	if salts == nil {
		return &mt.TLFutureSalts{}, nil
	}
	copySalts := *salts
	copySalts.Salts = append([]*mt.TLFutureSalt(nil), salts.Salts...)
	return &copySalts, nil
}

func normalizeFutureSaltsCount(num int32) int32 {
	switch {
	case num <= 0:
		return defaultFutureSaltsCount
	case num > maxFutureSaltsCount:
		return maxFutureSaltsCount
	default:
		return num
	}
}
