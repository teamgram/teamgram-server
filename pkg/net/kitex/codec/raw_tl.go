package codec

import (
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
)

// RawTLObject carries a complete TL encoded object, including its constructor.
type RawTLObject struct {
	Payload []byte
}

func NewRawTLObject(payload []byte) *RawTLObject {
	return &RawTLObject{Payload: payload}
}

func (m *RawTLObject) Encode(x *bin.Encoder, _ int32) error {
	x.PutRaw(m.Payload)
	return nil
}

func (m *RawTLObject) Decode(d *bin.Decoder) error {
	payload := d.RawRemaining()
	m.Payload = append(m.Payload[:0], payload...)
	return d.Skip(len(payload))
}

func (m *RawTLObject) ConstructorID() (uint32, error) {
	return bin.NewDecoder(m.Payload).PeekClazzID()
}
