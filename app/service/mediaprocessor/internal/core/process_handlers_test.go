package core

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/app/service/mediaprocessor/internal/processor"
	"github.com/teamgram/teamgram-server/v2/app/service/mediaprocessor/internal/svc"
	"github.com/teamgram/teamgram-server/v2/app/service/mediaprocessor/mediaprocessor"
	"github.com/teamgram/teamgram-server/v2/pkg/media/ffmpeg2"
	"github.com/teamgram/teamgram-server/v2/pkg/media/imaging2"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

var (
	errRead    = errors.New("read failed")
	errPut     = errors.New("put failed")
	errConvert = errors.New("convert failed")
	errProbe   = errors.New("probe failed")
	errCover   = errors.New("cover failed")
)

func TestProcessHandlersRejectInvalidArgument(t *testing.T) {
	c := New(context.Background(), &svc.ServiceContext{})

	tests := []struct {
		name string
		call func() error
	}{
		{
			name: "photo",
			call: func() error {
				_, err := c.MediaProcessorProcessPhoto(&mediaprocessor.TLMediaProcessorProcessPhoto{})
				return err
			},
		},
		{
			name: "gif",
			call: func() error {
				_, err := c.MediaProcessorProcessGif(&mediaprocessor.TLMediaProcessorProcessGif{})
				return err
			},
		},
		{
			name: "mp4",
			call: func() error {
				_, err := c.MediaProcessorProcessMp4(&mediaprocessor.TLMediaProcessorProcessMp4{})
				return err
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.call(); !errors.Is(err, mediaprocessor.ErrMediaProcessorInvalidArgument) {
				t.Fatalf("got err %v, want ErrMediaProcessorInvalidArgument", err)
			}
		})
	}
}

func TestProcessPhotoReadsResizesAndWritesDerivatives(t *testing.T) {
	dfsFake := &fakeDFS{readBytes: []byte("original")}
	procFake := &fakeProcessor{
		resized: []imaging2.PhotoSizeBytes{
			{Type: "s", W: 90, H: 60, Bytes: []byte("small")},
			{Type: "m", W: 320, H: 240, Bytes: []byte("medium")},
		},
	}
	c := New(context.Background(), &svc.ServiceContext{DfsClient: dfsFake, Processor: procFake})

	out, err := c.MediaProcessorProcessPhoto(&mediaprocessor.TLMediaProcessorProcessPhoto{
		OwnerId:   1001,
		ObjectId:  "original-photo",
		ReadLease: []byte("lease"),
		FileName:  "image.jpg",
		Profile:   tg.BoolTrueClazz,
	})
	if err != nil {
		t.Fatalf("MediaProcessorProcessPhoto() error = %v", err)
	}

	if got, want := string(dfsFake.readLease), "lease"; got != want {
		t.Fatalf("read lease = %q, want %q", got, want)
	}
	if string(procFake.resizeInput) != "original" || procFake.resizeExt != ".jpg" || !procFake.resizeProfile {
		t.Fatalf("resize call input=%q ext=%q profile=%t", procFake.resizeInput, procFake.resizeExt, procFake.resizeProfile)
	}
	if len(dfsFake.puts) != 2 {
		t.Fatalf("puts len = %d, want 2", len(dfsFake.puts))
	}
	assertPut(t, dfsFake.puts[0], "media_derivative", "s_image.jpg", "image/jpeg", []byte("small"))
	assertPut(t, dfsFake.puts[1], "media_derivative", "m_image.jpg", "image/jpeg", []byte("medium"))
	if out.OriginalObjectId != "original-photo" || len(out.Sizes) != 2 {
		t.Fatalf("out = %+v", out)
	}
	assertDerivative(t, out.Sizes[0], processor.DerivativePhotoSize, "put-1", "s_image.jpg", "image/jpeg", 1005, 90, 60, []byte("small"))
	assertDerivative(t, out.Sizes[1], processor.DerivativePhotoSize, "put-2", "m_image.jpg", "image/jpeg", 1006, 320, 240, []byte("medium"))
}

func TestProcessGifRejectsShortInputBeforeConvert(t *testing.T) {
	dfsFake := &fakeDFS{readBytes: make([]byte, 10239)}
	procFake := &fakeProcessor{}
	c := New(context.Background(), &svc.ServiceContext{DfsClient: dfsFake, Processor: procFake})

	_, err := c.MediaProcessorProcessGif(validGifRequest())
	if !errors.Is(err, mediaprocessor.ErrMediaProcessorInvalidArgument) {
		t.Fatalf("MediaProcessorProcessGif() error = %v, want ErrMediaProcessorInvalidArgument", err)
	}
	if procFake.convertCalled {
		t.Fatalf("ConvertGIFToMP4 was called for short input")
	}
}

func TestProcessGifConvertProbeAndWriteCoverFailureIsNonFatal(t *testing.T) {
	dfsFake := &fakeDFS{readBytes: make([]byte, 10240)}
	procFake := &fakeProcessor{
		mp4:      []byte("converted-mp4"),
		metadata: &ffmpeg2.VideoMetadata{Width: 640, Height: 360, Duration: 7},
		coverErr: errCover,
	}
	c := New(context.Background(), &svc.ServiceContext{DfsClient: dfsFake, Processor: procFake})

	out, err := c.MediaProcessorProcessGif(validGifRequest())
	if err != nil {
		t.Fatalf("MediaProcessorProcessGif() error = %v", err)
	}

	if len(dfsFake.puts) != 1 {
		t.Fatalf("puts len = %d, want 1", len(dfsFake.puts))
	}
	if !procFake.coverCalled {
		t.Fatalf("ExtractMP4Cover was not called")
	}
	assertPut(t, dfsFake.puts[0], "media_derivative", "anim.mp4", "video/mp4", []byte("converted-mp4"))
	if out.ObjectId != "put-1" || out.MimeType != "video/mp4" || out.Size2 != 1013 || len(out.Thumbs) != 0 {
		t.Fatalf("out = %+v", out)
	}
	assertAttributes(t, out.Attributes, true, "anim.mp4")
}

func TestProcessGifUsesSuppliedThumbReadLease(t *testing.T) {
	dfsFake := &fakeDFS{
		readByLease: map[string][]byte{
			"lease":       make([]byte, 10240),
			"thumb-lease": []byte("supplied-thumb"),
		},
	}
	procFake := &fakeProcessor{
		mp4:      []byte("converted-mp4"),
		metadata: &ffmpeg2.VideoMetadata{Width: 640, Height: 360, Duration: 7},
		coverErr: errCover,
	}
	c := New(context.Background(), &svc.ServiceContext{DfsClient: dfsFake, Processor: procFake})

	req := validGifRequest()
	req.ThumbObjectId = "supplied-thumb-object"
	req.ThumbReadLease = []byte("thumb-lease")
	out, err := c.MediaProcessorProcessGif(req)
	if err != nil {
		t.Fatalf("MediaProcessorProcessGif() error = %v", err)
	}

	if len(dfsFake.readLeases) != 2 || string(dfsFake.readLeases[0]) != "lease" || string(dfsFake.readLeases[1]) != "thumb-lease" {
		t.Fatalf("read leases = %q, want [lease thumb-lease]", dfsFake.readLeases)
	}
	if procFake.coverCalled {
		t.Fatalf("ExtractMP4Cover was called for supplied thumb")
	}
	if len(dfsFake.puts) != 2 {
		t.Fatalf("puts len = %d, want 2", len(dfsFake.puts))
	}
	assertPut(t, dfsFake.puts[0], "media_derivative", "anim.mp4", "video/mp4", []byte("converted-mp4"))
	assertPut(t, dfsFake.puts[1], "media_derivative", "anim_thumb.jpg", "image/jpeg", []byte("supplied-thumb"))
	if out.ObjectId != "put-1" || len(out.Thumbs) != 1 {
		t.Fatalf("out = %+v", out)
	}
	assertDerivative(t, out.Thumbs[0], processor.DerivativeDocumentThumb, "put-2", "anim_thumb.jpg", "image/jpeg", 1014, 640, 360, []byte("supplied-thumb"))
}

func TestProcessMp4ProbeCoverFailureIsNonFatal(t *testing.T) {
	original := []byte("original-mp4")
	dfsFake := &fakeDFS{readBytes: original}
	procFake := &fakeProcessor{
		metadata: &ffmpeg2.VideoMetadata{Width: 1920, Height: 1080, Duration: 12},
		coverErr: errCover,
	}
	c := New(context.Background(), &svc.ServiceContext{DfsClient: dfsFake, Processor: procFake})

	out, err := c.MediaProcessorProcessMp4(validMp4Request())
	if err != nil {
		t.Fatalf("MediaProcessorProcessMp4() error = %v", err)
	}

	if len(dfsFake.puts) != 0 {
		t.Fatalf("puts len = %d, want 0", len(dfsFake.puts))
	}
	if !procFake.coverCalled {
		t.Fatalf("ExtractMP4Cover was not called")
	}
	if out.ObjectId != "original-mp4" || out.MimeType != "video/mp4" || out.Size2 != int64(len(original)) || len(out.Thumbs) != 0 {
		t.Fatalf("out = %+v", out)
	}
	assertAttributes(t, out.Attributes, false, "video.mp4")
}

func TestProcessHandlersPropagateDFSReadAndPutErrors(t *testing.T) {
	t.Run("read", func(t *testing.T) {
		c := New(context.Background(), &svc.ServiceContext{
			DfsClient: &fakeDFS{readErr: errRead},
			Processor: &fakeProcessor{
				resized: []imaging2.PhotoSizeBytes{{Type: "s", W: 1, H: 1, Bytes: []byte("x")}},
			},
		})
		_, err := c.MediaProcessorProcessPhoto(validPhotoRequest())
		if !errors.Is(err, errRead) {
			t.Fatalf("MediaProcessorProcessPhoto() error = %v, want read error", err)
		}
	})

	t.Run("put", func(t *testing.T) {
		c := New(context.Background(), &svc.ServiceContext{
			DfsClient: &fakeDFS{readBytes: []byte("original"), putErr: errPut},
			Processor: &fakeProcessor{
				resized: []imaging2.PhotoSizeBytes{{Type: "s", W: 1, H: 1, Bytes: []byte("x")}},
			},
		})
		_, err := c.MediaProcessorProcessPhoto(validPhotoRequest())
		if !errors.Is(err, errPut) {
			t.Fatalf("MediaProcessorProcessPhoto() error = %v, want put error", err)
		}
	})
}

func TestProcessGifSuppliedThumbErrorsAreFatal(t *testing.T) {
	t.Run("read", func(t *testing.T) {
		c := New(context.Background(), &svc.ServiceContext{
			DfsClient: &fakeDFS{
				readByLease: map[string][]byte{"lease": make([]byte, 10240)},
				readErrByLease: map[string]error{
					"thumb-lease": errRead,
				},
			},
			Processor: &fakeProcessor{
				mp4:      []byte("converted-mp4"),
				metadata: &ffmpeg2.VideoMetadata{Width: 640, Height: 360, Duration: 7},
			},
		})
		req := validGifRequest()
		req.ThumbObjectId = "supplied-thumb-object"
		req.ThumbReadLease = []byte("thumb-lease")
		_, err := c.MediaProcessorProcessGif(req)
		if !errors.Is(err, errRead) {
			t.Fatalf("MediaProcessorProcessGif() error = %v, want read error", err)
		}
	})

	t.Run("put", func(t *testing.T) {
		c := New(context.Background(), &svc.ServiceContext{
			DfsClient: &fakeDFS{
				readByLease: map[string][]byte{
					"lease":       make([]byte, 10240),
					"thumb-lease": []byte("supplied-thumb"),
				},
				putErrAt: 2,
				putErr:   errPut,
			},
			Processor: &fakeProcessor{
				mp4:      []byte("converted-mp4"),
				metadata: &ffmpeg2.VideoMetadata{Width: 640, Height: 360, Duration: 7},
			},
		})
		req := validGifRequest()
		req.ThumbObjectId = "supplied-thumb-object"
		req.ThumbReadLease = []byte("thumb-lease")
		_, err := c.MediaProcessorProcessGif(req)
		if !errors.Is(err, errPut) {
			t.Fatalf("MediaProcessorProcessGif() error = %v, want put error", err)
		}
	})
}

func TestProcessHandlersPropagateProcessorErrors(t *testing.T) {
	t.Run("convert", func(t *testing.T) {
		c := New(context.Background(), &svc.ServiceContext{
			DfsClient: &fakeDFS{readBytes: make([]byte, 10240)},
			Processor: &fakeProcessor{convertErr: errConvert},
		})
		_, err := c.MediaProcessorProcessGif(validGifRequest())
		if !errors.Is(err, errConvert) {
			t.Fatalf("MediaProcessorProcessGif() error = %v, want convert error", err)
		}
	})

	t.Run("probe", func(t *testing.T) {
		c := New(context.Background(), &svc.ServiceContext{
			DfsClient: &fakeDFS{readBytes: []byte("mp4")},
			Processor: &fakeProcessor{probeErr: errProbe},
		})
		_, err := c.MediaProcessorProcessMp4(validMp4Request())
		if !errors.Is(err, errProbe) {
			t.Fatalf("MediaProcessorProcessMp4() error = %v, want probe error", err)
		}
	})
}

func TestProcessHandlersRejectMalformedDerivativeStores(t *testing.T) {
	tests := []struct {
		name string
		call func(*fakeDFS) error
	}{
		{
			name: "photo size nil result",
			call: func(dfsFake *fakeDFS) error {
				dfsFake.readBytes = []byte("original")
				dfsFake.putResults = []*dfs.FileFinalizedObject{nil}
				c := New(context.Background(), &svc.ServiceContext{
					DfsClient: dfsFake,
					Processor: &fakeProcessor{
						resized: []imaging2.PhotoSizeBytes{{Type: "s", W: 1, H: 1, Bytes: []byte("x")}},
					},
				})
				_, err := c.MediaProcessorProcessPhoto(validPhotoRequest())
				return err
			},
		},
		{
			name: "photo size empty object id",
			call: func(dfsFake *fakeDFS) error {
				dfsFake.readBytes = []byte("original")
				dfsFake.putResults = []*dfs.FileFinalizedObject{
					dfs.MakeTLFileFinalizedObject(&dfs.TLFileFinalizedObject{Size2: 1}).ToFileFinalizedObject(),
				}
				c := New(context.Background(), &svc.ServiceContext{
					DfsClient: dfsFake,
					Processor: &fakeProcessor{
						resized: []imaging2.PhotoSizeBytes{{Type: "s", W: 1, H: 1, Bytes: []byte("x")}},
					},
				})
				_, err := c.MediaProcessorProcessPhoto(validPhotoRequest())
				return err
			},
		},
		{
			name: "gif converted mp4 empty object id",
			call: func(dfsFake *fakeDFS) error {
				dfsFake.readBytes = make([]byte, 10240)
				dfsFake.putResults = []*dfs.FileFinalizedObject{
					dfs.MakeTLFileFinalizedObject(&dfs.TLFileFinalizedObject{}).ToFileFinalizedObject(),
				}
				c := New(context.Background(), &svc.ServiceContext{
					DfsClient: dfsFake,
					Processor: &fakeProcessor{
						mp4:      []byte("converted-mp4"),
						metadata: &ffmpeg2.VideoMetadata{Width: 640, Height: 360, Duration: 7},
					},
				})
				_, err := c.MediaProcessorProcessGif(validGifRequest())
				return err
			},
		},
		{
			name: "gif extracted thumb nil result",
			call: func(dfsFake *fakeDFS) error {
				dfsFake.readBytes = make([]byte, 10240)
				dfsFake.putResults = []*dfs.FileFinalizedObject{
					dfs.MakeTLFileFinalizedObject(&dfs.TLFileFinalizedObject{ObjectId: "mp4-object"}).ToFileFinalizedObject(),
					nil,
				}
				c := New(context.Background(), &svc.ServiceContext{
					DfsClient: dfsFake,
					Processor: &fakeProcessor{
						mp4:      []byte("converted-mp4"),
						metadata: &ffmpeg2.VideoMetadata{Width: 640, Height: 360, Duration: 7},
						cover:    []byte("cover"),
					},
				})
				_, err := c.MediaProcessorProcessGif(validGifRequest())
				return err
			},
		},
		{
			name: "mp4 thumb empty object id",
			call: func(dfsFake *fakeDFS) error {
				dfsFake.readBytes = []byte("original-mp4")
				dfsFake.putResults = []*dfs.FileFinalizedObject{
					dfs.MakeTLFileFinalizedObject(&dfs.TLFileFinalizedObject{}).ToFileFinalizedObject(),
				}
				c := New(context.Background(), &svc.ServiceContext{
					DfsClient: dfsFake,
					Processor: &fakeProcessor{
						metadata: &ffmpeg2.VideoMetadata{Width: 1920, Height: 1080, Duration: 12},
						cover:    []byte("cover"),
					},
				})
				_, err := c.MediaProcessorProcessMp4(validMp4Request())
				return err
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.call(&fakeDFS{}); !errors.Is(err, mediaprocessor.ErrMediaProcessorInvalidArgument) {
				t.Fatalf("got err %v, want ErrMediaProcessorInvalidArgument", err)
			}
		})
	}
}

func validPhotoRequest() *mediaprocessor.TLMediaProcessorProcessPhoto {
	return &mediaprocessor.TLMediaProcessorProcessPhoto{
		OwnerId:   1001,
		ObjectId:  "original-photo",
		ReadLease: []byte("lease"),
		FileName:  "image.jpg",
		Profile:   tg.BoolFalseClazz,
	}
}

func validGifRequest() *mediaprocessor.TLMediaProcessorProcessGif {
	return &mediaprocessor.TLMediaProcessorProcessGif{
		OwnerId:   1001,
		ObjectId:  "original-gif",
		ReadLease: []byte("lease"),
		FileName:  "anim.gif",
	}
}

func validMp4Request() *mediaprocessor.TLMediaProcessorProcessMp4 {
	return &mediaprocessor.TLMediaProcessorProcessMp4{
		OwnerId:    1001,
		ObjectId:   "original-mp4",
		ReadLease:  []byte("lease"),
		FileName:   "video.mp4",
		Attributes: []byte("ignored"),
	}
}

type fakeDFS struct {
	readBytes      []byte
	readByLease    map[string][]byte
	readErrByLease map[string]error
	readErr        error
	putErr         error
	putErrAt       int
	putResults     []*dfs.FileFinalizedObject
	readLease      []byte
	readLeases     [][]byte
	puts           []*dfs.TLDfsPutFile
}

func (f *fakeDFS) DfsGetFileByReadLease(ctx context.Context, in *dfs.TLDfsGetFileByReadLease) (*tg.UploadFile, error) {
	f.readLease = append([]byte(nil), in.ReadLease...)
	f.readLeases = append(f.readLeases, append([]byte(nil), in.ReadLease...))
	if err := f.readErrByLease[string(in.ReadLease)]; err != nil {
		return nil, err
	}
	if f.readErr != nil {
		return nil, f.readErr
	}
	if bytes, ok := f.readByLease[string(in.ReadLease)]; ok {
		return (&tg.TLUploadFile{Bytes: bytes}).ToUploadFile(), nil
	}
	return (&tg.TLUploadFile{Bytes: f.readBytes}).ToUploadFile(), nil
}

func (f *fakeDFS) DfsPutFile(ctx context.Context, in *dfs.TLDfsPutFile) (*dfs.FileFinalizedObject, error) {
	f.puts = append(f.puts, in)
	if f.putErr != nil && (f.putErrAt == 0 || f.putErrAt == len(f.puts)) {
		return nil, f.putErr
	}
	if len(f.putResults) >= len(f.puts) {
		return f.putResults[len(f.puts)-1], nil
	}
	return dfs.MakeTLFileFinalizedObject(&dfs.TLFileFinalizedObject{
		ObjectId: fmt.Sprintf("put-%d", len(f.puts)),
		Size2:    int64(1000 + len(in.Bytes)),
		MimeType: in.MimeType,
	}).ToFileFinalizedObject(), nil
}

type fakeProcessor struct {
	resized       []imaging2.PhotoSizeBytes
	resizeInput   []byte
	resizeExt     string
	resizeProfile bool
	resizeErr     error

	mp4           []byte
	convertErr    error
	convertCalled bool

	metadata *ffmpeg2.VideoMetadata
	probeErr error

	cover       []byte
	coverErr    error
	coverCalled bool
}

func (f *fakeProcessor) ResizePhoto(ctx context.Context, input []byte, ext string, profile bool) ([]imaging2.PhotoSizeBytes, error) {
	f.resizeInput = append([]byte(nil), input...)
	f.resizeExt = ext
	f.resizeProfile = profile
	if f.resizeErr != nil {
		return nil, f.resizeErr
	}
	return f.resized, nil
}

func (f *fakeProcessor) ConvertGIFToMP4(ctx context.Context, input []byte) ([]byte, error) {
	f.convertCalled = true
	if f.convertErr != nil {
		return nil, f.convertErr
	}
	return f.mp4, nil
}

func (f *fakeProcessor) ExtractMP4Cover(ctx context.Context, input []byte) ([]byte, error) {
	f.coverCalled = true
	if f.coverErr != nil {
		return nil, f.coverErr
	}
	return f.cover, nil
}

func (f *fakeProcessor) ProbeMP4(ctx context.Context, input []byte) (*ffmpeg2.VideoMetadata, error) {
	if f.probeErr != nil {
		return nil, f.probeErr
	}
	return f.metadata, nil
}

func assertPut(t *testing.T, put *dfs.TLDfsPutFile, purpose, fileName, mimeType string, bytes []byte) {
	t.Helper()
	if put.OwnerId != 1001 || put.Purpose != purpose || put.FileName != fileName || put.MimeType != mimeType || string(put.Bytes) != string(bytes) {
		t.Fatalf("put = %+v, bytes=%q", put, put.Bytes)
	}
}

func assertDerivative(t *testing.T, got mediaprocessor.ProcessorDerivativeClazz, kind, objectID, fileName, mimeType string, size int64, width, height int32, bytes []byte) {
	t.Helper()
	if got.Kind != kind || got.ObjectId != objectID || got.FileName != fileName || got.MimeType != mimeType ||
		got.Size2 != size || got.Width != width || got.Height != height || string(got.Bytes) != string(bytes) {
		t.Fatalf("derivative = %+v", got)
	}
}

func assertAttributes(t *testing.T, encoded []byte, animated bool, fileName string) {
	t.Helper()
	decoded, err := processor.DecodeDocumentAttributes(encoded)
	if err != nil {
		t.Fatalf("DecodeDocumentAttributes() error = %v", err)
	}
	wantLen := 2
	if animated {
		wantLen = 3
	}
	if len(decoded) != wantLen {
		t.Fatalf("decoded len = %d, want %d", len(decoded), wantLen)
	}
	if _, ok := decoded[0].(*tg.TLDocumentAttributeVideo); !ok {
		t.Fatalf("decoded[0] = %T, want video", decoded[0])
	}
	name, ok := decoded[1].(*tg.TLDocumentAttributeFilename)
	if !ok {
		t.Fatalf("decoded[1] = %T, want filename", decoded[1])
	}
	if filepath.Base(name.FileName) != fileName {
		t.Fatalf("filename = %q, want %q", name.FileName, fileName)
	}
	if animated {
		if _, ok := decoded[2].(*tg.TLDocumentAttributeAnimated); !ok {
			t.Fatalf("decoded[2] = %T, want animated", decoded[2])
		}
	}
}
