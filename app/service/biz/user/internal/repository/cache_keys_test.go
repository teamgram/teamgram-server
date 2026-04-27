package repository

import "testing"

func TestUserDataCacheKeyUsesV3Prefix(t *testing.T) {
	got := userDataCacheKey(123)
	if got != "user_data.v3#123" {
		t.Fatalf("unexpected key %q", got)
	}
}
