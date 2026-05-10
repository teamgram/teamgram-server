package ffmpeg2

import (
	"bytes"
	"context"
	"errors"
	"testing"
)

func TestExternalProcessorReportsUnavailable(t *testing.T) {
	processor := NewProcessor()
	if _, err := processor.ConvertToMP4(context.Background(), bytes.NewReader([]byte("gif"))); !errors.Is(err, ErrProcessorUnavailable) {
		t.Fatalf("ConvertToMP4() error = %v, want ErrProcessorUnavailable", err)
	}
	if _, err := processor.ExtractFirstFrame(context.Background(), bytes.NewReader([]byte("mp4"))); !errors.Is(err, ErrProcessorUnavailable) {
		t.Fatalf("ExtractFirstFrame() error = %v, want ErrProcessorUnavailable", err)
	}
	if _, err := processor.GetVideoMetadata(context.Background(), bytes.NewReader([]byte("mp4"))); !errors.Is(err, ErrProcessorUnavailable) {
		t.Fatalf("GetVideoMetadata() error = %v, want ErrProcessorUnavailable", err)
	}
}

func TestDecodeMetadataReturnsAudioOnlyDuration(t *testing.T) {
	metadata, err := decodeMetadata([]byte(`{"streams":[],"format":{"duration":"11.2"}}`))
	if err != nil {
		t.Fatalf("decodeMetadata() error = %v", err)
	}
	if metadata.Width != 0 || metadata.Height != 0 || metadata.Duration != 11 {
		t.Fatalf("metadata = %+v, want duration-only 11 seconds", metadata)
	}
}
