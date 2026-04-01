package kitex

import (
	"context"
	"errors"
	"testing"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/serviceinfo"
	kitexserver "github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/service"
)

type closableClient struct {
	closed bool
}

func (c *closableClient) Call(context.Context, string, interface{}, interface{}) error {
	return nil
}

func (c *closableClient) Close() error {
	c.closed = true
	return nil
}

func TestClientResourceCloseDelegates(t *testing.T) {
	cli := &closableClient{}
	resource := &client2{Client: cli}

	if err := resource.Close(); err != nil {
		t.Fatalf("close failed: %v", err)
	}
	if !cli.closed {
		t.Fatal("expected underlying client to be closed")
	}
}

func TestNewClientReturnsResolverError(t *testing.T) {
	orig := newEtcdResolver
	defer func() { newEtcdResolver = orig }()
	newEtcdResolver = func([]string, ...etcd.Option) (discovery.Resolver, error) {
		return nil, errors.New("resolver failed")
	}

	_, err := NewClient(RpcClientConf{
		DestService: "svc.test",
		ServiceName: "svc.test",
		Etcd: discov.EtcdConf{
			Hosts: []string{"127.0.0.1:2379"},
			Key:   "svc.test",
		},
	}, func(opts ...client.Option) (Client, error) {
		t.Fatal("new client should not be called when resolver setup fails")
		return nil, nil
	})
	if err == nil {
		t.Fatal("expected resolver error")
	}
}

func TestNewServerReturnsRegistryError(t *testing.T) {
	orig := newEtcdRegistry
	defer func() { newEtcdRegistry = orig }()
	newEtcdRegistry = func([]string, ...etcd.Option) (registry.Registry, error) {
		return nil, errors.New("registry failed")
	}

	_, err := NewServer(RpcServerConf{
		ServiceConf: service.ServiceConf{Name: "svc.test"},
		ListenOn:    "127.0.0.1:0",
		Etcd: EtcdConf{
			EtcdConf: discov.EtcdConf{
				Hosts: []string{"127.0.0.1:2379"},
				Key:   "svc.test",
			},
		},
	}, func(server kitexserver.Server) error {
		t.Fatal("register should not be called when registry setup fails")
		return nil
	})
	if err == nil {
		t.Fatal("expected registry error")
	}
}

func TestNewCachedClientReturnsCreationError(t *testing.T) {
	orig := newClientWithServiceInfo
	defer func() { newClientWithServiceInfo = orig }()
	newClientWithServiceInfo = func(RpcClientConf, *serviceinfo.ServiceInfo) (Client, error) {
		return nil, errors.New("client create failed")
	}
	iface.RegisterKitexServiceInfoForClient("svc.test.cached", &serviceinfo.ServiceInfo{})

	_, err := newCachedClient(RpcClientConf{
		ServiceName: "svc.test.cached",
		Endpoints:   []string{"127.0.0.1:1"},
	})
	if err == nil {
		t.Fatal("expected cached client creation error")
	}
}
