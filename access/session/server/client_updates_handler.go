// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
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

// Author: Benqi (wubenqi@gmail.com)

package server

import (
	"container/list"
	"fmt"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/mtproto/rpc"
	"reflect"
	"time"
)

const (
	kConnUnknown = 0
	kTcpConn     = 1
	kHttpConn    = 2
	kPushConn    = 3
)

type clientUpdatesHandler struct {
	session         *clientSessionHandler
	pushSession     *clientSessionHandler
	syncMessages    []*pendingMessage
	connState       int
	connType        int
	tcpConnID       ClientConnID
	rpcConnID       ClientConnID
	pushConnID      ClientConnID
	httpWaitConnIDs *list.List
}

func newClientUpdatesHandler() *clientUpdatesHandler {
	return &clientUpdatesHandler{
		syncMessages:    []*pendingMessage{},
		connState:       kConnUnknown,
		connType:        kConnUnknown,
		httpWaitConnIDs: list.New(),
	}
}

func (c *clientUpdatesHandler) String() string {
	return fmt.Sprintf("{sess: %s, conn_state: %d, conn_type: %d, tcp_conn_id: %s, push_session: %s, push_conn_id: %s}",
		c.session,
		c.connState,
		c.connType,
		c.tcpConnID,
		c.pushSession,
		c.pushConnID)
}

func (c *clientUpdatesHandler) SubscribeUpdates(sess *clientSessionHandler, connID ClientConnID) error {
	// TODO(@benqi): clear
	glog.Infof("subscribeUpdates -- {last_connID: {%s}, last_sess: {%s}, connID: {%s}}, sess: {%s}",
		c.tcpConnID, c.session, sess, connID)

	if connID.connType == mtproto.TRANSPORT_TCP {
		c.tcpConnID = connID
		c.connState = kTcpConn
		c.session = sess

		if len(c.syncMessages) > 0 {
			cntl := zrpc.NewController()
			c.session.sendPendingMessagesToClient(connID, cntl, c.syncMessages)
			c.syncMessages = []*pendingMessage{}
		}
		glog.Info("subscribeUpdates - c: ", c)
	} else if connID.connType == mtproto.TRANSPORT_HTTP {
		for e := c.httpWaitConnIDs.Front(); e != nil; e = e.Next() {
			connID2, _ := e.Value.(ClientConnID)
			if connID2.Equal(connID) {
				// TODO(@benqi): log..
				err := fmt.Errorf("http_wait conn existed: %v", connID)
				glog.Error(err)
				return err
			}
		}
		c.httpWaitConnIDs.PushBack(connID)
		c.connState = kHttpConn
		c.session = sess
	} else {
		err := fmt.Errorf("invalid connType: %v", connID)
		glog.Error(err)
		return err
	}
	return nil
}

func (c *clientUpdatesHandler) Subscribe2Updates(sess *clientSessionHandler, connID ClientConnID) error {
	// TODO(@benqi): clear
	glog.Infof("subscribe2Updates -- {last_connID: {%s}, last_sess: {%s}, connID: {%s}}, sess: {%s}",
		c.tcpConnID, c.session, connID, sess)

	// if connID.connType == mtproto.TRANSPORT_TCP {
		c.pushConnID = connID
		c.connState = kTcpConn
		c.pushSession = sess

		//if len(c.syncMessages) > 0 {
		//	cntl := zrpc.NewController()
		//	c.pushSession.sendPendingMessagesToClient(connID, cntl, c.syncMessages)
		//	c.syncMessages = []*pendingMessage{}
		//}
	// }

	glog.Info("subscribe2Updates - c: ", c)
	return nil
}

func (c *clientUpdatesHandler) getUpdatesConnID() (*ClientConnID, bool) {
	if c.connState == kTcpConn {
		if c.session != nil && c.pushSession != nil {
			if c.session.statusSyncTime + 60 > time.Now().Unix() {
				glog.Info("select updates")
				return &c.tcpConnID, false
			} else {
				glog.Info("select push")
				return &c.pushConnID, true
			}
		} else if c.session != nil {
			glog.Info("select updates")
			return &c.tcpConnID, false
		} else if c.pushSession != nil {
			glog.Info("select push")
			return &c.pushConnID, true
		} else {
		}
	} else if c.connState == kHttpConn {
		e := c.httpWaitConnIDs.Front()
		if e != nil {
			connID, _ := e.Value.(ClientConnID)
			c.httpWaitConnIDs.Remove(e)
			return &connID, false
		}
	}
	return nil, false
}

func (c *clientUpdatesHandler) UnSubscribeUpdates(connID ClientConnID) {
	// TODO(@benqi): clear

	if connID.connType == mtproto.TRANSPORT_TCP {
		if c.tcpConnID.Equal(connID) {
			c.session = nil
			// c.tcpConnID = connID
		} else if c.pushConnID.Equal(connID) {
			// frontendConnID == connID.frontendConnID && c.pushConnID.clientConnID == connID.clientConnID {
			c.pushSession = nil
			// c.pushConnID
		}
		if c.session == nil && c.pushSession == nil {
			c.connState = kConnUnknown
		}
	} else if connID.connType == mtproto.TRANSPORT_HTTP {
		for e := c.httpWaitConnIDs.Front(); e != nil; e = e.Next() {
			connID2, _ := e.Value.(ClientConnID)
			if connID2.Equal(connID) {
				c.httpWaitConnIDs.Remove(e)
			}
		}
		if c.httpWaitConnIDs.Len() == 0 {
			c.connState = kConnUnknown
		}
	}
}

func (c *clientUpdatesHandler) onSyncRpcResultData(cntl *zrpc.ZRpcController, data []byte) {
}

func (c *clientUpdatesHandler) onSyncData(isPush bool, cntl *zrpc.ZRpcController, obj mtproto.TLObject) {
	//switch obj.(type) {
	//case *mtproto.Updates:
	//default:
	//	glog.Error("invalid upadtes type, c: ", c, ", obj: ", obj)
	//	return
	//	// rpc result message
	//}

	//if isPush {
	//	if c.pushSession != nil {
	//		syncMessage := &pendingMessage{
	//			messageId: mtproto.GenerateMessageId(),
	//			confirm:   true,
	//			tl:        obj,
	//		}
	//		c.syncMessages = append(c.syncMessages, syncMessage)
	//
	//		glog.Infof("onSyncData - sendPending {c: {%v}, sess: {%v}, connID: {%v}}, pushObj: {%s}", c, c.session, c.pushConnID, reflect.TypeOf(obj))
	//		c.session.sendPendingMessagesToClient(c.pushConnID, cntl, c.syncMessages)
	//		c.syncMessages = []*pendingMessage{}
	//	}
	//} else {
	//
	//}
	//glog.Info("onSyncData - clientUpdatesHandler: ", c, obj)
	//connID, isPush := c.getUpdatesConnID()
	//
	////if c.session == nil || connID == nil {
	////	glog.Error("session not inited.")
	////	return
	////}

	// var pusObj mtproto.TLObject
	var (
		session = c.session
		connID *ClientConnID
	)

	if isPush {
		switch obj.(type) {
		case *mtproto.TLUpdatesTooLong:
			session = c.pushSession
			connID = &c.pushConnID
		//	pusObj = obj
		//case *mtproto.TLUpdateShortMessage:
		//	pusObj = mtproto.NewTLUpdatesTooLong().To_Updates()
		//case *mtproto.TLUpdateShortChatMessage:
		//	pusObj = mtproto.NewTLUpdatesTooLong().To_Updates()
		//case *mtproto.TLUpdates:
		//	upds := obj.(*mtproto.TLUpdates).GetUpdates()
		//	for _, upd := range upds {
		//		switch upd.GetConstructor() {
		//		case mtproto.TLConstructor_CRC32_updateNewMessage:
		//			pusObj = mtproto.NewTLUpdatesTooLong().To_Updates()
		//		}
		//	}
		//	if pusObj == nil {
		//		glog.Info("not push to client.")
		//		return
		//	}
		default:
			glog.Error("invalid push obj: ", obj)
			return
		//case *mtproto.Updates:
		//	switch obj.(*mtproto.Updates).GetConstructor() {
		//	case mtproto.TLConstructor_CRC32_updatesTooLong:
		//		pusObj = obj
		//	case mtproto.TLConstructor_CRC32_updateShortMessage,
		//		mtproto.TLConstructor_CRC32_updateShortChatMessage:
		//		pusObj = mtproto.NewTLUpdatesTooLong().To_Updates()
		//	case mtproto.TLConstructor_CRC32_updates:
		//		upds := obj.(*mtproto.Updates).GetData2().GetUpdates()
		//		for _, upd := range upds {
		//			switch upd.GetConstructor() {
		//			case mtproto.TLConstructor_CRC32_updateNewMessage:
		//				pusObj = mtproto.NewTLUpdatesTooLong().To_Updates()
		//			}
		//		}
		//		if pusObj == nil {
		//			glog.Info("not push to client.")
		//			return
		//		}
		//	default:
		//		glog.Info("not push to client.")
		//		return
		//	}
		//default:
		//	return
		}

		// session = c.pushSession
	} else {
		session = c.session
		connID = &c.tcpConnID

		// pusObj = obj
		//syncMessage := &pendingMessage{
		//	messageId: mtproto.GenerateMessageId(),
		//	confirm:   true,
		//	tl:        pusObj,
		//}
		//c.syncMessages = append(c.syncMessages, syncMessage)
		//
		//glog.Infof("onSyncData - sendPending {sess: {%v}, connID: {%v}}", c.session, connID)
		//c.session.sendPendingMessagesToClient(*connID, cntl, c.syncMessages)
		//c.syncMessages = []*pendingMessage{}
	}

	if session == nil || connID == nil {
		glog.Error("session not inited.")
		return
	}

	syncMessage := &pendingMessage{
		messageId: mtproto.GenerateMessageId(),
		confirm:   true,
		tl:        obj,
	}
	c.syncMessages = append(c.syncMessages, syncMessage)

	glog.Infof("onSyncData - sendPending {c: {%v}, sess: {%v}, connID: {%v}}, pushObj: {%s}", c, session, connID, reflect.TypeOf(obj))
	session.sendPendingMessagesToClient(*connID, cntl, c.syncMessages)
	c.syncMessages = []*pendingMessage{}

}
