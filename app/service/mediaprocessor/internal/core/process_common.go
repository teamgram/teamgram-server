package core

import (
	"bytes"
	"path"
	"strings"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/app/service/mediaprocessor/internal/processor"
	"github.com/teamgram/teamgram-server/v2/app/service/mediaprocessor/mediaprocessor"
	"github.com/teamgram/teamgram-server/v2/pkg/media/ffmpeg2"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const (
	mediaDerivativePurpose = "media_derivative"
	jpegMimeType           = "image/jpeg"
	mp4MimeType            = "video/mp4"
	minGifBytes            = 10240
	readLeaseChunkBytes    = 4 * 1024 * 1024
)

func validProcessInput(ownerID int64, objectID string, readLease []byte, fileName string) bool {
	return ownerID != 0 && objectID != "" && len(readLease) != 0 && fileName != ""
}

func (c *MediaProcessorCore) readOriginalBytes(readLease []byte) ([]byte, error) {
	var out bytes.Buffer
	for offset := int64(0); ; {
		original, err := c.svcCtx.DfsClient.DfsGetFileByReadLease(c.ctx, &dfs.TLDfsGetFileByReadLease{
			ReadLease: readLease,
			Offset:    offset,
			Limit:     readLeaseChunkBytes,
		})
		if err != nil {
			return nil, err
		}
		uploadFile, ok := original.ToUploadFile()
		if !ok {
			return nil, mediaprocessor.ErrMediaProcessorInvalidArgument
		}
		if len(uploadFile.Bytes) == 0 {
			break
		}
		if _, err := out.Write(uploadFile.Bytes); err != nil {
			return nil, err
		}
		offset += int64(len(uploadFile.Bytes))
		if len(uploadFile.Bytes) < readLeaseChunkBytes {
			break
		}
	}
	return out.Bytes(), nil
}

func makeDerivative(kind string, stored *dfs.FileFinalizedObject, fileName, mimeType string, size int64, width, height int32, bytes []byte) *mediaprocessor.ProcessorDerivative {
	if stored != nil {
		if stored.ObjectId != "" {
			objectID := stored.ObjectId
			if stored.Size2 > 0 {
				size = stored.Size2
			}
			return mediaprocessor.MakeTLProcessorDerivative(&mediaprocessor.TLProcessorDerivative{
				Kind:     kind,
				ObjectId: objectID,
				FileName: fileName,
				MimeType: mimeType,
				Size2:    size,
				Width:    width,
				Height:   height,
				Bytes:    bytes,
			}).ToProcessorDerivative()
		}
		if stored.Size2 > 0 {
			size = stored.Size2
		}
	}
	return mediaprocessor.MakeTLProcessorDerivative(&mediaprocessor.TLProcessorDerivative{
		Kind:     kind,
		FileName: fileName,
		MimeType: mimeType,
		Size2:    size,
		Width:    width,
		Height:   height,
		Bytes:    bytes,
	}).ToProcessorDerivative()
}

func putDerivative(c *MediaProcessorCore, ownerID int64, fileName, mimeType string, bytes []byte) (*dfs.FileFinalizedObject, error) {
	stored, err := c.svcCtx.DfsClient.DfsPutFile(c.ctx, &dfs.TLDfsPutFile{
		OwnerId:  ownerID,
		Purpose:  mediaDerivativePurpose,
		FileName: fileName,
		MimeType: mimeType,
		Bytes:    bytes,
	})
	if err != nil {
		return nil, err
	}
	if stored == nil || stored.ObjectId == "" {
		return nil, mediaprocessor.ErrMediaProcessorInvalidArgument
	}
	return stored, nil
}

func mp4FileName(fileName string) string {
	ext := path.Ext(fileName)
	if ext == "" {
		return fileName + ".mp4"
	}
	return strings.TrimSuffix(fileName, ext) + ".mp4"
}

func thumbFileName(fileName string) string {
	ext := path.Ext(fileName)
	if ext == "" {
		return fileName + "_thumb.jpg"
	}
	return strings.TrimSuffix(fileName, ext) + "_thumb.jpg"
}

func encodeVideoAttributes(metadata *ffmpeg2.VideoMetadata, fileName string, animated bool) ([]byte, error) {
	return processor.EncodeDocumentAttributes(metadata, fileName, animated)
}

func makeThumbDerivative(stored *dfs.FileFinalizedObject, fileName string, metadata *ffmpeg2.VideoMetadata, cover []byte) *mediaprocessor.ProcessorDerivative {
	var width, height int32
	if metadata != nil {
		width = metadata.Width
		height = metadata.Height
	}
	return makeDerivative(processor.DerivativeDocumentThumb, stored, fileName, jpegMimeType, int64(len(cover)), width, height, cover)
}

func profileBool(v tg.BoolClazz) bool {
	return tg.FromBoolClazz(v)
}
