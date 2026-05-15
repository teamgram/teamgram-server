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
	UserId        int64
	PermAuthKeyId int64
	AuthKeyId     int64
	AuthKeyType   int32
	SessionId     int64
	Layer         int32
	AuthKey       *crypto.AuthKey
	Writer        SessionWriter
	MainUpdates   bool
	registerSeq   uint64
}

type SessionUpdatesResult struct {
	OK            bool
	Reason        string
	PermAuthKeyId int64
	AuthKeyId     int64
	AuthKeyType   int32
	SessionId     int64
	UpdatesClass  string
}

type LocalWriter struct {
	mu          sync.RWMutex
	targets     map[sessionKey]LocalTarget
	registerSeq uint64
	policy      GenericUpdatesPolicy
}

type sessionKey struct {
	authKeyId   int64
	authKeyType int32
	sessionId   int64
}

type GenericUpdatesWriter = func(ctx context.Context, permAuthKeyId int64, updates tg.UpdatesClazz) (int, error)

type GenericUpdatesPolicy interface {
	HandleGenericUpdatesWithWriter(ctx context.Context, permAuthKeyId int64, updates tg.UpdatesClazz, write GenericUpdatesWriter) (int, error)
}

type genericUpdatesWriterBinder interface {
	SetGenericUpdatesWriter(GenericUpdatesWriter)
}

var allowedExactSessionUpdateNames = map[string]struct{}{
	tg.ClazzName_updateLoginToken:      {},
	tg.ClazzName_updateSentPhoneCode:   {},
	tg.ClazzName_updateDcOptions:       {},
	tg.ClazzName_updateConfig:          {},
	tg.ClazzName_updateLangPackTooLong: {},
	tg.ClazzName_updateLangPack:        {},
}

func NewLocalWriter() *LocalWriter {
	return &LocalWriter{targets: make(map[sessionKey]LocalTarget)}
}

func (w *LocalWriter) SetGenericUpdatesPolicy(policy GenericUpdatesPolicy) {
	if w == nil {
		return
	}
	w.mu.Lock()
	w.policy = policy
	w.mu.Unlock()
	if binder, ok := policy.(genericUpdatesWriterBinder); ok {
		binder.SetGenericUpdatesWriter(w.writeUpdatesDirect)
	}
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
	if target.MainUpdates && isNormalAuthKeyType(target.AuthKeyType) {
		for key, existing := range w.targets {
			if existing.PermAuthKeyId == target.PermAuthKeyId && existing.MainUpdates && isNormalAuthKeyType(key.authKeyType) {
				existing.MainUpdates = false
				w.targets[key] = existing
			}
		}
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
	w.mu.RLock()
	policy := w.policy
	w.mu.RUnlock()
	if policy != nil {
		return policy.HandleGenericUpdatesWithWriter(ctx, permAuthKeyId, updates, w.writeUpdatesDirect)
	}
	return w.writeUpdatesDirect(ctx, permAuthKeyId, updates)
}

func (w *LocalWriter) writeUpdatesDirect(ctx context.Context, permAuthKeyId int64, updates tg.UpdatesClazz) (int, error) {
	if w == nil || updates == nil {
		return 0, nil
	}
	targets := w.matchTargets(func(key sessionKey, target LocalTarget) bool {
		return target.PermAuthKeyId == permAuthKeyId &&
			target.MainUpdates &&
			isNormalAuthKeyType(key.authKeyType)
	})
	return writeUpdatesToTargets(ctx, targets, updates)
}

func (w *LocalWriter) WriteSessionUpdates(ctx context.Context, authKeyId int64, sessionId int64, updates tg.UpdatesClazz) (bool, error) {
	result, err := w.WriteSessionUpdatesDetailed(ctx, authKeyId, sessionId, updates)
	return result.OK, err
}

func (w *LocalWriter) WriteSessionUpdatesDetailed(ctx context.Context, authKeyId int64, sessionId int64, updates tg.UpdatesClazz) (SessionUpdatesResult, error) {
	result := SessionUpdatesResult{
		AuthKeyId:    authKeyId,
		SessionId:    sessionId,
		UpdatesClass: updatesClassName(updates),
	}
	if w == nil || updates == nil {
		result.Reason = "missing_writer_or_updates"
		return result, nil
	}
	targets := w.matchTargets(func(key sessionKey, target LocalTarget) bool {
		return key.authKeyId == authKeyId && key.sessionId == sessionId
	})
	if len(targets) != 1 {
		result.Reason = "no_local_session"
		return result, nil
	}
	target := targets[0]
	result.PermAuthKeyId = target.PermAuthKeyId
	result.AuthKeyType = target.AuthKeyType
	if target.AuthKeyType == tg.AuthKeyTypeMediaTemp {
		result.Reason = "media_temp_exact_session_update_rejected"
		return result, nil
	}
	if target.UserId <= 0 && !isAllowedExactSessionUpdates(updates) {
		result.Reason = "exact_session_update_not_allowed"
		return result, nil
	}
	_, err := writeUpdatesToTargets(ctx, targets, updates)
	result.OK = err == nil
	if err != nil {
		result.Reason = "write_failed"
	}
	return result, err
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

func isNormalAuthKeyType(authKeyType int32) bool {
	return authKeyType == tg.AuthKeyTypePerm || authKeyType == tg.AuthKeyTypeTemp
}

func isAllowedExactSessionUpdates(updates tg.UpdatesClazz) bool {
	if updates == nil {
		return false
	}
	switch updates.UpdatesClazzName() {
	case tg.ClazzName_updateShort:
		short, ok := updates.(*tg.TLUpdateShort)
		return ok && isAllowedExactSessionUpdate(short.Update)
	case tg.ClazzName_updates:
		group, ok := updates.(*tg.TLUpdates)
		return ok && areAllowedExactSessionUpdates(group.Updates)
	case tg.ClazzName_updatesCombined:
		group, ok := updates.(*tg.TLUpdatesCombined)
		return ok && areAllowedExactSessionUpdates(group.Updates)
	default:
		return false
	}
}

func areAllowedExactSessionUpdates(updates []tg.UpdateClazz) bool {
	for _, update := range updates {
		if !isAllowedExactSessionUpdate(update) {
			return false
		}
	}
	return true
}

func isAllowedExactSessionUpdate(update tg.UpdateClazz) bool {
	if update == nil {
		return false
	}
	_, ok := allowedExactSessionUpdateNames[update.UpdateClazzName()]
	return ok
}

func updatesClassName(updates tg.UpdatesClazz) string {
	if updates == nil {
		return ""
	}
	return updates.UpdatesClazzName()
}
