package repository

import (
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
	"github.com/zeromicro/go-zero/core/discov"
)

func TestHasRPCClientConfigRequiresServiceAndRoute(t *testing.T) {
	tests := []struct {
		name string
		conf kitex.RpcClientConf
		want bool
	}{
		{
			name: "empty",
			conf: kitex.RpcClientConf{},
			want: false,
		},
		{
			name: "service only",
			conf: kitex.RpcClientConf{DestService: "idgen"},
			want: false,
		},
		{
			name: "direct endpoints",
			conf: kitex.RpcClientConf{DestService: "idgen", Endpoints: []string{"127.0.0.1:20040"}},
			want: true,
		},
		{
			name: "target",
			conf: kitex.RpcClientConf{DestService: "idgen", Target: "127.0.0.1:20040"},
			want: true,
		},
		{
			name: "etcd",
			conf: kitex.RpcClientConf{
				DestService: "idgen",
				Etcd:        discov.EtcdConf{Hosts: []string{"127.0.0.1:2379"}, Key: "teamgram/idgen"},
			},
			want: true,
		},
		{
			name: "route without service",
			conf: kitex.RpcClientConf{Endpoints: []string{"127.0.0.1:20040"}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasRPCClientConfig(tt.conf); got != tt.want {
				t.Fatalf("hasRPCClientConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
