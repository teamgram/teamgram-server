package alloc

import (
	"context"
	"errors"
)

var ErrMongoStoreNotImplemented = errors.New("alloc: mongo store is reserved but not implemented")

type mongoStore struct{}

func NewMongoStoreFake() SeqStore {
	return mongoStore{}
}

func (mongoStore) Malloc(ctx context.Context, key string, size int64) (int64, error) {
	return 0, ErrMongoStoreNotImplemented
}

func (mongoStore) GetMaxSeq(ctx context.Context, key string) (int64, error) {
	return 0, ErrMongoStoreNotImplemented
}

func (mongoStore) SetMaxSeq(ctx context.Context, key string, seq int64) error {
	return ErrMongoStoreNotImplemented
}
