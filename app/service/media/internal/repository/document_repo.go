package repository

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	dfsapi "github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/app/service/media/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/app/service/mediaprocessor/mediaprocessor"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const documentThumbSizeType = "m"

func (r *Repository) GetDocument(ctx context.Context, id int64) (*tg.Document, error) {
	if id == 0 {
		return nil, media.ErrDocumentNotFound
	}
	return r.loadDocument(ctx, id)
}

func (r *Repository) mapDocumentResult(ctx context.Context, doc *tg.Document, err error) (*tg.Document, error) {
	if err != nil {
		if isServiceError(err) {
			return nil, err
		}
		return nil, wrapStorage("get document", err)
	}
	return doc, nil
}

func (r *Repository) loadDocument(ctx context.Context, id int64) (*tg.Document, error) {
	do, err := r.model.DocumentsModel.FindOneByDocumentId(ctx, id)
	if err != nil {
		if isNotFound(err) {
			return nil, media.ErrDocumentNotFound
		}
		return nil, wrapStorage("load document", err)
	}
	return r.mapDocumentWithThumbs(ctx, do)
}

func (r *Repository) GetDocumentByRequest(ctx context.Context, in *media.TLMediaGetDocument) (*tg.Document, error) {
	return r.GetDocument(ctx, in.Id)
}

func (r *Repository) UploadedDocumentMedia(ctx context.Context, in *media.TLMediaUploadedDocumentMedia) (*tg.MessageMedia, error) {
	if in == nil || in.Media == nil {
		return nil, media.ErrMediaInvalidArgument
	}
	uploaded, ok := in.Media.(*tg.TLInputMediaUploadedDocument)
	if !ok || uploaded.File == nil {
		return nil, media.ErrMediaInvalidArgument
	}
	if r.dfsClient == nil {
		return nil, wrapMediaDownstream("dfs upload document", media.ErrMediaDownstream)
	}
	classification, err := classifyUploadedDocument(uploaded)
	if err != nil {
		return nil, err
	}

	finalized, err := r.dfsClient.CommitUpload(ctx, &dfsapi.TLDfsCommitUpload{
		UploadSessionId: externalUploadSessionID(in.OwnerId, uploaded.File),
		OwnerId:         in.OwnerId,
		File:            uploaded.File,
		Purpose:         "media_original",
	})
	if err != nil {
		return nil, wrapDfsUploadError("dfs commit document upload", err)
	}
	if finalized == nil || finalized.ObjectId == "" {
		return nil, wrapMediaInvalidUploadedFile("dfs commit document upload", errors.New("missing finalized object"))
	}

	var doc *tg.Document
	documentObjectID := finalized.ObjectId
	thumbObjectIDs := map[string]string(nil)
	switch classification.RequiredTransform {
	case requiredDocumentTransformGifv:
		doc, documentObjectID, thumbObjectIDs, err = r.processUploadedGifDocument(ctx, in.OwnerId, finalized, uploaded)
	case requiredDocumentTransformMp4:
		doc, documentObjectID, thumbObjectIDs, err = r.processUploadedMp4Document(ctx, in.OwnerId, finalized, uploaded)
	default:
		doc, thumbObjectIDs, err = r.documentFromOriginalUpload(ctx, in.OwnerId, finalized, uploaded)
		documentObjectID = finalized.ObjectId
	}
	if err != nil {
		return nil, err
	}
	if err := r.saveDocumentAggregateWithPaths(ctx, uploadedFileName(uploaded), doc, documentObjectID, thumbObjectIDs); err != nil {
		return nil, err
	}
	docClazz, ok := doc.ToDocument()
	if !ok {
		return nil, media.ErrMediaInvalidArgument
	}
	return uploadedDocumentMessageMedia(uploaded, docClazz, classification, nil), nil
}

func uploadedDocumentMessageMedia(uploaded *tg.TLInputMediaUploadedDocument, doc tg.DocumentClazz, classification documentClassification, videoCover tg.PhotoClazz) *tg.MessageMedia {
	video, round, voice := documentMessageFlagsFromRenderKind(classification)
	return tg.MakeTLMessageMediaDocument(&tg.TLMessageMediaDocument{
		Spoiler:        uploaded.Spoiler,
		Video:          video,
		Round:          round,
		Voice:          voice,
		Document:       doc,
		VideoCover:     videoCover,
		VideoTimestamp: uploaded.VideoTimestamp,
		TtlSeconds:     uploaded.TtlSeconds,
	}).ToMessageMedia()
}

func (r *Repository) UploadedDocumentMediaViaLegacyDFS(ctx context.Context, in *media.TLMediaUploadedDocumentMedia) (*tg.MessageMedia, error) {
	if in == nil || in.Media == nil {
		return nil, media.ErrMediaInvalidArgument
	}
	uploaded, ok := in.Media.(*tg.TLInputMediaUploadedDocument)
	if !ok || uploaded.File == nil {
		return nil, media.ErrMediaInvalidArgument
	}
	if r.dfsClient == nil {
		return nil, wrapMediaDownstream("dfs legacy upload document", media.ErrMediaDownstream)
	}

	var (
		doc *tg.Document
		err error
	)
	switch {
	case isAnimatedGif(uploaded):
		doc, err = r.dfsClient.UploadGifDocumentMediaViaLegacyDFS(ctx, &dfsapi.TLDfsUploadGifDocumentMedia{Creator: in.OwnerId, Media: in.Media})
	case uploaded.MimeType == "video/mp4":
		doc, err = r.dfsClient.UploadMp4DocumentMediaViaLegacyDFS(ctx, &dfsapi.TLDfsUploadMp4DocumentMedia{Creator: in.OwnerId, Media: in.Media})
	default:
		doc, err = r.dfsClient.UploadDocumentFileV2ViaLegacyDFS(ctx, &dfsapi.TLDfsUploadDocumentFileV2{Creator: in.OwnerId, Media: in.Media})
	}
	if err != nil {
		return nil, wrapDfsUploadError("dfs legacy upload document", err)
	}
	if doc == nil {
		return nil, wrapMediaInvalidUploadedFile("dfs legacy upload document", errors.New("missing legacy document"))
	}
	if err := r.refreshDocumentFileReference(ctx, in.OwnerId, doc); err != nil {
		return nil, err
	}
	if err := r.saveDocumentAggregate(ctx, uploadedFileName(uploaded), doc); err != nil {
		return nil, err
	}
	docClazz, ok := doc.ToDocument()
	if !ok {
		return nil, media.ErrMediaInvalidArgument
	}
	return tg.MakeTLMessageMediaDocument(&tg.TLMessageMediaDocument{
		Spoiler:    uploaded.Spoiler,
		Document:   docClazz,
		TtlSeconds: uploaded.TtlSeconds,
	}).ToMessageMedia(), nil
}

func (r *Repository) processUploadedGifDocument(ctx context.Context, ownerID int64, finalized *dfsapi.FileFinalizedObject, uploaded *tg.TLInputMediaUploadedDocument) (*tg.Document, string, map[string]string, error) {
	if r.processorClient == nil {
		return nil, "", nil, wrapMediaDownstream("mediaprocessor process gif", media.ErrMediaDownstream)
	}
	if len(finalized.ReadLease) == 0 {
		return nil, "", nil, wrapMediaInvalidUploadedFile("dfs commit document upload", errors.New("missing finalized read lease"))
	}
	req := &mediaprocessor.TLMediaProcessorProcessGif{
		OwnerId:   ownerID,
		ObjectId:  finalized.ObjectId,
		ReadLease: finalized.ReadLease,
		FileName:  uploadedFileName(uploaded),
	}
	processed, err := r.processorClient.ProcessGif(ctx, req)
	if err != nil {
		return nil, "", nil, wrapMediaDownstream("mediaprocessor process gif", err)
	}
	doc, thumbObjectIDs, err := r.documentFromProcessedUpload(ctx, ownerID, finalized, processed)
	if err != nil {
		return nil, "", nil, err
	}
	return doc, processed.ObjectId, thumbObjectIDs, nil
}

func (r *Repository) processUploadedMp4Document(ctx context.Context, ownerID int64, finalized *dfsapi.FileFinalizedObject, uploaded *tg.TLInputMediaUploadedDocument) (*tg.Document, string, map[string]string, error) {
	if r.processorClient == nil {
		return nil, "", nil, wrapMediaDownstream("mediaprocessor process mp4", media.ErrMediaDownstream)
	}
	if len(finalized.ReadLease) == 0 {
		return nil, "", nil, wrapMediaInvalidUploadedFile("dfs commit document upload", errors.New("missing finalized read lease"))
	}
	attrs, err := encodeDocumentAttributeVector(uploaded.Attributes)
	if err != nil {
		return nil, "", nil, err
	}
	processed, err := r.processorClient.ProcessMp4(ctx, &mediaprocessor.TLMediaProcessorProcessMp4{
		OwnerId:    ownerID,
		ObjectId:   finalized.ObjectId,
		ReadLease:  finalized.ReadLease,
		FileName:   uploadedFileName(uploaded),
		Attributes: attrs,
	})
	if err != nil {
		return nil, "", nil, wrapMediaDownstream("mediaprocessor process mp4", err)
	}
	doc, thumbObjectIDs, err := r.documentFromProcessedUpload(ctx, ownerID, finalized, processed)
	if err != nil {
		return nil, "", nil, err
	}
	return doc, processed.ObjectId, thumbObjectIDs, nil
}

func (r *Repository) documentFromOriginalUpload(ctx context.Context, ownerID int64, finalized *dfsapi.FileFinalizedObject, uploaded *tg.TLInputMediaUploadedDocument) (*tg.Document, map[string]string, error) {
	if finalized == nil || finalized.ObjectId == "" || uploaded == nil {
		return nil, nil, wrapMediaInvalidUploadedFile("dfs commit document upload", errors.New("missing finalized object"))
	}
	if finalized.Size2 <= 0 {
		return nil, nil, wrapMediaInvalidUploadedFile("dfs commit document upload", errors.New("invalid finalized object size"))
	}
	now := r.repositoryNow()
	docID := stablePositiveID("document:" + finalized.ObjectId)
	accessHash := stablePositiveID("document-access:" + finalized.ObjectId)
	fileReference, err := r.generateDocumentFileReference(ctx, ownerID, docID, accessHash, finalized.ObjectId, now)
	if err != nil {
		return nil, nil, err
	}
	thumbs, thumbObjectIDs, err := r.documentThumbsFromUploadedThumb(ctx, uploaded.Thumb, ownerID)
	if err != nil {
		return nil, nil, err
	}
	return tg.MakeTLDocument(&tg.TLDocument{
		Id:            docID,
		AccessHash:    accessHash,
		FileReference: fileReference,
		Date:          int32(now.Unix()),
		MimeType:      uploaded.MimeType,
		Size2:         finalized.Size2,
		Thumbs:        thumbs,
		VideoThumbs:   nil,
		DcId:          finalized.DcId,
		Attributes:    uploaded.Attributes,
	}).ToDocument(), thumbObjectIDs, nil
}

func (r *Repository) documentThumbsFromUploadedThumb(ctx context.Context, thumb tg.InputFileClazz, ownerID int64) ([]tg.PhotoSizeClazz, map[string]string, error) {
	if thumb == nil {
		return nil, nil, nil
	}
	if r.processorClient == nil {
		return nil, nil, wrapMediaDownstream("mediaprocessor process document thumb", media.ErrMediaDownstream)
	}
	if r.dfsClient == nil {
		return nil, nil, wrapMediaDownstream("dfs upload document thumb", media.ErrMediaDownstream)
	}
	finalized, err := r.dfsClient.CommitUpload(ctx, &dfsapi.TLDfsCommitUpload{
		UploadSessionId: externalUploadSessionID(ownerID, thumb),
		OwnerId:         ownerID,
		File:            thumb,
		Purpose:         "media_thumbnail",
	})
	if err != nil {
		return nil, nil, wrapDfsUploadError("dfs commit document thumb upload", err)
	}
	if finalized == nil || finalized.ObjectId == "" || len(finalized.ReadLease) == 0 {
		return nil, nil, wrapMediaInvalidUploadedFile("dfs commit document thumb upload", errors.New("missing finalized thumb object"))
	}
	processed, err := r.processorClient.ProcessPhoto(ctx, &mediaprocessor.TLMediaProcessorProcessPhoto{
		OwnerId:   ownerID,
		ObjectId:  finalized.ObjectId,
		ReadLease: finalized.ReadLease,
		FileName:  inputFileName(thumb),
	})
	if err != nil {
		return nil, nil, wrapMediaDownstream("mediaprocessor process document thumb", err)
	}
	if processed == nil {
		return nil, nil, wrapMediaInvalidUploadedFile("mediaprocessor process document thumb", errors.New("missing processed thumb"))
	}
	sizes, objectIDs, err := mapProcessedPhotoSizes(processed.Sizes)
	if err != nil {
		return nil, nil, wrapMediaInvalidUploadedFile("mediaprocessor process document thumb", err)
	}
	var selected *tg.TLPhotoSize
	var selectedObjectID string
	for _, size := range sizes {
		normal, ok := size.(*tg.TLPhotoSize)
		if !ok || normal.W <= 0 || normal.H <= 0 {
			continue
		}
		objectID := objectIDs[normal.Type]
		if objectID == "" {
			continue
		}
		if selected == nil || int64(normal.W)*int64(normal.H) > int64(selected.W)*int64(selected.H) {
			selected = normal
			selectedObjectID = objectID
		}
	}
	if selected == nil {
		return nil, nil, wrapMediaInvalidUploadedFile("mediaprocessor process document thumb", errors.New("missing normal document thumb"))
	}
	return []tg.PhotoSizeClazz{
		tg.MakeTLPhotoSize(&tg.TLPhotoSize{
			Type:  documentThumbSizeType,
			W:     selected.W,
			H:     selected.H,
			Size2: selected.Size2,
		}),
	}, map[string]string{documentThumbSizeType: selectedObjectID}, nil
}

func (r *Repository) documentFromProcessedUpload(ctx context.Context, ownerID int64, finalized *dfsapi.FileFinalizedObject, processed *mediaprocessor.ProcessedDocument) (*tg.Document, map[string]string, error) {
	if finalized == nil || processed == nil || processed.ObjectId == "" || processed.MimeType == "" || processed.Size2 <= 0 {
		return nil, nil, wrapMediaInvalidUploadedFile("mediaprocessor process document", errors.New("missing processed document object"))
	}
	attrs, err := decodeDocumentAttributeVector(processed.Attributes)
	if err != nil {
		return nil, nil, err
	}
	if len(attrs) == 0 {
		return nil, nil, wrapMediaInvalidUploadedFile("mediaprocessor process document", errors.New("missing processed document attributes"))
	}
	thumbs, thumbObjectIDs, err := mapProcessedDocumentThumbs(processed.Thumbs)
	if err != nil {
		return nil, nil, wrapMediaInvalidUploadedFile("mediaprocessor process document", err)
	}
	now := r.repositoryNow()
	docID := stablePositiveID("document:" + processed.ObjectId)
	accessHash := stablePositiveID("document-access:" + processed.ObjectId)
	fileReference, err := r.generateDocumentFileReference(ctx, ownerID, docID, accessHash, processed.ObjectId, now)
	if err != nil {
		return nil, nil, err
	}
	return tg.MakeTLDocument(&tg.TLDocument{
		Id:            docID,
		AccessHash:    accessHash,
		FileReference: fileReference,
		Date:          int32(now.Unix()),
		MimeType:      processed.MimeType,
		Size2:         processed.Size2,
		Thumbs:        thumbs,
		VideoThumbs:   nil,
		DcId:          finalized.DcId,
		Attributes:    attrs,
	}).ToDocument(), thumbObjectIDs, nil
}

func (r *Repository) generateDocumentFileReference(ctx context.Context, ownerID, docID, accessHash int64, objectID string, now time.Time) ([]byte, error) {
	if r == nil || r.fileReferenceService == nil {
		return nil, media.ErrFileReferenceInvalid
	}
	ttl := r.fileReferenceTTL
	if ttl <= 0 {
		ttl = 24 * time.Hour
	}
	fileReference, err := r.fileReferenceService.Generate(ctx, FileReferenceClaims{
		MediaID:      docID,
		ObjectID:     objectID,
		OriginDomain: "document",
		OriginID:     ownerID,
		ExpireAt:     now.Add(ttl).Unix(),
		AccessHash:   accessHash,
	}, r)
	if err != nil {
		return nil, err
	}
	return fileReference, nil
}

func (r *Repository) refreshDocumentFileReference(ctx context.Context, ownerID int64, doc *tg.Document) error {
	if r == nil || r.fileReferenceService == nil {
		return media.ErrFileReferenceInvalid
	}
	do, ok := doc.ToDocument()
	if !ok {
		return media.ErrMediaInvalidArgument
	}
	now := r.repositoryNow()
	fileReference, err := r.generateDocumentFileReference(ctx, ownerID, do.Id, do.AccessHash, documentObjectPath(do.Id), now)
	if err != nil {
		return err
	}
	do.FileReference = fileReference
	return nil
}

func (r *Repository) repositoryNow() time.Time {
	if r != nil && r.fileReferenceService != nil {
		return r.fileReferenceService.now()
	}
	return time.Now()
}

func mapProcessedDocumentThumbs(derivatives []mediaprocessor.ProcessorDerivativeClazz) ([]tg.PhotoSizeClazz, map[string]string, error) {
	const (
		maxInt32                         = int64(^uint32(0) >> 1)
		processorDerivativeDocumentThumb = "document_thumb"
	)
	if len(derivatives) == 0 {
		return nil, nil, nil
	}
	sizes := make([]tg.PhotoSizeClazz, 0, len(derivatives))
	objectIDs := make(map[string]string, len(derivatives))
	for i, derivative := range derivatives {
		if derivative == nil {
			return nil, nil, errors.New("document thumb derivative is nil")
		}
		if derivative.Kind != processorDerivativeDocumentThumb {
			return nil, nil, fmt.Errorf("document thumb derivative %d has unknown kind %q", i, derivative.Kind)
		}
		if derivative.ObjectId == "" {
			return nil, nil, fmt.Errorf("document thumb derivative %d missing object id", i)
		}
		if derivative.Width <= 0 || derivative.Height <= 0 {
			return nil, nil, fmt.Errorf("document thumb derivative %d has invalid dimensions", i)
		}
		if derivative.Size2 <= 0 || derivative.Size2 > maxInt32 {
			return nil, nil, fmt.Errorf("document thumb derivative %d has invalid size", i)
		}
		if _, exists := objectIDs[documentThumbSizeType]; exists {
			return nil, nil, fmt.Errorf("document thumb derivative %d duplicates size type %q", i, documentThumbSizeType)
		}
		sizes = append(sizes, tg.MakeTLPhotoSize(&tg.TLPhotoSize{
			Type:  documentThumbSizeType,
			W:     derivative.Width,
			H:     derivative.Height,
			Size2: int32(derivative.Size2),
		}))
		objectIDs[documentThumbSizeType] = derivative.ObjectId
	}
	return sizes, objectIDs, nil
}

func (r *Repository) GetDocumentList(ctx context.Context, ids []int64) (*media.VectorDocument, error) {
	out := &media.VectorDocument{Datas: make([]tg.DocumentClazz, 0, len(ids))}
	for _, id := range ids {
		doc, err := r.GetDocument(ctx, id)
		if err != nil {
			if errorsIsDocumentNotFound(err) {
				continue
			}
			return nil, err
		}
		docClazz, ok := doc.ToDocument()
		if !ok {
			return nil, media.ErrMediaInvalidArgument
		}
		out.Datas = append(out.Datas, docClazz)
	}
	return out, nil
}

func (r *Repository) mapDocumentWithThumbs(ctx context.Context, doc *model.Documents) (*tg.Document, error) {
	var thumbs []model.PhotoSizes
	if doc.ThumbId != 0 {
		var err error
		thumbs, err = r.model.PhotoSizesModel.SelectListByPhotoSizeId(ctx, doc.ThumbId)
		if err != nil {
			return nil, wrapStorage("load document thumbs", err)
		}
	}
	var videoThumbs []model.VideoSizes
	if doc.VideoThumbId != 0 {
		var err error
		videoThumbs, err = r.model.VideoSizesModel.SelectListByVideoSizeId(ctx, doc.VideoThumbId)
		if err != nil {
			return nil, wrapStorage("load document video thumbs", err)
		}
	}
	fileReference, err := r.generateLoadedFileReference(ctx, "document", doc.DocumentId, doc.AccessHash, firstNonEmpty(doc.FilePath, fmt.Sprintf("document:%d", doc.DocumentId)))
	if err != nil {
		return nil, err
	}
	return mapDocumentAggregate(doc, thumbs, videoThumbs, fileReference)
}

func (r *Repository) saveDocumentAggregate(ctx context.Context, uploadedFileName string, doc *tg.Document) error {
	return r.saveDocumentAggregateWithPaths(ctx, uploadedFileName, doc, "", nil)
}

func (r *Repository) saveDocumentAggregateWithPaths(ctx context.Context, uploadedFileName string, doc *tg.Document, documentObjectID string, thumbObjectIDs map[string]string) error {
	if r == nil || r.model == nil || doc == nil {
		return nil
	}
	do, ok := doc.ToDocument()
	if !ok {
		return media.ErrMediaInvalidArgument
	}
	attrs, err := encodeDocumentAttributesForStorage(do.Attributes)
	if err != nil {
		return err
	}
	row := &model.Documents{
		DocumentId:       do.Id,
		AccessHash:       do.AccessHash,
		DcId:             do.DcId,
		FilePath:         firstNonEmpty(documentObjectID, documentObjectPath(do.Id)),
		FileSize:         do.Size2,
		UploadedFileName: uploadedFileName,
		Ext:              strings.TrimPrefix(strings.ToLower(filepath.Ext(uploadedFileName)), "."),
		MimeType:         do.MimeType,
		Attributes:       attrs,
		Date2:            int64(do.Date),
	}
	if len(do.Thumbs) > 0 {
		row.ThumbId = do.Id
	}
	if len(do.VideoThumbs) > 0 {
		row.VideoThumbId = do.Id
	}
	if _, err := r.model.DocumentsModel.Insert2(ctx, row); err != nil {
		return wrapStorage("save document", err)
	}
	for _, thumb := range do.Thumbs {
		if err := r.savePhotoSizeWithPath(ctx, do.Id, thumb, thumbObjectIDs); err != nil {
			return err
		}
	}
	for _, thumb := range do.VideoThumbs {
		if err := r.saveVideoSize(ctx, do.Id, thumb); err != nil {
			return err
		}
	}
	return nil
}

func isAnimatedGif(uploaded *tg.TLInputMediaUploadedDocument) bool {
	if uploaded.MimeType != "image/gif" {
		return false
	}
	for _, attr := range uploaded.Attributes {
		if _, ok := attr.(*tg.TLDocumentAttributeAnimated); ok {
			return true
		}
	}
	return false
}

func uploadedFileName(uploaded *tg.TLInputMediaUploadedDocument) string {
	if uploaded == nil {
		return ""
	}
	for _, attr := range uploaded.Attributes {
		if filename, ok := attr.(*tg.TLDocumentAttributeFilename); ok {
			return filename.FileName
		}
	}
	return inputFileName(uploaded.File)
}

func documentObjectPath(id int64) string {
	return strconv.FormatInt(id, 10) + ".dat"
}

func errorsIsDocumentNotFound(err error) bool {
	return errors.Is(err, media.ErrDocumentNotFound)
}
