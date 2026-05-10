package imaging2

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"strings"

	"github.com/disintegration/imaging"
	strippedjpeg "github.com/teamgram/teamgram-server/v2/pkg/media/imaging2/jpeg"
	"golang.org/x/image/bmp"
)

type Processor interface {
	ResizePhoto(ctx context.Context, original []byte, ext string, isABC bool) ([]PhotoSizeBytes, error)
	EncodeStripped(ctx context.Context, image []byte) ([]byte, error)
}

type PhotoSizeBytes struct {
	Type  string
	W     int32
	H     int32
	Bytes []byte
}

type ResizeInfo struct {
	Type string
	Size int
}

const (
	PhotoSZStrippedSize = 40
	PhotoSZSmallSize    = 90
	PhotoSZMediumSize   = 320
	PhotoSZXLargeSize   = 800
	PhotoSZYLargeSize   = 1280
	PhotoSZWLargeSize   = 2560
	PhotoSZASize        = 160
	PhotoSZBSize        = 320
	PhotoSZCSize        = 640
	PhotoSZDSize        = 1280

	maxImageBytes  = 32 << 20
	maxImagePixels = 64_000_000
)

var ErrImageTooLarge = errors.New("image too large")

var (
	PhotoSizes = []ResizeInfo{
		{Type: "s", Size: PhotoSZSmallSize},
		{Type: "m", Size: PhotoSZMediumSize},
		{Type: "x", Size: PhotoSZXLargeSize},
		{Type: "y", Size: PhotoSZYLargeSize},
		{Type: "w", Size: PhotoSZWLargeSize},
	}
	ProfilePhotoSizes = []ResizeInfo{
		{Type: "a", Size: PhotoSZASize},
		{Type: "b", Size: PhotoSZBSize},
		{Type: "c", Size: PhotoSZCSize},
		{Type: "d", Size: PhotoSZDSize},
	}
)

type ImagingProcessor struct{}

func NewProcessor() *ImagingProcessor {
	return &ImagingProcessor{}
}

func (p *ImagingProcessor) ResizePhoto(ctx context.Context, original []byte, ext string, isABC bool) ([]PhotoSizeBytes, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	img, err := decodeImage(original, ext)
	if err != nil {
		return nil, fmt.Errorf("decode image: %w", err)
	}
	if isABC {
		img = normalizeProfileImage(img)
	}
	sizeList := PhotoSizes
	if isABC {
		sizeList = ProfilePhotoSizes
	}
	resizeByWidth, maxSide := resizeAxis(img)
	out := make([]PhotoSizeBytes, 0, len(sizeList))
	for _, size := range sizeList {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		target := size.Size
		willBreak := false
		if target >= maxSide {
			target = maxSide
			willBreak = true
		}
		resized := resizeByLongestSide(img, resizeByWidth, target)
		data, err := encodeImage(resized, ext)
		if err != nil {
			return nil, fmt.Errorf("encode %s: %w", size.Type, err)
		}
		bounds := resized.Bounds()
		out = append(out, PhotoSizeBytes{
			Type:  size.Type,
			W:     int32(bounds.Dx()),
			H:     int32(bounds.Dy()),
			Bytes: data,
		})
		if willBreak {
			break
		}
	}
	return out, nil
}

func (p *ImagingProcessor) EncodeStripped(ctx context.Context, data []byte) ([]byte, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	img, err := decodeImageAuto(data)
	if err != nil {
		return nil, fmt.Errorf("decode stripped image: %w", err)
	}
	resizeByWidth, maxSide := resizeAxis(img)
	if maxSide > PhotoSZStrippedSize {
		img = resizeByLongestSide(img, resizeByWidth, PhotoSZStrippedSize)
	}
	var buf bytes.Buffer
	if err := strippedjpeg.EncodeStripped(&buf, img, &strippedjpeg.Options{Quality: 30}); err != nil {
		return nil, fmt.Errorf("encode stripped photo: %w", err)
	}
	return buf.Bytes(), nil
}

func validateImageBytes(data []byte, decodeConfig func(io.Reader) (image.Config, error)) error {
	if len(data) == 0 {
		return fmt.Errorf("empty image")
	}
	if len(data) > maxImageBytes {
		return ErrImageTooLarge
	}
	cfg, err := decodeConfig(bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("decode image config: %w", err)
	}
	if cfg.Width <= 0 || cfg.Height <= 0 {
		return fmt.Errorf("invalid image dimensions")
	}
	if cfg.Width > maxImagePixels/cfg.Height {
		return ErrImageTooLarge
	}
	return nil
}

func decodeImage(data []byte, ext string) (image.Image, error) {
	decode, decodeConfig, err := imageDecoders(ext)
	if err != nil {
		return nil, err
	}
	if err := validateImageBytes(data, decodeConfig); err != nil {
		return nil, err
	}
	return decode(bytes.NewReader(data))
}

func decodeImageAuto(data []byte) (image.Image, error) {
	for _, ext := range []string{".jpg", ".png", ".gif", ".bmp"} {
		img, err := decodeImage(data, ext)
		if err == nil {
			return img, nil
		}
	}
	return nil, imaging.ErrUnsupportedFormat
}

func imageDecoders(ext string) (func(io.Reader) (image.Image, error), func(io.Reader) (image.Config, error), error) {
	switch normalizeExt(ext) {
	case ".jpg", ".jpeg":
		return jpeg.Decode, jpeg.DecodeConfig, nil
	case ".png":
		return png.Decode, png.DecodeConfig, nil
	case ".gif":
		return gif.Decode, gif.DecodeConfig, nil
	case ".bmp":
		return bmp.Decode, bmp.DecodeConfig, nil
	default:
		return nil, nil, imaging.ErrUnsupportedFormat
	}
}

func normalizeProfileImage(img image.Image) image.Image {
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()
	if w >= PhotoSZDSize && h >= PhotoSZDSize {
		if w != h {
			return imaging.Fill(img, PhotoSZDSize, PhotoSZDSize, imaging.Center, imaging.Lanczos)
		}
		return img
	}
	if w <= PhotoSZCSize && h <= PhotoSZCSize {
		return imaging.Fill(img, PhotoSZCSize, PhotoSZCSize, imaging.Center, imaging.Lanczos)
	}
	if w != h {
		return imaging.Fill(img, PhotoSZCSize, PhotoSZCSize, imaging.Center, imaging.Lanczos)
	}
	return img
}

func resizeAxis(img image.Image) (byWidth bool, maxSide int) {
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()
	if w >= h {
		return true, w
	}
	return false, h
}

func resizeByLongestSide(img image.Image, byWidth bool, size int) *image.NRGBA {
	if byWidth {
		return imaging.Resize(img, size, 0, imaging.Lanczos)
	}
	return imaging.Resize(img, 0, size, imaging.Lanczos)
}

func encodeImage(img image.Image, ext string) ([]byte, error) {
	var buf bytes.Buffer
	format, err := imageFormat(ext)
	if err != nil {
		return nil, err
	}
	if err := imaging.Encode(&buf, img, format); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func imageFormat(ext string) (imaging.Format, error) {
	formats := map[string]imaging.Format{
		".jpg":  imaging.JPEG,
		".jpeg": imaging.JPEG,
		".png":  imaging.PNG,
		".gif":  imaging.GIF,
		".tif":  imaging.TIFF,
		".tiff": imaging.TIFF,
		".bmp":  imaging.BMP,
	}
	normalized := normalizeExt(ext)
	format, ok := formats[normalized]
	if !ok {
		return -1, imaging.ErrUnsupportedFormat
	}
	return format, nil
}

func normalizeExt(ext string) string {
	normalized := strings.ToLower(ext)
	if !strings.HasPrefix(normalized, ".") {
		normalized = "." + normalized
	}
	return normalized
}
