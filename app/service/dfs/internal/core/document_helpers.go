package core

import (
	"context"
	"errors"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository"
	minioadapter "github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/minio"
	"github.com/teamgram/teamgram-server/v2/pkg/media/ffmpeg2"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type documentRepository interface {
	NextDocumentID(ctx context.Context) (int64, error)
	NextEncryptedFileID(ctx context.Context) (int64, error)
	SaveDocumentObject(ctx context.Context, documentID int64, data []byte) (int64, error)
	SaveEncryptedObject(ctx context.Context, fileID int64, data []byte) (int64, error)
	SaveDocumentThumbs(ctx context.Context, documentID int64, image []byte, ext string) ([]repository.StoredDocumentThumb, error)
	ConvertDocumentToMP4(ctx context.Context, data []byte) ([]byte, error)
	ExtractDocumentFrame(ctx context.Context, data []byte) ([]byte, error)
	GetDocumentVideoMetadata(ctx context.Context, data []byte) (*ffmpeg2.VideoMetadata, error)
}

type uploadedDocumentData struct {
	file *uploadedPhotoFile
	data []byte
	date int32
	size int64
	ext  string
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

func (c *DfsCore) readUploadedDocumentData(creator int64, file *uploadedPhotoFile) (*uploadedDocumentData, error) {
	if file == nil {
		return nil, dfs.ErrDfsInvalidArgument
	}
	if err := checkFileParts(file.parts); err != nil {
		return nil, err
	}
	reader, err := c.uploadSessions().OpenUploadedFile(c.ctx, creator, file.id)
	if err != nil {
		return nil, err
	}
	data, err := readAllSeeker(reader)
	if err != nil {
		return nil, dfs.WrapDfsStorage("read uploaded document", err)
	}
	if err := checkMD5(data, file.md5Checksum); err != nil {
		return nil, err
	}
	date := nowUnix()
	size := int64(len(data))
	info, err := c.uploadSessions().LoadUploadedFileInfo(c.ctx, creator, file.id)
	if err != nil && !errors.Is(err, dfs.ErrDfsFileNotFound) {
		return nil, err
	}
	if info != nil {
		if info.Mtime > 0 {
			date = int32(info.Mtime)
		}
		if fileSize := info.FileSize(); fileSize > 0 {
			size = fileSize
		}
	}
	return &uploadedDocumentData{
		file: file,
		data: data,
		date: date,
		size: size,
		ext:  fileExt(file.name),
	}, nil
}

func (c *DfsCore) logNonFatalError(op string, err error) {
	if c != nil && c.Logger != nil && err != nil {
		c.Logger.Errorf("%s - non-fatal error: %v", op, err)
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
	return makeDocumentWithThumbs(id, ext, date, mimeType, size, nil, nil, attrs)
}

func makeDocumentWithThumbs(id int64, ext string, date int32, mimeType string, size int64, thumbs []tg.PhotoSizeClazz, videoThumbs []tg.VideoSizeClazz, attrs []tg.DocumentAttributeClazz) *tg.Document {
	return tg.MakeTLDocument(&tg.TLDocument{
		Id:            id,
		AccessHash:    minioadapter.MakeAccessHash(storageFileTypeConstructor(ext), rand32()),
		FileReference: []byte{},
		Date:          date,
		MimeType:      mimeType,
		Size2:         size,
		Thumbs:        thumbs,
		VideoThumbs:   videoThumbs,
		DcId:          1,
		Attributes:    attrs,
	}).ToDocument()
}

func photoSizesFromStored(stored []repository.StoredDocumentThumb) []tg.PhotoSizeClazz {
	if len(stored) == 0 {
		return nil
	}
	out := make([]tg.PhotoSizeClazz, 0, len(stored))
	for _, size := range stored {
		if size.Type == "i" {
			out = append(out, tg.MakeTLPhotoStrippedSize(&tg.TLPhotoStrippedSize{
				Type:  size.Type,
				Bytes: size.Bytes,
			}))
			continue
		}
		out = append(out, tg.MakeTLPhotoSize(&tg.TLPhotoSize{
			Type:  size.Type,
			W:     size.W,
			H:     size.H,
			Size2: size.Size,
		}))
	}
	return out
}

func imageSizeAttributeFromThumbs(stored []repository.StoredDocumentThumb) tg.DocumentAttributeClazz {
	for _, size := range stored {
		if size.W > 0 && size.H > 0 {
			return tg.MakeTLDocumentAttributeImageSize(&tg.TLDocumentAttributeImageSize{W: size.W, H: size.H})
		}
	}
	return tg.MakeTLDocumentAttributeImageSize(&tg.TLDocumentAttributeImageSize{})
}

func videoAttributesFromMetadata(metadata *ffmpeg2.VideoMetadata, fileName string) []tg.DocumentAttributeClazz {
	var duration int32
	var width int32
	var height int32
	if metadata != nil {
		duration = metadata.Duration
		width = metadata.Width
		height = metadata.Height
	}
	return []tg.DocumentAttributeClazz{
		tg.MakeTLDocumentAttributeVideo(&tg.TLDocumentAttributeVideo{
			SupportsStreaming: true,
			Duration:          float64(duration),
			W:                 width,
			H:                 height,
		}),
		tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{FileName: fileName}),
		tg.MakeTLDocumentAttributeAnimated(&tg.TLDocumentAttributeAnimated{}),
	}
}

func makeMp4DocumentAttributes(attrs []tg.DocumentAttributeClazz) []tg.DocumentAttributeClazz {
	out := make([]tg.DocumentAttributeClazz, 0, len(attrs))
	for _, attr := range attrs {
		switch a := attr.(type) {
		case *tg.TLDocumentAttributeFilename:
			out = append(out, attr)
		case *tg.TLDocumentAttributeVideo:
			out = append(out, tg.MakeTLDocumentAttributeVideo(&tg.TLDocumentAttributeVideo{
				RoundMessage:      a.RoundMessage,
				SupportsStreaming: true,
				Nosound:           a.Nosound,
				Duration:          a.Duration,
				W:                 a.W,
				H:                 a.H,
				PreloadPrefixSize: a.PreloadPrefixSize,
				VideoStartTs:      a.VideoStartTs,
				VideoCodec:        a.VideoCodec,
			}))
		case *tg.TLDocumentAttributeAnimated:
			out = append(out, attr)
		}
	}
	return out
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
