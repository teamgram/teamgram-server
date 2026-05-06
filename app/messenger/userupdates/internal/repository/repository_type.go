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

import "time"

const (
	PayloadCodecJSON int32 = 1

	OpTypeSendMessage int32 = 1

	EventTypeNewMessage          int32 = 1
	EventTypeReadHistory         int32 = 2
	EventTypeUpdatePinnedMessage int32 = 3
	EventTypeMarkDialogUnread    int32 = 4
	EventTypeScheduledMarker     int32 = 5
	EventTypeDeleteMessages      int32 = 6
	EventTypeDeleteHistory       int32 = 7
	EventTypeDialogPublicUpdate  int32 = 100

	OperationResultStatusCompleted int32 = 1

	MessageKindText      int32 = 1
	MessageStatusLive    int32 = 1
	MessageStatusDeleted int32 = 2

	PushTypeUserUpdate    int32 = 1
	PushTaskStatusPending int32 = 1
)

const (
	PushTaskStatusPublishing      int32 = 2
	PushTaskStatusPublished       int32 = 3
	PushTaskStatusFailedRetryable int32 = 4
	PushTaskStatusFailedTerminal  int32 = 5
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
	SourceMessageDate        time.Time
	SourcePeerSeq            int64
	SourceCanonicalMessageID int64
	ClearBeforeDate          time.Time
	PayloadSchemaVersion     int32
	Payload                  []byte
	PayloadHash              []byte
	Status                   int32
	AttemptCount             int32
	NextRetryAt              time.Time
	LeaseOwner               string
	LeaseUntil               time.Time
	LastErrorCode            string
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
}

type ApplyUserOperationResult struct {
	UserID          int64
	OperationID     string
	Pts             int64
	PtsCount        int32
	ResponsePayload []byte
	ResponseHash    []byte
	AlreadyApplied  bool
}

type OperationResult struct {
	UserID            int64
	OperationID       string
	OpType            int32
	Status            int32
	Pts               int64
	PtsCount          int32
	PayloadHash       []byte
	ResponsePayload   []byte
	ResponseHash      []byte
	TerminalErrorCode string
}

type UserState struct {
	UserID      int64
	Pts         int64
	Seq         int64
	Date        int32
	PartitionID int32
	OwnerEpoch  int64
	RowVersion  int64
}

type UserEvent struct {
	UserID             int64
	Pts                int64
	PtsCount           int32
	OperationID        string
	OpType             int32
	EventType          int32
	PeerType           int32
	PeerID             int64
	CanonicalMessageID int64
	PeerSeq            int64
	ActorUserID        int64
	EventSchemaVersion int32
	EventCodec         int32
	EventPayload       []byte
	EventPayloadHash   []byte
}

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
