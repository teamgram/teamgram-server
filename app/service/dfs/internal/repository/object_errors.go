package repository

import (
	"errors"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	minioadapter "github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/minio"
)

func mapObjectReadError(op string, err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, minioadapter.ErrObjectNotFound) {
		return fmt.Errorf("%w: %s: %w", dfs.ErrDfsFileNotFound, op, err)
	}
	return dfs.WrapDfsStorage(op, err)
}
