package core

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestReadHistoryDoesNotUseLogOnlyPeerOperation(t *testing.T) {
	forbidden := []string{
		"processPeer" + "ReadOutboxOperation",
		"peer read outbox update " + "failed",
	}

	entries, err := os.ReadDir(".")
	if err != nil {
		t.Fatalf("read core dir: %v", err)
	}
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") {
			continue
		}

		data, err := os.ReadFile(filepath.Clean(name))
		if err != nil {
			t.Fatalf("read %s: %v", name, err)
		}
		for _, marker := range forbidden {
			if strings.Contains(string(data), marker) {
				t.Fatalf("%s contains forbidden readHistory log-only peer path marker %q", name, marker)
			}
		}
	}
}
