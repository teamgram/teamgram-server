// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/msg/inbox/inbox"
	"github.com/teamgram/teamgram-server/app/messenger/msg/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

// InboxSendUserMessageToInboxV2
// inbox.sendUserMessageToInboxV2 flags:# user_id:long out:flags.0?true from_id:long peer_user_id:long inbox:MessageBox users:flags.1?Vector<ImmutableUser> = Void;
func (c *InboxCore) InboxSendUserMessageToInboxV2(in *inbox.TLInboxSendUserMessageToInboxV2) (*mtproto.Void, error) {
	if in.Out {
		isUseV3 := false
		if in.GetLayer() != nil {
			isUseV3 = true
		}
		if !isUseV3 && in.GetServerId() != nil {
			isUseV3 = true
		}
		if !isUseV3 && in.GetSessionId() != nil {
			isUseV3 = true
		}
		if !isUseV3 && in.GetClientReqMsgId() != nil {
			isUseV3 = true
		}

		//boxList := in.GetBoxList()
		for _, inBox := range in.GetBoxList() {
			if isUseV3 {
				inBox.Pts = c.svcCtx.Dao.IDGenClient2.NextPtsId(c.ctx, in.FromId)
				inBox.PtsCount = 1
			}

			err := c.svcCtx.Dao.SendMessageToOutboxV1(
				c.ctx,
				in.FromId,
				mtproto.MakePeerUtil(in.PeerType, in.PeerId),
				inBox)
			if err != nil {
				// TODO: handle error
				c.Logger.Errorf("inbox.sendUserMessageToInboxV2 - error: %v", err)
				return nil, err
			}

			// TODO: handle sendToSelfUser
			if in.PeerType == mtproto.PEER_USER && in.FromId == in.PeerId {
				peer2 := inBox.GetMessage().GetSavedPeerId()
				if peer2 == nil {
					c.Logger.Errorf("inbox.sendUserMessageToInboxV2 - error: sendToSelfUser")
				} else {
					peer := mtproto.FromPeer(peer2)
					_, _, _ = c.svcCtx.Dao.SavedDialogsDAO.InsertOrUpdate(
						c.ctx,
						&dataobject.SavedDialogsDO{
							UserId:     in.FromId,
							PeerType:   peer.PeerType,
							PeerId:     peer.PeerId,
							Pinned:     0,
							TopMessage: inBox.GetMessageId(),
						})
				}
			}

			_, err = c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(c.ctx, &sync.TLSyncUpdatesNotMe{
				UserId:        inBox.UserId,
				PermAuthKeyId: in.FromAuthKeyId,
				Updates: mtproto.MakeUpdatesByUpdatesUsersChats(
					in.Users,
					in.Chats,
					mtproto.MakeTLUpdateNewMessage(&mtproto.Update{
						Message_MESSAGE: inBox.GetMessage(),
						Pts_INT32:       inBox.Pts,
						PtsCount:        inBox.PtsCount,
					}).To_Update()),
			})
		}

		if isUseV3 {
			if len(in.GetBoxList()) == 1 {
				box := in.BoxList[0]

				rpcResult := &mtproto.TLRpcResult{
					ReqMsgId: in.GetClientReqMsgId().GetValue(),
					Result: mtproto.MakeReplyUpdates(
						func(idList []int64) []*mtproto.User {
							// TODO: check
							//users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(ctx,
							//	&userpb.TLUserGetMutableUsers{
							//		Id: idList,
							//	})
							return in.Users
						},
						func(idList []int64) []*mtproto.Chat {
							return []*mtproto.Chat{}
						},
						func(idList []int64) []*mtproto.Chat {
							// TODO
							return []*mtproto.Chat{}
						},
						mtproto.MakeTLUpdateNewMessage(&mtproto.Update{
							Pts_INT32:       box.Pts,
							PtsCount:        box.PtsCount,
							RandomId:        box.RandomId,
							Message_MESSAGE: box.Message,
						}).To_Update()),
				}
				// push
				x := mtproto.NewEncodeBuf(512)
				_ = rpcResult.Encode(x, in.GetLayer().GetValue())
				_, _ = c.svcCtx.Dao.SyncClient.SyncPushRpcResult(c.ctx, &sync.TLSyncPushRpcResult{
					UserId:         box.UserId,
					AuthKeyId:      in.GetAuthKeyId().GetValue(),
					PermAuthKeyId:  in.GetAuthKeyId().GetValue(),
					ServerId:       in.GetServerId().GetValue(),
					SessionId:      in.GetSessionId().GetValue(),
					ClientReqMsgId: in.GetClientReqMsgId().GetValue(),
					RpcResult:      x.GetBuf(),
				})
			}
		}
	} else {
		for _, inbox2 := range in.GetBoxList() {
			var (
				inBox *mtproto.MessageBox
				err   error
			)

			switch in.PeerType {
			case mtproto.PEER_USER:
				inBox, err = c.svcCtx.Dao.SendUserMessageToInbox(c.ctx,
					in.FromId,
					in.PeerId,
					inbox2.GetDialogMessageId(),
					inbox2.GetRandomId(),
					inbox2.GetMessage())
				if err != nil {
					c.Logger.Errorf("inbox.sendUserMessageToInboxV2 - error: %v", err)
					return nil, err
				}
			case mtproto.PEER_CHAT:
				inBox, err = c.svcCtx.Dao.SendChatMessageToInbox(c.ctx,
					in.FromId,
					in.PeerId,
					in.UserId,
					inbox2.GetDialogMessageId(),
					inbox2.GetRandomId(),
					inbox2.GetMessage())
				if err != nil {
					c.Logger.Errorf("inbox.sendUserMessageToInboxV2 - error: %v", err)
					return nil, err
				}
			default:
				c.Logger.Errorf("inbox.sendUserMessageToInboxV2 - error: invalid peerType")
				return mtproto.EmptyVoid, nil

			}

			if inBox.DialogMessageId == 1 &&
				(in.FromId != 42777 && in.FromId != 424000) {
				//isContact, _ := s.UserFacade.GetContactAndMutual(ctx, toId, fromId)
				//if !isContact {
				//	s.UserFacade.AddPeerSettings(ctx, toId, model.MakeUserPeerUtil(fromId), &mtproto.PeerSettings{
				//		AddContact:   true,
				//		BlockContact: true,
				//	})
				//}
			}

			pushUpdates := mtproto.MakeUpdatesByUpdatesUsersChats(
				in.Users,
				in.Chats,
				mtproto.MakeTLUpdateNewMessage(&mtproto.Update{
					Message_MESSAGE: inBox.GetMessage(),
					Pts_INT32:       inBox.Pts,
					PtsCount:        inBox.PtsCount,
				}).To_Update())

			if in.PeerType == mtproto.PEER_CHAT {
				switch inBox.GetMessage().GetAction().GetPredicateName() {
				case mtproto.Predicate_messageActionChatMigrateTo:
					_, _ = c.svcCtx.Dao.DialogClient.DialogInsertOrUpdateDialog(
						c.ctx,
						&dialog.TLDialogInsertOrUpdateDialog{
							UserId:          in.UserId,
							PeerType:        mtproto.PEER_CHAT,
							PeerId:          in.PeerId,
							TopMessage:      nil,
							ReadOutboxMaxId: nil,
							ReadInboxMaxId:  &wrapperspb.Int32Value{Value: inBox.MessageId},
							UnreadCount:     &wrapperspb.Int32Value{Value: 0},
							UnreadMark:      false,
							PinnedMsgId:     nil,
							Date2:           nil,
						})

					pushUpdates.PushFrontUpdate(mtproto.MakeTLUpdateReadHistoryInbox(&mtproto.Update{
						FolderId:         nil,
						Peer_PEER:        mtproto.MakePeerChat(in.PeerId),
						MaxId:            inBox.MessageId,
						StillUnreadCount: 0,
						Pts_INT32:        c.svcCtx.Dao.NextPtsId(c.ctx, in.UserId),
						PtsCount:         1,
					}).To_Update())
				}
			}

			var (
				isBot = false
			)

			for _, u := range in.GetUsers() {
				if u.GetId() == in.UserId {
					isBot = u.GetBot()
					break
				}
			}

			if isBot {
				if c.svcCtx.Dao.BotSyncClient != nil {
					_, err = c.svcCtx.Dao.BotSyncClient.SyncPushBotUpdates(c.ctx, &sync.TLSyncPushBotUpdates{
						UserId:  inBox.UserId,
						Updates: pushUpdates,
					})
				} else {
					// TODO: log
				}
			} else {
				_, err = c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx, &sync.TLSyncPushUpdates{
					UserId:  inBox.UserId,
					Updates: pushUpdates,
				})
			}
			if err != nil {
				c.Logger.Errorf("inbox.sendUserMessageToInboxV2 - error: %v", err)
			}
		}
	}

	return mtproto.EmptyVoid, nil
}
