package processor

import (
	"bytes"
	"context"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/pkg/media/ffmpeg2"
	"github.com/teamgram/teamgram-server/v2/pkg/media/imaging2"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const (
	DerivativePhotoSize     = "photo_size"
	DerivativePhotoStripped = "photo_stripped"
	DerivativeDocumentThumb = "document_thumb"
	DocumentAttributeLayer  = 224
)

type MediaProcessor interface {
	ResizePhoto(ctx context.Context, input []byte, ext string, profile bool) ([]imaging2.PhotoSizeBytes, error)
	BuildPhotoDerivatives(ctx context.Context, input []byte, ext string, profile bool) ([]imaging2.PhotoDerivativeBytes, error)
	ConvertGIFToMP4(ctx context.Context, input []byte) ([]byte, error)
	ExtractMP4Cover(ctx context.Context, input []byte) ([]byte, error)
	ProbeMP4(ctx context.Context, input []byte) (*ffmpeg2.VideoMetadata, error)
}

type Processor struct {
	imaging imaging2.Processor
	ffmpeg  ffmpeg2.Processor
}

func New() *Processor {
	return NewWithDeps(imaging2.NewProcessor(), ffmpeg2.NewProcessor())
}

func NewWithDeps(imaging imaging2.Processor, ffmpeg ffmpeg2.Processor) *Processor {
	return &Processor{
		imaging: imaging,
		ffmpeg:  ffmpeg,
	}
}

func (p *Processor) ResizePhoto(ctx context.Context, input []byte, ext string, profile bool) ([]imaging2.PhotoSizeBytes, error) {
	return p.imaging.ResizePhoto(ctx, input, ext, profile)
}

func (p *Processor) BuildPhotoDerivatives(ctx context.Context, input []byte, ext string, profile bool) ([]imaging2.PhotoDerivativeBytes, error) {
	return p.imaging.BuildPhotoDerivatives(ctx, input, ext, profile)
}

func (p *Processor) ConvertGIFToMP4(ctx context.Context, input []byte) ([]byte, error) {
	return p.ffmpeg.ConvertToMP4(ctx, bytes.NewReader(input))
}

func (p *Processor) ExtractMP4Cover(ctx context.Context, input []byte) ([]byte, error) {
	return p.ffmpeg.ExtractFirstFrame(ctx, bytes.NewReader(input))
}

func (p *Processor) ProbeMP4(ctx context.Context, input []byte) (*ffmpeg2.VideoMetadata, error) {
	return p.ffmpeg.GetVideoMetadata(ctx, bytes.NewReader(input))
}

func EncodeDocumentAttributes(metadata *ffmpeg2.VideoMetadata, fileName string, animated bool, sourceAttrs ...[]tg.DocumentAttributeClazz) ([]byte, error) {
	if metadata == nil {
		metadata = &ffmpeg2.VideoMetadata{}
	}
	sourceVideo := firstVideoAttribute(sourceAttrs...)
	attrs := []tg.DocumentAttributeClazz{
		tg.MakeTLDocumentAttributeVideo(&tg.TLDocumentAttributeVideo{
			SupportsStreaming: true,
			RoundMessage:      sourceVideo.RoundMessage,
			Nosound:           sourceVideo.Nosound,
			Duration:          float64(metadata.Duration),
			W:                 metadata.Width,
			H:                 metadata.Height,
			PreloadPrefixSize: sourceVideo.PreloadPrefixSize,
			VideoStartTs:      sourceVideo.VideoStartTs,
			VideoCodec:        sourceVideo.VideoCodec,
		}),
		tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{
			FileName: fileName,
		}),
	}
	if animated {
		attrs = append(attrs, tg.MakeTLDocumentAttributeAnimated(&tg.TLDocumentAttributeAnimated{}))
	}

	x := bin.NewEncoder()
	if err := iface.EncodeObjectList(x, attrs, DocumentAttributeLayer); err != nil {
		return nil, fmt.Errorf("encode document attributes: %w", err)
	}
	return x.Clone(), nil
}

func firstVideoAttribute(sourceAttrs ...[]tg.DocumentAttributeClazz) *tg.TLDocumentAttributeVideo {
	for _, attrs := range sourceAttrs {
		for _, attr := range attrs {
			if video, ok := attr.(*tg.TLDocumentAttributeVideo); ok {
				return video
			}
		}
	}
	return &tg.TLDocumentAttributeVideo{}
}

func DecodeDocumentAttributes(data []byte) ([]tg.DocumentAttributeClazz, error) {
	if len(data) == 0 {
		return []tg.DocumentAttributeClazz{}, nil
	}
	attrs, err := iface.DecodeObjectList[tg.DocumentAttributeClazz](bin.NewDecoder(data))
	if err != nil {
		return nil, fmt.Errorf("decode document attributes: %w", err)
	}
	return attrs, nil
}
