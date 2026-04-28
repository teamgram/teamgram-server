package repository

import "testing"

func TestIsValidUsername(t *testing.T) {
	valid := []string{"teamgram", "team_gram", "abc12", "a1_b2"}
	for _, username := range valid {
		if !isValidUsername(username) {
			t.Fatalf("expected valid username %q", username)
		}
	}

	invalid := []string{"", "ab", "12345", "_abcde", "abcde_", "a__bcde", "UPPER", "with-dash", "with.dot"}
	for _, username := range invalid {
		if isValidUsername(username) {
			t.Fatalf("expected invalid username %q", username)
		}
	}
}
