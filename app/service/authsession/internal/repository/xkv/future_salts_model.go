package xkv

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/kv"
)

const (
	cacheSaltPrefix = "salts"
)

const (
	saltTimeout = 30 * 60 // salt timeout
)

type FutureSaltsModel interface {
	PutSalts(ctx context.Context, keyId int64, salts []*tg.TLFutureSalt) (err error)
	GetSalts(ctx context.Context, keyId int64) (salts []*tg.TLFutureSalt, err error)
	PutSaltCache(ctx context.Context, keyId int64, salt *tg.TLFutureSalt) error
	GetFutureSalts(ctx context.Context, authKeyId int64, num int32) (*tg.TLFutureSalts, error)
	DeleteSalts(ctx context.Context, keyId int64) error
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

func (m *futureSaltsModel) getOrNotInsertSaltList(ctx context.Context, keyId int64, size int32) ([]*tg.TLFutureSalt, error) {
	var (
		salts = make([]*tg.TLFutureSalt, 0, size)

		date           = int32(time.Now().Unix())
		lastValidUntil = date
		saltsData      []*tg.TLFutureSalt
		lastSalt       *tg.TLFutureSalt
	)

	saltList, err := m.GetSalts(ctx, keyId)
	if err != nil {
		return nil, err
	}

	if len(saltList) > 0 {
		hasLastSalt := false
		for idx, salt := range saltList {
			if salt.ValidSince >= date {
				if !hasLastSalt {
					if idx > 0 {
						lastSalt = saltList[idx-1]
					}
					hasLastSalt = true
				}
				saltsData = append(saltsData, salt)
				if lastValidUntil < salt.ValidUntil {
					lastValidUntil = salt.ValidUntil
				}
			}
		}
		if !hasLastSalt {
			lastSalt = saltList[len(saltList)-1]
		}

		// check ValidUntil
		if lastSalt != nil && lastSalt.ValidUntil+300 < date {
			lastSalt = nil
		}
	}

	left := size - int32(len(saltsData))
	if left > 0 {
		for i := int32(0); i < size; i++ {
			salt := tg.MakeTLFutureSalt(&tg.TLFutureSalt{
				ValidSince: lastValidUntil,
				ValidUntil: lastValidUntil + saltTimeout,
				Salt:       rand.Int63(),
			})
			saltsData = append(saltsData, salt)
			lastValidUntil += saltTimeout
		}
	}

	for i := int32(0); i < size; i++ {
		salts = append(salts, saltsData[i])
	}

	var (
		salts2     []*tg.TLFutureSalt
		saltsData2 []*tg.TLFutureSalt
	)

	if lastSalt != nil {
		salts2 = append(salts2, lastSalt)
		saltsData2 = append(saltsData2, lastSalt)
	}

	salts2 = append(salts2, salts...)
	saltsData2 = append(saltsData2, saltsData...)

	if left > 0 {
		err = m.PutSalts(ctx, keyId, saltsData2)
		if err != nil {
			return nil, err
		}
	}
	return salts2, nil
}

func (m *futureSaltsModel) PutSalts(ctx context.Context, keyId int64, salts []*tg.TLFutureSalt) (err error) {
	var (
		b   []byte
		key = genCacheSaltKey(keyId)
	)

	if b, err = json.Marshal(salts); err != nil {
		logx.WithContext(ctx).Errorf("conn.SETEX(%s) error(%v)", key, err)
		return
	}

	// 误差 500
	if err = m.kv.SetexCtx(ctx, key, string(b), len(salts)*saltTimeout); err != nil {
		logx.WithContext(ctx).Errorf("conn.SETEX(%s) error(%v)", key, err)
	}

	return
}

func (m *futureSaltsModel) GetSalts(ctx context.Context, keyId int64) (salts []*tg.TLFutureSalt, err error) {
	var (
		key  = genCacheSaltKey(keyId)
		bBuf string
	)

	bBuf, err = m.kv.GetCtx(ctx, key)
	if err != nil {
		logx.WithContext(ctx).Errorf("conn.Do(GET %s) error(%v)", key, err)
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

func (m *futureSaltsModel) PutSaltCache(ctx context.Context, keyId int64, salt *tg.TLFutureSalt) error {
	return m.PutSalts(ctx, keyId, []*tg.TLFutureSalt{salt})
}

func (m *futureSaltsModel) GetFutureSalts(ctx context.Context, authKeyId int64, num int32) (*tg.TLFutureSalts, error) {
	pSalts, err := m.getOrNotInsertSaltList(ctx, authKeyId, num)
	if err != nil {
		return nil, err
	}
	salts := tg.MakeTLFutureSalts(&tg.TLFutureSalts{
		ReqMsgId: 0,
		Now:      0,
		Salts:    pSalts,
	})

	return salts, nil
}

func (m *futureSaltsModel) DeleteSalts(ctx context.Context, keyId int64) error {
	var (
		key = genCacheSaltKey(keyId)
	)

	_, err := m.kv.DelCtx(ctx, key)

	return err
}
