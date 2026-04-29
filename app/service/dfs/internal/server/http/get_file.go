package http

import (
	"context"
	"errors"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileOpener interface {
	OpenHttpUploadedFile(ctx context.Context, creator, fileID int64) (*core.HttpUploadedFile, error)
}

type pathParamKey struct{}

func GetDfsFile(opener FileOpener) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		creator, fileID, err := parseFileParam(fileParam(r))
		if err != nil {
			http.NotFound(w, r)
			return
		}
		file, err := opener.OpenHttpUploadedFile(r.Context(), creator, fileID)
		if err != nil {
			if errors.Is(err, dfs.ErrDfsFileNotFound) || errors.Is(err, dfs.ErrDfsInvalidArgument) {
				http.NotFound(w, r)
				return
			}
			logx.Errorf("getDfsFile(%d, %d) - error: %v", creator, fileID, err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		name := file.Name
		if name == "" {
			name = path.Base(r.URL.Path)
		}
		http.ServeContent(w, r, name, file.Modtime, file.ReadSeeker)
	}
}

func fileParam(r *http.Request) string {
	if v, ok := r.Context().Value(pathParamKey{}).(string); ok {
		return v
	}
	return path.Base(r.URL.Path)
}

func parseFileParam(raw string) (int64, int64, error) {
	parts := strings.Split(raw, "_")
	if len(parts) != 2 {
		return 0, 0, dfs.ErrDfsInvalidArgument
	}
	creator, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return 0, 0, dfs.ErrDfsInvalidArgument
	}
	filePart := strings.Split(parts[1], ".")[0]
	fileID, err := strconv.ParseInt(filePart, 10, 64)
	if err != nil {
		return 0, 0, dfs.ErrDfsInvalidArgument
	}
	return creator, fileID, nil
}
