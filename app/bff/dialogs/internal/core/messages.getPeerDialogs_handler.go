// Copyright 2022 Teamgram Authors
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
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/app/service/biz/updates/updates"
)

/*
	private func synchronizeGroupMessageStats(postbox: Postbox, network: Network, groupId: PeerGroupId, namespace: MessageId.Namespace) -> Signal<Void, NoError> {
		return postbox.transaction { transaction -> Signal<Void, NoError> in
			if namespace != Namespaces.Message.Cloud || groupId == .root {
				transaction.confirmSynchronizedPeerGroupMessageStats(groupId: groupId, namespace: namespace)
				return .complete()
			}

			if !transaction.doesChatListGroupContainHoles(groupId: groupId) {
				transaction.recalculateChatListGroupStats(groupId: groupId)
				return .complete()
			}

			return network.request(Api.functions.messages.getPeerDialogs(peers: [.inputDialogPeerFolder(folderId: groupId.rawValue)]))
			|> map(Optional.init)
			|> `catch` { _ -> Signal<Api.messages.PeerDialogs?, NoError> in
				return .single(nil)
			}
			|> mapToSignal { result -> Signal<Void, NoError> in
				return postbox.transaction { transaction in
					if let result = result {
						switch result {
						case let .peerDialogs(peerDialogs):
							for dialog in peerDialogs.dialogs {
								switch dialog {
									case let .dialogFolder(dialogFolder):
										transaction.resetPeerGroupSummary(groupId: groupId, namespace: namespace, summary: PeerGroupUnreadCountersSummary(all: PeerGroupUnreadCounters(messageCount: dialogFolder.unreadMutedMessagesCount, chatCount: dialogFolder.unreadMutedPeersCount)))
									case .dialog:
										assertionFailure()
										break
								}
							}
						}
					}
					transaction.confirmSynchronizedPeerGroupMessageStats(groupId: groupId, namespace: namespace)
				}
			}
		}
		|> switchToLatest
	}
*/

// MessagesGetPeerDialogs
// messages.getPeerDialogs#e470bcfd peers:Vector<InputDialogPeer> = messages.PeerDialogs;
func (c *DialogsCore) MessagesGetPeerDialogs(in *mtproto.TLMessagesGetPeerDialogs) (*mtproto.Messages_PeerDialogs, error) {
	var (
		peerDialogIdList []int64
		folderId         int32 = -1
	)
	for _, peer := range in.GetPeers() {
		switch peer.GetPredicateName() {
		case mtproto.Predicate_inputDialogPeer:
			p := mtproto.FromInputPeer2(c.MD.UserId, peer.Peer)
			switch p.PeerType {
			case mtproto.PEER_SELF:
			case mtproto.PEER_USER:
			case mtproto.PEER_CHAT:
			case mtproto.PEER_CHANNEL:
			default:
				err := mtproto.ErrPeerIdInvalid
				c.Logger.Errorf("messages.getPeerDialogs - getPeerDialogs error: %v", err)
				return nil, err
			}
			peerDialogIdList = append(peerDialogIdList, mtproto.MakePeerDialogId(p.PeerType, p.PeerId))
		case mtproto.Predicate_inputDialogPeerFolder:
			if folderId == -1 {
				folderId = peer.FolderId
			} else {
				// if has inputDialogPeerFolder, len(peers) == 1
				err := mtproto.ErrFolderIdInvalid
				c.Logger.Errorf("messages.getPeerDialogs - getPeerDialogs error: %v", err)
				return nil, err
			}
		default:
			err := mtproto.ErrPeerIdInvalid
			c.Logger.Errorf("messages.getPeerDialogs - getPeerDialogs error: %v", err)
			return nil, err
		}
	}

	var (
		err     error
		dialogs dialog.DialogExtList
		state   *mtproto.Updates_State
	)

	state, err = c.svcCtx.Dao.UpdatesClient.UpdatesGetState(c.ctx, &updates.TLUpdatesGetState{
		AuthKeyId: c.MD.AuthId,
		UserId:    c.MD.UserId,
	})
	if err != nil {
		c.Logger.Errorf("messages.getPeerDialogs - getState error: %v", err)
		return nil, err
	}

	if len(peerDialogIdList) > 0 {
		dList, err := c.svcCtx.Dao.DialogClient.DialogGetDialogsByIdList(c.ctx, &dialog.TLDialogGetDialogsByIdList{
			UserId: c.MD.UserId,
			IdList: peerDialogIdList,
		})
		if err != nil {
			c.Logger.Errorf("messages.getPeerDialogs - getPeerDialogs error: %v", err)
			return nil, err
		}
		dialogs = dList.GetDatas()
	} else if folderId != -1 {
		dList, err := c.svcCtx.Dao.DialogClient.DialogGetDialogFolder(c.ctx, &dialog.TLDialogGetDialogFolder{
			UserId:   c.MD.UserId,
			FolderId: folderId,
		})
		if err != nil {
			c.Logger.Errorf("messages.getPeerDialogs - getPeerDialogs error: %v", err)
			return nil, err
		}
		dialogs = dList.GetDatas()
	}

	return c.makeMessagesDialogs(dialogs).ToMessagesPeerDialogs(state), nil
}
