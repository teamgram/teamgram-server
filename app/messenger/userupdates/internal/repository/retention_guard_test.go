package repository

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func TestOperationalCleanupDoesNotDeleteDurableTables(t *testing.T) {
	durableTables := []string{
		"canonical_messages",
		"user_message_views",
		"user_dialogs",
		"user_pts_events",
		"user_operation_results",
	}

	sources := repositorySourceFiles(t)
	for _, table := range durableTables {
		t.Run(table, func(t *testing.T) {
			deleteFromTable := regexp.MustCompile(`(?i)\bdelete\s+from\s+` + "`?" + `\s*` + regexp.QuoteMeta(table) + `\s*` + "`?" + `\b`)
			for _, source := range sources {
				if deleteFromTable.MatchString(source.contents) {
					t.Fatalf("%s deletes from durable table %s", source.path, table)
				}
			}
		})
	}
}

func TestOperationalCleanupTablesAreNotDurable(t *testing.T) {
	durableTables := map[string]struct{}{
		"canonical_messages":     {},
		"user_message_views":     {},
		"user_dialogs":           {},
		"user_pts_events":        {},
		"user_operation_results": {},
	}
	operationalTables := []string{
		"affected_operation_outbox",
		"push_task_outbox",
		"dialog_side_effect_outbox",
	}

	for _, table := range operationalTables {
		t.Run(table, func(t *testing.T) {
			if _, ok := durableTables[table]; ok {
				t.Fatalf("operational cleanup table %s is misclassified as durable", table)
			}
		})
	}
}

type repositorySourceFile struct {
	path     string
	contents string
}

func repositorySourceFiles(t *testing.T) []repositorySourceFile {
	t.Helper()

	var sources []repositorySourceFile
	err := filepath.WalkDir(".", func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, "_test.go") {
			return nil
		}
		if !strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, ".xml") {
			return nil
		}

		data, err := os.ReadFile(filepath.Clean(path))
		if err != nil {
			return err
		}
		sources = append(sources, repositorySourceFile{
			path:     path,
			contents: string(data),
		})
		return nil
	})
	if err != nil {
		t.Fatalf("scan repository sources: %v", err)
	}
	return sources
}
