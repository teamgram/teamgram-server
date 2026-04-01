package tg

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
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

func TestTLInvokeAfterMsgCalcSizeAndRoundTripIncludeRawQuery(t *testing.T) {
	query := []byte{0x78, 0x56, 0x34, 0x12}
	obj := &TLInvokeAfterMsg{
		MsgId: 12345,
		Query: query,
	}

	size := obj.CalcSize(testLayer)
	data, err := iface.EncodeObject(obj, testLayer)
	if err != nil {
		t.Fatalf("encode object: %v", err)
	}
	if got := len(data); got != size {
		t.Fatalf("encoded len mismatch, got %d want %d", got, size)
	}
	if got := data[len(data)-len(query):]; string(got) != string(query) {
		t.Fatalf("expected encoded payload suffix %v, got %v", query, got)
	}

	var decoded TLInvokeAfterMsg
	if err := decoded.Decode(bin.NewDecoder(data)); err != nil {
		t.Fatalf("decode object: %v", err)
	}
	if string(decoded.Query) != string(query) {
		t.Fatalf("expected decoded query %v, got %v", query, decoded.Query)
	}
}

func TestTLInvokeAfterMsgMarshalJSONIncludesClazzName(t *testing.T) {
	obj := &TLInvokeAfterMsg{
		MsgId: 1,
		Query: []byte{1, 2, 3, 4},
	}

	data, err := json.Marshal(obj)
	if err != nil {
		t.Fatalf("marshal json: %v", err)
	}
	if !strings.Contains(string(data), "\"_name\":\"invokeAfterMsg\"") {
		t.Fatalf("expected _name in json, got %s", data)
	}
}

func TestValidateRecursesIntoRequiredObjectSlices(t *testing.T) {
	obj := &TLAccountPasskeys{
		Passkeys: []PasskeyClazz{nil},
	}

	if err := obj.Validate(testLayer); err == nil {
		t.Fatalf("expected recursive slice validation error")
	}
}
