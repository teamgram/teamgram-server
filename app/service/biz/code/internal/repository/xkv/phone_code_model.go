package xkv

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/teamgram/marmota/pkg/stores/kv"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/code/code"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	phoneCodeDefaultTTL = 180
)

type phoneCodeTransactionCacheData struct {
	AuthKeyId             int64  `json:"auth_key_id"`
	SessionId             int64  `json:"session_id"`
	Phone                 string `json:"phone"`
	PhoneNumberRegistered bool   `json:"phone_number_registered"`
	PhoneCode             string `json:"phone_code"`
	PhoneCodeHash         string `json:"phone_code_hash"`
	PhoneCodeExpired      int32  `json:"phone_code_expired"`
	PhoneCodeExtraData    string `json:"phone_code_extra_data"`
	SentCodeType          int32  `json:"sent_code_type"`
	FlashCallPattern      string `json:"flash_call_pattern"`
	NextCodeType          int32  `json:"next_code_type"`
	State                 int32  `json:"state"`
}

func phoneCodeTransactionCacheDataFromTL(data *code.PhoneCodeTransaction) *phoneCodeTransactionCacheData {
	if data == nil {
		return nil
	}
	return &phoneCodeTransactionCacheData{
		AuthKeyId:             data.AuthKeyId,
		SessionId:             data.SessionId,
		Phone:                 data.Phone,
		PhoneNumberRegistered: data.PhoneNumberRegistered,
		PhoneCode:             data.PhoneCode,
		PhoneCodeHash:         data.PhoneCodeHash,
		PhoneCodeExpired:      data.PhoneCodeExpired,
		PhoneCodeExtraData:    data.PhoneCodeExtraData,
		SentCodeType:          data.SentCodeType,
		FlashCallPattern:      data.FlashCallPattern,
		NextCodeType:          data.NextCodeType,
		State:                 data.State,
	}
}

func (data *phoneCodeTransactionCacheData) toTL() *code.PhoneCodeTransaction {
	if data == nil {
		return nil
	}
	return code.MakeTLPhoneCodeTransaction(&code.TLPhoneCodeTransaction{
		AuthKeyId:             data.AuthKeyId,
		SessionId:             data.SessionId,
		Phone:                 data.Phone,
		PhoneNumberRegistered: data.PhoneNumberRegistered,
		PhoneCode:             data.PhoneCode,
		PhoneCodeHash:         data.PhoneCodeHash,
		PhoneCodeExpired:      data.PhoneCodeExpired,
		PhoneCodeExtraData:    data.PhoneCodeExtraData,
		SentCodeType:          data.SentCodeType,
		FlashCallPattern:      data.FlashCallPattern,
		NextCodeType:          data.NextCodeType,
		State:                 data.State,
	})
}

// PhoneCodeModel abstracts KV operations for phone verification codes.
type PhoneCodeModel interface {
	GetPhoneCode(ctx context.Context, authKeyId int64, phone string) (*code.PhoneCodeTransaction, error)
	PutPhoneCode(ctx context.Context, authKeyId int64, phone string, data *code.PhoneCodeTransaction) error
	DeletePhoneCode(ctx context.Context, authKeyId int64, phone string) error
}

type phoneCodeModel struct {
	kv     kv.ExtStore
	prefix string
}

// NewPhoneCodeModel creates a kv-backed phone code model.
func NewPhoneCodeModel(kv kv.ExtStore, prefix string) PhoneCodeModel {
	return &phoneCodeModel{
		kv:     kv,
		prefix: prefix,
	}
}

func (m *phoneCodeModel) cacheKey(authKeyId int64, phone string) string {
	if m.prefix == "" {
		return fmt.Sprintf("phone_code#%d:%s", authKeyId, phone)
	}
	return fmt.Sprintf("%s:phone_code#%d:%s", m.prefix, authKeyId, phone)
}

func (m *phoneCodeModel) GetPhoneCode(ctx context.Context, authKeyId int64, phone string) (*code.PhoneCodeTransaction, error) {
	val, err := m.kv.GetCtx(ctx, m.cacheKey(authKeyId, phone))
	if err != nil {
		return nil, fmt.Errorf("phone_code.GetPhoneCode kv get: %w", err)
	}
	if val == "" {
		return nil, nil
	}

	txn, err := unmarshalPhoneCodeTransaction([]byte(val))
	if err != nil {
		logx.WithContext(ctx).Errorf("phone_code.GetPhoneCode json unmarshal error(%v)", err)
		return nil, fmt.Errorf("phone_code.GetPhoneCode json unmarshal: %w", err)
	}

	return txn, nil
}

func (m *phoneCodeModel) PutPhoneCode(ctx context.Context, authKeyId int64, phone string, data *code.PhoneCodeTransaction) error {
	if data == nil {
		return nil
	}

	b, err := marshalPhoneCodeTransaction(data)
	if err != nil {
		return fmt.Errorf("phone_code.PutPhoneCode json marshal: %w", err)
	}

	return m.kv.SetexCtx(ctx, m.cacheKey(authKeyId, phone), string(b), phoneCodeDefaultTTL)
}

func (m *phoneCodeModel) DeletePhoneCode(ctx context.Context, authKeyId int64, phone string) error {
	_, err := m.kv.DelCtx(ctx, m.cacheKey(authKeyId, phone))
	return err
}

func marshalPhoneCodeTransaction(data *code.PhoneCodeTransaction) ([]byte, error) {
	if data == nil {
		return []byte("null"), nil
	}

	return json.Marshal(phoneCodeTransactionCacheDataFromTL(data))
}

func unmarshalPhoneCodeTransaction(data []byte) (*code.PhoneCodeTransaction, error) {
	var txn phoneCodeTransactionCacheData
	if err := json.Unmarshal(data, &txn); err != nil {
		return nil, err
	}
	return txn.toTL(), nil
}
