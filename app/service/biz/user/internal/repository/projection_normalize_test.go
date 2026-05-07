package repository

import (
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/internal/config"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
)

func TestNormalizeProjectionRequestDedupesAndDropsInvalidTargets(t *testing.T) {
	cfg := normalizeProjectionConfig(ProjectionConfig{})
	got, err := normalizeProjectionRequest([]int64{2, 2, 1}, []int64{0, -1, 3, 3, 2}, true, cfg)
	if err != nil {
		t.Fatalf("normalizeProjectionRequest() error = %v", err)
	}
	if want := []int64{2, 1}; !sameInt64s(got.ViewerUserIds, want) {
		t.Fatalf("viewer ids = %v, want %v", got.ViewerUserIds, want)
	}
	if want := []int64{3, 2}; !sameInt64s(got.TargetUserIds, want) {
		t.Fatalf("target ids = %v, want %v", got.TargetUserIds, want)
	}
	if want := []int64{2, 1, 3}; !sameInt64s(got.HydrateUserIds, want) {
		t.Fatalf("hydrate ids = %v, want %v", got.HydrateUserIds, want)
	}
	if !got.WithFacts {
		t.Fatalf("with facts = false, want true")
	}
}

func TestNormalizeProjectionRequestRejectsMissingViewer(t *testing.T) {
	cfg := normalizeProjectionConfig(ProjectionConfig{})
	_, err := normalizeProjectionRequest(nil, []int64{1}, false, cfg)
	if !errors.Is(err, userpb.ErrUserInvalidArgument) {
		t.Fatalf("error = %v, want %v", err, userpb.ErrUserInvalidArgument)
	}
}

func TestNormalizeProjectionRequestRejectsTooManyPairs(t *testing.T) {
	cfg := normalizeProjectionConfig(ProjectionConfig{MaxViewerUserIds: 2, MaxTargetUserIds: 4, MaxProjectionPairs: 3})
	_, err := normalizeProjectionRequest([]int64{1, 2}, []int64{3, 4}, false, cfg)
	if !errors.Is(err, userpb.ErrUserInvalidArgument) {
		t.Fatalf("error = %v, want %v", err, userpb.ErrUserInvalidArgument)
	}
}

func TestNormalizeProjectionConfigClampsSQLChunkSize(t *testing.T) {
	if got := normalizeProjectionConfig(ProjectionConfig{SQLInChunkSize: 1}); got.SQLInChunkSize != 100 {
		t.Fatalf("low SQL chunk size = %d, want 100", got.SQLInChunkSize)
	}
	if got := normalizeProjectionConfig(ProjectionConfig{SQLInChunkSize: 1001}); got.SQLInChunkSize != 1000 {
		t.Fatalf("high SQL chunk size = %d, want 1000", got.SQLInChunkSize)
	}
	if got := normalizeProjectionConfig(ProjectionConfig{}); got.SQLInChunkSize != 500 {
		t.Fatalf("default SQL chunk size = %d, want 500", got.SQLInChunkSize)
	}
}

func TestChunkInt64sUsesConfiguredSize(t *testing.T) {
	got := chunkInt64s([]int64{1, 2, 3, 4, 5}, 2)
	if len(got) != 3 || len(got[0]) != 2 || len(got[1]) != 2 || len(got[2]) != 1 || got[2][0] != 5 {
		t.Fatalf("chunks = %#v", got)
	}
}

func TestProjectionConfigDefaultsContactMapCacheEnabled(t *testing.T) {
	got := projectionConfigFromConfig(config.ProjectionConf{})
	if !got.ContactMapCacheEnabled {
		t.Fatalf("contact map cache enabled = false, want true")
	}
}

func TestProjectionConfigMapsContactMapCacheDisabled(t *testing.T) {
	got := projectionConfigFromConfig(config.ProjectionConf{ContactMapCacheDisabled: true})
	if got.ContactMapCacheEnabled {
		t.Fatalf("contact map cache enabled = true, want false")
	}
}

func sameInt64s(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
