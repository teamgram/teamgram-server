package repository

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const documentAttributeStorageVersionV2 = 2
const documentAttributeStorageTLVector = "tl_object_vector"

type documentAttributeStorageEnvelope struct {
	SchemaVersion int    `json:"schema_version"`
	Encoding      string `json:"encoding"`
	Layer         int32  `json:"layer"`
	Bytes         string `json:"bytes"`
}

func encodeDocumentAttributesForStorage(attrs []tg.DocumentAttributeClazz) (string, error) {
	if len(attrs) == 0 {
		return "", nil
	}
	data, err := encodeDocumentAttributeVector(attrs)
	if err != nil {
		return "", err
	}
	envelope := documentAttributeStorageEnvelope{
		SchemaVersion: documentAttributeStorageVersionV2,
		Encoding:      documentAttributeStorageTLVector,
		Layer:         documentAttributeVectorLayer,
		Bytes:         base64.StdEncoding.EncodeToString(data),
	}
	b, err := json.Marshal(envelope)
	if err != nil {
		return "", fmt.Errorf("%w: encode document attribute storage envelope: %w", media.ErrMediaInvalidUploadedFile, err)
	}
	return string(b), nil
}

func decodeDocumentAttributesFromStorage(raw string) ([]tg.DocumentAttributeClazz, error) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return []tg.DocumentAttributeClazz{}, nil
	}
	if strings.HasPrefix(trimmed, "[") {
		return decodeLegacyDocumentAttributes(trimmed)
	}
	var envelope documentAttributeStorageEnvelope
	if err := json.Unmarshal([]byte(trimmed), &envelope); err != nil {
		return nil, fmt.Errorf("%w: decode document attribute storage envelope: %w", media.ErrMediaStorage, err)
	}
	if envelope.SchemaVersion != documentAttributeStorageVersionV2 {
		return nil, fmt.Errorf("%w: unsupported document attribute storage schema version %d", media.ErrMediaStorage, envelope.SchemaVersion)
	}
	if envelope.Encoding != documentAttributeStorageTLVector {
		return nil, fmt.Errorf("%w: unsupported document attribute storage encoding %q", media.ErrMediaStorage, envelope.Encoding)
	}
	data, err := base64.StdEncoding.DecodeString(envelope.Bytes)
	if err != nil {
		return nil, fmt.Errorf("%w: decode document attribute storage bytes: %w", media.ErrMediaStorage, err)
	}
	attrs, err := decodeDocumentAttributeVector(data)
	if err != nil {
		return nil, fmt.Errorf("%w: decode stored document attribute vector: %v", media.ErrMediaStorage, err)
	}
	return attrs, nil
}
