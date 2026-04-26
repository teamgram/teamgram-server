package repository

import (
	"encoding/json"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/status/status"
)

func TestGetUserKey(t *testing.T) {
	key := getUserKey(123)
	expected := "user_online_keys#123"
	if key != expected {
		t.Errorf("getUserKey(123) = %s, want %s", key, expected)
	}
}

func TestGetUserKey_Negative(t *testing.T) {
	key := getUserKey(-1)
	expected := "user_online_keys#-1"
	if key != expected {
		t.Errorf("getUserKey(-1) = %s, want %s", key, expected)
	}
}

func TestSessionEntryJSONRoundtrip(t *testing.T) {
	sess := &status.TLSessionEntry{
		UserId:        100,
		AuthKeyId:     12345,
		Gateway:       "gw1",
		Expired:       1700000000,
		Layer:         177,
		PermAuthKeyId: 12345,
		Client:        "Android",
	}

	// Match the repository write path: plain JSON without MarshalWithName wrapper.
	type plainSessionEntry status.TLSessionEntry
	data, err := json.Marshal((*plainSessionEntry)(sess))
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var decoded status.TLSessionEntry
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if decoded.AuthKeyId != sess.AuthKeyId || decoded.UserId != sess.UserId {
		t.Errorf("roundtrip mismatch: got %+v, want %+v", decoded, *sess)
	}
}
