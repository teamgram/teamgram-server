// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package dao

import (
	"context"
	"fmt"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

func SerializeToBuffer2(salt, sessionId int64, msg2 *mtproto.TLMessageRawData) []byte {
	x := mtproto.NewEncodeBuf(32 + len(msg2.Body))

	x.Long(salt)
	x.Long(sessionId)
	x.Long(msg2.MsgId)
	x.Int(msg2.Seqno)
	x.Int(msg2.Bytes)
	x.Bytes(msg2.Body)

	return x.GetBuf()
}

func (d *Dao) watchGateway(c zrpc.RpcClientConf) {
	sub, err := discov.NewSubscriber(c.Etcd.Hosts, c.Etcd.Key)
	if err != nil {
		logx.Errorf("watchGateway NewSubscriber(%+v) error: %v", c.Etcd, err)
		return
	}

	update := func() {
		values := sub.Values()
		if len(values) == 0 {
			return
		}

		d.gateMu.Lock()
		defer d.gateMu.Unlock()

		clients := map[string]*Gateway{}
		for _, v := range values {
			if old, ok := d.eGateServers[v]; ok {
				clients[v] = old
				continue
			}
			c.Endpoints = []string{v}
			cli, err := NewGateway(c)
			if err != nil {
				logx.Errorf("watchGateway NewGateway(%+v) error: %v", v, err)
				continue
			}
			clients[v] = cli
		}

		for key, old := range d.eGateServers {
			if _, ok := clients[key]; !ok {
				old.cancel()
				logx.Infof("watchGateway DelGateway: %s", key)
			}
		}

		d.eGateServers = clients
	}

	sub.AddListener(update)
	update()
}

func (d *Dao) SendDataToGateway(ctx context.Context, gatewayId string, authKeyId, salt, sessionId int64, msg *mtproto.TLMessageRawData) (bool, error) {
	d.gateMu.RLock()
	c, ok := d.eGateServers[gatewayId]
	d.gateMu.RUnlock()

	if ok {
		return c.SendDataToGate(ctx, authKeyId, sessionId, SerializeToBuffer2(salt, sessionId, msg))
	} else {
		logx.WithContext(ctx).Errorf("not found k: %s", gatewayId)
		return false, fmt.Errorf("not found k: %s", gatewayId)
	}
}

func (d *Dao) SendHttpDataToGateway(ctx context.Context, ch chan interface{}, authKeyId, salt, sessionId int64, msg *mtproto.TLMessageRawData) (bool, error) {
	select {
	case ch <- SerializeToBuffer2(salt, sessionId, msg):
		close(ch)
		return true, nil
	default:
		logx.WithContext(ctx).Errorf("Default fail !!!!! ch closed")
		return false, fmt.Errorf("ch closed")
	}
}
