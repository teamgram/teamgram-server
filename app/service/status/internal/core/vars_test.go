package core

import "testing"

func TestGetUserKey(t *testing.T) {
	tests := []struct {
		name string
		id   int64
		want string
	}{
		{"positive id", 12345, "user_online_keys#12345"},
		{"zero id", 0, "user_online_keys#0"},
		{"negative id", -1, "user_online_keys#-1"},
		{"large id", 1000000000, "user_online_keys#1000000000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getUserKey(tt.id)
			if got != tt.want {
				t.Errorf("getUserKey(%d) = %q, want %q", tt.id, got, tt.want)
			}
		})
	}
}
