package repository

import (
	"testing"

	"github.com/teamgram/marmota/pkg/stores/kv"
)

func TestRepositoryUsesMarmotaExtStore(t *testing.T) {
	var r Repository
	var _ kv.ExtStore = r.kv
}
