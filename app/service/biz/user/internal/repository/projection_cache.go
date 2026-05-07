package repository

import (
	"context"
	"encoding/json"
)

const projectionCacheSchemaVersion = 1

type projectionCacheDecodeStatus int

const (
	projectionCacheDecodeHit projectionCacheDecodeStatus = iota
	projectionCacheDecodeMiss
	projectionCacheDecodeStale
	projectionCacheDecodeCorrupt
)

type projectionCacheEnvelope[T any] struct {
	SchemaVersion int `json:"schema_version"`
	Data          T   `json:"data"`
}

type projectionUserCacheDTO struct {
	UserID                        int64                           `json:"user_id"`
	AccessHash                    int64                           `json:"access_hash"`
	UserType                      int32                           `json:"user_type,omitempty"`
	SecretKeyID                   int64                           `json:"secret_key_id,omitempty"`
	FirstName                     string                          `json:"first_name,omitempty"`
	LastName                      string                          `json:"last_name,omitempty"`
	Username                      string                          `json:"username,omitempty"`
	Phone                         string                          `json:"phone,omitempty"`
	CountryCode                   string                          `json:"country_code,omitempty"`
	Verified                      bool                            `json:"verified,omitempty"`
	Support                       bool                            `json:"support,omitempty"`
	Scam                          bool                            `json:"scam,omitempty"`
	Fake                          bool                            `json:"fake,omitempty"`
	Premium                       bool                            `json:"premium,omitempty"`
	About                         string                          `json:"about,omitempty"`
	Restricted                    bool                            `json:"restricted,omitempty"`
	Deleted                       bool                            `json:"deleted,omitempty"`
	StoriesMaxID                  int32                           `json:"stories_max_id,omitempty"`
	Birthday                      string                          `json:"birthday,omitempty"`
	PersonalChannelID             int64                           `json:"personal_channel_id,omitempty"`
	PhotoID                       int64                           `json:"photo_id,omitempty"`
	IsBot                         bool                            `json:"is_bot,omitempty"`
	EmojiStatusDocumentID         int64                           `json:"emoji_status_document_id,omitempty"`
	EmojiStatusUntil              int32                           `json:"emoji_status_until,omitempty"`
	Color                         int32                           `json:"color,omitempty"`
	ColorBackgroundEmojiID        int64                           `json:"color_background_emoji_id,omitempty"`
	ProfileColor                  int32                           `json:"profile_color,omitempty"`
	ProfileColorBackgroundEmojiID int64                           `json:"profile_color_background_emoji_id,omitempty"`
	Bot                           *projectionBotCacheDTO          `json:"bot,omitempty"`
	Usernames                     []projectionUsernameCacheDTO    `json:"usernames,omitempty"`
	RestrictionReasons            []projectionRestrictionCacheDTO `json:"restriction_reasons,omitempty"`
}

type projectionBotCacheDTO struct {
	ID                   int64   `json:"id"`
	BotType              int32   `json:"bot_type,omitempty"`
	Creator              int64   `json:"creator,omitempty"`
	Token                string  `json:"token,omitempty"`
	Description          string  `json:"description,omitempty"`
	BotChatHistory       bool    `json:"bot_chat_history,omitempty"`
	BotNochats           bool    `json:"bot_nochats,omitempty"`
	BotInlineGeo         bool    `json:"bot_inline_geo,omitempty"`
	BotInfoVersion       int32   `json:"bot_info_version,omitempty"`
	BotInlinePlaceholder *string `json:"bot_inline_placeholder,omitempty"`
	BotAttachMenu        bool    `json:"bot_attach_menu,omitempty"`
	AttachMenuEnabled    bool    `json:"attach_menu_enabled,omitempty"`
	BotCanEdit           bool    `json:"bot_can_edit,omitempty"`
	BotBusiness          bool    `json:"bot_business,omitempty"`
	BotHasMainApp        bool    `json:"bot_has_main_app,omitempty"`
	BotActiveUsers       *int32  `json:"bot_active_users,omitempty"`
}

type projectionUsernameCacheDTO struct {
	Username string `json:"username"`
	Editable bool   `json:"editable,omitempty"`
	Active   bool   `json:"active,omitempty"`
	Order    int32  `json:"order,omitempty"`
}

type projectionRestrictionCacheDTO struct {
	Platform string `json:"platform,omitempty"`
	Reason   string `json:"reason,omitempty"`
	Text     string `json:"text,omitempty"`
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
	UserID      int64 `json:"user_id"`
	HasPresence bool  `json:"has_presence,omitempty"`
	LastSeenAt  int64 `json:"last_seen_at,omitempty"`
	Expires     int32 `json:"expires,omitempty"`
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
	data, status := decodeProjectionCacheStatus[T](raw)
	return data, status == projectionCacheDecodeHit
}

func decodeProjectionCacheStatus[T any](raw string) (T, projectionCacheDecodeStatus) {
	var zero T
	if raw == "" {
		return zero, projectionCacheDecodeMiss
	}

	var env projectionCacheEnvelope[T]
	if err := json.Unmarshal([]byte(raw), &env); err != nil {
		return zero, projectionCacheDecodeCorrupt
	}
	if env.SchemaVersion != projectionCacheSchemaVersion {
		return zero, projectionCacheDecodeStale
	}
	return env.Data, projectionCacheDecodeHit
}

func (r *Repository) getProjectionComponentCache(ctx context.Context, key string, out interface{}) bool {
	var env projectionCacheEnvelope[json.RawMessage]
	if err := r.GetCache(ctx, key, &env); err != nil {
		return false
	}
	raw, status := decodeProjectionCacheStatus[json.RawMessage](mustMarshalProjectionCacheEnvelope(env))
	if status != projectionCacheDecodeHit {
		return false
	}
	if err := json.Unmarshal(raw, out); err != nil {
		return false
	}
	return true
}

func (r *Repository) setProjectionComponentCache(ctx context.Context, key string, data interface{}) {
	_ = r.SetCache(ctx, key, projectionCacheEnvelope[interface{}]{
		SchemaVersion: projectionCacheSchemaVersion,
		Data:          data,
	})
}

func mustMarshalProjectionCacheEnvelope(env projectionCacheEnvelope[json.RawMessage]) string {
	b, err := json.Marshal(env)
	if err != nil {
		return ""
	}
	return string(b)
}
