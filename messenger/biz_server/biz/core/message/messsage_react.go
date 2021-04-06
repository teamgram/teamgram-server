package message

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dataobject"
	"time"
)

func (m *MessageModel) PushReaction(messageId int64, reactId, senderUserId int32) (bool, error) {
	glog.Infoln("Push reaction")
	mReact := &dataobject.MessageReactDataDO{
		ReactDataId:   0,
		ReactId:       reactId,
		MessageDataID: messageId,
		SenderUserId:  senderUserId,
		Date3:         int32(time.Now().Unix()),
		Deleted:       0,
	}
	m.dao.MessageReactDataDAO.Insert(mReact)
	return true, nil
}
