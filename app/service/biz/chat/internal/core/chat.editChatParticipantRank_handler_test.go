package core

import (
	"testing"

	"github.com/teamgram/proto/mtproto"
)

func TestCanEditChatParticipantRank(t *testing.T) {
	tests := []struct {
		name          string
		operator      *mtproto.ImmutableChatParticipant
		participantID int64
		want          bool
	}{
		{
			name: "creator can edit another participant rank",
			operator: &mtproto.ImmutableChatParticipant{
				UserId:          1,
				State:           mtproto.ChatMemberStateNormal,
				ParticipantType: mtproto.ChatMemberCreator,
			},
			participantID: 2,
			want:          true,
		},
		{
			name: "admin with manage ranks can edit own rank",
			operator: &mtproto.ImmutableChatParticipant{
				UserId:          2,
				State:           mtproto.ChatMemberStateNormal,
				ParticipantType: mtproto.ChatMemberAdmin,
				AdminRights: mtproto.MakeTLChatAdminRights(&mtproto.ChatAdminRights{
					ManageRanks: true,
				}).To_ChatAdminRights(),
			},
			participantID: 2,
			want:          true,
		},
		{
			name: "admin with manage ranks cannot edit another participant rank",
			operator: &mtproto.ImmutableChatParticipant{
				UserId:          2,
				State:           mtproto.ChatMemberStateNormal,
				ParticipantType: mtproto.ChatMemberAdmin,
				AdminRights: mtproto.MakeTLChatAdminRights(&mtproto.ChatAdminRights{
					ManageRanks: true,
				}).To_ChatAdminRights(),
			},
			participantID: 3,
			want:          false,
		},
		{
			name: "normal participant cannot edit own rank",
			operator: &mtproto.ImmutableChatParticipant{
				UserId:          2,
				State:           mtproto.ChatMemberStateNormal,
				ParticipantType: mtproto.ChatMemberNormal,
			},
			participantID: 2,
			want:          false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canEditChatParticipantRank(tt.operator, tt.participantID); got != tt.want {
				t.Fatalf("canEditChatParticipantRank() = %v, want %v", got, tt.want)
			}
		})
	}
}
