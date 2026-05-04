package config

import (
	"testing"

	"github.com/teamgram/marmota/pkg/stores/kv"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func validMinimalConfig() Config {
	return Config{
		KV: kv.KvConf{
			cache.NodeConf{
				RedisConf: redis.RedisConf{
					Host: "127.0.0.1:6379",
					Type: "node",
				},
				Weight: 1,
			},
		},
		SessionExpiresSeconds:                  60,
		HashKeyTTLSeconds:                      600,
		CleanupOnWriteIntervalSeconds:          60,
		PresenceQueryDefaultQPSPerCaller:       50,
		PresenceGatewayDiagnosticsQPSPerCaller: 1,
	}
}

func TestValidateRejectsEmptyKV(t *testing.T) {
	cfg := validMinimalConfig()
	cfg.KV = nil
	if err := cfg.Validate(); err == nil {
		t.Fatal("Validate() error = nil, want error")
	}
}

func TestValidateRejectsAllZeroKVWeights(t *testing.T) {
	cfg := validMinimalConfig()
	cfg.KV[0].Weight = 0
	if err := cfg.Validate(); err == nil {
		t.Fatal("Validate() error = nil, want error")
	}
}

func TestValidateAcceptsValidMinimalConfig(t *testing.T) {
	cfg := validMinimalConfig()
	if err := cfg.Validate(); err != nil {
		t.Fatalf("Validate() error = %v, want nil", err)
	}
}
