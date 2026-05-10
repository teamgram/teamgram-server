package repository

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/service/media/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
)

func (r *Repository) SaveFileReference(ctx context.Context, token []byte, claims FileReferenceClaims) error {
	if r == nil || r.model == nil || r.model.FileReferencesModel == nil {
		return media.ErrMediaStorage
	}
	// Domain is the lookup partition. OriginDomain is retained separately as
	// the audit/refresh origin. They are equal for Phase 1 generated handles.
	_, _, err := r.model.FileReferencesModel.Insert(ctx, &model.FileReferences{
		RefHash:      append([]byte(nil), token...),
		Domain:       claims.OriginDomain,
		MediaId:      claims.MediaID,
		AccessHash:   claims.AccessHash,
		ObjectId:     claims.ObjectID,
		OriginDomain: claims.OriginDomain,
		OriginId:     claims.OriginID,
		ExpireAt:     claims.ExpireAt,
		RevokedAt:    0,
	})
	if err != nil {
		return wrapStorage("save file reference", err)
	}
	return nil
}

func (r *Repository) LoadFileReference(ctx context.Context, token []byte) (FileReferenceClaims, error) {
	if r == nil || r.model == nil || r.model.FileReferencesModel == nil {
		return FileReferenceClaims{}, media.ErrMediaStorage
	}
	row, err := r.model.FileReferencesModel.SelectByRefHash(ctx, token)
	if err != nil {
		if isNotFound(err) {
			return FileReferenceClaims{}, media.ErrFileReferenceInvalid
		}
		return FileReferenceClaims{}, wrapStorage("load file reference", err)
	}
	if row.RevokedAt > 0 {
		return FileReferenceClaims{}, media.ErrFileReferenceInvalid
	}
	return FileReferenceClaims{
		MediaID:      row.MediaId,
		ObjectID:     row.ObjectId,
		OriginDomain: row.OriginDomain,
		OriginID:     row.OriginId,
		ExpireAt:     row.ExpireAt,
		AccessHash:   row.AccessHash,
	}, nil
}
