/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package inbox

import (
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/tg"
)

var _ iface.TLObject
var _ fmt.Stringer
var _ *tg.Bool
var _ bin.Fields

// InboxMessageDataClazz <--
//   - TL_InboxMessageData
type InboxMessageDataClazz interface {
	iface.TLObject
	InboxMessageDataClazzName() string
}

func DecodeInboxMessageDataClazz(d *bin.Decoder) (InboxMessageDataClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_inboxMessageData:
		x := &TLInboxMessageData{ClazzID: id}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeInboxMessageData - unexpected clazzId: %d", id)
	}
}

// TLInboxMessageData <--
type TLInboxMessageData struct {
	ClazzID         uint32      `json:"_id"`
	RandomId        int64       `json:"random_id"`
	DialogMessageId int64       `json:"dialog_message_id"`
	Message         *tg.Message `json:"message"`
}

// InboxMessageDataClazzName <--
func (m *TLInboxMessageData) InboxMessageDataClazzName() string {
	return ClazzName_inboxMessageData
}

// ClazzName <--
func (m *TLInboxMessageData) ClazzName() string {
	return ClazzName_inboxMessageData
}

// ToInboxMessageData <--
func (m *TLInboxMessageData) ToInboxMessageData() *InboxMessageData {
	return MakeInboxMessageData(m)
}

// Encode <--
func (m *TLInboxMessageData) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x3bbdadd4: func() error {
			x.PutClazzID(0x3bbdadd4)

			x.PutInt64(m.RandomId)
			x.PutInt64(m.DialogMessageId)
			_ = m.Message.Encode(x, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_inboxMessageData, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inboxMessageData, layer)
	}
}

// Decode <--
func (m *TLInboxMessageData) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x3bbdadd4: func() (err error) {
			m.RandomId, err = d.Int64()
			m.DialogMessageId, err = d.Int64()

			m2 := &tg.Message{}
			_ = m2.Decode(d)
			m.Message = m2

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// InboxMessageData <--
type InboxMessageData struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	InboxMessageDataClazz
}

// MakeInboxMessageData <--
func MakeInboxMessageData(c InboxMessageDataClazz) *InboxMessageData {
	return &InboxMessageData{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		InboxMessageDataClazz: c,
	}
}

// Encode <--
func (m *InboxMessageData) Encode(x *bin.Encoder, layer int32) error {
	if m.InboxMessageDataClazz != nil {
		return m.InboxMessageDataClazz.Encode(x, layer)
	}

	return fmt.Errorf("InboxMessageData - invalid Clazz")
}

// Decode <--
func (m *InboxMessageData) Decode(d *bin.Decoder) (err error) {
	m.InboxMessageDataClazz, err = DecodeInboxMessageDataClazz(d)
	return
}

// Match <--
func (m *InboxMessageData) Match(f ...interface{}) {
	switch c := m.InboxMessageDataClazz.(type) {
	case *TLInboxMessageData:
		for _, v := range f {
			if f1, ok := v.(func(c *TLInboxMessageData) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToInboxMessageData <--
func (m *InboxMessageData) ToInboxMessageData() (*TLInboxMessageData, bool) {
	if m.InboxMessageDataClazz == nil {
		return nil, false
	}

	if x, ok := m.InboxMessageDataClazz.(*TLInboxMessageData); ok {
		return x, true
	}

	return nil, false
}

// InboxMessageIdClazz <--
//   - TL_InboxMessageId
type InboxMessageIdClazz interface {
	iface.TLObject
	InboxMessageIdClazzName() string
}

func DecodeInboxMessageIdClazz(d *bin.Decoder) (InboxMessageIdClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_inboxMessageId:
		x := &TLInboxMessageId{ClazzID: id}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeInboxMessageId - unexpected clazzId: %d", id)
	}
}

// TLInboxMessageId <--
type TLInboxMessageId struct {
	ClazzID         uint32 `json:"_id"`
	Id              int32  `json:"id"`
	DialogMessageId int64  `json:"dialog_message_id"`
}

// InboxMessageIdClazzName <--
func (m *TLInboxMessageId) InboxMessageIdClazzName() string {
	return ClazzName_inboxMessageId
}

// ClazzName <--
func (m *TLInboxMessageId) ClazzName() string {
	return ClazzName_inboxMessageId
}

// ToInboxMessageId <--
func (m *TLInboxMessageId) ToInboxMessageId() *InboxMessageId {
	return MakeInboxMessageId(m)
}

// Encode <--
func (m *TLInboxMessageId) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xc692c19f: func() error {
			x.PutClazzID(0xc692c19f)

			x.PutInt32(m.Id)
			x.PutInt64(m.DialogMessageId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_inboxMessageId, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inboxMessageId, layer)
	}
}

// Decode <--
func (m *TLInboxMessageId) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xc692c19f: func() (err error) {
			m.Id, err = d.Int32()
			m.DialogMessageId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// InboxMessageId <--
type InboxMessageId struct {
	// ClazzID   uint32 `json:"_id"`
	// ClazzName string `json:"_name"`
	InboxMessageIdClazz
}

// MakeInboxMessageId <--
func MakeInboxMessageId(c InboxMessageIdClazz) *InboxMessageId {
	return &InboxMessageId{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		InboxMessageIdClazz: c,
	}
}

// Encode <--
func (m *InboxMessageId) Encode(x *bin.Encoder, layer int32) error {
	if m.InboxMessageIdClazz != nil {
		return m.InboxMessageIdClazz.Encode(x, layer)
	}

	return fmt.Errorf("InboxMessageId - invalid Clazz")
}

// Decode <--
func (m *InboxMessageId) Decode(d *bin.Decoder) (err error) {
	m.InboxMessageIdClazz, err = DecodeInboxMessageIdClazz(d)
	return
}

// Match <--
func (m *InboxMessageId) Match(f ...interface{}) {
	switch c := m.InboxMessageIdClazz.(type) {
	case *TLInboxMessageId:
		for _, v := range f {
			if f1, ok := v.(func(c *TLInboxMessageId) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ToInboxMessageId <--
func (m *InboxMessageId) ToInboxMessageId() (*TLInboxMessageId, bool) {
	if m.InboxMessageIdClazz == nil {
		return nil, false
	}

	if x, ok := m.InboxMessageIdClazz.(*TLInboxMessageId); ok {
		return x, true
	}

	return nil, false
}
