package media

import "errors"

var (
	ErrMediaStorage     = errors.New("media: storage failure")
	ErrPhotoNotFound    = errors.New("media: photo not found")
	ErrDocumentNotFound = errors.New("media: document not found")
)
