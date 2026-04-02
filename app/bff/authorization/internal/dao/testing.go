package dao

import "github.com/zeromicro/go-zero/core/stores/kv"

// NewWithKVStore creates a Dao backed by the provided KV store.
func NewWithKVStore(store kv.Store) *Dao {
	return &Dao{kv: store}
}
