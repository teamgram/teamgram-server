package core

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/code/code"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/code/internal/svc"
)

// mockRepo implements the subset of repository methods used by core handlers.
type mockRepo struct {
	store  map[string][]byte
	getErr error
	putErr error
	delErr error
}

func (m *mockRepo) cacheKey(authKeyId int64, phone string) string {
	return fmt.Sprintf("phone_code#%d:%s", authKeyId, phone)
}

func (m *mockRepo) GetCachePhoneCode(_ context.Context, authKeyId int64, phone string) (*code.PhoneCodeTransaction, error) {
	if m.getErr != nil {
		return nil, m.getErr
	}
	val, ok := m.store[m.cacheKey(authKeyId, phone)]
	if !ok || val == nil {
		return nil, nil
	}
	// TLPhoneCodeTransaction.MarshalJSON uses iface.MarshalWithName which
	// nests the data under "_object". Unmarshal accordingly.
	var wrapper struct {
		Object code.PhoneCodeTransaction `json:"_object"`
	}
	if err := json.Unmarshal(val, &wrapper); err != nil {
		return nil, err
	}
	return &wrapper.Object, nil
}

func (m *mockRepo) PutCachePhoneCode(_ context.Context, authKeyId int64, phone string, data *code.PhoneCodeTransaction) error {
	if m.putErr != nil {
		return m.putErr
	}
	b, _ := json.Marshal(data)
	m.store[m.cacheKey(authKeyId, phone)] = b
	return nil
}

func (m *mockRepo) DeleteCachePhoneCode(_ context.Context, authKeyId int64, phone string) error {
	delete(m.store, m.cacheKey(authKeyId, phone))
	return m.delErr
}

func newTestCore(t *testing.T, repo svc.Repo) *CodeCore {
	t.Helper()
	return New(context.Background(), &svc.ServiceContext{
		Repo: repo,
	})
}

func TestCreatePhoneCode_NewCode(t *testing.T) {
	c := newTestCore(t, &mockRepo{store: make(map[string][]byte)})

	in := &code.TLCodeCreatePhoneCode{
		AuthKeyId:    1,
		SessionId:    100,
		Phone:        "1234567890",
		SentCodeType: 1,
	}

	result, err := c.CodeCreatePhoneCode(in)
	if err != nil {
		t.Fatalf("CodeCreatePhoneCode() err = %v", err)
	}
	if result == nil {
		t.Fatal("CodeCreatePhoneCode() result is nil")
	}
	if result.PhoneCode == "" {
		t.Error("PhoneCode is empty")
	}
	if result.PhoneCodeHash == "" {
		t.Error("PhoneCodeHash is empty")
	}
	if result.PhoneCodeExpired <= int32(time.Now().Unix()) {
		t.Error("PhoneCodeExpired is in the past")
	}
	if result.State != codeStateSend {
		t.Errorf("State = %d, want %d", result.State, codeStateSend)
	}
}

func TestCreatePhoneCode_SessionChanged(t *testing.T) {
	store := make(map[string][]byte)
	existing, _ := json.Marshal(&code.TLPhoneCodeTransaction{
		AuthKeyId:     1,
		SessionId:     200,
		Phone:         "1234567890",
		PhoneCode:     "12345",
		PhoneCodeHash: "abc",
	})
	store["phone_code#1:1234567890"] = existing

	c := newTestCore(t, &mockRepo{store: store})

	in := &code.TLCodeCreatePhoneCode{
		AuthKeyId:    1,
		SessionId:    100,
		Phone:        "1234567890",
		SentCodeType: 1,
	}

	result, err := c.CodeCreatePhoneCode(in)
	if err != nil {
		t.Fatalf("CodeCreatePhoneCode() err = %v", err)
	}
	if result.SessionId != 100 {
		t.Errorf("SessionId = %d, want 100", result.SessionId)
	}
}

func TestGetPhoneCode_Success(t *testing.T) {
	store := make(map[string][]byte)
	existing, _ := json.Marshal(&code.TLPhoneCodeTransaction{
		AuthKeyId:     1,
		Phone:         "1234567890",
		PhoneCodeHash: "abc",
	})
	store["phone_code#1:1234567890"] = existing

	c := newTestCore(t, &mockRepo{store: store})

	result, err := c.CodeGetPhoneCode(&code.TLCodeGetPhoneCode{
		AuthKeyId:     1,
		Phone:         "1234567890",
		PhoneCodeHash: "abc",
	})
	if err != nil {
		t.Fatalf("CodeGetPhoneCode() err = %v", err)
	}
	if result.PhoneCodeHash != "abc" {
		t.Errorf("PhoneCodeHash = %s, want abc", result.PhoneCodeHash)
	}
}

func TestGetPhoneCode_Expired(t *testing.T) {
	c := newTestCore(t, &mockRepo{store: make(map[string][]byte)})

	_, err := c.CodeGetPhoneCode(&code.TLCodeGetPhoneCode{
		AuthKeyId:     1,
		Phone:         "1234567890",
		PhoneCodeHash: "abc",
	})
	if !errors.Is(err, code.ErrPhoneCodeExpired) {
		t.Errorf("err = %v, want ErrPhoneCodeExpired", err)
	}
}

func TestGetPhoneCode_InvalidHash(t *testing.T) {
	store := make(map[string][]byte)
	existing, _ := json.Marshal(&code.TLPhoneCodeTransaction{
		AuthKeyId:     1,
		Phone:         "1234567890",
		PhoneCodeHash: "abc",
	})
	store["phone_code#1:1234567890"] = existing

	c := newTestCore(t, &mockRepo{store: store})

	_, err := c.CodeGetPhoneCode(&code.TLCodeGetPhoneCode{
		AuthKeyId:     1,
		Phone:         "1234567890",
		PhoneCodeHash: "wrong",
	})
	if !errors.Is(err, code.ErrPhoneCodeInvalid) {
		t.Errorf("err = %v, want ErrPhoneCodeInvalid", err)
	}
}

func TestGetPhoneCode_StorageError(t *testing.T) {
	storageErr := errors.New("redis connection refused")
	c := newTestCore(t, &mockRepo{store: make(map[string][]byte), getErr: storageErr})

	_, err := c.CodeGetPhoneCode(&code.TLCodeGetPhoneCode{
		AuthKeyId:     1,
		Phone:         "1234567890",
		PhoneCodeHash: "abc",
	})
	if !errors.Is(err, code.ErrPhoneCodeExpired) {
		t.Errorf("err = %v, want ErrPhoneCodeExpired", err)
	}
}

func TestDeletePhoneCode(t *testing.T) {
	c := newTestCore(t, &mockRepo{store: make(map[string][]byte)})

	result, err := c.CodeDeletePhoneCode(&code.TLCodeDeletePhoneCode{
		AuthKeyId:     1,
		Phone:         "1234567890",
		PhoneCodeHash: "abc",
	})
	if err != nil {
		t.Fatalf("CodeDeletePhoneCode() err = %v", err)
	}
	if result == nil {
		t.Fatal("CodeDeletePhoneCode() result is nil")
	}
}

func TestUpdatePhoneCodeData_StorageError(t *testing.T) {
	storageErr := errors.New("redis connection refused")
	c := newTestCore(t, &mockRepo{store: make(map[string][]byte), putErr: storageErr})

	_, err := c.CodeUpdatePhoneCodeData(&code.TLCodeUpdatePhoneCodeData{
		AuthKeyId: 1,
		Phone:     "1234567890",
		CodeData: code.MakeTLPhoneCodeTransaction(&code.TLPhoneCodeTransaction{
			AuthKeyId: 1,
			Phone:     "1234567890",
		}),
	})
	if !errors.Is(err, code.ErrCodeStorage) {
		t.Errorf("err = %v, want ErrCodeStorage", err)
	}
}
