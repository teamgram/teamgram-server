package payload

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type MessageOperationV1 struct {
	SchemaVersion             int               `json:"schema_version"`
	OperationKind             string            `json:"operation_kind"`
	CanonicalMessageID        int64             `json:"canonical_message_id"`
	PeerType                  int32             `json:"peer_type"`
	PeerID                    int64             `json:"peer_id"`
	PeerSeq                   int64             `json:"peer_seq"`
	FromUserID                int64             `json:"from_user_id"`
	ToUserID                  int64             `json:"to_user_id"`
	Date                      int32             `json:"date"`
	EditDate                  int32             `json:"edit_date,omitempty"`
	EditVersion               int32             `json:"edit_version,omitempty"`
	Out                       bool              `json:"out"`
	MessageText               string            `json:"message_text"`
	Entities                  []MessageEntityV1 `json:"entities,omitempty"`
	ReplyToCanonicalMessageID int64             `json:"reply_to_canonical_message_id,omitempty"`
	ReplyToPeerSeq            int64             `json:"reply_to_peer_seq,omitempty"`
	DependencyPts             []int64           `json:"dependency_pts,omitempty"`
	ClearDraft                bool              `json:"clear_draft,omitempty"`
	SourcePermAuthKeyID       int64             `json:"source_perm_auth_key_id,omitempty"`
	ClearDraftBeforeDate      int32             `json:"clear_draft_before_date,omitempty"`
	SavedDialogSideEffect     bool              `json:"saved_dialog_side_effect,omitempty"`
	ReadInboxMaxPeerSeq       int64             `json:"read_inbox_max_peer_seq,omitempty"`
	ReadOutboxMaxPeerSeq      int64             `json:"read_outbox_max_peer_seq,omitempty"`
	DeletePeerSeqs            []int64           `json:"delete_peer_seqs,omitempty"`
	DeleteMaxPeerSeq          int64             `json:"delete_max_peer_seq,omitempty"`
	JustClear                 bool              `json:"just_clear,omitempty"`
	Revoke                    bool              `json:"revoke,omitempty"`
	UnreadMark                *bool             `json:"unread_mark,omitempty"`
	PinnedPeerSeq             int64             `json:"pinned_peer_seq,omitempty"`
	PinnedCanonicalMessageID  int64             `json:"pinned_canonical_message_id,omitempty"`
	HasScheduled              *bool             `json:"has_scheduled,omitempty"`
}

type MessageEntityV1 struct {
	Offset int32  `json:"offset"`
	Length int32  `json:"length"`
	Kind   string `json:"kind"`
	URL    string `json:"url,omitempty"`
}

type OperationResponseV1 struct {
	SchemaVersion int    `json:"schema_version"`
	OperationID   string `json:"operation_id,omitempty"`
	Pts           int64  `json:"pts"`
	PtsCount      int32  `json:"pts_count"`
	EventType     string `json:"event_type,omitempty"`
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
