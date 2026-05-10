package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"hash/fnv"
	"strings"
	"time"

	dfsapi "github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/app/service/media/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/app/service/mediaprocessor/mediaprocessor"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func (r *Repository) GetPhoto(ctx context.Context, id int64) (*tg.Photo, error) {
	if id == 0 {
		return nil, media.ErrPhotoNotFound
	}
	return r.loadPhoto(ctx, id)
}

func (r *Repository) mapPhotoResult(ctx context.Context, photo *tg.Photo, err error) (*tg.Photo, error) {
	if err != nil {
		if isServiceError(err) {
			return nil, err
		}
		return nil, wrapStorage("get photo", err)
	}
	return photo, nil
}

func (r *Repository) loadPhoto(ctx context.Context, id int64) (*tg.Photo, error) {
	do, err := r.model.PhotosModel.FindOneByPhotoId(ctx, id)
	if err != nil {
		if isNotFound(err) {
			return nil, media.ErrPhotoNotFound
		}
		return nil, wrapStorage("load photo", err)
	}
	var sizes []model.PhotoSizes
	if do.SizeId != 0 {
		sizes, err = r.model.PhotoSizesModel.SelectListByPhotoSizeId(ctx, do.SizeId)
		if err != nil {
			return nil, wrapStorage("load photo sizes", err)
		}
	}
	var videoSizes []model.VideoSizes
	if do.VideoSizeId != 0 {
		videoSizes, err = r.model.VideoSizesModel.SelectListByVideoSizeId(ctx, do.VideoSizeId)
		if err != nil {
			return nil, wrapStorage("load video sizes", err)
		}
	}
	fileReference, err := r.generateLoadedFileReference("photo", do.PhotoId, do.AccessHash, fmt.Sprintf("photo:%d", do.PhotoId))
	if err != nil {
		return nil, err
	}
	return mapPhotoAggregate(do, sizes, videoSizes, fileReference)
}

func (r *Repository) GetPhotoByRequest(ctx context.Context, in *media.TLMediaGetPhoto) (*tg.Photo, error) {
	return r.GetPhoto(ctx, in.PhotoId)
}

func (r *Repository) UploadPhotoFile(ctx context.Context, in *media.TLMediaUploadPhotoFile) (*tg.Photo, error) {
	if in == nil || in.File == nil {
		return nil, media.ErrMediaInvalidArgument
	}
	if r.dfsClient == nil {
		return nil, wrapMediaDownstream("dfs upload photo", media.ErrMediaDownstream)
	}
	if r.processorClient == nil {
		return nil, wrapMediaDownstream("mediaprocessor process photo", media.ErrMediaDownstream)
	}
	finalized, err := r.dfsClient.CommitUpload(ctx, &dfsapi.TLDfsCommitUpload{
		UploadSessionId: externalUploadSessionID(in.OwnerId, in.File),
		OwnerId:         in.OwnerId,
		File:            in.File,
		Purpose:         "media_original",
	})
	if err != nil {
		return nil, wrapDfsUploadError("dfs commit photo upload", err)
	}
	if finalized == nil || finalized.ObjectId == "" || len(finalized.ReadLease) == 0 {
		return nil, wrapMediaInvalidUploadedFile("dfs commit photo upload", errors.New("missing finalized object"))
	}
	processed, err := r.processorClient.ProcessPhoto(ctx, &mediaprocessor.TLMediaProcessorProcessPhoto{
		OwnerId:   in.OwnerId,
		ObjectId:  finalized.ObjectId,
		ReadLease: finalized.ReadLease,
		FileName:  inputFileName(in.File),
		Profile:   tg.ToBoolClazz(false),
	})
	if err != nil {
		return nil, wrapMediaDownstream("mediaprocessor process photo", err)
	}
	photo, sizeObjectIDs, err := r.photoFromProcessedUpload(in.OwnerId, finalized, processed)
	if err != nil {
		return nil, err
	}
	if err := r.savePhotoAggregateWithSizePaths(ctx, inputFileName(in.File), photo, sizeObjectIDs); err != nil {
		return nil, err
	}
	return photo, nil
}

func (r *Repository) UploadPhotoFileViaLegacyDFS(ctx context.Context, in *media.TLMediaUploadPhotoFile) (*tg.Photo, error) {
	if in == nil || in.File == nil {
		return nil, media.ErrMediaInvalidArgument
	}
	if r.dfsClient == nil {
		return nil, wrapMediaDownstream("dfs legacy upload photo", media.ErrMediaDownstream)
	}
	photo, err := r.dfsClient.UploadPhotoFileV2ViaLegacyDFS(ctx, &dfsapi.TLDfsUploadPhotoFileV2{
		Creator: in.OwnerId,
		File:    in.File,
	})
	if err != nil {
		return nil, wrapDfsUploadError("dfs legacy upload photo", err)
	}
	if photo == nil {
		return nil, wrapMediaInvalidUploadedFile("dfs legacy upload photo", errors.New("missing legacy photo"))
	}
	if err := r.savePhotoAggregate(ctx, inputFileName(in.File), photo); err != nil {
		return nil, err
	}
	return photo, nil
}

func externalUploadSessionID(ownerID int64, file tg.InputFileClazz) string {
	switch f := file.(type) {
	case *tg.TLInputFile:
		return fmt.Sprintf("ext:%d:%d:%d", ownerID, f.Id, f.Parts)
	case *tg.TLInputFileBig:
		return fmt.Sprintf("ext:%d:%d:%d", ownerID, f.Id, f.Parts)
	default:
		return ""
	}
}

func inputFileName(file tg.InputFileClazz) string {
	switch f := file.(type) {
	case *tg.TLInputFile:
		return f.Name
	case *tg.TLInputFileBig:
		return f.Name
	default:
		return ""
	}
}

func (r *Repository) photoFromProcessedUpload(ownerID int64, finalized *dfsapi.FileFinalizedObject, processed *mediaprocessor.ProcessedPhoto) (*tg.Photo, map[string]string, error) {
	if r == nil || r.fileReferenceService == nil {
		return nil, nil, media.ErrFileReferenceInvalid
	}
	if finalized == nil || processed == nil {
		return nil, nil, media.ErrMediaInvalidArgument
	}
	photoID := stablePositiveID("photo:" + finalized.ObjectId)
	accessHash := stablePositiveID("photo-access:" + finalized.ObjectId)
	now := time.Now()
	if r.fileReferenceService != nil {
		now = r.fileReferenceService.now()
	}
	ttl := r.fileReferenceTTL
	if ttl <= 0 {
		ttl = 24 * time.Hour
	}
	fileReference, err := r.fileReferenceService.Generate(FileReferenceClaims{
		MediaID:      photoID,
		ObjectID:     firstNonEmpty(processed.OriginalObjectId, finalized.ObjectId),
		OriginDomain: "photo",
		OriginID:     ownerID,
		ExpireAt:     now.Add(ttl).Unix(),
		AccessHash:   accessHash,
	})
	if err != nil {
		return nil, nil, err
	}
	sizes, sizeObjectIDs, err := mapProcessedPhotoSizes(processed.Sizes)
	if err != nil {
		return nil, nil, wrapMediaInvalidUploadedFile("mediaprocessor process photo", err)
	}
	if len(sizes) == 0 {
		return nil, nil, wrapMediaInvalidUploadedFile("mediaprocessor process photo", errors.New("missing photo sizes"))
	}
	return tg.MakeTLPhoto(&tg.TLPhoto{
		HasStickers:   false,
		Id:            photoID,
		AccessHash:    accessHash,
		FileReference: fileReference,
		Date:          int32(now.Unix()),
		Sizes:         sizes,
		VideoSizes:    nil,
		DcId:          finalized.DcId,
	}).ToPhoto(), sizeObjectIDs, nil
}

func mapProcessedPhotoSizes(derivatives []mediaprocessor.ProcessorDerivativeClazz) ([]tg.PhotoSizeClazz, map[string]string, error) {
	const maxInt32 = int64(^uint32(0) >> 1)

	sizes := make([]tg.PhotoSizeClazz, 0, len(derivatives))
	objectIDs := make(map[string]string, len(derivatives))
	seenTypes := make(map[string]struct{}, len(derivatives))
	for i, derivative := range derivatives {
		if derivative == nil {
			return nil, nil, fmt.Errorf("photo derivative %d is nil", i)
		}
		switch derivative.Kind {
		case processorDerivativePhotoStripped:
			const strippedSizeType = "i"
			if derivative.Width <= 0 || derivative.Height <= 0 {
				return nil, nil, fmt.Errorf("photo derivative %d has invalid dimensions", i)
			}
			if len(derivative.Bytes) == 0 {
				return nil, nil, fmt.Errorf("photo derivative %d has empty stripped bytes", i)
			}
			if _, exists := seenTypes[strippedSizeType]; exists {
				return nil, nil, fmt.Errorf("photo derivative %d duplicates size type %q", i, strippedSizeType)
			}
			sizes = append(sizes, tg.MakeTLPhotoStrippedSize(&tg.TLPhotoStrippedSize{
				Type:  strippedSizeType,
				Bytes: append([]byte(nil), derivative.Bytes...),
			}))
			seenTypes[strippedSizeType] = struct{}{}
		case processorDerivativePhotoSize:
			sizeType, ok := photoSizeTypeFromDerivativeFileName(derivative.FileName)
			if !ok {
				return nil, nil, fmt.Errorf("photo derivative %d missing size type", i)
			}
			if derivative.ObjectId == "" {
				return nil, nil, fmt.Errorf("photo derivative %d missing object id", i)
			}
			if derivative.Width <= 0 || derivative.Height <= 0 {
				return nil, nil, fmt.Errorf("photo derivative %d has invalid dimensions", i)
			}
			if derivative.Size2 <= 0 || derivative.Size2 > maxInt32 {
				return nil, nil, fmt.Errorf("photo derivative %d has invalid size", i)
			}
			if _, exists := seenTypes[sizeType]; exists {
				return nil, nil, fmt.Errorf("photo derivative %d duplicates size type %q", i, sizeType)
			}
			if len(derivative.ProgressiveSizes) > 0 {
				if err := validateProgressiveSizes(derivative.ProgressiveSizes, derivative.Size2); err != nil {
					return nil, nil, fmt.Errorf("photo derivative %d has invalid progressive sizes: %w", i, err)
				}
				sizes = append(sizes, tg.MakeTLPhotoSizeProgressive(&tg.TLPhotoSizeProgressive{
					Type:  sizeType,
					W:     derivative.Width,
					H:     derivative.Height,
					Sizes: append([]int32(nil), derivative.ProgressiveSizes...),
				}))
			} else {
				sizes = append(sizes, tg.MakeTLPhotoSize(&tg.TLPhotoSize{
					Type:  sizeType,
					W:     derivative.Width,
					H:     derivative.Height,
					Size2: int32(derivative.Size2),
				}))
			}
			seenTypes[sizeType] = struct{}{}
			objectIDs[sizeType] = derivative.ObjectId
		default:
			return nil, nil, fmt.Errorf("photo derivative %d has unknown kind %q", i, derivative.Kind)
		}
	}
	return sizes, objectIDs, nil
}

const (
	processorDerivativePhotoSize     = "photo_size"
	processorDerivativePhotoStripped = "photo_stripped"
	photoSizeCachedTypeProgressive   = 4
)

func validateProgressiveSizes(sizes []int32, total int64) error {
	var previous int32
	for i, size := range sizes {
		if size <= 0 {
			return fmt.Errorf("scan %d is not positive", i)
		}
		if i > 0 && size <= previous {
			return fmt.Errorf("scan %d is not increasing", i)
		}
		if int64(size) > total {
			return fmt.Errorf("scan %d exceeds total size", i)
		}
		previous = size
	}
	if int64(sizes[len(sizes)-1]) != total {
		return fmt.Errorf("last scan size %d does not equal total size %d", sizes[len(sizes)-1], total)
	}
	return nil
}

func photoSizeTypeFromDerivativeFileName(fileName string) (string, bool) {
	sizeType, _, ok := strings.Cut(fileName, "_")
	if !ok || sizeType == "" {
		return "", false
	}
	for _, r := range sizeType {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r < '0' || r > '9') {
			return "", false
		}
	}
	return sizeType, true
}

func stablePositiveID(seed string) int64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(seed))
	id := int64(h.Sum64() & uint64(^uint64(0)>>1))
	if id == 0 {
		return 1
	}
	return id
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}

func (r *Repository) UploadProfilePhotoFile(ctx context.Context, in *media.TLMediaUploadProfilePhotoFile) (*tg.Photo, error) {
	if in == nil || (in.File == nil && in.Video == nil) {
		return nil, media.ErrMediaInvalidArgument
	}
	if r.dfsClient == nil {
		return nil, wrapMediaDownstream("dfs upload profile photo", media.ErrMediaDownstream)
	}
	if r.processorClient == nil {
		return nil, wrapMediaDownstream("mediaprocessor process profile photo", media.ErrMediaDownstream)
	}

	var (
		photo         *tg.Photo
		sizeObjectIDs map[string]string
		fileName      string
	)
	if in.File != nil {
		finalized, err := r.dfsClient.CommitUpload(ctx, &dfsapi.TLDfsCommitUpload{
			UploadSessionId: externalUploadSessionID(in.OwnerId, in.File),
			OwnerId:         in.OwnerId,
			File:            in.File,
			Purpose:         "media_profile_photo",
		})
		if err != nil {
			return nil, wrapDfsUploadError("dfs commit profile photo upload", err)
		}
		if finalized == nil || finalized.ObjectId == "" || len(finalized.ReadLease) == 0 {
			return nil, wrapMediaInvalidUploadedFile("dfs commit profile photo upload", errors.New("missing finalized object"))
		}
		fileName = inputFileName(in.File)
		processed, err := r.processorClient.ProcessPhoto(ctx, &mediaprocessor.TLMediaProcessorProcessPhoto{
			OwnerId:   in.OwnerId,
			ObjectId:  finalized.ObjectId,
			ReadLease: finalized.ReadLease,
			FileName:  fileName,
			Profile:   tg.ToBoolClazz(true),
		})
		if err != nil {
			return nil, wrapMediaDownstream("mediaprocessor process profile photo", err)
		}
		photo, sizeObjectIDs, err = r.photoFromProcessedUpload(in.OwnerId, finalized, processed)
		if err != nil {
			return nil, err
		}
	}

	if in.Video != nil {
		finalized, err := r.dfsClient.CommitUpload(ctx, &dfsapi.TLDfsCommitUpload{
			UploadSessionId: externalUploadSessionID(in.OwnerId, in.Video),
			OwnerId:         in.OwnerId,
			File:            in.Video,
			Purpose:         "media_profile_video",
		})
		if err != nil {
			return nil, wrapDfsUploadError("dfs commit profile video upload", err)
		}
		if finalized == nil || finalized.ObjectId == "" || len(finalized.ReadLease) == 0 {
			return nil, wrapMediaInvalidUploadedFile("dfs commit profile video upload", errors.New("missing finalized video object"))
		}
		videoName := inputFileName(in.Video)
		processed, err := r.processorClient.ProcessMp4(ctx, &mediaprocessor.TLMediaProcessorProcessMp4{
			OwnerId:   in.OwnerId,
			ObjectId:  finalized.ObjectId,
			ReadLease: finalized.ReadLease,
			FileName:  videoName,
		})
		if err != nil {
			return nil, wrapMediaDownstream("mediaprocessor process profile video", err)
		}
		videoSizes, err := profileVideoSizesFromProcessed(processed, in.VideoStartTs)
		if err != nil {
			return nil, wrapMediaInvalidUploadedFile("mediaprocessor process profile video", err)
		}
		if photo == nil {
			photo, sizeObjectIDs, err = r.photoFromProfileVideoUpload(in.OwnerId, finalized, processed)
			if err != nil {
				return nil, err
			}
			fileName = videoName
		}
		do, ok := photo.ToPhoto()
		if !ok {
			return nil, media.ErrMediaInvalidArgument
		}
		do.VideoSizes = videoSizes
	}
	if photo == nil {
		return nil, media.ErrMediaInvalidArgument
	}
	if err := r.savePhotoAggregateWithSizePaths(ctx, fileName, photo, sizeObjectIDs); err != nil {
		return nil, err
	}
	return photo, nil
}

func (r *Repository) photoFromProfileVideoUpload(ownerID int64, finalized *dfsapi.FileFinalizedObject, processed *mediaprocessor.ProcessedDocument) (*tg.Photo, map[string]string, error) {
	if r == nil || r.fileReferenceService == nil {
		return nil, nil, media.ErrFileReferenceInvalid
	}
	if finalized == nil || processed == nil || processed.ObjectId == "" {
		return nil, nil, wrapMediaInvalidUploadedFile("mediaprocessor process profile video", errors.New("missing processed profile video"))
	}
	sizes, sizeObjectIDs, err := mapProcessedDocumentThumbs(processed.Thumbs)
	if err != nil {
		return nil, nil, wrapMediaInvalidUploadedFile("mediaprocessor process profile video", err)
	}
	if len(sizes) == 0 {
		return nil, nil, wrapMediaInvalidUploadedFile("mediaprocessor process profile video", errors.New("missing profile video cover"))
	}
	photoID := stablePositiveID("profile-photo:" + processed.ObjectId)
	accessHash := stablePositiveID("profile-photo-access:" + processed.ObjectId)
	now := r.repositoryNow()
	ttl := r.fileReferenceTTL
	if ttl <= 0 {
		ttl = 24 * time.Hour
	}
	fileReference, err := r.fileReferenceService.Generate(FileReferenceClaims{
		MediaID:      photoID,
		ObjectID:     processed.ObjectId,
		OriginDomain: "photo",
		OriginID:     ownerID,
		ExpireAt:     now.Add(ttl).Unix(),
		AccessHash:   accessHash,
	})
	if err != nil {
		return nil, nil, err
	}
	return tg.MakeTLPhoto(&tg.TLPhoto{
		HasStickers:   false,
		Id:            photoID,
		AccessHash:    accessHash,
		FileReference: fileReference,
		Date:          int32(now.Unix()),
		Sizes:         sizes,
		DcId:          finalized.DcId,
	}).ToPhoto(), sizeObjectIDs, nil
}

func profileVideoSizesFromProcessed(processed *mediaprocessor.ProcessedDocument, startTs *float64) ([]tg.VideoSizeClazz, error) {
	const maxInt32 = int64(^uint32(0) >> 1)
	if processed == nil || processed.Size2 <= 0 || processed.Size2 > maxInt32 {
		return nil, errors.New("invalid profile video size")
	}
	attrs, err := decodeDocumentAttributeVector(processed.Attributes)
	if err != nil {
		return nil, err
	}
	for _, attr := range attrs {
		if video, ok := attr.(*tg.TLDocumentAttributeVideo); ok {
			return []tg.VideoSizeClazz{tg.MakeTLVideoSize(&tg.TLVideoSize{
				Type:         "v",
				W:            video.W,
				H:            video.H,
				Size2:        int32(processed.Size2),
				VideoStartTs: startTs,
			})}, nil
		}
	}
	return nil, errors.New("missing profile video attributes")
}

func (r *Repository) UploadedProfilePhoto(ctx context.Context, in *media.TLMediaUploadedProfilePhoto) (*tg.Photo, error) {
	if in == nil || in.PhotoId == 0 {
		return nil, media.ErrMediaInvalidArgument
	}
	if r.dfsClient == nil {
		return nil, wrapMediaDownstream("dfs uploaded profile photo", media.ErrMediaDownstream)
	}
	photo, err := r.dfsClient.UploadedProfilePhoto(ctx, &dfsapi.TLDfsUploadedProfilePhoto{
		Creator: in.OwnerId,
		PhotoId: in.PhotoId,
	})
	if err != nil {
		return nil, wrapDfsUploadError("dfs uploaded profile photo", err)
	}
	if photo == nil {
		return nil, wrapMediaInvalidUploadedFile("dfs uploaded profile photo", errors.New("missing uploaded profile photo"))
	}
	if err := r.savePhotoAggregate(ctx, "", photo); err != nil {
		return nil, err
	}
	return photo, nil
}

func (r *Repository) GetPhotoSizeList(ctx context.Context, sizeID int64) (*media.PhotoSizeList, error) {
	if sizeID == 0 {
		return nil, media.ErrMediaInvalidArgument
	}
	sizes, err := r.model.PhotoSizesModel.SelectListByPhotoSizeId(ctx, sizeID)
	if err != nil {
		return nil, wrapStorage("load photo size list", err)
	}
	photoSizes, err := mapPhotoSizes(sizes)
	if err != nil {
		return nil, err
	}
	return media.MakeTLPhotoSizeList(&media.TLPhotoSizeList{SizeId: sizeID, Sizes: photoSizes, DcId: 1}).ToPhotoSizeList(), nil
}

func (r *Repository) GetPhotoSizeListList(ctx context.Context, ids []int64) (*media.VectorPhotoSizeList, error) {
	out := &media.VectorPhotoSizeList{Datas: make([]media.PhotoSizeListClazz, 0, len(ids))}
	for _, id := range ids {
		list, err := r.GetPhotoSizeList(ctx, id)
		if err != nil {
			return nil, err
		}
		out.Datas = append(out.Datas, list)
	}
	return out, nil
}

func (r *Repository) GetVideoSizeList(ctx context.Context, sizeID int64) (*media.VideoSizeList, error) {
	if sizeID == 0 {
		return nil, media.ErrMediaInvalidArgument
	}
	sizes, err := r.model.VideoSizesModel.SelectListByVideoSizeId(ctx, sizeID)
	if err != nil {
		return nil, wrapStorage("load video size list", err)
	}
	return media.MakeTLVideoSizeList(&media.TLVideoSizeList{SizeId: sizeID, Sizes: mapVideoSizes(sizes), DcId: 1}).ToVideoSizeList(), nil
}

func (r *Repository) savePhotoAggregate(ctx context.Context, inputFileName string, photo *tg.Photo) error {
	return r.savePhotoAggregateWithSizePaths(ctx, inputFileName, photo, nil)
}

func (r *Repository) savePhotoAggregateWithSizePaths(ctx context.Context, inputFileName string, photo *tg.Photo, sizeObjectIDs map[string]string) error {
	if r == nil || r.model == nil || photo == nil {
		return nil
	}
	do, ok := photo.ToPhoto()
	if !ok {
		return media.ErrMediaInvalidArgument
	}
	photoRow := &model.Photos{
		PhotoId:       do.Id,
		AccessHash:    do.AccessHash,
		HasStickers:   do.HasStickers,
		DcId:          do.DcId,
		Date2:         int64(do.Date),
		InputFileName: inputFileName,
	}
	if len(do.Sizes) > 0 {
		photoRow.SizeId = do.Id
	}
	if len(do.VideoSizes) > 0 {
		photoRow.HasVideo = true
		photoRow.VideoSizeId = do.Id
	}
	if _, err := r.model.PhotosModel.Insert2(ctx, photoRow); err != nil {
		return wrapStorage("save photo", err)
	}
	for _, size := range do.Sizes {
		if err := r.savePhotoSizeWithPath(ctx, do.Id, size, sizeObjectIDs); err != nil {
			return err
		}
	}
	for _, size := range do.VideoSizes {
		if err := r.saveVideoSize(ctx, do.Id, size); err != nil {
			return err
		}
	}
	return nil
}

func (r *Repository) savePhotoSize(ctx context.Context, id int64, size tg.PhotoSizeClazz) error {
	return r.savePhotoSizeWithPath(ctx, id, size, nil)
}

func (r *Repository) savePhotoSizeWithPath(ctx context.Context, id int64, size tg.PhotoSizeClazz, sizeObjectIDs map[string]string) error {
	switch s := size.(type) {
	case *tg.TLPhotoSize:
		_, _, err := r.model.PhotoSizesModel.Insert(ctx, &model.PhotoSizes{PhotoSizeId: id, SizeType: s.Type, Width: s.W, Height: s.H, FileSize: s.Size2, FilePath: sizeObjectIDs[s.Type]})
		if err != nil {
			return wrapStorage("save photo size", err)
		}
	case *tg.TLPhotoStrippedSize:
		_, _, err := r.model.PhotoSizesModel.Insert(ctx, &model.PhotoSizes{PhotoSizeId: id, SizeType: s.Type, HasStripped: true, StrippedBytes: string(s.Bytes)})
		if err != nil {
			return wrapStorage("save stripped photo size", err)
		}
	case *tg.TLPhotoSizeProgressive:
		cachedBytes, err := json.Marshal(s.Sizes)
		if err != nil {
			return wrapMediaInvalidUploadedFile("save progressive photo size", err)
		}
		fileSize := int32(0)
		if len(s.Sizes) > 0 {
			fileSize = s.Sizes[len(s.Sizes)-1]
		}
		_, _, err = r.model.PhotoSizesModel.Insert(ctx, &model.PhotoSizes{PhotoSizeId: id, SizeType: s.Type, Width: s.W, Height: s.H, FileSize: fileSize, FilePath: sizeObjectIDs[s.Type], CachedType: photoSizeCachedTypeProgressive, CachedBytes: string(cachedBytes)})
		if err != nil {
			return wrapStorage("save progressive photo size", err)
		}
	}
	return nil
}

func (r *Repository) saveVideoSize(ctx context.Context, id int64, size tg.VideoSizeClazz) error {
	if s, ok := size.(*tg.TLVideoSize); ok {
		var startTs float64
		if s.VideoStartTs != nil {
			startTs = *s.VideoStartTs
		}
		_, _, err := r.model.VideoSizesModel.Insert(ctx, &model.VideoSizes{VideoSizeId: id, SizeType: s.Type, Width: s.W, Height: s.H, FileSize: s.Size2, VideoStartTs: startTs})
		if err != nil {
			return wrapStorage("save video size", err)
		}
	}
	return nil
}
