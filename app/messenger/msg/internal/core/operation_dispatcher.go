package core

import (
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
)

type DeliveryPolicy int32

const (
	DeliveryPolicyRequesterSync      DeliveryPolicy = 1
	DeliveryPolicyBrokerDurableAck   DeliveryPolicy = 2
	DeliveryPolicyDurableAsync       DeliveryPolicy = 3
	DeliveryPolicyApplyAck           DeliveryPolicy = 4
	DeliveryPolicyBestEffortPushOnly DeliveryPolicy = 5
)

type OperationEnvelope struct {
	UserID               int64
	OperationID          string
	OpType               int32
	OperationKind        string
	ActorUserID          int64
	AuthKeyID            *int64
	AuthKeyIDExclude     *int64
	PeerType             int32
	PeerID               int64
	CanonicalMessageID   *int64
	CanonicalPeerSeq     *int64
	CanonicalDate        *int64
	PayloadSchemaVersion int32
	PayloadCodec         int32
	PayloadHash          []byte
	Payload              []byte
	DependencyPts        []int64
	DeliveryPolicy       DeliveryPolicy
}

func (e OperationEnvelope) toTLUserOperation() *userupdates.TLUserOperation {
	route := payload.RouteUser(e.UserID)
	return userupdates.MakeTLUserOperation(&userupdates.TLUserOperation{
		UserId:               e.UserID,
		BucketId:             int32(route.BucketID),
		PartitionId:          int32(route.ReceiverPartitionID),
		OperationId:          e.OperationID,
		OpType:               e.OpType,
		OpSource:             0,
		ActorUserId:          e.ActorUserID,
		AuthKeyId:            e.AuthKeyID,
		AuthKeyIdExclude:     e.AuthKeyIDExclude,
		PeerType:             e.PeerType,
		PeerId:               e.PeerID,
		CanonicalMessageId:   e.CanonicalMessageID,
		CanonicalPeerSeq:     e.CanonicalPeerSeq,
		CanonicalDate:        e.CanonicalDate,
		PayloadSchemaVersion: e.PayloadSchemaVersion,
		PayloadCodec:         e.PayloadCodec,
		PayloadHash:          e.PayloadHash,
		Payload:              e.Payload,
		DependencyPts:        firstDependencyPts(e.DependencyPts),
	})
}

func (e OperationEnvelope) toReceiverOperation() repository.ReceiverOperation {
	route := payload.RouteUser(e.UserID)
	return repository.ReceiverOperation{
		UserID:        e.UserID,
		BucketID:      int32(route.BucketID),
		PartitionID:   int32(route.ReceiverPartitionID),
		OperationID:   e.OperationID,
		OpType:        e.OpType,
		PeerType:      e.PeerType,
		PeerID:        e.PeerID,
		PayloadCodec:  e.PayloadCodec,
		Payload:       e.Payload,
		PayloadHash:   e.PayloadHash,
		DependencyPts: append([]int64(nil), e.DependencyPts...),
	}
}

func (e OperationEnvelope) toAffectedUserOperation(requesterUserID int64) *userupdates.TLAffectedUserOperation {
	return userupdates.MakeTLAffectedUserOperation(&userupdates.TLAffectedUserOperation{
		RequesterUserId: requesterUserID,
		DeliveryPolicy:  int32(e.DeliveryPolicy),
		OperationKind:   e.OperationKind,
		Operation:       e.toTLUserOperation(),
	})
}

func (c *MsgCore) dispatchRequesterSync(requester OperationEnvelope, effects []OperationEnvelope) (*userupdates.UserOperationResult, error) {
	if requester.DeliveryPolicy != DeliveryPolicyRequesterSync {
		return nil, fmt.Errorf("%w: unsupported requester delivery policy=%d", msg.ErrSenderSyncFailed, requester.DeliveryPolicy)
	}
	if c == nil || c.svcCtx == nil || c.svcCtx.UserUpdates == nil {
		return nil, msg.ErrSenderSyncFailed
	}

	affected := make([]userupdates.AffectedUserOperationClazz, 0, len(effects))
	for _, effect := range effects {
		if effect.UserID == requester.UserID {
			continue
		}
		if effect.DeliveryPolicy != DeliveryPolicyDurableAsync {
			return nil, fmt.Errorf("%w: unsupported affected delivery policy=%d", msg.ErrSenderSyncFailed, effect.DeliveryPolicy)
		}
		affected = append(affected, effect.toAffectedUserOperation(requester.UserID))
	}

	if len(affected) == 0 {
		result, err := c.svcCtx.UserUpdates.UserupdatesProcessUserOperation(c.ctx, &userupdates.TLUserupdatesProcessUserOperation{
			Operation: requester.toTLUserOperation(),
		})
		if err != nil {
			return nil, fmt.Errorf("%w: %w", msg.ErrSenderSyncFailed, err)
		}
		if result == nil {
			return nil, msg.ErrSenderSyncFailed
		}
		return result, nil
	}

	result, err := c.svcCtx.UserUpdates.UserupdatesProcessUserOperationWithEffects(c.ctx, &userupdates.TLUserupdatesProcessUserOperationWithEffects{
		Operation:       requester.toTLUserOperation(),
		AffectedEffects: affected,
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %w", msg.ErrSenderSyncFailed, err)
	}
	if result == nil {
		return nil, msg.ErrSenderSyncFailed
	}
	return result, nil
}

func (c *MsgCore) dispatchRequesterBatchSync(requesters []OperationEnvelope) ([]*userupdates.UserOperationResult, error) {
	if len(requesters) == 0 {
		return []*userupdates.UserOperationResult{}, nil
	}
	if c == nil || c.svcCtx == nil || c.svcCtx.UserUpdates == nil {
		return nil, msg.ErrSenderSyncFailed
	}
	operations := make([]userupdates.UserOperationClazz, 0, len(requesters))
	for _, requester := range requesters {
		if requester.DeliveryPolicy != DeliveryPolicyRequesterSync {
			return nil, fmt.Errorf("%w: unsupported requester delivery policy=%d", msg.ErrSenderSyncFailed, requester.DeliveryPolicy)
		}
		operations = append(operations, requester.toTLUserOperation())
	}
	vector, err := c.svcCtx.UserUpdates.UserupdatesProcessUserOperationBatch(c.ctx, &userupdates.TLUserupdatesProcessUserOperationBatch{
		Operations: operations,
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %w", msg.ErrSenderSyncFailed, err)
	}
	if vector == nil || len(vector.Datas) != len(requesters) {
		return nil, msg.ErrSenderSyncFailed
	}
	out := make([]*userupdates.UserOperationResult, 0, len(vector.Datas))
	for _, item := range vector.Datas {
		result := item
		if result == nil {
			return nil, msg.ErrSenderSyncFailed
		}
		out = append(out, result)
	}
	return out, nil
}

func (c *MsgCore) dispatchBrokerDurableAck(effect OperationEnvelope) (repository.KafkaAck, error) {
	if effect.DeliveryPolicy != DeliveryPolicyBrokerDurableAck {
		return repository.KafkaAck{}, fmt.Errorf("%w: unsupported receiver delivery policy=%d", msg.ErrReceiverBackpressure, effect.DeliveryPolicy)
	}
	if c == nil || c.svcCtx == nil || c.svcCtx.ReceiverPublisher == nil {
		return repository.KafkaAck{}, msg.ErrReceiverBackpressure
	}
	ack, err := c.svcCtx.ReceiverPublisher.Publish(c.ctx, effect.toReceiverOperation())
	if err != nil {
		return repository.KafkaAck{}, fmt.Errorf("%w: %w", msg.ErrReceiverBackpressure, err)
	}
	return ack, nil
}

func firstDependencyPts(values []int64) *int64 {
	if len(values) == 0 {
		return nil
	}
	return &values[0]
}
