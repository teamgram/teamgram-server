package core

import (
	"context"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	minioadapter "github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/minio"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type documentRepository interface {
	NextDocumentID(ctx context.Context) (int64, error)
	NextEncryptedFileID(ctx context.Context) (int64, error)
	SaveDocumentObject(ctx context.Context, documentID int64, data []byte) (int64, error)
	SaveEncryptedObject(ctx context.Context, fileID int64, data []byte) (int64, error)
}

type uploadedEncryptedFile struct {
	id             int64
	parts          int32
	md5Checksum    string
	keyFingerprint int32
}

func (c *DfsCore) documents() documentRepository {
	if c.documentRepository != nil {
		return c.documentRepository
	}
	return c.svcCtx.Repo
}

func inputMediaUploadedDocument(in tg.InputMediaClazz) (*tg.TLInputMediaUploadedDocument, error) {
	if media, ok := in.(*tg.TLInputMediaUploadedDocument); ok {
		return media, nil
	}
	return nil, dfs.ErrDfsInvalidArgument
}

func inputEncryptedFile(in tg.InputEncryptedFileClazz) (*uploadedEncryptedFile, error) {
	switch f := in.(type) {
	case *tg.TLInputEncryptedFileUploaded:
		return &uploadedEncryptedFile{id: f.Id, parts: f.Parts, md5Checksum: f.Md5Checksum, keyFingerprint: f.KeyFingerprint}, nil
	case *tg.TLInputEncryptedFileBigUploaded:
		return &uploadedEncryptedFile{id: f.Id, parts: f.Parts, keyFingerprint: f.KeyFingerprint}, nil
	default:
		return nil, dfs.ErrDfsInvalidArgument
	}
}

func filterDocumentAttributes(mimeType string, attrs []tg.DocumentAttributeClazz) []tg.DocumentAttributeClazz {
	out := make([]tg.DocumentAttributeClazz, 0, len(attrs))
	for _, attr := range attrs {
		switch a := attr.(type) {
		case *tg.TLDocumentAttributeAnimated:
			continue
		case *tg.TLDocumentAttributeFilename:
			if a.FileName != "" {
				out = append(out, attr)
			}
		case *tg.TLDocumentAttributeAudio:
			if mimeType == "audio/ogg" {
				if a.Voice {
					out = append(out, attr)
				}
			} else {
				out = append(out, attr)
			}
		default:
			out = append(out, attr)
		}
	}
	return out
}

func makeDocument(id int64, ext string, date int32, mimeType string, size int64, attrs []tg.DocumentAttributeClazz) *tg.Document {
	return tg.MakeTLDocument(&tg.TLDocument{
		Id:            id,
		AccessHash:    minioadapter.MakeAccessHash(storageFileTypeConstructor(ext), rand32()),
		FileReference: []byte{},
		Date:          date,
		MimeType:      mimeType,
		Size2:         size,
		Thumbs:        nil,
		VideoThumbs:   nil,
		DcId:          1,
		Attributes:    attrs,
	}).ToDocument()
}

func makeEncryptedFile(id int64, accessHash int64, size int64, keyFingerprint int32) *tg.EncryptedFile {
	return tg.MakeTLEncryptedFile(&tg.TLEncryptedFile{
		Id:             id,
		AccessHash:     accessHash,
		Size2:          size,
		DcId:           1,
		KeyFingerprint: keyFingerprint,
	}).ToEncryptedFile()
}

func documentObjectPath(id int64) string {
	return fmt.Sprintf("%d.dat", id)
}
