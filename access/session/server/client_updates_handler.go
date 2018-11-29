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
)

const (
	kConnUnknown = 0
	kTcpConn     = 1
	kHttpConn    = 2
)

type clientUpdatesHandler struct {
	session         *clientSessionHandler
	syncMessages    []*pendingMessage
	connState       int
	connType        int
	tcpConnID       ClientConnID
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
	return fmt.Sprintf("{sess: %s, conn_state: %d, conn_type: %d, tcp_conn_id: %s}",
		c.session, c.connState, c.connType, c.tcpConnID)
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

func (c *clientUpdatesHandler) getUpdatesConnID() *ClientConnID {
	if c.connState == kTcpConn {
		return &c.tcpConnID
	} else if c.connState == kHttpConn {
		e := c.httpWaitConnIDs.Front()
		if e != nil {
			connID, _ := e.Value.(ClientConnID)
			c.httpWaitConnIDs.Remove(e)
			return &connID
			// return &e.Value.(ClientConnID)
		}
	}
	return nil
}

func (c *clientUpdatesHandler) UnSubscribeUpdates(connID ClientConnID) {
	// TODO(@benqi): clear

	if connID.connType == mtproto.TRANSPORT_TCP {
		c.tcpConnID = connID
		c.connState = kConnUnknown
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

func (c *clientUpdatesHandler) onSyncData(cntl *zrpc.ZRpcController, obj mtproto.TLObject) {
	//switch obj.(type) {
	//case *mtproto.Updates:
	syncMessage := &pendingMessage{
		messageId: mtproto.GenerateMessageId(),
		confirm:   true,
		tl:        obj,
	}
	c.syncMessages = append(c.syncMessages, syncMessage)
	//default:
	//	glog.Error("invalid upadtes type, c: ", c, ", obj: ", obj)
	//	return
	//	// rpc result message
	//}

	connID := c.getUpdatesConnID()

	if c.session == nil || connID == nil {
		glog.Error("session not inited.")
		return
	}

	glog.Infof("onSyncData - sendPending {sess: {%v}, connID: {%v}}", c.session, connID)

	c.session.sendPendingMessagesToClient(*connID, cntl, c.syncMessages)
	c.syncMessages = []*pendingMessage{}
}
