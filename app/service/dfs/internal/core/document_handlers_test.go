package core

import (
	"bytes"
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/ffmpeg2"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeDocumentRepository struct {
	nextDocumentID  int64
	nextEncryptedID int64
	documents       map[int64][]byte
	encrypted       map[int64][]byte
	thumbs          map[int64][]byte
	mp4             []byte
	frame           []byte
	metadata        *ffmpeg2.VideoMetadata
}

func newFakeDocumentRepository() *fakeDocumentRepository {
	return &fakeDocumentRepository{
		nextDocumentID:  7001,
		nextEncryptedID: 8001,
		documents:       make(map[int64][]byte),
		encrypted:       make(map[int64][]byte),
		thumbs:          make(map[int64][]byte),
		mp4:             []byte("converted-mp4"),
		frame:           []byte("frame-jpeg"),
		metadata:        &ffmpeg2.VideoMetadata{Width: 640, Height: 360, Duration: 7},
	}
}

func (f *fakeDocumentRepository) NextDocumentID(context.Context) (int64, error) {
	id := f.nextDocumentID
	f.nextDocumentID++
	return id, nil
}

func (f *fakeDocumentRepository) NextEncryptedFileID(context.Context) (int64, error) {
	id := f.nextEncryptedID
	f.nextEncryptedID++
	return id, nil
}

func (f *fakeDocumentRepository) SaveDocumentObject(_ context.Context, documentID int64, data []byte) (int64, error) {
	f.documents[documentID] = append([]byte(nil), data...)
	return int64(len(data)), nil
}

func (f *fakeDocumentRepository) SaveEncryptedObject(_ context.Context, fileID int64, data []byte) (int64, error) {
	f.encrypted[fileID] = append([]byte(nil), data...)
	return int64(len(data)), nil
}

func (f *fakeDocumentRepository) SaveDocumentThumbs(_ context.Context, documentID int64, image []byte, _ string) ([]repository.StoredDocumentThumb, error) {
	f.thumbs[documentID] = append([]byte(nil), image...)
	return []repository.StoredDocumentThumb{
		{Type: "i", Bytes: []byte("stripped")},
		{Type: "m", W: 320, H: 180, Size: 123},
	}, nil
}

func (f *fakeDocumentRepository) ConvertDocumentToMP4(context.Context, []byte) ([]byte, error) {
	return append([]byte(nil), f.mp4...), nil
}

func (f *fakeDocumentRepository) ExtractDocumentFrame(context.Context, []byte) ([]byte, error) {
	return append([]byte(nil), f.frame...), nil
}

func (f *fakeDocumentRepository) GetDocumentVideoMetadata(context.Context, []byte) (*ffmpeg2.VideoMetadata, error) {
	return f.metadata, nil
}

func TestDfsUploadDocumentFileV2CachesUploadRefAndReturnsDocument(t *testing.T) {
	core, uploads, docs := newDocumentTestCore(t)
	writeUploadedTestFile(t, uploads, 1001, 2002, []byte("document-bytes"))

	out, err := core.DfsUploadDocumentFileV2(&dfs.TLDfsUploadDocumentFileV2{
		Creator: 1001,
		Media: tg.MakeTLInputMediaUploadedDocument(&tg.TLInputMediaUploadedDocument{
			File:     tg.MakeTLInputFile(&tg.TLInputFile{Id: 2002, Parts: 1, Name: "report.pdf"}),
			MimeType: "application/pdf",
			Attributes: []tg.DocumentAttributeClazz{
				tg.MakeTLDocumentAttributeAnimated(&tg.TLDocumentAttributeAnimated{}),
				tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{FileName: ""}),
				tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{FileName: "report.pdf"}),
			},
		}),
	})
	if err != nil {
		t.Fatalf("DfsUploadDocumentFileV2() error = %v", err)
	}
	document, ok := out.ToDocument()
	if !ok {
		t.Fatalf("Document clazz = %s, want document", out.ClazzName())
	}
	if document.Id != 7001 {
		t.Fatalf("document.Id = %d, want 7001", document.Id)
	}
	if document.MimeType != "application/pdf" || document.Size2 != int64(len("document-bytes")) {
		t.Fatalf("document mime/size = %s/%d", document.MimeType, document.Size2)
	}
	if len(document.Attributes) != 1 {
		t.Fatalf("len(document.Attributes) = %d, want 1", len(document.Attributes))
	}
	if string(docs.documents[7001]) != "document-bytes" {
		t.Fatalf("stored document = %q, want document-bytes", docs.documents[7001])
	}
	info, err := uploads.LoadObjectCacheRef(context.Background(), 7001)
	if err != nil {
		t.Fatalf("LoadObjectCacheRef() error = %v", err)
	}
	if info.Creator != 1001 || info.FileID != 2002 {
		t.Fatalf("cache ref = %+v, want creator/file 1001/2002", info)
	}
}

func TestDfsUploadGifDocumentMediaUsesAnimatedGifPath(t *testing.T) {
	core, uploads, docs := newDocumentTestCore(t)
	writeUploadedTestFile(t, uploads, 1001, 4004, bytes.Repeat([]byte("g"), gifNeedConvertSize+1))

	out, err := core.DfsUploadGifDocumentMedia(&dfs.TLDfsUploadGifDocumentMedia{
		Creator: 1001,
		Media: tg.MakeTLInputMediaUploadedDocument(&tg.TLInputMediaUploadedDocument{
			File: tg.MakeTLInputFile(&tg.TLInputFile{Id: 4004, Parts: 1, Name: "loop.gif"}),
			Attributes: []tg.DocumentAttributeClazz{
				tg.MakeTLDocumentAttributeImageSize(&tg.TLDocumentAttributeImageSize{W: 640, H: 360}),
			},
		}),
	})
	if err != nil {
		t.Fatalf("DfsUploadGifDocumentMedia() error = %v", err)
	}
	document, ok := out.ToDocument()
	if !ok {
		t.Fatalf("Document clazz = %s, want document", out.ClazzName())
	}
	if document.MimeType != "video/mp4" {
		t.Fatalf("MimeType = %s, want video/mp4", document.MimeType)
	}
	if string(docs.documents[7001]) != "converted-mp4" {
		t.Fatalf("stored document = %q, want converted-mp4", docs.documents[7001])
	}
	if len(document.Thumbs) != 2 {
		t.Fatalf("len(Thumbs) = %d, want 2", len(document.Thumbs))
	}
	assertDocumentHasAttribute(t, document.Attributes, tg.ClazzName_documentAttributeVideo)
	assertDocumentHasAttribute(t, document.Attributes, tg.ClazzName_documentAttributeFilename)
	assertDocumentHasAttribute(t, document.Attributes, tg.ClazzName_documentAttributeAnimated)
}

func TestDfsUploadMp4DocumentMediaCreatesThumbs(t *testing.T) {
	core, uploads, docs := newDocumentTestCore(t)
	writeUploadedTestFile(t, uploads, 1001, 5005, []byte("mp4-bytes"))

	out, err := core.DfsUploadMp4DocumentMedia(&dfs.TLDfsUploadMp4DocumentMedia{
		Creator: 1001,
		Media: tg.MakeTLInputMediaUploadedDocument(&tg.TLInputMediaUploadedDocument{
			File:     tg.MakeTLInputFile(&tg.TLInputFile{Id: 5005, Parts: 1, Name: "movie.mp4"}),
			MimeType: "video/mp4",
			Attributes: []tg.DocumentAttributeClazz{
				tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{FileName: "movie.mp4"}),
				tg.MakeTLDocumentAttributeVideo(&tg.TLDocumentAttributeVideo{Duration: 7, W: 640, H: 360}),
			},
		}),
	})
	if err != nil {
		t.Fatalf("DfsUploadMp4DocumentMedia() error = %v", err)
	}
	if string(docs.documents[7001]) != "mp4-bytes" {
		t.Fatalf("stored document = %q, want mp4-bytes", docs.documents[7001])
	}
	if string(docs.thumbs[7001]) != "frame-jpeg" {
		t.Fatalf("stored thumb source = %q, want frame-jpeg", docs.thumbs[7001])
	}
	document, ok := out.ToDocument()
	if !ok {
		t.Fatalf("Document clazz = %s, want document", out.ClazzName())
	}
	if len(document.Thumbs) != 2 {
		t.Fatalf("len(Thumbs) = %d, want 2", len(document.Thumbs))
	}
	videoAttr := assertDocumentHasVideoAttribute(t, document.Attributes)
	if !videoAttr.SupportsStreaming {
		t.Fatal("video attribute SupportsStreaming = false, want true")
	}
}

func TestDfsUploadWallPaperFileCreatesDocumentAndThumb(t *testing.T) {
	core, uploads, docs := newDocumentTestCore(t)
	writeUploadedTestFile(t, uploads, 1001, 6006, []byte("wallpaper-bytes"))

	out, err := core.DfsUploadWallPaperFile(&dfs.TLDfsUploadWallPaperFile{
		Creator:  1001,
		File:     tg.MakeTLInputFile(&tg.TLInputFile{Id: 6006, Parts: 1, Name: "wallpaper.jpg"}),
		MimeType: "image/jpeg",
		Admin:    tg.BoolTrueClazz,
	})
	if err != nil {
		t.Fatalf("DfsUploadWallPaperFile() error = %v", err)
	}
	if string(docs.documents[7001]) != "wallpaper-bytes" {
		t.Fatalf("stored document = %q, want wallpaper-bytes", docs.documents[7001])
	}
	document, ok := out.ToDocument()
	if !ok {
		t.Fatalf("Document clazz = %s, want document", out.ClazzName())
	}
	if len(document.Thumbs) != 2 {
		t.Fatalf("len(Thumbs) = %d, want 2", len(document.Thumbs))
	}
	assertDocumentHasAttribute(t, document.Attributes, tg.ClazzName_documentAttributeImageSize)
	assertDocumentHasAttribute(t, document.Attributes, tg.ClazzName_documentAttributeFilename)
}

func TestDfsUploadThemeFileSupportsOptionalThumb(t *testing.T) {
	core, uploads, docs := newDocumentTestCore(t)
	writeUploadedTestFile(t, uploads, 1001, 7007, []byte("theme-bytes"))
	writeUploadedTestFile(t, uploads, 1001, 7008, []byte("theme-thumb"))

	out, err := core.DfsUploadThemeFile(&dfs.TLDfsUploadThemeFile{
		Creator:  1001,
		File:     tg.MakeTLInputFile(&tg.TLInputFile{Id: 7007, Parts: 1, Name: "theme.tdesktop-theme"}),
		Thumb:    tg.MakeTLInputFile(&tg.TLInputFile{Id: 7008, Parts: 1, Name: "theme.jpg"}),
		MimeType: "application/x-tgtheme",
		FileName: "theme.tdesktop-theme",
	})
	if err != nil {
		t.Fatalf("DfsUploadThemeFile() error = %v", err)
	}
	if string(docs.documents[7001]) != "theme-bytes" {
		t.Fatalf("stored document = %q, want theme-bytes", docs.documents[7001])
	}
	if string(docs.thumbs[7001]) != "theme-thumb" {
		t.Fatalf("stored thumb source = %q, want theme-thumb", docs.thumbs[7001])
	}
	document, ok := out.ToDocument()
	if !ok {
		t.Fatalf("Document clazz = %s, want document", out.ClazzName())
	}
	if len(document.Thumbs) != 2 {
		t.Fatalf("len(Thumbs) = %d, want 2", len(document.Thumbs))
	}
	assertDocumentHasAttribute(t, document.Attributes, tg.ClazzName_documentAttributeFilename)
}

func TestDfsUploadRingtoneFileSetsAudioAttributes(t *testing.T) {
	core, uploads, docs := newDocumentTestCore(t)
	writeUploadedTestFile(t, uploads, 1001, 8008, []byte("ringtone-bytes"))
	docs.metadata = &ffmpeg2.VideoMetadata{Duration: 11}

	out, err := core.DfsUploadRingtoneFile(&dfs.TLDfsUploadRingtoneFile{
		Creator:  1001,
		File:     tg.MakeTLInputFile(&tg.TLInputFile{Id: 8008, Parts: 1, Name: "ringtone.mp3"}),
		MimeType: "audio/mpeg",
		FileName: "tone",
	})
	if err != nil {
		t.Fatalf("DfsUploadRingtoneFile() error = %v", err)
	}
	if string(docs.documents[7001]) != "ringtone-bytes" {
		t.Fatalf("stored document = %q, want ringtone-bytes", docs.documents[7001])
	}
	document, ok := out.ToDocument()
	if !ok {
		t.Fatalf("Document clazz = %s, want document", out.ClazzName())
	}
	audioAttr := assertDocumentHasAudioAttribute(t, document.Attributes)
	if audioAttr.Duration != 11 || audioAttr.Title == nil || *audioAttr.Title != "tone" {
		t.Fatalf("audio attribute = %+v, want duration 11 title tone", audioAttr)
	}
	assertDocumentHasAttribute(t, document.Attributes, tg.ClazzName_documentAttributeFilename)
}

func TestDfsUploadEncryptedFileV2StoresEncryptedObject(t *testing.T) {
	core, uploads, docs := newDocumentTestCore(t)
	writeUploadedTestFile(t, uploads, 1001, 3003, []byte("secret-bytes"))

	out, err := core.DfsUploadEncryptedFileV2(&dfs.TLDfsUploadEncryptedFileV2{
		Creator: 1001,
		File: tg.MakeTLInputEncryptedFileUploaded(&tg.TLInputEncryptedFileUploaded{
			Id:             3003,
			Parts:          1,
			Md5Checksum:    "dbfc252afe1c84d312028ac2d558a2c3",
			KeyFingerprint: 1234,
		}),
	})
	if err != nil {
		t.Fatalf("DfsUploadEncryptedFileV2() error = %v", err)
	}
	encrypted, ok := out.ToEncryptedFile()
	if !ok {
		t.Fatalf("EncryptedFile clazz = %s, want encryptedFile", out.ClazzName())
	}
	if encrypted.Id != 8001 {
		t.Fatalf("encrypted.Id = %d, want 8001", encrypted.Id)
	}
	if encrypted.Size2 != int64(len("secret-bytes")) {
		t.Fatalf("encrypted.Size2 = %d, want %d", encrypted.Size2, len("secret-bytes"))
	}
	if encrypted.KeyFingerprint != 0 {
		t.Fatalf("encrypted.KeyFingerprint = %d, want 0", encrypted.KeyFingerprint)
	}
	if string(docs.encrypted[8001]) != "secret-bytes" {
		t.Fatalf("stored encrypted = %q, want secret-bytes", docs.encrypted[8001])
	}
	info, err := uploads.LoadObjectCacheRef(context.Background(), 8001)
	if err != nil {
		t.Fatalf("LoadObjectCacheRef() error = %v", err)
	}
	if info.Creator != 1001 || info.FileID != 3003 {
		t.Fatalf("cache ref = %+v, want creator/file 1001/3003", info)
	}
}

func assertDocumentHasAttribute(t *testing.T, attrs []tg.DocumentAttributeClazz, clazzName string) {
	t.Helper()
	for _, attr := range attrs {
		if attr.DocumentAttributeClazzName() == clazzName {
			return
		}
	}
	t.Fatalf("attributes missing %s: %#v", clazzName, attrs)
}

func assertDocumentHasVideoAttribute(t *testing.T, attrs []tg.DocumentAttributeClazz) *tg.TLDocumentAttributeVideo {
	t.Helper()
	for _, attr := range attrs {
		if video, ok := attr.(*tg.TLDocumentAttributeVideo); ok {
			return video
		}
	}
	t.Fatalf("attributes missing video: %#v", attrs)
	return nil
}

func assertDocumentHasAudioAttribute(t *testing.T, attrs []tg.DocumentAttributeClazz) *tg.TLDocumentAttributeAudio {
	t.Helper()
	for _, attr := range attrs {
		if audio, ok := attr.(*tg.TLDocumentAttributeAudio); ok {
			return audio
		}
	}
	t.Fatalf("attributes missing audio: %#v", attrs)
	return nil
}

func newDocumentTestCore(t *testing.T) (*DfsCore, *UploadSessionManager, *fakeDocumentRepository) {
	t.Helper()
	uploadRepo := newFakeUploadStateRepository()
	uploads := NewUploadSessionManager(uploadRepo)
	docs := newFakeDocumentRepository()
	return &DfsCore{
		ctx:                  context.Background(),
		uploadSessionManager: uploads,
		documentRepository:   docs,
	}, uploads, docs
}
