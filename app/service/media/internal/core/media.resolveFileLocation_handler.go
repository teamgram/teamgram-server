package core

import (
	"errors"

	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MediaResolveFileLocation
// media.resolveFileLocation location:InputFileLocation viewer_id:long = MediaResolvedFileObject;
func (c *MediaCore) MediaResolveFileLocation(in *media.TLMediaResolveFileLocation) (*media.MediaResolvedFileObject, error) {
	out, err := c.svcCtx.Repo.ResolveFileLocation(c.ctx, in)
	if err == nil {
		return out, nil
	}
	switch {
	case errors.Is(err, media.ErrFileReferenceEmpty):
		return nil, tg.ErrFileReferenceEmpty
	case errors.Is(err, media.ErrFileReferenceExpired):
		return nil, tg.ErrFileReferenceExpired
	case errors.Is(err, media.ErrFileReferenceInvalid):
		return nil, tg.ErrFileReferenceInvalid
	case errors.Is(err, media.ErrFileLocationInvalid),
		errors.Is(err, media.ErrDocumentNotFound),
		errors.Is(err, media.ErrPhotoNotFound),
		errors.Is(err, media.ErrMediaInvalidArgument):
		return nil, tg.ErrLocationInvalid
	default:
		return nil, err
	}
}
