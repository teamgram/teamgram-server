package push

import (
	"context"
	"sync"

	gmtproto "github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/mtproto"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/crypto"
)

type SessionWriter interface {
	NextSeqNo(contentRelated bool) int32
	WriteEncrypted(ctx context.Context, msg gmtproto.EncryptedMessage) error
}

type LocalTarget struct {
	AuthKeyId   int64
	AuthKeyType int32
	SessionId   int64
	AuthKey     *crypto.AuthKey
	Writer      SessionWriter
}

type LocalWriter struct {
	mu      sync.RWMutex
	targets map[sessionKey]LocalTarget
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
