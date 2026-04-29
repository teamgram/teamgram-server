package core

import (
	"context"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository"
	minioadapter "github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/minio"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakePhotoRepository struct {
	nextPhotoID int64
	photos      map[string][]byte
	videos      map[string][]byte
	mp4         []byte
	frame       []byte
}

func newFakePhotoRepository() *fakePhotoRepository {
	return &fakePhotoRepository{
		nextPhotoID: 9001,
		photos:      make(map[string][]byte),
		videos:      make(map[string][]byte),
		mp4:         []byte("mp4"),
		frame:       []byte("frame-jpeg"),
	}
}

func (f *fakePhotoRepository) NextPhotoID(context.Context) (int64, error) {
	id := f.nextPhotoID
	f.nextPhotoID++
	return id, nil
}

func (f *fakePhotoRepository) LoadOriginalPhotoBytes(_ context.Context, photoID int64) ([]byte, error) {
	path := "0/" + minioadapter.ObjectPath(photoID)
	data := f.photos[path]
	if data == nil {
		return nil, dfs.ErrDfsFileNotFound
	}
	return append([]byte(nil), data...), nil
}

func (f *fakePhotoRepository) SaveProfileVideoObject(_ context.Context, photoID int64, data []byte) (int64, error) {
	path := "v/" + minioadapter.ObjectPath(photoID)
	f.videos[path] = append([]byte(nil), data...)
	return int64(len(data)), nil
}

func (f *fakePhotoRepository) ConvertToMP4(context.Context, []byte) ([]byte, error) {
	return f.mp4, nil
}

func (f *fakePhotoRepository) ExtractFirstFrame(context.Context, []byte) ([]byte, error) {
	return f.frame, nil
}

func (f *fakePhotoRepository) SavePhotoObjects(_ context.Context, photoID int64, original []byte, _ string, _ bool, storeOriginal bool) (*repository.StoredPhoto, error) {
	if storeOriginal {
		f.photos["0/"+minioadapter.ObjectPath(photoID)] = append([]byte(nil), original...)
	}
	f.photos["s/"+minioadapter.ObjectPath(photoID)] = []byte("small")
	f.photos["m/"+minioadapter.ObjectPath(photoID)] = []byte("medium")
	return &repository.StoredPhoto{
		ID: photoID,
		Sizes: []repository.StoredPhotoSize{
			{Type: "s", W: 90, H: 60, Size: int32(len("small"))},
			{Type: "m", W: 320, H: 200, Size: int32(len("medium"))},
		},
	}, nil
}

func TestDfsUploadPhotoFileV2StoresOriginalAndSizes(t *testing.T) {
	core, uploads, photos := newPhotoTestCore(t)
	writeUploadedTestFile(t, uploads, 1001, 2002, []byte("photo-bytes"))

	out, err := core.DfsUploadPhotoFileV2(&dfs.TLDfsUploadPhotoFileV2{
		Creator: 1001,
		File: tg.MakeTLInputFile(&tg.TLInputFile{
			Id:          2002,
			Parts:       1,
			Name:        "avatar.jpg",
			Md5Checksum: "571c58e834fd876178aca15d610f4512",
		}),
	})
	if err != nil {
		t.Fatalf("DfsUploadPhotoFileV2() error = %v", err)
	}
	photo, ok := out.ToPhoto()
	if !ok {
		t.Fatalf("Photo clazz = %s, want photo", out.ClazzName())
	}
	if photo.Id != 9001 {
		t.Fatalf("photo.Id = %d, want 9001", photo.Id)
	}
	if got := minioadapter.StorageTypeFromAccessHash(photo.AccessHash); got != int32(tg.ClazzID_storage_fileJpeg) {
		t.Fatalf("storage type = %#x, want jpeg", uint32(got))
	}
	if string(photos.photos["0/9001.dat"]) != "photo-bytes" {
		t.Fatalf("original = %q, want photo-bytes", photos.photos["0/9001.dat"])
	}
	if string(photos.photos["s/9001.dat"]) != "small" || string(photos.photos["m/9001.dat"]) != "medium" {
		t.Fatalf("size objects = %#v", photos.photos)
	}
	if len(photo.Sizes) != 2 {
		t.Fatalf("len(photo.Sizes) = %d, want 2", len(photo.Sizes))
	}
}

func TestDfsUploadPhotoFileV2RejectsChecksumMismatch(t *testing.T) {
	core, uploads, _ := newPhotoTestCore(t)
	writeUploadedTestFile(t, uploads, 1001, 2002, []byte("photo-bytes"))

	_, err := core.DfsUploadPhotoFileV2(&dfs.TLDfsUploadPhotoFileV2{
		Creator: 1001,
		File: tg.MakeTLInputFile(&tg.TLInputFile{
			Id:          2002,
			Parts:       1,
			Name:        "avatar.jpg",
			Md5Checksum: "bad",
		}),
	})
	if !errors.Is(err, dfs.ErrDfsChecksumInvalid) {
		t.Fatalf("DfsUploadPhotoFileV2() error = %v, want ErrDfsChecksumInvalid", err)
	}
}

func TestDfsUploadProfilePhotoFileV2RequiresFileOrVideo(t *testing.T) {
	core, _, _ := newPhotoTestCore(t)
	_, err := core.DfsUploadProfilePhotoFileV2(&dfs.TLDfsUploadProfilePhotoFileV2{Creator: 1001})
	if !errors.Is(err, dfs.ErrDfsInvalidArgument) {
		t.Fatalf("DfsUploadProfilePhotoFileV2() error = %v, want ErrDfsInvalidArgument", err)
	}
}

func TestDfsUploadedProfilePhotoReadsOriginalObject(t *testing.T) {
	core, _, photos := newPhotoTestCore(t)
	photos.nextPhotoID = 9100
	photos.photos["0/3003.dat"] = []byte("original")

	out, err := core.DfsUploadedProfilePhoto(&dfs.TLDfsUploadedProfilePhoto{Creator: 1001, PhotoId: 3003})
	if err != nil {
		t.Fatalf("DfsUploadedProfilePhoto() error = %v", err)
	}
	photo, ok := out.ToPhoto()
	if !ok {
		t.Fatalf("Photo clazz = %s, want photo", out.ClazzName())
	}
	if photo.Id != 9100 {
		t.Fatalf("photo.Id = %d, want 9100", photo.Id)
	}
	if _, ok := photos.photos["0/9100.dat"]; ok {
		t.Fatal("uploadedProfilePhoto wrote a new original object; master only writes resized profile sizes")
	}
	if string(photos.photos["s/9100.dat"]) != "small" {
		t.Fatalf("profile size object = %q, want small", photos.photos["s/9100.dat"])
	}
}

func newPhotoTestCore(t *testing.T) (*DfsCore, *UploadSessionManager, *fakePhotoRepository) {
	t.Helper()
	uploadRepo := newFakeUploadStateRepository()
	uploads := NewUploadSessionManager(uploadRepo)
	photos := newFakePhotoRepository()
	return &DfsCore{
		ctx:                  context.Background(),
		uploadSessionManager: uploads,
		photoRepository:      photos,
	}, uploads, photos
}

func writeUploadedTestFile(t *testing.T, manager *UploadSessionManager, creator, fileID int64, data []byte) {
	t.Helper()
	totalParts := int32(1)
	if err := manager.WritePart(context.Background(), WritePartCommand{
		Creator:        creator,
		FileID:         fileID,
		FilePart:       0,
		Bytes:          data,
		FileTotalParts: &totalParts,
	}); err != nil {
		t.Fatalf("WritePart() error = %v", err)
	}
}
