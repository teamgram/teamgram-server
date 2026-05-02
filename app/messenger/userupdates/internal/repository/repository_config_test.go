package repository

import (
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/config"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
)

func TestNewRepositoryConfiguresIdgenClient(t *testing.T) {
	repo := NewRepository(config.Config{
		Idgen: kitex.RpcClientConf{
			DestService: "service.idgen",
			Endpoints:   []string{"127.0.0.1:20660"},
		},
	})

	if _, ok := repo.idgen.(*idgenRPCGenerator); !ok {
		t.Fatalf("idgen = %T, want *idgenRPCGenerator", repo.idgen)
	}
}

func TestNewRepositoryUsesUnavailableIdgenWithoutClientConfig(t *testing.T) {
	repo := NewRepository(config.Config{})

	if _, ok := repo.idgen.(unavailableIDGenerator); !ok {
		t.Fatalf("idgen = %T, want unavailableIDGenerator", repo.idgen)
	}
}
