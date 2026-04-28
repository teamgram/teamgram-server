package repository

import (
	"testing"

	"github.com/zeromicro/go-zero/core/jsonx"
)

func TestAvailableReactionsToStorageUsesJSONX(t *testing.T) {
	got := availableReactionsToStorage([]string{"like", "fire"})
	var decoded []string
	if err := jsonx.UnmarshalFromString(got, &decoded); err != nil {
		t.Fatalf("availableReactionsToStorage produced invalid jsonx payload %q: %v", got, err)
	}
	if len(decoded) != 2 || decoded[0] != "like" || decoded[1] != "fire" {
		t.Fatalf("decoded reactions = %#v", decoded)
	}
}

func TestChatAttributeMutationsNeedingExplicitVersionBump(t *testing.T) {
	for _, op := range []chatAttributeMutation{
		chatAttributeMutationAbout,
		chatAttributeMutationNoForwards,
		chatAttributeMutationTTLPeriod,
		chatAttributeMutationAvailableReactions,
		chatAttributeMutationAdmin,
	} {
		if !op.needsExplicitVersionBump() {
			t.Fatalf("%s should call UpdateVersionTx explicitly", op)
		}
	}
	for _, op := range []chatAttributeMutation{
		chatAttributeMutationTitle,
		chatAttributeMutationPhoto,
		chatAttributeMutationDefaultBannedRights,
	} {
		if op.needsExplicitVersionBump() {
			t.Fatalf("%s should rely on generated version bump", op)
		}
	}
}
