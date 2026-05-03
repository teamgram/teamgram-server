package repository

import (
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/users/internal/config"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
)

func TestNewRepositoryLeavesUserClientNilWithoutConfig(t *testing.T) {
	r := NewRepository(config.Config{})
	if r == nil {
		t.Fatal("NewRepository returned nil")
	}
	if r.UserClient != nil {
		t.Fatal("UserClient is non-nil without client config")
	}
}

func TestHasRPCClientConfig(t *testing.T) {
	if hasRPCClientConfig(kitex.RpcClientConf{}) {
		t.Fatal("empty RpcClientConf detected as configured")
	}
	if !hasRPCClientConfig(kitex.RpcClientConf{Endpoints: []string{"127.0.0.1:1234"}}) {
		t.Fatal("Endpoints RpcClientConf not detected as configured")
	}
	if !hasRPCClientConfig(kitex.RpcClientConf{Target: "direct://127.0.0.1:1234"}) {
		t.Fatal("Target RpcClientConf not detected as configured")
	}
}
