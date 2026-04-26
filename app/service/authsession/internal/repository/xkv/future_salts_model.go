package xkv

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/kv"
)

const (
	// saltMinTTL keeps an entry alive for at least one salt window so that
	// callers can still observe the current/just-expired salt while the next
	// rotation happens.
	saltMinTTL = 30 * 60

	// saltGraceTTL is the additional time we keep the cache valid after the
	// last salt expires, so the previous-salt compatibility window in the
	// repository layer still has data to read from.
	saltGraceTTL = 5 * 60
)

type FutureSaltsModel interface {
	PutSalts(ctx context.Context, keyId int64, salts []*FutureSaltRecord) (err error)
	GetSalts(ctx context.Context, keyId int64) (salts []*FutureSaltRecord, err error)
	DeleteSalts(ctx context.Context, keyId int64) error
}

type FutureSaltRecord struct {
	ValidSince int32 `json:"valid_since"`
	ValidUntil int32 `json:"valid_until"`
	Salt       int64 `json:"salt"`
}

type futureSaltsModel struct {
	kv     kv.Store
	prefix string
}

// NewFutureSaltsModel builds a kv-backed model. prefix is prepended to every
// key so multiple services can share the same kv store without collisions.
func NewFutureSaltsModel(kv kv.Store, prefix string) FutureSaltsModel {
	return &futureSaltsModel{
		kv:     kv,
		prefix: prefix,
	}
}

func (m *futureSaltsModel) cacheKey(id int64) string {
	if m.prefix == "" {
		return fmt.Sprintf("salts#%d", id)
	}
	return fmt.Sprintf("%s:salts#%d", m.prefix, id)
}

// nowUnix is overridable in tests.
var nowUnix = func() int32 { return int32(time.Now().Unix()) }

// computeSaltsTTL derives the cache TTL from the salt set so the cache stays
// alive until the last salt expires (plus a small grace window). An empty set
// is treated as "delete".
func computeSaltsTTL(salts []*FutureSaltRecord, now int32) int {
	if len(salts) == 0 {
		return 0
	}
	var maxValidUntil int32
	for _, s := range salts {
		if s == nil {
			continue
		}
		if s.ValidUntil > maxValidUntil {
			maxValidUntil = s.ValidUntil
		}
	}
	ttl := int(maxValidUntil-now) + saltGraceTTL
	if ttl < saltMinTTL {
		return saltMinTTL
	}
	return ttl
}

func (m *futureSaltsModel) PutSalts(ctx context.Context, keyId int64, salts []*FutureSaltRecord) (err error) {
	key := m.cacheKey(keyId)

	if len(salts) == 0 {
		_, err = m.kv.DelCtx(ctx, key)
		return
	}

	b, err := json.Marshal(salts)
	if err != nil {
		return err
	}

	ttl := computeSaltsTTL(salts, nowUnix())
	return m.kv.SetexCtx(ctx, key, string(b), ttl)
}

func (m *futureSaltsModel) GetSalts(ctx context.Context, keyId int64) (salts []*FutureSaltRecord, err error) {
	key := m.cacheKey(keyId)

	bBuf, err := m.kv.GetCtx(ctx, key)
	if err != nil {
		return nil, err
	}
	if bBuf == "" {
		return nil, nil
	}

	if err = jsonx.UnmarshalFromString(bBuf, &salts); err != nil {
		logx.WithContext(ctx).Errorf("getSalts jsonx.UnmarshalFromString(%s) error(%v)", bBuf, err)
		return nil, nil
	}

	return
}

func (m *futureSaltsModel) DeleteSalts(ctx context.Context, keyId int64) error {
	_, err := m.kv.DelCtx(ctx, m.cacheKey(keyId))
	return err
}
