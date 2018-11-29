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

package dataobject

type PhoneCallSessionsDO struct {
	Id                    int32  `db:"id"`
	CallSessionId         int64  `db:"call_session_id"`
	AdminId               int32  `db:"admin_id"`
	AdminAccessHash       int64  `db:"admin_access_hash"`
	ParticipantId         int32  `db:"participant_id"`
	ParticipantAccessHash int64  `db:"participant_access_hash"`
	UdpP2p                int8   `db:"udp_p2p"`
	UdpReflector          int8   `db:"udp_reflector"`
	MinLayer              int32  `db:"min_layer"`
	MaxLayer              int32  `db:"max_layer"`
	GA                    string `db:"g_a"`
	GB                    string `db:"g_b"`
	State                 int32  `db:"state"`
	AdminDebugData        string `db:"admin_debug_data"`
	ParticipantDebugData  string `db:"participant_debug_data"`
	Date                  int32  `db:"date"`
	CreatedAt             string `db:"created_at"`
	UpdatedAt             string `db:"updated_at"`
}
