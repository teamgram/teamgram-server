package repository

import "encoding/json"

const projectionCacheSchemaVersion = 1

type projectionCacheEnvelope[T any] struct {
	SchemaVersion int `json:"schema_version"`
	Data          T   `json:"data"`
}

type projectionUserCacheDTO struct {
	UserID     int64  `json:"user_id"`
	FirstName  string `json:"first_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	Username   string `json:"username,omitempty"`
	Phone      string `json:"phone,omitempty"`
	AccessHash int64  `json:"access_hash"`
	Deleted    bool   `json:"deleted,omitempty"`
	IsBot      bool   `json:"is_bot,omitempty"`
}

type projectionPrivacyCacheDTO struct {
	UserID int64                      `json:"user_id"`
	Rules  map[int32][]privacyRuleDTO `json:"rules"`
}

type projectionContactMapCacheDTO struct {
	OwnerUserID int64                           `json:"owner_user_id"`
	Contacts    map[int64]projectionContactFact `json:"contacts"`
}

type projectionPresenceCacheDTO struct {
	UserID     int64 `json:"user_id"`
	LastSeenAt int64 `json:"last_seen_at"`
	Expires    int32 `json:"expires"`
}

func encodeProjectionCache[T any](data T) (string, error) {
	b, err := json.Marshal(projectionCacheEnvelope[T]{
		SchemaVersion: projectionCacheSchemaVersion,
		Data:          data,
	})
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func decodeProjectionCache[T any](raw string) (T, bool) {
	var zero T
	if raw == "" {
		return zero, false
	}

	var env projectionCacheEnvelope[T]
	if err := json.Unmarshal([]byte(raw), &env); err != nil {
		return zero, false
	}
	if env.SchemaVersion != projectionCacheSchemaVersion {
		return zero, false
	}
	return env.Data, true
}
