package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestLoadConfigFallsBackWhenFileMissing(t *testing.T) {
	cfg := loadConfig("/tmp/teamgram-server-v2-missing-config.data")
	if cfg == nil {
		t.Fatal("expected fallback config, got nil")
	}
	if cfg.TestMode != tg.BoolFalseClazz {
		t.Fatalf("expected test_mode=false fallback, got %#v", cfg.TestMode)
	}
	if cfg.MeUrlPrefix != "https://t.me/" {
		t.Fatalf("expected fallback me_url_prefix, got %q", cfg.MeUrlPrefix)
	}
	if cfg.DcTxtDomainName == "" {
		t.Fatal("expected non-empty dc_txt_domain_name fallback")
	}
}

func TestHelpGetConfigUsesRuntimeTimestamps(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.HelpGetConfig(&tg.TLHelpGetConfig{})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected config result, got nil")
	}
	if result.Expires <= result.Date {
		t.Fatalf("expected expires > date, got %d <= %d", result.Expires, result.Date)
	}
	if result.MeUrlPrefix != "https://t.me/" {
		t.Fatalf("expected fallback me_url_prefix, got %q", result.MeUrlPrefix)
	}
}
