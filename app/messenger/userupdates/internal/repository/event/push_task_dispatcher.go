package event

import (
	"context"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/gateway"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/projection"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
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
		logx.WithContext(ctx).Errorf("push task degraded: task_id=%d user_id=%d code=user_projection_dependency_nil", msg.TaskID, msg.UserID)
		return nil
	}
	bundle, err := d.user.UserGetUserProjectionBundle(ctx, &userpb.TLUserGetUserProjectionBundle{
		ViewerUserIds: []int64{msg.UserID},
		TargetUserIds: ids,
	})
	if err != nil {
		logx.WithContext(ctx).Errorf("push task degraded: task_id=%d user_id=%d code=user_projection_failed err=%v", msg.TaskID, msg.UserID, err)
		return nil
	}
	if bundle == nil {
		logx.WithContext(ctx).Errorf("push task degraded: task_id=%d user_id=%d code=user_projection_nil_bundle", msg.TaskID, msg.UserID)
		return nil
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
	projected, err := projection.ProjectPushTask(msg)
	if err != nil {
		return nil, nil, err
	}
	return projected.Updates, projected.AuthKeyIDExclude, nil
}
