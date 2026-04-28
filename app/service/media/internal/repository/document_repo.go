package repository

import (
	"context"

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
	return documentFromModel(do), nil
}

func (r *Repository) GetDocumentByRequest(ctx context.Context, in *media.TLMediaGetDocument) (*tg.Document, error) {
	return r.GetDocument(ctx, in.Id)
}

func documentFromModel(do *model.Documents) *tg.Document {
	if do == nil {
		return nil
	}
	return tg.MakeTLDocument(&tg.TLDocument{
		Id:            do.DocumentId,
		AccessHash:    do.AccessHash,
		FileReference: []byte{},
		Date:          int32(do.Date2),
		MimeType:      do.MimeType,
		Size2:         do.FileSize,
		DcId:          do.DcId,
		Attributes:    []tg.DocumentAttributeClazz{},
	}).ToDocument()
}
