package push

import (
	"context"
	"sync"

	gmtproto "github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/mtproto"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/crypto"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type SessionWriter interface {
	NextSeqNo(contentRelated bool) int32
	WriteEncrypted(ctx context.Context, msg gmtproto.EncryptedMessage) error
}

type LocalTarget struct {
	PermAuthKeyId int64
	AuthKeyId     int64
	AuthKeyType   int32
	SessionId     int64
	Layer         int32
	AuthKey       *crypto.AuthKey
	Writer        SessionWriter
	registerSeq   uint64
}

type LocalWriter struct {
	mu          sync.RWMutex
	targets     map[sessionKey]LocalTarget
	registerSeq uint64
}

type sessionKey struct {
	authKeyId   int64
	authKeyType int32
	sessionId   int64
}

func NewLocalWriter() *LocalWriter {
	return &LocalWriter{targets: make(map[sessionKey]LocalTarget)}
}

func (w *LocalWriter) Register(target LocalTarget) {
	if w == nil || target.Writer == nil || target.AuthKey == nil {
		return
	}
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.targets == nil {
		w.targets = make(map[sessionKey]LocalTarget)
	}
	if target.PermAuthKeyId == 0 {
		target.PermAuthKeyId = target.AuthKeyId
	}
	w.registerSeq++
	target.registerSeq = w.registerSeq
	w.targets[sessionKey{authKeyId: target.AuthKeyId, authKeyType: target.AuthKeyType, sessionId: target.SessionId}] = target
}

func (w *LocalWriter) Unregister(authKeyId int64, authKeyType int32, sessionId int64) {
	if w == nil {
		return
	}
	w.mu.Lock()
	defer w.mu.Unlock()
	delete(w.targets, sessionKey{authKeyId: authKeyId, authKeyType: authKeyType, sessionId: sessionId})
}

func (w *LocalWriter) WriteRPCResult(ctx context.Context, authKeyId int64, sessionId int64, reqMsgId int64, rpcResultData []byte) (bool, error) {
	if w == nil {
		return false, nil
	}
	w.mu.RLock()
	var (
		target  LocalTarget
		matches int
	)
	for key, candidate := range w.targets {
		if key.authKeyId == authKeyId && key.sessionId == sessionId {
			target = candidate
			matches++
		}
	}
	w.mu.RUnlock()
	if matches != 1 {
		return false, nil
	}
	return true, target.Writer.WriteEncrypted(ctx, gmtproto.EncryptedMessage{
		AuthKeyId: target.AuthKeyId,
		SessionId: target.SessionId,
		MsgId:     gmtproto.NextServerMsgId(reqMsgId),
		SeqNo:     target.Writer.NextSeqNo(true),
		Body:      append([]byte(nil), rpcResultData...),
	})
}

func (w *LocalWriter) WriteUpdates(ctx context.Context, permAuthKeyId int64, updates tg.UpdatesClazz) (int, error) {
	if w == nil || updates == nil {
		return 0, nil
	}
	targets := w.matchTargets(func(key sessionKey, target LocalTarget) bool {
		return target.PermAuthKeyId == permAuthKeyId && isNormalAuthKeyType(key.authKeyType)
	})
	targets = newestTargetByAuthKey(targets)
	return writeUpdatesToTargets(ctx, targets, updates)
}

func (w *LocalWriter) WriteSessionUpdates(ctx context.Context, authKeyId int64, sessionId int64, updates tg.UpdatesClazz) (bool, error) {
	if w == nil || updates == nil {
		return false, nil
	}
	targets := w.matchTargets(func(key sessionKey, target LocalTarget) bool {
		return key.authKeyId == authKeyId && key.sessionId == sessionId
	})
	if len(targets) != 1 {
		return false, nil
	}
	_, err := writeUpdatesToTargets(ctx, targets, updates)
	return err == nil, err
}

func (w *LocalWriter) matchTargets(match func(sessionKey, LocalTarget) bool) []LocalTarget {
	w.mu.RLock()
	defer w.mu.RUnlock()
	targets := make([]LocalTarget, 0, len(w.targets))
	for key, target := range w.targets {
		if match(key, target) {
			targets = append(targets, target)
		}
	}
	return targets
}

func writeUpdatesToTargets(ctx context.Context, targets []LocalTarget, updates tg.UpdatesClazz) (int, error) {
	written := 0
	for _, target := range targets {
		body, err := iface.EncodeObject(updates, target.Layer)
		if err != nil {
			return written, err
		}
		if err := target.Writer.WriteEncrypted(ctx, gmtproto.EncryptedMessage{
			AuthKeyId: target.AuthKeyId,
			SessionId: target.SessionId,
			MsgId:     gmtproto.NextServerMsgId(0),
			SeqNo:     target.Writer.NextSeqNo(true),
			Body:      body,
		}); err != nil {
			return written, err
		}
		written++
	}
	return written, nil
}

func newestTargetByAuthKey(targets []LocalTarget) []LocalTarget {
	if len(targets) <= 1 {
		return targets
	}
	byAuthKey := make(map[int64]LocalTarget, len(targets))
	for _, target := range targets {
		current, ok := byAuthKey[target.AuthKeyId]
		if !ok || target.registerSeq > current.registerSeq {
			byAuthKey[target.AuthKeyId] = target
		}
	}
	if len(byAuthKey) == len(targets) {
		return targets
	}
	out := make([]LocalTarget, 0, len(byAuthKey))
	for _, target := range byAuthKey {
		out = append(out, target)
	}
	return out
}

func isNormalAuthKeyType(authKeyType int32) bool {
	return authKeyType == tg.AuthKeyTypePerm || authKeyType == tg.AuthKeyTypeTemp
}
