// Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package repository

import "github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/eventtypes"

const (
	PayloadCodecJSON int32 = eventtypes.PayloadCodecJSON

	OpTypeSendMessage int32 = 1

	DeliveryPolicyRequesterSync      int32 = 1
	DeliveryPolicyBrokerDurableAck   int32 = 2
	DeliveryPolicyDurableAsync       int32 = 3
	DeliveryPolicyApplyAck           int32 = 4
	DeliveryPolicyBestEffortPushOnly int32 = 5

	EventTypeNewMessage              int32 = eventtypes.EventTypeNewMessage
	EventTypeReadHistory             int32 = eventtypes.EventTypeReadHistory
	EventTypeUpdatePinnedMessage     int32 = eventtypes.EventTypeUpdatePinnedMessage
	EventTypeMarkDialogUnread        int32 = eventtypes.EventTypeMarkDialogUnread
	EventTypeScheduledMarker         int32 = eventtypes.EventTypeScheduledMarker
	EventTypeDeleteMessages          int32 = eventtypes.EventTypeDeleteMessages
	EventTypeDeleteHistory           int32 = eventtypes.EventTypeDeleteHistory
	EventTypeEditMessage             int32 = eventtypes.EventTypeEditMessage
	EventTypeChatParticipantsChanged int32 = eventtypes.EventTypeChatParticipantsChanged
	EventTypeDialogPublicUpdate      int32 = eventtypes.EventTypeDialogPublicUpdate

	OperationResultStatusCompleted int32 = 1

	MessageKindText      int32 = 1
	MessageKindMedia     int32 = 2
	MessageKindService   int32 = 3
	MessageStatusLive    int32 = 1
	MessageStatusDeleted int32 = 2

	PushTypeUserUpdate    int32 = 1
	PushTaskStatusPending int32 = 1
)

const (
	AuthSeqReplayPolicyDurableReplay = "durable_replay"
	AuthSeqReplayPolicyTTLReplay     = "ttl_replay"
	AuthSeqReplayPolicyRealtimeOnly  = "realtime_only"

	AuthSeqVisibilityAllUserAuthKeys        = "all_user_auth_keys"
	AuthSeqVisibilityNotSourcePermAuthKey   = "not_source_perm_auth_key"
	AuthSeqVisibilityExplicitPermAuthKeySet = "explicit_perm_auth_key_set"

	AuthSeqCodecTLBinary        int32 = 1
	AuthSeqLayer                int32 = 224
	AuthSeqMaxFutureSkewSeconds int64 = 86400
	AuthSeqPhoneCallTTLSeconds  int64 = 60
)

const (
	PushTaskStatusPublishing      int32 = 2
	PushTaskStatusPublished       int32 = 3
	PushTaskStatusFailedRetryable int32 = 4
	PushTaskStatusFailedTerminal  int32 = 5
)

const (
	AffectedOutboxStatusPending        int32 = 1
	AffectedOutboxStatusProcessing     int32 = 2
	AffectedOutboxStatusCompleted      int32 = 3
	AffectedOutboxStatusFailedTerminal int32 = 4
)

const (
	DialogSideEffectKindClearDraftAfterSend          = "clear_draft_after_send"
	DialogSideEffectKindUpsertSavedDialogFromMessage = "upsert_saved_dialog_from_message"
)

const (
	DialogSideEffectStatusPending         int32 = 1
	DialogSideEffectStatusPublishing      int32 = 2
	DialogSideEffectStatusCompleted       int32 = 3
	DialogSideEffectStatusFailedRetryable int32 = 4
	DialogSideEffectStatusBlocked         int32 = 5
)

const (
	InitialRetryDelaySeconds = 1
	MaxRetryDelaySeconds     = 300
	BlockedAfterAttempts     = 20
	BlockedAfterAgeSeconds   = 3600
	OutboxWorkerBatchSize    = 100
)

const (
	DeliveryFailedOperationStatusOpen      int32 = 1
	DeliveryFailedOperationStatusReplaying int32 = 2
	DeliveryFailedOperationStatusReplayed  int32 = 3
	DeliveryFailedOperationStatusIgnored   int32 = 4
	DeliveryFailedOperationStatusTerminal  int32 = 5
)

const (
	FailureCategoryControlFlow    int32 = 1
	FailureCategoryInfrastructure int32 = 2
	FailureCategoryCorruption     int32 = 3
)

type KafkaAck struct {
	Topic     string
	Partition int32
	Offset    int64
}

type ReceiverKafkaRecord struct {
	Topic     string
	Partition int32
	Offset    int64
	Key       []byte
	Value     []byte
}

type PushTask struct {
	TaskID          int64
	UserID          int64
	Pts             int64
	PushType        int32
	PeerType        int32
	PeerID          int64
	OperationID     string
	PushPartitionID int32
	TaskPayload     []byte
}

type DialogSideEffect struct {
	SideEffectID             int64
	Kind                     string
	UserID                   int64
	PeerType                 int32
	PeerID                   int64
	SourcePermAuthKeyID      int64
	SourceOperationID        string
	SourceMessageDate        int64
	SourcePeerSeq            int64
	SourceCanonicalMessageID int64
	ClearBeforeDate          int64
	PayloadSchemaVersion     int32
	Payload                  []byte
	PayloadHash              []byte
	Status                   int32
	AttemptCount             int32
	NextRetryAt              int64
	LeaseOwner               string
	LeaseUntil               int64
	LastErrorCode            string
}

type AffectedOutbox struct {
	RequesterUserID   int64
	TargetUserID      int64
	TargetBucketID    int32
	TargetPartitionID int32
	OperationID       string
	OpType            int32
	OperationKind     string
	PeerType          int32
	PeerID            int64
	PayloadCodec      int32
	Payload           []byte
	PayloadHash       []byte
	DeliveryPolicy    int32
	OwnerTokenPayload []byte
}

type ApplyUserOperationInput struct {
	UserID           int64
	OperationID      string
	OpType           int32
	PeerType         int32
	PeerID           int64
	PayloadCodec     int32
	Payload          []byte
	PayloadHash      []byte
	BucketID         int32
	PartitionID      int32
	DependencyPts    []int64
	AuthKeyIDExclude *int64
	AffectedOutboxes []AffectedOutbox
}

type ApplyUserOperationResult struct {
	UserID                int64
	OperationID           string
	Pts                   int64
	PtsCount              int32
	ResponseSchemaVersion int32
	ResponsePayload       []byte
	ResponseHash          []byte
	AlreadyApplied        bool
}

type OperationResult struct {
	UserID                int64
	OperationID           string
	OpType                int32
	Status                int32
	Pts                   int64
	PtsCount              int32
	PayloadHash           []byte
	ResponseSchemaVersion int32
	ResponsePayload       []byte
	ResponseHash          []byte
	TerminalErrorCode     string
}

type UserState struct {
	UserID      int64
	Pts         int64
	Seq         int64
	Date        int32
	UnreadCount int32
	PartitionID int32
	OwnerEpoch  int64
	RowVersion  int64
}

type UserEvent = eventtypes.UserEvent

type GetDifferenceInput struct {
	UserID        int64
	PermAuthKeyID int64
	Pts           int64
	Limit         int32
	Date          *int64
}

type OutboxReadDateInput struct {
	UserID   int64
	PeerType int32
	PeerID   int64
	MsgID    int32
}

type GetDifferenceResult struct {
	State         UserState
	Events        []UserEvent
	AuthSeqEvents []AuthSeqEvent
}

type DialogSideEffectAppendInput struct {
	UserID               int64
	SourcePermAuthKeyID  int64
	OperationID          string
	TargetAuthPolicy     string
	PublicUpdateType     string
	PeerType             int32
	PeerID               int64
	PayloadSchemaVersion int32
	Payload              []byte
	PayloadHash          []byte
}

type AuthSeqUpdateAppendInput struct {
	UserID               int64
	SourcePermAuthKeyID  int64
	TargetPermAuthKeyIDs []int64
	OperationID          string
	UpdateType           string
	ReplayPolicy         string
	VisibilityPolicy     string
	Layer                int32
	TLBytes              []byte
	PayloadHash          []byte
	ExpireAt             int64
	Now                  int64
}

type AuthSeqDeliveryEvent struct {
	UserID              int64
	PermAuthKeyID       int64
	Seq                 int64
	Date                int32
	PayloadID           string
	ReplayPolicy        string
	OperationID         string
	SourcePermAuthKeyID int64
	VisibilityPolicy    string
	TLBytes             []byte
	PayloadHash         []byte
	Layer               int32
}

type AuthSeqAppendResult struct {
	UserID         int64
	OperationID    string
	Seq            int64
	Date           int32
	AlreadyApplied bool
}

type PtsAppendResult struct {
	UserID         int64
	OperationID    string
	Pts            int64
	PtsCount       int32
	AlreadyApplied bool
}

type AuthSeqEvent struct {
	UserID              int64
	Seq                 int64
	Date                int32
	OperationID         string
	SourcePermAuthKeyID int64
	TargetAuthPolicy    string
	PublicUpdateType    string
	PeerType            int32
	PeerID              int64
	EventSchemaVersion  int32
	EventCodec          int32
	EventPayload        []byte
	EventPayloadHash    []byte
}

type RecordDeliveryFailureInput struct {
	FailedId        int64
	UserId          int64
	OperationId     string
	OpType          int32
	BucketId        int32
	KafkaTopic      string
	KafkaPartition  int32
	KafkaOffset     int64
	PayloadHash     []byte
	FailureCategory int32
	FailureCode     string
	FailureMessage  string
}
