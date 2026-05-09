package objectstore

import (
	"crypto/sha256"
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
	OwnerID         int64       `json:"owner_id,omitempty"`
	FileID          int64       `json:"file_id,omitempty"`
	Purpose         string      `json:"purpose,omitempty"`
	FileName        string      `json:"file_name,omitempty"`
	Bucket          string      `json:"bucket"`
	Key             string      `json:"key"`
	Size            int64       `json:"size"`
	MimeType        string      `json:"mime_type"`
	StorageType     int32       `json:"storage_type"`
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

func BuildHashChunks(data []byte, chunkSize int) []HashChunk {
	if chunkSize <= 0 || len(data) == 0 {
		return nil
	}
	chunks := make([]HashChunk, 0, (len(data)+chunkSize-1)/chunkSize)
	for offset := 0; offset < len(data); offset += chunkSize {
		end := offset + chunkSize
		if end > len(data) {
			end = len(data)
		}
		sum := sha256.Sum256(data[offset:end])
		chunks = append(chunks, HashChunk{
			Offset: int64(offset),
			Limit:  int32(end - offset),
			Hash:   append([]byte(nil), sum[:]...),
		})
	}
	return chunks
}

func FilterHashChunks(chunks []HashChunk, offset int64, limit int32) []HashChunk {
	if offset < 0 || limit < 0 {
		return nil
	}
	end := int64(0)
	if limit > 0 {
		end = offset + int64(limit)
	}
	out := make([]HashChunk, 0, len(chunks))
	for _, chunk := range chunks {
		chunkEnd := chunk.Offset + int64(chunk.Limit)
		if limit <= 0 {
			if chunkEnd > offset {
				out = append(out, chunk)
			}
			continue
		}
		if chunk.Offset < end && chunkEnd > offset {
			out = append(out, chunk)
		}
	}
	return out
}
