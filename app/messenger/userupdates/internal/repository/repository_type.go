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

const (
	PayloadCodecJSON int32 = 1

	OpTypeSendMessage int32 = 1

	EventTypeNewMessage int32 = 1

	OperationResultStatusCompleted int32 = 1

	MessageKindText   int32 = 1
	MessageStatusLive int32 = 1

	PushTypeUserUpdate    int32 = 1
	PushTaskStatusPending int32 = 1
)

type ApplyUserOperationInput struct {
	UserID        int64
	OperationID   string
	OpType        int32
	PeerType      int32
	PeerID        int64
	PayloadCodec  int32
	Payload       []byte
	PayloadHash   []byte
	BucketID      int32
	PartitionID   int32
	DependencyPts []int64
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
	UserID int64
	Pts    int64
	Limit  int32
}

type GetDifferenceResult struct {
	State  UserState
	Events []UserEvent
}
