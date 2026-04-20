/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package inbox

import (
	"encoding/json"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

var (
	_ iface.TLObject
	_ fmt.Stringer
	_ *tg.Bool
	_ bin.Fields
	_ json.Marshaler
)

// InboxMessageDataClazz <--
//   - TL_InboxMessageData
type InboxMessageDataClazz = *TLInboxMessageData

func DecodeInboxMessageDataClazz(d *bin.Decoder) (InboxMessageDataClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0x3bbdadd4:
		x := &TLInboxMessageData{ClazzID: id, ClazzName2: ClazzName_inboxMessageData}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeInboxMessageData - unexpected clazzId: %d", id)
	}

}

// TLInboxMessageData <--
type TLInboxMessageData struct {
	ClazzID         uint32          `json:"_id"`
	ClazzName2      string          `json:"_name"`
	RandomId        int64           `json:"random_id"`
	DialogMessageId int64           `json:"dialog_message_id"`
	Message         tg.MessageClazz `json:"message"`
}

func MakeTLInboxMessageData(m *TLInboxMessageData) *TLInboxMessageData {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_inboxMessageData

	return m
}

func (m *TLInboxMessageData) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *TLInboxMessageData) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("inboxMessageData", m)
}

// InboxMessageDataClazzName <--
func (m *TLInboxMessageData) InboxMessageDataClazzName() string {
	return ClazzName_inboxMessageData
}

// ClazzName <--
func (m *TLInboxMessageData) ClazzName() string {
	return m.ClazzName2
}

// ToInboxMessageData <--
func (m *TLInboxMessageData) ToInboxMessageData() *InboxMessageData {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLInboxMessageData) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inboxMessageData, int(layer)); clazzId {
	case 0x3bbdadd4:
		size := 4
		size += 8
		size += 8
		size += iface.CalcObjectSize(m.Message, layer)

		return size
	default:
		return 0
	}
}

func (m *TLInboxMessageData) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inboxMessageData, int(layer)); clazzId {
	case 0x3bbdadd4:
		if err := iface.ValidateRequiredObject("message", m.Message); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inboxMessageData, layer)
	}
}

// Encode <--
func (m *TLInboxMessageData) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inboxMessageData, int(layer)); clazzId {
	case 0x3bbdadd4:
		x.PutClazzID(0x3bbdadd4)

		x.PutInt64(m.RandomId)
		x.PutInt64(m.DialogMessageId)
		_ = m.Message.Encode(x, layer)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inboxMessageData, layer)
	}
}

// Decode <--
func (m *TLInboxMessageData) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0x3bbdadd4:
		m.RandomId, err = d.Int64()
		if err != nil {
			return err
		}
		m.DialogMessageId, err = d.Int64()
		if err != nil {
			return err
		}

		m.Message, err = tg.DecodeMessageClazz(d)
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// InboxMessageData <--
type InboxMessageData = TLInboxMessageData

// InboxMessageIdClazz <--
//   - TL_InboxMessageId
type InboxMessageIdClazz = *TLInboxMessageId

func DecodeInboxMessageIdClazz(d *bin.Decoder) (InboxMessageIdClazz, error) {
	// id, err := d.PeekClazzID()
	id, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	switch id {
	case 0xc692c19f:
		x := &TLInboxMessageId{ClazzID: id, ClazzName2: ClazzName_inboxMessageId}
		if err := x.Decode(d); err != nil {
			return nil, err
		}
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeInboxMessageId - unexpected clazzId: %d", id)
	}

}

// TLInboxMessageId <--
type TLInboxMessageId struct {
	ClazzID         uint32 `json:"_id"`
	ClazzName2      string `json:"_name"`
	Id              int32  `json:"id"`
	DialogMessageId int64  `json:"dialog_message_id"`
}

func MakeTLInboxMessageId(m *TLInboxMessageId) *TLInboxMessageId {
	if m == nil {
		return nil
	}
	m.ClazzName2 = ClazzName_inboxMessageId

	return m
}

func (m *TLInboxMessageId) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *TLInboxMessageId) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return iface.MarshalWithName("inboxMessageId", m)
}

// InboxMessageIdClazzName <--
func (m *TLInboxMessageId) InboxMessageIdClazzName() string {
	return ClazzName_inboxMessageId
}

// ClazzName <--
func (m *TLInboxMessageId) ClazzName() string {
	return m.ClazzName2
}

// ToInboxMessageId <--
func (m *TLInboxMessageId) ToInboxMessageId() *InboxMessageId {
	if m == nil {
		return nil
	}

	return m

}

func (m *TLInboxMessageId) CalcSize(layer int32) int {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inboxMessageId, int(layer)); clazzId {
	case 0xc692c19f:
		size := 4
		size += 4
		size += 8

		return size
	default:
		return 0
	}
}

func (m *TLInboxMessageId) Validate(layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inboxMessageId, int(layer)); clazzId {
	case 0xc692c19f:

		return nil
	default:
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inboxMessageId, layer)
	}
}

// Encode <--
func (m *TLInboxMessageId) Encode(x *bin.Encoder, layer int32) error {
	switch clazzId := iface.GetClazzIDByName(ClazzName_inboxMessageId, int(layer)); clazzId {
	case 0xc692c19f:
		x.PutClazzID(0xc692c19f)

		x.PutInt32(m.Id)
		x.PutInt64(m.DialogMessageId)

		return nil
	default:
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_inboxMessageId, layer)
	}
}

// Decode <--
func (m *TLInboxMessageId) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case 0xc692c19f:
		m.Id, err = d.Int32()
		if err != nil {
			return err
		}
		m.DialogMessageId, err = d.Int64()
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// InboxMessageId <--
type InboxMessageId = TLInboxMessageId
