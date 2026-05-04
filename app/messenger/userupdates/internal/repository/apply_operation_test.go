package repository

import (
	"encoding/json"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
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
