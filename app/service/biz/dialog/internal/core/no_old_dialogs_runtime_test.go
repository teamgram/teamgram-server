package core

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func TestNoOldDialogsRuntime(t *testing.T) {
	root := filepath.Clean("../../")
	pattern := regexp.MustCompile(`\bDialogsModel\b`)
	var hits []string
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			name := d.Name()
			if name == "dialog" || name == "dialogservice" || path == filepath.Join(root, "internal", "repository", "model") {
				return filepath.SkipDir
			}
			return nil
		}
		if strings.HasSuffix(path, "_test.go") || (filepath.Ext(path) != ".go" && filepath.Ext(path) != ".xml") {
			return nil
		}
		rel, err := filepath.Rel(".", path)
		if err != nil {
			return err
		}
		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if pattern.Match(b) {
			hits = append(hits, rel)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("scan old dialogs runtime: %v", err)
	}
	if len(hits) != 0 {
		t.Fatalf("old mixed dialogs runtime references remain: %v", hits)
	}
}
