package tg

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
)

const testLayer = 223

func TestTLAuthSendCodeCalcSizeMatchesEncodedLen(t *testing.T) {
	obj := &TLAuthSendCode{
		PhoneNumber: "+8613800138000",
		ApiId:       1000,
		ApiHash:     "hash",
		Settings:    &TLCodeSettings{},
	}

	size := obj.CalcSize(testLayer)
	if size <= 0 {
		t.Fatalf("expected positive size, got %d", size)
	}

	data, err := iface.EncodeObject(obj, testLayer)
	if err != nil {
		t.Fatalf("encode object: %v", err)
	}
	if got := len(data); got != size {
		t.Fatalf("encoded len mismatch, got %d want %d", got, size)
	}
}

func TestTLAuthSentCodeMarshalJSONIncludesClazzName(t *testing.T) {
	obj := &TLAuthSentCode{
		Type:          &TLAuthSentCodeTypeApp{Length: 6},
		PhoneCodeHash: "hash",
	}

	data, err := json.Marshal(obj)
	if err != nil {
		t.Fatalf("marshal json: %v", err)
	}
	if !strings.Contains(string(data), "\"_name\":\"auth_sentCode\"") {
		t.Fatalf("expected _name in json, got %s", data)
	}
}

func TestAuthSentCodeWrapperMarshalJSONIncludesConcreteClazzName(t *testing.T) {
	obj := (&TLAuthSentCode{
		Type:          &TLAuthSentCodeTypeApp{Length: 6},
		PhoneCodeHash: "hash",
	}).ToAuthSentCode()

	data, err := json.Marshal(obj)
	if err != nil {
		t.Fatalf("marshal wrapper json: %v", err)
	}
	if !strings.Contains(string(data), "\"_name\":\"auth_sentCode\"") {
		t.Fatalf("expected wrapper _name in json, got %s", data)
	}
}

func TestValidateRejectsMissingRequiredFields(t *testing.T) {
	if err := (&TLAuthSendCode{
		PhoneNumber: "",
		ApiId:       1000,
		ApiHash:     "hash",
		Settings:    &TLCodeSettings{},
	}).Validate(testLayer); err == nil {
		t.Fatalf("expected phone_number validation error")
	}

	if err := (&TLAuthSendCode{
		PhoneNumber: "+8613800138000",
		ApiId:       1000,
		ApiHash:     "hash",
	}).Validate(testLayer); err == nil {
		t.Fatalf("expected settings validation error")
	}
}

func TestValidateAllowsMissingFlagsFields(t *testing.T) {
	obj := &TLAuthSentCode{
		Type:          &TLAuthSentCodeTypeApp{Length: 6},
		PhoneCodeHash: "hash",
	}

	if err := obj.Validate(testLayer); err != nil {
		t.Fatalf("expected validation success, got %v", err)
	}
}
