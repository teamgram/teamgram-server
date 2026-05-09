package rpc

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeDfsClient struct {
	commitReq *dfs.TLDfsCommitUpload
	putReq    *dfs.TLDfsPutFile
	photoReq  *dfs.TLDfsUploadPhotoFileV2
	finalized *dfs.FileFinalizedObject
	photo     *tg.Photo
	err       error
}

func (f *fakeDfsClient) DfsCommitUpload(_ context.Context, in *dfs.TLDfsCommitUpload) (*dfs.FileFinalizedObject, error) {
	f.commitReq = in
	return f.finalized, f.err
}

func (f *fakeDfsClient) DfsPutFile(_ context.Context, in *dfs.TLDfsPutFile) (*dfs.FileFinalizedObject, error) {
	f.putReq = in
	return f.finalized, f.err
}

func (f *fakeDfsClient) DfsGetFileByReadLease(context.Context, *dfs.TLDfsGetFileByReadLease) (*tg.UploadFile, error) {
	return nil, f.err
}

func (f *fakeDfsClient) DfsGetFileHashesByReadLease(context.Context, *dfs.TLDfsGetFileHashesByReadLease) (*dfs.VectorFileHash, error) {
	return nil, f.err
}

func (f *fakeDfsClient) DfsUploadPhotoFileV2(_ context.Context, in *dfs.TLDfsUploadPhotoFileV2) (*tg.Photo, error) {
	f.photoReq = in
	return f.photo, f.err
}

func (f *fakeDfsClient) DfsUploadProfilePhotoFileV2(context.Context, *dfs.TLDfsUploadProfilePhotoFileV2) (*tg.Photo, error) {
	return nil, f.err
}

func (f *fakeDfsClient) DfsUploadDocumentFileV2(context.Context, *dfs.TLDfsUploadDocumentFileV2) (*tg.Document, error) {
	return nil, f.err
}

func (f *fakeDfsClient) DfsUploadGifDocumentMedia(context.Context, *dfs.TLDfsUploadGifDocumentMedia) (*tg.Document, error) {
	return nil, f.err
}

func (f *fakeDfsClient) DfsUploadMp4DocumentMedia(context.Context, *dfs.TLDfsUploadMp4DocumentMedia) (*tg.Document, error) {
	return nil, f.err
}

func (f *fakeDfsClient) DfsUploadedProfilePhoto(context.Context, *dfs.TLDfsUploadedProfilePhoto) (*tg.Photo, error) {
	return nil, f.err
}

func (f *fakeDfsClient) DfsWriteFilePartData(context.Context, *dfs.TLDfsWriteFilePartData) (*tg.Bool, error) {
	return nil, f.err
}

func (f *fakeDfsClient) DfsUploadEncryptedFileV2(context.Context, *dfs.TLDfsUploadEncryptedFileV2) (*tg.EncryptedFile, error) {
	return nil, f.err
}

func (f *fakeDfsClient) DfsDownloadFile(context.Context, *dfs.TLDfsDownloadFile) (*tg.UploadFile, error) {
	return nil, f.err
}

func (f *fakeDfsClient) DfsUploadWallPaperFile(context.Context, *dfs.TLDfsUploadWallPaperFile) (*tg.Document, error) {
	return nil, f.err
}

func (f *fakeDfsClient) DfsUploadThemeFile(context.Context, *dfs.TLDfsUploadThemeFile) (*tg.Document, error) {
	return nil, f.err
}

func (f *fakeDfsClient) DfsUploadRingtoneFile(context.Context, *dfs.TLDfsUploadRingtoneFile) (*tg.Document, error) {
	return nil, f.err
}

func TestDfsClientForwardsCommitUpload(t *testing.T) {
	finalized := &dfs.FileFinalizedObject{ObjectId: "object-1"}
	fake := &fakeDfsClient{finalized: finalized}
	client := NewDFSClient(fake)
	req := &dfs.TLDfsCommitUpload{UploadSessionId: "upload-1"}

	got, err := client.CommitUpload(context.Background(), req)
	if err != nil {
		t.Fatalf("CommitUpload() error = %v", err)
	}
	if got != finalized {
		t.Fatalf("CommitUpload() result = %#v, want finalized object", got)
	}
	if fake.commitReq != req {
		t.Fatalf("forwarded request = %#v, want original", fake.commitReq)
	}
}

func TestDfsClientForwardsPutFile(t *testing.T) {
	finalized := &dfs.FileFinalizedObject{ObjectId: "object-1"}
	fake := &fakeDfsClient{finalized: finalized}
	client := NewDFSClient(fake)
	req := &dfs.TLDfsPutFile{FileName: "file.bin"}

	got, err := client.PutFile(context.Background(), req)
	if err != nil {
		t.Fatalf("PutFile() error = %v", err)
	}
	if got != finalized {
		t.Fatalf("PutFile() result = %#v, want finalized object", got)
	}
	if fake.putReq != req {
		t.Fatalf("forwarded request = %#v, want original", fake.putReq)
	}
}

func TestDfsClientForwardsLegacyUploadPhoto(t *testing.T) {
	photo := tg.MakeTLPhoto(&tg.TLPhoto{Id: 10}).ToPhoto()
	fake := &fakeDfsClient{photo: photo}
	client := NewDFSClient(fake)
	req := &dfs.TLDfsUploadPhotoFileV2{Creator: 1001}

	got, err := client.UploadPhotoFileV2ViaLegacyDFS(context.Background(), req)
	if err != nil {
		t.Fatalf("UploadPhotoFileV2ViaLegacyDFS() error = %v", err)
	}
	if got != photo {
		t.Fatalf("UploadPhotoFileV2ViaLegacyDFS() result = %#v, want forwarded photo", got)
	}
	if fake.photoReq != req {
		t.Fatalf("forwarded request = %#v, want original", fake.photoReq)
	}
}
