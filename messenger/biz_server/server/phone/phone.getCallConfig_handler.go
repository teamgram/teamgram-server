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

package phone

import (
    "github.com/golang/glog"
    "golang.org/x/net/context"
    "github.com/nebula-chat/chatengine/pkg/grpc_util"
    "github.com/nebula-chat/chatengine/pkg/logger"
    "github.com/nebula-chat/chatengine/mtproto"
    "encoding/json"
)

/*
 "dataJSON":
 {
    "audio_frame_size": 60,
    "jitter_min_delay_20": 6,
    "jitter_min_delay_40": 4,
    "jitter_min_delay_60": 2,
    "jitter_max_delay_20": 25,
    "jitter_max_delay_40": 15,
    "jitter_max_delay_60": 10,
    "jitter_max_slots_20": 50,
    "jitter_max_slots_40": 30,
    "jitter_max_slots_60": 20,
    "jitter_losses_to_reset": 20,
    "jitter_resync_threshold": 0.5,
    "audio_congestion_window": 1024,
    "audio_max_bitrate": 20000,
    "audio_max_bitrate_edge": 16000,
    "audio_max_bitrate_gprs": 8000,
    "audio_max_bitrate_saving": 8000,
    "audio_init_bitrate": 16000,
    "audio_init_bitrate_edge": 8000,
    "audio_init_bitrate_gprs": 8000,
    "audio_init_bitrate_saving": 8000,
    "audio_bitrate_step_incr": 1000,
    "audio_bitrate_step_decr": 1000,
    "use_system_ns": true,
    "use_system_aec": true
 }
*/

type callConfigDataJSON struct {
    AudioFrameSize         int     `json:"audio_frame_size"`
    JitterMinDelay20       int     `json:"jitter_min_delay_20"`
    JitterMinDelay40       int     `json:"jitter_min_delay_40"`
    JitterMinDelay60       int     `json:"jitter_min_delay_60"`
    JitterMaxDelay20       int     `json:"jitter_max_delay_20"`
    JitterMaxDelay40       int     `json:"jitter_max_delay_40"`
    JitterMaxDelay60       int     `json:"jitter_max_delay_60"`
    JitterMaxSlots20       int     `json:"jitter_max_slots_20"`
    JitterMaxSlots40       int     `json:"jitter_max_slots_40"`
    JitterMaxSlots60       int     `json:"jitter_max_slots_60"`
    JitterLossesToReset    int     `json:"jitter_losses_to_reset"`
    JitterResyncThreshold  float32 `json:"jitter_resync_threshold"`
    AudioCongestionWindow  int     `json:"audio_congestion_window"`
    AudioMaxBitrate        int     `json:"audio_max_bitrate"`
    AudioMaxBitrateEdge    int     `json:"audio_max_bitrate_edge"`
    AudioMaxBitrateGprs    int     `json:"audio_max_bitrate_gprs"`
    AudioMaxBitrateSaving  int     `json:"audio_max_bitrate_saving"`
    AudioInitBitrate       int     `json:"audio_init_bitrate"`
    AudioInitBitrateEdge   int     `json:"audio_init_bitrate_edge"`
    AudioInitBitrateGrps   int     `json:"audio_init_bitrate_gprs"`
    AudioInitBitrateSaving int     `json:"audio_init_bitrate_saving"`
    AudioBitrateStepIncr   int     `json:"audio_bitrate_step_incr"`
    AudioBitrateStepDecr   int     `json:"audio_bitrate_step_decr"`
    UseSystemNs            bool    `json:"use_system_ns"`
    UseSystemAec           bool    `json:"us audioInitBitrateGrps inte_system_aec"`
}

// TODO(@benqi): 写死配置
func NewCallConfigDataJSON() *callConfigDataJSON {
    return &callConfigDataJSON{
        AudioFrameSize:         60,
        JitterMinDelay20:       6,
        JitterMinDelay40:       4,
        JitterMinDelay60:       2,
        JitterMaxDelay20:       25,
        JitterMaxDelay40:       15,
        JitterMaxDelay60:       10,
        JitterMaxSlots20:       50,
        JitterMaxSlots40:       30,
        JitterMaxSlots60:       20,
        JitterLossesToReset:    20,
        JitterResyncThreshold:  0.5,
        AudioCongestionWindow:  1024,
        AudioMaxBitrate:        20000,
        AudioMaxBitrateEdge:    16000,
        AudioMaxBitrateGprs:    8000,
        AudioMaxBitrateSaving:  8000,
        AudioInitBitrate:       16000,
        AudioInitBitrateEdge:   8000,
        AudioInitBitrateGrps:   8000,
        AudioInitBitrateSaving: 8000,
        AudioBitrateStepIncr:   1000,
        AudioBitrateStepDecr:   1000,
        UseSystemNs:            true,
        UseSystemAec:           true,
    }
}

// phone.getCallConfig#55451fa9 = DataJSON;
func (s *PhoneServiceImpl) PhoneGetCallConfig(ctx context.Context, request *mtproto.TLPhoneGetCallConfig) (*mtproto.DataJSON, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("phone.getCallConfig#55451fa9 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

    callConfig := NewCallConfigDataJSON()
    callConfigData, err := json.Marshal(callConfig)
    if err != nil {
        glog.Error(err)
        return nil, err
    }
    j := string(callConfigData)
    glog.Info("callConfigData: ", callConfig, ", ", j)

    reply := mtproto.NewTLDataJSON()
    reply.SetData(j)

    glog.Infof("phone.getCallConfig#55451fa9 - reply %s", logger.JsonDebugData(reply))
    return reply.To_DataJSON(), nil
}
