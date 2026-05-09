package repository

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/media/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/filelease"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestLocatorDocumentOriginalResolvesStoredObjectAndValidatesFileReference(t *testing.T) {
	now := time.Unix(1700000000, 0)
	repo := testLocatorRepository(now)
	repo.model.DocumentsModel = &captureDocumentsModel{byID: map[int64]*model.Documents{
		100: {
			DocumentId: 100,
			AccessHash: 200,
			DcId:       2,
			FilePath:   "doc-original-object",
			FileSize:   321,
			MimeType:   "application/pdf",
		},
	}}
	ref := testFileReference(t, repo, FileReferenceClaims{
		MediaID:      100,
		ObjectID:     "doc-original-object",
		OriginDomain: "document",
		OriginID:     7,
		AccessHash:   200,
		ExpireAt:     now.Add(time.Hour).Unix(),
	})

	resolved, err := repo.ResolveFileLocation(context.Background(), &media.TLMediaResolveFileLocation{
		Location: tg.MakeTLInputDocumentFileLocation(&tg.TLInputDocumentFileLocation{
			Id:            100,
			AccessHash:    200,
			FileReference: ref,
		}),
		ViewerId: 7,
	})
	if err != nil {
		t.Fatalf("ResolveFileLocation() error = %v", err)
	}
	if resolved.ObjectId != "doc-original-object" || resolved.Size2 != 321 || resolved.MimeType != "application/pdf" || resolved.DcId != 2 {
		t.Fatalf("resolved object = %#v", resolved)
	}
	claims, err := filelease.Verify("read-secret", resolved.ReadLease, now)
	if err != nil {
		t.Fatalf("read lease verify failed: %v", err)
	}
	if claims.Bucket != "documents" || claims.Key != "objects/doc-original-object.dat" || claims.ObjectID != "doc-original-object" {
		t.Fatalf("read lease claims = %#v", claims)
	}
	if uint32(claims.StorageType) != 0xae1e508d {
		t.Fatalf("storage type = %#x, want pdf", uint32(claims.StorageType))
	}
}

func TestLocatorDocumentOriginalStorageTypeMatchesDfsFileMapping(t *testing.T) {
	now := time.Unix(1700000000, 0)
	tests := []struct {
		name     string
		fileName string
		mimeType string
		wantType uint32
	}{
		{name: "mp3 mime with parameter", fileName: "audio.bin", mimeType: "audio/mpeg; charset=binary", wantType: tg.ClazzID_storage_fileMp3},
		{name: "mov filename fallback", fileName: "clip.mov", mimeType: "application/octet-stream", wantType: tg.ClazzID_storage_fileMov},
		{name: "unknown", fileName: "archive.bin", mimeType: "application/octet-stream", wantType: tg.ClazzID_storage_fileUnknown},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := testLocatorRepository(now)
			repo.model.DocumentsModel = &captureDocumentsModel{byID: map[int64]*model.Documents{
				100: {
					DocumentId:       100,
					AccessHash:       200,
					DcId:             2,
					FilePath:         "doc-original-object",
					FileSize:         321,
					MimeType:         tt.mimeType,
					UploadedFileName: tt.fileName,
				},
			}}
			ref := testFileReference(t, repo, FileReferenceClaims{
				MediaID:      100,
				ObjectID:     "doc-original-object",
				OriginDomain: "document",
				OriginID:     7,
				AccessHash:   200,
				ExpireAt:     now.Add(time.Hour).Unix(),
			})

			resolved, err := repo.ResolveFileLocation(context.Background(), &media.TLMediaResolveFileLocation{
				Location: tg.MakeTLInputDocumentFileLocation(&tg.TLInputDocumentFileLocation{
					Id:            100,
					AccessHash:    200,
					FileReference: ref,
				}),
				ViewerId: 7,
			})
			if err != nil {
				t.Fatalf("ResolveFileLocation() error = %v", err)
			}
			if uint32(resolved.StorageFileType) != tt.wantType {
				t.Fatalf("StorageFileType = %#x, want %#x", uint32(resolved.StorageFileType), tt.wantType)
			}
		})
	}
}

func TestLocatorDocumentThumbResolvesPhotoSizeObject(t *testing.T) {
	now := time.Unix(1700000000, 0)
	repo := testLocatorRepository(now)
	repo.model.DocumentsModel = &captureDocumentsModel{byID: map[int64]*model.Documents{
		101: {DocumentId: 101, AccessHash: 201, DcId: 3, FilePath: "doc-object", FileSize: 1000, MimeType: "video/mp4", ThumbId: 901},
	}}
	repo.model.PhotoSizesModel = &capturePhotoSizesModel{byID: []model.PhotoSizes{
		{PhotoSizeId: 901, SizeType: "m", Width: 320, Height: 240, FileSize: 111, FilePath: "doc-thumb-object"},
	}}
	ref := testFileReference(t, repo, FileReferenceClaims{
		MediaID:      101,
		ObjectID:     "doc-object",
		OriginDomain: "document",
		OriginID:     7,
		AccessHash:   201,
		ExpireAt:     now.Add(time.Hour).Unix(),
	})

	resolved, err := repo.ResolveFileLocation(context.Background(), &media.TLMediaResolveFileLocation{
		Location: tg.MakeTLInputDocumentFileLocation(&tg.TLInputDocumentFileLocation{
			Id:            101,
			AccessHash:    201,
			FileReference: ref,
			ThumbSize:     "m",
		}),
		ViewerId: 7,
	})
	if err != nil {
		t.Fatalf("ResolveFileLocation() error = %v", err)
	}
	if resolved.ObjectId != "doc-thumb-object" || resolved.Size2 != 111 || resolved.MimeType != "image/jpeg" {
		t.Fatalf("resolved thumb = %#v", resolved)
	}
}

func TestLocatorPhotoThumbResolvesPhotoSizeObjectAndValidatesReference(t *testing.T) {
	now := time.Unix(1700000000, 0)
	repo := testLocatorRepository(now)
	repo.model.PhotosModel = &capturePhotosModel{found: &model.Photos{PhotoId: 300, AccessHash: 400, DcId: 4, SizeId: 300}}
	repo.model.PhotoSizesModel = &capturePhotoSizesModel{byID: []model.PhotoSizes{
		{PhotoSizeId: 300, SizeType: "x", Width: 800, Height: 600, FileSize: 222, FilePath: "photo-x-object"},
	}}
	ref := testFileReference(t, repo, FileReferenceClaims{
		MediaID:      300,
		ObjectID:     "photo-original-object",
		OriginDomain: "photo",
		OriginID:     8,
		AccessHash:   400,
		ExpireAt:     now.Add(time.Hour).Unix(),
	})

	resolved, err := repo.ResolveFileLocation(context.Background(), &media.TLMediaResolveFileLocation{
		Location: tg.MakeTLInputPhotoFileLocation(&tg.TLInputPhotoFileLocation{
			Id:            300,
			AccessHash:    400,
			FileReference: ref,
			ThumbSize:     "x",
		}),
		ViewerId: 8,
	})
	if err != nil {
		t.Fatalf("ResolveFileLocation() error = %v", err)
	}
	if resolved.ObjectId != "photo-x-object" || resolved.StorageFileType != int32(0x7efe0e) {
		t.Fatalf("resolved photo = %#v", resolved)
	}
}

func TestLocatorPeerPhotoUsesPhotoSizeMetadataInsteadOfLegacyPath(t *testing.T) {
	now := time.Unix(1700000000, 0)
	repo := testLocatorRepository(now)
	repo.model.PhotosModel = &capturePhotosModel{found: &model.Photos{PhotoId: 500, AccessHash: 600, DcId: 5, SizeId: 500}}
	repo.model.PhotoSizesModel = &capturePhotoSizesModel{byID: []model.PhotoSizes{
		{PhotoSizeId: 500, SizeType: "s", Width: 120, Height: 120, FileSize: 50, FilePath: "small-profile-object"},
		{PhotoSizeId: 500, SizeType: "x", Width: 640, Height: 640, FileSize: 150, FilePath: "big-profile-object"},
	}}

	resolved, err := repo.ResolveFileLocation(context.Background(), &media.TLMediaResolveFileLocation{
		Location: tg.MakeTLInputPeerPhotoFileLocation(&tg.TLInputPeerPhotoFileLocation{
			Big:     true,
			Peer:    tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 8, AccessHash: 9}),
			PhotoId: 500,
		}),
		ViewerId: 8,
	})
	if err != nil {
		t.Fatalf("ResolveFileLocation() error = %v", err)
	}
	if resolved.ObjectId != "big-profile-object" {
		t.Fatalf("resolved object = %q, want metadata object", resolved.ObjectId)
	}
	if resolved.ObjectId == documentObjectPath(500) {
		t.Fatalf("resolved legacy path guess %q", resolved.ObjectId)
	}
}

func TestLocatorFileReferenceErrorsRemainSemantic(t *testing.T) {
	now := time.Unix(1700000000, 0)
	repo := testLocatorRepository(now)
	_, err := repo.ResolveFileLocation(context.Background(), &media.TLMediaResolveFileLocation{
		Location: tg.MakeTLInputDocumentFileLocation(&tg.TLInputDocumentFileLocation{Id: 100, AccessHash: 200}),
		ViewerId: 7,
	})
	if !errors.Is(err, media.ErrFileReferenceEmpty) {
		t.Fatalf("empty file_reference error = %v, want ErrFileReferenceEmpty", err)
	}

	expired := testFileReference(t, repo, FileReferenceClaims{
		MediaID:      100,
		ObjectID:     "doc-original-object",
		OriginDomain: "document",
		AccessHash:   200,
		ExpireAt:     now.Add(-time.Second).Unix(),
	})
	_, err = repo.ResolveFileLocation(context.Background(), &media.TLMediaResolveFileLocation{
		Location: tg.MakeTLInputDocumentFileLocation(&tg.TLInputDocumentFileLocation{Id: 100, AccessHash: 200, FileReference: expired}),
		ViewerId: 7,
	})
	if !errors.Is(err, media.ErrFileReferenceExpired) {
		t.Fatalf("expired file_reference error = %v, want ErrFileReferenceExpired", err)
	}
}

func testLocatorRepository(now time.Time) *Repository {
	return &Repository{
		model:                &model.Models{},
		fileReferenceService: NewFileReferenceService([]byte("file-ref-secret"), func() time.Time { return now }),
		readLeaseSecret:      []byte("read-secret"),
		readLeaseTTL:         time.Minute,
	}
}

func testFileReference(t *testing.T, r *Repository, claims FileReferenceClaims) []byte {
	t.Helper()
	token, err := r.fileReferenceService.Generate(claims)
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}
	return token
}
