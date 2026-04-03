package server

import (
	"testing"

	bffconfig "github.com/teamgram/teamgram-server/v2/app/bff/bff/internal/config"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func TestBuildAuthorizationConfigForwardsRuntimeDependencies(t *testing.T) {
	src := bffconfig.Config{
		RpcServerConf: kitex.RpcServerConf{
			ListenOn: "127.0.0.1:20010",
		},
		KV: kv.KvConf{
			{
				RedisConf: redis.RedisConf{
					Host: "127.0.0.1:6379",
				},
			},
		},
		BizServiceClient: kitex.RpcClientConf{
			DestService: "service.biz_service",
			ServiceName: "service.biz_service",
		},
		AuthSessionClient: kitex.RpcClientConf{
			DestService: "service.authsession",
			ServiceName: "service.authsession",
		},
		StatusClient: kitex.RpcClientConf{
			DestService: "service.status",
			ServiceName: "service.status",
		},
		MsgClient: kitex.RpcClientConf{
			DestService: "messenger.msg",
			ServiceName: "messenger.msg",
		},
	}

	got := buildAuthorizationConfig(src)

	if got.ListenOn != src.ListenOn {
		t.Fatalf("expected listen_on %q, got %q", src.ListenOn, got.ListenOn)
	}
	if len(got.KV) != 1 || got.KV[0].Host != "127.0.0.1:6379" {
		t.Fatalf("expected kv hosts to be forwarded, got %#v", got.KV)
	}
	if got.UserClient.DestService != "service.biz_service" {
		t.Fatalf("expected user client dest service to use biz client, got %#v", got.UserClient)
	}
	if got.UserClient.ServiceName != "RPCUser" {
		t.Fatalf("expected user client service name RPCUser, got %#v", got.UserClient)
	}
	if got.ChatClient.DestService != "service.biz_service" {
		t.Fatalf("expected chat client dest service to use biz client, got %#v", got.ChatClient)
	}
	if got.ChatClient.ServiceName != "RPCChat" {
		t.Fatalf("expected chat client service name RPCChat, got %#v", got.ChatClient)
	}
	if got.UsernameClient.DestService != "service.biz_service" {
		t.Fatalf("expected username client dest service to use biz client, got %#v", got.UsernameClient)
	}
	if got.AuthsessionClient.DestService != "service.authsession" {
		t.Fatalf("expected authsession client to be forwarded, got %#v", got.AuthsessionClient)
	}
	if got.AuthsessionClient.ServiceName != "RPCAuthsession" {
		t.Fatalf("expected authsession client service name RPCAuthsession, got %#v", got.AuthsessionClient)
	}
	if got.StatusClient.DestService != "service.status" {
		t.Fatalf("expected status client to be forwarded, got %#v", got.StatusClient)
	}
	if got.StatusClient.ServiceName != "RPCStatus" {
		t.Fatalf("expected status client service name RPCStatus, got %#v", got.StatusClient)
	}
	if got.MsgClient.DestService != "messenger.msg" {
		t.Fatalf("expected msg client to be forwarded, got %#v", got.MsgClient)
	}
	if got.MsgClient.ServiceName != "RPCMsg" {
		t.Fatalf("expected msg client service name RPCMsg, got %#v", got.MsgClient)
	}
}
