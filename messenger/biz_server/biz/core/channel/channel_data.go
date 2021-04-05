package channel

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/base"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dataobject"
	"github.com/nebula-chat/chatengine/mtproto"
)

const (
	kChannelParticipant        = 0
	kChannelParticipantCreator = 1
	kChannelParticipantAdmin   = 2
)

const (
	kChannelParticipantStateNormal = 0 // normal
	kChannelParticipantStateLeft   = 1 // left
	kChannelParticipantStateKicked = 2 // kicked
)

type channelLogicData struct {
	channel      *dataobject.ChannelsDO
	participants []dataobject.ChannelParticipantsDO
	dao          *channelsDAO
	cb           core.PhotoCallback
}

func makeChannelParticipantByDO(do *dataobject.ChannelParticipantsDO) (participant *mtproto.ChannelParticipant) {
	participant = &mtproto.ChannelParticipant{Data2: &mtproto.ChannelParticipant_Data{
		UserId: do.UserId,
		// InviterId: do.InviterUserId,
		// Date:      do.InvitedAt,
	}}

	switch do.ParticipantType {
	case kChannelParticipant:
		participant.Constructor = mtproto.TLConstructor_CRC32_chatParticipant
		participant.Data2.Date = do.InvitedAt
	case kChannelParticipantCreator:
		participant.Constructor = mtproto.TLConstructor_CRC32_channelParticipantCreator
	case kChannelParticipantAdmin:
		participant.Constructor = mtproto.TLConstructor_CRC32_channelParticipantAdmin
		participant.Data2.Date = do.InvitedAt
	default:
		panic("channelParticipant type error.")
	}

	return
}

// func (m *ChannelModel) NewChannelLogicById(channelId int32) (channelData *channelLogicData, err error) {
// 	channelDO := m.dao.ChannelsDAO.Select(channelId)
// 	if channelDO == nil {
// 		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_CHANNEL_ID_INVALID)
// 	} else {
// 		channelData = &channelLogicData{
// 			channel: channelDO,
// 			dao:     m.dao,
// 			cb:      m.photoCallback,
// 		}
// 		channelData.checkOrLoadChannelParticipantList()
// 	}
// 	return
// }

// func (m *ChannelModel) NewChannelLogicByLink(link string) (channelData *channelLogicData, err error) {
// 	if link == "" {
// 		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_INVITE_HASH_INVALID)
// 	}

// 	channelDO := m.dao.ChannelsDAO.SelectByLink(link)
// 	if channelDO == nil {
// 		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_INVITE_HASH_INVALID)
// 	} else {
// 		channelData = &channelLogicData{
// 			channel: channelDO,
// 			dao:     m.dao,
// 			cb:      m.photoCallback,
// 		}
// 	}
// 	return
// }
func (m *ChannelModel) CreateChannel(creatorId int32, userIdList []int32, title string) (channelData *channelLogicData, err error) {
	glog.Errorln("Create channel not impl yet")
	return nil, nil
	// date := int32(time.Now().Unix())
	// //  TODO(@benqi): Check FLOOD_WAIT_
	// chatDO := m.dao.ChatsDAO.SelectLastCreator(creatorId)
	// if chatDO != nil {
	// 	if date-chatDO.Date < kCreateChatFlood {
	// 		err = mtproto.NewFloodWaitX2(int(date - chatDO.Date))
	// 		glog.Error("create error: ", err, ". lastCreate = ", chatDO.Date)
	// 		return nil, err
	// 	}
	// }

	// // TODO(@benqi): 事务
	// chatData = &chatLogicData{
	// 	chat: &dataobject.ChatsDO{
	// 		CreatorUserId: creatorId,
	// 		AccessHash:    rand.Int63(),
	// 		// TODO(@benqi): use message_id is randomid
	// 		// RandomId:         helper.NextSnowflakeId(),
	// 		ParticipantCount: int32(1 + len(userIdList)),
	// 		Title:            title,
	// 		PhotoId:          0,
	// 		Version:          1,
	// 		Date:             int32(time.Now().Unix()),
	// 	},
	// 	participants: make([]dataobject.ChatParticipantsDO, 1+len(userIdList)),
	// 	dao:          m.dao,
	// 	cb:           m.photoCallback,
	// }
	// chatData.chat.Id = int32(m.dao.ChatsDAO.Insert(chatData.chat))

	// chatData.participants = make([]dataobject.ChatParticipantsDO, 1+len(userIdList))
	// chatData.participants[0].ChatId = chatData.chat.Id
	// chatData.participants[0].UserId = creatorId
	// chatData.participants[0].ParticipantType = kChatParticipantCreator
	// m.dao.ChatParticipantsDAO.Insert(&chatData.participants[0])

	// for i := 0; i < len(userIdList); i++ {
	// 	chatData.participants[i+1].ChatId = chatData.chat.Id
	// 	chatData.participants[i+1].UserId = userIdList[i]
	// 	chatData.participants[i+1].ParticipantType = kChatParticipant
	// 	chatData.participants[i+1].InviterUserId = creatorId
	// 	chatData.participants[i+1].InvitedAt = chatData.chat.Date
	// 	m.dao.ChatParticipantsDAO.Insert(&chatData.participants[i+1])
	// }
	// return
}

func (this *channelLogicData) GetPhotoId() int64 {
	return this.channel.PhotoId
}

func (this *channelLogicData) GetChannelId() int32 {
	return this.channel.Id
}

func (this *channelLogicData) GetVersion() int32 {
	return this.channel.Version
}

func (this *channelLogicData) GetLink() string {
	return this.channel.Link
}

func (this *channelLogicData) GetCreator() int32 {
	return this.channel.CreatorUserId
}

func (this *channelLogicData) checkOrLoadChannelParticipantList() {
	if len(this.participants) == 0 {
		this.participants = this.dao.ChannelParticipantsDAO.SelectList(this.channel.Id)
	}
}

func (this *channelLogicData) MakeMessageService(fromId int32, action *mtproto.MessageAction) *mtproto.Message {
	peer := &base.PeerUtil{
		PeerType: base.PEER_CHANNEL,
		PeerId:   this.channel.Id,
	}

	message := &mtproto.TLMessageService{Data2: &mtproto.Message_Data{
		Date:   this.channel.Date,
		FromId: fromId,
		ToId:   peer.ToPeer(),
		Action: action,
	}}
	return message.To_Message()
}

func (this *channelLogicData) MakeCreateChannelMessage(creatorId int32) *mtproto.Message {
	idList := this.GetChannelParticipantIdList()
	action := &mtproto.TLMessageActionChannelCreate{Data2: &mtproto.MessageAction_Data{
		Title: this.channel.Title,
		Users: idList,
	}}

	return this.MakeMessageService(creatorId, action.To_MessageAction())
}

func (this *channelLogicData) MakeAddUserMessage(inviterId, channelUserId int32) *mtproto.Message {
	// idList := this.GetChannelParticipantIdList()
	action := &mtproto.TLMessageActionChatAddUser{Data2: &mtproto.MessageAction_Data{
		Title: this.channel.Title,
		Users: []int32{channelUserId},
	}}

	return this.MakeMessageService(inviterId, action.To_MessageAction())
}

func (this *channelLogicData) MakeJoinedByLinkMessage(inviterId int32) *mtproto.Message {
	action := &mtproto.TLMessageActionChatJoinedByLink{Data2: &mtproto.MessageAction_Data{
		InviterId: inviterId,
	}}

	return this.MakeMessageService(inviterId, action.To_MessageAction())
}

func (this *channelLogicData) MakeDeleteUserMessage(operatorId, channelUserId int32) *mtproto.Message {
	// idList := this.GetChannelParticipantIdList()
	action := &mtproto.TLMessageActionChatDeleteUser{Data2: &mtproto.MessageAction_Data{
		Title:  this.channel.Title,
		UserId: channelUserId,
	}}

	return this.MakeMessageService(operatorId, action.To_MessageAction())
}

func (this *channelLogicData) MakeChannelEditTitleMessage(operatorId int32, title string) *mtproto.Message {
	// idList := this.GetChannelParticipantIdList()
	action := &mtproto.TLMessageActionChatEditTitle{Data2: &mtproto.MessageAction_Data{
		Title: title,
	}}

	return this.MakeMessageService(operatorId, action.To_MessageAction())
}

func (this *channelLogicData) GetChannelParticipantList() []*mtproto.ChannelParticipant {
	this.checkOrLoadChannelParticipantList()

	participantList := make([]*mtproto.ChannelParticipant, 0, len(this.participants))
	for i := 0; i < len(this.participants); i++ {
		if this.participants[i].State == 0 {
			participantList = append(participantList, makeChannelParticipantByDO(&this.participants[i]))
		}
	}
	return participantList
}

func (this *channelLogicData) GetChannelParticipantIdList() []int32 {
	this.checkOrLoadChannelParticipantList()

	idList := make([]int32, 0, len(this.participants))
	for i := 0; i < len(this.participants); i++ {
		if this.participants[i].State == 0 {
			idList = append(idList, this.participants[i].UserId)
		}
	}
	return idList
}

func (this *channelLogicData) GetChannelParticipants() *mtproto.TLChannelsChannelParticipants {
	this.checkOrLoadChannelParticipantList()
	glog.Errorln("GetChannelParticipants not impl yet")

	// TODO: TLChannelParticipants
	return &mtproto.TLChannelsChannelParticipants{Data2: &mtproto.Channels_ChannelParticipants_Data{
		// 		ChannelId:    this.channel.Id,
		Participants: this.GetChannelParticipantList(),
		// 		Version:      this.channel.Version,
	}}
}

// func (this *channelLogicData) AddChannelUser(inviterId, userId int32) error {
// 	this.checkOrLoadChannelParticipantList()

// 	// TODO(@benqi): check userId exisits
// 	var founded = -1
// 	for i := 0; i < len(this.participants); i++ {
// 		if userId == this.participants[i].UserId {
// 			if this.participants[i].State != kChannelParticipantStateNormal {
// 				founded = i
// 			} else {
// 				err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PARTICIPANT_EXISTED)
// 				glog.Errorf("participant existed: {channel_id: %d, invited by: %d, user_id: %d}, error: ",
// 					this.channel.Id, inviterId, userId, err)
// 				return err
// 			}
// 		}
// 	}

// 	var now = int32(time.Now().Unix())

// 	if founded != -1 {
// 		this.participants[founded].State = 0
// 		this.dao.ChannelParticipantsDAO.Update(int8(kChannelParticipant), inviterId, now, this.participants[founded].Id)
// 	} else {
// 		channelParticipant := &dataobject.ChannelParticipantsDO{
// 			ChannelId:       this.channel.Id,
// 			UserId:          userId,
// 			ParticipantType: kChannelParticipant,
// 			InviterUserId:   inviterId,
// 			InvitedAt:       now,
// 			// JoinedAt:        now,
// 		}
// 		channelParticipant.Id = int32(this.dao.ChannelParticipantsDAO.Insert(channelParticipant))
// 		this.participants = append(this.participants, *channelParticipant)
// 	}

// 	// update channel
// 	this.channel.ParticipantCount += 1
// 	this.channel.Version += 1
// 	this.channel.Date = now
// 	this.dao.ChannelsDAO.UpdateParticipantCount(this.channel.ParticipantCount, this.chat.Id)

// 	return nil
// }

func (this *channelLogicData) findChannelParticipant(selfUserId int32) (int, *dataobject.ChannelParticipantsDO) {
	for i := 0; i < len(this.participants); i++ {
		if this.participants[i].UserId == selfUserId {
			return i, &this.participants[i]
		}
	}
	return -1, nil
}

func (this *channelLogicData) ToChat(selfUserId int32) *mtproto.Chat {
	// TODO(@benqi): kicked:flags.1?true left:flags.2?true admins_enabled:flags.3?true admin:flags.4?true deactivated:flags.5?true

	var (
		forbidden        = false
		participant      *dataobject.ChannelParticipantsDO
		participantCount int32
	)

	for i := 0; i < len(this.participants); i++ {
		if this.participants[i].State == kChannelParticipantStateNormal {
			participantCount++
		}

		if selfUserId == this.participants[i].UserId {
			participant = &this.participants[i]
		}

		if this.participants[i].UserId == selfUserId && this.participants[i].State != kChannelParticipantStateNormal {
			forbidden = true
			// break
		}
	}

	//if participant == nil || participant.State == kChatParticipantStateLeft {
	//	chat := &mtproto.TLChatEmpty{Data2: &mtproto.Chat_Data{
	//		Id: this.chat.Id,
	//	}}
	//	return chat.To_Chat()
	//}

	if participant == nil || forbidden {
		channel := &mtproto.TLChannelForbidden{Data2: &mtproto.Chat_Data{
			Id:    this.channel.Id,
			Title: this.channel.Title,
		}}
		return channel.To_Chat()
	} else {
		channel := &mtproto.TLChannel{Data2: &mtproto.Chat_Data{
			Creator:           this.channel.CreatorUserId == selfUserId,
			Kicked:            participant.State == kChannelParticipantStateKicked,
			Left:              participant.State == kChannelParticipantStateLeft,
			AdminsEnabled:     this.channel.AdminsEnabled == 1,
			Deactivated:       this.channel.Deactivated == 1,
			ParticipantsCount: participantCount,
			Id:                this.channel.Id,
			Title:             this.channel.Title,
			Date:              this.channel.Date,
			Version:           this.channel.Version,
		}}

		if participant.State == kChannelParticipantStateNormal {
			if participant.ParticipantType == kChannelParticipantAdmin {
				channel.Data2.Admin = true
			}
		}

		if this.channel.PhotoId == 0 {
			channel.SetPhoto(mtproto.NewTLChatPhotoEmpty().To_ChatPhoto())
		} else {
			// sizeList, _ := nbfs_client.GetPhotoSizeList(this.channel.PhotoId)
			channel.SetPhoto(this.cb.GetChatPhoto(this.channel.PhotoId))
		}
		return channel.To_Chat()
	}
}
