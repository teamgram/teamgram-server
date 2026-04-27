package repository

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/authsession/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestAuthDataStateMapping(t *testing.T) {
	tests := []struct {
		name string
		data *cacheAuthData
		want int32
	}{
		{name: "nil aggregate", data: nil, want: tg.AuthStateNew},
		{name: "nil client", data: &cacheAuthData{}, want: tg.AuthStateWaitInit},
		{name: "client without user", data: &cacheAuthData{Client: &clientSessionCacheData{}}, want: tg.AuthStateUnauthorized},
		{name: "client and user", data: &cacheAuthData{Client: &clientSessionCacheData{}, BindUser: &bindUser{UserId: 42}}, want: tg.AuthStateNormal},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toAuthKeyStateData(1001, tt.data)
			if got.KeyState != tt.want {
				t.Fatalf("KeyState = %d, want %d", got.KeyState, tt.want)
			}
		})
	}
}

func TestAuthDataCachePayloadDoesNotUseTLDebugJSON(t *testing.T) {
	data := &cacheAuthData{
		Client: &clientSessionCacheData{
			AuthKeyId: 1001,
			Ip:        "127.0.0.1",
			Layer:     158,
			Params:    "{}",
		},
	}
	b, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Marshal cacheAuthData error: %v", err)
	}
	payload := string(b)
	if strings.Contains(payload, `"_name"`) || strings.Contains(payload, `"_id"`) {
		t.Fatalf("cache payload contains TL debug JSON metadata: %s", payload)
	}
	if !strings.Contains(payload, `"auth_key_id":1001`) || !strings.Contains(payload, `"params":"{}"`) {
		t.Fatalf("cache payload missing service-owned client fields: %s", payload)
	}
}

func TestAuthDataClientSessionMapping(t *testing.T) {
	got := toClientSession(1001, &model.Auths{
		AuthKeyId:      2002,
		ClientIp:       "127.0.0.1",
		Layer:          158,
		ApiId:          9,
		DeviceModel:    "device",
		SystemVersion:  "system",
		AppVersion:     "app",
		SystemLangCode: "en-US",
		LangPack:       "android",
		LangCode:       "en",
		Proxy:          "proxy",
		Params:         "{}",
	})
	if got.AuthKeyId != 1001 || got.Ip != "127.0.0.1" || got.Layer != 158 || got.Params != "{}" {
		t.Fatalf("mapped client session mismatch: %#v", got)
	}
}

func TestClientKindAndLangPackMapping(t *testing.T) {
	if got := normalizeLangPack("", "Telegram A"); got != "weba" {
		t.Fatalf("normalizeLangPack() = %q, want weba", got)
	}
	if got := normalizeLangPack("android", "Telegram TDLib"); got != "android" {
		t.Fatalf("normalizeLangPack() = %q, want android", got)
	}
}
