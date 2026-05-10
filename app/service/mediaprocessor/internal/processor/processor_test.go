package processor

import (
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/media/ffmpeg2"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestKindConstants(t *testing.T) {
	if DerivativePhotoSize != "photo_size" {
		t.Fatalf("DerivativePhotoSize = %q", DerivativePhotoSize)
	}
	if DerivativeDocumentThumb != "document_thumb" {
		t.Fatalf("DerivativeDocumentThumb = %q", DerivativeDocumentThumb)
	}
}

func TestEncodeDocumentAttributesRoundTrip(t *testing.T) {
	tests := []struct {
		name       string
		animated   bool
		wantLength int
	}{
		{name: "mp4", wantLength: 2},
		{name: "gif converted mp4", animated: true, wantLength: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded, err := EncodeDocumentAttributes(&ffmpeg2.VideoMetadata{
				Width:    640,
				Height:   360,
				Duration: 7,
			}, "clip.mp4", tt.animated)
			if err != nil {
				t.Fatalf("EncodeDocumentAttributes() error = %v", err)
			}

			decoded, err := DecodeDocumentAttributes(encoded)
			if err != nil {
				t.Fatalf("DecodeDocumentAttributes() error = %v", err)
			}
			if len(decoded) != tt.wantLength {
				t.Fatalf("decoded len = %d, want %d", len(decoded), tt.wantLength)
			}

			video, ok := decoded[0].(*tg.TLDocumentAttributeVideo)
			if !ok {
				t.Fatalf("decoded[0] = %T, want *tg.TLDocumentAttributeVideo", decoded[0])
			}
			if video.W != 640 || video.H != 360 || video.Duration != 7 || !video.SupportsStreaming {
				t.Fatalf("video attribute = %+v", video)
			}

			fileName, ok := decoded[1].(*tg.TLDocumentAttributeFilename)
			if !ok {
				t.Fatalf("decoded[1] = %T, want *tg.TLDocumentAttributeFilename", decoded[1])
			}
			if fileName.FileName != "clip.mp4" {
				t.Fatalf("filename = %q", fileName.FileName)
			}

			if tt.animated {
				if _, ok := decoded[2].(*tg.TLDocumentAttributeAnimated); !ok {
					t.Fatalf("decoded[2] = %T, want *tg.TLDocumentAttributeAnimated", decoded[2])
				}
			}
		})
	}
}
