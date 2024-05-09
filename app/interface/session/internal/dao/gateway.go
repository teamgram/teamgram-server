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
	sub, _ := discov.NewSubscriber(c.Etcd.Hosts, c.Etcd.Key)
	update := func() {
		values := sub.Values()
		if len(values) == 0 {
			return
		}

		clients := map[string]*Gateway{}
		for _, v := range values {
			if old, ok := d.eGateServers[v]; ok {
				clients[v] = old
				continue
			}
			c.Endpoints = []string{v}
			// cli, err := zrpc.NewClient(c)
			cli, err := NewGateway(c)
			if err != nil {
				logx.Error("watchComet NewClient(%+v) error(%v)", values, err)
				return
			}
			clients[v] = cli
		}

		for key, old := range d.eGateServers {
			if _, ok := clients[key]; !ok {
				old.cancel()
				logx.Infof("watchComet DelComet:%s", key)
			}
		}

		d.eGateServers = clients
	}

	sub.AddListener(update)
	update()
}

func (d *Dao) SendDataToGateway(ctx context.Context, gatewayId string, authKeyId, salt, sessionId int64, msg *mtproto.TLMessageRawData) (bool, error) {
	if c, ok := d.eGateServers[gatewayId]; ok {
		return c.SendDataToGate(ctx, authKeyId, sessionId, SerializeToBuffer2(salt, sessionId, msg))
	} else {
		logx.WithContext(ctx).Errorf("not found k: %s, %v", gatewayId, d.eGateServers)
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
