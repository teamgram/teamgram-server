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

package update

import (
	"github.com/nebula-chat/chatengine/pkg/util"
)

const (
	seqUpdatesNgenId        = "seq_updates_ngen_"
	ptsUpdatesNgenId        = "pts_updates_ngen_"
	qtsUpdatesNgenId        = "qts_updates_ngen_"
	channelPtsUpdatesNgenId = "channel_pts_updates_ngen_"
)

func (m *UpdateModel) NextSeqId(key int32) (seq int64) {
	seq, _ = m.dao.SeqIDGen.GetNextSeqID(seqUpdatesNgenId + util.Int32ToString(key))
	return
}

func (m *UpdateModel) CurrentSeqId(key int32) (seq int64) {
	var err error
	seq, _ = m.dao.SeqIDGen.GetCurrentSeqID(seqUpdatesNgenId + util.Int32ToString(key))

	if err != nil {
		seq = -1
	}
	return
}

func (m *UpdateModel) NextPtsId(key int32) (seq int64) {
	seq, _ = m.dao.SeqIDGen.GetNextSeqID(ptsUpdatesNgenId + util.Int32ToString(key))
	return
}

func (m *UpdateModel) CurrentPtsId(key int32) (seq int64) {
	seq, _ = m.dao.SeqIDGen.GetCurrentSeqID(ptsUpdatesNgenId + util.Int32ToString(key))
	return
}

func (m *UpdateModel) NextQtsId(key int32) (seq int64) {
	seq, _ = m.dao.SeqIDGen.GetNextSeqID(qtsUpdatesNgenId + util.Int32ToString(key))
	return
}

func (m *UpdateModel) CurrentQtsId(key int32) (seq int64) {
	seq, _ = m.dao.SeqIDGen.GetNextSeqID(qtsUpdatesNgenId + util.Int32ToString(key))
	return
}

func (m *UpdateModel) NextChannelPtsId(key int32) (seq int64) {
	seq, _ = m.dao.SeqIDGen.GetNextSeqID(channelPtsUpdatesNgenId + util.Int32ToString(key))
	return
}

func (m *UpdateModel) CurrentChannelPtsId(key int32) (seq int64) {
	seq, _ = m.dao.SeqIDGen.GetNextSeqID(channelPtsUpdatesNgenId + util.Int32ToString(key))
	return
}
