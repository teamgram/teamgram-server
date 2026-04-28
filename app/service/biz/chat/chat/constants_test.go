package chat

import "testing"

func TestChatMemberConstantsMatchLegacyStorageValues(t *testing.T) {
	tests := []struct {
		name string
		got  int32
		want int32
	}{
		{"ChatMemberNormal", ChatMemberNormal, 0},
		{"ChatMemberAdmin", ChatMemberAdmin, 1},
		{"ChatMemberCreator", ChatMemberCreator, 2},
		{"ChatMemberStateNormal", ChatMemberStateNormal, 0},
		{"ChatMemberStateLeft", ChatMemberStateLeft, 1},
		{"ChatMemberStateKicked", ChatMemberStateKicked, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Fatalf("%s = %d, want %d", tt.name, tt.got, tt.want)
			}
		})
	}
}
