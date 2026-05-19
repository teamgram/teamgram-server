package serviceaction

import (
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func Encode(action tg.MessageActionClazz) (*payload.ServiceActionRefV1, error) {
	if action == nil {
		return nil, fmt.Errorf("service action encode: nil action")
	}
	body, err := iface.EncodeObject(action, payload.ServiceActionLayer)
	if err != nil {
		return nil, fmt.Errorf("service action encode: %w", err)
	}
	return &payload.ServiceActionRefV1{
		SchemaVersion: payload.ServiceActionSchemaVersionV1,
		Codec:         payload.ServiceActionCodecTLBinary,
		Layer:         payload.ServiceActionLayer,
		ActionPayload: append([]byte(nil), body...),
	}, nil
}

func Decode(ref *payload.ServiceActionRefV1) (tg.MessageActionClazz, error) {
	if ref == nil {
		return nil, fmt.Errorf("service action decode: nil ref")
	}
	if ref.SchemaVersion != payload.ServiceActionSchemaVersionV1 {
		return nil, fmt.Errorf("service action decode: unsupported schema_version=%d", ref.SchemaVersion)
	}
	if ref.Codec != payload.ServiceActionCodecTLBinary {
		return nil, fmt.Errorf("service action decode: unsupported codec=%d", ref.Codec)
	}
	if ref.Layer != payload.ServiceActionLayer {
		return nil, fmt.Errorf("service action decode: unsupported layer=%d", ref.Layer)
	}
	if len(ref.ActionPayload) == 0 {
		return nil, fmt.Errorf("service action decode: empty payload")
	}
	decoder := bin.NewDecoder(ref.ActionPayload)
	obj, err := iface.DecodeObject(decoder)
	if err != nil {
		return nil, fmt.Errorf("service action decode: %w", err)
	}
	if remaining := decoder.Remaining(); remaining != 0 {
		return nil, fmt.Errorf("service action decode: trailing bytes=%d", remaining)
	}
	action, ok := obj.(tg.MessageActionClazz)
	if !ok {
		return nil, fmt.Errorf("service action decode: object %T is not MessageAction", obj)
	}
	return action, nil
}
