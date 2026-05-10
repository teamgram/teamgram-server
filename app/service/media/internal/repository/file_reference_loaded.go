package repository

import (
	"context"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
)

func (r *Repository) generateLoadedFileReference(ctx context.Context, domain string, mediaID, accessHash int64, objectID string) ([]byte, error) {
	if r == nil || r.fileReferenceService == nil {
		return nil, media.ErrFileReferenceInvalid
	}
	now := r.repositoryNow()
	ttl := r.fileReferenceTTL
	if ttl <= 0 {
		ttl = 24 * time.Hour
	}
	return r.fileReferenceService.Generate(ctx, FileReferenceClaims{
		MediaID:      mediaID,
		ObjectID:     objectID,
		OriginDomain: domain,
		OriginID:     0,
		ExpireAt:     now.Add(ttl).Unix(),
		AccessHash:   accessHash,
	}, r)
}
