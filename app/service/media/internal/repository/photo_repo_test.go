package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"testing"
	"time"

	dfsapi "github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/app/service/media/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/app/service/mediaprocessor/mediaprocessor"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const testLayer = 223

type photoModelNotFound struct {
	model.PhotosModel
}

func (photoModelNotFound) FindOneByPhotoId(context.Context, int64) (*model.Photos, error) {
	return nil, &model.NotFoundError{Resource: "photos", Key: "photo_id=10"}
}

type capturePhotosModel struct {
	model.PhotosModel
	inserted []*model.Photos
	found    *model.Photos
}

func (m *capturePhotosModel) Insert(_ context.Context, data *model.Photos) (int64, int64, error) {
	m.inserted = append(m.inserted, data)
	return 1, 1, nil
}

func (m *capturePhotosModel) Insert2(_ context.Context, data *model.Photos) (sql.Result, error) {
	m.inserted = append(m.inserted, data)
	return nil, nil
}

func (m *capturePhotosModel) FindOneByPhotoId(_ context.Context, _ int64) (*model.Photos, error) {
	if m.found == nil {
		return nil, &model.NotFoundError{Resource: "photos", Key: "photo_id"}
	}
	return m.found, nil
}

type capturePhotoSizesModel struct {
	model.PhotoSizesModel
	inserted []*model.PhotoSizes
	byID     []model.PhotoSizes
}

func (m *capturePhotoSizesModel) Insert(_ context.Context, data *model.PhotoSizes) (int64, int64, error) {
	m.inserted = append(m.inserted, data)
	return 1, 1, nil
}

func (m *capturePhotoSizesModel) SelectListByPhotoSizeId(_ context.Context, _ int64) ([]model.PhotoSizes, error) {
	return m.byID, nil
}

type captureVideoSizesModel struct {
	model.VideoSizesModel
	inserted []*model.VideoSizes
	byID     []model.VideoSizes
}

func (m *captureVideoSizesModel) Insert(_ context.Context, data *model.VideoSizes) (int64, int64, error) {
	m.inserted = append(m.inserted, data)
	return 1, 1, nil
}

func (m *captureVideoSizesModel) SelectListByVideoSizeId(_ context.Context, _ int64) ([]model.VideoSizes, error) {
	return m.byID, nil
}

type fakeDfsMediaClient struct {
	photo              *tg.Photo
	document           *tg.Document
	finalized          *dfsapi.FileFinalizedObject
	commitReq          *dfsapi.TLDfsCommitUpload
	putReq             *dfsapi.TLDfsPutFile
	uploadPhotoRequest *dfsapi.TLDfsUploadPhotoFileV2
	uploadProfileReq   *dfsapi.TLDfsUploadProfilePhotoFileV2
	uploadedProfileReq *dfsapi.TLDfsUploadedProfilePhoto
	uploadDocumentReq  *dfsapi.TLDfsUploadDocumentFileV2
	uploadGifReq       *dfsapi.TLDfsUploadGifDocumentMedia
	uploadMp4Req       *dfsapi.TLDfsUploadMp4DocumentMedia
}

func (c *fakeDfsMediaClient) CommitUpload(_ context.Context, in *dfsapi.TLDfsCommitUpload) (*dfsapi.FileFinalizedObject, error) {
	c.commitReq = in
	return c.finalized, nil
}

func (c *fakeDfsMediaClient) PutFile(_ context.Context, in *dfsapi.TLDfsPutFile) (*dfsapi.FileFinalizedObject, error) {
	c.putReq = in
	return c.finalized, nil
}

func (c *fakeDfsMediaClient) UploadPhotoFileV2ViaLegacyDFS(_ context.Context, in *dfsapi.TLDfsUploadPhotoFileV2) (*tg.Photo, error) {
	c.uploadPhotoRequest = in
	return c.photo, nil
}

func (c *fakeDfsMediaClient) UploadProfilePhotoFileV2ViaLegacyDFS(_ context.Context, in *dfsapi.TLDfsUploadProfilePhotoFileV2) (*tg.Photo, error) {
	c.uploadProfileReq = in
	return c.photo, nil
}

func (c *fakeDfsMediaClient) UploadDocumentFileV2ViaLegacyDFS(_ context.Context, in *dfsapi.TLDfsUploadDocumentFileV2) (*tg.Document, error) {
	c.uploadDocumentReq = in
	return c.document, nil
}

func (c *fakeDfsMediaClient) UploadGifDocumentMediaViaLegacyDFS(_ context.Context, in *dfsapi.TLDfsUploadGifDocumentMedia) (*tg.Document, error) {
	c.uploadGifReq = in
	return c.document, nil
}

func (c *fakeDfsMediaClient) UploadMp4DocumentMediaViaLegacyDFS(_ context.Context, in *dfsapi.TLDfsUploadMp4DocumentMedia) (*tg.Document, error) {
	c.uploadMp4Req = in
	return c.document, nil
}

func (c *fakeDfsMediaClient) UploadedProfilePhoto(_ context.Context, in *dfsapi.TLDfsUploadedProfilePhoto) (*tg.Photo, error) {
	c.uploadedProfileReq = in
	return c.photo, nil
}

type fakeMediaProcessorClient struct {
	photoReq *mediaprocessor.TLMediaProcessorProcessPhoto
	gifReq   *mediaprocessor.TLMediaProcessorProcessGif
	mp4Req   *mediaprocessor.TLMediaProcessorProcessMp4
	photo    *mediaprocessor.ProcessedPhoto
	document *mediaprocessor.ProcessedDocument
}

func (c *fakeMediaProcessorClient) ProcessPhoto(_ context.Context, in *mediaprocessor.TLMediaProcessorProcessPhoto) (*mediaprocessor.ProcessedPhoto, error) {
	c.photoReq = in
	return c.photo, nil
}

func (c *fakeMediaProcessorClient) ProcessGif(_ context.Context, in *mediaprocessor.TLMediaProcessorProcessGif) (*mediaprocessor.ProcessedDocument, error) {
	c.gifReq = in
	return c.document, nil
}

func (c *fakeMediaProcessorClient) ProcessMp4(_ context.Context, in *mediaprocessor.TLMediaProcessorProcessMp4) (*mediaprocessor.ProcessedDocument, error) {
	c.mp4Req = in
	return c.document, nil
}

func TestGetPhotoReturnsStorageError(t *testing.T) {
	r := &Repository{}
	errBoom := errors.New("db down")
	_, err := r.mapPhotoResult(context.Background(), nil, errBoom)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestGetPhotoMapsModelNotFound(t *testing.T) {
	r := &Repository{model: &model.Models{PhotosModel: photoModelNotFound{}}}
	_, err := r.GetPhoto(context.Background(), 10)
	if !errors.Is(err, media.ErrPhotoNotFound) {
		t.Fatalf("expected ErrPhotoNotFound, got %v", err)
	}
	if errors.Is(err, media.ErrMediaStorage) {
		t.Fatalf("expected semantic not found, got storage error: %v", err)
	}
}

func TestGetPhotoReturnsPhoto(t *testing.T) {
	r := &Repository{}
	photo := tg.MakeTLPhoto(&tg.TLPhoto{Id: 10}).ToPhoto()
	got, err := r.mapPhotoResult(context.Background(), photo, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	gotPhoto, ok := got.ToPhoto()
	if !ok {
		t.Fatalf("expected photo, got %#v", got)
	}
	if gotPhoto.Id != 10 {
		t.Fatalf("expected photo id 10, got %d", gotPhoto.Id)
	}
}

func TestPhotoFromModelBuildsValidMinimalPhoto(t *testing.T) {
	got, err := mapPhotoAggregate(&model.Photos{
		PhotoId:     10,
		AccessHash:  20,
		HasStickers: true,
		DcId:        4,
		Date2:       30,
	}, nil, nil, []byte("file-reference"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	gotPhoto, ok := got.ToPhoto()
	if !ok {
		t.Fatalf("expected photo, got %#v", got)
	}
	if gotPhoto.Id != 10 {
		t.Fatalf("expected photo id 10, got %d", gotPhoto.Id)
	}
	if gotPhoto.FileReference == nil {
		t.Fatal("expected non-nil file_reference")
	}
	if gotPhoto.Sizes == nil {
		t.Fatal("expected non-nil required sizes")
	}
	if gotPhoto.VideoSizes != nil {
		t.Fatalf("expected absent video_sizes, got %#v", gotPhoto.VideoSizes)
	}
	if err := gotPhoto.Validate(testLayer); err != nil {
		t.Fatalf("expected valid photo: %v", err)
	}
}

func TestUploadPhotoFileCommitsAndProcessesPhoto(t *testing.T) {
	photos := &capturePhotosModel{}
	photoSizes := &capturePhotoSizesModel{}
	dfsClient := &fakeDfsMediaClient{
		photo: testPhotoWithSizes(101, false),
		finalized: dfsapi.MakeTLFileFinalizedObject(&dfsapi.TLFileFinalizedObject{
			ObjectId:        "original-object",
			UploadSessionId: "ext:7:11:1",
			ReadLease:       []byte("read-lease"),
			DcId:            2,
		}).ToFileFinalizedObject(),
	}
	processorClient := &fakeMediaProcessorClient{
		photo: mediaprocessor.MakeTLProcessedPhoto(&mediaprocessor.TLProcessedPhoto{
			OriginalObjectId: "original-object",
			Sizes: []mediaprocessor.ProcessorDerivativeClazz{
				mediaprocessor.MakeTLProcessorDerivative(&mediaprocessor.TLProcessorDerivative{
					Kind:     "photo_size",
					ObjectId: "derivative-object-s",
					FileName: "s_avatar.jpg",
					MimeType: "image/jpeg",
					Size2:    600,
					Width:    160,
					Height:   120,
					Bytes:    []byte("preview"),
				}),
				mediaprocessor.MakeTLProcessorDerivative(&mediaprocessor.TLProcessorDerivative{
					Kind:     "photo_size",
					ObjectId: "derivative-object-m",
					FileName: "m_avatar.jpg",
					MimeType: "image/jpeg",
					Size2:    1200,
					Width:    320,
					Height:   240,
					Bytes:    []byte("preview"),
				}),
			},
		}).ToProcessedPhoto(),
	}
	r := &Repository{
		model:                &model.Models{PhotosModel: photos, PhotoSizesModel: photoSizes, VideoSizesModel: &captureVideoSizesModel{}},
		dfsClient:            dfsClient,
		processorClient:      processorClient,
		fileReferenceService: NewFileReferenceService([]byte("test-secret"), func() time.Time { return time.Unix(1700000000, 0) }),
		fileReferenceTTL:     time.Hour,
	}

	got, err := r.UploadPhotoFile(context.Background(), &media.TLMediaUploadPhotoFile{
		OwnerId: 7,
		File:    tg.MakeTLInputFile(&tg.TLInputFile{Id: 11, Parts: 1, Name: "avatar.jpg", Md5Checksum: "md5"}),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dfsClient.commitReq == nil {
		t.Fatal("expected dfs CommitUpload request")
	}
	if dfsClient.uploadPhotoRequest != nil {
		t.Fatalf("expected default path not to call UploadPhotoFileV2, got %#v", dfsClient.uploadPhotoRequest)
	}
	if dfsClient.commitReq.UploadSessionId != "ext:7:11:1" || dfsClient.commitReq.OwnerId != 7 || dfsClient.commitReq.Purpose != "media_original" {
		t.Fatalf("unexpected dfs commit request: %#v", dfsClient.commitReq)
	}
	if processorClient.photoReq == nil {
		t.Fatal("expected mediaprocessor ProcessPhoto request")
	}
	if processorClient.photoReq.OwnerId != 7 || processorClient.photoReq.ObjectId != "original-object" || string(processorClient.photoReq.ReadLease) != "read-lease" || processorClient.photoReq.FileName != "avatar.jpg" {
		t.Fatalf("unexpected process photo request: %#v", processorClient.photoReq)
	}
	gotPhoto, ok := got.ToPhoto()
	if !ok {
		t.Fatalf("expected photo, got %#v", got)
	}
	if len(gotPhoto.FileReference) == 0 {
		t.Fatal("expected non-empty file_reference")
	}
	if len(photos.inserted) != 1 || photos.inserted[0].InputFileName != "avatar.jpg" {
		t.Fatalf("expected saved photo with input file name, got %#v", photos.inserted)
	}
	if photos.inserted[0].PhotoId == 0 || photos.inserted[0].AccessHash == 0 || photos.inserted[0].DcId != 2 || photos.inserted[0].Date2 == 0 {
		t.Fatalf("expected saved photo identifiers, dc, and date, got %#v", photos.inserted[0])
	}
	if len(photoSizes.inserted) != 2 {
		t.Fatalf("expected two saved photo sizes, got %#v", photoSizes.inserted)
	}
	sizesByType := map[string]*model.PhotoSizes{}
	for _, size := range photoSizes.inserted {
		sizesByType[size.SizeType] = size
	}
	if sizesByType["s"] == nil || sizesByType["s"].FilePath != "derivative-object-s" || sizesByType["s"].PhotoSizeId != photos.inserted[0].PhotoId || sizesByType["s"].FileSize != 600 || sizesByType["s"].Width != 160 || sizesByType["s"].Height != 120 {
		t.Fatalf("expected saved s derivative metadata, got %#v", sizesByType["s"])
	}
	if sizesByType["m"] == nil || sizesByType["m"].FilePath != "derivative-object-m" || sizesByType["m"].PhotoSizeId != photos.inserted[0].PhotoId || sizesByType["m"].FileSize != 1200 || sizesByType["m"].Width != 320 || sizesByType["m"].Height != 240 {
		t.Fatalf("expected saved m derivative metadata, got %#v", sizesByType["m"])
	}
}

func TestUploadPhotoFileCallsDfsAndSavesPhotoSizes(t *testing.T) {
	photos := &capturePhotosModel{}
	photoSizes := &capturePhotoSizesModel{}
	dfsClient := &fakeDfsMediaClient{finalized: dfsapi.MakeTLFileFinalizedObject(&dfsapi.TLFileFinalizedObject{
		ObjectId:        "original-object-legacy-test",
		UploadSessionId: "ext:7:11:1",
		ReadLease:       []byte("read-lease"),
		DcId:            1,
	}).ToFileFinalizedObject()}
	processorClient := &fakeMediaProcessorClient{photo: mediaprocessor.MakeTLProcessedPhoto(&mediaprocessor.TLProcessedPhoto{
		OriginalObjectId: "original-object-legacy-test",
		Sizes: []mediaprocessor.ProcessorDerivativeClazz{
			mediaprocessor.MakeTLProcessorDerivative(&mediaprocessor.TLProcessorDerivative{Kind: "photo_size", ObjectId: "derivative-object-legacy-test", FileName: "m_avatar.jpg", Width: 320, Height: 240, Size2: 1000}),
		},
	}).ToProcessedPhoto()}
	r := &Repository{
		model:                &model.Models{PhotosModel: photos, PhotoSizesModel: photoSizes, VideoSizesModel: &captureVideoSizesModel{}},
		dfsClient:            dfsClient,
		processorClient:      processorClient,
		fileReferenceService: NewFileReferenceService([]byte("test-secret"), func() time.Time { return time.Unix(1700000000, 0) }),
		fileReferenceTTL:     time.Hour,
	}

	got, err := r.UploadPhotoFile(context.Background(), &media.TLMediaUploadPhotoFile{
		OwnerId: 7,
		File:    tg.MakeTLInputFile(&tg.TLInputFile{Id: 11, Parts: 1, Name: "avatar.jpg", Md5Checksum: "md5"}),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got == nil {
		t.Fatal("expected uploaded photo")
	}
	if dfsClient.commitReq == nil || dfsClient.commitReq.OwnerId != 7 {
		t.Fatalf("expected dfs commit request with owner 7, got %#v", dfsClient.commitReq)
	}
	if len(photos.inserted) != 1 {
		t.Fatalf("expected one photo row, got %d", len(photos.inserted))
	}
	if photos.inserted[0].InputFileName != "avatar.jpg" {
		t.Fatalf("expected input file name saved, got %q", photos.inserted[0].InputFileName)
	}
	if len(photoSizes.inserted) != 1 || photoSizes.inserted[0].PhotoSizeId != photos.inserted[0].PhotoId {
		t.Fatalf("expected saved photo size for generated photo id, got %#v", photoSizes.inserted)
	}
	if photoSizes.inserted[0].FilePath != "derivative-object-legacy-test" {
		t.Fatalf("expected saved derivative object id, got %#v", photoSizes.inserted[0])
	}
}

func TestUploadPhotoFileMapsStrippedAndProgressiveSizes(t *testing.T) {
	photos := &capturePhotosModel{}
	photoSizes := &capturePhotoSizesModel{}
	dfsClient := &fakeDfsMediaClient{finalized: dfsapi.MakeTLFileFinalizedObject(&dfsapi.TLFileFinalizedObject{
		ObjectId:        "original-object-photo-variants",
		UploadSessionId: "ext:7:11:1",
		ReadLease:       []byte("read-lease"),
		DcId:            2,
	}).ToFileFinalizedObject()}
	processorClient := &fakeMediaProcessorClient{photo: mediaprocessor.MakeTLProcessedPhoto(&mediaprocessor.TLProcessedPhoto{
		OriginalObjectId: "original-object-photo-variants",
		Sizes: []mediaprocessor.ProcessorDerivativeClazz{
			mediaprocessor.MakeTLProcessorDerivative(&mediaprocessor.TLProcessorDerivative{Kind: "photo_stripped", FileName: "i_avatar.jpg", Width: 40, Height: 30, Size2: 10, Bytes: []byte("stripped-bytes")}),
			mediaprocessor.MakeTLProcessorDerivative(&mediaprocessor.TLProcessorDerivative{Kind: "photo_size", ObjectId: "derivative-object-s", FileName: "s_avatar.jpg", Width: 160, Height: 120, Size2: 600}),
			mediaprocessor.MakeTLProcessorDerivative(&mediaprocessor.TLProcessorDerivative{Kind: "photo_size", ObjectId: "derivative-object-x", FileName: "x_avatar.jpg", Width: 800, Height: 600, Size2: 1200, ProgressiveSizes: []int32{400, 900, 1200}}),
		},
	}).ToProcessedPhoto()}
	r := &Repository{
		model:                &model.Models{PhotosModel: photos, PhotoSizesModel: photoSizes, VideoSizesModel: &captureVideoSizesModel{}},
		dfsClient:            dfsClient,
		processorClient:      processorClient,
		fileReferenceService: NewFileReferenceService([]byte("test-secret"), func() time.Time { return time.Unix(1700000000, 0) }),
		fileReferenceTTL:     time.Hour,
	}

	got, err := r.UploadPhotoFile(context.Background(), &media.TLMediaUploadPhotoFile{
		OwnerId: 7,
		File:    tg.MakeTLInputFile(&tg.TLInputFile{Id: 11, Parts: 1, Name: "avatar.jpg", Md5Checksum: "md5"}),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	gotPhoto, ok := got.ToPhoto()
	if !ok {
		t.Fatalf("expected photo, got %#v", got)
	}
	if len(gotPhoto.Sizes) != 3 {
		t.Fatalf("expected stripped, normal, and progressive sizes, got %#v", gotPhoto.Sizes)
	}
	if stripped, ok := gotPhoto.Sizes[0].(*tg.TLPhotoStrippedSize); !ok || stripped.Type != "i" || string(stripped.Bytes) != "stripped-bytes" {
		t.Fatalf("expected stripped size with inline bytes, got %#v", gotPhoto.Sizes[0])
	}
	if normal, ok := gotPhoto.Sizes[1].(*tg.TLPhotoSize); !ok || normal.Type != "s" || normal.Size2 != 600 {
		t.Fatalf("expected normal s size, got %#v", gotPhoto.Sizes[1])
	}
	if progressive, ok := gotPhoto.Sizes[2].(*tg.TLPhotoSizeProgressive); !ok || progressive.Type != "x" || progressive.W != 800 || progressive.H != 600 || len(progressive.Sizes) != 3 || progressive.Sizes[2] != 1200 {
		t.Fatalf("expected progressive x size, got %#v", gotPhoto.Sizes[2])
	}

	if len(photoSizes.inserted) != 3 {
		t.Fatalf("expected three saved photo sizes, got %#v", photoSizes.inserted)
	}
	sizesByType := map[string]*model.PhotoSizes{}
	for _, size := range photoSizes.inserted {
		sizesByType[size.SizeType] = size
	}
	if sizesByType["i"] == nil || !sizesByType["i"].HasStripped || sizesByType["i"].StrippedBytes != "stripped-bytes" || sizesByType["i"].FilePath != "" {
		t.Fatalf("expected stripped row without object id path, got %#v", sizesByType["i"])
	}
	if sizesByType["s"] == nil || sizesByType["s"].FilePath != "derivative-object-s" || sizesByType["s"].FileSize != 600 {
		t.Fatalf("expected normal row with object id path, got %#v", sizesByType["s"])
	}
	if sizesByType["x"] == nil || sizesByType["x"].FilePath != "derivative-object-x" || sizesByType["x"].CachedType != 4 || sizesByType["x"].CachedBytes == "" {
		t.Fatalf("expected progressive row with cached scan sizes, got %#v", sizesByType["x"])
	}
	var cachedSizes []int32
	if err := json.Unmarshal([]byte(sizesByType["x"].CachedBytes), &cachedSizes); err != nil {
		t.Fatalf("expected cached progressive sizes JSON: %v", err)
	}
	if len(cachedSizes) != 3 || cachedSizes[0] != 400 || cachedSizes[1] != 900 || cachedSizes[2] != 1200 {
		t.Fatalf("unexpected cached progressive sizes: %v", cachedSizes)
	}

	photos.found = photos.inserted[0]
	photoSizes.byID = make([]model.PhotoSizes, 0, len(photoSizes.inserted))
	for _, inserted := range photoSizes.inserted {
		photoSizes.byID = append(photoSizes.byID, *inserted)
	}
	reloaded, err := r.GetPhoto(context.Background(), photos.inserted[0].PhotoId)
	if err != nil {
		t.Fatalf("unexpected reload error: %v", err)
	}
	reloadedPhoto, ok := reloaded.ToPhoto()
	if !ok {
		t.Fatalf("expected reloaded photo, got %#v", reloaded)
	}
	if len(reloadedPhoto.Sizes) != 3 {
		t.Fatalf("expected three reloaded sizes, got %#v", reloadedPhoto.Sizes)
	}
	if stripped, ok := reloadedPhoto.Sizes[0].(*tg.TLPhotoStrippedSize); !ok || stripped.Type != "i" || string(stripped.Bytes) != "stripped-bytes" {
		t.Fatalf("expected reloaded stripped size, got %#v", reloadedPhoto.Sizes[0])
	}
	if progressive, ok := reloadedPhoto.Sizes[2].(*tg.TLPhotoSizeProgressive); !ok || progressive.Type != "x" || len(progressive.Sizes) != 3 || progressive.Sizes[2] != 1200 {
		t.Fatalf("expected reloaded progressive size, got %#v", reloadedPhoto.Sizes[2])
	}
}

func TestUploadPhotoFileViaLegacyDFSCallsLegacyWrapper(t *testing.T) {
	photos := &capturePhotosModel{}
	photoSizes := &capturePhotoSizesModel{}
	dfsClient := &fakeDfsMediaClient{photo: testPhotoWithSizes(707, false)}
	r := &Repository{
		model:     &model.Models{PhotosModel: photos, PhotoSizesModel: photoSizes, VideoSizesModel: &captureVideoSizesModel{}},
		dfsClient: dfsClient,
	}

	got, err := r.UploadPhotoFileViaLegacyDFS(context.Background(), &media.TLMediaUploadPhotoFile{
		OwnerId: 7,
		File:    tg.MakeTLInputFile(&tg.TLInputFile{Id: 11, Parts: 1, Name: "avatar.jpg", Md5Checksum: "md5"}),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got == nil {
		t.Fatal("expected uploaded photo")
	}
	if dfsClient.uploadPhotoRequest == nil || dfsClient.uploadPhotoRequest.Creator != 7 {
		t.Fatalf("expected legacy dfs photo upload request, got %#v", dfsClient.uploadPhotoRequest)
	}
	if dfsClient.commitReq != nil {
		t.Fatalf("expected legacy path not to commit upload, got %#v", dfsClient.commitReq)
	}
	if len(photos.inserted) != 1 || photos.inserted[0].PhotoId != 707 || photos.inserted[0].InputFileName != "avatar.jpg" {
		t.Fatalf("expected saved legacy photo row, got %#v", photos.inserted)
	}
	if len(photoSizes.inserted) != 1 || photoSizes.inserted[0].PhotoSizeId != 707 {
		t.Fatalf("expected saved legacy photo size, got %#v", photoSizes.inserted)
	}
}

func TestUploadPhotoFileViaLegacyDFSRejectsNilLegacyPhoto(t *testing.T) {
	r := &Repository{
		model:     &model.Models{PhotosModel: &capturePhotosModel{}, PhotoSizesModel: &capturePhotoSizesModel{}, VideoSizesModel: &captureVideoSizesModel{}},
		dfsClient: &fakeDfsMediaClient{},
	}

	_, err := r.UploadPhotoFileViaLegacyDFS(context.Background(), &media.TLMediaUploadPhotoFile{
		OwnerId: 7,
		File:    tg.MakeTLInputFile(&tg.TLInputFile{Id: 11, Parts: 1, Name: "avatar.jpg", Md5Checksum: "md5"}),
	})
	if !errors.Is(err, media.ErrMediaInvalidUploadedFile) {
		t.Fatalf("expected ErrMediaInvalidUploadedFile, got %v", err)
	}
}

func TestUploadPhotoFileRejectsInvalidProcessorDerivativeWithoutPersisting(t *testing.T) {
	photos := &capturePhotosModel{}
	photoSizes := &capturePhotoSizesModel{}
	dfsClient := &fakeDfsMediaClient{finalized: dfsapi.MakeTLFileFinalizedObject(&dfsapi.TLFileFinalizedObject{
		ObjectId:        "original-object-invalid-derivative",
		UploadSessionId: "ext:7:11:1",
		ReadLease:       []byte("read-lease"),
		DcId:            1,
	}).ToFileFinalizedObject()}
	processorClient := &fakeMediaProcessorClient{photo: mediaprocessor.MakeTLProcessedPhoto(&mediaprocessor.TLProcessedPhoto{
		OriginalObjectId: "original-object-invalid-derivative",
		Sizes: []mediaprocessor.ProcessorDerivativeClazz{
			mediaprocessor.MakeTLProcessorDerivative(&mediaprocessor.TLProcessorDerivative{Kind: "photo_size", ObjectId: "derivative-object-s", FileName: "s_avatar.jpg", Width: 160, Height: 120, Size2: 600}),
			mediaprocessor.MakeTLProcessorDerivative(&mediaprocessor.TLProcessorDerivative{Kind: "photo_size", ObjectId: "derivative-object-bad", FileName: "avatar.jpg", Width: 320, Height: 240, Size2: 1000}),
		},
	}).ToProcessedPhoto()}
	r := &Repository{
		model:                &model.Models{PhotosModel: photos, PhotoSizesModel: photoSizes, VideoSizesModel: &captureVideoSizesModel{}},
		dfsClient:            dfsClient,
		processorClient:      processorClient,
		fileReferenceService: NewFileReferenceService([]byte("test-secret"), func() time.Time { return time.Unix(1700000000, 0) }),
		fileReferenceTTL:     time.Hour,
	}

	_, err := r.UploadPhotoFile(context.Background(), &media.TLMediaUploadPhotoFile{
		OwnerId: 7,
		File:    tg.MakeTLInputFile(&tg.TLInputFile{Id: 11, Parts: 1, Name: "avatar.jpg", Md5Checksum: "md5"}),
	})
	if !errors.Is(err, media.ErrMediaInvalidUploadedFile) {
		t.Fatalf("expected invalid uploaded file error, got %v", err)
	}
	if len(photos.inserted) != 0 || len(photoSizes.inserted) != 0 {
		t.Fatalf("expected no persisted rows, got photos=%#v sizes=%#v", photos.inserted, photoSizes.inserted)
	}
}

func TestUploadProfilePhotoFileSavesVideoSizes(t *testing.T) {
	photos := &capturePhotosModel{}
	photoSizes := &capturePhotoSizesModel{}
	videos := &captureVideoSizesModel{}
	startTs := 1.25
	dfsClient := &fakeDfsMediaClient{finalized: testFinalizedObject("profile-video-object", 4096)}
	processorClient := &fakeMediaProcessorClient{document: testProcessedDocument(t, "profile-video-object", "video/mp4", 4096, []tg.DocumentAttributeClazz{
		tg.MakeTLDocumentAttributeVideo(&tg.TLDocumentAttributeVideo{Duration: 3, W: 640, H: 360}),
		tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{FileName: "profile.mp4"}),
	}, testDocumentThumbs())}
	r := &Repository{
		model:                &model.Models{PhotosModel: photos, PhotoSizesModel: photoSizes, VideoSizesModel: videos},
		dfsClient:            dfsClient,
		processorClient:      processorClient,
		fileReferenceService: NewFileReferenceService([]byte("test-secret"), func() time.Time { return time.Unix(1700000000, 0) }),
		fileReferenceTTL:     time.Hour,
	}

	_, err := r.UploadProfilePhotoFile(context.Background(), &media.TLMediaUploadProfilePhotoFile{
		OwnerId:      9,
		Video:        tg.MakeTLInputFile(&tg.TLInputFile{Id: 12, Parts: 1, Name: "profile.mp4", Md5Checksum: "md5"}),
		VideoStartTs: &startTs,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dfsClient.uploadProfileReq != nil {
		t.Fatalf("expected default profile path not to call legacy DFS profile upload, got %#v", dfsClient.uploadProfileReq)
	}
	if dfsClient.commitReq == nil || dfsClient.commitReq.OwnerId != 9 || dfsClient.commitReq.Purpose != "media_profile_video" {
		t.Fatalf("expected dfs commit request for profile video, got %#v", dfsClient.commitReq)
	}
	if processorClient.mp4Req == nil || processorClient.mp4Req.ObjectId != "profile-video-object" || processorClient.mp4Req.FileName != "profile.mp4" {
		t.Fatalf("expected mediaprocessor mp4 request, got %#v", processorClient.mp4Req)
	}
	if len(photos.inserted) != 1 || !photos.inserted[0].HasVideo {
		t.Fatalf("expected saved profile photo with video flag, got %#v", photos.inserted)
	}
	if len(photoSizes.inserted) != 1 || photoSizes.inserted[0].FilePath != "gif-thumb-object" {
		t.Fatalf("expected saved profile cover photo size, got %#v", photoSizes.inserted)
	}
	if len(videos.inserted) != 1 || videos.inserted[0].VideoSizeId != photos.inserted[0].PhotoId || videos.inserted[0].Width != 640 || videos.inserted[0].Height != 360 || videos.inserted[0].FileSize != 4096 || videos.inserted[0].VideoStartTs != startTs {
		t.Fatalf("expected saved video size for generated profile photo, got %#v", videos.inserted)
	}
}

func TestUploadProfilePhotoFileCommitsAndProcessesImage(t *testing.T) {
	photos := &capturePhotosModel{}
	photoSizes := &capturePhotoSizesModel{}
	dfsClient := &fakeDfsMediaClient{finalized: dfsapi.MakeTLFileFinalizedObject(&dfsapi.TLFileFinalizedObject{
		ObjectId:        "profile-photo-object",
		UploadSessionId: "ext:9:12:1",
		ReadLease:       []byte("read-lease"),
		DcId:            2,
	}).ToFileFinalizedObject()}
	processorClient := &fakeMediaProcessorClient{
		photo: mediaprocessor.MakeTLProcessedPhoto(&mediaprocessor.TLProcessedPhoto{
			OriginalObjectId: "profile-photo-object",
			Sizes: []mediaprocessor.ProcessorDerivativeClazz{
				mediaprocessor.MakeTLProcessorDerivative(&mediaprocessor.TLProcessorDerivative{Kind: "photo_size", ObjectId: "profile-derivative-m", FileName: "m_profile.jpg", Width: 320, Height: 240, Size2: 1000}),
			},
		}).ToProcessedPhoto(),
	}
	r := &Repository{
		model:                &model.Models{PhotosModel: photos, PhotoSizesModel: photoSizes, VideoSizesModel: &captureVideoSizesModel{}},
		dfsClient:            dfsClient,
		processorClient:      processorClient,
		fileReferenceService: NewFileReferenceService([]byte("test-secret"), func() time.Time { return time.Unix(1700000000, 0) }),
		fileReferenceTTL:     time.Hour,
	}

	_, err := r.UploadProfilePhotoFile(context.Background(), &media.TLMediaUploadProfilePhotoFile{
		OwnerId: 9,
		File:    tg.MakeTLInputFile(&tg.TLInputFile{Id: 12, Parts: 1, Name: "profile.jpg", Md5Checksum: "md5"}),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dfsClient.uploadProfileReq != nil {
		t.Fatalf("expected default profile path not to call legacy DFS profile upload, got %#v", dfsClient.uploadProfileReq)
	}
	if dfsClient.commitReq == nil || dfsClient.commitReq.OwnerId != 9 || dfsClient.commitReq.Purpose != "media_profile_photo" {
		t.Fatalf("expected dfs commit request for profile photo, got %#v", dfsClient.commitReq)
	}
	if processorClient.photoReq == nil || processorClient.photoReq.ObjectId != "profile-photo-object" || string(processorClient.photoReq.ReadLease) != "read-lease" || processorClient.photoReq.FileName != "profile.jpg" || processorClient.photoReq.Profile != tg.BoolTrueClazz {
		t.Fatalf("expected mediaprocessor profile photo request, got %#v", processorClient.photoReq)
	}
	if len(photos.inserted) != 1 || photos.inserted[0].InputFileName != "profile.jpg" {
		t.Fatalf("expected saved profile photo row, got %#v", photos.inserted)
	}
	if len(photoSizes.inserted) != 1 || photoSizes.inserted[0].FilePath != "profile-derivative-m" {
		t.Fatalf("expected saved profile derivative object id, got %#v", photoSizes.inserted)
	}
}

func TestGetPhotoLoadsSizes(t *testing.T) {
	photos := &capturePhotosModel{found: &model.Photos{PhotoId: 303, AccessHash: 404, SizeId: 303, VideoSizeId: 303, DcId: 2, Date2: 5}}
	photoSizes := &capturePhotoSizesModel{byID: []model.PhotoSizes{
		{PhotoSizeId: 303, SizeType: "m", Width: 320, Height: 240, FileSize: 1000},
		{PhotoSizeId: 303, SizeType: "x", Width: 800, Height: 600, FileSize: 1200, CachedType: 4, CachedBytes: "[400,900,1200]"},
	}}
	videoSizes := &captureVideoSizesModel{byID: []model.VideoSizes{{VideoSizeId: 303, SizeType: "v", Width: 320, Height: 240, FileSize: 2000, VideoStartTs: 1.5}}}
	r := &Repository{
		model:                &model.Models{PhotosModel: photos, PhotoSizesModel: photoSizes, VideoSizesModel: videoSizes},
		fileReferenceService: NewFileReferenceService([]byte("test-secret"), func() time.Time { return time.Unix(1700000000, 0) }),
		fileReferenceTTL:     time.Hour,
	}

	got, err := r.GetPhoto(context.Background(), 303)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	do, ok := got.ToPhoto()
	if !ok {
		t.Fatalf("expected photo, got %#v", got)
	}
	if len(do.Sizes) != 2 || len(do.VideoSizes) != 1 {
		t.Fatalf("expected two photo sizes and one video size, got %#v %#v", do.Sizes, do.VideoSizes)
	}
	if progressive, ok := do.Sizes[1].(*tg.TLPhotoSizeProgressive); !ok || progressive.Type != "x" || progressive.W != 800 || progressive.H != 600 || len(progressive.Sizes) != 3 || progressive.Sizes[2] != 1200 {
		t.Fatalf("expected cached_type=4 to reload as progressive photo size, got %#v", do.Sizes[1])
	}
	claims, err := r.fileReferenceService.Validate(do.FileReference)
	if err != nil {
		t.Fatalf("expected valid loaded photo file_reference: %v", err)
	}
	if claims.OriginDomain != "photo" || claims.MediaID != 303 || claims.AccessHash != 404 {
		t.Fatalf("unexpected loaded photo file_reference claims: %#v", claims)
	}
}

func TestUploadedProfilePhotoCallsDfsAndSavesNewPhoto(t *testing.T) {
	photos := &capturePhotosModel{}
	dfsClient := &fakeDfsMediaClient{photo: testPhotoWithSizes(404, false)}
	r := &Repository{
		model:     &model.Models{PhotosModel: photos, PhotoSizesModel: &capturePhotoSizesModel{}, VideoSizesModel: &captureVideoSizesModel{}},
		dfsClient: dfsClient,
	}

	_, err := r.UploadedProfilePhoto(context.Background(), &media.TLMediaUploadedProfilePhoto{OwnerId: 5, PhotoId: 44})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dfsClient.uploadedProfileReq == nil || dfsClient.uploadedProfileReq.PhotoId != 44 {
		t.Fatalf("expected uploaded profile dfs request for photo 44, got %#v", dfsClient.uploadedProfileReq)
	}
	if len(photos.inserted) != 1 || photos.inserted[0].PhotoId != 404 {
		t.Fatalf("expected saved new photo 404, got %#v", photos.inserted)
	}
}

func testPhotoWithSizes(id int64, withVideo bool) *tg.Photo {
	photo := tg.MakeTLPhoto(&tg.TLPhoto{
		Id:         id,
		AccessHash: 99,
		Date:       10,
		DcId:       1,
		Sizes: []tg.PhotoSizeClazz{
			tg.MakeTLPhotoSize(&tg.TLPhotoSize{Type: "m", W: 320, H: 240, Size2: 1000}),
		},
	}).ToPhoto()
	if withVideo {
		do, _ := photo.ToPhoto()
		do.VideoSizes = []tg.VideoSizeClazz{
			tg.MakeTLVideoSize(&tg.TLVideoSize{Type: "v", W: 320, H: 240, Size2: 2000}),
		}
	}
	return photo
}
