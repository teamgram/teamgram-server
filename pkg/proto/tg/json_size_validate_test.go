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

func TestMessageStringUsesFlatDebugJSON(t *testing.T) {
	msg := &TLMessage{
		ClazzID:    ClazzID_message_3ae56482,
		ClazzName2: ClazzName_message,
		Out:        true,
		Id:         100,
		PeerId: &TLPeerUser{
			ClazzID:    ClazzID_peerUser,
			ClazzName2: ClazzName_peerUser,
			UserId:     123456789,
		},
		Date:    1710000000,
		Message: "hello",
	}

	got := msg.String()
	if strings.Contains(got, "_object") || strings.Contains(got, "_clazz") || strings.Contains(got, "_name") {
		t.Fatalf("expected flat debug json, got %s", got)
	}

	var decoded map[string]any
	if err := json.Unmarshal([]byte(got), &decoded); err != nil {
		t.Fatalf("unmarshal debug json: %v: %s", err, got)
	}
	if decoded["@type"] != "message" {
		t.Fatalf("expected @type message, got %#v in %s", decoded["@type"], got)
	}
	if decoded["@id"] != "0x3ae56482" {
		t.Fatalf("expected @id 0x3ae56482, got %#v in %s", decoded["@id"], got)
	}
	if decoded["out"] != true {
		t.Fatalf("expected true bool field to be included, got %s", got)
	}
	if _, ok := decoded["mentioned"]; ok {
		t.Fatalf("expected false bool field to be omitted, got %s", got)
	}

	peer, ok := decoded["peer_id"].(map[string]any)
	if !ok {
		t.Fatalf("expected nested peer object, got %#v in %s", decoded["peer_id"], got)
	}
	if peer["@type"] != "peerUser" || peer["@id"] != "0x59511722" {
		t.Fatalf("expected typed peer object, got %#v in %s", peer, got)
	}
}

func TestMessageWrapperStringUnwrapsClazzForDebugJSON(t *testing.T) {
	msg := (&TLMessage{
		ClazzID:    ClazzID_message_3ae56482,
		ClazzName2: ClazzName_message,
		Id:         100,
		PeerId: &TLPeerUser{
			ClazzID:    ClazzID_peerUser,
			ClazzName2: ClazzName_peerUser,
			UserId:     123456789,
		},
		Date:    1710000000,
		Message: "hello",
	}).ToMessage()

	got := msg.String()
	if strings.Contains(got, "_object") || strings.Contains(got, "_clazz") || strings.Contains(got, "_name") {
		t.Fatalf("expected wrapper to unwrap to flat debug json, got %s", got)
	}

	var decoded map[string]any
	if err := json.Unmarshal([]byte(got), &decoded); err != nil {
		t.Fatalf("unmarshal debug json: %v: %s", err, got)
	}
	if decoded["@type"] != "message" || decoded["@id"] != "0x3ae56482" {
		t.Fatalf("expected wrapped message identity, got %#v in %s", decoded, got)
	}
	if _, ok := decoded["_clazz"]; ok {
		t.Fatalf("expected no _clazz field, got %s", got)
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
