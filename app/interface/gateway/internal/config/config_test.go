package config

import (
	"path/filepath"
	"testing"

	"github.com/zeromicro/go-zero/core/conf"
)

func TestGatewayExampleConfigLoads(t *testing.T) {
	var c Config
	path := filepath.Join("..", "..", "etc", "gateway.yaml")
	if err := conf.Load(path, &c); err != nil {
		t.Fatalf("load gateway config %s: %v", path, err)
	}
	if c.Transport.TCPListenOn == "" {
		t.Fatal("Transport.TCPListenOn is empty")
	}
}
