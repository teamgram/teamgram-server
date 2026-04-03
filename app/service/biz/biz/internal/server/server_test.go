package server

import (
	"testing"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	bizconfig "github.com/teamgram/teamgram-server/v2/app/service/biz/biz/internal/config"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func TestBuildUserConfigForwardsStorageAndMediaClient(t *testing.T) {
	src := bizconfig.Config{
		RpcServerConf: kitex.RpcServerConf{
			ListenOn: "127.0.0.1:20020",
		},
		Mysql: sqlx.Config{
			DSN: "root:@tcp(127.0.0.1:3306)/teamgram",
		},
		Cache: cache.CacheConf{
			{
				RedisConf: redis.RedisConf{
					Host: "127.0.0.1:6379",
				},
			},
		},
		MediaClient: kitex.RpcClientConf{
			DestService: "service.media",
			ServiceName: "service.media",
		},
	}

	got := buildUserConfig(src)

	if got.ListenOn != src.ListenOn {
		t.Fatalf("expected listen_on %q, got %q", src.ListenOn, got.ListenOn)
	}
	if got.Mysql.DSN != src.Mysql.DSN {
		t.Fatalf("expected mysql dsn to be forwarded, got %#v", got.Mysql)
	}
	if len(got.Cache) != 1 || got.Cache[0].Host != "127.0.0.1:6379" {
		t.Fatalf("expected cache hosts to be forwarded, got %#v", got.Cache)
	}
	if got.MediaClient.DestService != "service.media" {
		t.Fatalf("expected media client dest service to be forwarded, got %#v", got.MediaClient)
	}
	if got.MediaClient.ServiceName != "RPCMedia" {
		t.Fatalf("expected media client service name RPCMedia, got %#v", got.MediaClient)
	}
}
