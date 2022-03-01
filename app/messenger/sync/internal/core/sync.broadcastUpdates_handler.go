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
	// "github.com/teamgram/marmota/pkg/container2"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"
	// channelpb "github.com/teamgram/teamgram-server/app/service/biz/channel/channel"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
)

// SyncBroadcastUpdates
// sync.broadcastUpdates broadcast_type:int chat_id:long exclude_id_list:Vector<long> updates:Updates = Void;
func (c *SyncCore) SyncBroadcastUpdates(in *sync.TLSyncBroadcastUpdates) (*mtproto.Void, error) {
	pushUpdates := &sync.TLSyncPushUpdates{
		UserId:  0,
		Updates: in.Updates,
	}

	switch in.BroadcastType {
	case sync.BroadcastTypeChat:
		idList, _ := c.svcCtx.Dao.ChatClient.ChatGetChatParticipantIdList(c.ctx, &chatpb.TLChatGetChatParticipantIdList{
			ChatId: in.ChatId,
		})

		for _, id := range idList.GetDatas() {
			pushUpdates.UserId = id
			c.SyncPushUpdates(pushUpdates)
		}
	default:
		//ch, _ := c.svcCtx.Dao.ChannelClient.ChannelGetMutableChannelByPush(c.ctx, &channelpb.TLChannelGetMutableChannelByPush{
		//	ChannelId: in.ChatId,
		//	OffsetId:  0,
		//	Limit:     0,
		//})
		//
		//ch.WalkByPush(func(isAdmin bool, pushUserId int64) {
		//	if ok, _ := container2.Contains(pushUserId, in.ExcludeIdList); ok {
		//		return
		//	}
		//	switch in.BroadcastType {
		//	case sync.BroadcastTypeChannel:
		//	case sync.BroadcastTypeChannelAdmin:
		//		if !isAdmin {
		//			return
		//		}
		//	case sync.BroadcastTypeChannelNotAdmin:
		//		if isAdmin {
		//			return
		//		}
		//	default:
		//		return
		//	}
		//
		//	pushUpdates.UserId = pushUserId
		//	c.SyncPushUpdates(pushUpdates)
		//})
	}

	return mtproto.EmptyVoid, nil
}
