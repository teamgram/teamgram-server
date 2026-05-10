package repository

import (
	"errors"
	"fmt"

	dfsapi "github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
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

func wrapMediaInvalidUploadedFile(op string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%w: %s: %w", media.ErrMediaInvalidUploadedFile, op, err)
}

func wrapMediaInvalidArgument(op string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%w: %s: %w", media.ErrMediaInvalidArgument, op, err)
}

func wrapMediaChecksumInvalid(op string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%w: %s: %w", media.ErrMediaChecksumInvalid, op, err)
}

func wrapMediaDownstream(op string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%w: %s: %w", media.ErrMediaDownstream, op, err)
}

func wrapDfsUploadError(op string, err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, dfsapi.ErrDfsInvalidArgument) {
		return wrapMediaInvalidArgument(op, err)
	}
	if errors.Is(err, dfsapi.ErrDfsChecksumInvalid) {
		return wrapMediaChecksumInvalid(op, err)
	}
	var missing *dfsapi.MissingUploadPartError
	if errors.As(err, &missing) {
		return wrapMediaInvalidUploadedFile(op, err)
	}
	if errors.Is(err, dfsapi.ErrDfsFileNotFound) ||
		errors.Is(err, dfsapi.ErrDfsInvalidFilePart) ||
		errors.Is(err, dfsapi.ErrDfsImageProcessFailed) ||
		errors.Is(err, dfsapi.ErrDfsVideoProcessFailed) {
		return wrapMediaInvalidUploadedFile(op, err)
	}
	return wrapMediaDownstream(op, err)
}

func isServiceError(err error) bool {
	return errors.Is(err, media.ErrMediaStorage) ||
		errors.Is(err, media.ErrPhotoNotFound) ||
		errors.Is(err, media.ErrDocumentNotFound) ||
		errors.Is(err, media.ErrFileLocationInvalid) ||
		errors.Is(err, media.ErrFileReferenceEmpty) ||
		errors.Is(err, media.ErrFileReferenceExpired) ||
		errors.Is(err, media.ErrFileReferenceInvalid) ||
		errors.Is(err, media.ErrMediaInvalidArgument) ||
		errors.Is(err, media.ErrMediaInvalidUploadedFile) ||
		errors.Is(err, media.ErrMediaChecksumInvalid) ||
		errors.Is(err, media.ErrMediaDownstream) ||
		errors.Is(err, media.ErrMediaBlocked)
}
