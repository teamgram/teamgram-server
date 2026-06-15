package core

import (
	"errors"
	"os"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestServiceActionPolicyCoversGeneratedActions(t *testing.T) {
	body, err := os.ReadFile("../../../../../pkg/proto/tg/schema.tl.tg_gen.go")
	if err != nil {
		t.Fatalf("read generated tg schema: %v", err)
	}
	re := regexp.MustCompile(`type TLMessageAction([A-Za-z0-9]+)`)
	matches := re.FindAllStringSubmatch(string(body), -1)
	if len(matches) == 0 {
		t.Fatal("no generated TLMessageAction types found")
	}
	seen := make(map[string]struct{}, len(matches))
	var missing []string
	for _, match := range matches {
		name := "messageAction" + match[1]
		if _, ok := seen[name]; ok {
			continue
		}
		seen[name] = struct{}{}
		if _, ok := serviceActionPolicyByClazzName[name]; !ok {
			missing = append(missing, name)
		}
	}
	sort.Strings(missing)
	if len(missing) > 0 {
		t.Fatalf("service action policy missing generated actions: %s", strings.Join(missing, ", "))
	}
}

func TestValidateServiceActionPolicyAllowsDisplayOnly(t *testing.T) {
	err := validateServiceActionPolicy(tg.MakeTLMessageActionCustomAction(&tg.TLMessageActionCustomAction{Message: "notice"}), payload.PeerTypeChat, 55, 1001, nil)
	if err != nil {
		t.Fatalf("validateServiceActionPolicy() error = %v", err)
	}
}

func TestValidateServiceActionPolicyAllowsChatEditTitleAfterOwnerMutation(t *testing.T) {
	err := validateServiceActionPolicy(tg.MakeTLMessageActionChatEditTitle(&tg.TLMessageActionChatEditTitle{Title: "new"}), payload.PeerTypeChat, 55, 1001, nil)
	if err != nil {
		t.Fatalf("validateServiceActionPolicy() error = %v", err)
	}
}

func TestValidateServiceActionPolicyAllowsChatEditPhotoAfterOwnerMutation(t *testing.T) {
	err := validateServiceActionPolicy(
		tg.MakeTLMessageActionChatEditPhoto(&tg.TLMessageActionChatEditPhoto{Photo: tg.MakeTLPhotoEmpty(&tg.TLPhotoEmpty{})}),
		payload.PeerTypeChat,
		55,
		1001,
		nil,
	)
	if err != nil {
		t.Fatalf("validateServiceActionPolicy() error = %v", err)
	}
}

func TestValidateServiceActionPolicyAllowsChatDeletePhotoAfterOwnerMutation(t *testing.T) {
	err := validateServiceActionPolicy(tg.MakeTLMessageActionChatDeletePhoto(&tg.TLMessageActionChatDeletePhoto{}), payload.PeerTypeChat, 55, 1001, nil)
	if err != nil {
		t.Fatalf("validateServiceActionPolicy() error = %v", err)
	}
}

func TestValidateServiceActionPolicyRejectsStateClaim(t *testing.T) {
	err := validateServiceActionPolicy(tg.MakeTLMessageActionPinMessage(&tg.TLMessageActionPinMessage{}), payload.PeerTypeChat, 55, 1001, nil)
	if err == nil {
		t.Fatal("validateServiceActionPolicy() error = nil, want reject")
	}
	if !errors.Is(err, msg.ErrSendStateConflict) {
		t.Fatalf("error = %v, want ErrSendStateConflict", err)
	}
}

func TestValidateServiceActionPolicyRejectsMissingAction(t *testing.T) {
	err := validateServiceActionPolicy(nil, payload.PeerTypeChat, 55, 1001, nil)
	if err == nil {
		t.Fatal("validateServiceActionPolicy() error = nil, want reject")
	}
	if !errors.Is(err, msg.ErrSendStateConflict) {
		t.Fatalf("error = %v, want ErrSendStateConflict", err)
	}
}

func TestValidateServiceActionPolicyRejectsUnclassifiedAction(t *testing.T) {
	err := validateServiceActionPolicy(fakeMessageAction{}, payload.PeerTypeChat, 55, 1001, nil)
	if err == nil {
		t.Fatal("validateServiceActionPolicy() error = nil, want reject")
	}
	if !errors.Is(err, msg.ErrSendStateConflict) {
		t.Fatalf("error = %v, want ErrSendStateConflict", err)
	}
}

type fakeMessageAction struct{}

func (fakeMessageAction) Encode(*bin.Encoder, int32) error { return nil }
func (fakeMessageAction) Decode(*bin.Decoder) error        { return nil }
func (fakeMessageAction) MessageActionClazzName() string   { return "messageActionFuture" }

func TestValidateServiceActionPolicyRequiresParticipantsFact(t *testing.T) {
	err := validateServiceActionPolicy(tg.MakeTLMessageActionChatCreate(&tg.TLMessageActionChatCreate{Title: "team"}), payload.PeerTypeChat, 55, 1001, nil)
	if err == nil {
		t.Fatal("validateServiceActionPolicy() error = nil, want missing fact")
	}
	if !errors.Is(err, msg.ErrSendStateConflict) {
		t.Fatalf("error = %v, want ErrSendStateConflict", err)
	}
}

func TestValidateServiceActionPolicyAcceptsParticipantsFact(t *testing.T) {
	fact, err := payload.WrapFact(payload.FactKindChatParticipantsChanged, payload.ChatParticipantsChangedFactV1{
		SchemaVersion: 1,
		ChatID:        55,
		ActorUserID:   1001,
		Version:       1,
		Participants:  []payload.ChatParticipantFactV1{{UserID: 1001, Role: "creator", Date: 1_772_000_000}},
	})
	if err != nil {
		t.Fatalf("WrapFact() error = %v", err)
	}
	err = validateServiceActionPolicy(tg.MakeTLMessageActionChatCreate(&tg.TLMessageActionChatCreate{Title: "team", Users: []int64{1001}}), payload.PeerTypeChat, 55, 1001, []payload.UpdateFactV1{fact})
	if err != nil {
		t.Fatalf("validateServiceActionPolicy() error = %v", err)
	}
}

func TestValidateServiceActionPolicyRejectsUnrelatedParticipantsFact(t *testing.T) {
	fact, err := payload.WrapFact(payload.FactKindChatParticipantsChanged, payload.ChatParticipantsChangedFactV1{
		SchemaVersion: 1,
		ChatID:        56,
		ActorUserID:   1001,
		Version:       1,
		Participants:  []payload.ChatParticipantFactV1{{UserID: 1001, Role: "creator", Date: 1_772_000_000}},
	})
	if err != nil {
		t.Fatalf("WrapFact() error = %v", err)
	}
	err = validateServiceActionPolicy(tg.MakeTLMessageActionChatAddUser(&tg.TLMessageActionChatAddUser{Users: []int64{1002}}), payload.PeerTypeChat, 55, 1001, []payload.UpdateFactV1{fact})
	if err == nil || !errors.Is(err, msg.ErrSendStateConflict) {
		t.Fatalf("validateServiceActionPolicy() error = %v, want ErrSendStateConflict", err)
	}
}

func TestValidateServiceActionPolicyRejectsAddUserMissingFromParticipantsFact(t *testing.T) {
	fact, err := payload.WrapFact(payload.FactKindChatParticipantsChanged, payload.ChatParticipantsChangedFactV1{
		SchemaVersion: 1,
		ChatID:        55,
		ActorUserID:   1001,
		Version:       1,
		Participants:  []payload.ChatParticipantFactV1{{UserID: 1001, Role: "creator", Date: 1_772_000_000}},
	})
	if err != nil {
		t.Fatalf("WrapFact() error = %v", err)
	}
	err = validateServiceActionPolicy(tg.MakeTLMessageActionChatAddUser(&tg.TLMessageActionChatAddUser{Users: []int64{1002}}), payload.PeerTypeChat, 55, 1001, []payload.UpdateFactV1{fact})
	if err == nil || !errors.Is(err, msg.ErrSendStateConflict) {
		t.Fatalf("validateServiceActionPolicy() error = %v, want ErrSendStateConflict", err)
	}
}
