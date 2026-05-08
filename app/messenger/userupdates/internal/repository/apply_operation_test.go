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
	eventPayload, _, _, _, err := buildEventAndResponse(
		ApplyUserOperationInput{
			OperationID:      "op",
			AuthKeyIDExclude: &authKeyIDExclude,
		},
		payload.MessageOperationV1{
			SchemaVersion:      payload.MessageOperationSchemaVersion,
			OperationKind:      payload.OperationKindSendMessage,
			CanonicalMessageID: 7001,
			PeerType:           payload.PeerTypeUser,
			PeerID:             2002,
			PeerSeq:            9,
			FromUserID:         1001,
			ToUserID:           2002,
			Date:               1777781234,
			Out:                true,
			MessageText:        "hello",
		},
		38,
		1,
	)
	if err != nil {
		t.Fatalf("buildEventAndResponse() error = %v", err)
	}
	var event payload.MessageEventV1
	if err := json.Unmarshal(eventPayload, &event); err != nil {
		t.Fatalf("unmarshal event payload: %v", err)
	}
	if event.AuthKeyIdExclude == nil || *event.AuthKeyIdExclude != authKeyIDExclude {
		t.Fatalf("auth_key_id_exclude = %v, want %d", event.AuthKeyIdExclude, authKeyIDExclude)
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
