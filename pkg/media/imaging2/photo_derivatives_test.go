package imaging2

import (
	"bytes"
	"context"
	"errors"
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"testing"
)

func TestBuildPhotoDerivativesDoesNotUpscaleTinyImage(t *testing.T) {
	processor := NewProcessorWithProgressiveEncoder(fakeProgressiveJPEGEncoder{})
	derivatives, err := processor.BuildPhotoDerivatives(context.Background(), testSizedJPEG(t, 16, 14), "jpg", false)
	if err != nil {
		t.Fatalf("BuildPhotoDerivatives() error = %v", err)
	}
	assertPhotoDerivatives(t, derivatives, []wantPhotoDerivative{
		{typ: "i", w: 16, h: 14, stripped: true},
		{typ: "m", w: 16, h: 14},
	})
}

func TestBuildPhotoDerivativesUsesProgressiveXForMediumImage(t *testing.T) {
	processor := NewProcessorWithProgressiveEncoder(fakeProgressiveJPEGEncoder{})
	derivatives, err := processor.BuildPhotoDerivatives(context.Background(), testSizedJPEG(t, 606, 429), "jpg", false)
	if err != nil {
		t.Fatalf("BuildPhotoDerivatives() error = %v", err)
	}
	assertPhotoDerivatives(t, derivatives, []wantPhotoDerivative{
		{typ: "i", w: 40, h: 28, stripped: true},
		{typ: "m", w: 320, h: 227},
		{typ: "x", w: 606, h: 429, progressive: true},
	})
}

func TestBuildPhotoDerivativesUsesProgressiveYForLargeImage(t *testing.T) {
	processor := NewProcessorWithProgressiveEncoder(fakeProgressiveJPEGEncoder{})
	derivatives, err := processor.BuildPhotoDerivatives(context.Background(), testSizedJPEG(t, 1198, 646), "jpg", false)
	if err != nil {
		t.Fatalf("BuildPhotoDerivatives() error = %v", err)
	}
	assertPhotoDerivatives(t, derivatives, []wantPhotoDerivative{
		{typ: "i", w: 40, h: 22, stripped: true},
		{typ: "m", w: 320, h: 173},
		{typ: "x", w: 800, h: 431},
		{typ: "y", w: 1198, h: 646, progressive: true},
	})
}

func TestBuildPhotoDerivativesCapsYAt1280(t *testing.T) {
	processor := NewProcessorWithProgressiveEncoder(fakeProgressiveJPEGEncoder{})
	derivatives, err := processor.BuildPhotoDerivatives(context.Background(), testSizedJPEG(t, 2000, 616), "jpg", false)
	if err != nil {
		t.Fatalf("BuildPhotoDerivatives() error = %v", err)
	}
	assertPhotoDerivatives(t, derivatives, []wantPhotoDerivative{
		{typ: "i", w: 40, h: 12, stripped: true},
		{typ: "m", w: 320, h: 99},
		{typ: "x", w: 800, h: 246},
		{typ: "y", w: 1280, h: 394, progressive: true},
	})
}

func TestBuildPhotoDerivativesOmitsStrippedWhenStrippedEncodingFails(t *testing.T) {
	processor := NewProcessorWithProgressiveEncoder(fakeProgressiveJPEGEncoder{})
	processor.stripped = failingStrippedJPEGEncoder{}
	derivatives, err := processor.BuildPhotoDerivatives(context.Background(), testSizedJPEG(t, 606, 429), "jpg", false)
	if err != nil {
		t.Fatalf("BuildPhotoDerivatives() error = %v", err)
	}
	assertPhotoDerivatives(t, derivatives, []wantPhotoDerivative{
		{typ: "m", w: 320, h: 227},
		{typ: "x", w: 606, h: 429, progressive: true},
	})
}

func TestBuildPhotoDerivativesUsesProgressiveBytesForDimensions(t *testing.T) {
	processor := NewProcessorWithProgressiveEncoder(geometryChangingProgressiveJPEGEncoder{width: 123, height: 45})
	derivatives, err := processor.BuildPhotoDerivatives(context.Background(), testSizedJPEG(t, 606, 429), "jpg", false)
	if err != nil {
		t.Fatalf("BuildPhotoDerivatives() error = %v", err)
	}
	assertPhotoDerivatives(t, derivatives, []wantPhotoDerivative{
		{typ: "i", w: 40, h: 28, stripped: true},
		{typ: "m", w: 320, h: 227},
		{typ: "x", w: 123, h: 45, progressive: true},
	})
}

type wantPhotoDerivative struct {
	typ         string
	w           int32
	h           int32
	stripped    bool
	progressive bool
}

func assertPhotoDerivatives(t *testing.T, got []PhotoDerivativeBytes, want []wantPhotoDerivative) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("len(derivatives) = %d, want %d: %+v", len(got), len(want), got)
	}
	for i := range want {
		if got[i].Type != want[i].typ || got[i].W != want[i].w || got[i].H != want[i].h || got[i].Stripped != want[i].stripped {
			t.Fatalf("derivative[%d] = {type:%q w:%d h:%d stripped:%t}, want {type:%q w:%d h:%d stripped:%t}",
				i, got[i].Type, got[i].W, got[i].H, got[i].Stripped, want[i].typ, want[i].w, want[i].h, want[i].stripped)
		}
		if len(got[i].Bytes) == 0 {
			t.Fatalf("derivative[%d] has empty bytes", i)
		}
		if want[i].progressive {
			if len(got[i].ProgressiveSizes) != 2 {
				t.Fatalf("derivative[%d].ProgressiveSizes = %v, want two sizes", i, got[i].ProgressiveSizes)
			}
			if got[i].ProgressiveSizes[0] != int32(len(got[i].Bytes)/2) || got[i].ProgressiveSizes[1] != int32(len(got[i].Bytes)) {
				t.Fatalf("derivative[%d].ProgressiveSizes = %v, want half/full byte lengths", i, got[i].ProgressiveSizes)
			}
		} else if len(got[i].ProgressiveSizes) != 0 {
			t.Fatalf("derivative[%d].ProgressiveSizes = %v, want empty", i, got[i].ProgressiveSizes)
		}
	}
}

type fakeProgressiveJPEGEncoder struct{}

func (fakeProgressiveJPEGEncoder) EncodeProgressiveJPEG(ctx context.Context, input []byte, ext string, maxSide int) ([]byte, []int32, error) {
	if err := ctx.Err(); err != nil {
		return nil, nil, err
	}
	img, err := decodeImage(input, ext)
	if err != nil {
		return nil, nil, err
	}
	resizeByWidth, originalMaxSide := resizeAxis(img)
	if maxSide < originalMaxSide {
		img = resizeByLongestSide(img, resizeByWidth, maxSide)
	}
	data, err := encodeImage(img, ".jpg")
	if err != nil {
		return nil, nil, err
	}
	return data, []int32{int32(len(data) / 2), int32(len(data))}, nil
}

type failingStrippedJPEGEncoder struct{}

func (failingStrippedJPEGEncoder) EncodeStrippedJPEG(io.Writer, image.Image) error {
	return errors.New("stripped encode failed")
}

type geometryChangingProgressiveJPEGEncoder struct {
	width  int
	height int
}

func (e geometryChangingProgressiveJPEGEncoder) EncodeProgressiveJPEG(ctx context.Context, _ []byte, _ string, _ int) ([]byte, []int32, error) {
	if err := ctx.Err(); err != nil {
		return nil, nil, err
	}
	data := testJPEGBytes(e.width, e.height)
	return data, []int32{int32(len(data) / 2), int32(len(data))}, nil
}

func testSizedJPEG(t *testing.T, width, height int) []byte {
	t.Helper()
	data, err := encodeTestJPEG(width, height)
	if err != nil {
		t.Fatalf("encode jpeg fixture: %v", err)
	}
	return data
}

func testJPEGBytes(width, height int) []byte {
	data, err := encodeTestJPEG(width, height)
	if err != nil {
		panic(err)
	}
	return data
}

func encodeTestJPEG(width, height int) ([]byte, error) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.RGBA{R: uint8(x), G: uint8(y), B: 120, A: 255})
		}
	}
	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, img, nil); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
