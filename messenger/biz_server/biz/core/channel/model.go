package channel

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dao"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dao/mysql_dao"
)

type channelsDAO struct {
	*mysql_dao.ChannelDAO
	*mysql_dao.ChannelParticipantsDAO
}

type ChannelModel struct {
	dao           *channelsDAO
	photoCallback core.PhotoCallback
	// accountCallback core.AccountCallback
}

func (m *ChannelModel) RegisterCallback(cb interface{}) {
	switch cb.(type) {
	case core.PhotoCallback:
		glog.Info("chatModel - register core.PhotoCallback")
		m.photoCallback = cb.(core.PhotoCallback)
		// case core.AccountCallback:
		// 	glog.Info("chatModel - register core.AccountCallback")
		// 	m.accountCallback = cb.(core.AccountCallback)
	}
}

func (m *ChannelModel) InstallModel() {
	// m.dao.CommonDAO = dao.GetCommonDAO(dao.DB_MASTER)
	// m.dao.UsersDAO = dao.GetUsersDAO(dao.DB_MASTER)
	m.dao.ChannelDAO = dao.GetChannelsDAO(dao.DB_MASTER)
	m.dao.ChannelParticipantsDAO = dao.GetChannelParticipantsDAO(dao.DB_MASTER)
}

// func (m *ChannelModel) GetChatListBySelfAndIDList(selfUserId int32, idList []int32) (chats []*mtproto.Chat) {
// 	if len(idList) == 0 {
// 		return []*mtproto.Chat{}
// 	}

// 	chats = make([]*mtproto.Chat, 0, len(idList))

// 	// TODO(@benqi): 性能优化，从DB里一次性取出所有的chatList
// 	for _, id := range idList {
// 		chatData, err := m.NewChannelLogicById(id)
// 		if err != nil {
// 			glog.Error("getChannelListBySelfIDList - not find channel_id: ", id)
// 			chatEmpty := &mtproto.TLChatEmpty{Data2: &mtproto.Chat_Data{
// 				Id: id,
// 			}}
// 			chats = append(chats, chatEmpty.To_Chat())
// 		} else {ss
// 			chats = append(chats, chatData.ToChat(selfUserId))
// 		}
// 	}

// 	return
// }

func init() {
	core.RegisterCoreModel(&ChannelModel{dao: &channelsDAO{}})
}
