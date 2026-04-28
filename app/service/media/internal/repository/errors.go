package repository

import (
	"errors"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/service/media/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
)

func isNotFound(err error) bool {
	return errors.Is(err, model.ErrNotFound)
}

func wrapStorage(op string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%w: %s: %w", media.ErrMediaStorage, op, err)
}

func isServiceError(err error) bool {
	return errors.Is(err, media.ErrMediaStorage) ||
		errors.Is(err, media.ErrPhotoNotFound) ||
		errors.Is(err, media.ErrDocumentNotFound)
}
