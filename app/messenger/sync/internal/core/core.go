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
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/proto/mtproto/rpc/metadata"
	"github.com/teamgram/teamgram-server/app/interface/session/session"
	"github.com/teamgram/teamgram-server/app/messenger/sync/internal/svc"
	"github.com/teamgram/teamgram-server/app/service/status/status"

	"github.com/zeromicro/go-zero/core/logx"
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

	return false, nil
}

func (c *SyncCore) pushUpdatesToSession(syncType SyncType, userId, authKeyId, clientMsgId int64, pushData *mtproto.Updates, hasServerId string, notification bool) {
	if syncType == syncTypeUserMe && hasServerId != "" {
		logx.Infof("pushUpdatesToSession - pushData: {server_id: %d, auth_key_id: %d}", hasServerId, authKeyId)
		if clientMsgId != 0 {
			c.svcCtx.Dao.PushSessionUpdatesToSession(
				c.ctx,
				hasServerId,
				&session.TLSessionPushSessionUpdatesData{
					AuthKeyId: authKeyId,
					SessionId: clientMsgId,
					Updates:   pushData,
				})
		} else {
			c.svcCtx.Dao.PushUpdatesToSession(
				c.ctx,
				hasServerId,
				&session.TLSessionPushUpdatesData{
					AuthKeyId:    authKeyId,
					Notification: notification,
					Updates:      pushData,
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
		logx.Infof("statusList - #%v", statusList)
		for _, sess := range statusList.GetUserSessions() {
			if syncType == syncTypeUserNotMe && sess.AuthKeyId == authKeyId {
				continue
			}
			pushExcludeList = append(pushExcludeList, sess.AuthKeyId)
			if keyIdList, ok := serverIdKeyIdList[sess.Gateway]; ok {
				keyIdList = append(keyIdList, sess.AuthKeyId)
				serverIdKeyIdList[sess.Gateway] = keyIdList
			} else {
				serverIdKeyIdList[sess.Gateway] = []int64{sess.AuthKeyId}
			}
		}

		logx.Infof("serverIdKeyIdList - #%v", serverIdKeyIdList)
		for serverId, keyIdList := range serverIdKeyIdList {
			for _, keyId := range keyIdList {
				// log.Debugf("serverIdKeyIdList - #%v", serverIdKeyIdList)
				c.svcCtx.Dao.PushUpdatesToSession(
					c.ctx,
					serverId,
					&session.TLSessionPushUpdatesData{
						AuthKeyId:    keyId,
						Notification: notification,
						Updates:      pushData,
					})
			}
		}

		if syncType == syncTypeUser {
			if c.svcCtx.Dao.PushClient != nil {
				c.Logger.Infof("push PushClient...")
				c.svcCtx.Dao.PushClient.SyncPushUpdatesIfNot(c.ctx, &sync.TLSyncPushUpdatesIfNot{
					UserId:   userId,
					Excludes: pushExcludeList,
					Updates:  pushData,
				})
			}
		}
	}
}
