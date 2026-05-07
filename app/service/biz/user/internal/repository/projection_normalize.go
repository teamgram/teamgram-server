package repository

import userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"

type normalizedProjectionRequest struct {
	ViewerUserIds  []int64
	TargetUserIds  []int64
	HydrateUserIds []int64
	WithFacts      bool
}

func normalizeProjectionConfig(cfg ProjectionConfig) ProjectionConfig {
	if cfg.SQLInChunkSize == 0 {
		cfg.SQLInChunkSize = 500
	}
	if cfg.SQLInChunkSize < 100 {
		cfg.SQLInChunkSize = 100
	}
	if cfg.SQLInChunkSize > 1000 {
		cfg.SQLInChunkSize = 1000
	}
	if cfg.MaxViewerUserIds <= 0 {
		cfg.MaxViewerUserIds = 8
	}
	if cfg.MaxTargetUserIds <= 0 {
		cfg.MaxTargetUserIds = 1000
	}
	if cfg.MaxProjectionPairs <= 0 {
		cfg.MaxProjectionPairs = 5000
	}
	if cfg.ContactMapMaxEntries <= 0 {
		cfg.ContactMapMaxEntries = 1000
	}
	return cfg
}

func normalizeProjectionRequest(viewerIds []int64, targetIds []int64, withFacts bool, cfg ProjectionConfig) (normalizedProjectionRequest, error) {
	viewers := uniquePositiveInt64s(viewerIds)
	if len(viewers) == 0 || len(viewers) > cfg.MaxViewerUserIds {
		return normalizedProjectionRequest{}, userpb.ErrUserInvalidArgument
	}

	targets := uniquePositiveInt64s(targetIds)
	if len(targets) > cfg.MaxTargetUserIds {
		return normalizedProjectionRequest{}, userpb.ErrUserInvalidArgument
	}
	if len(targets) > 0 && len(viewers)*len(targets) > cfg.MaxProjectionPairs {
		return normalizedProjectionRequest{}, userpb.ErrUserInvalidArgument
	}

	hydrate := uniquePositiveInt64s(append(append([]int64{}, viewers...), targets...))
	return normalizedProjectionRequest{
		ViewerUserIds:  viewers,
		TargetUserIds:  targets,
		HydrateUserIds: hydrate,
		WithFacts:      withFacts,
	}, nil
}

func uniquePositiveInt64s(in []int64) []int64 {
	if len(in) == 0 {
		return []int64{}
	}
	seen := make(map[int64]struct{}, len(in))
	out := make([]int64, 0, len(in))
	for _, id := range in {
		if id <= 0 {
			continue
		}
		if _, ok := seen[id]; ok {
			continue
		}
		seen[id] = struct{}{}
		out = append(out, id)
	}
	return out
}

func chunkInt64s(in []int64, size int) [][]int64 {
	if len(in) == 0 {
		return nil
	}
	if size <= 0 {
		size = len(in)
	}

	chunks := make([][]int64, 0, (len(in)+size-1)/size)
	for start := 0; start < len(in); start += size {
		end := start + size
		if end > len(in) {
			end = len(in)
		}
		chunks = append(chunks, in[start:end])
	}
	return chunks
}
