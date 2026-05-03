package model

import (
	"os"
	"strings"
	"testing"
)

func TestAuthsInsertOrUpdatePersistsLayer(t *testing.T) {
	body, err := os.ReadFile("auths_model_biz_gen.go")
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}
	source := string(body)
	if !strings.Contains(source, "insert into auths(auth_key_id, layer, api_id, device_model") {
		t.Fatal("auths.InsertOrUpdate insert columns do not include layer")
	}
	if !strings.Contains(source, "layer = values(layer), api_id = values(api_id)") {
		t.Fatal("auths.InsertOrUpdate update clause does not include layer")
	}
}
