// Copyright 2022 Teamgram Authors
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
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package service

import (
	"context"
	"fmt"
	"github.com/teamgram/marmota/pkg/net/ip"
	"sync"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/interface/session/internal/config"
	"github.com/teamgram/teamgram-server/app/interface/session/internal/dao"

	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

var (
	//etcdPrefix is a etcd globe key prefix
	endpoints string
)

//
//func init() {
//	endpoints = os.Getenv("ETCD_ENDPOINTS")
//	if endpoints == "" {
//		panic(fmt.Errorf("invalid etcd config endpoints:%+v", endpoints))
//	}
//}

type Service struct {
	ac              config.Config
	mu              sync.RWMutex
	sessionsManager map[int64]*authSessions
	eGateServers    map[string]*Gateway
	reqCache        *RequestManager
	serverId        string
	*dao.Dao
}

func New(c config.Config) *Service {
	var (
		// ac  = &Config{}
		// err error
		s = new(Service)
	)

	//if err = paladin.Get("application.toml").UnmarshalTOML(&ac); err != nil {
	//	if err != paladin.ErrNotExist {
	//		panic(err)
	//	}
	//}

	s.ac = c
	s.Dao = dao.New(c)
	s.sessionsManager = make(map[int64]*authSessions)
	s.eGateServers = make(map[string]*Gateway)
	s.reqCache = NewRequestManager()
	s.serverId = ip.FigureOutListenOn(c.ListenOn)

	s.watchGateway(s.ac.GatewayClient)
	return s
}

func (s *Service) Close() error {
	for _, c := range s.eGateServers {
		if err := c.Close(); err != nil {
			logx.Errorf("c.Close() error(%v)", err)
		}
	}

	// s.Dao.Close()
	return nil
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	return nil
}

//func (s *Service) newAddress(insMap map[string][]*naming.Instance) error {
//	ins := insMap[env.Zone]
//	if len(ins) == 0 {
//		return fmt.Errorf("watchComet instance is empty")
//	}
//	eGates := map[string]*Gateway{}
//	options := gatewayOptions{
//		RoutineSize: s.ac.Routine.Size,
//		RoutineChan: s.ac.Routine.Chan,
//	}
//
//	for _, in := range ins {
//		if old, ok := s.eGateServers[in.Hostname]; ok {
//			eGates[in.Hostname] = old
//			continue
//		}
//		c, err := NewGateway(in, s.ac, options)
//		if err != nil {
//			log.Errorf("watchComet NewComet(%+v) error(%v)", in, err)
//			return err
//		}
//		eGates[in.Hostname] = c
//		log.Infof("watchComet AddComet grpc:%+v", in)
//	}
//	for key, old := range s.eGateServers {
//		if _, ok := eGates[key]; !ok {
//			_ = old
//			// old.cancel()
//			log.Infof("watchComet DelComet:%s", key)
//		}
//	}
//	s.eGateServers = eGates
//	return nil
//}

func (s *Service) watchGateway(c zrpc.RpcClientConf) {
	sub, _ := discov.NewSubscriber(c.Etcd.Hosts, c.Etcd.Key)
	update := func() {
		values := sub.Values()
		if len(values) == 0 {
			return
		}

		clients := map[string]*Gateway{}
		for _, v := range values {
			if old, ok := s.eGateServers[v]; ok {
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

		for key, old := range s.eGateServers {
			if _, ok := clients[key]; !ok {
				old.cancel()
				logx.Infof("watchComet DelComet:%s", key)
			}
		}

		s.eGateServers = clients
	}

	sub.AddListener(update)
	update()
}

func (s *Service) SendDataToGateway(ctx context.Context, gatewayId string, authKeyId, salt, sessionId int64, msg *mtproto.TLMessageRawData) (bool, error) {
	// log.Debugf("sendDataToGateway - %v", s.eGateServers)
	//k := fmt.Sprintf("/nebulaim/egate/node%d", serverId)
	if c, ok := s.eGateServers[gatewayId]; ok {
		return c.SendDataToGate(ctx, authKeyId, sessionId, SerializeToBuffer2(salt, sessionId, msg))
	} else {
		logx.WithContext(ctx).Errorf("not found k: %s, %v", gatewayId, s.eGateServers)
		return false, fmt.Errorf("not found k: %s", gatewayId)
	}
}

func (s *Service) SendHttpDataToGateway(ctx context.Context, ch chan interface{}, authKeyId, salt, sessionId int64, msg *mtproto.TLMessageRawData) (bool, error) {
	select {
	case ch <- SerializeToBuffer2(salt, sessionId, msg):
		close(ch)
		return true, nil
	default:
		logx.WithContext(ctx).Errorf("Default fail !!!!! ch closed")
		return false, fmt.Errorf("ch closed")
	}
}

//func (s *Service) PushUpdatesToNpns(ctx context.Context, authKeyId int64, updates *mtproto.Updates) {
//	s.npnsClient.PushUpdates(ctx, &npnspb.PushUpdates{
//		AuthKeyId: authKeyId,
//		Updates:   updates,
//	})
//}

func (s *Service) DeleteByAuthKeyId(authKeyId int64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if sessList, ok := s.sessionsManager[authKeyId]; ok {
		sessList.Stop()
		delete(s.sessionsManager, authKeyId)
	}
}
