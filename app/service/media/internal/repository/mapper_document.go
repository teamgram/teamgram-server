package repository

import (
	"encoding/json"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/service/media/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const documentAttributeVectorLayer = 224

func mapDocumentAggregate(doc *model.Documents, thumbs []model.PhotoSizes, videoThumbs []model.VideoSizes, fileReference []byte) (*tg.Document, error) {
	if doc == nil {
		return nil, media.ErrDocumentNotFound
	}
	attrs, err := decodeLegacyDocumentAttributes(doc.Attributes)
	if err != nil {
		return nil, err
	}
	return tg.MakeTLDocument(&tg.TLDocument{
		Id:            doc.DocumentId,
		AccessHash:    doc.AccessHash,
		FileReference: fileReference,
		Date:          int32(doc.Date2),
		MimeType:      doc.MimeType,
		Size2:         doc.FileSize,
		Thumbs:        mapOptionalPhotoSizes(thumbs),
		VideoThumbs:   mapVideoSizes(videoThumbs),
		DcId:          doc.DcId,
		Attributes:    attrs,
	}).ToDocument(), nil
}

func mapOptionalPhotoSizes(sizes []model.PhotoSizes) []tg.PhotoSizeClazz {
	if len(sizes) == 0 {
		return nil
	}
	return mapPhotoSizes(sizes)
}

func decodeLegacyDocumentAttributes(raw string) ([]tg.DocumentAttributeClazz, error) {
	if raw == "" {
		return []tg.DocumentAttributeClazz{}, nil
	}
	var items []legacyAttribute
	if err := json.Unmarshal([]byte(raw), &items); err != nil {
		return nil, fmt.Errorf("%w: decode document attributes: %w", media.ErrMediaStorage, err)
	}
	out := make([]tg.DocumentAttributeClazz, 0, len(items))
	for _, item := range items {
		switch item.Name {
		case tg.ClazzName_documentAttributeFilename:
			out = append(out, tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{FileName: item.Object.FileName}))
		case tg.ClazzName_documentAttributeImageSize:
			out = append(out, tg.MakeTLDocumentAttributeImageSize(&tg.TLDocumentAttributeImageSize{W: item.Object.W, H: item.Object.H}))
		case tg.ClazzName_documentAttributeAnimated:
			out = append(out, tg.MakeTLDocumentAttributeAnimated(&tg.TLDocumentAttributeAnimated{}))
		case tg.ClazzName_documentAttributeVideo:
			out = append(out, tg.MakeTLDocumentAttributeVideo(&tg.TLDocumentAttributeVideo{
				RoundMessage:      item.Object.RoundMessage,
				SupportsStreaming: item.Object.SupportsStreaming,
				Nosound:           item.Object.Nosound,
				Duration:          item.Object.Duration,
				W:                 item.Object.W,
				H:                 item.Object.H,
			}))
		case tg.ClazzName_documentAttributeAudio:
			out = append(out, tg.MakeTLDocumentAttributeAudio(&tg.TLDocumentAttributeAudio{
				Voice:     item.Object.Voice,
				Duration:  item.Object.DurationInt32(),
				Title:     item.Object.Title,
				Performer: item.Object.Performer,
				Waveform:  item.Object.Waveform,
			}))
		}
	}
	return out, nil
}

func encodeDocumentAttributeVector(attrs []tg.DocumentAttributeClazz) ([]byte, error) {
	x := bin.NewEncoder()
	if err := iface.EncodeObjectList(x, attrs, documentAttributeVectorLayer); err != nil {
		return nil, fmt.Errorf("%w: encode document attribute vector: %w", media.ErrMediaInvalidUploadedFile, err)
	}
	return x.Clone(), nil
}

func decodeDocumentAttributeVector(raw []byte) ([]tg.DocumentAttributeClazz, error) {
	if len(raw) == 0 {
		return []tg.DocumentAttributeClazz{}, nil
	}
	attrs, err := iface.DecodeObjectList[tg.DocumentAttributeClazz](bin.NewDecoder(raw))
	if err != nil {
		return nil, fmt.Errorf("%w: decode document attribute vector: %w", media.ErrMediaInvalidUploadedFile, err)
	}
	return attrs, nil
}

type legacyAttribute struct {
	Name   string                `json:"_name"`
	Object legacyAttributeObject `json:"_object"`
}

type legacyAttributeObject struct {
	FileName          string  `json:"file_name"`
	W                 int32   `json:"w"`
	H                 int32   `json:"h"`
	RoundMessage      bool    `json:"round_message"`
	SupportsStreaming bool    `json:"supports_streaming"`
	Nosound           bool    `json:"nosound"`
	Duration          float64 `json:"duration"`
	Voice             bool    `json:"voice"`
	Title             *string `json:"title"`
	Performer         *string `json:"performer"`
	Waveform          []byte  `json:"waveform"`
}

func (o legacyAttributeObject) DurationInt32() int32 {
	return int32(o.Duration)
}
