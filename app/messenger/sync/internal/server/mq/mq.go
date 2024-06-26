// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package mq

import (
	"context"
	"encoding/json"
	"fmt"

	kafka "github.com/teamgram/marmota/pkg/mq"
	"github.com/teamgram/teamgram-server/app/messenger/sync/internal/core"
	"github.com/teamgram/teamgram-server/app/messenger/sync/internal/svc"
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// New new a grpc server.
func New(svcCtx *svc.ServiceContext, conf kafka.KafkaConsumerConf) *kafka.ConsumerGroup {
	s := kafka.MustKafkaConsumer(&conf)
	s.RegisterHandlers(
		conf.Topics[0],
		func(ctx context.Context, method, key string, value []byte) {
			logx.WithContext(ctx).Debugf("method: %s, key: %s, value: %s", key, value)

			switch protoreflect.FullName(method) {
			case proto.MessageName((*sync.TLSyncUpdatesMe)(nil)):
				threading.RunSafe(func() {
					c := core.New(ctx, svcCtx)

					r := new(sync.TLSyncUpdatesMe)
					if err := json.Unmarshal(value, r); err != nil {
						c.Logger.Error(err.Error())
						return
					}
					c.Logger.Debugf("sync.updatesMe - request: %s", r)

					c.SyncUpdatesMe(r)
				})
			case proto.MessageName((*sync.TLSyncUpdatesNotMe)(nil)):
				threading.RunSafe(func() {
					c := core.New(ctx, svcCtx)

					r := new(sync.TLSyncUpdatesNotMe)
					if err := json.Unmarshal(value, r); err != nil {
						c.Logger.Error(err.Error())
						return
					}
					c.Logger.Debugf("sync.updatesNotMe - request: %s", r)

					c.SyncUpdatesNotMe(r)
				})
			case proto.MessageName((*sync.TLSyncPushUpdates)(nil)):
				threading.RunSafe(func() {
					c := core.New(ctx, svcCtx)

					r := new(sync.TLSyncPushUpdates)
					if err := json.Unmarshal(value, r); err != nil {
						c.Logger.Error(err.Error())
						return
					}
					c.Logger.Debugf("sync.pushUpdates - request: %s", r)

					c.SyncPushUpdates(r)
				})
			case proto.MessageName((*sync.TLSyncPushRpcResult)(nil)):
				threading.RunSafe(func() {
					c := core.New(ctx, svcCtx)

					r := new(sync.TLSyncPushRpcResult)
					if err := json.Unmarshal(value, r); err != nil {
						c.Logger.Error(err.Error())
						return
					}
					c.Logger.Debugf("sync.pushRpcResult - request: %s", r)

					c.SyncPushRpcResult(r)
				})
			case proto.MessageName((*sync.TLSyncBroadcastUpdates)(nil)):
				threading.RunSafe(func() {
					c := core.New(ctx, svcCtx)

					r := new(sync.TLSyncBroadcastUpdates)
					if err := json.Unmarshal(value, r); err != nil {
						c.Logger.Error(err.Error())
						return
					}
					c.Logger.Debugf("sync.broadcastUpdates - request: %s", r)

					c.SyncBroadcastUpdates(r)
				})
			default:
				err := fmt.Errorf("invalid key: %s", key)
				logx.Error(err.Error())
			}
		})
	return s
}
