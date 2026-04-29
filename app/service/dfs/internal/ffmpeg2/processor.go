package ffmpeg2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"os/exec"
	"strconv"
)

type Processor interface {
	ExtractFirstFrame(ctx context.Context, input io.Reader) ([]byte, error)
	ConvertToMP4(ctx context.Context, input io.Reader) ([]byte, error)
	GetVideoMetadata(ctx context.Context, input io.Reader) (*VideoMetadata, error)
}

type VideoMetadata struct {
	Width    int32
	Height   int32
	Duration int32
}

var ErrProcessorUnavailable = errors.New("ffmpeg processor unavailable")

type ExternalProcessor struct{}

func NewProcessor() *ExternalProcessor {
	return &ExternalProcessor{}
}

func (p *ExternalProcessor) ExtractFirstFrame(ctx context.Context, input io.Reader) ([]byte, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if input == nil {
		return nil, ErrProcessorUnavailable
	}
	in, cleanup, err := writeTempInput(input, "teamgram-ffmpeg-*")
	if err != nil {
		return nil, err
	}
	defer cleanup()
	out, err := tempOutput("teamgram-frame-*.jpg")
	if err != nil {
		return nil, err
	}
	defer os.Remove(out)
	if err := run(ctx, "ffmpeg", "-y", "-i", in, "-frames:v", "1", out); err != nil {
		return nil, err
	}
	return os.ReadFile(out)
}

func (p *ExternalProcessor) ConvertToMP4(ctx context.Context, input io.Reader) ([]byte, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if input == nil {
		return nil, ErrProcessorUnavailable
	}
	in, cleanup, err := writeTempInput(input, "teamgram-ffmpeg-*")
	if err != nil {
		return nil, err
	}
	defer cleanup()
	out, err := tempOutput("teamgram-video-*.mp4")
	if err != nil {
		return nil, err
	}
	defer os.Remove(out)
	if err := run(ctx, "ffmpeg", "-y", "-i", in, "-movflags", "faststart", out); err != nil {
		return nil, err
	}
	return os.ReadFile(out)
}

func (p *ExternalProcessor) GetVideoMetadata(ctx context.Context, input io.Reader) (*VideoMetadata, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if input == nil {
		return nil, ErrProcessorUnavailable
	}
	in, cleanup, err := writeTempInput(input, "teamgram-ffprobe-*")
	if err != nil {
		return nil, err
	}
	defer cleanup()
	out, err := exec.CommandContext(ctx, "ffprobe", "-v", "error", "-select_streams", "v:0", "-show_entries", "stream=width,height,duration", "-show_entries", "format=duration", "-of", "json", in).Output()
	if err != nil {
		return nil, fmt.Errorf("%w: ffprobe metadata: %w", ErrProcessorUnavailable, err)
	}
	var probed struct {
		Streams []struct {
			Width    int32  `json:"width"`
			Height   int32  `json:"height"`
			Duration string `json:"duration"`
		} `json:"streams"`
		Format struct {
			Duration string `json:"duration"`
		} `json:"format"`
	}
	if err := json.Unmarshal(out, &probed); err != nil {
		return nil, fmt.Errorf("ffprobe metadata decode: %w", err)
	}
	if len(probed.Streams) == 0 {
		return nil, fmt.Errorf("%w: no video stream", ErrProcessorUnavailable)
	}
	duration := parseDurationSeconds(probed.Streams[0].Duration)
	if duration == 0 {
		duration = parseDurationSeconds(probed.Format.Duration)
	}
	return &VideoMetadata{
		Width:    probed.Streams[0].Width,
		Height:   probed.Streams[0].Height,
		Duration: duration,
	}, nil
}

func writeTempInput(input io.Reader, pattern string) (string, func(), error) {
	f, err := os.CreateTemp("", pattern)
	if err != nil {
		return "", nil, fmt.Errorf("create temp input: %w", err)
	}
	cleanup := func() {
		os.Remove(f.Name())
	}
	if _, err := io.Copy(f, input); err != nil {
		f.Close()
		cleanup()
		return "", nil, fmt.Errorf("write temp input: %w", err)
	}
	if err := f.Close(); err != nil {
		cleanup()
		return "", nil, fmt.Errorf("close temp input: %w", err)
	}
	return f.Name(), cleanup, nil
}

func tempOutput(pattern string) (string, error) {
	f, err := os.CreateTemp("", pattern)
	if err != nil {
		return "", fmt.Errorf("create temp output: %w", err)
	}
	name := f.Name()
	if err := f.Close(); err != nil {
		os.Remove(name)
		return "", fmt.Errorf("close temp output: %w", err)
	}
	return name, nil
}

func run(ctx context.Context, name string, args ...string) error {
	if _, err := exec.LookPath(name); err != nil {
		return fmt.Errorf("%w: %s not found: %w", ErrProcessorUnavailable, name, err)
	}
	out, err := exec.CommandContext(ctx, name, args...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("%w: %s failed: %w: %s", ErrProcessorUnavailable, name, err, out)
	}
	return nil
}

func parseDurationSeconds(raw string) int32 {
	if raw == "" || raw == "N/A" {
		return 0
	}
	v, err := strconv.ParseFloat(raw, 64)
	if err != nil || v <= 0 {
		return 0
	}
	return int32(math.Round(v))
}
