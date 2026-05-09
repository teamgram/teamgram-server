package repository

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/config"
	minioadapter "github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/minio"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/core/stores/redis"
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

func TestNewRepositorySelectsSpoolWhenRootDirConfigured(t *testing.T) {
	repo := NewRepository(config.Config{
		Kv: minimalKVConf(),
		UploadSpool: config.UploadSpoolConf{
			RootDir:    t.TempDir(),
			NodeIDFile: "node_id",
		},
	})
	got := reflect.TypeOf(repo.uploadStateModel).String()
	if got != "*spool.UploadStateModel" {
		t.Fatalf("uploadStateModel type = %s, want *spool.UploadStateModel", got)
	}
}

func TestNewRepositoryUsesXKVWhenSpoolRootDirEmpty(t *testing.T) {
	repo := NewRepository(config.Config{Kv: minimalKVConf()})
	got := reflect.TypeOf(repo.uploadStateModel).String()
	if got != "*xkv.uploadStateModel" {
		t.Fatalf("uploadStateModel type = %s, want *xkv.uploadStateModel", got)
	}
}

func TestRepositoryMapsSpoolNotFoundToDfsFileNotFound(t *testing.T) {
	repo := NewRepository(config.Config{
		Kv: minimalKVConf(),
		UploadSpool: config.UploadSpoolConf{
			RootDir:    t.TempDir(),
			NodeIDFile: "node_id",
		},
	})
	_, err := repo.LoadUploadFileInfo(context.Background(), 1001, 2002)
	if !errors.Is(err, dfs.ErrDfsFileNotFound) {
		t.Fatalf("LoadUploadFileInfo() error = %v, want ErrDfsFileNotFound", err)
	}
	if errors.Is(err, dfs.ErrDfsStorage) {
		t.Fatalf("LoadUploadFileInfo() error = %v, should not be ErrDfsStorage", err)
	}
}

func TestRepositoryScanSpoolOnStartNoopsForXKVBackend(t *testing.T) {
	repo := NewRepository(config.Config{Kv: minimalKVConf()})
	if err := repo.ScanSpoolOnStart(context.Background()); err != nil {
		t.Fatalf("ScanSpoolOnStart() error = %v", err)
	}
	if err := repo.CleanupExpiredUploadSessions(context.Background(), time.Unix(1_700_010_000, 0)); err != nil {
		t.Fatalf("CleanupExpiredUploadSessions() error = %v", err)
	}
	if !repo.IsUploadNodeWritable() {
		t.Fatal("IsUploadNodeWritable() = false, want true for xkv backend")
	}
	repo.MarkUploadNodeDraining("xkv drain")
	if repo.IsUploadNodeWritable() {
		t.Fatal("IsUploadNodeWritable() after xkv drain = true, want false")
	}
	err := repo.SaveUploadPart(context.Background(), UploadRef{Creator: 1001, FileID: 2002}, 0, []byte("data"))
	if !errors.Is(err, dfs.ErrDfsStorage) {
		t.Fatalf("SaveUploadPart() after xkv drain error = %v, want ErrDfsStorage", err)
	}
}

func TestRepositoryUploadNodeDrainingRejectsWritesAsStorage(t *testing.T) {
	repo := NewRepository(config.Config{
		Kv: minimalKVConf(),
		UploadSpool: config.UploadSpoolConf{
			RootDir:    t.TempDir(),
			NodeIDFile: "node_id",
		},
	})
	if !repo.IsUploadNodeWritable() {
		t.Fatal("IsUploadNodeWritable() before draining = false, want true")
	}
	preDrainInfo := &DfsFileInfo{
		Creator:        1001,
		FileID:         2002,
		FileTotalParts: 1,
		Mtime:          time.Unix(1_700_010_000, 0).Unix(),
	}
	if err := repo.SaveUploadFileInfo(context.Background(), preDrainInfo); err != nil {
		t.Fatalf("SaveUploadFileInfo() before drain error = %v", err)
	}
	if err := repo.SaveObjectCacheRef(context.Background(), 3003, preDrainInfo); err != nil {
		t.Fatalf("SaveObjectCacheRef() before drain error = %v", err)
	}

	repo.MarkUploadNodeDraining("test drain")
	if repo.IsUploadNodeWritable() {
		t.Fatal("IsUploadNodeWritable() after draining = true, want false")
	}
	err := repo.SaveUploadPart(context.Background(), UploadRef{Creator: 1001, FileID: 2002}, 0, []byte("data"))
	if !errors.Is(err, dfs.ErrDfsStorage) {
		t.Fatalf("SaveUploadPart() error = %v, want ErrDfsStorage", err)
	}
	err = repo.SaveUploadFileInfo(context.Background(), &DfsFileInfo{
		Creator:        1001,
		FileID:         2002,
		FileTotalParts: 1,
		Mtime:          time.Unix(1_700_010_000, 0).Unix(),
	})
	if !errors.Is(err, dfs.ErrDfsStorage) {
		t.Fatalf("SaveUploadFileInfo() error = %v, want ErrDfsStorage", err)
	}
	err = repo.SaveObjectCacheRef(context.Background(), 3004, preDrainInfo)
	if !errors.Is(err, dfs.ErrDfsStorage) {
		t.Fatalf("SaveObjectCacheRef() error = %v, want ErrDfsStorage", err)
	}
	info, err := repo.LoadObjectCacheRef(context.Background(), 3003)
	if err != nil {
		t.Fatalf("LoadObjectCacheRef() after draining error = %v", err)
	}
	if info.Creator != preDrainInfo.Creator || info.FileID != preDrainInfo.FileID {
		t.Fatalf("LoadObjectCacheRef() after draining = %+v, want creator/file from pre-drain ref", info)
	}
}

func minimalKVConf() kv.KvConf {
	return kv.KvConf{
		cache.NodeConf{
			RedisConf: redis.RedisConf{
				Host: "127.0.0.1:6379",
				Type: "node",
			},
			Weight: 1,
		},
	}
}
