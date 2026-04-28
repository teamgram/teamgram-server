package repository

import (
	"encoding/json"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestEncodePrivacyRulesUsesStorageDTO(t *testing.T) {
	data, err := encodePrivacyRules([]tg.PrivacyRuleClazz{
		tg.MakeTLPrivacyValueAllowUsers(&tg.TLPrivacyValueAllowUsers{Users: []int64{1, 2}}).ToPrivacyRule().Clazz,
		tg.PrivacyValueDisallowBotsClazz,
	})
	if err != nil {
		t.Fatal(err)
	}

	var payload privacyRulesStorageDTO
	if err := json.Unmarshal([]byte(data), &payload); err != nil {
		t.Fatal(err)
	}
	if payload.SchemaVersion != 1 {
		t.Fatalf("schema_version = %d, want 1", payload.SchemaVersion)
	}
	if len(payload.Rules) != 2 {
		t.Fatalf("len(rules) = %d, want 2", len(payload.Rules))
	}
	if payload.Rules[0].Type != "allow_users" || len(payload.Rules[0].Users) != 2 {
		t.Fatalf("unexpected first rule: %+v", payload.Rules[0])
	}
	if payload.Rules[0].LegacyName != "" || payload.Rules[0].LegacyID != 0 {
		t.Fatalf("new storage DTO leaked TL identity: %+v", payload.Rules[0])
	}
	if payload.Rules[1].Type != "disallow_bots" {
		t.Fatalf("unexpected second rule: %+v", payload.Rules[1])
	}
}

func TestDecodePrivacyRulesAcceptsLegacyTLJSONShape(t *testing.T) {
	rules, err := decodePrivacyRules(`[{"_name":"privacyValueDisallowUsers","users":[42]},{"_name":"privacyValueAllowAll"}]`)
	if err != nil {
		t.Fatal(err)
	}
	if len(rules) != 2 {
		t.Fatalf("len(rules) = %d, want 2", len(rules))
	}
	disallowUsers, ok := rules[0].(*tg.TLPrivacyValueDisallowUsers)
	if !ok {
		t.Fatalf("rule[0] = %T, want *TLPrivacyValueDisallowUsers", rules[0])
	}
	if len(disallowUsers.Users) != 1 || disallowUsers.Users[0] != 42 {
		t.Fatalf("unexpected disallow users rule: %+v", disallowUsers)
	}
	if rules[1].PrivacyRuleClazzName() != tg.ClazzName_privacyValueAllowAll {
		t.Fatalf("rule[1] = %s, want %s", rules[1].PrivacyRuleClazzName(), tg.ClazzName_privacyValueAllowAll)
	}
}

func TestEvaluatePrivacyRulesWithContext(t *testing.T) {
	tests := []struct {
		name  string
		ctx   privacyEvaluationContext
		rules []tg.PrivacyRuleClazz
		want  bool
	}{
		{
			name:  "allow all permits peer",
			ctx:   privacyEvaluationContext{PeerID: 42},
			rules: []tg.PrivacyRuleClazz{tg.PrivacyValueAllowAllClazz},
			want:  true,
		},
		{
			name:  "disallow all rejects peer",
			ctx:   privacyEvaluationContext{PeerID: 42},
			rules: []tg.PrivacyRuleClazz{tg.PrivacyValueDisallowAllClazz},
			want:  false,
		},
		{
			name:  "allow contacts permits contact",
			ctx:   privacyEvaluationContext{PeerID: 42, IsContact: true},
			rules: []tg.PrivacyRuleClazz{tg.PrivacyValueAllowContactsClazz},
			want:  true,
		},
		{
			name:  "allow contacts rejects non contact",
			ctx:   privacyEvaluationContext{PeerID: 42},
			rules: []tg.PrivacyRuleClazz{tg.PrivacyValueAllowContactsClazz},
			want:  false,
		},
		{
			name:  "disallow contacts rejects contact under allow all",
			ctx:   privacyEvaluationContext{PeerID: 42, IsContact: true},
			rules: []tg.PrivacyRuleClazz{tg.PrivacyValueAllowAllClazz, tg.PrivacyValueDisallowContactsClazz},
			want:  false,
		},
		{
			name: "allow users permits listed peer",
			ctx:  privacyEvaluationContext{PeerID: 42},
			rules: []tg.PrivacyRuleClazz{
				tg.PrivacyValueDisallowAllClazz,
				tg.MakeTLPrivacyValueAllowUsers(&tg.TLPrivacyValueAllowUsers{Users: []int64{42}}).ToPrivacyRule().Clazz,
			},
			want: true,
		},
		{
			name: "disallow users rejects listed peer",
			ctx:  privacyEvaluationContext{PeerID: 42},
			rules: []tg.PrivacyRuleClazz{
				tg.PrivacyValueAllowAllClazz,
				tg.MakeTLPrivacyValueDisallowUsers(&tg.TLPrivacyValueDisallowUsers{Users: []int64{42}}).ToPrivacyRule().Clazz,
			},
			want: false,
		},
		{
			name: "unlisted peer keeps default decision",
			ctx:  privacyEvaluationContext{PeerID: 43},
			rules: []tg.PrivacyRuleClazz{
				tg.PrivacyValueAllowAllClazz,
				tg.MakeTLPrivacyValueDisallowUsers(&tg.TLPrivacyValueDisallowUsers{Users: []int64{42}}).ToPrivacyRule().Clazz,
			},
			want: true,
		},
		{
			name:  "close friends permits close friend",
			ctx:   privacyEvaluationContext{PeerID: 42, IsCloseFriend: true},
			rules: []tg.PrivacyRuleClazz{tg.PrivacyValueDisallowAllClazz, tg.PrivacyValueAllowCloseFriendsClazz},
			want:  true,
		},
		{
			name:  "premium permits premium peer",
			ctx:   privacyEvaluationContext{PeerID: 42, IsPremium: true},
			rules: []tg.PrivacyRuleClazz{tg.PrivacyValueDisallowAllClazz, tg.PrivacyValueAllowPremiumClazz},
			want:  true,
		},
		{
			name:  "disallow bots rejects bot under allow all",
			ctx:   privacyEvaluationContext{PeerID: 42, IsBot: true},
			rules: []tg.PrivacyRuleClazz{tg.PrivacyValueAllowAllClazz, tg.PrivacyValueDisallowBotsClazz},
			want:  false,
		},
		{
			name: "unsupported chat participant rule is conservative",
			ctx:  privacyEvaluationContext{PeerID: 42},
			rules: []tg.PrivacyRuleClazz{
				tg.PrivacyValueDisallowAllClazz,
				tg.MakeTLPrivacyValueAllowChatParticipants(&tg.TLPrivacyValueAllowChatParticipants{Chats: []int64{10}}).ToPrivacyRule().Clazz,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evaluatePrivacyRules(tt.rules, tt.ctx); got != tt.want {
				t.Fatalf("evaluatePrivacyRules() = %v, want %v", got, tt.want)
			}
		})
	}
}
