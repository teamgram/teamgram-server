// Copyright (c) 2024 The Teamgooo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package core

import (
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesGetPeerDialogs
// messages.getPeerDialogs#e470bcfd peers:Vector<InputDialogPeer> = messages.PeerDialogs;
func (c *DialogsCore) MessagesGetPeerDialogs(in *tg.TLMessagesGetPeerDialogs) (*tg.MessagesPeerDialogs, error) {
	if len(in.Peers) == 1 {
		if inputDialogPeer, ok := in.Peers[0].(*tg.TLInputDialogPeer); ok && inputDialogPeer.Peer != nil {
			peer := tg.FromInputPeer2(0, inputDialogPeer.Peer)
			if c.MD != nil {
				peer = tg.FromInputPeer2(c.MD.UserId, inputDialogPeer.Peer)
			}

			if c.svcCtx != nil && c.svcCtx.DialogClient != nil && c.MD != nil && c.MD.UserId != 0 {
				dialogExt, err := c.svcCtx.DialogClient.DialogGetDialogById(c.ctx, &dialog.TLDialogGetDialogById{
					UserId:   c.MD.UserId,
					PeerType: peer.PeerType,
					PeerId:   peer.PeerId,
				})
				if err != nil {
					c.Logger.Errorf("messages.getPeerDialogs - DialogGetDialogById error: %v", err)
					return nil, err
				}
				if dialogExt != nil && dialogExt.Dialog != nil {
					topMessage := extractDialogTopMessage(dialogExt.Dialog)
					messages := []tg.MessageClazz{
						makePlaceholderDialogMessage(peer.PeerId, topMessage),
					}
					if fetched := c.fetchDialogTopMessage(c.MD.UserId, dialogExt.Dialog); fetched != nil {
						messages = []tg.MessageClazz{fetched}
					}
					return tg.MakeTLMessagesPeerDialogs(&tg.TLMessagesPeerDialogs{
						Dialogs:  []tg.DialogClazz{dialogExt.Dialog},
						Messages: messages,
						Chats:    makeDialogPeerChats(peer),
						Users:    makeDialogPeerUsers(peer),
						State: tg.MakeTLUpdatesState(&tg.TLUpdatesState{
							Pts:  1,
							Qts:  0,
							Date: 10,
							Seq:  0,
						}),
					}).ToMessagesPeerDialogs(), nil
				}
			}

			if peer.PeerType == tg.PEER_SELF || peer.PeerType == tg.PEER_USER {
				return tg.MakeTLMessagesPeerDialogs(&tg.TLMessagesPeerDialogs{
					Dialogs: []tg.DialogClazz{
						makePlaceholderDialog(peer.PeerId, 10),
					},
					Messages: []tg.MessageClazz{
						makePlaceholderDialogMessage(peer.PeerId, 10),
					},
					Chats: []tg.ChatClazz{},
					Users: []tg.UserClazz{
						makePlaceholderUser(peer.PeerId),
					},
					State: tg.MakeTLUpdatesState(&tg.TLUpdatesState{
						Pts:  1,
						Qts:  0,
						Date: 10,
						Seq:  0,
					}),
				}).ToMessagesPeerDialogs(), nil
			}
		}
	}

	// Return an empty peer-dialogs envelope until dialog/update stores are wired.
	return tg.MakeTLMessagesPeerDialogs(&tg.TLMessagesPeerDialogs{
		Dialogs:  []tg.DialogClazz{},
		Messages: []tg.MessageClazz{},
		Chats:    []tg.ChatClazz{},
		Users:    []tg.UserClazz{},
		State: tg.MakeTLUpdatesState(&tg.TLUpdatesState{
			Pts:  0,
			Qts:  0,
			Date: 0,
			Seq:  0,
		}),
	}).ToMessagesPeerDialogs(), nil
}

func extractDialogTopMessage(dialog tg.DialogClazz) int32 {
	if dialog == nil {
		return 10
	}
	if d, ok := dialog.(*tg.TLDialog); ok {
		if d.TopMessage > 0 {
			return d.TopMessage
		}
	}
	return 10
}

func makeDialogPeerUsers(peer *tg.PeerUtil) []tg.UserClazz {
	if peer == nil {
		return []tg.UserClazz{}
	}
	switch peer.PeerType {
	case tg.PEER_SELF, tg.PEER_USER:
		return []tg.UserClazz{makePlaceholderUser(peer.PeerId)}
	default:
		return []tg.UserClazz{}
	}
}

func makeDialogPeerChats(peer *tg.PeerUtil) []tg.ChatClazz {
	if peer == nil {
		return []tg.ChatClazz{}
	}
	switch peer.PeerType {
	case tg.PEER_CHAT:
		return []tg.ChatClazz{tg.MakeTLChatEmpty(&tg.TLChatEmpty{Id: peer.PeerId})}
	case tg.PEER_CHANNEL:
		return []tg.ChatClazz{tg.MakeTLChannelForbidden(&tg.TLChannelForbidden{
			Id:         peer.PeerId,
			AccessHash: 0,
			Title:      "",
		})}
	default:
		return []tg.ChatClazz{}
	}
}

func makePlaceholderDialog(peerID int64, topMessage int32) tg.DialogClazz {
	return tg.MakeTLDialog(&tg.TLDialog{
		Peer: tg.MakeTLPeerUser(&tg.TLPeerUser{
			UserId: peerID,
		}),
		TopMessage:      topMessage,
		ReadInboxMaxId:  topMessage,
		ReadOutboxMaxId: topMessage,
		UnreadCount:     0,
		NotifySettings:  tg.MakeTLPeerNotifySettings(&tg.TLPeerNotifySettings{}),
	})
}

func makePlaceholderDialogMessage(peerID int64, messageID int32) tg.MessageClazz {
	return tg.MakeTLMessage(&tg.TLMessage{
		Out:     true,
		Id:      messageID,
		FromId:  tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: peerID}),
		PeerId:  tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: peerID}),
		Date:    messageID,
		Message: "placeholder",
	})
}

func makePlaceholderUser(userID int64) tg.UserClazz {
	return tg.MakeTLUserEmpty(&tg.TLUserEmpty{Id: userID})
}

func makePlaceholderDialogMessageID(randomID int64) int32 {
	if randomID < 0 {
		randomID = -randomID
	}
	id := int32(randomID % 0x7fffffff)
	if id == 0 {
		id = 1
	}
	return id
}
