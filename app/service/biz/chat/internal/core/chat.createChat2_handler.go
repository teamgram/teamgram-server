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
	"math/rand"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/internal/dal/dataobject"
)

// ChatCreateChat2
// chat.createChat2 creator_id:long user_id_list:Vector<long> title:string = MutableChat;
func (c *ChatCore) ChatCreateChat2(in *chat.TLChatCreateChat2) (*mtproto.MutableChat, error) {
	var (
		chatsDO    *dataobject.ChatsDO
		err        error
		date       = time.Now().Unix()
		creatorId  = in.CreatorId
		userIdList = in.UserIdList
		title      = in.Title
	)

	// TODO:
	if chatsDO, err = c.svcCtx.Dao.ChatsDAO.SelectLastCreator(c.ctx, creatorId); err != nil {
		c.Logger.Errorf("chat.createChat2 - error: %v", err)
		return nil, err
	} else if chatsDO != nil {
		if date-chatsDO.Date < createChatFlood {
			err = mtproto.NewErrFloodWaitX(int32(date - chatsDO.Date))
			c.Logger.Errorf("createChat error: %v. lastCreate = ", err, chatsDO.Date)
			return nil, err
		}
	}

	chatsDO = &dataobject.ChatsDO{
		Id:                   0,
		CreatorUserId:        creatorId,
		AccessHash:           rand.Int63(),
		RandomId:             0,
		ParticipantCount:     int32(1 + len(userIdList)),
		Title:                title,
		About:                "",
		PhotoId:              0,
		DefaultBannedRights:  int64(mtproto.MakeDefaultBannedRights().ToBannedRights()),
		MigratedToId:         0,
		MigratedToAccessHash: 0,
		Deactivated:          false,
		Version:              1,
		Date:                 date,
	}

	participantDOList := make([]*dataobject.ChatParticipantsDO, 0)
	for i := 0; i < len(userIdList)+1; i++ {
		if i == 0 {
			participantDOList = append(participantDOList, &dataobject.ChatParticipantsDO{
				UserId:          creatorId,
				ParticipantType: mtproto.ChatMemberCreator,
				Link:            chat.GenChatInviteHash(),
				InviterUserId:   0,
				InvitedAt:       date,
				Date2:           date,
				IsBot:           false,
			})
		} else {
			participantDOList = append(participantDOList, &dataobject.ChatParticipantsDO{
				UserId:          userIdList[i-1],
				ParticipantType: mtproto.ChatMemberNormal,
				Link:            "",
				InviterUserId:   creatorId,
				InvitedAt:       date,
				Date2:           date,
				IsBot:           false,
			})
		}
	}

	for _, id := range in.Bots {
		participantDOList = append(participantDOList, &dataobject.ChatParticipantsDO{
			UserId:          id,
			ParticipantType: mtproto.ChatMemberNormal,
			Link:            "",
			InviterUserId:   creatorId,
			InvitedAt:       date,
			Date2:           date,
			IsBot:           true,
		})
	}

	tR := sqlx.TxWrapper(c.ctx, c.svcCtx.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		// 1. insert chat
		chatsDO.Id, _, err = c.svcCtx.Dao.ChatsDAO.InsertTx(tx, chatsDO)
		if err != nil {
			result.Err = err
			return
		}
		//chatsDO.Id = chatId
		for i := 0; i < len(participantDOList); i++ {
			participantDOList[i].ChatId = chatsDO.Id
		}

		_, _, err = c.svcCtx.Dao.ChatParticipantsDAO.InsertBulkTx(tx, participantDOList)
		if err != nil {
			result.Err = err
			return
		}

		_, _, result.Err = c.svcCtx.Dao.ChatInvitesDAO.InsertTx(tx, &dataobject.ChatInvitesDO{
			ChatId:    chatsDO.Id,
			AdminId:   creatorId,
			Link:      participantDOList[0].Link,
			Permanent: true,
			Date2:     date,
		})
		return
	})

	if tR.Err != nil {
		err = tR.Err
		c.Logger.Errorf("chat.createChat2 - error: %v", tR.Err)
		return nil, tR.Err
	}

	chat2 := mtproto.MakeTLMutableChat(&mtproto.MutableChat{
		Chat:             c.svcCtx.Dao.MakeImmutableChatByDO(chatsDO),
		ChatParticipants: make([]*mtproto.ImmutableChatParticipant, 0, len(participantDOList)),
	}).To_MutableChat()

	for i := 0; i < len(participantDOList); i++ {
		chat2.ChatParticipants = append(chat2.ChatParticipants,
			c.svcCtx.Dao.MakeImmutableChatParticipant(participantDOList[i]))
	}

	chat2.Chat.ParticipantsCount = int32(len(participantDOList))

	// c.svcCtx.Dao.PutMutableChat(c.ctx, chat2)

	return chat2, nil
}
