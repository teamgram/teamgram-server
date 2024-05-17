/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package core

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/proto/mtproto/rpc/metadata"
	"github.com/teamgram/teamgram-server/app/interface/session/session"
	"github.com/teamgram/teamgram-server/app/messenger/sync/internal/svc"
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"
	"github.com/teamgram/teamgram-server/app/service/status/status"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type SyncType int

const (
	syncTypeUser      SyncType = 1 // 该用户所有设备
	syncTypeUserNotMe SyncType = 2 // 该用户除了某个设备
	syncTypeUserMe    SyncType = 3 // 该用户指定某个设备
)

type SyncCore struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	MD *metadata.RpcMetadata
}

func New(ctx context.Context, svcCtx *svc.ServiceContext) *SyncCore {
	return &SyncCore{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		MD:     metadata.RpcMetadataFromIncoming(ctx),
	}
}

func (c *SyncCore) processUpdates(syncType SyncType, userId int64, isBot bool, ups *mtproto.Updates) (needPush bool, err error) {
	mtproto.VisitUpdates(userId, ups, map[string]mtproto.UpdateVisitedFunc{
		mtproto.Predicate_updateNewMessage: func(
			userId int64,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32,
		) {
			c.svcCtx.Dao.AddToPtsQueue(c.ctx, userId, update.Pts_INT32, update.PtsCount, update)
			needPush = true
		},
		mtproto.Predicate_updateDeleteMessages: func(
			userId int64,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32,
		) {
			c.svcCtx.Dao.AddToPtsQueue(c.ctx, userId, update.Pts_INT32, update.PtsCount, update)
			needPush = true
		},
		mtproto.Predicate_updateReadHistoryInbox: func(
			userId int64,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32,
		) {
			c.svcCtx.Dao.AddToPtsQueue(c.ctx, userId, update.Pts_INT32, update.PtsCount, update)
			needPush = true
		},
		mtproto.Predicate_updateReadHistoryOutbox: func(
			userId int64,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32,
		) {
			c.svcCtx.Dao.AddToPtsQueue(c.ctx, userId, update.Pts_INT32, update.PtsCount, update)
			needPush = true
		},
		mtproto.Predicate_updateWebPage: func(
			userId int64,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32,
		) {
			c.svcCtx.Dao.AddToPtsQueue(c.ctx, userId, update.Pts_INT32, update.PtsCount, update)
			needPush = true
		},
		mtproto.Predicate_updateReadMessagesContents: func(
			userId int64,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32,
		) {
			c.svcCtx.Dao.AddToPtsQueue(c.ctx, userId, update.Pts_INT32, update.PtsCount, update)
			needPush = true
		},
		mtproto.Predicate_updateEditMessage: func(
			userId int64,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32,
		) {
			c.svcCtx.Dao.AddToPtsQueue(c.ctx, userId, update.Pts_INT32, update.PtsCount, update)
			needPush = true
		},
		mtproto.Predicate_updateFolderPeers: func(
			userId int64,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32,
		) {
			if syncType == syncTypeUserNotMe {
				c.svcCtx.Dao.AddToPtsQueue(c.ctx, userId, update.Pts_INT32, update.PtsCount, update)
			}
		},
		mtproto.Predicate_updatePinnedMessages: func(
			userId int64,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32,
		) {
			if syncType == syncTypeUserNotMe {
				c.svcCtx.Dao.AddToPtsQueue(c.ctx, userId, update.Pts_INT32, update.PtsCount, update)
			}
		},
		mtproto.Predicate_updatePhoneCall: func(
			userId int64,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32,
		) {
			if update.GetPhoneCall().GetPredicateName() == mtproto.Predicate_phoneCallRequested {
				// log.Debugf("recv phoneCallRequested")
				needPush = true
			}
		},
		mtproto.Predicate_updatePeerSettings: func(
			userId int64,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32,
		) {
			needPush = true
		},
	})

	return needPush, nil
}

func (c *SyncCore) pushUpdatesToSession(syncType SyncType, userId, permAuthKeyId int64, hasServerId *wrapperspb.StringValue, authKeyId, sessionId *wrapperspb.Int64Value, pushData *mtproto.Updates, notification bool) {
	if syncType == syncTypeUserMe && hasServerId != nil {
		c.Logger.Debugf("pushUpdatesToSession - pushData: {server_id: %v, auth_key_id: %v}", hasServerId, authKeyId)
		if sessionId != nil {
			c.svcCtx.Dao.PushSessionUpdatesToSession(
				c.ctx,
				hasServerId.GetValue(),
				&session.TLSessionPushSessionUpdatesData{
					PermAuthKeyId: permAuthKeyId,
					AuthKeyId:     authKeyId.GetValue(),
					SessionId:     sessionId.GetValue(),
					Updates:       pushData,
				})
		} else {
			c.svcCtx.Dao.PushUpdatesToSession(
				c.ctx,
				hasServerId.GetValue(),
				&session.TLSessionPushUpdatesData{
					PermAuthKeyId: permAuthKeyId,
					Notification:  notification,
					Updates:       pushData,
				})
		}
	} else {
		var (
			pushExcludeList   = make([]int64, 0)
			serverIdKeyIdList = make(map[string][]int64)
		)

		statusList, _ := c.svcCtx.Dao.StatusClient.StatusGetUserOnlineSessions(c.ctx, &status.TLStatusGetUserOnlineSessions{
			UserId: userId,
		})
		c.Logger.Debugf("statusList - #%v", statusList)
		for _, sess := range statusList.GetUserSessions() {
			if syncType == syncTypeUserNotMe && sess.AuthKeyId == permAuthKeyId {
				continue
			}
			pushExcludeList = append(pushExcludeList, sess.PermAuthKeyId)
			if keyIdList, ok := serverIdKeyIdList[sess.Gateway]; ok {
				keyIdList = append(keyIdList, sess.AuthKeyId)
				serverIdKeyIdList[sess.Gateway] = keyIdList
			} else {
				serverIdKeyIdList[sess.Gateway] = []int64{sess.AuthKeyId}
			}
		}

		c.Logger.Debugf("serverIdKeyIdList - #%v", serverIdKeyIdList)
		for serverId, keyIdList := range serverIdKeyIdList {
			for _, keyId := range keyIdList {
				// log.Debugf("serverIdKeyIdList - #%v", serverIdKeyIdList)
				c.svcCtx.Dao.PushUpdatesToSession(
					c.ctx,
					serverId,
					&session.TLSessionPushUpdatesData{
						PermAuthKeyId: keyId,
						Notification:  notification,
						Updates:       pushData,
					})
			}
		}

		if syncType == syncTypeUser {
			if c.svcCtx.Dao.PushClient != nil {
				c.Logger.Debugf("push PushClient...")
				c.svcCtx.Dao.PushClient.SyncPushUpdatesIfNot(c.ctx, &sync.TLSyncPushUpdatesIfNot{
					UserId:   userId,
					Excludes: pushExcludeList,
					Updates:  pushData,
				})
			}
		}
	}
}
