package messages

import (
	"context"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/base"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core"
	"github.com/nebula-chat/chatengine/messenger/sync/sync_client"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"time"
)

func (s *MessagesServiceImpl) MessagesSendReact(ctx context.Context, request *mtproto.TLMessagesSendReact) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.sendReact#136D148A - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		peer *base.PeerUtil
		err  error
	)

	if request.GetPeer().GetConstructor() == mtproto.TLConstructor_CRC32_inputPeerEmpty {
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error("messages.sendReact#136D148A - invalid peer", err)
		return mtproto.ToBool(false), err
	}

	peer = base.FromInputPeer(request.GetPeer())

	// Only sender
	outboxData, err := s.MessageModel.GetMessageBox3(peer.PeerType, request.GetId())
	if err != nil {
		glog.Infoln("messages.sendReact#136D148A - get message box2", err)
		return nil, err
	}

	s.PushReaction(request.GetId(), request.ReactId, md.UserId)

	if peer.PeerType == base.PEER_SELF || peer.PeerType == base.PEER_USER {
		react := &mtproto.TLUpdateNewReact{Data2: &mtproto.Update_Data{
			UserId:        md.UserId,
			ReactId:       int64(request.ReactId),
			MessageDataId: request.GetId(), //1879052288
			Pts:           int32(core.NextPtsId(md.UserId)),
			PtsCount:      1,
		}}

		updates := &mtproto.TLUpdates{Data2: &mtproto.Updates_Data{
			Updates: []*mtproto.Update{react.To_Update()},
			Seq:     0,
			Date:    int32(time.Now().Unix()),
		}}

		go sync_client.GetSyncClient().PushUpdates(md.UserId, updates.To_Updates())
	}

	// Receiver msg ???
	boxList := s.MessageModel.GetPeerMessageListByMessageDataId(outboxData.OwnerId, outboxData.MessageDataId)
	for _, box := range boxList {
		react := &mtproto.TLUpdateNewReact{Data2: &mtproto.Update_Data{
			UserId:   box.OwnerId,
			ReactId:  int64(request.ReactId),
			MessageDataId: request.GetId(),
			Pts:      int32(core.NextPtsId(box.OwnerId)),
			PtsCount: 1,
		}}

		updates := &mtproto.TLUpdates{Data2: &mtproto.Updates_Data{
			Updates: []*mtproto.Update{react.To_Update()},
			Date:    int32(time.Now().Unix()),
		}}
		go sync_client.GetSyncClient().PushUpdates(box.OwnerId, updates.To_Updates())
	}

	glog.Infof("messages.sendReact#136D148A - reply: %s", logger.JsonDebugData(mtproto.ToBool(true)))
	return mtproto.ToBool(true), nil
}
