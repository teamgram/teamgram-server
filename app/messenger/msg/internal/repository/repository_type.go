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

	MessageKindText   int32 = 1
	MessageStatusLive int32 = 1
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
	SendStateID        int64
	SenderUserID       int64
	PeerType           int32
	PeerID             int64
	ClientRandomID     int64
	RequestPayloadHash []byte
	MessageText        string
	MessageDate        int64
}

type CanonicalMessageResult struct {
	SendStateID        int64
	CanonicalMessageID int64
	PeerSeq            int64
	MessageDate        int64
	RequestPayloadHash []byte
	CreatedNew         bool
}

type ListHistoryMessagesInput struct {
	UserID    int64
	PeerType  int32
	PeerID    int64
	OffsetID  int32
	AddOffset int32
	MaxID     int32
	MinID     int32
	Limit     int32
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
	CanonicalMessageID int64
	PeerSeq            int64
	ReplyToPeerSeq     int64
	FromUserID         int64
	Outgoing           bool
	PeerType           int32
	PeerID             int64
	MessageKind        int32
	MessageText        string
	MessageDate        int64
}

type CanonicalMessage struct {
	CanonicalMessageID int64
	PeerSeq            int64
	FromUserID         int64
	PeerType           int32
	PeerID             int64
	MessageKind        int32
	MessageText        string
	MessageDate        int64
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
	GetCanonicalMessageByPeerSeq(ctx context.Context, userID int64, peerType int32, peerID int64, peerSeq int64) (*CanonicalMessage, error)
	EditCanonicalMessage(ctx context.Context, in EditCanonicalMessageInput) (*EditMessageResult, error)
	ListHistoryMessages(ctx context.Context, in ListHistoryMessagesInput) ([]HistoryMessage, error)
}

type MessageSendStateRepository interface {
	CreateOrLoadSendState(ctx context.Context, in CreateSendStateInput) (*SendState, error)
	MarkCanonicalCreated(ctx context.Context, sendStateID int64, canonicalMessageID int64, peerSeq int64) error
	MarkSenderCommitted(ctx context.Context, in MarkSenderCommittedInput) error
	MarkReceiverOpsAcked(ctx context.Context, sendStateID int64, receiverManifestID int64) error
	MarkCompleted(ctx context.Context, sendStateID int64) error
	MarkRetryableFailure(ctx context.Context, in MarkRetryableFailureInput) error
}
