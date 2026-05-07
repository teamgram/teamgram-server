package event

import (
	"context"
	"encoding/json"
	"fmt"
	"math"

	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/gateway"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
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

type PushTaskUserProjector interface {
	UserGetUserProjectionBundle(ctx context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error)
}

type PushTaskDispatcher struct {
	authsession PushTaskAuthKeyRouter
	gateway     PushTaskGateway
	user        PushTaskUserProjector
}

func NewPushTaskDispatcher(authsession PushTaskAuthKeyRouter, gateway PushTaskGateway, user PushTaskUserProjector) *PushTaskDispatcher {
	return &PushTaskDispatcher{authsession: authsession, gateway: gateway, user: user}
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
	if err := d.projectPushUsers(ctx, msg, updates); err != nil {
		return err
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

func (d *PushTaskDispatcher) projectPushUsers(ctx context.Context, msg *payload.PushTaskKafkaMessageV1, updates tg.UpdatesClazz) error {
	full, combined, wrapper := fullUpdatesFromClazz(updates)
	if wrapper == nil {
		return nil
	}
	ids := pushProjectionTargetIDs(msg.UserID, wrapper)
	if len(ids) == 0 {
		return nil
	}
	if d.user == nil {
		return fmt.Errorf("push task user projection dependency is nil")
	}
	bundle, err := d.user.UserGetUserProjectionBundle(ctx, &userpb.TLUserGetUserProjectionBundle{
		ViewerUserIds: []int64{msg.UserID},
		TargetUserIds: ids,
	})
	if err != nil {
		return fmt.Errorf("push task project users: task_id=%d user_id=%d: %w", msg.TaskID, msg.UserID, err)
	}
	if bundle == nil {
		return fmt.Errorf("push task project users: task_id=%d user_id=%d: nil bundle", msg.TaskID, msg.UserID)
	}
	if len(bundle.MissingUserIds) > 0 {
		logx.WithContext(ctx).Errorf("push task degraded: task_id=%d user_id=%d code=missing_user_refs count=%d", msg.TaskID, msg.UserID, len(bundle.MissingUserIds))
	}
	users := []tg.UserClazz{}
	for _, viewer := range bundle.ViewerUsers {
		if viewer != nil && viewer.ViewerUserId == msg.UserID {
			users = viewer.Users
			break
		}
	}
	if full != nil {
		full.Users = users
	}
	if combined != nil {
		combined.Users = users
	}
	return nil
}

func fullUpdatesFromClazz(updates tg.UpdatesClazz) (*tg.TLUpdates, *tg.TLUpdatesCombined, *tg.Updates) {
	switch u := updates.(type) {
	case *tg.TLUpdates:
		return u, nil, u.ToUpdates()
	case *tg.TLUpdatesCombined:
		return nil, u, u.ToUpdates()
	default:
		return nil, nil, nil
	}
}

func pushProjectionTargetIDs(viewerUserID int64, updates *tg.Updates) []int64 {
	ids := tg.CollectUserIDsFromUpdates(updates)
	if len(ids) == 0 {
		return ids
	}
	for _, id := range ids {
		if id == viewerUserID {
			return ids
		}
	}
	if viewerUserID > 0 {
		ids = append(ids, viewerUserID)
	}
	return ids
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
	date, err := userupdatesDateInt32FromUnixSeconds(int64(event.Date), "push message date")
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
			Date:     date,
			ReplyTo:  replyTo,
		}), event.AuthKeyIdExclude, nil
	}
	message := tg.MakeTLMessage(&tg.TLMessage{
		Out:     event.Out,
		Id:      messageID,
		FromId:  peerFromUser(event.FromUserID),
		PeerId:  peerFromEvent(event.PeerType, event.PeerID),
		ReplyTo: replyTo,
		Date:    date,
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
		Date:  date,
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
	date, err := userupdatesDateInt32FromUnixSeconds(int64(event.Date), "read history updates date")
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
		Date:    date,
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
	date, err := userupdatesDateInt32FromUnixSeconds(int64(event.Date), "edit message date")
	if err != nil {
		return nil, nil, err
	}
	editDate32, err := userupdatesDateInt32FromUnixSeconds(int64(editDate), "edit date")
	if err != nil {
		return nil, nil, err
	}
	updateDate, err := userupdatesDateInt32FromUnixSeconds(int64(editDate)-1, "edit updates date")
	if err != nil {
		return nil, nil, err
	}
	message := tg.MakeTLMessage(&tg.TLMessage{
		Out:      event.Out,
		Id:       messageID,
		FromId:   peerFromUser(event.FromUserID),
		PeerId:   peerFromEvent(event.PeerType, event.PeerID),
		Date:     date,
		Message:  event.MessageText,
		EditDate: &editDate32,
	})
	return tg.MakeTLUpdates(&tg.TLUpdates{
		Updates: []tg.UpdateClazz{tg.MakeTLUpdateEditMessage(&tg.TLUpdateEditMessage{
			Message:  message,
			Pts:      pts,
			PtsCount: 1,
		})},
		Users: []tg.UserClazz{},
		Chats: []tg.ChatClazz{},
		Date:  updateDate,
		Seq:   0,
	}), event.AuthKeyIdExclude, nil
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

func userupdatesDateInt32FromUnixSeconds(seconds int64, field string) (int32, error) {
	date, err := tg.DateInt32FromUnixSeconds(seconds)
	if err != nil {
		return 0, fmt.Errorf("%w: convert %s: %v", userupdates.ErrUserupdatesStorage, field, err)
	}
	return date, nil
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
