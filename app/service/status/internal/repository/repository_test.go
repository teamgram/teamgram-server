package repository

import (
	"encoding/json"
	"strings"
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

func TestSessionEntryCachePayloadDoesNotUseTLJSON(t *testing.T) {
	sess := &status.TLSessionEntry{
		UserId:        100,
		AuthKeyId:     12345,
		Gateway:       "gw1",
		Expired:       1700000000,
		Layer:         177,
		PermAuthKeyId: 12345,
		Client:        "Android",
	}

	cacheData := sessionEntryCacheDataFromTL(sess)
	data, err := json.Marshal(cacheData)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	payload := string(data)
	if strings.Contains(payload, `"_name"`) || strings.Contains(payload, `"_id"`) {
		t.Fatalf("cache payload contains TL JSON metadata: %s", payload)
	}

	var decoded sessionEntryCacheData
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	got := decoded.toTL()
	if got.AuthKeyId != sess.AuthKeyId || got.UserId != sess.UserId || got.Gateway != sess.Gateway {
		t.Errorf("roundtrip mismatch: got %+v, want %+v", got, *sess)
	}
}
