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
