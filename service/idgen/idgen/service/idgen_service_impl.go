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

package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/nebula-chat/chatengine/pkg/util"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"github.com/nebula-chat/chatengine/service/idgen/proto"
	"math/rand"
	"time"
)

const (
	PATH       = "/seqs/"
	UUID_KEY   = "/seqs/snowflake-uuid"
	BACKOFF    = 100  // max backoff delay millisecond
	CONCURRENT = 128  // max concurrent connections to etcd
	UUID_QUEUE = 1024 // uuid process queue
)

const (
	TS_MASK         = 0x1FFFFFFFFFF // 41bit
	SN_MASK         = 0xFFF         // 12bit
	MACHINE_ID_MASK = 0x3FF         // 10bit
)

type idgenServiceImpl struct {
	etcd      *Etcd
	machineID int64 // 10-bit machine id
	chProc    chan chan int64
}

func newIDGenServiceImpl(conf *etcdConf) (*idgenServiceImpl, error) {
	etcd, err := NewEtcd(conf)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	return &idgenServiceImpl{etcd: etcd}, nil
}

func (s *idgenServiceImpl) init() {
	s.chProc = make(chan chan int64, UUID_QUEUE)
	s.initMachineID()
	go s.uuidTask()
}

func (s *idgenServiceImpl) initMachineID() {
	// var prevIndex int64
	var prevValue int32
	for {
		// get the key
		resp, err := s.etcd.EtcCli.Get(context.Background(), UUID_KEY)
		if err != nil {
			glog.Error(err)
			panic(err)
		}
		for _, value := range resp.Kvs {
			prevValue, err = util.StringToInt32(string(value.Value))
			if err != nil {
				glog.Error(err)
				return
			}
			// prevIndex = value.ModRevision
			// glog.Info(prevValue)
			// glog.Info(prevIndex)
		}
		// glog.Info(prevValue)
		// glog.Info(prevIndex)
		// _, err = s.dao.Etcd.EtcCli.Put(context.Background(), UUID_KEY, fmt.Sprint(prevValue+1), clientv3.WithRev(0))
		_, err = s.etcd.EtcCli.Put(context.Background(), UUID_KEY, fmt.Sprint(prevValue+1))
		if err != nil {
			cas_delay()
			continue
		}
		// record serial number of this service, already shifted
		s.machineID = (int64(prevValue+1) & MACHINE_ID_MASK) << 12
		return
	}
}

// uuid generator
func (s *idgenServiceImpl) uuidTask() {
	var sn int64      // 12-bit serial no
	var last_ts int64 // last timestamp
	for {
		ret := <-s.chProc
		// get a correct serial number
		t := ts()
		if t < last_ts { // clock shift backward
			glog.Error("clock shift happened, waiting until the clock moving to the next millisecond.")
			t = s.wait_ms(last_ts)
		}
		if last_ts == t { // same millisecond
			sn = (sn + 1) & SN_MASK
			if sn == 0 { // serial number overflows, wait until next ms
				t = s.wait_ms(last_ts)
			}
		} else { // new millsecond, reset serial number to 0
			sn = 0
		}
		// remember last timestamp
		last_ts = t
		// generate uuid, format:
		//
		// 0		0.................0		0..............0	0........0
		// 1-bit	41bit timestamp			10bit machine-id	12bit sn
		var uuid int64
		uuid |= (int64(t) & TS_MASK) << 22
		uuid |= s.machineID
		uuid |= sn
		ret <- uuid
	}
}

// wait_ms will spin wait till next millisecond.
func (s *idgenServiceImpl) wait_ms(last_ts int64) int64 {
	t := ts()
	for t <= last_ts {
		t = ts()
	}
	return t
}

// random delay
func cas_delay() {
	<-time.After(time.Duration(rand.Int63n(BACKOFF)) * time.Millisecond)
}

// get timestamp
func ts() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// rpc GetUUID(String) returns (Int64);
func (s *idgenServiceImpl) GetUUID(ctx context.Context, request *seqsvr.Void) (reply *seqsvr.Int64, err error) {
	glog.Infof("idgen.GetUUID - request: %s", logger.JsonDebugData(request))

	req := make(chan int64, 1)

	s.chProc <- req
	reply = &seqsvr.Int64{<-req}

	glog.Infof("idgen.GetUUID - reply: {%v}", reply)
	return
}

// rpc GetCurrentSeqID(String) returns (Int64);
func (s *idgenServiceImpl) GetCurrentSeqID(ctx context.Context, request *seqsvr.String) (reply *seqsvr.Int64, err error) {
	glog.Infof("idgen.GetCurrentSeqID - request: %s", logger.JsonDebugData(request))

	var prevIndex int64
	var prevValue int64
	var resp *clientv3.GetResponse
	key := PATH + request.V

	glog.Info(s.etcd.EtcCli)
	if s.etcd.EtcCli == nil {
		glog.Error("s.etcd.EtcCli == nil")
		return
	}
	// get the key
	resp, err = s.etcd.EtcCli.Get(context.Background(), key)
	if err != nil {
		glog.Error(err)
		return nil, errors.New("key not exists, need to create first")
	}
	// get prevValue & prevIndex
	for _, value := range resp.Kvs {
		prevValue, err = util.StringToInt64(string(value.Value))
		if err != nil {
			glog.Error(err)
			return nil, errors.New("marlformed value")
		}
		prevIndex = value.ModRevision
	}

	glog.Info(prevValue)
	glog.Info(prevIndex)

	reply = &seqsvr.Int64{V: prevValue + 1}

	glog.Infof("idgen.GetCurrentSeqID - reply: {%v}", reply)
	return
}

// rpc GetNextSeqID(String) returns (Int64);
func (s *idgenServiceImpl) GetNextSeqID(ctx context.Context, request *seqsvr.String) (reply *seqsvr.Int64, err error) {
	glog.Infof("idgen.GetNextSeqID - request: %s", logger.JsonDebugData(request))

	var prevIndex int64
	var prevValue int64
	var resp *clientv3.GetResponse
	key := PATH + request.V
	for {
		glog.Info(s.etcd.EtcCli)
		if s.etcd.EtcCli == nil {
			glog.Error("s.etcd.EtcCli == nil")
			return
		}
		// get the key
		resp, err = s.etcd.EtcCli.Get(context.Background(), key)
		if err != nil {
			glog.Error(err)
			return nil, errors.New("key not exists, need to create first")
		}
		// get prevValue & prevIndex
		for _, value := range resp.Kvs {
			prevValue, err = util.StringToInt64(string(value.Value))
			if err != nil {
				glog.Error(err)
				return nil, errors.New("marlformed value")
			}
			prevIndex = value.ModRevision
		}
		glog.Info(prevValue)
		glog.Info(prevIndex)
		_, err = s.etcd.EtcCli.Put(context.Background(), key, fmt.Sprint(prevValue+1))
		if err != nil {
			cas_delay()
			continue
		}

		reply = &seqsvr.Int64{V: prevValue + 1}
		break
	}

	glog.Infof("idgen.GetNextSeqID - reply: {%v}", reply)
	return
}
