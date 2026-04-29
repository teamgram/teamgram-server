package media

import "errors"

var (
	ErrMediaStorage             = errors.New("media: storage failure")
	ErrPhotoNotFound            = errors.New("media: photo not found")
	ErrDocumentNotFound         = errors.New("media: document not found")
	ErrMediaInvalidArgument     = errors.New("media: invalid argument")
	ErrMediaInvalidUploadedFile = errors.New("media: invalid uploaded file")
	ErrMediaChecksumInvalid     = errors.New("media: checksum invalid")
	ErrMediaDownstream          = errors.New("media: downstream failure")
	ErrMediaBlocked             = errors.New("media: blocked")
)
