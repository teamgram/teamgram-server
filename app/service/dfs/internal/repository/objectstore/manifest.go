package objectstore

import (
	"errors"
	"fmt"
	"strings"
)

var ErrInvalidManifestKey = errors.New("objectstore: invalid manifest key")

type ManifestKeys struct {
	MetaPrefix string
}

type ObjectManifest struct {
	ObjectID        string      `json:"object_id"`
	UploadSessionID string      `json:"upload_session_id,omitempty"`
	Bucket          string      `json:"bucket"`
	Key             string      `json:"key"`
	Size            int64       `json:"size"`
	MimeType        string      `json:"mime_type"`
	SHA256          []byte      `json:"sha256"`
	DCID            int32       `json:"dc_id"`
	Chunks          []HashChunk `json:"chunks"`
}

type HashChunk struct {
	Offset int64  `json:"offset"`
	Limit  int32  `json:"limit"`
	Hash   []byte `json:"hash"`
}

func (k ManifestKeys) prefix() string {
	p := strings.Trim(k.MetaPrefix, "/")
	if p == "" {
		return "_meta"
	}
	return p
}

func (k ManifestKeys) Object(objectID string) (string, error) {
	if err := validateManifestID(objectID); err != nil {
		return "", err
	}
	return k.prefix() + "/objects/" + objectID + ".json", nil
}

func (k ManifestKeys) Upload(uploadSessionID string) (string, error) {
	if err := validateManifestID(uploadSessionID); err != nil {
		return "", err
	}
	return k.prefix() + "/uploads/" + uploadSessionID + ".json", nil
}

func (k ManifestKeys) Hashes(objectID string) (string, error) {
	if err := validateManifestID(objectID); err != nil {
		return "", err
	}
	return k.prefix() + "/hashes/" + objectID + "/v1.json", nil
}

func validateManifestID(id string) error {
	if id == "" || strings.TrimSpace(id) != id || strings.Contains(id, "/") ||
		strings.Contains(id, `\`) || strings.Contains(id, "..") {
		return fmt.Errorf("%w: %q", ErrInvalidManifestKey, id)
	}
	return nil
}
