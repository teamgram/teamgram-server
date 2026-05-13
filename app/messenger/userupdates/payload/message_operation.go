package payload

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type MessageOperationV1 struct {
	SchemaVersion              int               `json:"schema_version"`
	OperationKind              string            `json:"operation_kind"`
	CanonicalMessageID         int64             `json:"canonical_message_id"`
	PeerType                   int32             `json:"peer_type"`
	PeerID                     int64             `json:"peer_id"`
	PeerSeq                    int64             `json:"peer_seq"`
	FromUserID                 int64             `json:"from_user_id"`
	ToUserID                   int64             `json:"to_user_id"`
	Date                       int32             `json:"date"`
	EditDate                   int32             `json:"edit_date,omitempty"`
	EditVersion                int32             `json:"edit_version,omitempty"`
	Out                        bool              `json:"out"`
	MessageText                string            `json:"message_text"`
	Entities                   []MessageEntityV1 `json:"entities,omitempty"`
	UserMessageID              int64             `json:"user_message_id,omitempty"`
	ReplyToCanonicalMessageID  int64             `json:"reply_to_canonical_message_id,omitempty"`
	ReplyToPeerSeq             int64             `json:"reply_to_peer_seq,omitempty"`
	ReplyToUserMessageID       int64             `json:"reply_to_user_message_id,omitempty"`
	DependencyPts              []int64           `json:"dependency_pts,omitempty"`
	ClearDraft                 bool              `json:"clear_draft,omitempty"`
	SourcePermAuthKeyID        int64             `json:"source_perm_auth_key_id,omitempty"`
	ClearDraftBeforeDate       int32             `json:"clear_draft_before_date,omitempty"`
	SavedDialogSideEffect      bool              `json:"saved_dialog_side_effect,omitempty"`
	ReadMaxUserMessageID       int64             `json:"read_max_user_message_id,omitempty"`
	ReadInboxMaxPeerSeq        int64             `json:"read_inbox_max_peer_seq,omitempty"`
	ReadInboxMaxUserMessageID  int64             `json:"read_inbox_max_user_message_id,omitempty"`
	ReadOutboxMaxPeerSeq       int64             `json:"read_outbox_max_peer_seq,omitempty"`
	ReadOutboxMaxUserMessageID int64             `json:"read_outbox_max_user_message_id,omitempty"`
	DeletePeerSeqs             []int64           `json:"delete_peer_seqs,omitempty"`
	DeleteUserMessageIDs       []int64           `json:"delete_user_message_ids,omitempty"`
	DeleteMaxPeerSeq           int64             `json:"delete_max_peer_seq,omitempty"`
	JustClear                  bool              `json:"just_clear,omitempty"`
	Revoke                     bool              `json:"revoke,omitempty"`
	UnreadMark                 *bool             `json:"unread_mark,omitempty"`
	PinnedPeerSeq              int64             `json:"pinned_peer_seq,omitempty"`
	PinnedUserMessageID        int64             `json:"pinned_user_message_id,omitempty"`
	PinnedCanonicalMessageID   int64             `json:"pinned_canonical_message_id,omitempty"`
	HasScheduled               *bool             `json:"has_scheduled,omitempty"`
}

type MessageEntityV1 struct {
	Offset int32  `json:"offset"`
	Length int32  `json:"length"`
	Kind   string `json:"kind"`
	URL    string `json:"url,omitempty"`
	UserID int64  `json:"user_id,omitempty"`
}

type MediaRefV1 struct {
	SchemaVersion       int                      `json:"schema_version"`
	Kind                string                   `json:"kind"`
	ID                  int64                    `json:"id"`
	AccessHash          int64                    `json:"access_hash,omitempty"`
	FileReference       []byte                   `json:"file_reference,omitempty"`
	Date                int32                    `json:"date,omitempty"`
	DcID                int32                    `json:"dc_id,omitempty"`
	TTLSeconds          int32                    `json:"ttl_seconds,omitempty"`
	MimeType            string                   `json:"mime_type,omitempty"`
	Size                int64                    `json:"size,omitempty"`
	PhoneNumber         string                   `json:"phone_number,omitempty"`
	FirstName           string                   `json:"first_name,omitempty"`
	LastName            string                   `json:"last_name,omitempty"`
	Vcard               string                   `json:"vcard,omitempty"`
	UserID              int64                    `json:"user_id,omitempty"`
	PhotoSizes          []PhotoSizeRefV1         `json:"photo_sizes,omitempty"`
	DocumentThumbs      []PhotoSizeRefV1         `json:"document_thumbs,omitempty"`
	DocumentVideoThumbs []VideoSizeRefV1         `json:"document_video_thumbs,omitempty"`
	DocumentAttributes  []DocumentAttributeRefV1 `json:"document_attributes,omitempty"`
	DocumentMediaFlags  *DocumentMediaFlagsV1    `json:"document_media_flags,omitempty"`
	VideoCover          *PhotoRefV1              `json:"video_cover,omitempty"`
	VideoTimestamp      *int32                   `json:"video_timestamp,omitempty"`
}

type PhotoSizeRefV1 struct {
	Kind  string  `json:"kind"`
	Type  string  `json:"type,omitempty"`
	W     int32   `json:"w,omitempty"`
	H     int32   `json:"h,omitempty"`
	Size  int32   `json:"size,omitempty"`
	Bytes []byte  `json:"bytes,omitempty"`
	Sizes []int32 `json:"sizes,omitempty"`
}

type VideoSizeRefV1 struct {
	Kind             string           `json:"kind"`
	Type             string           `json:"type,omitempty"`
	W                int32            `json:"w,omitempty"`
	H                int32            `json:"h,omitempty"`
	Size             int32            `json:"size,omitempty"`
	VideoStartTs     *float64         `json:"video_start_ts,omitempty"`
	EmojiID          int64            `json:"emoji_id,omitempty"`
	StickerSet       *StickerSetRefV1 `json:"sticker_set,omitempty"`
	StickerID        int64            `json:"sticker_id,omitempty"`
	BackgroundColors []int32          `json:"background_colors,omitempty"`
}

type DocumentMediaFlagsV1 struct {
	Video   bool `json:"video"`
	Round   bool `json:"round"`
	Voice   bool `json:"voice"`
	Spoiler bool `json:"spoiler"`
}

type PhotoRefV1 struct {
	ID            int64            `json:"id"`
	AccessHash    int64            `json:"access_hash,omitempty"`
	FileReference []byte           `json:"file_reference,omitempty"`
	Date          int32            `json:"date,omitempty"`
	DcID          int32            `json:"dc_id,omitempty"`
	Sizes         []PhotoSizeRefV1 `json:"sizes,omitempty"`
	VideoSizes    []VideoSizeRefV1 `json:"video_sizes,omitempty"`
}

type StickerSetRefV1 struct {
	Kind       string `json:"kind"`
	ID         int64  `json:"id,omitempty"`
	AccessHash int64  `json:"access_hash,omitempty"`
	ShortName  string `json:"short_name,omitempty"`
	Emoticon   string `json:"emoticon,omitempty"`
}

type MaskCoordsRefV1 struct {
	N    int32   `json:"n"`
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
	Zoom float64 `json:"zoom"`
}

type DocumentAttributeRefV1 struct {
	Kind              string           `json:"kind"`
	W                 int32            `json:"w,omitempty"`
	H                 int32            `json:"h,omitempty"`
	FileName          string           `json:"file_name,omitempty"`
	Duration          int32            `json:"duration,omitempty"`
	DurationFloat     float64          `json:"duration_float,omitempty"`
	Title             *string          `json:"title,omitempty"`
	Performer         *string          `json:"performer,omitempty"`
	Waveform          []byte           `json:"waveform,omitempty"`
	Voice             bool             `json:"voice,omitempty"`
	RoundMessage      bool             `json:"round_message,omitempty"`
	SupportsStreaming bool             `json:"supports_streaming,omitempty"`
	NoSound           bool             `json:"nosound,omitempty"`
	PreloadPrefixSize *int32           `json:"preload_prefix_size,omitempty"`
	VideoStartTs      *float64         `json:"video_start_ts,omitempty"`
	VideoCodec        *string          `json:"video_codec,omitempty"`
	Alt               string           `json:"alt,omitempty"`
	StickerSet        *StickerSetRefV1 `json:"sticker_set,omitempty"`
	Mask              bool             `json:"mask,omitempty"`
	MaskCoords        *MaskCoordsRefV1 `json:"mask_coords,omitempty"`
	Free              bool             `json:"free,omitempty"`
	TextColor         bool             `json:"text_color,omitempty"`
}

type MessageAttrsV1 struct {
	SchemaVersion int   `json:"schema_version"`
	GroupedID     int64 `json:"grouped_id,omitempty"`
	Noforwards    bool  `json:"noforwards,omitempty"`
	Silent        bool  `json:"silent,omitempty"`
	InvertMedia   bool  `json:"invert_media,omitempty"`
}

type ForwardRefV1 struct {
	SchemaVersion      int    `json:"schema_version"`
	FromUserID         int64  `json:"from_user_id,omitempty"`
	FromName           string `json:"from_name,omitempty"`
	Date               int64  `json:"date"`
	SourcePeerType     int32  `json:"source_peer_type,omitempty"`
	SourcePeerID       int64  `json:"source_peer_id,omitempty"`
	SourceMessageID    int64  `json:"source_message_id,omitempty"`
	SavedFromPeerType  int32  `json:"saved_from_peer_type,omitempty"`
	SavedFromPeerID    int64  `json:"saved_from_peer_id,omitempty"`
	SavedFromMessageID int64  `json:"saved_from_message_id,omitempty"`
}

type ServiceActionRefV1 struct {
	SchemaVersion int     `json:"schema_version"`
	Kind          string  `json:"kind"`
	Title         string  `json:"title,omitempty"`
	Users         []int64 `json:"users,omitempty"`
}

type MessageOperationV3 struct {
	SchemaVersion              int                 `json:"schema_version"`
	OperationKind              string              `json:"operation_kind"`
	CanonicalMessageID         int64               `json:"canonical_message_id"`
	PeerType                   int32               `json:"peer_type"`
	PeerID                     int64               `json:"peer_id"`
	PeerSeq                    int64               `json:"peer_seq"`
	FromUserID                 int64               `json:"from_user_id"`
	ToUserID                   int64               `json:"to_user_id"`
	Date                       int32               `json:"date"`
	EditDate                   int32               `json:"edit_date,omitempty"`
	EditVersion                int32               `json:"edit_version,omitempty"`
	Out                        bool                `json:"out"`
	MessageText                string              `json:"message_text"`
	Entities                   []MessageEntityV1   `json:"entities,omitempty"`
	UserMessageID              int64               `json:"user_message_id,omitempty"`
	ReplyToCanonicalMessageID  int64               `json:"reply_to_canonical_message_id,omitempty"`
	ReplyToPeerSeq             int64               `json:"reply_to_peer_seq,omitempty"`
	ReplyToUserMessageID       int64               `json:"reply_to_user_message_id,omitempty"`
	DependencyPts              []int64             `json:"dependency_pts,omitempty"`
	ClearDraft                 bool                `json:"clear_draft,omitempty"`
	SourcePermAuthKeyID        int64               `json:"source_perm_auth_key_id,omitempty"`
	ClearDraftBeforeDate       int32               `json:"clear_draft_before_date,omitempty"`
	SavedDialogSideEffect      bool                `json:"saved_dialog_side_effect,omitempty"`
	ReadMaxUserMessageID       int64               `json:"read_max_user_message_id,omitempty"`
	ReadInboxMaxPeerSeq        int64               `json:"read_inbox_max_peer_seq,omitempty"`
	ReadInboxMaxUserMessageID  int64               `json:"read_inbox_max_user_message_id,omitempty"`
	ReadOutboxMaxPeerSeq       int64               `json:"read_outbox_max_peer_seq,omitempty"`
	ReadOutboxMaxUserMessageID int64               `json:"read_outbox_max_user_message_id,omitempty"`
	DeletePeerSeqs             []int64             `json:"delete_peer_seqs,omitempty"`
	DeleteUserMessageIDs       []int64             `json:"delete_user_message_ids,omitempty"`
	DeleteMaxPeerSeq           int64               `json:"delete_max_peer_seq,omitempty"`
	JustClear                  bool                `json:"just_clear,omitempty"`
	Revoke                     bool                `json:"revoke,omitempty"`
	UnreadMark                 *bool               `json:"unread_mark,omitempty"`
	PinnedPeerSeq              int64               `json:"pinned_peer_seq,omitempty"`
	PinnedUserMessageID        int64               `json:"pinned_user_message_id,omitempty"`
	PinnedCanonicalMessageID   int64               `json:"pinned_canonical_message_id,omitempty"`
	HasScheduled               *bool               `json:"has_scheduled,omitempty"`
	MediaRef                   *MediaRefV1         `json:"media_ref,omitempty"`
	Attrs                      *MessageAttrsV1     `json:"attrs,omitempty"`
	ForwardRef                 *ForwardRefV1       `json:"forward_ref,omitempty"`
	ServiceAction              *ServiceActionRefV1 `json:"service_action,omitempty"`
}

type UpdateFactV1 struct {
	SchemaVersion int             `json:"schema_version"`
	Kind          string          `json:"kind"`
	Payload       json.RawMessage `json:"payload"`
}

type NewMessageFactV1 struct {
	SchemaVersion             int                 `json:"schema_version"`
	CanonicalMessageID        int64               `json:"canonical_message_id"`
	PeerType                  int32               `json:"peer_type"`
	PeerID                    int64               `json:"peer_id"`
	PeerSeq                   int64               `json:"peer_seq,omitempty"`
	SenderUserID              int64               `json:"sender_user_id"`
	ToUserID                  int64               `json:"to_user_id,omitempty"`
	ClientRandomID            int64               `json:"client_random_id"`
	Date                      int32               `json:"date"`
	MessageText               string              `json:"message_text,omitempty"`
	Entities                  []MessageEntityV1   `json:"entities,omitempty"`
	ReplyToCanonicalMessageID int64               `json:"reply_to_canonical_message_id,omitempty"`
	ReplyToUserMessageID      int64               `json:"reply_to_user_message_id,omitempty"`
	MediaRef                  *MediaRefV1         `json:"media_ref,omitempty"`
	Attrs                     *MessageAttrsV1     `json:"attrs,omitempty"`
	ForwardRef                *ForwardRefV1       `json:"forward_ref,omitempty"`
	ServiceAction             *ServiceActionRefV1 `json:"service_action,omitempty"`
	ClearDraft                bool                `json:"clear_draft,omitempty"`
	SourcePermAuthKeyID       int64               `json:"source_perm_auth_key_id,omitempty"`
	ClearDraftBeforeDate      int32               `json:"clear_draft_before_date,omitempty"`
}

type ChatParticipantFactV1 struct {
	UserID        int64  `json:"user_id"`
	Role          string `json:"role"`
	InviterUserID int64  `json:"inviter_user_id,omitempty"`
	Date          int32  `json:"date,omitempty"`
}

type ChatParticipantsChangedFactV1 struct {
	SchemaVersion int                     `json:"schema_version"`
	ChatID        int64                   `json:"chat_id"`
	ActorUserID   int64                   `json:"actor_user_id"`
	Version       int32                   `json:"version"`
	Participants  []ChatParticipantFactV1 `json:"participants"`
}

type MessageOperationV4 struct {
	SchemaVersion int              `json:"schema_version"`
	OperationKind string           `json:"operation_kind"`
	MessageFact   NewMessageFactV1 `json:"message_fact"`
	AttachFacts   []UpdateFactV1   `json:"attach_facts,omitempty"`
}

type OperationResponseV1 struct {
	SchemaVersion int    `json:"schema_version"`
	OperationID   string `json:"operation_id,omitempty"`
	Pts           int64  `json:"pts"`
	PtsCount      int32  `json:"pts_count"`
	EventType     string `json:"event_type,omitempty"`
}

type OperationResponseV2 struct {
	SchemaVersion int    `json:"schema_version"`
	OperationID   string `json:"operation_id,omitempty"`
	Pts           int64  `json:"pts"`
	PtsCount      int32  `json:"pts_count"`
	EventType     string `json:"event_type,omitempty"`
	UserMessageID int64  `json:"user_message_id,omitempty"`
}

type OperationResponseV3 struct {
	SchemaVersion       int    `json:"schema_version"`
	OperationID         string `json:"operation_id,omitempty"`
	Pts                 int64  `json:"pts"`
	PtsCount            int32  `json:"pts_count"`
	EventType           string `json:"event_type,omitempty"`
	UserMessageID       int64  `json:"user_message_id,omitempty"`
	ClientRandomID      int64  `json:"client_random_id,omitempty"`
	ReplyEnvelope       []byte `json:"reply_envelope,omitempty"`
	ReplyEnvelopeCodec  int32  `json:"reply_envelope_codec,omitempty"`
	ReplyEnvelopeSchema int32  `json:"reply_envelope_schema,omitempty"`
}

type MessageEventV1 struct {
	SchemaVersion      int               `json:"schema_version"`
	EventKind          string            `json:"event_kind"`
	CanonicalMessageID int64             `json:"canonical_message_id"`
	MessageID          int64             `json:"message_id"`
	PeerType           int32             `json:"peer_type"`
	PeerID             int64             `json:"peer_id"`
	FromUserID         int64             `json:"from_user_id"`
	ToUserID           int64             `json:"to_user_id"`
	Date               int32             `json:"date"`
	EditDate           int32             `json:"edit_date,omitempty"`
	EditVersion        int32             `json:"edit_version,omitempty"`
	Out                bool              `json:"out"`
	MessageText        string            `json:"message_text"`
	Entities           []MessageEntityV1 `json:"entities,omitempty"`
	ReplyToPeerSeq     int64             `json:"reply_to_peer_seq,omitempty"`
	AuthKeyIdExclude   *int64            `json:"auth_key_id_exclude,omitempty"`
}

type MessageEventV2 struct {
	SchemaVersion        int               `json:"schema_version"`
	EventKind            string            `json:"event_kind"`
	CanonicalMessageID   int64             `json:"canonical_message_id"`
	PeerSeq              int64             `json:"peer_seq,omitempty"`
	MessageID            int64             `json:"message_id"`
	PeerType             int32             `json:"peer_type"`
	PeerID               int64             `json:"peer_id"`
	FromUserID           int64             `json:"from_user_id"`
	ToUserID             int64             `json:"to_user_id"`
	Date                 int32             `json:"date"`
	EditDate             int32             `json:"edit_date,omitempty"`
	EditVersion          int32             `json:"edit_version,omitempty"`
	Out                  bool              `json:"out"`
	MessageText          string            `json:"message_text"`
	Entities             []MessageEntityV1 `json:"entities,omitempty"`
	ReplyToUserMessageID int64             `json:"reply_to_user_message_id,omitempty"`
	ReadMaxUserMessageID int64             `json:"read_max_user_message_id,omitempty"`
	DeleteUserMessageIDs []int64           `json:"delete_user_message_ids,omitempty"`
	PinnedUserMessageID  int64             `json:"pinned_user_message_id,omitempty"`
	AuthKeyIdExclude     *int64            `json:"auth_key_id_exclude,omitempty"`
}

type MessageEventV3 struct {
	SchemaVersion        int                 `json:"schema_version"`
	EventKind            string              `json:"event_kind"`
	CanonicalMessageID   int64               `json:"canonical_message_id"`
	PeerSeq              int64               `json:"peer_seq,omitempty"`
	MessageID            int64               `json:"message_id"`
	PeerType             int32               `json:"peer_type"`
	PeerID               int64               `json:"peer_id"`
	FromUserID           int64               `json:"from_user_id"`
	ToUserID             int64               `json:"to_user_id"`
	Date                 int32               `json:"date"`
	EditDate             int32               `json:"edit_date,omitempty"`
	EditVersion          int32               `json:"edit_version,omitempty"`
	Out                  bool                `json:"out"`
	MessageText          string              `json:"message_text"`
	Entities             []MessageEntityV1   `json:"entities,omitempty"`
	ReplyToUserMessageID int64               `json:"reply_to_user_message_id,omitempty"`
	ReadMaxUserMessageID int64               `json:"read_max_user_message_id,omitempty"`
	DeleteUserMessageIDs []int64             `json:"delete_user_message_ids,omitempty"`
	PinnedUserMessageID  int64               `json:"pinned_user_message_id,omitempty"`
	AuthKeyIdExclude     *int64              `json:"auth_key_id_exclude,omitempty"`
	MediaRef             *MediaRefV1         `json:"media_ref,omitempty"`
	Attrs                *MessageAttrsV1     `json:"attrs,omitempty"`
	ForwardRef           *ForwardRefV1       `json:"forward_ref,omitempty"`
	ServiceAction        *ServiceActionRefV1 `json:"service_action,omitempty"`
}

type MessageEventV4 struct {
	SchemaVersion    int              `json:"schema_version"`
	EventKind        string           `json:"event_kind"`
	MessageFact      NewMessageFactV1 `json:"message_fact"`
	AttachFacts      []UpdateFactV1   `json:"attach_facts,omitempty"`
	MessageID        int64            `json:"message_id"`
	Pts              int64            `json:"pts"`
	PtsCount         int32            `json:"pts_count"`
	AuthKeyIdExclude *int64           `json:"auth_key_id_exclude,omitempty"`
}

func WrapFact(kind string, value any) (UpdateFactV1, error) {
	if kind == "" {
		return UpdateFactV1{}, fmt.Errorf("update fact kind is empty")
	}
	if value == nil {
		return UpdateFactV1{}, fmt.Errorf("update fact %s value is nil", kind)
	}
	body, err := json.Marshal(value)
	if err != nil {
		return UpdateFactV1{}, fmt.Errorf("marshal update fact %s: %w", kind, err)
	}
	if err := validateUpdateFactPayload(kind, body); err != nil {
		return UpdateFactV1{}, err
	}
	return UpdateFactV1{
		SchemaVersion: 1,
		Kind:          kind,
		Payload:       body,
	}, nil
}

func DecodeUpdateFact(f UpdateFactV1) (any, error) {
	if err := validateUpdateFactPayload(f.Kind, f.Payload); err != nil {
		return nil, err
	}
	switch f.Kind {
	case FactKindChatParticipantsChanged:
		var fact ChatParticipantsChangedFactV1
		if err := json.Unmarshal(f.Payload, &fact); err != nil {
			return nil, fmt.Errorf("decode update fact %s: %w", f.Kind, err)
		}
		return fact, nil
	case FactKindNewMessage:
		var fact NewMessageFactV1
		if err := json.Unmarshal(f.Payload, &fact); err != nil {
			return nil, fmt.Errorf("decode update fact %s: %w", f.Kind, err)
		}
		return fact, nil
	default:
		return nil, fmt.Errorf("unsupported update fact kind: %s", f.Kind)
	}
}

func validateUpdateFactPayload(kind string, payload []byte) error {
	trimmed := bytes.TrimSpace(payload)
	if len(trimmed) == 0 {
		return fmt.Errorf("update fact %s payload is empty", kind)
	}
	if bytes.Equal(trimmed, []byte("null")) {
		return fmt.Errorf("update fact %s payload is null", kind)
	}
	return nil
}

func HashBytes(b []byte) []byte {
	sum := sha256.Sum256(b)
	out := make([]byte, len(sum))
	copy(out, sum[:])
	return out
}

func HashHex(b []byte) string {
	return hex.EncodeToString(HashBytes(b))
}

func SenderOperationID(canonicalMessageID, senderUserID int64) string {
	return fmt.Sprintf("v1:msg:%d:sender:%d:out", canonicalMessageID, senderUserID)
}

func ReceiverOperationID(canonicalMessageID, receiverUserID int64) string {
	return fmt.Sprintf("v1:msg:%d:receiver:%d:in", canonicalMessageID, receiverUserID)
}
