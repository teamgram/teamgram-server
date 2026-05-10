package repository

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/media/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/filelease"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const (
	fileObjectBucket      = "documents"
	defaultReadLeaseTTL   = 300 * time.Second
	fileReferenceDocument = "document"
	fileReferencePhoto    = "photo"
	locatorThumbMimeType  = "image/jpeg"
	fileObjectKeyFormat   = "objects/%s.dat"
)

// ResolveFileLocation resolves a Telegram file location into a direct file object lease.
func (r *Repository) ResolveFileLocation(ctx context.Context, in *media.TLMediaResolveFileLocation) (*media.MediaResolvedFileObject, error) {
	if in == nil || in.Location == nil {
		return nil, media.ErrFileLocationInvalid
	}

	switch loc := in.Location.(type) {
	case *tg.TLInputDocumentFileLocation:
		return r.resolveDocumentLocation(ctx, loc)
	case *tg.TLInputPhotoFileLocation:
		return r.resolvePhotoLocation(ctx, loc)
	case *tg.TLInputPeerPhotoFileLocation:
		return r.resolvePeerPhotoLocation(ctx, loc)
	default:
		return nil, media.ErrFileLocationInvalid
	}
}

func (r *Repository) resolveDocumentLocation(ctx context.Context, loc *tg.TLInputDocumentFileLocation) (*media.MediaResolvedFileObject, error) {
	if loc == nil {
		return nil, media.ErrFileLocationInvalid
	}
	if _, err := r.validateFileReference(ctx, loc.FileReference, fileReferenceDocument, loc.Id, loc.AccessHash); err != nil {
		return nil, err
	}
	doc, err := r.findDocument(ctx, loc.Id)
	if err != nil {
		return nil, err
	}
	if doc.AccessHash != loc.AccessHash {
		return nil, media.ErrFileLocationInvalid
	}
	if loc.ThumbSize != "" {
		photoSize, photoErr := r.findPhotoSize(ctx, doc.ThumbId, loc.ThumbSize)
		videoSize, videoErr := r.findVideoSize(ctx, doc.VideoThumbId, loc.ThumbSize)
		photoFound := photoErr == nil
		videoFound := videoErr == nil
		if photoFound && videoFound {
			return nil, media.ErrFileLocationInvalid
		}
		if photoFound && doc.VideoThumbId != 0 && errors.Is(videoErr, media.ErrMediaStorage) {
			return nil, videoErr
		}
		if videoFound && doc.ThumbId != 0 && errors.Is(photoErr, media.ErrMediaStorage) {
			return nil, photoErr
		}
		if photoFound {
			return r.makeResolvedObject(photoSize.FilePath, int64(photoSize.FileSize), locatorThumbMimeType, doc.DcId, storageFileType(tg.ClazzID_storage_fileJpeg))
		}
		if videoFound {
			return r.makeResolvedObject(videoSize.FilePath, int64(videoSize.FileSize), doc.MimeType, doc.DcId, storageFileType(tg.ClazzID_storage_fileMp4))
		}
		if errors.Is(photoErr, media.ErrMediaStorage) {
			return nil, photoErr
		}
		if errors.Is(videoErr, media.ErrMediaStorage) {
			return nil, videoErr
		}
		return nil, media.ErrFileLocationInvalid
	}
	return r.makeResolvedObject(doc.FilePath, doc.FileSize, doc.MimeType, doc.DcId, storageTypeForFile(doc.UploadedFileName, doc.MimeType))
}

func (r *Repository) resolvePhotoLocation(ctx context.Context, loc *tg.TLInputPhotoFileLocation) (*media.MediaResolvedFileObject, error) {
	if loc == nil || loc.ThumbSize == "" {
		return nil, media.ErrFileLocationInvalid
	}
	if _, err := r.validateFileReference(ctx, loc.FileReference, fileReferencePhoto, loc.Id, loc.AccessHash); err != nil {
		return nil, err
	}
	photo, err := r.findPhoto(ctx, loc.Id)
	if err != nil {
		return nil, err
	}
	if photo.AccessHash != loc.AccessHash {
		return nil, media.ErrFileLocationInvalid
	}
	size, err := r.findPhotoSize(ctx, photo.SizeId, loc.ThumbSize)
	if err != nil {
		return nil, err
	}
	return r.makeResolvedObject(size.FilePath, int64(size.FileSize), locatorThumbMimeType, photo.DcId, storageFileType(tg.ClazzID_storage_fileJpeg))
}

func (r *Repository) resolvePeerPhotoLocation(ctx context.Context, loc *tg.TLInputPeerPhotoFileLocation) (*media.MediaResolvedFileObject, error) {
	if loc == nil || loc.PhotoId == 0 {
		return nil, media.ErrFileLocationInvalid
	}
	photo, err := r.findPhoto(ctx, loc.PhotoId)
	if err != nil {
		return nil, err
	}
	preferred := []string{"s", "m", "x", "y"}
	if loc.Big {
		preferred = []string{"x", "y", "m", "s"}
	}
	size, err := r.findFirstPhotoSize(ctx, photo.SizeId, preferred)
	if err != nil {
		return nil, err
	}
	return r.makeResolvedObject(size.FilePath, int64(size.FileSize), locatorThumbMimeType, photo.DcId, storageFileType(tg.ClazzID_storage_fileJpeg))
}

func (r *Repository) validateFileReference(ctx context.Context, token []byte, domain string, mediaID, accessHash int64) (FileReferenceClaims, error) {
	if r == nil || r.fileReferenceService == nil {
		return FileReferenceClaims{}, media.ErrFileReferenceInvalid
	}
	claims, err := r.fileReferenceService.Validate(ctx, token, r)
	if err != nil {
		return FileReferenceClaims{}, err
	}
	if claims.OriginDomain != domain || claims.MediaID != mediaID || claims.AccessHash != accessHash {
		return FileReferenceClaims{}, media.ErrFileReferenceInvalid
	}
	return claims, nil
}

func (r *Repository) findDocument(ctx context.Context, id int64) (*model.Documents, error) {
	if r == nil || r.model == nil || r.model.DocumentsModel == nil {
		return nil, media.ErrMediaStorage
	}
	doc, err := r.model.DocumentsModel.FindOneByDocumentId(ctx, id)
	if err != nil {
		if isNotFound(err) {
			return nil, media.ErrFileLocationInvalid
		}
		return nil, wrapStorage("find document for file location", err)
	}
	return doc, nil
}

func (r *Repository) findPhoto(ctx context.Context, id int64) (*model.Photos, error) {
	if r == nil || r.model == nil || r.model.PhotosModel == nil {
		return nil, media.ErrMediaStorage
	}
	photo, err := r.model.PhotosModel.FindOneByPhotoId(ctx, id)
	if err != nil {
		if isNotFound(err) {
			return nil, media.ErrFileLocationInvalid
		}
		return nil, wrapStorage("find photo for file location", err)
	}
	return photo, nil
}

func (r *Repository) findPhotoSize(ctx context.Context, sizeID int64, sizeType string) (model.PhotoSizes, error) {
	if sizeID == 0 || sizeType == "" {
		return model.PhotoSizes{}, media.ErrFileLocationInvalid
	}
	sizes, err := r.listPhotoSizes(ctx, sizeID)
	if err != nil {
		return model.PhotoSizes{}, err
	}
	for _, size := range sizes {
		if size.SizeType == sizeType {
			if size.FilePath == "" {
				return model.PhotoSizes{}, media.ErrFileLocationInvalid
			}
			return size, nil
		}
	}
	return model.PhotoSizes{}, media.ErrFileLocationInvalid
}

func (r *Repository) findVideoSize(ctx context.Context, sizeID int64, sizeType string) (model.VideoSizes, error) {
	if sizeID == 0 || sizeType == "" {
		return model.VideoSizes{}, media.ErrFileLocationInvalid
	}
	if r == nil || r.model == nil || r.model.VideoSizesModel == nil {
		return model.VideoSizes{}, media.ErrMediaStorage
	}
	sizes, err := r.model.VideoSizesModel.SelectListByVideoSizeId(ctx, sizeID)
	if err != nil {
		if isNotFound(err) {
			return model.VideoSizes{}, media.ErrFileLocationInvalid
		}
		return model.VideoSizes{}, wrapStorage("list video sizes for file location", err)
	}
	for _, size := range sizes {
		if size.SizeType == sizeType {
			if size.FilePath == "" {
				return model.VideoSizes{}, media.ErrFileLocationInvalid
			}
			return size, nil
		}
	}
	return model.VideoSizes{}, media.ErrFileLocationInvalid
}

func (r *Repository) findFirstPhotoSize(ctx context.Context, sizeID int64, preferred []string) (model.PhotoSizes, error) {
	if sizeID == 0 {
		return model.PhotoSizes{}, media.ErrFileLocationInvalid
	}
	sizes, err := r.listPhotoSizes(ctx, sizeID)
	if err != nil {
		return model.PhotoSizes{}, err
	}
	for _, want := range preferred {
		for _, size := range sizes {
			if size.SizeType == want && size.FilePath != "" {
				return size, nil
			}
		}
	}
	return model.PhotoSizes{}, media.ErrFileLocationInvalid
}

func (r *Repository) listPhotoSizes(ctx context.Context, sizeID int64) ([]model.PhotoSizes, error) {
	if r == nil || r.model == nil || r.model.PhotoSizesModel == nil {
		return nil, media.ErrMediaStorage
	}
	sizes, err := r.model.PhotoSizesModel.SelectListByPhotoSizeId(ctx, sizeID)
	if err != nil {
		if isNotFound(err) {
			return nil, media.ErrFileLocationInvalid
		}
		return nil, wrapStorage("list photo sizes for file location", err)
	}
	return sizes, nil
}

func (r *Repository) makeResolvedObject(objectID string, size int64, mimeType string, dcID int32, storageType int32) (*media.MediaResolvedFileObject, error) {
	if objectID == "" || size < 0 || mimeType == "" {
		return nil, media.ErrFileLocationInvalid
	}
	lease, err := r.signReadLease(objectID, size, mimeType, dcID, storageType)
	if err != nil {
		return nil, err
	}
	return media.MakeTLMediaResolvedFileObject(&media.TLMediaResolvedFileObject{
		ObjectId:        objectID,
		ReadLease:       lease,
		Size2:           size,
		MimeType:        mimeType,
		DcId:            dcID,
		StorageFileType: storageType,
	}).ToMediaResolvedFileObject(), nil
}

func (r *Repository) signReadLease(objectID string, size int64, mimeType string, dcID int32, storageType int32) ([]byte, error) {
	if r == nil {
		return nil, media.ErrMediaStorage
	}
	ttl := r.readLeaseTTL
	if ttl <= 0 {
		ttl = defaultReadLeaseTTL
	}
	token, err := filelease.Sign(string(r.readLeaseSecret), filelease.Claims{
		ObjectID:    objectID,
		Bucket:      fileObjectBucket,
		Key:         fmt.Sprintf(fileObjectKeyFormat, objectID),
		Size:        size,
		MimeType:    mimeType,
		StorageType: storageType,
		DCID:        dcID,
		ExpiresAt:   time.Now().Add(ttl).Unix(),
	})
	if err != nil {
		return nil, wrapStorage("sign read lease", err)
	}
	return token, nil
}

func storageTypeForFile(fileName, mimeType string) int32 {
	base := strings.ToLower(strings.TrimSpace(strings.Split(mimeType, ";")[0]))
	switch base {
	case "image/jpeg", "image/jpg":
		return storageFileType(tg.ClazzID_storage_fileJpeg)
	case "image/png":
		return storageFileType(tg.ClazzID_storage_filePng)
	case "image/gif":
		return storageFileType(tg.ClazzID_storage_fileGif)
	case "application/pdf":
		return storageFileType(tg.ClazzID_storage_filePdf)
	case "audio/mpeg", "audio/mp3":
		return storageFileType(tg.ClazzID_storage_fileMp3)
	case "video/quicktime", "video/mov":
		return storageFileType(tg.ClazzID_storage_fileMov)
	case "video/mp4":
		return storageFileType(tg.ClazzID_storage_fileMp4)
	case "image/webp":
		return storageFileType(tg.ClazzID_storage_fileWebp)
	}
	switch strings.ToLower(filepath.Ext(fileName)) {
	case ".jpg", ".jpeg":
		return storageFileType(tg.ClazzID_storage_fileJpeg)
	case ".png":
		return storageFileType(tg.ClazzID_storage_filePng)
	case ".gif":
		return storageFileType(tg.ClazzID_storage_fileGif)
	case ".pdf":
		return storageFileType(tg.ClazzID_storage_filePdf)
	case ".mp3":
		return storageFileType(tg.ClazzID_storage_fileMp3)
	case ".mov":
		return storageFileType(tg.ClazzID_storage_fileMov)
	case ".mp4":
		return storageFileType(tg.ClazzID_storage_fileMp4)
	case ".webp":
		return storageFileType(tg.ClazzID_storage_fileWebp)
	default:
		return storageFileType(tg.ClazzID_storage_fileUnknown)
	}
}

func storageFileType(clazzID uint32) int32 {
	return int32(clazzID)
}
