package tg

import (
	"encoding/json"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
)

type TLInputMessageReadMetric struct {
	ClazzID uint32 `json:"_id"`
	MsgId   int32  `json:"msg_id"`
	Count   int32  `json:"count"`
}

type InputMessageReadMetricClazz = *TLInputMessageReadMetric

func (m *TLInputMessageReadMetric) Encode(x *bin.Encoder, layer int32) error {
	x.PutClazzID(0x4067c5e6)
	x.PutInt32(m.MsgId)
	x.PutInt32(m.Count)
	return nil
}

func (m *TLInputMessageReadMetric) Decode(d *bin.Decoder) error {
	var err error
	m.MsgId, err = d.Int32()
	if err != nil {
		return err
	}
	m.Count, err = d.Int32()
	return err
}

type TLMessagesComposeMessageWithAI struct {
	ClazzID         uint32                `json:"_id"`
	Proofread       bool                  `json:"proofread"`
	Emojify         bool                  `json:"emojify"`
	Text            TextWithEntitiesClazz `json:"text"`
	TranslateToLang *string               `json:"translate_to_lang"`
	ChangeTone      *string               `json:"change_tone"`
}

func (m *TLMessagesComposeMessageWithAI) Encode(x *bin.Encoder, layer int32) error {
	x.PutClazzID(0xfd426afe)
	var flags uint32
	if m.Proofread {
		flags |= 1 << 0
	}
	if m.TranslateToLang != nil {
		flags |= 1 << 1
	}
	if m.ChangeTone != nil {
		flags |= 1 << 2
	}
	if m.Emojify {
		flags |= 1 << 3
	}
	x.PutUint32(flags)
	if m.Text != nil {
		_ = m.Text.Encode(x, layer)
	} else {
		_ = MakeTLTextWithEntities(&TLTextWithEntities{Text: "", Entities: []MessageEntityClazz{}}).Encode(x, layer)
	}
	if m.TranslateToLang != nil {
		x.PutString(*m.TranslateToLang)
	}
	if m.ChangeTone != nil {
		x.PutString(*m.ChangeTone)
	}
	return nil
}

func (m *TLMessagesComposeMessageWithAI) Decode(d *bin.Decoder) error {
	flags, err := d.Uint32()
	if err != nil {
		return err
	}
	if (flags & (1 << 0)) != 0 {
		m.Proofread = true
	}
	if (flags & (1 << 3)) != 0 {
		m.Emojify = true
	}
	m.Text, _ = DecodeTextWithEntitiesClazz(d)
	if (flags & (1 << 1)) != 0 {
		v, err := d.String()
		if err != nil {
			return err
		}
		m.TranslateToLang = &v
	}
	if (flags & (1 << 2)) != 0 {
		v, err := d.String()
		if err != nil {
			return err
		}
		m.ChangeTone = &v
	}
	return nil
}

type MessagesComposedMessageWithAI struct {
	Text TextWithEntitiesClazz `json:"text"`
}

func (m *MessagesComposedMessageWithAI) Encode(x *bin.Encoder, layer int32) error {
	if m.Text != nil {
		return m.Text.Encode(x, layer)
	}
	return MakeTLTextWithEntities(&TLTextWithEntities{Text: "", Entities: []MessageEntityClazz{}}).Encode(x, layer)
}

func (m *MessagesComposedMessageWithAI) Decode(d *bin.Decoder) error {
	text, err := DecodeTextWithEntitiesClazz(d)
	if err != nil {
		return err
	}
	m.Text = text
	return nil
}

func (m *MessagesComposedMessageWithAI) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

type TLMessagesReportReadMetrics struct {
	ClazzID uint32                        `json:"_id"`
	Peer    InputPeerClazz                `json:"peer"`
	Metrics []InputMessageReadMetricClazz `json:"metrics"`
}

func (m *TLMessagesReportReadMetrics) Encode(x *bin.Encoder, layer int32) error {
	x.PutClazzID(0x4067c5e6)
	if m.Peer != nil {
		_ = m.Peer.Encode(x, layer)
	}
	_ = iface.EncodeObjectList(x, m.Metrics, layer)
	return nil
}

func (m *TLMessagesReportReadMetrics) Decode(d *bin.Decoder) error {
	m.Peer, _ = DecodeInputPeerClazz(d)
	if err := d.ConsumeClazzID(iface.ClazzID_vector); err != nil {
		return err
	}
	n, err := d.Int()
	if err != nil {
		return err
	}
	m.Metrics = make([]InputMessageReadMetricClazz, n)
	for i := 0; i < n; i++ {
		msg := &TLInputMessageReadMetric{}
		msg.ClazzID, _ = d.ClazzID()
		if err = msg.Decode(d); err != nil {
			return err
		}
		m.Metrics[i] = msg
	}
	return nil
}

type TLMessagesReportMusicListen struct {
	ClazzID          uint32             `json:"_id"`
	Id               InputDocumentClazz `json:"id"`
	ListenedDuration int32              `json:"listened_duration"`
}

func (m *TLMessagesReportMusicListen) Encode(x *bin.Encoder, layer int32) error {
	x.PutClazzID(0xddbcd819)
	if m.Id != nil {
		_ = m.Id.Encode(x, layer)
	}
	x.PutInt32(m.ListenedDuration)
	return nil
}

func (m *TLMessagesReportMusicListen) Decode(d *bin.Decoder) error {
	var err error
	m.Id, _ = DecodeInputDocumentClazz(d)
	m.ListenedDuration, err = d.Int32()
	return err
}

type TLMessagesAddPollAnswer struct {
	ClazzID uint32          `json:"_id"`
	Peer    InputPeerClazz  `json:"peer"`
	MsgId   int32           `json:"msg_id"`
	Answer  PollAnswerClazz `json:"answer"`
}

func (m *TLMessagesAddPollAnswer) Encode(x *bin.Encoder, layer int32) error {
	x.PutClazzID(0x19bc4b6d)
	if m.Peer != nil {
		_ = m.Peer.Encode(x, layer)
	}
	x.PutInt32(m.MsgId)
	if m.Answer != nil {
		_ = m.Answer.Encode(x, layer)
	}
	return nil
}

func (m *TLMessagesAddPollAnswer) Decode(d *bin.Decoder) error {
	var err error
	m.Peer, _ = DecodeInputPeerClazz(d)
	m.MsgId, err = d.Int32()
	if err != nil {
		return err
	}
	m.Answer, _ = DecodePollAnswerClazz(d)
	return nil
}

type TLMessagesDeletePollAnswer struct {
	ClazzID uint32         `json:"_id"`
	Peer    InputPeerClazz `json:"peer"`
	MsgId   int32          `json:"msg_id"`
	Option  []byte         `json:"option"`
}

func (m *TLMessagesDeletePollAnswer) Encode(x *bin.Encoder, layer int32) error {
	x.PutClazzID(0xac8505a5)
	if m.Peer != nil {
		_ = m.Peer.Encode(x, layer)
	}
	x.PutInt32(m.MsgId)
	x.PutBytes(m.Option)
	return nil
}

func (m *TLMessagesDeletePollAnswer) Decode(d *bin.Decoder) error {
	var err error
	m.Peer, _ = DecodeInputPeerClazz(d)
	m.MsgId, err = d.Int32()
	if err != nil {
		return err
	}
	m.Option, err = d.Bytes()
	return err
}

type TLMessagesGetUnreadPollVotes struct {
	ClazzID   uint32         `json:"_id"`
	Peer      InputPeerClazz `json:"peer"`
	TopMsgId  *int32         `json:"top_msg_id"`
	OffsetId  int32          `json:"offset_id"`
	AddOffset int32          `json:"add_offset"`
	Limit     int32          `json:"limit"`
	MaxId     int32          `json:"max_id"`
	MinId     int32          `json:"min_id"`
}

func (m *TLMessagesGetUnreadPollVotes) Encode(x *bin.Encoder, layer int32) error {
	x.PutClazzID(0x43286cf2)
	var flags uint32
	if m.TopMsgId != nil {
		flags |= 1 << 0
	}
	x.PutUint32(flags)
	if m.Peer != nil {
		_ = m.Peer.Encode(x, layer)
	}
	if m.TopMsgId != nil {
		x.PutInt32(*m.TopMsgId)
	}
	x.PutInt32(m.OffsetId)
	x.PutInt32(m.AddOffset)
	x.PutInt32(m.Limit)
	x.PutInt32(m.MaxId)
	x.PutInt32(m.MinId)
	return nil
}

func (m *TLMessagesGetUnreadPollVotes) Decode(d *bin.Decoder) error {
	flags, err := d.Uint32()
	if err != nil {
		return err
	}
	m.Peer, _ = DecodeInputPeerClazz(d)
	if (flags & (1 << 0)) != 0 {
		v, err := d.Int32()
		if err != nil {
			return err
		}
		m.TopMsgId = &v
	}
	if m.OffsetId, err = d.Int32(); err != nil {
		return err
	}
	if m.AddOffset, err = d.Int32(); err != nil {
		return err
	}
	if m.Limit, err = d.Int32(); err != nil {
		return err
	}
	if m.MaxId, err = d.Int32(); err != nil {
		return err
	}
	m.MinId, err = d.Int32()
	return err
}

type TLMessagesReadPollVotes struct {
	ClazzID  uint32         `json:"_id"`
	Peer     InputPeerClazz `json:"peer"`
	TopMsgId *int32         `json:"top_msg_id"`
}

func (m *TLMessagesReadPollVotes) Encode(x *bin.Encoder, layer int32) error {
	x.PutClazzID(0x1720b4d8)
	var flags uint32
	if m.TopMsgId != nil {
		flags |= 1 << 0
	}
	x.PutUint32(flags)
	if m.Peer != nil {
		_ = m.Peer.Encode(x, layer)
	}
	if m.TopMsgId != nil {
		x.PutInt32(*m.TopMsgId)
	}
	return nil
}

func (m *TLMessagesReadPollVotes) Decode(d *bin.Decoder) error {
	flags, err := d.Uint32()
	if err != nil {
		return err
	}
	m.Peer, _ = DecodeInputPeerClazz(d)
	if (flags & (1 << 0)) != 0 {
		v, err := d.Int32()
		if err != nil {
			return err
		}
		m.TopMsgId = &v
	}
	return nil
}
