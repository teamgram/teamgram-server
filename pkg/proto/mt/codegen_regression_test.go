package mt

import (
	"errors"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
)

func TestTLReqPqDecodePropagatesClazzIDReadError(t *testing.T) {
	var m TLReqPq
	err := m.Decode(bin.NewDecoder(nil))
	if !errors.Is(err, bin.ErrUnexpectedEOF) {
		t.Fatalf("expected unexpected EOF, got %v", err)
	}
}

func TestGeneratedMTServiceUsesSwitchAndNoIgnoredClazzIDError(t *testing.T) {
	data, err := os.ReadFile("schema.tl.mt_service.pb.go")
	if err != nil {
		t.Fatalf("read generated file: %v", err)
	}

	src := string(data)
	for _, bad := range []string{
		"var encodeF = map[uint32]func() error",
		"var decodeF = map[uint32]func() error",
		"m.ClazzID, _ = d.ClazzID()",
	} {
		if strings.Contains(src, bad) {
			t.Fatalf("generated mt service code still contains banned pattern: %q", bad)
		}
	}
}

func TestGeneratedMTCodeDoesNotOverwriteDecodeErrors(t *testing.T) {
	files := []string{
		"schema.tl.mt.pb.go",
		"schema.tl.mt_service.pb.go",
	}
	re := regexp.MustCompile(`err = [^\n]+\n\s*err = `)

	for _, name := range files {
		data, err := os.ReadFile(name)
		if err != nil {
			t.Fatalf("read %s: %v", name, err)
		}
		if re.Find(data) != nil {
			t.Fatalf("generated file %s still contains consecutive err overwrites", name)
		}
	}
}

func TestValidateRecursesIntoRequiredObjectSlices(t *testing.T) {
	msg := &TLTlsClientHello{
		Blocks: []TlsBlockClazz{
			&TLTlsBlockScope{},
		},
	}

	if err := msg.Validate(223); err == nil {
		t.Fatalf("expected nested validation error, got nil")
	}
}
