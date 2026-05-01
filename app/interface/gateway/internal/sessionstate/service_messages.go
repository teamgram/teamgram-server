package sessionstate

import (
	"strings"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/mtproto"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/mt"
)

func (p *Processor) handleServiceMessage(obj iface.TLObject, msg mtproto.EncryptedMessage) ([]byte, bool, error) {
	switch request := obj.(type) {
	case *mt.TLPing:
		return encodeObject(&mt.TLPong{MsgId: msg.MsgId, PingId: request.PingId}), true, nil
	case *mt.TLPingDelayDisconnect:
		p.disconnectAt = time.Now().Add(time.Duration(request.DisconnectDelay) * time.Second)
		return encodeObject(&mt.TLPong{MsgId: msg.MsgId, PingId: request.PingId}), true, nil
	case *mt.TLMsgsAck:
		return nil, true, nil
	case *mt.TLMsgsStateReq:
		return encodeObject(&mt.TLMsgsStateInfo{
			ReqMsgId: msg.MsgId,
			Info:     strings.Repeat("\x01", len(request.MsgIds)),
		}), true, nil
	default:
		return nil, false, nil
	}
}
