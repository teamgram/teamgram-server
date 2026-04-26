package xkv

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/kv"
)

const (
	cacheSaltPrefix = "salts"
)

const (
	saltTimeout = 30 * 60
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

func NewFutureSaltsModel(kv kv.Store, prefix string) FutureSaltsModel {
	return &futureSaltsModel{
		kv:     kv,
		prefix: prefix,
	}
}

func genCacheSaltKey(id int64) string {
	return fmt.Sprintf("%s#%d", cacheSaltPrefix, id)
}

func (m *futureSaltsModel) formatKey(key string) string {
	return m.prefix + ":" + key
}

func (m *futureSaltsModel) PutSalts(ctx context.Context, keyId int64, salts []*FutureSaltRecord) (err error) {
	var (
		b   []byte
		key = genCacheSaltKey(keyId)
	)

	if b, err = json.Marshal(salts); err != nil {
		return
	}

	// 误差 500
	if err = m.kv.SetexCtx(ctx, key, string(b), len(salts)*saltTimeout); err != nil {
		return
	}

	return
}

func (m *futureSaltsModel) GetSalts(ctx context.Context, keyId int64) (salts []*FutureSaltRecord, err error) {
	var (
		key  = genCacheSaltKey(keyId)
		bBuf string
	)

	bBuf, err = m.kv.GetCtx(ctx, key)
	if err != nil {
		return
	} else if bBuf == "" {
		return nil, nil
	}

	if err = jsonx.UnmarshalFromString(bBuf, &salts); err != nil {
		logx.WithContext(ctx).Errorf("getSalts jsonx.UnmarshalFromString(%s) error(%v)", bBuf, err)
		return nil, nil
	}

	return
}

func (m *futureSaltsModel) DeleteSalts(ctx context.Context, keyId int64) error {
	var (
		key = genCacheSaltKey(keyId)
	)

	_, err := m.kv.DelCtx(ctx, key)

	return err
}
