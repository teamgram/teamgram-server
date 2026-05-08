package repository

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
)

func TestBuildEventAndResponseCarriesAuthKeyExclude(t *testing.T) {
	authKeyIDExclude := int64(9001)
	eventPayload, _, responsePayload, _, err := buildEventAndResponse(
		ApplyUserOperationInput{
			OperationID:      "op",
			AuthKeyIDExclude: &authKeyIDExclude,
		},
		payload.MessageOperationV1{
			SchemaVersion:        payload.MessageOperationSchemaVersion,
			OperationKind:        payload.OperationKindSendMessage,
			CanonicalMessageID:   7001,
			PeerType:             payload.PeerTypeUser,
			PeerID:               2002,
			PeerSeq:              9,
			UserMessageID:        101,
			FromUserID:           1001,
			ToUserID:             2002,
			Date:                 1777781234,
			Out:                  true,
			MessageText:          "hello",
			ReplyToUserMessageID: 77,
		},
		38,
		1,
	)
	if err != nil {
		t.Fatalf("buildEventAndResponse() error = %v", err)
	}
	var event payload.MessageEventV2
	if err := json.Unmarshal(eventPayload, &event); err != nil {
		t.Fatalf("unmarshal event payload: %v", err)
	}
	if event.AuthKeyIdExclude == nil || *event.AuthKeyIdExclude != authKeyIDExclude {
		t.Fatalf("auth_key_id_exclude = %v, want %d", event.AuthKeyIdExclude, authKeyIDExclude)
	}
	if event.MessageID != 101 || event.PeerSeq != 9 || event.ReplyToUserMessageID != 77 {
		t.Fatalf("event public/internal ids = %+v, want message_id=101 peer_seq=9 reply=77", event)
	}
	var response payload.OperationResponseV2
	if err := json.Unmarshal(responsePayload, &response); err != nil {
		t.Fatalf("unmarshal response payload: %v", err)
	}
	if response.UserMessageID != 101 {
		t.Fatalf("response user_message_id = %d, want 101", response.UserMessageID)
	}
}

func TestBuildEditEventAndResponseCarriesPublicMessageID(t *testing.T) {
	eventPayload, _, responsePayload, _, err := buildEventAndResponse(
		ApplyUserOperationInput{OperationID: "edit-op"},
		payload.MessageOperationV1{
			SchemaVersion:      payload.MessageOperationSchemaVersion,
			OperationKind:      payload.OperationKindEditMessage,
			CanonicalMessageID: 7001,
			PeerType:           payload.PeerTypeUser,
			PeerID:             2002,
			PeerSeq:            9,
			UserMessageID:      101,
			FromUserID:         1001,
			ToUserID:           2002,
			Date:               1777781234,
			EditDate:           1777782234,
			EditVersion:        2,
			Out:                true,
			MessageText:        "edited",
		},
		39,
		1,
	)
	if err != nil {
		t.Fatalf("buildEventAndResponse() error = %v", err)
	}
	var event payload.MessageEventV2
	if err := json.Unmarshal(eventPayload, &event); err != nil {
		t.Fatalf("unmarshal event payload: %v", err)
	}
	if event.EventKind != payload.OperationKindEditMessage || event.MessageID != 101 || event.PeerSeq != 9 {
		t.Fatalf("edit event ids = %+v, want public message_id=101 and peer_seq=9", event)
	}
	var response payload.OperationResponseV2
	if err := json.Unmarshal(responsePayload, &response); err != nil {
		t.Fatalf("unmarshal response payload: %v", err)
	}
	if response.UserMessageID != 101 {
		t.Fatalf("edit response user_message_id = %d, want 101", response.UserMessageID)
	}
}

func TestExtractHashTagsNormalizesAndDeduplicates(t *testing.T) {
	got := extractHashTags("hello #Go #go #team_gram #中文 ok #")
	want := []string{"go", "team_gram", "中文"}
	if len(got) != len(want) {
		t.Fatalf("extractHashTags len = %d, want %d: %#v", len(got), len(want), got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("extractHashTags[%d] = %q, want %q; all=%#v", i, got[i], want[i], got)
		}
	}
}

func TestValidateAffectedOutboxRejectsPayloadHashMismatch(t *testing.T) {
	err := validateAffectedOutbox(AffectedOutbox{
		TargetUserID:   1001,
		OperationID:    "affected-op",
		OperationKind:  payload.OperationKindSendMessage,
		DeliveryPolicy: DeliveryPolicyDurableAsync,
		PayloadCodec:   PayloadCodecJSON,
		Payload:        []byte(`{"ok":true}`),
		PayloadHash:    []byte("wrong"),
	})
	if !errors.Is(err, userupdates.ErrOperationPayloadConflict) {
		t.Fatalf("validateAffectedOutbox() error = %v, want ErrOperationPayloadConflict", err)
	}
}
