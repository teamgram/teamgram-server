package svc

import (
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
	"github.com/zeromicro/go-zero/core/discov"
)

func TestHasSessionClient(t *testing.T) {
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
			name: "dest_service_only",
			conf: kitex.RpcClientConf{DestService: "interface.session"},
			want: true,
		},
		{
			name: "direct_endpoints",
			conf: kitex.RpcClientConf{Endpoints: []string{"127.0.0.1:20120"}},
			want: true,
		},
		{
			name: "etcd",
			conf: kitex.RpcClientConf{Etcd: discov.EtcdConf{
				Hosts: []string{"127.0.0.1:2379"},
				Key:   "interface.session",
			}},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasSessionClient(tt.conf); got != tt.want {
				t.Fatalf("hasSessionClient(%+v) = %v, want %v", tt.conf, got, tt.want)
			}
		})
	}
}
