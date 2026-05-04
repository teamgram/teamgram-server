package identity

import (
	"context"

	"github.com/bytedance/gopkg/cloud/metainfo"
)

const CallerServiceKey = "teamgram_caller_service"

func WithCallerService(ctx context.Context, service string) context.Context {
	if service == "" {
		return ctx
	}
	return metainfo.WithPersistentValue(ctx, CallerServiceKey, service)
}

func CallerService(ctx context.Context) (string, bool) {
	service, ok := metainfo.GetPersistentValue(ctx, CallerServiceKey)
	return service, ok && service != ""
}
