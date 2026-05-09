package objectstore

import (
	"errors"
	"testing"
)

func TestManifestKeys(t *testing.T) {
	keys := ManifestKeys{MetaPrefix: "_meta"}
	got, err := keys.Object("obj-1")
	if err != nil {
		t.Fatalf("Object() error = %v", err)
	}
	if got != "_meta/objects/obj-1.json" {
		t.Fatalf("Object() = %q", got)
	}

	got, err = keys.Upload("ext:1001:2002:1")
	if err != nil {
		t.Fatalf("Upload() error = %v", err)
	}
	if got != "_meta/uploads/ext:1001:2002:1.json" {
		t.Fatalf("Upload() = %q", got)
	}

	got, err = keys.Hashes("obj-1")
	if err != nil {
		t.Fatalf("Hashes() error = %v", err)
	}
	if got != "_meta/hashes/obj-1/v1.json" {
		t.Fatalf("Hashes() = %q", got)
	}
}

func TestManifestKeysDefaultAndSlashedPrefix(t *testing.T) {
	got, err := (ManifestKeys{}).Object("obj-1")
	if err != nil {
		t.Fatalf("Object() error = %v", err)
	}
	if got != "_meta/objects/obj-1.json" {
		t.Fatalf("Object() with default prefix = %q", got)
	}

	got, err = (ManifestKeys{MetaPrefix: "/custom/meta/"}).Hashes("obj-1")
	if err != nil {
		t.Fatalf("Hashes() error = %v", err)
	}
	if got != "custom/meta/hashes/obj-1/v1.json" {
		t.Fatalf("Hashes() with slashed prefix = %q", got)
	}
}

func TestManifestKeysRejectInvalidIDs(t *testing.T) {
	keys := ManifestKeys{MetaPrefix: "_meta"}
	for _, tc := range []struct {
		name string
		id   string
	}{
		{name: "empty", id: ""},
		{name: "slash", id: "obj/1"},
		{name: "backslash", id: `obj\1`},
		{name: "dotdot", id: "obj..1"},
		{name: "whitespace", id: " obj-1"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := keys.Object(tc.id); !errors.Is(err, ErrInvalidManifestKey) {
				t.Fatalf("Object() error = %v, want ErrInvalidManifestKey", err)
			}
			if _, err := keys.Upload(tc.id); !errors.Is(err, ErrInvalidManifestKey) {
				t.Fatalf("Upload() error = %v, want ErrInvalidManifestKey", err)
			}
			if _, err := keys.Hashes(tc.id); !errors.Is(err, ErrInvalidManifestKey) {
				t.Fatalf("Hashes() error = %v, want ErrInvalidManifestKey", err)
			}
		})
	}
}
