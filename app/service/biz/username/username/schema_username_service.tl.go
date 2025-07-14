/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package username

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/tg"
)

var (
	_ iface.TLObject
	_ fmt.Stringer
	_ *tg.Bool
	_ bin.Fields
	_ json.Marshaler
)

// TLUsernameGetAccountUsername <--
type TLUsernameGetAccountUsername struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id"`
}

func (m *TLUsernameGetAccountUsername) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLUsernameGetAccountUsername) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x92ef8d5: func() error {
			x.PutClazzID(0x92ef8d5)

			x.PutInt64(m.UserId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_username_getAccountUsername, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_username_getAccountUsername, layer)
	}
}

// Decode <--
func (m *TLUsernameGetAccountUsername) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x92ef8d5: func() (err error) {
			m.UserId, err = d.Int64()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUsernameCheckAccountUsername <--
type TLUsernameCheckAccountUsername struct {
	ClazzID  uint32 `json:"_id"`
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
}

func (m *TLUsernameCheckAccountUsername) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLUsernameCheckAccountUsername) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x49f7f105: func() error {
			x.PutClazzID(0x49f7f105)

			x.PutInt64(m.UserId)
			x.PutString(m.Username)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_username_checkAccountUsername, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_username_checkAccountUsername, layer)
	}
}

// Decode <--
func (m *TLUsernameCheckAccountUsername) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x49f7f105: func() (err error) {
			m.UserId, err = d.Int64()
			m.Username, err = d.String()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUsernameGetChannelUsername <--
type TLUsernameGetChannelUsername struct {
	ClazzID   uint32 `json:"_id"`
	ChannelId int64  `json:"channel_id"`
}

func (m *TLUsernameGetChannelUsername) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLUsernameGetChannelUsername) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x868487d5: func() error {
			x.PutClazzID(0x868487d5)

			x.PutInt64(m.ChannelId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_username_getChannelUsername, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_username_getChannelUsername, layer)
	}
}

// Decode <--
func (m *TLUsernameGetChannelUsername) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x868487d5: func() (err error) {
			m.ChannelId, err = d.Int64()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUsernameCheckChannelUsername <--
type TLUsernameCheckChannelUsername struct {
	ClazzID   uint32 `json:"_id"`
	ChannelId int64  `json:"channel_id"`
	Username  string `json:"username"`
}

func (m *TLUsernameCheckChannelUsername) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLUsernameCheckChannelUsername) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x26d4be9d: func() error {
			x.PutClazzID(0x26d4be9d)

			x.PutInt64(m.ChannelId)
			x.PutString(m.Username)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_username_checkChannelUsername, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_username_checkChannelUsername, layer)
	}
}

// Decode <--
func (m *TLUsernameCheckChannelUsername) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x26d4be9d: func() (err error) {
			m.ChannelId, err = d.Int64()
			m.Username, err = d.String()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUsernameUpdateUsernameByPeer <--
type TLUsernameUpdateUsernameByPeer struct {
	ClazzID  uint32 `json:"_id"`
	PeerType int32  `json:"peer_type"`
	PeerId   int64  `json:"peer_id"`
	Username string `json:"username"`
}

func (m *TLUsernameUpdateUsernameByPeer) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLUsernameUpdateUsernameByPeer) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x6669bddc: func() error {
			x.PutClazzID(0x6669bddc)

			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)
			x.PutString(m.Username)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_username_updateUsernameByPeer, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_username_updateUsernameByPeer, layer)
	}
}

// Decode <--
func (m *TLUsernameUpdateUsernameByPeer) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x6669bddc: func() (err error) {
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()
			m.Username, err = d.String()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUsernameCheckUsername <--
type TLUsernameCheckUsername struct {
	ClazzID  uint32 `json:"_id"`
	Username string `json:"username"`
}

func (m *TLUsernameCheckUsername) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLUsernameCheckUsername) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x28caa6d5: func() error {
			x.PutClazzID(0x28caa6d5)

			x.PutString(m.Username)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_username_checkUsername, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_username_checkUsername, layer)
	}
}

// Decode <--
func (m *TLUsernameCheckUsername) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x28caa6d5: func() (err error) {
			m.Username, err = d.String()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUsernameUpdateUsername <--
type TLUsernameUpdateUsername struct {
	ClazzID  uint32 `json:"_id"`
	PeerType int32  `json:"peer_type"`
	PeerId   int64  `json:"peer_id"`
	Username string `json:"username"`
}

func (m *TLUsernameUpdateUsername) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLUsernameUpdateUsername) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x52d65433: func() error {
			x.PutClazzID(0x52d65433)

			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)
			x.PutString(m.Username)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_username_updateUsername, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_username_updateUsername, layer)
	}
}

// Decode <--
func (m *TLUsernameUpdateUsername) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x52d65433: func() (err error) {
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()
			m.Username, err = d.String()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUsernameDeleteUsername <--
type TLUsernameDeleteUsername struct {
	ClazzID  uint32 `json:"_id"`
	Username string `json:"username"`
}

func (m *TLUsernameDeleteUsername) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLUsernameDeleteUsername) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xc0777388: func() error {
			x.PutClazzID(0xc0777388)

			x.PutString(m.Username)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_username_deleteUsername, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_username_deleteUsername, layer)
	}
}

// Decode <--
func (m *TLUsernameDeleteUsername) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xc0777388: func() (err error) {
			m.Username, err = d.String()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUsernameResolveUsername <--
type TLUsernameResolveUsername struct {
	ClazzID  uint32 `json:"_id"`
	Username string `json:"username"`
}

func (m *TLUsernameResolveUsername) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLUsernameResolveUsername) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x77ba2cc6: func() error {
			x.PutClazzID(0x77ba2cc6)

			x.PutString(m.Username)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_username_resolveUsername, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_username_resolveUsername, layer)
	}
}

// Decode <--
func (m *TLUsernameResolveUsername) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x77ba2cc6: func() (err error) {
			m.Username, err = d.String()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUsernameGetListByUsernameList <--
type TLUsernameGetListByUsernameList struct {
	ClazzID uint32   `json:"_id"`
	Names   []string `json:"names"`
}

func (m *TLUsernameGetListByUsernameList) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLUsernameGetListByUsernameList) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x48a7974d: func() error {
			x.PutClazzID(0x48a7974d)

			iface.EncodeStringList(x, m.Names)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_username_getListByUsernameList, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_username_getListByUsernameList, layer)
	}
}

// Decode <--
func (m *TLUsernameGetListByUsernameList) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x48a7974d: func() (err error) {

			m.Names, err = iface.DecodeStringList(d)

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUsernameDeleteUsernameByPeer <--
type TLUsernameDeleteUsernameByPeer struct {
	ClazzID  uint32 `json:"_id"`
	PeerType int32  `json:"peer_type"`
	PeerId   int64  `json:"peer_id"`
}

func (m *TLUsernameDeleteUsernameByPeer) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLUsernameDeleteUsernameByPeer) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x1e44c06d: func() error {
			x.PutClazzID(0x1e44c06d)

			x.PutInt32(m.PeerType)
			x.PutInt64(m.PeerId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_username_deleteUsernameByPeer, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_username_deleteUsernameByPeer, layer)
	}
}

// Decode <--
func (m *TLUsernameDeleteUsernameByPeer) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x1e44c06d: func() (err error) {
			m.PeerType, err = d.Int32()
			m.PeerId, err = d.Int64()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLUsernameSearch <--
type TLUsernameSearch struct {
	ClazzID          uint32  `json:"_id"`
	Q                string  `json:"q"`
	ExcludedContacts []int64 `json:"excluded_contacts"`
	Limit            int32   `json:"limit"`
}

func (m *TLUsernameSearch) String() string {
	wrapper := iface.WithNameWrapper{"", m}
	return wrapper.String()
}

// Encode <--
func (m *TLUsernameSearch) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xe8a5a306: func() error {
			x.PutClazzID(0xe8a5a306)

			x.PutString(m.Q)

			iface.EncodeInt64List(x, m.ExcludedContacts)

			x.PutInt32(m.Limit)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_username_search, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_username_search, layer)
	}
}

// Decode <--
func (m *TLUsernameSearch) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xe8a5a306: func() (err error) {
			m.Q, err = d.String()

			m.ExcludedContacts, err = iface.DecodeInt64List(d)

			m.Limit, err = d.Int32()

			return nil
		},
	}

	if m.ClazzID == 0 {
		m.ClazzID, _ = d.ClazzID()
	}
	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// Vector api result type
// ----------------------------------------------------------------------------
// VectorResList <--

// VectorUsernameData <--
type VectorUsernameData struct {
	Datas []UsernameDataClazz `json:"_datas"`
}

func (m *VectorUsernameData) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

// Encode <--
func (m *VectorUsernameData) Encode(x *bin.Encoder, layer int32) error {
	_ = iface.EncodeObjectList(x, m.Datas, layer)

	return nil
}

// Decode <--
func (m *VectorUsernameData) Decode(d *bin.Decoder) (err error) {
	m.Datas, err = iface.DecodeObjectList[UsernameDataClazz](d)

	return err
}

// ----------------------------------------------------------------------------
// rpc

type RPCUsername interface {
	UsernameGetAccountUsername(ctx context.Context, in *TLUsernameGetAccountUsername) (*UsernameData, error)
	UsernameCheckAccountUsername(ctx context.Context, in *TLUsernameCheckAccountUsername) (*UsernameExist, error)
	UsernameGetChannelUsername(ctx context.Context, in *TLUsernameGetChannelUsername) (*UsernameData, error)
	UsernameCheckChannelUsername(ctx context.Context, in *TLUsernameCheckChannelUsername) (*UsernameExist, error)
	UsernameUpdateUsernameByPeer(ctx context.Context, in *TLUsernameUpdateUsernameByPeer) (*tg.Bool, error)
	UsernameCheckUsername(ctx context.Context, in *TLUsernameCheckUsername) (*UsernameExist, error)
	UsernameUpdateUsername(ctx context.Context, in *TLUsernameUpdateUsername) (*tg.Bool, error)
	UsernameDeleteUsername(ctx context.Context, in *TLUsernameDeleteUsername) (*tg.Bool, error)
	UsernameResolveUsername(ctx context.Context, in *TLUsernameResolveUsername) (*tg.Peer, error)
	UsernameGetListByUsernameList(ctx context.Context, in *TLUsernameGetListByUsernameList) (*VectorUsernameData, error)
	UsernameDeleteUsernameByPeer(ctx context.Context, in *TLUsernameDeleteUsernameByPeer) (*tg.Bool, error)
	UsernameSearch(ctx context.Context, in *TLUsernameSearch) (*VectorUsernameData, error)
}
