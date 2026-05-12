package core

import (
	"encoding/json"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
)

func (c *MsgCore) processUserDialogOperation(userID int64, authKeyID int64, peerType int32, peerID int64, operationID string, body []byte, effects []OperationEnvelope) (*userupdates.UserOperationResult, error) {
	hashBytes := payload.HashBytes(body)
	var op payload.MessageOperationV1
	if err := json.Unmarshal(body, &op); err != nil {
		return nil, fmt.Errorf("%w: decode dialog operation kind: %v", msg.ErrMsgStorage, err)
	}
	return c.dispatchRequesterSync(OperationEnvelope{
		UserID:               userID,
		OperationID:          operationID,
		OpType:               payload.OpTypeSendMessage,
		OperationKind:        op.OperationKind,
		ActorUserID:          userID,
		AuthKeyID:            &authKeyID,
		AuthKeyIDExclude:     &authKeyID,
		PeerType:             peerType,
		PeerID:               peerID,
		PayloadSchemaVersion: payload.MessageOperationSchemaVersion,
		PayloadCodec:         payload.PayloadCodecJSON,
		PayloadHash:          hashBytes,
		Payload:              body,
		DeliveryPolicy:       DeliveryPolicyRequesterSync,
	}, effects)
}
