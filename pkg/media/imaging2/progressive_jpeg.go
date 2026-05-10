package imaging2

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

const (
	defaultImageMagickTimeout = 30 * time.Second
	defaultProgressiveQuality = 85

	jpegSOI  = 0xd8
	jpegEOI  = 0xd9
	jpegSOS  = 0xda
	jpegSOF2 = 0xc2
)

type ImageMagickProgressiveEncoder struct {
	Binary  string
	Timeout time.Duration
	Quality int
}

func ResolveImageMagickBinary(configured string) (string, error) {
	if configured != "" {
		binary, err := exec.LookPath(configured)
		if err != nil {
			return "", fmt.Errorf("resolve configured ImageMagick binary %q: %w", configured, err)
		}
		base := filepath.Base(binary)
		if base != "magick" && base != "convert" {
			return "", fmt.Errorf("unsupported ImageMagick binary %q", base)
		}
		return binary, nil
	}

	for _, candidate := range []string{"magick", "convert"} {
		binary, err := exec.LookPath(candidate)
		if err == nil {
			return binary, nil
		}
	}
	return "", fmt.Errorf("resolve ImageMagick binary: %w", exec.ErrNotFound)
}

func (e ImageMagickProgressiveEncoder) EncodeProgressiveJPEG(ctx context.Context, input []byte, ext string, maxSide int) ([]byte, []int32, error) {
	if err := ctx.Err(); err != nil {
		return nil, nil, err
	}
	binary, err := ResolveImageMagickBinary(e.Binary)
	if err != nil {
		return nil, nil, err
	}
	timeout := e.Timeout
	if timeout <= 0 {
		timeout = defaultImageMagickTimeout
	}
	quality := e.Quality
	if quality <= 0 {
		quality = defaultProgressiveQuality
	}

	runCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	cmd := exec.CommandContext(runCtx, binary, imageMagickArgs(maxSide, quality)...)
	cmd.Stdin = bytes.NewReader(input)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		if runCtx.Err() != nil {
			return nil, nil, fmt.Errorf("encode progressive jpeg with ImageMagick timed out: %w", runCtx.Err())
		}
		if stderr.Len() > 0 {
			return nil, nil, fmt.Errorf("encode progressive jpeg with ImageMagick: %w: %s", err, stderr.String())
		}
		return nil, nil, fmt.Errorf("encode progressive jpeg with ImageMagick: %w", err)
	}
	data := stdout.Bytes()
	scanSizes, err := ProgressiveScanSizes(data)
	if err != nil {
		return nil, nil, fmt.Errorf("parse progressive jpeg scan sizes: %w", err)
	}
	return data, scanSizes, nil
}

func imageMagickArgs(maxSide int, quality int) []string {
	resizeArg := strconv.Itoa(maxSide) + "x" + strconv.Itoa(maxSide) + ">"
	return []string{
		"-",
		"-auto-orient",
		"-resize", resizeArg,
		"-strip",
		"-interlace", "Plane",
		"-quality", strconv.Itoa(quality),
		"jpg:-",
	}
}

func ProgressiveScanSizes(data []byte) ([]int32, error) {
	if len(data) < 2 || data[0] != 0xff || data[1] != jpegSOI {
		return nil, fmt.Errorf("missing jpeg SOI")
	}

	var scanSizes []int32
	progressive := false
	nonProgressiveSOF := false
	for pos := 2; pos < len(data); {
		markerPos, marker, err := nextMarker(data, pos)
		if err != nil {
			return nil, err
		}
		pos = markerPos + 2
		switch {
		case marker == jpegEOI:
			if !progressive {
				if nonProgressiveSOF {
					return nil, fmt.Errorf("jpeg is not progressive")
				}
				return nil, fmt.Errorf("missing progressive SOF marker")
			}
			if len(scanSizes) == 0 {
				return nil, fmt.Errorf("progressive jpeg has no scans")
			}
			if scanSizes[len(scanSizes)-1] != int32(len(data)) {
				scanSizes = appendIncreasingScanSize(scanSizes, len(data))
			}
			return scanSizes, nil
		case marker == jpegSOF2:
			progressive = true
		case isSOFMarker(marker):
			nonProgressiveSOF = true
		case marker == jpegSOS:
			if pos+2 > len(data) {
				return nil, fmt.Errorf("truncated jpeg SOS length")
			}
			segmentLen := int(data[pos])<<8 | int(data[pos+1])
			if segmentLen < 2 {
				return nil, fmt.Errorf("invalid jpeg SOS length %d", segmentLen)
			}
			scanStart := pos + segmentLen
			if scanStart > len(data) {
				return nil, fmt.Errorf("truncated jpeg SOS segment")
			}
			scanEnd, err := nextScanBoundary(data, scanStart)
			if err != nil {
				return nil, err
			}
			if progressive {
				scanSizes = appendIncreasingScanSize(scanSizes, scanEnd)
			}
			pos = scanEnd
			continue
		}

		if markerHasPayload(marker) {
			if pos+2 > len(data) {
				return nil, fmt.Errorf("truncated jpeg marker 0x%02x length", marker)
			}
			segmentLen := int(data[pos])<<8 | int(data[pos+1])
			if segmentLen < 2 {
				return nil, fmt.Errorf("invalid jpeg marker 0x%02x length %d", marker, segmentLen)
			}
			pos += segmentLen
			if pos > len(data) {
				return nil, fmt.Errorf("truncated jpeg marker 0x%02x segment", marker)
			}
		}
	}

	return nil, fmt.Errorf("missing jpeg EOI")
}

func nextMarker(data []byte, pos int) (int, byte, error) {
	for pos < len(data) && data[pos] != 0xff {
		pos++
	}
	if pos >= len(data) {
		return 0, 0, fmt.Errorf("missing jpeg marker")
	}
	for pos < len(data) && data[pos] == 0xff {
		pos++
	}
	if pos >= len(data) {
		return 0, 0, fmt.Errorf("truncated jpeg marker")
	}
	return pos - 1, data[pos], nil
}

func nextScanBoundary(data []byte, pos int) (int, error) {
	for pos < len(data) {
		if data[pos] != 0xff {
			pos++
			continue
		}
		if pos+1 >= len(data) {
			return 0, fmt.Errorf("truncated jpeg scan data")
		}
		next := data[pos+1]
		switch {
		case next == 0x00:
			pos += 2
		case next == 0xff:
			pos++
		case next >= 0xd0 && next <= 0xd7:
			pos += 2
		default:
			return pos, nil
		}
	}
	return 0, fmt.Errorf("missing jpeg scan boundary")
}

func appendIncreasingScanSize(scanSizes []int32, size int) []int32 {
	if size <= 0 {
		return scanSizes
	}
	if len(scanSizes) == 0 || int32(size) > scanSizes[len(scanSizes)-1] {
		return append(scanSizes, int32(size))
	}
	return scanSizes
}

func markerHasPayload(marker byte) bool {
	if marker == 0x01 {
		return false
	}
	if marker >= 0xd0 && marker <= 0xd9 {
		return false
	}
	return true
}

func isSOFMarker(marker byte) bool {
	switch marker {
	case 0xc0, 0xc1, 0xc3, 0xc5, 0xc6, 0xc7, 0xc9, 0xca, 0xcb, 0xcd, 0xce, 0xcf:
		return true
	default:
		return false
	}
}
