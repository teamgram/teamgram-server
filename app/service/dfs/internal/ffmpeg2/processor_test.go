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
