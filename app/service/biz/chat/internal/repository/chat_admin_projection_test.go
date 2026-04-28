package repository

import (
	"testing"

	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestApplyChatAdminMutationPreservesUnchangedParticipantFields(t *testing.T) {
	original := tg.MakeTLImmutableChatParticipant(&tg.TLImmutableChatParticipant{
		Id:            41,
		ChatId:        10,
		UserId:        2,
		State:         chatpb.ChatMemberStateNormal,
		Useage:        7,
		InviterUserId: 1,
		InvitedAt:     100,
		KickedAt:      200,
		LeftAt:        300,
		Date:          400,
		IsBot:         true,
	}).ToImmutableChatParticipant()

	promoted, adminRights, link := applyChatAdminMutation(original, true)
	if promoted == original {
		t.Fatal("applyChatAdminMutation returned the original participant pointer")
	}
	if promoted.Useage != 7 || promoted.InviterUserId != 1 || promoted.InvitedAt != 100 || promoted.KickedAt != 200 || promoted.LeftAt != 300 || promoted.Date != 400 || !promoted.IsBot {
		t.Fatalf("promoted participant lost unchanged fields: %#v", promoted)
	}
	if promoted.ParticipantType != chatpb.ChatMemberAdmin || promoted.AdminRights == nil || promoted.Link == "" {
		t.Fatalf("promoted participant did not update admin fields: %#v", promoted)
	}
	if adminRights == 0 || link == "" {
		t.Fatalf("storage admin rights/link = %d/%q, want populated", adminRights, link)
	}

	demoted, adminRights, link := applyChatAdminMutation(promoted, false)
	if demoted.Useage != 7 || demoted.InviterUserId != 1 || demoted.InvitedAt != 100 || demoted.KickedAt != 200 || demoted.LeftAt != 300 || demoted.Date != 400 || !demoted.IsBot {
		t.Fatalf("demoted participant lost unchanged fields: %#v", demoted)
	}
	if demoted.ParticipantType != chatpb.ChatMemberNormal || demoted.AdminRights != nil || demoted.Link != "" {
		t.Fatalf("demoted participant did not clear admin fields: %#v", demoted)
	}
	if adminRights != 0 || link != "" {
		t.Fatalf("demote storage admin rights/link = %d/%q, want cleared", adminRights, link)
	}
}
