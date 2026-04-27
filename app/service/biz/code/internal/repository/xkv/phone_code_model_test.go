package xkv

import (
	"context"
	"strings"
	"testing"

	"github.com/teamgram/marmota/pkg/stores/kv"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/code/code"
)

type fakeExtStore struct {
	kv.ExtStore
	values map[string]string
}

func (f *fakeExtStore) GetCtx(_ context.Context, key string) (string, error) {
	return f.values[key], nil
}

func (f *fakeExtStore) SetexCtx(_ context.Context, key, value string, _ int) error {
	f.values[key] = value
	return nil
}

func (f *fakeExtStore) DelCtx(_ context.Context, keys ...string) (int, error) {
	for _, key := range keys {
		delete(f.values, key)
	}
	return len(keys), nil
}

func TestPhoneCodeModelRoundTripUsesCacheDTO(t *testing.T) {
	ctx := context.Background()
	store := &fakeExtStore{values: make(map[string]string)}
	model := NewPhoneCodeModel(store, "test")

	want := code.MakeTLPhoneCodeTransaction(&code.TLPhoneCodeTransaction{
		AuthKeyId:             123,
		SessionId:             456,
		Phone:                 "13800138000",
		PhoneNumberRegistered: true,
		PhoneCode:             "12345",
		PhoneCodeHash:         "hash-value",
		PhoneCodeExpired:      180,
		PhoneCodeExtraData:    "{}",
		SentCodeType:          1,
		FlashCallPattern:      "*",
		NextCodeType:          2,
		State:                 1,
	})

	if err := model.PutPhoneCode(ctx, want.AuthKeyId, want.Phone, want); err != nil {
		t.Fatalf("PutPhoneCode() error = %v", err)
	}
	stored := store.values["test:phone_code#123:13800138000"]
	if strings.Contains(stored, `"_object"`) || strings.Contains(stored, `"_name"`) || strings.Contains(stored, `"_id"`) {
		t.Fatalf("stored phone code uses TL JSON metadata: %s", stored)
	}

	got, err := model.GetPhoneCode(ctx, want.AuthKeyId, want.Phone)
	if err != nil {
		t.Fatalf("GetPhoneCode() error = %v", err)
	}
	if got == nil {
		t.Fatal("GetPhoneCode() returned nil")
	}

	if got.AuthKeyId != want.AuthKeyId ||
		got.SessionId != want.SessionId ||
		got.Phone != want.Phone ||
		got.PhoneNumberRegistered != want.PhoneNumberRegistered ||
		got.PhoneCode != want.PhoneCode ||
		got.PhoneCodeHash != want.PhoneCodeHash ||
		got.PhoneCodeExpired != want.PhoneCodeExpired ||
		got.PhoneCodeExtraData != want.PhoneCodeExtraData ||
		got.SentCodeType != want.SentCodeType ||
		got.FlashCallPattern != want.FlashCallPattern ||
		got.NextCodeType != want.NextCodeType ||
		got.State != want.State {
		t.Fatalf("GetPhoneCode() = %+v, want %+v", got, want)
	}
}
