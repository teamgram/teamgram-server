package imaging2

import (
	"bytes"
	"context"
	"image/jpeg"
	"os"
	"reflect"
	"testing"
)

func TestProgressiveScanSizesRejectsBaselineJPEG(t *testing.T) {
	_, err := ProgressiveScanSizes(testSizedJPEG(t, 640, 480))
	if err == nil {
		t.Fatal("ProgressiveScanSizes(baseline) error = nil, want error")
	}
}

func TestProgressiveScanSizesRejectsProgressiveJPEGWithoutScans(t *testing.T) {
	data := []byte{
		0xff, 0xd8,
		0xff, 0xc2, 0x00, 0x02,
		0xff, 0xd9,
	}
	_, err := ProgressiveScanSizes(data)
	if err == nil {
		t.Fatal("ProgressiveScanSizes(SOF2 without SOS) error = nil, want error")
	}
}

func TestProgressiveScanSizesSkipsStuffedAndFillBytesInScanData(t *testing.T) {
	data := []byte{
		0xff, 0xd8,
		0xff, 0xc2, 0x00, 0x02,
		0xff, 0xda, 0x00, 0x02,
		0x11, 0xff, 0x00, 0x22, 0xff,
		0xff, 0xd9,
	}
	sizes, err := ProgressiveScanSizes(data)
	if err != nil {
		t.Fatalf("ProgressiveScanSizes() error = %v", err)
	}
	want := []int32{15, int32(len(data))}
	if !reflect.DeepEqual(sizes, want) {
		t.Fatalf("ProgressiveScanSizes() = %v, want %v", sizes, want)
	}
}

func TestProgressiveScanSizesRejectsInvalidInput(t *testing.T) {
	for _, tc := range []struct {
		name string
		data []byte
	}{
		{name: "empty"},
		{name: "missing SOI", data: []byte("not jpeg")},
		{name: "truncated marker", data: []byte{0xff, 0xd8, 0xff}},
		{name: "truncated segment", data: []byte{0xff, 0xd8, 0xff, 0xc2, 0x00}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ProgressiveScanSizes(tc.data)
			if err == nil {
				t.Fatal("ProgressiveScanSizes() error = nil, want error")
			}
		})
	}
}

func TestImageMagickArgsPreventsUpscaling(t *testing.T) {
	args := imageMagickArgs(1280, 85)
	for i := 0; i < len(args)-1; i++ {
		if args[i] == "-resize" {
			if args[i+1] != "1280x1280>" {
				t.Fatalf("resize arg = %q, want 1280x1280>", args[i+1])
			}
			return
		}
	}
	t.Fatalf("imageMagickArgs() = %v, want -resize argument", args)
}

func TestResolveImageMagickBinaryRejectsUnsupportedConfiguredExecutable(t *testing.T) {
	exe, err := os.Executable()
	if err != nil {
		t.Fatalf("os.Executable(): %v", err)
	}
	_, err = ResolveImageMagickBinary(exe)
	if err == nil {
		t.Fatal("ResolveImageMagickBinary(test binary) error = nil, want error")
	}
}

func TestImageMagickProgressiveEncoderProducesDecodeableProgressiveJPEG(t *testing.T) {
	binary, err := ResolveImageMagickBinary("")
	if err != nil {
		t.Skipf("ImageMagick binary unavailable: %v", err)
	}

	encoder := ImageMagickProgressiveEncoder{Binary: binary}
	data, scanSizes, err := encoder.EncodeProgressiveJPEG(context.Background(), testSizedJPEG(t, 2000, 616), "jpg", 1280)
	if err != nil {
		t.Fatalf("EncodeProgressiveJPEG() error = %v", err)
	}
	if len(data) == 0 {
		t.Fatal("EncodeProgressiveJPEG() returned empty data")
	}

	cfg, err := jpeg.DecodeConfig(bytes.NewReader(data))
	if err != nil {
		t.Fatalf("jpeg.DecodeConfig(progressive) error = %v", err)
	}
	if cfg.Width != 1280 || cfg.Height != 394 {
		t.Fatalf("progressive dimensions = %dx%d, want 1280x394", cfg.Width, cfg.Height)
	}
	if len(scanSizes) < 2 {
		t.Fatalf("scan sizes = %v, want at least two entries", scanSizes)
	}
	for i := range scanSizes {
		if scanSizes[i] <= 0 {
			t.Fatalf("scanSizes[%d] = %d, want positive", i, scanSizes[i])
		}
		if i > 0 && scanSizes[i] <= scanSizes[i-1] {
			t.Fatalf("scan sizes = %v, want strictly increasing", scanSizes)
		}
	}
	if got := scanSizes[len(scanSizes)-1]; got != int32(len(data)) {
		t.Fatalf("last scan size = %d, want full length %d", got, len(data))
	}

	parsed, err := ProgressiveScanSizes(data)
	if err != nil {
		t.Fatalf("ProgressiveScanSizes(encoded) error = %v", err)
	}
	if len(parsed) != len(scanSizes) {
		t.Fatalf("parsed scan sizes = %v, encoder scan sizes = %v", parsed, scanSizes)
	}
	for i := range parsed {
		if parsed[i] != scanSizes[i] {
			t.Fatalf("parsed scan sizes = %v, encoder scan sizes = %v", parsed, scanSizes)
		}
	}
}
