package svc

import (
	"testing"

	bffproxyclient "github.com/teamgram/teamgram-server/v2/app/bff/bff/client"
	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/config"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
)

func TestNewServiceContextWiresGatewayDependencies(t *testing.T) {
	c := config.Config{
		GatewayId: "gateway-test",
		Transport: config.TransportConf{
			TCPListenOn: "127.0.0.1:0",
		},
		AuthsessionClient: kitex.RpcClientConf{
			DestService: "service.authsession",
			ServiceName: "RPCAuthsession",
			Endpoints:   []string{"127.0.0.1:1"},
		},
		BffClient: bffproxyclient.BFFProxyClientListConf{
			Clients: []bffproxyclient.BFFProxyClientConf{{
				RpcClientConf: kitex.RpcClientConf{
					DestService: "bff.configuration",
					ServiceName: "RPCConfiguration",
					Endpoints:   []string{"127.0.0.1:2"},
				},
				ServiceNameList: []string{"RPCConfiguration", "RPCDialogs", "RPCMessages"},
			}},
		},
	}

	ctx := NewServiceContext(c)
	if ctx == nil {
		t.Fatal("NewServiceContext() returned nil")
	}
	if ctx.Repo == nil {
		t.Fatal("NewServiceContext().Repo is nil")
	}
	if ctx.BFF == nil {
		t.Fatal("NewServiceContext().BFF is nil")
	}
	if ctx.Repo.AuthsessionClient == nil {
		t.Fatal("NewServiceContext().Repo.AuthsessionClient is nil")
	}
	if ctx.BFF.RawClients["RPCDialogs"] == nil {
		t.Fatal("NewServiceContext().BFF.RawClients missing RPCDialogs")
	}
	if ctx.BFF.RawClients["RPCMessages"] == nil {
		t.Fatal("NewServiceContext().BFF.RawClients missing RPCMessages")
	}
}
