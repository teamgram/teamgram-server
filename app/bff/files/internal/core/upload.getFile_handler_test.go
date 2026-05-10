package core

import (
	"context"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/files/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/bff/files/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/bff/files/internal/svc"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	mediapb "github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestUploadGetFileUsesMediaLocatorAndReadLease(t *testing.T) {
	location := tg.MakeTLInputDocumentFileLocation(&tg.TLInputDocumentFileLocation{
		Id:            1001,
		AccessHash:    2002,
		FileReference: []byte("file-reference"),
	})
	dfsClient := &fakeFilesDfsClient{
		readResp: tg.MakeTLUploadFile(&tg.TLUploadFile{
			Type:  tg.MakeTLStorageFileUnknown(&tg.TLStorageFileUnknown{}),
			Mtime: 1,
			Bytes: []byte("part"),
		}).ToUploadFile(),
	}
	mediaClient := &fakeFilesMediaClient{
		resolveResp: mediapb.MakeTLMediaResolvedFileObject(&mediapb.TLMediaResolvedFileObject{
			ObjectId:        "object-1001",
			ReadLease:       []byte("read-lease"),
			Size2:           4,
			MimeType:        "application/octet-stream",
			DcId:            1,
			StorageFileType: testStorageFileType(tg.ClazzID_storage_fileUnknown),
		}).ToMediaResolvedFileObject(),
	}
	core := newUploadGetFileTestCore(dfsClient, mediaClient, false)

	got, err := core.UploadGetFile(&tg.TLUploadGetFile{
		Location: location,
		Offset:   11,
		Limit:    22,
	})
	if err != nil {
		t.Fatalf("UploadGetFile() error = %v", err)
	}
	if mediaClient.resolveReq == nil || mediaClient.resolveReq.Location != location || mediaClient.resolveReq.ViewerId != 1001 {
		t.Fatalf("resolve request = %#v", mediaClient.resolveReq)
	}
	if dfsClient.readReq == nil || string(dfsClient.readReq.ReadLease) != "read-lease" || dfsClient.readReq.Offset != 11 || dfsClient.readReq.Limit != 22 {
		t.Fatalf("read request = %#v", dfsClient.readReq)
	}
	if dfsClient.downloadReq != nil {
		t.Fatalf("legacy DfsDownloadFile called: %#v", dfsClient.downloadReq)
	}
	file, ok := got.ToUploadFile()
	if !ok || file.Type == nil || file.Type.StorageFileTypeClazzName() != tg.ClazzName_storage_filePartial {
		t.Fatalf("upload file type = %#v, want storage_filePartial", got)
	}
}

func TestUploadGetFileResolvesPhotoAndPeerPhotoLocations(t *testing.T) {
	tests := []struct {
		name     string
		location tg.InputFileLocationClazz
	}{
		{
			name: "photo",
			location: tg.MakeTLInputPhotoFileLocation(&tg.TLInputPhotoFileLocation{
				Id:            1001,
				AccessHash:    2002,
				FileReference: []byte("file-reference"),
				ThumbSize:     "m",
			}),
		},
		{
			name: "peer photo",
			location: tg.MakeTLInputPeerPhotoFileLocation(&tg.TLInputPeerPhotoFileLocation{
				Big:     true,
				PhotoId: 3003,
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dfsClient := &fakeFilesDfsClient{
				readResp: tg.MakeTLUploadFile(&tg.TLUploadFile{
					Type:  tg.MakeTLStorageFilePartial(&tg.TLStorageFilePartial{}),
					Mtime: 1,
					Bytes: []byte("part"),
				}).ToUploadFile(),
			}
			mediaClient := &fakeFilesMediaClient{
				resolveResp: mediapb.MakeTLMediaResolvedFileObject(&mediapb.TLMediaResolvedFileObject{
					ObjectId:        "object-photo",
					ReadLease:       []byte("read-lease"),
					Size2:           4,
					MimeType:        "image/jpeg",
					DcId:            1,
					StorageFileType: testStorageFileType(tg.ClazzID_storage_fileJpeg),
				}).ToMediaResolvedFileObject(),
			}

			_, err := newUploadGetFileTestCore(dfsClient, mediaClient, false).UploadGetFile(&tg.TLUploadGetFile{
				Location: tt.location,
				Offset:   7,
				Limit:    8,
			})
			if err != nil {
				t.Fatalf("UploadGetFile() error = %v", err)
			}
			if mediaClient.resolveReq == nil || mediaClient.resolveReq.Location != tt.location {
				t.Fatalf("resolve request = %#v", mediaClient.resolveReq)
			}
			if dfsClient.readReq == nil || string(dfsClient.readReq.ReadLease) != "read-lease" || dfsClient.readReq.Offset != 7 || dfsClient.readReq.Limit != 8 {
				t.Fatalf("read request = %#v", dfsClient.readReq)
			}
			if dfsClient.downloadReq != nil {
				t.Fatalf("legacy DfsDownloadFile called: %#v", dfsClient.downloadReq)
			}
		})
	}
}

func TestUploadGetFileLegacyFallbackIsConfigGated(t *testing.T) {
	location := tg.MakeTLInputEncryptedFileLocation(&tg.TLInputEncryptedFileLocation{Id: 1001, AccessHash: 2002})
	dfsClient := &fakeFilesDfsClient{
		downloadResp: tg.MakeTLUploadFile(&tg.TLUploadFile{
			Type:  tg.MakeTLStorageFilePartial(&tg.TLStorageFilePartial{}),
			Mtime: 1,
			Bytes: []byte("legacy"),
		}).ToUploadFile(),
	}
	mediaClient := &fakeFilesMediaClient{}

	_, err := newUploadGetFileTestCore(dfsClient, mediaClient, false).UploadGetFile(&tg.TLUploadGetFile{Location: location})
	if !errors.Is(err, tg.ErrLocationInvalid) && err != tg.ErrLocationInvalid {
		t.Fatalf("UploadGetFile() error = %v, want LOCATION_INVALID", err)
	}
	if dfsClient.downloadReq != nil || mediaClient.resolveReq != nil {
		t.Fatalf("fallback disabled calls: download=%#v resolve=%#v", dfsClient.downloadReq, mediaClient.resolveReq)
	}

	got, err := newUploadGetFileTestCore(dfsClient, mediaClient, true).UploadGetFile(&tg.TLUploadGetFile{
		Location: location,
		Offset:   3,
		Limit:    4,
	})
	if err != nil {
		t.Fatalf("UploadGetFile(legacy) error = %v", err)
	}
	if got == nil || dfsClient.downloadReq == nil || dfsClient.downloadReq.Location != location || dfsClient.downloadReq.Offset != 3 || dfsClient.downloadReq.Limit != 4 {
		t.Fatalf("legacy download request = %#v got=%#v", dfsClient.downloadReq, got)
	}
}

func TestUploadGetFileMapsMediaReferenceErrors(t *testing.T) {
	location := tg.MakeTLInputPhotoFileLocation(&tg.TLInputPhotoFileLocation{
		Id:            1001,
		AccessHash:    2002,
		FileReference: []byte("expired-reference"),
		ThumbSize:     "m",
	})
	core := newUploadGetFileTestCore(&fakeFilesDfsClient{}, &fakeFilesMediaClient{resolveErr: mediapb.ErrFileReferenceExpired}, false)

	_, err := core.UploadGetFile(&tg.TLUploadGetFile{Location: location})
	if !errors.Is(err, tg.ErrFileReferenceExpired) && err != tg.ErrFileReferenceExpired {
		t.Fatalf("UploadGetFile() error = %v, want FILE_REFERENCE_EXPIRED", err)
	}
}

func newUploadGetFileTestCore(dfsClient *fakeFilesDfsClient, mediaClient *fakeFilesMediaClient, legacyFallback bool) *FilesCore {
	return &FilesCore{
		ctx: context.Background(),
		svcCtx: &svc.ServiceContext{
			Config: config.Config{Download: config.DownloadConf{LegacyDownloadFallback: legacyFallback}},
			Repo: &repository.Repository{
				DfsClient:   dfsClient,
				MediaClient: mediaClient,
			},
		},
		MD: &metadata.RpcMetadata{UserId: 1001},
	}
}

func testStorageFileType(clazzID uint32) int32 {
	return int32(clazzID)
}

type fakeFilesDfsClient struct {
	readReq      *dfs.TLDfsGetFileByReadLease
	readResp     *tg.UploadFile
	readErr      error
	hashReq      *dfs.TLDfsGetFileHashesByReadLease
	hashResp     *dfs.VectorFileHash
	hashErr      error
	downloadReq  *dfs.TLDfsDownloadFile
	downloadResp *tg.UploadFile
	downloadErr  error
}

func (f *fakeFilesDfsClient) DfsCommitUpload(context.Context, *dfs.TLDfsCommitUpload) (*dfs.FileFinalizedObject, error) {
	return nil, errors.New("unexpected DfsCommitUpload")
}

func (f *fakeFilesDfsClient) DfsPutFile(context.Context, *dfs.TLDfsPutFile) (*dfs.FileFinalizedObject, error) {
	return nil, errors.New("unexpected DfsPutFile")
}

func (f *fakeFilesDfsClient) DfsGetFileByReadLease(_ context.Context, in *dfs.TLDfsGetFileByReadLease) (*tg.UploadFile, error) {
	f.readReq = in
	return f.readResp, f.readErr
}

func (f *fakeFilesDfsClient) DfsGetFileHashesByReadLease(_ context.Context, in *dfs.TLDfsGetFileHashesByReadLease) (*dfs.VectorFileHash, error) {
	f.hashReq = in
	return f.hashResp, f.hashErr
}

func (f *fakeFilesDfsClient) DfsWriteFilePartData(context.Context, *dfs.TLDfsWriteFilePartData) (*tg.Bool, error) {
	return tg.BoolTrue, nil
}

func (f *fakeFilesDfsClient) DfsUploadPhotoFileV2(context.Context, *dfs.TLDfsUploadPhotoFileV2) (*tg.Photo, error) {
	return nil, errors.New("unexpected DfsUploadPhotoFileV2")
}

func (f *fakeFilesDfsClient) DfsUploadProfilePhotoFileV2(context.Context, *dfs.TLDfsUploadProfilePhotoFileV2) (*tg.Photo, error) {
	return nil, errors.New("unexpected DfsUploadProfilePhotoFileV2")
}

func (f *fakeFilesDfsClient) DfsUploadEncryptedFileV2(context.Context, *dfs.TLDfsUploadEncryptedFileV2) (*tg.EncryptedFile, error) {
	return nil, errors.New("unexpected DfsUploadEncryptedFileV2")
}

func (f *fakeFilesDfsClient) DfsDownloadFile(_ context.Context, in *dfs.TLDfsDownloadFile) (*tg.UploadFile, error) {
	f.downloadReq = in
	return f.downloadResp, f.downloadErr
}

func (f *fakeFilesDfsClient) DfsUploadDocumentFileV2(context.Context, *dfs.TLDfsUploadDocumentFileV2) (*tg.Document, error) {
	return nil, errors.New("unexpected DfsUploadDocumentFileV2")
}

func (f *fakeFilesDfsClient) DfsUploadGifDocumentMedia(context.Context, *dfs.TLDfsUploadGifDocumentMedia) (*tg.Document, error) {
	return nil, errors.New("unexpected DfsUploadGifDocumentMedia")
}

func (f *fakeFilesDfsClient) DfsUploadMp4DocumentMedia(context.Context, *dfs.TLDfsUploadMp4DocumentMedia) (*tg.Document, error) {
	return nil, errors.New("unexpected DfsUploadMp4DocumentMedia")
}

func (f *fakeFilesDfsClient) DfsUploadWallPaperFile(context.Context, *dfs.TLDfsUploadWallPaperFile) (*tg.Document, error) {
	return nil, errors.New("unexpected DfsUploadWallPaperFile")
}

func (f *fakeFilesDfsClient) DfsUploadThemeFile(context.Context, *dfs.TLDfsUploadThemeFile) (*tg.Document, error) {
	return nil, errors.New("unexpected DfsUploadThemeFile")
}

func (f *fakeFilesDfsClient) DfsUploadRingtoneFile(context.Context, *dfs.TLDfsUploadRingtoneFile) (*tg.Document, error) {
	return nil, errors.New("unexpected DfsUploadRingtoneFile")
}

func (f *fakeFilesDfsClient) DfsUploadedProfilePhoto(context.Context, *dfs.TLDfsUploadedProfilePhoto) (*tg.Photo, error) {
	return nil, errors.New("unexpected DfsUploadedProfilePhoto")
}

type fakeFilesMediaClient struct {
	resolveReq      *mediapb.TLMediaResolveFileLocation
	resolveResp     *mediapb.MediaResolvedFileObject
	resolveErr      error
	photoReq        *mediapb.TLMediaGetPhoto
	photoResp       *tg.Photo
	photoErr        error
	uploadPhotoReq  *mediapb.TLMediaUploadPhotoFile
	uploadPhotoResp *tg.Photo
	uploadPhotoErr  error
}

func (f *fakeFilesMediaClient) MediaUploadPhotoFile(_ context.Context, in *mediapb.TLMediaUploadPhotoFile) (*tg.Photo, error) {
	f.uploadPhotoReq = in
	return f.uploadPhotoResp, f.uploadPhotoErr
}

func (f *fakeFilesMediaClient) MediaUploadProfilePhotoFile(context.Context, *mediapb.TLMediaUploadProfilePhotoFile) (*tg.Photo, error) {
	return nil, errors.New("unexpected MediaUploadProfilePhotoFile")
}

func (f *fakeFilesMediaClient) MediaGetPhoto(_ context.Context, in *mediapb.TLMediaGetPhoto) (*tg.Photo, error) {
	f.photoReq = in
	return f.photoResp, f.photoErr
}

func (f *fakeFilesMediaClient) MediaGetPhotoSizeList(context.Context, *mediapb.TLMediaGetPhotoSizeList) (*mediapb.PhotoSizeList, error) {
	return nil, errors.New("unexpected MediaGetPhotoSizeList")
}

func (f *fakeFilesMediaClient) MediaGetPhotoSizeListList(context.Context, *mediapb.TLMediaGetPhotoSizeListList) (*mediapb.VectorPhotoSizeList, error) {
	return nil, errors.New("unexpected MediaGetPhotoSizeListList")
}

func (f *fakeFilesMediaClient) MediaGetVideoSizeList(context.Context, *mediapb.TLMediaGetVideoSizeList) (*mediapb.VideoSizeList, error) {
	return nil, errors.New("unexpected MediaGetVideoSizeList")
}

func (f *fakeFilesMediaClient) MediaUploadedDocumentMedia(context.Context, *mediapb.TLMediaUploadedDocumentMedia) (*tg.MessageMedia, error) {
	return nil, errors.New("unexpected MediaUploadedDocumentMedia")
}

func (f *fakeFilesMediaClient) MediaGetDocument(context.Context, *mediapb.TLMediaGetDocument) (*tg.Document, error) {
	return nil, errors.New("unexpected MediaGetDocument")
}

func (f *fakeFilesMediaClient) MediaGetDocumentList(context.Context, *mediapb.TLMediaGetDocumentList) (*mediapb.VectorDocument, error) {
	return nil, errors.New("unexpected MediaGetDocumentList")
}

func (f *fakeFilesMediaClient) MediaUploadEncryptedFile(context.Context, *mediapb.TLMediaUploadEncryptedFile) (*tg.EncryptedFile, error) {
	return nil, errors.New("unexpected MediaUploadEncryptedFile")
}

func (f *fakeFilesMediaClient) MediaGetEncryptedFile(context.Context, *mediapb.TLMediaGetEncryptedFile) (*tg.EncryptedFile, error) {
	return nil, errors.New("unexpected MediaGetEncryptedFile")
}

func (f *fakeFilesMediaClient) MediaUploadWallPaperFile(context.Context, *mediapb.TLMediaUploadWallPaperFile) (*tg.Document, error) {
	return nil, errors.New("unexpected MediaUploadWallPaperFile")
}

func (f *fakeFilesMediaClient) MediaUploadThemeFile(context.Context, *mediapb.TLMediaUploadThemeFile) (*tg.Document, error) {
	return nil, errors.New("unexpected MediaUploadThemeFile")
}

func (f *fakeFilesMediaClient) MediaUploadStickerFile(context.Context, *mediapb.TLMediaUploadStickerFile) (*tg.Document, error) {
	return nil, errors.New("unexpected MediaUploadStickerFile")
}

func (f *fakeFilesMediaClient) MediaUploadRingtoneFile(context.Context, *mediapb.TLMediaUploadRingtoneFile) (*tg.Document, error) {
	return nil, errors.New("unexpected MediaUploadRingtoneFile")
}

func (f *fakeFilesMediaClient) MediaUploadedProfilePhoto(context.Context, *mediapb.TLMediaUploadedProfilePhoto) (*tg.Photo, error) {
	return nil, errors.New("unexpected MediaUploadedProfilePhoto")
}

func (f *fakeFilesMediaClient) MediaResolveFileLocation(_ context.Context, in *mediapb.TLMediaResolveFileLocation) (*mediapb.MediaResolvedFileObject, error) {
	f.resolveReq = in
	return f.resolveResp, f.resolveErr
}
