package alloc

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/teamgram/marmota/pkg/stores/kv"
)

type evalStore interface {
	EvalCtx(ctx context.Context, script, key string, args ...any) (any, error)
}

type xkvCache struct {
	kv       evalStore
	lockTTL  time.Duration
	dataTTL  time.Duration
	nowMilli func() int64
}

func NewXKVCache(kv kv.ExtStore) Cache {
	return NewXKVCacheWithStore(kv)
}

func NewXKVCacheWithStore(store evalStore) Cache {
	return &xkvCache{
		kv:       store,
		lockTTL:  3 * time.Second,
		dataTTL:  365 * 24 * time.Hour,
		nowMilli: func() int64 { return time.Now().UnixMilli() },
	}
}

func (c *xkvCache) Malloc(ctx context.Context, key string, size int64) ([]int64, error) {
	res, err := c.kv.EvalCtx(
		ctx,
		mallocScript,
		key,
		size,
		int64(c.lockTTL/time.Second),
		int64(c.dataTTL/time.Second),
		c.nowMilli(),
	)
	if err != nil {
		return nil, err
	}
	return toInt64Slice(res)
}

func (c *xkvCache) SetSeq(ctx context.Context, key string, owner, currSeq, lastSeq, mill int64) (int64, error) {
	if lastSeq < currSeq {
		return 0, errorsNew("lastSeq must be greater than or equal to currSeq")
	}
	res, err := c.kv.EvalCtx(
		ctx,
		setSeqScript,
		key,
		owner,
		int64(c.dataTTL/time.Second),
		currSeq,
		lastSeq,
		mill,
	)
	if err != nil {
		return 0, err
	}
	return toInt64(res)
}

func toInt64Slice(v any) ([]int64, error) {
	switch x := v.(type) {
	case []int64:
		return x, nil
	case []any:
		out := make([]int64, 0, len(x))
		for _, item := range x {
			n, err := toInt64(item)
			if err != nil {
				return nil, err
			}
			out = append(out, n)
		}
		return out, nil
	default:
		return nil, fmt.Errorf("alloc: invalid int64 slice type %T", v)
	}
}

func toInt64(v any) (int64, error) {
	switch x := v.(type) {
	case int:
		return int64(x), nil
	case int8:
		return int64(x), nil
	case int16:
		return int64(x), nil
	case int32:
		return int64(x), nil
	case int64:
		return x, nil
	case uint:
		return int64(x), nil
	case uint8:
		return int64(x), nil
	case uint16:
		return int64(x), nil
	case uint32:
		return int64(x), nil
	case uint64:
		return int64(x), nil
	case string:
		return strconv.ParseInt(x, 10, 64)
	case []byte:
		return strconv.ParseInt(string(x), 10, 64)
	default:
		return 0, fmt.Errorf("alloc: invalid int64 type %T", v)
	}
}

func errorsNew(s string) error {
	return fmt.Errorf("alloc: %s", s)
}
