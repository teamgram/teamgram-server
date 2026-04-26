package repository

import (
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
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
		{name: "client without user", data: &cacheAuthData{Client: &authsession.ClientSession{}}, want: tg.AuthStateUnauthorized},
		{name: "client and user", data: &cacheAuthData{Client: &authsession.ClientSession{}, BindUser: &bindUser{UserId: 42}}, want: tg.AuthStateNormal},
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
