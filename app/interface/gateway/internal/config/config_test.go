package config

import (
	"path/filepath"
	"slices"
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

func TestGatewayExampleConfigUsesUnifiedBFF(t *testing.T) {
	var c Config
	path := filepath.Join("..", "..", "etc", "gateway.yaml")
	if err := conf.Load(path, &c); err != nil {
		t.Fatalf("load gateway config %s: %v", path, err)
	}

	if got := len(c.BffClient.Clients); got != 1 {
		t.Fatalf("BffClient.Clients length = %d, want 1 unified bff.bff client", got)
	}
	client := c.BffClient.Clients[0]
	if client.DestService != "bff.bff" {
		t.Fatalf("BFF DestService = %q, want bff.bff", client.DestService)
	}
	if client.Etcd.Key != "bff.bff" {
		t.Fatalf("BFF Etcd.Key = %q, want bff.bff", client.Etcd.Key)
	}

	for _, serviceName := range []string{"RPCConfiguration", "RPCAuthorization", "RPCQrCode", "RPCDialogs", "RPCMessages"} {
		if !slices.Contains(client.ServiceNameList, serviceName) {
			t.Fatalf("ServiceNameList missing %s: %#v", serviceName, client.ServiceNameList)
		}
	}
}
