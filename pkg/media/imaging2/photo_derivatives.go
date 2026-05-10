package imaging2

import (
	"bytes"
	"context"
	"fmt"
	"image"

	strippedjpeg "github.com/teamgram/teamgram-server/v2/pkg/media/imaging2/jpeg"
)

func (p *ImagingProcessor) BuildPhotoDerivatives(ctx context.Context, original []byte, ext string, isABC bool) ([]PhotoDerivativeBytes, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if isABC {
		sizes, err := p.ResizePhoto(ctx, original, ext, true)
		if err != nil {
			return nil, fmt.Errorf("resize profile photo derivatives: %w", err)
		}
		out := make([]PhotoDerivativeBytes, 0, len(sizes))
		for _, size := range sizes {
			out = append(out, PhotoDerivativeBytes{
				Type:  size.Type,
				W:     size.W,
				H:     size.H,
				Bytes: size.Bytes,
			})
		}
		return out, nil
	}

	img, err := decodeImage(original, ext)
	if err != nil {
		return nil, fmt.Errorf("decode photo derivatives image: %w", err)
	}
	resizeByWidth, originalMaxSide := resizeAxis(img)

	stripped, err := buildStrippedDerivative(img, resizeByWidth, min(originalMaxSide, PhotoSZStrippedSize))
	if err != nil {
		return nil, err
	}

	downloadable := []ResizeInfo{{Type: "m", Size: min(originalMaxSide, PhotoSZMediumSize)}}
	if originalMaxSide > PhotoSZMediumSize {
		downloadable = append(downloadable, ResizeInfo{Type: "x", Size: min(originalMaxSide, PhotoSZXLargeSize)})
	}
	if originalMaxSide > PhotoSZXLargeSize {
		downloadable = append(downloadable, ResizeInfo{Type: "y", Size: min(originalMaxSide, PhotoSZYLargeSize)})
	}

	out := make([]PhotoDerivativeBytes, 0, 1+len(downloadable))
	out = append(out, stripped)
	for i, size := range downloadable {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		useProgressive := len(downloadable) > 1 && i == len(downloadable)-1
		derivative, err := p.buildDownloadableDerivative(ctx, original, ext, img, resizeByWidth, size, useProgressive)
		if err != nil {
			return nil, err
		}
		out = append(out, derivative)
	}
	return out, nil
}

func (p *ImagingProcessor) buildDownloadableDerivative(
	ctx context.Context,
	original []byte,
	ext string,
	img image.Image,
	resizeByWidth bool,
	size ResizeInfo,
	progressive bool,
) (PhotoDerivativeBytes, error) {
	resized := resizeByLongestSide(img, resizeByWidth, size.Size)
	bounds := resized.Bounds()
	derivative := PhotoDerivativeBytes{
		Type: size.Type,
		W:    int32(bounds.Dx()),
		H:    int32(bounds.Dy()),
	}
	if progressive {
		data, progressiveSizes, err := p.progressive.EncodeProgressiveJPEG(ctx, original, ext, size.Size)
		if err != nil {
			return PhotoDerivativeBytes{}, fmt.Errorf("encode progressive %s photo derivative: %w", size.Type, err)
		}
		derivative.Bytes = data
		derivative.ProgressiveSizes = progressiveSizes
		return derivative, nil
	}
	data, err := encodeImage(resized, ".jpg")
	if err != nil {
		return PhotoDerivativeBytes{}, fmt.Errorf("encode %s photo derivative: %w", size.Type, err)
	}
	derivative.Bytes = data
	return derivative, nil
}

func buildStrippedDerivative(img image.Image, resizeByWidth bool, maxSide int) (PhotoDerivativeBytes, error) {
	resized := resizeByLongestSide(img, resizeByWidth, maxSide)
	var buf bytes.Buffer
	if err := strippedjpeg.EncodeStripped(&buf, resized, &strippedjpeg.Options{Quality: 30}); err != nil {
		return PhotoDerivativeBytes{}, fmt.Errorf("encode stripped photo derivative: %w", err)
	}
	bounds := resized.Bounds()
	return PhotoDerivativeBytes{
		Type:     "i",
		W:        int32(bounds.Dx()),
		H:        int32(bounds.Dy()),
		Bytes:    buf.Bytes(),
		Stripped: true,
	}, nil
}
