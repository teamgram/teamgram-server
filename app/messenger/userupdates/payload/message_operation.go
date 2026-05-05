package payload

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type MessageOperationV1 struct {
	SchemaVersion         int               `json:"schema_version"`
	OperationKind         string            `json:"operation_kind"`
	CanonicalMessageID    int64             `json:"canonical_message_id"`
	PeerType              int32             `json:"peer_type"`
	PeerID                int64             `json:"peer_id"`
	PeerSeq               int64             `json:"peer_seq"`
	FromUserID            int64             `json:"from_user_id"`
	ToUserID              int64             `json:"to_user_id"`
	Date                  int32             `json:"date"`
	Out                   bool              `json:"out"`
	MessageText           string            `json:"message_text"`
	Entities              []MessageEntityV1 `json:"entities,omitempty"`
	DependencyPts         []int64           `json:"dependency_pts,omitempty"`
	ClearDraft            bool              `json:"clear_draft,omitempty"`
	SourcePermAuthKeyID   int64             `json:"source_perm_auth_key_id,omitempty"`
	ClearDraftBeforeDate  int32             `json:"clear_draft_before_date,omitempty"`
	SavedDialogSideEffect bool              `json:"saved_dialog_side_effect,omitempty"`
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
	Out                bool              `json:"out"`
	MessageText        string            `json:"message_text"`
	Entities           []MessageEntityV1 `json:"entities,omitempty"`
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
