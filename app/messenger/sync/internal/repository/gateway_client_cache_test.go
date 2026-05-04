package repository

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
)

func TestGatewayClientCacheRevalidatesAfterTTL(t *testing.T) {
	validations := 0
	policy := allowLocalhostPolicy()
	policy.Resolver = func(ctx context.Context, host string) ([]net.IP, error) {
		validations++
		return []net.IP{net.ParseIP("127.0.0.1")}, nil
	}
	cache := NewGatewayClientCache(kitex.RpcClientConf{}, policy, 10, time.Second, time.Minute)
	now := time.Unix(100, 0)
	cache.now = func() time.Time { return now }
	cache.newPusher = func(conf kitex.RpcClientConf) (GatewayPusher, error) {
		return &fakeGatewayPusher{}, nil
	}

	if _, err := cache.Get(context.Background(), "localhost:20110"); err != nil {
		t.Fatalf("first Get() error = %v", err)
	}
	if _, err := cache.Get(context.Background(), "localhost:20110"); err != nil {
		t.Fatalf("second Get() error = %v", err)
	}
	now = now.Add(time.Second)
	if _, err := cache.Get(context.Background(), "localhost:20110"); err != nil {
		t.Fatalf("third Get() error = %v", err)
	}
	if validations != 2 {
		t.Fatalf("validations = %d, want 2", validations)
	}
}

func TestGatewayClientCacheDialsPinnedValidatedIP(t *testing.T) {
	policy := AddressPolicy{
		AllowedCIDRs: []string{"10.0.0.0/8"},
		AllowedPorts: []int{20110},
		Resolver: func(ctx context.Context, host string) ([]net.IP, error) {
			if host != "gateway.internal" {
				t.Fatalf("resolver host = %q, want gateway.internal", host)
			}
			return []net.IP{net.ParseIP("10.20.30.40")}, nil
		},
	}
	cache := NewGatewayClientCache(kitex.RpcClientConf{}, policy, 10, time.Minute, time.Minute)
	var endpoint string
	cache.newPusher = func(conf kitex.RpcClientConf) (GatewayPusher, error) {
		if len(conf.Endpoints) != 1 {
			t.Fatalf("Endpoints = %v, want one endpoint", conf.Endpoints)
		}
		endpoint = conf.Endpoints[0]
		return &fakeGatewayPusher{}, nil
	}

	if _, err := cache.Get(context.Background(), "gateway.internal:20110"); err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if endpoint != "10.20.30.40:20110" {
		t.Fatalf("endpoint = %q, want pinned validated IP", endpoint)
	}
}

func TestGatewayClientCacheEvictsLRU(t *testing.T) {
	cache := NewGatewayClientCache(kitex.RpcClientConf{}, allowLocalhostPolicy(), 1, time.Minute, time.Minute)
	cache.newPusher = func(conf kitex.RpcClientConf) (GatewayPusher, error) {
		return &fakeGatewayPusher{}, nil
	}

	if _, err := cache.Get(context.Background(), "127.0.0.1:20110"); err != nil {
		t.Fatalf("first Get() error = %v", err)
	}
	if _, err := cache.Get(context.Background(), "127.0.0.2:20110"); err != nil {
		t.Fatalf("second Get() error = %v", err)
	}
	if len(cache.entries) != 1 {
		t.Fatalf("entries = %d, want 1", len(cache.entries))
	}
	if _, ok := cache.entries["127.0.0.1:20110"]; ok {
		t.Fatal("oldest entry was not evicted")
	}
}
