// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
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
	PreferenceScopeMainPin          = "main_pin"
	PreferenceScopeFolderPin        = "folder_pin"
	PreferenceScopeFolderMembership = "folder_membership"

	TargetAuthPolicyAll                  = "all"
	TargetAuthPolicyNotSourcePermAuthKey = "not_source_perm_auth_key"

	OutboxStatusPending         int32 = 1
	OutboxStatusPublishing      int32 = 2
	OutboxStatusPublished       int32 = 3
	OutboxStatusFailedRetryable int32 = 4
	OutboxStatusBlocked         int32 = 5

	DeliveryPathUserupdatesPTS     = "userupdates_pts"
	DeliveryPathUserupdatesAuthSeq = "userupdates_auth_seq"

	DefaultOutboxWorkerBatchSize    = 100
	DefaultOutboxWorkerLeaseSeconds = 30
	DefaultOutboxWorkerPollSeconds  = 1
	InitialRetryDelaySeconds        = 1
	OutboxWorkerMaxRetryDelay       = 300
	OutboxWorkerBlockedAttempts     = 20
	OutboxWorkerBlockedAgeSeconds   = 3600
)

type ToggleDialogPinInput struct {
	UserID              int64
	PeerType            int32
	PeerID              int64
	FolderID            int32
	Pinned              bool
	PinOrder            int64
	SourcePermAuthKeyID int64
	OperationID         string
	OutboxID            int64
	EventType           string
	Payload             []byte
}

type ReorderPinnedDialogsInput struct {
	UserID              int64
	FolderID            int32
	PeerOrder           []PeerRef
	SourcePermAuthKeyID int64
	OperationID         string
	OutboxID            int64
	EventType           string
	Payload             []byte
}

type EditPeerFoldersInput struct {
	UserID              int64
	PeerType            int32
	PeerID              int64
	OldFolderID         int32
	NewFolderID         int32
	SourcePermAuthKeyID int64
	OperationID         string
	OutboxID            int64
	PublicUpdateType    string
	Payload             []byte
}

type PreferenceMutationResult struct {
	UserID           int64
	PeerDialogID     int64
	AggregateVersion int64
}

type SaveDraftInput struct {
	UserID              int64
	PeerType            int32
	PeerID              int64
	DraftKind           int32
	Message             string
	EntitiesPayload     []byte
	ReplyToPeerSeq      int64
	DraftPayload        []byte
	Date                time.Time
	SourcePermAuthKeyID int64
	OperationID         string
	OutboxID            int64
	EventType           string
}

type ClearDraftInput struct {
	UserID              int64
	PeerType            int32
	PeerID              int64
	SourcePermAuthKeyID int64
	OperationID         string
	OutboxID            int64
	EventType           string
}

type ClearDraftAfterSendInput struct {
	UserID              int64
	PeerType            int32
	PeerID              int64
	ClearBeforeDate     time.Time
	SourcePermAuthKeyID int64
	OperationID         string
	OutboxID            int64
	EventType           string
	Payload             []byte
}

type ClearAllDraftsInput struct {
	UserID              int64
	SourcePermAuthKeyID int64
	OperationID         string
	OutboxIDs           []int64
}

type DraftMutationResult struct {
	UserID       int64
	PeerDialogID int64
	Cleared      bool
}

type DraftRecord struct {
	UserID          int64
	PeerType        int32
	PeerID          int64
	PeerDialogID    int64
	DraftKind       int32
	Message         string
	EntitiesPayload []byte
	ReplyToPeerSeq  int64
	DraftPayload    []byte
	Date            time.Time
}

type SavedDialogTopInput struct {
	UserID                int64
	PeerType              int32
	PeerID                int64
	TopPeerSeq            int64
	TopCanonicalMessageID int64
	TopMessageDate        time.Time
	Payload               []byte
}

type SavedDialogPinInput struct {
	UserID              int64
	PeerType            int32
	PeerID              int64
	Pinned              bool
	PinOrder            int64
	SourcePermAuthKeyID int64
	OperationID         string
	OutboxID            int64
	EventType           string
	Payload             []byte
}

type SavedDialogRecord struct {
	UserID                int64
	PeerType              int32
	PeerID                int64
	TopPeerSeq            int64
	TopCanonicalMessageID int64
	TopMessageDate        time.Time
	Pinned                bool
	PinOrder              int64
	SavedPayload          []byte
}

type ReorderPinnedSavedDialogsInput struct {
	UserID              int64
	Order               []PeerRef
	SourcePermAuthKeyID int64
	OperationID         string
	OutboxID            int64
	EventType           string
	Payload             []byte
}

type PrivatePeerPolicyInput struct {
	UserID              int64
	PeerUserID          int64
	TTLPeriod           int32
	ThemeEmoticon       string
	SourcePermAuthKeyID int64
	OperationID         string
	ActorOutboxID       int64
	PeerOutboxID        int64
	DeliveryPath        string
	PublicUpdateType    string
	Payload             []byte
}

type PrivatePeerPolicyResult struct {
	Scope PolicyScope
}

type PeerWallpaperInput struct {
	UserID              int64
	PeerType            int32
	PeerID              int64
	WallpaperID         int64
	WallpaperOverridden bool
	SourcePermAuthKeyID int64
	OperationID         string
	OutboxID            int64
	EventType           string
	Payload             []byte
}

type DialogFilterRecord struct {
	UserID              int64
	DialogFilterID      int32
	Slug                string
	Title               string
	OrderValue          int64
	Enabled             bool
	FilterSchemaVersion int32
	FilterPayload       []byte
}

type SaveDialogFilterInput struct {
	UserID              int64
	DialogFilterID      int32
	Slug                string
	Title               string
	OrderValue          int64
	Enabled             bool
	FilterSchemaVersion int32
	FilterPayload       []byte
	SourcePermAuthKeyID int64
	OperationID         string
	OutboxID            int64
	EventType           string
}

type DeleteDialogFilterInput struct {
	UserID              int64
	DialogFilterID      int32
	SourcePermAuthKeyID int64
	OperationID         string
	OutboxID            int64
	EventType           string
}

type ReorderDialogFiltersInput struct {
	UserID              int64
	Order               []int32
	SourcePermAuthKeyID int64
	OperationID         string
	OutboxID            int64
	EventType           string
}
