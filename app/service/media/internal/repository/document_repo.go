package repository

import (
	"context"
	"encoding/json"
	"errors"
	"path/filepath"
	"strconv"
	"strings"

	dfsapi "github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/app/service/media/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

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

func documentFromModel(do *model.Documents) *tg.Document {
	doc, err := mapDocumentAggregate(do, nil, nil)
	if err != nil {
		return nil
	}
	return doc
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

	var (
		doc *tg.Document
		err error
	)
	switch {
	case isAnimatedGif(uploaded):
		doc, err = r.dfsClient.UploadGifDocumentMedia(ctx, &dfsapi.TLDfsUploadGifDocumentMedia{Creator: in.OwnerId, Media: uploaded})
	case uploaded.MimeType == "video/mp4":
		doc, err = r.dfsClient.UploadMp4DocumentMedia(ctx, &dfsapi.TLDfsUploadMp4DocumentMedia{Creator: in.OwnerId, Media: uploaded})
	default:
		doc, err = r.dfsClient.UploadDocumentFileV2(ctx, &dfsapi.TLDfsUploadDocumentFileV2{Creator: in.OwnerId, Media: uploaded})
	}
	if err != nil {
		return nil, wrapDfsUploadError("dfs upload document", err)
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
	return mapDocumentAggregate(doc, thumbs, videoThumbs)
}

func (r *Repository) saveDocumentAggregate(ctx context.Context, uploadedFileName string, doc *tg.Document) error {
	if r == nil || r.model == nil || doc == nil {
		return nil
	}
	do, ok := doc.ToDocument()
	if !ok {
		return media.ErrMediaInvalidArgument
	}
	attrs, err := encodeLegacyDocumentAttributes(do.Attributes)
	if err != nil {
		return err
	}
	row := &model.Documents{
		DocumentId:       do.Id,
		AccessHash:       do.AccessHash,
		DcId:             do.DcId,
		FilePath:         documentObjectPath(do.Id),
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
		if err := r.savePhotoSize(ctx, do.Id, thumb); err != nil {
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

func encodeLegacyDocumentAttributes(attrs []tg.DocumentAttributeClazz) (string, error) {
	if len(attrs) == 0 {
		return "", nil
	}
	items := make([]legacyAttribute, 0, len(attrs))
	for _, attr := range attrs {
		item := legacyAttribute{}
		switch a := attr.(type) {
		case *tg.TLDocumentAttributeFilename:
			item.Name = tg.ClazzName_documentAttributeFilename
			item.Object.FileName = a.FileName
		case *tg.TLDocumentAttributeImageSize:
			item.Name = tg.ClazzName_documentAttributeImageSize
			item.Object.W = a.W
			item.Object.H = a.H
		case *tg.TLDocumentAttributeAnimated:
			item.Name = tg.ClazzName_documentAttributeAnimated
		case *tg.TLDocumentAttributeVideo:
			item.Name = tg.ClazzName_documentAttributeVideo
			item.Object.RoundMessage = a.RoundMessage
			item.Object.SupportsStreaming = a.SupportsStreaming
			item.Object.Nosound = a.Nosound
			item.Object.Duration = a.Duration
			item.Object.W = a.W
			item.Object.H = a.H
		case *tg.TLDocumentAttributeAudio:
			item.Name = tg.ClazzName_documentAttributeAudio
			item.Object.Voice = a.Voice
			item.Object.Duration = float64(a.Duration)
			item.Object.Title = a.Title
			item.Object.Performer = a.Performer
			item.Object.Waveform = a.Waveform
		default:
			continue
		}
		items = append(items, item)
	}
	if len(items) == 0 {
		return "", nil
	}
	b, err := json.Marshal(items)
	if err != nil {
		return "", wrapStorage("encode document attributes", err)
	}
	return string(b), nil
}

func errorsIsDocumentNotFound(err error) bool {
	return errors.Is(err, media.ErrDocumentNotFound)
}
