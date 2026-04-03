package dao

import (
	"errors"
	"testing"

	"github.com/zeromicro/go-zero/core/discov"
)

func TestRpcShardingManagerStartIgnoresSubscriberCreationFailure(t *testing.T) {
	orig := newShardingSubscriber
	t.Cleanup(func() {
		newShardingSubscriber = orig
	})

	newShardingSubscriber = func(_ []string, _ string) (shardingSubscriber, error) {
		return nil, errors.New("boom")
	}

	mgr := NewRpcShardingManager("127.0.0.1:20120", discov.EtcdConf{
		Hosts: []string{"127.0.0.1:2379"},
		Key:   "interface.session",
	})

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("Start should not panic when subscriber creation fails: %v", r)
		}
	}()

	mgr.Start()

	if got := len(mgr.shardingList); got != 0 {
		t.Fatalf("expected sharding list to remain empty, got %d", got)
	}
}
