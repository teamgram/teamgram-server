package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/teamgram/marmota/pkg/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
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
	OwnerUserID       int64                           `json:"owner_user_id"`
	Contacts          map[int64]projectionContactFact `json:"contacts"`
	CoveredContactIDs []int64                         `json:"covered_contact_ids,omitempty"`
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
		if !projectionCacheNotFound(err) {
			logx.WithContext(ctx).Errorf("user projection cache get failed: key=%s err=%v", key, err)
		}
		return false
	}
	raw, status := decodeProjectionCacheStatus[json.RawMessage](mustMarshalProjectionCacheEnvelope(env))
	if status != projectionCacheDecodeHit {
		r.logProjectionCacheDecodeMiss(ctx, key, status)
		return false
	}
	if err := json.Unmarshal(raw, out); err != nil {
		logx.WithContext(ctx).Errorf("user projection cache dto decode failed: key=%s err=%v", key, err)
		return false
	}
	return true
}

func getProjectionComponentCaches[T any](r *Repository, ctx context.Context, keys []string) map[string]T {
	out := make(map[string]T, len(keys))
	if len(keys) == 0 {
		return out
	}
	if err := r.QueryRows(ctx, func(ctx context.Context, _ *sqlx.DB, _ ...string) (map[string]interface{}, error) {
		return map[string]interface{}{}, nil
	}, func(key, raw string) (interface{}, error) {
		dto, status := decodeProjectionCacheStatus[T](raw)
		switch status {
		case projectionCacheDecodeHit:
			out[key] = dto
			return nil, nil
		case projectionCacheDecodeMiss:
			return nil, sql.ErrNoRows
		case projectionCacheDecodeStale:
			r.logProjectionCacheDecodeMiss(ctx, key, status)
			r.deleteProjectionComponentCaches(ctx, key)
			return nil, nil
		case projectionCacheDecodeCorrupt:
			r.logProjectionCacheDecodeMiss(ctx, key, status)
			r.deleteProjectionComponentCaches(ctx, key)
			return nil, nil
		default:
			r.logProjectionCacheDecodeMiss(ctx, key, status)
			r.deleteProjectionComponentCaches(ctx, key)
			return nil, nil
		}
	}, keys...); err != nil {
		logx.WithContext(ctx).Errorf("user projection cache bulk get failed: keys=%d err=%v", len(keys), err)
	}
	return out
}

func (r *Repository) setProjectionComponentCache(ctx context.Context, key string, data interface{}) {
	if err := r.SetCache(ctx, key, projectionCacheEnvelope[interface{}]{
		SchemaVersion: projectionCacheSchemaVersion,
		Data:          data,
	}); err != nil {
		logx.WithContext(ctx).Errorf("user projection cache set failed: key=%s err=%v", key, err)
	}
}

func (r *Repository) deleteProjectionComponentCaches(ctx context.Context, keys ...string) {
	if len(keys) == 0 {
		return
	}
	if err := r.DelCache(ctx, keys...); err != nil {
		logx.WithContext(ctx).Errorf("user projection cache delete failed: keys=%d err=%v", len(keys), err)
	}
}

func (r *Repository) invalidateProjectionFactCache(ctx context.Context, userID int64) {
	if userID > 0 {
		r.deleteProjectionComponentCaches(ctx, projectionFactsCacheKey(userID))
	}
}

func (r *Repository) invalidateProjectionPrivacyCache(ctx context.Context, userID int64) {
	if userID > 0 {
		r.deleteProjectionComponentCaches(ctx, projectionPrivacyCacheKey(userID))
	}
}

func (r *Repository) invalidateProjectionPresenceCache(ctx context.Context, userID int64) {
	if userID > 0 {
		r.deleteProjectionComponentCaches(ctx, projectionPresenceCacheKey(userID))
	}
}

func (r *Repository) invalidateProjectionContactMapCaches(ctx context.Context, ownerIDs ...int64) {
	keys := make([]string, 0, len(ownerIDs))
	seen := make(map[int64]struct{}, len(ownerIDs))
	for _, id := range ownerIDs {
		if id <= 0 {
			continue
		}
		if _, ok := seen[id]; ok {
			continue
		}
		seen[id] = struct{}{}
		keys = append(keys, projectionContactMapCacheKey(id))
	}
	r.deleteProjectionComponentCaches(ctx, keys...)
}

func (r *Repository) logProjectionCacheDecodeMiss(ctx context.Context, key string, status projectionCacheDecodeStatus) {
	switch status {
	case projectionCacheDecodeMiss:
		return
	case projectionCacheDecodeStale:
		logx.WithContext(ctx).Errorf("user projection cache stale schema: key=%s", key)
	case projectionCacheDecodeCorrupt:
		logx.WithContext(ctx).Errorf("user projection cache corrupt payload: key=%s", key)
	default:
		logx.WithContext(ctx).Errorf("user projection cache decode failed: key=%s status=%d", key, status)
	}
}

func (r *Repository) logProjectionCacheIdentityMismatch(ctx context.Context, key, component string, expected, got int64) {
	logx.WithContext(ctx).Errorf("user projection cache identity mismatch: component=%s key=%s expected=%d got=%d", component, key, expected, got)
	r.deleteProjectionComponentCaches(ctx, key)
}

func projectionCacheNotFound(err error) bool {
	return err == nil || errors.Is(err, sql.ErrNoRows) || errors.Is(err, sqlx.ErrNotFound)
}

func mustMarshalProjectionCacheEnvelope(env projectionCacheEnvelope[json.RawMessage]) string {
	b, err := json.Marshal(env)
	if err != nil {
		return ""
	}
	return string(b)
}
