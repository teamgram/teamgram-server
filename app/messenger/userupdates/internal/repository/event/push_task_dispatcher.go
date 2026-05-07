package event

import (
	"context"
	"encoding/json"
	"fmt"
	"math"

	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/gateway"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
	"github.com/zeromicro/go-zero/core/logx"
)

type PushTaskKafkaRecord struct {
	Topic     string
	Partition int32
	Offset    int64
	Key       []byte
	Value     []byte
}

type PushTaskAuthKeyRouter interface {
	AuthsessionGetPermAuthKeyIds(ctx context.Context, in *authsession.TLAuthsessionGetPermAuthKeyIds) (*authsession.VectorLong, error)
}

type PushTaskGateway interface {
	GatewayPushUpdatesData(ctx context.Context, in *gateway.TLGatewayPushUpdatesData) (*tg.Bool, error)
}

type PushTaskDispatcher struct {
	authsession PushTaskAuthKeyRouter
	gateway     PushTaskGateway
}

func NewPushTaskDispatcher(authsession PushTaskAuthKeyRouter, gateway PushTaskGateway) *PushTaskDispatcher {
	return &PushTaskDispatcher{authsession: authsession, gateway: gateway}
}

func (d *PushTaskDispatcher) HandlePushTaskKafkaRecord(ctx context.Context, record PushTaskKafkaRecord) error {
	msg, err := payload.UnmarshalPushTaskKafkaMessage(record.Value)
	if err != nil {
		logx.WithContext(ctx).Errorf("push task terminal: topic=%s partition=%d offset=%d code=payload_decode_failed err=%v", record.Topic, record.Partition, record.Offset, err)
		return nil
	}
	if msg.PushType != 1 {
		logx.WithContext(ctx).Errorf("push task terminal: task_id=%d user_id=%d code=unsupported_push_type push_type=%d", msg.TaskID, msg.UserID, msg.PushType)
		return nil
	}
	updates, authKeyIDExclude, err := pushTaskUpdates(msg)
	if err != nil {
		logx.WithContext(ctx).Errorf("push task terminal: task_id=%d user_id=%d code=payload_projection_failed err=%v", msg.TaskID, msg.UserID, err)
		return nil
	}
	if d.authsession == nil || d.gateway == nil {
		return fmt.Errorf("push task dispatcher dependencies are nil")
	}
	keys, err := d.authsession.AuthsessionGetPermAuthKeyIds(ctx, &authsession.TLAuthsessionGetPermAuthKeyIds{UserId: msg.UserID})
	if err != nil {
		return fmt.Errorf("push task route auth keys: task_id=%d user_id=%d: %w", msg.TaskID, msg.UserID, err)
	}
	if keys == nil || len(keys.Datas) == 0 {
		return nil
	}
	for _, permAuthKeyId := range keys.Datas {
		if authKeyIDExclude != nil && *authKeyIDExclude == permAuthKeyId {
			continue
		}
		if _, err := d.gateway.GatewayPushUpdatesData(ctx, &gateway.TLGatewayPushUpdatesData{
			PermAuthKeyId: permAuthKeyId,
			Updates:       updates,
		}); err != nil {
			return fmt.Errorf("push task gateway push: task_id=%d user_id=%d perm_auth_key_id=%d: %w", msg.TaskID, msg.UserID, permAuthKeyId, err)
		}
	}
	return nil
}

func pushTaskUpdates(msg *payload.PushTaskKafkaMessageV1) (tg.UpdatesClazz, *int64, error) {
	var event payload.MessageEventV1
	if err := json.Unmarshal(msg.Payload, &event); err != nil {
		return nil, nil, fmt.Errorf("decode message event: %w", err)
	}
	if event.SchemaVersion != payload.MessageEventSchemaVersion {
		return nil, nil, fmt.Errorf("unsupported message event schema=%d kind=%s", event.SchemaVersion, event.EventKind)
	}
	if event.EventKind == payload.OperationKindReadHistory {
		return readHistoryUpdates(msg, event)
	}
	if event.EventKind == payload.OperationKindEditMessage {
		return editMessageUpdates(msg, event)
	}
	if event.EventKind != payload.EventKindNewMessage {
		return nil, nil, fmt.Errorf("unsupported message event schema=%d kind=%s", event.SchemaVersion, event.EventKind)
	}
	messageID, err := int64ToInt32(event.MessageID, "message id")
	if err != nil {
		return nil, nil, err
	}
	pts, err := int64ToInt32(msg.Pts, "pts")
	if err != nil {
		return nil, nil, err
	}
	replyTo, err := replyHeaderFromPeerSeq(event.ReplyToPeerSeq)
	if err != nil {
		return nil, nil, err
	}
	if event.PeerType == payload.PeerTypeUser {
		return tg.MakeTLUpdateShortMessage(&tg.TLUpdateShortMessage{
			Out:      event.Out,
			Id:       messageID,
			UserId:   shortMessageUserID(event),
			Message:  event.MessageText,
			Pts:      pts,
			PtsCount: 1,
			Date:     event.Date,
			ReplyTo:  replyTo,
		}), event.AuthKeyIdExclude, nil
	}
	message := tg.MakeTLMessage(&tg.TLMessage{
		Out:     event.Out,
		Id:      messageID,
		FromId:  peerFromUser(event.FromUserID),
		PeerId:  peerFromEvent(event.PeerType, event.PeerID),
		ReplyTo: replyTo,
		Date:    event.Date,
		Message: event.MessageText,
	})
	return tg.MakeTLUpdates(&tg.TLUpdates{
		Updates: []tg.UpdateClazz{tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
			Message:  message,
			Pts:      pts,
			PtsCount: 1,
		})},
		Users: []tg.UserClazz{},
		Chats: []tg.ChatClazz{},
		Date:  event.Date,
		Seq:   pts,
	}), event.AuthKeyIdExclude, nil
}

func replyHeaderFromPeerSeq(peerSeq int64) (tg.MessageReplyHeaderClazz, error) {
	if peerSeq <= 0 {
		return nil, nil
	}
	replyToMsgID, err := int64ToInt32(peerSeq, "reply peer seq")
	if err != nil {
		return nil, err
	}
	return tg.MakeTLMessageReplyHeader(&tg.TLMessageReplyHeader{ReplyToMsgId: &replyToMsgID}), nil
}

func readHistoryUpdates(msg *payload.PushTaskKafkaMessageV1, event payload.MessageEventV1) (tg.UpdatesClazz, *int64, error) {
	maxID, err := int64ToInt32(event.MessageID, "message id")
	if err != nil {
		return nil, nil, err
	}
	pts, err := int64ToInt32(msg.Pts, "pts")
	if err != nil {
		return nil, nil, err
	}
	peer := peerFromEvent(event.PeerType, event.PeerID)
	var update tg.UpdateClazz
	if event.Out {
		update = tg.MakeTLUpdateReadHistoryOutbox(&tg.TLUpdateReadHistoryOutbox{
			Peer:     peer,
			MaxId:    maxID,
			Pts:      pts,
			PtsCount: 1,
		})
	} else {
		update = tg.MakeTLUpdateReadHistoryInbox(&tg.TLUpdateReadHistoryInbox{
			Peer:             peer,
			MaxId:            maxID,
			StillUnreadCount: 0,
			Pts:              pts,
			PtsCount:         1,
		})
	}
	return tg.MakeTLUpdates(&tg.TLUpdates{
		Updates: []tg.UpdateClazz{update},
		Users:   []tg.UserClazz{},
		Chats:   []tg.ChatClazz{},
		Date:    event.Date,
		Seq:     pts,
	}), event.AuthKeyIdExclude, nil
}

func editMessageUpdates(msg *payload.PushTaskKafkaMessageV1, event payload.MessageEventV1) (tg.UpdatesClazz, *int64, error) {
	messageID, err := int64ToInt32(event.MessageID, "message id")
	if err != nil {
		return nil, nil, err
	}
	pts, err := int64ToInt32(msg.Pts, "pts")
	if err != nil {
		return nil, nil, err
	}
	editDate := event.EditDate
	if editDate == 0 {
		editDate = event.Date
	}
	message := tg.MakeTLMessage(&tg.TLMessage{
		Out:      event.Out,
		Id:       messageID,
		FromId:   peerFromUser(event.FromUserID),
		PeerId:   peerFromEvent(event.PeerType, event.PeerID),
		Date:     event.Date,
		Message:  event.MessageText,
		EditDate: &editDate,
	})
	return tg.MakeTLUpdates(&tg.TLUpdates{
		Updates: []tg.UpdateClazz{tg.MakeTLUpdateEditMessage(&tg.TLUpdateEditMessage{
			Message:  message,
			Pts:      pts,
			PtsCount: 1,
		})},
		Users: editMessageUsers(event.FromUserID, event.ToUserID),
		Chats: []tg.ChatClazz{},
		Date:  editDate - 1,
		Seq:   0,
	}), event.AuthKeyIdExclude, nil
}

func editMessageUsers(fromUserID, toUserID int64) []tg.UserClazz {
	if fromUserID == toUserID {
		return []tg.UserClazz{tg.MakeTLUser(&tg.TLUser{Id: fromUserID})}
	}
	return []tg.UserClazz{
		tg.MakeTLUser(&tg.TLUser{Id: fromUserID}),
		tg.MakeTLUser(&tg.TLUser{Id: toUserID}),
	}
}

func shortMessageUserID(event payload.MessageEventV1) int64 {
	if event.Out {
		return event.PeerID
	}
	return event.FromUserID
}

func int64ToInt32(v int64, field string) (int32, error) {
	if v < math.MinInt32 || v > math.MaxInt32 {
		return 0, fmt.Errorf("%s out of int32 range", field)
	}
	return int32(v), nil
}

func peerFromUser(userID int64) tg.PeerClazz {
	if userID == 0 {
		return nil
	}
	return tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: userID})
}

func peerFromEvent(peerType int32, peerID int64) tg.PeerClazz {
	switch peerType {
	case payload.PeerTypeUser:
		return tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: peerID})
	case payload.PeerTypeChat:
		return tg.MakeTLPeerChat(&tg.TLPeerChat{ChatId: peerID})
	case payload.PeerTypeChannel:
		return tg.MakeTLPeerChannel(&tg.TLPeerChannel{ChannelId: peerID})
	default:
		return tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: peerID})
	}
}
