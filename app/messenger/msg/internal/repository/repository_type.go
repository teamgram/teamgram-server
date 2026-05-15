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

import "context"

const (
	SendStateStatusInitialized     int32 = 1
	SendStateStatusCanonical       int32 = 2
	SendStateStatusSenderCommitted int32 = 3
	SendStateStatusReceiverAcked   int32 = 4
	SendStateStatusCompleted       int32 = 5
	SendStateStatusFailedRetryable int32 = 6

	MessageKindText    int32 = 1
	MessageKindMedia   int32 = 2
	MessageKindService int32 = 3
	MessageStatusLive  int32 = 1
)

type KafkaAck struct {
	Topic     string
	Partition int32
	Offset    int64
}

type CreateSendStateInput struct {
	SenderUserID                int64
	PeerType                    int32
	PeerID                      int64
	ClientRandomID              int64
	RequestPayloadSchemaVersion int32
	RequestPayloadHash          []byte
	MessageText                 string
	ReplyToCanonicalMessageID   int64
}

type SendState struct {
	SendStateID                 int64
	SenderUserID                int64
	PeerType                    int32
	PeerID                      int64
	ClientRandomID              int64
	CanonicalMessageID          int64
	PeerSeq                     int64
	Status                      int32
	RequestPayloadSchemaVersion int32
	RequestPayloadHash          []byte
	SenderOperationID           string
	SenderPTS                   int64
	SenderPTSCount              int32
	SenderUpdateSchemaVersion   int32
	SenderUpdatePayload         []byte
	SenderUpdatePayloadHash     []byte
	ReceiverManifestID          int64
	RetryCount                  int32
}

type CreateCanonicalMessageInput struct {
	SendStateID                  int64
	SenderUserID                 int64
	PeerType                     int32
	PeerID                       int64
	ClientRandomID               int64
	RequestPayloadHash           []byte
	MessageText                  string
	MessageDate                  int64
	EntitiesPayloadSchemaVersion int32
	EntitiesPayload              []byte
	MediaRefSchemaVersion        int32
	MediaRefPayload              []byte
	MessageAttrsSchemaVersion    int32
	MessageAttrsPayload          []byte
	ForwardRefSchemaVersion      int32
	ForwardRefPayload            []byte
	ServiceActionSchemaVersion   int32
	ServiceActionPayload         []byte
}

type CreateCanonicalBatchInput struct {
	SenderUserID int64
	PeerType     int32
	PeerID       int64
	Items        []CreateCanonicalBatchItem
}

type CreateCanonicalBatchItem struct {
	ClientRandomID               int64
	RequestPayloadSchemaVersion  int32
	RequestPayloadHash           []byte
	MessageText                  string
	MessageDate                  int64
	MediaRefSchemaVersion        int32
	MediaRefPayload              []byte
	EntitiesPayloadSchemaVersion int32
	EntitiesPayload              []byte
	MessageAttrsSchemaVersion    int32
	MessageAttrsPayload          []byte
	ForwardRefSchemaVersion      int32
	ForwardRefPayload            []byte
	ServiceActionSchemaVersion   int32
	ServiceActionPayload         []byte
}

type CanonicalMessageResult struct {
	SendStateID                  int64
	CanonicalMessageID           int64
	PeerSeq                      int64
	MessageDate                  int64
	RequestPayloadHash           []byte
	SendStateStatus              int32
	SenderOperationID            string
	SenderPTS                    int64
	SenderPTSCount               int32
	SenderUpdateSchemaVersion    int32
	SenderUpdatePayload          []byte
	SenderUpdatePayloadHash      []byte
	EntitiesPayloadSchemaVersion int32
	EntitiesPayload              []byte
	MediaRefSchemaVersion        int32
	MediaRefPayload              []byte
	MessageAttrsSchemaVersion    int32
	MessageAttrsPayload          []byte
	ForwardRefSchemaVersion      int32
	ForwardRefPayload            []byte
	ServiceActionSchemaVersion   int32
	ServiceActionPayload         []byte
	CreatedNew                   bool
}

type CanonicalBatchResult struct {
	Items []CanonicalMessageResult
}

type ListHistoryMessagesInput struct {
	UserID               int64
	PeerType             int32
	PeerID               int64
	OffsetID             int32
	AddOffset            int32
	MaxID                int32
	MinID                int32
	Limit                int32
	CursorsResolved      bool
	ResolvedCursorBounds HistoryCursorBounds
}

type SearchHashTagMessagesInput struct {
	UserID   int64
	PeerType int32
	PeerID   int64
	HashTag  string
	OffsetID int32
	Limit    int32
}

type HistoryMessage struct {
	CanonicalMessageID   int64
	PeerSeq              int64
	UserMessageID        int64
	ReplyToPeerSeq       int64
	ReplyToUserMessageID int64
	FromUserID           int64
	Outgoing             bool
	PeerType             int32
	PeerID               int64
	MessageKind          int32
	MessageText          string
	MessageDate          int64
	ViewPayload          []byte
}

type UserMessageBox struct {
	UserID             int64
	UserMessageID      int64
	CanonicalMessageID int64
	PeerType           int32
	PeerID             int64
	PeerSeq            int64
	FromUserID         int64
	Outgoing           bool
	MessageText        string
	MessageDate        int64
	ViewPayload        []byte
}

type ForwardSourceIdentity struct {
	UserID             int64
	UserMessageID      int64
	CanonicalMessageID int64
}

type ForwardSourceLookup struct {
	UserID              int64
	SourcePeerType      int32
	SourcePeerID        int64
	SourceUserMessageID int64
}

type CanonicalMessage struct {
	CanonicalMessageID           int64
	PeerSeq                      int64
	FromUserID                   int64
	PeerType                     int32
	PeerID                       int64
	MessageKind                  int32
	MessageText                  string
	MessageDate                  int64
	EntitiesPayloadSchemaVersion int32
	EntitiesPayload              []byte
	MediaRefSchemaVersion        int32
	MediaRefPayload              []byte
	MessageAttrsSchemaVersion    int32
	MessageAttrsPayload          []byte
	ForwardRefSchemaVersion      int32
	ForwardRefPayload            []byte
	ServiceActionSchemaVersion   int32
	ServiceActionPayload         []byte
}

type EditCanonicalMessageInput struct {
	ActorUserID     int64
	PeerType        int32
	PeerID          int64
	PeerSeq         int64
	NewMessageText  string
	RequestEditDate int64
}

type EditMessageResult struct {
	CanonicalMessageID int64
	PeerSeq            int64
	FromUserID         int64
	PeerType           int32
	PeerID             int64
	MessageKind        int32
	MessageText        string
	MessageDate        int64
	EditDate           int64
	EditVersion        int32
}

type MarkSenderCommittedInput struct {
	SendStateID               int64
	SenderOperationID         string
	SenderPTS                 int64
	SenderPTSCount            int32
	SenderUpdateSchemaVersion int32
	SenderUpdatePayload       []byte
	SenderUpdatePayloadHash   []byte
}

type MarkRetryableFailureInput struct {
	SendStateID       int64
	LastErrorCategory int32
	LastErrorCode     string
	LastErrorMessage  string
}

type MessageRepository interface {
	CreateOrGetByClientRandom(ctx context.Context, in CreateCanonicalMessageInput) (*CanonicalMessageResult, error)
	CreateOrGetCanonicalBatchByClientRandom(ctx context.Context, in CreateCanonicalBatchInput) (*CanonicalBatchResult, error)
	GetCanonicalMessageByPeerSeq(ctx context.Context, userID int64, peerType int32, peerID int64, peerSeq int64) (*CanonicalMessage, error)
	ListRecentCanonicalMessagesBeforePeerSeq(ctx context.Context, peerType int32, peerID int64, beforePeerSeq int64, limit int32) ([]CanonicalMessage, error)
	EditCanonicalMessage(ctx context.Context, in EditCanonicalMessageInput) (*EditMessageResult, error)
	ListHistoryMessages(ctx context.Context, in ListHistoryMessagesInput) ([]HistoryMessage, error)
	GetUserMessage(ctx context.Context, userID int64, userMessageID int64) (*UserMessageBox, error)
	GetUserMessageList(ctx context.Context, userID int64, ids []int64) ([]UserMessageBox, error)
	ResolveForwardSourceIdentity(ctx context.Context, lookup ForwardSourceLookup) (*ForwardSourceIdentity, error)
	RevalidateForwardSources(ctx context.Context, sources []ForwardSourceIdentity) error
	ResolveMessageID(ctx context.Context, userID int64, peerType int32, peerID int64, userMessageID int64) (*ResolvedMessageID, error)
	ResolveMessageIDs(ctx context.Context, userID int64, userMessageIDs []int64) ([]ResolvedMessageID, error)
	ResolveMessageIDsForDelete(ctx context.Context, userID int64, userMessageIDs []int64) ([]ResolvedMessageID, error)
	ResolveHistoryCursorIDs(ctx context.Context, userID int64, peerType int32, peerID int64, offsetID int32, maxID int32, minID int32) (HistoryCursorBounds, error)
	ResolvePeerSeqToUserMessageID(ctx context.Context, userID int64, peerType int32, peerID int64, peerSeq int64) (int64, error)
}

type MessageSendStateRepository interface {
	CreateOrLoadSendState(ctx context.Context, in CreateSendStateInput) (*SendState, error)
	MarkCanonicalCreated(ctx context.Context, sendStateID int64, canonicalMessageID int64, peerSeq int64) error
	MarkSenderCommitted(ctx context.Context, in MarkSenderCommittedInput) error
	MarkReceiverOpsAcked(ctx context.Context, sendStateID int64, receiverManifestID int64) error
	MarkCompleted(ctx context.Context, sendStateID int64) error
	MarkRetryableFailure(ctx context.Context, in MarkRetryableFailureInput) error
}
