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

package core

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
)

type sendRequestPayloadV3 struct {
	SchemaVersion             int                         `json:"schema_version"`
	SenderUserID              int64                       `json:"sender_user_id"`
	PeerType                  int32                       `json:"peer_type"`
	PeerID                    int64                       `json:"peer_id"`
	ClientRandomID            int64                       `json:"client_random_id"`
	MessageText               string                      `json:"message_text,omitempty"`
	Entities                  []payload.MessageEntityV1   `json:"entities,omitempty"`
	ReplyToCanonicalMessageID int64                       `json:"reply_to_canonical_message_id,omitempty"`
	MediaRef                  *payload.MediaRefV1         `json:"media_ref,omitempty"`
	Attrs                     *payload.MessageAttrsV1     `json:"attrs,omitempty"`
	ForwardRef                *payload.ForwardRefV1       `json:"forward_ref,omitempty"`
	ServiceAction             *payload.ServiceActionRefV1 `json:"service_action,omitempty"`
	BatchSideEffectHash       string                      `json:"batch_side_effect_hash,omitempty"`
	ClearDraft                bool                        `json:"clear_draft,omitempty"`
	SourcePermAuthKeyID       int64                       `json:"source_perm_auth_key_id,omitempty"`
	ClearDraftBeforeDate      int32                       `json:"clear_draft_before_date,omitempty"`
}

func marshalSendRequestV3(normalized normalizedOutboxMessage, effects batchSideEffects) ([]byte, []byte, error) {
	request := sendRequestPayloadV3{
		SchemaVersion:             payload.MessageOperationSchemaVersionV3,
		SenderUserID:              normalized.FromUserID,
		PeerType:                  normalized.PeerType,
		PeerID:                    normalized.PeerID,
		ClientRandomID:            normalized.RandomID,
		MessageText:               normalized.MessageText,
		Entities:                  normalized.Entities,
		ReplyToCanonicalMessageID: normalized.ReplyToCanonicalMessageID,
		MediaRef:                  normalized.MediaRef,
		Attrs:                     normalized.attrsPtr(),
		ForwardRef:                normalized.ForwardRef,
		ServiceAction:             normalized.ServiceAction,
		ClearDraft:                effects.ClearDraft,
		SourcePermAuthKeyID:       effects.SourcePermAuthKeyID,
		ClearDraftBeforeDate:      effects.ClearDraftBeforeDate,
	}
	if effects.hasAny() {
		sideEffectsBody, err := json.Marshal(effects)
		if err != nil {
			return nil, nil, fmt.Errorf("%w: marshal send request side effects: %v", msg.ErrMsgStorage, err)
		}
		request.BatchSideEffectHash = hex.EncodeToString(payload.HashBytes(sideEffectsBody))
	}
	body, err := json.Marshal(request)
	if err != nil {
		return nil, nil, fmt.Errorf("%w: marshal send request v3: %v", msg.ErrMsgStorage, err)
	}
	return body, payload.HashBytes(body), nil
}

func buildNormalizedMessageOperationPayload(normalized normalizedOutboxMessage, toUserID int64, peerID int64, out bool, canonical *repository.CanonicalMessageResult, effects batchSideEffects) ([]byte, []byte, []byte, error) {
	date, err := msgDateInt32FromUnixSeconds(canonical.MessageDate, "message date")
	if err != nil {
		return nil, nil, nil, err
	}
	operation := payload.MessageOperationV3{
		SchemaVersion:             payload.MessageOperationSchemaVersionV3,
		OperationKind:             payload.OperationKindSendMessage,
		CanonicalMessageID:        canonical.CanonicalMessageID,
		PeerType:                  normalized.PeerType,
		PeerID:                    peerID,
		PeerSeq:                   canonical.PeerSeq,
		FromUserID:                normalized.FromUserID,
		ToUserID:                  toUserID,
		Date:                      date,
		Out:                       out,
		MessageText:               normalized.MessageText,
		Entities:                  normalized.Entities,
		ReplyToCanonicalMessageID: normalized.ReplyToCanonicalMessageID,
		ReplyToUserMessageID:      0,
		MediaRef:                  normalized.MediaRef,
		Attrs:                     normalized.attrsPtr(),
		ForwardRef:                normalized.ForwardRef,
		ServiceAction:             normalized.ServiceAction,
		ClearDraft:                effects.ClearDraft,
		SourcePermAuthKeyID:       effects.SourcePermAuthKeyID,
		ClearDraftBeforeDate:      effects.ClearDraftBeforeDate,
	}
	body, err := json.Marshal(operation)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("%w: marshal message operation from_user_id=%d peer_id=%d", msg.ErrMsgStorage, normalized.FromUserID, peerID)
	}
	hashBytes := payload.HashBytes(body)
	return body, hashBytes, hashBytes, nil
}

func normalizedCanonicalPayloads(normalized normalizedOutboxMessage) (repository.CreateCanonicalMessageInput, error) {
	var in repository.CreateCanonicalMessageInput
	if len(normalized.Entities) > 0 {
		body, err := json.Marshal(normalized.Entities)
		if err != nil {
			return in, fmt.Errorf("%w: marshal message entities: %v", msg.ErrMsgStorage, err)
		}
		in.EntitiesPayloadSchemaVersion = payload.MessageOperationSchemaVersionV3
		in.EntitiesPayload = body
	}
	if normalized.MediaRef != nil {
		body, err := json.Marshal(normalized.MediaRef)
		if err != nil {
			return in, fmt.Errorf("%w: marshal message media ref: %v", msg.ErrMsgStorage, err)
		}
		in.MediaRefSchemaVersion = payload.MediaRefSchemaVersionV1
		in.MediaRefPayload = body
	}
	if normalized.hasAttrs() {
		attrs := normalized.Attrs
		body, err := json.Marshal(attrs)
		if err != nil {
			return in, fmt.Errorf("%w: marshal message attrs: %v", msg.ErrMsgStorage, err)
		}
		in.MessageAttrsSchemaVersion = payload.MessageAttrsSchemaVersionV1
		in.MessageAttrsPayload = body
	}
	if normalized.ForwardRef != nil {
		body, err := json.Marshal(normalized.ForwardRef)
		if err != nil {
			return in, fmt.Errorf("%w: marshal message forward ref: %v", msg.ErrMsgStorage, err)
		}
		in.ForwardRefSchemaVersion = payload.ForwardRefSchemaVersionV1
		in.ForwardRefPayload = body
	}
	if normalized.ServiceAction != nil {
		body, err := json.Marshal(normalized.ServiceAction)
		if err != nil {
			return in, fmt.Errorf("%w: marshal message service action: %v", msg.ErrMsgStorage, err)
		}
		in.ServiceActionSchemaVersion = payload.ServiceActionSchemaVersionV1
		in.ServiceActionPayload = body
	}
	return in, nil
}

func (m normalizedOutboxMessage) attrsPtr() *payload.MessageAttrsV1 {
	if !m.hasAttrs() {
		return nil
	}
	attrs := m.Attrs
	return &attrs
}

func (m normalizedOutboxMessage) hasAttrs() bool {
	return m.Attrs.GroupedID != 0 || m.Attrs.Noforwards || m.Attrs.Silent || m.Attrs.InvertMedia
}

func (e batchSideEffects) hasAny() bool {
	return e.ClearDraft || e.SourcePermAuthKeyID != 0 || e.ClearDraftBeforeDate != 0
}
