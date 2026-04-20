package main_test

import (
	"os"
	"os/exec"
	"testing"
)

func TestIDGenCommandBuilds(t *testing.T) {
	t.Parallel()

	cmd := exec.Command("go", "build", ".")
	cmd.Dir = "."
	cmd.Env = append(os.Environ(),
		"GOCACHE="+t.TempDir(),
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("go build failed: %v\n%s", err, output)
	}
}
