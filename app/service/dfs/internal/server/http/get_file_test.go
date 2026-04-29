package http

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/core"
)

type fakeHTTPFileOpener struct {
	creator int64
	fileID  int64
	file    *core.HttpUploadedFile
	err     error
}

func (f *fakeHTTPFileOpener) OpenHttpUploadedFile(_ context.Context, creator, fileID int64) (*core.HttpUploadedFile, error) {
	f.creator = creator
	f.fileID = fileID
	if f.err != nil {
		return nil, f.err
	}
	return f.file, nil
}

func TestMiniHttpGetFileParsesCreatorAndFileID(t *testing.T) {
	opener := &fakeHTTPFileOpener{file: httpFile("payload")}
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/dfs/file/1001_2002.gif", nil)
	req = withFilePathParam(req, "1001_2002.gif")

	GetDfsFile(opener)(rr, req)

	if opener.creator != 1001 || opener.fileID != 2002 {
		t.Fatalf("OpenHttpUploadedFile creator/file = %d/%d, want 1001/2002", opener.creator, opener.fileID)
	}
	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rr.Code)
	}
}

func TestMiniHttpGetFileReturns404ForInvalidPath(t *testing.T) {
	opener := &fakeHTTPFileOpener{file: httpFile("payload")}
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/dfs/file/bad", nil)
	req = withFilePathParam(req, "bad")

	GetDfsFile(opener)(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want 404", rr.Code)
	}
	if opener.creator != 0 || opener.fileID != 0 {
		t.Fatalf("OpenHttpUploadedFile called for invalid path")
	}
}

func TestMiniHttpGetFileUsesCoreOpenHttpUploadedFile(t *testing.T) {
	opener := &fakeHTTPFileOpener{file: httpFile("served-bytes")}
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/dfs/file/1001_2002.jpg", nil)
	req = withFilePathParam(req, "1001_2002.jpg")

	GetDfsFile(opener)(rr, req)

	if rr.Body.String() != "served-bytes" {
		t.Fatalf("body = %q, want served-bytes", rr.Body.String())
	}
}

func TestMiniHttpGetFileMapsNotFoundTo404(t *testing.T) {
	opener := &fakeHTTPFileOpener{err: dfs.ErrDfsFileNotFound}
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/dfs/file/1001_2002.jpg", nil)
	req = withFilePathParam(req, "1001_2002.jpg")

	GetDfsFile(opener)(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want 404", rr.Code)
	}
}

func httpFile(data string) *core.HttpUploadedFile {
	return &core.HttpUploadedFile{
		Name:       "file.dat",
		Modtime:    time.Unix(100, 0),
		ReadSeeker: bytes.NewReader([]byte(data)),
	}
}

func withFilePathParam(r *http.Request, file string) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), pathParamKey{}, file))
}

var _ io.ReadSeeker = (*bytes.Reader)(nil)
