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
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/internal/dal/dataobject"
)

/*
# android source code:
## case1
```
    public void loadUsers() {
        if (usersLoading) {
            return;
        }

        boolean hasMoreJoinedUsers = invite.usage > joinedUsers.size();
        boolean hasMoreRequestedUsers = invite.request_needed && invite.requested > requestedUsers.size();
        boolean loadRequestedUsers;
        if (hasMoreJoinedUsers) {
            loadRequestedUsers = false;
        } else if (hasMoreRequestedUsers) {
            loadRequestedUsers = true;
        } else {
            return;
        }

        final List<TLRPC.TL_chatInviteImporter> importersList = loadRequestedUsers ? requestedUsers : joinedUsers;
        TLRPC.TL_messages_getChatInviteImporters req = new TLRPC.TL_messages_getChatInviteImporters();
        req.flags |= 2;
        req.link = invite.link;
        req.peer = MessagesController.getInstance(UserConfig.selectedAccount).getInputPeer(-chatId);
        req.requested = loadRequestedUsers;
        if (importersList.isEmpty()) {
            req.offset_user = new TLRPC.TL_inputUserEmpty();
        } else {
            TLRPC.TL_chatInviteImporter invitedUser = importersList.get(importersList.size() - 1);
            req.offset_user = MessagesController.getInstance(currentAccount).getInputUser(users.get(invitedUser.user_id));
            req.offset_date = invitedUser.date;
        }

        usersLoading = true;
        ConnectionsManager.getInstance(UserConfig.selectedAccount).sendRequest(req, (response, error) -> {
```

## case2
```
    public int getImporters(final long chatId, final String query, TLRPC.TL_chatInviteImporter lastImporter, LongSparseArray<TLRPC.User> users, RequestDelegate onComplete) {
        boolean isEmptyQuery = TextUtils.isEmpty(query);
        TLRPC.TL_messages_getChatInviteImporters req = new TLRPC.TL_messages_getChatInviteImporters();
        req.peer = MessagesController.getInstance(currentAccount).getInputPeer(-chatId);
        req.requested = true;
        req.limit = 30;
        if (!isEmptyQuery) {
            req.q = query;
            req.flags |= 4;
        }
        if (lastImporter == null) {
            req.offset_user = new TLRPC.TL_inputUserEmpty();
        } else {
            req.offset_user = getMessagesController().getInputUser(users.get(lastImporter.user_id));
            req.offset_date = lastImporter.date;
        }
```

## case3
```
    public void loadUsers(TLRPC.TL_chatInviteExported invite, long chatId) {
        if (invite == null) {
            setUsers(0, null);
            return;
        }
        setUsers(invite.usage, invite.importers);
        if (invite.usage > 0 && invite.importers == null && !loadingImporters) {
            TLRPC.TL_messages_getChatInviteImporters req = new TLRPC.TL_messages_getChatInviteImporters();
            req.link = invite.link;
            req.peer = MessagesController.getInstance(UserConfig.selectedAccount).getInputPeer(-chatId);
            req.offset_user = new TLRPC.TL_inputUserEmpty();
            req.limit = Math.min(invite.usage, 3);

            loadingImporters = true;
            ConnectionsManager.getInstance(UserConfig.selectedAccount).sendRequest(req, (response, error) -> {
```
*/

// ChatGetChatInviteImporters
// chat.getChatInviteImporters flags:# self_id:long chat_id:long requested:flags.0?true link:flags.1?string q:flags.2?string offset_date:int offset_user:long limit:int = Vector<ChatInviteImporter>;
func (c *ChatCore) ChatGetChatInviteImporters(in *chat.TLChatGetChatInviteImporters) (*chat.Vector_ChatInviteImporter, error) {
	var (
		rInvites []*mtproto.ChatInviteImporter
		link     = chat.GetInviteHashByLink(in.GetLink().GetValue())
		limit    = in.Limit
	)

	// TODO: see (case1, case2, case3)

	if limit == 0 {
		limit = 50
	}

	// TODO: q

	// TODO: see (case1, case2, case3)
	var (
		requested int32
	)
	if in.GetRequested() {
		requested = 1
	} else {
		requested = 0
	}

	if requested == 1 {
		c.svcCtx.Dao.ChatInviteParticipantsDAO.SelectRecentRequestedListWithCB(
			c.ctx,
			in.GetChatId(),
			func(sz, i int, v *dataobject.ChatInviteParticipantsDO) {
				rInvites = append(rInvites, mtproto.MakeTLChatInviteImporter(&mtproto.ChatInviteImporter{
					Requested:  v.Requested,
					UserId:     v.UserId,
					Date:       int32(v.Date2),
					About:      nil,
					ApprovedBy: nil,
				}).To_ChatInviteImporter())
			})
		if rInvites == nil {
			rInvites = []*mtproto.ChatInviteImporter{}
		}
	} else {
		c.svcCtx.Dao.ChatInviteParticipantsDAO.SelectListByLinkWithCB(
			c.ctx,
			link,
			requested,
			func(sz, i int, v *dataobject.ChatInviteParticipantsDO) {
				rInvites = append(rInvites, mtproto.MakeTLChatInviteImporter(&mtproto.ChatInviteImporter{
					Requested:  v.Requested,
					UserId:     v.UserId,
					Date:       int32(v.Date2),
					About:      nil,
					ApprovedBy: mtproto.MakeFlagsInt64(v.ApprovedBy),
				}).To_ChatInviteImporter())
			})

		if rInvites == nil {
			rInvites = []*mtproto.ChatInviteImporter{}
		}
	}

	var (
		offset = 0
	)

	for i, v := range rInvites {
		if in.OffsetUser == v.UserId && in.OffsetDate == v.Date {
			offset = i + 1
			break
		}
	}
	if len(rInvites) >= offset+int(limit) {
		rInvites = rInvites[offset : offset+int(limit)]
	} else {
		rInvites = rInvites[offset:]
	}

	// c.Logger.Errorf("offset: %d", offset)

	return &chat.Vector_ChatInviteImporter{
		Datas: rInvites,
	}, nil
}
