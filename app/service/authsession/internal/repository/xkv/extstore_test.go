package xkv

import (
	"testing"

	"github.com/teamgram/marmota/pkg/stores/kv"
)

func TestFutureSaltsModelUsesMarmotaExtStore(t *testing.T) {
	var m futureSaltsModel
	var _ kv.ExtStore = m.kv
}

func TestAuthKeyLifecycleModelUsesMarmotaExtStore(t *testing.T) {
	var m authKeyLifecycleModel
	var _ kv.ExtStore = m.kv
}
