package imaging2

import (
	"bytes"
	"context"
	"errors"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"testing"

	"github.com/disintegration/imaging"
)

func TestResizePhotoProducesConfiguredSizes(t *testing.T) {
	processor := NewProcessor()
	sizes, err := processor.ResizePhoto(context.Background(), testJPEG(t), "jpg", false)
	if err != nil {
		t.Fatalf("ResizePhoto() error = %v", err)
	}
	if len(sizes) != 3 {
		t.Fatalf("len(sizes) = %d, want 3", len(sizes))
	}
	for i, wantType := range []string{"s", "m", "x"} {
		size := sizes[i]
		if size.Type != wantType {
			t.Fatalf("size[%d].Type = %q, want %q", i, size.Type, wantType)
		}
		if size.W <= 0 || size.H <= 0 || len(size.Bytes) == 0 {
			t.Fatalf("size[%d] invalid: %+v bytes=%d", i, size, len(size.Bytes))
		}
	}
	if sizes[0].W != PhotoSZSmallSize {
		t.Fatalf("small width = %d, want %d", sizes[0].W, PhotoSZSmallSize)
	}
	if sizes[2].W != 640 || sizes[2].H != 480 {
		t.Fatalf("terminal size = %dx%d, want 640x480", sizes[2].W, sizes[2].H)
	}
}

func TestResizePhotoUsesProfileSizeSet(t *testing.T) {
	processor := NewProcessor()
	sizes, err := processor.ResizePhoto(context.Background(), testJPEG(t), "jpg", true)
	if err != nil {
		t.Fatalf("ResizePhoto(profile) error = %v", err)
	}
	if len(sizes) != 3 {
		t.Fatalf("len(sizes) = %d, want 3", len(sizes))
	}
	if sizes[0].Type != "a" {
		t.Fatalf("first profile size type = %q, want a", sizes[0].Type)
	}
	if got := sizes[len(sizes)-1]; got.Type != "c" || got.W != PhotoSZCSize || got.H != PhotoSZCSize {
		t.Fatalf("terminal profile size = %+v, want c 640x640", got)
	}
}

func TestEncodeStrippedProducesTelegramStrippedPayload(t *testing.T) {
	processor := NewProcessor()
	data, err := processor.EncodeStripped(context.Background(), testJPEG(t))
	if err != nil {
		t.Fatalf("EncodeStripped() error = %v", err)
	}
	if len(data) <= 3 {
		t.Fatalf("EncodeStripped() returned %d bytes, want payload", len(data))
	}
	if data[0] != 0x01 {
		t.Fatalf("stripped marker = %#x, want 0x01", data[0])
	}
	if int(data[1]) > PhotoSZStrippedSize || int(data[2]) > PhotoSZStrippedSize {
		t.Fatalf("stripped dimensions = %dx%d, want max side <= %d", data[2], data[1], PhotoSZStrippedSize)
	}
	if _, err := jpeg.Decode(bytes.NewReader(data)); err == nil {
		t.Fatal("stripped payload decoded as a full JPEG")
	}
}

func TestResizePhotoPreservesSupportedFormat(t *testing.T) {
	processor := NewProcessor()
	sizes, err := processor.ResizePhoto(context.Background(), testPNG(t), ".png", false)
	if err != nil {
		t.Fatalf("ResizePhoto(png) error = %v", err)
	}
	if len(sizes) == 0 {
		t.Fatal("ResizePhoto(png) returned no sizes")
	}
	if _, err := png.Decode(bytes.NewReader(sizes[0].Bytes)); err != nil {
		t.Fatalf("resized png is not png: %v", err)
	}
}

func TestResizePhotoRejectsUnsupportedDecodeFormat(t *testing.T) {
	processor := NewProcessor()
	_, err := processor.ResizePhoto(context.Background(), testJPEG(t), ".tiff", false)
	if !errors.Is(err, imaging.ErrUnsupportedFormat) {
		t.Fatalf("ResizePhoto(tiff) error = %v, want unsupported format", err)
	}
}

func TestResizePhotoRejectsOversizedInput(t *testing.T) {
	processor := NewProcessor()
	oversized := make([]byte, maxImageBytes+1)
	_, err := processor.ResizePhoto(context.Background(), oversized, "jpg", false)
	if !errors.Is(err, ErrImageTooLarge) {
		t.Fatalf("ResizePhoto() error = %v, want ErrImageTooLarge", err)
	}
}

func testJPEG(t *testing.T) []byte {
	t.Helper()
	return encodeFixture(t, func(w io.Writer, img image.Image) error {
		return jpeg.Encode(w, img, nil)
	})
}

func testPNG(t *testing.T) []byte {
	t.Helper()
	return encodeFixture(t, png.Encode)
}

func encodeFixture(t *testing.T, encode func(io.Writer, image.Image) error) []byte {
	t.Helper()
	img := image.NewRGBA(image.Rect(0, 0, 640, 480))
	for y := 0; y < 480; y++ {
		for x := 0; x < 640; x++ {
			img.Set(x, y, color.RGBA{R: uint8(x), G: uint8(y), B: 120, A: 255})
		}
	}
	var buf bytes.Buffer
	if err := encode(&buf, img); err != nil {
		t.Fatalf("encode fixture: %v", err)
	}
	return buf.Bytes()
}
