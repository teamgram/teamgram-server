package core

import (
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
)

func (c *MsgCore) processUserDialogOperation(userID int64, authKeyID int64, peerType int32, peerID int64, operationID string, body []byte) (*userupdates.UserOperationResult, error) {
	if c.svcCtx.UserUpdates == nil {
		return nil, msg.ErrSenderSyncFailed
	}
	route := payload.RouteUser(userID)
	hashBytes := payload.HashBytes(body)
	result, err := c.svcCtx.UserUpdates.UserupdatesProcessUserOperation(c.ctx, &userupdates.TLUserupdatesProcessUserOperation{
		Operation: userupdates.MakeTLUserOperation(&userupdates.TLUserOperation{
			UserId:               userID,
			BucketId:             int32(route.BucketID),
			PartitionId:          int32(route.ReceiverPartitionID),
			OperationId:          operationID,
			OpType:               payload.OpTypeSendMessage,
			OpSource:             0,
			ActorUserId:          userID,
			AuthKeyId:            &authKeyID,
			AuthKeyIdExclude:     &authKeyID,
			PeerType:             peerType,
			PeerId:               peerID,
			PayloadSchemaVersion: payload.MessageOperationSchemaVersion,
			PayloadCodec:         payload.PayloadCodecJSON,
			PayloadHash:          hashBytes,
			Payload:              body,
		}),
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", msg.ErrSenderSyncFailed, err)
	}
	if result == nil {
		return nil, msg.ErrSenderSyncFailed
	}
	return result, nil
}
