package repository

import (
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	minioadapter "github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/minio"
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

func TestMapObjectReadErrorMapsMinioMissToDfsFileNotFound(t *testing.T) {
	err := mapObjectReadError("get photo file", minioadapter.ErrObjectNotFound)
	if !errors.Is(err, dfs.ErrDfsFileNotFound) {
		t.Fatalf("mapObjectReadError() error = %v, want ErrDfsFileNotFound", err)
	}
	if errors.Is(err, dfs.ErrDfsStorage) {
		t.Fatalf("mapObjectReadError() error = %v, should not be ErrDfsStorage", err)
	}
}

func TestMapObjectReadErrorMapsOrdinaryErrorToStorage(t *testing.T) {
	cause := errors.New("connection reset")
	err := mapObjectReadError("get photo file", cause)
	if !errors.Is(err, dfs.ErrDfsStorage) {
		t.Fatalf("mapObjectReadError() error = %v, want ErrDfsStorage", err)
	}
	if !errors.Is(err, cause) {
		t.Fatalf("mapObjectReadError() error = %v, want original cause preserved", err)
	}
}
